"""Unit tests for JSON Schema to IR parser."""

import pytest

from xschema_core import parse
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
    StringNode,
    TupleNode,
    UnionNode,
)


class TestBooleanSchema:
    """Test parsing boolean schemas."""

    def test_true_schema_returns_any_node(self):
        result = parse(True)
        assert isinstance(result, AnyNode)

    def test_false_schema_returns_never_node(self):
        result = parse(False)
        assert isinstance(result, NeverNode)


class TestEmptySchema:
    """Test parsing empty schemas."""

    def test_empty_dict_returns_any_node(self):
        result = parse({})
        assert isinstance(result, AnyNode)


class TestStringParsing:
    """Test parsing string schemas with constraints."""

    def test_basic_string(self):
        result = parse({"type": "string"})
        assert isinstance(result, StringNode)
        assert result.kind == "string"
        assert result.format is None
        assert result.min_length is None
        assert result.max_length is None
        assert result.pattern is None

    def test_string_with_format(self):
        result = parse({"type": "string", "format": "email"})
        assert isinstance(result, StringNode)
        assert result.format == "email"

    def test_string_with_min_length(self):
        result = parse({"type": "string", "minLength": 1})
        assert isinstance(result, StringNode)
        assert result.min_length == 1

    def test_string_with_max_length(self):
        result = parse({"type": "string", "maxLength": 100})
        assert isinstance(result, StringNode)
        assert result.max_length == 100

    def test_string_with_pattern(self):
        result = parse({"type": "string", "pattern": "^[a-z]+$"})
        assert isinstance(result, StringNode)
        assert result.pattern == "^[a-z]+$"

    def test_string_with_all_constraints(self):
        result = parse({
            "type": "string",
            "format": "uri",
            "minLength": 5,
            "maxLength": 200,
            "pattern": "^https://",
        })
        assert isinstance(result, StringNode)
        assert result.format == "uri"
        assert result.min_length == 5
        assert result.max_length == 200
        assert result.pattern == "^https://"

    def test_string_with_description(self):
        result = parse({"type": "string", "description": "A user's email"})
        assert isinstance(result, StringNode)
        assert result.description == "A user's email"


class TestNumberParsing:
    """Test parsing number/integer schemas with constraints."""

    def test_basic_number(self):
        result = parse({"type": "number"})
        assert isinstance(result, NumberNode)
        assert result.kind == "number"
        assert result.integer is False

    def test_integer_type(self):
        result = parse({"type": "integer"})
        assert isinstance(result, NumberNode)
        assert result.integer is True

    def test_number_with_minimum(self):
        result = parse({"type": "number", "minimum": 0})
        assert isinstance(result, NumberNode)
        assert result.minimum == 0

    def test_number_with_maximum(self):
        result = parse({"type": "number", "maximum": 100})
        assert isinstance(result, NumberNode)
        assert result.maximum == 100

    def test_number_with_exclusive_minimum_draft6(self):
        """Draft-6+ style: exclusiveMinimum as number."""
        result = parse({"type": "number", "exclusiveMinimum": 0})
        assert isinstance(result, NumberNode)
        assert result.exclusive_minimum == 0
        assert result.minimum is None

    def test_number_with_exclusive_maximum_draft6(self):
        """Draft-6+ style: exclusiveMaximum as number."""
        result = parse({"type": "number", "exclusiveMaximum": 100})
        assert isinstance(result, NumberNode)
        assert result.exclusive_maximum == 100
        assert result.maximum is None

    def test_number_with_exclusive_minimum_draft4(self):
        """Draft-4 style: exclusiveMinimum as boolean with minimum."""
        result = parse({"type": "number", "minimum": 0, "exclusiveMinimum": True})
        assert isinstance(result, NumberNode)
        assert result.exclusive_minimum == 0
        assert result.minimum is None

    def test_number_with_exclusive_maximum_draft4(self):
        """Draft-4 style: exclusiveMaximum as boolean with maximum."""
        result = parse({"type": "number", "maximum": 100, "exclusiveMaximum": True})
        assert isinstance(result, NumberNode)
        assert result.exclusive_maximum == 100
        assert result.maximum is None

    def test_number_with_multiple_of(self):
        result = parse({"type": "integer", "multipleOf": 5})
        assert isinstance(result, NumberNode)
        assert result.multiple_of == 5

    def test_number_with_all_constraints(self):
        result = parse({
            "type": "number",
            "minimum": 0,
            "maximum": 100,
            "multipleOf": 0.5,
        })
        assert isinstance(result, NumberNode)
        assert result.minimum == 0
        assert result.maximum == 100
        assert result.multiple_of == 0.5

    def test_number_with_description(self):
        result = parse({"type": "number", "description": "Price in dollars"})
        assert isinstance(result, NumberNode)
        assert result.description == "Price in dollars"


