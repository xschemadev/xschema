"""XSchema example for Python runtime client usage."""

import _xschema
from xschema_client import create_xschema_client

xschema: "_xschema.XSchemaClient" = create_xschema_client(_xschema.schemas)

# ============================================
# Type extraction using type annotations
# ============================================

# Use "namespace:id" format to get schemas
tsconfig_schema = xschema("another:TSConfig")
calendar = xschema("user:Calendar")

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

user = xschema("user:User")
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
