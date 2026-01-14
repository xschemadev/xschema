"""JSON Schema to IR parser."""

from typing import Any, Literal

from xschema_core.ir import (
    AnyNode,
    ArrayConstraints,
    ArrayNode,
    BooleanNode,
    ConditionalNode,
    ContainsConstraint,
    Dependency,
    EnumNode,
    IntersectionNode,
    LiteralNode,
    NeverNode,
    NotNode,
    NullNode,
    NullableNode,
    NumberConstraints,
    NumberNode,
    ObjectNode,
    OneOfNode,
    PatternPropertyDef,
    PropertyDef,
    PropertyDependency,
    SchemaNode,
    SchemaDependency,
    StringConstraints,
    StringNode,
    TupleNode,
    TypeGuard,
    TypeGuardedNode,
    UnionNode,
)


def parse(schema: dict[str, Any] | bool) -> SchemaNode:
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

    # Handle not keyword
    if "not" in schema:
        inner = parse(schema["not"])
        return NotNode(schema=inner)

    # Handle conditional (if/then/else)
    if "if" in schema:
        if_schema = parse(schema["if"])
        then_schema = parse(schema["then"]) if "then" in schema else None
        else_schema = parse(schema["else"]) if "else" in schema else None
        return ConditionalNode(
            if_schema=if_schema, then_schema=then_schema, else_schema=else_schema
        )

    # Handle nullable (OpenAPI 3.0 style)
    if schema.get("nullable") is True:
        # Remove nullable and parse the rest
        inner_schema = {k: v for k, v in schema.items() if k != "nullable"}
        inner = parse(inner_schema)
        return NullableNode(inner=inner)

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
        variant_list: list[SchemaNode] = []
        for t in schema_type:
            variant_schema = {**schema, "type": t}
            variant_list.append(parse(variant_schema))
        return UnionNode(variants=tuple(variant_list))

    # Handle prefixItems (tuple)
    if "prefixItems" in schema:
        return _parse_tuple(schema)

    # Infer type from keywords if not specified
    if schema_type is None:
        # Check if this is a type-guarded schema (multiple type-specific keywords without type)
        guards = _detect_type_guards(schema)
        if len(guards) > 1:
            return TypeGuardedNode(guards=tuple(guards))

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


def _parse_string(schema: dict[str, Any]) -> StringNode:
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


def _parse_number(schema: dict[str, Any], integer: bool) -> NumberNode:
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


def _parse_object(schema: dict[str, Any]) -> ObjectNode:
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

    # Handle patternProperties
    pattern_props: list[PatternPropertyDef] = []
    for pattern, prop_schema in schema.get("patternProperties", {}).items():
        prop_node = (
            parse(prop_schema)
            if isinstance(prop_schema, dict)
            else (AnyNode() if prop_schema is True else NeverNode())
        )
        pattern_props.append(PatternPropertyDef(pattern=pattern, schema=prop_node))

    # Handle propertyNames
    property_names_node = None
    if "propertyNames" in schema:
        property_names_node = parse(schema["propertyNames"])

    # Handle dependencies (legacy draft-4/7 keyword)
    dependencies: list[tuple[str, Dependency]] = []
    for prop_name, dep_value in schema.get("dependencies", {}).items():
        if isinstance(dep_value, list):
            # Property dependency
            dependencies.append(
                (prop_name, PropertyDependency(required_properties=tuple(dep_value)))
            )
        elif isinstance(dep_value, dict):
            # Schema dependency
            dependencies.append((prop_name, SchemaDependency(schema=parse(dep_value))))

    # Handle dependentRequired (draft 2019-09+ keyword)
    for prop_name, required_props in schema.get("dependentRequired", {}).items():
        dependencies.append(
            (prop_name, PropertyDependency(required_properties=tuple(required_props)))
        )

    # Handle dependentSchemas (draft 2019-09+ keyword)
    for prop_name, dep_schema in schema.get("dependentSchemas", {}).items():
        dependencies.append((prop_name, SchemaDependency(schema=parse(dep_schema))))

    # Handle unevaluatedProperties
    unevaluated = schema.get("unevaluatedProperties")
    unevaluated_node: SchemaNode | Literal[False] | None
    if unevaluated is None:
        unevaluated_node = None
    elif unevaluated is False:
        unevaluated_node = False
    elif isinstance(unevaluated, bool):
        unevaluated_node = AnyNode() if unevaluated else False
    else:
        unevaluated_node = parse(unevaluated)

    return ObjectNode(
        properties=tuple(properties),
        required=tuple(schema.get("required", [])),
        additional_properties=additional_node,
        pattern_properties=tuple(pattern_props),
        property_names=property_names_node,
        min_properties=schema.get("minProperties"),
        max_properties=schema.get("maxProperties"),
        dependencies=tuple(dependencies),
        unevaluated_properties=unevaluated_node,
    )


