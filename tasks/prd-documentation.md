# PRD: xschema Documentation

## Introduction

Build comprehensive documentation for xschema - a tool that generates native validators (Zod, Pydantic, etc.) from JSON Schema with full type safety. The docs need to serve two user journeys (framework vs library mode), scale to multiple languages, and build trust through compliance transparency.

## Goals

- Get users to working code in under 5 minutes
- Clearly distinguish between framework mode (full integration) and library mode (direct adapter usage)
- Structure docs to scale to many languages (TypeScript, Python, Go, etc.)
- Build trust through transparent compliance reporting
- Provide language-specific getting started guides that don't require cross-referencing

## User Stories

### US-001: Restructure content directory

**Description:** As a developer, I need the docs file structure reorganized so content is logically grouped by scope (general vs language-specific).

**Acceptance Criteria:**
- [ ] Move existing adapter pages from `/typescript/adapters/` to `/typescript/adapters/`
- [ ] Create `/typescript/framework/` directory
- [ ] Create `/typescript/library/` directory
- [ ] Create `/internals/` directory
- [ ] Move compliance pages to correct locations
- [ ] Remove `test.mdx` placeholder file
- [ ] Verify build passes after restructure

**Research:** Check current structure at `web/content/docs/`

---

### US-002: Create docs introduction page

**Description:** As a new visitor, I want to understand what xschema does and why it exists so I can decide if it solves my problem.

**Acceptance Criteria:**
- [ ] Replace boilerplate index.mdx with xschema introduction
- [ ] Explain the problem: JSON Schema ecosystem is fragmented, non-compliant, poor DX
- [ ] Explain the solution: native validators with compliance testing
- [ ] Include "two ways to use xschema" overview (framework vs library)
- [ ] Link to getting-started and language sections
- [ ] Typecheck passes

**Research:**
- `web/DOCUMENTATION_PLAN.md` - vision and problem statement
- `typescript/packages/client/README.md` - product description

---

### US-003: Create general getting-started page

**Description:** As a new user, I want a quick overview that helps me choose my path (language + mode) so I don't waste time reading irrelevant docs.

**Acceptance Criteria:**
- [ ] Create `/docs/getting-started.mdx`
- [ ] Decision tree: pick language → pick mode (framework vs library)
- [ ] Brief explanation of when to use framework vs library mode
- [ ] Cards/links to each language's getting started
- [ ] Typecheck passes

**Research:**
- `cli/language/` - see what languages are supported (currently just typescript)
- `typescript/packages/adapters/` - list of available adapters

---

### US-004: Create TypeScript overview page

**Description:** As a TypeScript developer, I want a language-specific landing page that shows me all my options.

**Acceptance Criteria:**
- [ ] Update `/typescript/index.mdx` as language hub
- [ ] Brief intro to xschema for TS
- [ ] Links to framework getting-started, library getting-started, adapters comparison
- [ ] List available adapters with quick links
- [ ] Typecheck passes

**Research:**
- `typescript/packages/adapters/` - list adapters (zod, arktype, effect, valibot)
- `typescript/packages/client/` - client package info

---

### US-005: Create TypeScript framework getting-started

**Description:** As a developer wanting full xschema integration, I want a step-by-step guide to set up the framework mode.

**Acceptance Criteria:**
- [ ] Create `/typescript/framework/getting-started.mdx`
- [ ] Step 1: Install dependencies (`@xschemadev/client`, adapter package, CLI)
- [ ] Step 2: Create `.xschema.jsonc` config file with example
- [ ] Step 3: Run `xschema generate`
- [ ] Step 4: Use generated schemas with `createXSchemaClient`
- [ ] Step 5: Extract types with `XSchemaType<"namespace:id">`
- [ ] Include working code example
- [ ] Typecheck passes

**Research:**
- `typescript/example/` - working example project
- `typescript/example/user.xschema.jsonc` - config file example
- `typescript/example/main.ts` - client usage example
- `typescript/example/.xschema/xschema.gen.ts` - generated output example
- `typescript/packages/client/README.md` - client API docs

---

### US-006: Create TypeScript framework config reference

**Description:** As a framework user, I want to understand all config file options so I can customize my setup.

