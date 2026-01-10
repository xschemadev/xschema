# @xschemadev/valibot Compliance Report

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2019-09 | 769 | 90 | 0 | 89.5% |
| draft2020-12 | 782 | 89 | 0 | 89.8% |
| draft3 | 363 | 44 | 0 | 89.2% |
| draft4 | 475 | 76 | 0 | 86.2% |
| draft6 | 652 | 87 | 0 | 88.2% |
| draft7 | 728 | 87 | 0 | 89.3% |

## Badges

![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-89.5%25-yellow)
![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-89.8%25-yellow)
![draft3](https://img.shields.io/badge/draft3%20compliance-89.2%25-yellow)
![draft4](https://img.shields.io/badge/draft4%20compliance-86.2%25-yellow)
![draft6](https://img.shields.io/badge/draft6%20compliance-88.2%25-yellow)
![draft7](https://img.shields.io/badge/draft7%20compliance-89.3%25-yellow)

## draft2019-09

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 16/19 |
| additionalProperties | ⚠️ | 16/21 |
| allOf | ✅ | 30/30 |
| anchor | ✅ | 0/0 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ⚠️ | 19/21 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ⚠️ | 17/20 |
| dependentSchemas | ⚠️ | 16/20 |
| enum | ⚠️ | 43/45 |
| exclusiveMaximum | ⚠️ | 3/4 |
| exclusiveMinimum | ⚠️ | 3/4 |
| format | ✅ | 114/114 |
| if-then-else | ✅ | 26/26 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 16/28 |
| maxContains | ✅ | 12/12 |
| maxItems | ⚠️ | 5/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ⚠️ | 7/10 |
| maximum | ⚠️ | 7/8 |
| minContains | ✅ | 28/28 |
| minItems | ⚠️ | 5/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ⚠️ | 5/8 |
| minimum | ⚠️ | 9/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 0/0 |
| oneOf | ⚠️ | 23/27 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 21/23 |
| properties | ⚠️ | 19/28 |
| propertyNames | ⚠️ | 13/20 |
| recursiveRef | ✅ | 0/0 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 79/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ⚠️ | 67/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalItems - 3 failures</summary>

- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 5 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>contains - 2 failures</summary>

- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>dependentRequired - 3 failures</summary>

- **empty dependents**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>dependentSchemas - 4 failures</summary>

- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

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
<summary>exclusiveMaximum - 1 failure</summary>

- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMinimum - 1 failure</summary>

- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 12 failures</summary>

- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 1 failure</summary>

- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxProperties - 3 failures</summary>

- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maximum - 1 failure</summary>

- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minItems - 1 failure</summary>

- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minProperties - 3 failures</summary>

- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minimum - 2 failures</summary>

- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

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
<summary>pattern - 6 failures</summary>

- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 2 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 9 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>propertyNames - 7 failures</summary>

- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
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
<summary>required - 11 failures</summary>

- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 1 failure</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 2 failures</summary>

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
  - Test: no validation: invalid number, but it still validates
  - Expected: `valid`, Got: `false`

</details>

## draft2020-12

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 16/21 |
| allOf | ✅ | 30/30 |
| anchor | ✅ | 0/0 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ⚠️ | 19/21 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ⚠️ | 17/20 |
| dependentSchemas | ⚠️ | 16/20 |
| dynamicRef | ✅ | 0/0 |
| enum | ⚠️ | 43/45 |
| exclusiveMaximum | ⚠️ | 3/4 |
| exclusiveMinimum | ⚠️ | 3/4 |
| format | ✅ | 133/133 |
| if-then-else | ✅ | 26/26 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 20/29 |
| maxContains | ✅ | 12/12 |
| maxItems | ⚠️ | 5/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ⚠️ | 7/10 |
| maximum | ⚠️ | 7/8 |
| minContains | ✅ | 28/28 |
| minItems | ⚠️ | 5/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ⚠️ | 5/8 |
| minimum | ⚠️ | 9/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 0/0 |
| oneOf | ⚠️ | 23/27 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 21/23 |
| prefixItems | ⚠️ | 6/11 |
| properties | ⚠️ | 19/28 |
| propertyNames | ⚠️ | 13/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 79/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ⚠️ | 67/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 5 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>contains - 2 failures</summary>

- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>dependentRequired - 3 failures</summary>

- **empty dependents**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>dependentSchemas - 4 failures</summary>

- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `true`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `true`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

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
<summary>exclusiveMaximum - 1 failure</summary>

- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMinimum - 1 failure</summary>

- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 9 failures</summary>

- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`
- **prefixItems with no additional items allowed**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxItems - 1 failure</summary>

- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxProperties - 3 failures</summary>

- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maximum - 1 failure</summary>

- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minItems - 1 failure</summary>

- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minProperties - 3 failures</summary>

- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minimum - 2 failures</summary>

- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

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
<summary>pattern - 6 failures</summary>

- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 2 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>prefixItems - 5 failures</summary>

- **a schema given for prefixItems**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **a schema given for prefixItems**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **a schema given for prefixItems**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **prefixItems with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **prefixItems with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 9 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>propertyNames - 7 failures</summary>

- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
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
<summary>required - 11 failures</summary>

- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 1 failure</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 2 failures</summary>

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
  - Test: no validation: invalid number, but it still validates
  - Expected: `valid`, Got: `false`

</details>

## draft3

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 11/14 |
| additionalProperties | ⚠️ | 11/16 |
| default | ✅ | 7/7 |
| dependencies | ⚠️ | 15/18 |
| disallow | ✅ | 9/9 |
| divisibleBy | ⚠️ | 7/8 |
| enum | ⚠️ | 15/16 |
| extends | ✅ | 10/10 |
| format | ✅ | 60/60 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 6/7 |
| maxItems | ⚠️ | 3/4 |
| maxLength | ⚠️ | 4/5 |
| maximum | ⚠️ | 11/14 |
| minItems | ⚠️ | 3/4 |
| minLength | ⚠️ | 4/5 |
| minimum | ⚠️ | 10/13 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 16/17 |
| properties | ⚠️ | 13/15 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 8/8 |
| required | ⚠️ | 3/4 |
| type | ⚠️ | 72/80 |
| uniqueItems | ⚠️ | 60/62 |

### Failures

<details>
<summary>additionalItems - 3 failures</summary>

- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 5 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in extends are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 3 failures</summary>

- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>divisibleBy - 1 failure</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>enum - 1 failure</summary>

- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 1 failure</summary>

- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 1 failure</summary>

- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maximum - 3 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **maximum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minItems - 1 failure</summary>

- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minimum - 3 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>pattern - 6 failures</summary>

- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 1 failure</summary>

- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 2 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>required - 1 failure</summary>

- **required validation**
  - Test: present required property is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>type - 8 failures</summary>

- **applies a nested schema**
  - Test: an object is invalid otherwise
  - Expected: `invalid`, Got: `true`
- **object type matches objects**
  - Test: an array is not an object
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
<summary>uniqueItems - 2 failures</summary>

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
| additionalItems | ⚠️ | 14/17 |
| additionalProperties | ⚠️ | 11/16 |
| allOf | ✅ | 27/27 |
| anyOf | ✅ | 15/15 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 26/29 |
| enum | ⚠️ | 47/49 |
| format | ✅ | 36/36 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 12/21 |
| maxItems | ⚠️ | 3/4 |
| maxLength | ⚠️ | 4/5 |
| maxProperties | ⚠️ | 5/8 |
| maximum | ⚠️ | 11/14 |
| minItems | ⚠️ | 3/4 |
| minLength | ⚠️ | 4/5 |
| minProperties | ⚠️ | 3/6 |
| minimum | ⚠️ | 13/17 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 20/20 |
| oneOf | ⚠️ | 19/23 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 16/18 |
| properties | ⚠️ | 15/24 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 4/15 |
| type | ⚠️ | 78/79 |
| uniqueItems | ⚠️ | 67/69 |

### Failures

<details>
<summary>additionalItems - 3 failures</summary>

- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 5 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
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

- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **dependencies with escaped characters**
  - Test: invalid object 3
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
<summary>items - 9 failures</summary>

- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxItems - 1 failure</summary>

- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxProperties - 3 failures</summary>

- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maximum - 3 failures</summary>

- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **maximum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minItems - 1 failure</summary>

- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minProperties - 3 failures</summary>

- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minimum - 4 failures</summary>

- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

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
<summary>pattern - 6 failures</summary>

- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 2 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 9 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>required - 11 failures</summary>

- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 1 failure</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 2 failures</summary>

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
| additionalItems | ⚠️ | 16/19 |
| additionalProperties | ⚠️ | 11/16 |
| allOf | ✅ | 30/30 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ⚠️ | 17/19 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 32/36 |
| enum | ⚠️ | 43/45 |
| exclusiveMaximum | ⚠️ | 3/4 |
| exclusiveMinimum | ⚠️ | 3/4 |
| format | ✅ | 54/54 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 16/28 |
| maxItems | ⚠️ | 5/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ⚠️ | 7/10 |
| maximum | ⚠️ | 7/8 |
| minItems | ⚠️ | 5/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ⚠️ | 5/8 |
| minimum | ⚠️ | 9/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 21/23 |
| properties | ⚠️ | 19/28 |
| propertyNames | ⚠️ | 13/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 79/80 |
| uniqueItems | ⚠️ | 67/69 |

### Failures

<details>
<summary>additionalItems - 3 failures</summary>

- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 5 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>contains - 2 failures</summary>

- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 4 failures</summary>

- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **dependencies with empty array**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
- **dependencies with escaped characters**
  - Test: invalid object 3
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
<summary>exclusiveMaximum - 1 failure</summary>

- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMinimum - 1 failure</summary>

- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 12 failures</summary>

- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 1 failure</summary>

- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxProperties - 3 failures</summary>

- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maximum - 1 failure</summary>

- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minItems - 1 failure</summary>

- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minProperties - 3 failures</summary>

- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minimum - 2 failures</summary>

- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

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
<summary>pattern - 6 failures</summary>

- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 2 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 9 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>propertyNames - 7 failures</summary>

- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
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
<summary>required - 11 failures</summary>

- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 1 failure</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 2 failures</summary>

- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

## draft7

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 16/19 |
| additionalProperties | ⚠️ | 11/16 |
| allOf | ✅ | 30/30 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ⚠️ | 19/21 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 32/36 |
| enum | ⚠️ | 43/45 |
| exclusiveMaximum | ⚠️ | 3/4 |
| exclusiveMinimum | ⚠️ | 3/4 |
| format | ✅ | 102/102 |
| if-then-else | ✅ | 26/26 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 16/28 |
| maxItems | ⚠️ | 5/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ⚠️ | 7/10 |
| maximum | ⚠️ | 7/8 |
| minItems | ⚠️ | 5/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ⚠️ | 5/8 |
| minimum | ⚠️ | 9/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 38/38 |
| oneOf | ⚠️ | 23/27 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 21/23 |
| properties | ⚠️ | 19/28 |
| propertyNames | ⚠️ | 13/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 79/80 |
| uniqueItems | ⚠️ | 67/69 |

### Failures

<details>
<summary>additionalItems - 3 failures</summary>

- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 5 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `false`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>contains - 2 failures</summary>

- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 4 failures</summary>

- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **dependencies with empty array**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
- **dependencies with escaped characters**
  - Test: invalid object 3
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
<summary>exclusiveMaximum - 1 failure</summary>

- **exclusiveMaximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>exclusiveMinimum - 1 failure</summary>

- **exclusiveMinimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>items - 12 failures</summary>

- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 1 failure</summary>

- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxProperties - 3 failures</summary>

- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maximum - 1 failure</summary>

- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minItems - 1 failure</summary>

- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minProperties - 3 failures</summary>

- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minimum - 2 failures</summary>

- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>multipleOf - 1 failure</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

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
<summary>pattern - 6 failures</summary>

- **pattern validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores booleans
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores floats
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores integers
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **pattern validation**
  - Test: ignores objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>patternProperties - 2 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 9 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>propertyNames - 7 failures</summary>

- **propertyNames validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **propertyNames validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
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
<summary>required - 11 failures</summary>

- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
- **required with escaped characters**
  - Test: object with some properties missing is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>type - 1 failure</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 2 failures</summary>

- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

