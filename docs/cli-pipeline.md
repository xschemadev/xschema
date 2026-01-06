# CLI Pipeline

The xschema CLI pipeline is the core of the system. It orchestrates the entire process from parsing config files to writing generated code to disk.

## Pipeline Overview

The pipeline consists of 4 sequential steps:

```
┌─────────┐    ┌──────────┐    ┌───────────┐    ┌─────────┐
│ 1. Parse│ →  │ 2.Fetch  │ →  │ 3.Generate│ →  │4.Inject │
│ Configs │    │ Schemas  │    │ Code      │    │ Files   │
└─────────┘    └──────────┘    └───────────┘    └─────────┘
```

Each step is independent and can be understood separately, but they form a cohesive pipeline.

## Step 1: Parse

**Package:** `cli/parser/`

**Input:** Project root directory

**Output:** `ParseResult` containing declarations and detected language

### What Parser Does

1. **Finds all config files** in the project
   - Searches for `*.json` and `*.jsonc` files
   - Uses `git ls-files` (fast, respects .gitignore)
   - Falls back to directory walk if git unavailable

2. **Identifies xschema configs** by `$schema` field
   - Checks if `$schema` starts with `https://xschema.dev/schemas/`
   - Skips non-xschema JSON files

3. **Detects language** from schema URL
   - `ts.jsonc` → TypeScript
   - `py.jsonc` → Python
   - Errors if multiple languages detected (unless `--lang` flag)

4. **Extracts schema declarations** from each config
   - Each `schemas` array entry becomes a declaration
   - Namespace defaults to filename (or explicit override)
   - Validates no duplicate IDs in same namespace

### Code Flow

**Location:** `cli/parser/parser.go`

```go
func Parse(ctx context.Context, projectRoot string, langFilter string) (*ParseResult, error) {
    // 1. Find all JSON/JSONC files
    files := getConfigFiles(ctx, projectRoot)
    
    // 2. Parse each file and filter by xschema $schema
    var configs []ConfigFile
    for _, path := range files {
        config := parseConfigFile(path)  // Returns nil if not xschema
        if config != nil {
            configs = append(configs, *config)
        }
    }
    
    // 3. Check for language conflicts
    if multipleLanguagesDetected && langFilter == "" {
        return nil, fmt.Errorf("multiple languages detected. Use --lang to specify")
    }
    
    // 4. Merge declarations and validate
    declarations := mergeDeclarations(configs)
    
    return &ParseResult{
        Language:     detectedLanguage,
        Configs:      configs,
        Declarations: declarations,
    }
}
```

### Configuration File Format

**Basic Structure:**

```jsonc
{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  "namespace": "optional-override",
  "schemas": [
    {
      "id": "SchemaName",
      "sourceType": "url" | "file" | "json",
      "source": "...",
      "adapter": "adapter-name"
    }
  ]
}
```

**Complete Example:**

```jsonc
{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  
  // Optional: override namespace (defaults to filename)
  "namespace": "models",
  
  "schemas": [
    // URL-based schema
    {
      "id": "User",
      "sourceType": "url",
      "source": "https://json.schemastore.org/package.json",
      "adapter": "zod"
    },
    
    // File-based schema (relative to config file)
    {
      "id": "Calendar",
      "sourceType": "file",
      "source": "./schemas/calendar.json",
      "adapter": "zod"
    },
    
    // Inline JSON schema
    {
      "id": "Event",
      "sourceType": "json",
      "source": {
        "type": "object",
        "properties": {
          "name": { "type": "string" },
          "startDate": { "type": "string", "format": "date-time" }
        },
        "required": ["name", "startDate"]
      },
      "adapter": "zod"
    }
  ]
}
```

### Parser Data Types

