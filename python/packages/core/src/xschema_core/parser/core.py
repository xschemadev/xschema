"""Core JSON Schema to IR parsing logic.

This module contains:
- parse(): Main entry point that converts JSON Schema to IR
- _resolve_json_pointer(): JSON pointer resolution (#/$defs/Name)
- _resolve_ref(): $ref resolution with cycle detection
- _parse_with_ctx(): Main recursive parsing function
- _infer_type(): Type inference from keywords
- _detect_type_guards(): Type guard detection for typeless schemas
- _detect_discriminator(): Discriminated union detection for oneOf

The parse function expects schemas to be pre-bundled by the CLI.
All external $refs should be resolved before reaching this parser.
Internal JSON pointer refs (#/$defs/Name) are resolved during parsing.
"""

from typing import Any, Dict, List, Optional, Set, Union

from xschema_core.ir import (
    AnyNode,
    BooleanNode,
    IntersectionNode,
    NeverNode,
    NullNode,
    NullableNode,
    RefNode,
    SchemaNode,
    TypeGuard,
    TypeGuardedNode,
    UnionNode,
)
from xschema_core.parser.primitives import parse_number, parse_string
from xschema_core.parser.collections import (
    parse_array,
    parse_object,
    parse_tuple,
)
from xschema_core.parser.composition import (
    parse_all_of,
    parse_any_of,
    parse_conditional,
    parse_not,
    parse_one_of,
)
from xschema_core.parser.values import parse_enum, parse_literal
from xschema_core.parser.context import ParseContext, create_context


def parse(schema: Union[Dict[str, Any], bool]) -> SchemaNode:
    """Parse a JSON Schema dict into an IR SchemaNode.

    This is the main entry point for converting JSON Schema to IR.

    Args:
        schema: A JSON Schema as a dict, or a boolean schema (true/false).
                Schemas should be pre-bundled by the CLI - all external $refs
                should already be resolved. Internal JSON pointer refs
                (#/$defs/Name) are supported.

    Returns:
        A SchemaNode representing the schema in IR form.

    Example:
        >>> from xschema_core.parser import parse
        >>> node = parse({"type": "string", "minLength": 1})
        >>> type(node).__name__
        'StringNode'
        >>> node.constraints.min_length
        1
    """
    # Create context with root schema for internal $ref resolution
    ctx = create_context(schema)
    return _parse_with_ctx(schema, ctx)


def _resolve_json_pointer(ref: str, root: dict[str, Any]) -> Any:
    """Resolve a JSON pointer (fragment starting with #/) to a schema node.

    JSON Pointer is defined in RFC 6901. The pointer starts with # and uses
    / to separate path segments. Special characters are escaped:
    - ~0 represents ~
    - ~1 represents /
    - URI percent-encoding (e.g., %25 for %) is decoded first

    Args:
        ref: JSON pointer string like "#/$defs/Name" or "#/properties/foo%2Fbar"
        root: The root schema to resolve against

    Returns:
        The schema node at the pointer location

    Raises:
        ValueError: If the pointer path doesn't exist
    """
    from urllib.parse import unquote

    path_parts = ref[2:].split("/")  # Remove "#/" prefix
    current: Any = root

    for part in path_parts:
        # Handle URI percent-encoding first (e.g., %25 -> %)
        part = unquote(part)
        # Then handle JSON pointer escaping (~1 -> /, ~0 -> ~)
        part = part.replace("~1", "/").replace("~0", "~")

        if isinstance(current, dict) and part in current:
            current = current[part]
        elif isinstance(current, list):
            # Handle list index
            try:
                idx = int(part)
                if 0 <= idx < len(current):
                    current = current[idx]
                else:
                    raise ValueError(f"list index {idx} out of range")
            except ValueError:
                raise ValueError(f"invalid list index '{part}'")
        else:
            raise ValueError(f"path not found at '{part}'")

    return current


