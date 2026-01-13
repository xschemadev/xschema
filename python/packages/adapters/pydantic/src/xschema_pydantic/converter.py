"""Converter module - orchestrates parse -> render pipeline."""

from typing import Any

from xschema_core import parse


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
    """
    namespace = input_data.get("namespace", "")
    schema_id = input_data.get("id", "")
    var_name = input_data.get("varName", "")
    schema = input_data.get("schema", {})

    # Parse the JSON Schema to IR
    ir_node = parse(schema)

    # Generate class name from varName
    class_name = to_pascal_case(var_name)

    # TODO: Call renderer to generate Pydantic code (implemented in adapter-renderer-* tasks)
    # For now, generate a placeholder that shows the converter is working
    imports = ["from pydantic import BaseModel"]
    schema_code = f"class {class_name}(BaseModel):\n    pass  # TODO: implement rendering"
    type_expr = class_name

    return {
        "namespace": namespace,
        "id": schema_id,
        "varName": var_name,
        "imports": imports,
        "schema": schema_code,
        "type": type_expr,
    }
