# @xschemadev/typescript Compliance Report

Generated: 2026-01-10T19:15:44Z

## Type-Only Adapter

This adapter generates **type definitions only** - no runtime validation code is produced.

Runtime validation tests are **skipped** because:
1. Compliance tests require runtime validation (checking if data matches a schema)
2. Types exist only at compile-time and are erased at runtime
3. The generated code is validated via **TypeScript type-checking** (`tsc --noEmit`)

**Note:** Any failures shown below indicate TypeScript compilation errors in the generated types, not runtime validation failures.

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2020-12 | 0 | 0 | 1030 | N/A (type-only) |
| draft2019-09 | 0 | 0 | 1027 | N/A (type-only) |
| draft7 | 0 | 0 | 872 | N/A (type-only) |
| draft6 | 0 | 0 | 792 | N/A (type-only) |
| draft4 | 0 | 0 | 591 | N/A (type-only) |
| draft3 | 0 | 0 | 426 | N/A (type-only) |

## draft2020-12

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ❌ | 0/21 |
| allOf | ❌ | 0/30 |
| anchor | ✅ | 0/0 |
| anyOf | ❌ | 0/18 |
| boolean_schema | ❌ | 0/18 |
| const | ❌ | 0/54 |
| contains | ❌ | 0/21 |
| content | ❌ | 0/18 |
| default | ❌ | 0/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ❌ | 0/20 |
| dependentSchemas | ❌ | 0/20 |
| dynamicRef | ❌ | 0/19 |
| enum | ❌ | 0/45 |
| exclusiveMaximum | ❌ | 0/4 |
| exclusiveMinimum | ❌ | 0/4 |
| format | ❌ | 0/133 |
| if-then-else | ❌ | 0/30 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ❌ | 0/29 |
| maxContains | ❌ | 0/12 |
| maxItems | ❌ | 0/6 |
| maxLength | ❌ | 0/7 |
| maxProperties | ❌ | 0/10 |
| maximum | ❌ | 0/8 |
| minContains | ❌ | 0/28 |
| minItems | ❌ | 0/6 |
| minLength | ❌ | 0/7 |
| minProperties | ❌ | 0/8 |
| minimum | ❌ | 0/11 |
| multipleOf | ❌ | 0/10 |
| not | ❌ | 0/38 |
| oneOf | ❌ | 0/27 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/23 |
| prefixItems | ❌ | 0/11 |
| properties | ❌ | 0/28 |
| propertyNames | ❌ | 0/20 |
| ref | ❌ | 0/35 |
| refRemote | ❌ | 0/17 |
| required | ❌ | 0/16 |
| type | ❌ | 0/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ❌ | 0/46 |
| uniqueItems | ❌ | 0/69 |
| vocabulary | ❌ | 0/5 |

## draft2019-09

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ❌ | 0/19 |
| additionalProperties | ❌ | 0/21 |
| allOf | ❌ | 0/30 |
| anchor | ✅ | 0/0 |
| anyOf | ❌ | 0/18 |
| boolean_schema | ❌ | 0/18 |
| const | ❌ | 0/54 |
| contains | ❌ | 0/21 |
| content | ❌ | 0/18 |
| default | ❌ | 0/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ❌ | 0/20 |
| dependentSchemas | ❌ | 0/20 |
| enum | ❌ | 0/45 |
| exclusiveMaximum | ❌ | 0/4 |
| exclusiveMinimum | ❌ | 0/4 |
| format | ❌ | 0/114 |
| if-then-else | ❌ | 0/30 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ❌ | 0/28 |
| maxContains | ❌ | 0/12 |
| maxItems | ❌ | 0/6 |
| maxLength | ❌ | 0/7 |
| maxProperties | ❌ | 0/10 |
| maximum | ❌ | 0/8 |
| minContains | ❌ | 0/28 |
| minItems | ❌ | 0/6 |
| minLength | ❌ | 0/7 |
| minProperties | ❌ | 0/8 |
| minimum | ❌ | 0/11 |
| multipleOf | ❌ | 0/10 |
| not | ❌ | 0/38 |
| oneOf | ❌ | 0/27 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/23 |
| properties | ❌ | 0/28 |
| propertyNames | ❌ | 0/20 |
| recursiveRef | ❌ | 0/30 |
| ref | ❌ | 0/35 |
| refRemote | ❌ | 0/17 |
| required | ❌ | 0/16 |
| type | ❌ | 0/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ❌ | 0/44 |
| uniqueItems | ❌ | 0/69 |
| vocabulary | ❌ | 0/5 |

