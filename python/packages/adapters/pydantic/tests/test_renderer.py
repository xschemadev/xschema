"""Unit tests for Pydantic renderer.

Tests verify that IR nodes are correctly converted to Pydantic code.
"""

from xschema_core import (
    StringNode,
    StringConstraints,
    NumberNode,
    NumberConstraints,
    BooleanNode,
    NullNode,
    LiteralNode,
    EnumNode,
    AnyNode,
    NeverNode,
    ArrayNode,
    ArrayConstraints,
    ContainsConstraint,
    TupleNode,
    ObjectNode,
    PropertyDef,
    PatternPropertyDef,
    UnionNode,
    OneOfNode,
    IntersectionNode,
    NotNode,
    ConditionalNode,
    TypeGuardedNode,
    TypeGuard,
    NullableNode,
)
from xschema_pydantic.renderer import render


# =============================================================================
# Primitive Types
# =============================================================================


def test_render_string_unconstrained():
    """Test rendering unconstrained string."""
    node = StringNode()
    result = render(node, "Test")
    assert result.type_expr == "str"
    assert result.code == ""
    assert result.imports == set()


def test_render_string_with_length():
    """Test string with minLength/maxLength constraints."""
    node = StringNode(constraints=StringConstraints(min_length=1, max_length=100))
    result = render(node, "Test")
    assert (
        "Annotated[str, StringConstraints(min_length=1, max_length=100)]"
        in result.type_expr
    )
    assert "from typing import Annotated" in result.imports
    assert "from pydantic import StringConstraints" in result.imports


def test_render_string_with_pattern():
    """Test string with regex pattern."""
    node = StringNode(constraints=StringConstraints(pattern="^[A-Z]+$"))
    result = render(node, "Test")
    assert "Annotated[str, StringConstraints(pattern=r'^[A-Z]+$')]" in result.type_expr


def test_render_string_format_email():
    """Test string with email format."""
    node = StringNode(format="email")
    result = render(node, "Test")
    assert result.type_expr == "EmailStr"
    assert "from pydantic import EmailStr" in result.imports


def test_render_string_format_uuid():
    """Test string with uuid format."""
    node = StringNode(format="uuid")
    result = render(node, "Test")
    assert result.type_expr == "UUID"
    assert "from uuid import UUID" in result.imports


def test_render_string_format_date():
    """Test string with date format."""
    node = StringNode(format="date")
    result = render(node, "Test")
    assert result.type_expr == "date"
    assert "from datetime import date" in result.imports


def test_render_string_format_datetime():
    """Test string with date-time format."""
    node = StringNode(format="date-time")
    result = render(node, "Test")
    assert result.type_expr == "datetime"
    assert "from datetime import datetime" in result.imports


def test_render_number_unconstrained():
    """Test rendering unconstrained number (float)."""
    node = NumberNode(integer=False)
    result = render(node, "Test")
    assert result.type_expr == "float"
    assert result.code == ""


def test_render_number_integer():
    """Test rendering integer."""
    node = NumberNode(integer=True)
    result = render(node, "Test")
    assert result.type_expr == "int"


def test_render_number_with_minimum():
    """Test number with minimum constraint."""
    node = NumberNode(
        integer=True,
        constraints=NumberConstraints(minimum=0),
    )
    result = render(node, "Test")
    assert "Annotated[int, Ge(0)]" in result.type_expr
    assert "from annotated_types import Ge" in result.imports


def test_render_number_with_maximum():
    """Test number with maximum constraint."""
    node = NumberNode(
        integer=False,
        constraints=NumberConstraints(maximum=100.5),
    )
    result = render(node, "Test")
    assert "Annotated[float, Le(100.5)]" in result.type_expr
    assert "from annotated_types import Le" in result.imports


def test_render_number_with_exclusive():
    """Test number with exclusive minimum/maximum."""
    node = NumberNode(
        integer=False,
        constraints=NumberConstraints(
            exclusive_minimum=0,
            exclusive_maximum=100,
        ),
    )
    result = render(node, "Test")
    assert "Gt(0)" in result.type_expr
    assert "Lt(100)" in result.type_expr
    assert "from annotated_types import Gt" in result.imports
    assert "from annotated_types import Lt" in result.imports


