# xschema-pydantic Compliance Report

## Summary

| Draft | Passed | Failed | Skipped | Unsupported | Coverage |
| ----- | ------ | ------ | ------- | ----------- | -------- |
| draft2019-09 | 1024 | 32 | 0 | 178 | 97.0% |
| draft2020-12 | 1038 | 32 | 0 | 201 | 97.0% |
| draft3 | 423 | 11 | 0 | 0 | 97.5% |
| draft4 | 596 | 17 | 0 | 0 | 97.2% |
| draft6 | 807 | 25 | 0 | 0 | 97.0% |
| draft7 | 891 | 25 | 0 | 0 | 97.3% |

## Badges

![draft2019-09](https://img.shields.io/badge/draft2019-09%20compliance-97.0%25-brightgreen)
![draft2020-12](https://img.shields.io/badge/draft2020-12%20compliance-97.0%25-brightgreen)
![draft3](https://img.shields.io/badge/draft3%20compliance-97.5%25-brightgreen)
![draft4](https://img.shields.io/badge/draft4%20compliance-97.2%25-brightgreen)
![draft6](https://img.shields.io/badge/draft6%20compliance-97.0%25-brightgreen)
![draft7](https://img.shields.io/badge/draft7%20compliance-97.3%25-brightgreen)

## draft2019-09

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 19/19 |
| additionalProperties | ✅ | 21/21 |
| allOf | ✅ | 30/30 |
| anchor | ✅ | 8/8 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 21/21 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ❌ | 0/2 |
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
| not | ✅ | 38/38 |
| oneOf | ✅ | 27/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| properties | ✅ | 28/28 |
| propertyNames | ✅ | 20/20 |
| recursiveRef | ✅ | 0/0 |
| ref | ⚠️ | 53/79 |
| refRemote | ⚠️ | 27/31 |
| required | ✅ | 16/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 14/14 |
| unevaluatedProperties | ✅ | 25/25 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ✅ | 5/5 |

### Unsupported Features

These tests are intentionally excluded due to documented limitations.

<details>
<summary>$recursiveAnchor is not supported: Recursive references ($recursiveRef/$recursiveAnchor) require runtime scope tracking (at /$defs/inner) (4 tests)</summary>

- `draft2019-09/recursiveRef/dynamic $recursiveRef destination (not predictable at schema compile time)/integer node`
- `draft2019-09/recursiveRef/dynamic $recursiveRef destination (not predictable at schema compile time)/numeric node`
- `draft2019-09/recursiveRef/multiple dynamic paths to the $recursiveRef keyword/recurse to anyLeafNode - floats are allowed`
- `draft2019-09/recursiveRef/multiple dynamic paths to the $recursiveRef keyword/recurse to integerNode - floats are not allowed`

</details>

<details>
<summary>$recursiveAnchor is not supported: Recursive references ($recursiveRef/$recursiveAnchor) require runtime scope tracking (at /$defs/myobject) (10 tests)</summary>

- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor works like $ref/integer does not match as a property value`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor works like $ref/integer matches at the outer level`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor works like $ref/single level match`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor works like $ref/two levels, integer does not match as a property value`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor works like $ref/two levels, properties match with inner definition`
- `draft2019-09/recursiveRef/$recursiveRef without using nesting/integer does not match as a property value`
- `draft2019-09/recursiveRef/$recursiveRef without using nesting/integer matches at the outer level`
- `draft2019-09/recursiveRef/$recursiveRef without using nesting/single level match`
- `draft2019-09/recursiveRef/$recursiveRef without using nesting/two levels, no match`
- `draft2019-09/recursiveRef/$recursiveRef without using nesting/two levels, properties match with inner definition`

</details>

<details>
<summary>$recursiveAnchor is not supported: Recursive references ($recursiveRef/$recursiveAnchor) require runtime scope tracking (at /anyOf/1/additionalProperties) (3 tests)</summary>

- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor in the outer schema resource/leaf node does not match: recursion only uses inner schema`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor in the outer schema resource/leaf node does not match; no recursion`
- `draft2019-09/recursiveRef/$recursiveRef with no $recursiveAnchor in the outer schema resource/leaf node matches: recursion only uses inner schema`

</details>

<details>
<summary>$recursiveAnchor is not supported: Recursive references ($recursiveRef/$recursiveAnchor) require runtime scope tracking (at root) (19 tests)</summary>

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
- `draft2019-09/ref/$ref with $recursiveAnchor/extra items allowed for inner arrays`
- `draft2019-09/ref/$ref with $recursiveAnchor/extra items disallowed for root`
- `draft2019-09/unevaluatedItems/unevaluatedItems with $recursiveRef/with no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with $recursiveRef/with unevaluated items`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with $recursiveRef/with no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with $recursiveRef/with unevaluated properties`

</details>

<details>
<summary>$recursiveRef is not supported: Recursive references ($recursiveRef/$recursiveAnchor) require runtime scope tracking (at /properties/foo) (4 tests)</summary>

- `draft2019-09/recursiveRef/$recursiveRef without $recursiveAnchor works like $ref/match`
- `draft2019-09/recursiveRef/$recursiveRef without $recursiveAnchor works like $ref/mismatch`
- `draft2019-09/recursiveRef/$recursiveRef without $recursiveAnchor works like $ref/recursive match`
- `draft2019-09/recursiveRef/$recursiveRef without $recursiveAnchor works like $ref/recursive mismatch`

</details>

<details>
<summary>unevaluatedItems is not supported: inside applicator subschema cannot see sibling annotations (cousins problem) (at /allOf/1) (1 test)</summary>

- `draft2019-09/unevaluatedItems/unevaluatedItems can't see inside cousins/always fails`

</details>

<details>
<summary>unevaluatedItems is not supported: requires annotation tracking when combined with applicators (prefixItems, contains, allOf, anyOf, oneOf, if) (at /properties/foo) (2 tests)</summary>

- `draft2019-09/unevaluatedItems/item is evaluated in an uncle schema to unevaluatedItems/no extra items`
- `draft2019-09/unevaluatedItems/item is evaluated in an uncle schema to unevaluatedItems/uncle keyword evaluation is not significant`

</details>

<details>
<summary>unevaluatedItems is not supported: requires annotation tracking when combined with applicators (prefixItems, contains, allOf, anyOf, oneOf, if) (at root) (37 tests)</summary>

- `draft2019-09/unevaluatedItems/Evaluated items collection needs to consider instance location/with an unevaluated item that exists at another location`
- `draft2019-09/unevaluatedItems/unevaluatedItems before $ref/with no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems before $ref/with unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems can see annotations from if without then and else/invalid in case if is evaluated`
- `draft2019-09/unevaluatedItems/unevaluatedItems can see annotations from if without then and else/valid in case if is evaluated`
- `draft2019-09/unevaluatedItems/unevaluatedItems with $ref/with no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with $ref/with unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with anyOf/when one schema matches and has no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with anyOf/when one schema matches and has unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with anyOf/when two schemas match and has no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with anyOf/when two schemas match and has unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with boolean schemas/with no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with boolean schemas/with unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with if/then/else/when if doesn't match and it has no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with if/then/else/when if doesn't match and it has unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with if/then/else/when if matches and it has no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with if/then/else/when if matches and it has unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with ignored additionalItems/all valid under unevaluatedItems`
- `draft2019-09/unevaluatedItems/unevaluatedItems with ignored additionalItems/invalid under unevaluatedItems`
- `draft2019-09/unevaluatedItems/unevaluatedItems with ignored applicator additionalItems/all valid under unevaluatedItems`
- `draft2019-09/unevaluatedItems/unevaluatedItems with ignored applicator additionalItems/invalid under unevaluatedItems`
- `draft2019-09/unevaluatedItems/unevaluatedItems with items and additionalItems/unevaluatedItems doesn't apply`
- `draft2019-09/unevaluatedItems/unevaluatedItems with nested items and additionalItems/with additional items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with nested items and additionalItems/with no additional items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with nested items/with invalid additional item`
- `draft2019-09/unevaluatedItems/unevaluatedItems with nested items/with no additional items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with nested items/with only (valid) additional items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with nested tuple/with no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with nested tuple/with unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with nested unevaluatedItems/with additional items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with nested unevaluatedItems/with no additional items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with not/with unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with oneOf/with no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with oneOf/with unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with tuple/with no unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with tuple/with unevaluated items`
- `draft2019-09/unevaluatedItems/unevaluatedItems with uniform items/unevaluatedItems doesn't apply`

</details>

<details>
<summary>unevaluatedProperties is not supported: cyclic $ref with unevaluatedProperties requires recursive annotation tracking (at root) (7 tests)</summary>

- `draft2019-09/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Deep nested is valid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Empty is valid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Nested is valid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Single is valid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Unevaluated on 1st level is invalid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Unevaluated on 2nd level is invalid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Unevaluated on 3rd level is invalid`

</details>

<details>
<summary>unevaluatedProperties is not supported: inside applicator subschema cannot see sibling annotations (cousins problem) (at /allOf/0) (8 tests)</summary>

- `draft2019-09/unevaluatedProperties/cousin unevaluatedProperties, true and false, false with properties/with nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/cousin unevaluatedProperties, true and false, false with properties/with no nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/cousin unevaluatedProperties, true and false, true with properties/with nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/cousin unevaluatedProperties, true and false, true with properties/with no nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/in-place applicator siblings, allOf has unevaluated/base case: both properties present`
- `draft2019-09/unevaluatedProperties/in-place applicator siblings, allOf has unevaluated/in place applicator siblings, bar is missing`
- `draft2019-09/unevaluatedProperties/in-place applicator siblings, allOf has unevaluated/in place applicator siblings, foo is missing`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties can't see inside cousins (reverse order)/always fails`

</details>

<details>
<summary>unevaluatedProperties is not supported: inside applicator subschema cannot see sibling annotations (cousins problem) (at /allOf/1) (1 test)</summary>

- `draft2019-09/unevaluatedProperties/unevaluatedProperties can't see inside cousins/always fails`

</details>

<details>
<summary>unevaluatedProperties is not supported: inside applicator subschema cannot see sibling annotations (cousins problem) (at /anyOf/0) (3 tests)</summary>

- `draft2019-09/unevaluatedProperties/in-place applicator siblings, anyOf has unevaluated/base case: both properties present`
- `draft2019-09/unevaluatedProperties/in-place applicator siblings, anyOf has unevaluated/in place applicator siblings, bar is missing`
- `draft2019-09/unevaluatedProperties/in-place applicator siblings, anyOf has unevaluated/in place applicator siblings, foo is missing`

</details>

<details>
<summary>unevaluatedProperties is not supported: requires annotation tracking when combined with applicators (allOf, anyOf, oneOf, if, $ref, dependentSchemas, not) (at /not) (2 tests)</summary>

- `draft2019-09/not/collect annotations inside a 'not', even if collection is disabled/annotations are still collected inside a 'not'`
- `draft2019-09/not/collect annotations inside a 'not', even if collection is disabled/unevaluated property`

</details>

<details>
<summary>unevaluatedProperties is not supported: requires annotation tracking when combined with applicators (allOf, anyOf, oneOf, if, $ref, dependentSchemas, not) (at root) (77 tests)</summary>

- `draft2019-09/unevaluatedProperties/dependentSchemas with unevaluatedProperties/unevaluatedProperties doesn't consider dependentSchemas`
- `draft2019-09/unevaluatedProperties/dependentSchemas with unevaluatedProperties/unevaluatedProperties doesn't see bar when foo2 is absent`
- `draft2019-09/unevaluatedProperties/dependentSchemas with unevaluatedProperties/unevaluatedProperties sees bar when foo2 is present`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/Empty is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/a + b is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/a + c is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/a + d is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/a is valid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/all + a is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/all + foo is valid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/all is valid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/b + c is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/b + d is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/b is valid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/c + d is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/c is valid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/d is valid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/xx + a is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/xx + b is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/xx + c is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/xx + d is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/xx + foo is invalid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/xx + foox is valid`
- `draft2019-09/unevaluatedProperties/dynamic evalation inside nested refs/xx is valid`
- `draft2019-09/unevaluatedProperties/nested unevaluatedProperties, outer false, inner true, properties inside/with nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/nested unevaluatedProperties, outer false, inner true, properties inside/with no nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/nested unevaluatedProperties, outer false, inner true, properties outside/with nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/nested unevaluatedProperties, outer false, inner true, properties outside/with no nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/nested unevaluatedProperties, outer true, inner false, properties inside/with nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/nested unevaluatedProperties, outer true, inner false, properties inside/with no nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/nested unevaluatedProperties, outer true, inner false, properties outside/with nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/nested unevaluatedProperties, outer true, inner false, properties outside/with no nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/Empty is invalid (no x or y)`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and b and x and y are invalid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and b and x are valid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and b and y are valid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and b are invalid (no x or y)`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and x are valid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and y are valid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/x and y are invalid`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties before $ref/with no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties before $ref/with unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties can see annotations from if without then and else/invalid in case if is evaluated`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties can see annotations from if without then and else/valid in case if is evaluated`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with $ref/with no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with $ref/with unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with anyOf/when one matches and has no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with anyOf/when one matches and has unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with anyOf/when two match and has no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with anyOf/when two match and has unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with boolean schemas/with no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with boolean schemas/with unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with dependentSchemas/with no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with dependentSchemas/with unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else, else not defined/when if is false and has no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else, else not defined/when if is false and has unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else, else not defined/when if is true and has no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else, else not defined/when if is true and has unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else, then not defined/when if is false and has no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else, then not defined/when if is false and has unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else, then not defined/when if is true and has no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else, then not defined/when if is true and has unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else/when if is false and has no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else/when if is false and has unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else/when if is true and has no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with if/then/else/when if is true and has unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with nested additionalProperties/with additional properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with nested additionalProperties/with no additional properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with nested patternProperties/with additional properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with nested patternProperties/with no additional properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with nested properties/with additional properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with nested properties/with no additional properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with nested unevaluatedProperties/with nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with nested unevaluatedProperties/with no nested unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with not/with unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with oneOf/with no unevaluated properties`
- `draft2019-09/unevaluatedProperties/unevaluatedProperties with oneOf/with unevaluated properties`

</details>

### Unexpected Failures

<details>
<summary>defs - 2 failures</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/defs/group_0: encountered unresolved non-local $ref "https://json-schema.org/draft/2019-09/schema"`
- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/defs/group_0: encountered unresolved non-local $ref "https://json-schema.org/draft/2019-09/schema"`

</details>

<details>
<summary>ref - 26 failures</summary>

- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_11: recursive local $ref "#/$defs/node" is not supported`
- **Recursive references between schemas**
  - Test: valid tree
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_11: recursive local $ref "#/$defs/node" is not supported`
- **URN ref with nested pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_28: failed to resolve local $ref "#/$defs/bar": key "bar" not found`
- **URN ref with nested pointer ref**
  - Test: a string is valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_28: failed to resolve local $ref "#/$defs/bar": key "bar" not found`
- **order of evaluation: $id and $anchor and $ref**
  - Test: data is valid against first definition
  - Expected: `valid`, Got: `false`
- **order of evaluation: $id and $ref on nested schema**
  - Test: data is invalid against nested sibling
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json"`
- **order of evaluation: $id and $ref on nested schema**
  - Test: data is valid against nested sibling
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_20: encountered unresolved non-local $ref "./bar.json"`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_7: $ref value must be a string`
- **property named $ref that is not a reference**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_7: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_8: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_8: $ref value must be a string`
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_15: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_15: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **refs with relative uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_15: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_16: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_16: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_16: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_6: encountered unresolved non-local $ref "https://json-schema.org/draft/2019-09/schema"`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_6: encountered unresolved non-local $ref "https://json-schema.org/draft/2019-09/schema"`
- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_0: recursive local $ref "#" is not supported`
- **simple URN base URI with $ref via the URN**
  - Test: invalid under the URN IDed schema
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_21: recursive local $ref "#" is not supported`
- **simple URN base URI with $ref via the URN**
  - Test: valid under the URN IDed schema
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/ref/group_21: recursive local $ref "#" is not supported`

</details>

<details>
<summary>refRemote - 4 failures</summary>

- **anchor within remote ref**
  - Test: remote anchor invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_2: recursive local $ref "#/$defs/localhost_1234_draft2019_09_locationIndependentIdentifier_json__A" is not supported`
- **anchor within remote ref**
  - Test: remote anchor valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_2: recursive local $ref "#/$defs/localhost_1234_draft2019_09_locationIndependentIdentifier_json__A" is not supported`
- **fragment within remote ref**
  - Test: remote fragment invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_1: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported`
- **fragment within remote ref**
  - Test: remote fragment valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2019-09/refRemote/group_1: recursive local $ref "#/$defs/localhost_1234_draft2019_09_subSchemas_json__integer" is not supported`

</details>

## draft2020-12

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalProperties | ✅ | 21/21 |
| allOf | ✅ | 30/30 |
| anchor | ✅ | 8/8 |
| anyOf | ✅ | 18/18 |
| boolean_schema | ✅ | 18/18 |
| const | ✅ | 54/54 |
| contains | ✅ | 21/21 |
| content | ✅ | 18/18 |
| default | ✅ | 7/7 |
| defs | ❌ | 0/2 |
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
| not | ✅ | 38/38 |
| oneOf | ✅ | 27/27 |
| pattern | ✅ | 9/9 |
| patternProperties | ✅ | 23/23 |
| prefixItems | ✅ | 11/11 |
| properties | ✅ | 28/28 |
| propertyNames | ✅ | 20/20 |
| ref | ⚠️ | 53/79 |
| refRemote | ⚠️ | 27/31 |
| required | ✅ | 16/16 |
| type | ✅ | 80/80 |
| unevaluatedItems | ✅ | 14/14 |
| unevaluatedProperties | ✅ | 27/27 |
| uniqueItems | ✅ | 69/69 |
| vocabulary | ✅ | 5/5 |

### Unsupported Features

These tests are intentionally excluded due to documented limitations.

<details>
<summary>$dynamicAnchor is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /$defs/content) (2 tests)</summary>

- `draft2020-12/dynamicRef/$dynamicRef skips over intermediate resources - direct reference/integer property passes`
- `draft2020-12/dynamicRef/$dynamicRef skips over intermediate resources - direct reference/string property fails`

</details>

<details>
<summary>$dynamicAnchor is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /$defs/defaultItemType) (4 tests)</summary>

- `draft2020-12/dynamicRef/multiple dynamic paths to the $dynamicRef keyword/number list with number values`
- `draft2020-12/dynamicRef/multiple dynamic paths to the $dynamicRef keyword/number list with string values`
- `draft2020-12/dynamicRef/multiple dynamic paths to the $dynamicRef keyword/string list with number values`
- `draft2020-12/dynamicRef/multiple dynamic paths to the $dynamicRef keyword/string list with string values`

</details>

<details>
<summary>$dynamicAnchor is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /$defs/detached) (2 tests)</summary>

- `draft2020-12/dynamicRef/$ref to $dynamicRef finds detached $dynamicAnchor/non-number is invalid`
- `draft2020-12/dynamicRef/$ref to $dynamicRef finds detached $dynamicAnchor/number is valid`

</details>

<details>
<summary>$dynamicAnchor is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /$defs/elements) (6 tests)</summary>

- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $defs first/correct extended schema`
- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $defs first/incorrect extended schema`
- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $defs first/incorrect parent schema`
- `draft2020-12/dynamicRef/tests for implementation dynamic anchor and reference link/correct extended schema`
- `draft2020-12/dynamicRef/tests for implementation dynamic anchor and reference link/incorrect extended schema`
- `draft2020-12/dynamicRef/tests for implementation dynamic anchor and reference link/incorrect parent schema`

</details>

<details>
<summary>$dynamicAnchor is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /$defs/foo) (12 tests)</summary>

- `draft2020-12/dynamicRef/A $dynamicRef resolves to the first $dynamicAnchor still in scope that is encountered when the schema is evaluated/An array containing non-strings is invalid`
- `draft2020-12/dynamicRef/A $dynamicRef resolves to the first $dynamicAnchor still in scope that is encountered when the schema is evaluated/An array of strings is valid`
- `draft2020-12/dynamicRef/A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor/An array containing non-strings is invalid`
- `draft2020-12/dynamicRef/A $dynamicRef to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor/An array of strings is valid`
- `draft2020-12/dynamicRef/A $dynamicRef with a non-matching $dynamicAnchor in the same schema resource behaves like a normal $ref to $anchor/Any array is valid`
- `draft2020-12/dynamicRef/A $dynamicRef with intermediate scopes that don't include a matching $dynamicAnchor does not affect dynamic scope resolution/An array containing non-strings is invalid`
- `draft2020-12/dynamicRef/A $dynamicRef with intermediate scopes that don't include a matching $dynamicAnchor does not affect dynamic scope resolution/An array of strings is valid`
- `draft2020-12/dynamicRef/A $dynamicRef without a matching $dynamicAnchor in the same schema resource behaves like a normal $ref to $anchor/Any array is valid`
- `draft2020-12/dynamicRef/A $dynamicRef without anchor in fragment behaves identical to $ref/An array of numbers is valid`
- `draft2020-12/dynamicRef/A $dynamicRef without anchor in fragment behaves identical to $ref/An array of strings is invalid`
- `draft2020-12/dynamicRef/A $ref to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor/An array containing non-strings is invalid`
- `draft2020-12/dynamicRef/A $ref to a $dynamicAnchor in the same schema resource behaves like a normal $ref to an $anchor/An array of strings is valid`

</details>

<details>
<summary>$dynamicAnchor is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /$defs/list/$defs/items) (1 test)</summary>

- `draft2020-12/dynamicRef/An $anchor with the same name as a $dynamicAnchor is not used for dynamic scope resolution/Any array is valid`

</details>

<details>
<summary>$dynamicAnchor is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /$defs/second/$defs/length) (2 tests)</summary>