class TestBooleanParsing:
    """Test parsing boolean type."""

    def test_basic_boolean(self):
        result = parse({"type": "boolean"})
        assert isinstance(result, BooleanNode)
        assert result.kind == "boolean"

    def test_boolean_with_description(self):
        result = parse({"type": "boolean", "description": "Is active flag"})
        assert isinstance(result, BooleanNode)
        assert result.description == "Is active flag"


class TestNullParsing:
    """Test parsing null type."""

    def test_basic_null(self):
        result = parse({"type": "null"})
        assert isinstance(result, NullNode)
        assert result.kind == "null"

    def test_null_with_description(self):
        result = parse({"type": "null", "description": "Represents absence"})
        assert isinstance(result, NullNode)
        assert result.description == "Represents absence"


class TestObjectParsing:
    """Test parsing object schemas with properties and required."""

    def test_basic_object(self):
        result = parse({"type": "object"})
        assert isinstance(result, ObjectNode)
        assert result.kind == "object"
        assert result.properties == ()
        assert result.required == frozenset()

    def test_object_with_properties(self):
        result = parse({
            "type": "object",
            "properties": {
                "name": {"type": "string"},
                "age": {"type": "integer"},
            },
        })
        assert isinstance(result, ObjectNode)
        assert len(result.properties) == 2

        props_dict = dict(result.properties)
        assert "name" in props_dict
        assert "age" in props_dict
        assert isinstance(props_dict["name"].schema, StringNode)
        assert isinstance(props_dict["age"].schema, NumberNode)

    def test_object_with_required(self):
        result = parse({
            "type": "object",
            "properties": {
                "id": {"type": "string"},
                "name": {"type": "string"},
            },
            "required": ["id"],
        })
        assert isinstance(result, ObjectNode)
        assert result.required == frozenset(["id"])

        props_dict = dict(result.properties)
        assert props_dict["id"].required is True
        assert props_dict["name"].required is False

    def test_object_with_additional_properties_false(self):
        result = parse({
            "type": "object",
            "properties": {"name": {"type": "string"}},
            "additionalProperties": False,
        })
        assert isinstance(result, ObjectNode)
        assert result.additional_properties is False

    def test_object_with_additional_properties_true(self):
        result = parse({
            "type": "object",
            "additionalProperties": True,
        })
        assert isinstance(result, ObjectNode)
        assert result.additional_properties is True

    def test_object_with_additional_properties_schema(self):
        result = parse({
            "type": "object",
            "additionalProperties": {"type": "string"},
        })
        assert isinstance(result, ObjectNode)
        assert isinstance(result.additional_properties, StringNode)

    def test_object_with_description(self):
        result = parse({
            "type": "object",
            "description": "A user object",
            "properties": {"name": {"type": "string"}},
        })
        assert isinstance(result, ObjectNode)
        assert result.description == "A user object"

    def test_object_property_descriptions(self):
        result = parse({
            "type": "object",
            "properties": {
                "name": {"type": "string", "description": "The user's name"},
            },
        })
        assert isinstance(result, ObjectNode)
        props_dict = dict(result.properties)
        assert props_dict["name"].description == "The user's name"


class TestArrayParsing:
    """Test parsing array schemas with items."""

    def test_basic_array(self):
        result = parse({"type": "array"})
        assert isinstance(result, ArrayNode)
        assert result.kind == "array"
        assert isinstance(result.items, AnyNode)

    def test_array_with_items(self):
        result = parse({
            "type": "array",
            "items": {"type": "string"},
        })
        assert isinstance(result, ArrayNode)
        assert isinstance(result.items, StringNode)

    def test_array_with_min_items(self):
        result = parse({
            "type": "array",
            "items": {"type": "number"},
            "minItems": 1,
        })
        assert isinstance(result, ArrayNode)
        assert result.min_items == 1

    def test_array_with_max_items(self):
        result = parse({
            "type": "array",
            "items": {"type": "number"},
            "maxItems": 10,
        })
        assert isinstance(result, ArrayNode)
        assert result.max_items == 10

    def test_array_with_unique_items(self):
        result = parse({
            "type": "array",
            "items": {"type": "string"},
            "uniqueItems": True,
        })
        assert isinstance(result, ArrayNode)
        assert result.unique_items is True

    def test_array_with_all_constraints(self):
        result = parse({
            "type": "array",
            "items": {"type": "integer"},
            "minItems": 1,
            "maxItems": 100,
            "uniqueItems": True,
        })
        assert isinstance(result, ArrayNode)
        assert result.min_items == 1
        assert result.max_items == 100
        assert result.unique_items is True
        assert isinstance(result.items, NumberNode)

    def test_array_with_description(self):
        result = parse({
            "type": "array",
            "description": "List of tags",
            "items": {"type": "string"},
        })
        assert isinstance(result, ArrayNode)
        assert result.description == "List of tags"


