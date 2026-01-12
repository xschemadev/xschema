# PRD: Global Known Issues Refactor

## Introduction

Refactor known-issues from per-adapter to global. Known-issues should only contain actual xschema limitations (things that cannot be done in static generation), not adapter-specific quirks. Adapters will report pass/fail honestly; known-issues become global context explaining why certain tests are fundamentally unsupported.

## Goals

- Single source of truth for xschema limitations at `cli/compliance/known-issues.json`
- Compliance runner reads from cli folder, not per-adapter
- Web app gets `/compliance` landing page explaining what compliance means
- Web app gets `/compliance/known-issues` page showing all limitations with test lists
- Remove per-adapter known-issues.json files (already done in current branch)

## User Stories

### US-001: Update known-issues.json format

**Description:** As a developer, I need the known-issues file to have a clear structure with grouped limitations.

**Acceptance Criteria:**

- [ ] Format: `[{ "name": string, "reason": string, "tests": string[] }]`
- [ ] `name` is short identifier (e.g., "dynamic-refs", "recursive-refs")
- [ ] `reason` explains why this is fundamentally unsupported in static generation
- [ ] `tests` array contains full test paths
- [ ] File lives at `cli/compliance/known-issues.json`
- [ ] Typecheck passes

### US-002: Update compliance runner to use global known-issues

**Description:** As a developer, I want compliance to read known-issues from cli folder instead of per-adapter.

**Acceptance Criteria:**

- [ ] `LoadKnownIssues()` in `cli/compliance/runner.go` reads from `cli/compliance/known-issues.json`
- [ ] Remove adapter path dependency for known-issues loading
- [ ] Update `KnownIssueGroup` struct to include `Name` field
- [ ] All existing compliance commands work (`bun run compliance` in adapter folders)
- [ ] Typecheck/build passes

### US-003: Re-run compliance on all adapters

**Description:** As a developer, I need fresh compliance results after the refactor.

**Acceptance Criteria:**

- [ ] Run `bun run compliance` in `typescript/packages/adapters/zod/`
- [ ] Run `bun run compliance` in `typescript/packages/adapters/valibot/`
- [ ] Run `bun run compliance` in `typescript/packages/adapters/arktype/`
- [ ] Run `bun run compliance` in `typescript/packages/adapters/effect/`
- [ ] All results written to respective `compliance/results/` folders
- [ ] No runtime errors

### US-004: Create compliance landing page

**Description:** As a user, I want a `/compliance` page explaining what JSON Schema compliance means and linking to adapter-specific results.

**Acceptance Criteria:**

- [ ] Page at `/docs/compliance` (or similar route structure)
- [ ] Explains what JSON Schema compliance testing is
- [ ] Explains how xschema handles compliance (static generation approach)
- [ ] Links to `/docs/compliance/known-issues` for limitations
- [ ] Links to each adapter's compliance page
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### US-005: Create known-issues web page

**Description:** As a user, I want a `/compliance/known-issues` page showing all xschema limitations with their affected tests.

**Acceptance Criteria:**

- [ ] Page at `/docs/compliance/known-issues`
- [ ] Intro section explaining these are fundamental static-generation limitations
- [ ] Each limitation group displayed as accordion (similar to adapter compliance page component)
- [ ] Accordion header shows `name` and count of affected tests
- [ ] Accordion body shows `reason` explanation
- [ ] Accordion body lists all test names from `tests` array
- [ ] Data sourced from `cli/compliance/known-issues.json` (at build time)
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### US-006: Update adapter compliance page generation

**Description:** As a developer, I need the compliance MDX generator to reflect the new global known-issues approach.

**Acceptance Criteria:**

- [ ] `web/scripts/generate-compliance.ts` no longer reads per-adapter known-issues
- [ ] Known issues section removed from individual adapter compliance pages (they just show pass/fail/skip)
- [ ] Adapter pages link to global known-issues page for context
- [ ] Regenerate all adapter compliance MDX files
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

## Functional Requirements

- FR-1: Known-issues file format must be `[{ "name": string, "reason": string, "tests": string[] }]`
- FR-2: Compliance runner must load known-issues from `cli/compliance/known-issues.json` (relative to cli module, not adapter)
- FR-3: Known-issues page must generate test list from the JSON file at build time
- FR-4: Compliance landing page must be static MDX or generated content
- FR-5: Accordions on known-issues page must match visual style of adapter compliance pages

## Non-Goals

- No per-adapter known-issues files
- No runtime fetching of known-issues data in web app
- No changes to how compliance tests themselves are executed (just where config is read from)
- No changes to test discovery or JSON Schema test suite sources

## Technical Considerations

- Compliance runner is Go code in `cli/compliance/`
- Web app uses TanStack Router + Fumadocs with MDX
- Current accordion pattern in `generate-compliance.ts` can be reused
- Known-issues JSON will need to be imported/embedded in web build

## Success Metrics

- Single known-issues.json file at cli level
- All adapters run compliance without errors
- Known-issues page displays all limitations clearly
- Compliance landing page provides good context for new users

## Unresolved Questions

- should known-issues page be MDX (manual) or generated script like adapter compliance pages?
- exact route structure: `/docs/compliance/` vs `/compliance/` - follow existing fumadocs patterns?