def _resolve_ref(ref: str, ctx: ParseContext) -> SchemaNode:
    """Resolve a $ref to a SchemaNode.

    Only handles internal JSON pointer refs (e.g., #/$defs/Name, #/properties/foo).
    External refs should be bundled by the CLI before reaching the adapter.

    Cycle Detection:
        Uses ctx.resolving set to track refs being resolved. When a cycle is
        detected (ref already in resolving set), returns RefNode(resolved=None).
        This allows adapters to handle recursive types appropriately.

    Args:
        ref: The $ref string value
        ctx: Parse context with root schema and cycle tracking

    Returns:
        RefNode with the ref path and resolved schema (or None for cycles)

    Raises:
        ValueError: If ref is external (not starting with #) or unresolvable
    """
    if ctx.root is None:
        raise ValueError(
            f"Encountered $ref '{ref}' - schemas must be bundled by the Go CLI before processing. "
            "Run the schema through xschema generate to bundle all references."
        )

    # Only handle JSON pointer refs starting with #
    if not ref.startswith("#"):
        raise ValueError(
            f"Encountered external $ref '{ref}' - schemas must be bundled by the Go CLI before processing. "
            "Run the schema through xschema generate to bundle all references."
        )

    # Check for cycles - return RefNode with resolved=None to break infinite recursion
    if ref in ctx.resolving:
        return RefNode(path=ref, resolved=None)

    # Root ref - recursive to root schema
    if ref == "#":
        return RefNode(path=ref, resolved=None)

    # Resolve JSON pointer
    try:
        target_schema = _resolve_json_pointer(ref, ctx.root)
    except ValueError as e:
        raise ValueError(
            f"Failed to resolve $ref '{ref}': {e}. "
            "The schema may be malformed or the CLI bundler may have an issue."
        )

    # Mark as resolving to detect cycles
    ctx.resolving.add(ref)
    try:
        resolved = _parse_with_ctx(target_schema, ctx)
    finally:
        ctx.resolving.discard(ref)

    return RefNode(path=ref, resolved=resolved)


def _parse_with_ctx(
    schema: Union[Dict[str, Any], bool], ctx: ParseContext
) -> SchemaNode:
    """Parse a JSON Schema dict into an IR SchemaNode with context.

    This is the main recursive parsing function. It handles all JSON Schema
    keywords and dispatches to specialized parsers for different types.

    Parsing order matters - keywords are checked in this order:
    1. Boolean schemas (true/false)
    2. Empty schema
    3. $ref (terminates - no sibling keywords processed)
    4. const (literal value)
    5. enum (value set)
    6. not (negation)
    7. if/then/else (conditional)
    8. nullable (OpenAPI extension)
    9. Composition keywords (allOf, anyOf, oneOf)
    10. Type-based parsing
    """
    # Handle boolean schemas
    if schema is True:
        return AnyNode()
    if schema is False:
        return NeverNode()

    # Handle empty schema
    if not schema:
        return AnyNode()

    # Handle $ref - CLI should pre-bundle all schemas
    if "$ref" in schema:
        ref = schema["$ref"]
        return _resolve_ref(ref, ctx)

    # Handle const (literal value)
    if "const" in schema:
        return parse_literal(schema)

    # Handle enum
    if "enum" in schema:
        return parse_enum(schema)

    # Handle not keyword
    if "not" in schema:
        return parse_not(schema["not"], ctx, _parse_with_ctx)

    # Handle conditional (if/then/else)
    if "if" in schema:
        return parse_conditional(
            schema["if"],
            schema.get("then"),
            schema.get("else"),
            ctx,
            _parse_with_ctx,
        )

    # Handle nullable (OpenAPI 3.0 style)
    if schema.get("nullable") is True:
        # Remove nullable and parse the rest
        inner_schema = {k: v for k, v in schema.items() if k != "nullable"}
        inner = _parse_with_ctx(inner_schema, ctx)
        return NullableNode(inner=inner)

    # Handle composition keywords
    # Check if there are sibling validation keywords that need to be combined
    composition_keys = {"allOf", "anyOf", "oneOf"}
    meta_keys = {
        "$schema",
        "$id",
        "$ref",
        "$defs",
        "definitions",
        "$comment",
        "title",
        "description",
        "examples",
        "default",
        "deprecated",
        "readOnly",
        "writeOnly",
        "$anchor",
    }
    # Keys that are meaningful only with composition (not standalone validation)
    composition_only_keys = composition_keys | meta_keys

    has_composition = bool(composition_keys & set(schema.keys()))
    has_sibling_validation = bool(set(schema.keys()) - composition_only_keys)
    composition_count = len(composition_keys & set(schema.keys()))

    # Multiple composition keywords (allOf + anyOf, etc.) or composition + sibling validation
    # All must be combined with intersection
    if composition_count > 1 or (has_composition and has_sibling_validation):
        composition_nodes: list[SchemaNode] = []
        if "allOf" in schema:
            # Flatten allOf into intersection members
            all_of_node = parse_all_of(schema["allOf"], ctx, _parse_with_ctx)
            composition_nodes.extend(all_of_node.schemas)
        if "anyOf" in schema:
            composition_nodes.append(
                parse_any_of(schema["anyOf"], ctx, _parse_with_ctx)
            )
        if "oneOf" in schema:
            composition_nodes.append(
                parse_one_of(schema["oneOf"], ctx, _parse_with_ctx)
            )

        # Add sibling validation schema if present
        if has_sibling_validation:
            sibling_schema = {
                k: v for k, v in schema.items() if k not in composition_keys
            }
            sibling_node = _parse_with_ctx(sibling_schema, ctx)
            all_schemas = tuple(composition_nodes) + (sibling_node,)
        else:
            all_schemas = tuple(composition_nodes)

        return IntersectionNode(schemas=all_schemas)

    # Single composition keyword without sibling validation
    if "allOf" in schema:
        return parse_all_of(schema["allOf"], ctx, _parse_with_ctx)

    if "anyOf" in schema:
        return parse_any_of(schema["anyOf"], ctx, _parse_with_ctx)

    if "oneOf" in schema:
        return parse_one_of(schema["oneOf"], ctx, _parse_with_ctx)

    # Get type - can be string or array
    schema_type = schema.get("type")

    # Handle array of types (union)
    if isinstance(schema_type, list):
        variant_list: list[SchemaNode] = []
        for t in schema_type:
            variant_schema = {**schema, "type": t}
            variant_list.append(_parse_with_ctx(variant_schema, ctx))
        return UnionNode(variants=tuple(variant_list))

    # Handle prefixItems (tuple) only if type is explicitly "array"
    # If no type is specified, prefixItems should be wrapped in a type guard
    # so it only applies to arrays (not objects that look like arrays)
    if "prefixItems" in schema and schema_type == "array":
        return parse_tuple(schema, ctx)

    # Infer type from keywords if not specified
    if schema_type is None:
        # Check if this is a type-guarded schema (multiple type-specific keywords without type)
        guards = _detect_type_guards(schema, ctx)
        if len(guards) > 1:
            return TypeGuardedNode(guards=tuple(guards))

        # Special case: single array guard with prefixItems needs TypeGuardedNode
        # to allow non-arrays through (JSON Schema semantics: prefixItems only applies to arrays)
        if len(guards) == 1 and guards[0].check == "array" and "prefixItems" in schema:
            return TypeGuardedNode(guards=tuple(guards))

        schema_type = _infer_type(schema)

    # Parse based on type
    if schema_type == "string":
        return parse_string(schema)

    if schema_type == "number":
        return parse_number(schema, integer=False)

    if schema_type == "integer":
        return parse_number(schema, integer=True)

    if schema_type == "boolean":
        return BooleanNode()

    if schema_type == "null":
        return NullNode()

    if schema_type == "object":
        return parse_object(schema, ctx)

    if schema_type == "array":
        return parse_array(schema, ctx)

    # Unknown type - return AnyNode
    return AnyNode()


