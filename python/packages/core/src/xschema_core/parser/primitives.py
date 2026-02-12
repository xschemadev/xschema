"""Parsers for primitive types: string, number, boolean, null.

This module handles the parsing of JSON Schema primitive type constraints.
Each function converts JSON Schema keywords into strongly-typed IR nodes.

Key semantics:
- String constraints: minLength, maxLength, pattern, format
- Number constraints: minimum, maximum, exclusiveMinimum, exclusiveMaximum, multipleOf
- Draft compatibility: exclusiveMin/Max can be boolean (draft4) or number (draft6+)
  - Draft 4: {"minimum": 0, "exclusiveMinimum": true} means > 0
  - Draft 6+: {"exclusiveMinimum": 0} means > 0
  - This parser handles both styles, normalizing to IR format
"""

from typing import Any

from xschema_core.ir import NumberConstraints, NumberNode, StringConstraints, StringNode


def parse_string(schema: dict[str, Any]) -> StringNode:
    """Parse a string schema into a StringNode.

    Extracts string-specific constraints from JSON Schema keywords:
    - minLength: minimum character length (inclusive)
    - maxLength: maximum character length (inclusive)
    - pattern: ECMA-262 regular expression
    - format: semantic format hint (email, uri, date-time, etc.)

    Args:
        schema: JSON Schema object with type="string" or inferred string type

    Returns:
        StringNode with extracted constraints

    Example:
        >>> parse_string({"minLength": 1, "maxLength": 100, "pattern": "^[a-z]+$"})
        StringNode(constraints=StringConstraints(min_length=1, max_length=100, pattern="^[a-z]+$"), format=None)
    """
    constraints = StringConstraints(
        min_length=schema.get("minLength"),
        max_length=schema.get("maxLength"),
        pattern=schema.get("pattern"),
    )
    return StringNode(
        constraints=constraints,
        format=schema.get("format"),
    )


def parse_number(schema: dict[str, Any], integer: bool) -> NumberNode:
    """Parse a number/integer schema into a NumberNode.

    Handles both draft4 and draft6+ styles for exclusive bounds:
    - Draft 4: exclusiveMinimum/Maximum are booleans, values in minimum/maximum
    - Draft 6+: exclusiveMinimum/Maximum are numbers directly

    Number constraints:
    - minimum: inclusive lower bound
    - maximum: inclusive upper bound
    - exclusiveMinimum: exclusive lower bound (value must be strictly greater)
    - exclusiveMaximum: exclusive upper bound (value must be strictly less)
    - multipleOf: value must be a multiple of this number

    Args:
        schema: JSON Schema object with type="number" or type="integer"
        integer: True if this is an integer schema, False for number

    Returns:
        NumberNode with normalized constraints

    Example (draft6+):
        >>> parse_number({"exclusiveMinimum": 0, "maximum": 100}, integer=False)
        NumberNode(constraints=NumberConstraints(minimum=None, maximum=100, exclusive_minimum=0, ...), integer=False)

    Example (draft4):
        >>> parse_number({"minimum": 0, "exclusiveMinimum": true, "maximum": 100}, integer=True)
        NumberNode(constraints=NumberConstraints(minimum=None, maximum=100, exclusive_minimum=0, ...), integer=True)
    """
    # Handle exclusiveMinimum/exclusiveMaximum which can be boolean (draft4) or number (draft6+)
    exclusive_min = schema.get("exclusiveMinimum")
    exclusive_max = schema.get("exclusiveMaximum")

    # Draft 4 style: exclusiveMinimum is boolean, actual value is in minimum
    # Convert: {"minimum": 0, "exclusiveMinimum": true} -> exclusive_minimum=0, minimum=None
    if exclusive_min is True:
        exclusive_min = schema.get("minimum")
    elif exclusive_min is False:
        exclusive_min = None

    if exclusive_max is True:
        exclusive_max = schema.get("maximum")
    elif exclusive_max is False:
        exclusive_max = None

    # Build constraints, handling draft4 vs draft6+ differences
    constraints = NumberConstraints(
        # If exclusiveMinimum is boolean and true, ignore minimum (it's the exclusive bound)
        minimum=schema.get("minimum")
        if not isinstance(schema.get("exclusiveMinimum"), bool)
        or not schema.get("exclusiveMinimum")
        else None,
        # If exclusiveMaximum is boolean and true, ignore maximum (it's the exclusive bound)
        maximum=schema.get("maximum")
        if not isinstance(schema.get("exclusiveMaximum"), bool)
        or not schema.get("exclusiveMaximum")
        else None,
        # Normalized exclusive bounds (either from draft6+ number or draft4 conversion)
        exclusive_minimum=exclusive_min
        if not isinstance(exclusive_min, bool)
        else None,
        exclusive_maximum=exclusive_max
        if not isinstance(exclusive_max, bool)
        else None,
        multiple_of=schema.get("multipleOf"),
    )

    return NumberNode(
        constraints=constraints,
        integer=integer,
    )
