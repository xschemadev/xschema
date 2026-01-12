# XSchema CLI - Analysis & Proposed Improvements

## Executive Summary

The CLI is well-architected with clear separation of concerns. However, there are critical JSON Schema spec compliance issues, performance bottlenecks, and several edge cases not handled correctly.

---

## Critical Issues

### 1. Dynamic References Not Supported

**Location:** `bundler/bundler.go:263, 277-279`

**Problem:** `$dynamicRef`, `$dynamicAnchor`, `$recursiveRef`, `$recursiveAnchor` are explicitly rejected.

**Impact:** Breaks draft2020-12 schemas using dynamic references. These are used for extensible schemas where the reference target is determined at validation time.

**Example failing schema:**
```json
{
  "$dynamicRef": "#meta",
  "$defs": {
    "Meta": {"$dynamicAnchor": "meta", "type": "string"}
  }
}
```

**Options:**
1. Implement dynamic reference resolution (complex)
2. Document as unsupported with clear error message (recommended for now)

---

### 2. RFC 6901 JSON Pointer Unescaping Order Wrong

**Location:** `bundler/bundler.go:956-965`

**Problem:** Current order is URI decode → JSON pointer unescape. Spec says JSON pointer unescape first (`~1` → `/`, `~0` → `~`), then URI decode.

**Impact:** Refs with special characters fail silently or resolve incorrectly.

**Fix:**
```go
// Current (wrong):
segment = url.PathUnescape(segment)
segment = strings.ReplaceAll(segment, "~1", "/")
segment = strings.ReplaceAll(segment, "~0", "~")

// Should be:
segment = strings.ReplaceAll(segment, "~1", "/")
segment = strings.ReplaceAll(segment, "~0", "~")
segment = url.PathUnescape(segment)
```

---

### 3. Scope Handling After $id

**Location:** `bundler/bundler.go:286-298`

**Problem:** When `$id` is found, `currentScopePath` is reset to empty. But scope should only reset for descendants, not siblings.

**Example:**
```json
{
  "properties": {
    "a": {"$id": "a.json", "$ref": "#/something"},
    "b": {"$ref": "#/properties/a"}  // This ref might break
  }
}
```

**Impact:** Nested refs inside scoped schemas may not resolve correctly.

---

### 4. Vocabulary Filtering Ignores Nested $vocabulary

**Location:** `vocabulary/vocabulary.go:87-125`

**Problem:** Doesn't respect `$vocabulary` declarations in nested subschemas.

**Example:**
```json
{
  "$vocabulary": {"...validation": false},
  "properties": {
    "nested": {
      "$vocabulary": {"...validation": true},
      "minLength": 1  // Should NOT be filtered, but currently IS
    }
  }
}
```

---

### 5. Anchor Path Rewriting Incomplete

**Location:** `bundler/bundler.go:469-479`

**Problem:** When remote schema has anchor, deeply nested anchors may not rewrite correctly after flattening.

---

## JSON Schema Spec Compliance Issues

| Feature | Status | Issue |
|---------|--------|-------|
| `$dynamicRef` | ✗ Rejected | Not implemented |
| `$dynamicAnchor` | ✗ Rejected | Not implemented |
| `$recursiveRef` | ✗ Rejected | Not implemented |
| `$recursiveAnchor` | ✗ Rejected | Not implemented |
| RFC 6901 escaping | ⚠️ Wrong order | See issue #2 |
| Fragment-only `$id` | ⚠️ Ambiguous | `$id: "#foo"` treated as anchor |
| Scope changes | ⚠️ Overly aggressive | Resets for siblings too |
| Nested `$vocabulary` | ✗ Ignored | Not scoped correctly |

---

## Performance Issues

### 1. Sequential Fetching in Processor Stage 1

**Location:** `processor/processor.go:291-323`

**Problem:** External refs are fetched sequentially in a loop.

**Impact:** Slow for schemas with many external references.

**Fix:** Use `golang.org/x/sync/errgroup` (already in go.mod).

---

### 2. No Compiled Schema Caching in Validator

**Location:** `validator/validator.go`

**Problem:** Every validation recompiles the schema from scratch.

**Impact:** Repeated validations of same schema are slow.