- `draft2020-12/dynamicRef/$dynamicRef avoids the root of each schema, but scopes are still registered/data is not sufficient for schema at second#/$defs/length`
- `draft2020-12/dynamicRef/$dynamicRef avoids the root of each schema, but scopes are still registered/data is sufficient for schema at second#/$defs/length`

</details>

<details>
<summary>$dynamicAnchor is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /allOf/0/$defs/elements) (3 tests)</summary>

- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $ref first/correct extended schema`
- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $ref first/incorrect extended schema`
- `draft2020-12/dynamicRef/$ref and $dynamicAnchor are independent of order - $ref first/incorrect parent schema`

</details>

<details>
<summary>$dynamicAnchor is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at root) (5 tests)</summary>

- `draft2020-12/dynamicRef/A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor resolves to the first $dynamicAnchor in the dynamic scope/The recursive part is not valid against the root`
- `draft2020-12/dynamicRef/A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor resolves to the first $dynamicAnchor in the dynamic scope/The recursive part is valid against the root`
- `draft2020-12/dynamicRef/A $dynamicRef that initially resolves to a schema without a matching $dynamicAnchor behaves like a normal $ref to $anchor/The recursive part doesn't need to validate against the root`
- `draft2020-12/dynamicRef/strict-tree schema, guards against misspelled properties/instance with correct field`
- `draft2020-12/dynamicRef/strict-tree schema, guards against misspelled properties/instance with misspelled field`