```go
// cli/parser/types.go

type Declaration struct {
    Namespace  string          // e.g., "user"
    ID         string          // e.g., "Profile"
    SourceType SourceType      // "url", "file", "json"
    Source     json.RawMessage // URL/path string or inline JSON
    Adapter    string          // e.g., "zod"
    ConfigPath string          // Path to config file (for relative resolution)
}

type ConfigFile struct {
    Path      string             // Absolute path
    Namespace string             // From filename or explicit
    Language  *language.Language // Detected language
    Schemas   []SchemaEntryRaw   // Raw entries
}

type ParseResult struct {
    Language     *language.Language // Detected language
    Configs      []ConfigFile       // All parsed configs
    Declarations []Declaration      // Flattened declarations
}
```

### Rules and Validation

1. **Namespace Defaults**
   - Defaults to config filename without extension
   - Example: `user.jsonc` → namespace `"user"`
   - Can override with explicit `"namespace"` field

2. **Multiple Files Same Namespace**
   - Configs with same namespace are merged
   - Example: `user.jsonc` and `user-extra.jsonc` → one namespace
   - Schemas from both are combined

3. **Duplicate Detection**
   - Error if same `id` appears twice in same namespace
   - Example: Two schemas with `"id": "User"` in namespace `"user"` → error
   - Different namespaces are OK

4. **Multiple Languages**
   - Error if configs mix languages (TS + Python)
   - Use `--lang` flag to filter to one language
   - Single language per invocation

## Step 2: Retrieve

**Package:** `cli/retriever/`

**Input:** Declarations from parser

**Output:** Retrieved schemas (raw JSON Schema documents)

### What Retriever Does

1. **Fetches schemas** from various sources
   - **URL:** HTTP GET request with retry logic
   - **File:** Read from filesystem (relative to config file)
   - **JSON:** Use inline schema directly

2. **Validates schemas** are valid JSON

3. **Caches results** to avoid redundant fetches

4. **Handles concurrency** - fetches up to 10 schemas in parallel

### Fetch Strategies

#### URL Fetching

```go
func retrieveFromURL(ctx context.Context, url string, opts Options) (json.RawMessage, error) {
    // Retries with exponential backoff
    // Default: 3 attempts, 500ms + 1000ms + 2000ms delays
    // Custom timeouts: 30 seconds per request
    // User-Agent: "xschema-cli/1.0"
}
```

**Retry Logic:**
- Attempt 1: Immediate
- Attempt 2: After 500ms
- Attempt 3: After 1500ms (1000ms more)
- Server errors (5xx) trigger retry
- Client errors (4xx) fail immediately

#### File Fetching

```go
func retrieveFromFile(ctx context.Context, filePath string, configPath string) (json.RawMessage, error) {
    // Resolve path relative to config file's directory
    configDir := filepath.Dir(configPath)
    fullPath := filepath.Join(configDir, filePath)
    return os.ReadFile(fullPath)
}
```

**Example Resolution:**
- Config: `/project/schemas/user.jsonc`
- Source: `./models/calendar.json`
- Resolved: `/project/schemas/models/calendar.json`

#### JSON Fetching

Inline schemas are used directly (no fetching needed).

### Caching

Built-in in-memory cache per CLI invocation:

```go
cache := newSchemaCache()

// Cache key strategy
switch source.SourceType {
case SourceURL:
    cacheKey = "url:" + url
case SourceFile:
    cacheKey = "file:" + absolutePath
case SourceJSON:
    cacheKey = "json:" + namespace + ":" + id
}

// Cache hit
if cached, ok := cache.get(cacheKey); ok {
    return cached
}

// Cache miss -> fetch -> store
schema := fetchSchema(...)
cache.set(cacheKey, schema)
```

Can be disabled with `--no-cache` flag.

### Concurrency

Fetches up to 10 schemas in parallel using `golang.org/x/sync/errgroup`:

```go
g, ctx := errgroup.WithContext(ctx)
g.SetLimit(10)  // Max 10 concurrent fetches

for _, decl := range declarations {
    g.Go(func() error {
        schema := fetch(...)
        results[idx] = schema
        return nil
    })
}

err = g.Wait()  // Wait for all to complete
```

