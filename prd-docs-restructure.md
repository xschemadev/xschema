# PRD: Documentation Restructure

## Introduction

Restructure xschema documentation into 5 distinct sections with Fumadocs-style sidebar tabs. The current docs have duplicate content, confusing navigation paths, and mix concerns across sections. The new structure separates Framework (main experience), Runtime (programmatic usage), CLI (command reference), Adapters (library-specific docs), and Compliance (testing system).

## Goals

- Create 5 clearly separated documentation sections with sidebar tabs
- Eliminate duplicate and conflicting content
- Match user mental models and journeys
- Support multiple languages (TypeScript now, Python future) within sections
- Provide clear entry points for different user types

## User Stories

### US-001: Set up Fumadocs multi-section structure

**Description:** As a developer, I need the Fumadocs source configuration and folder structure set up correctly so that sidebar tabs work.

**Acceptance Criteria:**

- [ ] Create `content/docs/(framework)/` route group for default section
- [ ] Create `content/docs/runtime/`, `content/docs/cli/`, `content/docs/adapters/`, `content/docs/compliance/` folders
- [ ] Each section has `meta.json` with `root: true`, `title`, `description`, `icon`
- [ ] Root `content/docs/meta.json` defines section order
- [ ] Update `src/lib/source.ts` if needed for new structure
- [ ] Sidebar shows 5 tabs: Framework, Runtime, CLI, Adapters, Compliance
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-002: Create Framework section landing page

**Description:** As a user visiting /docs, I want to see an introduction with cards linking to all sections so I can navigate to what I need.

**Acceptance Criteria:**

- [ ] `(framework)/index.mdx` contains:
  - Brief intro (what xschema does in 2 sentences)
  - Cards linking to: Getting Started, What is XSchema, Adapters section, CLI section, Compliance section
  - Similar layout to Fumadocs landing page
- [ ] Cards use appropriate icons
- [ ] Page title is "Quick Start" or "Introduction"
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-003: Create "What is XSchema" deep-dive page

**Description:** As a user, I want to understand xschema's philosophy, architecture, and how it compares to alternatives so I can decide if it's right for me.

**Acceptance Criteria:**

- [ ] `(framework)/what-is-xschema.mdx` contains:
  - Philosophy section (why xschema exists, design principles)
  - Architecture overview (5-stage pipeline explained simply, link to internals for deep-dive)
  - Comparison with alternatives (json-schema-to-zod, typebox, etc.)
  - When to use xschema vs when not to
- [ ] Content adapted from current `internals/index.mdx` and `internals/architecture.mdx`
- [ ] Mermaid diagram showing the 5-stage pipeline (see US-004)
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-004: Set up Mermaid diagram support

**Description:** As a documentation author, I need Mermaid diagram support so I can visualize the xschema pipeline architecture.

**Acceptance Criteria:**

- [ ] Add `remarkMdxMermaid` plugin to `source.config.ts`:
  ```typescript
  import { remarkMdxMermaid } from 'fumadocs-core/mdx-plugins';
  import { defineConfig, defineDocs } from 'fumadocs-mdx/config';

  export const docs = defineDocs({
    dir: 'content/docs',
  });

  export default defineConfig({
    mdxOptions: {
      remarkPlugins: [remarkMdxMermaid],
    },
  });
  ```
- [ ] Create Mermaid component at `src/components/mdx/mermaid.tsx` (client component using mermaid library)
- [ ] Register Mermaid component in MDX components (in `src/routes/docs/$.tsx` or equivalent)
- [ ] Add pipeline diagram to `what-is-xschema.mdx`:
  - Deep dive into the architecture (see `internals/architecture.mdx` and CLI codebase)
  - Create the best possible Mermaid diagram representing the xschema pipeline
  - Show the flow from config/schemas through processing stages to generated output
  - Make it visually clear and useful for understanding the system
- [ ] Diagram renders correctly in browser
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-005: Create Framework getting started with language subpages

**Description:** As a user, I want a getting started guide for my language so I can set up xschema quickly.

**Acceptance Criteria:**

