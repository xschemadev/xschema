# @xschemadev/valibot Compliance Report

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2019-09 | 769 | 146 | 0 | 84.0% |
| draft2020-12 | 782 | 138 | 0 | 85.0% |
| draft3 | 365 | 42 | 0 | 89.7% |
| draft4 | 494 | 74 | 0 | 87.0% |
| draft6 | 673 | 95 | 0 | 87.6% |
| draft7 | 749 | 95 | 0 | 88.7% |

## Badges

![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-84.0%25-yellow)
![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-85.0%25-yellow)
![draft3](https://img.shields.io/badge/draft3%20compliance-89.7%25-yellow)
![draft4](https://img.shields.io/badge/draft4%20compliance-87.0%25-yellow)
![draft6](https://img.shields.io/badge/draft6%20compliance-87.6%25-yellow)
![draft7](https://img.shields.io/badge/draft7%20compliance-88.7%25-yellow)

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
| defs | ❌ | 0/2 |
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
| recursiveRef | ❌ | 0/34 |
| ref | ❌ | 0/12 |
| refRemote | ❌ | 0/4 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 79/80 |
| unevaluatedItems | ❌ | 0/2 |
| unevaluatedProperties | ❌ | 0/2 |
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
<summary>defs - 2 failures</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: bundling error: failed to bundle schema from "https://json-schema.org/draft/2019-09/schema": unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: bundling error: failed to bundle schema from "https://json-schema.org/draft/2019-09/schema": unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`

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
<summary>recursiveRef - 34 failures</summary>

- **$recursiveRef with $recursiveAnchor: false works like $ref**
  - Test: integer does not match as a property value
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with $recursiveAnchor: false works like $ref**
  - Test: integer matches at the outer level
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with $recursiveAnchor: false works like $ref**
  - Test: single level match
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with $recursiveAnchor: false works like $ref**
  - Test: two levels, integer does not match as a property value
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with $recursiveAnchor: false works like $ref**
  - Test: two levels, properties match with inner definition
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with nesting**
  - Test: integer matches at the outer level
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with nesting**
  - Test: integer now matches as a property value
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with nesting**
  - Test: single level match
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with nesting**
  - Test: two levels, properties match with $recursiveRef
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with nesting**
  - Test: two levels, properties match with inner definition
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor in the initial target schema resource**
  - Test: leaf node does not match: recursion uses the inner schema
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor in the initial target schema resource**
  - Test: leaf node does not match; no recursion
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor in the initial target schema resource**
  - Test: leaf node matches: recursion uses the inner schema
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor in the outer schema resource**
  - Test: leaf node does not match: recursion only uses inner schema
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor in the outer schema resource**
  - Test: leaf node does not match; no recursion
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor in the outer schema resource**
  - Test: leaf node matches: recursion only uses inner schema
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor works like $ref**
  - Test: integer does not match as a property value
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor works like $ref**
  - Test: integer matches at the outer level
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor works like $ref**
  - Test: single level match
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor works like $ref**
  - Test: two levels, integer does not match as a property value
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef with no $recursiveAnchor works like $ref**
  - Test: two levels, properties match with inner definition
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef without $recursiveAnchor works like $ref**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveRef": dynamic and recursive references are not supported`
- **$recursiveRef without $recursiveAnchor works like $ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveRef": dynamic and recursive references are not supported`
- **$recursiveRef without $recursiveAnchor works like $ref**
  - Test: recursive match
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveRef": dynamic and recursive references are not supported`
- **$recursiveRef without $recursiveAnchor works like $ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveRef": dynamic and recursive references are not supported`
- **$recursiveRef without using nesting**
  - Test: integer does not match as a property value
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef without using nesting**
  - Test: integer matches at the outer level
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef without using nesting**
  - Test: single level match
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef without using nesting**
  - Test: two levels, no match
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$recursiveRef without using nesting**
  - Test: two levels, properties match with inner definition
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **dynamic $recursiveRef destination (not predictable at schema compile time)**
  - Test: integer node
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **dynamic $recursiveRef destination (not predictable at schema compile time)**
  - Test: numeric node
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **multiple dynamic paths to the $recursiveRef keyword**
  - Test: recurse to anyLeafNode - floats are allowed
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **multiple dynamic paths to the $recursiveRef keyword**
  - Test: recurse to integerNode - floats are not allowed
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`

</details>

<details>
<summary>ref - 12 failures</summary>

- **$ref with $recursiveAnchor**
  - Test: extra items allowed for inner arrays
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **$ref with $recursiveAnchor**
  - Test: extra items disallowed for root
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **URN ref with nested pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/bar" points to missing target: key "bar" not found`
- **URN ref with nested pointer ref**
  - Test: a string is valid
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/bar" points to missing target: key "bar" not found`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **refs with relative uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to bundle schema from "https://json-schema.org/draft/2019-09/schema": unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to bundle schema from "https://json-schema.org/draft/2019-09/schema": unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`

</details>

<details>
<summary>refRemote - 4 failures</summary>

- **anchor within remote ref**
  - Test: remote anchor invalid
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_draft2019_09_locationIndependentIdentifier_jsonfoo" points to missing target: key "localhost_1234_draft2019_09_locationIndependentIdentifier_jsonfoo" not found`
- **anchor within remote ref**
  - Test: remote anchor valid
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_draft2019_09_locationIndependentIdentifier_jsonfoo" points to missing target: key "localhost_1234_draft2019_09_locationIndependentIdentifier_jsonfoo" not found`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_draft2019_09_nested_foo_ref_string_json/$defs/localhost_1234_draft2019_09_nested_string_json" points to missing target: key "$defs" not found`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: string is valid
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_draft2019_09_nested_foo_ref_string_json/$defs/localhost_1234_draft2019_09_nested_string_json" points to missing target: key "$defs" not found`

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
<summary>unevaluatedItems - 2 failures</summary>

- **unevaluatedItems with $recursiveRef**
  - Test: with no unevaluated items
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **unevaluatedItems with $recursiveRef**
  - Test: with unevaluated items
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`

</details>

<details>
<summary>unevaluatedProperties - 2 failures</summary>

- **unevaluatedProperties with $recursiveRef**
  - Test: with no unevaluated properties
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`
- **unevaluatedProperties with $recursiveRef**
  - Test: with unevaluated properties
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$recursiveAnchor": dynamic and recursive references are not supported`

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
| defs | ❌ | 0/2 |
| dependentRequired | ⚠️ | 17/20 |
| dependentSchemas | ⚠️ | 16/20 |
| dynamicRef | ❌ | 0/33 |
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
| ref | ❌ | 0/10 |
| refRemote | ❌ | 0/4 |
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
<summary>defs - 2 failures</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: bundling error: failed to bundle schema from "https://json-schema.org/draft/2020-12/schema": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: bundling error: failed to bundle schema from "https://json-schema.org/draft/2020-12/schema": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`

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
<summary>dynamicRef - 33 failures</summary>

- **$dynamicRef points to a boolean schema**
  - Test: follow $dynamicRef to a false schema
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicRef": dynamic and recursive references are not supported`
- **$dynamicRef points to a boolean schema**
  - Test: follow $dynamicRef to a true schema
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicRef": dynamic and recursive references are not supported`
- **$dynamicRef skips over intermediate resources - direct reference**
  - Test: integer property passes
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **$dynamicRef skips over intermediate resources - direct reference**
  - Test: string property fails
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **$ref and $dynamicAnchor are independent of order - $defs first**
  - Test: correct extended schema
  - Expected: `valid`, Got: `error: bundling error: failed to bundle schema from "http://localhost:1234/draft2020-12/extendible-dynamic-ref.json": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **$ref and $dynamicAnchor are independent of order - $defs first**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `error: bundling error: failed to bundle schema from "http://localhost:1234/draft2020-12/extendible-dynamic-ref.json": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **$ref and $dynamicAnchor are independent of order - $defs first**
  - Test: incorrect parent schema
  - Expected: `invalid`, Got: `error: bundling error: failed to bundle schema from "http://localhost:1234/draft2020-12/extendible-dynamic-ref.json": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **$ref and $dynamicAnchor are independent of order - $ref first**
  - Test: correct extended schema
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **$ref and $dynamicAnchor are independent of order - $ref first**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **$ref and $dynamicAnchor are independent of order - $ref first**
  - Test: incorrect parent schema
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **$ref to $dynamicRef finds detached $dynamicAnchor**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to bundle schema from "http://localhost:1234/draft2020-12/detached-dynamicref.json": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **$ref to $dynamicRef finds detached $dynamicAnchor**
  - Test: number is valid
  - Expected: `valid`, Got: `error: bundling error: failed to bundle schema from "http://localhost:1234/draft2020-12/detached-dynamicref.json": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor resolves to the first $dynamicAnchor in the dynamic scope**
  - Test: The recursive part is not valid against the root
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor resolves to the first $dynamicAnchor in the dynamic scope**
  - Test: The recursive part is valid against the root
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **A $dynamicRef that initially resolves to a schema without a matching $dynamicAnchor behaves like a normal $ref to $anchor**
  - Test: The recursive part doesn't need to validate against the root
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array of strings is valid
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **A $dynamicRef to an $anchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicRef": dynamic and recursive references are not supported`
- **A $dynamicRef to an $anchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array of strings is valid
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicRef": dynamic and recursive references are not supported`
- **A $ref to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **A $ref to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array of strings is valid
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **after leaving a dynamic scope, it is not used by a $dynamicRef**
  - Test: /then/$defs/thingy is the final stop for the $dynamicRef
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicRef": dynamic and recursive references are not supported`
- **after leaving a dynamic scope, it is not used by a $dynamicRef**
  - Test: first_scope is not in dynamic scope for the $dynamicRef
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicRef": dynamic and recursive references are not supported`
- **after leaving a dynamic scope, it is not used by a $dynamicRef**
  - Test: string matches /$defs/thingy, but the $dynamicRef does not stop here
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicRef": dynamic and recursive references are not supported`
- **multiple dynamic paths to the $dynamicRef keyword**
  - Test: number list with number values
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **multiple dynamic paths to the $dynamicRef keyword**
  - Test: number list with string values
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **multiple dynamic paths to the $dynamicRef keyword**
  - Test: string list with number values
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **multiple dynamic paths to the $dynamicRef keyword**
  - Test: string list with string values
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **strict-tree schema, guards against misspelled properties**
  - Test: instance with correct field
  - Expected: `valid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **strict-tree schema, guards against misspelled properties**
  - Test: instance with misspelled field
  - Expected: `invalid`, Got: `error: bundling error: unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **tests for implementation dynamic anchor and reference link**
  - Test: correct extended schema
  - Expected: `valid`, Got: `error: bundling error: failed to bundle schema from "http://localhost:1234/draft2020-12/extendible-dynamic-ref.json": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **tests for implementation dynamic anchor and reference link**
  - Test: incorrect extended schema
  - Expected: `invalid`, Got: `error: bundling error: failed to bundle schema from "http://localhost:1234/draft2020-12/extendible-dynamic-ref.json": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **tests for implementation dynamic anchor and reference link**
  - Test: incorrect parent schema
  - Expected: `invalid`, Got: `error: bundling error: failed to bundle schema from "http://localhost:1234/draft2020-12/extendible-dynamic-ref.json": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`

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
<summary>ref - 10 failures</summary>

- **URN ref with nested pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/bar" points to missing target: key "bar" not found`
- **URN ref with nested pointer ref**
  - Test: a string is valid
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/bar" points to missing target: key "bar" not found`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **refs with relative uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/inner" points to missing target: key "$defs" not found`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to bundle schema from "https://json-schema.org/draft/2020-12/schema": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to bundle schema from "https://json-schema.org/draft/2020-12/schema": unsupported keyword "$dynamicAnchor": dynamic and recursive references are not supported`

</details>

<details>
<summary>refRemote - 4 failures</summary>

- **anchor within remote ref**
  - Test: remote anchor invalid
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_draft2020_12_locationIndependentIdentifier_jsonfoo" points to missing target: key "localhost_1234_draft2020_12_locationIndependentIdentifier_jsonfoo" not found`
- **anchor within remote ref**
  - Test: remote anchor valid
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_draft2020_12_locationIndependentIdentifier_jsonfoo" points to missing target: key "localhost_1234_draft2020_12_locationIndependentIdentifier_jsonfoo" not found`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_draft2020_12_nested_foo_ref_string_json/$defs/localhost_1234_draft2020_12_nested_string_json" points to missing target: key "$defs" not found`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: string is valid
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_draft2020_12_nested_foo_ref_string_json/$defs/localhost_1234_draft2020_12_nested_string_json" points to missing target: key "$defs" not found`

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
| maximum | ⚠️ | 12/14 |
| minItems | ⚠️ | 3/4 |
| minLength | ⚠️ | 4/5 |
| minimum | ⚠️ | 11/13 |
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
<summary>maximum - 2 failures</summary>

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
<summary>minimum - 2 failures</summary>

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
| maximum | ⚠️ | 12/14 |
| minItems | ⚠️ | 3/4 |
| minLength | ⚠️ | 4/5 |
| minProperties | ⚠️ | 3/6 |
| minimum | ⚠️ | 14/17 |
| multipleOf | ⚠️ | 9/10 |
| not | ✅ | 20/20 |
| oneOf | ⚠️ | 19/23 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 16/18 |
| properties | ⚠️ | 15/24 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 17/17 |
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
<summary>maximum - 2 failures</summary>

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
<summary>minimum - 3 failures</summary>

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
| ref | ❌ | 0/6 |
| refRemote | ⚠️ | 21/23 |
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
<summary>ref - 6 failures</summary>

- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`
- **refs with relative uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`
- **relative refs with absolute uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`

</details>

<details>
<summary>refRemote - 2 failures</summary>

- **retrieved nested refs resolve relative to their URI not $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_nested_foo_ref_string_json/$defs/localhost_1234_nested_string_json" points to missing target: key "$defs" not found`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: string is valid
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_nested_foo_ref_string_json/$defs/localhost_1234_nested_string_json" points to missing target: key "$defs" not found`

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
| ref | ❌ | 0/6 |
| refRemote | ⚠️ | 21/23 |
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
<summary>ref - 6 failures</summary>

- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`
- **refs with relative uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`
- **relative refs with absolute uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: $ref "#/definitions/inner" points to missing target: key "definitions" not found`

</details>

<details>
<summary>refRemote - 2 failures</summary>

- **retrieved nested refs resolve relative to their URI not $id**
  - Test: number is invalid
  - Expected: `invalid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_nested_foo_ref_string_json/$defs/localhost_1234_nested_string_json" points to missing target: key "$defs" not found`
- **retrieved nested refs resolve relative to their URI not $id**
  - Test: string is valid
  - Expected: `valid`, Got: `error: bundling error: $ref "#/$defs/localhost_1234_nested_foo_ref_string_json/$defs/localhost_1234_nested_string_json" points to missing target: key "$defs" not found`

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

