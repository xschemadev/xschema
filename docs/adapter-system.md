# Adapter System

The adapter system is the core extensibility mechanism of xschema. Adapters are language-specific converters that transform JSON Schema to native validators (Zod, Pydantic, etc.).

## Adapter Protocol

The adapter protocol uses a simple **stdin/stdout JSON interface** for communication between the Go CLI and language-specific adapters. This design allows adapters to be written in any language and run via package managers.

### Input Format (stdin)

The CLI sends a JSON array of `ConvertInput` objects to the adapter's stdin:

```typescript
interface ConvertInput {
  namespace: string;  // e.g., "user"
  id: string;         // e.g., "Profile"
  schema: object;     // Raw JSON Schema
}
```

**Example Input:**
```json
[
  {
    "namespace": "user",
    "id": "Profile",
    "schema": {
      "type": "object",
      "properties": {
        "name": { "type": "string" },
        "age": { "type": "integer" }
      },
      "required": ["name"]
    }
  },
  {
    "namespace": "user",
    "id": "Settings",
    "schema": {
      "type": "object",
      "properties": {
        "theme": { "type": "string", "enum": ["light", "dark"] }
      }
    }
  }
]
```

### Output Format (stdout)

Adapters must return a JSON array of `ConvertResult` objects to stdout:

```typescript
interface ConvertResult {
  namespace: string;        // Same as input
  id: string;               // Same as input
  imports: string[];        // Array of import statements (full statements, no trailing semicolon)
  schema?: string;          // Generated validator code
  type?: string;            // Type expression for inference
  validate?: string;        // Validation function for compliance harness (optional)
  validateImports?: string[]; // Imports needed by validate function (optional)
}
```

**Field Details:**

- `imports`: Full import statements without trailing semicolons (e.g., `import { z } from "zod"`). The generator merges and deduplicates imports from all schemas.
- `schema`: The generated validator/schema code expression
- `type`: Type expression for TypeScript type inference
- `validate`: Optional validation function used by the compliance harness to test runtime validation. Takes `data` parameter and returns `boolean`. Can reference `schema` variable. Empty string = type-only adapter (skips runtime validation).
- `validateImports`: Optional imports needed by the `validate` function, same format as `imports`. Merged with `imports` for harness generation.

**Example Output:**
```json
[
  {
    "namespace": "user",
    "id": "Profile",
    "imports": ["import { z } from \"zod\""],
    "schema": "z.object({ name: z.string(), age: z.number().int() })",
    "type": "z.infer<typeof user_Profile>",
    "validate": "(data) => schema.safeParse(data).success"
  },
  {
    "namespace": "user",
    "id": "Settings",
    "imports": ["import { z } from \"zod\""],
    "schema": "z.object({ theme: z.enum([\"light\", \"dark\"]) })",
    "type": "z.infer<typeof user_Settings>",
    "validate": "(data) => schema.safeParse(data).success"
  }
]
```

### Go Types (Internal CLI)

The CLI uses these types internally for type-safe handling:

```go
// cli/generator/generator.go

type GenerateInput struct {
    Namespace string          `json:"namespace"`
    ID        string          `json:"id"`
    Schema    json.RawMessage `json:"schema"`
}

type GenerateOutput struct {
    Namespace       string   `json:"namespace"`
    ID              string   `json:"id"`
    Schema          string   `json:"schema"`                     // generated code
    Type            string   `json:"type"`                       // type expression
    Imports         []string `json:"imports"`                    // required imports
    Validate        string   `json:"validate,omitempty"`         // validation function (compliance)
    ValidateImports []string `json:"validateImports,omitempty"`  // imports for validate
}
```

## How Adapters Are Called

### Adapter Discovery and Invocation

**Location:** `cli/generator/generator.go`

The process:

1. **Parse adapter name** from config file (e.g., `"zod"`)
2. **Construct binary name** using language-specific prefix:
   - Adapter: `"zod"` → Binary: `xschema-zod`
   - Prefix is defined per language (TypeScript: `"xschema-"`)
3. **Detect runner** from lockfiles:
   - `bun.lock` → `bunx`
   - `pnpm-lock.yaml` → `pnpm exec`
   - `yarn.lock` → `yarn`
   - `package-lock.json` → `npx`
4. **Execute:** `{runner} {binary}` (e.g., `bunx xschema-zod`)
5. **Pipe schemas** via stdin and collect results from stdout

