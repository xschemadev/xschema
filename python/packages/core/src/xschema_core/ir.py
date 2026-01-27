"""
Intermediate Representation (IR) for JSON Schema.

This module defines a discriminated union of dataclass nodes that represent
JSON Schema constructs in a normalized, language-agnostic form. The IR serves
as the bridge between JSON Schema parsing and code generation:

    JSON Schema (input) → Parser → IR (this module) → Renderer → Target Code

Design Principles:
1. DISCRIMINATED UNION - Every node has a `kind` field enabling pattern matching
2. NORMALIZED - Different JSON Schema spellings map to same IR (e.g., draft4 and
   draft2020-12 exclusive bounds both become NumberConstraints.exclusive_minimum)
3. SELF-CONTAINED - All info needed for code generation is in the node (after bundling)
4. FROZEN DATACLASSES - Immutable nodes enable safe sharing and hashing

Why IR exists (vs passing raw JSON Schema to renderers):
- SIMPLICITY: Renderers don't need to handle JSON Schema's many equivalent spellings
- TYPE SAFETY: Static types catch adapter bugs at development time
- TESTABILITY: IR nodes are easy to construct in unit tests
- SEPARATION: Parser handles JSON Schema complexity, renderer handles target language

Note: The CLI pre-processes schemas before sending to adapters:
- External $refs are bundled (fetched and inlined)
- Boolean schemas (true/false) are normalized to AnyNode/NeverNode
- Draft differences are normalized (items→prefixItems, etc.)
"""

from dataclasses import dataclass
from typing import Any, Literal, Union

# JSON value type for const/enum (any valid JSON value).
# JSON Schema equality semantics: type matters (false != 0, null != "null").
# Renderers must use JSON equality, not Python equality (Python: False == 0 is True).
JsonValue = str | int | float | bool | None | list[Any] | dict[str, Any]


# ============================================
# Constraint Types
#
# These dataclasses hold validation constraints extracted from JSON Schema.
# Constraints are separate from nodes to keep node definitions clean and
# allow sharing constraint logic across similar node types.
# ============================================


@dataclass(frozen=True)
class StringConstraints:
    """
    String validation constraints from JSON Schema.

    JSON Schema keywords:
    - minLength: minimum string length (in Unicode codepoints, not bytes)
    - maxLength: maximum string length
    - pattern: ECMA-262 regex the string must match

    Note: 'format' is stored on StringNode, not here - it's metadata not validation.
    """

    min_length: int | None = None
    max_length: int | None = None
    pattern: str | None = None


@dataclass(frozen=True)
class NumberConstraints:
    """
    Number validation constraints from JSON Schema.

    JSON Schema keywords (normalized across drafts):
    - minimum/maximum: inclusive bounds
    - exclusiveMinimum/exclusiveMaximum: exclusive bounds

    Draft compatibility note:
    - Draft 4-6: exclusiveMinimum was a boolean modifying minimum
    - Draft 7+: exclusiveMinimum is the actual boundary value
    - Parser normalizes both forms to the numeric exclusive_minimum field

    Example: {"minimum": 0, "exclusiveMinimum": true} (draft4)
             → exclusive_minimum=0, minimum=None
    """

    minimum: float | None = None
    maximum: float | None = None
    exclusive_minimum: float | None = None
    exclusive_maximum: float | None = None
    multiple_of: float | None = None


@dataclass(frozen=True)
class ContainsConstraint:
    """
    Contains constraint for arrays (draft 2019-09+).

    JSON Schema keywords:
    - contains: schema that at least one item must match
    - minContains: minimum number of items that must match (default 1)
    - maxContains: maximum number of items that may match

    Example: {"contains": {"type": "number"}, "minContains": 2}
             → array must have at least 2 numbers
    """

    schema: "SchemaNode"
    min_contains: int = 1
    max_contains: int | None = None


