"""Integration tests for xschema-pydantic adapter."""

import json
import subprocess
import sys
import tempfile
from importlib.util import module_from_spec, spec_from_file_location
from pathlib import Path


def test_end_to_end_integration():
    """Test full pipeline: stdin -> adapter -> stdout -> generated code -> validation."""

    # Test fixtures with various schema types
    test_input = [
        {
            "namespace": "test",
            "id": "User",
            "varName": "test_user",
            "schema": {
                "type": "object",
                "properties": {
                    "name": {"type": "string", "minLength": 1},
                    "age": {"type": "integer", "minimum": 0, "maximum": 150},
                    "email": {"type": "string", "minLength": 3},
                },
                "required": ["name", "email"],
                "additionalProperties": False,
            },
        },
        {
            "namespace": "test",
            "id": "Product",
            "varName": "test_product",
            "schema": {
                "type": "object",
                "properties": {
                    "id": {"type": "string", "format": "uuid"},
                    "price": {"type": "number", "minimum": 0},
                    "tags": {
                        "type": "array",
                        "items": {"type": "string"},
                        "minItems": 1,
                        "uniqueItems": True,
                    },
                },
                "required": ["id", "price"],
            },
        },
        {
            "namespace": "test",
            "id": "Status",
            "varName": "test_status",
            "schema": {"enum": ["active", "inactive", "pending"]},
        },
        {
            "namespace": "test",
            "id": "Count",
            "varName": "test_count",
            "schema": {"type": "integer", "minimum": 0},
        },
        {
            "namespace": "test",
            "id": "StringArray",
            "varName": "test_string_array",
            "schema": {"type": "array", "items": {"type": "string"}},
        },
    ]

    # Run adapter via stdin/stdout
    input_json = json.dumps(test_input)
    result = subprocess.run(
        ["uv", "run", "xschema-pydantic"],
        input=input_json,
        capture_output=True,
        text=True,
        cwd=Path(__file__).parent.parent,
    )

    # Verify subprocess succeeded
    assert result.returncode == 0, f"Adapter failed: {result.stderr}"

    # Parse output JSON
    try:
        output = json.loads(result.stdout)
    except json.JSONDecodeError as e:
        raise AssertionError(f"Invalid JSON output: {result.stdout}") from e

    # Verify output structure
    assert isinstance(output, list), "Output should be a list"
    assert len(output) == 5, f"Expected 5 results, got {len(output)}"

    # Verify each result has required fields
    for i, item in enumerate(output):
        assert "namespace" in item, f"Result {i} missing 'namespace'"
        assert "id" in item, f"Result {i} missing 'id'"
        assert "varName" in item, f"Result {i} missing 'varName'"
        assert "imports" in item, f"Result {i} missing 'imports'"
        assert "schema" in item, f"Result {i} missing 'schema'"
        assert "type" in item, f"Result {i} missing 'type'"

        assert item["namespace"] == "test", f"Result {i} has wrong namespace"
        assert isinstance(item["imports"], list), f"Result {i} imports should be list"
        assert isinstance(item["schema"], str), f"Result {i} schema should be string"
        assert isinstance(item["type"], str), f"Result {i} type should be string"

    # Build generated module code
    generated_code_lines = []

    # Collect all imports (deduplicated)
    all_imports = set()
    for item in output:
        all_imports.update(item["imports"])

    # Add imports
    for imp in sorted(all_imports):
        generated_code_lines.append(imp)

    generated_code_lines.append("")  # blank line after imports

    # Add schemas
    for item in output:
        generated_code_lines.append(f"# {item['id']}")
        generated_code_lines.append(item["schema"])
        generated_code_lines.append("")

    generated_code = "\n".join(generated_code_lines)

    # Write generated code to temp file and import
    with tempfile.TemporaryDirectory() as tmpdir:
        module_path = Path(tmpdir) / "generated_schemas.py"
        module_path.write_text(generated_code)

        # Import the generated module
        spec = spec_from_file_location("generated_schemas", module_path)
        assert spec is not None, "Failed to create module spec"
        assert spec.loader is not None, "Module spec has no loader"

        module = module_from_spec(spec)
        sys.modules["generated_schemas"] = module
        spec.loader.exec_module(module)

        # Test User schema (class name is PascalCase of varName)
        TestUser = getattr(module, "TestUser")

        # Valid user data
        valid_user = {"name": "John Doe", "email": "john@example.com", "age": 30}
        user_instance = TestUser(**valid_user)
        assert user_instance.name == "John Doe"
        assert user_instance.email == "john@example.com"
        assert user_instance.age == 30

        # Invalid user: missing required field
        try:
            TestUser(name="Jane")
            assert False, "Should have raised validation error for missing email"
        except Exception:
            pass  # Expected

        # Invalid user: name too short
        try:
            TestUser(name="", email="te")
            assert False, (
                "Should have raised validation error for empty name or short email"
            )
        except Exception:
            pass  # Expected

        # Test Product schema
        TestProduct = getattr(module, "TestProduct")
        from uuid import uuid4

        valid_product = {
            "id": str(uuid4()),
            "price": 99.99,
            "tags": ["electronics", "gadget"],
        }
        product_instance = TestProduct(**valid_product)
        assert product_instance.price == 99.99
        assert len(product_instance.tags) == 2

        # Invalid product: negative price
        try:
            TestProduct(id=str(uuid4()), price=-10, tags=["test"])
            assert False, "Should have raised validation error for negative price"
        except Exception:
            pass  # Expected

        # Test Status enum (TypeAdapter uses var_name)
        test_status = getattr(module, "test_status")
        status_instance = test_status.validate_python("active")
        assert status_instance == "active"

        # Invalid status
        try:
            test_status.validate_python("invalid")
            assert False, "Should have raised validation error for invalid enum"
        except Exception:
            pass  # Expected

        # Test Count (non-object root type)
        test_count = getattr(module, "test_count")
        count_instance = test_count.validate_python(42)
        assert count_instance == 42

        # Invalid count: negative
        try:
            test_count.validate_python(-5)
            assert False, "Should have raised validation error for negative count"
        except Exception:
            pass  # Expected

        # Test StringArray (array root type)
        test_string_array = getattr(module, "test_string_array")
        array_instance = test_string_array.validate_python(["hello", "world"])
        assert array_instance == ["hello", "world"]

        # Clean up
        del sys.modules["generated_schemas"]


