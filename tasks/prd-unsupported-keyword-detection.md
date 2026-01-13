# PRD: Centralized Unsupported Keyword Detection

## Introduction

Move unsupported keyword detection from TypeScript adapters to the Go CLI, creating a single source of truth. The Go `unsupported` package already exists but isn't integrated into the pipeline. This change removes fragile TS checks, adds automatic compliance test skipping based on schema content, and provides clear error messages when users try to generate code from unsupported schemas.

## Goals

- Single source of truth for unsupported features in `cli/unsupported/`
- Remove keyword checks from TypeScript parser
- Compliance automatically skips tests containing unsupported keywords
- `xschema generate` errors early with clear messages for unsupported schemas
- Context-aware detection for `unevaluatedProperties`/`unevaluatedItems` (support standalone, unsupport with applicators)
- Keywords-only configuration (remove hardcoded test paths)

## User Stories

### US-001: Define UnsupportedKeywordError type

**Description:** As a developer, I need a typed error for unsupported keywords so callers can distinguish it from other errors.

**Acceptance Criteria:**

- [ ] Add `UnsupportedKeywordError` struct in `cli/unsupported/unsupported.go`
- [ ] Contains: Keyword string, Reason string, Path string (location in schema)
- [ ] Implements `error` interface with descriptive message
- [ ] `ValidateKeywords()` returns `*UnsupportedKeywordError` (not generic error)
- [ ] Typecheck passes (`go vet ./...`)

### US-002: Add context-aware unevaluated detection

**Description:** As a user, I want `unevaluatedProperties` to work when used standalone (no applicators) but error when combined with applicators.

**Acceptance Criteria:**

- [ ] `ValidateKeywords()` checks for unevaluated keywords at each schema level
- [ ] `unevaluatedProperties` standalone: allowed (no error)
- [ ] `unevaluatedProperties` with applicators (allOf, anyOf, oneOf, if, $ref, dependentSchemas, not): returns `UnsupportedKeywordError`
- [ ] Same logic for `unevaluatedItems` with array applicators (prefixItems, contains)
- [ ] Detection is recursive (checks nested schemas)
- [ ] Unit tests cover: standalone allowed, with-applicators rejected, nested cases
- [ ] `go test ./unsupported/` passes

### US-003: Update unsupported-features.json to keywords-only

**Description:** As a maintainer, I want unsupported features defined by keywords only, not hardcoded test paths.

**Acceptance Criteria:**

- [ ] Remove all entries from `tests` arrays in `unsupported-features.json`
- [ ] Add `unevaluated-with-applicators` group with keywords: `["unevaluatedProperties", "unevaluatedItems"]` and `context: "with-applicators"` flag
- [ ] Add `context` field to `FeatureGroup` struct to indicate context-aware detection
- [ ] Remove `ContainsTest()` method (no longer needed)
- [ ] Remove `TestPaths()` and `TestCount()` methods
- [ ] Update tests to not rely on test paths
- [ ] `go test ./unsupported/` passes

### US-004: Integrate validation into processor

**Description:** As a user running `xschema generate`, I want early errors when my schema contains unsupported keywords.

**Acceptance Criteria:**

- [ ] `processor.Process()` calls `unsupported.ValidateKeywords()` on each schema
- [ ] Validation happens AFTER bundling (so we validate the resolved schema)
- [ ] `UnsupportedKeywordError` propagates up to caller unchanged
- [ ] Other processor errors remain unchanged
- [ ] Existing processor tests pass
- [ ] `go test ./processor/` passes

### US-005: Compliance catches UnsupportedKeywordError and skips

**Description:** As a compliance runner, I want tests with unsupported keywords automatically skipped (not failed).

**Acceptance Criteria:**

- [ ] `bundleSchemas()` in `cli/compliance/runner.go` catches `UnsupportedKeywordError`
- [ ] When caught, mark group as unsupported (not failed)
- [ ] Add to `summary.UnsupportedFeatures` with keyword and reason
- [ ] All tests in that group counted as unsupported
- [ ] No adapter call for unsupported groups
- [ ] Existing compliance flow unchanged for supported schemas
- [ ] `go test ./compliance/` passes

### US-006: Remove TypeScript parser checks

**Description:** As a maintainer, I want the TS parser to not check for unsupported keywords since Go handles it.

**Acceptance Criteria:**

