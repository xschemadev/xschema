"""JSON Schema to IR parser."""

from typing import Any

from xschema_core.ir import (
    AnyNode,
    ArrayConstraints,
    ArrayNode,
    BooleanNode,
    EnumNode,
    IntersectionNode,
    LiteralNode,
    NeverNode,
    NullNode,
    NumberConstraints,
    NumberNode,
    ObjectNode,
    OneOfNode,
    PropertyDef,
    RefNode,
    SchemaNode,
    StringConstraints,
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

    # Handle $ref - CLI should bundle everything, throw if encountered
    if "$ref" in schema:
        raise ValueError(
            "Unexpected $ref in schema. CLI should pre-bundle all schemas."
        )

    # Handle const (literal value)
    if "const" in schema:
        return LiteralNode(value=schema["const"])

    # Handle enum
    if "enum" in schema:
        return EnumNode(values=tuple(schema["enum"]))

    # Handle composition keywords
    if "allOf" in schema:
        schemas = tuple(parse(s) for s in schema["allOf"])
        return IntersectionNode(schemas=schemas)

    if "anyOf" in schema:
        variants = tuple(parse(s) for s in schema["anyOf"])
        return UnionNode(variants=variants)

    if "oneOf" in schema:
        schemas_tuple = tuple(parse(s) for s in schema["oneOf"])
        return OneOfNode(schemas=schemas_tuple)

    # Get type - can be string or array
    schema_type = schema.get("type")

    # Handle array of types (union)
    if isinstance(schema_type, list):
        variants = []
        for t in schema_type:
            variant_schema = {**schema, "type": t}
            variants.append(parse(variant_schema))
        return UnionNode(variants=tuple(variants))

    # Handle prefixItems (tuple)
    if "prefixItems" in schema:
        return _parse_tuple(schema)

    # Infer type from keywords if not specified
    if schema_type is None:
        schema_type = _infer_type(schema)

    # Parse based on type
    if schema_type == "string":
        return _parse_string(schema)

    if schema_type == "number":
        return _parse_number(schema, integer=False)

    if schema_type == "integer":
        return _parse_number(schema, integer=True)

    if schema_type == "boolean":
        return BooleanNode()

    if schema_type == "null":
        return NullNode()

    if schema_type == "object":
        return _parse_object(schema)

    if schema_type == "array":
        return _parse_array(schema)

    # Unknown type - return AnyNode
    return AnyNode()


def _parse_string(schema: dict) -> StringNode:
    """Parse a string schema."""
    constraints = StringConstraints(
        min_length=schema.get("minLength"),
        max_length=schema.get("maxLength"),
        pattern=schema.get("pattern"),
    )
    return StringNode(
        constraints=constraints,
        format=schema.get("format"),
    )


def _parse_number(schema: dict, integer: bool) -> NumberNode:
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

    constraints = NumberConstraints(
        minimum=schema.get("minimum")
        if not isinstance(schema.get("exclusiveMinimum"), bool)
        or not schema.get("exclusiveMinimum")
        else None,
        maximum=schema.get("maximum")
        if not isinstance(schema.get("exclusiveMaximum"), bool)
        or not schema.get("exclusiveMaximum")
        else None,
        exclusive_minimum=exclusive_min
        if not isinstance(exclusive_min, bool)
        else None,
        exclusive_maximum=exclusive_max
        if not isinstance(exclusive_max, bool)
        else None,
        multiple_of=schema.get("multipleOf"),
    )

    return NumberNode(
        constraints=constraints,
        integer=integer,
    )


def _parse_object(schema: dict) -> ObjectNode:
    """Parse an object schema."""
    properties: list[tuple[str, PropertyDef]] = []
    required_set = frozenset(schema.get("required", []))

    for prop_name, prop_schema in schema.get("properties", {}).items():
        prop_node = (
            parse(prop_schema)
            if isinstance(prop_schema, dict)
            else (AnyNode() if prop_schema is True else NeverNode())
        )
        properties.append(
            (
                prop_name,
                PropertyDef(
                    schema=prop_node,
                    required=prop_name in required_set,
                ),
            )
        )

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
    )


def _parse_array(schema: dict) -> ArrayNode | TupleNode:
    """Parse an array schema."""
    items = schema.get("items")

    if items is None:
        items_node = AnyNode()
    elif isinstance(items, bool):
        items_node = AnyNode() if items else NeverNode()
    elif isinstance(items, list):
        # Legacy tuple syntax (items as array) - convert to tuple
        return _parse_legacy_tuple(schema)
    else:
        items_node = parse(items)

    constraints = ArrayConstraints(
        min_items=schema.get("minItems"),
        max_items=schema.get("maxItems"),
        unique_items=schema.get("uniqueItems", False),
    )

    return ArrayNode(
        items=items_node,
        constraints=constraints,
    )


def _parse_tuple(schema: dict) -> TupleNode:
    """Parse a tuple schema (using prefixItems)."""
    prefix_items = tuple(parse(item) for item in schema.get("prefixItems", []))

    # Rest items can be in 'items' when prefixItems is present
    rest = schema.get("items")
    if rest is None:
        rest_node = None
    elif rest is False:
        rest_node = False
    elif isinstance(rest, bool):
        rest_node = AnyNode() if rest else False
    else:
        rest_node = parse(rest)

    constraints = ArrayConstraints(
        min_items=schema.get("minItems"),
        max_items=schema.get("maxItems"),
        unique_items=schema.get("uniqueItems", False),
    )

    return TupleNode(
        prefix_items=prefix_items,
        rest_items=rest_node,
        constraints=constraints,
    )


def _parse_legacy_tuple(schema: dict) -> TupleNode:
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
    elif additional is False:
        rest_node = False
    elif isinstance(additional, bool):
        rest_node = AnyNode() if additional else False
    else:
        rest_node = parse(additional)

    constraints = ArrayConstraints(
        min_items=schema.get("minItems"),
        max_items=schema.get("maxItems"),
        unique_items=schema.get("uniqueItems", False),
    )

    return TupleNode(
        prefix_items=prefix_items,
        rest_items=rest_node,
        constraints=constraints,
    )


def _infer_type(schema: dict) -> str | None:
    """Infer the type from other keywords in the schema."""
    # String keywords
    if any(k in schema for k in ("minLength", "maxLength", "pattern", "format")):
        return "string"

    # Number keywords
    if any(
        k in schema
        for k in (
            "minimum",
            "maximum",
            "exclusiveMinimum",
            "exclusiveMaximum",
            "multipleOf",
        )
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
