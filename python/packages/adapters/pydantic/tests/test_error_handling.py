"""Tests for error handling and edge cases."""

import json
import subprocess
import sys
from pathlib import Path

import pytest

from xschema_pydantic.converter import convert
from xschema_pydantic.errors import (
    ConversionError,
    InvalidSchemaError,
    XSchemaError,
)


class TestCustomExceptions:
    """Test custom exception classes."""

    def test_xschema_error_without_path(self):
        """XSchemaError without schema path shows just message."""
        error = XSchemaError("Something went wrong")
        assert str(error) == "Something went wrong"
        assert error.schema_path is None

    def test_xschema_error_with_path(self):
        """XSchemaError with schema path includes path in message."""
        error = XSchemaError("Something went wrong", schema_path="$.user.User")
        assert str(error) == "Something went wrong (at $.user.User)"
        assert error.schema_path == "$.user.User"

    def test_invalid_schema_error(self):
        """InvalidSchemaError is a subclass of XSchemaError."""
        error = InvalidSchemaError("Invalid schema", schema_path="$.foo")
        assert isinstance(error, XSchemaError)
        assert str(error) == "Invalid schema (at $.foo)"

    def test_conversion_error(self):
        """ConversionError is a subclass of XSchemaError."""
        error = ConversionError("Conversion failed", schema_path="$.bar")
        assert isinstance(error, XSchemaError)
        assert str(error) == "Conversion failed (at $.bar)"


class TestConverterValidation:
    """Test converter input validation."""

    def test_missing_id(self):
        """Missing 'id' field raises InvalidSchemaError."""
        with pytest.raises(InvalidSchemaError, match="Missing required field 'id'"):
            convert({"namespace": "test", "varName": "foo", "schema": {}})

    def test_missing_var_name(self):
        """Missing 'varName' field raises InvalidSchemaError."""
        with pytest.raises(
            InvalidSchemaError, match="Missing required field 'varName'"
        ):
            convert({"namespace": "test", "id": "Foo", "schema": {}})

    def test_missing_schema(self):
        """Missing 'schema' field raises InvalidSchemaError."""
        with pytest.raises(InvalidSchemaError, match="Missing required field 'schema'"):
            convert({"namespace": "test", "id": "Foo", "varName": "foo"})

    def test_schema_wrong_type(self):
        """Schema must be dict or bool, not other types."""
        with pytest.raises(InvalidSchemaError, match="Schema must be a dict or bool"):
            convert(
                {
                    "namespace": "test",
                    "id": "Foo",
                    "varName": "foo",
                    "schema": "not a schema",
                }
            )

        with pytest.raises(InvalidSchemaError, match="Schema must be a dict or bool"):
            convert(
                {
                    "namespace": "test",
                    "id": "Foo",
                    "varName": "foo",
                    "schema": 123,
                }
            )

        with pytest.raises(InvalidSchemaError, match="Schema must be a dict or bool"):
            convert(
                {
                    "namespace": "test",
                    "id": "Foo",
                    "varName": "foo",
                    "schema": ["array"],
                }
            )

    def test_schema_path_includes_namespace(self):
        """Schema path includes namespace when present."""
        with pytest.raises(InvalidSchemaError) as exc_info:
            convert(
                {
                    "namespace": "users",
                    "id": "User",
                    "varName": "user",
                    "schema": "invalid",
                }
            )
        assert exc_info.value.schema_path == "$.users.User"

    def test_schema_path_without_namespace(self):
        """Schema path works without namespace."""
        with pytest.raises(InvalidSchemaError) as exc_info:
            convert(
                {
                    "namespace": "",
                    "id": "User",
                    "varName": "user",
                    "schema": "invalid",
                }
            )
        assert exc_info.value.schema_path == "$.User"

    def test_unexpected_ref_in_schema(self):
        """$ref in schema raises InvalidSchemaError (should be pre-bundled)."""
        with pytest.raises(InvalidSchemaError, match="Unexpected \\$ref"):
            convert(
                {
                    "namespace": "test",
                    "id": "Foo",
                    "varName": "foo",
                    "schema": {"$ref": "#/$defs/Something"},
                }
            )


