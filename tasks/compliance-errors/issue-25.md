# issue-25: ref / ref creates new scope when adjacent to keywords

## Error Signature

- **Keyword**: `ref`
- **Case**: `ref creates new scope when adjacent to keywords`
- **Test**: `referenced subschema doesn't see annotations from properties`
- **Normalized Got**: `true`
- **Expected**: `invalid` (false)

## Baseline (2 failures)

| # | Adapter | Draft | Test | Expected | Got | Report Path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | python pydantic | draft2019-09 | referenced subschema doesn't see annotations from properties | invalid | true | python/packages/adapters/pydantic/compliance/results/REPORT.md |
| 2 | python pydantic | draft2020-12 | referenced subschema doesn't see annotations from properties | invalid | true | python/packages/adapters/pydantic/compliance/results/REPORT.md |

## Root Cause

The Go processor's `resolveInternalRefs` wraps `$ref` target + sibling keywords in `allOf`:
```json
{"allOf": [{"unevaluatedProperties": false}, {"properties": {"prop1": {"type": "string"}}}]}
```

Both allOf sub-schemas parse as `ObjectNode` in the Python core parser. The pydantic renderer's `render_intersection` detects all-object allOf and calls `_merge_object_schemas`, which merges properties from all sub-schemas into one BaseModel class. However, `_merge_object_schemas` completely ignores `unevaluated_properties` from the input ObjectNodes — the `unevaluated_properties=False` on allOf[0] (which should forbid all properties since it declares none) is silently dropped.

## Fix

In `render_intersection`, detect when any allOf sub-schema has `unevaluated_properties is False` and skip the static merge path. Route to the runtime validation path (BeforeValidator + TypeAdapter) instead, which validates each sub-schema independently — preserving per-schema evaluation scope isolation.
