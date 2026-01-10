# @xschemadev/typebox Compliance Report

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2019-09 | 337 | 522 | 0 | 39.2% |
| draft2020-12 | 356 | 515 | 0 | 40.9% |
| draft3 | 161 | 246 | 0 | 39.6% |
| draft4 | 188 | 363 | 0 | 34.1% |
| draft6 | 282 | 457 | 0 | 38.2% |
| draft7 | 338 | 477 | 0 | 41.5% |

## Badges

![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-39.2%25-red)
![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-40.9%25-red)
![draft3](https://img.shields.io/badge/draft3%20compliance-39.6%25-red)
![draft4](https://img.shields.io/badge/draft4%20compliance-34.1%25-red)
![draft6](https://img.shields.io/badge/draft6%20compliance-38.2%25-red)
![draft7](https://img.shields.io/badge/draft7%20compliance-41.5%25-red)

## draft2019-09

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ❌ | 0/19 |
| additionalProperties | ❌ | 0/21 |
| allOf | ⚠️ | 11/30 |
| anchor | ✅ | 0/0 |
| anyOf | ⚠️ | 10/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 33/54 |
| contains | ❌ | 0/21 |
| content | ✅ | 18/18 |
| default | ⚠️ | 3/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ❌ | 0/20 |
| dependentSchemas | ❌ | 0/20 |
| enum | ⚠️ | 25/45 |
| exclusiveMaximum | ❌ | 0/4 |
| exclusiveMinimum | ❌ | 0/4 |
| format | ✅ | 114/114 |
| if-then-else | ⚠️ | 8/26 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ⚠️ | 3/28 |
| maxContains | ⚠️ | 2/12 |
| maxItems | ❌ | 0/6 |
| maxLength | ❌ | 0/7 |
| maxProperties | ❌ | 0/10 |
| maximum | ❌ | 0/8 |
| minContains | ⚠️ | 2/28 |
| minItems | ❌ | 0/6 |
| minLength | ❌ | 0/7 |
| minProperties | ❌ | 0/8 |
| minimum | ❌ | 0/11 |
| multipleOf | ⚠️ | 1/10 |
| not | ✅ | 0/0 |
| oneOf | ⚠️ | 7/27 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/23 |
| properties | ❌ | 0/28 |
| propertyNames | ❌ | 0/20 |
| recursiveRef | ✅ | 0/0 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ❌ | 0/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ❌ | 0/69 |
| vocabulary | ⚠️ | 2/5 |

### Failures

<details>
<summary>additionalItems - 19 failures</summary>

- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as false without items**
  - Test: items defaults to empty schema so everything is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems with heterogeneous array**
  - Test: valid instance
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: equal number of items present
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `error: Unknown type`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `error: Unknown type`
- **when items is schema, additionalItems does nothing**
  - Test: invalid with a array of mixed types
  - Expected: `invalid`, Got: `error: Unknown type`
- **when items is schema, additionalItems does nothing**
  - Test: valid with a array of type integers
  - Expected: `valid`, Got: `error: Unknown type`
- **when items is schema, boolean additionalItems does nothing**
  - Test: all items match schema
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>additionalProperties - 21 failures</summary>

- **additionalProperties are allowed by default**
  - Test: additional properties are allowed
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with propertyNames**
  - Test: Valid against both keywords
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with propertyNames**
  - Test: Valid against propertyNames, but not additionalProperties
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar even when foo2 is present
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties doesn't consider dependentSchemas
  - Expected: `invalid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: matching the pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>allOf - 19 failures</summary>

- **allOf**
  - Test: allOf
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: true
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf simple types**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>anyOf - 8 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf**
  - Test: second anyOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: both anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: first anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: second anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf with base schema**
  - Test: one anyOf valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>const - 21 failures</summary>

- **const with [false] does not match [0]**
  - Test: [0.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [false] does not match [0]**
  - Test: [0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [1.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [1] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with array**
  - Test: another array item is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with array**
  - Test: array with additional items is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with null**
  - Test: not null is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with null**
  - Test: null is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with object**
  - Test: another object is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with object**
  - Test: another type is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": 0.0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": 0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": false} is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": 1.0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": 1} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": true} is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>contains - 21 failures</summary>

- **contains keyword validation**
  - Test: array with item matching schema (5) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array with item matching schema (6) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array with two items matching schema (5, 6) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema true**
  - Test: any non-empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema true**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array with item 5 is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array with two items 5 is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array without item 5 is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains with false if subschema**
  - Test: any non-empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains with false if subschema**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains with null instance elements**
  - Test: allows null items
  - Expected: `valid`, Got: `error: Unknown type`
- **items + contains**
  - Test: does not match items, matches contains
  - Expected: `invalid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches both items and contains
  - Expected: `valid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches items, does not match contains
  - Expected: `invalid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches neither items nor contains
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>default - 4 failures</summary>

- **invalid string value for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid string value for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>dependentRequired - 20 failures</summary>

- **dependencies with escaped characters**
  - Test: CRLF
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted quotes
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `error: Unknown type`
- **empty dependents**
  - Test: empty object
  - Expected: `valid`, Got: `error: Unknown type`
- **empty dependents**
  - Test: non-object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **empty dependents**
  - Test: object with one property
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: nondependants
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **single dependency**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: nondependant
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: with dependency
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>dependentSchemas - 20 failures</summary>

- **boolean subschemas**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **boolean subschemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **boolean subschemas**
  - Test: object with property having schema false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **boolean subschemas**
  - Test: object with property having schema true is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted tab
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted tab invalid under dependent schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **single dependency**
  - Test: wrong type both
  - Expected: `invalid`, Got: `error: Unknown type`
- **single dependency**
  - Test: wrong type other
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>enum - 20 failures</summary>

- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [0] does not match [false]**
  - Test: [false] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [true] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [0.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [1.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [1] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: extra properties in object is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: objects are deep compared
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: null is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>exclusiveMaximum - 4 failures</summary>

- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: below the exclusiveMaximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>exclusiveMinimum - 4 failures</summary>

- **exclusiveMinimum validation**
  - Test: above the exclusiveMinimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>if-then-else - 18 failures</summary>

- **if and else without then**
  - Test: invalid through else
  - Expected: `invalid`, Got: `error: Unknown type`
- **if and else without then**
  - Test: valid through else
  - Expected: `valid`, Got: `error: Unknown type`
- **if and else without then**
  - Test: valid when if test passes
  - Expected: `valid`, Got: `error: Unknown type`
- **if and then without else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `error: Unknown type`
- **if and then without else**
  - Test: valid through then
  - Expected: `valid`, Got: `error: Unknown type`
- **if and then without else**
  - Test: valid when if test fails
  - Expected: `valid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: invalid redirects to else and fails
  - Expected: `invalid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: no redirects to then and fails
  - Expected: `invalid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: other redirects to else and passes
  - Expected: `valid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: yes redirects to then and passes
  - Expected: `valid`, Got: `error: Unknown type`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `error: Unknown type`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (valid)
  - Expected: `valid`, Got: `error: Unknown type`
- **if with boolean schema true**
  - Test: boolean schema true in if always chooses the then path (invalid)
  - Expected: `invalid`, Got: `error: Unknown type`
- **if with boolean schema true**
  - Test: boolean schema true in if always chooses the then path (valid)
  - Expected: `valid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: invalid through else
  - Expected: `invalid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: valid through else
  - Expected: `valid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: valid through then
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>infinite-loop-detection - 2 failures</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `error: Unknown type`
- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: passing case
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>items - 25 failures</summary>

- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: correct types
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `error: Unknown type`
- **array-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with boolean schema (false)**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schema (true)**
  - Test: any array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schema (true)**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **single-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxContains - 10 failures</summary>

- **maxContains with contains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains with contains**
  - Test: all elements match, valid maxContains
  - Expected: `valid`, Got: `error: Unknown type`
- **maxContains with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains with contains**
  - Test: some elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains with contains**
  - Test: some elements match, valid maxContains
  - Expected: `valid`, Got: `error: Unknown type`
- **maxContains with contains, value with a decimal**
  - Test: one element matches, valid maxContains
  - Expected: `valid`, Got: `error: Unknown type`
- **maxContains with contains, value with a decimal**
  - Test: too many elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains < maxContains**
  - Test: actual < minContains < maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains < maxContains**
  - Test: minContains < actual < maxContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains < maxContains**
  - Test: minContains < maxContains < actual
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxItems - 6 failures</summary>

- **maxItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxItems validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxLength - 7 failures</summary>

- **maxLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxProperties - 10 failures</summary>

- **maxProperties = 0 means the object is empty**
  - Test: no properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxProperties validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maximum - 8 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: below the maximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: below the maximum is invalid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point integer is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>minContains - 26 failures</summary>

- **maxContains < minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains < minContains**
  - Test: invalid maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains < minContains**
  - Test: invalid maxContains and minContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains < minContains**
  - Test: invalid minContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains = minContains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains = minContains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains = minContains**
  - Test: all elements match, valid maxContains and minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **maxContains = minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains = 0 with maxContains**
  - Test: empty data
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains = 0 with maxContains**
  - Test: not more than maxContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains = 0 with maxContains**
  - Test: too many
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains = 0 with no maxContains**
  - Test: empty data
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains = 0 with no maxContains**
  - Test: minContains = 0 makes contains always pass
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=1 with contains**
  - Test: all elements match, valid minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=1 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=1 with contains**
  - Test: no elements match
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=1 with contains**
  - Test: single element matches, valid minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=1 with contains**
  - Test: some elements match, valid minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: all elements match, valid minContains (exactly as needed)
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: all elements match, valid minContains (more than needed)
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: some elements match, invalid minContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: some elements match, valid minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=2 with contains with a decimal value**
  - Test: both elements match, valid minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=2 with contains with a decimal value**
  - Test: one element matches, invalid minContains
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minItems - 6 failures</summary>

- **minItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minItems validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minLength - 7 failures</summary>

- **minLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minLength validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minProperties - 8 failures</summary>

- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minProperties validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minimum - 11 failures</summary>

- **minimum validation**
  - Test: above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point with float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: negative above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: positive above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>multipleOf - 9 failures</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 4.5 is multiple of 1.5
  - Expected: `valid`, Got: `error: Unknown type`
- **by number**
  - Test: zero is multiple of anything
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.0075 is multiple of 0.0001
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `error: Unknown type`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 20 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with empty schema**
  - Test: one valid - valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>pattern - 9 failures</summary>

- **pattern is not anchored**
  - Test: matches a substring
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a matching pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>patternProperties - 23 failures</summary>

- **multiple simultaneous patternProperties are validated**
  - Test: a simultaneous match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: multiple matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple valid matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with a property matching both true and false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema true is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: non recognized members are ignored
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>properties - 28 failures</summary>

- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: both properties present and valid is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: doesn't invalidate other properties
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: both properties present is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: no property present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: only 'false' property present is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: only 'true' property present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with escaped characters**
  - Test: object with all numbers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty validates others
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>propertyNames - 20 failures</summary>

- **propertyNames validation**
  - Test: all property names valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: object without properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: matching property names valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: non-matching property name is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: object without properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema false**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with boolean schema true**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema true**
  - Test: object with any properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: object with property foo is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with property foo and bar is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with property foo is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>required - 16 failures</summary>

- **required default validation**
  - Test: not required by default
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: present required property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **required with empty array**
  - Test: property not required
  - Expected: `valid`, Got: `error: Unknown type`
- **required with escaped characters**
  - Test: object with all properties present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>uniqueItems - 69 failures</summary>

- **uniqueItems validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: different objects are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of strings is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": false} and {"a": 0} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": true} and {"a": 1} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>vocabulary - 3 failures</summary>

- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: applicator vocabulary still works
  - Expected: `invalid`, Got: `error: Unknown type`
- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: no validation: invalid number, but it still validates
  - Expected: `valid`, Got: `error: Unknown type`
- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: no validation: valid number
  - Expected: `valid`, Got: `error: Unknown type`

</details>

## draft2020-12

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ❌ | 0/21 |
| allOf | ⚠️ | 11/30 |
| anchor | ✅ | 0/0 |
| anyOf | ⚠️ | 10/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 33/54 |
| contains | ❌ | 0/21 |
| content | ✅ | 18/18 |
| default | ⚠️ | 3/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ❌ | 0/20 |
| dependentSchemas | ❌ | 0/20 |
| dynamicRef | ✅ | 0/0 |
| enum | ⚠️ | 25/45 |
| exclusiveMaximum | ❌ | 0/4 |
| exclusiveMinimum | ❌ | 0/4 |
| format | ✅ | 133/133 |
| if-then-else | ⚠️ | 8/26 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ⚠️ | 3/29 |
| maxContains | ⚠️ | 2/12 |
| maxItems | ❌ | 0/6 |
| maxLength | ❌ | 0/7 |
| maxProperties | ❌ | 0/10 |
| maximum | ❌ | 0/8 |
| minContains | ⚠️ | 2/28 |
| minItems | ❌ | 0/6 |
| minLength | ❌ | 0/7 |
| minProperties | ❌ | 0/8 |
| minimum | ❌ | 0/11 |
| multipleOf | ⚠️ | 1/10 |
| not | ✅ | 0/0 |
| oneOf | ⚠️ | 7/27 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/23 |
| prefixItems | ❌ | 0/11 |
| properties | ❌ | 0/28 |
| propertyNames | ❌ | 0/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ❌ | 0/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ❌ | 0/69 |
| vocabulary | ⚠️ | 2/5 |

### Failures

<details>
<summary>additionalProperties - 21 failures</summary>

- **additionalProperties are allowed by default**
  - Test: additional properties are allowed
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with propertyNames**
  - Test: Valid against both keywords
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with propertyNames**
  - Test: Valid against propertyNames, but not additionalProperties
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar even when foo2 is present
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties doesn't consider dependentSchemas
  - Expected: `invalid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: matching the pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>allOf - 19 failures</summary>

- **allOf**
  - Test: allOf
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: true
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf simple types**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>anyOf - 8 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf**
  - Test: second anyOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: both anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: first anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: second anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf with base schema**
  - Test: one anyOf valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>const - 21 failures</summary>

- **const with [false] does not match [0]**
  - Test: [0.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [false] does not match [0]**
  - Test: [0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [1.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [1] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with array**
  - Test: another array item is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with array**
  - Test: array with additional items is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with null**
  - Test: not null is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with null**
  - Test: null is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with object**
  - Test: another object is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with object**
  - Test: another type is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": 0.0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": 0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": false} is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": 1.0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": 1} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": true} is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>contains - 21 failures</summary>

- **contains keyword validation**
  - Test: array with item matching schema (5) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array with item matching schema (6) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array with two items matching schema (5, 6) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema true**
  - Test: any non-empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema true**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array with item 5 is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array with two items 5 is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array without item 5 is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains with false if subschema**
  - Test: any non-empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains with false if subschema**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains with null instance elements**
  - Test: allows null items
  - Expected: `valid`, Got: `error: Unknown type`
- **items + contains**
  - Test: does not match items, matches contains
  - Expected: `invalid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches both items and contains
  - Expected: `valid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches items, does not match contains
  - Expected: `invalid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches neither items nor contains
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>default - 4 failures</summary>

- **invalid string value for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid string value for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>dependentRequired - 20 failures</summary>

- **dependencies with escaped characters**
  - Test: CRLF
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted quotes
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `error: Unknown type`
- **empty dependents**
  - Test: empty object
  - Expected: `valid`, Got: `error: Unknown type`
- **empty dependents**
  - Test: non-object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **empty dependents**
  - Test: object with one property
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: nondependants
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependents required**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **single dependency**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: nondependant
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: with dependency
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>dependentSchemas - 20 failures</summary>

- **boolean subschemas**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **boolean subschemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **boolean subschemas**
  - Test: object with property having schema false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **boolean subschemas**
  - Test: object with property having schema true is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted tab
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: quoted tab invalid under dependent schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **single dependency**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **single dependency**
  - Test: wrong type both
  - Expected: `invalid`, Got: `error: Unknown type`
- **single dependency**
  - Test: wrong type other
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>enum - 20 failures</summary>

- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [0] does not match [false]**
  - Test: [false] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [true] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [0.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [1.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [1] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: extra properties in object is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: objects are deep compared
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: null is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>exclusiveMaximum - 4 failures</summary>

- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: below the exclusiveMaximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>exclusiveMinimum - 4 failures</summary>

- **exclusiveMinimum validation**
  - Test: above the exclusiveMinimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>if-then-else - 18 failures</summary>

- **if and else without then**
  - Test: invalid through else
  - Expected: `invalid`, Got: `error: Unknown type`
- **if and else without then**
  - Test: valid through else
  - Expected: `valid`, Got: `error: Unknown type`
- **if and else without then**
  - Test: valid when if test passes
  - Expected: `valid`, Got: `error: Unknown type`
- **if and then without else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `error: Unknown type`
- **if and then without else**
  - Test: valid through then
  - Expected: `valid`, Got: `error: Unknown type`
- **if and then without else**
  - Test: valid when if test fails
  - Expected: `valid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: invalid redirects to else and fails
  - Expected: `invalid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: no redirects to then and fails
  - Expected: `invalid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: other redirects to else and passes
  - Expected: `valid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: yes redirects to then and passes
  - Expected: `valid`, Got: `error: Unknown type`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `error: Unknown type`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (valid)
  - Expected: `valid`, Got: `error: Unknown type`
- **if with boolean schema true**
  - Test: boolean schema true in if always chooses the then path (invalid)
  - Expected: `invalid`, Got: `error: Unknown type`
- **if with boolean schema true**
  - Test: boolean schema true in if always chooses the then path (valid)
  - Expected: `valid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: invalid through else
  - Expected: `invalid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: valid through else
  - Expected: `valid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: valid through then
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>infinite-loop-detection - 2 failures</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `error: Unknown type`
- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: passing case
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>items - 26 failures</summary>

- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `error: Unknown type`
- **items does not look in applicators, valid case**
  - Test: prefixItems in allOf does not constrain items, invalid case
  - Expected: `invalid`, Got: `error: Unknown type`
- **items does not look in applicators, valid case**
  - Test: prefixItems in allOf does not constrain items, valid case
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with boolean schema (false)**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schema (true)**
  - Test: any array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schema (true)**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with heterogeneous array**
  - Test: valid instance
  - Expected: `valid`, Got: `error: Unknown type`
- **items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **prefixItems validation adjusts the starting index for items**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **prefixItems validation adjusts the starting index for items**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `error: Unknown type`
- **prefixItems with no additional items allowed**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `error: Unknown type`
- **prefixItems with no additional items allowed**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **prefixItems with no additional items allowed**
  - Test: equal number of items present
  - Expected: `valid`, Got: `error: Unknown type`
- **prefixItems with no additional items allowed**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `error: Unknown type`
- **prefixItems with no additional items allowed**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxContains - 10 failures</summary>

- **maxContains with contains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains with contains**
  - Test: all elements match, valid maxContains
  - Expected: `valid`, Got: `error: Unknown type`
- **maxContains with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains with contains**
  - Test: some elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains with contains**
  - Test: some elements match, valid maxContains
  - Expected: `valid`, Got: `error: Unknown type`
- **maxContains with contains, value with a decimal**
  - Test: one element matches, valid maxContains
  - Expected: `valid`, Got: `error: Unknown type`
- **maxContains with contains, value with a decimal**
  - Test: too many elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains < maxContains**
  - Test: actual < minContains < maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains < maxContains**
  - Test: minContains < actual < maxContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains < maxContains**
  - Test: minContains < maxContains < actual
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxItems - 6 failures</summary>

- **maxItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxItems validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxLength - 7 failures</summary>

- **maxLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxProperties - 10 failures</summary>

- **maxProperties = 0 means the object is empty**
  - Test: no properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxProperties validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maximum - 8 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: below the maximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: below the maximum is invalid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point integer is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>minContains - 26 failures</summary>

- **maxContains < minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains < minContains**
  - Test: invalid maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains < minContains**
  - Test: invalid maxContains and minContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains < minContains**
  - Test: invalid minContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains = minContains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains = minContains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxContains = minContains**
  - Test: all elements match, valid maxContains and minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **maxContains = minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains = 0**
  - Test: empty data
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains = 0**
  - Test: minContains = 0 makes contains always pass
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains = 0 with maxContains**
  - Test: empty data
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains = 0 with maxContains**
  - Test: not more than maxContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains = 0 with maxContains**
  - Test: too many
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=1 with contains**
  - Test: all elements match, valid minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=1 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=1 with contains**
  - Test: no elements match
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=1 with contains**
  - Test: single element matches, valid minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=1 with contains**
  - Test: some elements match, valid minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: all elements match, valid minContains (exactly as needed)
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: all elements match, valid minContains (more than needed)
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: some elements match, invalid minContains
  - Expected: `invalid`, Got: `error: Unknown type`
- **minContains=2 with contains**
  - Test: some elements match, valid minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=2 with contains with a decimal value**
  - Test: both elements match, valid minContains
  - Expected: `valid`, Got: `error: Unknown type`
- **minContains=2 with contains with a decimal value**
  - Test: one element matches, invalid minContains
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minItems - 6 failures</summary>

- **minItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minItems validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minLength - 7 failures</summary>

- **minLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minLength validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minProperties - 8 failures</summary>

- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minProperties validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minimum - 11 failures</summary>

- **minimum validation**
  - Test: above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point with float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: negative above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: positive above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>multipleOf - 9 failures</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 4.5 is multiple of 1.5
  - Expected: `valid`, Got: `error: Unknown type`
- **by number**
  - Test: zero is multiple of anything
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.0075 is multiple of 0.0001
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `error: Unknown type`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 20 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with empty schema**
  - Test: one valid - valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>pattern - 9 failures</summary>

- **pattern is not anchored**
  - Test: matches a substring
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a matching pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>patternProperties - 23 failures</summary>

- **multiple simultaneous patternProperties are validated**
  - Test: a simultaneous match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: multiple matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple valid matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with a property matching both true and false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema true is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: non recognized members are ignored
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>prefixItems - 11 failures</summary>

- **a schema given for prefixItems**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for prefixItems**
  - Test: array with additional items
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for prefixItems**
  - Test: correct types
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for prefixItems**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for prefixItems**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for prefixItems**
  - Test: wrong types
  - Expected: `invalid`, Got: `error: Unknown type`
- **additional items are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `error: Unknown type`
- **prefixItems with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **prefixItems with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **prefixItems with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **prefixItems with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>properties - 28 failures</summary>

- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: both properties present and valid is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: doesn't invalidate other properties
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: both properties present is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: no property present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: only 'false' property present is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: only 'true' property present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with escaped characters**
  - Test: object with all numbers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty validates others
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>propertyNames - 20 failures</summary>

- **propertyNames validation**
  - Test: all property names valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: object without properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: matching property names valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: non-matching property name is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: object without properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema false**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with boolean schema true**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema true**
  - Test: object with any properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: object with property foo is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with property foo and bar is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with property foo is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>required - 16 failures</summary>

- **required default validation**
  - Test: not required by default
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: present required property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **required with empty array**
  - Test: property not required
  - Expected: `valid`, Got: `error: Unknown type`
- **required with escaped characters**
  - Test: object with all properties present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>uniqueItems - 69 failures</summary>

- **uniqueItems validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: different objects are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of strings is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": false} and {"a": 0} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": true} and {"a": 1} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>vocabulary - 3 failures</summary>

- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: applicator vocabulary still works
  - Expected: `invalid`, Got: `error: Unknown type`
- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: no validation: invalid number, but it still validates
  - Expected: `valid`, Got: `error: Unknown type`
- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: no validation: valid number
  - Expected: `valid`, Got: `error: Unknown type`

</details>

## draft3

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ❌ | 0/14 |
| additionalProperties | ❌ | 0/16 |
| default | ⚠️ | 3/7 |
| dependencies | ❌ | 0/18 |
| disallow | ✅ | 9/9 |
| divisibleBy | ❌ | 0/8 |
| enum | ⚠️ | 10/16 |
| extends | ❌ | 0/10 |
| format | ✅ | 60/60 |
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
| ref | ✅ | 0/0 |
| refRemote | ⚠️ | 6/8 |
| required | ❌ | 0/4 |
| type | ⚠️ | 73/80 |
| uniqueItems | ❌ | 0/62 |

### Failures

<details>
<summary>additionalItems - 14 failures</summary>

- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as false without items**
  - Test: items defaults to empty schema so everything is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems with heterogeneous array**
  - Test: valid instance
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: equal number of items present
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `error: Unknown type`
- **when items is schema, additionalItems does nothing**
  - Test: all items match schema
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>additionalProperties - 16 failures</summary>

- **additionalProperties are allowed by default**
  - Test: additional properties are allowed
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties does not look in applicators**
  - Test: properties defined in extends are not examined
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: matching the pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>default - 4 failures</summary>

- **invalid string value for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid string value for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>dependencies - 18 failures</summary>

- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: nondependant
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: with dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: nondependants
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type both
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type other
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>divisibleBy - 8 failures</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 35 is not divisible by 1.5
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 4.5 is divisible by 1.5
  - Expected: `valid`, Got: `error: Unknown type`
- **by number**
  - Test: zero is divisible by anything (except 0)
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.0075 is divisible by 0.0001
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.00751 is not divisible by 0.0001
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>enum - 6 failures</summary>

- **heterogeneous enum validation**
  - Test: objects are deep compared
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: null is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>extends - 10 failures</summary>

- **extends**
  - Test: extends
  - Expected: `valid`, Got: `error: Unknown type`
- **extends**
  - Test: mismatch extended
  - Expected: `invalid`, Got: `error: Unknown type`
- **extends**
  - Test: mismatch extends
  - Expected: `invalid`, Got: `error: Unknown type`
- **extends**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **extends simple types**
  - Test: mismatch extends
  - Expected: `invalid`, Got: `error: Unknown type`
- **extends simple types**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple extends**
  - Test: mismatch both
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple extends**
  - Test: mismatch first extends
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple extends**
  - Test: mismatch second extends
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple extends**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>infinite-loop-detection - 2 failures</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `error: Unknown type`
- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: passing case
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>items - 7 failures</summary>

- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: correct types
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `error: Unknown type`
- **array-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxItems - 4 failures</summary>

- **maxItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxLength - 5 failures</summary>

- **maxLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>maximum - 14 failures</summary>

- **exclusiveMaximum validation**
  - Test: below the maximum is still valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: below the maximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation (explicit false exclusivity)**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation (explicit false exclusivity)**
  - Test: below the maximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation (explicit false exclusivity)**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: below the maximum is invalid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point integer is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>minItems - 4 failures</summary>

- **minItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minLength - 5 failures</summary>

- **minLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minimum - 13 failures</summary>

- **exclusiveMinimum validation**
  - Test: above the minimum is still valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point with float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: negative above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: positive above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>pattern - 9 failures</summary>

- **pattern is not anchored**
  - Test: matches a substring
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a matching pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>patternProperties - 17 failures</summary>

- **multiple simultaneous patternProperties are validated**
  - Test: a simultaneous match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: multiple matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple valid matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: non recognized members are ignored
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>properties - 15 failures</summary>

- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: both properties present and valid is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: doesn't invalidate other properties
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty validates others
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>refRemote - 2 failures</summary>

- **change resolution scope**
  - Test: changed scope ref invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **change resolution scope**
  - Test: changed scope ref valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>required - 4 failures</summary>

- **required default validation**
  - Test: not required by default
  - Expected: `valid`, Got: `error: Unknown type`
- **required explicitly false validation**
  - Test: not required if required is false
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: present required property is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>type - 7 failures</summary>

- **applies a nested schema**
  - Test: an object is invalid otherwise
  - Expected: `invalid`, Got: `true`
- **types can include schemas**
  - Test: a boolean is invalid
  - Expected: `invalid`, Got: `true`
- **types can include schemas**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`
- **types can include schemas**
  - Test: a string is invalid
  - Expected: `invalid`, Got: `true`
- **types can include schemas**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `true`
- **types can include schemas**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **types from separate schemas are merged**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 62 failures</summary>

- **uniqueItems validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of strings is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": false} and {"a": 0} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": true} and {"a": 1} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

## draft4

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ❌ | 0/17 |
| additionalProperties | ❌ | 0/16 |
| allOf | ⚠️ | 8/27 |
| anyOf | ⚠️ | 7/15 |
| default | ⚠️ | 3/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ❌ | 0/29 |
| enum | ⚠️ | 29/49 |
| format | ✅ | 36/36 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ⚠️ | 3/21 |
| maxItems | ❌ | 0/4 |
| maxLength | ❌ | 0/5 |
| maxProperties | ❌ | 0/8 |
| maximum | ❌ | 0/14 |
| minItems | ❌ | 0/4 |
| minLength | ❌ | 0/5 |
| minProperties | ❌ | 0/6 |
| minimum | ❌ | 0/17 |
| multipleOf | ⚠️ | 1/10 |
| not | ⚠️ | 18/20 |
| oneOf | ⚠️ | 3/23 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/18 |
| properties | ❌ | 0/24 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ❌ | 0/15 |
| type | ✅ | 79/79 |
| uniqueItems | ❌ | 0/69 |

### Failures

<details>
<summary>additionalItems - 17 failures</summary>

- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as false without items**
  - Test: items defaults to empty schema so everything is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems with heterogeneous array**
  - Test: valid instance
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: equal number of items present
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `error: Unknown type`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `error: Unknown type`
- **when items is schema, additionalItems does nothing**
  - Test: all items match schema
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>additionalProperties - 16 failures</summary>

- **additionalProperties are allowed by default**
  - Test: additional properties are allowed
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: matching the pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>allOf - 19 failures</summary>

- **allOf**
  - Test: allOf
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: true
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf simple types**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>anyOf - 8 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf**
  - Test: second anyOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: both anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: first anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: second anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf with base schema**
  - Test: one anyOf valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>default - 4 failures</summary>

- **invalid string value for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid string value for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 29 failures</summary>

- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: nondependant
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: with dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 1
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 2
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 4
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: valid object 1
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: valid object 2
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: valid object 3
  - Expected: `valid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: nondependants
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type both
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type other
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>enum - 20 failures</summary>

- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [0] does not match [false]**
  - Test: [false] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [true] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [0.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [1.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [1] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: extra properties in object is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: objects are deep compared
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: null is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>infinite-loop-detection - 2 failures</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `error: Unknown type`
- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: passing case
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>items - 18 failures</summary>

- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: correct types
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `error: Unknown type`
- **array-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxItems - 4 failures</summary>

- **maxItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxLength - 5 failures</summary>

- **maxLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxProperties - 8 failures</summary>

- **maxProperties = 0 means the object is empty**
  - Test: no properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maximum - 14 failures</summary>

- **exclusiveMaximum validation**
  - Test: below the maximum is still valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: below the maximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation (explicit false exclusivity)**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation (explicit false exclusivity)**
  - Test: below the maximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation (explicit false exclusivity)**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: below the maximum is invalid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point integer is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>minItems - 4 failures</summary>

- **minItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minLength - 5 failures</summary>

- **minLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minProperties - 6 failures</summary>

- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minimum - 17 failures</summary>

- **exclusiveMinimum validation**
  - Test: above the minimum is still valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation (explicit false exclusivity)**
  - Test: above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation (explicit false exclusivity)**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation (explicit false exclusivity)**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point with float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: negative above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: positive above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>multipleOf - 9 failures</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 4.5 is multiple of 1.5
  - Expected: `valid`, Got: `error: Unknown type`
- **by number**
  - Test: zero is multiple of anything
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.0075 is multiple of 0.0001
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `error: Unknown type`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>not - 2 failures</summary>

- **forbidden property**
  - Test: property absent
  - Expected: `valid`, Got: `error: Unknown type`
- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>oneOf - 20 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with empty schema**
  - Test: one valid - valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>pattern - 9 failures</summary>

- **pattern is not anchored**
  - Test: matches a substring
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a matching pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>patternProperties - 18 failures</summary>

- **multiple simultaneous patternProperties are validated**
  - Test: a simultaneous match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: multiple matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple valid matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: non recognized members are ignored
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>properties - 24 failures</summary>

- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: both properties present and valid is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: doesn't invalidate other properties
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with escaped characters**
  - Test: object with all numbers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty validates others
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>required - 15 failures</summary>

- **required default validation**
  - Test: not required by default
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: present required property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **required with escaped characters**
  - Test: object with all properties present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>uniqueItems - 69 failures</summary>

- **uniqueItems validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: different objects are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of strings is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": false} and {"a": 0} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": true} and {"a": 1} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

## draft6

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ❌ | 0/19 |
| additionalProperties | ❌ | 0/16 |
| allOf | ⚠️ | 11/30 |
| anyOf | ⚠️ | 10/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 33/54 |
| contains | ❌ | 0/19 |
| default | ⚠️ | 3/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ❌ | 0/36 |
| enum | ⚠️ | 25/45 |
| exclusiveMaximum | ❌ | 0/4 |
| exclusiveMinimum | ❌ | 0/4 |
| format | ✅ | 54/54 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ⚠️ | 3/28 |
| maxItems | ❌ | 0/6 |
| maxLength | ❌ | 0/7 |
| maxProperties | ❌ | 0/10 |
| maximum | ❌ | 0/8 |
| minItems | ❌ | 0/6 |
| minLength | ❌ | 0/7 |
| minProperties | ❌ | 0/8 |
| minimum | ❌ | 0/11 |
| multipleOf | ⚠️ | 1/10 |
| not | ⚠️ | 36/38 |
| oneOf | ⚠️ | 7/27 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/23 |
| properties | ❌ | 0/28 |
| propertyNames | ❌ | 0/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ❌ | 0/16 |
| type | ✅ | 80/80 |
| uniqueItems | ❌ | 0/69 |

### Failures

<details>
<summary>additionalItems - 19 failures</summary>

- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as false without items**
  - Test: items defaults to empty schema so everything is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems with heterogeneous array**
  - Test: valid instance
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: equal number of items present
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `error: Unknown type`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `error: Unknown type`
- **when items is schema, additionalItems does nothing**
  - Test: invalid with a array of mixed types
  - Expected: `invalid`, Got: `error: Unknown type`
- **when items is schema, additionalItems does nothing**
  - Test: valid with a array of type integers
  - Expected: `valid`, Got: `error: Unknown type`
- **when items is schema, boolean additionalItems does nothing**
  - Test: all items match schema
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>additionalProperties - 16 failures</summary>

- **additionalProperties are allowed by default**
  - Test: additional properties are allowed
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: matching the pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>allOf - 19 failures</summary>

- **allOf**
  - Test: allOf
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: true
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf simple types**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>anyOf - 8 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf**
  - Test: second anyOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: both anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: first anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: second anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf with base schema**
  - Test: one anyOf valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>const - 21 failures</summary>

- **const with [false] does not match [0]**
  - Test: [0.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [false] does not match [0]**
  - Test: [0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [1.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [1] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with array**
  - Test: another array item is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with array**
  - Test: array with additional items is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with null**
  - Test: not null is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with null**
  - Test: null is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with object**
  - Test: another object is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with object**
  - Test: another type is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": 0.0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": 0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": false} is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": 1.0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": 1} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": true} is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>contains - 19 failures</summary>

- **contains keyword validation**
  - Test: array with item matching schema (5) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array with item matching schema (6) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array with two items matching schema (5, 6) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema true**
  - Test: any non-empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema true**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array with item 5 is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array with two items 5 is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array without item 5 is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains with null instance elements**
  - Test: allows null items
  - Expected: `valid`, Got: `error: Unknown type`
- **items + contains**
  - Test: does not match items, matches contains
  - Expected: `invalid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches both items and contains
  - Expected: `valid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches items, does not match contains
  - Expected: `invalid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches neither items nor contains
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>default - 4 failures</summary>

- **invalid string value for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid string value for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 36 failures</summary>

- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: nondependant
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: with dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with boolean subschemas**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with boolean subschemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with boolean subschemas**
  - Test: object with property having schema false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with boolean subschemas**
  - Test: object with property having schema true is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with empty array**
  - Test: empty object
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with empty array**
  - Test: non-object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with empty array**
  - Test: object with one property
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 1
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 2
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 4
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: valid object 1
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: valid object 2
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: valid object 3
  - Expected: `valid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: nondependants
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type both
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type other
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>enum - 20 failures</summary>

- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [0] does not match [false]**
  - Test: [false] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [true] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [0.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [1.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [1] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: extra properties in object is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: objects are deep compared
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: null is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>exclusiveMaximum - 4 failures</summary>

- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: below the exclusiveMaximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>exclusiveMinimum - 4 failures</summary>

- **exclusiveMinimum validation**
  - Test: above the exclusiveMinimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>infinite-loop-detection - 2 failures</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `error: Unknown type`
- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: passing case
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>items - 25 failures</summary>

- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: correct types
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `error: Unknown type`
- **array-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with boolean schema (false)**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schema (true)**
  - Test: any array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schema (true)**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **single-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxItems - 6 failures</summary>

- **maxItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxItems validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxLength - 7 failures</summary>

- **maxLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxProperties - 10 failures</summary>

- **maxProperties = 0 means the object is empty**
  - Test: no properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxProperties validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maximum - 8 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: below the maximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: below the maximum is invalid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point integer is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>minItems - 6 failures</summary>

- **minItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minItems validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minLength - 7 failures</summary>

- **minLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minLength validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minProperties - 8 failures</summary>

- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minProperties validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minimum - 11 failures</summary>

- **minimum validation**
  - Test: above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point with float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: negative above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: positive above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>multipleOf - 9 failures</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 4.5 is multiple of 1.5
  - Expected: `valid`, Got: `error: Unknown type`
- **by number**
  - Test: zero is multiple of anything
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.0075 is multiple of 0.0001
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `error: Unknown type`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>not - 2 failures</summary>

- **forbidden property**
  - Test: property absent
  - Expected: `valid`, Got: `error: Unknown type`
- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>oneOf - 20 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with empty schema**
  - Test: one valid - valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>pattern - 9 failures</summary>

- **pattern is not anchored**
  - Test: matches a substring
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a matching pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>patternProperties - 23 failures</summary>

- **multiple simultaneous patternProperties are validated**
  - Test: a simultaneous match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: multiple matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple valid matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with a property matching both true and false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema true is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: non recognized members are ignored
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>properties - 28 failures</summary>

- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: both properties present and valid is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: doesn't invalidate other properties
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: both properties present is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: no property present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: only 'false' property present is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: only 'true' property present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with escaped characters**
  - Test: object with all numbers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty validates others
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>propertyNames - 20 failures</summary>

- **propertyNames validation**
  - Test: all property names valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: object without properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: matching property names valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: non-matching property name is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: object without properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema false**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with boolean schema true**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema true**
  - Test: object with any properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: object with property foo is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with property foo and bar is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with property foo is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>required - 16 failures</summary>

- **required default validation**
  - Test: not required by default
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: present required property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **required with empty array**
  - Test: property not required
  - Expected: `valid`, Got: `error: Unknown type`
- **required with escaped characters**
  - Test: object with all properties present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>uniqueItems - 69 failures</summary>

- **uniqueItems validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: different objects are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of strings is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": false} and {"a": 0} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": true} and {"a": 1} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

## draft7

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ❌ | 0/19 |
| additionalProperties | ❌ | 0/16 |
| allOf | ⚠️ | 11/30 |
| anyOf | ⚠️ | 10/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 33/54 |
| contains | ❌ | 0/21 |
| default | ⚠️ | 3/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ❌ | 0/36 |
| enum | ⚠️ | 25/45 |
| exclusiveMaximum | ❌ | 0/4 |
| exclusiveMinimum | ❌ | 0/4 |
| format | ✅ | 102/102 |
| if-then-else | ⚠️ | 8/26 |
| infinite-loop-detection | ❌ | 0/2 |
| items | ⚠️ | 3/28 |
| maxItems | ❌ | 0/6 |
| maxLength | ❌ | 0/7 |
| maxProperties | ❌ | 0/10 |
| maximum | ❌ | 0/8 |
| minItems | ❌ | 0/6 |
| minLength | ❌ | 0/7 |
| minProperties | ❌ | 0/8 |
| minimum | ❌ | 0/11 |
| multipleOf | ⚠️ | 1/10 |
| not | ⚠️ | 36/38 |
| oneOf | ⚠️ | 7/27 |
| pattern | ❌ | 0/9 |
| patternProperties | ❌ | 0/23 |
| properties | ❌ | 0/28 |
| propertyNames | ❌ | 0/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ❌ | 0/16 |
| type | ✅ | 80/80 |
| uniqueItems | ❌ | 0/69 |

### Failures

<details>
<summary>additionalItems - 19 failures</summary>

- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as false without items**
  - Test: items defaults to empty schema so everything is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalItems with heterogeneous array**
  - Test: valid instance
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalItems with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: equal number of items present
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `error: Unknown type`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `error: Unknown type`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `error: Unknown type`
- **when items is schema, additionalItems does nothing**
  - Test: invalid with a array of mixed types
  - Expected: `invalid`, Got: `error: Unknown type`
- **when items is schema, additionalItems does nothing**
  - Test: valid with a array of type integers
  - Expected: `valid`, Got: `error: Unknown type`
- **when items is schema, boolean additionalItems does nothing**
  - Test: all items match schema
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>additionalProperties - 16 failures</summary>

- **additionalProperties are allowed by default**
  - Test: additional properties are allowed
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties can exist by itself**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional invalid property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: matching the pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **non-ASCII pattern with additionalProperties**
  - Test: not matching the pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>allOf - 19 failures</summary>

- **allOf**
  - Test: allOf
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf**
  - Test: mismatch first
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf**
  - Test: mismatch second
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: false, anyOf: true, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: false, oneOf: true
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: false
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf combined with anyOf, oneOf**
  - Test: allOf: true, anyOf: true, oneOf: true
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf simple types**
  - Test: mismatch one
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf simple types**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch both
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch first allOf
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: mismatch second allOf
  - Expected: `invalid`, Got: `error: Unknown type`
- **allOf with base schema**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>anyOf - 8 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf**
  - Test: second anyOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: both anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: first anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf complex types**
  - Test: second anyOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **anyOf with base schema**
  - Test: one anyOf valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>const - 21 failures</summary>

- **const with [false] does not match [0]**
  - Test: [0.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [false] does not match [0]**
  - Test: [0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [1.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [1] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with array**
  - Test: another array item is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with array**
  - Test: array with additional items is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with null**
  - Test: not null is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with null**
  - Test: null is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with object**
  - Test: another object is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with object**
  - Test: another type is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": 0.0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": 0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": false} does not match {"a": 0}**
  - Test: {"a": false} is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": 1.0} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": 1} is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **const with {"a": true} does not match {"a": 1}**
  - Test: {"a": true} is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>contains - 21 failures</summary>

- **contains keyword validation**
  - Test: array with item matching schema (5) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array with item matching schema (6) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array with two items matching schema (5, 6) is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema true**
  - Test: any non-empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with boolean schema true**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array with item 5 is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array with two items 5 is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains keyword with const keyword**
  - Test: array without item 5 is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains with false if subschema**
  - Test: any non-empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **contains with false if subschema**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **contains with null instance elements**
  - Test: allows null items
  - Expected: `valid`, Got: `error: Unknown type`
- **items + contains**
  - Test: does not match items, matches contains
  - Expected: `invalid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches both items and contains
  - Expected: `valid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches items, does not match contains
  - Expected: `invalid`, Got: `error: Unknown type`
- **items + contains**
  - Test: matches neither items nor contains
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>default - 4 failures</summary>

- **invalid string value for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid string value for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: still valid when the invalid default is used
  - Expected: `valid`, Got: `error: Unknown type`
- **invalid type for default**
  - Test: valid when property is specified
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 36 failures</summary>

- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: nondependant
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies**
  - Test: with dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with boolean subschemas**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with boolean subschemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with boolean subschemas**
  - Test: object with property having schema false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with boolean subschemas**
  - Test: object with property having schema true is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with empty array**
  - Test: empty object
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with empty array**
  - Test: non-object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with empty array**
  - Test: object with one property
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 1
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 2
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: invalid object 4
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: valid object 1
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: valid object 2
  - Expected: `valid`, Got: `error: Unknown type`
- **dependencies with escaped characters**
  - Test: valid object 3
  - Expected: `valid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `error: Unknown type`
- **dependent subschema incompatible with root**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: nondependants
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: no dependency
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type both
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple dependencies subschema**
  - Test: wrong type other
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>enum - 20 failures</summary>

- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [0] does not match [false]**
  - Test: [false] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [1] does not match [true]**
  - Test: [true] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [0.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [1.0] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [1] is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: extra properties in object is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: objects are deep compared
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: null is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **heterogeneous enum-with-null validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>exclusiveMaximum - 4 failures</summary>

- **exclusiveMaximum validation**
  - Test: above the exclusiveMaximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: below the exclusiveMaximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>exclusiveMinimum - 4 failures</summary>

- **exclusiveMinimum validation**
  - Test: above the exclusiveMinimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: below the exclusiveMinimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>if-then-else - 18 failures</summary>

- **if and else without then**
  - Test: invalid through else
  - Expected: `invalid`, Got: `error: Unknown type`
- **if and else without then**
  - Test: valid through else
  - Expected: `valid`, Got: `error: Unknown type`
- **if and else without then**
  - Test: valid when if test passes
  - Expected: `valid`, Got: `error: Unknown type`
- **if and then without else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `error: Unknown type`
- **if and then without else**
  - Test: valid through then
  - Expected: `valid`, Got: `error: Unknown type`
- **if and then without else**
  - Test: valid when if test fails
  - Expected: `valid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: invalid redirects to else and fails
  - Expected: `invalid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: no redirects to then and fails
  - Expected: `invalid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: other redirects to else and passes
  - Expected: `valid`, Got: `error: Unknown type`
- **if appears at the end when serialized (keyword processing sequence)**
  - Test: yes redirects to then and passes
  - Expected: `valid`, Got: `error: Unknown type`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (invalid)
  - Expected: `invalid`, Got: `error: Unknown type`
- **if with boolean schema false**
  - Test: boolean schema false in if always chooses the else path (valid)
  - Expected: `valid`, Got: `error: Unknown type`
- **if with boolean schema true**
  - Test: boolean schema true in if always chooses the then path (invalid)
  - Expected: `invalid`, Got: `error: Unknown type`
- **if with boolean schema true**
  - Test: boolean schema true in if always chooses the then path (valid)
  - Expected: `valid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: invalid through else
  - Expected: `invalid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: invalid through then
  - Expected: `invalid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: valid through else
  - Expected: `valid`, Got: `error: Unknown type`
- **validate against correct branch, then vs else**
  - Test: valid through then
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>infinite-loop-detection - 2 failures</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `error: Unknown type`
- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: passing case
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>items - 25 failures</summary>

- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **a schema given for items**
  - Test: wrong type of items
  - Expected: `invalid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: correct types
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `error: Unknown type`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `error: Unknown type`
- **array-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: valid items
  - Expected: `valid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `error: Unknown type`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with boolean schema (false)**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schema (true)**
  - Test: any array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schema (true)**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **items with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **single-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxItems - 6 failures</summary>

- **maxItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxItems validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxLength - 7 failures</summary>

- **maxLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxLength validation**
  - Test: two graphemes is long enough
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxLength validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maxProperties - 10 failures</summary>

- **maxProperties = 0 means the object is empty**
  - Test: no properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maxProperties validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>maximum - 8 failures</summary>

- **maximum validation**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: below the maximum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: below the maximum is invalid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **maximum validation with unsigned integer**
  - Test: boundary point integer is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>minItems - 6 failures</summary>

- **minItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minItems validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minLength - 7 failures</summary>

- **minLength validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: one grapheme is not long enough
  - Expected: `invalid`, Got: `error: Unknown type`
- **minLength validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minLength validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minLength validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minProperties - 8 failures</summary>

- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minProperties validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>minimum - 11 failures</summary>

- **minimum validation**
  - Test: above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: boundary point with float is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: float below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: int below the minimum is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: negative above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **minimum validation with signed integer**
  - Test: positive above the minimum is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>multipleOf - 9 failures</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int
  - Expected: `valid`, Got: `error: Unknown type`
- **by int**
  - Test: int by int fail
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 35 is not multiple of 1.5
  - Expected: `invalid`, Got: `error: Unknown type`
- **by number**
  - Test: 4.5 is multiple of 1.5
  - Expected: `valid`, Got: `error: Unknown type`
- **by number**
  - Test: zero is multiple of anything
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.0075 is multiple of 0.0001
  - Expected: `valid`, Got: `error: Unknown type`
- **by small number**
  - Test: 0.00751 is not multiple of 0.0001
  - Expected: `invalid`, Got: `error: Unknown type`
- **small multiple of large integer**
  - Test: any integer is a multiple of 1e-8
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>not - 2 failures</summary>

- **forbidden property**
  - Test: property absent
  - Expected: `valid`, Got: `error: Unknown type`
- **forbidden property**
  - Test: property present
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>oneOf - 20 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: both oneOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: first oneOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: neither oneOf valid (complex)
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf complex types**
  - Test: second oneOf valid (complex)
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with base schema**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with base schema**
  - Test: one oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with empty schema**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with empty schema**
  - Test: one valid - valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: both invalid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: both valid - invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `error: Unknown type`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>pattern - 9 failures</summary>

- **pattern is not anchored**
  - Test: matches a substring
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a matching pattern is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: a non-matching pattern is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `error: Unknown type`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>patternProperties - 23 failures</summary>

- **multiple simultaneous patternProperties are validated**
  - Test: a simultaneous match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to both is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to one is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: an invalid due to the other is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **multiple simultaneous patternProperties are validated**
  - Test: multiple matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: a single valid match is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties validates properties matching a regex**
  - Test: multiple valid matches is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with a property matching both true and false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema false is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **patternProperties with boolean schemas**
  - Test: object with property matching schema true is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **patternProperties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: non recognized members are ignored
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: recognized members are accounted for
  - Expected: `invalid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive
  - Expected: `valid`, Got: `error: Unknown type`
- **regexes are not anchored by default and are case sensitive**
  - Test: regexes are case sensitive, 2
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>properties - 28 failures</summary>

- **object properties validation**
  - Test: both properties invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: both properties present and valid is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: doesn't invalidate other properties
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **object properties validation**
  - Test: one property invalid is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: Unknown type`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: both properties present is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: no property present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: only 'false' property present is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with boolean schema**
  - Test: only 'true' property present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with escaped characters**
  - Test: object with all numbers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **properties with escaped characters**
  - Test: object with strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties with null valued instance properties**
  - Test: allows null values
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty validates others
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `error: Unknown type`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>propertyNames - 20 failures</summary>

- **propertyNames validation**
  - Test: all property names valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: object without properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: matching property names valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: non-matching property name is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames validation with pattern**
  - Test: object without properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema false**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema false**
  - Test: object with any properties is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with boolean schema true**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with boolean schema true**
  - Test: object with any properties is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with const**
  - Test: object with property foo is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with any other property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with property foo and bar is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **propertyNames with enum**
  - Test: object with property foo is valid
  - Expected: `valid`, Got: `error: Unknown type`

</details>

<details>
<summary>required - 16 failures</summary>

- **required default validation**
  - Test: not required by default
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: Unknown type`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: Unknown type`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **required validation**
  - Test: present required property is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **required with empty array**
  - Test: property not required
  - Expected: `valid`, Got: `error: Unknown type`
- **required with escaped characters**
  - Test: object with all properties present is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

<details>
<summary>uniqueItems - 69 failures</summary>

- **uniqueItems validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: different objects are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [0] and [false] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: nested [1] and [true] are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique array of strings is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": false} and {"a": 0} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems validation**
  - Test: {"a": true} and {"a": 1} are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: non-unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: Unknown type`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: Unknown type`

</details>

