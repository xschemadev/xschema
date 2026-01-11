# Zod Adapter Discovery

## Baseline Per Draft

| Draft | Passed | Failed | Coverage |
| ----- | ------ | ------ | -------- |
| draft2020-12 | 869 | 67 | 92.8% |
| draft2019-09 | 857 | 74 | 92.1% |
| draft7 | 830 | 24 | 97.2% |
| draft6 | 754 | 24 | 96.9% |
| draft4 | 562 | 16 | 97.2% |
| draft3 | 396 | 17 | 95.9% |

## Failing Tests List

### draft2020-12 (67 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalProperties | additionalProperties does not look in applicators | properties defined in allOf are not examined |
| anchor | Location-independent identifier | match |
| anchor | Location-independent identifier | mismatch |
| defs | validate definition against metaschema | invalid definition schema |
| defs | validate definition against metaschema | valid definition schema |
| dynamicRef | $dynamicRef points to a boolean schema | follow $dynamicRef to a false schema |
| dynamicRef | $dynamicRef points to a boolean schema | follow $dynamicRef to a true schema |
| dynamicRef | $dynamicRef skips over intermediate resources - direct reference | integer property passes |
| dynamicRef | $dynamicRef skips over intermediate resources - direct reference | string property fails |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $defs first | correct extended schema |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $defs first | incorrect extended schema |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $defs first | incorrect parent schema |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $ref first | correct extended schema |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $ref first | incorrect extended schema |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $ref first | incorrect parent schema |
| dynamicRef | $ref to $dynamicRef finds detached $dynamicAnchor | non-number is invalid |
| dynamicRef | $ref to $dynamicRef finds detached $dynamicAnchor | number is valid |
| dynamicRef | A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor... | The recursive part is not valid against the root |
| dynamicRef | A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor... | The recursive part is valid against the root |
| dynamicRef | A $dynamicRef that initially resolves to a schema without a matching $dynamicAnchor... | The recursive part doesn't need to validate against the root |
| dynamicRef | A $dynamicRef to a $dynamicAnchor in the same schema resource... | An array containing non-strings is invalid |
| dynamicRef | A $dynamicRef to a $dynamicAnchor in the same schema resource... | An array of strings is valid |
| dynamicRef | A $dynamicRef to an $anchor in the same schema resource... | An array containing non-strings is invalid |
| dynamicRef | A $dynamicRef to an $anchor in the same schema resource... | An array of strings is valid |
| dynamicRef | A $ref to a $dynamicAnchor in the same schema resource... | An array containing non-strings is invalid |
| dynamicRef | A $ref to a $dynamicAnchor in the same schema resource... | An array of strings is valid |
| dynamicRef | after leaving a dynamic scope, it is not used by a $dynamicRef | /then/$defs/thingy is the final stop for the $dynamicRef |
| dynamicRef | after leaving a dynamic scope, it is not used by a $dynamicRef | first_scope is not in dynamic scope for the $dynamicRef |
| dynamicRef | after leaving a dynamic scope, it is not used by a $dynamicRef | string matches /$defs/thingy, but the $dynamicRef does not stop here |
| dynamicRef | multiple dynamic paths to the $dynamicRef keyword | number list with number values |
| dynamicRef | multiple dynamic paths to the $dynamicRef keyword | number list with string values |
| dynamicRef | multiple dynamic paths to the $dynamicRef keyword | string list with number values |
| dynamicRef | multiple dynamic paths to the $dynamicRef keyword | string list with string values |
| dynamicRef | strict-tree schema, guards against misspelled properties | instance with correct field |
| dynamicRef | strict-tree schema, guards against misspelled properties | instance with misspelled field |
| dynamicRef | tests for implementation dynamic anchor and reference link | correct extended schema |
| dynamicRef | tests for implementation dynamic anchor and reference link | incorrect extended schema |
| dynamicRef | tests for implementation dynamic anchor and reference link | incorrect parent schema |
| ref | URN ref with nested pointer ref | a non-string is invalid |
| ref | URN ref with nested pointer ref | a string is valid |
| ref | escaped pointer ref | percent invalid |
| ref | escaped pointer ref | percent valid |
| ref | escaped pointer ref | slash invalid |
| ref | escaped pointer ref | slash valid |
| ref | escaped pointer ref | tilde invalid |
| ref | escaped pointer ref | tilde valid |
| ref | order of evaluation: $id and $anchor and $ref | data is invalid against first definition |
| ref | order of evaluation: $id and $anchor and $ref | data is valid against first definition |
| ref | refs with quote | object with numbers is valid |
| ref | refs with quote | object with strings is invalid |
| ref | refs with relative uris and defs | invalid on inner field |
| ref | refs with relative uris and defs | invalid on outer field |
| ref | refs with relative uris and defs | valid on both fields |
| ref | relative refs with absolute uris and defs | invalid on inner field |
| ref | relative refs with absolute uris and defs | invalid on outer field |
| ref | relative refs with absolute uris and defs | valid on both fields |
| ref | remote ref, containing refs itself | remote ref invalid |
| ref | remote ref, containing refs itself | remote ref valid |
| refRemote | $ref to $ref finds detached $anchor | non-number is invalid |
| refRemote | $ref to $ref finds detached $anchor | number is valid |
| refRemote | Location-independent identifier in remote ref | integer is valid |
| refRemote | Location-independent identifier in remote ref | string is invalid |
| refRemote | anchor within remote ref | remote anchor invalid |
| refRemote | anchor within remote ref | remote anchor valid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | number is invalid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | string is valid |
| vocabulary | schema that uses custom metaschema with with no validation vocabulary | no validation: invalid number, but it still validates |