def _infer_type(schema: Dict[str, Any]) -> Optional[str]:
    """Infer the type from other keywords in the schema.

    JSON Schema allows omitting "type" when other keywords unambiguously
    indicate the type. This function checks for type-specific keywords
    and returns the implied type.

    Returns:
        Inferred type string or None if type cannot be determined
    """
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


def _detect_discriminator(variants: List[Any]) -> Optional[str]:
    """Detect if oneOf variants form a discriminated union.

    A discriminated union has all variants as objects with a common property
    that has a literal (const) value unique to each variant. This is useful
    for generating more efficient union types in target languages.

    Args:
        variants: List of oneOf variant schemas

    Returns:
        Property name that discriminates the union, or None if not discriminated

    Example:
        >>> variants = [
        ...     {"type": "object", "properties": {"kind": {"const": "a"}}},
        ...     {"type": "object", "properties": {"kind": {"const": "b"}}}
        ... ]
        >>> _detect_discriminator(variants)
        'kind'
    """
    if not variants:
        return None

    # Find common properties with const values across all variants
    common_discriminators: Optional[Set[str]] = None

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


def _detect_type_guards(schema: Dict[str, Any], ctx: ParseContext) -> List[TypeGuard]:
    """Detect type-specific constraints and create type guards.

    For schemas without an explicit type that have type-specific keywords,
    create guards that apply those constraints only to matching runtime types.

    This handles schemas like:
        {"minLength": 5, "minimum": 0}

    Which should validate:
        - Strings must have minLength >= 5
        - Numbers must be >= 0
        - Other types pass through (JSON Schema allows this)

    Args:
        schema: Schema dict to analyze
        ctx: Parse context for recursive parsing

    Returns:
        List of TypeGuard objects, one per detected type
    """
    guards = []

    # Check for string-specific keywords
    if any(k in schema for k in ("minLength", "maxLength", "pattern", "format")):
        string_schema = parse_string(schema)
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
        number_schema = parse_number(schema, integer=False)
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
        object_schema = parse_object(schema, ctx)
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
        # Use parse_tuple for prefixItems, parse_array for other array keywords
        if "prefixItems" in schema:
            array_schema = parse_tuple(schema, ctx)
        else:
            array_schema = parse_array(schema, ctx)
        guards.append(TypeGuard(check="array", schema=array_schema))

    return guards
