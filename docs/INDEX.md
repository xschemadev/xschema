# xschema Documentation Index

Complete documentation of the xschema codebase, architecture, and processes.

## Core Documentation

### [1. Adapter System](adapter-system.md)
**Learn how to extend xschema with new validators**

- Adapter protocol (stdin/stdout JSON interface)
- Input/output format specifications
- How adapters are discovered and invoked
- Complete guide to creating a new TypeScript adapter
- Example: @xschemadev/zod adapter walkthrough
- Testing and publishing adapters

**Key Sections:**
- Adapter Protocol - JSON input/output format
- How Adapters Are Called - discovery, runner detection, batch processing
- Creating a New Adapter - step-by-step guide with examples
- Testing Your Adapter - manual and automated testing

### [2. CLI Pipeline](cli-pipeline.md)
**Detailed walkthrough of the entire xschema generation pipeline**

The pipeline has 4 steps: Parse → Retrieve → Generate → Inject

- **Step 1: Parse** - Find and validate config files, detect language, extract declarations
- **Step 2: Retrieve** - Fetch schemas from URLs, files, or inline JSON with retry logic
- **Step 3: Generate** - Call adapters via CLI, batch processing
- **Step 4: Inject** - Merge imports, execute templates, write output files with stale cleanup

**Key Sections:**
- Pipeline Overview - visual diagram
- Each Step (1-4) - detailed explanation with code
- Configuration File Format - complete schema and examples
- Output File Structure - TypeScript examples with multi-file support
- Manifest-based stale file cleanup
- CLI Commands - usage and options
- Typical Workflow - real-world example project

### [3. Releasing](RELEASING.md)
**Complete guide to the automated release workflow**

- Release-please configuration and setup
- How version bumping works (Semantic Versioning)
- Build and publish workflows for all platforms
- Published packages (npm packages, GitHub releases)
- How to trigger a release
- Installation methods
- Troubleshooting

**Key Sections:**
- Overview - release-please basics
- Release Configuration - configuration files explained
- Release Flow - detailed 5-step process
- Build and Publish Workflows - job descriptions and platform support
- Published Packages - complete package catalog
- How to Trigger a Release - step-by-step instructions
- Installation - from npm or GitHub releases

## Quick Reference

### File Structure

```
xschema/
├── cli/                         # Go CLI
│   ├── cmd/                    # Cobra commands
│   ├── parser/                 # Config parsing
│   ├── retriever/              # Schema fetching
│   ├── generator/              # Adapter calling
│   ├── injector/               # File writing + manifest cleanup
│   └── language/               # Language system
│       ├── registry.go         # Language registry
│       ├── spec.go             # Core types (GeneratedFile, CommandSpec)
│       └── langs/              # Per-language implementations
│           └── typescript/     # TypeScript language
├── typescript/                  # TypeScript packages
│   ├── packages/
│   │   ├── core/               # Adapter protocol types
│   │   ├── client/             # Runtime client
│   │   └── adapters/           # Adapter implementations
│   │       ├── zod/            # Zod adapter
│   │       ├── valibot/        # Valibot adapter
│   │       └── ...             # Other adapters
├── .github/workflows/
│   └── release-please.yml      # Release automation
├── release-please-config.json  # Release configuration
└── docs/                        # Documentation (this directory)
    ├── adapter-system.md       # Adapter development
    ├── cli-pipeline.md         # Generation pipeline
    └── RELEASING.md            # Release workflow
```

### Key Concepts

**Adapter Protocol**
- Stdin: JSON array of `{namespace, id, schema}`
- Stdout: JSON array of `{namespace, id, imports[], schema?, type?, validate?, validationImports?[]}`
- Adapter refs: scoped npm packages like `@xschemadev/zod`
- Binary: derived from ref as `xschema-<package>` (e.g., `xschema-zod`)
- Runner: auto-detected from lockfiles (bunx, npm, pnpm, etc.)

