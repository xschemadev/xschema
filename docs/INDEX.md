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

### [2. Language Support](language-support.md)
**Understand how xschema detects and supports multiple programming languages**

- Language detection mechanism (via $schema URL)
- Language-agnostic core pipeline vs language-specific configuration
- Current implementations: TypeScript, Python
- Runner detection for package managers (bun, pnpm, npm, uv, poetry, etc.)
- Config file parsing and validation
- Requirements for adding new languages
- Detailed examples for adding Rust support

**Key Sections:**
- Language Detection - how $schema URLs map to languages
- Language-Specific Configuration - Language struct deep dive
- Runner Detection - for TypeScript and Python
- Adding a New Language - complete requirements checklist
- Configuration File Examples - TypeScript, Python, hypothetical Rust

### [3. CLI Pipeline](cli-pipeline.md)
**Detailed walkthrough of the entire xschema generation pipeline**

The pipeline has 4 steps: Parse → Retrieve → Generate → Inject

- **Step 1: Parse** - Find and validate config files, detect language, extract declarations
- **Step 2: Retrieve** - Fetch schemas from URLs, files, or inline JSON with retry logic
- **Step 3: Generate** - Call adapters via CLI, batch processing
- **Step 4: Inject** - Merge imports, execute templates, write output files

**Key Sections:**
- Pipeline Overview - visual diagram
- Each Step (1-4) - detailed explanation with code
- Configuration File Format - complete schema and examples
- Output File Structure - TypeScript and Python examples
- CLI Commands - usage and options
- Typical Workflow - real-world example project

### [4. Release Process](release-process.md)
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
│   ├── injector/               # File writing
│   └── language/               # Language configs
├── typescript/                  # TypeScript packages
│   ├── packages/
│   │   ├── core/               # Adapter protocol types
│   │   ├── client/             # Runtime client
│   │   └── zod/                # Zod adapter
├── .github/workflows/
│   └── release-please.yml      # Release automation
├── release-please-config.json  # Release configuration
└── docs/                        # Documentation (this directory)
    ├── adapter-system.md       # Adapter development
    ├── language-support.md     # Language support
    ├── cli-pipeline.md         # Generation pipeline
    ├── release-process.md      # Release workflow
    └── CONFORMANCE.md          # Compatibility specs
```

### Key Concepts

**Adapter Protocol**
- Stdin: JSON array of `{namespace, id, schema}`
- Stdout: JSON array of `{namespace, id, imports[], schema?, type?, validate?, validateImports?[]}`
- Binary discovery: `xschema-{adapter-name}` prefix
- Runner: auto-detected from lockfiles (bunx, npm, pnpm, etc.)

**Language Detection**
- Detected from `$schema` URL in config files
- Format: `https://xschema.dev/schemas/{ts|py}.jsonc`
- Can't mix languages in one project (use `--lang` flag)
- Language-specific: runner detection, output templates, imports

**Pipeline Steps**
1. **Parse**: Find configs, validate, extract declarations
2. **Retrieve**: Fetch schemas (with retry + cache + parallel)
3. **Generate**: Call adapters, batch by adapter name
4. **Inject**: Merge imports, execute template, write file

**Release Process**
- Conventional commits trigger semantic versioning
- Release-please creates PRs with version bumps + changelog
- Merge PR → automatic builds + publishes to npm + GitHub
- Multi-platform support: 6 Go binary platforms, 3 npm packages

## Document Size & Content

| Document | Lines | Topics | Audience |
|----------|-------|--------|----------|
| adapter-system.md | ~850 | Protocol, examples, testing, publishing | Adapter developers |
| language-support.md | ~600 | Detection, config, runner, adding languages | Language maintainers |
| cli-pipeline.md | ~900 | Each step, data types, examples, config format | Core developers |
| release-process.md | ~800 | Configuration, workflows, publishing, troubleshooting | Release managers |

**Total:** ~3,150 lines of documentation across 4 files

## Learning Path

**New to xschema?**
1. Read **CLI Pipeline** first - understand the 4 steps
2. Read **Language Support** - how languages are configured
3. Read **Adapter System** - how adapters work and extend it

**Building an adapter?**
1. Start with **Adapter System** - understand the protocol
2. Look at zod example in the document
3. Follow the step-by-step guide for your target library

**Adding language support?**
1. Read **Language Support** completely
2. Check the "Requirements for Adding a New Language" section
3. Use the Rust example as a template

**Managing releases?**
1. Read **Release Process** - understand the flow
2. See "How to Trigger a Release" section
3. Refer to troubleshooting for common issues

## Code References

Quick links to relevant source files:

**Adapter System:**
- Protocol types: `typescript/packages/core/src/types.ts`
- CLI helper: `typescript/packages/core/src/cli.ts`
- Zod example: `typescript/packages/zod/src/index.ts`

**Language Support:**
- Language config: `cli/language/language.go`
- Parser: `cli/parser/parser.go`
- Runner detection: `cli/language/language.go` (detectTSRunner, detectPythonRunner)

**CLI Pipeline:**
- Parse: `cli/parser/parser.go`
- Retrieve: `cli/retriever/retriever.go`
- Generate: `cli/generator/generator.go`
- Inject: `cli/injector/injector.go`
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
