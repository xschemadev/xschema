# **XSchema Specification**

## **The Problem**

You have a JSON Schema—in a URL, in a file, or directly in your code. You want:

1. A native validator in your language (Zod, Pydantic, Go structs, etc.)
2. Full type safety at compile time

Today's solutions force you to choose: universal schema (JSON Schema) with no types, or language-specific schemas (Zod) with no portability.

And if you try to convert between them? The converters are hard to build, often just random scripts someone wrote, and you never know if they actually handle all edge cases correctly. Does it support `allOf`? What about string formats? Nested refs? Usually you find out when something breaks in production.

## **The Goal**

```
JSON Schema → Native, fully-typed validator in ANY language
```

The result IS the native schema object. Not a wrapper. The actual Zod schema, Pydantic model, Go struct.

**Our solution to the converter quality problem:** A comprehensive test suite that all adapters must pass. Hundreds of JSON Schema test cases covering every feature—primitives, formats, constraints, combinators, refs, edge cases. When you use an xschema adapter, you know it works.

---

## **Architecture Overview**

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                User Code                                    │
│                                                                             │
│   import { zodAdapter } from '@xschema/adapter-zod'                         │
│   import { xschema } from '.xschema'                                        │
│                                                                             │
│   // From URL                                                               │
│   xschema.fromURL('User', 'https://api.example.com/user.json', zodAdapter)  │
│                                                                             │
│   // From file                                                              │
│   xschema.fromFile('Post', './schemas/post.json', zodAdapter)               │
│                                                                             │
│   // From inline schema                                                     │
│   xschema.fromSchema('Comment', { type: 'object', ... } as const, zodAdapter)│
│                                                                             │
│   // Use generated schemas                                                  │
│   xschema.User.parse(data)                                                          │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
                                      │ $ xschema generate
                                      ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                              xschema-cli                                    │
│                              (Go binary)                                    │
│                                                                             │
│   1. Parse codebase, find all FromURL/FromFile/FromSchema calls             │
│   2. Extract schemas (inline / file / URL) and adapter identifiers          │
│   3. Generate scripts in .xschema/generate/—one file per language needed:   │
│                                                                             │
│        .xschema/generate/                                                   │
│        ├── generate.ts    (for TypeScript adapters like zod, ajv)           │
│        ├── generate.py    (for Python adapters like pydantic)               │
│        ├── generate.sh    (for shell-based adapters)                        │
│        └── generate.go    (for Go adapters)                                 │
│                                                                             │
│   4. Run all generated scripts (bun/node/python/bash/go/etc.)               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
                                      │ outputs
                                      ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                              .xschema/                                      │
│                                                                             │
│   TypeScript:  .xschema/index.ts                                            │
│   Python:      .xschema/__init__.py                                         │
│   Go:          .xschema/xschema.go                                          │
│   Rust:        .xschema/mod.rs                                              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
                                      │ import
                                      ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                              Back to User Code                              │
│                                                                             │
│   // TypeScript                                                             │
│   import { xschema } from '.xschema'                                        │
│   xschema.fromURL('User', 'https://...', zodAdapter)                        │
│   xschema.User.parse(data)                  // ← Real Zod schema            │
│   type UserType = z.infer<typeof xschema.User>  // ← Type extraction works  │
│                                                                             │
│   // Python                                                                 │
│   from .xschema import xschema                                              │
│   xschema.from_url('User', 'https://...', pydantic_adapter)                 │
│   user = xschema.User(name="Alice")         # ← Full Pydantic API           │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## **Core Components**

### **1. xschema-cli (Go Binary)**

The orchestrator. Does NOT do schema conversion itself.

**Responsibilities:**

- Parse codebases in any language
- Find `xschema.fromURL()`, `xschema.fromFile()`, `xschema.fromSchema()` calls
- Extract schema name, source (URL/file/inline), and adapter
- Identify which adapter each call uses
- Generate the appropriate scripts in `.xschema/generate/`
- Execute all scripts using their respective runtimes

**Why Go:**

- Single binary, no runtime dependencies
- Fast parsing
- Easy cross-platform distribution

### **2. Adapters**

Each adapter is a package that tells the CLI how to generate code for a specific validation library.

**Adapter Structure:**

```
@xschema/adapter-{name}/
├── index.ts (or __init__.py, etc.)
│   ├── {name}Adapter    # Identifier used in user code
│   └── convert()        # Function that converts JSON Schema to native code
├── package.json (or pyproject.toml, etc.)
└── README.md
```

**Adapter Identifier Interface:**

```
interface XSchemaAdapter {
  readonly __brand: 'xschema-adapter';
  readonly name: string;           // e.g., 'zod', 'pydantic'
  readonly package: string;        // e.g., '@xschema/adapter-zod'
  readonly language: string;       // e.g., 'typescript', 'python'
  readonly runtime: string;        // e.g., 'bun', 'node', 'python3'
}
```

**Convert Function Interface:**

```
interface ConvertResult {
  code: string;        // The generated code expression
  imports: string[];   // Required import statements
}

type ConvertFunction = (schema: JSONSchema) => ConvertResult;
```