</details>

<details>
<summary>$dynamicRef is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /$defs/baseSchema) (4 tests)</summary>

- `draft2020-12/unevaluatedItems/unevaluatedItems with $dynamicRef/with no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with $dynamicRef/with unevaluated items`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with $dynamicRef/with no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with $dynamicRef/with unevaluated properties`

</details>

<details>
<summary>$dynamicRef is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /$defs/start) (3 tests)</summary>

- `draft2020-12/dynamicRef/after leaving a dynamic scope, it is not used by a $dynamicRef//then/$defs/thingy is the final stop for the $dynamicRef`
- `draft2020-12/dynamicRef/after leaving a dynamic scope, it is not used by a $dynamicRef/first_scope is not in dynamic scope for the $dynamicRef`
- `draft2020-12/dynamicRef/after leaving a dynamic scope, it is not used by a $dynamicRef/string matches /$defs/thingy, but the $dynamicRef does not stop here`

</details>

<details>
<summary>$dynamicRef is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /items) (2 tests)</summary>

- `draft2020-12/dynamicRef/A $dynamicRef to an $anchor in the same schema resource behaves like a normal $ref to an $anchor/An array containing non-strings is invalid`
- `draft2020-12/dynamicRef/A $dynamicRef to an $anchor in the same schema resource behaves like a normal $ref to an $anchor/An array of strings is valid`

