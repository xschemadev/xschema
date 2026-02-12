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


def test_schema_not_found():
    """Test error when schema doesn't exist."""
    schemas = {
        "user:Profile": MockValidator("Profile"),
    }
    xschema = create_xschema_client(schemas)

    # Non-existent key
    with pytest.raises(XSchemaError) as exc_info:
        xschema("user:NonExistent")
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


def test_nonexistent_namespace():
    """Test error when namespace doesn't exist."""
    schemas = {
        "user:Profile": MockValidator("Profile"),
        "another:Config": MockValidator("Config"),
    }
    xschema = create_xschema_client(schemas)

    # Non-existent namespace should fail
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
