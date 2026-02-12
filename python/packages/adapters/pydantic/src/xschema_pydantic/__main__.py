"""CLI entry point for xschema-pydantic adapter."""

import json
import sys
import warnings
from typing import Any

from xschema_pydantic.converter import convert
from xschema_pydantic.errors import (
    ConversionError,
    InvalidSchemaError,
    UnsupportedFeatureError,
    XSchemaError,
)


def format_error(error: Exception, schema_id: str | None = None) -> dict[str, Any]:
    """Format an error as a JSON object with helpful information.

    Args:
        error: The exception that occurred
        schema_id: Optional schema ID where error occurred

    Returns:
        Dict with error, type, and optional schemaId/schemaPath keys
    """
    error_dict: dict[str, Any] = {
        "error": str(error),
        "type": type(error).__name__,
    }

    if schema_id:
        error_dict["schemaId"] = schema_id

    # Add schema path if available
    if isinstance(error, XSchemaError) and error.schema_path:
        error_dict["schemaPath"] = error.schema_path

    return error_dict


def main() -> None:
    """Main entry point - reads JSON from stdin, writes results to stdout.

    Expected input format (JSON array):
        [{"namespace": "...", "id": "...", "varName": "...", "schema": {...}}, ...]

    Output format (JSON array):
        [{"namespace": "...", "id": "...", "varName": "...", "imports": [...], "schema": "...", "type": "..."}, ...]

    Errors are output as JSON to stderr with structure:
        {"error": "message", "type": "ErrorType", "schemaId": "...", "schemaPath": "..."}

    Exit codes:
        0: Success
        1: Invalid input (JSON decode error, wrong input format)
        2: Schema validation error (malformed schema)
        3: Conversion error (failed to generate code)
    """
    try:
        # Read JSON array from stdin
        try:
            inputs: list[dict[str, Any]] = json.load(sys.stdin)
        except json.JSONDecodeError as e:
            json.dump(
                {
                    "error": f"Invalid JSON input: {e.msg}",
                    "type": "JSONDecodeError",
                    "line": e.lineno,
                    "column": e.colno,
                },
                sys.stderr,
            )
            sys.exit(1)

        # Validate input is a list
        if not isinstance(inputs, list):
            json.dump(
                {
                    "error": f"Expected JSON array, got {type(inputs).__name__}",
                    "type": "ValueError",
                },
                sys.stderr,
            )
            sys.exit(1)

        # Process each input item
        results: list[dict[str, Any]] = []
        for i, item in enumerate(inputs):
            # Validate item is a dict
            if not isinstance(item, dict):
                json.dump(
                    {
                        "error": f"Item at index {i} is not a dict, got {type(item).__name__}",
                        "type": "ValueError",
                        "index": i,
                    },
                    sys.stderr,
                )
                sys.exit(1)

            schema_id = item.get("id", f"item[{i}]")

            try:
                # Call converter for each item
                result = convert(item)
                results.append(result)

            except UnsupportedFeatureError as e:
                # Unsupported features generate warnings, not errors
                # This allows partial conversion to succeed
                warnings.warn(str(e), UserWarning, stacklevel=2)
                # Write warning to stderr as JSON
                warning_dict = format_error(e, schema_id)
                warning_dict["severity"] = "warning"
                json.dump(warning_dict, sys.stderr)
                sys.stderr.write("\n")
                # Skip this schema and continue
                continue

            except InvalidSchemaError as e:
                # Schema validation errors exit with code 2
                json.dump(format_error(e, schema_id), sys.stderr)
                sys.exit(2)

            except ConversionError as e:
                # Conversion errors exit with code 3
                json.dump(format_error(e, schema_id), sys.stderr)
                sys.exit(3)

            except Exception as e:
                # Unexpected errors exit with code 3
                json.dump(
                    {
                        "error": f"Unexpected error: {str(e)}",
                        "type": type(e).__name__,
                        "schemaId": schema_id,
                    },
                    sys.stderr,
                )
                sys.exit(3)

        # Write JSON array to stdout
        json.dump(results, sys.stdout)

    except Exception as e:
        # Top-level unexpected errors
        json.dump(
            {
                "error": f"Fatal error: {str(e)}",
                "type": type(e).__name__,
            },
            sys.stderr,
        )
        sys.exit(1)


if __name__ == "__main__":
    main()
