"""
Generates the type probe fixture by running the pydantic adapter's convert function
against representative JSON schemas for each targeted construct.

Run from the pydantic adapter directory:
    uv run python ../../../tasks/type-fidelity-baseline/python/generate-probe.py
"""

import json
import os
import sys
from pathlib import Path

from xschema_pydantic.converter import convert

PROBE_CASES = [
    # --- mixed allOf ---
    {
        "name": "allOfMixed",
        "description": "allOf with object + string constraint - should ideally preserve object type",
        "schema": {
            "allOf": [
                {
                    "type": "object",
                    "properties": {
                        "name": {"type": "string"},
                        "value": {"type": "integer"},
                    },
                    "required": ["name"],
                },
                {
                    "properties": {"name": {"minLength": 1}},
                },
            ],
        },
    },
    {
        "name": "allOfObjectAndAdditional",
        "description": "TSConfig-like allOf with object + additionalProperties constraint",
        "schema": {
            "allOf": [
                {
                    "type": "object",
                    "properties": {"strict": {"type": "boolean"}},
                },
                {
                    "additionalProperties": {"type": "string"},
                },
            ],
        },
    },
    # --- oneOf ---
    {
        "name": "oneOfStringOrNumber",
        "description": "oneOf with string | number - should be str | int, currently Any",
        "schema": {
            "oneOf": [{"type": "string"}, {"type": "number"}],
        },
    },
    {
        "name": "oneOfObjects",
        "description": "oneOf with discriminated objects - currently Any",
        "schema": {
            "oneOf": [
                {
                    "type": "object",
                    "properties": {"kind": {"const": "a"}, "value": {"type": "string"}},
                    "required": ["kind", "value"],
                },
                {
                    "type": "object",
                    "properties": {"kind": {"const": "b"}, "value": {"type": "number"}},
                    "required": ["kind", "value"],
                },
            ],
        },
    },
    # --- not ---
    {
        "name": "notString",
        "description": "not string - negation can't narrow, currently Any",
        "schema": {
            "not": {"type": "string"},
        },
    },
    {
        "name": "notBoolean",
        "description": "not boolean - negation can't narrow, currently Any",
        "schema": {
            "not": {"type": "boolean"},
        },
    },
    # --- conditional (if/then/else) ---
    {
        "name": "conditionalIfThenElse",
        "description": "if/then/else - should ideally be union of then/else types, currently Any",
        "schema": {
            "if": {
                "type": "object",
                "properties": {"kind": {"const": "a"}},
                "required": ["kind"],
            },
            "then": {
                "type": "object",
                "properties": {"kind": {"const": "a"}, "value": {"type": "string"}},
                "required": ["kind", "value"],
            },
            "else": {
                "type": "object",
                "properties": {"kind": {"const": "b"}, "value": {"type": "number"}},
                "required": ["kind", "value"],
            },
        },
    },
    {
        "name": "conditionalIfThen",
        "description": "if/then only - currently Any",
        "schema": {
            "if": {"type": "string", "minLength": 1},
            "then": {"type": "string", "maxLength": 10},
        },
    },
    # --- typeGuarded ---
    {
        "name": "typeGuardedObject",
        "description": "type-guarded object without explicit type - currently Any",
        "schema": {
            "properties": {"name": {"type": "string"}},
            "required": ["name"],
        },
    },
    {
        "name": "typeGuardedArrayMinItems",
        "description": "type-guarded array constraint without explicit type - currently Any",
        "schema": {
            "minItems": 1,
        },
    },
    # --- tuple ---
    {
        "name": "tupleStringNumber",
        "description": "tuple [string, number] - should be tuple[str, int, ...], currently tuple[Any, ...]",
        "schema": {
            "type": "array",
            "prefixItems": [{"type": "string"}, {"type": "number"}],
        },
    },
    {
        "name": "tupleStringNumberClosed",
        "description": "closed tuple [string, number] - should be tuple[str, int], currently tuple[Any, ...]",
        "schema": {
            "type": "array",
            "prefixItems": [{"type": "string"}, {"type": "number"}],
            "items": False,
        },
    },
    {
        "name": "tupleMixedWithRest",
        "description": "tuple [string, number, ...boolean[]] - currently tuple[Any, ...]",
        "schema": {
            "type": "array",
            "prefixItems": [{"type": "string"}, {"type": "number"}],
            "items": {"type": "boolean"},
        },
    },
    # --- complex const ---
    {
        "name": "constObject",
        "description": "const object - should narrow, currently Any",
        "schema": {
            "const": {"name": "alice", "age": 30},
        },
    },
    {
        "name": "constArray",
        "description": "const array - should narrow, currently Any",
        "schema": {
            "const": [1, 2, 3],
        },
    },
    # --- complex enum ---
    {
        "name": "enumWithObjects",
        "description": "enum with object values - currently Any",
        "schema": {
            "enum": [{"a": 1}, {"b": 2}, "simple"],
        },
    },
    {
        "name": "enumWithArrays",
        "description": "enum with array values - currently Any",
        "schema": {
            "enum": [[1, 2], [3, 4], None],
        },
    },
]


