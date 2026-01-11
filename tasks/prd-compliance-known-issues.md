# PRD: Compliance Known Issues System

## Introduction

Improve compliance reports by excluding explicitly unsupported features from percentage calculations. Known issues are documented in a per-adapter JSON file, grouped by reason, and displayed separately in reports. This prevents "failing" tests for intentional limitations from skewing compliance metrics.

## Goals

- Exclude known/expected failures from compliance percentage calculations
- Document why specific tests are unsupported (language limitations, design decisions, etc.)
- Skip known-failing tests during execution (faster runs, cleaner output)
- Display known issues separately in web compliance pages
- Maintain transparency about what's supported vs intentionally unsupported

## User Stories

### US-001: Known issues file structure

**Description:** As a maintainer, I want a JSON file per adapter listing known issues so that I can document expected failures with explanations.

**Acceptance Criteria:**

- [ ] File location: `{adapter}/compliance/known-issues.json`
- [ ] Schema: array of groups, each with `reason` (string) and `tests` (array of test paths)
- [ ] Test path format: `{draft}/{keyword}/{group}/{test}` (e.g., `draft2020-12/additionalProperties/additionalProperties-are-allowed-by-default/no-additional-properties-is-valid`)
- [ ] Typecheck passes

### US-002: Skip known tests during execution

**Description:** As a developer, I want known-failing tests skipped during compliance runs so that execution is faster and output is cleaner.

**Acceptance Criteria:**

- [ ] Runner loads known-issues.json before test execution
- [ ] Tests matching known issue paths are not executed
- [ ] Skipped known tests don't appear in pass/fail/skip counts
- [ ] Log message indicates how many tests were skipped as known issues
- [ ] Typecheck passes

### US-003: Separate known issues in JSON results

**Description:** As a developer, I want JSON results to track known issues separately so that percentages reflect actual compliance.

**Acceptance Criteria:**

- [ ] Add `knownIssues` field to draft results (count + list)
- [ ] `summary.total` excludes known issue tests
- [ ] `summary.percentage` calculated without known issues
- [ ] Each known issue entry includes: test path, reason
- [ ] Typecheck passes

### US-004: Update dev report output

**Description:** As a developer, I want the CLI dev report to show known issues separately so I can see true compliance at a glance.

**Acceptance Criteria:**

- [ ] Summary line shows: "X passed, Y failed, Z skipped, W known issues"
- [ ] Percentage excludes known issues from denominator
- [ ] Known issues section lists reasons with test counts
- [ ] Typecheck passes

### US-005: Update web compliance page generation

**Description:** As a user viewing docs, I want the compliance page to clearly separate known limitations from unexpected failures.

**Acceptance Criteria:**

- [ ] "Known Behaviors" section displays known issues grouped by reason
- [ ] Each reason shows its test list
- [ ] "Issues" section only shows unexpected failures
- [ ] Summary table percentages exclude known issues
- [ ] Coverage stats reflect true compliance (known issues excluded)
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### US-006: Update REPORT.md generation

**Description:** As a developer, I want the markdown report to reflect the new known issues separation.

**Acceptance Criteria:**

- [ ] Summary table shows known issues count per draft
- [ ] Percentages exclude known issues
- [ ] "Known Issues" section renamed/split: "Known Behaviors" (from file) vs "Unexpected Failures"
- [ ] Typecheck passes

## Functional Requirements

- FR-1: Known issues file at `{adapter}/compliance/known-issues.json` with schema:
  ```json
  [
    {
      "reason": "Zod doesn't support recursive schemas without lazy()",
      "tests": [
        "draft2020-12/$ref/nested-refs/nested-ref-valid",
        "draft2020-12/$ref/nested-refs/nested-ref-invalid"
      ]
    }
  ]
  ```
- FR-2: Runner skips tests listed in known-issues.json entirely (no execution)
- FR-3: JSON results include `knownIssues: { count: number, items: Array<{ path: string, reason: string }> }`
- FR-4: All percentage calculations: `passed / (total - knownIssueCount) * 100`
- FR-5: CLI output shows known issues as separate category
- FR-6: Web page displays "Known Behaviors" section before "Issues" section
- FR-7: Empty known-issues.json (or missing file) = no known issues (backwards compatible)

## Non-Goals

- No UI for editing known issues (manual JSON editing only)
- No automatic detection of "should be known issue"
- No per-adapter override of reason display format
- No validation that listed tests actually exist in test suite

## Technical Considerations

- Runner is in Go (`cli/compliance/runner.go`) - needs to load JSON file
- Web generation is TypeScript (`web/scripts/generate-compliance.ts`)
- Existing types in `cli/compliance/types.go` need extension
- Known issues file should be optional (graceful handling when missing)

## Success Metrics

- Compliance percentages accurately reflect intentional support levels
- Known limitations documented with clear explanations
- No regression in compliance test execution time