def test_schema_with_advanced_features():
    """Test schemas with advanced JSON Schema features."""

    test_input = [
        {
            "namespace": "advanced",
            "id": "ComplexObject",
            "varName": "advanced_complex_object",
            "schema": {
                "type": "object",
                "properties": {
                    "id": {"type": "integer"},
                    "data": {
                        "anyOf": [
                            {"type": "string"},
                            {"type": "number"},
                            {"type": "boolean"},
                        ]
                    },
                },
                "required": ["id"],
                "minProperties": 1,
                "maxProperties": 5,
            },
        },
        {
            "namespace": "advanced",
            "id": "DateRange",
            "varName": "advanced_date_range",
            "schema": {
                "type": "object",
                "properties": {
                    "start": {"type": "string", "format": "date"},
                    "end": {"type": "string", "format": "date"},
                },
                "required": ["start", "end"],
            },
        },
    ]

    input_json = json.dumps(test_input)
    result = subprocess.run(
        ["uv", "run", "xschema-pydantic"],
        input=input_json,
        capture_output=True,
        text=True,
        cwd=Path(__file__).parent.parent,
    )

    assert result.returncode == 0, f"Adapter failed: {result.stderr}"

    try:
        output = json.loads(result.stdout)
    except json.JSONDecodeError as e:
        raise AssertionError(f"Invalid JSON output: {result.stdout}") from e

    assert len(output) == 2, f"Expected 2 results, got {len(output)}"

    # Build and test generated code
    generated_code_lines = []

    all_imports = set()
    for item in output:
        all_imports.update(item["imports"])

    for imp in sorted(all_imports):
        generated_code_lines.append(imp)

    generated_code_lines.append("")

    for item in output:
        generated_code_lines.append(f"# {item['id']}")
        generated_code_lines.append(item["schema"])
        generated_code_lines.append("")

    generated_code = "\n".join(generated_code_lines)

    with tempfile.TemporaryDirectory() as tmpdir:
        module_path = Path(tmpdir) / "advanced_schemas.py"
        module_path.write_text(generated_code)

        spec = spec_from_file_location("advanced_schemas", module_path)
        assert spec is not None
        assert spec.loader is not None

        module = module_from_spec(spec)
        sys.modules["advanced_schemas"] = module
        spec.loader.exec_module(module)

        # Test ComplexObject (class name is PascalCase)
        AdvancedComplexObject = getattr(module, "AdvancedComplexObject")

        # Valid with string data
        obj1 = AdvancedComplexObject(id=1, data="test")
        assert obj1.id == 1
        assert obj1.data == "test"

        # Valid with number data
        obj2 = AdvancedComplexObject(id=2, data=42.5)
        assert obj2.data == 42.5

        # Valid with boolean data
        obj3 = AdvancedComplexObject(id=3, data=True)
        assert obj3.data is True

        # Test DateRange (class name is PascalCase)
        # Note: Format is annotation-only per JSON Schema 2020-12, so dates are strings
        AdvancedDateRange = getattr(module, "AdvancedDateRange")

        date_range = AdvancedDateRange(start="2024-01-01", end="2024-12-31")
        assert date_range.start == "2024-01-01"
        assert date_range.end == "2024-12-31"

        del sys.modules["advanced_schemas"]
