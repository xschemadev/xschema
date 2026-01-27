"""Converter module - orchestrates the JSON Schema parse -> render pipeline.

This module is the entry point for converting JSON Schema to Pydantic code.
It coordinates two phases:

1. **Parse Phase** (via xschema_core.parse):
   - Input: JSON Schema (dict or bool)
   - Output: IR (Intermediate Representation) node tree
   - IR nodes like ObjectNode, ArrayNode, UnionNode represent schema semantics
   - Parsing is adapter-agnostic - same IR feeds Zod, Pydantic, etc.

2. **Render Phase** (via .renderer.render):
   - Input: IR node tree
   - Output: Pydantic code as string + type expression + imports
   - Converts IR to actual Pydantic constructs (BaseModel, TypeAdapter, validators)

The converter also handles:
- **Type guard generation**: JSON Schema keywords are type-specific. A schema like
  `{"minLength": 5}` should pass non-strings (no type constraint = no rejection).
  We detect these "type-only" schemas and wrap validation in isinstance() checks.

- **Helper function registry**: Some Pydantic patterns need shared helpers (e.g.
  _json_equals for deep equality, _try_validate for exception wrapping). The
  renderer tracks which helpers are needed, and we prepend them to the output.

- **Import collection**: Pydantic code needs imports (BaseModel, Field, etc.).
  ImportCollector deduplicates and sorts imports for clean output.

Protocol with CLI:
- CLI sends JSON via stdin: [{namespace, id, varName, schema}, ...]
- Adapter returns JSON via stdout: [{namespace, id, varName, imports, schema, type, validate}, ...]
- cli.py handles batching; this module converts one schema at a time
"""

from typing import Any

from xschema_core import parse

from .errors import ConversionError, InvalidSchemaError
from .import_collector import ImportCollector
from .renderer import render, _reset_helper_registry, get_needed_helpers


def _is_object_only_schema(ir_node: Any, original_schema: dict | bool) -> bool:
    """Check if schema only has object-specific keywords without explicit type.

    JSON Schema type inference problem:
    ------------------------------------
    JSON Schema keywords are type-specific. A schema like:
        {"properties": {"name": {"type": "string"}}}

    Has NO type constraint! Per JSON Schema spec, this schema:
    - VALIDATES "hello" (string) -> true (no type constraint, properties ignored)
    - VALIDATES 42 (number) -> true (no type constraint, properties ignored)
    - VALIDATES {"name": 123} -> false (IS object, so properties applies, wrong type)

    But if we parse this as ObjectNode and render to Pydantic BaseModel,
    Pydantic will reject ALL non-dict inputs. That's wrong!

    Solution:
    ---------
    We detect "typeless object schemas" - schemas where:
    1. Parser inferred ObjectNode (has object-specific keywords)
    2. Original schema lacks explicit "type": "object"

    These get wrapped in isinstance(v, dict) guards at validation time,
    allowing non-objects to pass through unchanged.

    Args:
        ir_node: The parsed IR node
        original_schema: The original JSON Schema

    Returns True if schema needs object-only guard wrapper.
    """
    if not hasattr(ir_node, "kind"):
        return False

    # Must be an ObjectNode
    if ir_node.kind != "object":
        return False

    # Must not have explicit type in original schema
    if isinstance(original_schema, bool):
        return False

    return "type" not in original_schema


def _is_number_only_schema(ir_node: Any, original_schema: dict | bool) -> bool:
    """Check if schema has number-specific keywords without explicit type.

    Same principle as _is_object_only_schema but for number keywords.

    Example: {"minimum": 0} should:
    - VALIDATE "hello" -> true (not a number, minimum doesn't apply)
    - VALIDATE -5 -> false (IS a number, violates minimum)

    Python gotcha: bool is subclass of int! We must exclude bools from
    "number" check, otherwise True/False get number validation applied.
    """
    if not hasattr(ir_node, "kind"):
        return False

    # Must be a NumberNode (number or integer)
    if ir_node.kind != "number":
        return False

    # Must not have explicit type in original schema
    if isinstance(original_schema, bool):
        return False

    return "type" not in original_schema


