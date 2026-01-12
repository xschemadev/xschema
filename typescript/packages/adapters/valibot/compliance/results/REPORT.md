# @xschemadev/valibot Compliance Report

## Summary

| Draft | Passed | Failed | Skipped | Unsupported | Coverage |
| ----- | ------ | ------ | ------- | ----------- | -------- |
| draft2019-09 | 849 | 10 | 0 | 40 | 98.8% |
| draft2020-12 | 861 | 10 | 0 | 33 | 98.9% |
| draft3 | 316 | 93 | 0 | 0 | 77.3% |
| draft4 | 492 | 78 | 0 | 0 | 86.3% |
| draft6 | 693 | 77 | 0 | 0 | 90.0% |
| draft7 | 769 | 77 | 0 | 0 | 90.9% |

## Badges

![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-98.8%25-brightgreen)
![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-98.9%25-brightgreen)
![draft3](https://img.shields.io/badge/draft3%20compliance-77.3%25-red)
![draft4](https://img.shields.io/badge/draft4%20compliance-86.3%25-yellow)
![draft6](https://img.shields.io/badge/draft6%20compliance-90.0%25-yellow)
![draft7](https://img.shields.io/badge/draft7%20compliance-90.9%25-yellow)

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
| vocabulary | ✅ | 5/5 |

### Unsupported Features

These tests are intentionally excluded due to documented limitations.

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
| vocabulary | ✅ | 5/5 |

### Unsupported Features

These tests are intentionally excluded due to documented limitations.

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
| additionalItems | ⚠️ | 4/14 |
| additionalProperties | ✅ | 16/16 |
| default | ✅ | 7/7 |
| dependencies | ⚠️ | 11/18 |
| disallow | ✅ | 9/9 |
| divisibleBy | ✅ | 8/8 |
| enum | ⚠️ | 10/16 |
| extends | ⚠️ | 2/10 |
| format | ✅ | 60/60 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 4/7 |
| maxItems | ✅ | 4/4 |
| maxLength | ✅ | 5/5 |
| maximum | ⚠️ | 8/14 |
| minItems | ✅ | 4/4 |
| minLength | ✅ | 5/5 |
| minimum | ⚠️ | 11/13 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 17/17 |
| properties | ✅ | 15/15 |
| ref | ❌ | 0/2 |
| refRemote | ✅ | 8/8 |
| required | ⚠️ | 1/4 |
| type | ⚠️ | 60/80 |
| uniqueItems | ⚠️ | 36/62 |

### Unexpected Failures

<details>
<summary>additionalItems - 10 failures</summary>

- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/additionalItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/additionalItems/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/additionalItems/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/additionalItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems with heterogeneous array**
  - Test: valid instance
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/additionalItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/additionalItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/additionalItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: equal number of items present
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/additionalItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/additionalItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/additionalItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

<details>
<summary>dependencies - 7 failures</summary>

- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/dependencies/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/dependencies/bar': 'anyOf' failed
  - at '/dependencies/bar': got string, want boolean or object
  - at '/dependencies/bar': got string, want array`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/dependencies/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/dependencies/bar': 'anyOf' failed
  - at '/dependencies/bar': got string, want boolean or object
  - at '/dependencies/bar': got string, want array`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/dependencies/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/dependencies/bar': 'anyOf' failed
  - at '/dependencies/bar': got string, want boolean or object
  - at '/dependencies/bar': got string, want array`
- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/dependencies/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/dependencies/bar': 'anyOf' failed
  - at '/dependencies/bar': got string, want boolean or object
  - at '/dependencies/bar': got string, want array`
- **dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/dependencies/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/dependencies/bar': 'anyOf' failed
  - at '/dependencies/bar': got string, want boolean or object
  - at '/dependencies/bar': got string, want array`
- **dependencies**
  - Test: nondependant
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/dependencies/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/dependencies/bar': 'anyOf' failed
  - at '/dependencies/bar': got string, want boolean or object
  - at '/dependencies/bar': got string, want array`
- **dependencies**
  - Test: with dependency
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/dependencies/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/dependencies/bar': 'anyOf' failed
  - at '/dependencies/bar': got string, want boolean or object
  - at '/dependencies/bar': got string, want array`

</details>

<details>
<summary>enum - 6 failures</summary>

- **enums in properties**
  - Test: both properties are valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/enum/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **enums in properties**
  - Test: missing all properties is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/enum/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **enums in properties**
  - Test: missing optional property is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/enum/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **enums in properties**
  - Test: missing required property is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/enum/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **enums in properties**
  - Test: wrong bar value
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/enum/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **enums in properties**
  - Test: wrong foo value
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/enum/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`

</details>

<details>
<summary>extends - 8 failures</summary>

- **extends**
  - Test: extends
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/extends/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **extends**
  - Test: mismatch extended
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/extends/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **extends**
  - Test: mismatch extends
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/extends/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **extends**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/extends/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **multiple extends**
  - Test: mismatch both
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/extends/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **multiple extends**
  - Test: mismatch first extends
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/extends/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **multiple extends**
  - Test: mismatch second extends
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/extends/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`
- **multiple extends**
  - Test: valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/extends/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/bar': 'allOf' failed
    - at '/properties/bar/required': got boolean, want array`

</details>

<details>
<summary>items - 3 failures</summary>

- **an array of schemas for items**
  - Test: correct types
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/items/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

<details>
<summary>maximum - 6 failures</summary>

- **exclusiveMaximum validation**
  - Test: below the maximum is still valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/maximum/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/maximum/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`
- **maximum validation (explicit false exclusivity)**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/maximum/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`
- **maximum validation (explicit false exclusivity)**
  - Test: below the maximum is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/maximum/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`
- **maximum validation (explicit false exclusivity)**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/maximum/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`
- **maximum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/maximum/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`

</details>

<details>
<summary>minimum - 2 failures</summary>

- **exclusiveMinimum validation**
  - Test: above the minimum is still valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/minimum/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMinimum': got boolean, want number`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/minimum/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMinimum': got boolean, want number`

</details>

<details>
<summary>ref - 2 failures</summary>

- **relative pointer ref to array**
  - Test: match array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/ref/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/ref/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

<details>
<summary>required - 3 failures</summary>

- **required explicitly false validation**
  - Test: not required if required is false
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/required/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/foo': 'allOf' failed
    - at '/properties/foo/required': got boolean, want array`
- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/required/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/foo': 'allOf' failed
    - at '/properties/foo/required': got boolean, want array`
- **required validation**
  - Test: present required property is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/required/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/properties/foo': 'allOf' failed
    - at '/properties/foo/required': got boolean, want array`

</details>

<details>
<summary>type - 20 failures</summary>

- **any type matches any type**
  - Test: any type includes array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type': got string, want array`
- **any type matches any type**
  - Test: any type includes boolean
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type': got string, want array`
- **any type matches any type**
  - Test: any type includes float
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type': got string, want array`
- **any type matches any type**
  - Test: any type includes integers
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type': got string, want array`
- **any type matches any type**
  - Test: any type includes null
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type': got string, want array`
- **any type matches any type**
  - Test: any type includes object
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type': got string, want array`
- **any type matches any type**
  - Test: any type includes string
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type': got string, want array`
- **applies a nested schema**
  - Test: an integer is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_10: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **applies a nested schema**
  - Test: an object is invalid otherwise
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_10: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **applies a nested schema**
  - Test: an object is valid only if it is fully valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_10: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **types can include schemas**
  - Test: a boolean is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_9: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **types can include schemas**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_9: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **types can include schemas**
  - Test: a string is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_9: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **types can include schemas**
  - Test: an array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_9: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **types can include schemas**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_9: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **types can include schemas**
  - Test: an object is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_9: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **types can include schemas**
  - Test: null is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_9: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **types from separate schemas are merged**
  - Test: a string is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_11: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type': validation failed
      - at '/type/0': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
      - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **types from separate schemas are merged**
  - Test: an array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_11: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type': validation failed
      - at '/type/0': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
      - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`
- **types from separate schemas are merged**
  - Test: an integer is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/type/group_11: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/type': 'anyOf' failed
    - at '/type': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
    - at '/type': validation failed
      - at '/type/0': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'
      - at '/type/1': value must be one of 'array', 'boolean', 'integer', 'null', 'number', 'object', 'string'`

</details>

<details>
<summary>uniqueItems - 26 failures</summary>

- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft3/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

## draft4

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 4/17 |
| additionalProperties | ✅ | 16/16 |
| allOf | ✅ | 27/27 |
| anyOf | ✅ | 15/15 |
| default | ✅ | 7/7 |
| definitions | ✅ | 0/0 |
| dependencies | ✅ | 29/29 |
| enum | ✅ | 49/49 |
| format | ✅ | 36/36 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 8/21 |
| maxItems | ✅ | 4/4 |
| maxLength | ✅ | 5/5 |
| maxProperties | ✅ | 8/8 |
| maximum | ⚠️ | 8/14 |
| minItems | ✅ | 4/4 |
| minLength | ✅ | 5/5 |
| minProperties | ✅ | 6/6 |
| minimum | ⚠️ | 11/17 |
| multipleOf | ✅ | 10/10 |
| not | ✅ | 20/20 |
| oneOf | ✅ | 23/23 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 18/18 |
| properties | ⚠️ | 19/24 |
| ref | ❌ | 0/4 |
| refRemote | ✅ | 17/17 |
| required | ⚠️ | 10/15 |
| type | ✅ | 79/79 |
| uniqueItems | ⚠️ | 43/69 |

### Unexpected Failures

<details>
<summary>additionalItems - 13 failures</summary>

- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '': validation failed
    - at '/allOf/0': 'allOf' failed
      - at '/allOf/0/items': got array, want boolean or object
    - at '/items': got array, want boolean or object`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems with heterogeneous array**
  - Test: valid instance
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: equal number of items present
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_6: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/additionalItems/group_6: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

<details>
<summary>items - 13 failures</summary>

- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: correct types
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: valid items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/items/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

<details>
<summary>maximum - 6 failures</summary>

- **exclusiveMaximum validation**
  - Test: below the maximum is still valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/maximum/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/maximum/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`
- **maximum validation (explicit false exclusivity)**
  - Test: above the maximum is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/maximum/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`
- **maximum validation (explicit false exclusivity)**
  - Test: below the maximum is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/maximum/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`
- **maximum validation (explicit false exclusivity)**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/maximum/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`
- **maximum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/maximum/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMaximum': got boolean, want number`

</details>

<details>
<summary>minimum - 6 failures</summary>

- **exclusiveMinimum validation**
  - Test: above the minimum is still valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/minimum/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMinimum': got boolean, want number`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/minimum/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMinimum': got boolean, want number`
- **minimum validation (explicit false exclusivity)**
  - Test: above the minimum is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/minimum/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMinimum': got boolean, want number`
- **minimum validation (explicit false exclusivity)**
  - Test: below the minimum is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/minimum/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMinimum': got boolean, want number`
- **minimum validation (explicit false exclusivity)**
  - Test: boundary point is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/minimum/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMinimum': got boolean, want number`
- **minimum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/minimum/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/exclusiveMinimum': got boolean, want number`

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
<summary>ref - 4 failures</summary>

- **Location-independent identifier**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/ref/group_12: invalid JSON Schema: anchor in "file:///home/trapani/dev/xschema/cli/schema.json#foo" not found in schema "file:///home/trapani/dev/xschema/cli/schema.json"`
- **Location-independent identifier**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/ref/group_12: invalid JSON Schema: anchor in "file:///home/trapani/dev/xschema/cli/schema.json#foo" not found in schema "file:///home/trapani/dev/xschema/cli/schema.json"`
- **relative pointer ref to array**
  - Test: match array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/ref/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/ref/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

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

<details>
<summary>uniqueItems - 26 failures</summary>

- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft4/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

## draft6

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 6/19 |
| additionalProperties | ✅ | 16/16 |
| allOf | ✅ | 30/30 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 19/19 |
| default | ✅ | 7/7 |
| definitions | ✅ | 0/0 |
| dependencies | ✅ | 36/36 |
| enum | ✅ | 45/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 54/54 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 12/28 |
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
| ref | ❌ | 0/10 |
| refRemote | ⚠️ | 21/23 |
| required | ⚠️ | 11/16 |
| type | ✅ | 80/80 |
| uniqueItems | ⚠️ | 43/69 |

### Unexpected Failures

<details>
<summary>additionalItems - 13 failures</summary>

- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_6: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '': validation failed
    - at '/allOf/0': 'allOf' failed
      - at '/allOf/0/items': got array, want boolean or object
    - at '/items': got array, want boolean or object`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_8: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems with heterogeneous array**
  - Test: valid instance
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_8: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: equal number of items present
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/additionalItems/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

<details>
<summary>items - 16 failures</summary>

- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: correct types
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_8: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: valid items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/items/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

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
<summary>ref - 10 failures</summary>

- **Location-independent identifier**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/ref/group_14: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Location-independent identifier**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/ref/group_14: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Location-independent identifier with base URI change in subschema**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/ref/group_16: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A/definitions/B': 'allOf' failed
  - at '/definitions/A/definitions/B/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/ref/group_16: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A/definitions/B': 'allOf' failed
  - at '/definitions/A/definitions/B/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Reference an anchor with a non-relative URI**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/ref/group_15: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Reference an anchor with a non-relative URI**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/ref/group_15: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **URN base URI with URN and anchor ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/ref/group_26: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/bar': 'allOf' failed
  - at '/definitions/bar/$id': '#something' does not match pattern '^[^#]*#?$'`
- **URN base URI with URN and anchor ref**
  - Test: a string is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/ref/group_26: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/bar': 'allOf' failed
  - at '/definitions/bar/$id': '#something' does not match pattern '^[^#]*#?$'`
- **relative pointer ref to array**
  - Test: match array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/ref/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/ref/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

<details>
<summary>refRemote - 2 failures</summary>

- **Location-independent identifier in remote ref**
  - Test: integer is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for external schema http://localhost:1234/draft6/locationIndependentIdentifier.json: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Location-independent identifier in remote ref**
  - Test: string is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for external schema http://localhost:1234/draft6/locationIndependentIdentifier.json: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`

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

<details>
<summary>uniqueItems - 26 failures</summary>

- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft6/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

## draft7

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 6/19 |
| additionalProperties | ✅ | 16/16 |
| allOf | ✅ | 30/30 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 21/21 |
| default | ✅ | 7/7 |
| definitions | ✅ | 0/0 |
| dependencies | ✅ | 36/36 |
| enum | ✅ | 45/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
| format | ✅ | 102/102 |
| if-then-else | ✅ | 26/26 |
| infinite-loop-detection | ✅ | 2/2 |
| items | ⚠️ | 12/28 |
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
| ref | ❌ | 0/10 |
| refRemote | ⚠️ | 21/23 |
| required | ⚠️ | 11/16 |
| type | ✅ | 80/80 |
| uniqueItems | ⚠️ | 43/69 |

### Unexpected Failures

<details>
<summary>additionalItems - 13 failures</summary>

- **additionalItems are allowed by default**
  - Test: only the first item is validated
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems as schema**
  - Test: additional items match schema
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_0: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems does not look in applicators, invalid case**
  - Test: items defined in allOf are not examined
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_6: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '': validation failed
    - at '/allOf/0': 'allOf' failed
      - at '/allOf/0/items': got array, want boolean or object
    - at '/items': got array, want boolean or object`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_8: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **additionalItems with heterogeneous array**
  - Test: valid instance
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_8: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: empty array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: equal number of items present
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (1)
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array of items with no additionalItems permitted**
  - Test: fewer number of items present (2)
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_3: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items validation adjusts the starting index for additionalItems**
  - Test: valid items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items validation adjusts the starting index for additionalItems**
  - Test: wrong type of second item
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/additionalItems/group_7: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

<details>
<summary>items - 16 failures</summary>

- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: array with additional items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: correct types
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **an array of schemas for items**
  - Test: wrong types
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **array-form items with null instance elements**
  - Test: allows null elements
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_8: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: too many items
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: too many sub-items
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: valid items
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: wrong item
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items and subitems**
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/item': 'allOf' failed
  - at '/definitions/item/items': got array, want boolean or object
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items with boolean schemas**
  - Test: array with two items is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **items with boolean schemas**
  - Test: empty array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/items/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

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
<summary>ref - 10 failures</summary>

- **Location-independent identifier**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/ref/group_14: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Location-independent identifier**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/ref/group_14: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Location-independent identifier with base URI change in subschema**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/ref/group_16: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A/definitions/B': 'allOf' failed
  - at '/definitions/A/definitions/B/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/ref/group_16: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A/definitions/B': 'allOf' failed
  - at '/definitions/A/definitions/B/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Reference an anchor with a non-relative URI**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/ref/group_15: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Reference an anchor with a non-relative URI**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/ref/group_15: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **URN base URI with URN and anchor ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/ref/group_27: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/bar': 'allOf' failed
  - at '/definitions/bar/$id': '#something' does not match pattern '^[^#]*#?$'`
- **URN base URI with URN and anchor ref**
  - Test: a string is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/ref/group_27: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/bar': 'allOf' failed
  - at '/definitions/bar/$id': '#something' does not match pattern '^[^#]*#?$'`
- **relative pointer ref to array**
  - Test: match array
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/ref/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **relative pointer ref to array**
  - Test: mismatch array
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/ref/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

<details>
<summary>refRemote - 2 failures</summary>

- **Location-independent identifier in remote ref**
  - Test: integer is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for external schema http://localhost:1234/draft7/locationIndependentIdentifier.json: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`
- **Location-independent identifier in remote ref**
  - Test: string is invalid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for external schema http://localhost:1234/draft7/locationIndependentIdentifier.json: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '/definitions/A': 'allOf' failed
  - at '/definitions/A/$id': '#foo' does not match pattern '^[^#]*#?$'`

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

<details>
<summary>uniqueItems - 26 failures</summary>

- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_1: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_2: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: non-unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_4: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: [true, true] from items array is valid
  - Expected: `valid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: bundling error: validation failed for compliance://draft7/uniqueItems/group_5: invalid JSON Schema: "file:///home/trapani/dev/xschema/cli/schema.json#" is not valid against metaschema: jsonschema validation failed with 'https://json-schema.org/draft/2020-12/schema#'
- at '': 'allOf' failed
  - at '/items': got array, want boolean or object`

</details>