def _parse_array(schema: dict[str, Any]) -> ArrayNode | TupleNode:
    """Parse an array schema."""
    items = schema.get("items")

    items_node: SchemaNode
    if items is None:
        items_node = AnyNode()
    elif isinstance(items, bool):
        items_node = AnyNode() if items else NeverNode()
    elif isinstance(items, list):
        # Legacy tuple syntax (items as array) - convert to tuple
        return _parse_legacy_tuple(schema)
    else:
        items_node = parse(items)

    # Handle contains constraint
    contains_constraint = None
    if "contains" in schema:
        contains_schema = parse(schema["contains"])
        min_contains = schema.get("minContains", 1)
        max_contains = schema.get("maxContains")
        contains_constraint = ContainsConstraint(
            schema=contains_schema,
            min_contains=min_contains,
            max_contains=max_contains,
        )

    constraints = ArrayConstraints(
        min_items=schema.get("minItems"),
        max_items=schema.get("maxItems"),
        unique_items=schema.get("uniqueItems", False),
        contains=contains_constraint,
    )

    # Handle unevaluatedItems
    unevaluated = schema.get("unevaluatedItems")
    unevaluated_node: SchemaNode | Literal[False] | None
    if unevaluated is None:
        unevaluated_node = None
    elif unevaluated is False:
        unevaluated_node = False
    elif isinstance(unevaluated, bool):
        unevaluated_node = AnyNode() if unevaluated else False
    else:
        unevaluated_node = parse(unevaluated)

    return ArrayNode(
        items=items_node,
        constraints=constraints,
        unevaluated_items=unevaluated_node,
    )


def _parse_tuple(schema: dict[str, Any]) -> TupleNode:
    """Parse a tuple schema (using prefixItems)."""
    prefix_items = tuple(parse(item) for item in schema.get("prefixItems", []))

    # Rest items can be in 'items' when prefixItems is present
    rest = schema.get("items")
    rest_node: SchemaNode | Literal[False] | None
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

    # Handle unevaluatedItems
    unevaluated = schema.get("unevaluatedItems")
    unevaluated_tuple_node: SchemaNode | Literal[False] | None
    if unevaluated is None:
        unevaluated_tuple_node = None
    elif unevaluated is False:
        unevaluated_tuple_node = False
    elif isinstance(unevaluated, bool):
        unevaluated_tuple_node = AnyNode() if unevaluated else False
    else:
        unevaluated_tuple_node = parse(unevaluated)

    return TupleNode(
        prefix_items=prefix_items,
        rest_items=rest_node,
        constraints=constraints,
        unevaluated_items=unevaluated_tuple_node,
    )


def _parse_legacy_tuple(schema: dict[str, Any]) -> TupleNode:
    """Parse legacy tuple syntax where items is an array."""
    items_list = schema.get("items", [])
    prefix_items = tuple(
        parse(item) if isinstance(item, dict) else (AnyNode() if item else NeverNode())
        for item in items_list
    )

    # additionalItems is the rest type for legacy tuples
    additional = schema.get("additionalItems")
    rest_legacy_node: SchemaNode | Literal[False] | None
    if additional is None:
        rest_legacy_node = None
    elif additional is False:
        rest_legacy_node = False
    elif isinstance(additional, bool):
        rest_legacy_node = AnyNode() if additional else False
    else:
        rest_legacy_node = parse(additional)

    constraints = ArrayConstraints(
        min_items=schema.get("minItems"),
        max_items=schema.get("maxItems"),
        unique_items=schema.get("uniqueItems", False),
    )

    # Handle unevaluatedItems
    unevaluated = schema.get("unevaluatedItems")
    unevaluated_legacy_node: SchemaNode | Literal[False] | None
    if unevaluated is None:
        unevaluated_legacy_node = None
    elif unevaluated is False:
        unevaluated_legacy_node = False
    elif isinstance(unevaluated, bool):
        unevaluated_legacy_node = AnyNode() if unevaluated else False
    else:
        unevaluated_legacy_node = parse(unevaluated)

    return TupleNode(
        prefix_items=prefix_items,
        rest_items=rest_legacy_node,
        constraints=constraints,
        unevaluated_items=unevaluated_legacy_node,
    )


def _infer_type(schema: dict[str, Any]) -> str | None:
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
        for k in (
            "properties",
            "required",
            "additionalProperties",
            "patternProperties",
            "propertyNames",
            "minProperties",
            "maxProperties",
            "dependentRequired",
            "dependentSchemas",
        )
    ):
        return "object"

    # Array keywords
    if any(
        k in schema
        for k in (
            "items",
            "minItems",
            "maxItems",
            "uniqueItems",
            "contains",
            "minContains",
            "maxContains",
            "prefixItems",
            "additionalItems",
            "unevaluatedItems",
        )
    ):
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


def _detect_type_guards(schema: dict[str, Any]) -> list[TypeGuard]:
    """Detect type-specific constraints and create type guards.

    For schemas without an explicit type that have type-specific keywords,
    create guards that apply those constraints only to matching runtime types.
    """
    guards = []

    # Check for string-specific keywords
    if any(k in schema for k in ("minLength", "maxLength", "pattern", "format")):
        string_schema = _parse_string(schema)
        guards.append(TypeGuard(check="string", schema=string_schema))

    # Check for number-specific keywords
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
        number_schema = _parse_number(schema, integer=False)
        guards.append(TypeGuard(check="number", schema=number_schema))

    # Check for object-specific keywords
    if any(
        k in schema
        for k in (
            "properties",
            "required",
            "additionalProperties",
            "patternProperties",
            "propertyNames",
            "minProperties",
            "maxProperties",
            "dependentRequired",
            "dependentSchemas",
        )
    ):
        object_schema = _parse_object(schema)
        guards.append(TypeGuard(check="object", schema=object_schema))

    # Check for array-specific keywords
    if any(
        k in schema
        for k in (
            "items",
            "minItems",
            "maxItems",
            "uniqueItems",
            "contains",
            "minContains",
            "maxContains",
            "prefixItems",
            "additionalItems",
            "unevaluatedItems",
        )
    ):
        array_schema = _parse_array(schema)
        guards.append(TypeGuard(check="array", schema=array_schema))

    return guards
