# PRD: Go CLI Language System Refactor

## Introduction/Overview

The Go CLI currently supports “languages” via a single `language.Language` struct and a hard-coded `Languages` slice. This setup has several issues:

- Adding a new language is not obviously scoped (logic is scattered: runner detection, templates, compliance, injection regexes, etc.).
- The system assumes **single-file output** (`OutputFile`), which won’t work for many languages.
- Adapter invocation is too coupled to Node-style runners (`bunx`/`npx`) and hardcodes how adapter strings map to runnable CLIs.
- Python support is incomplete but present, which creates false expectations and breaks docs/CLI UX.
- Documentation is currently inconsistent (broken links, outdated `cli/language/README.md` describing a removed tree-sitter approach).

This feature refactors language support to make setup easy, future-proof (multi-file outputs), and explicit. Python is removed for now; only TypeScript remains supported.

## Goals

- Make adding a new language straightforward: create a folder under `cli/language/langs/<language>/` and register it.
- Support **multi-file output generation** from the core pipeline.
- Treat `schemas[].adapter` as an **adapter reference string** (language-defined), and resolve it via a language-owned invoker using the convention `xschema-<package-basename>` (e.g. `@xschemadev/zod` → `xschema-zod`).
- Keep web config schema generation (`web/scripts/generate-schemas.ts`) in sync with the adapter reference format used by the CLI.
- Keep `--lang` flags; unknown languages should error.
- Remove incomplete Python language support (code + docs references) so the repo matches reality.
- Bring docs back into sync (fix missing/broken docs pages and outdated language docs).

## User Stories

### US-001: Define language spec + registry with validation
**Description:** As a maintainer, I want a validated language registry so languages are easy to add and impossible to register incorrectly.

**Acceptance Criteria:**
- [ ] Create core language types that separate **data** from **capabilities** (multi-file emitter, adapter invoker, optional compliance/injection).
- [ ] Provide registry API: `Register`, `ByName`, `BySchemaURL`, `AllIgnoreDirs`, and `SupportedLanguages` (names list).
- [ ] Registry validates at registration time:
  - unique canonical `Name`
  - unique schema URL mapping
  - required capabilities present for `generate` pipeline (invoker + emitter)
- [ ] Unknown language name returns `nil` (or error at callsites) with clear error messages in CLI commands.
- [ ] `go test ./...` passes
- [ ] `go vet ./...` passes

### US-002: Establish per-language folder structure
**Description:** As a contributor, I want a clear file layout so I can add a language without touching unrelated code.

**Acceptance Criteria:**
- [ ] Implement this directory structure:
  - `cli/language/spec/` (core types + capability interfaces)
  - `cli/language/registry/` (registration + lookup + validation)
  - `cli/language/langs/typescript/` (typescript implementation)
  - `cli/language/` remains the stable import path used by other packages (re-export/wrappers).
- [ ] The TypeScript language is registered automatically (e.g. via init/blank import from `cli/language`).
- [ ] `go test ./...` passes

### US-003: Multi-file output emission + injector support
**Description:** As a maintainer, I want languages to emit multiple files so future languages (rust/go/etc.) aren’t blocked.

**Acceptance Criteria:**
- [ ] Replace the single-file assumption with `[]GeneratedFile{ Path, Contents }` returned by a language emitter.
- [ ] Injector writes all emitted files under `--output` root.
- [ ] Path safety rules are enforced:
  - paths must be relative
  - no absolute paths
  - no `..` traversal after cleaning
  - no duplicate/overlapping paths (collision detection)
- [ ] Output ordering is deterministic (stable sorting by path).
- [ ] `go test ./...` passes

### US-004: Manifest-based stale file cleanup
**Description:** As a user, I want removed/renamed generated files to be cleaned up automatically without risking deletion of unrelated files.

**Acceptance Criteria:**
- [ ] Injector writes a manifest file at `--output/xschema.manifest.json` (fixed name).
- [ ] On subsequent runs, injector deletes **only** files that:
  - were previously generated (listed in old manifest)
  - are not generated in the current run
  - are inside `--output` root
