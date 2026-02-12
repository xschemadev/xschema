# issue-12: defs / validate definition against metaschema

## Error Signature

- **keyword**: `defs`
- **case**: `validate definition against metaschema`
- **normalized got**: `error: bundling error: failed to resolve internal refs for compliance://<draft>/defs/group_<n>: encountered unresolved non-local $ref "<metaschema-url>"`

## Root Cause

Same as issue-09. The bundler's `processRef` in `cli/bundler/bundler.go` detects metaschema URLs via `isMetaschema()` but returned the `$ref` object unchanged, leaving an external URL in the bundled output. `resolveInternalRefs` then rejected it as a non-local `$ref`. Fixed by replacing metaschema `$ref` with empty schema `{}`.

## Baseline (20 failures)

| # | Adapter | Draft | Test | Expected | Got | Report Path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript arktype | draft2019-09 | valid definition schema | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft2019-09.json` |
| 2 | typescript arktype | draft2019-09 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft2019-09.json` |
| 3 | typescript arktype | draft2020-12 | valid definition schema | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft2020-12.json` |
| 4 | typescript arktype | draft2020-12 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft2020-12.json` |
| 5 | typescript effect | draft2019-09 | valid definition schema | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft2019-09.json` |
| 6 | typescript effect | draft2019-09 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft2019-09.json` |
| 7 | typescript effect | draft2020-12 | valid definition schema | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft2020-12.json` |
| 8 | typescript effect | draft2020-12 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft2020-12.json` |
| 9 | python pydantic | draft2019-09 | valid definition schema | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft2019-09.json` |
| 10 | python pydantic | draft2019-09 | invalid definition schema | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft2019-09.json` |
| 11 | python pydantic | draft2020-12 | valid definition schema | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft2020-12.json` |
| 12 | python pydantic | draft2020-12 | invalid definition schema | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft2020-12.json` |
| 13 | typescript valibot | draft2019-09 | valid definition schema | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft2019-09.json` |
| 14 | typescript valibot | draft2019-09 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft2019-09.json` |
| 15 | typescript valibot | draft2020-12 | valid definition schema | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft2020-12.json` |
| 16 | typescript valibot | draft2020-12 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft2020-12.json` |
| 17 | typescript zod | draft2019-09 | valid definition schema | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft2019-09.json` |
| 18 | typescript zod | draft2019-09 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft2019-09.json` |
| 19 | typescript zod | draft2020-12 | valid definition schema | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft2020-12.json` |
| 20 | typescript zod | draft2020-12 | invalid definition schema | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft2020-12.json` |