### draft2019-09 (74 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalProperties | additionalProperties does not look in applicators | properties defined in allOf are not examined |
| anchor | Location-independent identifier | match |
| anchor | Location-independent identifier | mismatch |
| defs | validate definition against metaschema | invalid definition schema |
| defs | validate definition against metaschema | valid definition schema |
| recursiveRef | $recursiveRef with $recursiveAnchor: false works like $ref | integer does not match as a property value |
| recursiveRef | $recursiveRef with $recursiveAnchor: false works like $ref | integer matches at the outer level |
| recursiveRef | $recursiveRef with $recursiveAnchor: false works like $ref | single level match |
| recursiveRef | $recursiveRef with $recursiveAnchor: false works like $ref | two levels, integer does not match as a property value |
| recursiveRef | $recursiveRef with $recursiveAnchor: false works like $ref | two levels, properties match with inner definition |
| recursiveRef | $recursiveRef with nesting | integer matches at the outer level |
| recursiveRef | $recursiveRef with nesting | integer now matches as a property value |
| recursiveRef | $recursiveRef with nesting | single level match |
| recursiveRef | $recursiveRef with nesting | two levels, properties match with $recursiveRef |
| recursiveRef | $recursiveRef with nesting | two levels, properties match with inner definition |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the initial target schema resource | leaf node does not match: recursion uses the inner schema |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the initial target schema resource | leaf node does not match; no recursion |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the initial target schema resource | leaf node matches: recursion uses the inner schema |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the outer schema resource | leaf node does not match: recursion only uses inner schema |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the outer schema resource | leaf node does not match; no recursion |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the outer schema resource | leaf node matches: recursion only uses inner schema |
| recursiveRef | $recursiveRef with no $recursiveAnchor works like $ref | integer does not match as a property value |
| recursiveRef | $recursiveRef with no $recursiveAnchor works like $ref | integer matches at the outer level |
| recursiveRef | $recursiveRef with no $recursiveAnchor works like $ref | single level match |
| recursiveRef | $recursiveRef with no $recursiveAnchor works like $ref | two levels, integer does not match as a property value |
| recursiveRef | $recursiveRef with no $recursiveAnchor works like $ref | two levels, properties match with inner definition |
| recursiveRef | $recursiveRef without $recursiveAnchor works like $ref | match |
| recursiveRef | $recursiveRef without $recursiveAnchor works like $ref | mismatch |
| recursiveRef | $recursiveRef without $recursiveAnchor works like $ref | recursive match |
| recursiveRef | $recursiveRef without $recursiveAnchor works like $ref | recursive mismatch |
| recursiveRef | $recursiveRef without using nesting | integer does not match as a property value |
| recursiveRef | $recursiveRef without using nesting | integer matches at the outer level |
| recursiveRef | $recursiveRef without using nesting | single level match |
| recursiveRef | $recursiveRef without using nesting | two levels, no match |
| recursiveRef | $recursiveRef without using nesting | two levels, properties match with inner definition |
| recursiveRef | dynamic $recursiveRef destination (not predictable at schema compile time) | integer node |
| recursiveRef | dynamic $recursiveRef destination (not predictable at schema compile time) | numeric node |
| recursiveRef | multiple dynamic paths to the $recursiveRef keyword | recurse to anyLeafNode - floats are allowed |
| recursiveRef | multiple dynamic paths to the $recursiveRef keyword | recurse to integerNode - floats are not allowed |
| ref | $ref with $recursiveAnchor | extra items allowed for inner arrays |
| ref | $ref with $recursiveAnchor | extra items disallowed for root |
| ref | URN ref with nested pointer ref | a non-string is invalid |
| ref | URN ref with nested pointer ref | a string is valid |
| ref | escaped pointer ref | percent invalid |
| ref | escaped pointer ref | percent valid |
| ref | escaped pointer ref | slash invalid |
| ref | escaped pointer ref | slash valid |
| ref | escaped pointer ref | tilde invalid |
| ref | escaped pointer ref | tilde valid |
| ref | order of evaluation: $id and $anchor and $ref | data is invalid against first definition |
| ref | order of evaluation: $id and $anchor and $ref | data is valid against first definition |
| ref | refs with quote | object with numbers is valid |
| ref | refs with quote | object with strings is invalid |
| ref | refs with relative uris and defs | invalid on inner field |
| ref | refs with relative uris and defs | invalid on outer field |
| ref | refs with relative uris and defs | valid on both fields |
| ref | relative refs with absolute uris and defs | invalid on inner field |
| ref | relative refs with absolute uris and defs | invalid on outer field |
| ref | relative refs with absolute uris and defs | valid on both fields |
| ref | remote ref, containing refs itself | remote ref invalid |
| ref | remote ref, containing refs itself | remote ref valid |
| refRemote | $ref to $ref finds detached $anchor | non-number is invalid |
| refRemote | $ref to $ref finds detached $anchor | number is valid |
| refRemote | Location-independent identifier in remote ref | integer is valid |
| refRemote | Location-independent identifier in remote ref | string is invalid |
| refRemote | anchor within remote ref | remote anchor invalid |
| refRemote | anchor within remote ref | remote anchor valid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | number is invalid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | string is valid |
| unevaluatedItems | unevaluatedItems with $recursiveRef | with no unevaluated items |
| unevaluatedItems | unevaluatedItems with $recursiveRef | with unevaluated items |
| unevaluatedProperties | unevaluatedProperties with $recursiveRef | with no unevaluated properties |
| unevaluatedProperties | unevaluatedProperties with $recursiveRef | with unevaluated properties |
| vocabulary | schema that uses custom metaschema with with no validation vocabulary | no validation: invalid number, but it still validates |