## draft7

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ❌ | 0/19 |
| additionalProperties | ❌ | 0/16 |
| allOf | ❌ | 0/30 |
| anyOf | ❌ | 0/18 |
| boolean_schema | ❌ | 0/18 |
| const | ❌ | 0/54 |
| contains | ❌ | 0/21 |
| default | ❌ | 0/7 |
| definitions | ❌ | 0/2 |
| dependencies | ❌ | 0/36 |
| enum | ❌ | 0/45 |
| exclusiveMaximum | ❌ | 0/4 |
| exclusiveMinimum | ❌ | 0/4 |
| format | ❌ | 0/102 |
| if-then-else | ❌ | 0/30 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ❌ | 0/28 |
| maxItems | ❌ | 0/6 |
| maxLength | ❌ | 0/7 |
| maxProperties | ❌ | 0/10 |
| maximum | ❌ | 0/8 |
| minItems | ❌ | 0/6 |
| minLength | ❌ | 0/7 |
| minProperties | ❌ | 0/8 |
| minimum | ❌ | 0/11 |
| multipleOf | ❌ | 0/10 |
| not | ❌ | 0/38 |
| oneOf | ❌ | 0/27 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/23 |
| properties | ❌ | 0/28 |
| propertyNames | ❌ | 0/20 |
| ref | ❌ | 0/36 |
| refRemote | ❌ | 0/17 |
| required | ❌ | 0/16 |
| type | ❌ | 0/80 |
| uniqueItems | ❌ | 0/69 |

## draft6

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ❌ | 0/19 |
| additionalProperties | ❌ | 0/16 |
| allOf | ❌ | 0/30 |
| anyOf | ❌ | 0/18 |
| boolean_schema | ❌ | 0/18 |
| const | ❌ | 0/54 |
| contains | ❌ | 0/19 |
| default | ❌ | 0/7 |
| definitions | ❌ | 0/2 |
| dependencies | ❌ | 0/36 |
| enum | ❌ | 0/45 |
| exclusiveMaximum | ❌ | 0/4 |
| exclusiveMinimum | ❌ | 0/4 |
| format | ❌ | 0/54 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ❌ | 0/28 |
| maxItems | ❌ | 0/6 |
| maxLength | ❌ | 0/7 |
| maxProperties | ❌ | 0/10 |
| maximum | ❌ | 0/8 |
| minItems | ❌ | 0/6 |
| minLength | ❌ | 0/7 |
| minProperties | ❌ | 0/8 |
| minimum | ❌ | 0/11 |
| multipleOf | ❌ | 0/10 |
| not | ❌ | 0/38 |
| oneOf | ❌ | 0/27 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/23 |
| properties | ❌ | 0/28 |
| propertyNames | ❌ | 0/20 |
| ref | ❌ | 0/36 |
| refRemote | ❌ | 0/17 |
| required | ❌ | 0/16 |
| type | ❌ | 0/80 |
| uniqueItems | ❌ | 0/69 |

## draft4

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ❌ | 0/17 |
| additionalProperties | ❌ | 0/16 |
| allOf | ❌ | 0/27 |
| anyOf | ❌ | 0/15 |
| default | ❌ | 0/7 |
| definitions | ❌ | 0/2 |
| dependencies | ❌ | 0/29 |
| enum | ❌ | 0/49 |
| format | ❌ | 0/36 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ❌ | 0/21 |
| maxItems | ❌ | 0/4 |
| maxLength | ❌ | 0/5 |
| maxProperties | ❌ | 0/8 |
| maximum | ❌ | 0/14 |
| minItems | ❌ | 0/4 |
| minLength | ❌ | 0/5 |
| minProperties | ❌ | 0/6 |
| minimum | ❌ | 0/17 |
| multipleOf | ❌ | 0/10 |
| not | ❌ | 0/20 |
| oneOf | ❌ | 0/23 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/18 |
| properties | ❌ | 0/24 |
| ref | ❌ | 0/25 |
| refRemote | ❌ | 0/15 |
| required | ❌ | 0/15 |
| type | ❌ | 0/79 |
| uniqueItems | ❌ | 0/69 |

## draft3

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ❌ | 0/14 |
| additionalProperties | ❌ | 0/16 |
| default | ❌ | 0/7 |
| dependencies | ❌ | 0/18 |
| disallow | ❌ | 0/9 |
| divisibleBy | ❌ | 0/8 |
| enum | ❌ | 0/16 |
| extends | ❌ | 0/10 |
| format | ❌ | 0/60 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ❌ | 0/7 |
| maxItems | ❌ | 0/4 |
| maxLength | ❌ | 0/5 |
| maximum | ❌ | 0/14 |
| minItems | ❌ | 0/4 |
| minLength | ❌ | 0/5 |
| minimum | ❌ | 0/13 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/17 |
| properties | ❌ | 0/15 |
| ref | ❌ | 0/19 |
| refRemote | ❌ | 0/8 |
| required | ❌ | 0/4 |
| type | ❌ | 0/80 |
| uniqueItems | ❌ | 0/62 |

