"""Composition keyword parsing: allOf, anyOf, oneOf, not, if/then/else.

JSON Schema provides four composition keywords that combine schemas:

1. allOf (Intersection): ALL schemas must validate
   - Schema: {"allOf": [{min: 0}, {max: 10}]}
   - Valid: 5 (passes both), Invalid: 15 (fails max)
   - Analogous to TypeScript's intersection types: A & B

2. anyOf (Union): AT LEAST ONE schema must validate
   - Schema: {"anyOf": [{"type": "string"}, {"type": "number"}]}
   - Valid: "hello" (passes first), 42 (passes second), Invalid: null
   - Analogous to TypeScript's union types: A | B

3. oneOf (Exclusive): EXACTLY ONE schema must validate
   - Schema: {"oneOf": [{"multipleOf": 3}, {"multipleOf": 5}]}
   - Valid: 3 (only first), 5 (only second), Invalid: 15 (both pass!)
   - No direct TypeScript analog - requires runtime exactly-one check

4. not (Negation): Schema must NOT validate
   - Schema: {"not": {"type": "string"}}
   - Valid: 42, null, [], Invalid: "hello"
   - Often used with allOf for subtractive patterns

Conditional keywords (if/then/else) provide schema-level branching:
   - Schema: {"if": {"type": "string"}, "then": {"minLength": 1}, "else": {"minimum": 0}}
   - If "if" validates, apply "then"; otherwise apply "else"
   - All three are optional; missing then/else means no constraint

IMPORTANT SEMANTICS:
- Composition keywords evaluate INDEPENDENTLY. Each sub-schema sees the full instance.
- Results are combined after evaluation (all pass for allOf, any pass for anyOf, etc.)
- This affects error messages and short-circuit optimization opportunities.
- Properties defined in allOf branches do NOT exempt them from additionalProperties
  in sibling schemas - each schema evaluates independently.

Cross-keyword composition (allOf + anyOf + sibling validation):
- When multiple composition keywords appear together, all must pass (implicit intersection)
- Example: {"allOf": [A], "anyOf": [B, C], "properties": {...}}
  = allOf must pass AND at least one anyOf variant must pass AND properties must pass
- Parser wraps these in IntersectionNode to express this requirement
"""

from typing import Any, Callable, Protocol, Union, runtime_checkable

from xschema_core.ir import (
    ConditionalNode,
    IntersectionNode,
    NotNode,
    OneOfNode,
    SchemaNode,
    UnionNode,
)


@runtime_checkable
class ParseContext(Protocol):
    """Protocol for parse context - allows ref resolution and cycle detection."""

    root: Union[dict[str, Any], None]
    resolving: set[str]


# Type alias for the recursive parse function passed to composition parsers
# Using Any to avoid circular imports - actual type would be:
# Callable[[Union[dict[str, Any], bool], ParseContext], SchemaNode]
ParseFn = Callable[..., SchemaNode]


def parse_all_of(
    schemas: list[Union[dict[str, Any], bool]],
    ctx: ParseContext,
    parse_fn: ParseFn,
) -> IntersectionNode:
    """Parse allOf composition keyword.

    allOf requires ALL schemas to validate. The result is the intersection of all
    constraints - the instance must satisfy every sub-schema.

    Example:
        {"allOf": [{"type": "object", "properties": {"a": {}}},
                   {"required": ["a", "b"]}]}
        -> IntersectionNode with both object definitions

    Args:
        schemas: List of sub-schemas from allOf keyword
        ctx: Parse context for ref resolution
        parse_fn: Recursive parse function for sub-schemas

    Returns:
        IntersectionNode containing all parsed sub-schemas

    Note:
        allOf is the ONLY composition keyword that merges object types in renderers.
        anyOf/oneOf create runtime branches; allOf creates static merging.
    """
    parsed = tuple(parse_fn(s, ctx) for s in schemas)
    return IntersectionNode(schemas=parsed)