- [ ] Remove `unevaluatedItems` check from `typescript/packages/core/src/parser/index.ts`
- [ ] Remove `unevaluatedProperties` check (lines 55-68)
- [ ] `bun run build` passes (from `typescript/` directory)
- [ ] `bun run typecheck` passes

### US-007: Establish compliance baseline (BEFORE removing TS checks)

**Description:** As a maintainer, I need to capture current compliance state before changes to detect regressions.

**Acceptance Criteria:**

- [ ] Run `bun run compliance` on ALL adapters BEFORE any changes
- [ ] Save/note the current passed/failed/unsupported counts per adapter per draft
- [ ] Adapters: zod, arktype, valibot, effect
- [ ] This establishes the baseline for regression detection

### US-008: Verify compliance exhaustiveness (AFTER all changes)

**Description:** As a maintainer, I want compliance to pass with 0 unexpected failures after these changes.

**Acceptance Criteria:**

- [ ] Run `bun run compliance` on ALL adapters:
  - `typescript/packages/adapters/zod/`
  - `typescript/packages/adapters/arktype/`
  - `typescript/packages/adapters/valibot/`
  - `typescript/packages/adapters/effect/`
- [ ] All drafts show failed=0 in summary for each adapter
- [ ] Unsupported count includes dynamic refs, recursive refs, and unevaluated-with-applicators
- [ ] No regressions: passed count >= baseline, failed count <= baseline
- [ ] Any new failures investigated - either fix or add to unsupported with justification

### US-009: Unit tests for keyword detection

**Description:** As a developer, I need comprehensive tests for the keyword detection logic.

**Acceptance Criteria:**

- [ ] Test: `$dynamicRef` detected at root
- [ ] Test: `$dynamicRef` detected when nested in properties
- [ ] Test: `unevaluatedProperties` standalone allowed
- [ ] Test: `unevaluatedProperties` + allOf rejected
- [ ] Test: `unevaluatedProperties` + $ref rejected
- [ ] Test: `unevaluatedItems` standalone allowed
- [ ] Test: `unevaluatedItems` + prefixItems rejected
- [ ] Test: supported schema passes validation
- [ ] `go test ./unsupported/ -v` passes

### US-010: Opposite compliance test

**Description:** As a maintainer, I want a test that verifies unsupported schemas actually produce errors.

**Acceptance Criteria:**

- [ ] Add test in `cli/compliance/` or `cli/unsupported/`
- [ ] Test loads schemas known to be unsupported (from JSON Schema Test Suite)
- [ ] Calls `unsupported.ValidateKeywords()` on each
- [ ] Asserts all return `UnsupportedKeywordError`
- [ ] Covers: dynamicRef, recursiveRef, unevaluatedProperties+allOf
- [ ] `go test ./...` passes

## Functional Requirements

- FR-1: `UnsupportedKeywordError` type with Keyword, Reason, Path fields
- FR-2: `ValidateKeywords(schema any) *UnsupportedKeywordError` returns typed error or nil
- FR-3: Context-aware detection: `unevaluatedProperties`/`unevaluatedItems` only unsupported when combined with applicators at same schema level
- FR-4: Applicators for properties: allOf, anyOf, oneOf, if, $ref, dependentSchemas, not
- FR-5: Applicators for items: prefixItems, contains (and potentially allOf/anyOf/oneOf on array schemas)
- FR-6: `unsupported-features.json` uses keywords array only, no test paths
- FR-7: Processor validates schemas after bundling, propagates `UnsupportedKeywordError`
- FR-8: Compliance catches `UnsupportedKeywordError` in `bundleSchemas()`, marks as unsupported
- FR-9: TypeScript parser has no keyword validation (deferred to Go)

## Non-Goals

- No changes to adapter protocol
- No support for `$dynamicRef`/`$dynamicAnchor` (fundamentally incompatible with static codegen)
- No support for `$recursiveRef`/`$recursiveAnchor` (same limitation)
- No support for `unevaluatedProperties`/`unevaluatedItems` WITH applicators (requires annotation tracking)
- No context-aware detection beyond unevaluated keywords (other keywords are either always supported or always unsupported)

## Risks and Mitigations

| Risk | Mitigation |
|------|------------|
| Keyword detection misses semantic edge cases | US-007 baseline + US-008 comparison catches regressions as failures |
| Some tests unsupported for non-keyword reasons | If discovered, can add test path fallback (keep `tests` array capability but empty by default) |
| Adapters behave differently | Run compliance on ALL adapters, not just zod |
| TS removal breaks something unexpected | Remove TS checks AFTER Go detection is verified working |

