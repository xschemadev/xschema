# PRD: Adapters Optimization (Idiomatic + High Compliance)

## Introduction/Overview

Improve the TypeScript JSON Schema pipeline (`@xschemadev/core` IR/parser + each TS adapter renderer) to:

- represent **all JSON Schema features** in core/IR (lossless / no hard throws)
- maximize **test-suite compliance** across all drafts in the test suite, without intentionally breaking older drafts (unless explicitly justified + documented)
- keep adapters **idiomatic** (no “mini JSON Schema engine per adapter”), prioritizing readable output and good runtime performance
- make remaining gaps **explicit** via manual, accurate documentation in each adapter README

Guiding rule: implement features to the maximum extent that each target library supports *in its own best practices*. If a feature is not supported (or would require forced emulation), prefer a safe **widening** behavior (accept more data) and document the gap.

## Goals

- Improve compliance across all drafts in the test suite, without intentionally regressing older drafts.
- Make core parsing **total**: it should not throw just because a keyword exists.
- Make “unsupported feature” behavior predictable and consistent (widen + document).
- Improve generated code readability and runtime performance, using adapter-library best practices.

## Operating Principles (Critical)

- **Draft priority:** treat all drafts as first-class targets. Only prioritize a newer draft (e.g. draft2020-12) when there is a real tradeoff; otherwise prefer solutions that improve or maintain all drafts.
- **Idiomatic over forced:** adapters may use library-native constructs and *idiomatic* refinements. Avoid building per-adapter “JSON Schema engines”.
- **Unsupported = widen:** for features not supported by the adapter library, generated validation should default to accepting more (skip the check) rather than rejecting more.
- **Failures must be documented:** any known limitation must be explained in the adapter README (manual writing), with examples and expected compliance implications.

## Definitions

- **Widening behavior:** generated validator accepts some inputs that JSON Schema would reject because the adapter cannot express that keyword/interaction.
- **Forced emulation:** large inline generated code that re-implements JSON Schema semantics beyond what is typical for the adapter library.
- **New failure:** a (draft, keyword, group, test) that was passing in the US-001 baseline and fails after a change.
- **Expected regression:** a new failure that was explicitly pre-declared during discovery (with rationale) and then documented in the adapter README Limitations section.

## User Stories

### US-001: Establish baseline + improvement rubric (repo-wide)

**Description:** As a maintainer, I want a clear baseline of compliance + a consistent rubric for “supported vs forced” so improvements stay aligned.

**Acceptance Criteria:**

- [ ] For each adapter, run `bun run compliance` and record a baseline (used for later diffing). Baseline record must include, per draft:
  - overall pass %
  - top 10 failing keywords (by failed tests)
  - list of failing tests with identifiers: `(keyword, group, test)`
  - any keywords with unexpected runtime errors
- [ ] Define a rubric (short, explicit) for classifying a failure into:
  - `core-missing` (IR/parser cannot represent it or behaves incorrectly)
  - `adapter-native` (the library can express it cleanly)
  - `forced-emulation` (possible but would require heavy custom logic; avoid)
  - `not-supported` (library can’t do it reasonably)
- [ ] Identify at least 3 “high leverage” themes (cross-adapter) and at least 3 per-adapter themes.
- [ ] `bun run build` and `bun run typecheck` succeed from `typescript/`.

### US-002: Make core parsing lossless (no keyword-based hard throws)

**Description:** As an adapter author, I want core parsing to represent any JSON Schema keyword without throwing so adapters can decide what to do.

**Acceptance Criteria:**

- [ ] `@xschemadev/core` parser never throws *only because a keyword exists*.
- [ ] Previously-blocking cases are represented in IR rather than throwing (examples: `unevaluatedItems`, `unevaluatedProperties` with applicators).
- [ ] IR preserves enough information to:
  - render adapter-native code where supported
  - detect/report unsupported keywords where not supported
  - produce accurate README limitation notes
