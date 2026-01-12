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

### Pipeline Flow

1. **Parser**: finds JSON/JSONC files with xschema.dev `$schema`, extracts declarations
2. **Retriever**: fetches schemas from URL/file or passes inline JSON through
3. **Generator**: calls adapter CLIs via stdin/stdout with schema batches
4. **Injector**: writes generated code using language templates

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
- `generator.GenerateOutput`: Namespace, ID, Schema (code), Type (expression), Imports

## Common Gotchas

- Config file paths are relative to the config file's directory, not cwd
- Namespace defaults to filename without extension if not specified
- Same ID in same namespace across files = error
- Multiple languages without `--lang` flag = error
- `json.RawMessage` preserves raw JSON; don't re-marshal inline schemas
- Language detected from `$schema` URL: `ts.jsonc` -> ts, `py.jsonc` -> py
- Runner auto-detected from lockfiles: `bun.lock` -> bunx, `pnpm-lock.yaml` -> pnpm exec
- Always run `bun run build` from typescript/ dir before testing adapters
- Language registry is global; tests that register languages must call `language.ResetForTests()` (and `t.Cleanup(language.ResetForTests)`)

## Key File Locations

- CLI entry: `cli/main.go`
- Commands: `cli/cmd/generate.go`, `cli/cmd/compliance.go`
- Core IR types: `typescript/packages/core/src/ir/`
- Adapter example: `typescript/packages/adapters/zod/src/index.ts`
- Language configs: `cli/language/typescript.go`, `cli/language/python.go`