@dataclass(frozen=True)
class ArrayConstraints:
    """
    Array validation constraints (apply to both ArrayNode and TupleNode).

    JSON Schema keywords:
    - minItems/maxItems: array length bounds
    - uniqueItems: all items must be distinct (JSON equality)
    - contains: at least N items must match a schema

    Note: item-type constraints (items, prefixItems) are on the node,
    not in ArrayConstraints - they affect the type, not just validation.
    """

    min_items: int | None = None
    max_items: int | None = None
    unique_items: bool = False
    contains: ContainsConstraint | None = None


@dataclass(frozen=True)
class PropertyDef:
    """
    Object property definition.

    Combines the schema for a property with its required status.
    This allows ObjectNode.properties to be a single tuple rather than
    separate properties/required fields.
    """

    schema: "SchemaNode"
    required: bool = False


@dataclass(frozen=True)
class PatternPropertyDef:
    """
    Pattern property definition for regex-matched keys.

    JSON Schema keyword: patternProperties
    Keys matching the regex pattern are validated against the schema.

    Example: {"patternProperties": {"^S_": {"type": "string"}}}
             → keys starting with "S_" must have string values
    """

    pattern: str
    schema: "SchemaNode"


@dataclass(frozen=True)
class PropertyDependency:
    """
    Property dependency: if key exists, other keys must also exist.

    JSON Schema keyword: dependencies (draft 4-7) or dependentRequired (draft 2019-09+)

    Example: {"dependencies": {"credit_card": ["billing_address"]}}
             → if "credit_card" is present, "billing_address" must be too
    """

    kind: Literal["property"] = "property"
    required_properties: tuple[str, ...] = ()


@dataclass(frozen=True)
class SchemaDependency:
    """
    Schema dependency: if key exists, object must validate against schema.

    JSON Schema keyword: dependencies (draft 4-7) or dependentSchemas (draft 2019-09+)

    Example: {"dependencies": {"credit_card": {"properties": {"billing_address": {}}}}}
             → if "credit_card" present, object must have billing_address property
    """

    kind: Literal["schema"] = "schema"
    schema: "SchemaNode" = None  # type: ignore


Dependency = Union[PropertyDependency, SchemaDependency]


@dataclass(frozen=True)
class TypeGuard:
    """
    Type guard for runtime type checking in TypeGuardedNode.

    Pairs a type check with the schema to apply when that check passes.
    Multiple guards allow type-specific validation (e.g., minLength for
    strings, minimum for numbers - both from same typeless schema).
    """

    check: Literal["object", "array", "string", "number"]
    schema: "SchemaNode"


# ============================================
# Schema Nodes
#
# Each node represents a distinct JSON Schema construct.
# The `kind` field enables exhaustive pattern matching.
# Nodes are frozen dataclasses for immutability.
# ============================================


@dataclass(frozen=True)
class StringNode:
    """
    String type with optional validation constraints.

    JSON Schema: {"type": "string"} with optional minLength, maxLength, pattern, format.

    The format field is metadata (email, uri, date-time, etc.) - its validation
    behavior varies by implementation. Renderers may ignore it, use library
    validators, or generate format-specific types.
    """

    kind: Literal["string"] = "string"
    constraints: StringConstraints = StringConstraints()
    format: str | None = None


@dataclass(frozen=True)
class NumberNode:
    """
    Number or integer type with optional validation constraints.

    JSON Schema: {"type": "number"} or {"type": "integer"}
    The integer field distinguishes these - both use NumberConstraints.

    Renderer note: JSON/JavaScript don't distinguish int/float, but target
    languages may. Pydantic uses StrictInt vs StrictFloat for correct
    isinstance checks (Python: isinstance(1.0, int) is False).
    """

    kind: Literal["number"] = "number"
    constraints: NumberConstraints = NumberConstraints()
    integer: bool = False


@dataclass(frozen=True)
class BooleanNode:
    """
    Boolean type.

    JSON Schema: {"type": "boolean"}
    Only true and false are valid - 0, 1, "true", "false" are NOT booleans.
    """

    kind: Literal["boolean"] = "boolean"