def parse_any_of(
    schemas: list[Union[dict[str, Any], bool]],
    ctx: ParseContext,
    parse_fn: ParseFn,
) -> UnionNode:
    """Parse anyOf composition keyword.

    anyOf requires AT LEAST ONE schema to validate. This is the JSON Schema
    equivalent of a union type - the instance can be any of the variants.

    Example:
        {"anyOf": [{"type": "string"}, {"type": "number"}]}
        -> UnionNode with string and number variants

    Args:
        schemas: List of variant schemas from anyOf keyword
        ctx: Parse context for ref resolution
        parse_fn: Recursive parse function for sub-schemas

    Returns:
        UnionNode containing all variants

    Note:
        anyOf is permissive - if multiple variants match, that's fine.
        Use oneOf when exactly-one matching is required.
    """
    variants = tuple(parse_fn(s, ctx) for s in schemas)
    return UnionNode(variants=variants)


def parse_one_of(
    schemas: list[Union[dict[str, Any], bool]],
    ctx: ParseContext,
    parse_fn: ParseFn,
) -> OneOfNode:
    """Parse oneOf composition keyword.

    oneOf requires EXACTLY ONE schema to validate. If zero match or more than
    one match, validation fails.

    Example:
        {"oneOf": [{"type": "string"}, {"type": "number"}]}
        -> Valid: "hello" (only string matches), 42 (only number matches)
        -> If both could match (impossible here), validation would fail

    Common use case - discriminated unions:
        {"oneOf": [
            {"properties": {"type": {"const": "a"}, ...}},
            {"properties": {"type": {"const": "b"}, ...}}
        ]}
        Each variant has unique discriminator value, so exactly one matches.

    Args:
        schemas: List of schemas from oneOf keyword
        ctx: Parse context for ref resolution
        parse_fn: Recursive parse function for sub-schemas

    Returns:
        OneOfNode containing all schemas

    Note:
        Renderers may optimize discriminated unions (detect common const property)
        to avoid validating all variants at runtime.
    """
    parsed = tuple(parse_fn(s, ctx) for s in schemas)
    return OneOfNode(schemas=parsed)


def parse_not(
    schema: Union[dict[str, Any], bool],
    ctx: ParseContext,
    parse_fn: ParseFn,
) -> NotNode:
    """Parse not composition keyword.

    not inverts validation - the instance must NOT match the inner schema.

    Example:
        {"not": {"type": "string"}}
        -> Valid: 42, null, [], {} (anything except string)
        -> Invalid: "hello"

    Common patterns:
        - {"allOf": [A, {"not": B}]} - matches A but not B (set difference)
        - {"not": {"enum": [...]}} - anything except specific values

    Args:
        schema: The schema that must NOT match
        ctx: Parse context for ref resolution
        parse_fn: Recursive parse function

    Returns:
        NotNode wrapping the parsed inner schema

    Warning:
        Not is challenging for static type systems since "not string" can't
        be expressed in most type languages. Renderers typically use runtime
        validation with Any type annotation.
    """
    inner = parse_fn(schema, ctx)
    return NotNode(schema=inner)


def parse_conditional(
    if_schema: Union[dict[str, Any], bool],
    then_schema: Union[dict[str, Any], bool, None],
    else_schema: Union[dict[str, Any], bool, None],
    ctx: ParseContext,
    parse_fn: ParseFn,
) -> ConditionalNode:
    """Parse if/then/else conditional keywords.

    Conditional validation applies different schemas based on whether the
    instance matches the "if" schema:
    - If "if" validates -> apply "then" (if present)
    - If "if" fails -> apply "else" (if present)

    Example:
        {"if": {"type": "string"},
         "then": {"minLength": 1},
         "else": {"minimum": 0}}
        -> Strings must be non-empty; numbers must be >= 0

    Args:
        if_schema: The condition schema
        then_schema: Schema to apply when if matches (None if not present)
        else_schema: Schema to apply when if fails (None if not present)
        ctx: Parse context for ref resolution
        parse_fn: Recursive parse function

    Returns:
        ConditionalNode with parsed if/then/else schemas

    Note:
        then_schema or else_schema can be None (omitted in original schema).
        When omitted, that branch imposes no constraints (equivalent to true).
        The "if" schema is always present if this function is called.
    """
    parsed_if = parse_fn(if_schema, ctx)
    parsed_then = parse_fn(then_schema, ctx) if then_schema is not None else None
    parsed_else = parse_fn(else_schema, ctx) if else_schema is not None else None
    return ConditionalNode(
        if_schema=parsed_if,
        then_schema=parsed_then,
        else_schema=parsed_else,
    )
