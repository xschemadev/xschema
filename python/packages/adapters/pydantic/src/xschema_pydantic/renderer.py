"""Pydantic Renderer - Converts SchemaNode IR to Pydantic code strings."""

from dataclasses import dataclass, field
from typing import Any

from xschema_core import (
    SchemaNode,
    StringNode,
    NumberNode,
    BooleanNode,
    NullNode,
    LiteralNode,
    EnumNode,
    AnyNode,
    NeverNode,
    ArrayNode,
    TupleNode,
    ObjectNode,
    UnionNode,
    OneOfNode,
    IntersectionNode,
    NotNode,
    ConditionalNode,
    TypeGuardedNode,
    NullableNode,
    RefNode,
)


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
            # Placeholder: treat like union for now
            return render_union(UnionNode(variants=node.schemas), name)
        case "intersection":
            return render_intersection(node, name)
        case "not":
            # Placeholder: will implement in later task
            return RenderResult(
                code="", type_expr="Any", imports={"from typing import Any"}
            )
        case "conditional":
            # Placeholder: will implement in later task
            return RenderResult(
                code="", type_expr="Any", imports={"from typing import Any"}
            )
        case "typeGuarded":
            # Placeholder: will implement in later task
            return RenderResult(
                code="", type_expr="Any", imports={"from typing import Any"}
            )
        case "nullable":
            # Handle nullable by wrapping inner type with | None
            inner_result = render(node.inner, name)
            if inner_result.type_expr.endswith(" | None"):
                return inner_result
            return RenderResult(
                code=inner_result.code,
                type_expr=f"{inner_result.type_expr} | None",
                imports=inner_result.imports,
            )
        case "ref":
            return render_ref(node, name)
        case _:
            # Fallback for unknown node types
            return RenderResult(
                code="", type_expr="Any", imports={"from typing import Any"}
            )


def render_string(node: StringNode) -> RenderResult:
    """Render StringNode to Pydantic type with format support."""
    imports: set[str] = set()

    # Handle format first - format types override base str type
    if node.format is not None:
        format_result = _render_format(node.format, node.constraints)
        if format_result is not None:
            return format_result
        # Unknown formats fall through to str with constraints

    # No format or unknown format - use str with constraints
    has_constraints = node.constraints is not None and (
        node.constraints.min_length is not None
        or node.constraints.max_length is not None
        or node.constraints.pattern is not None
    )

    if has_constraints:
        # Use Annotated with StringConstraints for constrained strings
        imports.add("from typing import Annotated")
        imports.add("from pydantic import StringConstraints")

        constraints: list[str] = []
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
        type_expr = "str"

    return RenderResult(code="", type_expr=type_expr, imports=imports)


def _render_format(fmt: str, constraints) -> RenderResult | None:
    """Render format string types to Pydantic types.

    Returns None for unknown formats (caller should use str).
    """
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


def render_number(node: NumberNode) -> RenderResult:
    """Render NumberNode to Pydantic type."""
    imports: set[str] = set()
    base_type = "int" if node.integer else "float"

    has_constraints = node.constraints is not None and (
        node.constraints.minimum is not None
        or node.constraints.maximum is not None
        or node.constraints.exclusive_minimum is not None
        or node.constraints.exclusive_maximum is not None
        or node.constraints.multiple_of is not None
    )

    if has_constraints:
        # Use Annotated with annotated-types constraints
        imports.add("from typing import Annotated")
        annotations: list[str] = []

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
        if node.constraints.multiple_of is not None:
            imports.add("from annotated_types import MultipleOf")
            annotations.append(f"MultipleOf({node.constraints.multiple_of})")

        annotation_str = ", ".join(annotations)
        type_expr = f"Annotated[{base_type}, {annotation_str}]"
    else:
        type_expr = base_type

    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_boolean(node: BooleanNode) -> RenderResult:
    """Render BooleanNode to Pydantic type."""
    return RenderResult(code="", type_expr="bool", imports=set())


def render_null(node: NullNode) -> RenderResult:
    """Render NullNode to Pydantic type."""
    return RenderResult(code="", type_expr="None", imports=set())


