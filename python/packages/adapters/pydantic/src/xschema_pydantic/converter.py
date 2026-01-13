"""Converter module - orchestrates parse -> render pipeline."""

from typing import Any

from xschema_core import parse

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
    """
    namespace = input_data.get("namespace", "")
    schema_id = input_data.get("id", "")
    var_name = input_data.get("varName", "")
    schema = input_data.get("schema", {})

    # Parse the JSON Schema to IR
    ir_node = parse(schema)

    # Generate class name from varName
    class_name = to_pascal_case(var_name)

    # Render the IR to Pydantic code
    result = render(ir_node, class_name)

    # Convert imports set to sorted list for consistent output
    imports = sorted(result.imports)

    # Determine if this is a class definition or a type expression
    # Class definitions (objects) have code that starts with "class "
    is_class_definition = result.code.startswith("class ")

    if is_class_definition:
        schema_code = result.code
    else:
        # Non-object types get TypeAdapter
        imports.append("from pydantic import TypeAdapter")
        imports = sorted(set(imports))

        # If there's helper code (like for Never type), prepend it
        type_adapter_stmt = f"{var_name} = TypeAdapter({result.type_expr})"
        if result.code:
            schema_code = f"{result.code}\n\n{type_adapter_stmt}"
        else:
            schema_code = type_adapter_stmt

    return {
        "namespace": namespace,
        "id": schema_id,
        "varName": var_name,
        "imports": imports,
        "schema": schema_code,
        "type": result.type_expr,
    }
