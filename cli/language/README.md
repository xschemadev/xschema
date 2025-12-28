# Language Package

Defines language-specific configurations for parsing xschema calls and generating output.

## Adding a New Language

1. Add entry to `Languages` slice in `language.go`
2. Add tree-sitter queries in `queries.go`
3. Add template in `templates.go`
4. Add import merger in `imports.go`
5. Add test files in `parser/testdata/{lang}/`

---

## Language Struct

```go
type Language struct {
    Name          string                        // e.g., "typescript"
    Extensions    []string                      // e.g., [".ts", ".tsx"]
    GetSitterLang func() *sitter.Language       // tree-sitter language
    Query         string                        // tree-sitter query for xschema calls
    ImportQuery   string                        // query for adapter imports
    MethodMapping map[string]SourceType         // method name -> URL/File
    DetectRunner  func() (string, []string, error) // detect runtime (optional)
    
    // Client detection
    ClientPackage   string                      // e.g., "@xschema/client"
    ClientFactory   string                      // e.g., "createXSchemaClient"
    ClientQuery     string                      // query to find client variable
    ConfigQuery     string                      // query to extract config from client call
    ClientCallQuery string                      // query to find config object for injection
    
    // Client injection (after generation)
    BuildSchemasImport func(importPath string) string    // build import for schemas
    ImportPattern      string                            // regex to find import lines
    InjectSchemasKey   func(configContent string) string // inject "schemas" into config
    
    // Output generation
    OutputFile     string                       // e.g., "index.ts"
    Template       string                       // Go text/template
    MergeImports   func([]string) string        // dedupe/format imports
    BuildHeader    func(outDir string, schemas []SchemaEntry) string
    BuildFooter    func(outDir string, schemas []SchemaEntry) string
}
```

---

## Tree-Sitter Queries

Query must capture:
- `@method` - method name (fromURL, from_url, FromURL)
- `@name` - schema name (string literal)
- `@source` - URL or file path (string literal)
- `@adapter` - adapter identifier

Example (TypeScript):
```
(call_expression
  function: (member_expression
    object: (identifier) @obj
    property: (property_identifier) @method)
  arguments: (arguments 
    . [(string) (template_string)] @name 
    . [(string) (template_string)] @source 
    . (identifier) @adapter .)
  (#not-match? @name "\\$\\{")
  (#not-match? @source "\\$\\{"))
```

Note: The `@obj` capture is filtered in Go code to match the client variable name (found via `ClientQuery`).

### Client Query

The `ClientQuery` finds the variable name assigned to `createXSchemaClient()`. This allows users to name their client anything.

Captures required:
- `@client_name` - the variable name (e.g., `xschema`, `myClient`)

Example (TypeScript):
```
(lexical_declaration
  (variable_declarator
    name: (identifier) @client_name
    value: (call_expression
      function: (identifier) @_fn
      (#eq? @_fn "createXSchemaClient"))))
```
Matches: `const xschema = createXSchemaClient({});`

Example (Python):
```
(assignment
  left: (identifier) @client_name
  right: (call
    function: (identifier) @_fn
    (#eq? @_fn "create_xschema_client")))
```
Matches: `xschema = create_xschema_client({})`

### Config Query

The `ConfigQuery` extracts configuration options from the first argument of `createXSchemaClient()`.

Captures required:
- `@config_key` - property/key name
- `@config_value` - value (must be literal, variables not supported)

Example (TypeScript):
```
(call_expression
  function: (identifier) @_fn
  arguments: (arguments
    (object
      (pair
        key: (property_identifier) @config_key
        value: (_) @config_value)))
  (#eq? @_fn "createXSchemaClient"))
```
Matches config in: `createXSchemaClient({ output: ".generated", concurrency: 5 })`

Supported config keys:
- `output` - output directory (default: ".xschema")
- `concurrency` - max concurrent HTTP requests (default: 10)
- `httpTimeout` / `http_timeout` - HTTP timeout in ms (default: 30000)
- `retries` - max retry attempts (default: 3)

### Client Call Query

The `ClientCallQuery` finds the config object for injection. After generation, the CLI injects `schemas` import and adds it to the config.

Captures required:
- `@config` - the config object node
- `@schemas_key` (optional) - if schemas key already exists

Example (TypeScript):
```
(call_expression
  function: (identifier) @_fn
  arguments: (arguments
    (object) @config
    (pair
      key: (property_identifier) @_key
      (#eq? @_key "schemas"))? @schemas_key)
  (#eq? @_fn "createXSchemaClient"))
```

### Import Query

The `ImportQuery` captures adapter imports to populate `AdapterRef.Package`. This maps adapter variable names to their source package.

Captures required:
- `@package` - the import source/module (e.g., `"@xschema/adapter-zod"`)
- `@imported_name` - the imported identifier (e.g., `zodAdapter`)

The parser builds a map `importedName -> package` and uses it to populate `AdapterRef.Package` in each `Declaration`.

Example (TypeScript):
```
(import_statement
  source: (string) @package
  (import_specifier
    name: (identifier) @imported_name))
```
Matches: `import { zodAdapter } from "@xschema/adapter-zod"`

Example (Python):
```
(import_from_statement
  module_name: (dotted_name) @package
  (dotted_name) @imported_name)
```
Matches: `from xschema_adapter_pydantic import pydantic_adapter`

---

## Templates

Templates use Go's `text/template` to generate the final output file for each language.

