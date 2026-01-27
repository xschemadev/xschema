"""JSON Schema to IR parser."""

from typing import Any, Literal

from xschema_core.ir import (
    AnyNode,
    BooleanNode,
    EnumNode,
    IntersectionNode,
    LiteralNode,
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


class _ParseContext:
    """Context for parsing, holds root schema for internal $ref resolution."""

    def __init__(self, root: dict[str, Any] | None = None):
        self.root = root
        # Track refs being resolved to detect cycles
        self.resolving: set[str] = set()


def parse(schema: dict[str, Any] | bool) -> SchemaNode:
    """Parse a JSON Schema dict into an IR SchemaNode.

    Expects schemas to be pre-bundled by the CLI - all external $refs
    should be resolved. Internal JSON pointer refs (#/$defs/Name) are supported.
    """
    # Create context with root schema for internal $ref resolution
    ctx = _ParseContext(root=schema if isinstance(schema, dict) else None)
    return _parse_with_ctx(schema, ctx)


def _resolve_json_pointer(ref: str, root: dict[str, Any]) -> Any:
    """Resolve a JSON pointer (fragment starting with #/) to a schema node."""
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


def _resolve_ref(ref: str, ctx: _ParseContext) -> SchemaNode:
    """Resolve a $ref to a SchemaNode.

    Only handles internal JSON pointer refs (e.g., #/$defs/Name, #/properties/foo).
    External refs should be bundled by the CLI before reaching the adapter.
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

    # Check for cycles
    if ref in ctx.resolving:
        # Return a RefNode for cyclic references
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


def _parse_with_ctx(schema: dict[str, Any] | bool, ctx: _ParseContext) -> SchemaNode:
    """Parse a JSON Schema dict into an IR SchemaNode with context."""
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
        return LiteralNode(value=schema["const"])

    # Handle enum
    if "enum" in schema:
        return EnumNode(values=tuple(schema["enum"]))

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


def _detect_type_guards(schema: dict[str, Any], ctx: _ParseContext) -> list[TypeGuard]:
    """Detect type-specific constraints and create type guards.

    For schemas without an explicit type that have type-specific keywords,
    create guards that apply those constraints only to matching runtime types.
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