</details>

<details>
<summary>$dynamicRef is not supported: Dynamic references ($dynamicRef/$dynamicAnchor) require runtime scope tracking (at /properties/false) (2 tests)</summary>

- `draft2020-12/dynamicRef/$dynamicRef points to a boolean schema/follow $dynamicRef to a false schema`
- `draft2020-12/dynamicRef/$dynamicRef points to a boolean schema/follow $dynamicRef to a true schema`

</details>

<details>
<summary>unevaluatedItems is not supported: inside applicator subschema cannot see sibling annotations (cousins problem) (at /allOf/1) (1 test)</summary>

- `draft2020-12/unevaluatedItems/unevaluatedItems can't see inside cousins/always fails`

</details>

<details>
<summary>unevaluatedItems is not supported: requires annotation tracking when combined with applicators (prefixItems, contains, allOf, anyOf, oneOf, if) (at /properties/foo) (2 tests)</summary>

- `draft2020-12/unevaluatedItems/item is evaluated in an uncle schema to unevaluatedItems/no extra items`
- `draft2020-12/unevaluatedItems/item is evaluated in an uncle schema to unevaluatedItems/uncle keyword evaluation is not significant`

</details>

<details>
<summary>unevaluatedItems is not supported: requires annotation tracking when combined with applicators (prefixItems, contains, allOf, anyOf, oneOf, if) (at root) (52 tests)</summary>