class TestCLIErrorHandling:
    """Test CLI error handling via subprocess."""

    def run_cli(self, input_data: str) -> tuple[int, str, str]:
        """Run the CLI with given input and return exit code, stdout, stderr.

        Args:
            input_data: JSON string to pass via stdin

        Returns:
            Tuple of (exit_code, stdout, stderr)
        """
        # Get the package directory to run 'uv run xschema-pydantic'
        adapter_dir = Path(__file__).parent.parent
        result = subprocess.run(
            ["uv", "run", "xschema-pydantic"],
            input=input_data,
            capture_output=True,
            text=True,
            cwd=adapter_dir,
        )
        return result.returncode, result.stdout, result.stderr

    def test_invalid_json(self):
        """Invalid JSON input returns exit code 1 with error."""
        code, stdout, stderr = self.run_cli("{not valid json")
        assert code == 1
        assert stdout == ""

        error = json.loads(stderr)
        assert "Invalid JSON input" in error["error"]
        assert error["type"] == "JSONDecodeError"
        assert "line" in error

    def test_not_array(self):
        """Input must be an array, not object or primitive."""
        code, stdout, stderr = self.run_cli('{"not": "array"}')
        assert code == 1
        assert stdout == ""

        error = json.loads(stderr)
        assert "Expected JSON array" in error["error"]
        assert error["type"] == "ValueError"

    def test_array_with_non_dict(self):
        """Array items must be dicts."""
        code, stdout, stderr = self.run_cli('["string", 123]')
        assert code == 1
        assert stdout == ""

        error = json.loads(stderr)
        assert "not a dict" in error["error"]
        assert error["index"] == 0

    def test_missing_required_field(self):
        """Missing required field returns exit code 2 (schema validation error)."""
        code, stdout, stderr = self.run_cli(
            '[{"namespace": "test", "varName": "foo", "schema": {}}]'
        )
        assert code == 2
        assert stdout == ""

        error = json.loads(stderr)
        assert "Missing required field 'id'" in error["error"]
        assert error["type"] == "InvalidSchemaError"

    def test_invalid_schema_type(self):
        """Invalid schema type returns exit code 2."""
        code, stdout, stderr = self.run_cli(
            '[{"namespace": "test", "id": "Foo", "varName": "foo", "schema": "invalid"}]'
        )
        assert code == 2
        assert stdout == ""

        error = json.loads(stderr)
        assert "Schema must be a dict or bool" in error["error"]
        assert error["schemaId"] == "Foo"

    def test_unexpected_ref(self):
        """Unexpected $ref returns exit code 2."""
        code, stdout, stderr = self.run_cli(
            '[{"namespace": "test", "id": "Foo", "varName": "foo", "schema": {"$ref": "#/defs/X"}}]'
        )
        assert code == 2
        assert stdout == ""

        error = json.loads(stderr)
        assert "Unexpected $ref" in error["error"]
        assert error["schemaId"] == "Foo"
        assert "schemaPath" in error

    def test_valid_input_succeeds(self):
        """Valid input returns exit code 0 with results."""
        code, stdout, stderr = self.run_cli(
            '[{"namespace": "test", "id": "Str", "varName": "test_str", "schema": {"type": "string"}}]'
        )
        assert code == 0
        assert stderr == ""

        results = json.loads(stdout)
        assert len(results) == 1
        assert results[0]["namespace"] == "test"
        assert results[0]["id"] == "Str"
        assert results[0]["varName"] == "test_str"
        assert "imports" in results[0]
        assert "schema" in results[0]
        assert "type" in results[0]

    def test_multiple_schemas_first_fails(self):
        """When first schema fails, should exit immediately."""
        code, stdout, stderr = self.run_cli(
            """[
                {"namespace": "test", "id": "Bad", "varName": "bad", "schema": "invalid"},
                {"namespace": "test", "id": "Good", "varName": "good", "schema": {"type": "string"}}
            ]"""
        )
        assert code == 2
        assert stdout == ""

        error = json.loads(stderr)
        assert error["schemaId"] == "Bad"

    def test_error_includes_schema_path(self):
        """Error messages include schema path for debugging."""
        code, stdout, stderr = self.run_cli(
            '[{"namespace": "users", "id": "User", "varName": "user", "schema": {"$ref": "#/defs/X"}}]'
        )
        assert code == 2

        error = json.loads(stderr)
        assert error["schemaPath"] == "$.users.User"


class TestEdgeCases:
    """Test edge cases and corner scenarios."""

    def test_empty_schema(self):
        """Empty schema {} is valid and means 'any'."""
        result = convert(
            {
                "namespace": "test",
                "id": "Any",
                "varName": "test_any",
                "schema": {},
            }
        )
        assert result["type"] == "Any"

    def test_boolean_schema_true(self):
        """Boolean schema 'true' means 'any'."""
        result = convert(
            {
                "namespace": "test",
                "id": "Any",
                "varName": "test_any",
                "schema": True,
            }
        )
        assert result["type"] == "Any"

    def test_boolean_schema_false(self):
        """Boolean schema 'false' means 'never'."""
        result = convert(
            {
                "namespace": "test",
                "id": "Never",
                "varName": "test_never",
                "schema": False,
            }
        )
        # Never type uses a validator that always rejects
        assert (
            "never" in result["schema"].lower() or "BeforeValidator" in result["type"]
        )

    def test_empty_namespace(self):
        """Empty namespace is allowed."""
        result = convert(
            {
                "namespace": "",
                "id": "Foo",
                "varName": "foo",
                "schema": {"type": "string"},
            }
        )
        assert result["namespace"] == ""

    def test_complex_var_name(self):
        """Complex varName gets converted to PascalCase for classes."""
        result = convert(
            {
                "namespace": "test",
                "id": "Obj",
                "varName": "my_complex_schema_name",
                "schema": {"type": "object", "properties": {}},
            }
        )
        assert "class MyComplexSchemaName" in result["schema"]

    def test_special_characters_in_id(self):
        """Schema with special characters in ID should work."""
        result = convert(
            {
                "namespace": "test",
                "id": "User-Profile_v2",
                "varName": "user_profile",
                "schema": {"type": "string"},
            }
        )
        assert result["id"] == "User-Profile_v2"