@dataclass(frozen=True)
class NullNode:
    """
    Null type.

    JSON Schema: {"type": "null"}
    Only JSON null is valid - empty string, 0, false are NOT null.
    """

    kind: Literal["null"] = "null"


@dataclass(frozen=True)
class LiteralNode:
    """
    Literal/const value - schema matches exactly one value.

    JSON Schema: {"const": <value>}
    The value can be any JSON type: string, number, boolean, null, array, object.

    Renderer note: Use JSON equality for validation (type matters).
    Python: False == 0 is True, but JSON Schema: false != 0.
    """

    kind: Literal["literal"] = "literal"
    value: JsonValue = None


@dataclass(frozen=True)
class EnumNode:
    """
    Enum type - schema matches any of the listed values.

    JSON Schema: {"enum": [<value1>, <value2>, ...]}
    Values can be mixed types: {"enum": [1, "one", null]} is valid.

    Renderer note: Use JSON equality for validation (type matters).
    Values are stored in a tuple to preserve order and allow hashing.
    """

    kind: Literal["enum"] = "enum"
    values: tuple[JsonValue, ...] = ()


@dataclass(frozen=True)
class AnyNode:
    """
    Any type - accepts all valid JSON values.

    JSON Schema: {} or true (boolean schema)
    The empty schema places no constraints - everything validates.

    Generated code typically uses the target language's "any" or "unknown" type.
    """

    kind: Literal["any"] = "any"


@dataclass(frozen=True)
class NeverNode:
    """
    Never type - accepts no values (always fails validation).

    JSON Schema: false (boolean schema) or {"not": {}}
    No value can ever validate against this schema.

    Generated code typically uses the target language's "never" or bottom type.
    """

    kind: Literal["never"] = "never"


@dataclass(frozen=True)
class ArrayNode:
    """
    Homogeneous array type - all items validated against same schema.

    JSON Schema: {"type": "array", "items": <schema>}

    items field:
    - Schema that ALL array items must match
    - Parser sets to AnyNode when items keyword absent (any items allowed)
    - items: false would be TupleNode with no prefix and rest=False

    unevaluated_items field:
    - Only valid for simple cases (no allOf/anyOf/oneOf applicators)
    - CLI blocks complex unevaluated* schemas - adapters see simple cases only
    - False: no items beyond those evaluated by items are allowed
    - Schema: unevaluated items must match this schema

    Note: ArrayNode vs TupleNode distinction:
    - ArrayNode: homogeneous, all items same type (items: <schema>)
    - TupleNode: positional, specific type per index (prefixItems: [...])
    """

    kind: Literal["array"] = "array"
    items: "SchemaNode" = None  # type: ignore
    constraints: ArrayConstraints = ArrayConstraints()
    unevaluated_items: "SchemaNode | Literal[False] | None" = None


@dataclass(frozen=True)
class TupleNode:
    """
    Positional array type - items at specific positions have specific schemas.

    JSON Schema (draft 2020-12): {"prefixItems": [...], "items": <rest>}
    JSON Schema (draft 4-7): {"items": [...], "additionalItems": <rest>}
    Parser normalizes both to TupleNode.

    prefix_items: schemas for positional items (index 0, 1, 2, ...)
    rest_items: schema for items beyond prefix, or False to forbid extras

    Examples:
    - {"prefixItems": [{"type": "string"}, {"type": "number"}]}
      → first item must be string, second must be number, rest can be anything
    - {"prefixItems": [...], "items": false}
      → only prefix items allowed, no additional items
    - {"prefixItems": [...], "items": {"type": "boolean"}}
      → items beyond prefix must be booleans

    unevaluated_items: same semantics as ArrayNode (simple cases only)
    """

    kind: Literal["tuple"] = "tuple"
    prefix_items: tuple["SchemaNode", ...] = ()
    rest_items: "SchemaNode | Literal[False] | None" = None
    constraints: ArrayConstraints = ArrayConstraints()
    unevaluated_items: "SchemaNode | Literal[False] | None" = None


