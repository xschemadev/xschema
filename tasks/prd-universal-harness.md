# PRD: Universal Compliance Harness

## Introduction

Replace per-adapter harness files with one harness template per language. Currently each adapter (zod, arktype, valibot, etc.) has its own `compliance/harness.ts` with hardcoded imports and validation logic.

Solution: adapter provides a `validate` function, Go uses `text/template` to generate harness from a per-language template.

## Goals

- One harness template per language (not per adapter)
- Adapter output includes `validate` function in target language syntax
- Use Go `text/template` (same as existing code generation)
- Design supports future languages (Python, Rust, Go)
- Remove per-adapter harness files

## User Stories

### US-001: Add Validate field to AdapterOutput

**Description:** As the compliance runner, I need adapters to provide a validate function.

**Acceptance Criteria:**

- [ ] Add `Validate string` field to `AdapterOutput` struct in `cli/compliance/types.go`
- [ ] Add `ValidateImports []string` field for imports needed by validate function (optional)
- [ ] `Validate` contains function in target language syntax
- [ ] Function signature: takes `data`, returns bool, can reference `schema` variable
- [ ] Empty `Validate` = type-only adapter (skip runtime validation)

### US-002: Update TypeScript adapters to return validate

**Description:** As the harness generator, I need each TS adapter to provide its validate function.

**Acceptance Criteria:**

- [ ] Each adapter returns `validate` function: `(data) => boolean`
- [ ] Each adapter returns `validateImports` if needed for the validate function
- [ ] typescript (type-only): empty validate (triggers "skipped" behavior)
- [ ] All adapters pass existing compliance tests

### US-003: Create TypeScript harness template

**Description:** As the compliance system, I need a single TS harness template.

**Acceptance Criteria:**

- [ ] Uses Go text/template (same pattern as existing language templates)
- [ ] Outputs JSON array to stdout (same format as current harness)
- [ ] Handles type-only adapters (empty Validate → all results "skipped")
- [ ] Write Go tests for template rendering

### US-004: Update GenerateTempHarness to use templates

**Description:** As the compliance runner, I need GenerateTempHarness to use Go text/template.

**Acceptance Criteria:**

- [ ] Use `template.Execute` instead of `strings.ReplaceAll`
- [ ] Select template based on language
- [ ] Existing compliance tests pass
- [ ] Write thoughtful Go tests for new template generation logic

### US-005: Remove per-adapter harness files

**Description:** As a maintainer, I want to remove duplicate harness files.

**Acceptance Criteria:**

- [ ] Delete all `packages/adapters/*/compliance/harness.ts` files
- [ ] Update `FindHarness` logic (no longer needs per-adapter file)
- [ ] Verify no other code references deleted files (web app, shared functions, imports)
- [ ] All compliance tests pass
- [ ] Full test suite passes (no regressions)

### US-006: Add Python harness template (future-ready)

**Description:** Prepare template for future Python adapters.

**Acceptance Criteria:**

- [ ] Python harness template created (same pattern as TypeScript)
- [ ] Expects `validate` as full `def` block (Python lambdas can't have try/except)
- [ ] Same output format as TypeScript harness

## Functional Requirements

- FR-1: `AdapterOutput` must include `Validate` and `ValidateImports` fields (both optional)
- FR-2: Use Go `text/template` (same as existing language templates)
- FR-3: Empty `Validate` → all results "skipped"
- FR-4: Output format unchanged: `[{index, expected, actual, error?}]`

## Non-Goals

- No changes to adapter input format
- No changes to compliance CLI interface
- No Python/Rust/Go adapters implemented (just templates ready)

## Technical Considerations

- Follow existing patterns in `cli/language/` for templates and language config

## Success Metrics

- Adding new TS adapter requires only: adapter code + validate function (no harness.ts)
- All existing adapters pass compliance with universal harness

## Open Questions

None.
