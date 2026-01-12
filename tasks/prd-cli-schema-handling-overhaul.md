# PRD: CLI Schema Handling Overhaul

## Introduction

Fix critical JSON Schema spec compliance issues and performance bottlenecks in the bundler, processor, validator, and vocabulary modules. The current implementation has incorrect RFC 6901 unescaping order, overly aggressive scope resets, incomplete anchor path rewriting, sequential external ref fetching, redundant caching, and quadratic flattening complexity.

## Goals

- Fix RFC 6901 JSON pointer unescaping order (spec compliance)
- Fix scope handling so `$id` only affects descendants, not siblings
- Complete anchor path rewriting for deeply nested anchors
- Fix vocabulary filtering to respect nested `$vocabulary` declarations
- Parallelize processor's external ref crawl/fetch phase
- Cache compiled schemas in validator
- Unify all schema caching in fetcher module (retriever, processor, metaschema)
- Fix quadratic complexity in `$defs` flattening
- Early validation of declared schemas before crawling external refs
- Support multiple output files per language (for Python, Rust, etc.)

## User Stories

### US-001: Fix RFC 6901 unescaping order

**Description:** As a user with schemas containing special characters in refs, I want refs to resolve correctly so my schemas bundle properly.

**Acceptance Criteria:**

- [ ] JSON pointer unescape (`~1` → `/`, `~0` → `~`) happens before URI decode
- [ ] Test case: ref with `%2F` (encoded slash) and `~1` resolves correctly
- [ ] Existing bundler tests pass
- [ ] Zod adapter compliance still 100% (`bun run compliance` in zod adapter folder)
- [ ] Typecheck/lint passes

### US-002: Fix scope handling after $id

**Description:** As a user with nested `$id` declarations, I want refs in sibling properties to resolve correctly so complex schemas work.

**Acceptance Criteria:**

- [ ] `$id` creates new scope only for its descendants
- [ ] Sibling properties retain parent scope
- [ ] Test case: schema with `$id` in one property, `$ref` to parent path in sibling property
- [ ] Existing bundler tests pass
- [ ] Zod adapter compliance still 100% (`bun run compliance` in zod adapter folder)
- [ ] Typecheck/lint passes

### US-003: Fix anchor path rewriting after flattening

**Description:** As a user with schemas using `$anchor` in external refs, I want anchors to resolve correctly after bundling flattens `$defs`.

**Acceptance Criteria:**

- [ ] Anchors in external schemas rewrite to correct flattened paths
- [ ] Deeply nested anchors (3+ levels) rewrite correctly
- [ ] Test case: external schema with nested `$defs` containing `$anchor`
- [ ] Existing bundler tests pass
- [ ] Zod adapter compliance still 100% (`bun run compliance` in zod adapter folder)
- [ ] Typecheck/lint passes

### US-004: Implement vocabulary scope tracking

**Description:** As a user with nested `$vocabulary` declarations, I want keyword filtering to respect each scope so validation keywords aren't incorrectly stripped.

**Acceptance Criteria:**

- [ ] Vocabulary filtering uses scope stack
- [ ] Nested `$vocabulary` overrides parent for its subtree
- [ ] Test case: parent disables validation vocab, child re-enables it, `minLength` preserved in child
- [ ] Existing vocabulary tests pass
- [ ] Zod adapter compliance still 100% (`bun run compliance` in zod adapter folder)
- [ ] Typecheck/lint passes

### US-005: Parallelize processor crawl/fetch

**Description:** As a user with schemas having many external refs, I want faster bundling so generation completes quickly.

**Acceptance Criteria:**

- [ ] Frontier URIs fetched in parallel using errgroup
- [ ] Respects existing concurrency limit from options
- [ ] Wave-based: collect frontier → parallel fetch → extract refs → repeat
- [ ] Test case: schema with 5+ external refs, verify all fetched (mock fetcher tracks call count)
- [ ] Existing processor tests pass
- [ ] Typecheck/lint passes

### US-006: Cache compiled schemas in validator

**Description:** As a developer, I want the validator to cache compiled schemas so repeated validations of the same schema are fast.

**Acceptance Criteria:**

- [ ] Compiled schemas cached by content hash or URI
- [ ] Cache is thread-safe (sync.Map or mutex)
- [ ] Cache hit skips recompilation
- [ ] Test case: validate same schema twice, second call doesn't recompile (mock or instrument compiler)
- [ ] Test case: validate two different schemas, both compiled separately
- [ ] Test case: concurrent validations of same schema are safe (no race with -race flag)
- [ ] Existing validator tests pass
- [ ] Typecheck/lint passes

### US-007: Unify schema caching in fetcher module

**Description:** As a developer, I want a single cache for all schema fetching so no schema is fetched twice regardless of how it's referenced.

**Acceptance Criteria:**

- [ ] `fetcher.Cache` is the single source of truth for all fetched schemas
- [ ] Retriever reads/writes to `fetcher.Cache` when fetching declared schemas
- [ ] Processor.crawl reads/writes to same cache (checks before fetch, stores after fetch)
- [ ] Metaschema reads/writes to same cache (checks before fetch, stores after fetch)
- [ ] No duplicate HTTP requests for any URI (declared, external ref, or metaschema)
- [ ] Test case: declared schema also referenced via `$ref`, fetched only once
- [ ] Test case: custom `$schema` URI fetched only once across all schemas
- [ ] Existing retriever, processor, and metaschema tests pass
- [ ] Typecheck/lint passes