class TestTupleParsing:
    """Test parsing tuple schemas with prefixItems."""

    def test_tuple_with_prefix_items(self):
        result = parse({
            "type": "array",
            "prefixItems": [
                {"type": "string"},
                {"type": "number"},
            ],
        })
        assert isinstance(result, TupleNode)
        assert result.kind == "tuple"
        assert len(result.prefix_items) == 2
        assert isinstance(result.prefix_items[0], StringNode)
        assert isinstance(result.prefix_items[1], NumberNode)

    def test_tuple_with_rest_items(self):
        result = parse({
            "type": "array",
            "prefixItems": [
                {"type": "string"},
            ],
            "items": {"type": "number"},
        })
        assert isinstance(result, TupleNode)
        assert len(result.prefix_items) == 1
        assert isinstance(result.rest_items, NumberNode)

    def test_tuple_with_no_rest_items(self):
        result = parse({
            "type": "array",
            "prefixItems": [
                {"type": "string"},
                {"type": "boolean"},
            ],
            "items": False,
        })
        assert isinstance(result, TupleNode)
        assert result.rest_items is None

    def test_legacy_tuple_items_array(self):
        """Test legacy tuple syntax where items is an array."""
        result = parse({
            "type": "array",
            "items": [
                {"type": "string"},
                {"type": "number"},
            ],
        })
        assert isinstance(result, TupleNode)
        assert len(result.prefix_items) == 2
        assert isinstance(result.prefix_items[0], StringNode)
        assert isinstance(result.prefix_items[1], NumberNode)

    def test_legacy_tuple_with_additional_items(self):
        """Test legacy tuple with additionalItems."""
        result = parse({
            "type": "array",
            "items": [{"type": "string"}],
            "additionalItems": {"type": "number"},
        })
        assert isinstance(result, TupleNode)
        assert isinstance(result.rest_items, NumberNode)

    def test_tuple_with_description(self):
        result = parse({
            "description": "Coordinate pair",
            "prefixItems": [
                {"type": "number"},
                {"type": "number"},
            ],
        })
        assert isinstance(result, TupleNode)
        assert result.description == "Coordinate pair"


class TestEnumParsing:
    """Test parsing enum schemas."""

    def test_string_enum(self):
        result = parse({"enum": ["red", "green", "blue"]})
        assert isinstance(result, EnumNode)
        assert result.kind == "enum"
        assert result.values == ("red", "green", "blue")

    def test_mixed_enum(self):
        result = parse({"enum": [1, "two", True, None]})
        assert isinstance(result, EnumNode)
        assert result.values == (1, "two", True, None)

    def test_numeric_enum(self):
        result = parse({"enum": [0, 1, 2, 3]})
        assert isinstance(result, EnumNode)
        assert result.values == (0, 1, 2, 3)

    def test_enum_with_description(self):
        result = parse({
            "enum": ["pending", "active", "completed"],
            "description": "Status values",
        })
        assert isinstance(result, EnumNode)
        assert result.description == "Status values"


class TestConstParsing:
    """Test parsing const (literal) schemas."""

    def test_string_const(self):
        result = parse({"const": "fixed_value"})
        assert isinstance(result, LiteralNode)
        assert result.kind == "literal"
        assert result.value == "fixed_value"

    def test_number_const(self):
        result = parse({"const": 42})
        assert isinstance(result, LiteralNode)
        assert result.value == 42

    def test_boolean_const(self):
        result = parse({"const": True})
        assert isinstance(result, LiteralNode)
        assert result.value is True

    def test_null_const(self):
        result = parse({"const": None})
        assert isinstance(result, LiteralNode)
        assert result.value is None

    def test_const_with_description(self):
        result = parse({"const": "v1", "description": "API version"})
        assert isinstance(result, LiteralNode)
        assert result.description == "API version"


