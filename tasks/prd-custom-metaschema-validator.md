# PRD: Linear Pipeline Architecture & Full Compliance

## Introduction

Refactor the CLI to use a linear pipeline architecture with clear state flow. Create a Processor layer that orchestrates validation and bundling with pre-fetched external refs. This makes the bundler a pure function (no I/O), improves testability, and surfaces errors early. Also adds `--verbose` flag for detailed CLI output. Achieves 100% compliance on all supported tests.

## Goals

- Linear pipeline: Parser → Retriever → Processor → Generator → Injector
- Processor iteratively crawls and fetches external refs until none remain
- Bundler becomes pure (no network calls, uses pre-fetched cache)
- Unified metaschema fetching (shared cache)
- `--verbose` flag for detailed step-by-step CLI output
- Extract and respect `$vocabulary` for selective keyword disabling
- Achieve 100% compliance for zod, effect, arktype (valibot: all except JS prototype issues)

## Architecture

### Generate Command Flow

```
Parser → Retriever → Processor → Generator → Injector
                         │
                         ▼
              ┌─────────────────────┐
              │     Processor       │
              │                     │
              │  ┌───────────────┐  │
              │  │ CrawlAndFetch │◄─┼──┐ (iterative until no new URIs)
              │  │  (iterative)  │──┼──┘
              │  └───────┬───────┘  │
              │          ▼          │
              │  ┌───────────────┐  │
              │  │   Validate    │  │  (all schemas: declared + external)
              │  └───────┬───────┘  │
              │          ▼          │
              │  ┌───────────────┐  │
              │  │ Bundle (pure) │  │  (no I/O, uses cache)
              │  └───────────────┘  │
              └─────────────────────┘
```

**Processor phases:**
1. `crawlAndFetch()` - iteratively find and fetch external $ref URIs until no new ones found
2. `validateAll()` - validate all schemas (declared + external) just before bundling
3. `bundleAll()` - pure bundling using pre-fetched cache

### Compliance Command Flow

```
LoadTestSuite → ProcessKeyword (per keyword, parallel) → Report
                      │
     ┌────────────────┼────────────────┐
     ▼                ▼                ▼
FilterUnsupported  BundleSchemas   GenerateHarness
                      │                │
                      ▼                ▼
              CallAdapterBatch   ExecuteHarness
                      │                │
                      └───────┬────────┘
                              ▼
                      ProcessResults
```

**Detailed compliance flow:**
```
xschema compliance --adapter zod
    │
    ▼
Fetch test suite from GitHub (cached)
    │
    ▼
For each draft (draft2020-12, draft7, etc.):
    │
    ▼
For each keyword (parallel, bounded concurrency):
    │
    ├─► Phase 1: Filter unsupported tests
    │   └─ Check against unsupported-features.json
    │
    ├─► Phase 2: Bundle schemas
    │   └─ bundler.Bundle() with localhost:1234 → remotes/ mapping
    │
    ├─► Phase 3: Batch adapter call
    │   └─ stdin: [{namespace, id, varName, schema}, ...]
    │   └─ stdout: [{schema, type, validate, imports}, ...]
    │
    ├─► Phase 4: Generate harness
    │   └─ Merge imports + embed test data + validation calls
    │   └─ Create temp .ts file
    │
    ├─► Phase 5: Execute harness
    │   └─ bun run xschema-harness-xxx.ts
    │   └─ Output: [{groupId, index, expected, actual}, ...]
    │
    └─► Phase 6: Process results
        └─ Map back to test cases, count pass/fail/skip
    │
    ▼
Generate report (markdown/JSON)
```

**Key difference from generate:**
- Compliance uses bundler directly (not through Processor) because it has special localhost:1234 mapping
- Compliance bundler changes: make it accept a `Fetcher` interface so Processor and compliance can provide different implementations

## User Stories

### US-001: Create Processor package

**Description:** As a developer, I want a Processor layer that orchestrates schema processing so the pipeline is linear and clear.

**Acceptance Criteria:**
- [ ] Create `cli/processor/processor.go`
- [ ] `Process(ctx, schemas, opts) ([]ProcessedSchema, error)`
- [ ] ProcessedSchema contains: Namespace, ID, Schema (bundled), Adapter, SourceURI, Vocabulary
- [ ] Implements 3 phases: crawlAndFetch (iterative) → validate → bundle
- [ ] Accepts Options with Verbose flag and OnProgress callback
- [ ] `go build ./...` passes

### US-002: Implement iterative crawl and fetch

**Description:** As a developer, I want to iteratively crawl and fetch external refs until no new URIs are found.

