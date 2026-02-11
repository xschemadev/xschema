# ArkType Adapter Compliance Baseline

Captured: 2026-02-11

## Summary by Draft

| Draft | Passed | Failed | Skipped | Unsupported | Coverage |
| ----- | ------ | ------ | ------- | ----------- | -------- |
| draft3 | 397 | 0 | 0 | 0 | 100.0% |
| draft4 | 526 | 0 | 0 | 0 | 100.0% |
| draft6 | 707 | 0 | 0 | 0 | 100.0% |
| draft7 | 783 | 0 | 0 | 0 | 100.0% |
| draft2019-09 | 906 | 0 | 0 | 178 | 100.0% |
| draft2020-12 | 919 | 0 | 0 | 201 | 100.0% |

**All 6 drafts at 100% runtime compliance.** No failures, no skips.

## Unsupported Features

- `$recursiveRef/$recursiveAnchor` (draft2019-09): 40 tests - requires runtime scope tracking
- `$dynamicRef/$dynamicAnchor` (draft2020-12): 48 tests - requires runtime scope tracking
- `unevaluatedItems` inside applicators (draft2019-09): 40 tests - cousins problem / annotation tracking
- `unevaluatedItems` inside applicators (draft2020-12): 55 tests - cousins problem / annotation tracking
- `unevaluatedProperties` inside applicators (draft2019-09): 98 tests - cousins/cyclic/annotation tracking
- `unevaluatedProperties` inside applicators (draft2020-12): 98 tests - cousins/cyclic/annotation tracking

## Keywords by Draft (draft2020-12, representative)

| Keyword | Pass/Total | Status |
| ------- | ---------- | ------ |
| additionalProperties | 21/21 | pass |
| allOf | 30/30 | pass |
| anyOf | 18/18 | pass |
| boolean_schema | 18/18 | pass |
| const | 54/54 | pass |
| contains | 21/21 | pass |
| content | 18/18 | pass |
| default | 7/7 | pass |
| dependentRequired | 20/20 | pass |
| dependentSchemas | 20/20 | pass |
| enum | 45/45 | pass |
| exclusiveMaximum | 4/4 | pass |
| exclusiveMinimum | 4/4 | pass |
| format | 133/133 | pass |
| if-then-else | 26/26 | pass |
| maxContains | 12/12 | pass |
| maxItems | 6/6 | pass |
| maxLength | 7/7 | pass |
| maxProperties | 10/10 | pass |
| maximum | 8/8 | pass |
| minContains | 28/28 | pass |
| minItems | 6/6 | pass |
| minLength | 7/7 | pass |
| minProperties | 8/8 | pass |
| minimum | 11/11 | pass |
| multipleOf | 10/10 | pass |
| not | 38/38 | pass |
| oneOf | 27/27 | pass |
| pattern | 9/9 | pass |
| patternProperties | 23/23 | pass |
| prefixItems | 11/11 | pass |
| properties | 28/28 | pass |
| propertyNames | 20/20 | pass |
| required | 16/16 | pass |
| type | 80/80 | pass |
| unevaluatedItems | 14/14 | pass |
| unevaluatedProperties | 27/27 | pass |
| uniqueItems | 69/69 | pass |
| vocabulary | 5/5 | pass |

## Type Fidelity Issues (type.unknown usage in renderer)

The following renderer paths use `type.unknown` as the base type, causing `typeof schema.infer`
to degrade to `unknown` instead of a narrower TypeScript type:

### 1. `renderOneOf` (renderer.ts:705)
- **Pattern**: `type.unknown.narrow(...)` for oneOf with 2+ schemas
- **Impact**: `typeof schema.infer` -> `unknown` for all oneOf schemas
- **Example**: `{ "oneOf": [{ "type": "string" }, { "type": "number" }] }` -> `unknown` instead of `string | number`

### 2. `renderNot` (renderer.ts:716)
- **Pattern**: `type.unknown.narrow(...)` for not schemas
- **Impact**: `typeof schema.infer` -> `unknown` for all not schemas
- **Example**: `{ "not": { "type": "string" } }` -> `unknown` (acceptable - not can't narrow positively)

### 3. `renderConditional` (renderer.ts:772)
- **Pattern**: `type.unknown.narrow(...)` for if/then/else
- **Impact**: `typeof schema.infer` -> `unknown` for all conditional schemas
- **Example**: `{ "if": {...}, "then": {...}, "else": {...} }` -> `unknown`

### 4. `renderTypeGuarded` (renderer.ts:830)
- **Pattern**: `type.unknown.narrow(...)` for type-guarded schemas
- **Impact**: `typeof schema.infer` -> `unknown` for all typeGuarded schemas
- **Example**: `{ "properties": {"a": {"type": "string"}}, "minLength": 1 }` -> `unknown`

### 5. `renderTuple` (renderer.ts:567)
- **Pattern**: `type.unknown.array().narrow(...)` for tuple schemas
- **Impact**: `typeof schema.infer` -> `unknown[]` instead of `[string, number, ...]`
- **Example**: `{ "prefixItems": [{"type": "string"}, {"type": "number"}] }` -> `unknown[]`

### 6. `renderLiteral` for complex values (renderer.ts:730)
- **Pattern**: `type.unknown.narrow(...)` for non-primitive const values
- **Impact**: `typeof schema.infer` -> `unknown` instead of the literal type
- **Example**: `{ "const": [1, 2, 3] }` -> `unknown`

### 7. `renderEnum` for complex values (renderer.ts:754)
- **Pattern**: `type.unknown.narrow(...)` for enums with non-primitive values
- **Impact**: `typeof schema.infer` -> `unknown`
- **Example**: `{ "enum": [{"a": 1}, {"b": 2}] }` -> `unknown`

### 8. `renderObjectWithProtoProps` (renderer.ts:405)
- **Pattern**: `type.object.narrow(...)` for objects with prototype property names
- **Impact**: `typeof schema.infer` -> `object` (overly broad)
- **Note**: Slightly better than zod (object vs any), but still loses property info