def _is_string_only_schema(ir_node: Any, original_schema: dict | bool) -> bool:
    """Check if schema has string-specific keywords without explicit type.

    Same principle as _is_object_only_schema but for string keywords.

    Example: {"minLength": 3} should:
    - VALIDATE 12345 -> true (not a string, minLength doesn't apply)
    - VALIDATE "ab" -> false (IS a string, violates minLength)
    """
    if not hasattr(ir_node, "kind"):
        return False

    # Must be a StringNode
    if ir_node.kind != "string":
        return False

    # Must not have explicit type in original schema
    if isinstance(original_schema, bool):
        return False

    return "type" not in original_schema


def _is_array_only_schema(ir_node: Any, original_schema: dict | bool) -> bool:
    """Check if schema has array-specific keywords without explicit type.

    Same principle as _is_object_only_schema but for array keywords.

    Example: {"minItems": 1} should:
    - VALIDATE "hello" -> true (not an array, minItems doesn't apply)
    - VALIDATE [] -> false (IS an array, violates minItems)

    Also applies to tuple schemas (items as array) - they should also
    pass non-array values through unchanged.
    """
    if not hasattr(ir_node, "kind"):
        return False

    # Must be an ArrayNode or TupleNode
    # TupleNode is created when items is an array (positional validation)
    if ir_node.kind not in ("array", "tuple"):
        return False

    # Must not have explicit type in original schema
    if isinstance(original_schema, bool):
        return False

    return "type" not in original_schema


def to_pascal_case(name: str) -> str:
    """Convert a variable name to PascalCase for class names."""
    # Handle snake_case
    parts = name.split("_")
    # Capitalize each part
    return "".join(part.capitalize() for part in parts if part)


