# xschema Landing Page — Text Content & UX Notes

---

## Section 1: Hero (Already Implemented)

**Pill badge:** "Bring your JSON Schemas to life."

**Headline:**

> Cross-language type-safety,
> but it's _not_ a f@#!ing jungle.

**Description (below hero box):**

> xschema is an ecosystem of tools to bring cross-language type-safety validation in your codebases. It leverages the power of JSON Schema to validate your data across languages.

**CTAs:**

- Primary: "Getting Started" -> `/docs`
- Secondary: "Open in StackBlitz" -> StackBlitz link

---

## Section 2: One pipeline to rule them all

**Heading:** `One pipeline to rule them all.`

### 2a: The Problem

**Title:** `The jungle problem.`

**Subheading:** `Today, every library has its own converter.`

**Body:**

> Need to convert JSON Schema to a Zod validator? There's a package for that. Pydantic? Another one. ArkType? You're on your own.
>
> Each converter follows a different flow, supports a different subset of the spec, and there's no standard way to verify correctness. You end up with:

- **No shared pipeline** — each tool re-implements parsing, normalization, and reference resolution from scratch.
- **No correctness guarantees** — none of them test against the official JSON Schema Test Suite.
- **No transparency** — unsupported features fail silently or produce wrong output.
- **No portability** — switching validators means switching converters, configs, and workflows.

### 2b: The xschema Approach

**Title:** `The xschema approach.`

**Subheading:** `One pipeline. Any adapter. Verified.`

**Body:**

> xschema runs one pipeline for every adapter: normalize the schema, resolve all references, validate it, then hand it off to the adapter for code generation. The adapter only does one thing — convert a clean schema to native code.
>
> Every adapter is tested against the official JSON Schema Test Suite.

**Comparison:**

| What you get             | xschema                                       | One-off converters               |
|--------------------------|-----------------------------------------------|----------------------------------|
| Shared pipeline          | One CLI for all adapters                      | Each tool builds its own         |
| Compliance testing       | Official JSON Schema Test Suite               | Often absent or incomplete       |
| Supported features       | Documented per adapter with exact coverage    | Undocumented or incomplete       |
| Unsupported features     | Explicitly listed with rationale              | Silent failures                  |
| Validators               | Zod, ArkType, Effect, Valibot, Pydantic, more | One converter per library        |
| Languages                | TypeScript, Python, Go, Rust                  | One tool per language            |

### 2c: Compliance

**Callout:**

**Title:** `Built for transparency.`

**Body:**

> **Compliance report**: Tested against the official JSON Schema Test Suite. Every adapter reports its exact coverage. Supported features, unsupported features, and the technical reason behind each gap — all documented, all tested. Not a footnote, a whole documentation page for each adapter.

## Section 3: Built For

**Heading:** `Built for.`

### Card 1 — Application Developers

**Title:** `Developers shipping real projects.`

**Body:**

> You have JSON Schemas — from an API contract, an OpenAPI spec, a schema registry, or a local definition — and you need native validators in your codebase.
>
> xschema converts them to Zod, ArkType, Valibot, Pydantic, and others. Define your schemas once, run `xschema generate`, use the output directly. TypeScript, Python, Go, Rust — one unified workflow.

**Key point:** `No manual translation. No drift between schema and validator.`

**Framework Mode**

- Heading: `Framework Mode`
- Badge: `Build-time`
- Body: Schemas are declared in config files and converted at build time. Run `xschema generate` to produce native validators with full static type inference. Zero runtime cost.
- Bullets:
  - Full type inference and autocomplete
  - Tree-shakeable output
  - No runtime conversion cost
  - Schemas from files, URLs, or inline definitions

### Card 2 — Library & Framework Maintainers

**Title:** `Library maintainers who need codegen.`

**Body:**

> You maintain a library or framework that needs to generate validators from JSON Schema, but you don't want to own the conversion logic.
>
> xschema adapters handle it. Each one is tested against the official JSON Schema Test Suite, with clear reporting of what is supported and what isn't. Use Framework Mode to generate at build time, or Runtime Mode to convert programmatically inside your library.

**Key point:** `Offload the conversion. Focus on your domain.`

**Runtime Mode**

- Heading: `Runtime Mode`
- Badge: `Programmatic`
- Body: Convert schemas programmatically at runtime. When codegen is part of the workflow.
- Bullets:
  - Programmatic `convert()` API
  - Dynamic schema sources
  - Per-schema conversion cost

---

## Section 4: Under the Hood

**Heading:** `Under the hood.`

### 4a: Pipeline

Display as a horizontal flow (stacked on mobile):

```text
Parse -> Retrieve -> Process -> Generate -> Inject
```

