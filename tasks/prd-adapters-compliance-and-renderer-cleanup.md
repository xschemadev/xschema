# PRD: adapters compliance closure + renderer cleanliness

## introduction/overview

we want every adapter’s compliance failures to converge to only the things that are **impossible with static code generation** (given the current bundler already blocks unsupported ref features). everything else should be fixed.

after compliance closure, we do a second pass focused on **renderer cleanliness**: generated code should use bare/idiomatic library apis as much as possible and avoid “weird” constructs that exist only to satisfy compliance.

**in scope adapters:** `zod`, `valibot`, `effect`, `arktype`.

## goals

- for each adapter, reduce compliance failures until the remaining failures are **only** items that are impossible with static code generation.
- record the remaining “known limitations” (the impossible ones) in:
  - root `progress.txt` (running log)
  - each adapter’s `README.md` (stable reference)
- after renderer cleanup, keep compliance results **not worse** than the post-closure baseline.
- keep generated code idiomatic:
  - prefer library-native api usage over custom emulation
  - avoid “strange” constructs whose only purpose is to pass tests
  - priority: idiomatic apis first, then performance

## user stories

### us-001: zod — compliance closure to static-codegen limits

**description:** as a maintainer, i want the zod adapter to pass compliance except for limitations that are impossible with static code generation, so zod output is correct and predictable.

**acceptance criteria:**

- [ ] run `bun run compliance` in `typescript/packages/adapters/zod/` and capture baseline per draft.
- [ ] use the full report in `typescript/packages/adapters/zod/compliance/results/` to inspect failing tests.
- [ ] for every failing test, decide:
  - [ ] **fixable bug** (core/bundler/adapter behavior that can be implemented without runtime evaluation), or
  - [ ] **known limitation** because it is impossible with static code generation.
- [ ] fix every failure categorized as “fixable bug”, rerunning compliance until no fixable failures remain.
- [ ] update root `progress.txt` with: baseline, fixes made, and the final remaining known limitations summary.
- [ ] update `typescript/packages/adapters/zod/README.md` with a “known limitations” section describing the remaining failures and why they’re impossible with static code generation.

### us-002: zod — renderer cleanliness pass

**description:** as a maintainer, i want the zod renderer to produce clean/idiomatic zod code, so generated schemas look like something a human would write.

**acceptance criteria:**

- [ ] audit the zod renderer output patterns for common schema shapes (object/array/union/intersection/refinements).
- [ ] replace obviously weird/non-idiomatic constructs with more idiomatic zod api usage, while preserving json schema semantics.
- [ ] do not add “forced emulation” features that are out of scope for static code generation.
- [ ] run `bun run compliance` in `typescript/packages/adapters/zod/` and confirm results are not worse than us-001 baseline.
- [ ] update root `progress.txt` with a brief summary of what changed and why it’s more idiomatic.

### us-003: valibot — compliance closure to static-codegen limits

**description:** as a maintainer, i want the valibot adapter to pass compliance except for limitations that are impossible with static code generation, so valibot output is correct and predictable.

**acceptance criteria:**

- [ ] run `bun run compliance` in `typescript/packages/adapters/valibot/` and capture baseline per draft.
- [ ] use the full report in `typescript/packages/adapters/valibot/compliance/results/` to inspect failing tests.
- [ ] for every failing test, decide:
  - [ ] **fixable bug** (core/bundler/adapter behavior that can be implemented without runtime evaluation), or
  - [ ] **known limitation** because it is impossible with static code generation.
- [ ] fix every failure categorized as “fixable bug”, rerunning compliance until no fixable failures remain.
- [ ] update root `progress.txt` with: baseline, fixes made, and the final remaining known limitations summary.
- [ ] update `typescript/packages/adapters/valibot/README.md` with a “known limitations” section describing the remaining failures and why they’re impossible with static code generation.

### us-004: valibot — renderer cleanliness pass

**description:** as a maintainer, i want the valibot renderer to produce clean/idiomatic valibot code, so generated schemas are readable and avoid compliance-only hacks.

**acceptance criteria:**

- [ ] audit the valibot renderer output patterns for common schema shapes (object/array/union/intersection/pipelines/checks).
- [ ] replace obviously weird/non-idiomatic constructs with more idiomatic valibot api usage, while preserving json schema semantics.
- [ ] prioritize idiomatic apis first, then performance (avoid big perf regressions).
- [ ] run `bun run compliance` in `typescript/packages/adapters/valibot/` and confirm results are not worse than us-003 baseline.
- [ ] update root `progress.txt` with a brief summary of what changed and why it’s more idiomatic.

### us-005: effect — compliance closure to static-codegen limits

**description:** as a maintainer, i want the effect adapter to pass compliance except for limitations that are impossible with static code generation, so effect output is correct and predictable.

**acceptance criteria:**

- [ ] run `bun run compliance` in `typescript/packages/adapters/effect/` and capture baseline per draft.
- [ ] use the full report in `typescript/packages/adapters/effect/compliance/results/` to inspect failing tests.
- [ ] for every failing test, decide:
  - [ ] **fixable bug** (core/bundler/adapter behavior that can be implemented without runtime evaluation), or
  - [ ] **known limitation** because it is impossible with static code generation.
