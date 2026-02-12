# issue-08: ref / Recursive references between schemas

## Error Signature

- **Keyword**: `ref`
- **Case**: `Recursive references between schemas`
- **Normalized got**: `error: bundling error: failed to resolve internal refs for compliance://<draft>/ref/group_<n>: recursive local $ref "#" is not supported`

## Root Cause

Two-layer problem:

1. **Go processor** (already fixed by issue-01): `resolveInternalRefs` errored on recursive `$ref: "#"` instead of passing it through to adapters. This was fixed in issue-01 by returning the `$ref` object unchanged for cycle-detected refs.

2. **Pydantic adapter**: `render_ref` in `python/packages/adapters/pydantic/src/xschema_pydantic/renderer.py` emitted `Any` for recursive `$ref: "#"` refs. This meant no validation occurred at recursive positions — "valid tree" passed but "invalid tree" also passed (accepted invalid data).

## Fix

1. **Go processor** — already fixed by issue-01 (return `$ref` object unchanged for cycles)
2. **TS adapters** (zod, effect, valibot, arktype) — already fixed by issue-01 (`z.lazy()`, `S.suspend()`, `v.lazy()`, `type.unknown.narrow()`)
3. **Pydantic adapter** (`python/packages/adapters/pydantic/src/xschema_pydantic/renderer.py`):
   - Added `set_root_class_name(name, is_class)` to track the root class name for forward ref resolution
   - `render_ref` now emits `'ClassName'` (Pydantic forward reference) instead of `Any` when the root schema produces a typed BaseModel class (`type: "object"`)
   - For typeless object schemas (no explicit `type`), falls back to `Any` since non-object values must pass validation at recursive positions
4. **Pydantic converter** (`python/packages/adapters/pydantic/src/xschema_pydantic/converter.py`):
   - Sets `root_is_typed_object` before rendering to enable forward ref resolution
   - Emits `ClassName.model_rebuild()` after class definition when forward refs are used, to resolve them at runtime

## Baseline (36 failures)

| # | Adapter | Draft | Test | Expected | Got | Report Path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript arktype | draft4 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 2 | typescript arktype | draft4 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 3 | typescript arktype | draft6 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 4 | typescript arktype | draft6 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 5 | typescript arktype | draft7 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 6 | typescript arktype | draft7 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 7 | typescript arktype | draft2019-09 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 8 | typescript effect | draft4 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 9 | typescript effect | draft4 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 10 | typescript effect | draft6 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 11 | typescript effect | draft6 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 12 | typescript effect | draft7 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 13 | typescript effect | draft7 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 14 | typescript effect | draft2019-09 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 15 | python pydantic | draft4 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_11: recursive local $ref "#" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 16 | python pydantic | draft4 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_11: recursive local $ref "#" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 17 | python pydantic | draft6 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_11: recursive local $ref "#" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 18 | python pydantic | draft6 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_11: recursive local $ref "#" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 19 | python pydantic | draft7 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 20 | python pydantic | draft7 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 21 | python pydantic | draft2019-09 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_11: recursive local $ref "#" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 22 | python pydantic | draft2020-12 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_11: recursive local $ref "#" is not supported | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 23 | typescript valibot | draft4 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 24 | typescript valibot | draft4 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 25 | typescript valibot | draft6 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 26 | typescript valibot | draft6 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 27 | typescript valibot | draft7 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 28 | typescript valibot | draft7 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 29 | typescript valibot | draft2019-09 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 30 | typescript valibot | draft2020-12 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 31 | typescript zod | draft4 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 32 | typescript zod | draft4 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 33 | typescript zod | draft6 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 34 | typescript zod | draft6 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_11: recursive local $ref "#" is not supported | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 35 | typescript zod | draft7 | valid tree | valid | error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 36 | typescript zod | draft7 | invalid tree | invalid | error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported | typescript/packages/adapters/zod/compliance/results/REPORT.md |
