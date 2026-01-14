"""Tests for xschema_client."""

import pytest
from xschema_client import XSchemaError, create_xschema_client


class MockValidator:
    """Mock validator for testing."""

    def __init__(self, name: str):
        self.name = name

    def model_validate(self, data):
        return data


def test_full_key_lookup():
    """Test looking up schemas with full namespace:id keys."""
    schemas = {
        "user:Profile": MockValidator("Profile"),
        "another:TSConfig": MockValidator("TSConfig"),
    }
    xschema = create_xschema_client(schemas)

    # Should find by full key
    assert xschema("user:Profile").name == "Profile"
    assert xschema("another:TSConfig").name == "TSConfig"


def test_shorthand_with_default_namespace():
    """Test shorthand lookup when default_namespace is set."""
    schemas = {
        "user:Profile": MockValidator("Profile"),
        "user:Calendar": MockValidator("Calendar"),
        "another:TSConfig": MockValidator("TSConfig"),
    }
    xschema = create_xschema_client(schemas, default_namespace="user")

    # Shorthand should prepend default namespace
    assert xschema("Profile").name == "Profile"
    assert xschema("Calendar").name == "Calendar"

    # Full keys should still work
    assert xschema("user:Profile").name == "Profile"
    assert xschema("another:TSConfig").name == "TSConfig"


def test_shorthand_without_default_namespace():
    """Test that shorthand fails without default_namespace."""
    schemas = {
        "user:Profile": MockValidator("Profile"),
    }
    xschema = create_xschema_client(schemas)

    # Shorthand without default namespace should fail
    with pytest.raises(XSchemaError) as exc_info:
        xschema("Profile")

    assert "Unknown schema: Profile" in str(exc_info.value)
    assert "xschema generate" in str(exc_info.value)


def test_schema_not_found():
    """Test error when schema doesn't exist."""
    schemas = {
        "user:Profile": MockValidator("Profile"),
    }
    xschema = create_xschema_client(schemas, default_namespace="user")

    # Non-existent full key
    with pytest.raises(XSchemaError) as exc_info:
        xschema("user:NonExistent")
    assert "Unknown schema: user:NonExistent" in str(exc_info.value)

    # Non-existent shorthand
    with pytest.raises(XSchemaError) as exc_info:
        xschema("NonExistent")
    assert "Unknown schema: user:NonExistent" in str(exc_info.value)


def test_empty_schemas():
    """Test client with empty schemas dict."""
    xschema = create_xschema_client({})

    with pytest.raises(XSchemaError):
        xschema("anything")


def test_type_inference():
    """Test that client preserves type information."""
    schemas = {
        "user:Profile": MockValidator("Profile"),
    }
    xschema = create_xschema_client(schemas)

    # Should return the exact validator instance
    validator = xschema("user:Profile")
    assert isinstance(validator, MockValidator)
    assert validator.name == "Profile"


def test_colon_in_key_always_uses_full_key():
    """Test that any key with ':' is treated as full key, even with default_namespace."""
    schemas = {
        "user:Profile": MockValidator("Profile"),
        "another:Config": MockValidator("Config"),
    }
    xschema = create_xschema_client(schemas, default_namespace="user")

    # Key with ':' should never prepend default namespace
    assert xschema("another:Config").name == "Config"

    # Even if the namespace doesn't exist
    with pytest.raises(XSchemaError) as exc_info:
        xschema("nonexistent:Schema")
    assert "Unknown schema: nonexistent:Schema" in str(exc_info.value)


def test_client_is_callable():
    """Test that returned client is callable."""
    schemas = {"user:Profile": MockValidator("Profile")}
    xschema = create_xschema_client(schemas)

    # Should be callable
    assert callable(xschema)

    # Should work like a function
    result = xschema("user:Profile")
    assert result.name == "Profile"


def test_multiple_clients():
    """Test creating multiple independent clients."""
    schemas1 = {"user:Profile": MockValidator("Profile1")}
    schemas2 = {"user:Profile": MockValidator("Profile2")}

    xschema1 = create_xschema_client(schemas1)
    xschema2 = create_xschema_client(schemas2)

    # Should be independent
    assert xschema1("user:Profile").name == "Profile1"
    assert xschema2("user:Profile").name == "Profile2"