---

### 3. Redundant Metaschema Fetching

**Location:** `processor/processor.go`, `metaschema/metaschema.go`

**Problem:** Processor crawls and caches schemas, but metaschema module has separate cache. Custom metaschema may be fetched twice.

---

### 4. Quadratic Complexity in $defs Flattening

**Location:** `bundler/bundler.go` flattenDefs

**Problem:** Deeply nested `$defs` cause O(d²) rewrite operations where d = nesting depth.

---

## Missing Features

| Feature | Impact | Priority |
|---------|--------|----------|
| Dynamic refs | Breaks draft2020-12 advanced use | HIGH |
| Fetch retry in metaschema | Network hiccups fail immediately | MEDIUM |
| Parallel processor fetching | Performance | MEDIUM |
| Compiled schema cache | Performance | LOW |
| Vocabulary requirement checking | Spec compliance | LOW |
| Bidirectional refs | Rare schemas fail | LOW |

---

## Compliance Command Issues

### 1. No Incremental Caching

**Location:** `compliance/runner.go`

**Problem:** Each run re-bundles all schemas from scratch.

**Impact:** Repeated runs are slow.

### 2. Single Harness Per Keyword

**Location:** `compliance/runner.go:processKeyword`

**Problem:** All groups in keyword go into one harness. If harness execution fails, entire keyword fails.

**Impact:** No partial failure recovery.

### 3. Temp File Cleanup on Failure

**Location:** `compliance/harness.go:591`

**Problem:** `defer os.Remove(tempHarness)` deletes file even on error.

**Impact:** Can't debug failed harness execution.

---

## Proposed Improvements

### Improved Processor Pipeline

```mermaid
flowchart TD
    subgraph Stage1["Stage 1: Parallel Crawl & Fetch"]
        A1[RetrievedSchema array] --> A2[Extract all $refs]
        A2 --> A3[Deduplicate URIs]
        A3 --> A4[Parallel fetch with errgroup]
        A4 --> A5[Stream results to cache]
        A5 --> A6{More refs discovered?}
        A6 -->|yes| A3
        A6 -->|no| A7[Complete cache]
    end

    subgraph Stage2["Stage 2: Validate with Cache"]
        A7 --> B1[Build metaschemas from cache]
        B1 --> B2[Compile metaschemas once]
        B2 --> B3[Parallel validation with compiled]
        B3 --> B4[Collect errors]
    end

    subgraph Stage3["Stage 3: Bundle with Scope Tracking"]
        B4 --> C1[Create scope-aware bundler]
        C1 --> C2[Track $id scope boundaries]
        C2 --> C3[Bundle with correct scope]
        C3 --> C4[Vocabulary-aware filtering]
        C4 --> C5[ProcessedSchema array]
    end

    style Stage1 fill:#e3f2fd
    style Stage2 fill:#fff8e1
    style Stage3 fill:#e8f5e9
```

### Improved Bundler Flow

```mermaid
flowchart TD
    subgraph Init["Improved Initialization"]
        A[Bundle entry] --> B[collectIDsAndAnchors with scope map]
        B --> C[Build scope tree structure]
        C --> D[Validate anchor uniqueness per scope]
    end

    subgraph Process["Scope-Aware Processing"]
        D --> E[processNode with scope context]
        E --> E1{$id found?}
        E1 -->|yes| E2[Push new scope]
        E1 -->|no| E3[Continue current scope]
        E2 --> E4[Process children in new scope]
        E4 --> E5[Pop scope after children]
        E3 --> E6[Process children]
        E5 --> E7[Merge results]
        E6 --> E7
    end

    subgraph RefResolution["Improved Ref Resolution"]
        E7 --> F{Ref type?}
        F -->|$dynamicRef| F1[Track for runtime]
        F -->|$ref local| F2[Resolve in current scope]
        F -->|$ref external| F3[Fetch and bundle recursively]
        F1 --> F4[Mark schema as dynamic]
        F2 --> F5[Rewrite with scope-aware path]
        F3 --> F6[Add to $defs with scope prefix]
    end

    subgraph Flatten["Linear Flattening"]
        F4 --> G[Single-pass flatten]
        F5 --> G
        F6 --> G
        G --> G1[Build flat $defs in one pass]
        G1 --> G2[Update refs during build]
        G2 --> H[Output bundled schema]
    end

    style Init fill:#e3f2fd
    style Process fill:#fff8e1
    style RefResolution fill:#e8f5e9
    style Flatten fill:#f3e5f5
```

