"""Pydantic Renderer - Converts SchemaNode IR to Pydantic code strings.

Design Philosophy
-----------------
This renderer generates Python code that validates data at runtime using Pydantic.
Key design decisions:

1. STRICT TYPES: We use StrictInt, StrictFloat, StrictBool to prevent Python's
   loose type coercion. Python treats 1 == True and 0 == False, but JSON Schema
   does not - integers and booleans are distinct types. Without strict types,
   Pydantic would accept True for {"type": "integer"} schemas.

2. HELPER FUNCTION REGISTRY: Complex validations (const, enum with deep equality)
   need helper functions. Rather than emit helpers inline everywhere, we register
   which helpers are needed during rendering and emit them once at module top.
   This keeps generated code clean and avoids duplication.

3. VALIDATOR STRATEGIES: We use different Pydantic features based on the schema:
   - Annotated[T, Field(...)] for simple constraints (minLength, maximum)
   - AfterValidator for checks that need the validated value (uniqueItems)
   - BeforeValidator for checks that might transform/reject early (oneOf)
   - model_validator for object-level checks (dependencies, minProperties)

4. TYPE EXPRESSIONS vs CODE: Each render function returns both a type_expr
   (e.g., "list[int]") and code (class definitions, validators). Type expressions
   are used inline; code is emitted at module level.

5. JSON SCHEMA SEMANTICS: Many edge cases stem from JSON Schema's quirks:
   - Format is annotation-only by default (no validation)
   - additionalProperties applies to properties NOT in 'properties' keyword
   - unevaluatedProperties requires tracking what was "evaluated"
   - oneOf requires EXACTLY one match (not just first match like Union)

Generated Code Structure
------------------------
For a complex schema, the generated module looks like:

    # Helper functions (if needed)
    def _json_equals(a, b): ...

    # Nested classes (bottom-up order - dependencies first)
    class AddressStreet(BaseModel): ...
    class Address(BaseModel): ...

    # Main schema
    class Person(BaseModel): ...
    Person_Validator = TypeAdapter(Person)
"""

from collections.abc import Sequence
from dataclasses import dataclass, field
from typing import Any

from xschema_pydantic.errors import ConversionError

from xschema_core import (
    AnyNode,
    ArrayNode,
    BooleanNode,
    ConditionalNode,
    EnumNode,
    IntersectionNode,
    LiteralNode,
    NeverNode,
    NotNode,
    NullableNode,
    NullNode,
    NumberNode,
    ObjectNode,
    OneOfNode,
    PropertyDef,
    PropertyDependency,
    RefNode,
    SchemaDependency,
    SchemaNode,
    StringNode,
    TupleNode,
    TypeGuardedNode,
    UnionNode,
)


def _escape_for_single_quotes(name: str) -> str:
    """Escape a string for use inside single-quoted Python string literals.

    Escapes backslashes, single quotes, and control characters.
    """
    import json

    # json.dumps gives us proper escaping for control chars, but uses double quotes
    escaped = json.dumps(name)
    # Remove surrounding double quotes
    content = escaped[1:-1]
    # json.dumps escapes " as \" - we need to unescape that for single-quoted strings
    # But we do need to escape single quotes
    content = content.replace('\\"', '"')  # unescape double quotes
    content = content.replace("'", "\\'")  # escape single quotes
    return content


def _escape_for_double_quotes(name: str) -> str:
    """Escape a string for use inside double-quoted Python string literals.

    Escapes backslashes, double quotes, and control characters.
    """
    import json

    # json.dumps gives us proper escaping already for double-quoted strings
    escaped = json.dumps(name)
    # Remove surrounding double quotes, but keep all internal escaping
    return escaped[1:-1]


def _make_nullable(type_expr: str) -> str:
    """Make a type expression nullable by adding '| None'.

    Handles forward reference strings correctly. When type_expr is a quoted
    string like 'T', we need to keep the union inside the quotes: 'T | None'.
    Otherwise 'T' | None would be invalid Python (string | None).

    Examples:
        'T' -> 'T | None'
        StrictInt -> StrictInt | None
        'T | None' -> 'T | None' (already nullable)
        StrictInt | None -> StrictInt | None (already nullable)
    """
    # Skip if already nullable
    if type_expr.endswith(" | None") or type_expr == "None":
        return type_expr

    # Check if it's a forward reference (starts and ends with single quotes)
    if type_expr.startswith("'") and type_expr.endswith("'"):
        # Extract the inner type and make it nullable inside the quotes
        inner = type_expr[1:-1]
        return f"'{inner} | None'"

    # Regular type - just append | None
    return f"{type_expr} | None"


# Python builtins that shadow types used in type annotations
# When a field is named "list", it shadows Python's builtin list type,
# causing "list[int]" in type annotations to fail with "'NoneType' is not subscriptable"
# Only includes types that are actually used in generic type annotations (list[T], dict[K,V], etc.)
_PYTHON_TYPE_BUILTINS = frozenset(
    {
        "list",
        "dict",
        "set",
        "tuple",
        "type",
        "frozenset",
    }
)


def _is_valid_python_identifier(name: str) -> bool:
    """Check if a string is a valid Python identifier for Pydantic fields.

    A valid Python identifier for Pydantic:
    - Starts with a letter (a-z, A-Z) or underscore
    - Contains only letters, digits, or underscores
    - Is not a Python keyword
    - Does NOT start with double underscore (Pydantic treats dunder names specially)
    - Does NOT shadow Python builtins used in type annotations
    """
    import keyword

    if not name:
        return False
    if not name.isidentifier():
        return False
    if keyword.iskeyword(name):
        return False
    # Pydantic treats dunder names specially, so they need aliases
    if name.startswith("__"):
        return False
    # Builtins used in type annotations would be shadowed by field names
    if name in _PYTHON_TYPE_BUILTINS:
        return False
    return True


def _sanitize_property_name(name: str, used_names: set[str] | None = None) -> str:
    """Sanitize a property name to be a valid Python identifier.

    Replaces invalid characters with underscores and ensures the name
    doesn't start with a digit. If used_names is provided, ensures uniqueness
    by appending a counter suffix.

    Pydantic doesn't allow field names starting with double underscores,
    so we prefix those with 'f_' (field).

    Args:
        name: The original property name
        used_names: Set of already used identifiers to avoid collisions
    """
    import re
    import keyword

    if not name:
        base = "field_"
    else:
        # Replace any non-alphanumeric/underscore characters with underscore
        base = re.sub(r"[^a-zA-Z0-9_]", "_", name)

        # If it starts with a digit, prefix with underscore
        if base and base[0].isdigit():
            base = "_" + base

        # If empty after sanitization, use a placeholder
        if not base:
            base = "field_"

        # If it's a Python keyword, append underscore
        if keyword.iskeyword(base):
            base = base + "_"

        # If it's a Python builtin that shadows types, append underscore
        if base in _PYTHON_TYPE_BUILTINS:
            base = base + "_"

        # Pydantic doesn't allow field names starting with underscore
        # (both single _ and double __)
        # Prefix with 'f' to make it valid
        if base.startswith("_"):
            base = "f" + base

    # Ensure uniqueness if used_names provided
    if used_names is None:
        return base

    sanitized = base
    counter = 2
    while sanitized in used_names:
        sanitized = f"{base}_{counter}"
        counter += 1

    used_names.add(sanitized)
    return sanitized


# HELPER FUNCTION REGISTRY
# ========================
# Instead of emitting helper code every time it's needed, we register helpers
# during rendering and emit them once at the end. This avoids code duplication
# and keeps the generated output clean.
#
# Pattern: Call _register_helper("_json_equals") when generating code that needs it.
# Then get_needed_helpers() returns all helper code to emit at module top.

_helper_registry: set[str] = set()
_root_class_name: str = ""
_root_is_class: bool = False
_needs_model_rebuild: bool = False


def _reset_helper_registry() -> None:
    """Reset the helper registry. Called before rendering each schema."""
    global _helper_registry
    _helper_registry = set()


def set_root_class_name(name: str, is_class: bool) -> None:
    """Set the root class name for forward reference resolution.

    Called by the converter before rendering so that recursive $ref: "#"
    can be rendered as a forward reference to the root class instead of Any.
    Forward refs only work when the root schema generates a BaseModel class.
    """
    global _root_class_name, _root_is_class, _needs_model_rebuild
    _root_class_name = name
    _root_is_class = is_class
    _needs_model_rebuild = False


def needs_model_rebuild() -> bool:
    """Check if the rendered schema uses forward references requiring model_rebuild()."""
    return _needs_model_rebuild


def _register_helper(helper_name: str) -> None:
    """Register a helper function as needed."""
    _helper_registry.add(helper_name)


# HELPER FUNCTION DEFINITIONS
# ===========================
# These are emitted once at module top if registered during rendering.
#
# _json_equals: Critical for const/enum validation. Python's == operator has
# surprising behavior: False == 0 is True, True == 1 is True. JSON Schema
# requires false != 0 and true != 1. This helper implements correct semantics.
#
# _validates_against: Used for contains/oneOf/not validation where we need to
# check if a value matches a schema without raising an exception.
#
# _make_const_validator / _make_enum_validator: Factory functions that create
# validators using _json_equals. Using factories avoids closure variable issues
# and makes generated code more readable.

_HELPER_FUNCTIONS = {
    "_json_equals": """def _json_equals(a, b):
    \"\"\"Check JSON Schema equality semantics: false != 0, true != 1, deep equality.\"\"\"
    if isinstance(a, bool) != isinstance(b, bool):
        return False
    if isinstance(a, list):
        if not isinstance(b, list) or len(a) != len(b):
            return False
        return all(_json_equals(x, y) for x, y in zip(a, b))
    if isinstance(a, dict):
        if not isinstance(b, dict) or set(a.keys()) != set(b.keys()):
            return False
        return all(_json_equals(a[k], b[k]) for k in a)
    return a == b""",
    "_validates_against": """def _validates_against(validator, value):
    \"\"\"Check if a value validates against a TypeAdapter without raising.\"\"\"
    try:
        validator.validate_python(value)
        return True
    except Exception:
        return False""",
    "_make_const_validator": """def _make_const_validator(expected):
    \"\"\"Create a validator that checks for exact JSON Schema equality with a constant.\"\"\"
    def validator(v):
        if not _json_equals(v, expected):
            raise ValueError(f"Value must equal {expected}")
        return v
    return validator""",
    "_make_enum_validator": """def _make_enum_validator(allowed_values):
    \"\"\"Create a validator that checks if value is in enum using JSON Schema equality.\"\"\"
    def validator(v):
        if not any(_json_equals(v, allowed) for allowed in allowed_values):
            raise ValueError(f"Value must be one of {allowed_values}")
        return v
    return validator""",
}


def get_needed_helpers() -> str:
    """Get all needed helper functions as a single code block."""
    if not _helper_registry:
        return ""

    # Resolve dependencies: _make_const_validator and _make_enum_validator need _json_equals
    if (
        "_make_const_validator" in _helper_registry
        or "_make_enum_validator" in _helper_registry
    ):
        _register_helper("_json_equals")

    # Emit helpers in dependency order
    helpers = []
    if "_json_equals" in _helper_registry:
        helpers.append(_HELPER_FUNCTIONS["_json_equals"])
    if "_validates_against" in _helper_registry:
        helpers.append(_HELPER_FUNCTIONS["_validates_against"])
    if "_make_const_validator" in _helper_registry:
        helpers.append(_HELPER_FUNCTIONS["_make_const_validator"])
    if "_make_enum_validator" in _helper_registry:
        helpers.append(_HELPER_FUNCTIONS["_make_enum_validator"])

    return "\n\n\n".join(helpers)


@dataclass
class RenderResult:
    """Result of rendering a schema node."""

    code: str  # Generated Python code (class definition or type alias)
    type_expr: str  # Type expression (e.g., "str", "MyClass", "list[int]")
    imports: set[str] = field(default_factory=set)  # Required imports