### draft7 (24 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalProperties | additionalProperties does not look in applicators | properties defined in allOf are not examined |
| definitions | validate definition against metaschema | invalid definition schema |
| ref | Location-independent identifier | match |
| ref | Location-independent identifier | mismatch |
| ref | escaped pointer ref | percent invalid |
| ref | escaped pointer ref | percent valid |
| ref | escaped pointer ref | slash invalid |
| ref | escaped pointer ref | slash valid |
| ref | escaped pointer ref | tilde invalid |
| ref | escaped pointer ref | tilde valid |
| ref | refs with quote | object with numbers is valid |
| ref | refs with quote | object with strings is invalid |
| ref | refs with relative uris and defs | invalid on inner field |
| ref | refs with relative uris and defs | invalid on outer field |
| ref | refs with relative uris and defs | valid on both fields |
| ref | relative refs with absolute uris and defs | invalid on inner field |
| ref | relative refs with absolute uris and defs | invalid on outer field |
| ref | relative refs with absolute uris and defs | valid on both fields |
| refRemote | $ref to $ref finds location-independent $id | non-number is invalid |
| refRemote | $ref to $ref finds location-independent $id | number is valid |
| refRemote | Location-independent identifier in remote ref | integer is valid |
| refRemote | Location-independent identifier in remote ref | string is invalid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | number is invalid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | string is valid |