### US-008: Fix quadratic $defs flattening

**Description:** As a user with deeply nested `$defs`, I want bundling to complete in linear time so complex schemas don't hang.

**Acceptance Criteria:**

- [ ] Flattening is O(n) where n = total nodes, not O(d²) where d = nesting depth
- [ ] Single-pass: build flat structure and rewrite refs in one traversal
- [ ] Test case: schema with 5+ levels of nested `$defs`, all refs resolve correctly after flattening
- [ ] Test case: benchmark with 10 levels of nesting, verify linear scaling (not quadratic)
- [ ] Existing bundler tests pass
- [ ] Zod adapter compliance still 100% (`bun run compliance` in zod adapter folder)
- [ ] Typecheck/lint passes

### US-009: Early validation of declared schemas

**Description:** As a user, I want invalid schemas to fail fast before crawling external refs so I don't waste time on doomed operations.

**Acceptance Criteria:**

- [ ] Declared schemas validated immediately after retrieval, before processor.crawlAndFetch
- [ ] Invalid declared schema fails with clear error before any external ref fetching
- [ ] Valid schemas proceed to crawl phase as normal
- [ ] Test case: invalid declared schema (e.g., `{"type": "invalid"}`) fails before any `$ref` in it is fetched
- [ ] Test case: valid schema with external refs proceeds to crawl
- [ ] Existing processor tests pass
- [ ] Typecheck/lint passes

### US-010: Support multiple output files per language

**Description:** As a language implementer, I want to define multiple output files with separate templates so languages like Python and Rust can generate the required file structure.

**Acceptance Criteria:**

- [ ] `Language` struct supports `OutputFiles []OutputFileSpec` instead of single `OutputFile`/`Template`
- [ ] `OutputFileSpec` has `Path` (string) and `Template` (string)
- [ ] Injector iterates over `OutputFiles`, executes each template with same `TemplateData`
- [ ] Each template produces one `GeneratedFile`, all written to disk
- [ ] Manifest tracks all generated files for cleanup
- [ ] Backward compatible: TypeScript migrated to use `OutputFiles` with single entry
- [ ] Test case: language with 3 output files generates all 3 correctly
- [ ] Test case: stale files from previous run cleaned up correctly
- [ ] Existing injector tests pass
- [ ] TypeScript still works: Zod adapter compliance still 100% (`bun run compliance` in zod adapter folder)
- [ ] Typecheck/lint passes

## Functional Requirements

- FR-1: In `bundler.go`, apply JSON pointer unescaping (`~1`→`/`, `~0`→`~`) before URI decoding
- FR-2: Track scope as a tree/stack structure where `$id` pushes new scope for descendants only
- FR-3: When rewriting anchor paths after flattening, resolve anchor location relative to flattened structure
- FR-4: In `vocabulary.go`, maintain vocabulary stack and push/pop when encountering `$vocabulary`
- FR-5: In `processor.go` `crawlAndFetch`, batch frontier URIs and fetch with `errgroup.SetLimit()`
- FR-6: In `validator.go`, cache compiled `*jsonschema.Schema` by schema hash
- FR-7: Make `fetcher.Cache` the unified cache; all modules (retriever, processor.crawl, metaschema) read and write to it
- FR-8: In `bundler.go` `flattenDefs`, collect all nested defs and build rewrite map in single pass
- FR-9: Validate declared schemas after retrieval, before `crawlAndFetch` begins
- FR-10: Replace `Language.OutputFile`/`Template` with `OutputFiles []OutputFileSpec`; injector loops over array

## Non-Goals

- No support for `$dynamicRef`/`$dynamicAnchor` (already documented as unsupported)
- No changes to compliance command (caching, harness per group, temp file retention)
- No changes to generator module

## Technical Considerations

- `golang.org/x/sync/errgroup` already in go.mod, use for parallel fetching
- Scope tracking needs to handle both object traversal and ref resolution contexts
- Vocabulary stack must handle missing `$vocabulary` (inherit from parent)
- Anchor rewriting must account for key collision suffixes (`_2`, `_3`) from flattening
- Compiled schema cache: use content hash (SHA256 of JSON bytes) as key
- Unified cache: single `fetcher.Cache` instance created at start, all modules read/write to it
- Linear flattening: build `oldPath → newPath` map during single DFS traversal
- Early validation: call `validator.ValidateSchema` on each declared schema before processor.Process()
- Multi-file output: `OutputFileSpec{Path, Template}` array; empty template = empty file (for markers like `py.typed`)

## Success Metrics

- All existing tests pass
- New test cases for each fix pass
- Schemas with 10+ external refs bundle 3-5x faster (parallel fetch)
- Repeated validation of same schema 10x+ faster (compiled cache)
- Deeply nested schemas (5+ levels) bundle in O(n) time
- No duplicate HTTP requests for any schema URI (unified cache)
- Invalid declared schemas fail before any external ref is fetched
- Languages can define multiple output files (tested with mock 3-file language)
- No regressions in JSON Schema Test Suite compliance results
