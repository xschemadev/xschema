# xschema Agent Guidelines

JSON Schema to native validators (Zod, Pydantic, etc.) with full type safety.

## Quality Requirements

Before submitting changes, ensure ALL checks pass:

### Go CLI (from `cli/` directory)

```bash
go build -o xschema .                   # must compile
go vet ./...                            # must pass (no warnings)
go test ./...                           # must pass (all tests green)
```

### TypeScript (from `typescript/` directory)

```bash
bun run build                           # must compile
bun run typecheck                       # must pass (no type errors)
```

### Adapter Changes (from adapter directory, e.g. `typescript/packages/adapters/zod/`)

```bash
bun run compliance                      # must pass (if adapter code changed)
```

When I ask "run compliance on an adapter", that's the command that should be run

### Commits

```bash
bunx commitlint --from HEAD~1 --to HEAD # commit message must be valid
```

**Run relevant checks for the code you changed.** If you modified Go code, run Go checks. If you modified TypeScript, run TS checks. Always verify builds succeed.

## Project Structure

```
cli/                    # Go CLI (github.com/xschemadev/xschema)
  cmd/                  # cobra commands (root, generate, compliance)
  parser/               # parses JSON/JSONC config files
  retriever/            # fetches schemas from URL/file/inline
  generator/            # calls adapters to convert schemas
  injector/             # writes generated code
  language/             # language-specific config (TS, Python)
  bundler/              # bundles schemas
  compliance/           # adapter compliance testing
  ui/                   # terminal output helpers

typescript/             # TS packages (bun workspace)
  packages/core/        # @xschemadev/core - IR types, parser, utils
  packages/adapters/    # adapter packages (zod, arktype)
  packages/client/      # @xschemadev/client - runtime client
  example/              # example project
```

## Build/Test/Lint Commands

### Go CLI (run from `cli/` directory)

```bash
go test ./...                           # run all tests
go test ./parser/                       # run single package tests
go test ./parser/ -run TestParse        # run single test by name
go test ./parser/ -run TestParse -v     # verbose output
go test ./... -short                    # skip integration tests
go test . -run TestIntegration          # integration tests only
go build -o xschema .                   # build binary
go vet ./...                            # vet (run in CI)
go fmt ./...                            # format code
```

### TypeScript (run from `typescript/` directory)

```bash
bun install                             # install deps
bun run build                           # build all packages (core first)
bun run typecheck                       # type check all packages
bunx tsc --noEmit                       # type check single package (in pkg dir)
```

### Web (run from `web/` directory)

```bash
bun install                             # install deps
bun test                                # run bun tests
bun run types:check                     # run tsc (requires fumadocs-mdx)
bun run generate:schemas                # regenerate web json schemas
```

Note: `web/tsconfig.json` pins `compilerOptions.types`, so include both `bun` and `node` when scripts/tests import `bun:test` or `node:*`.

### Root (commitlint/husky)

```bash
bun install                             # install commitlint + husky
bunx commitlint --from HEAD~1 --to HEAD # validate last commit
```

## Code Style Guidelines

### Go Code

**Imports**: stdlib first, external second, internal third (blank lines between):

```go
import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "github.com/xschemadev/xschema/language"
    "github.com/xschemadev/xschema/ui"
)
```

**Naming**: Exported `PascalCase`, unexported `camelCase`, acronyms uppercase (`URL`, `ID`)

**Error handling**: wrap with context `fmt.Errorf("failed to X: %w", err)`, return early

**Logging**: use `ui.Verbosef()` for debug output (shown with `--verbose` flag)

**Functions**: accept `context.Context` as first param for cancellable operations

**Types**: define close to usage, use `json.RawMessage` for arbitrary JSON, prefer structs over maps

### TypeScript Code

**Imports**: external packages first, then relative (blank line between):

```typescript
import type { ConvertInput, ConvertResult } from "@xschemadev/core";
import { parse } from "@xschemadev/core";

import { render } from "./renderer.js";
```

**Naming**: types/interfaces `PascalCase`, functions/variables `camelCase`

**Types**: use `interface` for object shapes, `type` for unions/aliases

**Exports**: named exports preferred, use `export type { ... }` for type re-exports

**File extensions**: always use `.js` extension in imports (ESM)

## Commit Conventions

Conventional commits with enforced scopes (commitlint + husky):

- `cli` - Go CLI changes
- `ts` - TypeScript packages
- `py` - Python packages (future)
- `deps` - dependency updates
- `release` - release commits

Format: `type(scope): message` (lowercase)

Examples: `feat(cli): add watch mode`, `fix(ts): handle null schemas`

## Testing Patterns

### Go Tests

- Use `t.TempDir()` for temp files (auto-cleaned)
- Table-driven tests: `tests := []struct{...}{}` with `t.Run(tt.name, ...)`
- Integration tests: prefix with `TestIntegration_`, skip with `if testing.Short() { t.Skip() }`
- Use `t.Errorf` for non-fatal assertions, `t.Fatalf` for fatal errors

### TypeScript

No test framework - validation via `bun run typecheck` and `bun run build`

## Architecture Notes

### Design Principle