### draft6 (24 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalProperties | additionalProperties does not look in applicators | properties defined in allOf are not examined |
| definitions | validate definition against metaschema | invalid definition schema |
| ref | Location-independent identifier | match |
| ref | Location-independent identifier | mismatch |
| ref | escaped pointer ref | percent invalid |
| ref | escaped pointer ref | percent valid |
| ref | escaped pointer ref | slash invalid |
| ref | escaped pointer ref | slash valid |
| ref | escaped pointer ref | tilde invalid |
| ref | escaped pointer ref | tilde valid |
| ref | refs with quote | object with numbers is valid |
| ref | refs with quote | object with strings is invalid |
| ref | refs with relative uris and defs | invalid on inner field |
| ref | refs with relative uris and defs | invalid on outer field |
| ref | refs with relative uris and defs | valid on both fields |
| ref | relative refs with absolute uris and defs | invalid on inner field |
| ref | relative refs with absolute uris and defs | invalid on outer field |
| ref | relative refs with absolute uris and defs | valid on both fields |
| refRemote | $ref to $ref finds location-independent $id | non-number is invalid |
| refRemote | $ref to $ref finds location-independent $id | number is valid |
| refRemote | Location-independent identifier in remote ref | integer is valid |
| refRemote | Location-independent identifier in remote ref | string is invalid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | number is invalid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | string is valid |

### draft4 (16 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalProperties | additionalProperties does not look in applicators | properties defined in allOf are not examined |
| definitions | validate definition against metaschema | invalid definition schema |
| maximum | exclusiveMaximum validation | boundary point is invalid |
| minimum | exclusiveMinimum validation | boundary point is invalid |
| ref | Location-independent identifier | match |
| ref | Location-independent identifier | mismatch |
| ref | escaped pointer ref | percent invalid |
| ref | escaped pointer ref | percent valid |
| ref | escaped pointer ref | slash invalid |
| ref | escaped pointer ref | slash valid |
| ref | escaped pointer ref | tilde invalid |
| ref | escaped pointer ref | tilde valid |
| ref | refs with quote | object with numbers is valid |
| ref | refs with quote | object with strings is invalid |
| refRemote | Location-independent identifier in remote ref | integer is valid |
| refRemote | Location-independent identifier in remote ref | string is invalid |

### draft3 (17 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalProperties | additionalProperties does not look in applicators | properties defined in extends are not examined |
| dependencies | dependencies | missing dependency |
| maximum | exclusiveMaximum validation | boundary point is invalid |
| minimum | exclusiveMinimum validation | boundary point is invalid |
| ref | escaped pointer ref | percent invalid |
| ref | escaped pointer ref | percent valid |
| ref | escaped pointer ref | slash invalid |
| ref | escaped pointer ref | slash valid |
| ref | escaped pointer ref | tilde invalid |
| ref | escaped pointer ref | tilde valid |
| type | applies a nested schema | an object is invalid otherwise |
| type | types can include schemas | a boolean is invalid |
| type | types can include schemas | a float is invalid |
| type | types can include schemas | a string is invalid |
| type | types can include schemas | an integer is invalid |
| type | types can include schemas | null is invalid |
| type | types from separate schemas are merged | an integer is invalid |

