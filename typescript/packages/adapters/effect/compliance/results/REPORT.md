# @xschemadev/effect Compliance Report

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2019-09 | 821 | 38 | 0 | 95.6% |
| draft2020-12 | 833 | 38 | 0 | 95.6% |
| draft3 | 383 | 24 | 0 | 94.1% |
| draft4 | 515 | 36 | 0 | 93.5% |
| draft6 | 704 | 35 | 0 | 95.3% |
| draft7 | 780 | 35 | 0 | 95.7% |

## Badges

![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-95.6%25-brightgreen)
![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-95.6%25-brightgreen)
![draft3](https://img.shields.io/badge/draft3%20compliance-94.1%25-yellow)
![draft4](https://img.shields.io/badge/draft4%20compliance-93.5%25-yellow)
![draft6](https://img.shields.io/badge/draft6%20compliance-95.3%25-brightgreen)
![draft7](https://img.shields.io/badge/draft7%20compliance-95.7%25-brightgreen)

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
| if-then-else | ✅ | 26/26 |
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
| not | ✅ | 0/0 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ⚠️ | 24/28 |
| propertyNames | ✅ | 20/20 |
| recursiveRef | ✅ | 0/0 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 10/16 |
| type | ⚠️ | 71/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 7 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar even when foo2 is present
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties doesn't consider dependentSchemas
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
  - Test: matches both
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
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

- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 4 failures</summary>

- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>required - 6 failures</summary>

- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: a boolean is not an object
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
  - Test: an integer is not an object
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
<summary>vocabulary - 1 failure</summary>

- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: no validation: invalid number, but it still validates
  - Expected: `valid`, Got: `false`

</details>

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
| dynamicRef | ✅ | 0/0 |
| enum | ⚠️ | 43/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 133/133 |
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
| not | ✅ | 0/0 |
| oneOf | ⚠️ | 23/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| prefixItems | ✅ | 11/11 |
| properties | ⚠️ | 24/28 |
| propertyNames | ✅ | 20/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 10/16 |
| type | ⚠️ | 71/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 7 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar even when foo2 is present
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties doesn't consider dependentSchemas
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
  - Test: matches both
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
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

- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 4 failures</summary>

- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>required - 6 failures</summary>

- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: a boolean is not an object
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
  - Test: an integer is not an object
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
<summary>vocabulary - 1 failure</summary>

- **schema that uses custom metaschema with with no validation vocabulary**
  - Test: no validation: invalid number, but it still validates
  - Expected: `valid`, Got: `false`

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
| ref | ✅ | 0/0 |
| refRemote | ✅ | 8/8 |
| required | ⚠️ | 3/4 |
| type | ⚠️ | 68/80 |
| uniqueItems | ✅ | 62/62 |

### Failures

<details>
<summary>additionalProperties - 4 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in extends are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`

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
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>required - 1 failure</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 12 failures</summary>

- **applies a nested schema**
  - Test: an object is invalid otherwise
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: a boolean is not an object
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
  - Test: an integer is not an object
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
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 9/15 |
| type | ⚠️ | 70/79 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 4 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`

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
  - Test: matches both
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
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

- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 4 failures</summary>

- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>required - 6 failures</summary>

- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: a boolean is not an object
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
  - Test: an integer is not an object
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
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 10/16 |
| type | ⚠️ | 71/80 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 4 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>definitions - 2 failures</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | 
2 | import { Schema as S } from "effect"
3 | 
4 | const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
5 |   "group_0": (() => {
6 |     const schema = S.Union(S.Struct({ "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at <anonymous> (/home/trapani/dev/xschema/typescript/packages/adapters/effect/xschema-harness-1115516538.ts:6:73)
      at /home/trapani/dev/xschema/typescript/packages/adapters/effect/xschema-harness-1115516538.ts:39:115
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | 
2 | import { Schema as S } from "effect"
3 | 
4 | const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
5 |   "group_0": (() => {
6 |     const schema = S.Union(S.Struct({ "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at <anonymous> (/home/trapani/dev/xschema/typescript/packages/adapters/effect/xschema-harness-1115516538.ts:6:73)
      at /home/trapani/dev/xschema/typescript/packages/adapters/effect/xschema-harness-1115516538.ts:39:115
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
  - Test: matches both
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
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

- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 4 failures</summary>

- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>required - 6 failures</summary>

- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: a boolean is not an object
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
  - Test: an integer is not an object
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
| if-then-else | ✅ | 26/26 |
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
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 10/16 |
| type | ⚠️ | 71/80 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 4 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **additionalProperties with schema**
  - Test: an additional valid property is valid
  - Expected: `valid`, Got: `false`
- **additionalProperties with schema**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>definitions - 2 failures</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | 
2 | import { Schema as S } from "effect"
3 | 
4 | const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
5 |   "group_0": (() => {
6 |     const schema = S.Union(S.Struct({ "$comment": S.optional(S.String), "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "contentEncoding": S.optional(S.String), "contentMediaType": S.optional(S.String), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at <anonymous> (/home/trapani/dev/xschema/typescript/packages/adapters/effect/xschema-harness-4173276304.ts:6:107)
      at /home/trapani/dev/xschema/typescript/packages/adapters/effect/xschema-harness-4173276304.ts:39:115
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | 
2 | import { Schema as S } from "effect"
3 | 
4 | const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
5 |   "group_0": (() => {
6 |     const schema = S.Union(S.Struct({ "$comment": S.optional(S.String), "$id": S.optional(S.String.pipe(S.url())), "$ref": S.optional(S.String.pipe(S.url())), "$schema": S.optional(S.String.pipe(S.url())), "additionalItems": S.optional(S.Unknown), "additionalProperties": S.optional(S.Unknown), "allOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "anyOf": S.optional(S.Array(S.Unknown).pipe(S.minItems(1))), "const": S.optional(S.Unknown), "contains": S.optional(S.Unknown), "contentEncoding": S.optional(S.String), "contentMediaType": S.optional(S.String), "default": S.optional(S.Unknown), "definitions": S.optional(S.Record({ key: S.String, value: S.Unknown })), "dependencies": S.optional(S.Record({ key: S.String, value: S.Union(S.Unknown, S.Array(S.String).pipe(S.filter((arr) => {

TypeError: S.url is not a function. (In 'S.url()', 'S.url' is undefined)
      at <anonymous> (/home/trapani/dev/xschema/typescript/packages/adapters/effect/xschema-harness-4173276304.ts:6:107)
      at /home/trapani/dev/xschema/typescript/packages/adapters/effect/xschema-harness-4173276304.ts:39:115
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
  - Test: matches both
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches root
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

- **oneOf with missing optional property**
  - Test: first oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with missing optional property**
  - Test: second oneOf valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: first valid - valid
  - Expected: `valid`, Got: `false`
- **oneOf with required**
  - Test: second valid - valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 4 failures</summary>

- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>required - 6 failures</summary>

- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `true`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `true`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 9 failures</summary>

- **object type matches objects**
  - Test: a boolean is not an object
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
  - Test: an integer is not an object
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