- [ ] Deleting stale files does not delete files not listed in the old manifest.
- [ ] Manifest writing is robust (write temp + atomic rename).
- [ ] `go test ./...` passes

### US-005: Adapter reference semantics + CLI convention
**Description:** As a user, I want adapter references to be explicit and portable, so adapter invocation works across languages without hardcoding runtime assumptions in the core pipeline.

**Acceptance Criteria:**
- [ ] Config `schemas[].adapter` is treated as an **adapter reference string** and passed through to the language adapter invoker (core does not parse it beyond storing it).
- [ ] The TypeScript language validates adapter refs are **scoped npm package references** (e.g. `@xschemadev/zod`) and errors with migration guidance for legacy values (e.g. `"zod"`).
- [ ] The TypeScript adapter invoker derives the adapter CLI binary name from the reference:
  - basename = last path segment (e.g. `@xschemadev/zod` → `zod`)
  - if basename already starts with `xschema-`, use it; otherwise prefix `xschema-`
  - example: `@xschemadev/xschema-zod` → binary `xschema-zod` (no double prefix)
- [ ] Runner detection is performed relative to the **project root** (not current working dir).
- [ ] `go test ./...` passes

### US-006: Keep compliance command; error on unsupported languages
**Description:** As a maintainer, I want `xschema compliance` to remain, but to error clearly if a language is unknown/unsupported.

**Acceptance Criteria:**
- [ ] `xschema compliance --lang <unknown>` errors: `unknown language: <name>`.
- [ ] `xschema compliance --lang typescript` continues to work.
- [ ] Removing Python means `xschema compliance --lang python` errors.
- [ ] `go test ./...` passes

### US-007: Remove Python language support cleanly
**Description:** As a maintainer, I want Python removed so the codebase doesn’t claim support it can’t provide.

**Acceptance Criteria:**
- [ ] Remove Python language registration and implementation.
- [ ] Remove/adjust Python-specific templates/import logic and any docs that claim Python works.
- [ ] Parser multi-language tests remain covered by registering a test-only fake language (not by shipping Python).
- [ ] `go test ./...` passes

### US-008: Testing strategy per language + registry contract tests
**Description:** As a maintainer, I want tests that make it obvious what each language must implement.

**Acceptance Criteria:**
- [ ] Add registry/spec contract tests:
  - registration validation (required fields/capabilities)
  - unique name/schema mapping enforcement
  - safe path enforcement for emitted files
  - manifest cleanup behavior
- [ ] Add TypeScript language tests (in `cli/language/langs/typescript/`):
  - runner detection against fixture directories
  - adapter command building from package name
  - emitted file(s) content shape is deterministic
- [ ] Existing CLI unit tests are updated and remain green.
- [ ] `go test ./...` passes

### US-009: Documentation cleanup and restructuring
**Description:** As a contributor, I want documentation to match the actual system so I can add a language without guesswork.

**Acceptance Criteria:**
- [ ] Fix `docs/INDEX.md` broken references (e.g. missing `language-support.md`, wrong release doc name).
- [ ] Replace or rewrite `cli/language/README.md` to describe the new system (no tree-sitter references).
- [ ] Update `docs/cli-pipeline.md` language section to reflect:
  - only TypeScript supported
  - adapter field uses adapter package reference format (e.g. `@xschemadev/zod`) and the `xschema-<basename>` CLI convention
  - multi-file output capability
- [ ] `go test ./...` passes

### US-010: Web schema generation uses adapter refs
**Description:** As a user, I want the published JSON Schemas (used by IDEs) to validate and autocomplete adapter values in the same format the CLI expects, so configs don’t “typecheck” in the IDE but fail in the CLI (or vice versa).