**Acceptance Criteria:**
- [ ] Create `/typescript/framework/configuration.mdx`
- [ ] Document `$schema` URL requirement
- [ ] Document `schemas` array structure
- [ ] Document `id`, `sourceType`, `source`, `adapter` fields
- [ ] Document `sourceType` options: `file`, `url`, `json` (inline)
- [ ] Document namespace behavior (defaults to filename)
- [ ] Include examples for each source type
- [ ] Typecheck passes

**Research:**
- `cli/parser/parser.go` - config parsing, namespace derivation logic
- `cli/parser/types.go` - ConfigFileRaw, SchemaEntryRaw structs
- `cli/retriever/retriever.go` - sourceType handling, header substitution
- `typescript/example/user.xschema.jsonc` - file source example
- `typescript/example/subdir/another.xschema.jsonc` - URL source example
- `web/scripts/generate-schemas.ts` - JSON Schema for config validation

---

### US-007: Create TypeScript framework client reference

**Description:** As a framework user, I want to understand how to use the xschema client and type helpers.

**Acceptance Criteria:**
- [ ] Create `/typescript/framework/client.mdx`
- [ ] Document `createXSchemaClient` function
- [ ] Document `defaultNamespace` option
- [ ] Document schema lookup: `xschema("namespace:id")` and `xschema("id")` with default
- [ ] Document `XSchemaType<"namespace:id">` type helper
- [ ] Show examples with different adapters (Zod `.parse()`, ArkType `()`)
- [ ] Typecheck passes

**Research:**
- `typescript/packages/client/src/index.ts` - client implementation
- `typescript/packages/client/README.md` - usage examples
- `typescript/example/main.ts` - real usage example

---

### US-008: Create TypeScript library placeholder (TODO)

**Description:** As a developer who just wants to generate code without full framework integration, I want a placeholder indicating library mode is coming.

**Acceptance Criteria:**
- [ ] Create `/typescript/library/index.mdx` with TODO placeholder
- [ ] Brief explanation that library mode (direct adapter usage without client) is coming soon
- [ ] Link back to framework mode as current recommendation
- [ ] Typecheck passes

> **Note:** Library mode is not yet implemented. This is a placeholder for future work.

---

### US-009: Create TypeScript adapters comparison page

**Description:** As a TypeScript developer, I want to compare available adapters so I can choose the right one for my project.

**Acceptance Criteria:**
- [ ] Update `/typescript/adapters/index.mdx`
- [ ] Comparison table: adapter name, compliance %, bundle size (if known), unique features
- [ ] Brief description of each adapter's philosophy
- [ ] Links to each adapter's detail page
- [ ] Recommendation guidance (e.g., "Zod for most projects, ArkType for TS-native syntax")
- [ ] Typecheck passes

**Research:**
- `typescript/packages/adapters/*/package.json` - package names, versions, peer deps
- `typescript/packages/adapters/*/compliance/results/` - compliance percentages per draft
- `typescript/packages/adapters/*/src/index.ts` - adapter implementation patterns

---

### US-010: Update Zod adapter page

**Description:** As a developer considering Zod, I want complete information about the adapter.

**Acceptance Criteria:**
- [ ] Review and update `/typescript/adapters/zod/index.mdx`
- [ ] Ensure installation instructions are correct
- [ ] Ensure usage examples work with both framework and library mode
- [ ] Link to compliance page
- [ ] Document any Zod-specific features or limitations
- [ ] Typecheck passes

**Research:**
- `typescript/packages/adapters/zod/package.json` - package name, peer deps
- `typescript/packages/adapters/zod/src/` - implementation details
- `typescript/packages/adapters/zod/compliance/results/` - compliance data

---

### US-011: Update ArkType adapter page

**Description:** As a developer considering ArkType, I want complete information about the adapter.

**Acceptance Criteria:**
- [ ] Review and update `/typescript/adapters/arktype/index.mdx`
- [ ] Ensure installation instructions are correct
- [ ] Ensure usage examples work with both framework and library mode
- [ ] Link to compliance page
- [ ] Document any ArkType-specific features or limitations
- [ ] Typecheck passes

**Research:**
- `typescript/packages/adapters/arktype/package.json` - package name, peer deps
- `typescript/packages/adapters/arktype/src/` - implementation, uses `.narrow()` extensively
- `typescript/packages/adapters/arktype/compliance/results/` - compliance data

---

### US-012: Update Effect adapter page

**Description:** As a developer considering Effect Schema, I want complete information about the adapter.

