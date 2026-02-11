"""
Pydantic type-fidelity harness.

Regenerates the probe fixture from the adapter's convert(), inspects the
rendered type expressions for Any degradation, and checks each probe against
its expectation. Exits non-zero if any probe regresses.

Run from adapter dir: uv run python type-probe/type-fidelity.py
"""

import re
import sys
from pathlib import Path

from xschema_pydantic.converter import convert

# ── probe definitions ────────────────────────────────────────────────

PROBE_CASES = [
    # --- mixed allOf ---
    {
        "name": "allOfMixed",
        "description": "allOf with object + string constraint",
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
        "expectAny": False,
    },
    {
        "name": "allOfObjectAndAdditional",
        "description": "TSConfig-like allOf with object + additionalProperties",
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
        "expectAny": False,
    },
    # --- oneOf ---
    {
        "name": "oneOfStringOrNumber",
        "description": "oneOf string | number",
        "schema": {"oneOf": [{"type": "string"}, {"type": "number"}]},
        "expectAny": False,
    },
    {
        "name": "oneOfObjects",
        "description": "oneOf discriminated objects",
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
        "expectAny": False,
    },
    # --- not ---
    {
        "name": "notString",
        "description": "not string",
        "schema": {"not": {"type": "string"}},
        "expectAny": True,
    },
    {
        "name": "notBoolean",
        "description": "not boolean",
        "schema": {"not": {"type": "boolean"}},
        "expectAny": True,
    },
    # --- conditional ---
    {
        "name": "conditionalIfThenElse",
        "description": "if/then/else",
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
        "expectAny": False,
    },
    {
        "name": "conditionalIfThen",
        "description": "if/then only",
        "schema": {
            "if": {"type": "string", "minLength": 1},
            "then": {"type": "string", "maxLength": 10},
        },
        "expectAny": True,
    },
    # --- typeGuarded ---
    {
        "name": "typeGuardedObject",
        "description": "type-guarded object",
        "schema": {
            "properties": {"name": {"type": "string"}},
            "required": ["name"],
        },
        "expectAny": False,
    },
    {
        "name": "typeGuardedArrayMinItems",
        "description": "type-guarded array constraint",
        "schema": {"minItems": 1},
        "expectAny": True,
    },
    # --- tuple ---
    {
        "name": "tupleStringNumber",
        "description": "open tuple [string, number]",
        "schema": {
            "type": "array",
            "prefixItems": [{"type": "string"}, {"type": "number"}],
        },
        "expectAny": True,
    },
    {
        "name": "tupleStringNumberClosed",
        "description": "closed tuple [string, number]",
        "schema": {
            "type": "array",
            "prefixItems": [{"type": "string"}, {"type": "number"}],
            "items": False,
        },
        "expectAny": False,
    },
    {
        "name": "tupleMixedWithRest",
        "description": "tuple [string, number, ...boolean[]]",
        "schema": {
            "type": "array",
            "prefixItems": [{"type": "string"}, {"type": "number"}],
            "items": {"type": "boolean"},
        },
        "expectAny": False,
    },
    # --- complex const ---
    {
        "name": "constObject",
        "description": "const object",
        "schema": {"const": {"name": "alice", "age": 30}},
        "expectAny": False,
    },
    {
        "name": "constArray",
        "description": "const array",
        "schema": {"const": [1, 2, 3]},
        "expectAny": False,
    },
    # --- complex enum ---
    {
        "name": "enumWithObjects",
        "description": "enum with objects",
        "schema": {"enum": [{"a": 1}, {"b": 2}, "simple"]},
        "expectAny": False,
    },
    {
        "name": "enumWithArrays",
        "description": "enum with arrays",
        "schema": {"enum": [[1, 2], [3, 4], None]},
        "expectAny": False,
    },
]


# ── step 1: generate probe fixture and collect type expressions ──────


def _has_any(type_str: str) -> bool:
    """Check if a type expression contains Any (standalone token, not part of a class name)."""
    return bool(re.search(r"\bAny\b", type_str))


