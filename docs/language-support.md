# Language Support

xschema supports multiple languages through a language-agnostic core pipeline and language-specific configuration. This guide covers how language support works and how to add a new language.

## Language Detection

Language detection happens at the config file parsing stage and is based on the `$schema` URL.

### How Language Detection Works

**Location:** `cli/language/language.go`, `cli/parser/parser.go`

Each xschema config file declares its language via the `$schema` field:

```jsonc
{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  "schemas": [...]
}
```

**Detection Flow:**

1. **Parse config file** as JSON/JSONC (using `hujson` library)
2. **Extract `$schema` URL** from root object
3. **Check if URL** starts with `https://xschema.dev/schemas/`
4. **Extract schema extension** from URL (e.g., `ts.jsonc`, `py.jsonc`)
5. **Map extension to language** via lookup table

**Code:**

```go
// cli/language/language.go
const XSchemaBaseURL = "https://xschema.dev/schemas/"

func BySchemaURL(url string) *Language {
    if !strings.HasPrefix(url, XSchemaBaseURL) {
        return nil  // Not an xschema config
    }
    ext := strings.TrimPrefix(url, XSchemaBaseURL)  // "ts.jsonc"
    return languageBySchemaExt[ext]  // Lookup table
}

func IsXSchemaURL(url string) bool {
    return strings.HasPrefix(url, XSchemaBaseURL)
}
```

### Multi-Language Projects

If a project has config files for multiple languages (e.g., both TypeScript and Python), the CLI will reject it unless `--lang` flag is provided:

```bash
# Error: multiple languages detected (typescript, python). Use --lang to specify which one
xschema generate

# OK: Filter to TypeScript only
xschema generate --lang typescript

# OK: Filter to Python only
xschema generate --lang python
```

## Language-Agnostic vs Language-Specific

### Language-Agnostic (Core Pipeline)

The core xschema pipeline is fully language-agnostic:

1. **Parser** (`cli/parser/`):
   - Finds JSON/JSONC files
   - Validates `$schema` format
   - Detects language from URL (but doesn't require specific language)
   - Extracts schema declarations

2. **Retriever** (`cli/retriever/`):
   - Fetches schemas from URLs, files, or inline JSON
   - Returns raw JSON Schema objects
   - No language-specific processing

3. **Generator** (`cli/generator/`):
   - Calls adapter binaries via stdin/stdout
   - Receives JSON input/output
   - No language-specific logic

### Language-Specific Configuration

All language-specific behavior is encapsulated in the `Language` struct:

**Location:** `cli/language/language.go`

```go
type Language struct {
    // Basic info
    Name             string     // "typescript", "python"
    SchemaURL        string     // Full $schema URL
    SchemaExt        string     // Just the extension ("ts.jsonc")
    AdapterBinPrefix string     // Prefix for adapter binaries ("xschema-")
    
    // Runner detection
    DetectRunner     func() (cmd string, args []string, err error)
    
    // Output generation
    OutputFile       string     // e.g., "xschema.gen.ts", "__init__.py"
    Template         string     // Go text/template for output file
    MergeImports     func([]string) string
    BuildVarName     func(namespace, id string) string
    BuildHeader      func(outDir string, schemas []SchemaEntry) string
    BuildFooter      func(outDir string, schemas []SchemaEntry) string
    
    // Client injection
    BuildSchemasImport   func(importPath string) string
    ImportPattern        string
    InjectSchemasKey     func(configContent string) string
    ClientFactoryPattern string
    
    // Parser fallback
    IgnoreDirs       []string   // Directories to skip during walk
}
```

### Current Language Implementations

#### TypeScript

```go
Language{
    Name:                 "typescript",
    SchemaURL:            "https://xschema.dev/schemas/typescript.jsonc",
    SchemaExt:            "ts.jsonc",
    AdapterBinPrefix:     "xschema-",
    DetectRunner:         detectTSRunner,
    OutputFile:           "xschema.gen.ts",
    Template:             TSTemplate,
    MergeImports:         MergeTSImports,
    BuildVarName:         buildVarNameUnderscore,
    BuildSchemasImport:   buildTSSchemasImport,
    ImportPattern:        `(?m)^import\s+.*$`,
    InjectSchemasKey:     injectSchemasKeyBrace,
    ClientFactoryPattern: `createXSchemaClient\s*\(\s*(\{[^}]*\})\s*\)`,
    IgnoreDirs:           []string{"node_modules", "dist", "build"},
}
```

#### Python