- [ ] Lossless requirement is explicit: unknown/unsupported keywords must be retained (keyword name + raw value + draft/version context), not silently dropped.
- [ ] Add tests that cover: schemas containing these keywords do not throw, and the resulting IR retains those keyword payloads.
- [ ] `bun run build` and `bun run typecheck` succeed from `typescript/`.

### US-003: Add a capability-aware rendering contract (core ↔ adapters)

**Description:** As a maintainer, I want adapters to declare what they support so unsupported features can widen + be documented consistently.

**Acceptance Criteria:**

- [ ] Define an adapter-facing capability model (e.g. `AdapterCapabilities`) that can represent:
  - supported keywords
  - supported interactions (e.g. object strictness across `allOf`)
  - supported validation modes (e.g. “strict missing keys” vs “coerce/ignore”)
- [ ] Define a standard policy for unsupported features:
  - widening runtime behavior
  - structured “limitation markers” available to README authors (manual writing still required)
- [ ] Adapters can adopt the capability model incrementally.
- [ ] `bun run build` and `bun run typecheck` succeed from `typescript/`.

### US-004: Version-aware behavior (prioritize latest without breaking others)

**Description:** As a maintainer, I want draft differences honored so improvements help all drafts without accidental regressions.

**Acceptance Criteria:**

- [ ] IR carries draft/version context and draft-specific keyword semantics where they differ.
- [ ] If `$schema` is absent/unknown, default to latest draft behavior.
- [ ] If a change causes a regression in older drafts, it must be either:
  - accidental (then fix), or
  - intentional (then document in README(s) and provide rationale).

---

## Adapter Work Pattern (must be followed)

Every adapter gets two stories: **discovery** then **implementation**.

Discovery story output must include:

- the categorized failure list (rubric above)
- a proposed follow-up backlog (titles + 1–2 sentences each)
- explicit dependencies (e.g. “blocked by US-002 core change”)
- a list of **expected regressions** (forced-emulation removals) that the implementation is allowed to introduce

If discovery reveals new cross-adapter/core needs, the implementer should add new backlog items (or expand US-002/US-003) before starting implementation.

This is important: the plan is expected to evolve. Implementation is expected to add new backlog items when new information appears.

## Compliance Diff Format (required)

After each implementation story, produce a compliance diff summary for all drafts in scope (draft3/4/6/7/2019-09/2020-12):

- overall pass % before → after
- newly failing tests: list `(keyword, group, test)`
- resolved tests: list `(keyword, group, test)`
- top 10 failing keywords before → after

New failures are only allowed if they are expected regressions and documented in the adapter README.

---

### US-101: Zod adapter discovery (plan)

**Description:** As a maintainer, I want to map Zod failures to “native vs forced” so we only implement what’s idiomatic.

**Acceptance Criteria:**

- [ ] Run `bun run compliance` in `typescript/packages/adapters/zod/`.
- [ ] Produce a categorized list of top failures (all drafts) using the rubric.
- [ ] Propose follow-up backlog items with dependencies (e.g. core IR changes vs adapter-only).

### US-102: Zod adapter improvements (implement)

**Description:** As a user, I want Zod output that is idiomatic, readable, and fast, while improving compliance where Zod supports it.

**Acceptance Criteria:**

- [ ] Implement only `adapter-native` items (and optional small `forced-emulation` items only if clearly justified as still idiomatic).
- [ ] Generated output becomes simpler to read (prefer named helpers over repeated inline mega-closures when it improves readability and perf).
- [ ] Re-run `bun run compliance` in `typescript/packages/adapters/zod/` and produce a compliance diff summary (see Compliance Diff Format) vs the baseline from US-001.
- [ ] Any new failures must be either fixed, or match a pre-declared expected regression from discovery and be documented in the README Limitations section.
- [ ] If any older-draft regressions happen, they must be justified and documented.
- [ ] Update `typescript/packages/adapters/zod/README.md` with a manual **Limitations** section:
  - list unsupported keywords/interactions
  - describe widening behavior
  - include concrete examples