@dataclass(frozen=True)
class ObjectNode:
    """
    Object type with named properties and validation constraints.

    JSON Schema: {"type": "object", "properties": {...}, ...}

    PROPERTY VALIDATION ORDER (per JSON Schema spec):
    1. properties: named keys with specific schemas
    2. patternProperties: regex-matched keys with schemas
    3. additionalProperties: keys not matched by above two

    A key can match multiple patterns - ALL matching schemas must validate.
    A key in properties AND matching a pattern validates against BOTH.

    FIELD EXPLANATIONS:

    properties: tuple of (name, PropertyDef) pairs
        PropertyDef includes both schema and required status.
        Example: {"properties": {"name": {"type": "string"}}, "required": ["name"]}

    required: tuple of required property names
        Names here may NOT be in properties (required unknown key pattern).
        Example: {"required": ["id"]} without properties defining "id"

    additional_properties:
        - None: not specified, extra keys allowed (common default)
        - True: explicitly allows extra keys (marks them as "evaluated")
        - False: forbids keys not in properties/patternProperties
        - Schema: extra keys must match this schema

    pattern_properties: tuple of PatternPropertyDef
        Keys matching regex pattern must validate against schema.
        Multiple patterns can match same key - all must validate.

    property_names: schema that all keys must match (usually StringNode)
        Example: {"propertyNames": {"pattern": "^[a-z]+$"}} - lowercase keys only

    dependencies / dependent_required / dependent_schemas:
        If key X exists, then either:
        - Other keys must exist (PropertyDependency)
        - Object must validate against schema (SchemaDependency)

    unevaluated_properties:
        Only valid for simple cases (no allOf/anyOf/oneOf applicators).
        CLI blocks complex cases - adapters see simple cases only.
        - False: no properties beyond those "evaluated" are allowed
        - Schema: unevaluated properties must match this schema
        "Evaluated" means: matched by properties, patternProperties, or additionalProperties.

    NOTE ON JSON SCHEMA SEMANTICS:
    additionalProperties does NOT look inside applicators (allOf, anyOf, etc).
    {"allOf": [{"properties": {"foo": {}}}], "additionalProperties": false}
    → "foo" is NOT exempt from additionalProperties (it wasn't evaluated at this level).
    The CLI may merge allOf objects for simpler rendering, but tracks this correctly.
    """

    kind: Literal["object"] = "object"
    properties: tuple[tuple[str, PropertyDef], ...] = ()
    required: tuple[
        str, ...
    ] = ()  # Required property names (may include names not in properties)
    additional_properties: "SchemaNode | bool | None" = None
    pattern_properties: tuple[PatternPropertyDef, ...] = ()
    property_names: "SchemaNode | None" = None
    min_properties: int | None = None
    max_properties: int | None = None
    dependencies: tuple[tuple[str, Dependency], ...] = ()
    unevaluated_properties: "SchemaNode | Literal[False] | None" = None


@dataclass(frozen=True)
class UnionNode:
    """
    Union type - at least one variant must match (JSON Schema: anyOf).

    JSON Schema: {"anyOf": [<schema1>, <schema2>, ...]}
    Valid if the value matches ANY of the variant schemas.

    Code generation strategies:
    - TypeScript/Zod: union type (A | B | C)
    - Python/Pydantic: Union[A, B, C] or discriminated union if detectable

    Discriminated unions:
    Parser may detect when all variants are objects with a common literal
    property (e.g., {"type": "cat"} vs {"type": "dog"}). Renderers can use
    this for efficient tagged unions instead of try-each-variant validation.

    anyOf vs oneOf:
    - anyOf: at least one must match (permissive, may match multiple)
    - oneOf: exactly one must match (exclusive, validated at runtime)
    """

    kind: Literal["union"] = "union"
    variants: tuple["SchemaNode", ...] = ()