def generate_and_collect() -> tuple[str, dict[str, tuple[str, bool]]]:
    """Generate probe fixture and collect type expressions from convert results.

    Returns (fixture_content, type_map) where type_map maps
    var_name -> (type_expr, has_any).
    """
    all_imports: set[str] = set()
    code_blocks: list[str] = []
    type_map: dict[str, tuple[str, bool]] = {}

    for probe in PROBE_CASES:
        var_name = f"probe_{probe['name']}"
        result = convert(
            {
                "namespace": "probe",
                "id": probe["name"],
                "varName": var_name,
                "schema": probe["schema"],
            }
        )

        for imp in result["imports"]:
            all_imports.add(imp)

        code_blocks.append(f"# {probe['description']}")
        code_blocks.append(result["schema"])
        code_blocks.append("")

        # the convert result's "type" field is the rendered type expression
        type_expr = result["type"]
        type_map[var_name] = (type_expr, _has_any(type_expr))

    lines = [
        '"""',
        "Pydantic Type Probe Fixture - AUTO-GENERATED by type-fidelity harness",
        "Do NOT edit manually.",
        '"""',
        "",
    ]

    for imp in sorted(all_imports):
        lines.append(imp)
    lines.append("")
    lines.append("")

    for block in code_blocks:
        lines.append(block)

    lines.append("")

    return "\n".join(lines), type_map


# ── step 2: check expectations ──────────────────────────────────────


def check_expectations(
    type_map: dict[str, tuple[str, bool]],
) -> list[dict]:
    results = []

    for probe in PROBE_CASES:
        var_name = f"probe_{probe['name']}"
        ext = type_map.get(var_name)

        if ext is None:
            results.append(
                {
                    "probe": var_name,
                    "description": probe["description"],
                    "type": "(not found)",
                    "hasAny": True,
                    "expectAny": probe["expectAny"],
                    "status": "fail",
                }
            )
            continue

        type_str, has_any = ext

        if not probe["expectAny"] and has_any:
            status = "fail"
        elif probe["expectAny"] and not has_any:
            status = "improved"
        else:
            status = "pass"

        results.append(
            {
                "probe": var_name,
                "description": probe["description"],
                "type": type_str,
                "hasAny": has_any,
                "expectAny": probe["expectAny"],
                "status": status,
            }
        )

    return results


# ── main ─────────────────────────────────────────────────────────────


def main():
    fixture_dir = Path(__file__).parent
    fixture_path = fixture_dir / "probe-fixture.py"

    # step 1: generate fixture and collect type expressions
    fixture_content, type_map = generate_and_collect()
    fixture_path.write_text(fixture_content)

    # step 2: check expectations against rendered type expressions
    results = check_expectations(type_map)

    failures = [r for r in results if r["status"] == "fail"]
    improvements = [r for r in results if r["status"] == "improved"]
    any_count = sum(1 for r in results if r["hasAny"])

    # report
    print("pydantic type-fidelity report")
    print("=============================\n")

    print("| probe | type | any? | expected | status |")
    print("| ----- | ---- | ---- | -------- | ------ |")
    for r in results:
        any_str = "YES" if r["hasAny"] else "no"
        expect_str = "any ok" if r["expectAny"] else "NO any"
        icon = {"fail": "FAIL", "improved": "IMPROVED", "pass": "ok"}[r["status"]]
        print(f"| {r['probe']} | `{r['type']}` | {any_str} | {expect_str} | {icon} |")

    print(
        f"\ntotal: {len(results)} | any: {any_count} | any-free: {len(results) - any_count}"
    )

    if improvements:
        print(
            f"\n{len(improvements)} probe(s) improved - update expectAny to False in type-fidelity.py to lock in:"
        )
        for r in improvements:
            print(f"  - {r['probe']}: now `{r['type']}`")

    if failures:
        print(f"\nFAILED: {len(failures)} probe(s) regressed to Any:")
        for r in failures:
            print(f"  - {r['probe']}: got `{r['type']}`, expected no Any")
        sys.exit(1)

    print("\nall checks passed")


if __name__ == "__main__":
    main()
