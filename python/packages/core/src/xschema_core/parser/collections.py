"""
Parsers for collection types: object, array, tuple.

This module handles parsing JSON Schema collection types into IR nodes:
- Objects: properties, additionalProperties, patternProperties, etc.
- Arrays: items validation for variable-length arrays
- Tuples: prefixItems for fixed-length arrays with per-position schemas

JSON Schema draft compatibility notes:
- draft4/draft6/draft7: items as array means tuple, additionalItems for rest
- draft2019-09+: prefixItems for tuple, items for rest, additionalItems ignored
- dependencies keyword (draft4-7) splits into dependentRequired + dependentSchemas (draft2019-09+)
"""

from typing import Any, Literal, Union

from xschema_core.ir import (
    AnyNode,
    ArrayConstraints,
    ArrayNode,
    ContainsConstraint,
    Dependency,
    NeverNode,
    ObjectNode,
    PatternPropertyDef,
    PropertyDef,
    PropertyDependency,
    SchemaNode,
    SchemaDependency,
    TupleNode,
)


def parse_object(schema: dict[str, Any], ctx: Any) -> ObjectNode:
    """Parse an object schema into an ObjectNode.

    Handles:
    - properties: named property schemas
    - required: required property names
    - additionalProperties: schema for properties not in properties/patternProperties
    - patternProperties: regex-keyed property schemas
    - propertyNames: schema that all property names must match
    - minProperties/maxProperties: count constraints
    - dependencies/dependentRequired/dependentSchemas: property interdependencies
    - unevaluatedProperties: schema for properties not validated by other keywords

    additionalProperties interaction with patternProperties:
    - A property is "additional" only if it doesn't match ANY patternProperty pattern
    - patternProperties and additionalProperties can coexist
    - If both match a key, BOTH schemas must validate

    Example schemas:
    - Object with known properties: {"type": "object", "properties": {"name": {"type": "string"}}}
    - Strict object (no extra keys): {"type": "object", "additionalProperties": false}
    - String-keyed dict: {"type": "object", "additionalProperties": {"type": "number"}}
    - Pattern properties: {"patternProperties": {"^[Ss]": {"type": "string"}}}
    """
    from xschema_core.parser.core import _parse_with_ctx

    properties: list[tuple[str, PropertyDef]] = []
    required_set = frozenset(schema.get("required", []))

    # Parse properties
    for prop_name, prop_schema in schema.get("properties", {}).items():
        prop_node = (
            _parse_with_ctx(prop_schema, ctx)
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
    # - None/undefined: additional properties allowed with any type (open object)
    # - false: no additional properties allowed (closed object)
    # - true: explicitly allow any additional properties
    # - schema: additional properties must match schema
    additional = schema.get("additionalProperties")
    if additional is None:
        additional_node: Union[SchemaNode, bool, None] = None
    elif isinstance(additional, bool):
        additional_node = additional
    else:
        additional_node = _parse_with_ctx(additional, ctx)

    # Handle patternProperties
    # Validates properties whose names match regex patterns
    # A property can match multiple patterns - all must validate
    pattern_props: list[PatternPropertyDef] = []
    for pattern, prop_schema in schema.get("patternProperties", {}).items():
        prop_node = (
            _parse_with_ctx(prop_schema, ctx)
            if isinstance(prop_schema, dict)
            else (AnyNode() if prop_schema is True else NeverNode())
        )
        pattern_props.append(PatternPropertyDef(pattern=pattern, schema=prop_node))

    # Handle propertyNames
    # Schema that all property names (keys) must validate against
    # Example: {"propertyNames": {"pattern": "^[A-Z]"}} requires uppercase keys
    property_names_node = None
    if "propertyNames" in schema:
        property_names_node = _parse_with_ctx(schema["propertyNames"], ctx)

    # Handle dependencies
    # Legacy draft4-7: "dependencies" keyword (array = properties, object = schema)
    # Modern draft2019-09+: "dependentRequired" + "dependentSchemas"
    # Both are supported here for compatibility
    dependencies: list[tuple[str, Dependency]] = []

    # Legacy dependencies keyword (draft4-7)
    for prop_name, dep_value in schema.get("dependencies", {}).items():
        if isinstance(dep_value, list):
            # Property dependency: if prop_name present, these properties required
            dependencies.append(
                (prop_name, PropertyDependency(required_properties=tuple(dep_value)))
            )
        elif isinstance(dep_value, dict):
            # Schema dependency: if prop_name present, schema must validate
            dependencies.append(
                (prop_name, SchemaDependency(schema=_parse_with_ctx(dep_value, ctx)))
            )

    # Modern dependentRequired (draft2019-09+)
    # If property X exists, properties [Y, Z] must also exist
    for prop_name, required_props in schema.get("dependentRequired", {}).items():
        dependencies.append(
            (prop_name, PropertyDependency(required_properties=tuple(required_props)))
        )

    # Modern dependentSchemas (draft2019-09+)
    # If property X exists, schema must validate against the object
    for prop_name, dep_schema in schema.get("dependentSchemas", {}).items():
        dependencies.append(
            (prop_name, SchemaDependency(schema=_parse_with_ctx(dep_schema, ctx)))
        )

    # Handle unevaluatedProperties
    # Applies to properties not evaluated by properties/patternProperties/etc
    # More strict than additionalProperties - considers properties from allOf/anyOf/oneOf
    # Note: Complex unevaluatedProperties with applicators is validated by Go CLI
    unevaluated = schema.get("unevaluatedProperties")
    unevaluated_node: Union[SchemaNode, Literal[False], None]
    if unevaluated is None:
        unevaluated_node = None
    elif unevaluated is False:
        unevaluated_node = False
    elif isinstance(unevaluated, bool):
        unevaluated_node = AnyNode() if unevaluated else False
    else:
        unevaluated_node = _parse_with_ctx(unevaluated, ctx)

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


def parse_array(schema: dict[str, Any], ctx: Any) -> Union[ArrayNode, TupleNode]:
    """Parse an array schema into an ArrayNode or TupleNode.

    Returns TupleNode if:
    - prefixItems keyword present (draft2019-09+ tuple syntax)
    - items is an array (draft4-7 legacy tuple syntax)

    Returns ArrayNode otherwise (variable-length array with single item schema).

    Array constraints (apply to both ArrayNode and TupleNode):
    - minItems/maxItems: length constraints
    - uniqueItems: all items must be distinct (deep equality)
    - contains: at least one item must match schema (with min/maxContains)

    items keyword semantics by draft:
    - draft4-7: array means tuple (per-position schemas), object/bool means array items
    - draft2019-09+: only object/bool (array items), tuples use prefixItems

    additionalItems vs items in legacy drafts:
    - When items is array (tuple), additionalItems applies to extra positions
    - When items is object/bool (array), additionalItems ignored
    - draft2019-09+ removes additionalItems, items now means "rest items" after prefixItems

    Example schemas:
    - String array: {"type": "array", "items": {"type": "string"}}
    - Tuple: {"prefixItems": [{"type": "number"}, {"type": "string"}]}
    - Legacy tuple: {"items": [{"type": "number"}, {"type": "string"}], "additionalItems": false}
    - Contains: {"contains": {"type": "number"}, "minContains": 2}
    """
    from xschema_core.parser.core import _parse_with_ctx

    # Check for tuple syntax (prefixItems or legacy items as array)
    if "prefixItems" in schema:
        return parse_tuple(schema, ctx)

    items = schema.get("items")
    if isinstance(items, list):
        # Legacy tuple syntax (items as array) - convert to tuple
        return parse_legacy_tuple(schema, ctx)

    # Regular array - single schema for all items
    items_node: SchemaNode
    if items is None:
        items_node = AnyNode()
    elif isinstance(items, bool):
        items_node = AnyNode() if items else NeverNode()
    else:
        items_node = _parse_with_ctx(items, ctx)

    # Handle contains constraint
    # Requires at least minContains (default 1) items match the schema
    # At most maxContains (if specified) items can match
    contains_constraint = None
    if "contains" in schema:
        contains_schema = _parse_with_ctx(schema["contains"], ctx)
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
    # Applies to items not evaluated by items/prefixItems/contains
    # More strict than items - considers items from allOf/anyOf/oneOf
    unevaluated = schema.get("unevaluatedItems")
    unevaluated_node: Union[SchemaNode, Literal[False], None]
    if unevaluated is None:
        unevaluated_node = None
    elif unevaluated is False:
        unevaluated_node = False
    elif isinstance(unevaluated, bool):
        unevaluated_node = AnyNode() if unevaluated else False
    else:
        unevaluated_node = _parse_with_ctx(unevaluated, ctx)

    return ArrayNode(
        items=items_node,
        constraints=constraints,
        unevaluated_items=unevaluated_node,
    )


def parse_tuple(schema: dict[str, Any], ctx: Any) -> TupleNode:
    """Parse a tuple schema using prefixItems (draft2019-09+ syntax).

    Tuple structure:
    - prefixItems: array of schemas for positions 0, 1, 2, ...
    - items: schema for positions beyond prefixItems (rest items)
    - items: false means no extra items allowed (closed tuple)

    Example schemas:
    - Fixed 2-tuple: {"prefixItems": [{"type": "number"}, {"type": "string"}], "items": false}
    - Open tuple: {"prefixItems": [{"type": "number"}], "items": {"type": "string"}}
    - Pair with no extra: {"prefixItems": [{}, {}], "items": false}
    """
    from xschema_core.parser.core import _parse_with_ctx

    prefix_items = tuple(
        _parse_with_ctx(item, ctx) for item in schema.get("prefixItems", [])
    )

    # Rest items schema (applies to positions beyond prefixItems)
    # - None/undefined: any additional items allowed
    # - false: no additional items allowed (closed tuple)
    # - schema: additional items must match schema
    rest = schema.get("items")
    rest_node: Union[SchemaNode, Literal[False], None]
    if rest is None:
        rest_node = None
    elif rest is False:
        rest_node = False
    elif isinstance(rest, bool):
        rest_node = AnyNode() if rest else False
    else:
        rest_node = _parse_with_ctx(rest, ctx)

    constraints = ArrayConstraints(
        min_items=schema.get("minItems"),
        max_items=schema.get("maxItems"),
        unique_items=schema.get("uniqueItems", False),
    )

    # Handle unevaluatedItems for tuples
    unevaluated = schema.get("unevaluatedItems")
    unevaluated_tuple_node: Union[SchemaNode, Literal[False], None]
    if unevaluated is None:
        unevaluated_tuple_node = None
    elif unevaluated is False:
        unevaluated_tuple_node = False
    elif isinstance(unevaluated, bool):
        unevaluated_tuple_node = AnyNode() if unevaluated else False
    else:
        unevaluated_tuple_node = _parse_with_ctx(unevaluated, ctx)

    return TupleNode(
        prefix_items=prefix_items,
        rest_items=rest_node,
        constraints=constraints,
        unevaluated_items=unevaluated_tuple_node,
    )


def parse_legacy_tuple(schema: dict[str, Any], ctx: Any) -> TupleNode:
    """Parse legacy tuple syntax where items is an array (draft4-7).

    Legacy tuple structure:
    - items: array of schemas for positions [0, 1, 2, ...]
    - additionalItems: schema for positions beyond items array
    - additionalItems: false means no extra items allowed

    This syntax is deprecated in draft2019-09+ but still widely used.
    Modern equivalent uses prefixItems + items instead of items + additionalItems.

    Example schemas:
    - Legacy tuple: {"items": [{"type": "number"}, {"type": "string"}], "additionalItems": false}
    - Modern equivalent: {"prefixItems": [{"type": "number"}, {"type": "string"}], "items": false}
    """
    from xschema_core.parser.core import _parse_with_ctx

    items_list = schema.get("items", [])
    prefix_items = tuple(
        _parse_with_ctx(item, ctx)
        if isinstance(item, dict)
        else (AnyNode() if item else NeverNode())
        for item in items_list
    )

    # additionalItems is the rest type for legacy tuples
    # Semantically equivalent to items in prefixItems-based tuples
    additional = schema.get("additionalItems")
    rest_legacy_node: Union[SchemaNode, Literal[False], None]
    if additional is None:
        rest_legacy_node = None
    elif additional is False:
        rest_legacy_node = False
    elif isinstance(additional, bool):
        rest_legacy_node = AnyNode() if additional else False
    else:
        rest_legacy_node = _parse_with_ctx(additional, ctx)

    constraints = ArrayConstraints(
        min_items=schema.get("minItems"),
        max_items=schema.get("maxItems"),
        unique_items=schema.get("uniqueItems", False),
    )

    # Handle unevaluatedItems for legacy tuples
    unevaluated = schema.get("unevaluatedItems")
    unevaluated_legacy_node: Union[SchemaNode, Literal[False], None]
    if unevaluated is None:
        unevaluated_legacy_node = None
    elif unevaluated is False:
        unevaluated_legacy_node = False
    elif isinstance(unevaluated, bool):
        unevaluated_legacy_node = AnyNode() if unevaluated else False
    else:
        unevaluated_legacy_node = _parse_with_ctx(unevaluated, ctx)

    return TupleNode(
        prefix_items=prefix_items,
        rest_items=rest_legacy_node,
        constraints=constraints,
        unevaluated_items=unevaluated_legacy_node,
    )
