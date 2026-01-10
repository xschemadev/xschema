# @xschemadev/valibot Compliance Report

Generated: 2026-01-09T00:59:49Z

## Summary

| Draft | Passed | Failed | Skipped | Coverage |
| ----- | ------ | ------ | ------- | -------- |
| draft2020-12 | 728 | 302 | 0 | 70.7% |
| draft2019-09 | 721 | 306 | 0 | 70.2% |
| draft7 | 619 | 253 | 0 | 71.0% |
| draft6 | 549 | 243 | 0 | 69.3% |
| draft4 | 343 | 256 | 0 | 57.3% |
| draft3 | 287 | 141 | 0 | 67.1% |
| v1 | 487 | 392 | 0 | 55.4% |

## Badges

![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-70.7%25-red)
![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-70.2%25-red)
![draft7](https://img.shields.io/badge/draft7%20compliance-71.0%25-red)
![draft6](https://img.shields.io/badge/draft6%20compliance-69.3%25-red)
![draft4](https://img.shields.io/badge/draft4%20compliance-57.3%25-red)
![draft3](https://img.shields.io/badge/draft3%20compliance-67.1%25-red)
![v1](https://img.shields.io/badge/v1%20compliance-55.4%25-red)

## draft2020-12

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 12/21 |
| allOf | ⚠️ | 11/30 |
| anchor | ✅ | 0/0 |
| anyOf | ⚠️ | 13/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 33/54 |
| contains | ⚠️ | 17/21 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ⚠️ | 2/20 |
| dependentSchemas | ⚠️ | 8/20 |
| dynamicRef | ⚠️ | 9/19 |
| enum | ⚠️ | 23/45 |
| exclusiveMaximum | ⚠️ | 3/4 |
| exclusiveMinimum | ⚠️ | 3/4 |
| format | ✅ | 133/133 |
| if-then-else | ⚠️ | 20/30 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 18/29 |
| maxContains | ⚠️ | 8/12 |
| maxItems | ⚠️ | 5/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ⚠️ | 5/10 |
| maximum | ⚠️ | 7/8 |
| minContains | ⚠️ | 19/28 |
| minItems | ⚠️ | 5/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ⚠️ | 2/8 |
| minimum | ⚠️ | 9/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ⚠️ | 24/38 |
| oneOf | ⚠️ | 13/27 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 11/23 |
| prefixItems | ⚠️ | 6/11 |
| properties | ⚠️ | 17/28 |
| propertyNames | ⚠️ | 13/20 |
| ref | ⚠️ | 28/35 |
| refRemote | ⚠️ | 16/17 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 69/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 24/46 |
| uniqueItems | ⚠️ | 65/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalProperties - 9 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
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
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
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
<summary>allOf - 19 failures</summary>

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
- **allOf with boolean schemas, some false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with boolean schemas, all false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the first empty schema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the last empty schema**
  - Test: string is invalid
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
<summary>anyOf - 5 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf with boolean schemas, all false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>const - 21 failures</summary>

- **const validation**
  - Test: same value is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `false`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `false`
- **const with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **const with true does not match 1**
  - Test: true is valid
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
- **const with 0 does not match other zero-like types**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **const with 0 does not match other zero-like types**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **const with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **const with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **const with -2.0 matches integer and float types**
  - Test: integer -2 is valid
  - Expected: `valid`, Got: `false`
- **const with -2.0 matches integer and float types**
  - Test: float -2.0 is valid
  - Expected: `valid`, Got: `false`
- **float and integers are equal up to 64-bit representation limits**
  - Test: integer is valid
  - Expected: `valid`, Got: `false`
- **float and integers are equal up to 64-bit representation limits**
  - Test: float is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation but different codepoint**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation, but different number of codepoints**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>contains - 4 failures</summary>

- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with const keyword**
  - Test: array with item 5 is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with const keyword**
  - Test: array with two items 5 is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>dependentRequired - 18 failures</summary>

- **single dependency**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: nondependant
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: with dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2945116195.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **empty dependents**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
- **multiple dependents required**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: nondependants
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3128185357.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: CRLF
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property foo
           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1219795437.ts:8:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property foo
           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1219795437.ts:8:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property foo
           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1219795437.ts:8:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property foo
           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1219795437.ts:8:8

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>dependentSchemas - 12 failures</summary>

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
<summary>dynamicRef - 10 failures</summary>

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
<summary>enum - 22 failures</summary>

- **simple enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `true`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: both properties are valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: missing optional property is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 1 is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 2 is valid
  - Expected: `valid`, Got: `false`
- **enum with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with true does not match 1**
  - Test: true is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
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
<summary>items - 11 failures</summary>

- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
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
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`
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
<summary>maxContains - 4 failures</summary>

- **maxContains with contains**
  - Test: all elements match, valid maxContains
  - Expected: `valid`, Got: `false`
- **maxContains with contains**
  - Test: some elements match, valid maxContains
  - Expected: `valid`, Got: `false`
- **maxContains with contains, value with a decimal**
  - Test: one element matches, valid maxContains
  - Expected: `valid`, Got: `false`
- **minContains < maxContains**
  - Test: minContains < actual < maxContains
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
<summary>maxProperties - 5 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>maximum - 1 failure</summary>

- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minContains - 9 failures</summary>

- **minContains=1 with contains**
  - Test: single element matches, valid minContains
  - Expected: `valid`, Got: `false`
- **minContains=1 with contains**
  - Test: some elements match, valid minContains
  - Expected: `valid`, Got: `false`
- **minContains=1 with contains**
  - Test: all elements match, valid minContains
  - Expected: `valid`, Got: `false`
- **minContains=2 with contains**
  - Test: all elements match, valid minContains (exactly as needed)
  - Expected: `valid`, Got: `false`
- **minContains=2 with contains**
  - Test: all elements match, valid minContains (more than needed)
  - Expected: `valid`, Got: `false`
- **minContains=2 with contains**
  - Test: some elements match, valid minContains
  - Expected: `valid`, Got: `false`
- **minContains=2 with contains with a decimal value**
  - Test: both elements match, valid minContains
  - Expected: `valid`, Got: `false`
- **maxContains = minContains**
  - Test: all elements match, valid maxContains and minContains
  - Expected: `valid`, Got: `false`
- **minContains = 0 with maxContains**
  - Test: too many
  - Expected: `invalid`, Got: `true`

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
<summary>minProperties - 6 failures</summary>

- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `false`
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
  - Test: longer is valid
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
<summary>not - 14 failures</summary>

- **not**
  - Test: allowed
  - Expected: `valid`, Got: `false`
- **not multiple types**
  - Test: valid
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: match
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: string is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: boolean true is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: boolean false is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: null is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: object is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: empty object is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: array is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`
- **double negation**
  - Test: any value is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 14 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: mismatch base schema
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
- **oneOf with boolean schemas, all false**
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
<summary>pattern - 6 failures</summary>

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
<summary>patternProperties - 12 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>prefixItems - 5 failures</summary>

- **a schema given for prefixItems**
  - Test: incomplete array of items
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

</details>

<details>
<summary>properties - 11 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`
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
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>propertyNames - 7 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
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
<summary>ref - 7 failures</summary>

- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `false`
- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>refRemote - 1 failure</summary>

- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
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
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>type - 11 failures</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an object is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an array is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a boolean is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 22 failures</summary>

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
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 4 failures</summary>

- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
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
  - Test: no validation: invalid number, but it still validates
  - Expected: `valid`, Got: `false`

</details>

## draft2019-09

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 13/19 |
| additionalProperties | ⚠️ | 12/21 |
| allOf | ⚠️ | 11/30 |
| anchor | ✅ | 0/0 |
| anyOf | ⚠️ | 13/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 33/54 |
| contains | ⚠️ | 17/21 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ✅ | 0/0 |
| dependentRequired | ⚠️ | 2/20 |
| dependentSchemas | ⚠️ | 8/20 |
| enum | ⚠️ | 23/45 |
| exclusiveMaximum | ⚠️ | 3/4 |
| exclusiveMinimum | ⚠️ | 3/4 |
| format | ✅ | 114/114 |
| if-then-else | ⚠️ | 20/30 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 16/28 |
| maxContains | ⚠️ | 8/12 |
| maxItems | ⚠️ | 5/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ⚠️ | 5/10 |
| maximum | ⚠️ | 7/8 |
| minContains | ⚠️ | 19/28 |
| minItems | ⚠️ | 5/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ⚠️ | 2/8 |
| minimum | ⚠️ | 9/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ⚠️ | 24/38 |
| oneOf | ⚠️ | 13/27 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 11/23 |
| properties | ⚠️ | 17/28 |
| propertyNames | ⚠️ | 13/20 |
| recursiveRef | ⚠️ | 18/30 |
| ref | ⚠️ | 28/35 |
| refRemote | ⚠️ | 16/17 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 69/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 22/44 |
| uniqueItems | ⚠️ | 65/69 |
| vocabulary | ⚠️ | 4/5 |

### Failures

<details>
<summary>additionalItems - 6 failures</summary>

- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
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
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
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
<summary>allOf - 19 failures</summary>

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
- **allOf with boolean schemas, some false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with boolean schemas, all false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the first empty schema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the last empty schema**
  - Test: string is invalid
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
<summary>anyOf - 5 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf with boolean schemas, all false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>const - 21 failures</summary>

- **const validation**
  - Test: same value is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `false`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `false`
- **const with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **const with true does not match 1**
  - Test: true is valid
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
- **const with 0 does not match other zero-like types**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **const with 0 does not match other zero-like types**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **const with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **const with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **const with -2.0 matches integer and float types**
  - Test: integer -2 is valid
  - Expected: `valid`, Got: `false`
- **const with -2.0 matches integer and float types**
  - Test: float -2.0 is valid
  - Expected: `valid`, Got: `false`
- **float and integers are equal up to 64-bit representation limits**
  - Test: integer is valid
  - Expected: `valid`, Got: `false`
- **float and integers are equal up to 64-bit representation limits**
  - Test: float is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation but different codepoint**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation, but different number of codepoints**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>contains - 4 failures</summary>

- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with const keyword**
  - Test: array with item 5 is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with const keyword**
  - Test: array with two items 5 is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>dependentRequired - 18 failures</summary>

- **single dependency**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: nondependant
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: with dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2686447966.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **empty dependents**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
- **multiple dependents required**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: nondependants
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1657000819.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: CRLF
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property foo
           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1503847160.ts:8:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property foo
           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1503847160.ts:8:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property foo
           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1503847160.ts:8:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property foo
           ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1503847160.ts:8:8

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>dependentSchemas - 12 failures</summary>

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
<summary>enum - 22 failures</summary>

- **simple enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `true`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: both properties are valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: missing optional property is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 1 is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 2 is valid
  - Expected: `valid`, Got: `false`
- **enum with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with true does not match 1**
  - Test: true is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
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
<summary>items - 12 failures</summary>

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
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: empty array is valid
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
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxContains - 4 failures</summary>

- **maxContains with contains**
  - Test: all elements match, valid maxContains
  - Expected: `valid`, Got: `false`
- **maxContains with contains**
  - Test: some elements match, valid maxContains
  - Expected: `valid`, Got: `false`
- **maxContains with contains, value with a decimal**
  - Test: one element matches, valid maxContains
  - Expected: `valid`, Got: `false`
- **minContains < maxContains**
  - Test: minContains < actual < maxContains
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
<summary>maxProperties - 5 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>maximum - 1 failure</summary>

- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minContains - 9 failures</summary>

- **minContains=1 with contains**
  - Test: single element matches, valid minContains
  - Expected: `valid`, Got: `false`
- **minContains=1 with contains**
  - Test: some elements match, valid minContains
  - Expected: `valid`, Got: `false`
- **minContains=1 with contains**
  - Test: all elements match, valid minContains
  - Expected: `valid`, Got: `false`
- **minContains=2 with contains**
  - Test: all elements match, valid minContains (exactly as needed)
  - Expected: `valid`, Got: `false`
- **minContains=2 with contains**
  - Test: all elements match, valid minContains (more than needed)
  - Expected: `valid`, Got: `false`
- **minContains=2 with contains**
  - Test: some elements match, valid minContains
  - Expected: `valid`, Got: `false`
- **minContains=2 with contains with a decimal value**
  - Test: both elements match, valid minContains
  - Expected: `valid`, Got: `false`
- **maxContains = minContains**
  - Test: all elements match, valid maxContains and minContains
  - Expected: `valid`, Got: `false`
- **minContains = 0 with maxContains**
  - Test: too many
  - Expected: `invalid`, Got: `true`

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
<summary>minProperties - 6 failures</summary>

- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `false`
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
  - Test: longer is valid
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
<summary>not - 14 failures</summary>

- **not**
  - Test: allowed
  - Expected: `valid`, Got: `false`
- **not multiple types**
  - Test: valid
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: match
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: string is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: boolean true is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: boolean false is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: null is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: object is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: empty object is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: array is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`
- **double negation**
  - Test: any value is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 14 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: mismatch base schema
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
- **oneOf with boolean schemas, all false**
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
<summary>pattern - 6 failures</summary>

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
<summary>patternProperties - 12 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>properties - 11 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`
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
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>propertyNames - 7 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
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
<summary>ref - 7 failures</summary>

- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `false`
- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref valid, maxItems invalid
  - Expected: `invalid`, Got: `true`
- **ref applies alongside sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>refRemote - 1 failure</summary>

- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
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
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>type - 11 failures</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an object is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an array is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a boolean is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 22 failures</summary>

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
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores null
  - Expected: `valid`, Got: `false`
- **Evaluated properties collection needs to consider instance location**
  - Test: with an unevaluated property that exists at another location
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 4 failures</summary>

- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
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
  - Test: no validation: invalid number, but it still validates
  - Expected: `valid`, Got: `false`

</details>

## draft7

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 13/19 |
| additionalProperties | ⚠️ | 10/16 |
| allOf | ⚠️ | 11/30 |
| anyOf | ⚠️ | 13/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 33/54 |
| contains | ⚠️ | 17/21 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 8/36 |
| enum | ⚠️ | 23/45 |
| exclusiveMaximum | ⚠️ | 3/4 |
| exclusiveMinimum | ⚠️ | 3/4 |
| format | ✅ | 102/102 |
| if-then-else | ⚠️ | 20/30 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 16/28 |
| maxItems | ⚠️ | 5/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ⚠️ | 5/10 |
| maximum | ⚠️ | 7/8 |
| minItems | ⚠️ | 5/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ⚠️ | 2/8 |
| minimum | ⚠️ | 9/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ⚠️ | 24/38 |
| oneOf | ⚠️ | 13/27 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 11/23 |
| properties | ⚠️ | 17/28 |
| propertyNames | ⚠️ | 13/20 |
| ref | ⚠️ | 30/36 |
| refRemote | ⚠️ | 16/17 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 69/80 |
| uniqueItems | ⚠️ | 65/69 |

### Failures

<details>
<summary>additionalItems - 6 failures</summary>

- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
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
<summary>additionalProperties - 6 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
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
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 19 failures</summary>

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
- **allOf with boolean schemas, some false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with boolean schemas, all false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the first empty schema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the last empty schema**
  - Test: string is invalid
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
<summary>anyOf - 5 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf with boolean schemas, all false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>const - 21 failures</summary>

- **const validation**
  - Test: same value is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `false`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `false`
- **const with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **const with true does not match 1**
  - Test: true is valid
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
- **const with 0 does not match other zero-like types**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **const with 0 does not match other zero-like types**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **const with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **const with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **const with -2.0 matches integer and float types**
  - Test: integer -2 is valid
  - Expected: `valid`, Got: `false`
- **const with -2.0 matches integer and float types**
  - Test: float -2.0 is valid
  - Expected: `valid`, Got: `false`
- **float and integers are equal up to 64-bit representation limits**
  - Test: integer is valid
  - Expected: `valid`, Got: `false`
- **float and integers are equal up to 64-bit representation limits**
  - Test: float is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation but different codepoint**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation, but different number of codepoints**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>contains - 4 failures</summary>

- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with const keyword**
  - Test: array with item 5 is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with const keyword**
  - Test: array with two items 5 is valid
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
<summary>dependencies - 28 failures</summary>

- **dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: nondependant
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: with dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3276726842.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies with empty array**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
- **multiple dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: nondependants
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854705441.ts:8:46

Bun v1.3.5 (Linux x64)
`
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
  - Test: valid object 1
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4271889835.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: valid object 2
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4271889835.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: valid object 3
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4271889835.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 1
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4271889835.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 2
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4271889835.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4271889835.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 4
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4271889835.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 22 failures</summary>

- **simple enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `true`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: both properties are valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: missing optional property is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 1 is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 2 is valid
  - Expected: `valid`, Got: `false`
- **enum with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with true does not match 1**
  - Test: true is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
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
<summary>items - 12 failures</summary>

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
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: empty array is valid
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
- **items and subitems**
  - Test: fewer items is valid
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
<summary>maxProperties - 5 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>minProperties - 6 failures</summary>

- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `false`
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
  - Test: longer is valid
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
<summary>not - 14 failures</summary>

- **not**
  - Test: allowed
  - Expected: `valid`, Got: `false`
- **not multiple types**
  - Test: valid
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: match
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: string is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: boolean true is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: boolean false is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: null is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: object is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: empty object is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: array is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`
- **double negation**
  - Test: any value is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 14 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: mismatch base schema
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
- **oneOf with boolean schemas, all false**
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
<summary>pattern - 6 failures</summary>

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
<summary>patternProperties - 12 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>properties - 11 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`
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
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>propertyNames - 7 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
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
<summary>ref - 6 failures</summary>

- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `false`
- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>refRemote - 1 failure</summary>

- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
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
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>type - 11 failures</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an object is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an array is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a boolean is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 4 failures</summary>

- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
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
| additionalProperties | ⚠️ | 10/16 |
| allOf | ⚠️ | 11/30 |
| anyOf | ⚠️ | 13/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 33/54 |
| contains | ⚠️ | 15/19 |
| default | ✅ | 7/7 |
| definitions | ⚠️ | 1/2 |
| dependencies | ⚠️ | 8/36 |
| enum | ⚠️ | 23/45 |
| exclusiveMaximum | ⚠️ | 3/4 |
| exclusiveMinimum | ⚠️ | 3/4 |
| format | ✅ | 54/54 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 16/28 |
| maxItems | ⚠️ | 5/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ⚠️ | 5/10 |
| maximum | ⚠️ | 7/8 |
| minItems | ⚠️ | 5/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ⚠️ | 2/8 |
| minimum | ⚠️ | 9/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ⚠️ | 24/38 |
| oneOf | ⚠️ | 13/27 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 11/23 |
| properties | ⚠️ | 17/28 |
| propertyNames | ⚠️ | 13/20 |
| ref | ⚠️ | 30/36 |
| refRemote | ⚠️ | 16/17 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 69/80 |
| uniqueItems | ⚠️ | 65/69 |

### Failures

<details>
<summary>additionalItems - 6 failures</summary>

- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
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
<summary>additionalProperties - 6 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
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
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 19 failures</summary>

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
- **allOf with boolean schemas, some false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with boolean schemas, all false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the first empty schema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the last empty schema**
  - Test: string is invalid
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
<summary>anyOf - 5 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf with boolean schemas, all false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>const - 21 failures</summary>

- **const validation**
  - Test: same value is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `false`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `false`
- **const with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **const with true does not match 1**
  - Test: true is valid
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
- **const with 0 does not match other zero-like types**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **const with 0 does not match other zero-like types**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **const with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **const with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **const with -2.0 matches integer and float types**
  - Test: integer -2 is valid
  - Expected: `valid`, Got: `false`
- **const with -2.0 matches integer and float types**
  - Test: float -2.0 is valid
  - Expected: `valid`, Got: `false`
- **float and integers are equal up to 64-bit representation limits**
  - Test: integer is valid
  - Expected: `valid`, Got: `false`
- **float and integers are equal up to 64-bit representation limits**
  - Test: float is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation but different codepoint**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation, but different number of codepoints**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>contains - 4 failures</summary>

- **contains keyword validation**
  - Test: not array is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with const keyword**
  - Test: array with item 5 is valid
  - Expected: `valid`, Got: `false`
- **contains keyword with const keyword**
  - Test: array with two items 5 is valid
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
<summary>dependencies - 28 failures</summary>

- **dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: nondependant
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: with dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property bar requires "foo""));
                                   ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:32

8 | 				}, "Property bar requires "foo""));
                                      ^
error: Expected ")" but found """"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:35

8 | 				}, "Property bar requires "foo""));
                                        ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:37

8 | 				}, "Property bar requires "foo""));
                                         ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3098664549.ts:8:38

Bun v1.3.5 (Linux x64)
`
- **dependencies with empty array**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
- **multiple dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: nondependants
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:46

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 8 | 				}, "Property quux requires "foo", "bar""));
                                    ^
error: Expected ")" but found "foo"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:33

8 | 				}, "Property quux requires "foo", "bar""));
                                       ^
error: Expected ")" but found "", ""
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:36

8 | 				}, "Property quux requires "foo", "bar""));
                                           ^
error: Expected ";" but found "bar"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:40

8 | 				}, "Property quux requires "foo", "bar""));
                                                ^
error: Expected ";" but found ")"
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:45

8 | 				}, "Property quux requires "foo", "bar""));
                                                 ^
error: Unexpected )
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2073624455.ts:8:46

Bun v1.3.5 (Linux x64)
`
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
  - Test: valid object 1
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3021895774.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: valid object 2
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3021895774.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: valid object 3
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3021895774.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 1
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3021895774.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 2
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3021895774.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3021895774.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 4
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 14 | 				}, "Property foo
            ^
error: Unterminated string literal
    at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3021895774.ts:14:8

Bun v1.3.5 (Linux x64)
`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `true`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>enum - 22 failures</summary>

- **simple enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `true`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: both properties are valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: missing optional property is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 1 is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 2 is valid
  - Expected: `valid`, Got: `false`
- **enum with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with true does not match 1**
  - Test: true is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
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
<summary>infinite-loop-detection - 1 failure</summary>

- **evaluating the same schema location against the same data location twice is not a sign of an infinite loop**
  - Test: failing case
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>items - 12 failures</summary>

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
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schema (false)**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `true`
- **items with boolean schemas**
  - Test: array with one item is valid
  - Expected: `valid`, Got: `false`
- **items with boolean schemas**
  - Test: empty array is valid
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
- **items and subitems**
  - Test: fewer items is valid
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
<summary>maxProperties - 5 failures</summary>

- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>minProperties - 6 failures</summary>

- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `false`
- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `false`
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
  - Test: longer is valid
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
<summary>not - 14 failures</summary>

- **not**
  - Test: allowed
  - Expected: `valid`, Got: `false`
- **not multiple types**
  - Test: valid
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: match
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: string is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: boolean true is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: boolean false is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: null is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: object is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: empty object is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: array is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`
- **double negation**
  - Test: any value is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 14 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: mismatch base schema
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
- **oneOf with boolean schemas, all false**
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
<summary>pattern - 6 failures</summary>

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
<summary>patternProperties - 12 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: a single invalid match is invalid
  - Expected: `invalid`, Got: `true`
- **patternProperties validates properties matching a regex**
  - Test: multiple invalid matches is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>properties - 11 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `true`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `true`
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
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>propertyNames - 7 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
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
<summary>ref - 6 failures</summary>

- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `false`
- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `true`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `true`
- **ref overrides any sibling keywords**
  - Test: ref invalid
  - Expected: `invalid`, Got: `true`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>refRemote - 1 failure</summary>

- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
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
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>type - 11 failures</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an object is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an array is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a boolean is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 4 failures</summary>

- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `true`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
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
| additionalItems | ⚠️ | 11/17 |
| additionalProperties | ⚠️ | 10/16 |
| allOf | ⚠️ | 10/27 |
| anyOf | ⚠️ | 11/15 |
| default | ✅ | 7/7 |
| definitions | ❌ | 0/2 |
| dependencies | ❌ | 0/29 |
| enum | ⚠️ | 25/49 |
| format | ✅ | 36/36 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 12/21 |
| maxItems | ❌ | 0/4 |
| maxLength | ⚠️ | 4/5 |
| maxProperties | ❌ | 0/8 |
| maximum | ⚠️ | 11/14 |
| minItems | ❌ | 0/4 |
| minLength | ⚠️ | 4/5 |
| minProperties | ❌ | 0/6 |
| minimum | ⚠️ | 13/17 |
| multipleOf | ⚠️ | 9/10 |
| not | ⚠️ | 15/20 |
| oneOf | ⚠️ | 12/23 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 16/18 |
| properties | ⚠️ | 8/24 |
| ref | ⚠️ | 20/33 |
| refRemote | ⚠️ | 6/15 |
| required | ⚠️ | 4/15 |
| type | ⚠️ | 68/79 |
| uniqueItems | ⚠️ | 27/69 |

### Failures

<details>
<summary>additionalItems - 6 failures</summary>

- **additionalItems as schema**
  - Test: additional items do not match schema
  - Expected: `invalid`, Got: `true`
- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
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
<summary>additionalProperties - 6 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `true`
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
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>allOf - 17 failures</summary>

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
- **allOf with the first empty schema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the last empty schema**
  - Test: string is invalid
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
<summary>definitions - 2 failures</summary>

- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "$schema": v.optional(v.string()), "additionalItems": v.optional(v.union([v.any()])), "additionalProperties": v.optional(v.union([v.any()])), "allOf": v.optional(v.array(v.any()).pipe(v.minLength(1))), "anyOf": v.optional(v.array(v.any()).pipe(v.minLength(1))), "default": v.optional(v.any()), "definitions": v.optional(v.record(v.string(), v.any())), "dependencies": v.optional(v.record(v.string(), v.union([v.any()]))), "description": v.optional(v.string()), "enum": v.optional(v.array(v.any()).pipe(v.minLength(1), v.check((arr) => {

TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3330576141.ts:3:211
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "$schema": v.optional(v.string()), "additionalItems": v.optional(v.union([v.any()])), "additionalProperties": v.optional(v.union([v.any()])), "allOf": v.optional(v.array(v.any()).pipe(v.minLength(1))), "anyOf": v.optional(v.array(v.any()).pipe(v.minLength(1))), "default": v.optional(v.any()), "definitions": v.optional(v.record(v.string(), v.any())), "dependencies": v.optional(v.record(v.string(), v.union([v.any()]))), "description": v.optional(v.string()), "enum": v.optional(v.array(v.any()).pipe(v.minLength(1), v.check((arr) => {

TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3330576141.ts:3:211
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>dependencies - 29 failures</summary>

- **dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3734815486.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: nondependant
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3734815486.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: with dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3734815486.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3734815486.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3734815486.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3734815486.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3734815486.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1817642068.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: nondependants
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1817642068.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1817642068.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1817642068.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1817642068.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1817642068.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies subschema**
  - Test: valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2765475523.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies subschema**
  - Test: no dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2765475523.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies subschema**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2765475523.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies subschema**
  - Test: wrong type other
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2765475523.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies subschema**
  - Test: wrong type both
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2765475523.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: valid object 1
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2762743360.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: valid object 2
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2762743360.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: valid object 3
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2762743360.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 1
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2762743360.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 2
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2762743360.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 3
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2762743360.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: invalid object 4
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2762743360.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                 ^
TypeError: v.looseObject({ foo: v.optional(v.any()) }).pipe is not a function. (In 'v.looseObject({ foo: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.strictObject({ bar: v.optional(v.any()) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({ foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2298278053.ts:3:62
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependent subschema incompatible with root**
  - Test: matches dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                 ^
TypeError: v.looseObject({ foo: v.optional(v.any()) }).pipe is not a function. (In 'v.looseObject({ foo: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.strictObject({ bar: v.optional(v.any()) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({ foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2298278053.ts:3:62
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                 ^
TypeError: v.looseObject({ foo: v.optional(v.any()) }).pipe is not a function. (In 'v.looseObject({ foo: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.strictObject({ bar: v.optional(v.any()) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({ foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2298278053.ts:3:62
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependent subschema incompatible with root**
  - Test: no dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                 ^
TypeError: v.looseObject({ foo: v.optional(v.any()) }).pipe is not a function. (In 'v.looseObject({ foo: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.strictObject({ bar: v.optional(v.any()) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({ foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2298278053.ts:3:62
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>enum - 24 failures</summary>

- **simple enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `true`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: both properties are valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: missing optional property is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 1 is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 2 is valid
  - Expected: `valid`, Got: `false`
- **enum with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with true does not match 1**
  - Test: true is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation but different codepoint**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation, but different number of codepoints**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`

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
- **an array of schemas for items**
  - Test: incomplete array of items
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: empty array
  - Expected: `valid`, Got: `false`
- **an array of schemas for items**
  - Test: JavaScript pseudo-array is valid
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
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 4 failures</summary>

- **maxItems validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2393486046.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2393486046.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2393486046.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2393486046.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxProperties - 8 failures</summary>

- **maxProperties validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2790540017.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2790540017.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2790540017.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2790540017.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2790540017.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2790540017.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties = 0 means the object is empty**
  - Test: no properties is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 0, "Object must have at most 0 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 0, "Object must have at most 0 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2439808663.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 0, "Object must have at most 0 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 0, "Object must have at most 0 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2439808663.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>maximum - 3 failures</summary>

- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **maximum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 4 failures</summary>

- **minItems validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3425990279.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3425990279.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3425990279.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3425990279.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minProperties - 6 failures</summary>

- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4007304194.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4007304194.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4007304194.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4007304194.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4007304194.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4007304194.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>minimum - 4 failures</summary>

- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **minimum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
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
<summary>not - 5 failures</summary>

- **not**
  - Test: allowed
  - Expected: `valid`, Got: `false`
- **not multiple types**
  - Test: valid
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: match
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **double negation**
  - Test: any value is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 11 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: mismatch base schema
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
<summary>pattern - 6 failures</summary>

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
<summary>patternProperties - 2 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 16 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3885738694.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3885738694.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3885738694.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3885738694.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3885738694.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3885738694.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty validates others
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3885738694.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3885738694.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
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
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>ref - 13 failures</summary>

- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `false`
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
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "$schema": v.optional(v.string()), "additionalItems": v.optional(v.union([v.any()])), "additionalProperties": v.optional(v.union([v.any()])), "allOf": v.optional(v.array(v.any()).pipe(v.minLength(1))), "anyOf": v.optional(v.array(v.any()).pipe(v.minLength(1))), "default": v.optional(v.any()), "definitions": v.optional(v.record(v.string(), v.any())), "dependencies": v.optional(v.record(v.string(), v.union([v.any()]))), "description": v.optional(v.string()), "enum": v.optional(v.array(v.any()).pipe(v.minLength(1), v.check((arr) => {

TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2387805056.ts:3:211
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "$schema": v.optional(v.string()), "additionalItems": v.optional(v.union([v.any()])), "additionalProperties": v.optional(v.union([v.any()])), "allOf": v.optional(v.array(v.any()).pipe(v.minLength(1))), "anyOf": v.optional(v.array(v.any()).pipe(v.minLength(1))), "default": v.optional(v.any()), "definitions": v.optional(v.record(v.string(), v.any())), "dependencies": v.optional(v.record(v.string(), v.union([v.any()]))), "description": v.optional(v.string()), "enum": v.optional(v.array(v.any()).pipe(v.minLength(1), v.check((arr) => {

TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2387805056.ts:3:211
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
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
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>type - 11 failures</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an object is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an array is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a boolean is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>uniqueItems - 42 failures</summary>

- **uniqueItems validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of strings is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: [1] and [true] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: [0] and [false] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: nested [1] and [true] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: nested [0] and [false] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: different objects are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: {"a": false} and {"a": 0} are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: {"a": true} and {"a": 1} are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-402407004.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4154162493.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4154162493.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4154162493.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4154162493.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4154162493.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4154162493.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4154162493.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4154162493.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3632223244.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3632223244.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3632223244.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3632223244.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3632223244.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

## draft3

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ⚠️ | 11/14 |
| additionalProperties | ⚠️ | 9/16 |
| default | ✅ | 7/7 |
| dependencies | ❌ | 0/18 |
| disallow | ⚠️ | 5/9 |
| divisibleBy | ⚠️ | 7/8 |
| enum | ⚠️ | 9/16 |
| extends | ⚠️ | 3/10 |
| format | ✅ | 60/60 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 6/7 |
| maxItems | ❌ | 0/4 |
| maxLength | ⚠️ | 4/5 |
| maximum | ⚠️ | 11/14 |
| minItems | ❌ | 0/4 |
| minLength | ⚠️ | 4/5 |
| minimum | ⚠️ | 10/13 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 16/17 |
| properties | ⚠️ | 6/15 |
| ref | ⚠️ | 14/21 |
| refRemote | ⚠️ | 6/8 |
| required | ⚠️ | 3/4 |
| type | ⚠️ | 67/80 |
| uniqueItems | ⚠️ | 25/62 |

### Failures

<details>
<summary>additionalItems - 3 failures</summary>

- **array of items with no additionalItems permitted**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **additionalItems as false without items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **additionalItems with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>additionalProperties - 7 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1572304645.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1572304645.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1572304645.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1572304645.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1572304645.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1572304645.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties does not look in applicators**
  - Test: properties defined in extends are not examined
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>dependencies - 18 failures</summary>

- **dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.any(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1233146469.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: nondependant
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.any(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1233146469.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: with dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.any(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1233146469.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.any(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1233146469.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.any(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1233146469.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.any(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1233146469.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.any(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1233146469.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-494095977.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: nondependants
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-494095977.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-494095977.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-494095977.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-494095977.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-494095977.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies subschema**
  - Test: valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2936836502.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies subschema**
  - Test: no dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2936836502.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies subschema**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2936836502.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies subschema**
  - Test: wrong type other
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2936836502.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependencies subschema**
  - Test: wrong type both
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2936836502.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>disallow - 4 failures</summary>

- **disallow**
  - Test: allowed
  - Expected: `valid`, Got: `false`
- **multiple disallow**
  - Test: valid
  - Expected: `valid`, Got: `false`
- **multiple disallow subschema**
  - Test: match
  - Expected: `valid`, Got: `false`
- **multiple disallow subschema**
  - Test: other match
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>divisibleBy - 1 failure</summary>

- **by int**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>enum - 7 failures</summary>

- **simple enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `true`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: both properties are valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: missing optional property is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
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
<summary>items - 1 failure</summary>

- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxItems - 4 failures</summary>

- **maxItems validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3470219431.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3470219431.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3470219431.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3470219431.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maximum - 3 failures</summary>

- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **maximum validation (explicit false exclusivity)**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **exclusiveMaximum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>minItems - 4 failures</summary>

- **minItems validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2222440669.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2222440669.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2222440669.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2222440669.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minimum - 3 failures</summary>

- **minimum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`
- **exclusiveMinimum validation**
  - Test: boundary point is invalid
  - Expected: `invalid`, Got: `true`
- **minimum validation with signed integer**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>pattern - 6 failures</summary>

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
<summary>patternProperties - 1 failure</summary>

- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>properties - 9 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3376680493.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3376680493.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3376680493.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3376680493.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3376680493.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3376680493.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty validates others
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3376680493.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3376680493.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>ref - 7 failures</summary>

- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `false`
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
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "$ref": v.optional(v.pipe(v.string(), v.url())), "$schema": v.optional(v.pipe(v.string(), v.url())), "additionalItems": v.optional(v.union([v.any()])), "additionalProperties": v.optional(v.union([v.any()])), "default": v.optional(v.any()), "dependencies": v.optional(v.record(v.string(), v.union([v.any()]))), "description": v.optional(v.string()), "disallow": v.optional(v.union([v.any()])), "divisibleBy": v.optional(v.pipe(v.number(), v.minValue(0))), "enum": v.optional(v.array(v.any()).pipe(v.minLength(1), v.check((arr) => {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1), v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3530757780.ts:3:523
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "$ref": v.optional(v.pipe(v.string(), v.url())), "$schema": v.optional(v.pipe(v.string(), v.url())), "additionalItems": v.optional(v.union([v.any()])), "additionalProperties": v.optional(v.union([v.any()])), "default": v.optional(v.any()), "dependencies": v.optional(v.record(v.string(), v.union([v.any()]))), "description": v.optional(v.string()), "disallow": v.optional(v.union([v.any()])), "divisibleBy": v.optional(v.pipe(v.number(), v.minValue(0))), "enum": v.optional(v.array(v.any()).pipe(v.minLength(1), v.check((arr) => {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1), v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3530757780.ts:3:523
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
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
  - Test: present required property is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>type - 13 failures</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an object is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an array is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a boolean is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: null is invalid
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
<summary>uniqueItems - 37 failures</summary>

- **uniqueItems validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of strings is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: [1] and [true] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: [0] and [false] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: nested [1] and [true] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: nested [0] and [false] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: {"a": false} and {"a": 0} are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: {"a": true} and {"a": 1} are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2623840102.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3430275728.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3430275728.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3430275728.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3430275728.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3430275728.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3430275728.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3430275728.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3430275728.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2164889839.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2164889839.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2164889839.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2164889839.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2164889839.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

## v1

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ⚠️ | 11/21 |
| allOf | ⚠️ | 11/30 |
| anchor | ✅ | 0/0 |
| anyOf | ⚠️ | 13/18 |
| boolean_schema | ✅ | 18/18 |
| const | ⚠️ | 33/54 |
| contains | ❌ | 0/25 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ❌ | 0/2 |
| dependentRequired | ⚠️ | 2/20 |
| dependentSchemas | ❌ | 0/20 |
| dynamicRef | ⚠️ | 6/13 |
| enum | ⚠️ | 23/45 |
| exclusiveMaximum | ⚠️ | 3/4 |
| exclusiveMinimum | ⚠️ | 3/4 |
| if-then-else | ⚠️ | 18/26 |
| infinite-loop-detection | ⚠️ | 1/2 |
| items | ⚠️ | 19/29 |
| maxContains | ⚠️ | 2/12 |
| maxItems | ❌ | 0/6 |
| maxLength | ⚠️ | 6/7 |
| maxProperties | ❌ | 0/10 |
| maximum | ⚠️ | 7/8 |
| minContains | ⚠️ | 2/28 |
| minItems | ❌ | 0/6 |
| minLength | ⚠️ | 6/7 |
| minProperties | ❌ | 0/8 |
| minimum | ⚠️ | 9/11 |
| multipleOf | ⚠️ | 9/10 |
| not | ⚠️ | 24/38 |
| oneOf | ⚠️ | 13/27 |
| pattern | ⚠️ | 3/9 |
| patternProperties | ⚠️ | 21/23 |
| prefixItems | ⚠️ | 6/11 |
| properties | ⚠️ | 12/28 |
| propertyNames | ⚠️ | 6/10 |
| ref | ⚠️ | 30/38 |
| refRemote | ⚠️ | 16/17 |
| required | ⚠️ | 5/16 |
| type | ⚠️ | 69/80 |
| unevaluatedItems | ✅ | 0/0 |
| unevaluatedProperties | ⚠️ | 28/44 |
| uniqueItems | ⚠️ | 27/69 |

### Failures

<details>
<summary>additionalProperties - 10 failures</summary>

- **additionalProperties being false does not allow other properties**
  - Test: no additional properties is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2924754762.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties being false does not allow other properties**
  - Test: an additional property is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2924754762.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties being false does not allow other properties**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2924754762.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties being false does not allow other properties**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2924754762.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties being false does not allow other properties**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2924754762.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties being false does not allow other properties**
  - Test: patternProperties are not additional properties
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "bar": v.optional(v.any()), "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                                              ^
TypeError: v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe(v.check((val) => {
  for (const [key, value] of Object.entries(val))
    if (new RegExp("^v").test(key)) {
      if (!v.safeParse(v.any(), value).success)
        return !1;
    }
  return !0;
}, "Pattern property validation failed"))', 'v.strictObject({ bar: v.optional(v.any()), foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2924754762.ts:3:91
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **additionalProperties does not look in applicators**
  - Test: properties defined in allOf are not examined
  - Expected: `invalid`, Got: `true`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties doesn't consider dependentSchemas
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "foo2": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                   ^
TypeError: v.strictObject({ foo2: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ foo2: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.any(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.strictObject({ foo2: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2916969113.ts:3:64
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "foo2": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                   ^
TypeError: v.strictObject({ foo2: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ foo2: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.any(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.strictObject({ foo2: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2916969113.ts:3:64
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependentSchemas with additionalProperties**
  - Test: additionalProperties can't see bar even when foo2 is present
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.strictObject({ "foo2": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                   ^
TypeError: v.strictObject({ foo2: v.optional(v.any()) }).pipe is not a function. (In 'v.strictObject({ foo2: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.any(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.strictObject({ foo2: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2916969113.ts:3:64
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>allOf - 19 failures</summary>

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
- **allOf with boolean schemas, some false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with boolean schemas, all false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the first empty schema**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **allOf with the last empty schema**
  - Test: string is invalid
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
<summary>anyOf - 5 failures</summary>

- **anyOf**
  - Test: neither anyOf valid
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: mismatch base schema
  - Expected: `invalid`, Got: `true`
- **anyOf with base schema**
  - Test: both anyOf invalid
  - Expected: `invalid`, Got: `true`
- **anyOf with boolean schemas, all false**
  - Test: any value is invalid
  - Expected: `invalid`, Got: `true`
- **anyOf complex types**
  - Test: neither anyOf valid (complex)
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>const - 21 failures</summary>

- **const validation**
  - Test: same value is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object is valid
  - Expected: `valid`, Got: `false`
- **const with object**
  - Test: same object with different property order is valid
  - Expected: `valid`, Got: `false`
- **const with array**
  - Test: same array is valid
  - Expected: `valid`, Got: `false`
- **const with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **const with true does not match 1**
  - Test: true is valid
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
- **const with 0 does not match other zero-like types**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **const with 0 does not match other zero-like types**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **const with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **const with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **const with -2.0 matches integer and float types**
  - Test: integer -2 is valid
  - Expected: `valid`, Got: `false`
- **const with -2.0 matches integer and float types**
  - Test: float -2.0 is valid
  - Expected: `valid`, Got: `false`
- **float and integers are equal up to 64-bit representation limits**
  - Test: integer is valid
  - Expected: `valid`, Got: `false`
- **float and integers are equal up to 64-bit representation limits**
  - Test: float is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation but different codepoint**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`
- **characters with the same visual representation, but different number of codepoints**
  - Test: character uses the same codepoint
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>contains - 25 failures</summary>

- **contains keyword validation**
  - Test: array with item matching schema (5) is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.pipe(v.number(), v.minValue(5)), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4265063736.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword validation**
  - Test: array with item matching schema (6) is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.pipe(v.number(), v.minValue(5)), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4265063736.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword validation**
  - Test: array with two items matching schema (5, 6) is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.pipe(v.number(), v.minValue(5)), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4265063736.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword validation**
  - Test: array without items matching schema is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.pipe(v.number(), v.minValue(5)), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4265063736.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword validation**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.pipe(v.number(), v.minValue(5)), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4265063736.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword validation**
  - Test: not array or object is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.pipe(v.number(), v.minValue(5)), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4265063736.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with const keyword**
  - Test: array with item 5 is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2212654499.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with const keyword**
  - Test: array with two items 5 is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2212654499.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with const keyword**
  - Test: array without item 5 is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2212654499.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with boolean schema true**
  - Test: any non-empty array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.any(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2182759228.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with boolean schema true**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.any(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2182759228.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with boolean schema false**
  - Test: any non-empty array is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.never(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2261942869.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with boolean schema false**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.never(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2261942869.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid - string
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.never(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2261942869.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid - object
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.never(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2261942869.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid - number
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.never(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2261942869.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid - boolean
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.never(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2261942869.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains keyword with boolean schema false**
  - Test: non-arrays are valid - null
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.never(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2261942869.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **items + contains**
  - Test: matches items, does not match contains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.union([v.record(v.string(), v.pipe(v.number(), v.multipleOf(2))), v.array(v.pipe(v.number(), v.multipleOf(2))).pipe(v.check((arr) => {
                                                                                                                                    ^
TypeError: v.array(v.pipe(v.number(), v.multipleOf(2))).pipe is not a function. (In 'v.array(v.pipe(v.number(), v.multipleOf(2))).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.pipe(v.number(), v.multipleOf(3)), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.pipe(v.number(), v.multipleOf(2))).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-730942245.ts:3:129
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **items + contains**
  - Test: does not match items, matches contains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.union([v.record(v.string(), v.pipe(v.number(), v.multipleOf(2))), v.array(v.pipe(v.number(), v.multipleOf(2))).pipe(v.check((arr) => {
                                                                                                                                    ^
TypeError: v.array(v.pipe(v.number(), v.multipleOf(2))).pipe is not a function. (In 'v.array(v.pipe(v.number(), v.multipleOf(2))).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.pipe(v.number(), v.multipleOf(3)), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.pipe(v.number(), v.multipleOf(2))).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-730942245.ts:3:129
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **items + contains**
  - Test: matches both items and contains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.union([v.record(v.string(), v.pipe(v.number(), v.multipleOf(2))), v.array(v.pipe(v.number(), v.multipleOf(2))).pipe(v.check((arr) => {
                                                                                                                                    ^
TypeError: v.array(v.pipe(v.number(), v.multipleOf(2))).pipe is not a function. (In 'v.array(v.pipe(v.number(), v.multipleOf(2))).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.pipe(v.number(), v.multipleOf(3)), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.pipe(v.number(), v.multipleOf(2))).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-730942245.ts:3:129
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **items + contains**
  - Test: matches neither items nor contains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.union([v.record(v.string(), v.pipe(v.number(), v.multipleOf(2))), v.array(v.pipe(v.number(), v.multipleOf(2))).pipe(v.check((arr) => {
                                                                                                                                    ^
TypeError: v.array(v.pipe(v.number(), v.multipleOf(2))).pipe is not a function. (In 'v.array(v.pipe(v.number(), v.multipleOf(2))).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.pipe(v.number(), v.multipleOf(3)), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.pipe(v.number(), v.multipleOf(2))).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-730942245.ts:3:129
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains with false if subschema**
  - Test: any non-empty array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.any(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1700720773.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains with false if subschema**
  - Test: empty array is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.any(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1700720773.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **contains with null instance elements**
  - Test: allows null items
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.null_(), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2758533729.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

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
<summary>dependentRequired - 18 failures</summary>

- **single dependency**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3539769130.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: nondependant
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3539769130.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: with dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3539769130.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3539769130.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3539769130.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3539769130.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return Object.hasOwn(val, "foo");
  return !0;
}, "Property bar requires foo"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3539769130.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **empty dependents**
  - Test: non-object is valid
  - Expected: `valid`, Got: `false`
- **multiple dependents required**
  - Test: neither
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-426079751.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: nondependants
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-426079751.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: with dependencies
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-426079751.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: missing dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-426079751.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: missing other dependency
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-426079751.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **multiple dependents required**
  - Test: missing both dependencies
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "quux"))
    return Object.hasOwn(val, "foo") && Object.hasOwn(val, "bar");
  return !0;
}, "Property quux requires foo, bar"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-426079751.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: CRLF
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, `foo
bar`))
    return Object.hasOwn(val, "foo\rbar");
  return !0;
}, `Property foo
bar requires foo\rbar`))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1640473545.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, `foo
bar`))
    return Object.hasOwn(val, "foo\rbar");
  return !0;
}, `Property foo
bar requires foo\rbar`))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1640473545.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: CRLF missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, `foo
bar`))
    return Object.hasOwn(val, "foo\rbar");
  return !0;
}, `Property foo
bar requires foo\rbar`))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1640473545.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quotes missing dependent
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, `foo
bar`))
    return Object.hasOwn(val, "foo\rbar");
  return !0;
}, `Property foo
bar requires foo\rbar`))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1640473545.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>dependentSchemas - 20 failures</summary>

- **single dependency**
  - Test: valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3017203551.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: no dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3017203551.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: wrong type
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3017203551.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: wrong type other
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3017203551.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: wrong type both
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3017203551.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3017203551.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3017203551.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **single dependency**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.looseObject({ bar: v.optional(v.pipe(v.number(), v.integer())), foo: v.optional(v.pipe(v.number(), v.integer())) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-3017203551.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **boolean subschemas**
  - Test: object with property having schema true is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.never(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1100266434.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **boolean subschemas**
  - Test: object with property having schema false is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.never(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1100266434.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **boolean subschemas**
  - Test: object with both properties is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.never(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1100266434.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **boolean subschemas**
  - Test: empty object is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "bar"))
    return v.safeParse(v.never(), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1100266434.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted tab
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-378178152.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quote
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-378178152.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted tab invalid under dependent schema
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-378178152.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependencies with escaped characters**
  - Test: quoted quote invalid under dependent schema
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => {
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo\tbar"))
    return v.safeParse(v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 4, "Object must have at least 4 properties")), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-378178152.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependent subschema incompatible with root**
  - Test: matches root
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                 ^
TypeError: v.looseObject({ foo: v.optional(v.any()) }).pipe is not a function. (In 'v.looseObject({ foo: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.strictObject({ bar: v.optional(v.any()) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({ foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1492552084.ts:3:62
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependent subschema incompatible with root**
  - Test: matches dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                 ^
TypeError: v.looseObject({ foo: v.optional(v.any()) }).pipe is not a function. (In 'v.looseObject({ foo: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.strictObject({ bar: v.optional(v.any()) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({ foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1492552084.ts:3:62
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependent subschema incompatible with root**
  - Test: matches both
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                 ^
TypeError: v.looseObject({ foo: v.optional(v.any()) }).pipe is not a function. (In 'v.looseObject({ foo: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.strictObject({ bar: v.optional(v.any()) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({ foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1492552084.ts:3:62
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **dependent subschema incompatible with root**
  - Test: no dependency
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({ "foo": v.optional(v.any()) }).pipe(v.check((val) => {
                                                                 ^
TypeError: v.looseObject({ foo: v.optional(v.any()) }).pipe is not a function. (In 'v.looseObject({ foo: v.optional(v.any()) }).pipe(v.check((val) => {
  if (Object.hasOwn(val, "foo"))
    return v.safeParse(v.strictObject({ bar: v.optional(v.any()) }), val).success;
  return !0;
}, "Schema dependency validation failed"))', 'v.looseObject({ foo: v.optional(v.any()) }).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1492552084.ts:3:62
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>dynamicRef - 7 failures</summary>

- **A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor**
  - Test: An array containing non-strings is invalid
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
<summary>enum - 22 failures</summary>

- **simple enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: one of the enum is valid
  - Expected: `valid`, Got: `false`
- **heterogeneous enum validation**
  - Test: something else is invalid
  - Expected: `invalid`, Got: `true`
- **heterogeneous enum validation**
  - Test: valid object matches
  - Expected: `valid`, Got: `false`
- **heterogeneous enum-with-null validation**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: both properties are valid
  - Expected: `valid`, Got: `false`
- **enums in properties**
  - Test: missing optional property is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 1 is valid
  - Expected: `valid`, Got: `false`
- **enum with escaped characters**
  - Test: member 2 is valid
  - Expected: `valid`, Got: `false`
- **enum with false does not match 0**
  - Test: false is valid
  - Expected: `valid`, Got: `false`
- **enum with [false] does not match [0]**
  - Test: [false] is valid
  - Expected: `valid`, Got: `false`
- **enum with true does not match 1**
  - Test: true is valid
  - Expected: `valid`, Got: `false`
- **enum with [true] does not match [1]**
  - Test: [true] is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: integer zero is valid
  - Expected: `valid`, Got: `false`
- **enum with 0 does not match false**
  - Test: float zero is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0] is valid
  - Expected: `valid`, Got: `false`
- **enum with [0] does not match [false]**
  - Test: [0.0] is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: integer one is valid
  - Expected: `valid`, Got: `false`
- **enum with 1 does not match true**
  - Test: float one is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1] is valid
  - Expected: `valid`, Got: `false`
- **enum with [1] does not match [true]**
  - Test: [1.0] is valid
  - Expected: `valid`, Got: `false`
- **nul characters in strings**
  - Test: match string with nul
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
<summary>items - 10 failures</summary>

- **a schema given for items**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `false`
- **a schema given for items**
  - Test: JavaScript pseudo-array is valid
  - Expected: `valid`, Got: `false`
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
  - Test: wrong sub-item
  - Expected: `invalid`, Got: `true`
- **items and subitems**
  - Test: fewer items is valid
  - Expected: `valid`, Got: `false`
- **prefixItems with no additional items allowed**
  - Test: additional items are not permitted
  - Expected: `invalid`, Got: `true`
- **items does not look in applicators, valid case**
  - Test: prefixItems in allOf does not constrain items, invalid case
  - Expected: `invalid`, Got: `true`
- **items with heterogeneous array**
  - Test: heterogeneous invalid instance
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>maxContains - 10 failures</summary>

- **maxContains with contains**
  - Test: empty array
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1 && count <= 1;
}, "Array must contain between 1 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1320946900.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains with contains**
  - Test: all elements match, valid maxContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1 && count <= 1;
}, "Array must contain between 1 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1320946900.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains with contains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1 && count <= 1;
}, "Array must contain between 1 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1320946900.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains with contains**
  - Test: some elements match, valid maxContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1 && count <= 1;
}, "Array must contain between 1 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1320946900.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains with contains**
  - Test: some elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1 && count <= 1;
}, "Array must contain between 1 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1320946900.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains with contains, value with a decimal**
  - Test: one element matches, valid maxContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1 && count <= 1;
}, "Array must contain between 1 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1016411087.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains with contains, value with a decimal**
  - Test: too many elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1 && count <= 1;
}, "Array must contain between 1 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1016411087.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains < maxContains**
  - Test: array with actual < minContains < maxContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1 && count <= 3;
}, "Array must contain between 1 and 3 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-143323582.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains < maxContains**
  - Test: array with minContains < actual < maxContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1 && count <= 3;
}, "Array must contain between 1 and 3 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-143323582.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains < maxContains**
  - Test: array with minContains < maxContains < actual
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1 && count <= 3;
}, "Array must contain between 1 and 3 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-143323582.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>maxItems - 6 failures</summary>

- **maxItems validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-682507477.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-682507477.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-682507477.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-682507477.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2788506411.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxItems validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.maxLength(2));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(2))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2788506411.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>maxLength - 1 failure</summary>

- **maxLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>maxProperties - 10 failures</summary>

- **maxProperties validation**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-643438638.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-643438638.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-643438638.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-643438638.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-643438638.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-643438638.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation with a decimal**
  - Test: shorter is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4288418518.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties validation with a decimal**
  - Test: too long is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 2, "Object must have at most 2 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-4288418518.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties = 0 means the object is empty**
  - Test: no properties is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 0, "Object must have at most 0 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 0, "Object must have at most 0 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-807656817.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxProperties = 0 means the object is empty**
  - Test: one property is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 0, "Object must have at most 0 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length <= 0, "Object must have at most 0 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-807656817.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>maximum - 1 failure</summary>

- **maximum validation**
  - Test: ignores non-numbers
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minContains - 26 failures</summary>

- **minContains=1 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2767911535.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=1 with contains**
  - Test: no elements match
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2767911535.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=1 with contains**
  - Test: single element matches, valid minContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2767911535.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=1 with contains**
  - Test: some elements match, valid minContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2767911535.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=1 with contains**
  - Test: all elements match, valid minContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 1;
}, "Array must contain at least 1 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2767911535.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=2 with contains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2;
}, "Array must contain at least 2 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1840712362.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=2 with contains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2;
}, "Array must contain at least 2 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1840712362.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=2 with contains**
  - Test: some elements match, invalid minContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2;
}, "Array must contain at least 2 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1840712362.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=2 with contains**
  - Test: all elements match, valid minContains (exactly as needed)
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2;
}, "Array must contain at least 2 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1840712362.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=2 with contains**
  - Test: all elements match, valid minContains (more than needed)
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2;
}, "Array must contain at least 2 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1840712362.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=2 with contains**
  - Test: some elements match, valid minContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2;
}, "Array must contain at least 2 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1840712362.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=2 with contains with a decimal value**
  - Test: one element matches, invalid minContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2;
}, "Array must contain at least 2 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2607948915.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains=2 with contains with a decimal value**
  - Test: both elements match, valid minContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2;
}, "Array must contain at least 2 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2607948915.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains = minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2 && count <= 2;
}, "Array must contain between 2 and 2 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1691716555.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains = minContains**
  - Test: all elements match, invalid minContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2 && count <= 2;
}, "Array must contain between 2 and 2 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1691716555.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains = minContains**
  - Test: all elements match, invalid maxContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2 && count <= 2;
}, "Array must contain between 2 and 2 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1691716555.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains = minContains**
  - Test: all elements match, valid maxContains and minContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 2 && count <= 2;
}, "Array must contain between 2 and 2 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1691716555.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains < minContains**
  - Test: empty data
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 3 && count <= 1;
}, "Array must contain between 3 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854266377.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains < minContains**
  - Test: invalid minContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 3 && count <= 1;
}, "Array must contain between 3 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854266377.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains < minContains**
  - Test: invalid maxContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 3 && count <= 1;
}, "Array must contain between 3 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854266377.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **maxContains < minContains**
  - Test: invalid maxContains and minContains
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 3 && count <= 1;
}, "Array must contain between 3 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1854266377.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains = 0**
  - Test: empty data
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 0;
}, "Array must contain at least 0 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1252490661.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains = 0**
  - Test: minContains = 0 makes contains always pass
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 0;
}, "Array must contain at least 0 item(s) matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1252490661.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains = 0 with maxContains**
  - Test: empty data
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 0 && count <= 1;
}, "Array must contain between 0 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1296484592.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains = 0 with maxContains**
  - Test: not more than maxContains
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 0 && count <= 1;
}, "Array must contain between 0 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1296484592.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minContains = 0 with maxContains**
  - Test: too many
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  let count = 0;
  for (const item of arr)
    if (v.safeParse(v.literal(null), item).success)
      count++;
  return count >= 0 && count <= 1;
}, "Array must contain between 0 and 1 items matching schema"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1296484592.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>minItems - 6 failures</summary>

- **minItems validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-934969662.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-934969662.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-934969662.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation**
  - Test: ignores non-arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-934969662.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-267734821.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minItems validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.minLength(1));
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.minLength(1))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-267734821.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

</details>

<details>
<summary>minLength - 1 failure</summary>

- **minLength validation**
  - Test: ignores non-strings
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>minProperties - 8 failures</summary>

- **minProperties validation**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1969578109.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation**
  - Test: exact length is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1969578109.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1969578109.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation**
  - Test: ignores arrays
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1969578109.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation**
  - Test: ignores strings
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1969578109.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1969578109.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation with a decimal**
  - Test: longer is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1612352554.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **minProperties validation with a decimal**
  - Test: too short is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"));
                                     ^
TypeError: v.looseObject({}).pipe is not a function. (In 'v.looseObject({}).pipe(v.check((val) => Object.keys(val).length >= 1, "Object must have at least 1 properties"))', 'v.looseObject({}).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1612352554.ts:3:34
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`

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
<summary>not - 14 failures</summary>

- **not**
  - Test: allowed
  - Expected: `valid`, Got: `false`
- **not multiple types**
  - Test: valid
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: match
  - Expected: `valid`, Got: `false`
- **not more complex schema**
  - Test: other match
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: number is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: string is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: boolean true is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: boolean false is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: null is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: object is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: empty object is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: array is valid
  - Expected: `valid`, Got: `false`
- **allow everything with boolean schema false**
  - Test: empty array is valid
  - Expected: `valid`, Got: `false`
- **double negation**
  - Test: any value is valid
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>oneOf - 14 failures</summary>

- **oneOf**
  - Test: both oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf**
  - Test: neither oneOf valid
  - Expected: `invalid`, Got: `true`
- **oneOf with base schema**
  - Test: mismatch base schema
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
- **oneOf with boolean schemas, all false**
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
<summary>pattern - 6 failures</summary>

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
<summary>patternProperties - 2 failures</summary>

- **patternProperties validates properties matching a regex**
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **patternProperties validates properties matching a regex**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>prefixItems - 5 failures</summary>

- **a schema given for prefixItems**
  - Test: incomplete array of items
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

</details>

<details>
<summary>properties - 16 failures</summary>

- **object properties validation**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property validates property
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2382617915.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: property invalidates property
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2382617915.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates property
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2382617915.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty validates nonproperty
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2382617915.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: patternProperty invalidates nonproperty
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2382617915.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty ignores property
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2382617915.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty validates others
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2382617915.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **properties, patternProperties, additionalProperties interaction**
  - Test: additionalProperty invalidates others
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.objectWithRest({ "bar": v.optional(v.array(v.any())), "foo": v.optional(v.array(v.any()).pipe(v.maxLength(3))) }, v.pipe(v.number(), v.integer())).pipe(v.check((val) => {
                                                                                                              ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.maxLength(3))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-2382617915.ts:3:107
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
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
  - Test: __proto__ not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: toString not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: constructor not valid
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **properties whose names are Javascript object property names**
  - Test: all present and valid
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>propertyNames - 4 failures</summary>

- **propertyNames validation**
  - Test: some property names invalid
  - Expected: `invalid`, Got: `true`
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
<summary>ref - 8 failures</summary>

- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `false`
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
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **URN base URI with f-component**
  - Test: is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to fetch "https://json-schema.org/v1": failed to fetch https://json-schema.org/v1: status 404`

</details>

<details>
<summary>refRemote - 1 failure</summary>

- **root ref in remote ref**
  - Test: object is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>required - 11 failures</summary>

- **required validation**
  - Test: non-present required property is invalid
  - Expected: `invalid`, Got: `true`
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
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: ignores other non-objects
  - Expected: `valid`, Got: `false`
- **required properties whose names are Javascript object property names**
  - Test: none of the properties mentioned
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: __proto__ present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: toString present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: constructor present
  - Expected: `invalid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`
- **required properties whose names are Javascript object property names**
  - Test: all present
  - Expected: `valid`, Got: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

</details>

<details>
<summary>type - 11 failures</summary>

- **object type matches objects**
  - Test: an array is not an object
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a float is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an object is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: an array is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: a boolean is invalid
  - Expected: `invalid`, Got: `true`
- **multiple types can be specified in an array**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`
- **type: array or object**
  - Test: null is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: number is invalid
  - Expected: `invalid`, Got: `true`
- **type: array, object or null**
  - Test: string is invalid
  - Expected: `invalid`, Got: `true`

</details>

<details>
<summary>unevaluatedProperties - 16 failures</summary>

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
  - Test: ignores strings
  - Expected: `valid`, Got: `false`
- **non-object instances are valid**
  - Test: ignores null
  - Expected: `valid`, Got: `false`

</details>

<details>
<summary>uniqueItems - 42 failures</summary>

- **uniqueItems validation**
  - Test: unique array of integers is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of integers is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of more than two integers is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: numbers are unique if mathematically unequal
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: false is not equal to zero
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: true is not equal to one
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of strings is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of strings is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of objects is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of objects is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: property order of array of objects is ignored
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of nested objects is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of nested objects is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique array of arrays is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of arrays is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique array of more than two arrays is invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: 1 and true are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: 0 and false are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: [1] and [true] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: [0] and [false] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: nested [1] and [true] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: nested [0] and [false] are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: unique heterogeneous types are valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: non-unique heterogeneous types are invalid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: different objects are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: objects are non-unique despite key order
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: {"a": false} and {"a": 0} are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems validation**
  - Test: {"a": true} and {"a": 1} are unique
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.array(v.any()).pipe(v.check((arr) => {
                                    ^
TypeError: v.array(v.any()).pipe is not a function. (In 'v.array(v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.array(v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1388223108.ts:3:33
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1543725197.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1543725197.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1543725197.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1543725197.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: unique array extended from [false, true] is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1543725197.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: unique array extended from [true, false] is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1543725197.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [false, true] is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1543725197.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items**
  - Test: non-unique array extended from [true, false] is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
                                                                        ^
TypeError: v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe is not a function. (In 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tupleWithRest([v.boolean(), v.boolean()], v.any()).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1543725197.ts:3:69
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, true] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1182126475.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, false] from items array is valid
  - Expected: `valid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1182126475.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [false, false] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1182126475.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: [true, true] from items array is not valid
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1182126475.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `error: harness execution error: harness execution failed: exit status 1
stderr: 1 | import * as v from "valibot";
2 | 
3 | const schema = v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
                                                       ^
TypeError: v.tuple([v.boolean(), v.boolean()]).pipe is not a function. (In 'v.tuple([v.boolean(), v.boolean()]).pipe(v.check((arr) => {
  const seen = new Set;
  for (const item of arr) {
    const key = JSON.stringify(item);
    if (seen.has(key))
      return !1;
    seen.add(key);
  }
  return !0;
}, "Array items must be unique"))', 'v.tuple([v.boolean(), v.boolean()]).pipe' is undefined)
      at /home/trapani/dev/xschema-adapters-ts/typescript/packages/adapters/valibot/xschema-harness-1182126475.ts:3:52
      at loadAndEvaluateModule (2:1)

Bun v1.3.5 (Linux x64)
`
- **uniqueItems=false with an array of items and additionalItems=false**
  - Test: extra items are invalid even if unique
  - Expected: `invalid`, Got: `true`

</details>