### Retriever Data Types

```go
// cli/retriever/retriever.go

type RetrievedSchema struct {
    Namespace string
    ID        string
    Schema    json.RawMessage  // Raw JSON Schema
    Adapter   string           // Adapter name
}

type Options struct {
    Concurrency int
    HTTPTimeout time.Duration
    Retries     int
    NoCache     bool
}

func DefaultOptions() Options {
    return Options{
        Concurrency: 10,
        HTTPTimeout: 30 * time.Second,
        Retries:     3,
        NoCache:     false,
    }
}
```

## Step 3: Generate

**Package:** `cli/generator/`

**Input:** Retrieved schemas, language name

**Output:** Generated code from adapters

### What Generator Does

1. **Groups schemas by adapter**
   - Schemas using same adapter are batched together
   - Example: 3 schemas with `"adapter": "zod"` → one batch

2. **For each adapter batch:**
   - Detects package manager/runner
   - Constructs adapter command
   - Sends schemas via stdin
   - Reads generated code from stdout

3. **Collects and returns** all outputs

### Adapter Invocation

**Code Flow:**

```go
func Generate(ctx context.Context, input GenerateBatchInput) ([]GenerateOutput, error) {
    // 1. Get language config
    lang := language.ByName(input.Language)
    
    // 2. Detect runner
    runner, args, err := lang.DetectRunner()  // "bunx", [], nil
    
    // 3. Build adapter binary name
    binName := lang.AdapterBinPrefix + input.Adapter  // "xschema-zod"
    
    // 4. Execute: runner [args] binName
    cmdArgs := append(args, binName)
    cmd := exec.CommandContext(ctx, runner, cmdArgs...)
    
    // 5. Pipe schemas as JSON to stdin
    stdinData := json.Marshal(input.Schemas)
    cmd.Stdin = stdinData
    
    // 6. Capture stdout
    stdout := cmd.Output()
    
    // 7. Parse outputs
    var outputs []GenerateOutput
    json.Unmarshal(stdout, &outputs)
    
    return outputs, nil
}
```

**Example Execution:**

```bash
# For TypeScript with bun.lock
bunx xschema-zod

# Input (stdin):
[{
  "namespace": "user",
  "id": "Profile",
  "schema": { "type": "object", ... }
}]

# Output (stdout):
[{
  "namespace": "user",
  "id": "Profile",
  "imports": ["import { z } from \"zod\""],
  "schema": "z.object({ ... })",
  "type": "z.infer<typeof user_Profile>"
}]
```

### Batch Processing

Schemas are automatically grouped for efficiency:

```go
groups := GroupByAdapter(schemas)
// Result: { "zod": [schema1, schema2], "yup": [schema3] }

adapters := SortedAdapters(groups)  // Deterministic order

for _, adapter := range adapters {
    schemas := groups[adapter]
    outputs := Generate(ctx, GenerateBatchInput{
        Adapter:  adapter,
        Language: langName,
        Schemas:  schemas,  // Multiple schemas in one call
    })
}
```

### Generator Data Types

```go
// cli/generator/generator.go

type GenerateInput struct {
    Namespace string          `json:"namespace"`
    ID        string          `json:"id"`
    Schema    json.RawMessage `json:"schema"`
}

type GenerateOutput struct {
    Namespace string   `json:"namespace"`
    ID        string   `json:"id"`
    Schema    string   `json:"schema"`   // Generated code
    Type      string   `json:"type"`     // Type expression
    Imports   []string `json:"imports"`  // Import statements
}

type GenerateBatchInput struct {
    Adapter  string                      // e.g., "zod"
    Language string                      // e.g., "typescript"
    Schemas  []retriever.RetrievedSchema // Batch of schemas
}
```

## Step 4: Inject

**Package:** `cli/injector/`

**Input:** Generated outputs, language, output directory

**Output:** Written files (e.g., `.xschema/xschema.gen.ts`)