### Improved Vocabulary Filtering

```mermaid
flowchart TD
    subgraph Input
        A[Schema + Root vocabulary] --> B[Parse schema]
    end

    subgraph ScopeTracking["Vocabulary Scope Tracking"]
        B --> C[Initialize vocabulary stack]
        C --> D[Push root vocabulary]
        D --> E[Walk schema tree]
    end

    subgraph NodeProcessing["Per-Node Processing"]
        E --> F{Node has $vocabulary?}
        F -->|yes| F1[Push new vocabulary scope]
        F -->|no| F2[Use current scope]
        F1 --> G[Filter keywords by active vocab]
        F2 --> G
        G --> H[Process children recursively]
        H --> I{Was $vocabulary pushed?}
        I -->|yes| I1[Pop vocabulary scope]
        I -->|no| I2[Continue]
        I1 --> J[Return filtered node]
        I2 --> J
    end

    subgraph Output
        J --> K[Reassemble filtered schema]
    end

    style Input fill:#e3f2fd
    style ScopeTracking fill:#fff8e1
    style NodeProcessing fill:#e8f5e9
    style Output fill:#f3e5f5
```

### Improved Compliance Runner

```mermaid
flowchart TD
    subgraph Setup
        A[Load test suite] --> B[Load schema cache if exists]
        B --> C{Cache valid?}
        C -->|yes| D[Use cached bundles]
        C -->|no| E[Bundle all schemas]
        E --> F[Save to cache]
        D --> G[Prepare test execution]
        F --> G
    end

    subgraph Execution["Fault-Tolerant Execution"]
        G --> H[For each keyword parallel]
        H --> I[For each group]
        I --> J[Generate individual harness]
        J --> K{Execute harness}
        K -->|success| L[Record results]
        K -->|failure| M[Preserve harness file]
        M --> N[Mark group as error]
        L --> O[Continue to next group]
        N --> O
    end

    subgraph Aggregation
        O --> P[Aggregate per-group results]
        P --> Q[Build keyword result]
        Q --> R[Build draft result]
    end

    style Setup fill:#e3f2fd
    style Execution fill:#fff8e1
    style Aggregation fill:#e8f5e9
```

---

## Priority Matrix

| Issue | Severity | Effort | Priority |
|-------|----------|--------|----------|
| RFC 6901 order | HIGH | LOW | P0 |
| Scope handling | HIGH | MEDIUM | P0 |
| Dynamic refs (document) | MEDIUM | LOW | P1 |
| Parallel fetching | MEDIUM | MEDIUM | P1 |
| Vocabulary scoping | MEDIUM | HIGH | P2 |
| Anchor path rewriting | MEDIUM | MEDIUM | P2 |
| Compiled schema cache | LOW | MEDIUM | P3 |
| Compliance caching | LOW | HIGH | P3 |

---

## Recommended Action Plan

### Phase 1: Critical Fixes (Low Effort, High Impact)

1. **Fix RFC 6901 unescaping order** - 1 line fix
2. **Add clear error for dynamic refs** - Document limitation
3. **Fix scope reset logic** - Track scope per subtree

### Phase 2: Performance Improvements

1. **Parallel fetching in processor** - Use existing errgroup
2. **Shared cache between processor and metaschema** - Refactor to single cache

### Phase 3: Spec Compliance

1. **Vocabulary scoping** - Requires scope tracking infrastructure
2. **Anchor path rewriting** - Build on scope tracking

### Phase 4: Nice to Have

1. **Compiled schema cache** - Thread-safe cache with LRU
2. **Incremental compliance caching** - Hash-based cache invalidation

---

## Questions to Resolve

- dynamic ref support: full implementation vs documented limitation?
- vocabulary scoping: strict spec compliance needed for target use cases?
- compliance caching: worth complexity for CI/CD speed?
- older drafts (draft-03, draft-04): full support needed?