@dataclass(frozen=True)
class OneOfNode:
    """
    Exclusive union type - exactly one schema must match (JSON Schema: oneOf).

    JSON Schema: {"oneOf": [<schema1>, <schema2>, ...]}
    Valid if the value matches EXACTLY ONE schema. If zero or 2+ match, invalid.

    IMPORTANT SEMANTIC DIFFERENCE from anyOf:
    - anyOf: match A OR B OR both → valid
    - oneOf: match A OR B but NOT both → valid

    This requires runtime validation to ensure exactly one match.
    Some renderers implement this as anyOf + runtime check, others use
    discriminated union optimization when possible.

    Example where oneOf matters:
    {"oneOf": [{"multipleOf": 2}, {"multipleOf": 3}]}
    - 4: matches first only → valid
    - 9: matches second only → valid
    - 6: matches BOTH → INVALID (not exactly one)
    - 5: matches NONE → INVALID
    """

    kind: Literal["oneOf"] = "oneOf"
    schemas: tuple["SchemaNode", ...] = ()


@dataclass(frozen=True)
class IntersectionNode:
    """
    Intersection type - all schemas must match (JSON Schema: allOf).

    JSON Schema: {"allOf": [<schema1>, <schema2>, ...]}
    Valid if the value matches ALL schemas simultaneously.

    INDEPENDENT EVALUATION:
    Each schema evaluates independently against the complete value.
    {"allOf": [{"properties": {"a": {}}}, {"additionalProperties": false}]}
    → "a" is NOT exempt from additionalProperties (evaluated independently)

    Code generation strategies:
    - For objects: merge properties from all schemas (with conflict resolution)
    - For non-objects: intersection type (TypeScript: A & B)
    - May require runtime validation when static merge isn't possible

    Common patterns:
    - Extending base schemas: allOf with $ref and additional properties
    - Combining constraints: allOf with multiple validation rules
    - Mixin composition: combining independent object schemas
    """

    kind: Literal["intersection"] = "intersection"
    schemas: tuple["SchemaNode", ...] = ()


@dataclass(frozen=True)
class NotNode:
    """
    Negation type - value must NOT match the schema (JSON Schema: not).

    JSON Schema: {"not": <schema>}
    Valid if the value FAILS validation against the inner schema.

    CHALLENGING FOR STATIC TYPING:
    "not string" doesn't map cleanly to a single type - it's
    number | boolean | null | array | object | ...

    Renderer strategies:
    - Generate runtime validator that rejects matches
    - Use "any" type with validation (loses type safety)
    - Combine with other constraints for meaningful negation

    Example: {"type": "string", "not": {"const": ""}}
    → non-empty string (type constraint + not constraint)
    """

    kind: Literal["not"] = "not"
    schema: "SchemaNode" = None  # type: ignore


@dataclass(frozen=True)
class ConditionalNode:
    """
    Conditional validation - if/then/else branching (JSON Schema: if/then/else).

    JSON Schema: {"if": <test>, "then": <pass-schema>, "else": <fail-schema>}
    - If value matches "if" schema → also validate against "then"
    - If value doesn't match "if" schema → validate against "else"
    - Missing then/else means no additional validation for that branch

    IMPORTANT: The "if" schema is for TESTING, not filtering.
    {"if": {"type": "string"}, "then": {"minLength": 1}}
    - String "hello": if passes, then validates → valid
    - String "": if passes, then fails → INVALID
    - Number 42: if fails, no else → valid (no additional constraints)

    Code generation is tricky - may need runtime branching or type narrowing.
    """

    kind: Literal["conditional"] = "conditional"
    if_schema: "SchemaNode" = None  # type: ignore
    then_schema: "SchemaNode | None" = None
    else_schema: "SchemaNode | None" = None