### What Injector Does

1. **Merges imports** (deduplicates)
   - Language-specific merging (handles TS vs Python syntax)
   - Deduplicates identical imports
   - Maintains import order

2. **Builds template data**
   - Converts outputs to template-friendly format
   - Generates variable names
   - Builds language-specific headers/footers

3. **Executes template** using Go `text/template`

4. **Writes output file** to disk

### Template Data

**Structure:**

```go
type TemplateData struct {
    Imports string                 // Merged import statements
    Schemas []language.SchemaEntry // Individual schemas
    Header  string                 // Language-specific header
    Footer  string                 // Language-specific footer
}

type SchemaEntry struct {
    Namespace string  // "user"
    ID        string  // "Profile"
    VarName   string  // "user_Profile"
    Code      string  // Generated validator code
    Type      string  // Type expression
}
```

### TypeScript Template

**Location:** `cli/language/templates.go`

```go
const TSTemplate = `// Generated by xschema - DO NOT EDIT
// https://xschema.dev/docs
{{.Imports}}

{{range .Schemas}}
{{- if and .Type (not .Code)}}
export type {{.VarName}} = {{.Type}};
{{- else if and .Code (not .Type)}}
const {{.VarName}} = {{.Code}};
{{- else if and .Code .Type}}
const {{.VarName}} = {{.Code}};
type {{.VarName}}Type = {{.Type}};
{{- end}}
{{end}}

export const schemas = {
{{- range .Schemas}}
  "{{.Key}}": {{.VarName}},
{{- end}}
} as const;

export type SchemaTypes = {
{{- range .Schemas}}
  "{{.Key}}": {{.VarName}}{{if and .Code .Type}}Type{{end}};
{{- end}}
};

declare module '@xschemadev/client' {
  interface Register {
    schemas: typeof schemas;
    schemaTypes: SchemaTypes;
  }
}
`
```

**Generated Output Example:**

```typescript
// Generated by xschema - DO NOT EDIT
import { z } from "zod"

const user_Profile = z.object({ name: z.string() });
type user_ProfileType = z.infer<typeof user_Profile>;

const user_Settings = z.object({ theme: z.enum(["light", "dark"]) });

export const schemas = {
  "user:Profile": user_Profile,
  "user:Settings": user_Settings,
} as const;

export type SchemaTypes = {
  "user:Profile": user_ProfileType;
  "user:Settings": typeof user_Settings;
};

declare module '@xschemadev/client' {
  interface Register {
    schemas: typeof schemas;
    schemaTypes: SchemaTypes;
  }
}
```

### Python Template

**Location:** `cli/language/templates.go`

```go
const PyTemplate = `# Generated by xschema - DO NOT EDIT
# https://xschema.dev/docs
from typing import Literal, overload
{{.Imports}}
from xschema import XSchemaBase, XSchemaAdapter

{{range .Schemas}}
{{.Code}}
{{end}}

_schemas: dict[str, type] = {
{{- range .Schemas}}
  "{{.Key}}": {{.VarName}},
{{- end}}
}

class xschema(XSchemaBase):
{{- range .Schemas}}
    {{.VarName}} = {{.VarName}}
{{- end}}

{{.Footer}}
`
```

### Import Merging

Language-specific deduplication and formatting:

#### TypeScript Import Merging

```go
func MergeTSImports(imports []string) string {
    // Input: ["import { z } from \"zod\"", "import { z } from \"zod\""]
    // Output: "import { z } from \"zod\""
    
    // Groups imports by source:
    // "zod" -> ["z"]
    // "zod-extensions" -> ["refine"]
    
    // Merges multiple imports from same source:
    // [import { z }, import { refine }] from "zod"
    // -> import { z, refine } from "zod"
}
```

#### Python Import Merging

```go
func MergePyImports(imports []string) string {
    // Input: ["from pydantic import BaseModel", "from pydantic import Field"]
    // Output: "from pydantic import BaseModel, Field"
}
```

