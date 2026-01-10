# @xschemadev/zod Compliance Report

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2019-09 | 857 | 2 | 0 | 99.8% |
| draft2020-12 | 869 | 2 | 0 | 99.8% |
| draft3 | 396 | 11 | 0 | 97.3% |
| draft4 | 547 | 4 | 0 | 99.3% |
| draft6 | 737 | 2 | 0 | 99.7% |
| draft7 | 813 | 2 | 0 | 99.8% |

## Badges

![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-99.8%25-brightgreen)
![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-99.8%25-brightgreen)
![draft3](https://img.shields.io/badge/draft3%20compliance-97.3%25-brightgreen)
![draft4](https://img.shields.io/badge/draft4%20compliance-99.3%25-brightgreen)
![draft6](https://img.shields.io/badge/draft6%20compliance-99.7%25-brightgreen)
![draft7](https://img.shields.io/badge/draft7%20compliance-99.8%25-brightgreen)

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
| ref | ✅ | 0/0 |
| refRemote | ✅ | 8/8 |
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
| ref | ✅ | 0/0 |
| refRemote | ✅ | 0/0 |
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