General Guideline:
Start with the goal you have (what the generated file will look like), and then think about that could be generated from our template

### TemplateData

```go
type TemplateData struct {
    Imports string         // merged import statements
    Schemas []SchemaEntry  // all schemas to generate
    Header  string         // from BuildHeader()
    Footer  string         // from BuildFooter()
}

type SchemaEntry struct {
    Name string  // schema name from user code
    Code string  // generated validator code (e.g., "z.object({...})")
    Type string  // type expression (e.g., "z.infer<typeof User>")
}
```

### Template Syntax

```
{{.Header}}           - insert header (package decl, etc)
{{.Imports}}          - insert merged imports
{{.Footer}}           - insert footer
{{range .Schemas}}    - iterate over schemas
  {{.Name}}           - schema name
  {{.Code}}           - validator code
  {{.Type}}           - type expression
{{end}}

# Build comma-separated list inline:
{{range $i, $s := .Schemas}}{{if $i}}, {{end}}{{$s.Name}}{{end}}
```

### Conditional Output

Adapters can return Code only, Type only, or both. Handle all cases:

```
{{range .Schemas}}
{{- if and .Type (not .Code)}}
export type {{.Name}} = {{.Type}};
{{- else if and .Code (not .Type)}}
export const {{.Name}} = {{.Code}};
{{- else if and .Code .Type}}
export const {{.Name}} = {{.Code}};
export type {{.Name}}Type = {{.Type}};
{{- end}}
{{end}}
```

### Example: TypeScript Template

```go
const TSTemplate = `// Generated by xschema - DO NOT EDIT
{{- if .Header}}
{{.Header}}
{{- end}}
{{.Imports}}

{{range .Schemas}}
{{- if and .Code .Type}}
export const {{.Name}} = {{.Code}};
export type {{.Name}}Type = {{.Type}};
{{- end}}
{{end}}
export const schemas = { {{range $i, $s := .Schemas}}{{if $i}}, {{end}}{{$s.Name}}{{end}} } as const;

declare module '@xschema/client' {
  interface Register {
    schemas: typeof schemas
  }
}
`
```

The following are used to know what to put in the template, based on the `GeneratedOutput`

### Import Mergers

Adapters return import statements (e.g., `import { z } from "zod"`). Mergers dedupe and format them.

```go
func MergeTSImports(imports []string) string  // handles named/default imports
func MergePyImports(imports []string) string  // handles from X import Y
//...
```

---

## BuildHeader / BuildFooter

Generate language-specific preamble/postamble.

```go
// Python needs @overload stubs for type inference
func BuildPythonFooter(_ string, schemas []SchemaEntry) string {
    // Generate @overload stubs for from_url/from_file
}
```

---

## Testing

### Test Structure

```
cli/parser/testdata/
├── common/           # Expected results (shared across all langs)
│   ├── basic.json
│   ├── edge_cases.json
│   └── invalid.json
├── typescript/       # TS source files
│   ├── basic.ts
│   ├── edge_cases.ts
│   ├── invalid.ts
│   └── strings.ts    # language-specific tests
└── python/
    ├── basic.py
    ├── edge_cases.py
    ├── invalid.py
    └── strings.py
```

### Common Tests

Tests in `common/` are run against ALL languages. Each JSON defines expected parse results:

```json
// common/basic.json
[
  {"name": "User", "source": "url", "location": "https://api.example.com/user.json"},
  {"name": "Post", "source": "file", "location": "./schemas/post.json"}
]
```

Each language needs a corresponding source file that produces these results:

```typescript
// typescript/basic.ts
import { createXSchemaClient } from "@xschema/client";
import { zodAdapter } from "@xschema/zod";

const xschema = createXSchemaClient({ output: ".xschema" });

xschema.fromURL("User", "https://api.example.com/user.json", zodAdapter);
xschema.fromFile("Post", "./schemas/post.json", zodAdapter);
```

```python
# python/basic.py
from xschema import create_xschema_client
from xschema_pydantic import pydantic_adapter

xschema = create_xschema_client({ "output": ".xschema" })

xschema.from_url("User", "https://api.example.com/user.json", pydantic_adapter)
xschema.from_file("Post", "./schemas/post.json", pydantic_adapter)
```

### Adding a Common Test

1. Create `common/{test_name}.json` with expected results
2. Create `{lang}/{test_name}.{ext}` for each language
3. `TestCommon` in `parser_test.go` auto-discovers and runs them

### Language-Specific Tests

For features unique to a language (string syntax, etc), add:
- Source file: `{lang}/{test_name}.{ext}`
- Test function in `parser_test.go`

Example - testing Python string variants:

```python
# python/strings.py
xschema.from_url("DoubleQuote", "https://example.com/a.json", adapter)
xschema.from_url('SingleQuote', 'https://example.com/b.json', adapter)
xschema.from_url("""TripleDouble""", """https://example.com/c.json""", adapter)
xschema.from_file(r"RawString", r"./schemas/raw.json", adapter)
```

```go
// parser_test.go
func TestPythonStrings(t *testing.T) {
    lang := language.ByExtension(".py")
    decls, err := parseFile(context.Background(), "testdata/python/strings.py", lang)
    // ...assertions...
}
```

### Running Tests

```bash
go test ./language ./parser           # all tests
go test ./parser -run TestCommon      # common tests only
go test ./parser -run TestPython      # python-specific tests
```