### File Writing

**Location:** `cli/injector/injector.go`

```go
func Inject(input InjectInput) error {
    lang := language.ByName(input.Language)
    
    // Build template data
    data := buildTemplateData(input, lang)
    
    // Parse and execute template
    tmpl := template.New("inject").Parse(lang.Template)
    buf := tmpl.Execute(data)
    
    // Create output directory
    os.MkdirAll(input.OutDir, 0755)
    
    // Write file
    outPath := filepath.Join(input.OutDir, lang.OutputFile)
    os.WriteFile(outPath, buf.Bytes(), 0644)
    
    return nil
}
```

**Default output locations:**
- TypeScript: `.xschema/xschema.gen.ts`
- Python: `.xschema/__init__.py`

## CLI Entry Point

**Location:** `cli/cmd/generate.go`

The main `xschema generate` command orchestrates the entire pipeline:

```bash
xschema generate [options]

Options:
  -p, --project <dir>    Project root directory (default: current)
  -o, --output <dir>     Output directory (default: .xschema)
  --lang <language>      Filter to specific language
  -v, --verbose          Verbose output
  --dry-run              Show what would be generated
  -w, --watch            Watch for changes (TODO)
```

### Command Flow

```go
func runGenerate(cmd *cobra.Command, args []string) error {
    // Step 1: Parse
    result := parser.Parse(ctx, projectRoot, langFilter)
    
    // Step 2: Retrieve
    schemas := retriever.Retrieve(ctx, result.Declarations, retrieverOpts)
    
    // Handle dry-run
    if dryRun {
        return printDryRunOutput(schemas)
    }
    
    // Step 3: Generate
    outputs := generator.GenerateAll(ctx, schemas, result.Language.Name)
    
    // Step 4: Inject
    injector.Inject(injector.InjectInput{
        Language: result.Language.Name,
        Outputs:  outputs,
        OutDir:   outDir,
    })
    
    // Print summary
    printSummary(schemas, outDir, duration)
}
```

### Dry Run Mode

With `--dry-run` flag, the CLI shows what would be generated without writing files:

```bash
$ xschema generate --dry-run

Schemas that would be generated:
  zod
    • user:Profile
    • user:Settings
  pydantic
    • product:Item
```

## Error Handling

Each step has specific error handling:

1. **Parse errors:** Missing files, invalid JSON, language conflicts
2. **Retrieve errors:** HTTP failures, file not found, invalid schemas
3. **Generate errors:** Adapter not found, adapter crash, invalid output
4. **Inject errors:** Can't write files, template execution failure

All errors include context about which operation failed.

## Typical Workflow

**Example Project:**

```
my-project/
├── user.jsonc        # Declares User, Settings schemas
├── product.jsonc     # Declares Item, Cart schemas
├── package.json
└── bun.lock
```

**Step 1: Parse**
```
Found configs:
  • user.jsonc (namespace: user)
    - User (adapter: zod, source: url)
    - Settings (adapter: zod, source: url)
  • product.jsonc (namespace: product)
    - Item (adapter: zod, source: url)
    - Cart (adapter: zod, source: url)

Detected language: typescript
```

**Step 2: Retrieve**
```
Fetching 4 schemas:
  • user:User (from URL)
  • user:Settings (from URL)
  • product:Item (from URL)
  • product:Cart (from URL)

All schemas fetched successfully
```

**Step 3: Generate**
```
Grouped by adapter:
  • zod (4 schemas)

Running: bunx xschema-zod
  ✓ Generated 4 validators
```

**Step 4: Inject**
```
Writing: .xschema/xschema.gen.ts

Output file contains:
  - 4 validators (user_User, user_Settings, product_Item, product_Cart)
  - 2 import statements (zod)
  - 1 const export (schemas)
  - 1 type export (SchemaTypes)
```

**Result:**
```
.xschema/xschema.gen.ts created with 500 bytes
```