def test_render_number_with_multiple_of():
    """Test number with multipleOf constraint."""
    node = NumberNode(
        integer=True,
        constraints=NumberConstraints(multiple_of=5),
    )
    result = render(node, "Test")
    assert "MultipleOf(5)" in result.type_expr
    assert "from annotated_types import MultipleOf" in result.imports


def test_render_boolean():
    """Test rendering boolean."""
    node = BooleanNode()
    result = render(node, "Test")
    assert result.type_expr == "bool"
    assert result.code == ""


def test_render_null():
    """Test rendering null."""
    node = NullNode()
    result = render(node, "Test")
    assert result.type_expr == "None"
    assert result.code == ""


def test_render_literal_string():
    """Test rendering literal string."""
    node = LiteralNode(value="active")
    result = render(node, "Test")
    # json.dumps is used for strings, so double quotes
    assert result.type_expr == 'Literal["active"]'
    assert "from typing import Literal" in result.imports


def test_render_literal_number():
    """Test rendering literal number."""
    node = LiteralNode(value=42)
    result = render(node, "Test")
    assert result.type_expr == "Literal[42]"


def test_render_literal_boolean():
    """Test rendering literal boolean.

    Booleans use custom validator to prevent Python's True==1 and False==0 coercion.
    """
    node = LiteralNode(value=True)
    result = render(node, "Test")
    # booleans need strict type checking, so we use BeforeValidator
    assert "_make_const_validator(True)" in result.type_expr
    assert "from pydantic import BeforeValidator" in result.imports


def test_render_enum():
    """Test rendering enum (multiple literals)."""
    node = EnumNode(values=("red", "green", "blue"))
    result = render(node, "Test")
    # json.dumps is used for strings, so double quotes
    assert result.type_expr == 'Literal["red", "green", "blue"]'
    assert "from typing import Literal" in result.imports


def test_render_any():
    """Test rendering Any type."""
    node = AnyNode()
    result = render(node, "Test")
    assert result.type_expr == "Any"
    assert "from typing import Any" in result.imports


def test_render_never():
    """Test rendering Never type (always fails validation)."""
    node = NeverNode()
    result = render(node, "Test")
    assert "Annotated[Any, BeforeValidator(_never_validator)]" in result.type_expr
    assert "def _never_validator" in result.code
    assert "from pydantic import BeforeValidator" in result.imports


# =============================================================================
# Array Types
# =============================================================================


def test_render_array_simple():
    """Test rendering simple array without constraints."""
    node = ArrayNode(
        items=StringNode(),
        constraints=ArrayConstraints(),
    )
    result = render(node, "Test")
    assert result.type_expr == "list[str]"


def test_render_array_with_length_constraints():
    """Test array with minItems/maxItems."""
    node = ArrayNode(
        items=NumberNode(integer=True),
        constraints=ArrayConstraints(min_items=1, max_items=10),
    )
    result = render(node, "Test")
    assert (
        "Annotated[list[int], Field(min_length=1, max_length=10)]" in result.type_expr
    )
    assert "from pydantic import Field" in result.imports


def test_render_array_with_unique_items():
    """Test array with uniqueItems constraint."""
    node = ArrayNode(
        items=StringNode(),
        constraints=ArrayConstraints(unique_items=True),
    )
    result = render(node, "Test")
    assert "AfterValidator(_unique_test)" in result.type_expr
    assert "def _unique_test" in result.code
    assert "from pydantic import AfterValidator" in result.imports


def test_render_array_with_contains():
    """Test array with contains constraint."""
    node = ArrayNode(
        items=AnyNode(),
        constraints=ArrayConstraints(
            contains=ContainsConstraint(
                schema=StringNode(),
                min_contains=1,
            ),
        ),
    )
    result = render(node, "Test")
    assert "AfterValidator(_contains_test)" in result.type_expr
    assert "def _contains_test" in result.code
    assert "def _validates_against" in result.code


# =============================================================================
# Tuple Types
# =============================================================================


