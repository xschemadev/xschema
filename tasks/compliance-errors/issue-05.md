# issue-05: ref / remote ref, containing refs itself

## Error Signature

- **keyword**: `ref`
- **case**: `remote ref, containing refs itself`
- **normalized got**: `error: bundling error: failed to resolve internal refs for compliance://<draft>/ref/group_<n>: encountered unresolved non-local $ref "<metaschema-url>"`

## Root Cause

The bundler's `processRef` in `cli/bundler/bundler.go` detects metaschema URLs via `isMetaschema()` but returned the `$ref` object unchanged, leaving an external URL in the bundled output. The subsequent `resolveInternalRefs` in `cli/processor/local_refs.go` rejects any non-local `$ref` (not starting with `#`), producing this error.

Fixed by issue-09: metaschema `$ref`s are now replaced with empty schema `{}` in the bundler, and root-level metaschema refs are detected by `checkMetaschemaRef` in the processor and classified as `UnsupportedKeywordError`.

## Baseline (60 failures)

| # | Adapter | Draft | Test | Expected | Got | Report Path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript arktype | draft2019-09 | remote ref valid | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft2019-09.json` |
| 2 | typescript arktype | draft2019-09 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft2019-09.json` |
| 3 | typescript arktype | draft2020-12 | remote ref valid | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft2020-12.json` |
| 4 | typescript arktype | draft2020-12 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft2020-12.json` |
| 5 | typescript arktype | draft3 | remote ref valid | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft3.json` |
| 6 | typescript arktype | draft3 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft3.json` |
| 7 | typescript arktype | draft4 | remote ref valid | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft4.json` |
| 8 | typescript arktype | draft4 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft4.json` |
| 9 | typescript arktype | draft6 | remote ref valid | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft6.json` |
| 10 | typescript arktype | draft6 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft6.json` |
| 11 | typescript arktype | draft7 | remote ref valid | `true` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft7.json` |
| 12 | typescript arktype | draft7 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/arktype/compliance/results/draft7.json` |
| 13 | typescript effect | draft2019-09 | remote ref valid | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft2019-09.json` |
| 14 | typescript effect | draft2019-09 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft2019-09.json` |
| 15 | typescript effect | draft2020-12 | remote ref valid | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft2020-12.json` |
| 16 | typescript effect | draft2020-12 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft2020-12.json` |
| 17 | typescript effect | draft3 | remote ref valid | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft3.json` |
| 18 | typescript effect | draft3 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft3.json` |
| 19 | typescript effect | draft4 | remote ref valid | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft4.json` |
| 20 | typescript effect | draft4 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft4.json` |
| 21 | typescript effect | draft6 | remote ref valid | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft6.json` |
| 22 | typescript effect | draft6 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft6.json` |
| 23 | typescript effect | draft7 | remote ref valid | `true` | `error` | `typescript/packages/adapters/effect/compliance/results/draft7.json` |
| 24 | typescript effect | draft7 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/effect/compliance/results/draft7.json` |
| 25 | python pydantic | draft2019-09 | remote ref valid | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft2019-09.json` |
| 26 | python pydantic | draft2019-09 | remote ref invalid | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft2019-09.json` |
| 27 | python pydantic | draft2020-12 | remote ref valid | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft2020-12.json` |
| 28 | python pydantic | draft2020-12 | remote ref invalid | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft2020-12.json` |
| 29 | python pydantic | draft3 | remote ref valid | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft3.json` |
| 30 | python pydantic | draft3 | remote ref invalid | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft3.json` |
| 31 | python pydantic | draft4 | remote ref valid | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft4.json` |
| 32 | python pydantic | draft4 | remote ref invalid | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft4.json` |
| 33 | python pydantic | draft6 | remote ref valid | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft6.json` |
| 34 | python pydantic | draft6 | remote ref invalid | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft6.json` |
| 35 | python pydantic | draft7 | remote ref valid | `true` | `error` | `python/packages/adapters/pydantic/compliance/results/draft7.json` |
| 36 | python pydantic | draft7 | remote ref invalid | `false` | `error` | `python/packages/adapters/pydantic/compliance/results/draft7.json` |
| 37 | typescript valibot | draft2019-09 | remote ref valid | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft2019-09.json` |
| 38 | typescript valibot | draft2019-09 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft2019-09.json` |
| 39 | typescript valibot | draft2020-12 | remote ref valid | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft2020-12.json` |
| 40 | typescript valibot | draft2020-12 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft2020-12.json` |
| 41 | typescript valibot | draft3 | remote ref valid | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft3.json` |
| 42 | typescript valibot | draft3 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft3.json` |
| 43 | typescript valibot | draft4 | remote ref valid | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft4.json` |
| 44 | typescript valibot | draft4 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft4.json` |
| 45 | typescript valibot | draft6 | remote ref valid | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft6.json` |
| 46 | typescript valibot | draft6 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft6.json` |
| 47 | typescript valibot | draft7 | remote ref valid | `true` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft7.json` |
| 48 | typescript valibot | draft7 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/valibot/compliance/results/draft7.json` |
| 49 | typescript zod | draft2019-09 | remote ref valid | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft2019-09.json` |
| 50 | typescript zod | draft2019-09 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft2019-09.json` |
| 51 | typescript zod | draft2020-12 | remote ref valid | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft2020-12.json` |
| 52 | typescript zod | draft2020-12 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft2020-12.json` |
| 53 | typescript zod | draft3 | remote ref valid | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft3.json` |
| 54 | typescript zod | draft3 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft3.json` |
| 55 | typescript zod | draft4 | remote ref valid | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft4.json` |
| 56 | typescript zod | draft4 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft4.json` |
| 57 | typescript zod | draft6 | remote ref valid | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft6.json` |
| 58 | typescript zod | draft6 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft6.json` |
| 59 | typescript zod | draft7 | remote ref valid | `true` | `error` | `typescript/packages/adapters/zod/compliance/results/draft7.json` |
| 60 | typescript zod | draft7 | remote ref invalid | `false` | `error` | `typescript/packages/adapters/zod/compliance/results/draft7.json` |
