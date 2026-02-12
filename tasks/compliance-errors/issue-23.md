# issue-23: ref / order of evaluation: $id and $anchor and $ref

## Error Signature

- **Keyword**: `ref`
- **Case**: `order of evaluation: $id and $anchor and $ref`
- **Normalized Got**: `false`

## Root Cause

The bundler's `collectIDsAndAnchors` stores `$anchor` values in a flat `map[string]string` keyed only by anchor name, without scoping to the `$id`-based base URI. When two sub-schemas under different `$id`s define the same `$anchor` name (e.g., both `$defs/bigint` and `$defs/smallint` define `$anchor: "bigint"`), the last one collected overwrites the first. `processRef` then resolves `$ref: "#bigint"` to the wrong definition.

Per the JSON Schema spec, `$anchor` is scoped to the base URI of the schema resource that defines it. `$ref: "#bigint"` from the root should resolve to the anchor under the root's base URI, not one under a different `$id`.

## Baseline (7 failures)

Note: PRD originally specified 9 failures across 5 adapters, but actual reports show 7 — pydantic has no `$anchor` failures for this case, and valibot only fails in draft2019-09 (not draft2020-12).

| # | Adapter | Draft | Test | Expected | Got | Report Path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript arktype | draft2019-09 | data is valid against first definition | valid | false | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 2 | typescript arktype | draft2020-12 | data is valid against first definition | valid | false | typescript/packages/adapters/arktype/compliance/results/REPORT.md |
| 3 | typescript effect | draft2019-09 | data is valid against first definition | valid | false | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 4 | typescript effect | draft2020-12 | data is valid against first definition | valid | false | typescript/packages/adapters/effect/compliance/results/REPORT.md |
| 5 | typescript valibot | draft2019-09 | data is valid against first definition | valid | false | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 6 | typescript zod | draft2019-09 | data is valid against first definition | valid | false | typescript/packages/adapters/zod/compliance/results/REPORT.md |
| 7 | typescript zod | draft2020-12 | data is valid against first definition | valid | false | typescript/packages/adapters/zod/compliance/results/REPORT.md |

## Fix

Changed `bundleContext.anchors` from flat `map[string]string` to URI-scoped `map[string]map[string]string` in `cli/bundler/bundler.go`. Added `storeAnchor(baseURI, anchor, path)` and `lookupAnchor(baseURI, anchor)` helpers. All 5 anchor access sites updated to use scoped storage/lookup.