def test_render_tuple_fixed():
    """Test rendering fixed-size tuple."""
    node = TupleNode(
        prefix_items=(
            StringNode(),
            NumberNode(integer=True),
            BooleanNode(),
        ),
        rest_items=False,
        constraints=ArrayConstraints(),
    )
    result = render(node, "Test")
    assert result.type_expr == "tuple[str, int, bool]"


def test_render_tuple_with_rest():
    """Test rendering tuple with rest items."""
    node = TupleNode(
        prefix_items=(
            StringNode(),
            NumberNode(integer=True),
        ),
        rest_items=BooleanNode(),
        constraints=ArrayConstraints(),
    )
    result = render(node, "Test")
    assert (
        "Annotated[tuple[Any, ...], BeforeValidator(_tuple_test)]" in result.type_expr
    )
    assert "def _tuple_test" in result.code


def test_render_tuple_with_constraints():
    """Test tuple with minItems/maxItems."""
    node = TupleNode(
        prefix_items=(StringNode(),),
        rest_items=False,
        constraints=ArrayConstraints(min_items=1, max_items=1),
    )
    result = render(node, "Test")
    assert "AfterValidator(_tuple_constraints_test)" in result.type_expr
    assert "def _tuple_constraints_test" in result.code


# =============================================================================
# Object Types
# =============================================================================


def test_render_object_simple():
    """Test rendering simple object with properties."""
    node = ObjectNode(
        properties=(
            ("name", PropertyDef(schema=StringNode(), required=True)),
            ("age", PropertyDef(schema=NumberNode(integer=True), required=False)),
        ),
    )
    result = render(node, "Test")
    assert "class Test(BaseModel):" in result.code
    assert "name: str" in result.code
    assert "age: int | None = None" in result.code
    assert "from pydantic import BaseModel" in result.imports


def test_render_object_nested():
    """Test rendering nested objects."""
    inner = ObjectNode(
        properties=(("city", PropertyDef(schema=StringNode(), required=True)),),
    )
    outer = ObjectNode(
        properties=(("address", PropertyDef(schema=inner, required=True)),),
    )
    result = render(outer, "Test")
    assert "class TestAddress(BaseModel):" in result.code
    assert "class Test(BaseModel):" in result.code
    assert "address: TestAddress" in result.code


def test_render_object_additional_properties_false():
    """Test object with additionalProperties: false."""
    node = ObjectNode(
        properties=(("name", PropertyDef(schema=StringNode(), required=True)),),
        additional_properties=False,
    )
    result = render(node, "Test")
    assert "model_config = ConfigDict(extra='forbid')" in result.code
    assert "from pydantic import ConfigDict" in result.imports


def test_render_object_with_pattern_properties():
    """Test object with patternProperties."""
    node = ObjectNode(
        pattern_properties=(
            PatternPropertyDef(pattern="^num_", schema=NumberNode(integer=True)),
        ),
    )
    result = render(node, "Test")
    assert "@model_validator(mode='after')" in result.code
    assert "def _validate_advanced" in result.code
    assert "re.compile" in result.code
    assert "from pydantic import model_validator" in result.imports


def test_render_object_with_min_max_properties():
    """Test object with minProperties/maxProperties."""
    node = ObjectNode(
        min_properties=1,
        max_properties=5,
    )
    result = render(node, "Test")
    assert "@model_validator(mode='after')" in result.code
    assert "prop_count = len(self.__dict__)" in result.code
    assert "if prop_count < 1:" in result.code
    assert "if prop_count > 5:" in result.code


# =============================================================================
# Union & Composition Types
# =============================================================================


def test_render_union_simple():
    """Test rendering simple union (anyOf)."""
    node = UnionNode(
        variants=(
            StringNode(),
            NumberNode(integer=False),
        )
    )
    result = render(node, "Test")
    assert result.type_expr == "str | float"