**Commands orchestrate, packages execute.** The `cmd` package (generate, compliance) handles core logic and orchestration—calling packages in sequence, handling errors, coordinating data flow. Sub-packages are single-responsibility units that do one thing well and expose clean interfaces.

### Package Reference

Each package's role and how it connects to others:

| Package | Responsibility | Depends On | Called By |
|---------|---------------|------------|-----------|
| `cmd` | CLI entry points, orchestrates pipeline | all packages | `main.go` |
| `parser` | Finds config files, extracts declarations, detects language | `language`, `ui` | `cmd` |
| `retriever` | Fetches schemas from URL/file/inline with caching | `fetcher`, `ui` | `cmd` |
| `processor` | Validates, crawls refs, bundles schemas into self-contained units | `bundler`, `fetcher`, `validator`, `refextractor`, `metaschema`, `vocabulary` | `cmd`, `compliance` |
| `generator` | Calls adapter CLIs via stdin/stdout, groups by adapter | `adapter`, `language`, `ui` | `cmd` |
| `injector` | Writes output files using language templates | `adapter`, `language`, `ui` | `cmd` |
| `bundler` | Resolves `$ref`, flattens `$defs`, normalizes drafts | `fetcher`, `ui` | `processor` |
| `language` | Language config registry (TS, Python), templates, adapters | (none, leaf) | `parser`, `generator`, `injector`, `compliance` |
| `compliance` | Runs JSON Schema Test Suite against adapters | `adapter`, `fetcher`, `language`, `processor` | `cmd` |
| `ui` | Terminal output, colors, spinners, verbose logging | (external only) | all packages |
| `adapter` | Protocol types for adapter communication | (none, pure types) | `generator`, `injector`, `compliance` |
| `fetcher` | Fetcher interface, URI resolution, caching | (none, leaf) | `processor`, `bundler`, `retriever` |
| `validator` | Validates schemas against meta-schemas | `unsupported` | `processor` |
| `refextractor` | Discovers external `$ref` URIs for fetching | `fetcher`, `metaschema` | `processor` |
| `metaschema` | Fetches/caches meta-schemas, extracts `$vocabulary` | `fetcher` | `processor`, `bundler` |
| `vocabulary` | Filters schemas by enabled vocabularies | (none, leaf) | `processor` |
| `unsupported` | Defines keywords that can't be statically compiled | (none, embeds data) | `validator`, `compliance` |
| `config` | Loads .env files for header variable substitution | `ui` | `cmd` |

### Pipeline Flow

```
cmd/generate orchestrates:

  parse → retriever → processor → generator → injector
    │         │           │            │          │
    │         │           ├─ bundler   │          │
    │         │           ├─ validator │          │
    │         │           └─ refs...   │          │
    │         │                        │          │
    └─language                    language     language
```

1. **Parser**: finds config files with xschema.dev `$schema`, extracts declarations
2. **Retriever**: fetches raw schemas from URL/file/inline
3. **Processor**: validates, crawls external refs, bundles into self-contained schemas
4. **Generator**: calls adapter CLIs via stdin/stdout with schema batches
5. **Injector**: writes output files using language templates

### Bundler vs Adapters (responsibilities)

- **Bundler** blocks "wrong" schemas early (e.g. invalid/unresolvable refs, forbidden ref mechanisms), so adapters don't need to defend against broken inputs.
- **Adapters** decide what to do with the remaining valid bundled schemas based on what the target library can express idiomatically; adapter-specific gaps become documented adapter limitations or bugs to fix.

### Adapter Protocol

Adapters receive JSON array via stdin, output JSON array via stdout:

```typescript
// Input: [{ namespace: "user", id: "User", schema: {...} }]
// Output: [{ namespace: "user", id: "User", schema: "z.object(...)", type: "...", imports: [...] }]
```

### Key Types

- `parser.Declaration`: Namespace, ID, Adapter, ConfigPath, SourceType, Source
- `processor.ProcessedSchema`: Namespace, ID, bundled schema, adapter, sourceURI
- `adapter.ConvertResult`: Namespace, ID, Schema (code), Type (expression), Imports

## Common Gotchas

- Config file paths are relative to the config file's directory, not cwd
- Namespace defaults to filename without extension if not specified
- Same ID in same namespace across files = error
- Multiple languages without `--lang` flag = error
- `json.RawMessage` preserves raw JSON; don't re-marshal inline schemas
- Language detected from `$schema` URL: `typescript.jsonc` -> typescript, `python.jsonc` -> python
- Runner auto-detected from lockfiles: `bun.lock` -> bunx, `pnpm-lock.yaml` -> pnpm exec
- Always run `bun run build` from typescript/ dir before testing adapters
- Language registry is global; tests that register languages must call `language.ResetForTests()` (and `t.Cleanup(language.ResetForTests)`)
- **Compliance must be deterministic**: when iterating over Go maps, always sort keys first. Map iteration order is random by design. This affects error messages, report output, and test paths. See `unsupported/unsupported.go` for the pattern.

## Key File Locations

- CLI entry: `cli/main.go`
- Commands: `cli/cmd/generate.go`, `cli/cmd/compliance.go`
- Core IR types: `typescript/packages/core/src/ir/`
- Adapter example: `typescript/packages/adapters/zod/src/index.ts`
- Language configs: `cli/language/typescript.go`, `cli/language/python.go`
