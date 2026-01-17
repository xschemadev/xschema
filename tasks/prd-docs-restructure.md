# PRD: Documentation Structure Restructure

## Introduction

Reorganize the (framework) documentation to improve navigation and reduce redundancy. Configuration reference moves to a more prominent position, and language-specific content is unified into single comprehensive pages per language.

## Goals

- Simplify navigation by removing nested folders
- Position configuration reference as a core concept (not buried in getting started)
- Unify setup + client reference into single language pages
- Add proper package manager tabs for all installation steps

## User Stories

### Move configuration reference in navigation

**Description:** As a user, I want configuration reference easily accessible after learning what xschema is, so I can understand the config format before diving into language-specific setup.

**Acceptance Criteria:**

- [ ] `configuration.mdx` appears directly after `what-is-xschema` in sidebar
- [ ] No section header between `what-is-xschema` and `configuration`
- [ ] All internal links to configuration still work
- [ ] Typecheck passes

### Create unified TypeScript language page

**Description:** As a TypeScript developer, I want one comprehensive page covering setup and client usage, so I don't have to navigate between multiple pages.

**Acceptance Criteria:**

- [ ] New `typescript.mdx` at `(framework)/typescript.mdx` (not in subfolder)
- [ ] Contains in-depth setup flow with package manager tabs (bun/npm/pnpm/yarn)
- [ ] Contains full client reference after setup section
- [ ] "Type-only" section no longer says "future"
- [ ] Links to configuration reference where appropriate
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### Create Python language placeholder page

**Description:** As a user exploring xschema, I want to see Python is planned so I know multi-language support is coming.

**Acceptance Criteria:**

- [ ] New `python.mdx` at `(framework)/python.mdx`
- [ ] Contains "coming soon" callout
- [ ] Shows planned workflow preview
- [ ] Links to TypeScript page and GitHub
- [ ] Typecheck passes

### Update navigation structure

**Description:** As a user, I want clean sidebar navigation that reflects the new structure.

**Acceptance Criteria:**

- [ ] `meta.json` updated with new page order
- [ ] "---Languages---" section header added
- [ ] Old `getting-started/` folder removed
- [ ] Old `typescript/` folder (client) removed
- [ ] No broken links in sidebar
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### Add package manager tabs to installation steps

**Description:** As a developer, I want to see installation commands for my preferred package manager without scrolling through alternatives.

**Acceptance Criteria:**

- [ ] Tabs component used for all install commands
- [ ] Tab options: bun, npm, pnpm, yarn
- [ ] Tabs for: package installation, running generator
- [ ] Consistent tab order across all instances
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

## Functional Requirements

- FR-1: Update `(framework)/meta.json` to new structure:
  ```json
  {
    "pages": [
      "index",
      "what-is-xschema",
      "configuration",
      "---Languages---",
      "typescript",
      "python",
      "---Help---",
      "faq",
      "troubleshooting"
    ]
  }
  ```
- FR-2: Create `(framework)/typescript.mdx` combining setup and client content
- FR-3: Create `(framework)/python.mdx` with coming soon content
- FR-4: Delete `(framework)/getting-started/` folder and its contents
- FR-5: Delete `(framework)/typescript/` folder and its contents
- FR-6: Use Tabs component from fumadocs-ui for package manager selection
- FR-7: Remove "(future)" from "Type-only" heading and description text

## Non-Goals

- No changes to adapter-specific usage section in client reference
- No changes to configuration.mdx content (only its position)
- No changes to other docs sections (cli, adapters, compliance, runtime)
- No new content beyond combining existing pages

## Design Considerations

- Use fumadocs `<Tabs>` component for package manager selection
- Maintain existing code examples and explanations
- Keep the same icon assignments where applicable
- TypeScript page should flow: Prerequisites -> Installation -> Config -> Generate -> Client Usage

## Technical Considerations

- Fumadocs uses file-based routing; folder structure = URL structure
- `meta.json` controls sidebar order and section headers
- Internal links may need updating if paths change
- Existing `(framework)/typescript/client.mdx` -> new `(framework)/typescript.mdx`

## Success Metrics

- Users can find configuration reference within 1 click from landing
- TypeScript setup + usage on single page (no navigation required)
- All package managers have equal visibility via tabs