@dataclass(frozen=True)
class TypeGuardedNode:
    """
    Type-guarded schema - applies different validation based on runtime type.

    Used when JSON Schema has type-specific keywords but no explicit type.
    Example: {"minLength": 1, "minimum": 0}
    - No "type" keyword, so could be any JSON value
    - minLength only applies to strings
    - minimum only applies to numbers
    - Objects, arrays, booleans, null pass with no constraints

    Parser creates TypeGuardedNode with guards:
    - {check: "string", schema: StringNode(minLength=1)}
    - {check: "number", schema: NumberNode(minimum=0)}

    Renderer generates code that:
    1. Checks runtime type
    2. Applies appropriate schema if type matches
    3. Passes through unchanged if no guard matches

    WHY THIS EXISTS:
    JSON Schema's type-specific keywords are implicitly guarded by type.
    Without explicit "type", the keywords only constrain matching types.
    A typeless {"minLength": 1} accepts numbers, objects, etc. - only
    strings are constrained to be non-empty.
    """

    kind: Literal["typeGuarded"] = "typeGuarded"
    guards: tuple[TypeGuard, ...] = ()


@dataclass(frozen=True)
class NullableNode:
    """
    Nullable wrapper - value can be the inner type OR null.

    Created from:
    - OpenAPI 3.0 "nullable: true" property
    - JSON Schema union with null: {"anyOf": [<schema>, {"type": "null"}]}

    Parser may optimize Union[T, NullNode] → NullableNode(T) for cleaner output.

    Code generation: Optional[T] or T | None in most languages.
    """

    kind: Literal["nullable"] = "nullable"
    inner: "SchemaNode" = None  # type: ignore


@dataclass(frozen=True)
class RefNode:
    """
    Reference to another schema definition.

    JSON Schema: {"$ref": "#/$defs/Name"} or {"$ref": "other.json#/..."}

    PRE-BUNDLING:
    The CLI resolves external refs before sending to adapters.
    Adapters should only see internal refs (#/$defs/...) or fully resolved nodes.

    resolved field:
    - Contains the dereferenced SchemaNode for non-recursive refs
    - Is None for recursive refs (to break cycles)

    CYCLE DETECTION:
    For recursive schemas like {"$ref": "#"} or mutually recursive defs,
    parser sets resolved=None to prevent infinite recursion. Renderer must
    handle None by generating a forward reference or lazy evaluation.

    path field:
    - JSON Pointer to the definition (e.g., "#/$defs/Person")
    - Used for generating meaningful names in output code
    - Also used when resolved=None to generate forward references
    """

    kind: Literal["ref"] = "ref"
    path: str = ""
    resolved: "SchemaNode | None" = None


# ============================================
# SchemaNode Union Type
#
# The discriminated union of all IR node types. Use the `kind` field
# for exhaustive pattern matching in renderers:
#
#   def render(node: SchemaNode) -> str:
#       match node.kind:
#           case "string": return render_string(node)
#           case "number": return render_number(node)
#           case "object": return render_object(node)
#           ...
#
# Node categories:
# - Primitives: StringNode, NumberNode, BooleanNode, NullNode
# - Values: LiteralNode, EnumNode
# - Collections: ArrayNode, TupleNode, ObjectNode
# - Composition: UnionNode, OneOfNode, IntersectionNode, NotNode, ConditionalNode
# - Special: TypeGuardedNode, NullableNode, RefNode
# - Boundaries: AnyNode (accepts all), NeverNode (rejects all)
# ============================================
SchemaNode = Union[
    # Primitives
    StringNode,
    NumberNode,
    BooleanNode,
    NullNode,
    # Values
    LiteralNode,
    EnumNode,
    # Collections
    ArrayNode,
    TupleNode,
    ObjectNode,
    # Composition
    UnionNode,
    OneOfNode,
    IntersectionNode,
    NotNode,
    ConditionalNode,
    # Special
    TypeGuardedNode,
    NullableNode,
    RefNode,
    # Boundaries
    AnyNode,
    NeverNode,
]
