# @xschemadev/zod Compliance Report

Generated: 2026-01-05T13:32:46Z

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2020-12 | 904 | 371 | 0 | 70.9% |
| draft2019-09 | 883 | 355 | 0 | 71.3% |
| draft7 | 691 | 229 | 0 | 75.1% |
| draft6 | 616 | 216 | 0 | 74.0% |
| draft4 | 450 | 163 | 0 | 73.4% |
| draft3 | 330 | 104 | 0 | 76.0% |
| v1 | 751 | 358 | 0 | 67.7% |

## Badges

![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-70.9%25-red)
![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-71.3%25-red)
![draft7](https://img.shields.io/badge/draft7%20compliance-75.1%25-red)
![draft6](https://img.shields.io/badge/draft6%20compliance-74.0%25-red)
![draft4](https://img.shields.io/badge/draft4%20compliance-73.4%25-red)
![draft3](https://img.shields.io/badge/draft3%20compliance-76.0%25-red)
![v1](https://img.shields.io/badge/v1%20compliance-67.7%25-red)

## draft2020-12

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 12/21 |
| allOf | ⚠️ | 15/30 |
| anchor | ⚠️ | 4/8 |
| anyOf | ⚠️ | 14/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 47/54 |
| contains | ⚠️ | 11/21 |
| content | ✅ | 18/18 |
| default | ⚠️ | 6/7 |
| defs | ⚠️ | 1/2 |
| dependentRequired | ⚠️ | 14/20 |
| dependentSchemas | ⚠️ | 10/20 |
| dynamicRef | ⚠️ | 22/44 |
| enum | ⚠️ | 37/45 |
| exclusiveMaximum | ⚠️ | 2/4 |
| exclusiveMinimum | ⚠️ | 2/4 |
| format | ✅ | 133/133 |
| if-then-else | ⚠️ | 22/30 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 19/29 |
| maxContains | ⚠️ | 6/12 |
| maxItems | ⚠️ | 4/6 |
| maxLength | ⚠️ | 5/7 |
| maxProperties | ⚠️ | 7/10 |
| maximum | ⚠️ | 6/8 |
| minContains | ⚠️ | 14/28 |
| minItems | ⚠️ | 4/6 |
| minLength | ⚠️ | 4/7 |
| minProperties | ⚠️ | 6/8 |
| minimum | ⚠️ | 8/11 |
| multipleOf | ⚠️ | 6/10 |
| not | ⚠️ | 38/40 |
| oneOf | ⚠️ | 18/27 |
| pattern | ⚠️ | 8/9 |
| patternProperties | ⚠️ | 13/23 |
| prefixItems | ⚠️ | 9/11 |
| properties | ⚠️ | 16/28 |
| propertyNames | ⚠️ | 15/20 |
| ref | ⚠️ | 38/79 |
| refRemote | ⚠️ | 16/31 |
| required | ⚠️ | 10/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ⚠️ | 42/71 |
| unevaluatedProperties | ⚠️ | 69/125 |
| uniqueItems | ⚠️ | 50/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 9 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with propertyNames**
  - Test: Valid against propertyNames, but not additionalProperties
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
<summary>allOf - 15 failures</summary>

- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `true`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `true`
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
<summary>anchor - 4 failures</summary>

- **Location-independent identifier**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier with absolute URI**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **same $anchor with different base uri**
  - Test: $ref does not resolve to /$defs/A/allOf/0
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>anyOf - 4 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>const - 7 failures</summary>

- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `false`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `false`
- **const with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **const with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": false} is valid
  - Expected: `valid`, Got: `false`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": true} is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>contains - 10 failures</summary>

- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with const keyword**
  - Test: array without item 5 is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema true**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: matches items, does not match contains
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: does not match items, matches contains
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: matches neither items nor contains
  - Expected: `invalid`, Got: `true`
- **contains with false if subschema**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>default - 1 failure</summary>

- **the default keyword does not do anything if the property is missing**
  - Test: missing properties are not filled in with the default
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>defs - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentRequired - 6 failures</summary>

- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependents required**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependents required**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependents required**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentSchemas - 10 failures</summary>

- **single dependency**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type other
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type both
  - Expected: `invalid`, Got: `true`
- **boolean subschemas**
  - Test: object with property having schema false is invalid
  - Expected: `invalid`, Got: `true`
- **boolean subschemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted tab invalid under dependent schema
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
<summary>dynamicRef - 22 failures</summary>

- **A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **A $dynamicRef to an $anchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **A $ref to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **A $dynamicRef resolves to the first $dynamicAnchor still in scope that is encountered when the schema is evaluated**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **A $dynamicRef without anchor in fragment behaves identical to $ref**
  - Test: An array of strings is invalid
  - Expected: `invalid`, Got: `true`
- **A $dynamicRef with intermediate scopes that don't include a matching $dynamicAnchor does not affect dynamic scope resolution**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor resolves to the first $dynamicAnchor in the dynamic scope**
  - Test: The recursive part is not valid against the root
  - Expected: `invalid`, Got: `true`
- **multiple dynamic paths to the $dynamicRef keyword**
  - Test: number list with string values
  - Expected: `invalid`, Got: `true`
- **multiple dynamic paths to the $dynamicRef keyword**
  - Test: string list with number values
  - Expected: `invalid`, Got: `true`
- **after leaving a dynamic scope, it is not used by a $dynamicRef**
  - Test: string matches /$defs/thingy, but the $dynamicRef does not stop here
  - Expected: `invalid`, Got: `true`
- **after leaving a dynamic scope, it is not used by a $dynamicRef**
  - Test: first_scope is not in dynamic scope for the $dynamicRef
  - Expected: `invalid`, Got: `true`
- **strict-tree schema, guards against misspelled properties**
  - Test: instance with misspelled field
  - Expected: `invalid`, Got: `true`
- **tests for implementation dynamic anchor and reference link**
  - Test: incorrect parent schema
  - Expected: `invalid`, Got: `true`
- **tests for implementation dynamic anchor and reference link**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $defs first**
  - Test: incorrect parent schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $defs first**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $ref first**
  - Test: incorrect parent schema
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
- **$dynamicRef skips over intermediate resources - direct reference**
  - Test: string property fails
  - Expected: `invalid`, Got: `true`
- **$dynamicRef avoids the root of each schema, but scopes are still registered**
  - Test: data is not sufficient for schema at second#/$defs/length
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 8 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMaximum - 2 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>exclusiveMinimum - 2 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>if-then-else - 8 failures</summary>

- **if and then without else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `true`
- **if and else without then**
  - Test: invalid through else
  - Expected: `invalid`, Got: `true`
- **validate against correct branch, then vs else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `true`
- **validate against correct branch, then vs else**
  - Test: invalid through else
  - Expected: `invalid`, Got: `true`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: other redirects to else and passes
  - Expected: `valid`, Got: `false`
- **then: false fails when condition matches**
  - Test: matches if → then=false → invalid
  - Expected: `invalid`, Got: `true`
- **else: false fails when condition does not match**
  - Test: does not match if → else executes → invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>infinite-loop-detection - 1 failure</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>items - 10 failures</summary>

- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `true`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **prefixItems with no additional items allowed**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **items does not look in applicators, valid case**
  - Test: prefixItems in allOf does not constrain items, invalid case
  - Expected: `invalid`, Got: `true`
- **prefixItems validation adjusts the starting index for items**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `true`
- **items with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxContains - 6 failures</summary>

- **maxContains with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **maxContains with contains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains with contains**
  - Test: some elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains with contains, value with a decimal**
  - Test: too many elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **minContains < maxContains**
  - Test: actual < minContains < maxContains
  - Expected: `invalid`, Got: `true`
- **minContains < maxContains**
  - Test: minContains < maxContains < actual
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxItems - 2 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 2 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 3 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 2 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minContains - 14 failures</summary>

- **minContains=1 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **minContains=1 with contains**
  - Test: no elements match
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains**
  - Test: some elements match, invalid minContains
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains with a decimal value**
  - Test: one element matches, invalid minContains
  - Expected: `invalid`, Got: `true`
- **maxContains = minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **maxContains = minContains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `true`
- **maxContains = minContains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: invalid minContains
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: invalid maxContains and minContains
  - Expected: `invalid`, Got: `true`
- **minContains = 0 with maxContains**
  - Test: too many
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 2 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 3 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 2 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 3 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 4 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>not - 2 failures</summary>

- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`
- **collect annotations inside a 'not', even if collection is disabled**
  - Test: unevaluated property
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 9 failures</summary>

- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `false`
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
<summary>pattern - 1 failure</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>patternProperties - 10 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema false is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with a property matching both true and false is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>prefixItems - 2 failures</summary>

- **a schema given for prefixItems**
  - Test: wrong types
  - Expected: `invalid`, Got: `true`
- **prefixItems with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 12 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`
- **properties with boolean schema**
  - Test: only 'false' property present is invalid
  - Expected: `invalid`, Got: `true`
- **properties with boolean schema**
  - Test: both properties present is invalid
  - Expected: `invalid`, Got: `true`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>propertyNames - 5 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames validation with pattern**
  - Test: non-matching property name is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with const**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with enum**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 41 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: slash invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: tilde invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: percent invalid
  - Expected: `invalid`, Got: `true`
- **nested refs**
  - Test: nested ref invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref to boolean schema false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `true`
- **refs with quote**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `true`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `true`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `true`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `true`
- **$id must be resolved against nearest parent, not just immediate parent**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **order of evaluation: $id and $ref**
  - Test: data is invalid against first definition
  - Expected: `invalid`, Got: `true`
- **order of evaluation: $id and $anchor and $ref**
  - Test: data is invalid against first definition
  - Expected: `invalid`, Got: `true`
- **order of evaluation: $id and $ref on nested schema**
  - Test: data is invalid against nested sibling
  - Expected: `invalid`, Got: `true`
- **simple URN base URI with $ref via the URN**
  - Test: invalid under the URN IDed schema
  - Expected: `invalid`, Got: `true`
- **simple URN base URI with JSON pointer**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with NSS**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with r-component**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with q-component**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with URN and JSON pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with URN and anchor ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN ref with nested pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **ref to if**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref to then**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref to else**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref with absolute-path-reference**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `true`
- **$id with file URI still resolves pointers - *nix**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **$id with file URI still resolves pointers - windows**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **empty tokens in $ref json-pointer**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>refRemote - 15 failures</summary>

- **remote ref**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **fragment within remote ref**
  - Test: remote fragment invalid
  - Expected: `invalid`, Got: `true`
- **anchor within remote ref**
  - Test: remote anchor invalid
  - Expected: `invalid`, Got: `true`
- **ref within remote ref**
  - Test: ref within ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change**
  - Test: base URI change ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder in subschema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`
- **remote ref with ref to defs**
  - Test: invalid
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier in remote ref**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **remote HTTP ref with different $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **remote HTTP ref with different URN $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **remote HTTP ref with nested absolute ref**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **$ref to $ref finds detached $anchor**
  - Test: non-number is invalid
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
<summary>unevaluatedItems - 29 failures</summary>

- **unevaluatedItems false**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems as schema**
  - Test: with invalid unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with tuple**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with items**
  - Test: invalid under items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with nested tuple**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with nested items**
  - Test: with invalid additional item
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with anyOf**
  - Test: when one schema matches and has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with anyOf**
  - Test: when two schemas match and has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with oneOf**
  - Test: with no unevaluated items
  - Expected: `valid`, Got: `false`
- **unevaluatedItems with not**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with if/then/else**
  - Test: when if matches and it has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with if/then/else**
  - Test: when if doesn't match and it has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with boolean schemas**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with $ref**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems before $ref**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with $dynamicRef**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems can't see inside cousins**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **item is evaluated in an uncle schema to unevaluatedItems**
  - Test: uncle keyword evaluation is not significant
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems depends on adjacent contains**
  - Test: contains fails, second item is not evaluated
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems depends on adjacent contains**
  - Test: contains passes, second item is not evaluated
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems depends on multiple nested contains**
  - Test: 7 not evaluated, fails unevaluatedItems
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems and contains interact to control item dependency relationship**
  - Test: only b's are invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems and contains interact to control item dependency relationship**
  - Test: only c's are invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems and contains interact to control item dependency relationship**
  - Test: only b's and c's are invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems and contains interact to control item dependency relationship**
  - Test: only a's and c's are invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with minContains = 0**
  - Test: no items evaluated by contains
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with minContains = 0**
  - Test: some but not all items evaluated by contains
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems can see annotations from if without then and else**
  - Test: invalid in case if is evaluated
  - Expected: `invalid`, Got: `true`
- **Evaluated items collection needs to consider instance location**
  - Test: with an unevaluated item that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 56 failures</summary>

- **unevaluatedProperties schema**
  - Test: with invalid unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties false**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with adjacent properties**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with adjacent patternProperties**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with adjacent non-bool additionalProperties**
  - Test: with invalid additional properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with nested properties**
  - Test: with additional properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with nested patternProperties**
  - Test: with additional properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with anyOf**
  - Test: when one matches and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with anyOf**
  - Test: when two match and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with oneOf**
  - Test: with no unevaluated properties
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties with not**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else**
  - Test: when if is true and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else**
  - Test: when if is false and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, then not defined**
  - Test: when if is true and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, then not defined**
  - Test: when if is false and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, else not defined**
  - Test: when if is true and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, else not defined**
  - Test: when if is false and has no unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, else not defined**
  - Test: when if is false and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with dependentSchemas**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with boolean schemas**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with $ref**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties before $ref**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with $dynamicRef**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can't see inside cousins**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can't see inside cousins (reverse order)**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **nested unevaluatedProperties, outer true, inner false, properties outside**
  - Test: with no nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **nested unevaluatedProperties, outer true, inner false, properties outside**
  - Test: with nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **nested unevaluatedProperties, outer true, inner false, properties inside**
  - Test: with nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **cousin unevaluatedProperties, true and false, true with properties**
  - Test: with no nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **cousin unevaluatedProperties, true and false, true with properties**
  - Test: with nested unevaluated properties
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
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and x are valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and y are valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and b and x are valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and b and y are valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: a is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: b is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: c is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: d is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: xx is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: xx + foox is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: all is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: all + foo is valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties not affected by propertyNames**
  - Test: string property is invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can see annotations from if without then and else**
  - Test: invalid in case if is evaluated
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with unevaluatedProperties**
  - Test: unevaluatedProperties doesn't consider dependentSchemas
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with unevaluatedProperties**
  - Test: unevaluatedProperties doesn't see bar when foo2 is absent
  - Expected: `invalid`, Got: `true`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 19 failures</summary>

- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>vocabulary - 1 failure</summary>

- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: applicator vocabulary still works
  - Expected: `invalid`, Got: `true`

</details>

## draft2019-09

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 13/19 |
| additionalProperties | ⚠️ | 12/21 |
| allOf | ⚠️ | 15/30 |
| anchor | ⚠️ | 4/8 |
| anyOf | ⚠️ | 14/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 47/54 |
| contains | ⚠️ | 11/21 |
| content | ✅ | 18/18 |
| default | ⚠️ | 6/7 |
| defs | ⚠️ | 1/2 |
| dependentRequired | ⚠️ | 14/20 |
| dependentSchemas | ⚠️ | 10/20 |
| enum | ⚠️ | 37/45 |
| exclusiveMaximum | ⚠️ | 2/4 |
| exclusiveMinimum | ⚠️ | 2/4 |
| format | ✅ | 114/114 |
| if-then-else | ⚠️ | 22/30 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 20/28 |
| maxContains | ⚠️ | 6/12 |
| maxItems | ⚠️ | 4/6 |
| maxLength | ⚠️ | 5/7 |
| maxProperties | ⚠️ | 7/10 |
| maximum | ⚠️ | 6/8 |
| minContains | ⚠️ | 14/28 |
| minItems | ⚠️ | 4/6 |
| minLength | ⚠️ | 4/7 |
| minProperties | ⚠️ | 6/8 |
| minimum | ⚠️ | 8/11 |
| multipleOf | ⚠️ | 6/10 |
| not | ⚠️ | 38/40 |
| oneOf | ⚠️ | 18/27 |
| pattern | ⚠️ | 8/9 |
| patternProperties | ⚠️ | 13/23 |
| properties | ⚠️ | 16/28 |
| propertyNames | ⚠️ | 15/20 |
| recursiveRef | ⚠️ | 22/34 |
| ref | ⚠️ | 39/81 |
| refRemote | ⚠️ | 16/31 |
| required | ⚠️ | 10/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ⚠️ | 35/56 |
| unevaluatedProperties | ⚠️ | 68/123 |
| uniqueItems | ⚠️ | 50/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalItems - 6 failures</summary>

- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `true`
- **when items is schema, additionalItems does nothing**
  - Test: invalid with a array of mixed types
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `true`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 9 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with propertyNames**
  - Test: Valid against propertyNames, but not additionalProperties
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
<summary>allOf - 15 failures</summary>

- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `true`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `true`
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
<summary>anchor - 4 failures</summary>

- **Location-independent identifier**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier with absolute URI**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **same $anchor with different base uri**
  - Test: $ref does not resolve to /$defs/A/allOf/0
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>anyOf - 4 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>const - 7 failures</summary>

- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `false`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `false`
- **const with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **const with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": false} is valid
  - Expected: `valid`, Got: `false`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": true} is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>contains - 10 failures</summary>

- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with const keyword**
  - Test: array without item 5 is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema true**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: matches items, does not match contains
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: does not match items, matches contains
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: matches neither items nor contains
  - Expected: `invalid`, Got: `true`
- **contains with false if subschema**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>default - 1 failure</summary>

- **the default keyword does not do anything if the property is missing**
  - Test: missing properties are not filled in with the default
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>defs - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentRequired - 6 failures</summary>

- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependents required**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependents required**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependents required**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentSchemas - 10 failures</summary>

- **single dependency**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type other
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type both
  - Expected: `invalid`, Got: `true`
- **boolean subschemas**
  - Test: object with property having schema false is invalid
  - Expected: `invalid`, Got: `true`
- **boolean subschemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted tab invalid under dependent schema
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
<summary>enum - 8 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMaximum - 2 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>exclusiveMinimum - 2 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>if-then-else - 8 failures</summary>

- **if and then without else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `true`
- **if and else without then**
  - Test: invalid through else
  - Expected: `invalid`, Got: `true`
- **validate against correct branch, then vs else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `true`
- **validate against correct branch, then vs else**
  - Test: invalid through else
  - Expected: `invalid`, Got: `true`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: other redirects to else and passes
  - Expected: `valid`, Got: `false`
- **then: false fails when condition matches**
  - Test: matches if → then=false → invalid
  - Expected: `invalid`, Got: `true`
- **else: false fails when condition does not match**
  - Test: does not match if → else executes → invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>infinite-loop-detection - 1 failure</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>items - 8 failures</summary>

- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `true`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `true`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxContains - 6 failures</summary>

- **maxContains with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **maxContains with contains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains with contains**
  - Test: some elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains with contains, value with a decimal**
  - Test: too many elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **minContains < maxContains**
  - Test: actual < minContains < maxContains
  - Expected: `invalid`, Got: `true`
- **minContains < maxContains**
  - Test: minContains < maxContains < actual
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxItems - 2 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 2 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 3 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 2 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minContains - 14 failures</summary>

- **minContains=1 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **minContains=1 with contains**
  - Test: no elements match
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains**
  - Test: some elements match, invalid minContains
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains with a decimal value**
  - Test: one element matches, invalid minContains
  - Expected: `invalid`, Got: `true`
- **maxContains = minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **maxContains = minContains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `true`
- **maxContains = minContains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: invalid minContains
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: invalid maxContains and minContains
  - Expected: `invalid`, Got: `true`
- **minContains = 0 with maxContains**
  - Test: too many
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 2 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 3 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 2 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 3 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 4 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>not - 2 failures</summary>

- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`
- **collect annotations inside a 'not', even if collection is disabled**
  - Test: unevaluated property
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 9 failures</summary>

- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `false`
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
<summary>pattern - 1 failure</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>patternProperties - 10 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema false is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with a property matching both true and false is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 12 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`
- **properties with boolean schema**
  - Test: only 'false' property present is invalid
  - Expected: `invalid`, Got: `true`
- **properties with boolean schema**
  - Test: both properties present is invalid
  - Expected: `invalid`, Got: `true`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>propertyNames - 5 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames validation with pattern**
  - Test: non-matching property name is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with const**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with enum**
  - Test: object with any other property is invalid
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
  - Test: leaf node does not match: recursion uses the inner schema
  - Expected: `invalid`, Got: `true`
- **$recursiveRef with no $recursiveAnchor in the outer schema resource**
  - Test: leaf node does not match: recursion only uses inner schema
  - Expected: `invalid`, Got: `true`
- **multiple dynamic paths to the $recursiveRef keyword**
  - Test: recurse to integerNode - floats are not allowed
  - Expected: `invalid`, Got: `true`
- **dynamic $recursiveRef destination (not predictable at schema compile time)**
  - Test: integer node
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 42 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: slash invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: tilde invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: percent invalid
  - Expected: `invalid`, Got: `true`
- **nested refs**
  - Test: nested ref invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref to boolean schema false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `true`
- **refs with quote**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `true`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `true`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `true`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `true`
- **$id must be resolved against nearest parent, not just immediate parent**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **order of evaluation: $id and $ref**
  - Test: data is invalid against first definition
  - Expected: `invalid`, Got: `true`
- **order of evaluation: $id and $anchor and $ref**
  - Test: data is invalid against first definition
  - Expected: `invalid`, Got: `true`
- **order of evaluation: $id and $ref on nested schema**
  - Test: data is invalid against nested sibling
  - Expected: `invalid`, Got: `true`
- **simple URN base URI with $ref via the URN**
  - Test: invalid under the URN IDed schema
  - Expected: `invalid`, Got: `true`
- **simple URN base URI with JSON pointer**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with NSS**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with r-component**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with q-component**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with URN and JSON pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with URN and anchor ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN ref with nested pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **ref to if**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref to then**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref to else**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref with absolute-path-reference**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `true`
- **$id with file URI still resolves pointers - *nix**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **$id with file URI still resolves pointers - windows**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **empty tokens in $ref json-pointer**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **$ref with $recursiveAnchor**
  - Test: extra items disallowed for root
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>refRemote - 15 failures</summary>

- **remote ref**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **fragment within remote ref**
  - Test: remote fragment invalid
  - Expected: `invalid`, Got: `true`
- **anchor within remote ref**
  - Test: remote anchor invalid
  - Expected: `invalid`, Got: `true`
- **ref within remote ref**
  - Test: ref within ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change**
  - Test: base URI change ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder in subschema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`
- **remote ref with ref to defs**
  - Test: invalid
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier in remote ref**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **remote HTTP ref with different $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **remote HTTP ref with different URN $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **remote HTTP ref with nested absolute ref**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **$ref to $ref finds detached $anchor**
  - Test: non-number is invalid
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
<summary>unevaluatedItems - 21 failures</summary>

- **unevaluatedItems false**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems as schema**
  - Test: with invalid unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with tuple**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with ignored additionalItems**
  - Test: invalid under unevaluatedItems
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with ignored applicator additionalItems**
  - Test: invalid under unevaluatedItems
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with nested tuple**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with nested items**
  - Test: with invalid additional item
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with anyOf**
  - Test: when one schema matches and has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with anyOf**
  - Test: when two schemas match and has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with oneOf**
  - Test: with no unevaluated items
  - Expected: `valid`, Got: `false`
- **unevaluatedItems with not**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with if/then/else**
  - Test: when if matches and it has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with if/then/else**
  - Test: when if doesn't match and it has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with boolean schemas**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with $ref**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems before $ref**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with $recursiveRef**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems can't see inside cousins**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **item is evaluated in an uncle schema to unevaluatedItems**
  - Test: uncle keyword evaluation is not significant
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems can see annotations from if without then and else**
  - Test: invalid in case if is evaluated
  - Expected: `invalid`, Got: `true`
- **Evaluated items collection needs to consider instance location**
  - Test: with an unevaluated item that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 55 failures</summary>

- **unevaluatedProperties schema**
  - Test: with invalid unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties false**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with adjacent properties**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with adjacent patternProperties**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with nested properties**
  - Test: with additional properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with nested patternProperties**
  - Test: with additional properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with anyOf**
  - Test: when one matches and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with anyOf**
  - Test: when two match and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with oneOf**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with not**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else**
  - Test: when if is true and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else**
  - Test: when if is false and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, then not defined**
  - Test: when if is true and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, then not defined**
  - Test: when if is false and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, else not defined**
  - Test: when if is true and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, else not defined**
  - Test: when if is false and has no unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, else not defined**
  - Test: when if is false and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with dependentSchemas**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with boolean schemas**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with $ref**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties before $ref**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with $recursiveRef**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can't see inside cousins**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can't see inside cousins (reverse order)**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **nested unevaluatedProperties, outer true, inner false, properties outside**
  - Test: with no nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **nested unevaluatedProperties, outer true, inner false, properties outside**
  - Test: with nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **nested unevaluatedProperties, outer true, inner false, properties inside**
  - Test: with nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **cousin unevaluatedProperties, true and false, true with properties**
  - Test: with no nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **cousin unevaluatedProperties, true and false, true with properties**
  - Test: with nested unevaluated properties
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
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and x are valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and y are valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and b and x are valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and b and y are valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: a is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: b is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: c is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: d is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: xx is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: xx + foox is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: all is valid
  - Expected: `valid`, Got: `false`
- **dynamic evalation inside nested refs**
  - Test: all + foo is valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties not affected by propertyNames**
  - Test: string property is invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can see annotations from if without then and else**
  - Test: invalid in case if is evaluated
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with unevaluatedProperties**
  - Test: unevaluatedProperties doesn't consider dependentSchemas
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with unevaluatedProperties**
  - Test: unevaluatedProperties doesn't see bar when foo2 is absent
  - Expected: `invalid`, Got: `true`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 19 failures</summary>

- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>vocabulary - 1 failure</summary>

- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: applicator vocabulary still works
  - Expected: `invalid`, Got: `true`

</details>

## draft7

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 13/19 |
| additionalProperties | ⚠️ | 11/16 |
| allOf | ⚠️ | 15/30 |
| anyOf | ⚠️ | 14/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 47/54 |
| contains | ⚠️ | 11/21 |
| default | ⚠️ | 6/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 21/36 |
| enum | ⚠️ | 37/45 |
| exclusiveMaximum | ⚠️ | 2/4 |
| exclusiveMinimum | ⚠️ | 2/4 |
| format | ✅ | 102/102 |
| if-then-else | ⚠️ | 22/30 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 20/28 |
| maxItems | ⚠️ | 4/6 |
| maxLength | ⚠️ | 5/7 |
| maxProperties | ⚠️ | 7/10 |
| maximum | ⚠️ | 6/8 |
| minItems | ⚠️ | 4/6 |
| minLength | ⚠️ | 4/7 |
| minProperties | ⚠️ | 6/8 |
| minimum | ⚠️ | 8/11 |
| multipleOf | ⚠️ | 6/10 |
| not | ⚠️ | 37/38 |
| oneOf | ⚠️ | 18/27 |
| pattern | ⚠️ | 8/9 |
| patternProperties | ⚠️ | 13/23 |
| properties | ⚠️ | 16/28 |
| propertyNames | ⚠️ | 15/20 |
| ref | ⚠️ | 39/78 |
| refRemote | ⚠️ | 12/23 |
| required | ⚠️ | 10/16 |
| type | ✅ | 80/80 |
| uniqueItems | ⚠️ | 50/69 |

### Failures

<details>
<summary>additionalItems - 6 failures</summary>

- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `true`
- **when items is schema, additionalItems does nothing**
  - Test: invalid with a array of mixed types
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `true`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 5 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 15 failures</summary>

- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `true`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `true`
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
<summary>anyOf - 4 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>const - 7 failures</summary>

- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `false`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `false`
- **const with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **const with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": false} is valid
  - Expected: `valid`, Got: `false`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": true} is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>contains - 10 failures</summary>

- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with const keyword**
  - Test: array without item 5 is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema true**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: matches items, does not match contains
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: does not match items, matches contains
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: matches neither items nor contains
  - Expected: `invalid`, Got: `true`
- **contains with false if subschema**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>default - 1 failure</summary>

- **the default keyword does not do anything if the property is missing**
  - Test: missing properties are not filled in with the default
  - Expected: `valid`, Got: `false`

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
<summary>enum - 8 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMaximum - 2 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>exclusiveMinimum - 2 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>if-then-else - 8 failures</summary>

- **if and then without else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `true`
- **if and else without then**
  - Test: invalid through else
  - Expected: `invalid`, Got: `true`
- **validate against correct branch, then vs else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `true`
- **validate against correct branch, then vs else**
  - Test: invalid through else
  - Expected: `invalid`, Got: `true`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: other redirects to else and passes
  - Expected: `valid`, Got: `false`
- **then: false fails when condition matches**
  - Test: matches if → then=false → invalid
  - Expected: `invalid`, Got: `true`
- **else: false fails when condition does not match**
  - Test: does not match if → else executes → invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>infinite-loop-detection - 1 failure</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>items - 8 failures</summary>

- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `true`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `true`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 2 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 2 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 3 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 2 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 2 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 3 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 2 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 3 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 4 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>not - 1 failure</summary>

- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>oneOf - 9 failures</summary>

- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `false`
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
<summary>pattern - 1 failure</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>patternProperties - 10 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema false is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with a property matching both true and false is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 12 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`
- **properties with boolean schema**
  - Test: only 'false' property present is invalid
  - Expected: `invalid`, Got: `true`
- **properties with boolean schema**
  - Test: both properties present is invalid
  - Expected: `invalid`, Got: `true`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>propertyNames - 5 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames validation with pattern**
  - Test: non-matching property name is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with const**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with enum**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 39 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: slash invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: tilde invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: percent invalid
  - Expected: `invalid`, Got: `true`
- **nested refs**
  - Test: nested ref invalid
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref prevents a sibling $id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref to boolean schema false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `true`
- **refs with quote**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Reference an anchor with a non-relative URI**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `true`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `true`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `true`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `true`
- **$id must be resolved against nearest parent, not just immediate parent**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **simple URN base URI with $ref via the URN**
  - Test: invalid under the URN IDed schema
  - Expected: `invalid`, Got: `true`
- **simple URN base URI with JSON pointer**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with NSS**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with r-component**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with q-component**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with URN and JSON pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with URN and anchor ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **ref to if**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref to then**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref to else**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref with absolute-path-reference**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `true`
- **$id with file URI still resolves pointers - *nix**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **$id with file URI still resolves pointers - windows**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **empty tokens in $ref json-pointer**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>refRemote - 11 failures</summary>

- **remote ref**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **fragment within remote ref**
  - Test: remote fragment invalid
  - Expected: `invalid`, Got: `true`
- **ref within remote ref**
  - Test: ref within ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change**
  - Test: base URI change ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder in subschema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`
- **remote ref with ref to definitions**
  - Test: invalid
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier in remote ref**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **$ref to $ref finds location-independent $id**
  - Test: non-number is invalid
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
<summary>uniqueItems - 19 failures</summary>

- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

## draft6

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 13/19 |
| additionalProperties | ⚠️ | 11/16 |
| allOf | ⚠️ | 15/30 |
| anyOf | ⚠️ | 14/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 47/54 |
| contains | ⚠️ | 10/19 |
| default | ⚠️ | 6/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 21/36 |
| enum | ⚠️ | 37/45 |
| exclusiveMaximum | ⚠️ | 2/4 |
| exclusiveMinimum | ⚠️ | 2/4 |
| format | ✅ | 54/54 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 20/28 |
| maxItems | ⚠️ | 4/6 |
| maxLength | ⚠️ | 5/7 |
| maxProperties | ⚠️ | 7/10 |
| maximum | ⚠️ | 6/8 |
| minItems | ⚠️ | 4/6 |
| minLength | ⚠️ | 4/7 |
| minProperties | ⚠️ | 6/8 |
| minimum | ⚠️ | 8/11 |
| multipleOf | ⚠️ | 6/10 |
| not | ⚠️ | 37/38 |
| oneOf | ⚠️ | 18/27 |
| pattern | ⚠️ | 8/9 |
| patternProperties | ⚠️ | 13/23 |
| properties | ⚠️ | 16/28 |
| propertyNames | ⚠️ | 15/20 |
| ref | ⚠️ | 35/70 |
| refRemote | ⚠️ | 12/23 |
| required | ⚠️ | 10/16 |
| type | ✅ | 80/80 |
| uniqueItems | ⚠️ | 50/69 |

### Failures

<details>
<summary>additionalItems - 6 failures</summary>

- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `true`
- **when items is schema, additionalItems does nothing**
  - Test: invalid with a array of mixed types
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `true`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 5 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 15 failures</summary>

- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `true`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `true`
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
<summary>anyOf - 4 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>const - 7 failures</summary>

- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `false`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `false`
- **const with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **const with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": false} is valid
  - Expected: `valid`, Got: `false`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": true} is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>contains - 9 failures</summary>

- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with const keyword**
  - Test: array without item 5 is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema true**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: matches items, does not match contains
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: does not match items, matches contains
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: matches neither items nor contains
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>default - 1 failure</summary>

- **the default keyword does not do anything if the property is missing**
  - Test: missing properties are not filled in with the default
  - Expected: `valid`, Got: `false`

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
<summary>enum - 8 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMaximum - 2 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>exclusiveMinimum - 2 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>infinite-loop-detection - 1 failure</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>items - 8 failures</summary>

- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `true`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `true`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 2 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 2 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 3 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 2 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 2 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 3 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 2 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 3 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 4 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>not - 1 failure</summary>

- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>oneOf - 9 failures</summary>

- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `false`
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
<summary>pattern - 1 failure</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>patternProperties - 10 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema false is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with a property matching both true and false is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 12 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`
- **properties with boolean schema**
  - Test: only 'false' property present is invalid
  - Expected: `invalid`, Got: `true`
- **properties with boolean schema**
  - Test: both properties present is invalid
  - Expected: `invalid`, Got: `true`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>propertyNames - 5 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames validation with pattern**
  - Test: non-matching property name is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with const**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with enum**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 35 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: slash invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: tilde invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: percent invalid
  - Expected: `invalid`, Got: `true`
- **nested refs**
  - Test: nested ref invalid
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref prevents a sibling $id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref to boolean schema false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `true`
- **refs with quote**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Reference an anchor with a non-relative URI**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `true`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `true`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `true`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `true`
- **simple URN base URI with $ref via the URN**
  - Test: invalid under the URN IDed schema
  - Expected: `invalid`, Got: `true`
- **simple URN base URI with JSON pointer**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with NSS**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with r-component**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with q-component**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with URN and JSON pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with URN and anchor ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **ref with absolute-path-reference**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `true`
- **$id with file URI still resolves pointers - *nix**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **$id with file URI still resolves pointers - windows**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **empty tokens in $ref json-pointer**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>refRemote - 11 failures</summary>

- **remote ref**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **fragment within remote ref**
  - Test: remote fragment invalid
  - Expected: `invalid`, Got: `true`
- **ref within remote ref**
  - Test: ref within ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change**
  - Test: base URI change ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder in subschema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`
- **remote ref with ref to definitions**
  - Test: invalid
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier in remote ref**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **$ref to $ref finds location-independent $id**
  - Test: non-number is invalid
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
<summary>uniqueItems - 19 failures</summary>

- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

## draft4

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 12/17 |
| additionalProperties | ⚠️ | 11/16 |
| allOf | ⚠️ | 12/27 |
| anyOf | ⚠️ | 11/15 |
| default | ⚠️ | 6/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 16/29 |
| enum | ⚠️ | 41/49 |
| format | ✅ | 36/36 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 15/21 |
| maxItems | ⚠️ | 3/4 |
| maxLength | ⚠️ | 4/5 |
| maxProperties | ⚠️ | 6/8 |
| maximum | ⚠️ | 10/14 |
| minItems | ⚠️ | 3/4 |
| minLength | ⚠️ | 3/5 |
| minProperties | ⚠️ | 5/6 |
| minimum | ⚠️ | 12/17 |
| multipleOf | ⚠️ | 6/10 |
| not | ⚠️ | 19/20 |
| oneOf | ⚠️ | 14/23 |
| pattern | ⚠️ | 8/9 |
| patternProperties | ⚠️ | 11/18 |
| properties | ⚠️ | 14/24 |
| ref | ⚠️ | 23/45 |
| refRemote | ⚠️ | 9/17 |
| required | ⚠️ | 9/15 |
| type | ✅ | 79/79 |
| uniqueItems | ⚠️ | 50/69 |

### Failures

<details>
<summary>additionalItems - 5 failures</summary>

- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `true`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 5 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 15 failures</summary>

- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `true`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `true`
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
<summary>anyOf - 4 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>default - 1 failure</summary>

- **the default keyword does not do anything if the property is missing**
  - Test: missing properties are not filled in with the default
  - Expected: `valid`, Got: `false`

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
<summary>enum - 8 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>infinite-loop-detection - 1 failure</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>items - 6 failures</summary>

- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `true`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 1 failure</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 2 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 4 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation (explicit false exclusivity)**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 1 failure</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 2 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 1 failure</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 5 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation (explicit false exclusivity)**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 4 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>not - 1 failure</summary>

- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>oneOf - 9 failures</summary>

- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `false`
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
<summary>pattern - 1 failure</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>patternProperties - 7 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 10 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 22 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: slash invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: tilde invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: percent invalid
  - Expected: `invalid`, Got: `true`
- **nested refs**
  - Test: nested ref invalid
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `true`
- **refs with quote**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **id must be resolved against nearest parent, not just immediate parent**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **id with file URI still resolves pointers - *nix**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **id with file URI still resolves pointers - windows**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **empty tokens in $ref json-pointer**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>refRemote - 8 failures</summary>

- **remote ref**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **fragment within remote ref**
  - Test: remote fragment invalid
  - Expected: `invalid`, Got: `true`
- **ref within remote ref**
  - Test: ref within ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change**
  - Test: base URI change ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder in subschema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier in remote ref**
  - Test: string is invalid
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
<summary>uniqueItems - 19 failures</summary>

- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

## draft3

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 11/14 |
| additionalProperties | ⚠️ | 11/16 |
| default | ⚠️ | 6/7 |
| dependencies | ⚠️ | 11/18 |
| disallow | ⚠️ | 4/9 |
| divisibleBy | ⚠️ | 5/8 |
| enum | ⚠️ | 15/16 |
| extends | ⚠️ | 3/10 |
| format | ✅ | 60/60 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 5/7 |
| maxItems | ⚠️ | 3/4 |
| maxLength | ⚠️ | 4/5 |
| maximum | ⚠️ | 10/14 |
| minItems | ⚠️ | 3/4 |
| minLength | ⚠️ | 3/5 |
| minimum | ⚠️ | 9/13 |
| pattern | ⚠️ | 8/9 |
| patternProperties | ⚠️ | 10/17 |
| properties | ⚠️ | 9/15 |
| ref | ⚠️ | 14/27 |
| refRemote | ⚠️ | 4/8 |
| required | ⚠️ | 3/4 |
| type | ⚠️ | 73/80 |
| uniqueItems | ⚠️ | 45/62 |

### Failures

<details>
<summary>additionalItems - 3 failures</summary>

- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 5 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in extends are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>default - 1 failure</summary>

- **the default keyword does not do anything if the property is missing**
  - Test: missing properties are not filled in with the default
  - Expected: `valid`, Got: `false`

</details>

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
<summary>enum - 1 failure</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>extends - 7 failures</summary>

- **extends**
  - Test: mismatch extends
  - Expected: `invalid`, Got: `true`
- **extends**
  - Test: mismatch extended
  - Expected: `invalid`, Got: `true`
- **extends**
  - Test: wrong type
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
<summary>infinite-loop-detection - 1 failure</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>items - 2 failures</summary>

- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `true`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxItems - 1 failure</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 4 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation (explicit false exclusivity)**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 1 failure</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 2 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 4 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>pattern - 1 failure</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>patternProperties - 7 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 6 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 13 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: slash invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: tilde invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: percent invalid
  - Expected: `invalid`, Got: `true`
- **nested refs**
  - Test: nested ref invalid
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>refRemote - 4 failures</summary>

- **remote ref**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **fragment within remote ref**
  - Test: remote fragment invalid
  - Expected: `invalid`, Got: `true`
- **ref within remote ref**
  - Test: ref within ref invalid
  - Expected: `invalid`, Got: `true`
- **change resolution scope**
  - Test: changed scope ref invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 1 failure</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`

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

<details>
<summary>uniqueItems - 17 failures</summary>

- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

## v1

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 12/21 |
| allOf | ⚠️ | 15/30 |
| anchor | ⚠️ | 4/8 |
| anyOf | ⚠️ | 14/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 47/54 |
| contains | ⚠️ | 15/25 |
| content | ✅ | 18/18 |
| default | ⚠️ | 6/7 |
| defs | ⚠️ | 1/2 |
| dependentRequired | ⚠️ | 14/20 |
| dependentSchemas | ⚠️ | 10/20 |
| dynamicRef | ⚠️ | 12/27 |
| enum | ⚠️ | 37/45 |
| exclusiveMaximum | ⚠️ | 2/4 |
| exclusiveMinimum | ⚠️ | 2/4 |
| if-then-else | ⚠️ | 20/26 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 19/29 |
| maxContains | ⚠️ | 6/12 |
| maxItems | ⚠️ | 4/6 |
| maxLength | ⚠️ | 5/7 |
| maxProperties | ⚠️ | 7/10 |
| maximum | ⚠️ | 6/8 |
| minContains | ⚠️ | 14/28 |
| minItems | ⚠️ | 4/6 |
| minLength | ⚠️ | 4/7 |
| minProperties | ⚠️ | 6/8 |
| minimum | ⚠️ | 8/11 |
| multipleOf | ⚠️ | 6/10 |
| not | ⚠️ | 38/40 |
| oneOf | ⚠️ | 18/27 |
| pattern | ⚠️ | 8/9 |
| patternProperties | ⚠️ | 13/23 |
| prefixItems | ⚠️ | 9/11 |
| properties | ⚠️ | 16/28 |
| propertyNames | ⚠️ | 8/10 |
| ref | ⚠️ | 38/80 |
| refRemote | ⚠️ | 16/31 |
| required | ⚠️ | 10/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ⚠️ | 42/71 |
| unevaluatedProperties | ⚠️ | 68/123 |
| uniqueItems | ⚠️ | 50/69 |

### Failures

<details>
<summary>additionalProperties - 9 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with propertyNames**
  - Test: Valid against propertyNames, but not additionalProperties
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
<summary>allOf - 15 failures</summary>

- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `true`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `true`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `true`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `true`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `true`
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
<summary>anchor - 4 failures</summary>

- **Location-independent identifier**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier with absolute URI**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **same $anchor with different base uri**
  - Test: $ref does not resolve to /$defs/A/allOf/0
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>anyOf - 4 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>const - 7 failures</summary>

- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `false`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `false`
- **const with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **const with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": false} is valid
  - Expected: `valid`, Got: `false`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": true} is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>contains - 10 failures</summary>

- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with const keyword**
  - Test: array without item 5 is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema true**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: matches items, does not match contains
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: does not match items, matches contains
  - Expected: `invalid`, Got: `true`
- **items + contains**
  - Test: matches neither items nor contains
  - Expected: `invalid`, Got: `true`
- **contains with false if subschema**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>default - 1 failure</summary>

- **the default keyword does not do anything if the property is missing**
  - Test: missing properties are not filled in with the default
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>defs - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentRequired - 6 failures</summary>

- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependents required**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependents required**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `true`
- **multiple dependents required**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentSchemas - 10 failures</summary>

- **single dependency**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type other
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type both
  - Expected: `invalid`, Got: `true`
- **boolean subschemas**
  - Test: object with property having schema false is invalid
  - Expected: `invalid`, Got: `true`
- **boolean subschemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted tab invalid under dependent schema
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
<summary>dynamicRef - 15 failures</summary>

- **A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **A $dynamicRef resolves to the first $dynamicAnchor still in scope that is encountered when the schema is evaluated**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **A $dynamicRef with intermediate scopes that don't include a matching $dynamicAnchor does not affect dynamic scope resolution**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `true`
- **after leaving a dynamic scope, it is not used by a $dynamicRef**
  - Test: string matches /$defs/thingy, but the $dynamicRef does not stop here
  - Expected: `invalid`, Got: `true`
- **after leaving a dynamic scope, it is not used by a $dynamicRef**
  - Test: first_scope is not in dynamic scope for the $dynamicRef
  - Expected: `invalid`, Got: `true`
- **strict-tree schema, guards against misspelled properties**
  - Test: instance with misspelled field
  - Expected: `invalid`, Got: `true`
- **tests for implementation dynamic anchor and reference link**
  - Test: incorrect parent schema
  - Expected: `invalid`, Got: `true`
- **tests for implementation dynamic anchor and reference link**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $defs first**
  - Test: incorrect parent schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $defs first**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $ref first**
  - Test: incorrect parent schema
  - Expected: `invalid`, Got: `true`
- **$ref and $dynamicAnchor are independent of order - $ref first**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `true`
- **$ref to $dynamicRef finds detached $dynamicAnchor**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **$dynamicRef skips over intermediate resources - direct reference**
  - Test: string property fails
  - Expected: `invalid`, Got: `true`
- **$dynamicRef avoids the root of each schema, but scopes are still registered**
  - Test: data is not sufficient for schema at second#/$defs/length
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 8 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMaximum - 2 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>exclusiveMinimum - 2 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>if-then-else - 6 failures</summary>

- **if and then without else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `true`
- **if and else without then**
  - Test: invalid through else
  - Expected: `invalid`, Got: `true`
- **validate against correct branch, then vs else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `true`
- **validate against correct branch, then vs else**
  - Test: invalid through else
  - Expected: `invalid`, Got: `true`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: other redirects to else and passes
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>infinite-loop-detection - 1 failure</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>items - 10 failures</summary>

- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `true`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **prefixItems with no additional items allowed**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **items does not look in applicators, valid case**
  - Test: prefixItems in allOf does not constrain items, invalid case
  - Expected: `invalid`, Got: `true`
- **prefixItems validation adjusts the starting index for items**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `true`
- **items with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxContains - 6 failures</summary>

- **maxContains with contains**
  - Test: empty array
  - Expected: `invalid`, Got: `true`
- **maxContains with contains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains with contains**
  - Test: some elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains with contains, value with a decimal**
  - Test: too many elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **minContains < maxContains**
  - Test: array with actual < minContains < maxContains
  - Expected: `invalid`, Got: `true`
- **minContains < maxContains**
  - Test: array with minContains < maxContains < actual
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxItems - 2 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 2 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 3 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 2 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minContains - 14 failures</summary>

- **minContains=1 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **minContains=1 with contains**
  - Test: no elements match
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains**
  - Test: some elements match, invalid minContains
  - Expected: `invalid`, Got: `true`
- **minContains=2 with contains with a decimal value**
  - Test: one element matches, invalid minContains
  - Expected: `invalid`, Got: `true`
- **maxContains = minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **maxContains = minContains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `true`
- **maxContains = minContains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: invalid minContains
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: invalid maxContains
  - Expected: `invalid`, Got: `true`
- **maxContains < minContains**
  - Test: invalid maxContains and minContains
  - Expected: `invalid`, Got: `true`
- **minContains = 0 with maxContains**
  - Test: too many
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 2 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 3 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 2 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 3 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>multipleOf - 4 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>not - 2 failures</summary>

- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`
- **collect annotations inside a 'not', even if collection is disabled**
  - Test: unevaluated property
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 9 failures</summary>

- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `false`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `false`
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
<summary>pattern - 1 failure</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>patternProperties - 10 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `true`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `true`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema false is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties with boolean schemas**
  - Test: object with a property matching both true and false is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>prefixItems - 2 failures</summary>

- **a schema given for prefixItems**
  - Test: wrong types
  - Expected: `invalid`, Got: `true`
- **prefixItems with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 12 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`
- **properties with boolean schema**
  - Test: only 'false' property present is invalid
  - Expected: `invalid`, Got: `true`
- **properties with boolean schema**
  - Test: both properties present is invalid
  - Expected: `invalid`, Got: `true`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `true`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>propertyNames - 2 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 42 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: slash invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: tilde invalid
  - Expected: `invalid`, Got: `true`
- **escaped pointer ref**
  - Test: percent invalid
  - Expected: `invalid`, Got: `true`
- **nested refs**
  - Test: nested ref invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref to boolean schema false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `true`
- **refs with quote**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `true`
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `true`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `true`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `true`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `true`
- **$id must be resolved against nearest parent, not just immediate parent**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **order of evaluation: $id and $ref**
  - Test: data is invalid against first definition
  - Expected: `invalid`, Got: `true`
- **order of evaluation: $id and $anchor and $ref**
  - Test: data is invalid against first definition
  - Expected: `invalid`, Got: `true`
- **order of evaluation: $id and $ref on nested schema**
  - Test: data is invalid against nested sibling
  - Expected: `invalid`, Got: `true`
- **simple URN base URI with $ref via the URN**
  - Test: invalid under the URN IDed schema
  - Expected: `invalid`, Got: `true`
- **simple URN base URI with JSON pointer**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with NSS**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with r-component**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with q-component**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with f-component**
  - Test: is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with URN and JSON pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN base URI with URN and anchor ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **URN ref with nested pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `true`
- **ref to if**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref to then**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref to else**
  - Test: a non-integer is invalid due to the $ref
  - Expected: `invalid`, Got: `true`
- **ref with absolute-path-reference**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `true`
- **$id with file URI still resolves pointers - *nix**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **$id with file URI still resolves pointers - windows**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`
- **empty tokens in $ref json-pointer**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>refRemote - 15 failures</summary>

- **remote ref**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **fragment within remote ref**
  - Test: remote fragment invalid
  - Expected: `invalid`, Got: `true`
- **anchor within remote ref**
  - Test: remote anchor invalid
  - Expected: `invalid`, Got: `true`
- **ref within remote ref**
  - Test: ref within ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change**
  - Test: base URI change ref invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **base URI change - change folder in subschema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`
- **remote ref with ref to defs**
  - Test: invalid
  - Expected: `invalid`, Got: `true`
- **Location-independent identifier in remote ref**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **remote HTTP ref with different $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **remote HTTP ref with different URN $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **remote HTTP ref with nested absolute ref**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **$ref to $ref finds detached $anchor**
  - Test: non-number is invalid
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
<summary>unevaluatedItems - 29 failures</summary>

- **unevaluatedItems false**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems as schema**
  - Test: with invalid unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with tuple**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with items**
  - Test: invalid under items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with nested tuple**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with nested items**
  - Test: with invalid additional item
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with anyOf**
  - Test: when one schema matches and has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with anyOf**
  - Test: when two schemas match and has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with oneOf**
  - Test: with no unevaluated items
  - Expected: `valid`, Got: `false`
- **unevaluatedItems with not**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with if/then/else**
  - Test: when if matches and it has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with if/then/else**
  - Test: when if doesn't match and it has unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with boolean schemas**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with $ref**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems before $ref**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with $dynamicRef**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems can't see inside cousins**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **item is evaluated in an uncle schema to unevaluatedItems**
  - Test: uncle keyword evaluation is not significant
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems depends on adjacent contains**
  - Test: contains fails, second item is not evaluated
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems depends on adjacent contains**
  - Test: contains passes, second item is not evaluated
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems depends on multiple nested contains**
  - Test: 7 not evaluated, fails unevaluatedItems
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems and contains interact to control item dependency relationship**
  - Test: only b's are invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems and contains interact to control item dependency relationship**
  - Test: only c's are invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems and contains interact to control item dependency relationship**
  - Test: only b's and c's are invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems and contains interact to control item dependency relationship**
  - Test: only a's and c's are invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with minContains = 0**
  - Test: no items evaluated by contains
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems with minContains = 0**
  - Test: some but not all items evaluated by contains
  - Expected: `invalid`, Got: `true`
- **unevaluatedItems can see annotations from if without then and else**
  - Test: invalid in case if is evaluated
  - Expected: `invalid`, Got: `true`
- **Evaluated items collection needs to consider instance location**
  - Test: with an unevaluated item that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 55 failures</summary>

- **unevaluatedProperties schema**
  - Test: with invalid unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties false**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with adjacent properties**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with adjacent patternProperties**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with nested properties**
  - Test: with additional properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with nested patternProperties**
  - Test: with additional properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with anyOf**
  - Test: when one matches and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with anyOf**
  - Test: when two match and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with oneOf**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with not**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else**
  - Test: when if is true and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else**
  - Test: when if is false and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, then not defined**
  - Test: when if is true and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, then not defined**
  - Test: when if is false and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, else not defined**
  - Test: when if is true and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, else not defined**
  - Test: when if is false and has no unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with if/then/else, else not defined**
  - Test: when if is false and has unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with dependentSchemas**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with boolean schemas**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with $ref**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties before $ref**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties with $dynamicRef**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can't see inside cousins**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can't see inside cousins (reverse order)**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **nested unevaluatedProperties, outer true, inner false, properties outside**
  - Test: with no nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **nested unevaluatedProperties, outer true, inner false, properties outside**
  - Test: with nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **nested unevaluatedProperties, outer true, inner false, properties inside**
  - Test: with nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **cousin unevaluatedProperties, true and false, true with properties**
  - Test: with no nested unevaluated properties
  - Expected: `invalid`, Got: `true`
- **cousin unevaluatedProperties, true and false, true with properties**
  - Test: with nested unevaluated properties
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
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and x are valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and y are valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and b and x are valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties + ref inside allOf / oneOf**
  - Test: a and b and y are valid
  - Expected: `valid`, Got: `false`
- **dynamic evaluation inside nested refs**
  - Test: a is valid
  - Expected: `valid`, Got: `false`
- **dynamic evaluation inside nested refs**
  - Test: b is valid
  - Expected: `valid`, Got: `false`
- **dynamic evaluation inside nested refs**
  - Test: c is valid
  - Expected: `valid`, Got: `false`
- **dynamic evaluation inside nested refs**
  - Test: d is valid
  - Expected: `valid`, Got: `false`
- **dynamic evaluation inside nested refs**
  - Test: xx is valid
  - Expected: `valid`, Got: `false`
- **dynamic evaluation inside nested refs**
  - Test: xx + foox is valid
  - Expected: `valid`, Got: `false`
- **dynamic evaluation inside nested refs**
  - Test: all is valid
  - Expected: `valid`, Got: `false`
- **dynamic evaluation inside nested refs**
  - Test: all + foo is valid
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties not affected by propertyNames**
  - Test: string property is invalid
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can see annotations from if without then and else**
  - Test: invalid in case if is evaluated
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with unevaluatedProperties**
  - Test: unevaluatedProperties doesn't consider dependentSchemas
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with unevaluatedProperties**
  - Test: unevaluatedProperties doesn't see bar when foo2 is absent
  - Expected: `invalid`, Got: `true`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 19 failures</summary>

- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `true`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

