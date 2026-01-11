# CLI Agent Guidelines

Go CLI for xschema - converts JSON Schema to native validators.

## Quality Requirements

Before submitting changes, ALL of these must pass:

```bash
go build -o xschema .                   # must compile without errors
go vet ./...                            # must pass with no warnings
go test ./...                           # must pass (all tests green)
```

**Always run these checks after making changes.** Do not submit code that fails any of these.

## Build/Test Commands

```bash
go test ./...                           # run all tests
go test ./parser/                       # run single package tests
go test ./parser/ -run TestParse        # run single test by name
go test ./parser/ -run TestParse -v     # verbose output
go test ./... -short                    # skip integration tests
go test . -run TestIntegration          # integration tests only (root dir)
go build -o xschema .                   # build binary
go vet ./...                            # lint/vet
go fmt ./...                            # format
```

## Package Structure

```
cmd/        # cobra commands (root, generate, compliance)
parser/     # parses JSON/JSONC config files, extracts declarations
retriever/  # fetches schemas from URL/file/inline
generator/  # calls adapter CLIs via stdin/stdout
injector/   # writes generated code to output files
language/   # language configs (typescript)
bundler/    # bundles multiple schemas
compliance/ # adapter compliance testing harness
ui/         # terminal output helpers (colors, verbose logging)
```

## Code Style

### Language Specific Code

Language specific code should NEVER be written inside the cli logic, but it should be embedded into the language struct (language.go)

### Imports

Stdlib first, external second, internal third (blank lines between):

```go
import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/spf13/cobra"
    "github.com/tailscale/hujson"

    "github.com/xschemadev/xschema/language"
    "github.com/xschemadev/xschema/ui"
)
```

### Naming

- Exported: `PascalCase` (e.g., `ParseResult`, `Declaration`)
- Unexported: `camelCase` (e.g., `parseConfigFile`, `mergeDeclarations`)
- Acronyms uppercase: `URL`, `ID`, `JSON`, `HTTP`
- Constants: `camelCase` for unexported, `PascalCase` for exported

### Error Handling

- Wrap errors with context: `fmt.Errorf("failed to parse config: %w", err)`
- Return early on errors
- Use `t.Fatalf` for fatal test errors, `t.Errorf` for non-fatal

```go
config, err := parseConfigFile(path)
if err != nil {
    return nil, fmt.Errorf("failed to parse %s: %w", path, err)
}
```

### Context

Accept `context.Context` as first param for cancellable operations:

```go
func Parse(ctx context.Context, projectRoot string, langFilter string) (*ParseResult, error) {
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
    }
    // ...
}
```

### Logging

Use `ui.Verbosef()` for debug output (shown with `--verbose` flag):

```go
ui.Verbosef("parsing project: root=%s, langFilter=%s", projectRoot, langFilter)
ui.Verbosef("found potential config files: count=%d", len(files))
```

### Types

- Define close to usage
- Use `json.RawMessage` for arbitrary JSON (preserves raw bytes)
- Prefer structs over maps for known shapes
- Use type aliases sparingly

## Testing Patterns

### Table-Driven Tests

```go
func TestParseConfigFileMalformedJSON(t *testing.T) {
    tests := []struct {
        name    string
        content string
    }{
        {"truncated", `{"$schema": "https://xschema.dev/..."`},
        {"invalid syntax", `{"$schema": "...", schemas: []}`},
        {"empty file", ``},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test logic
        })
    }
}
```

### Temp Files

Use `t.TempDir()` - auto-cleaned after test:

```go
func TestParse(t *testing.T) {
    tmpDir := t.TempDir()
    configPath := filepath.Join(tmpDir, "test.jsonc")
    if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
        t.Fatalf("failed to write config: %v", err)
    }
    // ...
}
```

### Integration Tests

Prefix with `TestIntegration_`, place in root dir:

```go
func TestIntegration_MultipleConfigs(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    }
    // ...
}
```

## Key Types

- `parser.Declaration`: Namespace, ID, Adapter, ConfigPath, SourceType, Source
- `parser.ParseResult`: Language, Configs, Declarations
- `retriever.SchemaWithMeta`: extends Declaration with resolved Schema bytes
- `generator.GenerateOutput`: Namespace, ID, Schema (code), Type, Imports

## Common Patterns

### JSON Schema Detection

Check `$schema` URL for xschema.dev prefix:

```go
if !language.IsXSchemaURL(raw.Schema) {
    return nil, nil  // not an xschema config
}
```

### Namespace Derivation

Defaults to filename without extension:

```go
namespace := raw.Namespace
if namespace == "" {
    base := filepath.Base(path)
    namespace = strings.TrimSuffix(base, filepath.Ext(base))
}
```

### Grouping by Adapter

```go
byAdapter := result.DeclarationsByAdapter()
for adapter, decls := range byAdapter {
    // process each adapter's schemas
}
```

## Gotchas

- `json.RawMessage` must not be re-marshaled for inline schemas
- File paths in configs are relative to config file dir, not cwd
- Multiple languages require `--lang` flag
- Built-in languages register via `github.com/xschemadev/xschema/language/langs`; tests that call `language.ResetForTests()` should call `langs.RegisterBuiltins()` to restore built-ins
- Context cancellation should be checked in loops