- [ ] `bun run build` and `bun run typecheck` succeed from `typescript/`.

---

### US-201: Valibot adapter discovery (plan)

**Description:** As a maintainer, I want to decide which Valibot failures are idiomatic to implement.

**Acceptance Criteria:**

- [ ] Run `bun run compliance` in `typescript/packages/adapters/valibot/`.
- [ ] Categorize top failures (all drafts) and propose backlog items with dependencies.

### US-202: Valibot adapter improvements (implement)

**Description:** As a user, I want Valibot output that stays idiomatic and readable while improving correctness.

**Acceptance Criteria:**

- [ ] Implement only idiomatic Valibot constructs (avoid building a generic evaluator).
- [ ] Prefer output readability first, then runtime performance, then minimal size.
- [ ] Re-run `bun run compliance` in `typescript/packages/adapters/valibot/` and produce a compliance diff summary (see Compliance Diff Format) vs the baseline from US-001.
- [ ] Any new failures must be either fixed, or match a pre-declared expected regression from discovery and be documented in the README Limitations section.
- [ ] Update `typescript/packages/adapters/valibot/README.md` Limitations section with examples.
- [ ] `bun run build` and `bun run typecheck` succeed from `typescript/`.

---

### US-301: TypeBox adapter discovery (plan)

**Description:** As a maintainer, I want to confirm which features TypeBox can represent and which its runtime validator enforces.

**Acceptance Criteria:**

- [ ] Run `bun run compliance` in `typescript/packages/adapters/typebox/`.
- [ ] Categorize top failures (all drafts), explicitly separating “expressible schema” vs “runtime enforced”.
- [ ] Propose backlog items with dependencies.

### US-302: TypeBox adapter improvements (implement)

**Description:** As a user, I want TypeBox output that is idiomatic and fast, improving compliance where TypeBox supports it.

**Acceptance Criteria:**

- [ ] Prefer TypeBox-native constructs and simple patterns.
- [ ] Avoid deep inline validation closures.
- [ ] Re-run `bun run compliance` in `typescript/packages/adapters/typebox/` and produce a compliance diff summary (see Compliance Diff Format) vs the baseline from US-001.
- [ ] Any new failures must be either fixed, or match a pre-declared expected regression from discovery and be documented in the README Limitations section.
- [ ] Update `typescript/packages/adapters/typebox/README.md` Limitations section with examples.
- [ ] `bun run build` and `bun run typecheck` succeed from `typescript/`.

---

### US-401: Effect adapter discovery (plan)

**Description:** As a maintainer, I want to ensure Effect Schema options align with JSON Schema expectations where Effect supports them.

**Acceptance Criteria:**

- [ ] Run `bun run compliance` in `typescript/packages/adapters/effect/`.
- [ ] Categorize failures, paying special attention to:
  - missing required property handling
  - excess property handling
  - exactness options
- [ ] Propose backlog items with dependencies.

### US-402: Effect adapter improvements (implement)

**Description:** As a user, I want Effect output that is idiomatic and performant, using Effect primitives/options instead of emulation.

**Acceptance Criteria:**

- [ ] Use Effect Schema’s built-in options/constructs where they map cleanly to JSON Schema.
- [ ] Default to widening behavior when Effect cannot express semantics.
- [ ] Re-run `bun run compliance` in `typescript/packages/adapters/effect/` and produce a compliance diff summary (see Compliance Diff Format) vs the baseline from US-001.
- [ ] Any new failures must be either fixed, or match a pre-declared expected regression from discovery and be documented in the README Limitations section.
- [ ] Update `typescript/packages/adapters/effect/README.md` Limitations section with examples.
- [ ] `bun run build` and `bun run typecheck` succeed from `typescript/`.

---

### US-501: ArkType adapter discovery (plan)

**Description:** As a maintainer, I want to keep ArkType output clean and avoid forced emulation.

**Acceptance Criteria:**

