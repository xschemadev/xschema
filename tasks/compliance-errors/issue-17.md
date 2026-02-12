# issue-17: refRemote / fragment within remote ref

## Error Signature

- **keyword**: refRemote
- **case**: fragment within remote ref
- **normalized got**: `error: bundling error: failed to resolve internal refs for compliance://<draft>/refRemote/group_<n>: recursive local $ref "#/$defs/localhost_1234_<draft>_subSchemas_json__integer" is not supported`

## Root Cause

False positive recursion detection in `resolveInternalRefs` (`cli/processor/local_refs.go`). The bundler correctly fetched the remote schema at `localhost:1234/<draft>/subSchemas.json`, resolved the `#/integer` fragment to `$defs/localhost_1234_<draft>_subSchemas_json__integer`, and flattened defs. But `resolveRefObject` marked the target as "resolving", then when processing sibling `$defs` (draft2019-09+), the same flattened target was encountered again, triggering false recursion guard.

## Fix

Already resolved by prior fixes:
- issue-01: changed cycle detection from hard error to pass-through (preserves `$ref` for adapter lazy/suspend handling)
- issue-15: fixed `processObject` to process sibling keys after `processRef` returns, ensuring `$defs` children get fully resolved

## Baseline (20 failures)

| # | adapter | draft | test | expected | got | report path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript arktype | draft2019-09 | remote fragment valid | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 2 | typescript arktype | draft2019-09 | remote fragment invalid | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 3 | typescript arktype | draft2020-12 | remote fragment valid | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 4 | typescript arktype | draft2020-12 | remote fragment invalid | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 5 | typescript effect | draft2019-09 | remote fragment valid | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 6 | typescript effect | draft2019-09 | remote fragment invalid | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 7 | typescript effect | draft2020-12 | remote fragment valid | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 8 | typescript effect | draft2020-12 | remote fragment invalid | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 9 | python pydantic | draft2019-09 | remote fragment valid | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 10 | python pydantic | draft2019-09 | remote fragment invalid | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 11 | python pydantic | draft2020-12 | remote fragment valid | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 12 | python pydantic | draft2020-12 | remote fragment invalid | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 13 | typescript valibot | draft2019-09 | remote fragment valid | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 14 | typescript valibot | draft2019-09 | remote fragment invalid | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 15 | typescript valibot | draft2020-12 | remote fragment valid | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 16 | typescript valibot | draft2020-12 | remote fragment invalid | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 17 | typescript zod | draft2019-09 | remote fragment valid | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 18 | typescript zod | draft2019-09 | remote fragment invalid | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 19 | typescript zod | draft2020-12 | remote fragment valid | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 20 | typescript zod | draft2020-12 | remote fragment invalid | invalid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_3: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported | typescript/packages/adapters/zod/compliance/results/REPORT.md |