**Acceptance Criteria:**
- [ ] Review and update `/typescript/adapters/effect/index.mdx`
- [ ] Ensure installation instructions are correct
- [ ] Ensure usage examples work with both framework and library mode
- [ ] Link to compliance page
- [ ] Document any Effect-specific features or limitations
- [ ] Typecheck passes

**Research:**
- `typescript/packages/adapters/effect/package.json` - package name, peer deps
- `typescript/packages/adapters/effect/src/` - implementation, uses `Schema as S`
- `typescript/packages/adapters/effect/compliance/results/` - compliance data

---

### US-013: Update Valibot adapter page

**Description:** As a developer considering Valibot, I want complete information about the adapter.

**Acceptance Criteria:**
- [ ] Review and update `/typescript/adapters/valibot/index.mdx`
- [ ] Ensure installation instructions are correct
- [ ] Ensure usage examples work with both framework and library mode
- [ ] Link to compliance page
- [ ] Document any Valibot-specific features or limitations
- [ ] Typecheck passes

**Research:**
- `typescript/packages/adapters/valibot/package.json` - package name, peer deps
- `typescript/packages/adapters/valibot/src/` - implementation
- `typescript/packages/adapters/valibot/compliance/results/` - 98.9% due to prototype bug

---

### US-014: Update compliance overview page

**Description:** As a user evaluating xschema, I want to understand what compliance means and why it matters.

**Acceptance Criteria:**
- [ ] Review `/compliance/index.mdx`
- [ ] Explain JSON Schema Test Suite
- [ ] Explain static generation approach and its benefits
- [ ] Explain coverage calculation (excludes unsupported features)
- [ ] Link to all adapter compliance pages
- [ ] Link to unsupported features page
- [ ] Typecheck passes

**Research:**
- `web/content/docs/compliance/index.mdx` - existing content
- `cli/compliance/` - how compliance testing works
- `cli/unsupported/unsupported-features.json` - what's excluded and why

---

### US-015: Verify compliance pages render correctly

**Description:** As a maintainer, I need to ensure auto-generated compliance pages work after restructure.

**Acceptance Criteria:**
- [ ] Verify Zod compliance page renders with tabs and accordions
- [ ] Verify ArkType compliance page renders
- [ ] Verify Effect compliance page renders
- [ ] Verify Valibot compliance page renders
- [ ] Verify unsupported-features page renders
- [ ] Verify in browser using dev server

**Research:**
- `web/scripts/generate-compliance.ts` - generation script, may need path updates
- `web/scripts/generate-unsupported-features.ts` - unsupported features generation

---

### US-016: Add fumadocs meta.json files for navigation

**Description:** As a user navigating the docs, I want logical sidebar structure that groups related content.

**Acceptance Criteria:**
- [ ] Create `meta.json` files to control sidebar order and grouping
- [ ] Group general pages (intro, getting-started, compliance)
- [ ] Group TypeScript section with sub-navigation
- [ ] Ensure adapters appear as expandable section
- [ ] Verify navigation works in browser

**Research:**
- `web/source.config.ts` - fumadocs configuration
- Fumadocs docs at https://fumadocs.vercel.app/docs/navigation - meta.json format

---

### US-017: Create internals overview page

**Description:** As a curious developer or contributor, I want to understand how xschema works under the hood.

**Acceptance Criteria:**
- [ ] Create `/docs/internals/index.mdx`
- [ ] High-level architecture diagram/explanation (parse → retrieve → process → generate → inject)
- [ ] Explain why this architecture makes adapters simple (bundled schemas, normalized drafts)
- [ ] Link to detailed pages (architecture, adapters, compliance-testing)
- [ ] Typecheck passes

**Research:**
- `cli/cmd/generate.go` - pipeline orchestration (parse → retrieve → process → generate → inject)
- `AGENTS.md` - architecture overview table

---

### US-018: Document URL source with headers

**Description:** As a developer fetching schemas from authenticated APIs, I want to understand how to configure headers with environment variables.

**Acceptance Criteria:**
- [ ] Add section to configuration.mdx about `sourceType: "url"`
- [ ] Document `headers` field (only valid for URL sources)
- [ ] Document `${VAR_NAME}` substitution syntax
- [ ] Document automatic `.env` loading from project root
- [ ] Document `--env-file` flag for custom env file path
- [ ] Show example with Authorization header
- [ ] Explain error message when env var is missing
- [ ] Typecheck passes