- `draft2020-12/unevaluatedItems/Evaluated items collection needs to consider instance location/with an unevaluated item that exists at another location`
- `draft2020-12/unevaluatedItems/unevaluatedItems and contains interact to control item dependency relationship/a's and b's are valid`
- `draft2020-12/unevaluatedItems/unevaluatedItems and contains interact to control item dependency relationship/a's, b's and c's are valid`
- `draft2020-12/unevaluatedItems/unevaluatedItems and contains interact to control item dependency relationship/empty array is valid`
- `draft2020-12/unevaluatedItems/unevaluatedItems and contains interact to control item dependency relationship/only a's and c's are invalid`
- `draft2020-12/unevaluatedItems/unevaluatedItems and contains interact to control item dependency relationship/only a's are valid`
- `draft2020-12/unevaluatedItems/unevaluatedItems and contains interact to control item dependency relationship/only b's and c's are invalid`
- `draft2020-12/unevaluatedItems/unevaluatedItems and contains interact to control item dependency relationship/only b's are invalid`
- `draft2020-12/unevaluatedItems/unevaluatedItems and contains interact to control item dependency relationship/only c's are invalid`
- `draft2020-12/unevaluatedItems/unevaluatedItems before $ref/with no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems before $ref/with unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems can see annotations from if without then and else/invalid in case if is evaluated`
- `draft2020-12/unevaluatedItems/unevaluatedItems can see annotations from if without then and else/valid in case if is evaluated`
- `draft2020-12/unevaluatedItems/unevaluatedItems depends on adjacent contains/contains fails, second item is not evaluated`
- `draft2020-12/unevaluatedItems/unevaluatedItems depends on adjacent contains/contains passes, second item is not evaluated`
- `draft2020-12/unevaluatedItems/unevaluatedItems depends on adjacent contains/second item is evaluated by contains`
- `draft2020-12/unevaluatedItems/unevaluatedItems depends on multiple nested contains/5 not evaluated, passes unevaluatedItems`
- `draft2020-12/unevaluatedItems/unevaluatedItems depends on multiple nested contains/7 not evaluated, fails unevaluatedItems`
- `draft2020-12/unevaluatedItems/unevaluatedItems with $ref/with no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with $ref/with unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with anyOf/when one schema matches and has no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with anyOf/when one schema matches and has unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with anyOf/when two schemas match and has no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with anyOf/when two schemas match and has unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with boolean schemas/with no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with boolean schemas/with unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with if/then/else/when if doesn't match and it has no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with if/then/else/when if doesn't match and it has unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with if/then/else/when if matches and it has no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with if/then/else/when if matches and it has unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with items and prefixItems/unevaluatedItems doesn't apply`
- `draft2020-12/unevaluatedItems/unevaluatedItems with items/invalid under items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with items/valid under items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with minContains = 0/all items evaluated by contains`
- `draft2020-12/unevaluatedItems/unevaluatedItems with minContains = 0/empty array is valid`
- `draft2020-12/unevaluatedItems/unevaluatedItems with minContains = 0/no items evaluated by contains`
- `draft2020-12/unevaluatedItems/unevaluatedItems with minContains = 0/some but not all items evaluated by contains`
- `draft2020-12/unevaluatedItems/unevaluatedItems with nested items/with invalid additional item`
- `draft2020-12/unevaluatedItems/unevaluatedItems with nested items/with no additional items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with nested items/with only (valid) additional items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with nested prefixItems and items/with additional items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with nested prefixItems and items/with no additional items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with nested tuple/with no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with nested tuple/with unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with nested unevaluatedItems/with additional items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with nested unevaluatedItems/with no additional items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with not/with unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with oneOf/with no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with oneOf/with unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with tuple/with no unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with tuple/with unevaluated items`
- `draft2020-12/unevaluatedItems/unevaluatedItems with uniform items/unevaluatedItems doesn't apply`