def render(node: SchemaNode, name: str) -> RenderResult:
    """Render a SchemaNode to Pydantic code.

    Args:
        node: The IR node to render
        name: The name to use for the generated type (PascalCase)

    Returns:
        RenderResult with generated code, type expression, and required imports
    """
    match node.kind:
        case "string":
            return render_string(node)
        case "number":
            return render_number(node)
        case "boolean":
            return render_boolean(node)
        case "null":
            return render_null(node)
        case "literal":
            return render_literal(node)
        case "enum":
            return render_enum(node)
        case "any":
            return render_any(node)
        case "never":
            return render_never(node)
        case "array":
            return render_array(node, name)
        case "tuple":
            return render_tuple(node, name)
        case "object":
            return render_object(node, name)
        case "union":
            return render_union(node, name)
        case "oneOf":
            return render_oneof(node, name)
        case "intersection":
            return render_intersection(node, name)
        case "not":
            return render_not(node, name)
        case "conditional":
            return render_conditional(node, name)
        case "typeGuarded":
            return render_type_guarded(node, name)
        case "nullable":
            # Handle nullable by wrapping inner type with | None
            inner_result = render(node.inner, name)
            return RenderResult(
                code=inner_result.code,
                type_expr=_make_nullable(inner_result.type_expr),
                imports=inner_result.imports,
            )
        case "ref":
            return render_ref(node, name)
        case _:
            raise ConversionError(
                f"Unknown IR node kind '{node.kind}' — "
                "the adapter does not know how to render this construct"
            )


def render_string(node: StringNode) -> RenderResult:
    """Render StringNode to Pydantic type with format support.

    Uses StrictStr to ensure JSON Schema semantics where non-strings
    are rejected without coercion.
    """
    imports: set[str] = set()

    # Handle format first - format types override base str type
    if node.format is not None:
        format_result = _render_format(node.format, node.constraints)
        if format_result is not None:
            return format_result
        # Unknown formats fall through to str with constraints

    # No format or unknown format - use StrictStr with constraints
    has_constraints = node.constraints is not None and (
        node.constraints.min_length is not None
        or node.constraints.max_length is not None
        or node.constraints.pattern is not None
    )

    if has_constraints:
        # Use Annotated with StringConstraints for constrained strings
        # Note: StringConstraints with strict=True ensures no coercion
        imports.add("from typing import Annotated")
        imports.add("from pydantic import StringConstraints")

        constraints: list[str] = ["strict=True"]
        if node.constraints.min_length is not None:
            constraints.append(f"min_length={node.constraints.min_length}")
        if node.constraints.max_length is not None:
            constraints.append(f"max_length={node.constraints.max_length}")
        if node.constraints.pattern is not None:
            # Escape the pattern for Python string
            escaped_pattern = node.constraints.pattern.replace("\\", "\\\\").replace(
                "'", "\\'"
            )
            constraints.append(f"pattern=r'{escaped_pattern}'")

        constraint_str = ", ".join(constraints)
        type_expr = f"Annotated[str, StringConstraints({constraint_str})]"
    else:
        imports.add("from pydantic import StrictStr")
        type_expr = "StrictStr"

    return RenderResult(code="", type_expr=type_expr, imports=imports)


def _render_format(fmt: str, constraints: Any) -> RenderResult | None:
    """Render format string types to Pydantic types.

    Returns None to treat format as annotation-only (no validation).

    JSON Schema draft 2020-12 specifies that the format keyword is an
    annotation by default, not a validation keyword. Format validation
    only occurs when explicitly enabled via the format-assertion vocabulary
    in the meta-schema.

    For compliance with the JSON Schema spec, we treat format as annotation-only.
    This means all format values are valid - format is just metadata, not validation.

    Users who want format validation can:
    1. Add explicit validation in their application code
    2. Use Pydantic's built-in types directly (EmailStr, AnyUrl, etc.)
    3. Create custom validators for specific formats

    See: https://json-schema.org/draft/2020-12/json-schema-validation#name-format
    """
    # Always return None to treat format as annotation-only
    # This ensures compliance with JSON Schema draft 2020-12 semantics
    return None

    # NOTE: The code below is preserved for reference if format validation
    # is ever made configurable. For now, it's unreachable.

    imports: set[str] = set()

    # Email format
    if fmt == "email":
        imports.add("from pydantic import EmailStr")
        return RenderResult(code="", type_expr="EmailStr", imports=imports)

    # URI/URL formats
    if fmt == "uri" or fmt == "uri-reference":
        imports.add("from pydantic import AnyUrl")
        return RenderResult(code="", type_expr="AnyUrl", imports=imports)

    if fmt == "iri" or fmt == "iri-reference":
        # IRI is like URI but allows international characters
        imports.add("from pydantic import AnyUrl")
        return RenderResult(code="", type_expr="AnyUrl", imports=imports)

    if fmt == "uri-template":
        # URI template - no built-in Pydantic type, use str
        return None

    # UUID format
    if fmt == "uuid":
        imports.add("from uuid import UUID")
        return RenderResult(code="", type_expr="UUID", imports=imports)

    # Date/time formats
    if fmt == "date":
        imports.add("from datetime import date")
        return RenderResult(code="", type_expr="date", imports=imports)

    if fmt == "date-time":
        imports.add("from datetime import datetime")
        return RenderResult(code="", type_expr="datetime", imports=imports)

    if fmt == "time":
        imports.add("from datetime import time")
        return RenderResult(code="", type_expr="time", imports=imports)

    if fmt == "duration":
        # ISO 8601 duration - no built-in Python type, use str
        # Could use timedelta but duration is more complex
        return None

    # IP address formats
    if fmt == "ipv4":
        imports.add("from pydantic import IPvAnyAddress")
        imports.add("from typing import Annotated")
        imports.add("from pydantic import AfterValidator")
        code = """def _ipv4_validator(v):
    if v.version != 4:
        raise ValueError("Must be IPv4 address")
    return v"""
        return RenderResult(
            code=code,
            type_expr="Annotated[IPvAnyAddress, AfterValidator(_ipv4_validator)]",
            imports=imports,
        )

    if fmt == "ipv6":
        imports.add("from pydantic import IPvAnyAddress")
        imports.add("from typing import Annotated")
        imports.add("from pydantic import AfterValidator")
        code = """def _ipv6_validator(v):
    if v.version != 6:
        raise ValueError("Must be IPv6 address")
    return v"""
        return RenderResult(
            code=code,
            type_expr="Annotated[IPvAnyAddress, AfterValidator(_ipv6_validator)]",
            imports=imports,
        )

    # Hostname format
    if fmt == "hostname" or fmt == "idn-hostname":
        # RFC 1123 hostname validation via regex
        # Note: Simplified regex without lookahead/lookbehind (not supported by pydantic_core)
        imports.add("from typing import Annotated")
        imports.add("from pydantic import StringConstraints")
        # Hostname: labels separated by dots, each label 1-63 chars alphanumeric+hyphen
        hostname_pattern = r"^[A-Za-z0-9]([A-Za-z0-9-]{0,61}[A-Za-z0-9])?(\\.[A-Za-z0-9]([A-Za-z0-9-]{0,61}[A-Za-z0-9])?)*$"
        return RenderResult(
            code="",
            type_expr=f"Annotated[str, StringConstraints(pattern=r'{hostname_pattern}')]",
            imports=imports,
        )

    # Email with internationalization
    if fmt == "idn-email":
        # Pydantic EmailStr handles IDN
        imports.add("from pydantic import EmailStr")
        return RenderResult(code="", type_expr="EmailStr", imports=imports)

    # JSON Pointer
    if fmt == "json-pointer" or fmt == "relative-json-pointer":
        # No built-in type, use str
        return None

    # Regex format (string containing a regex)
    if fmt == "regex":
        # Could validate with re.compile, but just use str for now
        return None

    # Unknown format - return None to fall back to str
    return None


def _is_integer_value(val: float | int) -> bool:
    """Check if a numeric value is effectively an integer."""
    return isinstance(val, int) or (isinstance(val, float) and val == int(val))


def render_number(node: NumberNode) -> RenderResult:
    """Render NumberNode to Pydantic type.

    WHY StrictInt / StrictFloat:
    Python's bool is a subclass of int (True == 1, False == 0). Without strict
    types, Pydantic accepts True for {"type": "integer"} which violates JSON
    Schema semantics. StrictInt/StrictFloat explicitly reject bools.

    WHY NOT StrictInt | StrictFloat for 'number':
    JSON Schema 'number' accepts both ints and floats. You might think we'd use
    StrictInt | StrictFloat, but Pydantic's MultipleOf validator has precision
    bugs when applied to unions. StrictFloat alone works because it accepts
    integer inputs (42 → 42.0) - mathematically equivalent for validation.

    INTEGER + FLOAT multipleOf EDGE CASE:
    Schema {"type": "integer", "multipleOf": 0.123456789} is legal but weird.
    Pydantic's MultipleOf doesn't support non-integer multipleOf for int types.
    We handle this with a custom BeforeValidator that checks both int-ness
    and multipleOf constraint manually.
    """
    imports: set[str] = set()

    # Use strict types to prevent coercion
    if node.integer:
        imports.add("from pydantic import StrictInt")
        base_type = "StrictInt"
    else:
        # JSON Schema 'number' accepts both integers and floats
        # StrictFloat accepts integers too (converts to float), which is fine
        # for validation purposes
        imports.add("from pydantic import StrictFloat")
        base_type = "StrictFloat"

    has_constraints = node.constraints is not None and (
        node.constraints.minimum is not None
        or node.constraints.maximum is not None
        or node.constraints.exclusive_minimum is not None
        or node.constraints.exclusive_maximum is not None
        or node.constraints.multiple_of is not None
    )

    # Check if we have integer type with float multipleOf
    # Pydantic's MultipleOf doesn't support non-integer multipleOf for int types
    needs_custom_validator = (
        node.integer
        and node.constraints is not None
        and node.constraints.multiple_of is not None
        and not _is_integer_value(node.constraints.multiple_of)
    )

    if needs_custom_validator:
        # Use float base type with custom integer+multipleOf check
        # This handles edge cases like type:integer, multipleOf:0.123456789
        imports.add("from typing import Annotated")
        imports.add("from pydantic import BeforeValidator")

        multiple_of = node.constraints.multiple_of
        annotations: list[str] = []

        # Add other constraints if present
        if node.constraints.minimum is not None:
            imports.add("from annotated_types import Ge")
            annotations.append(f"Ge({node.constraints.minimum})")
        if node.constraints.exclusive_minimum is not None:
            imports.add("from annotated_types import Gt")
            annotations.append(f"Gt({node.constraints.exclusive_minimum})")
        if node.constraints.maximum is not None:
            imports.add("from annotated_types import Le")
            annotations.append(f"Le({node.constraints.maximum})")
        if node.constraints.exclusive_maximum is not None:
            imports.add("from annotated_types import Lt")
            annotations.append(f"Lt({node.constraints.exclusive_maximum})")

        # Create inline validator for integer + float multipleOf
        # Uses float base to avoid Pydantic's integer multipleOf restriction
        validator_code = (
            f"BeforeValidator(lambda v: v if (isinstance(v, (int, float)) and "
            f"not isinstance(v, bool) and float(v) == int(v) and "
            f"(v / {multiple_of}) % 1 == 0) else (_ for _ in ()).throw(ValueError('not valid')))"
        )
        annotations.append(validator_code)

        annotation_str = ", ".join(annotations)
        type_expr = f"Annotated[float, {annotation_str}]"
    elif has_constraints:
        # Use Annotated with annotated-types constraints
        imports.add("from typing import Annotated")
        constraint_annotations: list[str] = []

        if node.constraints.minimum is not None:
            imports.add("from annotated_types import Ge")
            constraint_annotations.append(f"Ge({node.constraints.minimum})")
        if node.constraints.exclusive_minimum is not None:
            imports.add("from annotated_types import Gt")
            constraint_annotations.append(f"Gt({node.constraints.exclusive_minimum})")
        if node.constraints.maximum is not None:
            imports.add("from annotated_types import Le")
            constraint_annotations.append(f"Le({node.constraints.maximum})")
        if node.constraints.exclusive_maximum is not None:
            imports.add("from annotated_types import Lt")
            constraint_annotations.append(f"Lt({node.constraints.exclusive_maximum})")
        if node.constraints.multiple_of is not None:
            imports.add("from annotated_types import MultipleOf")
            constraint_annotations.append(f"MultipleOf({node.constraints.multiple_of})")

        annotation_str = ", ".join(constraint_annotations)
        type_expr = f"Annotated[{base_type}, {annotation_str}]"
    else:
        type_expr = base_type

    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_boolean(node: BooleanNode) -> RenderResult:
    """Render BooleanNode to Pydantic type.

    Uses StrictBool to ensure proper JSON Schema semantics where integers
    like 0 and 1 are NOT valid booleans.
    """
    return RenderResult(
        code="", type_expr="StrictBool", imports={"from pydantic import StrictBool"}
    )