**Research:**
- `cli/retriever/retriever.go` - header substitution logic, `resolveHeaderValue()`
- `cli/config/env.go` - .env file loading
- `cli/parser/parser.go` - validation that headers only allowed for URL sourceType

---

### US-019: Document multiple config files behavior

**Description:** As a developer with a large project, I want to understand how multiple `.xschema.jsonc` files are handled.

**Acceptance Criteria:**
- [ ] Add section to configuration.mdx about multi-file projects
- [ ] Explain namespace merging (same namespace from different files merges)
- [ ] Explain duplicate ID detection (error with file paths shown)
- [ ] Explain output consolidation (single `xschema.gen.ts`)
- [ ] Show directory structure example
- [ ] Typecheck passes

**Research:**
- `cli/parser/parser.go` - `mergeDeclarations()`, duplicate detection
- `typescript/example/` - multi-file example (user.xschema.jsonc + subdir/another.xschema.jsonc)

---

### US-020: Create CLI reference page

**Description:** As a power user, I want a complete reference of all CLI commands and flags.

**Acceptance Criteria:**
- [ ] Create `/docs/cli.mdx`
- [ ] Document `xschema generate` command with all flags:
  - `--project` / `-p` (project root)
  - `--output` / `-o` (output directory, default `.xschema`)
  - `--lang` (filter by language)
  - `--verbose` / `-v`
  - `--dry-run` (show what would be generated)
  - `--concurrency` / `-c` (parallel fetches)
  - `--env-file` (custom .env path)
- [ ] Note that `--watch` is planned but not yet implemented
- [ ] Document config file discovery (git-based vs directory walk)
- [ ] Document runner auto-detection (bun.lock → bunx, pnpm-lock.yaml → pnpm exec)
- [ ] Typecheck passes

**Research:**
- `cli/cmd/generate.go` - all flags defined here
- `cli/cmd/root.go` - root command, version info
- `cli/parser/parser.go` - config discovery (git ls-files vs walk)
- `cli/language/langs/typescript/runner.go` - runner auto-detection

---

### US-021: Document JSON Schema draft support

**Description:** As a developer with legacy schemas, I want to understand which JSON Schema drafts are supported and how they're handled.

**Acceptance Criteria:**
- [ ] Add section to CLI reference or configuration page
- [ ] List supported drafts: draft3, draft4, draft6, draft7, draft2019-09, draft2020-12
- [ ] Explain that schemas are normalized to draft2020-12 internally
- [ ] Explain that `$schema` is auto-injected if missing
- [ ] Note that adapters receive normalized schemas
- [ ] Typecheck passes

**Research:**
- `cli/bundler/bundler.go` - `draftToSchemaURI` map, `normalizeLegacySyntax()`
- `cli/bundler/normalize.go` - draft normalization logic

---

### US-022: Document Valibot known limitation

**Description:** As a Valibot user, I want to know about the prototype property name bug so I can avoid it.

**Acceptance Criteria:**
- [ ] Add callout/warning to Valibot adapter page
- [ ] Explain that properties named `__proto__`, `constructor`, `toString` cause runtime errors
- [ ] Clarify this is a Valibot library bug, not xschema
- [ ] Recommend avoiding these property names or using a different adapter
- [ ] Typecheck passes

**Research:**
- `typescript/packages/adapters/valibot/compliance/results/` - see failing tests
- `typescript/packages/core/src/utils/primitives.ts` - `PROTOTYPE_PROPERTY_NAMES` constant

---

### US-023: Document output file structure

**Description:** As a developer, I want to understand what files xschema generates and where.

**Acceptance Criteria:**
- [ ] Add section to framework getting-started or configuration
- [ ] Document default output directory (`.xschema/`)
- [ ] Document main output file (`xschema.gen.ts`)
- [ ] Document manifest file (`xschema.manifest.json`) and its purpose
- [ ] Explain stale file cleanup on regenerate
- [ ] Explain `// Generated by xschema - DO NOT EDIT` header
- [ ] Typecheck passes

**Research:**
- `cli/injector/injector.go` - manifest handling, stale file cleanup
- `cli/language/langs/typescript/templates.go` - output template with header
- `typescript/example/.xschema/` - example output

