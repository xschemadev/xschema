# Language Package

The language package defines language-specific configurations and provides a registry for looking up languages by name or `$schema` URL.

## Architecture

```
language/
├── registry.go    # Language registry (Register, ByName, BySchemaURL)
├── language.go    # Language struct + registry wrappers
├── spec.go        # Core types (GeneratedFile, CommandSpec, AdapterInvoker)
├── paths.go       # Path normalization + validation
├── order.go       # Deterministic file ordering
├── imports.go     # Shared import helpers
├── templates.go   # Shared template helpers
└── langs/         # Per-language implementations
    ├── langs.go         # RegisterBuiltins()
    └── typescript/      # TypeScript language
        └── typescript.go
```

## Language Registry

Languages are registered globally via `language.Register()` and looked up by:
- **Name**: `language.ByName("typescript")`
- **Schema URL**: `language.BySchemaURL("https://xschema.dev/schemas/typescript.jsonc")`

```go
// Register a language (usually via init() in the language package)
err := language.Register(typescript.Language())

// Look up by name or schema URL
lang := language.ByName("typescript")
lang := language.BySchemaURL("https://xschema.dev/schemas/typescript.jsonc")

// List supported languages
names := language.SupportedLanguages() // []string{"typescript"}

// Get all ignore dirs (for directory walking)
dirs := language.AllIgnoreDirs() // map[string]bool{"node_modules": true, ...}

// Reset for tests (prevents global state leakage)
language.ResetForTests()
```

## Language Struct

```go
type Language struct {
    Name             string
    Extensions       []string  // source file extensions (for injector)
    SchemaURL        string    // e.g., "https://xschema.dev/schemas/typescript.jsonc"
    SchemaExt        string    // e.g., "typescript.jsonc" (derived from SchemaURL)
    AdapterBinPrefix string    // e.g., "xschema-" - prefix for adapter binaries
    DetectRunner     func() (cmd string, args []string, err error)
    AdapterInvoker   AdapterInvoker  // resolves adapter refs to CommandSpec

    // Output generation
    OutputFile   string
    Template     string
    MergeImports func([]string) string
    BuildHeader  func(outDir string, schemas []SchemaEntry) string
    BuildFooter  func(outDir string, schemas []SchemaEntry) string
    BuildVarName func(namespace, id string) string

    // Parser (for directory walking)
    IgnoreDirs []string  // e.g., ["node_modules", "dist"]

    // Compliance testing
    DetectHarnessRunner func(dir string) (cmd, args, err)
    GetPackageName      func(dir string) string
    AdapterCLIPath      func(adapterPath string) string
    HarnessExtension    string
    HarnessTemplate     string
}
```

## Core Types (spec.go)

### GeneratedFile

Represents one file emitted during generation:

```go
type GeneratedFile struct {
    Path     string  // relative to output root (forward slashes)
    Contents string
}
```

### CommandSpec

A fully-resolved command invocation:

```go
type CommandSpec struct {
    Cmd  string
    Args []string
    Dir  string   // working directory
    Env  []string // optional environment variables
}
```

### AdapterInvoker

Interface for resolving adapter refs to runnable commands:

```go
type AdapterInvoker interface {
    BuildAdapterCommand(ctx context.Context, input AdapterCommandInput) (CommandSpec, error)
}

type AdapterCommandInput struct {
    ProjectRoot string
    AdapterRef  string  // e.g., "@xschemadev/zod"
}
```

## Adding a New Language

1. **Create the language package** in `language/langs/<name>/`:

```go
package mylang

import "github.com/xschemadev/xschema/language"

func init() {
    if err := language.Register(Language()); err != nil {
        panic(err)
    }
}

func Language() language.Language {
    return language.Language{
        Name:             "mylang",
        Extensions:       []string{".ml"},
        SchemaURL:        language.XSchemaBaseURL + "mylang.jsonc",
        AdapterBinPrefix: "xschema-",
        AdapterInvoker:   invoker{},
        // ... other fields
    }
}

type invoker struct{}

func (invoker) BuildAdapterCommand(ctx context.Context, input language.AdapterCommandInput) (language.CommandSpec, error) {
    // Validate adapter ref, detect runner, return CommandSpec
}
```

2. **Register the language** in `language/langs/langs.go`:

```go
import (
    "github.com/xschemadev/xschema/language"
    "github.com/xschemadev/xschema/language/langs/mylang"
    "github.com/xschemadev/xschema/language/langs/typescript"
)

func RegisterBuiltins() error {
    if err := language.Register(typescript.Language()); err != nil {
        return err
    }
    if err := language.Register(mylang.Language()); err != nil {
        return err
    }
    return nil
}
```

3. **Import `langs` from main** (already done in `cli/main.go`):

```go
import _ "github.com/xschemadev/xschema/language/langs"
```

4. **Add web schema** in `web/public/schemas/mylang.jsonc`

## Path Helpers

Safe relative path handling for multi-file outputs:

```go
// Normalize path to forward slashes, reject unsafe paths
path, err := language.NormalizeRelativePath("foo/bar.ts")
// Returns: "foo/bar.ts", nil
// Rejects: absolute paths, ".." segments

// Sort files deterministically by path
language.SortGeneratedFiles(files)

// Validate no duplicates after normalization
err := language.ValidateGeneratedFiles(files)
```

## Testing

Tests that modify global registry state must use `ResetForTests()`:

```go
func TestSomething(t *testing.T) {
    language.ResetForTests()
    t.Cleanup(language.ResetForTests)

    // For tests needing built-in languages:
    langs.RegisterBuiltins()

    // Register test-only language
    err := language.Register(language.Language{
        Name:      "fake",
        SchemaURL: language.XSchemaBaseURL + "fake.jsonc",
    })
}
```