</details>

<details>
<summary>unevaluatedProperties is not supported: cyclic $ref with unevaluatedProperties requires recursive annotation tracking (at root) (7 tests)</summary>

- `draft2020-12/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Deep nested is valid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Empty is valid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Nested is valid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Single is valid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Unevaluated on 1st level is invalid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Unevaluated on 2nd level is invalid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + single cyclic ref/Unevaluated on 3rd level is invalid`

</details>

<details>
<summary>unevaluatedProperties is not supported: inside applicator subschema cannot see sibling annotations (cousins problem) (at /allOf/0) (8 tests)</summary>

- `draft2020-12/unevaluatedProperties/cousin unevaluatedProperties, true and false, false with properties/with nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/cousin unevaluatedProperties, true and false, false with properties/with no nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/cousin unevaluatedProperties, true and false, true with properties/with nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/cousin unevaluatedProperties, true and false, true with properties/with no nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/in-place applicator siblings, allOf has unevaluated/base case: both properties present`
- `draft2020-12/unevaluatedProperties/in-place applicator siblings, allOf has unevaluated/in place applicator siblings, bar is missing`
- `draft2020-12/unevaluatedProperties/in-place applicator siblings, allOf has unevaluated/in place applicator siblings, foo is missing`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties can't see inside cousins (reverse order)/always fails`

</details>

<details>
<summary>unevaluatedProperties is not supported: inside applicator subschema cannot see sibling annotations (cousins problem) (at /allOf/1) (1 test)</summary>

- `draft2020-12/unevaluatedProperties/unevaluatedProperties can't see inside cousins/always fails`

</details>

<details>
<summary>unevaluatedProperties is not supported: inside applicator subschema cannot see sibling annotations (cousins problem) (at /anyOf/0) (3 tests)</summary>

- `draft2020-12/unevaluatedProperties/in-place applicator siblings, anyOf has unevaluated/base case: both properties present`
- `draft2020-12/unevaluatedProperties/in-place applicator siblings, anyOf has unevaluated/in place applicator siblings, bar is missing`
- `draft2020-12/unevaluatedProperties/in-place applicator siblings, anyOf has unevaluated/in place applicator siblings, foo is missing`

</details>

<details>
<summary>unevaluatedProperties is not supported: requires annotation tracking when combined with applicators (allOf, anyOf, oneOf, if, $ref, dependentSchemas, not) (at /not) (2 tests)</summary>

- `draft2020-12/not/collect annotations inside a 'not', even if collection is disabled/annotations are still collected inside a 'not'`
- `draft2020-12/not/collect annotations inside a 'not', even if collection is disabled/unevaluated property`

</details>

<details>
<summary>unevaluatedProperties is not supported: requires annotation tracking when combined with applicators (allOf, anyOf, oneOf, if, $ref, dependentSchemas, not) (at root) (77 tests)</summary>