---

### US-024: Document adapter naming requirements

**Description:** As a developer, I want to understand the adapter naming convention to avoid errors.

**Acceptance Criteria:**
- [ ] Add note to configuration or adapter pages
- [ ] Explain adapters must use scoped names: `@xschemadev/zod` not `zod`
- [ ] Show the error message for invalid adapter refs
- [ ] List all valid adapter names for TypeScript
- [ ] Typecheck passes

**Research:**
- `cli/language/langs/typescript/typescript.go` - `ValidateAdapterRef()` function, error messages
- `typescript/packages/adapters/*/xschema.adapter.json` - adapter declarations

---

### US-025: Create architecture deep-dive page

**Description:** As a curious developer, I want to understand the full processing pipeline in detail.

**Acceptance Criteria:**
- [ ] Create `/docs/internals/architecture.mdx`
- [ ] Document the 5-stage pipeline: parse → retrieve → process → generate → inject
- [ ] **Parser**: config discovery (git vs walk), namespace derivation, language detection
- [ ] **Retriever**: fetch from URL/file/inline, header substitution, caching
- [ ] **Processor**: ref extraction, external fetch loop, validation, bundling
- [ ] **Bundler**: $ref resolution, draft normalization (all → 2020-12), $defs embedding
- [ ] **Generator**: adapter invocation via stdin/stdout, parallel processing
- [ ] **Injector**: template rendering, manifest tracking, stale file cleanup
- [ ] Explain why adapters receive "easy" input (bundled, normalized, validated)
- [ ] Typecheck passes

**Research:**
- `cli/cmd/generate.go` - pipeline orchestration
- `cli/parser/parser.go` - config discovery, namespace derivation
- `cli/retriever/retriever.go` - schema fetching
- `cli/processor/processor.go` - crawl-fetch-validate-bundle pipeline
- `cli/bundler/bundler.go` - $ref resolution, $defs embedding
- `cli/bundler/normalize.go` - draft normalization
- `cli/generator/generator.go` - adapter invocation
- `cli/injector/injector.go` - file writing, manifest
- `AGENTS.md` - architecture table and package descriptions

---

### US-026: Create adapter development guide

**Description:** As a contributor wanting to build a new adapter, I want to understand the full adapter development process.

**Acceptance Criteria:**
- [ ] Create `/docs/internals/adapters.mdx`
- [ ] Document adapter protocol (stdin/stdout JSON)
- [ ] Document input format: `[{ namespace, id, varName, schema }]`
- [ ] Document output format: `[{ namespace, id, varName, schema, type, imports }]`
- [ ] Document @xschemadev/core package and its exports
- [ ] Document IR (Intermediate Representation) - the 18 node kinds
- [ ] Document `parse()` function (JSON Schema → IR)
- [ ] Document `createAdapterCLI()` helper
- [ ] Document key utilities: `escapeString`, `isPrimitive`, `hasPrototypeProperties`
- [ ] Explain that schemas are pre-bundled (no $ref resolution needed)
- [ ] Link to existing adapter source code as examples
- [ ] Typecheck passes

**Research:**
- `cli/adapter/types.go` - ConvertInput, ConvertResult structs
- `typescript/packages/core/src/ir/` - IR type definitions (SchemaNode, all 18 kinds)
- `typescript/packages/core/src/parser/` - parse() function
- `typescript/packages/core/src/cli.ts` - createAdapterCLI()
- `typescript/packages/core/src/utils/` - utilities (primitives, code-builder)
- `typescript/packages/adapters/zod/src/index.ts` - example adapter implementation

---

### US-027: Create compliance testing guide

**Description:** As an adapter developer, I want to understand how to run compliance tests.

**Acceptance Criteria:**
- [ ] Create `/docs/internals/compliance-testing.mdx`
- [ ] Document `xschema compliance` command
- [ ] Document flags: `--draft`, `--keyword`, `--dev-report`, `--profile`, `--adapter-path`
- [ ] Explain that it must be run from adapter directory (or use `--adapter-path`)
- [ ] Explain test suite download and caching (`~/.cache/xschema/`)
- [ ] Explain `--dev-report` writes results for docs generation
- [ ] Explain how compliance results flow into auto-generated docs
- [ ] Typecheck passes

