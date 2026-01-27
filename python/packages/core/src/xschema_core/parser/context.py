"""Parse context for tracking state during JSON Schema parsing.

The parse context is responsible for:
1. Holding the root schema for resolving internal $ref pointers
2. Tracking which refs are currently being resolved (cycle detection)
3. Providing a clean interface for the recursive parse functions

Why Cycle Detection Matters:
---------------------------
JSON Schema allows recursive schemas via $ref. For example:

    {
        "$defs": {
            "Node": {
                "type": "object",
                "properties": {
                    "children": {"type": "array", "items": {"$ref": "#/$defs/Node"}}
                }
            }
        },
        "$ref": "#/$defs/Node"
    }

Without cycle detection, the parser would infinitely recurse. When we encounter
a ref that's already being resolved (it's in the `resolving` set), we return a
RefNode with `resolved=None` to break the cycle. The renderer can then handle
this by generating a ForwardRef or similar construct.

Root Schema:
-----------
The root schema is kept as a reference to resolve JSON pointers like #/$defs/Name.
All schemas processed by the parser are expected to be pre-bundled by the CLI -
external $refs (URLs, file paths) should already be resolved and inlined. Only
internal JSON pointer refs are supported at parse time.
"""

from dataclasses import dataclass, field
from typing import Any, Dict, Optional, Set, Union


@dataclass
class ParseContext:
    """Context for parsing, holds root schema for internal $ref resolution.

    Attributes:
        root: The root schema dict, used for resolving JSON pointer refs (#/$defs/...).
              Should be None only for boolean schemas.
        resolving: Set of $ref paths currently being resolved. Used to detect and
                   break reference cycles. When a ref is encountered that's already
                   in this set, the parser returns RefNode(resolved=None) to indicate
                   a cycle.

    Example:
        >>> ctx = ParseContext(root=schema)
        >>> ctx.resolving.add("#/$defs/User")
        >>> "#/$defs/User" in ctx.resolving
        True
        >>> ctx.resolving.discard("#/$defs/User")
    """

    root: Optional[Dict[str, Any]] = None
    resolving: Set[str] = field(default_factory=set)


def create_context(root: Union[Dict[str, Any], bool, None] = None) -> ParseContext:
    """Create a new parse context from a root schema.

    Args:
        root: The root JSON Schema. Can be a dict (the common case), a bool
              (true = any, false = never), or None.

    Returns:
        A ParseContext with the root schema stored (if it's a dict) and an
        empty resolving set.

    Example:
        >>> schema = {"type": "object", "properties": {"name": {"type": "string"}}}
        >>> ctx = create_context(schema)
        >>> ctx.root is schema
        True
    """
    if isinstance(root, dict):
        return ParseContext(root=root)
    return ParseContext(root=None)
