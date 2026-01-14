"""Converter module - orchestrates parse -> render pipeline."""

from typing import Any

from xschema_core import parse

from .errors import ConversionError, InvalidSchemaError
from .import_collector import ImportCollector
from .renderer import render


def _is_object_only_schema(ir_node: Any, original_schema: dict | bool) -> bool:
    """Check if schema only has object-specific keywords without explicit type.

    JSON Schema object keywords (properties, additionalProperties, etc.) should
    ignore non-object inputs - they pass validation automatically. This function
    detects schemas that need this "ignore non-objects" behavior.

    Args:
        ir_node: The parsed IR node
        original_schema: The original JSON Schema

    Returns True if:
    - The node is an ObjectNode (has object-specific constraints)
    - AND the schema doesn't have an explicit 'type' keyword
      (if it has type: object, Pydantic's rejection of non-objects is correct)
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

    JSON Schema number keywords (multipleOf, minimum, maximum, etc.) should
    ignore non-number inputs when there's no explicit type constraint.

    Args:
        ir_node: The parsed IR node
        original_schema: The original JSON Schema

    Returns True if:
    - The node is a NumberNode
    - AND the original schema doesn't have explicit 'type' keyword
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

    JSON Schema string keywords (minLength, maxLength, pattern, format) should
    ignore non-string inputs when there's no explicit type constraint.

    Args:
        ir_node: The parsed IR node
        original_schema: The original JSON Schema

    Returns True if:
    - The node is a StringNode
    - AND the original schema doesn't have explicit 'type' keyword
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

    JSON Schema array keywords (items, minItems, maxItems, etc.) should
    ignore non-array inputs when there's no explicit type constraint.

    Args:
        ir_node: The parsed IR node
        original_schema: The original JSON Schema

    Returns True if:
    - The node is an ArrayNode
    - AND the original schema doesn't have explicit 'type' keyword
    """
    if not hasattr(ir_node, "kind"):
        return False

    # Must be an ArrayNode
    if ir_node.kind != "array":
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
    """Convert a schema input to Pydantic code output.

    Args:
        input_data: Dict with namespace, id, varName, schema keys

    Returns:
        Dict with namespace, id, varName, imports, schema, type keys

    Raises:
        InvalidSchemaError: When required fields are missing or schema is malformed
        ConversionError: When schema cannot be converted to Pydantic code
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

    # Render the IR to Pydantic code
    try:
        result = render(ir_node, class_name)
    except Exception as e:
        raise ConversionError(
            f"Failed to render Pydantic code: {e}", schema_path=schema_path
        ) from e

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
        schema_code = result.code
        if type_guard_wrapper:
            # Wrap with isinstance check - non-matching types pass validation
            validate_expr = f"(lambda v: True if not {type_guard_wrapper} else _try_validate({class_name}.model_validate)(v))"
        else:
            # For simple class definitions, use model_validate directly
            validate_expr = f"_try_validate({class_name}.model_validate)"
    else:
        # For all other cases (primitives, unions, oneOf, etc.), use TypeAdapter
        collector.add("from pydantic import TypeAdapter")

        # If there's helper code, prepend it
        type_adapter_stmt = f"{var_name} = TypeAdapter({result.type_expr})"
        if result.code:
            schema_code = f"{result.code}\n\n{type_adapter_stmt}"
        else:
            schema_code = type_adapter_stmt

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