class TestAnyOfUnion:
    """Test parsing anyOf unions."""

    def test_anyof_primitives(self):
        result = parse({
            "anyOf": [
                {"type": "string"},
                {"type": "number"},
            ],
        })
        assert isinstance(result, UnionNode)
        assert result.kind == "union"
        assert len(result.variants) == 2
        assert isinstance(result.variants[0], StringNode)
        assert isinstance(result.variants[1], NumberNode)
        assert result.discriminator is None

    def test_anyof_with_null(self):
        """Common pattern: nullable type."""
        result = parse({
            "anyOf": [
                {"type": "string"},
                {"type": "null"},
            ],
        })
        assert isinstance(result, UnionNode)
        assert len(result.variants) == 2
        assert isinstance(result.variants[0], StringNode)
        assert isinstance(result.variants[1], NullNode)

    def test_anyof_with_description(self):
        result = parse({
            "anyOf": [{"type": "string"}, {"type": "number"}],
            "description": "String or number",
        })
        assert isinstance(result, UnionNode)
        assert result.description == "String or number"


class TestOneOfWithDiscriminator:
    """Test parsing oneOf with discriminator detection."""

    def test_oneof_without_discriminator(self):
        """oneOf with primitives - no discriminator."""
        result = parse({
            "oneOf": [
                {"type": "string"},
                {"type": "number"},
            ],
        })
        assert isinstance(result, UnionNode)
        assert result.discriminator is None

    def test_oneof_with_discriminator(self):
        """oneOf with discriminated objects."""
        result = parse({
            "oneOf": [
                {
                    "type": "object",
                    "properties": {
                        "type": {"const": "dog"},
                        "breed": {"type": "string"},
                    },
                    "required": ["type"],
                },
                {
                    "type": "object",
                    "properties": {
                        "type": {"const": "cat"},
                        "color": {"type": "string"},
                    },
                    "required": ["type"],
                },
            ],
        })
        assert isinstance(result, UnionNode)
        assert result.discriminator == "type"
        assert len(result.variants) == 2

    def test_oneof_discriminator_detection_no_common_const(self):
        """oneOf with objects but no common const property."""
        result = parse({
            "oneOf": [
                {
                    "type": "object",
                    "properties": {
                        "kind": {"const": "a"},
                        "value": {"type": "string"},
                    },
                },
                {
                    "type": "object",
                    "properties": {
                        "type": {"const": "b"},  # Different property name
                        "value": {"type": "number"},
                    },
                },
            ],
        })
        assert isinstance(result, UnionNode)
        assert result.discriminator is None

    def test_oneof_with_description(self):
        result = parse({
            "oneOf": [{"type": "string"}, {"type": "null"}],
            "description": "Optional string",
        })
        assert isinstance(result, UnionNode)
        assert result.description == "Optional string"


class TestAllOfIntersection:
    """Test parsing allOf intersections."""

    def test_allof_objects(self):
        result = parse({
            "allOf": [
                {
                    "type": "object",
                    "properties": {"name": {"type": "string"}},
                },
                {
                    "type": "object",
                    "properties": {"age": {"type": "number"}},
                },
            ],
        })
        assert isinstance(result, IntersectionNode)
        assert result.kind == "intersection"
        assert len(result.schemas) == 2
        assert isinstance(result.schemas[0], ObjectNode)
        assert isinstance(result.schemas[1], ObjectNode)

    def test_allof_with_ref_and_object(self):
        """Common pattern: extend a base schema."""
        result = parse({
            "allOf": [
                {"$ref": "#/definitions/Base"},
                {
                    "type": "object",
                    "properties": {"extra": {"type": "string"}},
                },
            ],
        })
        assert isinstance(result, IntersectionNode)
        assert len(result.schemas) == 2
        assert isinstance(result.schemas[0], RefNode)
        assert isinstance(result.schemas[1], ObjectNode)

    def test_allof_with_description(self):
        result = parse({
            "allOf": [{"type": "object"}, {"type": "object"}],
            "description": "Combined schema",
        })
        assert isinstance(result, IntersectionNode)
        assert result.description == "Combined schema"


