# @xschemadev/zod Compliance Report

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2019-09 | 857 | 2 | 0 | 99.8% |
| draft2020-12 | 869 | 2 | 0 | 99.8% |
| draft3 | 394 | 15 | 0 | 96.3% |
| draft4 | 547 | 21 | 0 | 96.3% |
| draft6 | 737 | 2 | 0 | 99.7% |
| draft7 | 813 | 2 | 0 | 99.8% |
| v1 | 746 | 10 | 0 | 98.7% |

## Badges

![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-99.8%25-brightgreen)
![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-99.8%25-brightgreen)
![draft3](https://img.shields.io/badge/draft3%20compliance-96.3%25-brightgreen)
![draft4](https://img.shields.io/badge/draft4%20compliance-96.3%25-brightgreen)
![draft6](https://img.shields.io/badge/draft6%20compliance-99.7%25-brightgreen)
![draft7](https://img.shields.io/badge/draft7%20compliance-99.8%25-brightgreen)
![v1](https://img.shields.io/badge/v1%20compliance-98.7%25-brightgreen)

## draft2019-09

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ⚠️ | 20/21 |
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
| properties | ✅ | 28/28 |
| propertyNames | ✅ | 20/20 |
| recursiveRef | ✅ | 0/0 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ✅ | 16/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 1 failure</summary>

- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
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
| additionalProperties | ⚠️ | 20/21 |
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
| properties | ✅ | 28/28 |
| propertyNames | ✅ | 20/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ✅ | 16/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 1 failure</summary>

- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
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
| additionalProperties | ⚠️ | 15/16 |
| default | ✅ | 7/7 |
| dependencies | ⚠️ | 17/18 |
| disallow | ✅ | 9/9 |
| divisibleBy | ✅ | 8/8 |
| enum | ✅ | 16/16 |
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
| properties | ✅ | 15/15 |
| ref | ❌ | 0/2 |
| refRemote | ⚠️ | 6/8 |
| required | ✅ | 4/4 |
| type | ⚠️ | 73/80 |
| uniqueItems | ✅ | 62/62 |

### Failures

<details>
<summary>additionalProperties - 1 failure</summary>

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
<summary>ref - 2 failures</summary>

- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`

</details>

<details>
<summary>refRemote - 2 failures</summary>

- **change resolution scope**
  - Test: changed scope ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`
- **change resolution scope**
  - Test: changed scope ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`

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

## draft4

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 17/17 |
| additionalProperties | ⚠️ | 15/16 |
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
| maximum | ⚠️ | 13/14 |
| minItems | ✅ | 4/4 |
| minLength | ✅ | 5/5 |
| minProperties | ✅ | 6/6 |
| minimum | ⚠️ | 16/17 |
| multipleOf | ✅ | 10/10 |
| not | ✅ | 20/20 |
| oneOf | ✅ | 23/23 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 18/18 |
| properties | ✅ | 24/24 |
| ref | ❌ | 0/8 |
| refRemote | ❌ | 0/9 |
| required | ✅ | 15/15 |
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
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
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
<summary>ref - 8 failures</summary>

- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "foo.json": failed to read foo.json: open foo.json: no such file or directory`
- **Location-independent identifier with base URI change in subschema**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "http://localhost:1234/nested.json": failed to read $CACHE/json-schema-test-suite-Test-JSON-Schema-Acceptance-1.035/remotes/nested.json: open $CACHE/json-schema-test-suite-Test-JSON-Schema-Acceptance-1.035/remotes/nested.json: no such file or directory`
- **Location-independent identifier with base URI change in subschema**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "http://localhost:1234/nested.json": failed to read $CACHE/json-schema-test-suite-Test-JSON-Schema-Acceptance-1.035/remotes/nested.json: open $CACHE/json-schema-test-suite-Test-JSON-Schema-Acceptance-1.035/remotes/nested.json: no such file or directory`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "tree": failed to read tree: open tree: no such file or directory`
- **Recursive references between schemas**
  - Test: valid tree
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "tree": failed to read tree: open tree: no such file or directory`
- **id must be resolved against nearest parent, not just immediate parent**
  - Test: non-number is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "http://example.com/b/d.json": failed to fetch http://example.com/b/d.json: status 404`
- **id must be resolved against nearest parent, not just immediate parent**
  - Test: number is valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "http://example.com/b/d.json": failed to fetch http://example.com/b/d.json: status 404`

</details>

<details>
<summary>refRemote - 9 failures</summary>

- **base URI change**
  - Test: base URI change ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`
- **base URI change**
  - Test: base URI change ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "folderInteger.json": failed to read folderInteger.json: open folderInteger.json: no such file or directory`
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
  - Test: null is valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "draft4/name.json": failed to read draft4/name.json: open draft4/name.json: no such file or directory`
- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "draft4/name.json": failed to read draft4/name.json: open draft4/name.json: no such file or directory`
- **root ref in remote ref**
  - Test: string is valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "draft4/name.json": failed to read draft4/name.json: open draft4/name.json: no such file or directory`

</details>

## draft6

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ⚠️ | 15/16 |
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
| properties | ✅ | 28/28 |
| propertyNames | ✅ | 20/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ✅ | 16/16 |
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
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

## draft7

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ⚠️ | 15/16 |
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
| properties | ✅ | 28/28 |
| propertyNames | ✅ | 20/20 |
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
| required | ✅ | 16/16 |
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
<summary>definitions - 1 failure</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `true`

</details>

## v1

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 23/24 |
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
| dependentSchemas | ✅ | 20/20 |
| dynamicRef | ✅ | 0/0 |
| enum | ✅ | 45/45 |
| exclusiveMaximum | ✅ | 4/4 |
| exclusiveMinimum | ✅ | 4/4 |
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
| properties | ✅ | 28/28 |
| propertyDependencies | ⚠️ | 17/21 |
| propertyNames | ✅ | 10/10 |
| ref | ❌ | 0/3 |
| refRemote | ✅ | 0/0 |
| required | ✅ | 16/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ✅ | 0/0 |
| uniqueItems | ✅ | 69/69 |

### Failures

<details>
<summary>additionalProperties - 1 failure</summary>

- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>defs - 2 failures</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`
- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`

</details>

<details>
<summary>propertyDependencies - 4 failures</summary>

- **multiple options selects the right one**
  - Test: bar with fewer than 2 properties is invalid
  - Expected: `invalid`, Got: `true`
- **multiple options selects the right one**
  - Test: bar with more than 2 properties is invalid
  - Expected: `invalid`, Got: `true`
- **multiple options selects the right one**
  - Test: baz with other properties is invalid
  - Expected: `invalid`, Got: `true`
- **multiple options selects the right one**
  - Test: quux is disallowed
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>ref - 3 failures</summary>

- **URN base URI with f-component**
  - Test: is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`

</details>

