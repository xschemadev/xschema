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
│                                User Code                                    │
│                                                                             │
│   import { zodAdapter } from '@xschema/adapter-zod'                         │
│   import { xschema } from '.xschema'                                        │
│                                                                             │
│   // From inline schema                                                     │
│   const User = xschema({                                                    │
│     type: 'object',                                                         │
│     properties: { name: { type: 'string' } },                               │
│     required: ['name']                                                      │
│   } as const, zodAdapter)                                                   │
│                                                                             │
│   // From file                                                              │
│   const Post = xschema('./schemas/post.json', zodAdapter)                   │
│                                                                             │
│   // From URL                                                               │
│   const ApiResponse = xschema('https://api.example.com/schema.json', zodAdapter)
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
                                      │ $ xschema generate
                                      ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                              xschema-cli                                    │
│                              (Go binary)                                    │
│                                                                             │
│   1. Parse codebase, find all xschema() calls                               │
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
│   Python:      .xschema/__init__.py + .xschema/__init__.pyi                 │
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
│   // TypeScript                                                             │
│   import { xschema } from '.xschema'                                        │
│   const User = xschema(schema, zodAdapter)  // ← Now fully typed!           │
│   User.parse(data)                          // ← Real Zod schema            │
│   type UserType = z.infer<typeof User>      // ← Type extraction works      │
│                                                                             │
│   // Python                                                                 │
│   from .xschema import xschema                                              │
│   User = xschema(schema, pydantic_adapter)  # ← Now a real Pydantic model   │
│   user = User(name="Alice")                 # ← Full Pydantic API           │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## **Core Components**

### **1. xschema-cli (Go Binary)**

The orchestrator. Does NOT do schema conversion itself.

**Responsibilities:**

- Parse codebases in any language
- Find `xschema()` calls
- Extract schema content (inline, file, URL)
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

Three ways to provide schemas:

```
import { zodAdapter } from '@xschema/adapter-zod';
import { xschema } from '.xschema';

// 1. Inline literal (must use 'as const')
const User = xschema({
  type: 'object',
  properties: {
    name: { type: 'string' },
    age: { type: 'number' }
  },
  required: ['name']
} as const, zodAdapter);

// 2. Local file
const Post = xschema('./schemas/post.json', zodAdapter);

// 3. Remote URL
const ApiResponse = xschema('https://api.example.com/schema.json', zodAdapter);
```

---

