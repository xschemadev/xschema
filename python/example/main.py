"""
XSchema Example - Python

This example shows how to use xschema to:
1. Define JSON Schema sources in config files (*.jsonc)
2. Generate type-safe validators (Pydantic)
3. Use the generated schemas with full type inference

Run `uv run generate` to regenerate schemas, then `uv run python main.py` to run this file.
"""

from xschema_client import create_xschema_client

# Import generated schemas
from _xschema import schemas

# Create the xschema client with generated schemas
# defaultNamespace allows shorthand lookups: xschema("Calendar") instead of xschema("user:Calendar")
xschema = create_xschema_client(schemas, default_namespace="user")

# ============================================
# Type extraction using type annotations
# ============================================

# Use full "namespace:id" to get schemas from any namespace
tsconfig_schema = xschema("another:TSConfig")
calendar_schema = xschema("user:Calendar")

# When defaultNamespace is set, you can omit it for that namespace
calendar = xschema("Calendar")  # Same as xschema("user:Calendar")

# Validate with Pydantic TypeAdapter
valid_calendar = calendar.validate_python(
    {
        "dtstart": "2024-01-01",
        "summary": "New Year's Day",
    }
)

print(f"Valid calendar event: {valid_calendar.summary}")

# Try invalid data
try:
    calendar.validate_python({"invalid": "data"})
except Exception as e:
    print(f"Validation failed as expected: {type(e).__name__}")

# ============================================
# User schema
# ============================================

user = xschema("User")
valid_user = user.validate_python(
    {
        "id": "123",
        "email": "alice@example.com",
        "name": "Alice",
        "age": 30,
    }
)

print(f"Valid user: {valid_user.name} ({valid_user.email})")

# ============================================
# Native schema (inline JSON Schema)
# ============================================

native = xschema("another:Native")
valid_native = native.validate_python(
    {
        "fruits": ["apple", "orange"],
        "vegetables": [
            {"veggieName": "carrot", "veggieLike": True},
            {"veggieName": "broccoli", "veggieLike": False},
        ],
    }
)

print(f"Valid native: {valid_native.fruits}")