**Acceptance Criteria:**
- [ ] `crawlAndFetch(ctx, schemas, cache) (map[URI]Schema, error)`
- [ ] Loop: crawl all schemas (declared + cached) for $ref URIs
- [ ] Fetch any URIs not yet in cache (parallel)
- [ ] Repeat until no new URIs found
- [ ] Detect circular refs to avoid infinite loops (track visited URIs)
- [ ] Error if any fetch fails (fail fast)
- [ ] `go build ./...` passes

### US-003: Refactor Bundler to pure with Fetcher interface

**Description:** As a developer, I want the bundler to be a pure function that accepts a Fetcher interface.

**Acceptance Criteria:**
- [ ] Create `Fetcher` interface: `Fetch(uri string) (json.RawMessage, error)`
- [ ] Change Bundle signature to accept `BundleInput` struct with Fetcher
- [ ] Processor provides `CacheFetcher` that looks up in pre-fetched cache
- [ ] Compliance provides `LocalhostFetcher` that maps localhost:1234 → remotes/
- [ ] Remove hardcoded retriever calls from bundler
- [ ] `go build ./...` passes

### US-004: Create Metaschema package

**Description:** As a developer, I want a dedicated metaschema package for shared fetching and vocabulary extraction.

**Acceptance Criteria:**
- [ ] Create `cli/metaschema/metaschema.go`
- [ ] `Get(uri string) (*Metaschema, error)` fetches and caches
- [ ] `ExtractVocabulary(schema any) map[string]bool` extracts $vocabulary
- [ ] `IsStandardDraft(uri string) bool` detects json-schema.org URLs
- [ ] Fetched during crawlAndFetch phase when custom $schema encountered
- [ ] `go build ./...` passes

### US-005: Update Validator for metaschema pre-loading

**Description:** As a developer, I want the validator to accept pre-loaded metaschemas.

**Acceptance Criteria:**
- [ ] Add `ValidateOptions` with Metaschemas map
- [ ] Use `compiler.AddResource()` to pre-load custom metaschemas
- [ ] Handle cross-draft schemas (each resource can have different $schema)
- [ ] Return vocabulary along with validation result
- [ ] `go build ./...` passes

### US-006: Update Generator

**Description:** As a developer, I want Generator to accept ProcessedSchema (already bundled) instead of doing bundling itself.

**Acceptance Criteria:**
- [ ] Accept `[]ProcessedSchema` instead of `[]RetrievedSchema`
- [ ] Remove bundling logic (moved to Processor)
- [ ] Simplify to: group by adapter → call adapter CLI
- [ ] `go build ./...` passes

### US-007: Add --verbose flag

**Description:** As a user, I want `--verbose` to show detailed processing steps.

**Acceptance Criteria:**
- [ ] Add `--verbose` flag to generate command
- [ ] Default output stays compact (current behavior)
- [ ] Verbose shows: crawl iterations, refs found, schemas fetched, validation results, bundling progress
- [ ] Pass OnProgress callback to Processor
- [ ] Works correctly

### US-008: Implement vocabulary-aware keyword skipping

**Description:** As a schema author, I want the TypeScript parser to skip disabled vocabulary keywords.

