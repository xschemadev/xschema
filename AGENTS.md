# xschema Agent Guidelines

JSON Schema to native validators (Zod, Pydantic, etc.) with full type safety.

## Project Structure

```
cli/                    # Go CLI (github.com/xschemadev/xschema)
  cmd/                  # cobra commands
  parser/               # parses JSON/JSONC config files
  retriever/            # fetches schemas from URL/file/inline
  generator/            # calls adapters to convert schemas
  injector/             # writes generated code
  language/             # language-specific config (TS, Python)
  ui/                   # terminal output helpers

typescript/             # TS packages (bun workspace)
  packages/core/        # @xschemadev/core - shared types
  packages/zod/         # @xschemadev/zod - zod adapter
  packages/client/      # @xschemadev/client - runtime client
  example/              # example project
```

## Build/Test/Lint Commands

### Go CLI (run from `cli/` directory)

```bash
go test ./...                           # run all tests
go test ./parser/                       # run single package tests
go test ./parser/ -run TestParse        # run single test by name
go test ./parser/ -v                    # verbose output
go test ./... -short                    # skip integration tests
go test . -run TestIntegration          # run integration tests only
go build -o xschema .                   # build binary
go fmt ./...                            # format code
golangci-lint run                       # lint (if installed)
```

### TypeScript (run from `typescript/` directory)

```bash
bun install                             # install deps
bun run build                           # build all packages
bun run typecheck                       # type check all packages
bunx tsc --noEmit                       # type check single package
```

## Code Style Guidelines

### Go Code

**Imports**: stdlib first, external second, internal third (grouped with blank lines):
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

**Error handling**: wrap with context `fmt.Errorf("failed to X: %w", err)`, return early, use `ui.Verbosef()` for debug logging

**Functions**: accept `context.Context` as first param for cancellable ops

**Types**: define close to usage, use `json.RawMessage` for arbitrary JSON, prefer structs over maps

### TypeScript Code

**Imports**: external packages first, then relative (blank line between):
```typescript
import type { ConvertInput, ConvertResult } from "@xschemadev/core";
import { jsonSchemaToZod } from "json-schema-to-zod";

import { convert } from "./index";
```

**Naming**: types/interfaces `PascalCase`, functions `camelCase`

**Types**: use `interface` for object shapes, `type` for unions/aliases, use `object` not `Object`

**Exports**: named exports preferred, re-export types with `export type { ... }`

## Commit Conventions

Uses conventional commits with these scopes:
- `cli` - Go CLI changes
- `ts` - TypeScript packages
- `py` - Python packages (future)
- `deps` - dependency updates
- `release` - release commits

Examples: `feat(cli): add watch mode`, `fix(ts): handle null schemas`

## Architecture Notes

### Config File Format

```jsonc
{
  "$schema": "https://xschema.dev/schemas/ts.jsonc",
  "namespace": "api",  // optional, defaults to filename
  "schemas": [
    {
      "id": "User",
      "sourceType": "url",     // "url" | "file" | "json"
      "source": "https://...", // or "./path.json" or {...}
      "adapter": "@xschemadev/zod"
    }
  ]
}
```

### Pipeline Flow

1. **Parser**: finds JSON/JSONC files with xschema.dev `$schema`, extracts declarations
2. **Retriever**: fetches schemas from URL/file or passes inline JSON through
3. **Generator**: calls adapter CLIs via stdin/stdout with schema batches
4. **Injector**: writes generated code using language templates

### Adapter Protocol

Adapters receive JSON array via stdin, output JSON array via stdout:
```typescript
// Input: [{ namespace: "user", id: "User", schema: {...} }]
// Output: [{ namespace: "user", id: "User", schema: "z.object(...)", type: "...", imports: [...] }]
```

### Key Types

- `parser.Declaration`: Namespace, ID, Adapter, ConfigPath, SourceType, Source (json.RawMessage)
- `generator.GenerateOutput`: Namespace, ID, Schema (code), Type (expression), Imports

## Testing Patterns

- Use `t.TempDir()` for temp files
- Table-driven tests: `tests := []struct{...}{}` with `t.Run(tt.name, ...)`
- Integration tests skip with `if testing.Short() { t.Skip() }`

## Common Gotchas

- Config file paths are relative to the config file's directory, not cwd
- Namespace defaults to filename without extension if not specified
- Same ID in same namespace across files = error
- Multiple languages without `--lang` flag = error
- `json.RawMessage` preserves raw JSON; don't re-marshal inline schemas
- Language detected from `$schema` URL: `ts.jsonc` -> typescript, `py.jsonc` -> python
- Runner auto-detected from lockfiles (bun.lock -> bunx, pnpm-lock.yaml -> pnpm exec)