- [ ] Run `bun run compliance` in `typescript/packages/adapters/arktype/`.
- [ ] Categorize failures and propose backlog items with dependencies.

### US-502: ArkType adapter improvements (implement)

**Description:** As a user, I want ArkType output that remains clean and reasonably fast while improving coverage where supported.

**Acceptance Criteria:**

- [ ] Implement only idiomatic ArkType patterns.
- [ ] Re-run `bun run compliance` in `typescript/packages/adapters/arktype/` and produce a compliance diff summary (see Compliance Diff Format) vs the baseline from US-001.
- [ ] Any new failures must be either fixed, or match a pre-declared expected regression from discovery and be documented in the README Limitations section.
- [ ] Update `typescript/packages/adapters/arktype/README.md` Limitations section with examples.
- [ ] `bun run build` and `bun run typecheck` succeed from `typescript/`.

---

### US-601: TypeScript adapter discovery + simplification (plan)

**Description:** As a maintainer, I want the TypeScript adapter output to stay readable without multi-file helpers unless unavoidable.

**Acceptance Criteria:**

- [ ] Run `bun run compliance` in `typescript/packages/adapters/typescript/`.
- [ ] Identify forced type-level patterns that reduce readability without meaningful benefit.
- [ ] Propose backlog items with dependencies.

### US-602: TypeScript adapter simplification (implement)

**Description:** As a user, I want generated TS types that are stable and easy to read.

**Acceptance Criteria:**

- [ ] Prefer readability first, then runtime performance (where applicable), then minimal size.
- [ ] Avoid multi-file helpers unless truly necessary.
- [ ] Re-run `bun run compliance` in `typescript/packages/adapters/typescript/` and produce a compliance diff summary (see Compliance Diff Format) vs the baseline from US-001 (skips/limits must be documented).
- [ ] Any new failures must be either fixed, or match a pre-declared expected regression from discovery and be documented in the README Limitations section.
- [ ] Update `typescript/packages/adapters/typescript/README.md` Limitations section with examples.
- [ ] `bun run build` and `bun run typecheck` succeed from `typescript/`.

---

### US-701: Document limitations per adapter (manual)

**Description:** As a maintainer, I want each adapter README to accurately describe limitations and widening behavior.

**Acceptance Criteria:**

- [ ] Each adapter README contains a **Limitations** section that is manually written and includes, per major limitation:
  - the keyword/interaction name (draft-specific if needed)
  - a clear statement of widening behavior ("this adapter may accept inputs JSON Schema would reject")
  - at least 1 concrete example schema + example data
  - (when applicable) a reference to at least one failing test-suite `(keyword, group)` name so it stays honest
  - guidance to run `bun run compliance` for current status
- [ ] If the adapter intentionally chooses readability/performance over compliance for a forced-emulation case, the README explains why.

## Functional Requirements

- FR-1: Core parser must represent any JSON Schema keyword without throwing solely due to keyword presence.
- FR-2: Core IR must preserve enough information to render or intentionally skip unsupported semantics.
- FR-3: Core must expose draft/version context.
- FR-4: Adapters must declare capabilities and follow the widening fallback policy for unsupported features.
- FR-5: Adapters must prioritize readability and performance in idiomatic ways.
- FR-6: Every adapter change must be validated by `bun run compliance` (from the adapter folder) and compared against the baseline to ensure regressions are expected and documented.

## Non-Goals (Out of Scope)

- No per-adapter full JSON Schema evaluation engine.
- No promise of 100% JSON Schema Test Suite compliance for every adapter.
- No new heavy runtime deps by default.

## Success Metrics

- Pass rates improve in aggregate across adapters and drafts, or stay stable while readability/performance improves for prioritized areas.
- No unexplained regressions in older drafts.
- Every adapter README Limitations section matches observed behavior (manual + accurate).
- Generated output becomes noticeably cleaner (less repeated inline logic, more consistent structure) while staying within adapter best practices.
