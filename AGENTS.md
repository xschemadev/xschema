# xschema Agent Guidelines

JSON Schema to native validators (Zod, Pydantic, etc.) with full type safety.

## Project Structure

```
cli/                    # Go CLI (main tool)
  cmd/                  # cobra commands
  parser/               # parses JSON/JSONC config files
  retriever/            # fetches schemas from URL/file/inline
  generator/            # calls adapters to convert schemas
  injector/             # writes generated code
  language/             # language-specific config (TS, Python)
  ui/                   # terminal output helpers

packages/typescript/    # TS packages (bun workspace)
  core/                 # shared types (ConvertInput, ConvertResult)
  client/               # runtime client
  adapters/zod/         # @xschema/zod adapter
  example/              # example project

packages/python/        # Python packages (future)
```

## Build/Test/Lint Commands

### Go CLI

```bash
# run all tests
cd cli && go test ./...

# run single test file
cd cli && go test ./parser/

# run single test by name
cd cli && go test ./parser/ -run TestParseConfigFile

# run with verbose output
cd cli && go test ./parser/ -v

# run short tests only (skip integration)
cd cli && go test ./... -short

# run integration tests
cd cli && go test . -run TestIntegration

# build CLI
cd cli && go build -o xschema .

# format code
cd cli && go fmt ./...

# lint (if golangci-lint installed)
cd cli && golangci-lint run
```

### TypeScript Packages

```bash
# install deps (from packages/typescript)
cd packages/typescript && bun install

# build all packages
cd packages/typescript && bun run build

# run example
cd packages/typescript/example && bun run main.ts

# type check (in any package)
bunx tsc --noEmit
```

## Code Style Guidelines

### Go Code

**Imports**: stdlib first, then external, then internal. Grouped with blank lines:
```go
import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "github.com/xschema/cli/language"
    "github.com/xschema/cli/ui"
)
```

**Naming**:
- Exported: `PascalCase` (e.g., `Parse`, `GenerateOutput`)
- Unexported: `camelCase` (e.g., `parseConfigFile`)
- Constants: `camelCase` or `PascalCase` based on export
- Acronyms: keep uppercase (`URL`, `HTTP`, `ID`)

**Error handling**:
- Wrap errors with context: `fmt.Errorf("failed to X: %w", err)`
- Return early on errors
- Use `ui.Verbosef()` for debug logging before returning errors

**Functions**:
- Accept `context.Context` as first param for cancellable ops
- Keep functions focused and small
- Use named return values sparingly (only when clarifying)

**Types**:
- Define types close to where they're used
- Use `json.RawMessage` for arbitrary JSON
- Prefer structs over maps for known shapes

### TypeScript Code

**Imports**: external packages first, then relative:
```typescript
import type { ConvertInput, ConvertResult } from "@xschema/core";
import { jsonSchemaToZod } from "json-schema-to-zod";

import { convert } from "./index";
```

**Naming**:
- Interfaces/Types: `PascalCase` (e.g., `XSchemaAdapter`, `ConvertResult`)
- Functions: `camelCase` (e.g., `convert`, `createAdapterCLI`)
- Constants: `camelCase` or `SCREAMING_SNAKE` for true constants

**Types**:
- Use `interface` for object shapes
- Use `type` for unions, intersections, aliases
- Mark readonly when appropriate: `readonly __brand: 'xschema-adapter'`
- Use `object` (not `Object`) in type annotations

**Exports**:
- Named exports preferred over default exports
- Re-export types with `export type { ... }` when possible

## Architecture Notes

### Config File Format

xschema configs are JSON/JSONC files with `$schema` pointing to xschema.dev:
```jsonc
{
  "$schema": "https://xschema.dev/schemas/ts.jsonc",
  "namespace": "api",  // optional, defaults to filename
  "schemas": [
    {
      "id": "User",
      "sourceType": "url",     // "url" | "file" | "json"
      "source": "https://...", // or "./path.json" or {...}
      "adapter": "@xschema/zod"
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
// Input
[{ namespace: "user", id: "User", schema: {...} }]

// Output
[{ namespace: "user", id: "User", schema: "z.object(...)", type: "z.infer<typeof user_User>", imports: ["import { z } from 'zod'"] }]
```

### Key Types

```go
// parser/types.go
type Declaration struct {
    Namespace  string
    ID         string
    SourceType SourceType // "url" | "file" | "json"
    Source     json.RawMessage
    Adapter    string
    ConfigPath string
}

// generator/generator.go
type GenerateOutput struct {
    Namespace string
    ID        string
    Schema    string   // generated code
    Type      string   // type expression
    Imports   []string
}
```

### Language Detection

- Detected from `$schema` URL: `https://xschema.dev/schemas/ts.jsonc` -> typescript
- Runner auto-detected from lockfiles (bun.lock -> bunx, pnpm-lock.yaml -> pnpm exec, etc.)

## Testing Patterns

**Unit tests**: create temp directories with test configs
```go
func TestParseConfigFile(t *testing.T) {
    tmpDir := t.TempDir()
    // write test files, run parser, assert
}
```

**Integration tests**: test full pipeline (Parse -> Retrieve -> Generate -> Inject)
```go
func TestIntegration_FullPipeline(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test in short mode")
    }
    // requires adapter to be installed
}
```

**Table-driven tests**: for testing multiple cases
```go
tests := []struct {
    name    string
    content string
}{
    {"truncated", `{...`},
    {"empty", ``},
}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) { ... })
}
```

## Common Gotchas

- Config file paths are relative to the config file's directory, not cwd
- Namespace defaults to filename without extension if not specified
- Same ID in same namespace across files = error
- Multiple languages in project without `--lang` flag = error
- `json.RawMessage` preserves raw JSON; don't re-marshal inline schemas