## Categorized Failures

### not-supported (bundler architectural limits)

These failures are caused by bundler limitations that require significant architectural changes. Not planned for implementation.

| Category | Tests | Description |
| -------- | ----- | ----------- |
| $dynamicRef/$dynamicAnchor | 33 | draft2020-12 dynamic reference semantics require runtime evaluation context |
| $recursiveRef/$recursiveAnchor | 36 | draft2019-09 recursive reference semantics require runtime evaluation context |
| $anchor / location-independent $id | ~20 | bundler doesn't resolve `#anchor` fragment refs (e.g., `#foo`) |
| Escaped pointer refs | ~30 | bundler doesn't decode URI-encoded segments in JSON pointers (`percent%25field`) |
| Refs with quotes | ~10 | bundler doesn't handle `%22` (quote) in pointer segments |
| Relative URI with $id scoping | ~18 | bundler flattens schemas, losing `$id`-scoped relative refs |
| URN refs with nested pointers | ~4 | bundler doesn't resolve URN-based refs |
| Nested $defs after remote fetch | ~8 | bundler's key-sanitization creates invalid lookup paths |
| Remote refs containing dynamic keywords | ~6 | fetching JSON Schema metaschema fails due to dynamic keywords |

### core-missing (blocked by core parser/bundler)

| Category | Tests | Description |
| -------- | ----- | ----------- |
| draft3/4 exclusiveMaximum/Minimum boolean | 4 | draft3/4 use `{maximum: 3, exclusiveMaximum: true}` - core normalizer handles this correctly, but compliance tests lack `$schema` field, so version defaults to draft-2020-12 and boolean form isn't converted |

**Root cause**: The adapter protocol (`ConvertInput`) doesn't include draft version. Bundler should inject `$schema` when test schema lacks one.

### adapter-native (fixable in zod adapter)

None identified.

### forced-emulation (non-idiomatic, not recommended)

| Category | Tests | Description | Reason to skip |
| -------- | ----- | ----------- | -------------- |
| additionalProperties + applicators | 6 | requires tracking evaluated properties across allOf/extends | needs unevaluatedProperties semantics, not expressible in Zod |
| validate definition against metaschema | 5 | requires validating schema definitions against metaschema | meta-validation out of scope |
| custom vocabulary (no validation) | 2 | custom metaschema disabling validation vocabulary | edge case, no real use |
| draft3 type with schemas | 7 | draft3 allows `type: [{...schema...}]` | legacy draft3 feature, rarely used |
| draft3 dependencies edge case | 1 | missing dependency check | same root cause as exclusiveMax/Min: no `$schema` in test, version defaults wrong |

## Expected Regressions

None expected. All failures are either:
1. Bundler limitations (not-supported)
2. Edge cases we intentionally don't support
3. Version detection issues (no `$schema` in compliance tests)

## Observations

1. **High baseline compliance**: draft7+ at 97%+, modern drafts at 92%+ despite dynamic ref limitations
2. **Most failures are bundler-side**: 100+ of ~220 total failures are bundler ref-handling issues
3. **Zod adapter itself is solid**: no adapter-specific validation bugs found
4. **draft3/4 failures are version-detection bugs**: the compliance tests don't include `$schema`, so version defaults to draft-2020-12 and boolean exclusiveMax/Min aren't normalized

## Recommended Implementation Stories

Based on this discovery, the following implementation stories would provide the most value:

1. **Bundler: inject $schema for compliance tests** - would fix draft3/4 boolean exclusiveMax/Min failures (~5 tests) and draft3 dependencies
2. **Bundler: URI-decode pointer segments** - would fix ~30 "escaped pointer ref" tests across all drafts
3. **Bundler: resolve $anchor fragments** - would fix ~20 "location-independent identifier" tests

Note: Dynamic/recursive ref support is explicitly out of scope (architectural decision in US-004/US-011).
