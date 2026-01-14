"""Converter module - orchestrates parse -> render pipeline."""

from typing import Any

from xschema_core import parse

from .errors import ConversionError, InvalidSchemaError
from .import_collector import ImportCollector
from .renderer import render


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

    if is_simple_class:
        schema_code = result.code
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
