"""JSON Schema to IR parser.

This package converts JSON Schema objects into strongly-typed Intermediate Representation (IR) nodes.
The IR is designed to be adapter-agnostic, capturing JSON Schema semantics in a way that adapters
(Pydantic, etc.) can render to target validation libraries.

Public API:
- parse: Main entry point - converts JSON Schema dict to SchemaNode
- ParseContext, create_context: Context for parsing with cycle detection
- parse_string, parse_number: Primitive parsers (exported for testing/internal use)
- parse_object, parse_array, parse_tuple, parse_legacy_tuple: Collection parsers
- parse_all_of, parse_any_of, parse_one_of, parse_not, parse_conditional: Composition parsers
- parse_literal, parse_enum: Value parsers

Internal structure:
- primitives: String, number, boolean, null type parsing
- collections: Object, array, tuple type parsing
- composition: allOf, anyOf, oneOf, not, if/then/else parsing
- values: const, enum parsing
- context: Parse context for cycle detection and ref resolution
"""

# Import from temporary location - will be moved to parser/core.py in future task
from xschema_core.parser_old import parse  # noqa: F401
from xschema_core.parser.context import ParseContext, create_context
from xschema_core.parser.primitives import parse_number, parse_string
from xschema_core.parser.collections import (
    parse_array,
    parse_legacy_tuple,
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

__all__ = [
    "parse",
    # Context
    "ParseContext",
    "create_context",
    # Primitive parsers
    "parse_string",
    "parse_number",
    # Collection parsers
    "parse_object",
    "parse_array",
    "parse_tuple",
    "parse_legacy_tuple",
    # Composition parsers
    "parse_all_of",
    "parse_any_of",
    "parse_one_of",
    "parse_not",
    "parse_conditional",
    # Value parsers
    "parse_literal",
    "parse_enum",
]
