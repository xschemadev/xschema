# @xschemadev/effect Compliance Report

Generated: 2026-01-09T09:55:50Z

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2020-12 | 971 | 59 | 0 | 94.3% |
| draft2019-09 | 964 | 63 | 0 | 93.9% |
| draft7 | 832 | 40 | 0 | 95.4% |
| draft6 | 752 | 40 | 0 | 94.9% |
| draft4 | 543 | 56 | 0 | 90.7% |
| draft3 | 395 | 33 | 0 | 92.3% |
| v1 | 819 | 60 | 0 | 93.2% |

## Badges

![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-94.3%25-yellow)
![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-93.9%25-yellow)
![draft7](https://img.shields.io/badge/draft7%20compliance-95.4%25-brightgreen)
![draft6](https://img.shields.io/badge/draft6%20compliance-94.9%25-yellow)
![draft4](https://img.shields.io/badge/draft4%20compliance-90.7%25-yellow)
![draft3](https://img.shields.io/badge/draft3%20compliance-92.3%25-yellow)
![v1](https://img.shields.io/badge/v1%20compliance-93.2%25-yellow)

## draft2020-12

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 14/21 |
| allOf | ✅ | 30/30 |
| anchor | ✅ | 0/0 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 21/21 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ✅ | 20/20 |
| dependentSchemas | ⚠️ | 16/20 |
| dynamicRef | ⚠️ | 11/19 |
| enum | ⚠️ | 43/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 133/133 |
| if-then-else | ✅ | 30/30 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 28/29 |
| maxContains | ✅ | 12/12 |
| maxItems | ✅ | 6/6 |
| maxLength | ✅ | 7/7 |
| maxProperties | ✅ | 10/10 |
| maximum | ✅ | 8/8 |
| minContains | ✅ | 28/28 |
| minItems | ✅ | 6/6 |
| minLength | ✅ | 7/7 |
| minProperties | ✅ | 8/8 |
| minimum | ✅ | 11/11 |
| multipleOf | ✅ | 10/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| prefixItems | ✅ | 11/11 |
| properties | ⚠️ | 24/28 |
| propertyNames | ✅ | 20/20 |
| ref | ⚠️ | 33/35 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/16 |
| type | ⚠️ | 71/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 35/46 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 7 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties doesn't consider dependentSchemas
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar even when foo2 is present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentSchemas - 4 failures</summary>

- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
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
<summary>enum - 2 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

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
<summary>properties - 4 failures</summary>

- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 2 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
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
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: an integer is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a float is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a string is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a boolean is not an object
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 11 failures</summary>

- **unevaluatedProperties with adjacent properties**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **cousin unevaluatedProperties, true and false, false with properties**
  - Test: with nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **property is evaluated in an uncle schema to unevaluatedProperties**
  - Test: uncle keyword evaluation is not significant
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, allOf has unevaluated**
  - Test: base case: both properties present
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, allOf has unevaluated**
  - Test: in place applicator siblings, foo is missing
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, anyOf has unevaluated**
  - Test: base case: both properties present
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, anyOf has unevaluated**
  - Test: in place applicator siblings, bar is missing
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties + single cyclic ref**
  - Test: Unevaluated on 1st level is invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties + single cyclic ref**
  - Test: Unevaluated on 2nd level is invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties + single cyclic ref**
  - Test: Unevaluated on 3rd level is invalid
  - Expected: `invalid`, Got: `true`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
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
| additionalProperties | ⚠️ | 14/21 |
| allOf | ✅ | 30/30 |
| anchor | ✅ | 0/0 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 21/21 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ✅ | 20/20 |
| dependentSchemas | ⚠️ | 16/20 |
| enum | ⚠️ | 43/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 114/114 |
| if-then-else | ✅ | 30/30 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 27/28 |
| maxContains | ✅ | 12/12 |
| maxItems | ✅ | 6/6 |
| maxLength | ✅ | 7/7 |
| maxProperties | ✅ | 10/10 |
| maximum | ✅ | 8/8 |
| minContains | ✅ | 28/28 |
| minItems | ✅ | 6/6 |
| minLength | ✅ | 7/7 |
| minProperties | ✅ | 8/8 |
| minimum | ✅ | 11/11 |
| multipleOf | ✅ | 10/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ⚠️ | 24/28 |
| propertyNames | ✅ | 20/20 |
| recursiveRef | ⚠️ | 18/30 |
| ref | ⚠️ | 33/35 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/16 |
| type | ⚠️ | 71/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 33/44 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 7 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties doesn't consider dependentSchemas
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar even when foo2 is present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentSchemas - 4 failures</summary>

- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 2 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

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
<summary>properties - 4 failures</summary>

- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>recursiveRef - 12 failures</summary>

- **$recursiveRef without $recursiveAnchor works like $ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
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
  - Test: leaf node does not match; no recursion
  - Expected: `invalid`, Got: `true`
- **$recursiveRef with no $recursiveAnchor in the initial target schema resource**
  - Test: leaf node does not match: recursion uses the inner schema
  - Expected: `invalid`, Got: `true`
- **$recursiveRef with no $recursiveAnchor in the outer schema resource**
  - Test: leaf node does not match; no recursion
  - Expected: `invalid`, Got: `true`
- **$recursiveRef with no $recursiveAnchor in the outer schema resource**
  - Test: leaf node does not match: recursion only uses inner schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 2 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
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
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: an integer is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a float is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a string is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a boolean is not an object
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 11 failures</summary>

- **unevaluatedProperties with adjacent properties**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **cousin unevaluatedProperties, true and false, false with properties**
  - Test: with nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **property is evaluated in an uncle schema to unevaluatedProperties**
  - Test: uncle keyword evaluation is not significant
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, allOf has unevaluated**
  - Test: base case: both properties present
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, allOf has unevaluated**
  - Test: in place applicator siblings, foo is missing
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, anyOf has unevaluated**
  - Test: base case: both properties present
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, anyOf has unevaluated**
  - Test: in place applicator siblings, bar is missing
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties + single cyclic ref**
  - Test: Unevaluated on 1st level is invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties + single cyclic ref**
  - Test: Unevaluated on 2nd level is invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties + single cyclic ref**
  - Test: Unevaluated on 3rd level is invalid
  - Expected: `invalid`, Got: `true`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
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
| additionalProperties | ⚠️ | 12/16 |
| allOf | ✅ | 30/30 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 21/21 |
| default | ✅ | 7/7 |
| definitions | ❌ | 0/2 |
| dependencies | ⚠️ | 33/36 |
| enum | ⚠️ | 43/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 102/102 |
| if-then-else | ✅ | 30/30 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 27/28 |
| maxItems | ✅ | 6/6 |
| maxLength | ✅ | 7/7 |
| maxProperties | ✅ | 10/10 |
| maximum | ✅ | 8/8 |
| minItems | ✅ | 6/6 |
| minLength | ✅ | 7/7 |
| minProperties | ✅ | 8/8 |
| minimum | ✅ | 11/11 |
| multipleOf | ✅ | 10/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ⚠️ | 24/28 |
| propertyNames | ✅ | 20/20 |
| ref | ⚠️ | 31/36 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/16 |
| type | ⚠️ | 71/80 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 4 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>definitions - 2 failures</summary>

- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import { Schema as S } from "effect";
2 | 
3 | const schema = S.Union(S.Struct({ "$comment": S.optional(S.String), "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "contentEncoding": S.optional(S.String), "contentMediaType": S.optional(S.String), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/effect/xschema-harness-1898311083.ts:3:103
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import { Schema as S } from "effect";
2 | 
3 | const schema = S.Union(S.Struct({ "$comment": S.optional(S.String), "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "contentEncoding": S.optional(S.String), "contentMediaType": S.optional(S.String), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/effect/xschema-harness-1898311083.ts:3:103
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>dependencies - 3 failures</summary>

- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 2 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

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
<summary>properties - 4 failures</summary>

- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 5 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref valid, maxItems ignored
  - Expected: `valid`, Got: `false`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import { Schema as S } from "effect";
2 | 
3 | const schema = S.Union(S.Struct({ "$comment": S.optional(S.String), "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "contentEncoding": S.optional(S.String), "contentMediaType": S.optional(S.String), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/effect/xschema-harness-3576832253.ts:3:103
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import { Schema as S } from "effect";
2 | 
3 | const schema = S.Union(S.Struct({ "$comment": S.optional(S.String), "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "contentEncoding": S.optional(S.String), "contentMediaType": S.optional(S.String), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/effect/xschema-harness-3576832253.ts:3:103
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

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
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: an integer is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a float is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a string is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a boolean is not an object
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

## draft6

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ⚠️ | 12/16 |
| allOf | ✅ | 30/30 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 19/19 |
| default | ✅ | 7/7 |
| definitions | ❌ | 0/2 |
| dependencies | ⚠️ | 33/36 |
| enum | ⚠️ | 43/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 54/54 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 27/28 |
| maxItems | ✅ | 6/6 |
| maxLength | ✅ | 7/7 |
| maxProperties | ✅ | 10/10 |
| maximum | ✅ | 8/8 |
| minItems | ✅ | 6/6 |
| minLength | ✅ | 7/7 |
| minProperties | ✅ | 8/8 |
| minimum | ✅ | 11/11 |
| multipleOf | ✅ | 10/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ⚠️ | 24/28 |
| propertyNames | ✅ | 20/20 |
| ref | ⚠️ | 31/36 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/16 |
| type | ⚠️ | 71/80 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 4 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>definitions - 2 failures</summary>

- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import { Schema as S } from "effect";
2 | 
3 | const schema = S.Union(S.Struct({ "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/effect/xschema-harness-962781202.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import { Schema as S } from "effect";
2 | 
3 | const schema = S.Union(S.Struct({ "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/effect/xschema-harness-962781202.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>dependencies - 3 failures</summary>

- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 2 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

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
<summary>properties - 4 failures</summary>

- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 5 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref valid, maxItems ignored
  - Expected: `valid`, Got: `false`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import { Schema as S } from "effect";
2 | 
3 | const schema = S.Union(S.Struct({ "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/effect/xschema-harness-1687770095.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import { Schema as S } from "effect";
2 | 
3 | const schema = S.Union(S.Struct({ "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/effect/xschema-harness-1687770095.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

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
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: an integer is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a float is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a string is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a boolean is not an object
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

## draft4

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 17/17 |
| additionalProperties | ⚠️ | 12/16 |
| allOf | ✅ | 27/27 |
| anyOf | ✅ | 15/15 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 26/29 |
| enum | ⚠️ | 47/49 |
| format | ✅ | 36/36 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 20/21 |
| maxItems | ✅ | 4/4 |
| maxLength | ✅ | 5/5 |
| maxProperties | ✅ | 8/8 |
| maximum | ⚠️ | 13/14 |
| minItems | ✅ | 4/4 |
| minLength | ✅ | 5/5 |
| minProperties | ✅ | 6/6 |
| minimum | ⚠️ | 16/17 |
| multipleOf | ✅ | 10/10 |
| not | ✅ | 20/20 |
| oneOf | ⚠️ | 19/23 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 18/18 |
| properties | ⚠️ | 20/24 |
| ref | ⚠️ | 22/33 |
| refRemote | ⚠️ | 6/15 |
| required | ⚠️ | 9/15 |
| type | ⚠️ | 70/79 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 4 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 3 failures</summary>

- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 2 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 1 failure</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 1 failure</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`

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
<summary>properties - 4 failures</summary>

- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 11 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref valid, maxItems ignored
  - Expected: `valid`, Got: `false`
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
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "http://localhost:1234/nested.json": failed to read $CACHE/json-schema-test-suite/remotes/nested.json: open $CACHE/json-schema-test-suite/remotes/nested.json: no such file or directory`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "http://localhost:1234/nested.json": failed to read $CACHE/json-schema-test-suite/remotes/nested.json: open $CACHE/json-schema-test-suite/remotes/nested.json: no such file or directory`
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

<details>
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: an integer is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a float is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a string is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a boolean is not an object
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

## draft3

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 14/14 |
| additionalProperties | ⚠️ | 12/16 |
| default | ✅ | 7/7 |
| dependencies | ⚠️ | 17/18 |
| disallow | ✅ | 9/9 |
| divisibleBy | ✅ | 8/8 |
| enum | ⚠️ | 15/16 |
| extends | ✅ | 10/10 |
| format | ✅ | 60/60 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ✅ | 7/7 |
| maxItems | ✅ | 4/4 |
| maxLength | ✅ | 5/5 |
| maximum | ⚠️ | 13/14 |
| minItems | ✅ | 4/4 |
| minLength | ✅ | 5/5 |
| minimum | ⚠️ | 12/13 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 17/17 |
| properties | ⚠️ | 12/15 |
| ref | ⚠️ | 14/21 |
| refRemote | ⚠️ | 6/8 |
| required | ⚠️ | 3/4 |
| type | ⚠️ | 68/80 |
| uniqueItems | ✅ | 62/62 |

### Failures

<details>
<summary>additionalProperties - 4 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in extends are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 1 failure</summary>

- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 1 failure</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maximum - 1 failure</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 1 failure</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 3 failures</summary>

- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>ref - 7 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: remote ref valid, maxItems ignored
  - Expected: `valid`, Got: `false`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import { Schema as S } from "effect";
2 | 
3 | const schema = S.Struct({ "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Union(S.Unknown, S.Boolean)), "additionalProperties": S.optional(S.Union(S.Unknown, S.Boolean)), "default": S.optional(S.Unknown), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.String, S.Array(S.String), S.Unknown) })), "description": S.optional(S.String), "disallow": S.optional(S.Union(S.String, S.Array(S.Union(S.String, S.Unknown)).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/effect/xschema-harness-581643292.ts:3:62
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import { Schema as S } from "effect";
2 | 
3 | const schema = S.Struct({ "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Union(S.Unknown, S.Boolean)), "additionalProperties": S.optional(S.Union(S.Unknown, S.Boolean)), "default": S.optional(S.Unknown), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.String, S.Array(S.String), S.Unknown) })), "description": S.optional(S.String), "disallow": S.optional(S.Union(S.String, S.Array(S.Union(S.String, S.Unknown)).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/effect/xschema-harness-581643292.ts:3:62
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

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
<summary>required - 1 failure</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 12 failures</summary>

- **object type matches objects**
  - Test: an integer is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a float is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a string is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a boolean is not an object
  - Expected: `invalid`, Got: `true`
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
| additionalProperties | ⚠️ | 14/21 |
| allOf | ✅ | 30/30 |
| anchor | ✅ | 0/0 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 25/25 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ❌ | 0/2 |
| dependentRequired | ✅ | 20/20 |
| dependentSchemas | ⚠️ | 16/20 |
| dynamicRef | ⚠️ | 8/13 |
| enum | ⚠️ | 43/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| if-then-else | ✅ | 26/26 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 28/29 |
| maxContains | ✅ | 12/12 |
| maxItems | ✅ | 6/6 |
| maxLength | ✅ | 7/7 |
| maxProperties | ✅ | 10/10 |
| maximum | ✅ | 8/8 |
| minContains | ✅ | 28/28 |
| minItems | ✅ | 6/6 |
| minLength | ✅ | 7/7 |
| minProperties | ✅ | 8/8 |
| minimum | ✅ | 11/11 |
| multipleOf | ✅ | 10/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| prefixItems | ✅ | 11/11 |
| properties | ⚠️ | 24/28 |
| propertyNames | ✅ | 10/10 |
| ref | ⚠️ | 33/38 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/16 |
| type | ⚠️ | 71/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 33/44 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 7 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties doesn't consider dependentSchemas
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar even when foo2 is present
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
<summary>dependentSchemas - 4 failures</summary>

- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
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
<summary>enum - 2 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 1 failure</summary>

- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

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
<summary>properties - 4 failures</summary>

- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 5 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
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
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: an integer is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a float is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a string is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a boolean is not an object
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 11 failures</summary>

- **unevaluatedProperties with adjacent properties**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **cousin unevaluatedProperties, true and false, false with properties**
  - Test: with nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **property is evaluated in an uncle schema to unevaluatedProperties**
  - Test: uncle keyword evaluation is not significant
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, allOf has unevaluated**
  - Test: base case: both properties present
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, allOf has unevaluated**
  - Test: in place applicator siblings, foo is missing
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, anyOf has unevaluated**
  - Test: base case: both properties present
  - Expected: `invalid`, Got: `true`
- **in-place applicator siblings, anyOf has unevaluated**
  - Test: in place applicator siblings, bar is missing
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties + single cyclic ref**
  - Test: Unevaluated on 1st level is invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties + single cyclic ref**
  - Test: Unevaluated on 2nd level is invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties + single cyclic ref**
  - Test: Unevaluated on 3rd level is invalid
  - Expected: `invalid`, Got: `true`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