### **3. Generated Output (.xschema/)**

The CLI generates everything into the `.xschema/` directory:

```
.xschema/
├── generate/
│   ├── generate.ts      # Script for TypeScript adapters
│   └── generate.py      # Script for Python adapters (if needed)
├── index.ts             # Output for TypeScript
└── __init__.py          # Output for Python (if needed)
```

---

**Schema Sources**

Three functions to provide schemas, all with `name` as first argument. Access via `xschema` namespace with language-appropriate casing:

```typescript
// TypeScript - camelCase
import { zodAdapter } from '@xschema/adapter-zod';
import { xschema } from '.xschema';

xschema.fromURL('User', 'https://api.example.com/user.json', zodAdapter);
xschema.fromFile('Post', './schemas/post.json', zodAdapter);
xschema.fromSchema('Comment', { type: 'object', ... } as const, zodAdapter);

xschema.User.parse(data);
type UserType = z.infer<typeof xschema.User>;
```

```python
# Python - snake_case
from .xschema import xschema

xschema.from_url('User', 'https://...', pydantic_adapter)
xschema.from_file('Post', './schemas/post.json', pydantic_adapter)
xschema.from_schema('Comment', {...}, pydantic_adapter)

user = xschema.User(name="Alice")
```

```go
// Go - PascalCase (via package namespace)
import "example/xschema"

xschema.FromURL("User", "https://...", adapter)
xschema.FromFile("Post", "./schemas/post.json", adapter)

user := xschema.User{Name: "Alice"}
```

**Signature:**
```
// TypeScript
xschema.fromURL(name, url, adapter)
xschema.fromFile(name, path, adapter)
xschema.fromSchema(name, schema, adapter)

// Python
xschema.from_url(name, url, adapter)
xschema.from_file(name, path, adapter)
xschema.from_schema(name, schema, adapter)

// Go
xschema.FromURL(name, url, adapter)
xschema.FromFile(name, path, adapter)
// FromSchema not available in Go
```

The `name` is always required and determines the exported schema name.

---

Examples can be found at /examples

---

## **CLI Commands**

```
# Generate schemas
xschema generate

# Init xschema
xschema init

# Generate with custom output directory
xschema generate --output .xschema

# Watch mode for development
xschema generate --watch

# Check if generation is needed (for CI)
xschema generate --check

# List found schemas in codebase
xschema list

# Verify an adapter works correctly
xschema verify @xschema/adapter-zod
```

---

## **Verifier (Test Suite)**

The verifier is a comprehensive test suite that validates adapters work correctly.

**What it does:**

1. Takes an adapter as input
2. Runs the adapter's `convert()` function against hundreds of JSON Schema test cases
3. For each test case, validates that:
    - Valid inputs pass validation
    - Invalid inputs fail validation
    - Edge cases are handled correctly
4. Reports which schemas pass/fail and why

**Test categories:**

- Primitive types (string, number, boolean, null)
- String formats (email, uuid, uri, date-time, etc.)
- String constraints (minLength, maxLength, pattern)
- Number constraints (minimum, maximum, multipleOf)
- Arrays (items, minItems, maxItems, uniqueItems)
- Objects (properties, required, additionalProperties)
- Nested schemas
- Combinators (allOf, anyOf, oneOf, not)
- References ($ref)
- Edge cases and error conditions

**Usage:**

```
xschema verify @xschema/adapter-zod
xschema verify xschema-adapter-pydantic
```

## **Project Structure**

```
xschema/
├── cli/                          # Go CLI
│   ├── cmd/
│   │   └── generate.go
│   ├── parser/
│   │   ├── typescript.go
│   │   ├── python.go
│   │   └── go.go
│   └── main.go
│
├── adapters/
│   ├── typescript/
│   │   └── zod/                  # @xschema/adapter-zod
│   │       ├── index.ts
│   │       └── package.json
│   │
│   └── python/
│       └── pydantic/             # xschema-adapter-pydantic
│           ├── __init__.py
│           └── pyproject.toml
│
├── verifier/                     # Test suite for adapters
│   └── ...
│
└── docs/
    └── README.md
```

---

## **Summary**

**xschema is:**

- A CLI that finds `xschema.fromURL/from_url/FromURL()` calls in your code
- Extracts JSON Schemas (inline, file, URL)
- Uses adapters to generate native validator code
- Outputs typed schemas you import directly by name

**Adapters are:**

- Packages with an identifier + convert function
- Identifier used in user code to specify which adapter
- Convert function called during generation to produce native code
- Specify their own language/runtime for the generate script

**What gets generated (in `.xschema/`):**

- `generate/generate.{ts,py,go,sh,...}` - The conversion scripts (one per language needed)
- `index.{ts,py,go,...}` - The output with native schemas + typed `xschema()` function

The `name` is always required and determines the exported schema name.

**The value:**

- One JSON Schema → native validators in any language
- Full type safety
- No wrapper objects—you get the real Zod schema, Pydantic model, etc.
- Adapters are verified against a comprehensive test suite
- Custom adapters are easy to create