- [ ] `(framework)/getting-started/index.mdx` - overview that routes to language pages
- [ ] `(framework)/getting-started/typescript.mdx` - full TS setup guide (adapted from current `typescript/framework/getting-started.mdx`)
- [ ] `(framework)/getting-started/python.mdx` - "Coming soon" placeholder with brief overview
- [ ] TypeScript guide covers: prerequisites, installation, config file, run generate, use output
- [ ] Includes working code examples user can copy
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-006: Consolidate configuration documentation

**Description:** As a user, I want a single authoritative configuration reference so I don't encounter conflicting information.

**Acceptance Criteria:**

- [ ] `(framework)/configuration.mdx` contains complete config reference
- [ ] Remove duplicate `/docs/configuration.mdx` (was at root)
- [ ] Covers: $schema, namespace, schemas array, source types (file/url/json), headers, env vars
- [ ] Includes TypeScript and Python config examples (Python marked coming soon)
- [ ] Adapter options table for each language
- [ ] Content merged from current `configuration.mdx` and `typescript/framework/configuration.mdx`
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-007: Create Framework TypeScript client page

**Description:** As a TypeScript user, I want client reference documentation so I can use the generated schemas effectively.

**Acceptance Criteria:**

- [ ] `(framework)/typescript/client.mdx` covers:
  - `createXSchemaClient` API and options
  - `XSchemaType` utility type
  - Usage examples for each adapter (Zod, ArkType, Effect, Valibot)
  - Schema lookup patterns (full key vs default namespace)
  - Error handling
- [ ] Content adapted from current `typescript/framework/client.mdx`
- [ ] Fix any inconsistent API examples (especially Valibot)
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-008: Consolidate troubleshooting documentation

**Description:** As a user encountering errors, I want a single troubleshooting page with all error messages so I can find solutions quickly.

**Acceptance Criteria:**

- [ ] `(framework)/troubleshooting.mdx` contains all errors:
  - CLI errors (config not found, multiple languages, duplicate IDs, etc.)
  - Client errors (unknown schema - merged from `typescript/framework/troubleshooting.mdx`)
  - Schema validation errors
  - Compliance errors
- [ ] Remove `typescript/framework/troubleshooting.mdx` (merged into main)
- [ ] Organized by category with clear headings
- [ ] Each error has: message, cause, fix
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-009: Create FAQ page

**Description:** As a user with common questions, I want a FAQ page so I can find quick answers.

**Acceptance Criteria:**

- [ ] `(framework)/faq.mdx` contains common questions:
  - "Which adapter should I use?" (link to adapters comparison)
  - "Does xschema support $ref?" (yes, explain bundling)
  - "Why is feature X unsupported?" (link to compliance/unsupported-features)
  - "Can I use xschema with OpenAPI schemas?"
  - "How do I migrate from json-schema-to-zod?"
  - "Does xschema work at runtime?" (explain Framework vs Runtime modes)
  - "What JSON Schema drafts are supported?"
- [ ] Collapsible format for Q&A
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-010: Create Runtime section

**Description:** As a user who wants programmatic schema conversion, I want Runtime documentation so I can use xschema without a build step.

**Acceptance Criteria:**

- [ ] `runtime/meta.json` with `root: true`, title "Runtime", icon
- [ ] `runtime/index.mdx` - overview explaining:
  - What runtime mode is (programmatic conversion, no config file)
  - When to use it (library authors, dynamic schemas, build tools)
  - "Coming soon" callout for the actual implementation
  - Brief API preview showing the intended usage
- [ ] `runtime/getting-started/typescript.mdx` - placeholder with coming soon
- [ ] `runtime/getting-started/python.mdx` - placeholder with coming soon
- [ ] `runtime/api-reference.mdx` - placeholder documenting planned API:
  ```typescript
  import { convert } from '@xschema/runtime'
  const result = convert(jsonSchema, zodAdapter, { options })
  ```
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-011: Create CLI section

**Description:** As a user, I want complete CLI documentation so I can use all available commands and options.

**Acceptance Criteria:**

- [ ] `cli/meta.json` with `root: true`, title "CLI", icon "Terminal"
- [ ] `cli/index.mdx` - overview of CLI purpose and installation
- [ ] `cli/generate.mdx` - complete `xschema generate` reference:
  - All flags with descriptions
  - Usage examples
  - Common workflows
- [ ] `cli/compliance.mdx` - complete `xschema compliance` reference:
  - All flags with descriptions
  - How to run against adapters
  - Report generation (`--dev-report`)
  - Adapted from current `internals/compliance-testing.mdx`
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-012: Create Adapters section with comparison