def render_null(node: NullNode) -> RenderResult:
    """Render NullNode to Pydantic type."""
    return RenderResult(code="", type_expr="None", imports=set())


def _is_primitive(value: Any) -> bool:
    """Check if a value is a JSON primitive (str, int, float, bool, None)."""
    return isinstance(value, (str, int, float, bool, type(None))) and not isinstance(
        value, (list, dict)
    )


def _format_literal_value(value: Any) -> str:
    """Format a primitive value for use in Literal[...].

    Uses json.dumps for strings to properly escape all control characters.
    """
    import json

    if value is None:
        return "None"
    elif isinstance(value, bool):
        return "True" if value else "False"
    elif isinstance(value, str):
        # use json.dumps to properly escape control characters like \x00
        return json.dumps(value)
    else:
        return repr(value)


def _format_json_value(value: Any) -> str:
    """Format any JSON value as Python literal for code generation."""
    import json

    if value is None:
        return "None"
    elif isinstance(value, bool):
        return "True" if value else "False"
    elif isinstance(value, str):
        # use json.dumps to handle escaping properly
        return json.dumps(value)
    elif isinstance(value, (list, dict)):
        # convert to Python literal representation
        return repr(value)
    else:
        return repr(value)


def _json_value_type(value: Any) -> str:
    """Map a JSON value to its narrowest Python type string.

    Used to replace Any with a concrete type in const/enum annotations where
    Literal[] can't be used (complex values, bool/int conflation).
    """
    if isinstance(value, bool):
        return "bool"
    if isinstance(value, int):
        return "int"
    if isinstance(value, float):
        return "float"
    if isinstance(value, str):
        return "str"
    if value is None:
        return "None"
    if isinstance(value, list):
        return "list"
    if isinstance(value, dict):
        return "dict"
    return "Any"


def _json_values_union_type(values: Sequence[Any]) -> str:
    """Compute the union type for a list of JSON values.

    Deduplicates and orders types deterministically. Falls back to Any only
    if any individual value can't be typed.
    """
    types: list[str] = []
    seen: set[str] = set()
    for v in values:
        t = _json_value_type(v)
        if t == "Any":
            return "Any"
        if t not in seen:
            seen.add(t)
            types.append(t)
    if not types:
        return "Any"
    if len(types) == 1:
        return types[0]
    return " | ".join(types)


def render_literal(node: LiteralNode) -> RenderResult:
    """Render LiteralNode (JSON Schema 'const') to Pydantic type.

    WHY NOT ALWAYS USE Literal[value]:
    Python's typing.Literal has a flaw: Literal[False] accepts 0, and Literal[0]
    accepts False, because bool is a subclass of int. JSON Schema requires
    {"const": false} to reject 0 and {"const": 0} to reject false.

    Solution: For booleans and 0/1 integers, we use a custom validator that
    checks both value AND type equality using _json_equals helper.

    WHY _make_const_validator:
    For complex values (lists, dicts), Literal[] doesn't work at all.
    We use a factory function that creates a validator with the expected value
    baked in. This avoids closure issues and generates readable code like:
    `BeforeValidator(_make_const_validator([1, 2, 3]))`
    """
    value = node.value

    # for booleans and integers (0, 1, etc), we need strict type checking
    # because Python's Literal[False] accepts 0 and Literal[0] accepts False
    needs_strict_check = isinstance(value, bool) or (
        isinstance(value, int) and not isinstance(value, bool) and value in (0, 1)
    )

    if _is_primitive(value) and not needs_strict_check:
        # simple case - use Literal for primitives that don't need strict type check
        imports: set[str] = {"from typing import Literal"}
        literal_repr = _format_literal_value(value)
        type_expr = f"Literal[{literal_repr}]"
        return RenderResult(code="", type_expr=type_expr, imports=imports)

    # complex value or needs strict check - use custom validator with deep equality
    # use narrowest representable type instead of Any
    narrow_type = _json_value_type(value)
    imports = {
        "from typing import Annotated",
        "from pydantic import BeforeValidator",
    }
    if narrow_type == "Any":
        imports.add("from typing import Any")

    expected_repr = _format_json_value(value)

    # Register helper for const validation
    _register_helper("_make_const_validator")

    type_expr = f"Annotated[{narrow_type}, BeforeValidator(_make_const_validator({expected_repr}))]"

    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_enum(node: EnumNode) -> RenderResult:
    """Render EnumNode (JSON Schema 'enum') to Pydantic type.

    Same issues as render_literal - Python's Literal[] treats False==0, True==1.
    Additionally, enum can contain complex values like [1,2,3] or {"a": 1}.

    Strategy:
    - Simple primitives without 0/1/bools: Literal[val1, val2, ...]
    - Has bools or 0/1: _make_enum_validator with _json_equals
    - Has complex values: _make_enum_validator with _json_equals

    The _json_equals helper handles deep structural comparison with proper
    bool/int distinction at every level of nesting.
    """
    # check if all values are primitives and don't need strict type checking
    all_primitive = all(_is_primitive(v) for v in node.values)

    # check if any value needs strict type checking (bool or 0/1)
    needs_strict_check = any(
        isinstance(v, bool)
        or (isinstance(v, int) and not isinstance(v, bool) and v in (0, 1))
        for v in node.values
    )

    if all_primitive and not needs_strict_check:
        # simple case - use Literal for all primitives
        imports: set[str] = {"from typing import Literal"}
        literal_values = [_format_literal_value(v) for v in node.values]
        values_str = ", ".join(literal_values)
        type_expr = f"Literal[{values_str}]"
        return RenderResult(code="", type_expr=type_expr, imports=imports)

    # has complex values or needs strict type checking - use custom validator
    # use narrowest representable union type instead of Any
    narrow_type = _json_values_union_type(node.values)
    imports = {
        "from typing import Annotated",
        "from pydantic import BeforeValidator",
    }
    if narrow_type == "Any":
        imports.add("from typing import Any")

    # Register helper for enum validation
    _register_helper("_make_enum_validator")

    # generate list of allowed values for the validator
    values_repr = "[" + ", ".join(_format_json_value(v) for v in node.values) + "]"
    type_expr = f"Annotated[{narrow_type}, BeforeValidator(_make_enum_validator({values_repr}))]"

    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_any(node: AnyNode) -> RenderResult:
    """Render AnyNode to Pydantic type."""
    imports: set[str] = {"from typing import Any"}
    return RenderResult(code="", type_expr="Any", imports=imports)


def _never_validator(v: Any) -> Any:
    """Validator that always fails - used for Never type."""
    raise ValueError("Never type: no value is valid")


def render_never(node: NeverNode) -> RenderResult:
    """Render NeverNode to Pydantic type.

    Note: Pydantic's TypeAdapter doesn't support typing.Never directly.
    We use Annotated with BeforeValidator that always fails for runtime validation.
    """
    imports: set[str] = {
        "from typing import Annotated, Any",
        "from pydantic import BeforeValidator",
    }
    # The validator function needs to be defined in the generated code
    # We generate it inline since we need it to be available
    code = """def _never_validator(v):
    raise ValueError("Never type: no value is valid")"""
    type_expr = "Annotated[Any, BeforeValidator(_never_validator)]"
    return RenderResult(code=code, type_expr=type_expr, imports=imports)


# Placeholder implementations for complex types (to be implemented in subsequent tasks)


def render_array(node: ArrayNode, name: str) -> RenderResult:
    """Render ArrayNode to Pydantic type with constraints."""
    # Render item type
    item_result = render(node.items, f"{name}Item")
    imports = item_result.imports.copy()
    code_parts = [item_result.code] if item_result.code else []

    base_type = f"list[{item_result.type_expr}]"

    # Check if we have constraints to apply
    has_length_constraints = (
        node.constraints.min_items is not None or node.constraints.max_items is not None
    )
    needs_validators = (
        node.constraints.unique_items or node.constraints.contains is not None
    )

    if has_length_constraints or needs_validators:
        # Need Annotated type with Field or validators
        imports.add("from typing import Annotated")
        annotations = []

        # Length constraints via Field
        if has_length_constraints:
            imports.add("from pydantic import Field")
            field_args = []
            if node.constraints.min_items is not None:
                field_args.append(f"min_length={node.constraints.min_items}")
            if node.constraints.max_items is not None:
                field_args.append(f"max_length={node.constraints.max_items}")
            annotations.append(f"Field({', '.join(field_args)})")

        # uniqueItems constraint via custom validator
        if node.constraints.unique_items:
            imports.add("from pydantic import AfterValidator")
            _register_helper("_json_equals")
            validator_name = f"_unique_{name.lower()}"
            # Uses _json_equals for proper JSON Schema equality semantics
            validator_code = f"""def {validator_name}(v: list) -> list:
    for i in range(len(v)):
        for j in range(i + 1, len(v)):
            if _json_equals(v[i], v[j]):
                raise ValueError("Items must be unique")
    return v"""
            code_parts.append(validator_code)
            annotations.append(f"AfterValidator({validator_name})")

        # contains constraint via custom validator
        if node.constraints.contains is not None:
            imports.add("from pydantic import AfterValidator")
            contains_result = render(
                node.constraints.contains.schema, f"{name}Contains"
            )
            imports.update(contains_result.imports)
            if contains_result.code:
                code_parts.append(contains_result.code)

            validator_name = f"_contains_{name.lower()}"
            min_c = node.constraints.contains.min_contains
            max_c = node.constraints.contains.max_contains

            # Generate validator based on whether we need TypeAdapter
            imports.add("from pydantic import TypeAdapter")
            validator_code = f"""def {validator_name}(v: list) -> list:
    validator = TypeAdapter({contains_result.type_expr})
    matching = sum(1 for item in v if _validates_against(validator, item))
    if matching < {min_c}:
        raise ValueError(f"At least {min_c} items must match contains schema, found {{matching}}")"""
            if max_c is not None:
                validator_code += f"""
    if matching > {max_c}:
        raise ValueError(f"At most {max_c} items must match contains schema, found {{matching}}")"""
            validator_code += "\n    return v"

            code_parts.append(validator_code)
            annotations.append(f"AfterValidator({validator_name})")
            _register_helper("_validates_against")

        type_expr = f"Annotated[{base_type}, {', '.join(annotations)}]"
    else:
        type_expr = base_type

    # Handle unevaluatedItems for simple cases (no applicators - CLI filters complex ones)
    # unevaluatedItems applies to items NOT evaluated by 'items' keyword
    # - If 'items' is AnyNode (no explicit items keyword), items are NOT evaluated
    # - If 'items' is an explicit schema, ALL items are evaluated → unevaluatedItems has no effect
    # We can detect "no items keyword" when items is AnyNode (default)
    if node.unevaluated_items is not None and isinstance(node.items, AnyNode):
        # No explicit 'items' keyword → items are not evaluated
        if node.unevaluated_items is False:
            # unevaluatedItems: false - no unevaluated items allowed
            # Since no items are evaluated, array must be empty
            imports.add("from typing import Annotated")
            imports.add("from pydantic import AfterValidator")
            validator_name = f"_unevaluated_{name.lower()}"
            validator_code = f"""def {validator_name}(v: list) -> list:
    if len(v) > 0:
        raise ValueError("unevaluatedItems: false - no items allowed")
    return v"""
            code_parts.append(validator_code)
            # Add validator to type expression
            if "Annotated" in type_expr:
                # Already has annotations, add another
                type_expr = (
                    type_expr.rstrip("]") + f", AfterValidator({validator_name})]"
                )
            else:
                type_expr = f"Annotated[{type_expr}, AfterValidator({validator_name})]"
        else:
            # unevaluatedItems is a schema - validate all items against it
            # Since no items are evaluated by 'items', ALL items are unevaluated
            imports.add("from typing import Annotated")
            imports.add("from pydantic import AfterValidator, TypeAdapter")
            unevaluated_result = render(node.unevaluated_items, f"{name}Unevaluated")
            imports.update(unevaluated_result.imports)
            if unevaluated_result.code:
                code_parts.append(unevaluated_result.code)

            validator_name = f"_unevaluated_{name.lower()}"
            validator_code = f"""def {validator_name}(v: list) -> list:
    validator = TypeAdapter({unevaluated_result.type_expr})
    for i, item in enumerate(v):
        try:
            validator.validate_python(item)
        except Exception as e:
            raise ValueError(f"Item {{i}} failed unevaluatedItems schema: {{e}}")
    return v"""
            code_parts.append(validator_code)
            # Add validator to type expression
            if "Annotated" in type_expr:
                type_expr = (
                    type_expr.rstrip("]") + f", AfterValidator({validator_name})]"
                )
            else:
                type_expr = f"Annotated[{type_expr}, AfterValidator({validator_name})]"

    code = "\n\n\n".join(code_parts) if code_parts else ""
    return RenderResult(code=code, type_expr=type_expr, imports=imports)