```go
Language{
    Name:                 "python",
    SchemaURL:            "https://xschema.dev/schemas/py.jsonc",
    SchemaExt:            "py.jsonc",
    DetectRunner:         detectPythonRunner,
    OutputFile:           "__init__.py",
    Template:             PyTemplate,
    MergeImports:         MergePyImports,
    BuildVarName:         buildVarNameUnderscore,
    BuildSchemasImport:   buildPySchemasImport,
    ImportPattern:        `(?m)^(?:import\s+|from\s+).*$`,
    InjectSchemasKey:     injectSchemasKeyBrace,
    ClientFactoryPattern: `create_xschema_client\s*\(\s*(\{[^}]*\})\s*\)`,
    BuildFooter:          BuildPythonFooter,
    IgnoreDirs:           []string{"__pycache__", ".venv", "venv"},
}
```

## Runner Detection

The xschema CLI needs to know how to execute adapters. Different languages/package managers require different runners.

### TypeScript Runner Detection

**Location:** `cli/language/language.go` - `detectTSRunner()`

Detection order:

1. **Check package.json** for `packageManager` field:
   ```json
   { "packageManager": "bun@1.0.0" }
   ```

2. **Check for lockfiles** in order:
   - `bun.lock` or `bun.lockb` → `bunx`
   - `pnpm-lock.yaml` → `pnpm exec`
   - `yarn.lock` → `yarn`
   - `package-lock.json` → `npx`

3. **Check system PATH** for available runners:
   - Try: `bunx`, `pnpm`, `yarn`, `npx` in that order

4. **Default to npx**

**Code:**

```go
func detectTSRunner() (string, []string, error) {
    // 1. Try package.json packageManager field
    if content, _ := os.ReadFile("package.json"); content != nil {
        pm := detectPackageManager(string(content))
        if pm == "bun" { return "bunx", nil, nil }
        if pm == "pnpm" { return "pnpm", []string{"exec"}, nil }
        // ...
    }
    
    // 2. Check lockfiles
    lockfileCmds := map[string][]string{
        "bun.lock":          {"bunx"},
        "pnpm-lock.yaml":    {"pnpm", "exec"},
        "yarn.lock":         {"yarn"},
        "package-lock.json": {"npx"},
    }
    
    for lf, cmd := range lockfileCmds {
        if _, err := os.Stat(lf); err == nil {
            return cmd[0], cmd[1:], nil
        }
    }
    
    // 3. Check PATH
    // 4. Default
    return "npx", nil, nil
}
```

### Python Runner Detection

**Location:** `cli/language/language.go` - `detectPythonRunner()`

Detection order:

1. **Check lockfiles** in order:
   - `uv.lock` → `uv run`
   - `poetry.lock` → `poetry run`
   - `Pipfile` → `pipenv run`

2. **Check pyproject.toml** for build system:
   - `build-backend` mentions `uv` → `uv run`
   - mentions `poetry-core` → `poetry run`

3. **Default to python -m**

**Code:**

```go
func detectPythonRunner() (string, []string, error) {
    lockfileCmds := map[string][]string{
        "uv.lock":     {"uv", "run"},
        "poetry.lock": {"poetry", "run"},
        "Pipfile":     {"pipenv", "run"},
    }
    
    for lf, cmd := range lockfileCmds {
        if _, err := os.Stat(lf); err == nil {
            return cmd[0], cmd[1:], nil
        }
    }
    
    // Check pyproject.toml
    // ...
    
    return "python", []string{"-m"}, nil
}
```

## Config File Parsing

### Configuration File Detection

The parser finds all JSON and JSONC files in the project using:

1. **Git (preferred):** `git ls-files *.json *.jsonc`
   - Respects `.gitignore`
   - Fast for large projects

2. **Directory walk (fallback):**
   - Manual filesystem traversal
   - Skips language-specific directories (`node_modules`, `__pycache__`, etc.)

### Configuration File Structure

Both TypeScript and Python configs follow the same schema:

```jsonc
{
  "$schema": "https://xschema.dev/schemas/{ts|py}.jsonc",
  "namespace": "optional-override",  // Defaults to filename
  "schemas": [
    {
      "id": "SchemaName",
      "sourceType": "url" | "file" | "json",
      "source": "...",  // URL string, file path, or inline object
      "adapter": "adapter-name"
    }
  ]
}
```

### Validation Rules

1. **Namespace defaults to filename** (without extension)
   - `user.jsonc` → namespace `"user"`
   - Can be overridden with `"namespace"` field

2. **Same namespace from different files is merged**
   - `user.jsonc` and `user-v2.jsonc` are merged into one namespace