| Stage        | Description                                              |
|--------------|----------------------------------------------------------|
| **Parse**    | Discovers config files and extracts schema declarations. |
| **Retrieve** | Fetches schemas from files, URLs, or inline definitions. |
| **Process**  | Normalizes, validates, resolves references, bundles.     |
| **Generate** | Delegates code generation to the target adapter.         |
| **Inject**   | Writes generated validators and types to your project.   |

**Caption:**

> The CLI does the heavy lifting. Adapters receive clean, self-contained schemas and only handle code generation.

### 4b: Components (grid of 4 cards)

**Card 1 — CLI**

- Title: `CLI`
- Body: Orchestrates the full pipeline. Parses configs, fetches schemas, normalizes drafts, resolves references, and delegates generation to adapters.
- Detail: `npx xschema generate`

**Card 2 — Adapters**

- Title: `Adapters`
- Body: Each adapter targets one validation library and converts normalized schemas into native validators. One adapter per library, one protocol for all.
- Detail: `Zod · ArkType · Effect Schema · Valibot · Pydantic · and more`

**Card 3 — Client**

- Title: `Client`
- Body: Language-specific packages that expose generated schemas with full type inference and autocomplete.
- Detail: `@xschemadev/client`

**Card 4 — Config**

- Title: `Config`
- Body: Declarative config files define which schemas to process, where to fetch them, and which adapter to use.
- Detail: `*.xschema.jsonc`

---

## Section 5: Try It

**Heading:** `Try it.`

**Subheading:** `Write a schema. Pick an adapter. See the output.`

**Implementation:** StackBlitz embed with a pre-configured xschema project.

**Layout:** Side-by-side (tabbed on mobile). Left: editable JSON Schema input. Right: generated validator output.

**Pre-filled example:**

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "properties": {
    "id": { "type": "string", "format": "uuid" },
    "name": { "type": "string", "minLength": 1 },
    "email": { "type": "string", "format": "email" },
    "role": { "enum": ["admin", "user", "viewer"] }
  },
  "required": ["id", "name", "email"]
}
```

**CTAs:**

- Primary: "Read the docs" -> `/docs`
- Secondary: "View on GitHub" -> repo link

---

## Section 6: Open Source

**Heading:** `Open source. Open process.`

**Body:**

> xschema is open source. Contributions, adapter proposals, and bug reports are greatly appreciated.

**CTAs:**

- "Read the docs" -> `/docs`
- "GitHub" -> repo link

---

## UI/UX Suggestions

### Comparison grid (Section 2b)

Render as a styled grid, not a raw HTML table. Two columns: xschema (highlighted) vs one-off converters (muted). Check/cross icons per row convey the difference faster than text in cells.
Use and modify the already present comparison-table component.

### Compliance callout (Section 2c)

Callout/banner with a distinct background.

### Built For cards (Section 3)

Two equal-width cards. The "key point" at the bottom of each card should be visually separated — border-top, slightly bolder text, or a different background shade. It's the takeaway line. In mobile, render with selectable tabs

### Pipeline visual (Section 4a)

Consider an animated or step-through diagram similar to the existing `<CliAnimation />`. On hover or scroll, each stage highlights and its description appears. This makes the pipeline feel tangible without requiring interaction.
Experiment with cpu-architecture.tsx, database-with-rest-api.tsx and google-gemini-effect.tsx.
The idea is to have the animated pipeline on the left, interactable, and the description of each component that appears on the right.
The small components cards are connected with highlighted lines, with a light that moves from one to the other (see the examples).
In mobile, they should be one under the other.

### Component cards (Section 4b)

Anchor each card with a monospace detail line at the bottom (e.g. `npx xschema generate`). Something concrete that the reader can immediately recognize.
Use the features-8 component. Rename it and modify it to fit current style.

### Try It embed (Section 5) SKIP FOR NOW

Include an adapter selector dropdown above the output panel so users can switch between Zod, ArkType, Valibot, etc. without editing config. Pre-run generation on load so output is visible immediately.

### Open Source footer (Section 6)

Minimal. One line of text, two buttons. No contributor grids or star counters — those add noise at this stage.

### General layout

- Max width `container` with horizontal padding (matching current hero container).
- Section headings: `text-brand`, uppercase monospace, small — matching the "One pipeline to rule them all." style already in the hero area.
- Generous vertical spacing between sections (`mt-16` to `mt-24`).
- Cards: use existing `cardVariants` for visual consistency.

### Tone

No "blazing fast", "revolutionary", "seamless", "powerful", "game-changing". Precise technical terms: validator, adapter, code generation, type inference, compliance, normalization. Short sentences. Declarative. The architecture speaks for itself.