### Batch Processing

Schemas are automatically grouped by adapter for efficient processing:

```go
// Schemas are grouped: { "zod": [schema1, schema2], "yup": [schema3] }
groups := retriever.GroupByAdapter(schemas)

for adapter, schemas := range groups {
    outputs := Generate(ctx, GenerateBatchInput{
        Adapter:  adapter,
        Language: langName,
        Schemas:  schemas,  // Batched together
    })
}
```

This allows a single adapter process to validate multiple schemas at once, improving performance.

## Example Adapter: @xschemadev/zod

The Zod adapter demonstrates the recommended adapter structure.

### Package Structure

```
typescript/packages/zod/
├── src/
│   ├── index.ts      # Core convert() function
│   └── cli.ts        # CLI entry point (executable)
├── package.json
├── tsconfig.json
└── dist/             # Compiled output
```

### Core Implementation (index.ts)

```typescript
import type { ConvertInput, ConvertResult } from "@xschemadev/core";
import { jsonSchemaToZod } from "json-schema-to-zod";

export function convert(input: ConvertInput): ConvertResult {
  const { namespace, id, schema } = input;
  const schemaCode = jsonSchemaToZod(schema);
  const varName = `${namespace}_${id}`;

  return {
    namespace,
    id,
    imports: ['import { z } from "zod"'],
    schema: schemaCode,
    type: `z.infer<typeof ${varName}>`,
    validate: "(data) => schema.safeParse(data).success",
  };
}
```

**Key Points:**
- Simple pure function that converts input to output
- Uses existing library (`json-schema-to-zod`) for heavy lifting
- Returns imports array (full import statements, no trailing semicolons)
- Type expression uses variable name pattern
- `validate` provides runtime validation for compliance testing

### CLI Entry Point (cli.ts)

```typescript
#!/usr/bin/env node
import { createAdapterCLI } from "@xschemadev/core";
import { convert } from "./index.js";

createAdapterCLI(convert);
```

**What `createAdapterCLI` does:**
1. Reads JSON array from stdin
2. Calls your `convert()` function for each item
3. Outputs JSON array to stdout
4. Handles errors gracefully

### Package Configuration (package.json)

```json
{
  "name": "@xschemadev/zod",
  "version": "0.0.1",
  "type": "module",
  "description": "xschema Zod adapter - convert JSON Schema to Zod validators",
  "exports": {
    ".": {
      "types": "./dist/index.d.ts",
      "import": "./dist/index.js"
    }
  },
  "bin": {
    "xschema-zod": "./dist/cli.js"
  },
  "files": ["dist"],
  "scripts": {
    "build": "tsc",
    "typecheck": "tsc --noEmit"
  },
  "dependencies": {
    "@xschemadev/core": "workspace:*",
    "json-schema-to-zod": "^2.4.1"
  },
  "peerDependencies": {
    "zod": "^3.0.0"
  }
}
```

**Important Fields:**
- `bin`: Maps `xschema-zod` command to executable script
- `dependencies`: Include `@xschemadev/core`
- `peerDependencies`: List the target library (Zod, Pydantic, etc.)
- `exports`: Makes functions available to other packages

## Creating a New Adapter

### Step-by-Step Guide

#### 1. Create Package Structure

```bash
mkdir -p packages/my-adapter/src
cd packages/my-adapter
```

#### 2. Create package.json

Choose an appropriate name and binary:

```json
{
  "name": "@xschemadev/my-adapter",
  "version": "0.0.1",
  "type": "module",
  "description": "xschema adapter for MyLib - convert JSON Schema to MyLib validators",
  "exports": {
    ".": {
      "types": "./dist/index.d.ts",
      "import": "./dist/index.js"
    }
  },
  "bin": {
    "xschema-my-adapter": "./dist/cli.js"
  },
  "files": ["dist"],
  "scripts": {
    "build": "tsc",
    "typecheck": "tsc --noEmit"
  },
  "dependencies": {
    "@xschemadev/core": "workspace:*"
  },
  "peerDependencies": {
    "my-lib": "^1.0.0"
  }
}
```

**Naming Convention:**
- Package: `@xschemadev/{library-name}` (e.g., `@xschemadev/pydantic`)
- Binary: `xschema-{library-name}` (e.g., `xschema-pydantic`)

#### 3. Implement Core Conversion (src/index.ts)