**Language Detection**
- Detected from `$schema` URL in config files
- Format: `https://xschema.dev/schemas/typescript.jsonc`
- Currently TypeScript-only (see `cli/language/langs/` for adding languages)
- Language-specific: runner detection, output templates, imports, adapter invocation

**Pipeline Steps**
1. **Parse**: Find configs, validate, extract declarations
2. **Retrieve**: Fetch schemas (with retry + cache + parallel)
3. **Generate**: Call adapters via `AdapterInvoker`, batch by adapter ref
4. **Inject**: Merge imports, execute template, write files with manifest cleanup

**Multi-File Outputs**
- Output written via `WriteGeneratedFiles(outDir, []GeneratedFile)`
- Manifest at `outDir/xschema.manifest.json` tracks generated files
- Stale files from previous runs are automatically deleted
- Manifest write is atomic (temp + rename)

**Release Process**
- Conventional commits trigger semantic versioning
- Release-please creates PRs with version bumps + changelog
- Merge PR → automatic builds + publishes to npm + GitHub
- Multi-platform support: 6 Go binary platforms, npm packages

## Document Size & Content

| Document | Lines | Topics | Audience |
|----------|-------|--------|----------|
| adapter-system.md | ~850 | Protocol, examples, testing, publishing | Adapter developers |
| cli-pipeline.md | ~900 | Each step, data types, examples, config format | Core developers |
| RELEASING.md | ~240 | Configuration, workflows, publishing, troubleshooting | Release managers |

## Learning Path

**New to xschema?**
1. Read **CLI Pipeline** first - understand the 4 steps
2. Read **Adapter System** - how adapters work and extend it
3. Explore `cli/language/` - see how languages plug in

**Building an adapter?**
1. Start with **Adapter System** - understand the protocol
2. Look at zod example: `typescript/packages/adapters/zod/`
3. Each adapter needs `xschema.adapter.json` with `{"ref": "@xschemadev/<name>"}`

**Adding language support?**
1. See TypeScript implementation: `cli/language/langs/typescript/`
2. Implement `Language` struct with `AdapterInvoker` interface
3. Register via `language.Register()` in `cli/language/langs/langs.go`

**Managing releases?**
1. Read **RELEASING.md** - understand the flow
2. See "How to Trigger a Release" section
3. Refer to troubleshooting for common issues

## Code References

Quick links to relevant source files:

**Adapter System:**
- Protocol types: `typescript/packages/core/src/types.ts`
- CLI helper: `typescript/packages/core/src/cli.ts`
- Zod adapter: `typescript/packages/adapters/zod/`
- Adapter metadata: `typescript/packages/adapters/*/xschema.adapter.json`

**Language System:**
- Registry: `cli/language/registry.go`
- Core types: `cli/language/spec.go` (GeneratedFile, CommandSpec, AdapterInvoker)
- Path helpers: `cli/language/paths.go`, `cli/language/order.go`
- TypeScript impl: `cli/language/langs/typescript/typescript.go`

**CLI Pipeline:**
- Parse: `cli/parser/parser.go`
- Retrieve: `cli/retriever/retriever.go`
- Generate: `cli/generator/generator.go`
- Inject: `cli/injector/injector.go` (includes WriteGeneratedFiles + manifest)
- Main command: `cli/cmd/generate.go`

**Release Process:**
- Workflow: `.github/workflows/release-please.yml`
- Config: `release-please-config.json`
- Manifest: `.release-please-manifest.json`

## External Resources

- [Conventional Commits](https://www.conventionalcommits.org/)
- [Release Please](https://github.com/googleapis/release-please)
- [json-schema-to-zod](https://github.com/StefanTerdell/json-schema-to-zod)
- [JSON Schema Specification](https://json-schema.org/)

## Contributing

When contributing to xschema, keep documentation in sync:

1. **Code changes** → Update relevant docs
2. **New adapter** → Document in Adapter System
3. **New language** → Document in Language Support
4. **Pipeline changes** → Update CLI Pipeline
5. **Release updates** → Update Release Process

Documentation lives in `/docs/` directory.
