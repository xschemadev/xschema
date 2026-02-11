# Zod Adapter Compliance Baseline

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

## Type Fidelity Issues (z.any() usage in renderer)

The following renderer paths use `z.any()` as the base type, causing `z.infer<typeof schema>`
to degrade to `any` instead of a narrower TypeScript type:

### 1. `renderOneOf` (renderer.ts:726)
- **Pattern**: `z.any().superRefine(...)` for oneOf with 2+ schemas
- **Impact**: `z.infer` → `any` for all oneOf schemas
- **Example**: `{ "oneOf": [{ "type": "string" }, { "type": "number" }] }` → `any` instead of `string | number`

### 2. `renderNot` (renderer.ts:746)
- **Pattern**: `z.any().refine(...)` for not schemas
- **Impact**: `z.infer` → `any` for all not schemas
- **Example**: `{ "not": { "type": "string" } }` → `any` instead of `unknown`

### 3. `renderConditional` (renderer.ts:812)
- **Pattern**: `z.any().superRefine(...)` for if/then/else
- **Impact**: `z.infer` → `any` for all conditional schemas
- **Example**: `{ "if": {...}, "then": {...}, "else": {...} }` → `any` instead of `unknown`

### 4. `renderTypeGuarded` (renderer.ts:895)
- **Pattern**: `z.any().superRefine(...)` for type-guarded schemas
- **Impact**: `z.infer` → `any` for all typeGuarded schemas
- **Example**: `{ "properties": {"a": {"type": "string"}}, "minLength": 1 }` → `any` instead of `unknown`

### 5. `renderTuple` (renderer.ts:579)
- **Pattern**: `z.array(z.any()).superRefine(...)` for tuple schemas
- **Impact**: `z.infer` → `any[]` instead of `[string, number, ...]`
- **Example**: `{ "prefixItems": [{"type": "string"}, {"type": "number"}] }` → `any[]`

### 6. `renderLiteral` for complex values (renderer.ts:756-759)
- **Pattern**: `z.array(z.any()).refine(...)` or `z.object({}).passthrough().refine(...)`
- **Impact**: `z.infer` → `any[]` or `{ [x: string]: unknown }` instead of the literal type
- **Example**: `{ "const": [1, 2, 3] }` → `any[]` instead of readonly tuple type

### 7. `renderEnum` for complex values (renderer.ts:790)
- **Pattern**: `z.any().refine(...)` for enums with non-primitive values
- **Impact**: `z.infer` → `any`
- **Example**: `{ "enum": [{"a": 1}, {"b": 2}] }` → `any`

### 8. `renderObjectWithProtoProps` (renderer.ts:413)
- **Pattern**: `z.any().superRefine(...)` for objects with prototype property names
- **Impact**: `z.infer` → `any` for objects with `__proto__`, `constructor`, `prototype` keys
- **Note**: This is a correctness issue too - non-objects should be rejected
