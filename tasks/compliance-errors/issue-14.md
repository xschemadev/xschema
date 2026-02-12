# issue-14: ref / naive replacement of $ref with its destination is not correct

## Error Signature

- **Keyword**: `ref`
- **Case**: `naive replacement of $ref with its destination is not correct`
- **Normalized got**: `false`

## Root Cause

The bundler's `normalizeLegacySyntax` (in `cli/bundler/normalize.go`) recursed into all object values unconditionally, including data-only keywords like `enum`, `const`, `default`. When a legacy-draft schema had an `enum` containing a literal `{"$ref": "#/definitions/a_string"}`, normalization rewrote the `$ref` value from `#/definitions/...` to `#/$defs/...` and renamed the `definitions` key to `$defs`. This corrupted the enum literal, so deep-equal comparison against the original data `{"$ref": "#/definitions/a_string"}` failed.

Draft 2019-09/2020-12 were unaffected because they already use `$defs` and no normalization runs.

Fix: skip normalization recursion for `nonSchemaKeywords` (`enum`, `const`, `default`, `example`, `examples`).

## Baseline (20 failures)

| # | Adapter | Draft | Test | Expected | Got | Report Path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript arktype | draft3 | match the enum exactly | valid | `false` | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 2 | typescript arktype | draft4 | match the enum exactly | valid | `false` | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 3 | typescript arktype | draft6 | match the enum exactly | valid | `false` | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 4 | typescript arktype | draft7 | match the enum exactly | valid | `false` | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 5 | typescript effect | draft3 | match the enum exactly | valid | `false` | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 6 | typescript effect | draft4 | match the enum exactly | valid | `false` | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 7 | typescript effect | draft6 | match the enum exactly | valid | `false` | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 8 | typescript effect | draft7 | match the enum exactly | valid | `false` | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 9 | python pydantic | draft3 | match the enum exactly | valid | `false` | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 10 | python pydantic | draft4 | match the enum exactly | valid | `false` | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 11 | python pydantic | draft6 | match the enum exactly | valid | `false` | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 12 | python pydantic | draft7 | match the enum exactly | valid | `false` | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 13 | typescript valibot | draft3 | match the enum exactly | valid | `false` | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 14 | typescript valibot | draft4 | match the enum exactly | valid | `false` | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 15 | typescript valibot | draft6 | match the enum exactly | valid | `false` | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 16 | typescript valibot | draft7 | match the enum exactly | valid | `false` | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 17 | typescript zod | draft3 | match the enum exactly | valid | `false` | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 18 | typescript zod | draft4 | match the enum exactly | valid | `false` | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 19 | typescript zod | draft6 | match the enum exactly | valid | `false` | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 20 | typescript zod | draft7 | match the enum exactly | valid | `false` | typescript/packages/adapters/zod/compliance/results/REPORT.md |