## Implementation Order

Critical: Some stories must be done in sequence to avoid breaking things.

```
US-007 (baseline)           # FIRST: capture current state
    ↓
US-001 (error type)         # Can start immediately
US-002 (context detection)  # Depends on US-001
US-003 (json config)        # Can parallel with US-001/002
US-009 (unit tests)         # After US-001, US-002, US-003
    ↓
US-004 (processor)          # After detection logic works
US-005 (compliance catch)   # After processor integration
US-010 (opposite test)      # After US-005
    ↓
US-006 (remove TS)          # ONLY after Go path fully working
    ↓
US-008 (final compliance)   # Verify detection works
    ↓
US-011 (cleanup)            # Remove deprecated methods
    ↓
US-012 (zod adapter)        # Fix adapters in parallel
US-013 (arktype adapter)    # Can run in parallel
US-014 (valibot adapter)    # Can run in parallel
US-015 (effect adapter)     # Can run in parallel
    ↓
US-016 (web app)            # After adapters fixed
US-017 (verbose output)     # Independent, can be done anytime
```

## Technical Design

### Architecture Decision: JSON Config vs Code

**Decision:** JSON config defines WHAT is unsupported, Go code defines HOW to detect it.

- `unsupported-features.json`: Declares keyword groups and reasons (data)
- `unsupported.go`: Contains detection logic, including context-aware rules (code)

The JSON doesn't need to encode detection logic (like "check for sibling applicators"). That complexity lives in Go code. The JSON is for:
1. Listing always-unsupported keywords (dynamic refs, recursive refs)
2. Documenting reasons for each group
3. Future: could add test path overrides if keyword detection isn't enough

### Error Flow Diagram

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              PROCESSOR PIPELINE                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  schemas ──► normalize ──► validateDeclared ──► crawlAndFetch ──► bundle   │
│                                                                      │      │
│                                                          ┌───────────┘      │
│                                                          ▼                  │
│                                              ┌─────────────────────┐        │
│                                              │ ValidateKeywords()  │        │
│                                              │  (NEW - after bundle)│        │
│                                              └──────────┬──────────┘        │
│                                                         │                   │
│                              ┌───────────────┬──────────┴──────────┐        │
│                              ▼               ▼                     ▼        │
│                           nil (ok)    *UnsupportedKeywordError   other err  │
│                              │               │                     │        │
└──────────────────────────────┼───────────────┼─────────────────────┼────────┘
                               │               │                     │
                               ▼               ▼                     ▼
                    ┌──────────────────────────────────────────────────────────┐
                    │                        CALLER                             │
                    ├──────────────────────────────────────────────────────────┤
                    │                                                          │
                    │  IF caller == "xschema generate":                        │
                    │      UnsupportedKeywordError → print error, exit 1       │
                    │      other error → print error, exit 1                   │
                    │                                                          │
                    │  IF caller == "compliance runner":                       │
                    │      UnsupportedKeywordError → mark group unsupported    │
                    │      other error → mark group failed                     │
                    │                                                          │
                    └──────────────────────────────────────────────────────────┘
```

### Struct Definitions

#### cli/unsupported/unsupported.go

```go
// UnsupportedKeywordError is returned when a schema contains keywords
// that cannot be statically compiled. Callers can type-assert to handle
// this differently from other errors (e.g., compliance skips vs generate fails).
type UnsupportedKeywordError struct {
    Keyword string // the unsupported keyword (e.g., "$dynamicRef")
    Reason  string // why it's unsupported
    Path    string // JSON pointer to location in schema (e.g., "/properties/foo")
}

func (e *UnsupportedKeywordError) Error() string {
    loc := "root"
    if e.Path != "" {
        loc = e.Path
    }
    return fmt.Sprintf("unsupported keyword %q at %s: %s", e.Keyword, loc, e.Reason)
}