- `draft2020-12/unevaluatedProperties/dependentSchemas with unevaluatedProperties/unevaluatedProperties doesn't consider dependentSchemas`
- `draft2020-12/unevaluatedProperties/dependentSchemas with unevaluatedProperties/unevaluatedProperties doesn't see bar when foo2 is absent`
- `draft2020-12/unevaluatedProperties/dependentSchemas with unevaluatedProperties/unevaluatedProperties sees bar when foo2 is present`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/Empty is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/a + b is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/a + c is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/a + d is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/a is valid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/all + a is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/all + foo is valid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/all is valid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/b + c is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/b + d is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/b is valid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/c + d is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/c is valid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/d is valid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/xx + a is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/xx + b is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/xx + c is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/xx + d is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/xx + foo is invalid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/xx + foox is valid`
- `draft2020-12/unevaluatedProperties/dynamic evalation inside nested refs/xx is valid`
- `draft2020-12/unevaluatedProperties/nested unevaluatedProperties, outer false, inner true, properties inside/with nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/nested unevaluatedProperties, outer false, inner true, properties inside/with no nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/nested unevaluatedProperties, outer false, inner true, properties outside/with nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/nested unevaluatedProperties, outer false, inner true, properties outside/with no nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/nested unevaluatedProperties, outer true, inner false, properties inside/with nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/nested unevaluatedProperties, outer true, inner false, properties inside/with no nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/nested unevaluatedProperties, outer true, inner false, properties outside/with nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/nested unevaluatedProperties, outer true, inner false, properties outside/with no nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/Empty is invalid (no x or y)`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and b and x and y are invalid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and b and x are valid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and b and y are valid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and b are invalid (no x or y)`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and x are valid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/a and y are valid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties + ref inside allOf / oneOf/x and y are invalid`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties before $ref/with no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties before $ref/with unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties can see annotations from if without then and else/invalid in case if is evaluated`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties can see annotations from if without then and else/valid in case if is evaluated`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with $ref/with no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with $ref/with unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with anyOf/when one matches and has no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with anyOf/when one matches and has unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with anyOf/when two match and has no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with anyOf/when two match and has unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with boolean schemas/with no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with boolean schemas/with unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with dependentSchemas/with no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with dependentSchemas/with unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else, else not defined/when if is false and has no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else, else not defined/when if is false and has unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else, else not defined/when if is true and has no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else, else not defined/when if is true and has unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else, then not defined/when if is false and has no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else, then not defined/when if is false and has unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else, then not defined/when if is true and has no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else, then not defined/when if is true and has unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else/when if is false and has no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else/when if is false and has unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else/when if is true and has no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with if/then/else/when if is true and has unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with nested additionalProperties/with additional properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with nested additionalProperties/with no additional properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with nested patternProperties/with additional properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with nested patternProperties/with no additional properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with nested properties/with additional properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with nested properties/with no additional properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with nested unevaluatedProperties/with nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with nested unevaluatedProperties/with no nested unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with not/with unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with oneOf/with no unevaluated properties`
- `draft2020-12/unevaluatedProperties/unevaluatedProperties with oneOf/with unevaluated properties`

</details>

### Unexpected Failures

<details>
<summary>defs - 2 failures</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/defs/group_0: encountered unresolved non-local $ref "https://json-schema.org/draft/2020-12/schema"`
- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/defs/group_0: encountered unresolved non-local $ref "https://json-schema.org/draft/2020-12/schema"`

</details>

<details>
<summary>ref - 26 failures</summary>

- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_11: recursive local $ref "#" is not supported`
- **Recursive references between schemas**
  - Test: valid tree
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_11: recursive local $ref "#" is not supported`
- **URN ref with nested pointer ref**
  - Test: a non-string is invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_28: failed to resolve local $ref "#/$defs/bar": key "bar" not found`
- **URN ref with nested pointer ref**
  - Test: a string is valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_28: failed to resolve local $ref "#/$defs/bar": key "bar" not found`
- **order of evaluation: $id and $anchor and $ref**
  - Test: data is valid against first definition
  - Expected: `valid`, Got: `false`
- **order of evaluation: $id and $ref on nested schema**
  - Test: data is invalid against nested sibling
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json"`
- **order of evaluation: $id and $ref on nested schema**
  - Test: data is valid against nested sibling
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_20: encountered unresolved non-local $ref "./bar.json"`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_7: $ref value must be a string`
- **property named $ref that is not a reference**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_7: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_8: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_8: $ref value must be a string`
- **ref creates new scope when adjacent to keywords**
  - Test: referenced subschema doesn't see annotations from properties
  - Expected: `invalid`, Got: `true`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_15: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_15: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **refs with relative uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_15: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_16: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_16: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_16: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_6: encountered unresolved non-local $ref "https://json-schema.org/draft/2020-12/schema"`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_6: encountered unresolved non-local $ref "https://json-schema.org/draft/2020-12/schema"`
- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_0: recursive local $ref "#" is not supported`
- **simple URN base URI with $ref via the URN**
  - Test: invalid under the URN IDed schema
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_21: recursive local $ref "#" is not supported`
- **simple URN base URI with $ref via the URN**
  - Test: valid under the URN IDed schema
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/ref/group_21: recursive local $ref "#" is not supported`

</details>

<details>
<summary>refRemote - 4 failures</summary>

- **anchor within remote ref**
  - Test: remote anchor invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_2: recursive local $ref "#/$defs/localhost_1234_draft2020_12_locationIndependentIdentifier_json__A" is not supported`
- **anchor within remote ref**
  - Test: remote anchor valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_2: recursive local $ref "#/$defs/localhost_1234_draft2020_12_locationIndependentIdentifier_json__A" is not supported`
- **fragment within remote ref**
  - Test: remote fragment invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_1: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported`
- **fragment within remote ref**
  - Test: remote fragment valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft2020-12/refRemote/group_1: recursive local $ref "#/$defs/localhost_1234_draft2020_12_subSchemas_json__integer" is not supported`

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
| ref | ⚠️ | 16/27 |
| refRemote | ✅ | 8/8 |
| required | ✅ | 4/4 |
| type | ✅ | 80/80 |
| uniqueItems | ✅ | 62/62 |

### Unexpected Failures

<details>
<summary>ref - 11 failures</summary>

- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `true`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `false`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft3/ref/group_6: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft3/ref/group_6: $ref value must be a string`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft3/ref/group_8: encountered unresolved non-local $ref "http://json-schema.org/draft-03/schema#"`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft3/ref/group_8: encountered unresolved non-local $ref "http://json-schema.org/draft-03/schema#"`
- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft3/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft3/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft3/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft3/ref/group_0: recursive local $ref "#" is not supported`

</details>

## draft4

