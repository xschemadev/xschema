"""CLI entry point for xschema-pydantic adapter."""

import json
import sys
from typing import Any

from xschema_pydantic.converter import convert


def main() -> None:
    """Main entry point - reads JSON from stdin, writes results to stdout.

    Expected input format (JSON array):
        [{"namespace": "...", "id": "...", "varName": "...", "schema": {...}}, ...]

    Output format (JSON array):
        [{"namespace": "...", "id": "...", "varName": "...", "imports": [...], "schema": "...", "type": "..."}, ...]

    Errors are output as JSON to stderr:
        {"error": "error message"}
    """
    try:
        # Read JSON array from stdin
        inputs: list[dict[str, Any]] = json.load(sys.stdin)

        # Process each input item
        results: list[dict[str, Any]] = []
        for item in inputs:
            # Validate required keys
            if not isinstance(item, dict):
                raise ValueError(f"Expected dict, got {type(item).__name__}")

            # Call converter for each item
            result = convert(item)
            results.append(result)

        # Write JSON array to stdout
        json.dump(results, sys.stdout)

    except json.JSONDecodeError as e:
        # Handle invalid JSON input
        json.dump({"error": f"Invalid JSON input: {e}"}, sys.stderr)
        sys.exit(1)
    except Exception as e:
        # Handle other errors gracefully
        json.dump({"error": str(e)}, sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()
