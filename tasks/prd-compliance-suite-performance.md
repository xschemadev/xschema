# PRD: Compliance Suite Performance

## 1. Introduction/Overview

The adapter compliance suite is currently slow in both local development and CI. The main cause is high per-test overhead: for each test group, the runner bundles the schema, spawns an adapter process, writes a temporary TypeScript harness file, then spawns a runtime process to execute that harness.

This PRD defines a “do it properly” performance overhaul focused on drastically reducing process spawns while keeping results deterministic and the `--dev-report` output stable across repeated runs.

## 2. Goals

- Reduce end-to-end compliance runtime for a typical TypeScript adapter (e.g. `@xschemadev/zod`) by **≥50% in CI** and **≥50% locally**.
- Reduce process spawns by batching so that, per draft, the runner executes:
  - adapter conversion in **O(keywords)** (or better), not **O(test groups)**
  - harness runtime execution in **O(keywords)** (or better), not **O(test groups)**
- Add a `--profile` mode that outputs a clear timing breakdown to identify regressions and validate improvements.
- Ensure deterministic results ordering and stable outputs:
  - repeated runs of `--dev-report` produce **byte-for-byte identical** outputs (no timestamps; stable ordering; stable path normalization).
- Default to safe, deterministic execution: `--jobs` defaults to **1**.

## 3. User Stories

### US-001: Add `--profile` timing breakdown

**Description:** As a developer, I want a timing breakdown for compliance runs so that I can see where time is spent and verify performance improvements.

**Acceptance Criteria:**

- [ ] Add a `--profile` boolean flag to the `xschema compliance` command.
- [ ] When `--profile` is enabled, print a timing summary that includes at least:
  - [ ] test suite load time
  - [ ] schema bundling time
  - [ ] adapter invocation time
  - [ ] harness generation time
  - [ ] harness execution time
- [ ] Save current output times to proggress.txt
- [ ] Timing output is clearly labeled and stable (plain text; stable labels + units).
- [ ] Timing output does not change the computed compliance results.
- [ ] Typecheck/build passes.

### US-002: Guarantee deterministic `--dev-report` outputs across runs

**Description:** As an adapter developer, I want `--dev-report` to produce identical output across runs so that diffs represent real changes, not noise.

**Acceptance Criteria:**

- [ ] Running `xschema compliance --dev-report` twice (same adapter build, same suite cache, same flags) produces byte-for-byte identical file contents in `compliance/results/`.
- [ ] Ensure stable ordering in all outputs:
  - [ ] drafts are written in a deterministic order
  - [ ] keywords are written in a deterministic order
  - [ ] failures are written in a deterministic order
- [ ] Remove non-deterministic timestamps from `--dev-report` outputs:
  - [ ] Remove `generatedAt` from the JSON report(s) written to `compliance/results/`.
  - [ ] Remove the `Generated:` line from `compliance/results/REPORT.md`.
- [ ] CI can validate results with a plain diff check (no timestamp regex special-casing).
- [ ] Typecheck/build passes.

### US-003: Reduce harness runtime overhead for bun

**Description:** As a developer, I want harness execution to avoid unnecessary overhead so that the suite runs faster without changing semantics.

**Acceptance Criteria:**

- [ ] For TypeScript harness execution under bun, prefer `bun <file>` over `bun run <file>`.
- [ ] Existing harness behavior remains the same (stdout JSON format, working directory resolution, error handling).
- [ ] Typecheck/build passes.
- [ ] --profile output performance is better than before (previous in progress.txt)

### US-004: Speed up keyword-only runs by loading only the requested test file

**Description:** As a developer iterating on a specific keyword, I want `--keyword` runs to load only the relevant test JSON file so iteration is faster.

**Acceptance Criteria:**

- [ ] When `--keyword <k>` is set, load only `tests/<draft>/<k>.json` for that draft.
- [ ] If `<k>.json` does not exist for that draft, return a clear error listing available keywords (or count).
- [ ] Results for `--keyword` match the subset of a full run for that draft.
- [ ] Typecheck/build passes.
- [ ] --profile output performance is better than before (previous in progress.txt)

### US-005: Batch adapter conversion and harness execution per keyword

**Description:** As a developer, I want the compliance runner to batch work per keyword to dramatically reduce process spawns and runtime.

**Acceptance Criteria:**

- [ ] For each (draft, keyword), the runner calls the adapter with a batch of schemas (one per group) in a single adapter invocation.
- [ ] For each (draft, keyword), the runner executes a single harness runtime process that evaluates all groups and testcases for that keyword.
- [ ] The resulting `ComplianceReport` totals (passed/failed/skipped/total) match the current implementation for the same adapter and suite inputs.
- [ ] Failure records include group + test descriptions and remain deterministic across runs.
- [ ] Typecheck/build passes.
- [ ] --profile output performance is better than before (previous in progress.txt)

