# Pydantic Adapter Compliance Baseline

Captured: 2026-02-11

## Summary by Draft

| Draft | Passed | Failed | Skipped | Unsupported | Coverage |
| ----- | ------ | ------ | ------- | ----------- | -------- |
| draft3 | 407 | 0 | 0 | 0 | 100.0% |
| draft4 | 566 | 0 | 0 | 0 | 100.0% |
| draft6 | 760 | 0 | 0 | 0 | 100.0% |
| draft7 | 836 | 0 | 0 | 0 | 100.0% |
| draft2019-09 | 936 | 0 | 0 | 178 | 100.0% |
| draft2020-12 | 950 | 0 | 0 | 201 | 100.0% |

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
| items | 29/29 | pass |
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

## Type Fidelity Issues (Any usage in renderer)

The following renderer paths use `Any` as the type expression, causing `TypeAdapter[Any]`
degradation instead of a narrower Python type:

### 1. `render_oneof` (renderer.py:2300)
- **Pattern**: `Annotated[Any, BeforeValidator(validator)]` for oneOf with 2+ schemas
- **Impact**: All oneOf schemas produce `TypeAdapter[Any]`
- **Example**: `{ "oneOf": [{ "type": "string" }, { "type": "number" }] }` → `TypeAdapter[Any]` instead of `TypeAdapter[str | int]`

### 2. `render_not` (renderer.py:2345)
- **Pattern**: `Annotated[Any, BeforeValidator(validator)]` for not schemas
- **Impact**: All not schemas produce `TypeAdapter[Any]`
- **Best possible**: `TypeAdapter[Any]` (negation can't narrow positively) — but could use `object` to signal "validated"

### 3. `render_conditional` (renderer.py:2437)
- **Pattern**: `Annotated[Any, BeforeValidator(validator)]` for if/then/else
- **Impact**: All conditional schemas produce `TypeAdapter[Any]`
- **Best possible**: Union of then/else types

### 4. `render_type_guarded` (renderer.py:2493)
- **Pattern**: `Annotated[Any, BeforeValidator(validator)]` for multi-type guard schemas
- **Impact**: Type-guarded schemas produce `TypeAdapter[Any]`

### 5. `render_intersection` mixed allOf (renderer.py:2093)
- **Pattern**: `Annotated[Any, BeforeValidator(validator)]` for mixed-type allOf
- **Impact**: Mixed allOf schemas (e.g. TSConfig-like) produce `TypeAdapter[Any]`
- **Example**: allOf with object + constraints → `TypeAdapter[Any]` instead of `TypeAdapter[SomeModel]`

### 6. `render_tuple` (renderer.py:1118, 1163, 1209)
- **Pattern**: `Annotated[tuple[Any, ...], BeforeValidator(validator)]` for prefix/closed/open tuples
- **Impact**: All non-trivial tuple schemas lose positional type info
- **Example**: `prefixItems: [{type: "string"}, {type: "number"}]` → `TypeAdapter[tuple[Any, ...]]` instead of `TypeAdapter[tuple[str, int]]`

### 7. `render_literal` complex const (renderer.py:816-818)
- **Pattern**: `Annotated[Any, BeforeValidator(_make_const_validator(...))]` for bool/0/1/list/dict const values
- **Impact**: Complex const values produce `TypeAdapter[Any]`

### 8. `render_enum` complex values (renderer.py:865-866)
- **Pattern**: `Annotated[Any, BeforeValidator(_make_enum_validator(...))]` for enums with bools/0/1/lists/dicts
- **Impact**: Mixed enums produce `TypeAdapter[Any]`

### 9. `render` fallback (renderer.py:408-410)
- **Pattern**: Bare `Any` for unrecognized IR node kinds
- **Impact**: Defensive fallback, should never trigger in practice

### 10. `render_ref` unresolved (renderer.py:2543)
- **Pattern**: Bare `Any` for unresolved `$ref` (including recursive `$ref: "#"`)
- **Impact**: Recursive/unresolved refs produce `TypeAdapter[Any]`