| Keyword | Status | Pass/Total |
| ------- | ------ | ---------- |
| additionalItems | ✅ | 17/17 |
| additionalProperties | ✅ | 16/16 |
| allOf | ✅ | 27/27 |
| anyOf | ✅ | 15/15 |
| default | ✅ | 7/7 |
| definitions | ❌ | 0/2 |
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
| properties | ✅ | 24/24 |
| ref | ⚠️ | 30/45 |
| refRemote | ✅ | 17/17 |
| required | ✅ | 15/15 |
| type | ✅ | 79/79 |
| uniqueItems | ✅ | 69/69 |

### Unexpected Failures

<details>
<summary>definitions - 2 failures</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/definitions/group_0: encountered unresolved non-local $ref "http://json-schema.org/draft-04/schema#"`
- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/definitions/group_0: encountered unresolved non-local $ref "http://json-schema.org/draft-04/schema#"`

</details>

<details>
<summary>ref - 15 failures</summary>

- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `true`
- **$ref prevents a sibling id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `false`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_10: recursive local $ref "#/$defs/node" is not supported`
- **Recursive references between schemas**
  - Test: valid tree
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_10: recursive local $ref "#/$defs/node" is not supported`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_8: $ref value must be a string`
- **property named $ref that is not a reference**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_8: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_9: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_9: $ref value must be a string`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_7: encountered unresolved non-local $ref "http://json-schema.org/draft-04/schema#"`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_7: encountered unresolved non-local $ref "http://json-schema.org/draft-04/schema#"`
- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft4/ref/group_0: recursive local $ref "#" is not supported`

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
| definitions | ❌ | 0/2 |
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
| ref | ⚠️ | 47/70 |
| refRemote | ✅ | 23/23 |
| required | ✅ | 16/16 |
| type | ✅ | 80/80 |
| uniqueItems | ✅ | 69/69 |

### Unexpected Failures

<details>
<summary>definitions - 2 failures</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/definitions/group_0: encountered unresolved non-local $ref "http://json-schema.org/draft-06/schema#"`
- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/definitions/group_0: encountered unresolved non-local $ref "http://json-schema.org/draft-06/schema#"`

</details>

<details>
<summary>ref - 23 failures</summary>

- **$ref prevents a sibling $id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `true`
- **$ref prevents a sibling $id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `false`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_12: recursive local $ref "#" is not supported`
- **Recursive references between schemas**
  - Test: valid tree
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_12: recursive local $ref "#" is not supported`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_8: $ref value must be a string`
- **property named $ref that is not a reference**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_8: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_9: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_9: $ref value must be a string`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_18: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_18: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **refs with relative uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_18: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_19: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_19: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_19: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_7: encountered unresolved non-local $ref "http://json-schema.org/draft-06/schema#"`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_7: encountered unresolved non-local $ref "http://json-schema.org/draft-06/schema#"`
- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_0: recursive local $ref "#" is not supported`
- **simple URN base URI with $ref via the URN**
  - Test: invalid under the URN IDed schema
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_20: recursive local $ref "#" is not supported`
- **simple URN base URI with $ref via the URN**
  - Test: valid under the URN IDed schema
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft6/ref/group_20: recursive local $ref "#" is not supported`

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
| definitions | ❌ | 0/2 |
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
| ref | ⚠️ | 55/78 |
| refRemote | ✅ | 23/23 |
| required | ✅ | 16/16 |
| type | ✅ | 80/80 |
| uniqueItems | ✅ | 69/69 |

### Unexpected Failures

<details>
<summary>definitions - 2 failures</summary>

- **validate definition against metaschema**
  - Test: invalid definition schema
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/definitions/group_0: encountered unresolved non-local $ref "http://json-schema.org/draft-07/schema#"`
- **validate definition against metaschema**
  - Test: valid definition schema
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/definitions/group_0: encountered unresolved non-local $ref "http://json-schema.org/draft-07/schema#"`

</details>

<details>
<summary>ref - 23 failures</summary>

- **$ref prevents a sibling $id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data does not validate
  - Expected: `invalid`, Got: `true`
- **$ref prevents a sibling $id from changing the base uri**
  - Test: $ref resolves to /definitions/base_foo, data validates
  - Expected: `valid`, Got: `false`
- **Recursive references between schemas**
  - Test: invalid tree
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported`
- **Recursive references between schemas**
  - Test: valid tree
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_12: recursive local $ref "#" is not supported`
- **naive replacement of $ref with its destination is not correct**
  - Test: match the enum exactly
  - Expected: `valid`, Got: `false`
- **property named $ref that is not a reference**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_8: $ref value must be a string`
- **property named $ref that is not a reference**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_8: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_9: $ref value must be a string`
- **property named $ref, containing an actual $ref**
  - Test: property named $ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_9: $ref value must be a string`
- **refs with relative uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_18: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **refs with relative uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_18: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **refs with relative uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_18: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on inner field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_19: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: invalid on outer field
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_19: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **relative refs with absolute uris and defs**
  - Test: valid on both fields
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_19: failed to resolve local $ref "#/$defs/inner": key "$defs" not found`
- **remote ref, containing refs itself**
  - Test: remote ref invalid
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_7: encountered unresolved non-local $ref "http://json-schema.org/draft-07/schema#"`
- **remote ref, containing refs itself**
  - Test: remote ref valid
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_7: encountered unresolved non-local $ref "http://json-schema.org/draft-07/schema#"`
- **root pointer ref**
  - Test: match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive match
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_0: recursive local $ref "#" is not supported`
- **root pointer ref**
  - Test: recursive mismatch
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_0: recursive local $ref "#" is not supported`
- **simple URN base URI with $ref via the URN**
  - Test: invalid under the URN IDed schema
  - Expected: `invalid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_21: recursive local $ref "#" is not supported`
- **simple URN base URI with $ref via the URN**
  - Test: valid under the URN IDed schema
  - Expected: `valid`, Got: `error: bundling error: failed to resolve internal refs for compliance://draft7/ref/group_21: recursive local $ref "#" is not supported`

</details>

