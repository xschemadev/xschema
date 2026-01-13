"""XSchema Core - IR types and JSON Schema parser."""

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
from xschema_core.parser import parse

__all__ = [
    "parse",
    "SchemaNode",
    "StringNode",
    "NumberNode",
    "BooleanNode",
    "NullNode",
    "LiteralNode",
    "EnumNode",
    "ArrayNode",
    "TupleNode",
    "ObjectNode",
    "UnionNode",
    "IntersectionNode",
    "RefNode",
    "AnyNode",
    "NeverNode",
    "PropertyDef",
]
