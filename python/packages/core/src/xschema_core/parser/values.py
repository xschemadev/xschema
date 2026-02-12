"""Parsers for value types: literal (const), enum.

This module handles JSON Schema keywords that constrain to specific values:

CONST (`const` keyword):
    Validates that a value is exactly equal to the const value.
    Uses JSON equality semantics (see below).

    JSON Schema: {"const": "hello"}
    Matches: "hello"
    Rejects: "Hello", "hello ", 123

ENUM (`enum` keyword):
    Validates that a value is one of the allowed values.
    Uses JSON equality semantics (see below).

    JSON Schema: {"enum": ["red", "green", "blue"]}
    Matches: "red", "green", "blue"
    Rejects: "Red", "yellow", 123


JSON Schema Equality Semantics:
-------------------------------
JSON Schema uses JSON-native equality which differs from Python's truthiness:

1. TYPE MATTERS: false != 0, true != 1, null != 0
   - {"const": false} rejects 0 (different types)
   - {"const": 0} rejects false (different types)
   - This is different from Python where False == 0

2. DEEP EQUALITY: Arrays and objects compare structurally
   - {"const": [1, 2, 3]} validates [1, 2, 3] exactly
   - {"const": {"a": 1}} validates {"a": 1} exactly

3. NUMBER EQUALITY: 1 == 1.0 (JSON doesn't distinguish int/float)
   - {"const": 1} validates both 1 and 1.0

Why This Matters for Code Generation:
------------------------------------
Renderers must be careful when generating validation code:
- Python `False == 0` is True, but JSON Schema `false != 0`
- Must use type-aware equality checks, not Python's ==
- See renderer's `_json_equals` helper for implementation

Example edge case:
    {"enum": [false, 0]}  # Both are valid, distinct values
    - false matches false (not 0)
    - 0 matches 0 (not false)
"""

from typing import Any

from xschema_core.ir import EnumNode, LiteralNode


def parse_literal(schema: dict[str, Any]) -> LiteralNode:
    """Parse a const schema into a LiteralNode.

    The `const` keyword declares that a value must be exactly equal to
    the specified value using JSON equality semantics.

    Args:
        schema: JSON Schema dict containing `const` keyword

    Returns:
        LiteralNode containing the const value

    Example:
        >>> parse_literal({"const": "hello"})
        LiteralNode(kind='literal', value='hello')

        >>> parse_literal({"const": {"type": "user", "id": 123}})
        LiteralNode(kind='literal', value={'type': 'user', 'id': 123})

    Note:
        The value can be ANY valid JSON value: string, number, boolean,
        null, array, or object. Renderers must handle all cases.
    """
    return LiteralNode(value=schema["const"])


def parse_enum(schema: dict[str, Any]) -> EnumNode:
    """Parse an enum schema into an EnumNode.

    The `enum` keyword declares that a value must be one of the specified
    values using JSON equality semantics.

    Args:
        schema: JSON Schema dict containing `enum` keyword

    Returns:
        EnumNode containing tuple of allowed values

    Example:
        >>> parse_enum({"enum": ["draft", "published", "archived"]})
        EnumNode(kind='enum', values=('draft', 'published', 'archived'))

        >>> parse_enum({"enum": [1, 2, 3, null]})
        EnumNode(kind='enum', values=(1, 2, 3, None))

    Note:
        - Values are stored as tuple (immutable) not list
        - Values can be mixed types: {"enum": [1, "one", true, null]}
        - Duplicate values are preserved (schema author's responsibility)
        - Renderers must use JSON equality, not Python equality
    """
    return EnumNode(values=tuple(schema["enum"]))
