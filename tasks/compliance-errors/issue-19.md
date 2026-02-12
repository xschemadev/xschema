# issue-19: ref / $ref prevents a sibling $id from changing the base uri

## Error Signature

- **Keyword**: `ref`
- **Case**: `$ref prevents a sibling $id from changing the base uri`
- **Normalized Got**: `false`

## Root Cause

In pre-2019-09 drafts (draft6, draft7), `$ref` consumes the entire object — sibling keywords including `$id` are ignored. But the bundler's `processObject` applied the sibling `$id` base URI change before resolving `$ref`, causing the ref to resolve against the wrong base URI and match the wrong definition.

## Baseline (10 failures)

| # | Adapter | Draft | Test | Expected | Got | Report Path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript arktype | draft6 | $ref resolves to /definitions/base_foo, data validates | valid | false | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 2 | typescript arktype | draft7 | $ref resolves to /definitions/base_foo, data validates | valid | false | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 3 | typescript effect | draft6 | $ref resolves to /definitions/base_foo, data validates | valid | false | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 4 | typescript effect | draft7 | $ref resolves to /definitions/base_foo, data validates | valid | false | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 5 | python pydantic | draft6 | $ref resolves to /definitions/base_foo, data validates | valid | false | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 6 | python pydantic | draft7 | $ref resolves to /definitions/base_foo, data validates | valid | false | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 7 | typescript valibot | draft6 | $ref resolves to /definitions/base_foo, data validates | valid | false | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 8 | typescript valibot | draft7 | $ref resolves to /definitions/base_foo, data validates | valid | false | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 9 | typescript zod | draft6 | $ref resolves to /definitions/base_foo, data validates | valid | false | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 10 | typescript zod | draft7 | $ref resolves to /definitions/base_foo, data validates | valid | false | typescript/packages/adapters/zod/compliance/results/REPORT.md |

## Fix

In `cli/bundler/bundler.go` `processObject`: skip sibling `$id` base URI change when `$ref` is present and draft is pre-2019-09 (`needsNormalization(b.draft)`). This ensures `$ref` resolves against the parent's base URI, not the sibling `$id`'s modified URI.