// FeatureGroup represents a group of unsupported features (existing, unchanged)
type FeatureGroup struct {
    Name        string   `json:"name"`
    Reason      string   `json:"reason"`
    Keywords    []string `json:"keywords"`
    Explanation string   `json:"explanation,omitempty"` // existing field
}
```

#### cli/compliance/runner.go (line ~363)

```go
// bundledGroup holds a test group with its pre-bundled schema
type bundledGroup struct {
    group          TestGroup
    bundledSchema  RawSchema
    bundleErr      error                              // existing: other errors
    unsupportedErr *unsupported.UnsupportedKeywordError // NEW: unsupported keyword error
}
```

### New JSON Config Structure

**File:** `cli/unsupported/unsupported-features.json`

```json
[
  {
    "name": "dynamic-refs",
    "reason": "Dynamic references require runtime scope tracking",
    "keywords": ["$dynamicRef", "$dynamicAnchor"],
    "explanation": "... existing explanation ..."
  },
  {
    "name": "recursive-refs", 
    "reason": "Recursive references require runtime scope tracking",
    "keywords": ["$recursiveRef", "$recursiveAnchor"],
    "explanation": "... existing explanation ..."
  },
  {
    "name": "unevaluated-with-applicators",
    "reason": "unevaluatedProperties/Items with applicators requires annotation tracking",
    "keywords": []
  }
]
```

**Note:** `unevaluated-with-applicators` has empty `keywords` because detection is context-aware (handled in code, not by keyword presence). The JSON documents the group exists; the code implements detection.

### Exact Code Locations

#### 1. cli/unsupported/unsupported.go

**Add after line 22** (after FeatureGroup struct):
```go
// UnsupportedKeywordError struct definition (see above)
```

**Replace ValidateKeywords() at line 94-99** with new implementation:
```go
// ValidateKeywords checks a schema for unsupported keywords.
// Returns *UnsupportedKeywordError if found, nil otherwise.
// Detection is context-aware: some keywords (like unevaluatedProperties)
// are only unsupported when combined with applicators.
func ValidateKeywords(schema any) *UnsupportedKeywordError {
    return validateNode(schema, "")
}
```

**Replace validateNode() and validateObject() at lines 101-132** with:
```go
func validateNode(node any, path string) *UnsupportedKeywordError {
    switch v := node.(type) {
    case map[string]any:
        return validateObject(v, path)
    case []any:
        for i, item := range v {
            if err := validateNode(item, fmt.Sprintf("%s/%d", path, i)); err != nil {
                return err
            }
        }
    }
    return nil
}

func validateObject(obj map[string]any, path string) *UnsupportedKeywordError {
    // 1. Check always-unsupported keywords (from JSON config)
    for keyword, reason := range keywords {
        if _, ok := obj[keyword]; ok {
            return &UnsupportedKeywordError{
                Keyword: keyword,
                Reason:  reason,
                Path:    path,
            }
        }
    }

    // 2. Context-aware: unevaluatedProperties with applicators
    if _, has := obj["unevaluatedProperties"]; has {
        if hasPropertyApplicators(obj) {
            return &UnsupportedKeywordError{
                Keyword: "unevaluatedProperties",
                Reason:  "unevaluatedProperties with applicators requires annotation tracking",
                Path:    path,
            }
        }
    }

    // 3. Context-aware: unevaluatedItems with applicators
    if _, has := obj["unevaluatedItems"]; has {
        if hasItemApplicators(obj) {
            return &UnsupportedKeywordError{
                Keyword: "unevaluatedItems",
                Reason:  "unevaluatedItems with applicators requires annotation tracking",
                Path:    path,
            }
        }
    }

    // 4. Recurse into nested schemas
    for k, v := range obj {
        if err := validateNode(v, path+"/"+k); err != nil {
            return err
        }
    }
    return nil
}

// Property applicators that create evaluation context for unevaluatedProperties
func hasPropertyApplicators(obj map[string]any) bool {
    applicators := []string{"allOf", "anyOf", "oneOf", "if", "$ref", "dependentSchemas", "not"}
    for _, a := range applicators {
        if _, exists := obj[a]; exists {
            return true
        }
    }
    return false
}

// Item applicators that create evaluation context for unevaluatedItems  
func hasItemApplicators(obj map[string]any) bool {
    applicators := []string{"prefixItems", "contains", "allOf", "anyOf", "oneOf", "if"}
    for _, a := range applicators {
        if _, exists := obj[a]; exists {
            return true
        }
    }
    return false
}
```

**Remove these methods** (no longer needed):
- `ContainsTest()` at lines 64-74
- `TestPaths()` at lines 76-83
- `TestCount()` at lines 85-92

#### 2. cli/processor/processor.go

**Add import** at line 20:
```go
"github.com/xschemadev/xschema/unsupported"
```

**Add validation in bundleAll() after line 263** (after vocabulary filtering, before building result):
```go
        // Validate for unsupported keywords (after bundling so we check resolved schema)
        var schemaMap map[string]any
        if err := json.Unmarshal(bundled, &schemaMap); err == nil {
            if unsupportedErr := unsupported.ValidateKeywords(schemaMap); unsupportedErr != nil {
                return nil, fmt.Errorf("schema %s contains unsupported keywords: %w", s.SourceURI, unsupportedErr)
            }
        }