def convert(input_data: dict[str, Any]) -> dict[str, Any]:
    """Convert a JSON Schema to Pydantic code - main entry point.

    Pipeline:
    ---------
    1. Validate input (required fields, schema type)
    2. Parse schema to IR via xschema_core.parse()
    3. Render IR to Pydantic code via renderer.render()
    4. Wrap in TypeAdapter for runtime validation
    5. Generate type guard wrapper if needed (typeless schemas)
    6. Collect imports and helpers

    Input format (from CLI):
        {
            "namespace": "users",      # optional, for grouping
            "id": "User",              # schema identifier
            "varName": "user_schema",  # Python variable name
            "schema": {...}            # JSON Schema dict or bool
        }

    Output format (to CLI):
        {
            "namespace": "users",
            "id": "User",
            "varName": "user_schema",
            "imports": ["from pydantic import ..."],
            "schema": "class User(BaseModel): ...",  # executable Python code
            "type": "User",                          # type expression for annotations
            "validate": "_try_validate(...)"         # validation expression
        }

    TypeAdapter usage:
    ------------------
    All schemas get wrapped in TypeAdapter, even simple classes. This ensures:
    - Consistent validation API (validate_python method)
    - Proper error handling via _try_validate wrapper
    - Works for both BaseModel classes and type expressions (Union, Annotated, etc.)

    Raises:
        InvalidSchemaError: Missing required fields or malformed schema
        ConversionError: Schema parsed but failed to render to Pydantic
    """
    # Validate required fields
    namespace = input_data.get("namespace", "")
    schema_id = input_data.get("id", "")
    var_name = input_data.get("varName", "")

    if not schema_id:
        raise InvalidSchemaError("Missing required field 'id'")
    if not var_name:
        raise InvalidSchemaError("Missing required field 'varName'")

    # Build schema path for error messages (do this before schema validation)
    schema_path = f"$.{namespace}.{schema_id}" if namespace else f"$.{schema_id}"

    schema = input_data.get("schema")
    if schema is None:
        raise InvalidSchemaError(
            f"Missing required field 'schema' for {schema_id}", schema_path=schema_path
        )

    # Validate schema is dict or bool (JSON Schema allows true/false)
    if not isinstance(schema, (dict, bool)):
        raise InvalidSchemaError(
            f"Schema must be a dict or bool, got {type(schema).__name__} for {schema_id}",
            schema_path=schema_path,
        )

    # Parse the JSON Schema to IR
    try:
        ir_node = parse(schema)
    except ValueError as e:
        raise InvalidSchemaError(str(e), schema_path=schema_path) from e
    except Exception as e:
        raise ConversionError(
            f"Failed to parse schema: {e}", schema_path=schema_path
        ) from e

    # Generate class name from varName
    class_name = to_pascal_case(var_name)

    # Reset helper registry before rendering
    _reset_helper_registry()

    # Render the IR to Pydantic code
    try:
        result = render(ir_node, class_name)
    except Exception as e:
        raise ConversionError(
            f"Failed to render Pydantic code: {e}", schema_path=schema_path
        ) from e

    # Get needed helper functions
    helpers = get_needed_helpers()

    # Collect and deduplicate imports using ImportCollector
    collector = ImportCollector()
    for import_stmt in result.imports:
        collector.add(import_stmt)

    # Determine if this is a simple class definition (just a BaseModel subclass)
    # A simple class definition:
    # 1. Code starts with "class "
    # 2. The class name matches the expected class_name (not a helper like "ClassNameOption0")
    # 3. The type_expr is just the class name (not an Annotated type or union)
    is_simple_class = (
        result.code.startswith(f"class {class_name}(")
        and result.type_expr == class_name
    )

    # Check if this is a type-specific schema that should ignore non-matching types
    # JSON Schema semantics: type-specific keywords only apply to their respective types
    is_object_only = _is_object_only_schema(ir_node, schema)
    is_number_only = _is_number_only_schema(ir_node, schema)
    is_string_only = _is_string_only_schema(ir_node, schema)
    is_array_only = _is_array_only_schema(ir_node, schema)

    # Determine the type guard wrapper if needed
    # These expressions check if a value IS the expected type (for validation)
    type_guard_wrapper = None
    if is_object_only:
        type_guard_wrapper = "isinstance(v, dict)"
    elif is_number_only:
        # Numbers: int or float, but not bool (bool is subclass of int in Python)
        # Wrap in parens so negation works correctly
        type_guard_wrapper = "(isinstance(v, (int, float)) and not isinstance(v, bool))"
    elif is_string_only:
        type_guard_wrapper = "isinstance(v, str)"
    elif is_array_only:
        type_guard_wrapper = "isinstance(v, list)"

    if is_simple_class:
        # For simple classes, create TypeAdapter for consistency
        # The template expects all schemas to have TypeAdapter assignments
        collector.add("from pydantic import TypeAdapter")

        # Build schema code with helpers, class, and type adapter
        type_adapter_stmt = f"{var_name} = TypeAdapter({class_name})"
        parts = []
        if helpers:
            parts.append(helpers)
        parts.append(result.code)
        parts.append(type_adapter_stmt)
        schema_code = "\n\n".join(parts)

        if type_guard_wrapper:
            # Wrap with isinstance check - non-matching types pass validation
            validate_expr = f"(lambda v: True if not {type_guard_wrapper} else _try_validate({var_name}.validate_python)(v))"
        else:
            # For TypeAdapter, use validate_python
            validate_expr = f"_try_validate({var_name}.validate_python)"
    else:
        # For all other cases (primitives, unions, oneOf, etc.), use TypeAdapter
        collector.add("from pydantic import TypeAdapter")

        # Build schema code with helpers and type adapter
        type_adapter_stmt = f"{var_name} = TypeAdapter({result.type_expr})"
        parts = []
        if helpers:
            parts.append(helpers)
        if result.code:
            parts.append(result.code)
        parts.append(type_adapter_stmt)
        schema_code = "\n\n".join(parts)

        if type_guard_wrapper:
            # Wrap with isinstance check - non-matching types pass validation
            validate_expr = f"(lambda v: True if not {type_guard_wrapper} else _try_validate({var_name}.validate_python)(v))"
        else:
            # For TypeAdapter, use validate_python
            validate_expr = f"_try_validate({var_name}.validate_python)"

    # Convert imports to sorted, deduplicated list
    imports = collector.to_list()

    return {
        "namespace": namespace,
        "id": schema_id,
        "varName": var_name,
        "imports": imports,
        "schema": schema_code,
        "type": result.type_expr,
        "validate": validate_expr,
    }