**Description:** As a user, I want to compare adapters and find documentation for my chosen adapter.

**Acceptance Criteria:**

- [ ] `adapters/meta.json` with `root: true`, title "Adapters", icon
- [ ] `adapters/index.mdx` - comparison table and recommendations:
  - Table: adapter, peer dep, bundle size, compliance %
  - "Which to choose" guidance
  - Links to individual adapter pages
- [ ] Move adapter pages from `typescript/adapters/` to `adapters/`:
  - `adapters/zod/index.mdx` and `compliance.mdx`
  - `adapters/arktype/index.mdx` and `compliance.mdx`
  - `adapters/valibot/index.mdx` and `compliance.mdx`
  - `adapters/effect/index.mdx` and `compliance.mdx`
- [ ] Each adapter page covers: overview, installation, usage, generated output
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-013: Create "Building Adapters" guide in Adapters section

**Description:** As a contributor, I want documentation on building new adapters so I can extend xschema.

**Acceptance Criteria:**

- [ ] `adapters/building-adapters.mdx` contains:
  - Adapter protocol (stdin/stdout JSON format)
  - Input/output types (ConvertInput, ConvertResult)
  - Using @xschemadev/core (createAdapterCLI, parse, utilities)
  - Intermediate Representation (IR) node types
  - Step-by-step guide to create new adapter
  - Testing with compliance suite
- [ ] Content adapted from current `internals/adapters.mdx`
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-014: Create Compliance section

**Description:** As a user, I want to understand the compliance system so I can trust xschema and know its limitations.

**Acceptance Criteria:**

- [ ] `compliance/meta.json` with `root: true`, title "Compliance", icon
- [ ] `compliance/index.mdx` - what compliance means:
  - JSON Schema Test Suite explanation
  - Why compliance matters
  - How to read compliance reports
  - Link to adapter-specific reports
- [ ] `compliance/unsupported-features.mdx` - move from current location (auto-generated)
- [ ] `compliance/how-it-works.mdx` - internals of compliance testing:
  - The testing pipeline
  - How unsupported features are detected
  - Coverage calculation
  - Adapted from parts of `internals/compliance-testing.mdx` and `internals/architecture.mdx`
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-015: Update navigation meta.json files

**Description:** As a user, I want logical sidebar navigation within each section so I can find content easily.

**Acceptance Criteria:**

- [ ] Root `meta.json` orders sections: `["(framework)", "runtime", "cli", "adapters", "compliance"]`
- [ ] `(framework)/meta.json` navigation:
  ```
  ---Introduction---
  index, what-is-xschema
  ---Getting Started---
  getting-started
  ---Reference---
  configuration
  ---TypeScript---
  typescript/client
  ---Python---
  python (future)
  ---Help---
  troubleshooting, faq
  ```
- [ ] Each section has logical navigation order
- [ ] Section separators (`---Name---`) used appropriately
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

### US-016: Remove deprecated/duplicate files

**Description:** As a maintainer, I want to remove old files so the codebase stays clean.

**Acceptance Criteria:**

- [ ] Remove `/docs/configuration.mdx` (consolidated into framework)
- [ ] Remove `/docs/getting-started.mdx` (replaced by framework structure)
- [ ] Remove `/docs/cli.mdx` (moved to cli section)
- [ ] Remove `/docs/typescript/framework/troubleshooting.mdx` (merged)
- [ ] Remove `/docs/typescript/library/index.mdx` (placeholder, runtime replaces concept)
- [ ] Remove `/docs/typescript/index.mdx` (no longer needed, adapters are language-agnostic section)
- [ ] Remove `/docs/typescript/framework/` folder (content moved)
- [ ] Remove `/docs/typescript/adapters/` folder (content moved to adapters section)
- [ ] Remove `/docs/internals/` folder (content distributed to relevant sections)
- [ ] Verify no broken links
- [ ] Typecheck passes

---

### US-017: Fix Valibot client example inconsistency

**Description:** As a user, I want consistent API examples across all adapter docs so I don't get confused.

**Acceptance Criteria:**

- [ ] Valibot adapter page uses same client API pattern as other adapters:
  ```typescript
  const xschema = createXSchemaClient({
    schemas,
    defaultNamespace: "user",
  });
  ```