```

#### 3. cli/compliance/runner.go

**Update bundledGroup struct** at line 362-367:
```go
type bundledGroup struct {
    group          TestGroup
    bundledSchema  RawSchema
    bundleErr      error
    unsupportedErr *unsupported.UnsupportedKeywordError // ADD THIS LINE
}
```

**Update bundleSchemas()** at lines 501-524. Replace error handling:
```go
        processed, err := processor.Process(ctx, toProcess, processor.Options{
            Fetcher: localhostFetcher,
            Draft:   opts.draft,
        })
        opts.timing.addSchemaBundling(time.Since(bundleStart))

        if err != nil {
            // Check if it's an unsupported keyword error
            var unsupportedErr *unsupported.UnsupportedKeywordError
            if errors.As(err, &unsupportedErr) {
                bundled[i] = bundledGroup{
                    group:          group,
                    unsupportedErr: unsupportedErr,
                }
                continue
            }
            // Other errors are real failures
            bundled[i] = bundledGroup{group: group, bundleErr: err}
            continue
        }
```

**Update markBundleErrors()** at lines 529-536 to handle unsupported separately:
```go
func markBundleErrors(bundled []bundledGroup, groupFilters []groupFilter, keywordResult *KeywordResult, summary *DraftSummary) {
    for i, bg := range bundled {
        if groupFilters[i].allKnown {
            continue
        }
        
        // Handle unsupported keyword errors (mark as unsupported, not failed)
        if bg.unsupportedErr != nil {
            for _, tc := range bg.group.Tests {
                testPath := fmt.Sprintf("%s/%s/%s", keywordResult.Keyword, bg.group.Description, tc.Description)
                summary.UnsupportedFeatures.Count++
                summary.UnsupportedFeatures.Items = append(summary.UnsupportedFeatures.Items, UnsupportedFeatureItem{
                    Path:   testPath,
                    Reason: bg.unsupportedErr.Reason,
                })
            }
            continue
        }
        
        // Handle other bundle errors (mark as failed)
        if bg.bundleErr != nil {
            markAllFailed(keywordResult, summary, bg.group, fmt.Sprintf("bundling error: %v", bg.bundleErr))
        }
    }
}
```

#### 4. typescript/packages/core/src/parser/index.ts

**Delete lines 51-68** (the entire unevaluated check block):
```typescript
// DELETE THIS BLOCK:
	// Fail on unevaluated keywords - they require annotation tracking we can't do statically
	if (schema.unevaluatedItems !== undefined) {
		throw new Error("unevaluatedItems is not supported");
	}
	if (schema.unevaluatedProperties !== undefined) {
		const hasApplicators =
			schema.allOf ||
			schema.anyOf ||
			schema.oneOf ||
			schema.if ||
			schema.$ref ||
			schema.dependentSchemas ||
			schema.not;
		if (hasApplicators) {
			throw new Error(
				"unevaluatedProperties with applicators is not supported",
			);
		}
	}
