"""IR node types for representing JSON Schema as a typed AST."""

from dataclasses import dataclass
from typing import Literal, Union


# ============================================
# Constraint Types
# ============================================


@dataclass(frozen=True)
class StringConstraints:
    """String validation constraints."""

    min_length: int | None = None
    max_length: int | None = None
    pattern: str | None = None


@dataclass(frozen=True)
class NumberConstraints:
    """Number validation constraints."""

    minimum: float | None = None
    maximum: float | None = None
    exclusive_minimum: float | None = None
    exclusive_maximum: float | None = None
    multiple_of: float | None = None


@dataclass(frozen=True)
class ContainsConstraint:
    """Contains constraint for arrays."""

    schema: "SchemaNode"
    min_contains: int = 1
    max_contains: int | None = None


@dataclass(frozen=True)
class ArrayConstraints:
    """Array validation constraints."""

    min_items: int | None = None
    max_items: int | None = None
    unique_items: bool = False
    contains: ContainsConstraint | None = None


@dataclass(frozen=True)
class PropertyDef:
    """Object property definition."""

    schema: "SchemaNode"
    required: bool = False


@dataclass(frozen=True)
class PatternPropertyDef:
    """Pattern property definition."""

    pattern: str
    schema: "SchemaNode"


@dataclass(frozen=True)
class PropertyDependency:
    """Property dependency (if property X exists, properties Y,Z must exist)."""

    kind: Literal["property"] = "property"
    required_properties: tuple[str, ...] = ()


@dataclass(frozen=True)
class SchemaDependency:
    """Schema dependency (if property X exists, schema must validate)."""

    kind: Literal["schema"] = "schema"
    schema: "SchemaNode" = None  # type: ignore


Dependency = Union[PropertyDependency, SchemaDependency]


@dataclass(frozen=True)
class TypeGuard:
    """Type guard with runtime type check and schema."""

    check: Literal["object", "array", "string", "number"]
    schema: "SchemaNode"


# ============================================
# Schema Nodes
# ============================================


@dataclass(frozen=True)
class StringNode:
    """String type with optional constraints."""

    kind: Literal["string"] = "string"
    constraints: StringConstraints = StringConstraints()
    format: str | None = None


@dataclass(frozen=True)
class NumberNode:
    """Number or integer type with optional constraints."""

    kind: Literal["number"] = "number"
    constraints: NumberConstraints = NumberConstraints()
    integer: bool = False


@dataclass(frozen=True)
class BooleanNode:
    """Boolean type."""

    kind: Literal["boolean"] = "boolean"


@dataclass(frozen=True)
class NullNode:
    """Null type."""

    kind: Literal["null"] = "null"


@dataclass(frozen=True)
class LiteralNode:
    """Literal/const value."""

    kind: Literal["literal"] = "literal"
    value: str | int | float | bool | None = None


@dataclass(frozen=True)
class EnumNode:
    """Enum type with list of allowed values."""

    kind: Literal["enum"] = "enum"
    values: tuple[str | int | float | bool | None, ...] = ()


@dataclass(frozen=True)
class AnyNode:
    """Any type (accepts all values)."""

    kind: Literal["any"] = "any"


@dataclass(frozen=True)
class NeverNode:
    """Never type (accepts no values)."""

    kind: Literal["never"] = "never"


@dataclass(frozen=True)
class ArrayNode:
    """Array type with item schema and constraints."""

    kind: Literal["array"] = "array"
    items: "SchemaNode" = None  # type: ignore
    constraints: ArrayConstraints = ArrayConstraints()
    unevaluated_items: "SchemaNode | Literal[False] | None" = None


@dataclass(frozen=True)
class TupleNode:
    """Tuple type with fixed prefix items and optional rest."""

    kind: Literal["tuple"] = "tuple"
    prefix_items: tuple["SchemaNode", ...] = ()
    rest_items: "SchemaNode | Literal[False] | None" = None
    constraints: ArrayConstraints = ArrayConstraints()
    unevaluated_items: "SchemaNode | Literal[False] | None" = None


@dataclass(frozen=True)
class ObjectNode:
    """Object type with properties and constraints."""

    kind: Literal["object"] = "object"
    properties: tuple[tuple[str, PropertyDef], ...] = ()
    additional_properties: "SchemaNode | bool | None" = None
    pattern_properties: tuple[PatternPropertyDef, ...] = ()
    property_names: "SchemaNode | None" = None
    min_properties: int | None = None
    max_properties: int | None = None
    dependencies: tuple[tuple[str, Dependency], ...] = ()
    unevaluated_properties: "SchemaNode | Literal[False] | None" = None


@dataclass(frozen=True)
class UnionNode:
    """Union type (anyOf - at least one must match)."""

    kind: Literal["union"] = "union"
    variants: tuple["SchemaNode", ...] = ()


@dataclass(frozen=True)
class OneOfNode:
    """OneOf type (exactly one must match)."""

    kind: Literal["oneOf"] = "oneOf"
    schemas: tuple["SchemaNode", ...] = ()


@dataclass(frozen=True)
class IntersectionNode:
    """Intersection type (allOf - all must match)."""

    kind: Literal["intersection"] = "intersection"
    schemas: tuple["SchemaNode", ...] = ()


@dataclass(frozen=True)
class NotNode:
    """Not type (must not match)."""

    kind: Literal["not"] = "not"
    schema: "SchemaNode" = None  # type: ignore


@dataclass(frozen=True)
class ConditionalNode:
    """Conditional type (if/then/else)."""

    kind: Literal["conditional"] = "conditional"
    if_schema: "SchemaNode" = None  # type: ignore
    then_schema: "SchemaNode | None" = None
    else_schema: "SchemaNode | None" = None


@dataclass(frozen=True)
class TypeGuardedNode:
    """Type-guarded schema (applies schema based on runtime type)."""

    kind: Literal["typeGuarded"] = "typeGuarded"
    guards: tuple[TypeGuard, ...] = ()


@dataclass(frozen=True)
class NullableNode:
    """Nullable wrapper (OpenAPI 3.0 nullable or union with null)."""

    kind: Literal["nullable"] = "nullable"
    inner: "SchemaNode" = None  # type: ignore


@dataclass(frozen=True)
class RefNode:
    """Reference to another schema (should be pre-bundled by CLI)."""

    kind: Literal["ref"] = "ref"
    path: str = ""
    resolved: "SchemaNode | None" = None


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
    OneOfNode,
    IntersectionNode,
    NotNode,
    ConditionalNode,
    TypeGuardedNode,
    NullableNode,
    RefNode,
    AnyNode,
    NeverNode,
]