class TestNestedSchemas:
    """Test parsing deeply nested schemas."""

    def test_nested_object_in_array(self):
        result = parse({
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "name": {"type": "string"},
                    "tags": {
                        "type": "array",
                        "items": {"type": "string"},
                    },
                },
            },
        })
        assert isinstance(result, ArrayNode)
        assert isinstance(result.items, ObjectNode)

        props_dict = dict(result.items.properties)
        assert isinstance(props_dict["name"].schema, StringNode)
        assert isinstance(props_dict["tags"].schema, ArrayNode)
        assert isinstance(props_dict["tags"].schema.items, StringNode)

    def test_nested_union_in_object(self):
        result = parse({
            "type": "object",
            "properties": {
                "value": {
                    "anyOf": [
                        {"type": "string"},
                        {"type": "number"},
                        {"type": "null"},
                    ],
                },
            },
        })
        assert isinstance(result, ObjectNode)
        props_dict = dict(result.properties)
        assert isinstance(props_dict["value"].schema, UnionNode)
        assert len(props_dict["value"].schema.variants) == 3

    def test_deeply_nested_structure(self):
        result = parse({
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "items": {
                            "type": "array",
                            "items": {
                                "type": "object",
                                "properties": {
                                    "id": {"type": "string"},
                                },
                            },
                        },
                    },
                },
            },
        })
        assert isinstance(result, ObjectNode)

        # Navigate: data -> items -> item -> id
        data_prop = dict(result.properties)["data"]
        assert isinstance(data_prop.schema, ObjectNode)

        items_prop = dict(data_prop.schema.properties)["items"]
        assert isinstance(items_prop.schema, ArrayNode)
        assert isinstance(items_prop.schema.items, ObjectNode)

        id_prop = dict(items_prop.schema.items.properties)["id"]
        assert isinstance(id_prop.schema, StringNode)


class TestRefParsing:
    """Test parsing $ref schemas."""

    def test_ref_basic(self):
        result = parse({"$ref": "#/definitions/User"})
        assert isinstance(result, RefNode)
        assert result.kind == "ref"
        assert result.path == "#/definitions/User"

    def test_ref_with_description(self):
        result = parse({
            "$ref": "#/definitions/Address",
            "description": "User's address",
        })
        assert isinstance(result, RefNode)
        assert result.description == "User's address"

    def test_ref_external(self):
        result = parse({"$ref": "./common.json#/definitions/UUID"})
        assert isinstance(result, RefNode)
        assert result.path == "./common.json#/definitions/UUID"


class TestTypeInference:
    """Test type inference from keywords when type is not specified."""

    def test_infer_string_from_min_length(self):
        result = parse({"minLength": 1})
        assert isinstance(result, StringNode)

    def test_infer_string_from_max_length(self):
        result = parse({"maxLength": 100})
        assert isinstance(result, StringNode)

    def test_infer_string_from_pattern(self):
        result = parse({"pattern": "^[a-z]+$"})
        assert isinstance(result, StringNode)

    def test_infer_string_from_format(self):
        result = parse({"format": "email"})
        assert isinstance(result, StringNode)

    def test_infer_number_from_minimum(self):
        result = parse({"minimum": 0})
        assert isinstance(result, NumberNode)

    def test_infer_number_from_maximum(self):
        result = parse({"maximum": 100})
        assert isinstance(result, NumberNode)

    def test_infer_number_from_multiple_of(self):
        result = parse({"multipleOf": 5})
        assert isinstance(result, NumberNode)

    def test_infer_object_from_properties(self):
        result = parse({"properties": {"name": {"type": "string"}}})
        assert isinstance(result, ObjectNode)

    def test_infer_object_from_required(self):
        result = parse({"required": ["name"]})
        assert isinstance(result, ObjectNode)

    def test_infer_object_from_additional_properties(self):
        result = parse({"additionalProperties": False})
        assert isinstance(result, ObjectNode)

    def test_infer_array_from_items(self):
        result = parse({"items": {"type": "string"}})
        assert isinstance(result, ArrayNode)

    def test_infer_array_from_min_items(self):
        result = parse({"minItems": 1})
        assert isinstance(result, ArrayNode)

    def test_infer_array_from_unique_items(self):
        result = parse({"uniqueItems": True})
        assert isinstance(result, ArrayNode)


class TestTypeArrayUnion:
    """Test parsing type as array (union of types)."""

    def test_nullable_string(self):
        result = parse({"type": ["string", "null"]})
        assert isinstance(result, UnionNode)
        assert len(result.variants) == 2
        assert isinstance(result.variants[0], StringNode)
        assert isinstance(result.variants[1], NullNode)

    def test_multiple_types(self):
        result = parse({"type": ["string", "number", "boolean"]})
        assert isinstance(result, UnionNode)
        assert len(result.variants) == 3
        assert isinstance(result.variants[0], StringNode)
        assert isinstance(result.variants[1], NumberNode)
        assert isinstance(result.variants[2], BooleanNode)

    def test_type_array_with_constraints(self):
        """Constraints should apply to each type."""
        result = parse({
            "type": ["string", "null"],
            "minLength": 1,
        })
        assert isinstance(result, UnionNode)
        assert isinstance(result.variants[0], StringNode)
        assert result.variants[0].min_length == 1
