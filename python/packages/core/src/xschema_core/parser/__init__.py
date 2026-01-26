"""JSON Schema to IR parser.

This package converts JSON Schema objects into strongly-typed Intermediate Representation (IR) nodes.
The IR is designed to be adapter-agnostic, capturing JSON Schema semantics in a way that adapters
(Pydantic, etc.) can render to target validation libraries.

Public API:
- parse: Main entry point - converts JSON Schema dict to SchemaNode
- parse_string, parse_number: Primitive parsers (exported for testing/internal use)

Internal structure:
- primitives: String, number, boolean, null type parsing
- collections: Object, array, tuple type parsing (TODO)
- composition: allOf, anyOf, oneOf, not parsing (TODO)
- values: const, enum parsing (TODO)
- context: Parse context for cycle detection and ref resolution (TODO)
"""

# Import from temporary location - will be moved to parser/core.py in future task
from xschema_core.parser_old import parse  # noqa: F401
from xschema_core.parser.primitives import parse_number, parse_string

__all__ = ["parse", "parse_string", "parse_number"]