Following some examples. Go directly to [summary](https://www.notion.so/XSchema-2d43965fc6ad80608c3af8d317806837?pvs=21)

## **Adapter Examples**

### **TypeScript + Zod**

**Adapter: `@xschema/adapter-zod`**

```
// @xschema/adapter-zod/index.ts
import { jsonSchemaToZod } from 'json-schema-to-zod';

export const zodAdapter = {
  __brand: 'xschema-adapter' as const,
  name: 'zod',
  package: '@xschema/adapter-zod',
  language: 'typescript',
  runtime: 'bun',  // or 'node', 'tsx', etc.
};

export function convert(schema: JSONSchema): ConvertResult {
  const zodCode = jsonSchemaToZod(schema);
  return {
    code: zodCode,
    imports: ["import { z } from 'zod'"],
  };
}
```

**User Code:**

```
// src/models.ts
import { zodAdapter } from '@xschema/adapter-zod';
import { xschema } from '.xschema';

const userSchema = {
  type: 'object',
  properties: {
    id: { type: 'string', format: 'uuid' },
    name: { type: 'string', minLength: 1 },
    email: { type: 'string', format: 'email' },
    age: { type: 'number', minimum: 0 },
  },
  required: ['id', 'name', 'email'],
} as const;

// Before CLI: User is PleaseRunXSchemaCli
// After CLI: User is z.ZodObject<...>
const User = xschema(userSchema, zodAdapter);

// Full Zod API available
const result = User.safeParse(data);
if (result.success) {
  console.log(result.data.name);  // typed as string
}

// Type extraction works
type UserType = z.infer<typeof User>;
// { id: string; name: string; email: string; age?: number }
```

**Generated: `.xschema/generate/generate.ts`**

```
// Generated by xschema-cli - DO NOT EDIT
import * as fs from 'fs';
import { convert as zodConvert } from '@xschema/adapter-zod';

const schemas = [
  {
    hash: 'a1b2c3',
    adapter: 'zod',
    schema: {"type":"object","properties":{"id":{"type":"string","format":"uuid"},"name":{"type":"string","minLength":1},"email":{"type":"string","format":"email"},"age":{"type":"number","minimum":0}},"required":["id","name","email"]},
    overloadType: `{
      readonly type: 'object';
      readonly properties: {
        readonly id: { readonly type: 'string'; readonly format: 'uuid' };
        readonly name: { readonly type: 'string'; readonly minLength: 1 };
        readonly email: { readonly type: 'string'; readonly format: 'email' };
        readonly age: { readonly type: 'number'; readonly minimum: 0 };
      };
      readonly required: readonly ['id', 'name', 'email'];
    }`,
  },
];

const converters = { zod: zodConvert };

// Convert and generate output
const results = schemas.map(({ hash, adapter, schema }) => ({
  hash,
  ...converters[adapter](schema),
}));

// Build output file
let output = `// Generated by xschema - DO NOT EDIT
import { z } from 'zod';
import { zodAdapter } from '@xschema/adapter-zod';

`;

for (const { hash, code } of results) {
  output += `const schema_${hash} = ${code};\n\n`;
}

for (const s of schemas) {
  output += `export function xschema(
  schema: ${s.overloadType},
  adapter: typeof zodAdapter
): typeof schema_${s.hash};\n\n`;
}

output += `export function xschema(schema: unknown, adapter: unknown): PleaseRunXSchemaCli;

export function xschema(schema: any, adapter: any) {
  const hash = computeHash(schema, adapter.name);
  const registry: Record<string, unknown> = {
    ${schemas.map(s => `'${s.hash}': schema_${s.hash}`).join(',\n    ')}
  };
  return registry[hash] ?? (() => { throw new Error('Run xschema generate'); })();
}

export type PleaseRunXSchemaCli = {
  readonly __brand: unique symbol;
  readonly __error: 'Run \`xschema generate\`';
};

function computeHash(schema: unknown, adapterName: string): string {
  const crypto = require('crypto');
  return crypto.createHash('sha256')
    .update(JSON.stringify(schema) + adapterName)
    .digest('hex')
    .slice(0, 6);
}
`;

fs.writeFileSync('.xschema/index.ts', output);
console.log('✓ Generated .xschema/index.ts');
```

**Generated: `.xschema/index.ts`**

```
// Generated by xschema - DO NOT EDIT
import { z } from 'zod';
import { zodAdapter } from '@xschema/adapter-zod';

const schema_a1b2c3 = z.object({
  id: z.string().uuid(),
  name: z.string().min(1),
  email: z.string().email(),
  age: z.number().min(0).optional(),
});

export function xschema(
  schema: {
    readonly type: 'object';
    readonly properties: {
      readonly id: { readonly type: 'string'; readonly format: 'uuid' };
      readonly name: { readonly type: 'string'; readonly minLength: 1 };
      readonly email: { readonly type: 'string'; readonly format: 'email' };
      readonly age: { readonly type: 'number'; readonly minimum: 0 };
    };
    readonly required: readonly ['id', 'name', 'email'];
  },
  adapter: typeof zodAdapter
): typeof schema_a1b2c3;

export function xschema(schema: unknown, adapter: unknown): PleaseRunXSchemaCli;

export function xschema(schema: any, adapter: any) {
  const hash = computeHash(schema, adapter.name);
  const registry: Record<string, unknown> = {
    'a1b2c3': schema_a1b2c3,
  };
  return registry[hash] ?? (() => { throw new Error('Run xschema generate'); })();
}

export type PleaseRunXSchemaCli = {
  readonly __brand: unique symbol;
  readonly __error: 'Run `xschema generate`';
};

function computeHash(schema: unknown, adapterName: string): string {
  const crypto = require('crypto');
  return crypto.createHash('sha256')
    .update(JSON.stringify(schema) + adapterName)
    .digest('hex')
    .slice(0, 6);
}
```

---

### **Python + Pydantic**

**Adapter: `xschema-adapter-pydantic`**

```
# xschema_adapter_pydantic/__init__.py
from dataclasses import dataclass
from typing import Any

@dataclass(frozen=True)
class XSchemaAdapter:
    __brand: str = "xschema-adapter"
    name: str = ""
    package: str = ""
    language: str = ""
    runtime: str = ""

pydantic_adapter = XSchemaAdapter(
    name="pydantic",
    package="xschema-adapter-pydantic",
    language="python",
    runtime="python3",
)

def convert(schema: dict) -> dict:
    """Convert JSON Schema to Pydantic model code."""
    code, class_name = _json_schema_to_pydantic(schema)
    return {
        "code": code,
        "imports": [
            "from pydantic import BaseModel, Field, EmailStr",
            "from typing import Optional",
            "from uuid import UUID",
        ],
        "class_name": class_name,
    }

def _json_schema_to_pydantic(schema: dict) -> tuple[str, str]:
    # Implementation converts JSON Schema to Pydantic model
    ...
```

**User Code:**

```
# src/models.py
from xschema_adapter_pydantic import pydantic_adapter
from .xschema import xschema

user_schema = {
    "type": "object",
    "properties": {
        "id": {"type": "string", "format": "uuid"},
        "name": {"type": "string", "minLength": 1},
        "email": {"type": "string", "format": "email"},
        "age": {"type": "integer", "minimum": 0},
    },
    "required": ["id", "name", "email"],
}

# Before CLI: User is PleaseRunXSchemaCli type
# After CLI: User is a Pydantic model class
User = xschema(user_schema, pydantic_adapter)

# Full Pydantic API available
user = User(
    id="123e4567-e89b-12d3-a456-426614174000",
    name="Alice",
    email="alice@example.com"
)
print(user.name)  # typed in IDE

# Validation
try:
    User(id="invalid", name="", email="not-an-email")
except ValidationError as e:
    print(e.errors())

# Serialization
print(user.model_dump_json())
```

**Generated: `.xschema/generate/generate.py`**

```
# Generated by xschema-cli - DO NOT EDIT
import json
from pathlib import Path
from xschema_adapter_pydantic import convert as pydantic_convert

schemas = [
    {
        "hash": "a1b2c3",
        "adapter": "pydantic",
        "schema": {"type":"object","properties":{"id":{"type":"string","format":"uuid"},"name":{"type":"string","minLength":1},"email":{"type":"string","format":"email"},"age":{"type":"integer","minimum":0}},"required":["id","name","email"]},
    },
]

converters = {"pydantic": pydantic_convert}

results = [
    {"hash": s["hash"], **converters[s["adapter"]](s["schema"])}
    for s in schemas
]

# Collect imports
all_imports = set()
for r in results:
    all_imports.update(r["imports"])

# Build output
output = '''# Generated by xschema - DO NOT EDIT
{imports}

{classes}

_registry = {{
{registry}
}}

def xschema(schema, adapter):
    hash_key = _compute_hash(schema, adapter.name)
    result = _registry.get(hash_key)
    if result is None:
        raise RuntimeError("Run 'xschema generate'")
    return result

def _compute_hash(schema, adapter_name: str) -> str:
    import hashlib
    content = json.dumps(schema, sort_keys=True) + adapter_name
    return hashlib.sha256(content.encode()).hexdigest()[:6]
'''.format(
    imports="\n".join(sorted(all_imports)),
    classes="\n\n".join(r["code"] for r in results),
    registry=",\n".join(f'    "{s["hash"]}": Schema_{s["hash"]}' for s in schemas),
)

Path(".xschema/__init__.py").write_text(output)
print("✓ Generated .xschema/__init__.py")
```

**Generated: `.xschema/__init__.py`**

```
# Generated by xschema - DO NOT EDIT
from pydantic import BaseModel, Field, EmailStr
from typing import Optional
from uuid import UUID
import json

class Schema_a1b2c3(BaseModel):
    id: UUID
    name: str = Field(min_length=1)
    email: EmailStr
    age: Optional[int] = Field(default=None, ge=0)

_registry = {
    "a1b2c3": Schema_a1b2c3,
}

def xschema(schema, adapter):
    hash_key = _compute_hash(schema, adapter.name)
    result = _registry.get(hash_key)
    if result is None:
        raise RuntimeError("Run 'xschema generate'")
    return result

def _compute_hash(schema, adapter_name: str) -> str:
    import hashlib
    content = json.dumps(schema, sort_keys=True) + adapter_name
    return hashlib.sha256(content.encode()).hexdigest()[:6]
```

**Generated: `.xschema/__init__.pyi`** (Type stub for IDE support)

```
# Generated by xschema - DO NOT EDIT
from typing import Type, overload
from xschema_adapter_pydantic import pydantic_adapter

class Schema_a1b2c3:
    id: str
    name: str
    email: str
    age: int | None
    def __init__(self, *, id: str, name: str, email: str, age: int | None = None) -> None: ...
    def model_dump_json(self) -> str: ...

class PleaseRunXSchemaCli:
    ...

@overload
def xschema(
    schema: dict,
    adapter: type[pydantic_adapter]
) -> Type[Schema_a1b2c3]: ...

@overload
def xschema(schema: object, adapter: object) -> Type[PleaseRunXSchemaCli]: ...
```

---

## **Output Examples by Language**

| Language | Output Files |
| --- | --- |
| TypeScript | `.xschema/index.ts` |
| Python | `.xschema/__init__.py` + `.xschema/__init__.pyi` |
| Go | `.xschema/xschema.go` |
| Rust | `.xschema/mod.rs` |

---

## **CLI Commands**

```
# Generate schemas
xschema generate

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

---

## **Development Workflow**

### **Manual Generation**

```
# Run once after changing schemas
xschema generate
```

### **Watch Mode**

```
# Automatically regenerate on file changes
xschema generate --watch
```

### **Build Tool Integrations**

We provide plugins that run generation automatically, so you never have to think about it:

**Vite:**

```
// vite.config.ts
import { xschemaPlugin } from '@xschema/vite-plugin';

export default {
  plugins: [xschemaPlugin()],
};
```

**Next.js:**

```
// next.config.js
const { withXSchema } = require('@xschema/next');

module.exports = withXSchema({
  // your next config
});
```

**Webpack:**

```
// webpack.config.js
const { XSchemaPlugin } = require('@xschema/webpack-plugin');

module.exports = {
  plugins: [new XSchemaPlugin()],
};
```

These plugins watch for changes and regenerate `.xschema/` automatically during development, and ensure generation runs before production builds.

---

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
├── plugins/
│   ├── vite/                     # @xschema/vite-plugin
│   ├── next/                     # @xschema/next
│   └── webpack/                  # @xschema/webpack-plugin
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

- A CLI that finds `xschema()` calls in your code
- Extracts JSON Schemas (inline, file, URL)
- Uses adapters to generate native validator code
- Outputs a typed `xschema()` function you import

**Adapters are:**

- Packages with an identifier + convert function
- Identifier used in user code to specify which adapter
- Convert function called during generation to produce native code
- Specify their own language/runtime for the generate script

**What gets generated (in `.xschema/`):**

- `generate/generate.{ts,py,go,sh,...}` - The conversion scripts (one per language needed)
- `index.{ts,py,go,...}` - The output with native schemas + typed `xschema()` function

**The value:**

- One JSON Schema → native validators in any language
- Full type safety
- No wrapper objects—you get the real Zod schema, Pydantic model, etc.
- Adapters are verified against a comprehensive test suite
- Custom adapters are easy to create