```typescript
import type { ConvertInput, ConvertResult } from "@xschemadev/core";
import { convertToMyLib } from "my-lib-converter";

/**
 * Converts a JSON Schema to a MyLib validator
 * @param input - Schema input with namespace, id, and raw JSON Schema
 * @returns Generated code and type information
 */
export function convert(input: ConvertInput): ConvertResult {
  const { namespace, id, schema } = input;

  // Your conversion logic
  const validatorCode = convertToMyLib(schema);
  const varName = `${namespace}_${id}`;

  return {
    namespace,
    id,
    imports: [
      'import { MyValidator } from "my-lib"',
      'import type { MyType } from "my-lib"',
    ],
    schema: validatorCode,
    type: `MyType<typeof ${varName}>`,
    validate: "(data) => MyValidator.check(schema, data)",
  };
}
```

**Best Practices:**
- Handle errors gracefully (throw descriptive errors)
- Import format: full import statements without trailing semicolons (e.g., `import { z } from "zod"`)
- Use consistent variable naming
- Return valid code that will compile in generated files
- Include `validate` function for runtime compliance testing (optional for type-only adapters)

#### 4. Create CLI Entry (src/cli.ts)

```typescript
#!/usr/bin/env node
import { createAdapterCLI } from "@xschemadev/core";
import { convert } from "./index.js";

createAdapterCLI(convert);
```

Make sure to include the shebang for executable permissions.

#### 5. Create tsconfig.json

```json
{
  "extends": "../../tsconfig.json",
  "compilerOptions": {
    "outDir": "./dist",
    "declaration": true,
    "declarationMap": true,
    "sourceMap": true
  },
  "include": ["src"],
  "exclude": ["node_modules", "dist"]
}
```

#### 6. Register in Workspace

Update `typescript/package.json`:

```json
{
  "workspaces": [
    "packages/core",
    "packages/client",
    "packages/zod",
    "packages/my-adapter"
  ]
}
```

#### 7. Build and Test

```bash
# Install dependencies
bun install

# Build
bun run build

# Test with real schema
echo '[{"namespace":"test","id":"User","schema":{"type":"object","properties":{"name":{"type":"string"}},"required":["name"]}}]' \
  | bunx xschema-my-adapter
```

Expected output:
```json
[{
  "namespace": "test",
  "id": "User",
  "imports": ["import { MyValidator } from \"my-lib\""],
  "schema": "MyValidator.object({ name: MyValidator.string() })",
  "type": "MyType<typeof test_User>",
  "validate": "(data) => MyValidator.check(schema, data)"
}]
```

## Core Utility: createAdapterCLI

**Location:** `typescript/packages/core/src/cli.ts`

The `createAdapterCLI` helper handles all CLI plumbing, so adapters only need to implement the conversion function.

```typescript
/**
 * Creates a CLI handler for xschema adapters.
 * Reads JSON array of ConvertInput from stdin, calls convert for each,
 * outputs JSON array of ConvertResult to stdout.
 *
 * @example
 * ```ts
 * #!/usr/bin/env node
 * import { createAdapterCLI } from "@xschemadev/core";
 * import { convert } from "./index";
 *
 * createAdapterCLI(convert);
 * ```
 */
export function createAdapterCLI(
  convert: (input: ConvertInput) => ConvertResult
): void {
  const chunks: string[] = [];
  process.stdin.on("data", (chunk) => chunks.push(String(chunk)));
  process.stdin.on("end", () => {
    try {
      const inputs: ConvertInput[] = JSON.parse(chunks.join(""));
      const outputs = inputs.map(convert);
      console.log(JSON.stringify(outputs));
    } catch (err) {
      console.error(err instanceof Error ? err.message : err);
      process.exit(1);
    }
  });
}
```

**Responsibilities:**
- Accumulates stdin data
- Parses JSON input
- Maps each input through your convert function
- Serializes output to JSON
- Handles errors and exits with proper code

## Testing Your Adapter

### Manual Testing

```bash
# Create test input
cat > test-input.json << 'EOF'
[{
  "namespace": "example",
  "id": "Person",
  "schema": {
    "type": "object",
    "properties": {
      "name": { "type": "string" },
      "email": { "type": "string", "format": "email" },
      "age": { "type": "integer", "minimum": 0 }
    },
    "required": ["name", "email"]
  }
}]
EOF

# Run adapter
cat test-input.json | bunx xschema-my-adapter
```