**Research:**
- `cli/cmd/compliance.go` - all flags defined here
- `cli/compliance/fetcher.go` - test suite download, caching
- `cli/compliance/runner.go` - test execution
- `cli/compliance/report.go` - results output, --dev-report handling
- `web/scripts/generate-compliance.ts` - how results become docs

---

### US-028: Create troubleshooting page

**Description:** As a developer debugging issues, I want to understand common error messages and how to fix them.

**Acceptance Criteria:**
- [ ] Create `/docs/troubleshooting.mdx`
- [ ] Document "no xschema config files found" - need $schema URL in config
- [ ] Document "multiple languages detected" - use --lang flag
- [ ] Document "duplicate schema ID" - same ID in same namespace across files
- [ ] Document "headers are only allowed for sourceType url"
- [ ] Document "missing env var" - set env var or use --env-file
- [ ] Document "Unknown schema: X" client error - run `xschema generate`
- [ ] Document adapter CLI not found - ensure adapter is installed/built
- [ ] Typecheck passes

**Research:**
- `cli/parser/parser.go` - parser error messages
- `cli/retriever/retriever.go` - retriever error messages
- `cli/generator/generator.go` - generator error messages
- `typescript/packages/client/src/index.ts` - client error messages

---

### US-029: Add link to internals from introduction

**Description:** As a curious user reading the introduction, I want to easily find how xschema works internally.

**Acceptance Criteria:**
- [ ] Add "How it works" or "Under the hood" section/link in index.mdx or getting-started.mdx
- [ ] Brief teaser about the architecture (schemas are bundled, normalized, then passed to adapters)
- [ ] Link to `/docs/internals/` for deeper exploration
- [ ] Typecheck passes

---

### US-030: Add proper icons to navigation

**Description:** As a user browsing the docs, I want visual icons to help identify sections quickly.

**Acceptance Criteria:**
- [ ] Add TypeScript icon to `/typescript/` section
- [ ] Add Zod icon/logo to `/typescript/adapters/zod/`
- [ ] Add ArkType icon/logo to `/typescript/adapters/arktype/`
- [ ] Add Effect icon/logo to `/typescript/adapters/effect/`
- [ ] Add Valibot icon/logo to `/typescript/adapters/valibot/`
- [ ] Add appropriate icons for other sections (compliance, internals, CLI, etc.)
- [ ] Icons defined in frontmatter `icon` field or meta.json
- [ ] Verify icons render in sidebar navigation

**Research:**
- Fumadocs docs - how to add icons to navigation
- Look for existing icon usage in `web/content/docs/*/meta.json`
- SVG icons can be placed in `web/public/` or use lucide-react icons

---

## Functional Requirements