- [ ] Import style matches other adapters:
  ```typescript
  import { schemas } from "./.xschema/xschema.gen";
  ```
- [ ] All adapter pages have consistent example structure
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

---

## Functional Requirements

- FR-1: Documentation must use Fumadocs route groups (`(framework)`) for default section
- FR-2: Each section must have `meta.json` with `root: true` to appear as sidebar tab
- FR-3: Landing page at `/docs` must show cards linking to all major sections
- FR-4: Configuration reference must be single source of truth (no duplicates)
- FR-5: All code examples must be copy-pasteable and working
- FR-6: Language-specific content must live in subfolders (e.g., `getting-started/typescript.mdx`)
- FR-7: "Coming soon" features must be clearly marked with callouts
- FR-8: Auto-generated files (compliance reports) must be preserved in new locations
- FR-9: All internal links must be updated to new paths
- FR-10: Sidebar navigation must be logical and match user expectations

## Non-Goals

- No actual implementation of Runtime mode (docs only describe planned API)
- No Python documentation content (only placeholder pages)
- No new features or code changes outside documentation
- No design/styling changes to Fumadocs theme
- No internationalization (English only)

## Design Considerations

- Use Fumadocs `<Cards>` component for landing page navigation
- Use Fumadocs `<Callout>` for "coming soon" notices
- Use Fumadocs `<Tabs>` for language-specific code examples where inline
- Maintain current styling and theme
- Icons should use Lucide icon names in meta.json

## Technical Considerations

- Route group `(framework)` maps to `/docs/` (parentheses hidden from URL)
- Other sections map to `/docs/{section}/`
- Source config in `src/lib/source.ts` may need adjustment for new structure
- Auto-generated compliance files need script updates for new paths
- Internal links use relative paths or absolute from `/docs/`

## Success Metrics

- Zero duplicate content across documentation
- All 5 sections visible as sidebar tabs
- Users can complete getting started in under 5 minutes
- All code examples are copy-paste functional
- No broken internal links

## File Structure

```
content/docs/
├── meta.json                           # Section order
├── (framework)/                        # Default section (URL: /docs/)
│   ├── meta.json                       # root: true, navigation
│   ├── index.mdx                       # Landing with cards
│   ├── what-is-xschema.mdx            # Philosophy, architecture, comparison
│   ├── getting-started/
│   │   ├── index.mdx                   # Routes to language pages
│   │   ├── typescript.mdx              # TS setup guide
│   │   └── python.mdx                  # Coming soon
│   ├── configuration.mdx               # Complete config reference
│   ├── typescript/
│   │   └── client.mdx                  # Client API reference
│   ├── python/                         # Future
│   ├── troubleshooting.mdx             # All errors consolidated
│   └── faq.mdx                         # Common questions
│
├── runtime/                            # Programmatic usage (URL: /docs/runtime/)
│   ├── meta.json                       # root: true
│   ├── index.mdx                       # Overview + coming soon
│   ├── getting-started/
│   │   ├── typescript.mdx              # Coming soon
│   │   └── python.mdx                  # Coming soon
│   └── api-reference.mdx               # Planned API docs
│
├── cli/                                # CLI reference (URL: /docs/cli/)
│   ├── meta.json                       # root: true
│   ├── index.mdx                       # Overview
│   ├── generate.mdx                    # generate command
│   └── compliance.mdx                  # compliance command
│
├── adapters/                           # Adapter docs (URL: /docs/adapters/)
│   ├── meta.json                       # root: true
│   ├── index.mdx                       # Comparison table
│   ├── zod/
│   │   ├── index.mdx
│   │   └── compliance.mdx              # Auto-generated
│   ├── arktype/
│   │   ├── index.mdx
│   │   └── compliance.mdx
│   ├── valibot/
│   │   ├── index.mdx
│   │   └── compliance.mdx
│   ├── effect/
│   │   ├── index.mdx
│   │   └── compliance.mdx
│   └── building-adapters.mdx           # Contributor guide
│
└── compliance/                         # Compliance system (URL: /docs/compliance/)
    ├── meta.json                       # root: true
    ├── index.mdx                       # What compliance means
    ├── unsupported-features.mdx        # Auto-generated
    └── how-it-works.mdx                # Testing internals
```
