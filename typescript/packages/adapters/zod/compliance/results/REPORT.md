# @xschemadev/zod Compliance Report

Generated: 2026-01-05T21:34:05Z

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2020-12 | 944 | 42 | 0 | 95.7% |
| draft2019-09 | 942 | 43 | 0 | 95.6% |
| draft7 | 833 | 39 | 0 | 95.5% |
| draft6 | 753 | 39 | 0 | 95.1% |
| draft4 | 548 | 51 | 0 | 91.5% |
| draft3 | 388 | 38 | 0 | 91.1% |
| v1 | 794 | 43 | 0 | 94.9% |

## Badges

![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-95.7%25-brightgreen)
![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-95.6%25-brightgreen)
![draft7](https://img.shields.io/badge/draft7%20compliance-95.5%25-brightgreen)
![draft6](https://img.shields.io/badge/draft6%20compliance-95.1%25-brightgreen)
![draft4](https://img.shields.io/badge/draft4%20compliance-91.5%25-yellow)
![draft3](https://img.shields.io/badge/draft3%20compliance-91.1%25-yellow)
![v1](https://img.shields.io/badge/v1%20compliance-94.9%25-yellow)

## draft2020-12

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 19/21 |
| allOf | ⚠️ | 27/30 |
| anchor | ✅ | 0/0 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ⚠️ | 19/21 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ⚠️ | 16/20 |
| dependentSchemas | ⚠️ | 18/20 |
| dynamicRef | ⚠️ | 11/19 |
| enum | ⚠️ | 44/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 133/133 |
| if-then-else | ✅ | 30/30 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 28/29 |
| maxContains | ✅ | 12/12 |
| maxItems | ✅ | 6/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ✅ | 10/10 |
| maximum | ✅ | 8/8 |
| minContains | ✅ | 28/28 |
| minItems | ✅ | 6/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ✅ | 8/8 |
| minimum | ✅ | 11/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| prefixItems | ✅ | 11/11 |
| properties | ⚠️ | 27/28 |
| propertyNames | ⚠️ | 19/20 |
| ref | ⚠️ | 33/34 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 1/3 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 2 failures</summary>

- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with propertyNames**
  - Test: Valid against propertyNames, but not additionalProperties
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 3 failures</summary>

- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>contains - 2 failures</summary>

- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentRequired - 4 failures</summary>

- **dependencies with escaped characters**
  - Test: CRLF
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-1022714504.ts:10:23

Bun v1.3.3 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-1022714504.ts:10:23

Bun v1.3.3 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-1022714504.ts:10:23

Bun v1.3.3 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-1022714504.ts:10:23

Bun v1.3.3 (Linux x64)
`

</details>

<details>
<summary>dependentSchemas - 2 failures</summary>

- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dynamicRef - 8 failures</summary>

- **A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **A $dynamicRef to an $anchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **A $ref to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **tests for implementation dynamic anchor and reference link**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $defs first**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $ref first**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref to $dynamicRef finds detached $dynamicAnchor**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **$dynamicRef points to a boolean schema**
  - Test: follow $dynamicRef to a false schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 1 failure</summary>

- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: null is not an object (evaluating 'Object.keys(val)')`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 4 failures</summary>

- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 1 failure</summary>

- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>propertyNames - 1 failure</summary>

- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 1 failure</summary>

- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 6 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 2 failures</summary>

- **in-place applicator siblings, anyOf has unevaluated**
  - Test: base case: both properties present
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, anyOf has unevaluated**
  - Test: in place applicator siblings, bar is missing
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>vocabulary - 1 failure</summary>

- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: no validation: invalid number, but it still validates
  - Expected: `valid`, Got: `false`

</details>

## draft2019-09

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ⚠️ | 19/21 |
| allOf | ⚠️ | 27/30 |
| anchor | ✅ | 0/0 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ⚠️ | 19/21 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ⚠️ | 16/20 |
| dependentSchemas | ⚠️ | 18/20 |
| enum | ⚠️ | 44/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 114/114 |
| if-then-else | ✅ | 30/30 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 27/28 |
| maxContains | ✅ | 12/12 |
| maxItems | ✅ | 6/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ✅ | 10/10 |
| maximum | ✅ | 8/8 |
| minContains | ✅ | 28/28 |
| minItems | ✅ | 6/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ✅ | 8/8 |
| minimum | ✅ | 11/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ⚠️ | 27/28 |
| propertyNames | ⚠️ | 19/20 |
| recursiveRef | ⚠️ | 21/30 |
| ref | ⚠️ | 33/34 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 1/3 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 2 failures</summary>

- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with propertyNames**
  - Test: Valid against propertyNames, but not additionalProperties
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 3 failures</summary>

- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>contains - 2 failures</summary>

- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentRequired - 4 failures</summary>

- **dependencies with escaped characters**
  - Test: CRLF
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-2836760364.ts:10:23

Bun v1.3.3 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-2836760364.ts:10:23

Bun v1.3.3 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-2836760364.ts:10:23

Bun v1.3.3 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-2836760364.ts:10:23

Bun v1.3.3 (Linux x64)
`

</details>

<details>
<summary>dependentSchemas - 2 failures</summary>

- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 1 failure</summary>

- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: null is not an object (evaluating 'Object.keys(val)')`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 4 failures</summary>

- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 1 failure</summary>

- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>propertyNames - 1 failure</summary>

- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>recursiveRef - 9 failures</summary>

- **$recursiveRef without $recursiveAnchor works like $ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **$recursiveRef without using nesting**
  - Test: integer does not match as a property value
  - Expected: `invalid`, Got: `true`
- **$recursiveRef without using nesting**
  - Test: two levels, no match
  - Expected: `invalid`, Got: `true`
- **$recursiveRef with $recursiveAnchor: false works like $ref**
  - Test: integer does not match as a property value
  - Expected: `invalid`, Got: `true`
- **$recursiveRef with $recursiveAnchor: false works like $ref**
  - Test: two levels, integer does not match as a property value
  - Expected: `invalid`, Got: `true`
- **$recursiveRef with no $recursiveAnchor works like $ref**
  - Test: integer does not match as a property value
  - Expected: `invalid`, Got: `true`
- **$recursiveRef with no $recursiveAnchor works like $ref**
  - Test: two levels, integer does not match as a property value
  - Expected: `invalid`, Got: `true`
- **$recursiveRef with no $recursiveAnchor in the initial target schema resource**
  - Test: leaf node does not match: recursion uses the inner schema
  - Expected: `invalid`, Got: `true`
- **$recursiveRef with no $recursiveAnchor in the outer schema resource**
  - Test: leaf node does not match: recursion only uses inner schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 1 failure</summary>

- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 6 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 2 failures</summary>

- **in-place applicator siblings, anyOf has unevaluated**
  - Test: base case: both properties present
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, anyOf has unevaluated**
  - Test: in place applicator siblings, bar is missing
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>vocabulary - 1 failure</summary>

- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: no validation: invalid number, but it still validates
  - Expected: `valid`, Got: `false`

</details>

## draft7

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ⚠️ | 15/16 |
| allOf | ⚠️ | 27/30 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ⚠️ | 19/21 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 21/36 |
| enum | ⚠️ | 44/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 102/102 |
| if-then-else | ✅ | 30/30 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 27/28 |
| maxItems | ✅ | 6/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ✅ | 10/10 |
| maximum | ✅ | 8/8 |
| minItems | ✅ | 6/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ✅ | 8/8 |
| minimum | ✅ | 11/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ⚠️ | 27/28 |
| propertyNames | ⚠️ | 19/20 |
| ref | ✅ | 36/36 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/16 |
| type | ✅ | 80/80 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 1 failure</summary>

- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 3 failures</summary>

- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>contains - 2 failures</summary>

- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 15 failures</summary>

- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type other
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type both
  - Expected: `invalid`, Got: `true`
- **dependencies with boolean subschemas**
  - Test: object with property having schema false is invalid
  - Expected: `invalid`, Got: `true`
- **dependencies with boolean subschemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 1
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 2
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 4
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 1 failure</summary>

- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: null is not an object (evaluating 'Object.keys(val)')`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 4 failures</summary>

- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 1 failure</summary>

- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>propertyNames - 1 failure</summary>

- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 6 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`

</details>

## draft6

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ⚠️ | 15/16 |
| allOf | ⚠️ | 27/30 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ⚠️ | 17/19 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 21/36 |
| enum | ⚠️ | 44/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 54/54 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 27/28 |
| maxItems | ✅ | 6/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ✅ | 10/10 |
| maximum | ✅ | 8/8 |
| minItems | ✅ | 6/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ✅ | 8/8 |
| minimum | ✅ | 11/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ⚠️ | 27/28 |
| propertyNames | ⚠️ | 19/20 |
| ref | ✅ | 36/36 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/16 |
| type | ✅ | 80/80 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 1 failure</summary>

- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 3 failures</summary>

- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>contains - 2 failures</summary>

- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 15 failures</summary>

- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type other
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type both
  - Expected: `invalid`, Got: `true`
- **dependencies with boolean subschemas**
  - Test: object with property having schema false is invalid
  - Expected: `invalid`, Got: `true`
- **dependencies with boolean subschemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 1
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 2
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 4
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 1 failure</summary>

- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: null is not an object (evaluating 'Object.keys(val)')`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 4 failures</summary>

- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 1 failure</summary>

- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>propertyNames - 1 failure</summary>

- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 6 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`

</details>

## draft4

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 17/17 |
| additionalProperties | ⚠️ | 15/16 |
| allOf | ⚠️ | 24/27 |
| anyOf | ✅ | 15/15 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 16/29 |
| enum | ⚠️ | 48/49 |
| format | ✅ | 36/36 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 20/21 |
| maxItems | ✅ | 4/4 |
| maxLength | ⚠️ | 4/5 |
| maxProperties | ✅ | 8/8 |
| maximum | ✅ | 14/14 |
| minItems | ✅ | 4/4 |
| minLength | ⚠️ | 4/5 |
| minProperties | ✅ | 6/6 |
| minimum | ✅ | 17/17 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 20/20 |
| oneOf | ⚠️ | 19/23 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 18/18 |
| properties | ⚠️ | 23/24 |
| ref | ⚠️ | 25/33 |
| refRemote | ⚠️ | 6/15 |
| required | ⚠️ | 9/15 |
| type | ✅ | 79/79 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 1 failure</summary>

- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 3 failures</summary>

- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 13 failures</summary>

- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type other
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type both
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 1
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 2
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: invalid object 4
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 1 failure</summary>

- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: null is not an object (evaluating 'Object.keys(val)')`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 4 failures</summary>

- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 1 failure</summary>

- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>ref - 8 failures</summary>

- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **Recursive references between schemas**
  - Test: valid tree
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "tree": failed to read tree: open tree: no such file or directory`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "tree": failed to read tree: open tree: no such file or directory`
- **Location-independent identifier with base URI change in subschema**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "http://localhost:1234/nested.json": failed to read /home/trapani/.cache/xschema/json-schema-test-suite/remotes/nested.json: open /home/trapani/.cache/xschema/json-schema-test-suite/remotes/nested.json: no such file or directory`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "http://localhost:1234/nested.json": failed to read /home/trapani/.cache/xschema/json-schema-test-suite/remotes/nested.json: open /home/trapani/.cache/xschema/json-schema-test-suite/remotes/nested.json: no such file or directory`
- **id must be resolved against nearest parent, not just immediate parent**
  - Test: number is valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "http://example.com/b/d.json": failed to fetch http://example.com/b/d.json: status 404`
- **id must be resolved against nearest parent, not just immediate parent**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "http://example.com/b/d.json": failed to fetch http://example.com/b/d.json: status 404`

</details>

<details>
<summary>refRemote - 9 failures</summary>

- **base URI change**
  - Test: base URI change ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`
- **base URI change**
  - Test: base URI change ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`
- **base URI change - change folder**
  - Test: number is valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`
- **base URI change - change folder**
  - Test: string is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`
- **base URI change - change folder in subschema**
  - Test: number is valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`
- **base URI change - change folder in subschema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`
- **root ref in remote ref**
  - Test: string is valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "draft4/name.json": failed to read draft4/name.json: open draft4/name.json: no such file or directory`
- **root ref in remote ref**
  - Test: null is valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "draft4/name.json": failed to read draft4/name.json: open draft4/name.json: no such file or directory`
- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "draft4/name.json": failed to read draft4/name.json: open draft4/name.json: no such file or directory`

</details>

<details>
<summary>required - 6 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`

</details>

## draft3

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 14/14 |
| additionalProperties | ✅ | 16/16 |
| default | ✅ | 7/7 |
| dependencies | ⚠️ | 11/18 |
| disallow | ⚠️ | 4/9 |
| divisibleBy | ⚠️ | 5/8 |
| enum | ⚠️ | 13/16 |
| extends | ⚠️ | 4/10 |
| format | ✅ | 60/60 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ✅ | 7/7 |
| maxItems | ✅ | 4/4 |
| maxLength | ⚠️ | 4/5 |
| maximum | ✅ | 14/14 |
| minItems | ✅ | 4/4 |
| minLength | ⚠️ | 4/5 |
| minimum | ✅ | 13/13 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 17/17 |
| properties | ✅ | 15/15 |
| ref | ⚠️ | 18/21 |
| refRemote | ⚠️ | 6/8 |
| required | ✅ | 2/2 |
| type | ⚠️ | 73/80 |
| uniqueItems | ✅ | 62/62 |

### Failures

<details>
<summary>dependencies - 7 failures</summary>

- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type other
  - Expected: `invalid`, Got: `true`
- **multiple dependencies subschema**
  - Test: wrong type both
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>disallow - 5 failures</summary>

- **disallow**
  - Test: disallowed
  - Expected: `invalid`, Got: `true`
- **multiple disallow**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **multiple disallow**
  - Test: other mismatch
  - Expected: `invalid`, Got: `true`
- **multiple disallow subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **multiple disallow subschema**
  - Test: other mismatch
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>divisibleBy - 3 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by number**
  - Test: 35 is not divisible by 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not divisible by 0.0001
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 3 failures</summary>

- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: null is not an object (evaluating 'Object.keys(val)')`
- **enums in properties**
  - Test: missing required property is invalid
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing all properties is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>extends - 6 failures</summary>

- **extends**
  - Test: mismatch extends
  - Expected: `invalid`, Got: `true`
- **extends**
  - Test: mismatch extended
  - Expected: `invalid`, Got: `true`
- **multiple extends**
  - Test: mismatch first extends
  - Expected: `invalid`, Got: `true`
- **multiple extends**
  - Test: mismatch second extends
  - Expected: `invalid`, Got: `true`
- **multiple extends**
  - Test: mismatch both
  - Expected: `invalid`, Got: `true`
- **extends simple types**
  - Test: mismatch extends
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 3 failures</summary>

- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>refRemote - 2 failures</summary>

- **change resolution scope**
  - Test: changed scope ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`
- **change resolution scope**
  - Test: changed scope ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`

</details>

<details>
<summary>type - 7 failures</summary>

- **types can include schemas**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `true`
- **types can include schemas**
  - Test: a string is invalid
  - Expected: `invalid`, Got: `true`
- **types can include schemas**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`
- **types can include schemas**
  - Test: a boolean is invalid
  - Expected: `invalid`, Got: `true`
- **types can include schemas**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **applies a nested schema**
  - Test: an object is invalid otherwise
  - Expected: `invalid`, Got: `true`
- **types from separate schemas are merged**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `true`

</details>

## v1

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 19/21 |
| allOf | ⚠️ | 27/30 |
| anchor | ✅ | 0/0 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ⚠️ | 23/25 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ❌ | 0/2 |
| dependentRequired | ⚠️ | 16/20 |
| dependentSchemas | ⚠️ | 18/20 |
| dynamicRef | ⚠️ | 8/13 |
| enum | ⚠️ | 44/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| if-then-else | ✅ | 26/26 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 28/29 |
| maxContains | ✅ | 12/12 |
| maxItems | ✅ | 6/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ✅ | 10/10 |
| maximum | ✅ | 8/8 |
| minContains | ✅ | 28/28 |
| minItems | ✅ | 6/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ✅ | 8/8 |
| minimum | ✅ | 11/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| prefixItems | ✅ | 11/11 |
| properties | ⚠️ | 27/28 |
| propertyNames | ⚠️ | 9/10 |
| ref | ⚠️ | 33/37 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 1/3 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 2 failures</summary>

- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with propertyNames**
  - Test: Valid against propertyNames, but not additionalProperties
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 3 failures</summary>

- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>contains - 2 failures</summary>

- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>defs - 2 failures</summary>

- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`
- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`

</details>

<details>
<summary>dependentRequired - 4 failures</summary>

- **dependencies with escaped characters**
  - Test: CRLF
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-2568105979.ts:10:23

Bun v1.3.3 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-2568105979.ts:10:23

Bun v1.3.3 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-2568105979.ts:10:23

Bun v1.3.3 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 10 |         }, { message: "Property foo
                           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema/typescript/packages/adapters/zod/xschema-harness-2568105979.ts:10:23

Bun v1.3.3 (Linux x64)
`

</details>

<details>
<summary>dependentSchemas - 2 failures</summary>

- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dynamicRef - 5 failures</summary>

- **A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **tests for implementation dynamic anchor and reference link**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $defs first**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $ref first**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref to $dynamicRef finds detached $dynamicAnchor**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 1 failure</summary>

- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: null is not an object (evaluating 'Object.keys(val)')`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 4 failures</summary>

- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 1 failure</summary>

- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>propertyNames - 1 failure</summary>

- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 4 failures</summary>

- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`
- **URN base URI with f-component**
  - Test: is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`

</details>

<details>
<summary>required - 6 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 2 failures</summary>

- **in-place applicator siblings, anyOf has unevaluated**
  - Test: base case: both properties present
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, anyOf has unevaluated**
  - Test: in place applicator siblings, bar is missing
  - Expected: `invalid`, Got: `true`

</details>

