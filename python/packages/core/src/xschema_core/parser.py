"""JSON Schema to IR parser."""

from typing import Any

from xschema_core.ir import (
    AnyNode,
    ArrayNode,
    BooleanNode,
    EnumNode,
    IntersectionNode,
    LiteralNode,
    NeverNode,
    NullNode,
    NumberNode,
    ObjectNode,
    PropertyDef,
    RefNode,
    SchemaNode,
    StringNode,
    TupleNode,
    UnionNode,
)


def parse(schema: dict | bool) -> SchemaNode:
    """Parse a JSON Schema dict into an IR SchemaNode."""
    # Handle boolean schemas
    if schema is True:
        return AnyNode()
    if schema is False:
        return NeverNode()

    # Handle empty schema
    if not schema:
        return AnyNode()

    description = schema.get("description")

    # Handle $ref - refs are pre-resolved by CLI, store path and resolved schema
    if "$ref" in schema:
        ref_path = schema["$ref"]
        # If the schema has been resolved inline by the CLI, it would be in a
        # different field. For now, we just store the path.
        return RefNode(path=ref_path, description=description)

    # Handle const (literal value)
    if "const" in schema:
        return LiteralNode(value=schema["const"], description=description)

    # Handle enum
    if "enum" in schema:
        return EnumNode(values=tuple(schema["enum"]), description=description)

    # Handle composition keywords
    if "allOf" in schema:
        schemas = tuple(parse(s) for s in schema["allOf"])
        return IntersectionNode(schemas=schemas, description=description)

    if "anyOf" in schema:
        variants = tuple(parse(s) for s in schema["anyOf"])
        return UnionNode(variants=variants, description=description)

    if "oneOf" in schema:
        variants = tuple(parse(s) for s in schema["oneOf"])
        # Detect discriminator: check if all variants have same property with Literal values
        discriminator = _detect_discriminator(schema["oneOf"])
        return UnionNode(
            variants=variants, discriminator=discriminator, description=description
        )

    # Get type - can be string or array
    schema_type = schema.get("type")

    # Handle array of types (union)
    if isinstance(schema_type, list):
        variants = []
        for t in schema_type:
            variant_schema = {**schema, "type": t}
            variants.append(parse(variant_schema))
        return UnionNode(variants=tuple(variants), description=description)

    # Handle prefixItems (tuple)
    if "prefixItems" in schema:
        return _parse_tuple(schema, description)

    # Infer type from keywords if not specified
    if schema_type is None:
        schema_type = _infer_type(schema)

    # Parse based on type
    if schema_type == "string":
        return _parse_string(schema, description)

    if schema_type == "number":
        return _parse_number(schema, integer=False, description=description)

    if schema_type == "integer":
        return _parse_number(schema, integer=True, description=description)

    if schema_type == "boolean":
        return BooleanNode(description=description)

    if schema_type == "null":
        return NullNode(description=description)

    if schema_type == "object":
        return _parse_object(schema, description)

    if schema_type == "array":
        return _parse_array(schema, description)

    # Unknown type - return AnyNode
    return AnyNode(description=description)


def _parse_string(schema: dict, description: str | None) -> StringNode:
    """Parse a string schema."""
    return StringNode(
        format=schema.get("format"),
        min_length=schema.get("minLength"),
        max_length=schema.get("maxLength"),
        pattern=schema.get("pattern"),
        description=description,
    )


def _parse_number(
    schema: dict, integer: bool, description: str | None
) -> NumberNode:
    """Parse a number/integer schema."""
    # Handle exclusiveMinimum/exclusiveMaximum which can be boolean (draft-4) or number (draft-6+)
    exclusive_min = schema.get("exclusiveMinimum")
    exclusive_max = schema.get("exclusiveMaximum")

    # Draft-4 style: exclusiveMinimum is boolean, actual value is in minimum
    if exclusive_min is True:
        exclusive_min = schema.get("minimum")
    elif exclusive_min is False:
        exclusive_min = None

    if exclusive_max is True:
        exclusive_max = schema.get("maximum")
    elif exclusive_max is False:
        exclusive_max = None

    return NumberNode(
        integer=integer,
        minimum=schema.get("minimum") if not isinstance(schema.get("exclusiveMinimum"), bool) or not schema.get("exclusiveMinimum") else None,
        maximum=schema.get("maximum") if not isinstance(schema.get("exclusiveMaximum"), bool) or not schema.get("exclusiveMaximum") else None,
        exclusive_minimum=exclusive_min if not isinstance(exclusive_min, bool) else None,
        exclusive_maximum=exclusive_max if not isinstance(exclusive_max, bool) else None,
        multiple_of=schema.get("multipleOf"),
        description=description,
    )