def render_tuple(node: TupleNode, name: str) -> RenderResult:
    """Render TupleNode to Pydantic type with rest items."""
    imports: set[str] = set()
    code_parts: list[str] = []
    item_types: list[str] = []

    # Render prefix items (fixed positions)
    for i, item in enumerate(node.prefix_items):
        item_result = render(item, f"{name}Item{i}")
        imports.update(item_result.imports)
        if item_result.code:
            code_parts.append(item_result.code)
        item_types.append(item_result.type_expr)

    # Handle rest_items (additional items beyond prefix)
    if node.rest_items is not None and node.rest_items is not False:
        # rest_items is a schema - means variable length tuple
        rest_result = render(node.rest_items, f"{name}Rest")
        imports.update(rest_result.imports)
        if rest_result.code:
            code_parts.append(rest_result.code)

        imports.add("from typing import Annotated")
        imports.add("from pydantic import BeforeValidator, TypeAdapter")

        validator_name = f"_tuple_{name.lower()}"

        # Build validator that checks prefix items and rest items
        validator_lines = [f"def {validator_name}(v) -> tuple:"]
        validator_lines.append("    if not isinstance(v, tuple):")
        validator_lines.append(
            "        v = tuple(v) if hasattr(v, '__iter__') else (v,)"
        )
        validator_lines.append(f"    if len(v) < {len(node.prefix_items)}:")
        validator_lines.append(
            f"        raise ValueError(f'Tuple must have at least {len(node.prefix_items)} items, got {{len(v)}}')"
        )
        validator_lines.append("    validated = []")

        # Validate prefix items
        for i, item_type in enumerate(item_types):
            validator_lines.append(f"    # Validate item {i}")
            validator_lines.append(
                f"    prefix_{i}_validator = TypeAdapter({item_type})"
            )
            validator_lines.append(f"    try:")
            validator_lines.append(
                f"        validated.append(prefix_{i}_validator.validate_python(v[{i}]))"
            )
            validator_lines.append(f"    except Exception as e:")
            validator_lines.append(
                f"        raise ValueError(f'Item at index {i} invalid: {{e}}')"
            )

        # Validate rest items
        validator_lines.append(f"    # Validate rest items")
        validator_lines.append(
            f"    rest_validator = TypeAdapter({rest_result.type_expr})"
        )
        validator_lines.append(f"    for i in range({len(node.prefix_items)}, len(v)):")
        validator_lines.append(f"        try:")
        validator_lines.append(
            f"            validated.append(rest_validator.validate_python(v[i]))"
        )
        validator_lines.append(f"        except Exception as e:")
        validator_lines.append(
            f"            raise ValueError(f'Item at index {{i}} must match rest schema: {{e}}')"
        )
        validator_lines.append("    return tuple(validated)")

        code_parts.append("\n".join(validator_lines))

        # narrower static type: union of all prefix types + rest type
        all_elem_types = _union_base_type(item_types + [rest_result.type_expr])
        if all_elem_types == "Any":
            imports.add("from typing import Any")
        type_expr = f"Annotated[tuple[{all_elem_types}, ...], BeforeValidator({validator_name})]"
    elif node.rest_items is False:
        # rest_items is False - no additional items allowed beyond prefix
        # JSON Schema: prefixItems + items:false means 0 to N items, where N = len(prefixItems)
        # Each item that exists must match its corresponding prefixItem schema
        if item_types:
            imports.add("from typing import Annotated")
            imports.add("from pydantic import BeforeValidator, TypeAdapter")

            validator_name = f"_tuple_{name.lower()}"
            prefix_len = len(node.prefix_items)

            # Build validator that allows 0 to N items
            validator_lines = [f"def {validator_name}(v) -> tuple:"]
            validator_lines.append("    if not isinstance(v, tuple):")
            validator_lines.append(
                "        v = tuple(v) if hasattr(v, '__iter__') else (v,)"
            )
            validator_lines.append(f"    if len(v) > {prefix_len}:")
            validator_lines.append(
                f"        raise ValueError(f'Tuple must have at most {prefix_len} items, got {{len(v)}}')"
            )
            validator_lines.append("    validated = []")

            # Validate prefix items that exist
            for i, item_type in enumerate(item_types):
                validator_lines.append(f"    # Validate item {i} if present")
                validator_lines.append(f"    if len(v) > {i}:")
                validator_lines.append(
                    f"        prefix_{i}_validator = TypeAdapter({item_type})"
                )
                validator_lines.append(f"        try:")
                validator_lines.append(
                    f"            validated.append(prefix_{i}_validator.validate_python(v[{i}]))"
                )
                validator_lines.append(f"        except Exception as e:")
                validator_lines.append(
                    f"            raise ValueError(f'Item at index {i} invalid: {{e}}')"
                )

            validator_lines.append("    return tuple(validated)")

            code_parts.append("\n".join(validator_lines))

            # closed tuple allows 0..N items, so positional types (which require
            # exact length) would reject shorter arrays. use union of prefix types.
            elem_union = _union_base_type(item_types)
            if elem_union == "Any":
                imports.add("from typing import Any")
            type_expr = f"Annotated[tuple[{elem_union}, ...], BeforeValidator({validator_name})]"
        else:
            type_expr = "tuple[()]"  # Empty tuple (no items allowed)
    else:
        # rest_items is None - JSON Schema default: additional items ARE allowed
        # When prefixItems is present without items/additionalItems keyword, items can be:
        # - empty array (valid)
        # - partial array (valid, validates items that exist)
        # - full array (valid, validates all prefix items)
        # - extra items (valid, no validation on extras)
        if item_types:
            imports.add("from typing import Annotated, Any")
            imports.add("from pydantic import BeforeValidator, TypeAdapter")

            validator_name = f"_tuple_{name.lower()}"

            # Build validator that validates items that exist, allows shorter/longer arrays
            validator_lines = [f"def {validator_name}(v) -> tuple:"]
            validator_lines.append("    if not isinstance(v, tuple):")
            validator_lines.append(
                "        v = tuple(v) if hasattr(v, '__iter__') else (v,)"
            )
            validator_lines.append("    validated = list(v)")  # Start with all items

            # Validate prefix items that exist (allow shorter arrays)
            for i, item_type in enumerate(item_types):
                validator_lines.append(f"    # Validate item {i} if present")
                validator_lines.append(f"    if len(v) > {i}:")
                validator_lines.append(
                    f"        prefix_{i}_validator = TypeAdapter({item_type})"
                )
                validator_lines.append(f"        try:")
                validator_lines.append(
                    f"            validated[{i}] = prefix_{i}_validator.validate_python(v[{i}])"
                )
                validator_lines.append(f"        except Exception as e:")
                validator_lines.append(
                    f"            raise ValueError(f'Item at index {i} invalid: {{e}}')"
                )

            # Rest items are allowed (any type) - no validation needed
            validator_lines.append("    return tuple(validated)")

            code_parts.append("\n".join(validator_lines))

            # open tuple: extras untyped → tuple[Any, ...] (can't narrow)
            type_expr = f"Annotated[tuple[Any, ...], BeforeValidator({validator_name})]"
        else:
            # No prefix items, rest_items is None - empty tuple or any tuple allowed
            # In JSON Schema, no items constraint means all arrays valid
            type_expr = "tuple[Any, ...]"

    # Handle tuple constraints (min/max items, uniqueItems, contains)
    needs_validators = (
        node.constraints.min_items is not None
        or node.constraints.max_items is not None
        or node.constraints.unique_items
        or node.constraints.contains is not None
    )

    if needs_validators:
        imports.add("from typing import Annotated")
        imports.add("from pydantic import AfterValidator")

        constraint_validator_name = f"_tuple_constraints_{name.lower()}"
        validator_lines = [f"def {constraint_validator_name}(v: tuple) -> tuple:"]

        if node.constraints.min_items is not None:
            validator_lines.append(f"    if len(v) < {node.constraints.min_items}:")
            validator_lines.append(
                f"        raise ValueError(f'Tuple must have at least {node.constraints.min_items} items, got {{len(v)}}')"
            )

        if node.constraints.max_items is not None:
            validator_lines.append(f"    if len(v) > {node.constraints.max_items}:")
            validator_lines.append(
                f"        raise ValueError(f'Tuple must have at most {node.constraints.max_items} items, got {{len(v)}}')"
            )

        if node.constraints.unique_items:
            # Use JSON Schema equality semantics
            _register_helper("_json_equals")
            validator_lines.append("    for i in range(len(v)):")
            validator_lines.append("        for j in range(i + 1, len(v)):")
            validator_lines.append("            if _json_equals(v[i], v[j]):")
            validator_lines.append(
                "                raise ValueError('Tuple items must be unique')"
            )

        if node.constraints.contains is not None:
            imports.add("from pydantic import TypeAdapter")
            contains_result = render(
                node.constraints.contains.schema, f"{name}Contains"
            )
            imports.update(contains_result.imports)
            if contains_result.code:
                code_parts.append(contains_result.code)

            min_c = node.constraints.contains.min_contains
            max_c = node.constraints.contains.max_contains

            validator_lines.append(
                f"    validator = TypeAdapter({contains_result.type_expr})"
            )
            validator_lines.append(
                "    matching = sum(1 for item in v if _validates_against(validator, item))"
            )
            validator_lines.append(f"    if matching < {min_c}:")
            validator_lines.append(
                f"        raise ValueError(f'At least {min_c} items must match contains schema, found {{matching}}')"
            )
            if max_c is not None:
                validator_lines.append(f"    if matching > {max_c}:")
                validator_lines.append(
                    f"        raise ValueError(f'At most {max_c} items must match contains schema, found {{matching}}')"
                )

            _register_helper("_validates_against")

        validator_lines.append("    return v")
        code_parts.append("\n".join(validator_lines))

        type_expr = (
            f"Annotated[{type_expr}, AfterValidator({constraint_validator_name})]"
        )

    code = "\n\n\n".join(code_parts) if code_parts else ""
    return RenderResult(code=code, type_expr=type_expr, imports=imports)


