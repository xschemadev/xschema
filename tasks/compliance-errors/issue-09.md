# issue-09: definitions / validate definition against metaschema

## Error Signature

- **keyword**: `definitions`
- **case**: `validate definition against metaschema`
- **normalized got**: `error: bundling error: failed to resolve internal refs for compliance://<draft>/definitions/group_0: encountered unresolved non-local $ref "<metaschema-url>"`

## Root Cause

The bundler's `processRef` in `cli/bundler/bundler.go` detects metaschema URLs via `isMetaschema()` and skips bundling — but returns the `$ref` object unchanged, leaving an external URL in the bundled output. The subsequent `resolveInternalRefs` in `cli/processor/local_refs.go` rejects any non-local `$ref` (not starting with `#`), producing this error.

## Baseline (30 failures)

| # | Adapter | Draft | Test | Expected | Got | Report Path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript arktype | draft4 | valid definition schema | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft4.json` |
| 2 | typescript arktype | draft4 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft4.json` |
| 3 | typescript arktype | draft6 | valid definition schema | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft6.json` |
| 4 | typescript arktype | draft6 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft6.json` |
| 5 | typescript arktype | draft7 | valid definition schema | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft7.json` |
| 6 | typescript arktype | draft7 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft7.json` |
| 7 | typescript effect | draft4 | valid definition schema | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft4.json` |
| 8 | typescript effect | draft4 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft4.json` |
| 9 | typescript effect | draft6 | valid definition schema | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft6.json` |
| 10 | typescript effect | draft6 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft6.json` |
| 11 | typescript effect | draft7 | valid definition schema | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft7.json` |
| 12 | typescript effect | draft7 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft7.json` |
| 13 | python pydantic | draft4 | valid definition schema | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft4.json` |
| 14 | python pydantic | draft4 | invalid definition schema | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft4.json` |
| 15 | python pydantic | draft6 | valid definition schema | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft6.json` |
| 16 | python pydantic | draft6 | invalid definition schema | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft6.json` |
| 17 | python pydantic | draft7 | valid definition schema | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft7.json` |
| 18 | python pydantic | draft7 | invalid definition schema | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft7.json` |
| 19 | typescript valibot | draft4 | valid definition schema | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft4.json` |
| 20 | typescript valibot | draft4 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft4.json` |
| 21 | typescript valibot | draft6 | valid definition schema | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft6.json` |
| 22 | typescript valibot | draft6 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft6.json` |
| 23 | typescript valibot | draft7 | valid definition schema | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft7.json` |
| 24 | typescript valibot | draft7 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft7.json` |
| 25 | typescript zod | draft4 | valid definition schema | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft4.json` |
| 26 | typescript zod | draft4 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft4.json` |
| 27 | typescript zod | draft6 | valid definition schema | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft6.json` |
| 28 | typescript zod | draft6 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft6.json` |
| 29 | typescript zod | draft7 | valid definition schema | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft7.json` |
| 30 | typescript zod | draft7 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft7.json` |