def render_literal(node: LiteralNode) -> RenderResult:
    """Render LiteralNode to Pydantic type."""
    imports: set[str] = {"from typing import Literal"}

    # Format the literal value appropriately
    value = node.value
    if value is None:
        literal_repr = "None"
    elif isinstance(value, bool):
        literal_repr = "True" if value else "False"
    elif isinstance(value, str):
        # Escape quotes in string
        escaped = value.replace("\\", "\\\\").replace("'", "\\'")
        literal_repr = f"'{escaped}'"
    else:
        literal_repr = repr(value)

    type_expr = f"Literal[{literal_repr}]"
    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_enum(node: EnumNode) -> RenderResult:
    """Render EnumNode to Pydantic type using Literal."""
    imports: set[str] = {"from typing import Literal"}

    # Build Literal type with all enum values
    literal_values: list[str] = []
    for value in node.values:
        if value is None:
            literal_values.append("None")
        elif isinstance(value, bool):
            literal_values.append("True" if value else "False")
        elif isinstance(value, str):
            escaped = value.replace("\\", "\\\\").replace("'", "\\'")
            literal_values.append(f"'{escaped}'")
        else:
            literal_values.append(repr(value))

    values_str = ", ".join(literal_values)
    type_expr = f"Literal[{values_str}]"
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
            validator_name = f"_unique_{name.lower()}"
            validator_code = f"""def {validator_name}(v: list) -> list:
    if len(v) != len(set(map(lambda x: x if isinstance(x, (str, int, float, bool)) else id(x), v))):
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

            # Add helper for validation checking
            if "_validates_against" not in [
                part for part in code_parts if "_validates_against" in part
            ]:
                helper_code = """def _validates_against(validator, value):
    try:
        validator.validate_python(value)
        return True
    except Exception:
        return False"""
                code_parts.insert(0, helper_code)

        type_expr = f"Annotated[{base_type}, {', '.join(annotations)}]"
    else:
        type_expr = base_type

    # Handle unevaluatedItems (simple case: treat like items for now)
    # In full JSON Schema, this is complex; we simplify since CLI filters hard cases

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

        # For variable-length tuples with heterogeneous types, use tuple[Any, ...]
        # and validate each position in custom validator
        imports.add("from typing import Annotated, Any")
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

        # Type as tuple[Any, ...] with validator doing the actual type checking
        type_expr = f"Annotated[tuple[Any, ...], BeforeValidator({validator_name})]"
    elif node.rest_items is False:
        # rest_items is False - no additional items allowed (strict tuple)
        if item_types:
            type_expr = f"tuple[{', '.join(item_types)}]"
        else:
            type_expr = "tuple[()]"  # Empty tuple
    else:
        # rest_items is None - default behavior (fixed-size tuple)
        if item_types:
            type_expr = f"tuple[{', '.join(item_types)}]"
        else:
            type_expr = "tuple[()]"

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
            validator_lines.append("    if len(v) != len(set(v)):")
            validator_lines.append(
                "        raise ValueError('Tuple items must be unique')"
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

            # Add helper if not already present
            if not any("_validates_against" in part for part in code_parts):
                helper_code = """def _validates_against(validator, value):
    try:
        validator.validate_python(value)
        return True
    except Exception:
        return False"""
                code_parts.insert(0, helper_code)

        validator_lines.append("    return v")
        code_parts.append("\n".join(validator_lines))

        type_expr = (
            f"Annotated[{type_expr}, AfterValidator({constraint_validator_name})]"
        )

    code = "\n\n\n".join(code_parts) if code_parts else ""
    return RenderResult(code=code, type_expr=type_expr, imports=imports)


def render_object(node: ObjectNode, name: str) -> RenderResult:
    """Render ObjectNode to Pydantic BaseModel class."""
    imports: set[str] = {"from pydantic import BaseModel"}
    nested_classes: list[str] = []
    field_lines: list[str] = []

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

        if is_required:
            if is_nullable:
                # Required but nullable: field: Type | None (no default)
                if not type_expr.endswith(" | None"):
                    type_expr = f"{type_expr} | None"
                field_lines.append(f"    {prop_name}: {type_expr}")
            else:
                # Required and not nullable: field: Type
                field_lines.append(f"    {prop_name}: {type_expr}")
        else:
            # Optional: field: Type | None = None
            if not type_expr.endswith(" | None") and type_expr != "None":
                type_expr = f"{type_expr} | None"
            field_lines.append(f"    {prop_name}: {type_expr} = None")

    # Handle additionalProperties
    config_line = None
    if node.additional_properties is False:
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='forbid')"
    elif node.additional_properties is True:
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='allow')"
    # If additional_properties is a schema, we allow extra but don't constrain type
    # (Pydantic doesn't support typed extra fields directly in model_config)
    elif (
        node.additional_properties is not None
        and node.additional_properties is not False
    ):
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='allow')"

    # Build the class definition
    class_lines: list[str] = [f"class {name}(BaseModel):"]

    # Add config if present
    if config_line:
        class_lines.append(config_line)

    # Add fields or pass if empty
    if field_lines:
        class_lines.extend(field_lines)
    elif not config_line:
        class_lines.append("    pass")

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
    """Render UnionNode to Pydantic type (placeholder)."""
    # Will be fully implemented in adapter-renderer-unions task
    imports: set[str] = set()
    variant_types: list[str] = []

    for i, variant in enumerate(node.variants):
        variant_result = render(variant, f"{name}Variant{i}")
        imports.update(variant_result.imports)
        variant_types.append(variant_result.type_expr)

    type_expr = " | ".join(variant_types)
    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_intersection(node: IntersectionNode, name: str) -> RenderResult:
    """Render IntersectionNode to Pydantic type (placeholder)."""
    # Will be fully implemented in adapter-renderer-unions task
    # For now, just render the first schema
    if node.schemas:
        return render(node.schemas[0], name)
    imports: set[str] = {"from typing import Any"}
    return RenderResult(code="", type_expr="Any", imports=imports)


def render_ref(node: RefNode, name: str) -> RenderResult:
    """Render RefNode to Pydantic type."""
    # If resolved, render the resolved schema
    if node.resolved is not None:
        return render(node.resolved, name)

    # Otherwise, just use the path as a forward reference
    # Extract the type name from the path (last segment after /)
    ref_name = node.path.split("/")[-1]
    return RenderResult(code="", type_expr=f"'{ref_name}'", imports=set())