### Automated Testing

```typescript
import { describe, it, expect } from "bun:test";
import { convert } from "./index";

describe("MyAdapter", () => {
  it("converts basic object schema", () => {
    const result = convert({
      namespace: "test",
      id: "User",
      schema: {
        type: "object",
        properties: {
          name: { type: "string" }
        }
      }
    });

    expect(result.namespace).toBe("test");
    expect(result.id).toBe("User");
    expect(result.imports).toContain('import { MyValidator } from "my-lib"');
    expect(result.schema).toBeDefined();
    expect(result.type).toBeDefined();
  });
});
```

## Adapter Interface (Core Types)

**Location:** `typescript/packages/core/src/types.ts`

```typescript
export interface XSchemaAdapter {
  readonly __brand: "xschema-adapter";
  readonly name: string;
  readonly language: string;
}

export interface ConvertInput {
  namespace: string;
  id: string;
  schema: object;
}

export interface ConvertResult {
  namespace: string;
  id: string;
  imports: string[];           // Full import statements (no trailing semicolons)
  schema?: string;             // Generated validator code
  type?: string;               // Type expression
  validate?: string;           // Validation function (compliance testing)
  validateImports?: string[];  // Imports for validate function
}
```

These are the authoritative types for the adapter protocol.

**Import Format:**
- Each entry in `imports` and `validateImports` must be a full import statement
- No trailing semicolons (generator handles formatting)
- Examples: `import { z } from "zod"`, `import * as v from "valibot"`, `import type { Schema } from "effect"`

## Publishing Your Adapter

1. **Build the adapter:**
   ```bash
   bun run build
   ```

2. **Test the CLI:**
   ```bash
   echo '[{"namespace":"test","id":"T","schema":{"type":"object"}}]' | bunx xschema-my-adapter
   ```

3. **Publish to npm:**
   ```bash
   cd packages/my-adapter
   npm publish --access public
   ```

4. **Create a release PR** using conventional commits:
   ```bash
   git commit -m "feat: add my-adapter"
   git push
   ```

5. **Merge the PR** created by release-please

## Common Patterns

### Multiple Imports per Schema

Some schemas might need multiple imports:

```typescript
export function convert(input: ConvertInput): ConvertResult {
  const imports = [
    'import { z } from "zod"',
    'import { refine } from "zod-extensions"',
  ];

  return {
    namespace: input.namespace,
    id: input.id,
    imports,
    schema: 'z.object(...).pipe(refine(...))',
    type: `z.infer<typeof ${input.namespace}_${input.id}>`,
  };
}
```

### Conditional Imports

```typescript
export function convert(input: ConvertInput): ConvertResult {
  const imports = ['import { z } from "zod"'];
  
  // Add conditional imports based on schema features
  const schemaStr = JSON.stringify(input.schema);
  if (schemaStr.includes('format')) {
    imports.push('import { z } from "zod"'); // Already added
  }
  if (schemaStr.includes('$ref')) {
    imports.push('import { fromUrl } from "json-schema-ref-parser"');
  }

  return {
    namespace: input.namespace,
    id: input.id,
    imports,
    schema: generateSchema(input.schema),
    type: generateType(input),
  };
}
```

### Error Handling

```typescript
export function convert(input: ConvertInput): ConvertResult {
  try {
    if (!input.schema || typeof input.schema !== 'object') {
      throw new Error(`Invalid schema for ${input.id}: must be an object`);
    }
    
    const code = convertToMyLib(input.schema);
    
    return {
      namespace: input.namespace,
      id: input.id,
      imports: ['import { MyValidator } from "my-lib"'],
      schema: code,
      type: `MyType<typeof ${input.namespace}_${input.id}>`,
    };
  } catch (err) {
    throw new Error(
      `Failed to convert ${input.namespace}:${input.id}: ${err instanceof Error ? err.message : String(err)}`
    );
  }
}
```

The CLI will catch any thrown errors and display them to the user.

## Release Configuration

After creating your adapter, you must configure it for automated releases. See the [Adding a New Adapter](./RELEASING.md#adding-a-new-adapter) section in RELEASING.md for the 3 files you need to update:

1. `release-please-config.json` - add package entry
2. `.release-please-manifest.json` - add initial version
3. `.github/workflows/release-please.yml` - add output, condition, and publish step