- FR-1: All pages must use fumadocs MDX components (Cards, Tabs, Accordions, etc.)
- FR-2: Code examples must be syntactically valid and copy-pasteable
- FR-3: Internal links must use relative paths that work after restructure
- FR-4: Auto-generated compliance pages must not be manually edited
- FR-5: Each language section must be self-contained (user shouldn't need to jump between languages)
- FR-6: Navigation sidebar must reflect the hierarchical structure

## Non-Goals

- Marketing/landing page (handled separately at `/`)
- Python documentation (future, same structure)
- Video tutorials or interactive examples
- Blog or changelog section

## Design Considerations

- Use fumadocs built-in components for consistency
- Keep code examples short and focused
- Use tabs for showing multiple approaches (npm/bun/pnpm, framework/library mode examples)
- Use cards for navigation and feature highlights
- Compliance tables should use color indicators (green/yellow/red)

## Technical Considerations

- Fumadocs with TanStack Start (already set up)
- `meta.json` files control fumadocs sidebar navigation
- Build must pass: `bun run build` in `web/` directory

### Auto-Generated Pages

Two scripts in `web/scripts/` generate documentation from source data:

**`generate-compliance.ts`**
- Discovers adapters from `content/docs/{lang}/adapters/{adapter}/index.mdx`
- Requires compliance results at `{lang}/packages/adapters/{adapter}/compliance/results/`
- Generates `compliance.mdx` with summary tables, tabs per draft, accordions for keywords
- **Build fails if adapter docs exist but compliance results are missing**

**`generate-unsupported-features.ts`**
- Reads from `cli/unsupported/unsupported-features.json`
- Generates `unsupported-features.mdx` with accordions for each limitation
- Currently 3 categories: dynamic-refs, recursive-refs, unevaluated-with-applicators

**`generate-schemas.ts`**
- Generates JSON Schema files for config validation
- Discovers adapters from `xschema.adapter.json` files
- Outputs to `web/public/schemas/`

### Adapter Discovery

The compliance script discovers adapters by scanning the docs structure, not the source packages. This means:
1. Add adapter docs first (`/typescript/adapters/{name}/index.mdx`)
2. Run compliance tests (`xschema compliance --dev-report`)
3. Build generates compliance.mdx automatically

## Success Metrics

- User can go from zero to working code in under 5 minutes following getting-started
- Clear path for both framework and library users
- All adapter compliance pages accessible and accurate
- No broken internal links
- Build passes without errors

## File Structure After Implementation

```
web/content/docs/
├── index.mdx                    # Introduction: what is xschema
├── getting-started.mdx          # Decision tree: language + mode
├── cli.mdx                      # CLI reference (generate command)
├── troubleshooting.mdx          # Common errors and solutions
├── meta.json
│
├── compliance/
│   ├── index.mdx                # Why compliance matters
│   ├── unsupported-features.mdx # (auto-generated)
│   └── meta.json
│
├── internals/                   # How xschema works under the hood
│   ├── index.mdx                # Overview + architecture intro
│   ├── architecture.mdx         # Deep-dive: parse → process → generate
│   ├── adapters.mdx             # Building new adapters
│   ├── compliance-testing.mdx   # Running compliance tests
│   └── meta.json
│
└── typescript/                  # (icon: TypeScript logo)
    ├── index.mdx                # TypeScript overview + links
    ├── meta.json
    │
    ├── framework/
    │   ├── getting-started.mdx  # Full setup guide
    │   ├── configuration.mdx    # Config file reference
    │   ├── client.mdx           # Client API reference
    │   └── meta.json
    │
    ├── library/
    │   ├── index.mdx            # TODO placeholder (not yet implemented)
    │   └── meta.json
    │
    └── adapters/
        ├── index.mdx            # Comparison table
        ├── meta.json
        ├── zod/                  # (icon: Zod logo)
        │   ├── index.mdx
        │   └── compliance.mdx   # (auto-generated)
        ├── arktype/             # (icon: ArkType logo)
        │   ├── index.mdx
        │   └── compliance.mdx
        ├── effect/              # (icon: Effect logo)
        │   ├── index.mdx
        │   └── compliance.mdx
        └── valibot/             # (icon: Valibot logo)
            ├── index.mdx        # includes known limitation warning
            └── compliance.mdx
```

## Content Details

### Key Sections by Page

**index.mdx (Introduction)**
- Problem: JSON Schema ecosystem is fragmented
- Solution: Native validators with compliance testing
- Two modes: Framework (full DX) vs Library (coming soon)
- Link to internals for curious readers

**configuration.mdx (Framework Config)**
- `$schema` URL requirement
- `sourceType` options: file, url, json
- `headers` with `${VAR}` substitution
- Namespace derivation rules
- Multiple config file behavior
- Output structure (.xschema/, manifest)

**cli.mdx (CLI Reference)**
- `xschema generate` with all flags
- Config discovery (git vs walk)
- Runner detection (bun/pnpm/npm)
- Draft support and normalization

**adapters/index.mdx (Comparison)**
- Table: name, compliance %, features
- Zod: most popular, 100% compliance
- ArkType: TS-native syntax, 100% compliance
- Effect: Effect ecosystem, 100% compliance  
- Valibot: smallest bundle, 98.9% (prototype bug)

**internals/index.mdx (Overview)**
- High-level pipeline: parse → retrieve → process → generate → inject
- Why this makes adapters simple
- Links to detailed pages

**internals/architecture.mdx (Deep Dive)**
- Parser: config discovery, namespace derivation
- Retriever: URL/file/inline fetching, caching
- Processor: ref crawling, validation, bundling
- Bundler: $ref resolution, draft normalization to 2020-12
- Generator: adapter stdin/stdout protocol
- Injector: template rendering, manifest tracking

**internals/adapters.mdx (Building Adapters)**
- Adapter protocol (stdin/stdout JSON)
- @xschemadev/core: IR types, parse(), createAdapterCLI()
- Key utilities for code generation
- Link to existing adapters as examples