**Acceptance Criteria:**
- [ ] Add vocabulary-to-keywords mapping in parser
- [ ] When vocabulary disabled, simply omit those keywords from IR (don't add constraints)
- [ ] Example: `{ type: "string", minLength: 5 }` with validation disabled → `StringNode` with no minLength constraint
- [ ] Continue parsing keywords from enabled vocabularies
- [ ] Applicator keywords (`properties`, `items`, etc.) still work when validation disabled
- [ ] `bun run build` passes in typescript/

### US-009: Fix bundler nested $defs flattening

**Description:** As a schema author, I want refs to nested definitions in remote schemas to resolve correctly.

**Acceptance Criteria:**
- [ ] Recursively extract nested `$defs` AND `definitions` (draft-04 keyword) to root level
- [ ] Nested defs get prefixed keys using `__` separator: `remote__foo__bar`
- [ ] Handle collision: if `remote__foo` already exists, append counter: `remote__foo__2`
- [ ] Refs inside flattened defs rewritten to point to flattened locations
- [ ] Fragment refs to definitions work (e.g., `other.json#/$defs/Foo`)
- [ ] URN refs with nested pointer refs work
- [ ] `go build ./...` passes

### US-010: Verify 100% compliance

**Description:** As a maintainer, I want to verify all fixes achieve 100% compliance.

**Acceptance Criteria:**
- [ ] Run `cd typescript/packages/zod && bun run compliance` - 100% pass rate
- [ ] Run `cd typescript/packages/effect && bun run compliance` - 100% pass rate
- [ ] Run `cd typescript/packages/arktype && bun run compliance` - 100% pass rate
- [ ] Valibot excluded from 100% requirement (JS prototype bug is upstream)
- [ ] Only `dynamic-refs` and `recursive-refs` remain in `unsupported-features.json`

## Functional Requirements

- FR-1: Pipeline must be linear: Parser → Retriever → Processor → Generator → Injector
- FR-2: Processor must iteratively crawl and fetch until no new external refs found
- FR-3: Bundler must be pure (uses Fetcher interface, no direct I/O)
- FR-4: Fetch failure must error immediately (fail fast)
- FR-5: Validation happens just before bundling (validates declared + external schemas)
- FR-6: Metaschema package must cache fetched metaschemas
- FR-7: Validator must accept pre-loaded metaschemas via AddResource
- FR-8: Cross-draft schemas supported (each resource can declare own $schema)
- FR-9: `--verbose` flag must show detailed processing steps
- FR-10: Parser must skip keywords whose vocabulary is disabled
- FR-11: Bundler must correctly flatten nested $defs
- FR-12: Compliance command continues using bundler with LocalhostFetcher
- FR-13: All compliance tests must pass except dynamic-refs, recursive-refs, valibot JS prototype

## Non-Goals

- Supporting `$dynamicRef`/`$dynamicAnchor` (requires runtime scope tracking)
- Supporting `$recursiveRef`/`$recursiveAnchor` (requires runtime scope tracking)
- Fixing Valibot's JS prototype property name bug (upstream issue)
- Verbose output by default (keep behind flag)
- Changing compliance command to use Processor (it has special localhost mapping needs)

## Implementation Order

Suggested order (dependencies noted):

1. **US-004: Metaschema package** - no dependencies, can be done first
2. **US-003: Bundler Fetcher interface** - no dependencies, enables US-001
3. **US-009: Nested $defs flattening** - works on existing bundler
4. **US-001: Processor package** - depends on US-003, US-004
5. **US-002: Iterative crawl/fetch** - part of US-001
6. **US-005: Validator metaschema pre-loading** - depends on US-004
7. **US-006: Generator update** - depends on US-001 (needs ProcessedSchema type)
8. **US-007: --verbose flag** - depends on US-001
9. **US-008: Vocabulary keyword skipping** - TypeScript side, can parallel with Go work
10. **US-010: Verify compliance** - last, after all fixes

## Edge Cases

**Circular refs between external schemas:**
A.json refs B.json, B.json refs A.json. Crawler marks URI as visited BEFORE fetching to prevent infinite loops. Both schemas end up in cache.

**Self-referential $ref (`#`):**
Fragment-only ref, always local. Never triggers external fetch.

**$id changes base URI:**
When external schema has `$id: "http://other.com/schema"`, refs inside it resolve against that URI, not the fetch URL. Crawler must track effective base URI per schema.

**Metaschema refs in schemas:**
If schema has `$ref: "https://json-schema.org/draft/2020-12/schema"`, crawler fetches it. Bundler will fail later if metaschema uses unsupported keywords ($recursiveAnchor). This is correct - fail at usage, not spec.

**Empty $vocabulary:**
If metaschema has `$vocabulary: {}`, all vocabularies disabled. Parser should produce minimal IR (just structure, no validation).

## Technical Considerations

**Generate command data flow:**
```
Parser
   ↓ Declarations
Retriever
   ↓ RetrievedSchema[] (raw)
Processor
   ├─ Phase 1: crawlAndFetch() [iterative]
   │   ├─ Iteration 1: crawl declared schemas → find URIs → fetch
   │   ├─ Iteration 2: crawl fetched schemas → find new URIs → fetch
   │   ├─ Iteration N: no new URIs found → done
   │   └─ Also fetch custom metaschemas during crawl
   │
   ├─ Phase 2: validateAll()
   │   └─ Validate declared + all external schemas
   │
   └─ Phase 3: bundleAll()
       └─ Pure bundling with CacheFetcher
   ↓ ProcessedSchema[] (bundled + vocabulary)
Generator
   ↓ ConvertResult[]
Injector
   ↓ files
```

**Compliance command data flow:**
```
LoadTestSuite
   ↓ map[keyword][]TestGroup
For each keyword (parallel):
   ├─ Filter unsupported
   ├─ Bundle with LocalhostFetcher (localhost:1234 → remotes/)
   ├─ Batch adapter call
   ├─ Generate harness (.ts file)
   ├─ Execute harness (bun run)
   └─ Process results
   ↓ KeywordResult[]
ComplianceReport
```

**Key types to create:**

- **ProcessedSchema** (processor package): Contains Namespace, ID, Schema (bundled JSON), Adapter, SourceURI, and Vocabulary map
- **Options** (processor package): Contains Verbose bool and OnProgress callback
- **Fetcher interface** (bundler package): Single method `Fetch(uri) -> (schema, error)`
- **BundleInput** (bundler package): Contains Schema, SourceURI, Fetcher, and Metaschema

**Fetcher implementations:**

- **CacheFetcher** (processor package): Looks up URI in pre-fetched cache map. Returns error if URI not found (indicates crawler bug - everything should be cached).
- **LocalhostFetcher** (compliance package): Maps `http://localhost:1234/*` to local remotes/ directory. Returns error for any other URL (compliance tests should never fetch real external URLs).

**URI resolution (per RFC 3986):**
Refs are resolved against the current base URI:
- Initial base URI = schema's SourceURI
- `$id` keyword changes base URI for its subtree
- Relative refs resolved per RFC 3986: `../common.json` against `http://example.com/schemas/user.json` → `http://example.com/common.json`
- Fragment-only refs (`#/$defs/Foo`, `#anchor`) are local, not fetched
- URI with fragment (`other.json#/$defs/Foo`) → fetch `other.json`, fragment resolved after bundling

**Iterative crawl+fetch algorithm:**

1. Initialize empty cache and visited set
2. Mark all declared schema URIs as visited (we already have their content)
3. Loop until no new URIs found:
   - Crawl all schemas (declared + cached) for external $ref URIs
   - External ref = has scheme (http/https/file) or is relative path, NOT fragment-only (#...)
   - Resolve relative refs against each schema's base URI
   - Collect URIs not yet visited, mark them visited immediately (prevents duplicates)
   - If no new URIs, exit loop
   - Fetch all new URIs in parallel using existing retriever concurrency
   - Fail fast on any fetch error
   - Add fetched schemas to cache
4. Return cache containing all external schemas

**Cross-draft schema handling:**
Per JSON Schema spec, each schema resource can declare its own `$schema`. When a schema references an external schema with a different draft, the validator should switch processing modes. Our validator (santhosh-tekuri/jsonschema) handles this automatically.

**Vocabulary detection:**

The `$vocabulary` object in metaschemas maps vocabulary URIs to booleans (e.g., `"https://json-schema.org/draft/2020-12/vocab/validation": true`).

To check if a vocabulary is enabled:
- Search vocabulary keys for substring match (e.g., "validation", "applicator")
- If found, use its boolean value
- If not found, default to enabled (true)

When vocabulary is disabled, omit its keywords from IR:
- **validation** disabled: skip type, enum, const, minimum, maximum, minLength, maxLength, pattern, multipleOf, etc.
- **applicator** disabled: skip properties, items, allOf, anyOf, oneOf, if/then/else, additionalProperties, etc.

If schema has no `$vocabulary` (standard drafts), treat all vocabularies as enabled.

**CLI output with --verbose:**
```
◉ Parsing config files...
  Found 3 declarations

◉ Retrieving schemas...
  ✓ user/Profile (file)
  ✓ user/Settings (URL)

◉ Processing schemas...
  Crawling external refs...
    Iteration 1: found 2 refs
    ✓ http://example.com/common.json
    ✓ http://example.com/address.json
    Iteration 2: found 1 ref
    ✓ http://example.com/country.json
    Iteration 3: no new refs
  Validating...
    ✓ user/Profile
    ✓ user/Settings
    ✓ 3 external schemas
  Bundling...
    ✓ user/Profile (3 refs embedded)
    ✓ user/Settings (0 refs embedded)

◉ Generating...
```

**Files to create:**
- `cli/processor/processor.go`
- `cli/metaschema/metaschema.go`

**Files to modify:**
- `cli/bundler/bundler.go` (add Fetcher interface, make pure)
- `cli/validator/validator.go` (accept pre-loaded metaschemas)
- `cli/generator/generator.go` (accept ProcessedSchema, remove bundling)
- `cli/cmd/generate.go` (add Processor step, --verbose flag)
- `cli/compliance/runner.go` (use LocalhostFetcher with bundler)
- `typescript/packages/core/src/parser/index.ts` (vocabulary-aware)

**Breaking changes (update all call sites in same PR):**
- `bundler.Bundle(schema, sourceURI)` → `bundler.Bundle(BundleInput)` - update calls in generator and compliance
- Generator input type changes from `RetrievedSchema` to `ProcessedSchema` - update call in generate command

## Success Metrics

- Linear pipeline visible in code and CLI output
- Bundler has zero direct I/O (uses Fetcher interface)
- All network errors surface in crawlAndFetch phase
- Iterative crawl handles arbitrary nesting depth
- 100% compliance pass rate for zod, effect, arktype
- Valibot compliance matches other adapters minus JS prototype failures
- `--verbose` shows clear iteration breakdown
- Compliance command unchanged in behavior (just uses LocalhostFetcher)

## References

- [JSON Schema $ref handling](https://json-schema.org/understanding-json-schema/structuring)
- [Cross-draft schema resources](https://json-schema.org/draft/2020-12/release-notes) - each resource can declare own $schema
