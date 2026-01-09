# @xschemadev/typebox Compliance Report

Generated: 2026-01-09T11:00:25Z

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2020-12 | 710 | 320 | 0 | 68.9% |
| draft2019-09 | 706 | 321 | 0 | 68.7% |
| draft7 | 615 | 257 | 0 | 70.5% |
| draft6 | 546 | 246 | 0 | 68.9% |
| draft4 | 388 | 211 | 0 | 64.8% |
| draft3 | 299 | 129 | 0 | 69.9% |
| v1 | 560 | 319 | 0 | 63.7% |

## Badges

![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-68.9%25-red)
![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-68.7%25-red)
![draft7](https://img.shields.io/badge/draft7%20compliance-70.5%25-red)
![draft6](https://img.shields.io/badge/draft6%20compliance-68.9%25-red)
![draft4](https://img.shields.io/badge/draft4%20compliance-64.8%25-red)
![draft3](https://img.shields.io/badge/draft3%20compliance-69.9%25-red)
![v1](https://img.shields.io/badge/v1%20compliance-63.7%25-red)

## draft2020-12

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 9/21 |
| allOf | ⚠️ | 15/30 |
| anchor | ✅ | 0/0 |
| anyOf | ⚠️ | 15/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 47/54 |
| contains | ⚠️ | 9/21 |
| content | ✅ | 18/18 |
| default | ⚠️ | 6/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ⚠️ | 10/20 |
| dependentSchemas | ⚠️ | 7/20 |
| dynamicRef | ⚠️ | 8/19 |
| enum | ⚠️ | 33/45 |
| exclusiveMaximum | ⚠️ | 1/4 |
| exclusiveMinimum | ⚠️ | 1/4 |
| format | ✅ | 133/133 |
| if-then-else | ⚠️ | 20/30 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 20/29 |
| maxContains | ⚠️ | 6/12 |
| maxItems | ⚠️ | 3/6 |
| maxLength | ⚠️ | 4/7 |
| maxProperties | ⚠️ | 4/10 |
| maximum | ⚠️ | 5/8 |
| minContains | ⚠️ | 14/28 |
| minItems | ⚠️ | 3/6 |
| minLength | ⚠️ | 3/7 |
| minProperties | ⚠️ | 3/8 |
| minimum | ⚠️ | 6/11 |
| multipleOf | ⚠️ | 5/10 |
| not | ⚠️ | 36/38 |
| oneOf | ⚠️ | 15/27 |
| pattern | ⚠️ | 2/9 |
| patternProperties | ⚠️ | 10/23 |
| prefixItems | ⚠️ | 4/11 |
| properties | ⚠️ | 12/28 |
| propertyNames | ⚠️ | 12/20 |
| ref | ⚠️ | 22/35 |
| refRemote | ⚠️ | 15/17 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 78/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 20/46 |
| uniqueItems | ⚠️ | 48/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 12 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>anyOf - 3 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
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
<summary>contains - 12 failures</summary>

- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
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
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `false`
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
  - Test: an explicit property value is checked against maximum (failing)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentRequired - 10 failures</summary>

- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **empty dependents**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
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
<summary>dependentSchemas - 13 failures</summary>

- **single dependency**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type other
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type both
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>dynamicRef - 11 failures</summary>

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

</details>

<details>
<summary>enum - 12 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: wrong foo value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: wrong bar value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing required property is invalid
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing all properties is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>exclusiveMaximum - 3 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMinimum - 3 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>if-then-else - 10 failures</summary>

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
- **if with boolean schema true**
  - Test: boolean schema true in if always chooses the then path (invalid)
  - Expected: `invalid`, Got: `true`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: no redirects to then and fails
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: invalid redirects to else and fails
  - Expected: `invalid`, Got: `true`
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
<summary>items - 9 failures</summary>

- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`
- **prefixItems with no additional items allowed**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **prefixItems with no additional items allowed**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `false`
- **prefixItems with no additional items allowed**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `false`
- **items does not look in applicators, valid case**
  - Test: prefixItems in allOf does not constrain items, valid case
  - Expected: `valid`, Got: `false`
- **prefixItems validation adjusts the starting index for items**
  - Test: valid items
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
<summary>maxItems - 3 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 3 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 6 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 3 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
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
<summary>minItems - 3 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 4 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 5 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 5 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 5 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **float division = inf**
  - Test: always invalid, but naive implementations may raise an overflow error
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>not - 2 failures</summary>

- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>oneOf - 12 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with boolean schemas, all true**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with boolean schemas, more than one true**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>pattern - 7 failures</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 13 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>prefixItems - 7 failures</summary>

- **a schema given for prefixItems**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **a schema given for prefixItems**
  - Test: array with additional items
  - Expected: `valid`, Got: `false`
- **a schema given for prefixItems**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **a schema given for prefixItems**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **prefixItems with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **prefixItems with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`
- **additional items are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 16 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>propertyNames - 8 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
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

</details>

<details>
<summary>refRemote - 2 failures</summary>

- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`
- **remote ref with ref to defs**
  - Test: invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>type - 2 failures</summary>

- **integer type matches integers**
  - Test: a float is not an integer
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 26 failures</summary>

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
- **unevaluatedProperties can't see inside cousins**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can't see inside cousins (reverse order)**
  - Test: always fails
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
- **non-object instances are valid**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties not affected by propertyNames**
  - Test: string property is invalid
  - Expected: `invalid`, Got: `true`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 21 failures</summary>

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
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`

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
| additionalItems | ⚠️ | 12/19 |
| additionalProperties | ⚠️ | 9/21 |
| allOf | ⚠️ | 15/30 |
| anchor | ✅ | 0/0 |
| anyOf | ⚠️ | 15/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 47/54 |
| contains | ⚠️ | 9/21 |
| content | ✅ | 18/18 |
| default | ⚠️ | 6/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ⚠️ | 10/20 |
| dependentSchemas | ⚠️ | 7/20 |
| enum | ⚠️ | 33/45 |
| exclusiveMaximum | ⚠️ | 1/4 |
| exclusiveMinimum | ⚠️ | 1/4 |
| format | ✅ | 114/114 |
| if-then-else | ⚠️ | 20/30 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 18/28 |
| maxContains | ⚠️ | 6/12 |
| maxItems | ⚠️ | 3/6 |
| maxLength | ⚠️ | 4/7 |
| maxProperties | ⚠️ | 4/10 |
| maximum | ⚠️ | 5/8 |
| minContains | ⚠️ | 14/28 |
| minItems | ⚠️ | 3/6 |
| minLength | ⚠️ | 3/7 |
| minProperties | ⚠️ | 3/8 |
| minimum | ⚠️ | 6/11 |
| multipleOf | ⚠️ | 5/10 |
| not | ⚠️ | 36/38 |
| oneOf | ⚠️ | 15/27 |
| pattern | ⚠️ | 2/9 |
| patternProperties | ⚠️ | 10/23 |
| properties | ⚠️ | 12/28 |
| propertyNames | ⚠️ | 12/20 |
| recursiveRef | ⚠️ | 18/30 |
| ref | ⚠️ | 22/35 |
| refRemote | ⚠️ | 15/17 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 78/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 19/44 |
| uniqueItems | ⚠️ | 48/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalItems - 7 failures</summary>

- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `false`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `false`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>additionalProperties - 12 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>anyOf - 3 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
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
<summary>contains - 12 failures</summary>

- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
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
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `false`
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
  - Test: an explicit property value is checked against maximum (failing)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependentRequired - 10 failures</summary>

- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **empty dependents**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
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
<summary>dependentSchemas - 13 failures</summary>

- **single dependency**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type other
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type both
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>enum - 12 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: wrong foo value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: wrong bar value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing required property is invalid
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing all properties is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>exclusiveMaximum - 3 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMinimum - 3 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>if-then-else - 10 failures</summary>

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
- **if with boolean schema true**
  - Test: boolean schema true in if always chooses the then path (invalid)
  - Expected: `invalid`, Got: `true`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: no redirects to then and fails
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: invalid redirects to else and fails
  - Expected: `invalid`, Got: `true`
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
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`
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
<summary>maxItems - 3 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 3 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 6 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 3 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
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
<summary>minItems - 3 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 4 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 5 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 5 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 5 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **float division = inf**
  - Test: always invalid, but naive implementations may raise an overflow error
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>not - 2 failures</summary>

- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>oneOf - 12 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with boolean schemas, all true**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with boolean schemas, more than one true**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>pattern - 7 failures</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 13 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>properties - 16 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>propertyNames - 8 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
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

</details>

<details>
<summary>refRemote - 2 failures</summary>

- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`
- **remote ref with ref to defs**
  - Test: invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>type - 2 failures</summary>

- **integer type matches integers**
  - Test: a float is not an integer
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 25 failures</summary>

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
- **unevaluatedProperties can't see inside cousins**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can't see inside cousins (reverse order)**
  - Test: always fails
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
- **non-object instances are valid**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties not affected by propertyNames**
  - Test: string property is invalid
  - Expected: `invalid`, Got: `true`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 21 failures</summary>

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
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`

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
| additionalItems | ⚠️ | 12/19 |
| additionalProperties | ⚠️ | 8/16 |
| allOf | ⚠️ | 15/30 |
| anyOf | ⚠️ | 15/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 47/54 |
| contains | ⚠️ | 9/21 |
| default | ⚠️ | 6/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 17/36 |
| enum | ⚠️ | 33/45 |
| exclusiveMaximum | ⚠️ | 1/4 |
| exclusiveMinimum | ⚠️ | 1/4 |
| format | ✅ | 102/102 |
| if-then-else | ⚠️ | 20/30 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 18/28 |
| maxItems | ⚠️ | 3/6 |
| maxLength | ⚠️ | 4/7 |
| maxProperties | ⚠️ | 4/10 |
| maximum | ⚠️ | 5/8 |
| minItems | ⚠️ | 3/6 |
| minLength | ⚠️ | 3/7 |
| minProperties | ⚠️ | 3/8 |
| minimum | ⚠️ | 6/11 |
| multipleOf | ⚠️ | 5/10 |
| not | ⚠️ | 36/38 |
| oneOf | ⚠️ | 15/27 |
| pattern | ⚠️ | 2/9 |
| patternProperties | ⚠️ | 10/23 |
| properties | ⚠️ | 12/28 |
| propertyNames | ⚠️ | 12/20 |
| ref | ⚠️ | 24/36 |
| refRemote | ⚠️ | 13/17 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 78/80 |
| uniqueItems | ⚠️ | 48/69 |

### Failures

<details>
<summary>additionalItems - 7 failures</summary>

- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `false`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `false`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>additionalProperties - 8 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>anyOf - 3 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
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
<summary>contains - 12 failures</summary>

- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
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
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `false`
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
  - Test: an explicit property value is checked against maximum (failing)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 19 failures</summary>

- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **dependencies with empty array**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
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
<summary>enum - 12 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: wrong foo value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: wrong bar value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing required property is invalid
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing all properties is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>exclusiveMaximum - 3 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMinimum - 3 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>if-then-else - 10 failures</summary>

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
- **if with boolean schema true**
  - Test: boolean schema true in if always chooses the then path (invalid)
  - Expected: `invalid`, Got: `true`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: no redirects to then and fails
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: invalid redirects to else and fails
  - Expected: `invalid`, Got: `true`
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
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 3 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 3 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 6 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 3 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 3 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 4 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 5 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 5 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 5 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **float division = inf**
  - Test: always invalid, but naive implementations may raise an overflow error
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>not - 2 failures</summary>

- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>oneOf - 12 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with boolean schemas, all true**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with boolean schemas, more than one true**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>pattern - 7 failures</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 13 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>properties - 16 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>propertyNames - 8 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>ref - 12 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
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
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
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

</details>

<details>
<summary>refRemote - 4 failures</summary>

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

</details>

<details>
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>type - 2 failures</summary>

- **integer type matches integers**
  - Test: a float is not an integer
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 21 failures</summary>

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
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`

</details>

## draft6

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 12/19 |
| additionalProperties | ⚠️ | 8/16 |
| allOf | ⚠️ | 15/30 |
| anyOf | ⚠️ | 15/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 47/54 |
| contains | ⚠️ | 8/19 |
| default | ⚠️ | 6/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 17/36 |
| enum | ⚠️ | 33/45 |
| exclusiveMaximum | ⚠️ | 1/4 |
| exclusiveMinimum | ⚠️ | 1/4 |
| format | ✅ | 54/54 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 18/28 |
| maxItems | ⚠️ | 3/6 |
| maxLength | ⚠️ | 4/7 |
| maxProperties | ⚠️ | 4/10 |
| maximum | ⚠️ | 5/8 |
| minItems | ⚠️ | 3/6 |
| minLength | ⚠️ | 3/7 |
| minProperties | ⚠️ | 3/8 |
| minimum | ⚠️ | 6/11 |
| multipleOf | ⚠️ | 5/10 |
| not | ⚠️ | 36/38 |
| oneOf | ⚠️ | 15/27 |
| pattern | ⚠️ | 2/9 |
| patternProperties | ⚠️ | 10/23 |
| properties | ⚠️ | 12/28 |
| propertyNames | ⚠️ | 12/20 |
| ref | ⚠️ | 24/36 |
| refRemote | ⚠️ | 13/17 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 78/80 |
| uniqueItems | ⚠️ | 48/69 |

### Failures

<details>
<summary>additionalItems - 7 failures</summary>

- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `false`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `false`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>additionalProperties - 8 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>anyOf - 3 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
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
<summary>contains - 11 failures</summary>

- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
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
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `false`
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
  - Test: an explicit property value is checked against maximum (failing)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 19 failures</summary>

- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **dependencies with empty array**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
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
<summary>enum - 12 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: wrong foo value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: wrong bar value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing required property is invalid
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing all properties is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>exclusiveMaximum - 3 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMinimum - 3 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: ignores non-numbers
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
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 3 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 3 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 6 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 3 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 3 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 4 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 5 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 5 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 5 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **float division = inf**
  - Test: always invalid, but naive implementations may raise an overflow error
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>not - 2 failures</summary>

- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>oneOf - 12 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with boolean schemas, all true**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with boolean schemas, more than one true**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>pattern - 7 failures</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 13 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>properties - 16 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>propertyNames - 8 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
- **propertyNames validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>ref - 12 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
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
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
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

</details>

<details>
<summary>refRemote - 4 failures</summary>

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

</details>

<details>
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>type - 2 failures</summary>

- **integer type matches integers**
  - Test: a float is not an integer
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 21 failures</summary>

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
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`

</details>

## draft4

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 10/17 |
| additionalProperties | ⚠️ | 8/16 |
| allOf | ⚠️ | 12/27 |
| anyOf | ⚠️ | 12/15 |
| default | ⚠️ | 6/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 13/29 |
| enum | ⚠️ | 37/49 |
| format | ✅ | 36/36 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 13/21 |
| maxItems | ⚠️ | 2/4 |
| maxLength | ⚠️ | 3/5 |
| maxProperties | ⚠️ | 3/8 |
| maximum | ⚠️ | 8/14 |
| minItems | ⚠️ | 2/4 |
| minLength | ⚠️ | 2/5 |
| minProperties | ⚠️ | 2/6 |
| minimum | ⚠️ | 9/17 |
| multipleOf | ⚠️ | 5/10 |
| not | ⚠️ | 18/20 |
| oneOf | ⚠️ | 13/23 |
| pattern | ⚠️ | 2/9 |
| patternProperties | ⚠️ | 8/18 |
| properties | ⚠️ | 10/24 |
| ref | ⚠️ | 17/33 |
| refRemote | ⚠️ | 6/15 |
| required | ⚠️ | 4/15 |
| type | ⚠️ | 77/79 |
| uniqueItems | ⚠️ | 48/69 |

### Failures

<details>
<summary>additionalItems - 7 failures</summary>

- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `false`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `false`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>additionalProperties - 8 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>anyOf - 3 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
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
  - Test: an explicit property value is checked against maximum (failing)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 16 failures</summary>

- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>enum - 12 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: wrong foo value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: wrong bar value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing required property is invalid
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing all properties is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>items - 8 failures</summary>

- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
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
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxLength - 2 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxProperties - 5 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 6 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation (explicit false exclusivity)**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 2 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 3 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 4 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minimum - 8 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation (explicit false exclusivity)**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 5 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **float division = inf**
  - Test: invalid, but naive implementations may raise an overflow error
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>not - 2 failures</summary>

- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>oneOf - 10 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>pattern - 7 failures</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 10 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>properties - 14 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>ref - 16 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
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
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
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
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>type - 2 failures</summary>

- **integer type matches integers**
  - Test: a float is not an integer
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 21 failures</summary>

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
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`

</details>

## draft3

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 8/14 |
| additionalProperties | ⚠️ | 8/16 |
| default | ⚠️ | 6/7 |
| dependencies | ⚠️ | 8/18 |
| disallow | ⚠️ | 8/9 |
| divisibleBy | ⚠️ | 4/8 |
| enum | ⚠️ | 11/16 |
| extends | ⚠️ | 3/10 |
| format | ✅ | 60/60 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 6/7 |
| maxItems | ⚠️ | 2/4 |
| maxLength | ⚠️ | 3/5 |
| maximum | ⚠️ | 8/14 |
| minItems | ⚠️ | 2/4 |
| minLength | ⚠️ | 2/5 |
| minimum | ⚠️ | 7/13 |
| pattern | ⚠️ | 2/9 |
| patternProperties | ⚠️ | 8/17 |
| properties | ⚠️ | 7/15 |
| ref | ⚠️ | 12/21 |
| refRemote | ⚠️ | 6/8 |
| required | ⚠️ | 3/4 |
| type | ⚠️ | 71/80 |
| uniqueItems | ⚠️ | 43/62 |

### Failures

<details>
<summary>additionalItems - 6 failures</summary>

- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `false`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `false`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>additionalProperties - 8 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
  - Test: an explicit property value is checked against maximum (failing)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 10 failures</summary>

- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>disallow - 1 failure</summary>

- **multiple disallow subschema**
  - Test: other match
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>divisibleBy - 4 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **by number**
  - Test: 35 is not divisible by 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not divisible by 0.0001
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 5 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: wrong foo value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: wrong bar value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing required property is invalid
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing all properties is invalid
  - Expected: `invalid`, Got: `true`

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
<summary>items - 1 failure</summary>

- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 2 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxLength - 2 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maximum - 6 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation (explicit false exclusivity)**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 2 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 3 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 6 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>pattern - 7 failures</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 9 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>properties - 8 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>ref - 9 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`

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
<summary>type - 9 failures</summary>

- **integer type matches integers**
  - Test: a float is not an integer
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
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
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`

</details>

## v1

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 9/21 |
| allOf | ⚠️ | 15/30 |
| anchor | ✅ | 0/0 |
| anyOf | ⚠️ | 15/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 47/54 |
| contains | ⚠️ | 9/25 |
| content | ✅ | 18/18 |
| default | ⚠️ | 6/7 |
| defs | ❌ | 0/2 |
| dependentRequired | ⚠️ | 10/20 |
| dependentSchemas | ⚠️ | 7/20 |
| dynamicRef | ⚠️ | 5/13 |
| enum | ⚠️ | 33/45 |
| exclusiveMaximum | ⚠️ | 1/4 |
| exclusiveMinimum | ⚠️ | 1/4 |
| if-then-else | ⚠️ | 18/26 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 20/29 |
| maxContains | ⚠️ | 6/12 |
| maxItems | ⚠️ | 3/6 |
| maxLength | ⚠️ | 4/7 |
| maxProperties | ⚠️ | 4/10 |
| maximum | ⚠️ | 5/8 |
| minContains | ⚠️ | 14/28 |
| minItems | ⚠️ | 3/6 |
| minLength | ⚠️ | 3/7 |
| minProperties | ⚠️ | 3/8 |
| minimum | ⚠️ | 6/11 |
| multipleOf | ⚠️ | 5/10 |
| not | ⚠️ | 36/38 |
| oneOf | ⚠️ | 15/27 |
| pattern | ⚠️ | 2/9 |
| patternProperties | ⚠️ | 10/23 |
| prefixItems | ⚠️ | 4/11 |
| properties | ⚠️ | 12/28 |
| propertyNames | ⚠️ | 5/10 |
| ref | ⚠️ | 22/38 |
| refRemote | ⚠️ | 15/17 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 78/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 19/44 |
| uniqueItems | ⚠️ | 48/69 |

### Failures

<details>
<summary>additionalProperties - 12 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>anyOf - 3 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
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
<summary>contains - 16 failures</summary>

- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `true`
- **contains keyword validation**
  - Test: not array or object is valid
  - Expected: `valid`, Got: `false`
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
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid - string
  - Expected: `valid`, Got: `false`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid - object
  - Expected: `valid`, Got: `false`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid - number
  - Expected: `valid`, Got: `false`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid - boolean
  - Expected: `valid`, Got: `false`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid - null
  - Expected: `valid`, Got: `false`
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
  - Test: an explicit property value is checked against maximum (failing)
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
<summary>dependentRequired - 10 failures</summary>

- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **empty dependents**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
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
<summary>dependentSchemas - 13 failures</summary>

- **single dependency**
  - Test: wrong type
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type other
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: wrong type both
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>dynamicRef - 8 failures</summary>

- **A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
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

</details>

<details>
<summary>enum - 12 failures</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: wrong foo value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: wrong bar value
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing required property is invalid
  - Expected: `invalid`, Got: `true`
- **enums in properties**
  - Test: missing all properties is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>exclusiveMaximum - 3 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMinimum - 3 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `true`
- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

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
- **if with boolean schema true**
  - Test: boolean schema true in if always chooses the then path (invalid)
  - Expected: `invalid`, Got: `true`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: no redirects to then and fails
  - Expected: `invalid`, Got: `true`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: invalid redirects to else and fails
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>infinite-loop-detection - 1 failure</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>items - 9 failures</summary>

- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`
- **prefixItems with no additional items allowed**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **prefixItems with no additional items allowed**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `false`
- **prefixItems with no additional items allowed**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `false`
- **items does not look in applicators, valid case**
  - Test: prefixItems in allOf does not constrain items, valid case
  - Expected: `valid`, Got: `false`
- **prefixItems validation adjusts the starting index for items**
  - Test: valid items
  - Expected: `valid`, Got: `false`

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
<summary>maxItems - 3 failures</summary>

- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxLength - 3 failures</summary>

- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxProperties - 6 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maximum - 3 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
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
<summary>minItems - 3 failures</summary>

- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minLength - 4 failures</summary>

- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `true`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minProperties - 5 failures</summary>

- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minimum - 5 failures</summary>

- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 5 failures</summary>

- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `true`
- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `true`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `true`
- **float division = inf**
  - Test: always invalid, but naive implementations may raise an overflow error
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>not - 2 failures</summary>

- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>oneOf - 12 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with boolean schemas, all true**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with boolean schemas, more than one true**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `true`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>pattern - 7 failures</summary>

- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `true`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 13 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>prefixItems - 7 failures</summary>

- **a schema given for prefixItems**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **a schema given for prefixItems**
  - Test: array with additional items
  - Expected: `valid`, Got: `false`
- **a schema given for prefixItems**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **a schema given for prefixItems**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **prefixItems with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **prefixItems with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`
- **additional items are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 16 failures</summary>

- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `true`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
- **propertyNames validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 16 failures</summary>

- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **relative pointer ref to object**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `true`
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
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
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`

</details>

<details>
<summary>refRemote - 2 failures</summary>

- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`
- **remote ref with ref to defs**
  - Test: invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
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
<summary>type - 2 failures</summary>

- **integer type matches integers**
  - Test: a float is not an integer
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 25 failures</summary>

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
- **unevaluatedProperties can't see inside cousins**
  - Test: always fails
  - Expected: `invalid`, Got: `true`
- **unevaluatedProperties can't see inside cousins (reverse order)**
  - Test: always fails
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
- **non-object instances are valid**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **unevaluatedProperties not affected by propertyNames**
  - Test: string property is invalid
  - Expected: `invalid`, Got: `true`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 21 failures</summary>

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
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `false`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `false`

</details>