def _parse_object(schema: dict, description: str | None) -> ObjectNode:
    """Parse an object schema."""
    properties: list[tuple[str, PropertyDef]] = []
    required_set = frozenset(schema.get("required", []))

    for prop_name, prop_schema in schema.get("properties", {}).items():
        prop_node = parse(prop_schema) if isinstance(prop_schema, dict) else (
            AnyNode() if prop_schema is True else NeverNode()
        )
        prop_description = prop_schema.get("description") if isinstance(prop_schema, dict) else None
        properties.append((
            prop_name,
            PropertyDef(
                schema=prop_node,
                required=prop_name in required_set,
                description=prop_description,
            ),
        ))

    # Handle additionalProperties
    additional = schema.get("additionalProperties")
    if additional is None:
        additional_node: SchemaNode | bool | None = None
    elif isinstance(additional, bool):
        additional_node = additional
    else:
        additional_node = parse(additional)

    return ObjectNode(
        properties=tuple(properties),
        additional_properties=additional_node,
        required=required_set,
        description=description,
    )


def _parse_array(schema: dict, description: str | None) -> ArrayNode | TupleNode:
    """Parse an array schema."""
    items = schema.get("items")

    if items is None:
        items_node = AnyNode()
    elif isinstance(items, bool):
        items_node = AnyNode() if items else NeverNode()
    elif isinstance(items, list):
        # Legacy tuple syntax (items as array) - convert to tuple
        return _parse_legacy_tuple(schema, description)
    else:
        items_node = parse(items)

    return ArrayNode(
        items=items_node,
        min_items=schema.get("minItems"),
        max_items=schema.get("maxItems"),
        unique_items=schema.get("uniqueItems", False),
        description=description,
    )


def _parse_tuple(schema: dict, description: str | None) -> TupleNode:
    """Parse a tuple schema (using prefixItems)."""
    prefix_items = tuple(parse(item) for item in schema.get("prefixItems", []))

    # Rest items can be in 'items' when prefixItems is present
    rest = schema.get("items")
    if rest is None:
        rest_node = None
    elif isinstance(rest, bool):
        rest_node = AnyNode() if rest else None
    else:
        rest_node = parse(rest)

    return TupleNode(
        prefix_items=prefix_items,
        rest_items=rest_node,
        description=description,
    )


def _parse_legacy_tuple(schema: dict, description: str | None) -> TupleNode:
    """Parse legacy tuple syntax where items is an array."""
    items_list = schema.get("items", [])
    prefix_items = tuple(
        parse(item) if isinstance(item, dict) else (AnyNode() if item else NeverNode())
        for item in items_list
    )

    # additionalItems is the rest type for legacy tuples
    additional = schema.get("additionalItems")
    if additional is None:
        rest_node = None
    elif isinstance(additional, bool):
        rest_node = AnyNode() if additional else None
    else:
        rest_node = parse(additional)

    return TupleNode(
        prefix_items=prefix_items,
        rest_items=rest_node,
        description=description,
    )


def _infer_type(schema: dict) -> str | None:
    """Infer the type from other keywords in the schema."""
    # String keywords
    if any(k in schema for k in ("minLength", "maxLength", "pattern", "format")):
        return "string"

    # Number keywords
    if any(
        k in schema
        for k in ("minimum", "maximum", "exclusiveMinimum", "exclusiveMaximum", "multipleOf")
    ):
        return "number"

    # Object keywords
    if any(
        k in schema
        for k in ("properties", "required", "additionalProperties", "patternProperties")
    ):
        return "object"

    # Array keywords
    if any(k in schema for k in ("items", "minItems", "maxItems", "uniqueItems")):
        return "array"

    return None


def _detect_discriminator(variants: list[Any]) -> str | None:
    """Detect if oneOf variants form a discriminated union.

    A discriminated union has all variants as objects with a common property
    that has a literal (const) value unique to each variant.
    """
    if not variants:
        return None

    # Find common properties with const values across all variants
    common_discriminators: set[str] | None = None

    for variant in variants:
        if not isinstance(variant, dict):
            return None
        if variant.get("type") != "object" and "properties" not in variant:
            # Could still be an object without explicit type
            if "properties" not in variant:
                return None

        properties = variant.get("properties", {})
        discriminator_props = set()

        for prop_name, prop_schema in properties.items():
            if isinstance(prop_schema, dict) and "const" in prop_schema:
                discriminator_props.add(prop_name)

        if common_discriminators is None:
            common_discriminators = discriminator_props
        else:
            common_discriminators &= discriminator_props

        if not common_discriminators:
            return None

    # Return the first common discriminator property
    if common_discriminators:
        return next(iter(common_discriminators))
    return None
