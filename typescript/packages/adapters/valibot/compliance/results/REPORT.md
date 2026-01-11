# @xschemadev/valibot Compliance Report

## Summary

| Draft | Passed | Failed | Skipped | Known | Coverage |
| ----- | ------ | ------ | ------- | ----- | -------- |
| draft2019-09 | 848 | 10 | 0 | 57 | 98.8% |
| draft2020-12 | 860 | 10 | 0 | 50 | 98.9% |
| draft3 | 400 | 0 | 0 | 7 | 100.0% |
| draft4 | 557 | 11 | 0 | 0 | 98.1% |
| draft6 | 749 | 19 | 0 | 0 | 97.5% |
| draft7 | 825 | 19 | 0 | 0 | 97.7% |

## Badges

![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-98.8%25-brightgreen)
![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-98.9%25-brightgreen)
![draft3](https://img.shields.io/badge/draft3%20compliance-100.0%25-brightgreen)
![draft4](https://img.shields.io/badge/draft4%20compliance-98.1%25-brightgreen)
![draft6](https://img.shields.io/badge/draft6%20compliance-97.5%25-brightgreen)
![draft7](https://img.shields.io/badge/draft7%20compliance-97.7%25-brightgreen)

## draft2019-09

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ✅ | 21/21 |
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
| dependentSchemas | ✅ | 20/20 |
| enum | ✅ | 45/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 114/114 |
| if-then-else | ✅ | 26/26 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ✅ | 28/28 |
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
| oneOf | ✅ | 27/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ⚠️ | 23/28 |
| propertyNames | ✅ | 20/20 |
| recursiveRef | ✅ | 0/0 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 11/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ✅ | 4/4 |

### Known Behaviors

These tests are intentionally excluded due to documented limitations.

<details>
<summary>Bundler limitation: relative URIs with $id combinations (12 tests)</summary>

- `draft2019-09/ref/URN ref with nested pointer ref/a non-string is invalid`
- `draft2019-09/ref/URN ref with nested pointer ref/a string is valid`
- `draft2019-09/ref/refs with relative uris and defs/invalid on inner field`
- `draft2019-09/ref/refs with relative uris and defs/invalid on outer field`
- `draft2019-09/ref/refs with relative uris and defs/valid on both fields`
- `draft2019-09/ref/relative refs with absolute uris and defs/invalid on inner field`
- `draft2019-09/ref/relative refs with absolute uris and defs/invalid on outer field`
- `draft2019-09/ref/relative refs with absolute uris and defs/valid on both fields`
- `draft2019-09/refRemote/anchor within remote ref/remote anchor invalid`
- `draft2019-09/refRemote/anchor within remote ref/remote anchor valid`
- `draft2019-09/refRemote/retrieved nested refs resolve relative to their URI not $id/number is invalid`
- `draft2019-09/refRemote/retrieved nested refs resolve relative to their URI not $id/string is valid`

</details>

<details>
<summary>Metaschema validation requires fetching schemas containing dynamic refs (4 tests)</summary>

- `draft2019-09/defs/validate definition against metaschema/invalid definition schema`
- `draft2019-09/defs/validate definition against metaschema/valid definition schema`
- `draft2019-09/ref/remote ref, containing refs itself/remote ref invalid`
- `draft2019-09/ref/remote ref, containing refs itself/remote ref valid`

</details>

<details>
<summary>Recursive references ($recursiveRef/$recursiveAnchor) require runtime scope tracking (40 tests)</summary>

- `draft2019-09/recursiveRef/$recursiveRef with $recursiveAnchor: false works like $ref/integer does not match as a property value`
- `draft2019-09/recursiveRef/$recursiveRef with $recursiveAnchor: false works like $ref/integer matches at the outer level`
- `draft2019-09/recursiveRef/$recursiveRef with $recursiveAnchor: false works like $ref/single level match`
- `draft2019-09/recursiveRef/$recursiveRef with $recursiveAnchor: false works like $ref/two levels, integer does not match as a property value`
- `draft2019-09/recursiveRef/$recursiveRef with $recursiveAnchor: false works like $ref/two levels, properties match with inner definition`
- `draft2019-09/recursiveRef/$recursiveRef with nesting/integer matches at the outer level`
- `draft2019-09/recursiveRef/$recursiveRef with nesting/integer now matches as a property value`
- `draft2019-09/recursiveRef/$recursiveRef with nesting/single level match`
- `draft2019-09/recursiveRef/$recursiveRef with nesting/two levels, properties match with $recursiveRef`
- `draft2019-09/recursiveRef/$recursiveRef with nesting/two levels, properties match with inner definition`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor in the initial target schema resource/leaf node does not match: recursion uses the inner schema`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor in the initial target schema resource/leaf node does not match; no recursion`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor in the initial target schema resource/leaf node matches: recursion uses the inner schema`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor in the outer schema resource/leaf node does not match: recursion only uses inner schema`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor in the outer schema resource/leaf node does not match; no recursion`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor in the outer schema resource/leaf node matches: recursion only uses inner schema`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor works like $ref/integer does not match as a property value`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor works like $ref/integer matches at the outer level`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor works like $ref/single level match`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor works like $ref/two levels, integer does not match as a property value`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor works like $ref/two levels, properties match with inner definition`
- `draft2019-09/recursiveRef/$recursiveRef without $recursiveAnchor works like $ref/match`
- `draft2019-09/recursiveRef/$recursiveRef without $recursiveAnchor works like $ref/mismatch`
- `draft2019-09/recursiveRef/$recursiveRef without $recursiveAnchor works like $ref/recursive match`
- `draft2019-09/recursiveRef/$recursiveRef without $recursiveAnchor works like $ref/recursive mismatch`
- `draft2019-09/recursiveRef/$recursiveRef without using nesting/integer does not match as a property value`
- `draft2019-09/recursiveRef/$recursiveRef without using nesting/integer matches at the outer level`
- `draft2019-09/recursiveRef/$recursiveRef without using nesting/single level match`
- `draft2019-09/recursiveRef/$recursiveRef without using nesting/two levels, no match`
- `draft2019-09/recursiveRef/$recursiveRef without using nesting/two levels, properties match with inner definition`
- `draft2019-09/recursiveRef/dynamic $recursiveRef destination (not predictable at schema compile time)/integer node`
- `draft2019-09/recursiveRef/dynamic $recursiveRef destination (not predictable at schema compile time)/numeric node`
- `draft2019-09/recursiveRef/multiple dynamic paths to the $recursiveRef keyword/recurse to anyLeafNode - floats are allowed`
- `draft2019-09/recursiveRef/multiple dynamic paths to the $recursiveRef keyword/recurse to integerNode - floats are not allowed`
- `draft2019-09/ref/$ref with $recursiveAnchor/extra items allowed for inner arrays`
- `draft2019-09/ref/$ref with $recursiveAnchor/extra items disallowed for root`
- `draft2019-09/unevaluatedItems/unevaluatedItems with $recursiveRef/with no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with $recursiveRef/with unevaluated items`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with $recursiveRef/with no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with $recursiveRef/with unevaluated properties`

</details>

<details>
<summary>Vocabulary control: custom metaschemas disabling validation vocabulary not supported (1 test)</summary>

- `draft2019-09/vocabulary/schema that uses custom metaschema with with no validation vocabulary/no validation: invalid number, but it still validates`

</details>

### Unexpected Failures

<details>
<summary>properties - 5 failures</summary>

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
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>required - 5 failures</summary>

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
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

## draft2020-12

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ✅ | 21/21 |
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
| dependentSchemas | ✅ | 20/20 |
| dynamicRef | ✅ | 0/0 |
| enum | ✅ | 45/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 133/133 |
| if-then-else | ✅ | 26/26 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ✅ | 29/29 |
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
| oneOf | ✅ | 27/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| prefixItems | ✅ | 11/11 |
| properties | ⚠️ | 23/28 |
| propertyNames | ✅ | 20/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ⚠️ | 11/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ✅ | 4/4 |

### Known Behaviors

These tests are intentionally excluded due to documented limitations.

<details>
<summary>Bundler limitation: relative URIs with $id combinations (12 tests)</summary>

- `draft2020-12/ref/URN ref with nested pointer ref/a non-string is invalid`
- `draft2020-12/ref/URN ref with nested pointer ref/a string is valid`
- `draft2020-12/ref/refs with relative uris and defs/invalid on inner field`
- `draft2020-12/ref/refs with relative uris and defs/invalid on outer field`
- `draft2020-12/ref/refs with relative uris and defs/valid on both fields`
- `draft2020-12/ref/relative refs with absolute uris and defs/invalid on inner field`
- `draft2020-12/ref/relative refs with absolute uris and defs/invalid on outer field`
- `draft2020-12/ref/relative refs with absolute uris and defs/valid on both fields`
- `draft2020-12/refRemote/anchor within remote ref/remote anchor invalid`
- `draft2020-12/refRemote/anchor within remote ref/remote anchor valid`
- `draft2020-12/refRemote/retrieved nested refs resolve relative to their URI not $id/number is invalid`
- `draft2020-12/refRemote/retrieved nested refs resolve relative to their URI not $id/string is valid`

</details>

<details>
<summary>Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (33 tests)</summary>

- `draft2020-12/dynamicRef/$dynamicRef points to a boolean schema/follow $dynamicRef to a false schema`
- `draft2020-12/dynamicRef/$dynamicRef points to a boolean schema/follow $dynamicRef to a true schema`
- `draft2020-12/dynamicRef/$dynamicRef skips over intermediate resources - direct reference/integer property passes`
- `draft2020-12/dynamicRef/$dynamicRef skips over intermediate resources - direct reference/string property fails`
- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $defs first/correct extended schema`
- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $defs first/incorrect extended schema`
- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $defs first/incorrect parent schema`
- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $ref first/correct extended schema`
- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $ref first/incorrect extended schema`
- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $ref first/incorrect parent schema`
- `draft2020-12/dynamicRef/$ref to $dynamicRef finds detached $dynamicAnchor/non-number is invalid`
- `draft2020-12/dynamicRef/$ref to $dynamicRef finds detached $dynamicAnchor/number is valid`
- `draft2020-12/dynamicRef/A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor resolves to the first $dynamicAnchor in the dynamic scope/The recursive part is not valid against the root`
- `draft2020-12/dynamicRef/A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor resolves to the first $dynamicAnchor in the dynamic scope/The recursive part is valid against the root`
- `draft2020-12/dynamicRef/A $dynamicRef that initially resolves to a schema without a matching $dynamicAnchor behaves like a normal $ref to $anchor/The recursive part doesn't need to validate against the root`
- `draft2020-12/dynamicRef/A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor/An array containing non-strings is invalid`
- `draft2020-12/dynamicRef/A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor/An array of strings is valid`
- `draft2020-12/dynamicRef/A $dynamicRef to an $anchor in the same schema resource behaves like a normal $ref to an $anchor/An array containing non-strings is invalid`
- `draft2020-12/dynamicRef/A $dynamicRef to an $anchor in the same schema resource behaves like a normal $ref to an $anchor/An array of strings is valid`
- `draft2020-12/dynamicRef/A $ref to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor/An array containing non-strings is invalid`
- `draft2020-12/dynamicRef/A $ref to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor/An array of strings is valid`
- `draft2020-12/dynamicRef/after leaving a dynamic scope, it is not used by a $dynamicRef//then/$defs/thingy is the final stop for the $dynamicRef`
- `draft2020-12/dynamicRef/after leaving a dynamic scope, it is not used by a $dynamicRef/first_scope is not in dynamic scope for the $dynamicRef`
- `draft2020-12/dynamicRef/after leaving a dynamic scope, it is not used by a $dynamicRef/string matches /$defs/thingy, but the $dynamicRef does not stop here`
- `draft2020-12/dynamicRef/multiple dynamic paths to the $dynamicRef keyword/number list with number values`
- `draft2020-12/dynamicRef/multiple dynamic paths to the $dynamicRef keyword/number list with string values`
- `draft2020-12/dynamicRef/multiple dynamic paths to the $dynamicRef keyword/string list with number values`
- `draft2020-12/dynamicRef/multiple dynamic paths to the $dynamicRef keyword/string list with string values`
- `draft2020-12/dynamicRef/strict-tree schema, guards against misspelled properties/instance with correct field`
- `draft2020-12/dynamicRef/strict-tree schema, guards against misspelled properties/instance with misspelled field`
- `draft2020-12/dynamicRef/tests for implementation dynamic anchor and reference link/correct extended schema`
- `draft2020-12/dynamicRef/tests for implementation dynamic anchor and reference link/incorrect extended schema`
- `draft2020-12/dynamicRef/tests for implementation dynamic anchor and reference link/incorrect parent schema`

</details>

<details>
<summary>Metaschema validation requires fetching schemas containing dynamic refs (4 tests)</summary>

- `draft2020-12/defs/validate definition against metaschema/invalid definition schema`
- `draft2020-12/defs/validate definition against metaschema/valid definition schema`
- `draft2020-12/ref/remote ref, containing refs itself/remote ref invalid`
- `draft2020-12/ref/remote ref, containing refs itself/remote ref valid`

</details>

<details>
<summary>Vocabulary control: custom metaschemas disabling validation vocabulary not supported (1 test)</summary>

- `draft2020-12/vocabulary/schema that uses custom metaschema with with no validation vocabulary/no validation: invalid number, but it still validates`

</details>

### Unexpected Failures

<details>
<summary>properties - 5 failures</summary>

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
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>required - 5 failures</summary>

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
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

## draft3

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 14/14 |
| additionalProperties | ✅ | 16/16 |
| default | ✅ | 7/7 |
| dependencies | ✅ | 18/18 |
| disallow | ✅ | 9/9 |
| divisibleBy | ✅ | 8/8 |
| enum | ✅ | 16/16 |
| extends | ✅ | 10/10 |
| format | ✅ | 60/60 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ✅ | 7/7 |
| maxItems | ✅ | 4/4 |
| maxLength | ✅ | 5/5 |
| maximum | ✅ | 14/14 |
| minItems | ✅ | 4/4 |
| minLength | ✅ | 5/5 |
| minimum | ✅ | 13/13 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 17/17 |
| properties | ✅ | 15/15 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 8/8 |
| required | ✅ | 4/4 |
| type | ✅ | 73/73 |
| uniqueItems | ✅ | 62/62 |

### Known Behaviors

These tests are intentionally excluded due to documented limitations.

<details>
<summary>Draft3 type with inline schemas is legacy feature removed in draft4+ (7 tests)</summary>

- `draft3/type/applies a nested schema/an object is invalid otherwise`
- `draft3/type/types can include schemas/a boolean is invalid`
- `draft3/type/types can include schemas/a float is invalid`
- `draft3/type/types can include schemas/a string is invalid`
- `draft3/type/types can include schemas/an integer is invalid`
- `draft3/type/types can include schemas/null is invalid`
- `draft3/type/types from separate schemas are merged/an integer is invalid`

</details>

## draft4

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 17/17 |
| additionalProperties | ✅ | 16/16 |
| allOf | ✅ | 27/27 |
| anyOf | ✅ | 15/15 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ✅ | 29/29 |
| enum | ✅ | 49/49 |
| format | ✅ | 36/36 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ✅ | 21/21 |
| maxItems | ✅ | 4/4 |
| maxLength | ✅ | 5/5 |
| maxProperties | ✅ | 8/8 |
| maximum | ✅ | 14/14 |
| minItems | ✅ | 4/4 |
| minLength | ✅ | 5/5 |
| minProperties | ✅ | 6/6 |
| minimum | ✅ | 17/17 |
| multipleOf | ✅ | 10/10 |
| not | ✅ | 20/20 |
| oneOf | ✅ | 23/23 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 18/18 |
| properties | ⚠️ | 19/24 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/15 |
| type | ✅ | 79/79 |
| uniqueItems | ✅ | 69/69 |

### Unexpected Failures

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 5 failures</summary>

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
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>required - 5 failures</summary>

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
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

## draft6

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ✅ | 16/16 |
| allOf | ✅ | 30/30 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 19/19 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ✅ | 36/36 |
| enum | ✅ | 45/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 54/54 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ✅ | 28/28 |
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
| oneOf | ✅ | 27/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ⚠️ | 23/28 |
| propertyNames | ✅ | 20/20 |
| ref | ❌ | 0/6 |
| refRemote | ⚠️ | 21/23 |
| required | ⚠️ | 11/16 |
| type | ✅ | 80/80 |
| uniqueItems | ✅ | 69/69 |

### Unexpected Failures

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 5 failures</summary>

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
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

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
<summary>required - 5 failures</summary>

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
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

## draft7

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ✅ | 16/16 |
| allOf | ✅ | 30/30 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 21/21 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ✅ | 36/36 |
| enum | ✅ | 45/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 102/102 |
| if-then-else | ✅ | 26/26 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ✅ | 28/28 |
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
| oneOf | ✅ | 27/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ⚠️ | 23/28 |
| propertyNames | ✅ | 20/20 |
| ref | ❌ | 0/6 |
| refRemote | ⚠️ | 21/23 |
| required | ⚠️ | 11/16 |
| type | ✅ | 80/80 |
| uniqueItems | ✅ | 69/69 |

### Unexpected Failures

<details>
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>properties - 5 failures</summary>

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
  - Test: none of the properties mentioned
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

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
<summary>required - 5 failures</summary>

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
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