**Acceptance Criteria:**
- [ ] `web/scripts/generate-schemas.ts` derives adapter values from adapter metadata files, not directory names.
- [ ] Each adapter directory contains `xschema.adapter.json` (minimum shape: `{ "ref": "..." }`).
- [ ] For adapters in `*/packages/adapters/*`, adapter reference is read from `xschema.adapter.json` `ref` (not from `package.json`).
- [ ] For TypeScript config schemas, generated `schemas[].adapter` enum/autocomplete uses these adapter refs (e.g. `@xschemadev/zod`).
- [ ] `web/scripts/generate-schemas.ts` errors with a clear message if any adapter:
  - is missing `xschema.adapter.json`
  - has invalid JSON
  - lacks `ref`
  - duplicates another adapter `ref`
- [ ] Running `bun run generate:schemas` from `web/` produces `web/public/schemas/typescript.jsonc` where `schemas[].adapter` suggests values like `@xschemadev/zod` (not `zod`).
- [ ] `bun run types:check` from `web/` passes (if script/types changed).

## Functional Requirements

- FR-1: The system must expose a single registry to register and lookup languages by canonical name.
- FR-2: The system must resolve language from `$schema` URLs like `https://xschema.dev/schemas/typescript.jsonc`.
- FR-3: `xschema generate` must error when `$schema` references an unknown language.
- FR-4: `schemas[].adapter` must be an adapter reference string; each language must map it to a runnable adapter CLI (default convention: `xschema-<package-basename>`); unsupported formats must error with migration guidance.
- FR-5: Adapter invocation must be language-defined (via invoker capability) and provide full command spec (cmd, args, working dir, env).
- FR-6: Output emission must support multiple files and must be validated for safe, relative paths.
- FR-7: Injector must write all emitted files and maintain `--output/xschema.manifest.json` for stale file cleanup.
- FR-8: Injector must delete stale generated files listed in the old manifest but not in the new manifest.
- FR-9: `xschema compliance --lang <name>` must error for unknown languages and work for `typescript`.
- FR-10: Documentation must reflect the actual supported languages and the adapter field format.
- FR-11: Web schema generation must use adapter reference strings for `schemas[].adapter` sourced from `xschema.adapter.json` files (not directory names).

## Non-Goals (Out of Scope)

- No Python language support in this iteration.
- No new languages added beyond TypeScript.
- No change to the adapter stdin/stdout protocol.
- No watch mode implementation.
- No tree-sitter-based parsing or client injection improvements beyond what already exists.
- No automatic migration tool for existing configs (manual update is acceptable).
- No object-form adapter config; `schemas[].adapter` remains a string in this iteration.

## Design Considerations (Optional)

- **Language capabilities as fields:** Prefer a `Language` struct with required fields plus capability slots (interfaces) for invoker/emitter. This keeps language definitions readable and avoids sprawling Go interface hierarchies.
- **Output freedom:** Languages can emit any relative paths inside `--output`. This supports future language conventions.
- **Templates:** Each language owns its own templates/rendering approach:
  - Simple languages can use Go `text/template` per output file.
  - Complex languages can build files programmatically.
  - Shared helpers should be limited to generic utilities (path validation, deterministic ordering, template rendering helper).

## Technical Considerations (Optional)

- **Global registry state:** Provide a `ResetForTests()` (or similar) to avoid test order coupling.
- **Atomic writes:** Write output files and the manifest using temp files + rename to reduce corruption on interruption.
- **Windows path behavior:** Validate paths using `filepath.Clean` and explicitly block volume/drive/absolute paths.
- **Stale deletion safety:** Only delete files listed in manifest and ensure they still resolve under `--output`.
- **Determinism:** Sort emitted files by path; sort manifest entries.

## Success Metrics

- A new language can be scaffolded by adding `cli/language/langs/<name>/` and registering it, without editing generator/injector/parser internals.
- Multi-file generation works end-to-end: removing an emitted file from output causes it to be removed on the next run via manifest cleanup.
- `xschema generate` and `xschema compliance` error clearly for unknown languages.
- Documentation contains no broken links for language docs and no claims of Python support.
- Generated config JSON schemas validate/suggest adapters using the same adapter reference format as the CLI.

## Open Questions

- None for this iteration (manifest name fixed; no `--clean`; adapter config remains string-only).