def main():
    out_dir = Path(__file__).parent
    fixture_path = out_dir / "probe-fixture.py"

    all_imports: set[str] = set()
    code_blocks: list[str] = []

    for probe in PROBE_CASES:
        result = convert(
            {
                "namespace": "probe",
                "id": probe["name"],
                "varName": f"probe_{probe['name']}",
                "schema": probe["schema"],
            }
        )

        for imp in result["imports"]:
            all_imports.add(imp)

        code_blocks.append(f"# {probe['description']}")
        code_blocks.append(result["schema"])
        code_blocks.append("")

    # Build the fixture
    lines = [
        '"""',
        "Pydantic Type Probe Fixture - AUTO-GENERATED",
        "",
        "This file contains generated Pydantic code for representative schemas of each",
        "targeted construct. The inferred types from pyright reveal where TypeAdapter[Any]",
        "degrades the type information.",
        "",
        f"Generated: {__import__('datetime').date.today().isoformat()}",
        "Do NOT edit manually - regenerate with:",
        "    cd python/packages/adapters/pydantic",
        "    uv run python ../../../../tasks/type-fidelity-baseline/python/generate-probe.py",
        '"""',
        "",
    ]

    # Sort imports
    sorted_imports = sorted(all_imports)
    for imp in sorted_imports:
        lines.append(imp)
    lines.append("")
    lines.append("")

    # Add helper functions needed by the generated code
    lines.append("# Helper functions needed by generated code")
    lines.append("def _json_equals(a: object, b: object) -> bool:")
    lines.append("    import json")
    lines.append(
        "    return json.dumps(a, sort_keys=True) == json.dumps(b, sort_keys=True)"
    )
    lines.append("")
    lines.append("")
    lines.append("def _try_validate(fn):")  # noqa: E501
    lines.append("    def wrapper(v):")
    lines.append("        try:")
    lines.append("            fn(v)")
    lines.append("            return True")
    lines.append("        except Exception:")
    lines.append("            return False")
    lines.append("    return wrapper")
    lines.append("")
    lines.append("")

    for block in code_blocks:
        lines.append(block)

    # Add reveal_type calls for pyright inspection
    lines.append("")
    lines.append("# reveal_type calls for pyright type inspection")
    for probe in PROBE_CASES:
        var_name = f"probe_{probe['name']}"
        lines.append(f"reveal_type({var_name})  # noqa: F821")

    lines.append("")

    fixture_path.write_text("\n".join(lines))
    print(f"Wrote probe fixture to: {fixture_path}")
    print(f"Probe cases: {len(PROBE_CASES)}")

    # Also write a JSON summary of what was generated
    summary = []
    for probe in PROBE_CASES:
        result = convert(
            {
                "namespace": "probe",
                "id": probe["name"],
                "varName": f"probe_{probe['name']}",
                "schema": probe["schema"],
            }
        )
        summary.append(
            {
                "name": probe["name"],
                "description": probe["description"],
                "type_expr": result["type"],
                "has_any": "Any" in result["type"],
            }
        )

    summary_path = out_dir / "probe-summary.json"
    summary_path.write_text(json.dumps(summary, indent=2) + "\n")
    print(f"Wrote probe summary to: {summary_path}")


if __name__ == "__main__":
    main()