def render_object(node: ObjectNode, name: str) -> RenderResult:
    """Render ObjectNode to Pydantic BaseModel class.

    CONFIGDICT extra='forbid' vs 'allow':
    - additionalProperties: false → extra='forbid' (Pydantic handles it)
    - additionalProperties: schema → extra='allow' + model_validator
    - patternProperties + additionalProperties: false → extra='allow' + manual check
      (can't use forbid because pattern-matching props would be rejected)
    - unevaluatedProperties: similar logic to additionalProperties

    WHY model_validator FOR ADVANCED FEATURES:
    Pydantic's field-level validation can't express constraints like:
    - "if property X exists, property Y must also exist" (dependencies)
    - "property names must match this pattern" (propertyNames)
    - "extra properties must match this schema" (additionalProperties: schema)

    We use @model_validator(mode='after') which runs after all fields are set,
    giving access to self.model_dump() for cross-field validation.

    PROPERTY NAME ALIASING:
    JSON Schema allows property names like "foo-bar" or "123" which aren't
    valid Python identifiers. We sanitize these to field names (foo_bar, _123)
    and use Field(alias="original-name") to preserve the JSON key.
    """
    imports: set[str] = {"from pydantic import BaseModel"}
    nested_classes: list[str] = []
    field_lines: list[str] = []
    validators: list[str] = []

    # Track used field names to avoid collisions from sanitized names
    used_field_names: set[str] = set()

    # Process each property
    for prop_name, prop_def in node.properties:
        # Render the property's schema
        # Use PascalCase nested class name based on property name
        nested_name = f"{name}{_to_pascal_case(prop_name)}"
        prop_result = render(prop_def.schema, nested_name)
        imports.update(prop_result.imports)

        # Collect nested class definitions
        if prop_result.code:
            nested_classes.append(prop_result.code)

        # Determine if property is required
        is_required = prop_def.required

        # Check if the property schema is a union containing null (nullable)
        is_nullable = _is_nullable_type(prop_def.schema)

        # Build the field definition
        type_expr = prop_result.type_expr

        # Check if property name needs aliasing (not a valid Python identifier)
        needs_alias = not _is_valid_python_identifier(prop_name)
        if needs_alias:
            field_name = _sanitize_property_name(prop_name, used_field_names)
        else:
            field_name = prop_name
            used_field_names.add(prop_name)
        escaped_alias = ""

        if needs_alias:
            imports.add("from pydantic import Field")
            # Escape property name for alias string
            escaped_alias = _escape_for_double_quotes(prop_name)

        if is_required:
            if is_nullable:
                # Required but nullable: field: Type | None (no default)
                type_expr = _make_nullable(type_expr)
                if needs_alias:
                    field_lines.append(
                        f'    {field_name}: {type_expr} = Field(alias="{escaped_alias}")'
                    )
                else:
                    field_lines.append(f"    {field_name}: {type_expr}")
            else:
                # Required and not nullable: field: Type
                if needs_alias:
                    field_lines.append(
                        f'    {field_name}: {type_expr} = Field(alias="{escaped_alias}")'
                    )
                else:
                    field_lines.append(f"    {field_name}: {type_expr}")
        else:
            # Optional: field: Type | None = None
            type_expr = _make_nullable(type_expr)
            if needs_alias:
                field_lines.append(
                    f'    {field_name}: {type_expr} = Field(default=None, alias="{escaped_alias}")'
                )
            else:
                field_lines.append(f"    {field_name}: {type_expr} = None")

    # Determine if there are required fields not covered by properties
    property_names = {prop_name for prop_name, _ in node.properties}
    uncovered_required = [r for r in node.required if r not in property_names]

    # Handle additionalProperties and unevaluatedProperties
    # unevaluatedProperties: false forbids any properties not evaluated by properties/patternProperties
    # When additionalProperties is not set and unevaluatedProperties: false, forbid extra properties
    config_line = None
    if node.additional_properties is False:
        # When patternProperties exist, we need to allow extra fields and validate manually
        # because Pydantic's extra='forbid' would reject pattern-matching properties
        if node.pattern_properties:
            imports.add("from pydantic import ConfigDict")
            config_line = "    model_config = ConfigDict(extra='allow')"
        else:
            imports.add("from pydantic import ConfigDict")
            config_line = "    model_config = ConfigDict(extra='forbid')"
    elif node.additional_properties is True:
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='allow')"
    # If additional_properties is a schema, we allow extra but don't constrain type
    # (Pydantic doesn't support typed extra fields directly in model_config)
    elif node.additional_properties is not None and not isinstance(
        node.additional_properties, bool
    ):
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='allow')"
    # unevaluatedProperties: false - forbid any properties not covered by properties/patternProperties
    # This is the "simple case" without applicators (allOf, anyOf, etc.) which CLI filters out
    elif node.unevaluated_properties is False:
        if node.pattern_properties:
            # Need to allow extra for patternProperties matching, validate manually
            imports.add("from pydantic import ConfigDict")
            config_line = "    model_config = ConfigDict(extra='allow')"
        else:
            imports.add("from pydantic import ConfigDict")
            config_line = "    model_config = ConfigDict(extra='forbid')"
    # unevaluatedProperties with a schema - need to allow extras and validate them
    elif node.unevaluated_properties is not None:
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='allow')"
    # If there are required fields without properties, allow extra fields to accept them
    elif uncovered_required:
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='allow')"
    # patternProperties, propertyNames, dependencies, min/maxProperties need access to extra fields
    elif (
        node.pattern_properties
        or node.property_names is not None
        or node.dependencies
        or node.min_properties is not None
        or node.max_properties is not None
    ):
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='allow')"

    # Generate validators for advanced object features
    needs_validator = (
        node.pattern_properties
        or node.property_names is not None
        or node.min_properties is not None
        or node.max_properties is not None
        or node.dependencies
        or node.unevaluated_properties is not None
        or uncovered_required  # Need validator for required fields without properties
        or (
            node.additional_properties is not None
            and not isinstance(node.additional_properties, bool)
        )  # Need validator for additionalProperties schema
        or (
            node.additional_properties is False and node.pattern_properties
        )  # Need validator when additionalProperties=false with patternProperties
    )

    if needs_validator:
        imports.add("from pydantic import model_validator")
        validator_lines = ["    @model_validator(mode='after')"]
        validator_lines.append("    def _validate_advanced(self):")
        validator_lines.append("        # Advanced object validations")

        # patternProperties: validate extra properties against regex patterns
        if node.pattern_properties:
            imports.add("import re")
            imports.add("from pydantic import TypeAdapter")

            for pattern_def in node.pattern_properties:
                # Render the pattern schema
                pattern_result = render(
                    pattern_def.schema, f"{name}Pattern{len(nested_classes)}"
                )
                imports.update(pattern_result.imports)
                if pattern_result.code:
                    nested_classes.append(pattern_result.code)

                # Generate validation code for this pattern
                escaped_pattern = pattern_def.pattern.replace("\\", "\\\\").replace(
                    "'", "\\'"
                )
                validator_lines.append(
                    f"        # Validate properties matching pattern: {escaped_pattern}"
                )
                validator_lines.append(
                    f"        pattern_{len(validators)} = re.compile(r'{escaped_pattern}')"
                )
                validator_lines.append(
                    f"        validator_{len(validators)} = TypeAdapter({pattern_result.type_expr})"
                )
                validator_lines.append(
                    "        for key, value in self.model_dump(exclude_unset=True).items():"
                )
                validator_lines.append(
                    f"            if pattern_{len(validators)}.search(key):"
                )
                validator_lines.append("                try:")
                validator_lines.append(
                    f"                    validator_{len(validators)}.validate_python(value)"
                )
                validator_lines.append("                except Exception as e:")
                validator_lines.append(
                    f"                    raise ValueError(f'Property {{key}} must match pattern {escaped_pattern}: {{e}}')"
                )

        # additionalProperties with schema: validate extra properties that don't match properties or patternProperties
        if node.additional_properties is not None and not isinstance(
            node.additional_properties, bool
        ):
            imports.add("from pydantic import TypeAdapter")

            # Render the additionalProperties schema
            additional_result = render(node.additional_properties, f"{name}Additional")
            imports.update(additional_result.imports)
            if additional_result.code:
                nested_classes.append(additional_result.code)

            # Build set of declared property names
            declared_props = {prop_name for prop_name, _ in node.properties}

            validator_lines.append("        # Validate additionalProperties")
            validator_lines.append(
                f"        declared_props = {{{', '.join(repr(p) for p in declared_props)}}}"
            )

            # Compile patternProperties patterns if they exist
            if node.pattern_properties:
                imports.add("import re")
                validator_lines.append("        # Compile patternProperties patterns")
                validator_lines.append("        pattern_regexes = [")
                for pattern_def in node.pattern_properties:
                    escaped_pattern = pattern_def.pattern.replace("\\", "\\\\").replace(
                        "'", "\\'"
                    )
                    validator_lines.append(
                        f"            re.compile(r'{escaped_pattern}'),"
                    )
                validator_lines.append("        ]")

            validator_lines.append(
                f"        additional_validator = TypeAdapter({additional_result.type_expr})"
            )
            validator_lines.append(
                "        for key, value in self.model_dump(exclude_unset=True).items():"
            )
            validator_lines.append(
                "            # Skip properties defined in 'properties'"
            )
            validator_lines.append("            if key in declared_props:")
            validator_lines.append("                continue")

            # Skip properties matching patternProperties
            if node.pattern_properties:
                validator_lines.append(
                    "            # Skip properties matching patternProperties"
                )
                validator_lines.append(
                    "            if any(pattern.search(key) for pattern in pattern_regexes):"
                )
                validator_lines.append("                continue")

            validator_lines.append(
                "            # Validate additional property against schema"
            )
            validator_lines.append("            try:")
            validator_lines.append(
                "                additional_validator.validate_python(value)"
            )
            validator_lines.append("            except Exception as e:")
            validator_lines.append(
                "                raise ValueError(f'Additional property {key} invalid: {e}')"
            )

        # additionalProperties=false with patternProperties: forbid properties not matching patterns or properties
        if node.additional_properties is False and node.pattern_properties:
            imports.add("import re")

            # Build set of declared property names
            declared_props = {prop_name for prop_name, _ in node.properties}

            validator_lines.append(
                "        # Validate additionalProperties=false with patternProperties"
            )
            validator_lines.append(
                f"        declared_props = {{{', '.join(repr(p) for p in declared_props)}}}"
            )

            # Compile patternProperties patterns
            validator_lines.append("        # Compile patternProperties patterns")
            validator_lines.append("        pattern_regexes = [")
            for pattern_def in node.pattern_properties:
                escaped_pattern = pattern_def.pattern.replace("\\", "\\\\").replace(
                    "'", "\\'"
                )
                validator_lines.append(f"            re.compile(r'{escaped_pattern}'),")
            validator_lines.append("        ]")

            validator_lines.append(
                "        for key in self.model_dump(exclude_unset=True).keys():"
            )
            validator_lines.append(
                "            # Skip properties defined in 'properties'"
            )
            validator_lines.append("            if key in declared_props:")
            validator_lines.append("                continue")
            validator_lines.append(
                "            # Skip properties matching patternProperties"
            )
            validator_lines.append(
                "            if any(pattern.search(key) for pattern in pattern_regexes):"
            )
            validator_lines.append("                continue")
            validator_lines.append(
                "            # Reject all other properties (additionalProperties=false)"
            )
            validator_lines.append(
                "            raise ValueError(f'Additional property {key} not allowed')"
            )

        # propertyNames: validate all property keys against schema
        if node.property_names is not None:
            imports.add("from pydantic import TypeAdapter")
            names_result = render(node.property_names, f"{name}PropertyNames")
            imports.update(names_result.imports)
            if names_result.code:
                nested_classes.append(names_result.code)

            validator_lines.append("        # Validate all property names")
            validator_lines.append(
                f"        names_validator = TypeAdapter({names_result.type_expr})"
            )
            validator_lines.append(
                "        for key in self.model_dump(exclude_unset=True).keys():"
            )
            validator_lines.append("            try:")
            validator_lines.append(
                "                names_validator.validate_python(key)"
            )
            validator_lines.append("            except Exception as e:")
            validator_lines.append(
                "                raise ValueError(f'Property name {key} invalid: {e}')"
            )

        # minProperties / maxProperties: count properties
        if node.min_properties is not None or node.max_properties is not None:
            validator_lines.append("        # Validate property count")
            validator_lines.append(
                "        prop_count = len(self.model_dump(exclude_unset=True))"
            )

            if node.min_properties is not None:
                validator_lines.append(
                    f"        if prop_count < {node.min_properties}:"
                )
                validator_lines.append(
                    f"            raise ValueError(f'Object must have at least {node.min_properties} properties, got {{prop_count}}')"
                )

            if node.max_properties is not None:
                validator_lines.append(
                    f"        if prop_count > {node.max_properties}:"
                )
                validator_lines.append(
                    f"            raise ValueError(f'Object must have at most {node.max_properties} properties, got {{prop_count}}')"
                )

        # dependencies: if property exists, require other properties or validate schema
        if node.dependencies:
            imports.add("from pydantic import TypeAdapter")

            # Use model_dump(exclude_unset=True) to only include properties that were actually set
            # This correctly distinguishes between "property not in input" vs "property set to null"
            validator_lines.append(
                "        _all_props = self.model_dump(exclude_unset=True)"
            )

            for prop_name, dependency in node.dependencies:
                # Use single-quote escaping for conditions (in single-quoted strings)
                escaped_prop_single = _escape_for_single_quotes(prop_name)
                validator_lines.append(
                    f"        # Dependency for property: {escaped_prop_single}"
                )
                # JSON Schema: check if property EXISTS (not if it has non-null value)
                validator_lines.append(
                    f"        if '{escaped_prop_single}' in _all_props:"
                )

                if dependency.kind == "property":
                    # Property dependency: require other properties to exist
                    if not dependency.required_properties:
                        # Empty dependency list - always valid, add pass for valid Python syntax
                        validator_lines.append("            pass")
                    else:
                        for required_prop in dependency.required_properties:
                            escaped_req_single = _escape_for_single_quotes(
                                required_prop
                            )
                            # JSON Schema: required property just needs to exist (even with null value)
                            validator_lines.append(
                                f"            if '{escaped_req_single}' not in _all_props:"
                            )
                            validator_lines.append(
                                f"                raise ValueError(f'When {escaped_prop_single} is present, {escaped_req_single} is required')"
                            )
                elif dependency.kind == "schema":
                    # Schema dependency: validate entire object against schema
                    dep_result = render(
                        dependency.schema, f"{name}Dependency{len(nested_classes)}"
                    )
                    imports.update(dep_result.imports)
                    if dep_result.code:
                        nested_classes.append(dep_result.code)

                    validator_lines.append(
                        f"            dep_validator = TypeAdapter({dep_result.type_expr})"
                    )
                    validator_lines.append("            try:")
                    validator_lines.append(
                        "                dep_validator.validate_python(_all_props)"
                    )
                    validator_lines.append("            except Exception as e:")
                    validator_lines.append(
                        f"                raise ValueError(f'When {escaped_prop_single} is present, object must match dependency schema: {{e}}')"
                    )

        # unevaluatedProperties: properties not "evaluated" by properties/patternProperties
        # In simple cases (no applicators), "evaluated" means:
        # - Declared in 'properties' keyword
        # - Matches a pattern in 'patternProperties' keyword
        # CLI filters complex cases with applicators (allOf, anyOf, etc.)
        if node.unevaluated_properties is not None:
            # Get declared property names (these are "evaluated")
            declared_props = {prop_name for prop_name, _ in node.properties}
            patterns = (
                [p.pattern for p in node.pattern_properties]
                if node.pattern_properties
                else []
            )

            if node.unevaluated_properties is False:
                # unevaluatedProperties: false - forbid any property not in properties/patternProperties
                # When there are patternProperties, we need manual validation (can't use extra='forbid')
                if patterns:
                    imports.add("import re")
                    validator_lines.append(
                        "        # Validate unevaluatedProperties: false"
                    )
                    validator_lines.append(
                        f"        _declared_props = {{{', '.join(repr(p) for p in declared_props)}}}"
                    )
                    validator_lines.append(f"        _patterns = {patterns!r}")
                    validator_lines.append(
                        "        for key in self.model_dump(exclude_unset=True).keys():"
                    )
                    validator_lines.append("            if key in _declared_props:")
                    validator_lines.append(
                        "                continue  # evaluated by properties"
                    )
                    validator_lines.append(
                        "            if any(re.search(p, key) for p in _patterns):"
                    )
                    validator_lines.append(
                        "                continue  # evaluated by patternProperties"
                    )
                    validator_lines.append(
                        "            raise ValueError(f'Unevaluated property not allowed: {key}')"
                    )
                # else: extra='forbid' already handles this case (set above)
            else:
                # unevaluatedProperties is a schema - validate unevaluated properties against it
                imports.add("from pydantic import TypeAdapter")
                if patterns:
                    imports.add("import re")
                unevaluated_result = render(
                    node.unevaluated_properties, f"{name}Unevaluated"
                )
                imports.update(unevaluated_result.imports)
                if unevaluated_result.code:
                    nested_classes.append(unevaluated_result.code)

                validator_lines.append(
                    "        # Validate unevaluated properties against schema"
                )
                validator_lines.append(
                    f"        _declared_props = {{{', '.join(repr(p) for p in declared_props)}}}"
                )
                if patterns:
                    validator_lines.append(f"        _patterns = {patterns!r}")
                validator_lines.append(
                    f"        _unevaluated_validator = TypeAdapter({unevaluated_result.type_expr})"
                )
                validator_lines.append(
                    "        for key, value in self.model_dump(exclude_unset=True).items():"
                )
                validator_lines.append("            if key in _declared_props:")
                validator_lines.append(
                    "                continue  # evaluated by properties"
                )
                if patterns:
                    validator_lines.append(
                        "            if any(re.search(p, key) for p in _patterns):"
                    )
                    validator_lines.append(
                        "                continue  # evaluated by patternProperties"
                    )
                validator_lines.append("            try:")
                validator_lines.append(
                    "                _unevaluated_validator.validate_python(value)"
                )
                validator_lines.append("            except Exception as e:")
                validator_lines.append(
                    "                raise ValueError(f'Unevaluated property {key} invalid: {e}')"
                )

        # required fields without corresponding properties
        # model_dump(exclude_unset=True) includes only properties that were actually set
        if uncovered_required:
            validator_lines.append("        # Validate required fields")
            validator_lines.append(
                "        _all_props = self.model_dump(exclude_unset=True)"
            )
            for req_field in uncovered_required:
                # For condition checks, use single-quoted strings
                escaped_single = _escape_for_single_quotes(req_field)
                # For error messages using double quotes, escape differently
                escaped_double = _escape_for_double_quotes(req_field)
                validator_lines.append(
                    f"        if '{escaped_single}' not in _all_props:"
                )
                validator_lines.append(
                    f"            raise ValueError(\"Missing required property: '{escaped_double}'\")"
                )

        validator_lines.append("        return self")
        validators.append("\n".join(validator_lines))

    # Build the class definition
    class_lines: list[str] = [f"class {name}(BaseModel):"]

    # Add config if present
    if config_line:
        class_lines.append(config_line)

    # Add fields or pass if empty
    if field_lines:
        class_lines.extend(field_lines)
    elif not config_line and not validators:
        class_lines.append("    pass")

    # Add validators
    if validators:
        class_lines.append("")  # Blank line before validators
        for validator in validators:
            class_lines.append(validator)

    # Combine nested classes and main class
    class_code = "\n".join(class_lines)
    if nested_classes:
        code = "\n\n\n".join(nested_classes) + "\n\n\n" + class_code
    else:
        code = class_code

    return RenderResult(code=code, type_expr=name, imports=imports)


