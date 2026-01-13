"""IR node types for representing JSON Schema as a typed AST."""

from dataclasses import dataclass
from typing import Literal, Union


@dataclass(frozen=True)
class StringNode:
    """String type with optional constraints."""

    kind: Literal["string"] = "string"
    format: str | None = None
    min_length: int | None = None
    max_length: int | None = None
    pattern: str | None = None
    description: str | None = None


@dataclass(frozen=True)
class NumberNode:
    """Number or integer type with optional constraints."""

    kind: Literal["number"] = "number"
    integer: bool = False
    minimum: float | None = None
    maximum: float | None = None
    exclusive_minimum: float | None = None
    exclusive_maximum: float | None = None
    multiple_of: float | None = None
    description: str | None = None


@dataclass(frozen=True)
class BooleanNode:
    """Boolean type."""

    kind: Literal["boolean"] = "boolean"
    description: str | None = None


@dataclass(frozen=True)
class NullNode:
    """Null type."""

    kind: Literal["null"] = "null"
    description: str | None = None


@dataclass(frozen=True)
class LiteralNode:
    """Literal/const value."""

    value: str | int | float | bool | None
    kind: Literal["literal"] = "literal"
    description: str | None = None


@dataclass(frozen=True)
class EnumNode:
    """Enum type with list of allowed values."""

    values: tuple[str | int | float | bool | None, ...]
    kind: Literal["enum"] = "enum"
    description: str | None = None


@dataclass(frozen=True)
class AnyNode:
    """Any type (accepts all values)."""

    kind: Literal["any"] = "any"
    description: str | None = None


@dataclass(frozen=True)
class NeverNode:
    """Never type (accepts no values)."""

    kind: Literal["never"] = "never"
    description: str | None = None


@dataclass(frozen=True)
class PropertyDef:
    """Object property definition."""

    schema: "SchemaNode"
    required: bool = False
    description: str | None = None


# Forward references for recursive types
@dataclass(frozen=True)
class ArrayNode:
    """Array type with item schema and constraints."""

    items: "SchemaNode"
    kind: Literal["array"] = "array"
    min_items: int | None = None
    max_items: int | None = None
    unique_items: bool = False
    description: str | None = None


@dataclass(frozen=True)
class TupleNode:
    """Tuple type with fixed prefix items and optional rest."""

    prefix_items: tuple["SchemaNode", ...]
    kind: Literal["tuple"] = "tuple"
    rest_items: "SchemaNode | None" = None
    description: str | None = None


@dataclass(frozen=True)
class ObjectNode:
    """Object type with properties and constraints."""

    kind: Literal["object"] = "object"
    properties: tuple[tuple[str, PropertyDef], ...] = ()
    additional_properties: "SchemaNode | bool | None" = None
    required: frozenset[str] = frozenset()
    description: str | None = None


@dataclass(frozen=True)
class UnionNode:
    """Union type (anyOf/oneOf)."""

    variants: tuple["SchemaNode", ...]
    kind: Literal["union"] = "union"
    discriminator: str | None = None
    description: str | None = None


@dataclass(frozen=True)
class IntersectionNode:
    """Intersection type (allOf)."""

    schemas: tuple["SchemaNode", ...]
    kind: Literal["intersection"] = "intersection"
    description: str | None = None


@dataclass(frozen=True)
class RefNode:
    """Reference to another schema."""

    path: str
    kind: Literal["ref"] = "ref"
    resolved: "SchemaNode | None" = None
    description: str | None = None


# Union of all schema node types
SchemaNode = Union[
    StringNode,
    NumberNode,
    BooleanNode,
    NullNode,
    LiteralNode,
    EnumNode,
    ArrayNode,
    TupleNode,
    ObjectNode,
    UnionNode,
    IntersectionNode,
    RefNode,
    AnyNode,
    NeverNode,
]
