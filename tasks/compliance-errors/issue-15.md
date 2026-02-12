# issue-15: ref / order of evaluation: $id and $ref on nested schema

## Error Signature

- **keyword**: ref
- **case**: order of evaluation: $id and $ref on nested schema
- **normalized got**: `error: bundling error: failed to resolve internal refs for compliance://<draft>/ref/group_<n>: encountered unresolved non-local $ref "./bar.json"`

## Root Cause

The bundler's `processObject` returned early when encountering a `$ref`, skipping processing of sibling keys. For schemas with both `$ref` and `$defs` (valid in draft2019-09+), the `$defs` children containing their own `$ref` values were never processed through `processNode`, leaving relative refs like `./bar.json` unresolved.

## Fix

Modified `processObject` in `cli/bundler/bundler.go` to process sibling keys after `processRef` returns. When `processRef` returns a map result, the remaining keys (excluding `$ref` and non-schema keywords) are processed through `processNode` with the correct base URI.

## Baseline (20 failures)

| # | adapter | draft | test | expected | got | report path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript arktype | draft2019-09 | data is valid against nested sibling | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 2 | typescript arktype | draft2019-09 | data is invalid against nested sibling | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 3 | typescript arktype | draft2020-12 | data is valid against nested sibling | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 4 | typescript arktype | draft2020-12 | data is invalid against nested sibling | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 5 | typescript effect | draft2019-09 | data is valid against nested sibling | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 6 | typescript effect | draft2019-09 | data is invalid against nested sibling | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 7 | typescript effect | draft2020-12 | data is valid against nested sibling | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 8 | typescript effect | draft2020-12 | data is invalid against nested sibling | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 9 | python pydantic | draft2019-09 | data is valid against nested sibling | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json" | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 10 | python pydantic | draft2019-09 | data is invalid against nested sibling | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json" | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 11 | python pydantic | draft2020-12 | data is valid against nested sibling | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json" | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 12 | python pydantic | draft2020-12 | data is invalid against nested sibling | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json" | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 13 | typescript valibot | draft2019-09 | data is valid against nested sibling | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 14 | typescript valibot | draft2019-09 | data is invalid against nested sibling | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 15 | typescript valibot | draft2020-12 | data is valid against nested sibling | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 16 | typescript valibot | draft2020-12 | data is invalid against nested sibling | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 17 | typescript zod | draft2019-09 | data is valid against nested sibling | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 18 | typescript zod | draft2019-09 | data is invalid against nested sibling | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 19 | typescript zod | draft2020-12 | data is valid against nested sibling | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 20 | typescript zod | draft2020-12 | data is invalid against nested sibling | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json" | typescript/packages/adapters/zod/compliance/results/REPORT.md |