def _to_pascal_case(name: str) -> str:
    """Convert a name to PascalCase for nested class names."""
    # Handle snake_case and camelCase
    parts = name.replace("-", "_").split("_")
    return "".join(part.capitalize() for part in parts if part)


def _is_nullable_type(node: SchemaNode) -> bool:
    """Check if a schema node represents a nullable type (union with null)."""
    if node.kind == "null":
        return True
    if node.kind == "union":
        # Check if any variant is null
        for variant in node.variants:
            if variant.kind == "null":
                return True
    return False


def render_union(node: UnionNode, name: str) -> RenderResult:
    """Render UnionNode (JSON Schema 'anyOf') to Pydantic union type.

    ANYOF vs ONEOF:
    - anyOf: at least one schema must match (this function)
    - oneOf: exactly one schema must match (render_oneof)

    Pydantic's Union type naturally implements anyOf - it tries each variant
    in order until one validates. No special handling needed.

    DISCRIMINATED UNION OPTIMIZATION:
    When all variants are objects with a common property that has different
    literal values in each variant (like {"type": "cat"} vs {"type": "dog"}),
    Pydantic can use that property to directly select the right variant
    instead of trying each one. This is both faster and gives better errors.

    We detect this pattern automatically by scanning all variant schemas
    for a common property with Literal values. If found, we emit:
    `Annotated[Union[Cat, Dog], Field(discriminator='type')]`

    This is an adapter-level optimization - JSON Schema doesn't have explicit
    discriminators, we just recognize the pattern.
    """
    imports: set[str] = set()
    code_parts: list[str] = []
    variant_types: list[str] = []

    # Render all variants
    for i, variant in enumerate(node.variants):
        variant_result = render(variant, f"{name}Variant{i}")
        imports.update(variant_result.imports)
        variant_types.append(variant_result.type_expr)
        if variant_result.code:
            code_parts.append(variant_result.code)

    # Deduplicate variant types to avoid "None | None" errors
    unique_variant_types = []
    seen = set()
    for vtype in variant_types:
        if vtype not in seen:
            unique_variant_types.append(vtype)
            seen.add(vtype)

    # Detect if this is a discriminated union (all variants are objects with a common discriminator)
    discriminator_field = _detect_discriminator(node.variants)

    if discriminator_field is not None and len(unique_variant_types) > 1:
        # Discriminated union - use Annotated with Field(discriminator=...)
        imports.add("from typing import Annotated, Union")
        imports.add("from pydantic import Field")

        # Build Union[variant1, variant2, ...]
        union_type = f"Union[{', '.join(unique_variant_types)}]"
        type_expr = (
            f"Annotated[{union_type}, Field(discriminator='{discriminator_field}')]"
        )
    else:
        # Simple union using | syntax
        # If only one unique type after deduplication, use it directly
        if len(unique_variant_types) == 1:
            type_expr = unique_variant_types[0]
        else:
            type_expr = " | ".join(unique_variant_types)

    code = "\n\n\n".join(code_parts) if code_parts else ""
    return RenderResult(code=code, type_expr=type_expr, imports=imports)