3. **Duplicate schema IDs are an error**
   - Two schemas with `"id": "User"` in the same namespace is rejected

4. **All configs in a project must be same language**
   - Error if some files have `ts.jsonc` and others have `py.jsonc`
   - Use `--lang` flag to filter to one language

## Adding a New Language

### Directory Convention

All languages follow the same directory structure:

```
{lang}/packages/adapters/{adapter}/
```

Where:
- `{lang}` = language name (must match `Language.Name` in Go, e.g., `typescript`, `python`, `rust`)
- `{adapter}` = adapter name (e.g., `zod`, `pydantic`, `serde`)

Compliance harnesses are generated from Go templates in `cli/compliance` at runtime. Each adapter stores compliance output under `compliance/results/`, which is created by the CLI when running `xschema compliance --dev-report`.

This convention enables automatic discovery by both the CLI and CI workflows.

### Requirements

To add support for a new language (e.g., Go, Rust, C#), implement:

#### 1. Language Configuration

Add entry to `Languages` slice in `cli/language/language.go`:

```go
{
    Name:             "rust",
    SchemaURL:        "https://xschema.dev/schemas/rs.jsonc",
    SchemaExt:        "rs.jsonc",
    AdapterBinPrefix: "xschema-",
    DetectRunner:     detectRustRunner,
    OutputFile:       "xschema.gen.rs",
    Template:         RustTemplate,
    MergeImports:     MergeRustImports,
    BuildVarName:     buildVarNameUpperCase,
    IgnoreDirs:       []string{"target"},
}
```

#### 2. Runner Detection Function

```go
func detectRustRunner() (string, []string, error) {
    // Check for Cargo.toml
    if _, err := os.Stat("Cargo.toml"); err == nil {
        return "cargo", []string{"run", "--bin"}, nil
    }
    return "", nil, fmt.Errorf("no Rust project found")
}
```

#### 3. Output Template

Add Go text/template string in `cli/language/templates.go`:

```go
const RustTemplate = `// Generated by xschema - DO NOT EDIT
// https://xschema.dev/docs
{{.Imports}}
{{range .Schemas}}
pub const {{.VarName}}: ... = {{.Code}};
{{end}}
`
```

**Template Data Structure:**

```go
type TemplateData struct {
    Imports string                 // Merged import statements
    Schemas []language.SchemaEntry // Schema entries with code
    Header  string                 // Language-specific header
    Footer  string                 // Language-specific footer
}
```

**Available Properties:**
```go
type SchemaEntry struct {
    Namespace string  // "user"
    ID        string  // "Profile"
    VarName   string  // "user_Profile"
    Code      string  // Generated validator code
    Type      string  // Type expression
}
```

#### 4. Import Merging Function

```go
func MergeRustImports(imports []string) string {
    // Dedup and format Rust use statements
    // e.g., "use serde::{Deserialize, Serialize};"
    // Input: []string{"use serde::Deserialize;", "use serde::Serialize;"}
    // Output: "use serde::{Deserialize, Serialize};"
}
```

#### 5. Variable Name Builder (if needed)

```go
func buildVarNameUpperCase(namespace, id string) string {
    // Build Rust CONSTANT_CASE: "USER_PROFILE"
    return strings.ToUpper(namespace + "_" + id)
}
```

#### 6. Create Adapter Package

Create an adapter package following the directory convention:

```
rust/packages/adapters/serde/
├── src/
│   └── ...              # Language-specific implementation
├── compliance/
│   └── results/         # Compliance reports (generated)
└── Cargo.toml           # Or appropriate manifest
```

The adapter must:
- Accept JSON Schema via stdin
- Output generated code via stdout
- Follow the adapter protocol (see [adapter-system.md](adapter-system.md))

#### 7. Register in build system

If using release-please, add to `release-please-config.json`:

```json
{
  "packages": {
    "rust/packages/adapters/serde": {
      "release-type": "simple",
      "component": "serde"
    }
  }
}
```

#### 8. Update compliance workflow

Add setup and build steps for the new language in `.github/workflows/compliance.yml`:

```yaml
- uses: actions-rust-lang/setup-rust-toolchain@v1
  if: matrix.lang == 'rust'

- name: Install Rust dependencies
  if: matrix.lang == 'rust'
  run: cargo fetch
  working-directory: rust

- name: Build Rust packages
  if: matrix.lang == 'rust'
  run: cargo build --release
  working-directory: rust
```

The CI workflow automatically discovers adapters by scanning adapter directories under `*/packages/adapters/*` and checking for a language-specific package manifest (for example `package.json`, `pyproject.toml`, or `Cargo.toml`). The path triggers use `*/packages/adapters/**` which covers all languages.

### Example: Full Rust Support

**Step 1: Config (cli/language/rust.go)**

```go
var rust = Language{
    Name:                "rust",
    Extensions:          []string{".rs"},
    SchemaURL:           XSchemaBaseURL + "rs.jsonc",
    SchemaExt:           "rs.jsonc",
    AdapterBinPrefix:    "xschema-",
    DetectRunner:        detectRustRunner,
    OutputFile:          "xschema.gen.rs",
    Template:            RustTemplate,
    MergeImports:        MergeRustImports,
    BuildVarName:        buildVarNameUpperCase,
    IgnoreDirs:          []string{"target", ".cargo"},
    DetectHarnessRunner: detectRustHarnessRunner,
    GetPackageName:      getRustPackageName,
    AdapterCLIPath:      getRustAdapterCLIPath,
}
```

Don't forget to add `rust` to the `Languages` slice in `language.go`.

**Step 2: Runner detection (cli/language/rust.go)**

```go
func detectRustRunner() (string, []string, error) {
    if _, err := os.Stat("Cargo.toml"); err == nil {
        return "cargo", []string{"run", "--bin"}, nil
    }
    return "", nil, fmt.Errorf("Cargo.toml not found")
}

func detectRustHarnessRunner(dir string) (string, []string, error) {
    return "cargo", []string{"run", "--quiet", "--"}, nil
}

func getRustPackageName(dir string) string {
    // Parse Cargo.toml to get package name
    return filepath.Base(dir)
}

func getRustAdapterCLIPath(adapterPath string) string {
    return filepath.Join(adapterPath, "target", "release", "cli")
}
```

**Step 3: Template (cli/language/templates.go)**

```go
const RustTemplate = `// Generated by xschema - DO NOT EDIT
// https://xschema.dev/docs
{{.Imports}}

{{range .Schemas}}
pub mod {{.VarName}} {
    {{.Code}}
}
{{end}}

pub mod schemas {
    use super::*;
    {{range .Schemas}}
    pub use {{.VarName}}::*;
    {{end}}
}
`
```

**Step 4: Import merging (cli/language/imports.go)**

```go
func MergeRustImports(imports []string) string {
    seen := make(map[string]bool)
    var result []string
    for _, imp := range imports {
        if !seen[imp] {
            result = append(result, imp)
            seen[imp] = true
        }
    }
    return strings.Join(result, "\n")
}
```

**Step 5: Create adapter directory structure**

Following the convention `{lang}/packages/adapters/{adapter}/`:

```bash
mkdir -p rust/packages/adapters/serde/src
mkdir -p rust/packages/adapters/serde/compliance
```

**Step 6: Add compliance harness template**

Define a Go text/template in `cli/compliance/harness_template.go` and wire it in `GetHarnessTemplate`/`GetHarnessExtension`. The CLI now generates harness files on the fly, so no per-adapter `compliance/harness.{ext}` file is required.

**Step 7: Add CI setup steps**

Add to `.github/workflows/compliance.yml`:

```yaml
- uses: actions-rust-lang/setup-rust-toolchain@v1
  if: matrix.lang == 'rust'

- name: Build Rust packages
  if: matrix.lang == 'rust'
  run: cargo build --release
  working-directory: rust
```

## Configuration File Examples

### TypeScript Example

```jsonc
{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  "schemas": [
    {
      "id": "User",
      "sourceType": "url",
      "source": "https://example.com/schemas/user.json",
      "adapter": "zod"
    },
    {
      "id": "Config",
      "sourceType": "file",
      "source": "./config.json",
      "adapter": "zod"
    },
    {
      "id": "Event",
      "sourceType": "json",
      "source": {
        "type": "object",
        "properties": {
          "name": { "type": "string" }
        }
      },
      "adapter": "zod"
    }
  ]
}
```

### Python Example

```jsonc
{
  "$schema": "https://xschema.dev/schemas/py.jsonc",
  "namespace": "models",
  "schemas": [
    {
      "id": "User",
      "sourceType": "url",
      "source": "https://example.com/schemas/user.json",
      "adapter": "pydantic"
    }
  ]
}
```

### Hypothetical Rust Example

```jsonc
{
  "$schema": "https://xschema.dev/schemas/rs.jsonc",
  "schemas": [
    {
      "id": "User",
      "sourceType": "url",
      "source": "https://example.com/schemas/user.json",
      "adapter": "serde"
    }
  ]
}
```