```

### Key Files Summary

| File | Line(s) | Change |
|------|---------|--------|
| `cli/unsupported/unsupported.go` | 22+ | Add `UnsupportedKeywordError` struct |
| `cli/unsupported/unsupported.go` | 94-132 | Replace `ValidateKeywords()` with context-aware version |
| `cli/unsupported/unsupported.go` | 64-92 | Remove `ContainsTest()`, `TestPaths()`, `TestCount()` |
| `cli/unsupported/unsupported-features.json` | all | Remove `tests` arrays, add unevaluated group |
| `cli/processor/processor.go` | 20 | Add unsupported import |
| `cli/processor/processor.go` | ~263 | Add `ValidateKeywords()` call after vocabulary filtering |
| `cli/compliance/runner.go` | 362-367 | Add `unsupportedErr` field to `bundledGroup` |
| `cli/compliance/runner.go` | 501-524 | Catch `UnsupportedKeywordError` in `bundleSchemas()` |
| `cli/compliance/runner.go` | 529-536 | Update `markBundleErrors()` to handle unsupported |
| `typescript/.../parser/index.ts` | 51-68 | Delete unevaluated checks |

### US-012: Fix zod adapter for standalone unevaluated keywords

**Description:** As a user, I want the zod adapter to properly handle standalone unevaluatedProperties/unevaluatedItems.

**Acceptance Criteria:**

- [ ] `unevaluatedProperties: false` → generates `z.object({}).strict()` or equivalent
- [ ] `unevaluatedProperties: { type: "string" }` → generates `z.object({}).catchall(z.string())`
- [ ] `{ properties: {...}, unevaluatedProperties: false }` → generates `.strict()` on the object
- [ ] `unevaluatedItems: false` → generates tuple that rejects extra items
- [ ] `unevaluatedItems: { type: "string" }` → generates tuple with string rest
- [ ] Run `bun run compliance` from `typescript/packages/adapters/zod/`
- [ ] Verify unevaluatedProperties/unevaluatedItems tests that don't involve applicators now pass
- [ ] No regressions in other tests

### US-013: Fix arktype adapter for standalone unevaluated keywords

**Description:** As a user, I want the arktype adapter to properly handle standalone unevaluatedProperties/unevaluatedItems.

**Acceptance Criteria:**

- [ ] Same patterns as US-012 but using arktype syntax
- [ ] Run `bun run compliance` from `typescript/packages/adapters/arktype/`
- [ ] Verify unevaluatedProperties/unevaluatedItems standalone tests pass
- [ ] No regressions in other tests

### US-014: Fix valibot adapter for standalone unevaluated keywords

**Description:** As a user, I want the valibot adapter to properly handle standalone unevaluatedProperties/unevaluatedItems.

**Acceptance Criteria:**

- [ ] Same patterns as US-012 but using valibot syntax
- [ ] Run `bun run compliance` from `typescript/packages/adapters/valibot/`
- [ ] Verify unevaluatedProperties/unevaluatedItems standalone tests pass
- [ ] No regressions in other tests (valibot baseline had 50 failures)

### US-015: Fix effect adapter for standalone unevaluated keywords

**Description:** As a user, I want the effect adapter to properly handle standalone unevaluatedProperties/unevaluatedItems.

**Acceptance Criteria:**

- [ ] Same patterns as US-012 but using effect schema syntax
- [ ] Run `bun run compliance` from `typescript/packages/adapters/effect/`
- [ ] Verify unevaluatedProperties/unevaluatedItems standalone tests pass
- [ ] No regressions in other tests

### US-016: Update web app for new unsupported features format

**Description:** As a user viewing the docs, I want the unsupported features page to display correctly with the new keyword-based format.

**Acceptance Criteria:**

- [ ] Update `web/scripts/generate-unsupported-features.ts` to read from `cli/unsupported/unsupported-features.json`
- [ ] Update script to use `keywords` array instead of `tests` array
- [ ] Generate page showing keyword groups with explanations
- [ ] Show which keywords trigger each limitation (e.g., "$dynamicRef, $dynamicAnchor")
- [ ] Verify in browser using chrome-devtools that the page renders correctly
- [ ] `bun run generate:schemas` (or equivalent) succeeds
- [ ] No build errors in web app

### US-017: Make unsupported features output verbose-only

**Description:** As a CLI user, I want the detailed unsupported features list to only show with --verbose flag.

**Acceptance Criteria:**

- [ ] Move "Unsupported Features:" print block in `cli/cmd/compliance.go` behind verbose check
- [ ] Summary line still shows unsupported count (e.g., "181 unsupported")
- [ ] With `--verbose`, show the detailed breakdown by reason
- [ ] `go build -o xschema . && go vet ./...` passes
- [ ] Test: run compliance without --verbose, confirm no "Unsupported Features:" section
- [ ] Test: run compliance with --verbose, confirm "Unsupported Features:" section appears

## Success Metrics

- Compliance shows 0 failed tests across all drafts for zod, arktype, effect
- Valibot shows ≤50 failed tests (baseline failures, not unevaluated-related)
- Unsupported count correctly includes all dynamic/recursive/unevaluated-with-applicators tests
- `xschema generate` with unsupported schema shows clear error message
- No TS parser changes needed when adding new unsupported keywords
- Web app unsupported features page renders correctly with new format
- CLI output is clean by default, detailed with --verbose