def _detect_discriminator(variants: tuple[SchemaNode, ...]) -> str | None:
    """Detect if a union has a discriminator field.

    Returns the discriminator field name if:
    1. All variants are objects
    2. All variants have the same property
    3. That property has a literal type (different value in each variant)

    This enables Pydantic's discriminated union optimization.
    """
    if len(variants) < 2:
        return None

    # Check all variants are objects
    if not all(v.kind == "object" for v in variants):
        return None

    # Cast to ObjectNode for type checker
    object_variants: list[ObjectNode] = [v for v in variants if v.kind == "object"]

    # Find common properties across all variants
    first_obj = object_variants[0]
    prop_names = {prop_name for prop_name, _ in first_obj.properties}

    for variant in object_variants[1:]:
        variant_props = {prop_name for prop_name, _ in variant.properties}
        prop_names &= variant_props

    if not prop_names:
        return None

    # Check each common property to see if it's a literal with different values
    for prop_name in prop_names:
        # Get the property schemas for this property in all variants
        prop_schemas = []
        for variant in object_variants:
            for name, prop_def in variant.properties:
                if name == prop_name:
                    prop_schemas.append(prop_def.schema)
                    break

        # Check if all are literals with different values
        if len(prop_schemas) != len(object_variants):
            continue

        if not all(s.kind == "literal" for s in prop_schemas):
            continue

        # Check all values are different
        literal_schemas: list[LiteralNode] = [
            s for s in prop_schemas if s.kind == "literal"
        ]
        values = [s.value for s in literal_schemas]
        if len(set(map(str, values))) == len(values):  # Use str() for hashability
            # Found discriminator!
            return prop_name

    return None


def _intersection_base_type(schemas: tuple[SchemaNode, ...]) -> str:
    """Compute the narrowest representable Python type for an intersection.

    When all sub-schemas map to the same base Python type we use that instead of
    Any so pyright (and users) see meaningful type information even though a
    BeforeValidator still enforces the full intersection at runtime.
    """
    _KIND_BASE: dict[str, str] = {
        "string": "str",
        "boolean": "bool",
        "null": "None",
        "array": "list",
        "tuple": "tuple",
    }

    base: str | None = None
    for s in schemas:
        kind = s.kind
        if kind == "number":
            t = "int" if getattr(s, "integer", False) else "float"
        elif kind == "any":
            # any doesn't constrain — skip; it won't widen the intersection
            continue
        elif kind in _KIND_BASE:
            t = _KIND_BASE[kind]
        else:
            # object, union, oneOf, intersection, not, conditional, typeGuarded,
            # literal, enum, ref, nullable, never — can't reduce to a single bare type
            return "Any"

        if base is None:
            base = t
        elif base != t:
            return "Any"

    return base if base is not None else "Any"


def _union_base_type(type_exprs: list[str]) -> str:
    """Compute a union base type from rendered type expressions.

    Used by oneOf and conditional to replace Any with a narrower union when
    all branches produce representable types. If any branch uses Any, the
    whole union collapses to Any (Union[str, Any] == Any in Python's type system).
    """
    import re

    for expr in type_exprs:
        if re.search(r"\bAny\b", expr):
            return "Any"

    if len(type_exprs) == 1:
        return type_exprs[0]
    return " | ".join(type_exprs)


def render_intersection(node: IntersectionNode, name: str) -> RenderResult:
    """Render IntersectionNode (JSON Schema 'allOf') to Pydantic type.

    ALLOF SEMANTICS:
    A value must validate against ALL schemas in the allOf array.
    Each schema validates independently - they don't "see" each other.

    OBJECT MERGING STRATEGY:
    When all allOf schemas are objects, we merge their properties into one class.
    This is a static optimization that produces cleaner code than runtime checks.
    Property conflicts (same name in multiple schemas) are handled by creating
    an intersection of those property schemas.

    ADDITIONALPROPERTIES + ALLOF GOTCHA:
    JSON Schema says additionalProperties validates independently per schema.
    If schema A has {additionalProperties: {type: "boolean"}} and schema B
    defines property "foo", then "foo" is still an "additional property" to A
    and must be boolean. We handle this with a model_validator that applies
    each additionalProperties schema to ALL properties.

    MIXED TYPES:
    If allOf contains both objects and primitives (unusual but legal), we
    can't merge statically. Instead, we generate a validator that runs each
    schema's validation in sequence, failing if any rejects the value.
    """
    imports: set[str] = set()
    code_parts: list[str] = []

    if not node.schemas:
        imports.add("from typing import Any")
        return RenderResult(code="", type_expr="Any", imports=imports)

    # Separate objects from non-objects
    object_schemas: list[ObjectNode] = [s for s in node.schemas if s.kind == "object"]
    non_object_schemas = [s for s in node.schemas if s.kind != "object"]

    # If all are objects, merge them
    if len(object_schemas) == len(node.schemas):
        return _merge_object_schemas(object_schemas, name)

    # If mixed types, render each and combine
    # For primitives, use custom validator checking all schemas
    if non_object_schemas:
        imports.add("from typing import Annotated")
        imports.add("from pydantic import BeforeValidator, TypeAdapter")

        # Render each schema and collect results
        rendered: list[RenderResult] = []
        for i, schema in enumerate(node.schemas):
            schema_result = render(schema, f"{name}Part{i}")
            rendered.append(schema_result)
            imports.update(schema_result.imports)
            if schema_result.code:
                code_parts.append(schema_result.code)

        # Compute narrowest static base type from sub-schemas
        base_type = _intersection_base_type(node.schemas)
        if base_type == "Any":
            imports.add("from typing import Any")

        # Create validator that checks all schemas
        validator_name = f"_intersection_{name.lower()}"
        validator_lines = [f"def {validator_name}(v):"]

        for i, schema_result in enumerate(rendered):
            validator_lines.append(
                f"    validator_{i} = TypeAdapter({schema_result.type_expr})"
            )
            validator_lines.append(f"    try:")
            validator_lines.append(f"        validator_{i}.validate_python(v)")
            validator_lines.append(f"    except Exception as e:")
            validator_lines.append(
                f"        raise ValueError(f'Failed intersection schema {i}: {{e}}')"
            )

        validator_lines.append("    return v")
        code_parts.append("\n".join(validator_lines))

        type_expr = f"Annotated[{base_type}, BeforeValidator({validator_name})]"
        code = "\n\n\n".join(code_parts)
        return RenderResult(code=code, type_expr=type_expr, imports=imports)

    # Fallback: just use first schema
    return render(node.schemas[0], name)


def _merge_object_schemas(schemas: list[ObjectNode], name: str) -> RenderResult:
    """Merge multiple object schemas into a single Pydantic model.

    Combines all properties, using the most restrictive required setting.
    """
    imports: set[str] = {"from pydantic import BaseModel"}
    nested_classes: list[str] = []
    field_lines: list[str] = []

    # Collect all properties from all schemas
    all_properties: dict[str, list[tuple[PropertyDef, ObjectNode]]] = {}

    for schema in schemas:
        for prop_name, prop_def in schema.properties:
            if prop_name not in all_properties:
                all_properties[prop_name] = []
            all_properties[prop_name].append((prop_def, schema))

    # Render each unique property
    for prop_name, prop_defs in all_properties.items():
        # If property appears in multiple schemas, use intersection
        if len(prop_defs) == 1:
            prop_def, _ = prop_defs[0]
            nested_name = f"{name}{_to_pascal_case(prop_name)}"
            prop_result = render(prop_def.schema, nested_name)
            imports.update(prop_result.imports)

            if prop_result.code:
                nested_classes.append(prop_result.code)

            # Property is required if marked required in any schema
            is_required = prop_def.required
            is_nullable = _is_nullable_type(prop_def.schema)
            type_expr = prop_result.type_expr

            if is_required:
                if is_nullable:
                    type_expr = _make_nullable(type_expr)
                field_lines.append(f"    {prop_name}: {type_expr}")
            else:
                type_expr = _make_nullable(type_expr)
                field_lines.append(f"    {prop_name}: {type_expr} = None")
        else:
            # Property in multiple schemas - intersect the schemas
            prop_schemas = [pd.schema for pd, _ in prop_defs]
            intersection = IntersectionNode(schemas=tuple(prop_schemas))
            nested_name = f"{name}{_to_pascal_case(prop_name)}"
            prop_result = render(intersection, nested_name)
            imports.update(prop_result.imports)

            if prop_result.code:
                nested_classes.append(prop_result.code)

            # Required if any schema marks it required
            is_required = any(pd.required for pd, _ in prop_defs)
            type_expr = prop_result.type_expr

            if is_required:
                field_lines.append(f"    {prop_name}: {type_expr}")
            else:
                type_expr = _make_nullable(type_expr)
                field_lines.append(f"    {prop_name}: {type_expr} = None")

    # Check additionalProperties - use most restrictive
    forbid_extra = any(s.additional_properties is False for s in schemas)

    # Check if any schema has additionalProperties as a schema (not just true/false)
    # When merging with allOf, additionalProperties schema must validate ALL properties
    additional_props_schemas = [
        s.additional_properties
        for s in schemas
        if s.additional_properties is not None
        and not isinstance(s.additional_properties, bool)
    ]

    validators: list[str] = []

    # If we have additionalProperties schemas, create validator to apply to all properties
    # This handles the case where allOf + sibling additionalProperties need independent validation
    if additional_props_schemas:
        imports.add("from pydantic import model_validator, TypeAdapter")

        # Render each additionalProperties schema
        for i, add_prop_schema in enumerate(additional_props_schemas):
            add_prop_result = render(add_prop_schema, f"{name}AdditionalProps{i}")
            imports.update(add_prop_result.imports)
            if add_prop_result.code:
                nested_classes.append(add_prop_result.code)

            # Create validator that validates ALL properties against additionalProperties schema
            validator_lines = ["    @model_validator(mode='after')"]
            validator_lines.append(f"    def _validate_additional_props_{i}(self):")
            validator_lines.append(
                f"        # Validate all properties against additionalProperties schema from allOf merge"
            )
            validator_lines.append(
                f"        validator = TypeAdapter({add_prop_result.type_expr})"
            )
            validator_lines.append(
                "        for key, value in self.model_dump(exclude_unset=True).items():"
            )
            validator_lines.append("            try:")
            validator_lines.append("                validator.validate_python(value)")
            validator_lines.append("            except Exception as e:")
            validator_lines.append(
                f"                raise ValueError(f'Property {{key}} must match additionalProperties schema: {{e}}')"
            )
            validator_lines.append("        return self")
            validators.append("\n".join(validator_lines))

    class_lines: list[str] = [f"class {name}(BaseModel):"]

    if forbid_extra:
        imports.add("from pydantic import ConfigDict")
        class_lines.append("    model_config = ConfigDict(extra='forbid')")

    if field_lines:
        class_lines.extend(field_lines)
    elif not forbid_extra and not validators:
        class_lines.append("    pass")

    # Add validators
    if validators:
        class_lines.append("")  # Blank line before validators
        for validator in validators:
            class_lines.append(validator)

    class_code = "\n".join(class_lines)
    if nested_classes:
        code = "\n\n\n".join(nested_classes) + "\n\n\n" + class_code
    else:
        code = class_code

    return RenderResult(code=code, type_expr=name, imports=imports)


