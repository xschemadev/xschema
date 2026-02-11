"""
Pydantic type-fidelity harness.

Regenerates the probe fixture from the adapter's convert(), runs pyright
to extract inferred types from reveal_type() calls, and checks each probe
against its expectation. Exits non-zero if any probe regresses.

Run from adapter dir: uv run python type-probe/type-fidelity.py
"""

import json
import re
import subprocess
import sys
import textwrap
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
        "expectAny": True,
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
        "expectAny": True,
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
        "expectAny": True,
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
        "expectAny": True,
    },
    {
        "name": "tupleMixedWithRest",
        "description": "tuple [string, number, ...boolean[]]",
        "schema": {
            "type": "array",
            "prefixItems": [{"type": "string"}, {"type": "number"}],
            "items": {"type": "boolean"},
        },
        "expectAny": True,
    },
    # --- complex const ---
    {
        "name": "constObject",
        "description": "const object",
        "schema": {"const": {"name": "alice", "age": 30}},
        "expectAny": True,
    },
    {
        "name": "constArray",
        "description": "const array",
        "schema": {"const": [1, 2, 3]},
        "expectAny": True,
    },
    # --- complex enum ---
    {
        "name": "enumWithObjects",
        "description": "enum with objects",
        "schema": {"enum": [{"a": 1}, {"b": 2}, "simple"]},
        "expectAny": True,
    },
    {
        "name": "enumWithArrays",
        "description": "enum with arrays",
        "schema": {"enum": [[1, 2], [3, 4], None]},
        "expectAny": True,
    },
]


# ── step 1: generate probe fixture ──────────────────────────────────


def generate_fixture() -> str:
    """Generate a Python file with probe variables and reveal_type() calls."""
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
    lines.append("# reveal_type calls for pyright type inspection")
    for probe in PROBE_CASES:
        var_name = f"probe_{probe['name']}"
        lines.append(f"reveal_type({var_name})")

    lines.append("")

    return "\n".join(lines)


# ── step 2: extract types via pyright ────────────────────────────────

_REVEAL_TYPE_RE = re.compile(r'^Type of "([^"]+)" is "(.+)"$')


def _has_any(type_str: str) -> bool:
    """Check if a pyright type string indicates Any/Unknown degradation.

    pyright reports Annotated[Any, ...] as TypeAdapter[Unknown], so we
    check for both Any and Unknown (but not when part of a class name).
    """
    # TypeAdapter[Unknown] is pyright's rendering of TypeAdapter[Annotated[Any, ...]]
    if "TypeAdapter[Unknown]" in type_str:
        return True
    # list[Any], tuple[Any, ...] etc
    if "Any" in type_str:
        # avoid false positives from class names containing "Any"
        # check it appears as a standalone type token
        return bool(re.search(r"\bAny\b", type_str))
    return False


def extract_types(fixture_path: str) -> dict[str, tuple[str, bool]]:
    """Run pyright on fixture, parse reveal_type diagnostics.

    Returns dict mapping var_name -> (type_string, has_any).
    """
    result = subprocess.run(
        ["pyright", "--outputjson", fixture_path],
        capture_output=True,
        text=True,
    )

    try:
        data = json.loads(result.stdout)
    except json.JSONDecodeError:
        print("error: pyright did not produce valid JSON output", file=sys.stderr)
        print(result.stderr, file=sys.stderr)
        sys.exit(1)

    types: dict[str, tuple[str, bool]] = {}

    for diag in data.get("generalDiagnostics", []):
        if diag.get("severity") != "information":
            continue
        m = _REVEAL_TYPE_RE.match(diag.get("message", ""))
        if not m:
            continue
        var_name = m.group(1)
        type_str = m.group(2)
        types[var_name] = (type_str, _has_any(type_str))

    return types


# ── step 3: check expectations ──────────────────────────────────────


def check_expectations(
    extracted: dict[str, tuple[str, bool]],
) -> list[dict]:
    results = []

    for probe in PROBE_CASES:
        var_name = f"probe_{probe['name']}"
        ext = extracted.get(var_name)

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

    # step 1: generate
    fixture_content = generate_fixture()
    fixture_path.write_text(fixture_content)

    # step 2: extract types
    extracted = extract_types(str(fixture_path))

    # step 3: check
    results = check_expectations(extracted)

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