### US-006: Add `--jobs` for optional bounded concurrency (default 1)

**Description:** As a developer, I want optional concurrency so that I can speed up runs on powerful machines/CI, while keeping the default deterministic and low-resource.

**Acceptance Criteria:**

- [ ] Add a `--jobs <n>` flag where `n >= 1`.
- [ ] Default is `--jobs 1`.
- [ ] When `--jobs > 1`, execution is parallelized at a safe unit of work (recommended: per keyword within a draft).
- [ ] Important: Output remains deterministic (final report ordering is stable regardless of execution order).
- [ ] Typecheck/build passes.
- [ ] Pretty progress output remains readable with concurrency (no interleaved garbage output).
- [ ] --profile output performance is better than before (previous in progress.txt)

### US-007: Pin JSON Schema Test Suite version for CI stability

**Description:** As a maintainer, I want the JSON Schema Test Suite pinned to a specific tag or commit so that CI results are stable and not affected by upstream changes on `main`.

**Acceptance Criteria:**

- [ ] Replace the current `main`-tracking download URL with a pinned reference (tag or commit SHA).
- [ ] Cache directory includes the pinned version so different versions don’t collide.
- [ ] CI uses the pinned suite version by default (no extra workflow complexity required).
- [ ] Provide a clear, low-effort process to bump the pinned version (single constant change + results regen).
- [ ] Typecheck/build passes.

## 4. Functional Requirements

- FR-1: The system must support a `--profile` flag that prints a timing breakdown for major phases of the compliance run.
- FR-2: The system must keep `--dev-report` outputs byte-for-byte stable across repeated runs with identical inputs.
- FR-3: The system must remove non-deterministic timestamps from `--dev-report` outputs (no `generatedAt`, no `Generated:` line in `REPORT.md`).
- FR-4: The system must execute TypeScript harnesses efficiently under bun (avoid unnecessary `bun run` overhead).
- FR-5: When `--keyword` is provided, the system must load only the corresponding keyword test file for each draft.
- FR-6: The system must batch adapter conversion per (draft, keyword), sending multiple schema inputs in one adapter invocation.
- FR-7: The system must batch harness execution per (draft, keyword), running one harness runtime process per keyword.
- FR-8: The system must provide an optional `--jobs` flag for bounded parallel execution; default must be `1`.
- FR-9: The system must preserve existing result semantics (pass/fail/skipped logic) and error normalization.
- FR-10: The system must keep report ordering deterministic (draft order, keyword order, failure ordering).
- FR-11: CI must be able to validate `compliance/results/` freshness using a simple `git diff --exit-code` check (no regex filtering).
- FR-12: The system must pin the JSON Schema Test Suite version (tag/commit) to keep CI stable; it must not silently track upstream `main`.

## 5. Non-Goals (Out of Scope)

- No changes to the meaning of the JSON Schema Test Suite or its expected outcomes.
- No changes to adapter correctness logic (this is performance + determinism work).
- No UI work.
- No changing the existing restriction that `--dev-report` cannot be combined with `--draft` or `--keyword`.
- No introducing new external dependencies unless clearly justified.

## 6. Design Considerations (Optional)

- CLI output should remain readable; `--profile` should add optional extra output without disrupting existing progress output.
- Prefer a single concise timing summary at the end of the run.
- If adding any structured `--profile` output format (e.g. JSON), keep it stable for CI parsing.

## 7. Technical Considerations (Optional)

- Primary bottleneck today is repeated process spawning:
  - adapter calls per group
  - harness executions per group
- Batching strategy:
  - Adapter protocol already supports arrays; use a unique `id` per group to map outputs back to inputs.
  - Generate one harness per keyword that contains an array of (schema, validate fn, tests) and executes them in-process.
- Determinism requirements:
  - Sort drafts and keywords before writing.
  - Ensure batching does not change ordering of keyword results.
  - Ensure `--dev-report` does not embed timestamps (remove `generatedAt` / `Generated:`) or machine-specific paths.
  - After timestamps are removed, simplify `.github/workflows/compliance.yml` to use a plain diff check on `compliance/results/` (no special-case filtering).
- Concurrency (`--jobs`):
  - Default 1 to preserve deterministic behavior and avoid resource spikes.
  - If enabled, parallelize at the keyword level to keep batching effective.
  - Collect results then sort before writing/reporting.
- Performance guardrails:
  - Assume no chunking in v1 (batch whole keyword); if a real keyword proves too large, revisit with a concrete repro and keep output deterministic.

## 8. Success Metrics

- Reduce wall-clock runtime by ≥50% for `xschema compliance` in CI for a typical adapter.
- Reduce wall-clock runtime by ≥50% locally.
- Reduce total spawned runtime processes by at least an order of magnitude (target: from “per group” to “per keyword”).
- `--dev-report` output is stable across repeated runs, enabling clean diffs.

## 9. Open Questions

- None for this scope