def render_oneof(node: OneOfNode, name: str) -> RenderResult:
    """Render OneOfNode (JSON Schema 'oneOf') to Pydantic type.

    ONEOF vs ANYOF - THE CRITICAL DIFFERENCE:
    - anyOf: value is valid if AT LEAST ONE schema matches
    - oneOf: value is valid if EXACTLY ONE schema matches

    Pydantic's Union can't express oneOf because it stops at first match.
    Consider oneOf: [{type: number}, {type: integer}]. The value 5 matches
    BOTH schemas (integers are numbers), so it should FAIL oneOf validation.
    But Union would happily accept it on first match.

    SOLUTION:
    Generate a BeforeValidator that:
    1. Tries each schema, counting matches
    2. If matches == 0: reject (no schema matched)
    3. If matches > 1: reject (ambiguous)
    4. If matches == 1: accept

    This requires trying ALL schemas even after finding a match, which is
    less efficient than Union but semantically correct.

    TYPE NARROWING:
    Instead of Annotated[Any, ...], we use Annotated[Union[T1, T2], ...] when
    all sub-schemas produce representable types. This gives pyright useful type
    info while the BeforeValidator still enforces exact-one-match at runtime.
    """
    imports: set[str] = {
        "from typing import Annotated",
        "from pydantic import BeforeValidator, TypeAdapter",
    }
    code_parts: list[str] = []
    rendered: list[RenderResult] = []

    # Render all schemas
    for i, schema in enumerate(node.schemas):
        schema_result = render(schema, f"{name}Option{i}")
        rendered.append(schema_result)
        imports.update(schema_result.imports)
        if schema_result.code:
            code_parts.append(schema_result.code)

    # Create validator that checks exactly one schema matches
    validator_name = f"_oneof_{name.lower()}"
    validator_lines = [f"def {validator_name}(v):"]
    validator_lines.append("    matches = []")

    for i, schema_result in enumerate(rendered):
        validator_lines.append(
            f"    validator_{i} = TypeAdapter({schema_result.type_expr})"
        )
        validator_lines.append("    try:")
        validator_lines.append(f"        validator_{i}.validate_python(v)")
        validator_lines.append(f"        matches.append({i})")
        validator_lines.append("    except Exception:")
        validator_lines.append("        pass")

    validator_lines.append("    if len(matches) == 0:")
    validator_lines.append(
        "        raise ValueError('Value must match at least one schema')"
    )
    validator_lines.append("    if len(matches) > 1:")
    validator_lines.append(
        f"        raise ValueError(f'Value must match exactly one schema, but matched {{len(matches)}} schemas')"
    )
    validator_lines.append("    return v")

    code_parts.append("\n".join(validator_lines))

    # narrow the static base type from sub-schema type expressions
    base_type = _union_base_type([r.type_expr for r in rendered])
    if base_type == "Any":
        imports.add("from typing import Any")

    type_expr = f"Annotated[{base_type}, BeforeValidator({validator_name})]"
    code = "\n\n\n".join(code_parts)
    return RenderResult(code=code, type_expr=type_expr, imports=imports)


def render_not(node: NotNode, name: str) -> RenderResult:
    """Render NotNode (not) to Pydantic type with negation validation.

    JSON Schema not requires the value to NOT match the schema.
    We use a custom validator that rejects values matching the schema.
    """
    imports: set[str] = {
        "from typing import Annotated, Any",
        "from pydantic import BeforeValidator, TypeAdapter",
    }
    code_parts: list[str] = []

    # Render the schema to negate
    schema_result = render(node.schema, f"{name}Not")
    imports.update(schema_result.imports)
    if schema_result.code:
        code_parts.append(schema_result.code)

    # Create validator that rejects matching values
    # Key insight: if validate_python succeeds, value matched → reject
    # If it raises ANY exception, value didn't match → accept
    # We use a flag to separate success from exception, avoiding the trap of
    # raising inside try-except (which would catch our own error on nested not)
    validator_name = f"_not_{name.lower()}"
    validator_lines = [f"def {validator_name}(v) -> Any:"]
    validator_lines.append(f"    validator = TypeAdapter({schema_result.type_expr})")
    validator_lines.append("    matched = False")
    validator_lines.append("    try:")
    validator_lines.append("        validator.validate_python(v)")
    validator_lines.append("        matched = True")
    validator_lines.append("    except Exception:")
    validator_lines.append("        pass")
    validator_lines.append("    if matched:")
    validator_lines.append(
        "        raise ValueError('Value must NOT match the schema')"
    )
    validator_lines.append("    return v")

    code_parts.append("\n".join(validator_lines))

    type_expr = f"Annotated[Any, BeforeValidator({validator_name})]"
    code = "\n\n\n".join(code_parts)
    return RenderResult(code=code, type_expr=type_expr, imports=imports)


def render_conditional(node: ConditionalNode, name: str) -> RenderResult:
    """Render ConditionalNode (if/then/else) to Pydantic type.

    JSON Schema if/then/else allows conditional validation:
    - If 'if' matches and 'then' exists, validate against 'then'
    - If 'if' doesn't match and 'else' exists, validate against 'else'
    - Otherwise, the original value is valid

    TYPE NARROWING:
    When both then and else exist, the value must be one of the two branch
    types, so we use Union[then_type, else_type] instead of Any. When only
    one branch exists, the other path passes any value through unchanged,
    so we fall back to Any.
    """
    imports: set[str] = {
        "from typing import Annotated",
        "from pydantic import BeforeValidator, TypeAdapter",
    }
    code_parts: list[str] = []

    # Render the if schema
    if_result = render(node.if_schema, f"{name}If")
    imports.update(if_result.imports)
    if if_result.code:
        code_parts.append(if_result.code)

    # Render then schema if present
    then_result = None
    if node.then_schema is not None:
        then_result = render(node.then_schema, f"{name}Then")
        imports.update(then_result.imports)
        if then_result.code:
            code_parts.append(then_result.code)

    # Render else schema if present
    else_result = None
    if node.else_schema is not None:
        else_result = render(node.else_schema, f"{name}Else")
        imports.update(else_result.imports)
        if else_result.code:
            code_parts.append(else_result.code)

    # Create validator with if/then/else logic
    validator_name = f"_conditional_{name.lower()}"
    validator_lines = [f"def {validator_name}(v):"]
    validator_lines.append(f"    # Check if 'if' schema matches")
    validator_lines.append(f"    if_validator = TypeAdapter({if_result.type_expr})")
    validator_lines.append("    try:")
    validator_lines.append("        if_validator.validate_python(v)")
    validator_lines.append("        if_matches = True")
    validator_lines.append("    except Exception:")
    validator_lines.append("        if_matches = False")

    # Handle then branch
    if then_result is not None:
        validator_lines.append(
            "    # If 'if' matches and 'then' exists, validate against 'then'"
        )
        validator_lines.append("    if if_matches:")
        validator_lines.append(
            f"        then_validator = TypeAdapter({then_result.type_expr})"
        )
        validator_lines.append("        try:")
        validator_lines.append("            return then_validator.validate_python(v)")
        validator_lines.append("        except Exception as e:")
        validator_lines.append(
            "            raise ValueError(f'Value must match then schema: {e}')"
        )

    # Handle else branch
    if else_result is not None:
        validator_lines.append(
            "    # If 'if' doesn't match and 'else' exists, validate against 'else'"
        )
        validator_lines.append("    if not if_matches:")
        validator_lines.append(
            f"        else_validator = TypeAdapter({else_result.type_expr})"
        )
        validator_lines.append("        try:")
        validator_lines.append("            return else_validator.validate_python(v)")
        validator_lines.append("        except Exception as e:")
        validator_lines.append(
            "            raise ValueError(f'Value must match else schema: {e}')"
        )

    # Default: value is valid
    validator_lines.append("    # Otherwise, value is valid as-is")
    validator_lines.append("    return v")

    code_parts.append("\n".join(validator_lines))

    # narrow: if both then and else exist, value must be one or the other
    if then_result is not None and else_result is not None:
        base_type = _union_base_type([then_result.type_expr, else_result.type_expr])
    else:
        # only one branch — the other path accepts anything
        base_type = "Any"

    if base_type == "Any":
        imports.add("from typing import Any")

    type_expr = f"Annotated[{base_type}, BeforeValidator({validator_name})]"
    code = "\n\n\n".join(code_parts)
    return RenderResult(code=code, type_expr=type_expr, imports=imports)


def render_type_guarded(node: TypeGuardedNode, name: str) -> RenderResult:
    """Render TypeGuardedNode to Pydantic type.

    TypeGuardedNode represents schemas with multiple type-specific keywords but no
    explicit type. Each guard checks the runtime type and applies the corresponding schema.

    Example: {minLength: 1, minimum: 0} has guards for both string and number types.

    We use a custom validator that applies the appropriate schema based on runtime type.
    """
    imports: set[str] = {
        "from typing import Annotated, Any",
        "from pydantic import BeforeValidator, TypeAdapter",
    }
    code_parts: list[str] = []

    # Render each guard's schema
    for i, guard in enumerate(node.guards):
        guard_result = render(guard.schema, f"{name}Guard{i}")
        imports.update(guard_result.imports)
        if guard_result.code:
            code_parts.append(guard_result.code)

    # Create validator that applies schema based on runtime type
    validator_name = f"_type_guarded_{name.lower()}"
    validator_lines = [f"def {validator_name}(v) -> Any:"]
    validator_lines.append("    # Apply schema based on runtime type")

    for i, guard in enumerate(node.guards):
        guard_result = render(guard.schema, f"{name}Guard{i}")

        # Map type guard check to Python type check
        check_code = _type_guard_check_to_python(guard.check)

        validator_lines.append(f"    if {check_code}:")
        validator_lines.append(
            f"        validator = TypeAdapter({guard_result.type_expr})"
        )
        validator_lines.append("        try:")
        validator_lines.append("            return validator.validate_python(v)")
        validator_lines.append("        except Exception as e:")
        validator_lines.append(
            f"            raise ValueError(f'Value must match {guard.check} schema: {{e}}')"
        )

    # If no guard matches, value is valid as-is
    validator_lines.append("    # No type guard matched, value is valid as-is")
    validator_lines.append("    return v")

    code_parts.append("\n".join(validator_lines))

    type_expr = f"Annotated[Any, BeforeValidator({validator_name})]"
    code = "\n\n\n".join(code_parts)
    return RenderResult(code=code, type_expr=type_expr, imports=imports)


def _type_guard_check_to_python(check: str) -> str:
    """Convert type guard check string to Python isinstance check.

    Maps IR type names to Python runtime checks.
    """
    if check == "string":
        return "isinstance(v, str)"
    elif check == "number":
        return "isinstance(v, (int, float)) and not isinstance(v, bool)"
    elif check == "integer":
        return "isinstance(v, int) and not isinstance(v, bool)"
    elif check == "boolean":
        return "isinstance(v, bool)"
    elif check == "null":
        return "v is None"
    elif check == "array":
        return "isinstance(v, (list, tuple))"
    elif check == "object":
        return "isinstance(v, dict)"
    else:
        # Unknown check - always false
        return "False"


def render_ref(node: RefNode, name: str) -> RenderResult:
    """Render RefNode to Pydantic type.

    Resolved refs are rendered inline. Recursive refs (resolved=None, path
    starts with '#') use a forward reference string to the root class so
    Pydantic validates the recursive structure correctly at runtime.
    Truly unresolvable refs (should never reach the renderer — the parser
    raises ValueError for those) produce a ConversionError.
    """
    if node.resolved is not None:
        return render(node.resolved, name)

    # Recursive ref (cycle detected by the parser). Use a forward reference
    # to the root class so Pydantic validates recursively at runtime.
    # Only works when the root schema generates a BaseModel class.
    if node.path.startswith("#"):
        global _needs_model_rebuild
        if _root_class_name and _root_is_class:
            _needs_model_rebuild = True
            return RenderResult(
                code="", type_expr=f"'{_root_class_name}'", imports=set()
            )
        # Root is not a class (TypeGuarded, union, etc.) — can't use forward refs
        return RenderResult(
            code="", type_expr="Any", imports={"from typing import Any"}
        )

    # Genuinely unresolved ref — the parser should have caught this.
    raise ConversionError(
        f"Unresolved $ref '{node.path}' — "
        "external refs must be bundled before reaching the adapter"
    )