- [ ] fix every failure categorized as “fixable bug”, rerunning compliance until no fixable failures remain.
- [ ] update root `progress.txt` with: baseline, fixes made, and the final remaining known limitations summary.
- [ ] update `typescript/packages/adapters/effect/README.md` with a “known limitations” section describing the remaining failures and why they’re impossible with static code generation.

### us-006: effect — renderer cleanliness pass

**description:** as a maintainer, i want the effect renderer to use effect’s native schema apis as directly as possible, so generated code is idiomatic and maintainable.

**acceptance criteria:**

- [ ] audit the effect renderer output patterns for common schema shapes.
- [ ] replace obviously weird/non-idiomatic constructs with more idiomatic effect schema apis, while preserving json schema semantics.
- [ ] avoid adding custom runtime behavior just to pass a test (keep it static-codegen friendly).
- [ ] run `bun run compliance` in `typescript/packages/adapters/effect/` and confirm results are not worse than us-005 baseline.
- [ ] update root `progress.txt` with a brief summary of what changed and why it’s more idiomatic.

### us-007: arktype — compliance closure to static-codegen limits

**description:** as a maintainer, i want the arktype adapter to pass compliance except for limitations that are impossible with static code generation, so arktype output is correct and predictable.

**acceptance criteria:**

- [ ] run `bun run compliance` in `typescript/packages/adapters/arktype/` and capture baseline per draft.
- [ ] use the full report in `typescript/packages/adapters/arktype/compliance/results/` to inspect failing tests.
- [ ] for every failing test, decide:
  - [ ] **fixable bug** (core/bundler/adapter behavior that can be implemented without runtime evaluation), or
  - [ ] **known limitation** because it is impossible with static code generation.
- [ ] fix every failure categorized as “fixable bug”, rerunning compliance until no fixable failures remain.
- [ ] update root `progress.txt` with: baseline, fixes made, and the final remaining known limitations summary.
- [ ] update `typescript/packages/adapters/arktype/README.md` with a “known limitations” section describing the remaining failures and why they’re impossible with static code generation.

### us-008: arktype — renderer cleanliness pass

**description:** as a maintainer, i want the arktype renderer to use arktype’s native syntax/constructs as directly as possible, so generated code stays readable and not overly clever.

**acceptance criteria:**

- [ ] audit the arktype renderer output patterns for common schema shapes.
- [ ] replace obviously weird/non-idiomatic constructs with more idiomatic arktype patterns, while preserving json schema semantics.
- [ ] stop when only “obviously weird hacks” are removed; keep pragmatic patterns if they’re common arktype usage.
- [ ] run `bun run compliance` in `typescript/packages/adapters/arktype/` and confirm results are not worse than us-007 baseline.
- [ ] update root `progress.txt` with a brief summary of what changed and why it’s more idiomatic.

## functional requirements

- fr-1: for each adapter (`zod`, `valibot`, `effect`, `arktype`), the workflow must be: run compliance → classify failures → fix non-limitation failures → repeat until only static-codegen-impossible failures remain.
- fr-2: “known limitation” must mean: impossible to achieve with static code generation while preserving json schema semantics (no hidden runtime evaluation engine).
- fr-3: every time a limitation is accepted as “known”, it must be recorded in root `progress.txt` and in the adapter’s `README.md`.
- fr-4: renderer cleanup changes must not introduce new fixable failures; compliance after cleanup must be >= the post-closure baseline for that adapter.
- fr-5: renderer cleanup must prefer idiomatic library apis; avoid custom/strange constructs that exist only to satisfy compliance.

## non-goals (out of scope)

- no new support for features already blocked by the bundler (e.g. dynamic/recursive refs).
- no “runtime validator engine” inside adapters that executes full json schema evaluation beyond what the target library naturally provides.
- no rewriting the compliance harness, unless required to fix incorrect test wiring.
- no adding new adapters in this prd.

## design considerations (optional)

- each adapter `README.md` should have a clearly labeled “known limitations” section with short examples when helpful.
- prefer small, local renderer refactors over large rewrites.

## technical considerations (optional)

- compliances run per adapter via `bun run compliance` in the adapter package directory.
- after a compliance run, the full report is in `compliance/results/` inside that adapter package (e.g. `typescript/packages/adapters/zod/compliance/results/`).
- if core types/apis change, run `bun run build` from `typescript/` so workspace packages see updated `dist`.
- changes to go bundler or cli code should be validated with `go build` + `go test ./...` from `cli/`.

## success metrics

- for each adapter, the remaining compliance failures are fully explainable as static-codegen-impossible limitations (and documented in both `progress.txt` and adapter `README.md`).
- no adapter has “mystery failures” (i.e. failures that are neither fixed nor justified as impossible).
- renderer outputs (spot-checked on common schema shapes) use mostly library-native apis and avoid compliance-only hacks.
- compliance results do not regress after renderer cleanup passes.
