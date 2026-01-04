# JSON Schema Conformance Testing

This document explains how xschema uses the official JSON-Schema-Test-Suite to verify that adapters correctly implement JSON Schema specifications.

## What is JSON-Schema-Test-Suite?

[JSON-Schema-Test-Suite](https://github.com/json-schema-org/JSON-Schema-Test-Suite) is the official test suite maintained by the JSON Schema organization. It contains thousands of test cases that verify correct implementation of JSON Schema validation.

**Why use it:**

- **Authoritative** - maintained by the JSON Schema spec authors
- **Comprehensive** - covers all keywords across all draft versions (draft-04 through draft-2020-12)
- **Language-agnostic** - pure JSON format, usable by any implementation
- **Industry standard** - used by ajv, jsonschema (Python), and 50+ other validators

**Test structure:**

```json
{
  "description": "type keyword validation",
  "schema": { "type": "string" },
  "tests": [
    { "description": "a string is valid", "data": "foo", "valid": true },
    { "description": "an integer is invalid", "data": 1, "valid": false }
  ]
}
```

Each test case has a schema and multiple test instances with expected validity.

## xschema Adapter Testing Strategy

xschema adapters (zod, valibot, pydantic, etc.) convert JSON Schema to native validators. We test them in two phases:

### Phase 1: Code Generation

Verify the adapter produces valid code from the schema:

```
JSON Schema -> Adapter -> Generated Code (must parse/compile)
```

### Phase 2: Runtime Validation

Verify the generated validator produces correct results:

```
Generated Validator + Test Data -> Validation Result == Expected
```

## Integration Architecture

### Directory Structure

```
conformance/
  runner/
    typescript.ts      # TS adapter runner
    python.py          # Python adapter runner
  results/
    zod.json           # Test results for zod
    valibot.json       # Test results for valibot
    pydantic.json      # Test results for pydantic
  expected-failures/
    zod.json           # Known unsupported features
    valibot.json
    pydantic.json
```

### TypeScript Runner Example

```typescript
// conformance/runner/typescript.ts
import { readdir, readFile } from "fs/promises";
import { join } from "path";

interface TestCase {
  description: string;
  schema: object;
  tests: Array<{
    description: string;
    data: unknown;
    valid: boolean;
  }>;
}

interface TestResult {
  file: string;
  testCase: string;
  test: string;
  expected: boolean;
  actual: boolean | "error";
  passed: boolean;
  error?: string;
}

async function runConformanceTests(
  adapter: string,
  draft: string
): Promise<TestResult[]> {
  const suitePath = `./JSON-Schema-Test-Suite/tests/${draft}`;
  const results: TestResult[] = [];

  const files = await readdir(suitePath);

  for (const file of files.filter((f) => f.endsWith(".json"))) {
    const content = await readFile(join(suitePath, file), "utf-8");
    const testCases: TestCase[] = JSON.parse(content);

    for (const testCase of testCases) {
      // Generate validator code from schema
      const generatedCode = await generateValidator(adapter, testCase.schema);

      // Create and evaluate the validator
      let validator: (data: unknown) => boolean;
      try {
        validator = await evalValidator(generatedCode);
      } catch (err) {
        // Code generation or compilation failed
        for (const test of testCase.tests) {
          results.push({
            file,
            testCase: testCase.description,
            test: test.description,
            expected: test.valid,
            actual: "error",
            passed: false,
            error: String(err),
          });
        }
        continue;
      }

      // Run each test
      for (const test of testCase.tests) {
        let actual: boolean | "error";
        try {
          actual = validator(test.data);
        } catch {
          actual = "error";
        }

        results.push({
          file,
          testCase: testCase.description,
          test: test.description,
          expected: test.valid,
          actual,
          passed: actual === test.valid,
        });
      }
    }
  }

  return results;
}

async function generateValidator(
  adapter: string,
  schema: object
): Promise<string> {
  // Call xschema adapter to generate code
  const input = JSON.stringify([{ namespace: "test", id: "Test", schema }]);

  const { stdout } = await $`bunx ${adapter} <<< ${input}`;
  const [result] = JSON.parse(stdout);

  return `
    ${result.imports.join("\n")}
    const validator = ${result.schema};
    export default (data) => validator.safeParse(data).success;
  `;
}
```

### Python Runner Example

```python
# conformance/runner/python.py
import json
import subprocess
from pathlib import Path
from dataclasses import dataclass

@dataclass
class TestResult:
    file: str
    test_case: str
    test: str
    expected: bool
    actual: bool | str
    passed: bool
    error: str | None = None

def run_conformance_tests(adapter: str, draft: str) -> list[TestResult]:
    suite_path = Path(f"./JSON-Schema-Test-Suite/tests/{draft}")
    results = []

    for file in suite_path.glob("*.json"):
        test_cases = json.loads(file.read_text())

        for test_case in test_cases:
            # Generate validator code
            try:
                code = generate_validator(adapter, test_case["schema"])
                validator = compile_validator(code)
            except Exception as e:
                for test in test_case["tests"]:
                    results.append(TestResult(
                        file=file.name,
                        test_case=test_case["description"],
                        test=test["description"],
                        expected=test["valid"],
                        actual="error",
                        passed=False,
                        error=str(e),
                    ))
                continue

            # Run tests
            for test in test_case["tests"]:
                try:
                    actual = validator(test["data"])
                except Exception:
                    actual = "error"

                results.append(TestResult(
                    file=file.name,
                    test_case=test_case["description"],
                    test=test["description"],
                    expected=test["valid"],
                    actual=actual,
                    passed=actual == test["valid"],
                ))

    return results

def generate_validator(adapter: str, schema: dict) -> str:
    input_data = json.dumps([{"namespace": "test", "id": "Test", "schema": schema}])
    result = subprocess.run(
        ["python", "-m", adapter],
        input=input_data,
        capture_output=True,
        text=True,
    )
    output = json.loads(result.stdout)
    return output[0]["schema"]
```

## When to Run Conformance Tests

### Recommendation: Tiered Approach

| Trigger | What to Run | Rationale |
|---------|-------------|-----------|
| PR touching adapters | Core keywords only (~200 tests) | Fast feedback, catches regressions |
| Merge to main | Full suite for changed adapters | Complete validation before release |
| Weekly scheduled | Full suite, all adapters, all drafts | Catch upstream changes |
| Before release | Full suite + generate reports | Update conformance badges |

### Path-based Triggers

```yaml
# Adapter changes trigger conformance tests
on:
  pull_request:
    paths:
      - "typescript/src/adapters/**"
      - "python/src/adapters/**"
      - "conformance/**"
```

### Test Subsets for PRs

For faster PR feedback, run only essential keywords:

```typescript
const CORE_KEYWORDS = [
  "type.json",
  "properties.json",
  "required.json",
  "additionalProperties.json",
  "items.json",
  "enum.json",
  "const.json",
  "allOf.json",
  "anyOf.json",
  "oneOf.json",
];

// Skip optional/ subdirectory in PRs
const skipOptional = process.env.CI_FULL_SUITE !== "true";
```

## Displaying Conformance Results

### README Badges

Use shields.io with dynamic JSON endpoint:

```markdown
![Zod Conformance](https://img.shields.io/endpoint?url=https://xschema.dev/badges/zod.json)
![Valibot Conformance](https://img.shields.io/endpoint?url=https://xschema.dev/badges/valibot.json)
![Pydantic Conformance](https://img.shields.io/endpoint?url=https://xschema.dev/badges/pydantic.json)
```

Badge JSON format:

```json
{
  "schemaVersion": 1,
  "label": "zod conformance",
  "message": "94.2%",
  "color": "green"
}
```

### Conformance Table

```markdown
## JSON Schema Conformance

| Adapter | draft-2020-12 | draft-2019-09 | draft-07 |
|---------|--------------|---------------|----------|
| zod | 94.2% (1,847/1,961) | 92.1% | 95.3% |
| valibot | 91.8% | 89.4% | 93.2% |
| pydantic | 96.5% | 94.2% | 97.1% |

Last updated: 2026-01-04
```

### Detailed Results Page

Host a detailed breakdown at `xschema.dev/conformance`:

```
Keyword Coverage: @xschema/zod (draft-2020-12)

type             100% (54/54)
properties       100% (89/89)
required         100% (34/34)
additionalProps   98% (47/48)
  - additionalProperties with null valued instance properties [FAIL]
items            100% (67/67)
$ref              87% (52/60)
  - remote ref resolution [8 SKIP - requires network]
format            75% (optional)
...
```

## GitHub Actions Workflow

```yaml
# .github/workflows/conformance.yml
name: Conformance Tests

on:
  push:
    branches: [master]
    paths:
      - "typescript/src/adapters/**"
      - "conformance/**"
  pull_request:
    paths:
      - "typescript/src/adapters/**"
      - "conformance/**"
  schedule:
    # Weekly on Sunday at midnight
    - cron: "0 0 * * 0"
  workflow_dispatch:
    inputs:
      full_suite:
        description: "Run full test suite"
        type: boolean
        default: false

jobs:
  conformance:
    runs-on: ubuntu-latest
    strategy:
      fail-fast: false
      matrix:
        adapter: [zod, valibot]
        draft: [draft-2020-12, draft-2019-09, draft-07]
        include:
          - adapter: zod
            lang: typescript
          - adapter: valibot
            lang: typescript

    steps:
      - uses: actions/checkout@v4

      - name: Checkout JSON-Schema-Test-Suite
        uses: actions/checkout@v4
        with:
          repository: json-schema-org/JSON-Schema-Test-Suite
          path: JSON-Schema-Test-Suite
          ref: main

      - uses: oven-sh/setup-bun@v2
        if: matrix.lang == 'typescript'

      - name: Install dependencies
        run: bun install
        working-directory: typescript

      - name: Build adapters
        run: bun run build
        working-directory: typescript

      - name: Run conformance tests
        id: test
        run: |
          bun run conformance/runner/typescript.ts \
            --adapter ${{ matrix.adapter }} \
            --draft ${{ matrix.draft }} \
            --output results.json \
            ${{ github.event.inputs.full_suite == 'true' && '--full' || '' }}
        env:
          CI_FULL_SUITE: ${{ github.event_name == 'schedule' || github.event_name == 'push' }}

      - name: Upload results
        uses: actions/upload-artifact@v4
        with:
          name: conformance-${{ matrix.adapter }}-${{ matrix.draft }}
          path: results.json

      - name: Check for regressions
        run: |
          bun run conformance/check-regressions.ts \
            --baseline conformance/results/${{ matrix.adapter }}.json \
            --current results.json

  aggregate-results:
    needs: conformance
    runs-on: ubuntu-latest
    if: github.event_name == 'push' || github.event_name == 'schedule'
    steps:
      - uses: actions/checkout@v4

      - uses: actions/download-artifact@v4
        with:
          path: artifacts

      - uses: oven-sh/setup-bun@v2

      - name: Generate conformance report
        run: bun run conformance/generate-report.ts

      - name: Update badges
        run: bun run conformance/update-badges.ts

      - name: Commit results
        uses: stefanzweifel/git-auto-commit-action@v5
        with:
          commit_message: "chore: update conformance results"
          file_pattern: "conformance/results/*.json docs/conformance.md"

  # Python adapters (future)
  conformance-python:
    if: false # Enable when Python adapters are ready
    runs-on: ubuntu-latest
    strategy:
      matrix:
        adapter: [pydantic]
        draft: [draft-2020-12]
    steps:
      - uses: actions/checkout@v4

      - name: Checkout JSON-Schema-Test-Suite
        uses: actions/checkout@v4
        with:
          repository: json-schema-org/JSON-Schema-Test-Suite
          path: JSON-Schema-Test-Suite

      - uses: actions/setup-python@v5
        with:
          python-version: "3.12"

      - name: Install dependencies
        run: pip install -e ".[dev]"
        working-directory: python

      - name: Run conformance tests
        run: |
          python -m conformance.runner \
            --adapter ${{ matrix.adapter }} \
            --draft ${{ matrix.draft }}
```

## Handling Expected Failures

Not all JSON Schema features can be supported by every adapter. Document and track these explicitly.

### Expected Failures File

```json
// conformance/expected-failures/zod.json
{
  "$comment": "Features not supported by zod adapter",
  "version": "1.0.0",
  "adapter": "@xschema/zod",
  "failures": [
    {
      "file": "refRemote.json",
      "reason": "Remote $ref resolution requires runtime fetch - not supported in code generation",
      "tests": ["*"]
    },
    {
      "file": "format.json",
      "testCase": "validation of IRIs",
      "reason": "Zod doesn't have built-in IRI validation",
      "tests": ["*"]
    },
    {
      "file": "unevaluatedProperties.json",
      "reason": "unevaluatedProperties requires runtime tracking not available in Zod",
      "tests": ["*"]
    },
    {
      "file": "dynamicRef.json",
      "reason": "$dynamicRef requires runtime resolution",
      "tests": ["*"]
    }
  ]
}
```

### Failure Categories

| Category | Example | Handling |
|----------|---------|----------|
| **Unsupported keyword** | `$dynamicRef`, `unevaluatedProperties` | Skip, document in adapter README |
| **Runtime-only** | Remote `$ref` resolution | Skip or require setup |
| **Format validation** | `iri`, `uri-template` | Optional, depends on adapter |
| **Edge cases** | Unicode normalization | Document limitation |
| **Upstream bugs** | Bugs in underlying libraries | Track issue link, skip test |

### Checking Against Expected Failures

```typescript
// conformance/check-regressions.ts
interface ExpectedFailure {
  file: string;
  testCase?: string;
  tests: string[] | "*";
  reason: string;
}

function isExpectedFailure(
  result: TestResult,
  failures: ExpectedFailure[]
): boolean {
  return failures.some((f) => {
    if (f.file !== result.file) return false;
    if (f.testCase && f.testCase !== result.testCase) return false;
    if (f.tests === "*") return true;
    return f.tests.includes(result.test);
  });
}

function checkRegressions(
  baseline: TestResult[],
  current: TestResult[],
  expectedFailures: ExpectedFailure[]
) {
  const regressions: TestResult[] = [];
  const improvements: TestResult[] = [];

  for (const curr of current) {
    const base = baseline.find(
      (b) =>
        b.file === curr.file &&
        b.testCase === curr.testCase &&
        b.test === curr.test
    );

    // New failure that's not expected
    if (!curr.passed && base?.passed && !isExpectedFailure(curr, expectedFailures)) {
      regressions.push(curr);
    }

    // Previously failing test now passes
    if (curr.passed && base && !base.passed) {
      improvements.push(curr);
    }
  }

  if (regressions.length > 0) {
    console.error(`${regressions.length} regressions detected!`);
    process.exit(1);
  }

  if (improvements.length > 0) {
    console.log(`${improvements.length} tests improved - update baseline!`);
  }
}
```

### PR Annotations

```yaml
- name: Annotate PR with failures
  if: failure() && github.event_name == 'pull_request'
  uses: actions/github-script@v7
  with:
    script: |
      const results = require('./results.json');
      const failures = results.filter(r => !r.passed);

      for (const f of failures.slice(0, 10)) {
        core.error(
          `${f.file}: ${f.testCase} - ${f.test}`,
          { title: 'Conformance Failure' }
        );
      }

      if (failures.length > 10) {
        core.error(`... and ${failures.length - 10} more failures`);
      }
```

## Quick Start

1. **Add test suite as submodule:**

   ```bash
   git submodule add https://github.com/json-schema-org/JSON-Schema-Test-Suite.git
   ```

2. **Create conformance runner for your adapter**

3. **Define expected failures for unsupported features**

4. **Add GitHub workflow**

5. **Run locally:**

   ```bash
   bun run conformance/runner/typescript.ts --adapter zod --draft draft-2020-12
   ```

## Resources

- [JSON-Schema-Test-Suite](https://github.com/json-schema-org/JSON-Schema-Test-Suite)
- [JSON Schema Specification](https://json-schema.org/specification)
- [npm package: @json-schema-org/tests](https://www.npmjs.com/package/@json-schema-org/tests)
- [Test Suite Schema](https://github.com/json-schema-org/JSON-Schema-Test-Suite/blob/main/test-schema.json)