def test_render_union_discriminated():
    """Test rendering discriminated union."""
    cat = ObjectNode(
        properties=(
            ("type", PropertyDef(schema=LiteralNode(value="cat"), required=True)),
            ("meow", PropertyDef(schema=BooleanNode(), required=True)),
        ),
    )
    dog = ObjectNode(
        properties=(
            ("type", PropertyDef(schema=LiteralNode(value="dog"), required=True)),
            ("bark", PropertyDef(schema=BooleanNode(), required=True)),
        ),
    )
    node = UnionNode(variants=(cat, dog))
    result = render(node, "Test")
    assert "Annotated[Union[" in result.type_expr
    assert "Field(discriminator='type')" in result.type_expr
    assert "from pydantic import Field" in result.imports


def test_render_oneof():
    """Test rendering oneOf (exactly one match)."""
    node = OneOfNode(
        schemas=(
            StringNode(),
            NumberNode(integer=False),
        )
    )
    result = render(node, "Test")
    assert "Annotated[Any, BeforeValidator(_oneof_test)]" in result.type_expr
    assert "def _oneof_test" in result.code
    assert "exactly one schema" in result.code


def test_render_intersection_objects():
    """Test rendering intersection of objects (allOf)."""
    obj1 = ObjectNode(
        properties=(("name", PropertyDef(schema=StringNode(), required=True)),),
    )
    obj2 = ObjectNode(
        properties=(
            ("age", PropertyDef(schema=NumberNode(integer=True), required=True)),
        ),
    )
    node = IntersectionNode(schemas=(obj1, obj2))
    result = render(node, "Test")
    assert "class Test(BaseModel):" in result.code
    assert "name: str" in result.code
    assert "age: int" in result.code


def test_render_intersection_primitives():
    """Test rendering intersection of primitives."""
    node = IntersectionNode(
        schemas=(
            StringNode(),
            StringNode(
                constraints=StringConstraints(min_length=1),
            ),
        )
    )
    result = render(node, "Test")
    assert "Annotated[Any, BeforeValidator(_intersection_test)]" in result.type_expr
    assert "def _intersection_test" in result.code


def test_render_not():
    """Test rendering not (negation)."""
    node = NotNode(schema=NullNode())
    result = render(node, "Test")
    assert "Annotated[Any, BeforeValidator(_not_test)]" in result.type_expr
    assert "def _not_test" in result.code


# =============================================================================
# Advanced Types
# =============================================================================


def test_render_conditional():
    """Test rendering if/then/else conditional."""
    node = ConditionalNode(
        if_schema=ObjectNode(
            properties=(
                (
                    "type",
                    PropertyDef(schema=LiteralNode(value="premium"), required=True),
                ),
            ),
        ),
        then_schema=ObjectNode(
            properties=(
                (
                    "discount",
                    PropertyDef(
                        schema=NumberNode(
                            integer=False,
                            constraints=NumberConstraints(minimum=0.1),
                        ),
                        required=True,
                    ),
                ),
            ),
        ),
        else_schema=None,
    )
    result = render(node, "Test")
    assert "Annotated[Any, BeforeValidator(_conditional_test)]" in result.type_expr
    assert "def _conditional_test" in result.code


def test_render_type_guarded():
    """Test rendering type-guarded schema (typeless with type-specific keywords)."""
    node = TypeGuardedNode(
        guards=(
            TypeGuard(
                check="string",
                schema=StringNode(
                    constraints=StringConstraints(min_length=1),
                ),
            ),
            TypeGuard(
                check="number",
                schema=NumberNode(
                    integer=False,
                    constraints=NumberConstraints(minimum=0),
                ),
            ),
        )
    )
    result = render(node, "Test")
    assert "Annotated[Any, BeforeValidator(_type_guarded_test)]" in result.type_expr
    assert "def _type_guarded_test" in result.code
    assert "isinstance" in result.code


def test_render_nullable():
    """Test rendering nullable type (OpenAPI 3.0 nullable: true)."""
    inner = StringNode()
    node = NullableNode(inner=inner)
    result = render(node, "Test")
    assert result.type_expr == "str | None"


def test_render_nullable_already_nullable():
    """Test rendering nullable when inner type is already nullable."""
    inner = UnionNode(
        variants=(
            StringNode(),
            NullNode(),
        )
    )
    node = NullableNode(inner=inner)
    result = render(node, "Test")
    # Should not add duplicate | None
    assert result.type_expr.count(" | None") == 1
