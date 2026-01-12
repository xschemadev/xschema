# PRD: Compliance Improvements

## Introduction

Fix remaining compliance failures in xschema adapters to achieve 100% pass rate on all supported tests. Currently, $ref bundling issues in the Go CLI and missing vocabulary support in TypeScript cause ~14 test failures across all adapters. Valibot has additional failures due to a Valibot library bug with JS prototype property names - this is out of scope.

## Goals

- Achieve 100% compliance on supported tests for zod, effect, and arktype adapters
- Fix $ref bundling issues in CLI bundler (Go)
- Add $vocabulary support in TypeScript parser
- Clean up adapter READMEs to remove redundant limitation docs

## User Stories

### US-001: Skip metaschema bundling

**Description:** As a schema author, I want refs to official JSON Schema metaschemas to work so that I can validate schemas against metaschemas.

**Acceptance Criteria:**
- [ ] Bundler detects refs to `json-schema.org/draft-*` URLs
- [ ] These refs are left as-is, not fetched/bundled
- [ ] No more "unsupported keyword $recursiveAnchor" errors for metaschema refs
- [ ] Typecheck/build passes

### US-002: Flatten nested definitions

**Description:** As a schema author, I want remote schemas with their own `$defs`/`definitions` to be bundled correctly so that refs to nested definitions resolve.

**Acceptance Criteria:**
- [ ] When embedding remote schema, extract its `$defs`/`definitions` to root level
- [ ] Prefix extracted def keys with parent key (e.g., `remote__inner`)
- [ ] Rewrite all refs to point to flattened locations
- [ ] Remove empty `$defs`/`definitions` from embedded schema
- [ ] Typecheck/build passes

### US-003: Fix anchor paths for embedded schemas

**Description:** As a schema author, I want anchors in remote schemas to resolve correctly after embedding so that anchor refs work.

**Acceptance Criteria:**
- [ ] `collectIDsAndAnchors` accepts path prefix parameter
- [ ] Anchors from fetched schemas stored with `/$defs/key` prefix
- [ ] Anchor refs in original schema resolve to correct embedded location
- [ ] Typecheck/build passes

### US-004: Handle fragment refs to flattened definitions

**Description:** As a schema author, I want refs like `remote.json#/definitions/Foo` to work so that I can reference specific definitions in remote schemas.

**Acceptance Criteria:**
- [ ] Detect when fragment points to `$defs/X` or `definitions/X`
- [ ] Rewrite to flattened location `#/$defs/key__X`
- [ ] Non-definition fragments still work (e.g., `#/properties/foo`)
- [ ] Typecheck/build passes

### US-005: Support $vocabulary disabling validation

**Description:** As a schema author using custom metaschemas, I want $vocabulary to be respected so that disabling validation vocabulary works.

**Acceptance Criteria:**
- [ ] Add `$vocabulary` field to JSONSchema TypeScript type
- [ ] Parser checks if validation vocabulary is disabled
- [ ] When disabled, return `{ kind: "any" }` (no validation)
- [ ] Typecheck/build passes

### US-006: Clean up adapter READMEs

**Description:** As a developer, I want READMEs to be concise so that limitations are documented in one place (compliance reports).

**Acceptance Criteria:**
- [ ] Remove "Known Limitations" section from zod README
- [ ] Remove "Known Limitations" section from effect README
- [ ] Remove "Known Limitations" section from arktype README
- [ ] Valibot README keeps ONLY the JS prototype property names limitation
- [ ] All other content in READMEs preserved

### US-007: Verify 100% compliance

**Description:** As a maintainer, I want to verify all fixes work so that compliance is actually 100%.

**Acceptance Criteria:**
- [ ] Run `xschema compliance --adapter zod` - 100% pass rate
- [ ] Run `xschema compliance --adapter effect` - 100% pass rate
- [ ] Run `xschema compliance --adapter arktype` - 100% pass rate
- [ ] Valibot passes all tests except JS prototype issues

## Functional Requirements

- FR-1: Bundler must skip fetching/bundling refs to `http://json-schema.org/draft-*` and `https://json-schema.org/draft/*` URLs
- FR-2: Bundler must extract nested `$defs` and `definitions` from embedded schemas to root level with prefixed keys
- FR-3: Bundler must rewrite all internal refs to point to flattened definition locations
- FR-4: Bundler must prefix anchor paths with embedding location when collecting anchors from fetched schemas
- FR-5: Bundler must handle fragment refs pointing to definitions by rewriting to flattened key
- FR-6: TypeScript parser must check `$vocabulary` and skip validation when validation vocabulary is disabled
- FR-7: Zod, Effect, and ArkType READMEs must not contain "Known Limitations" sections
- FR-8: Valibot README must document only the JS prototype property names limitation

## Non-Goals

- Fixing Valibot's JS prototype property name bug (upstream issue)
- Supporting `$dynamicRef`/`$recursiveRef` (fundamentally incompatible with static codegen)
- Supporting `unevaluatedProperties`/`unevaluatedItems` with applicators (requires runtime tracking)
- Supporting circular `$ref` with constraints (requires runtime lazy evaluation)

## Technical Considerations

- Bundler changes are in Go: `cli/bundler/bundler.go`
- Vocabulary changes are in TypeScript: `typescript/packages/core/src/parser/index.ts` and `typescript/packages/core/src/schema/json-schema.ts`
- Must maintain backward compatibility - existing schemas should still work
- Flattening uses `__` as separator to avoid conflicts with valid def names

## Success Metrics

- 100% compliance pass rate for zod, effect, arktype (excluding unsupported features)
- Valibot compliance matches other adapters minus JS prototype failures (~10 tests)
