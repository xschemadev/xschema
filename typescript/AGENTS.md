# TypeScript Agent Guidelines

TypeScript packages for xschema - core IR, adapters, and runtime client.

## Quality Requirements

Before submitting changes, ALL of these must pass:

```bash
bun run build                           # must compile without errors
bun run typecheck                       # must pass with no type errors
```

### Adapter Changes

If you modified an adapter (zod, arktype), also run compliance tests from the adapter directory:

```bash
cd packages/adapters/zod && bun run compliance    # test zod adapter
cd packages/adapters/arktype && bun run compliance # test arktype adapter
```

**Always run these checks after making changes.** Do not submit code that fails any of these.

## Build/Test Commands

```bash
bun install                             # install dependencies
bun run build                           # build all packages (core first)
bun run typecheck                       # type check all packages
bunx tsc --noEmit                       # type check single package (run in pkg dir)
```

### Single Package Commands

```bash
cd packages/core && bunx tsc --noEmit   # typecheck core only
cd packages/adapters/zod && bun run build  # build zod adapter
```

## Package Structure

```
packages/
  core/           # @xschemadev/core - IR types, JSON Schema parser, utils
    src/
      ir/         # intermediate representation types
      parser/     # JSON Schema -> IR parser
      schema/     # JSON Schema types and normalization
      utils/      # code builder, string helpers
      cli.ts      # adapter CLI helper
      types.ts    # adapter protocol types

  adapters/
    zod/          # @xschemadev/zod - Zod adapter
    arktype/      # @xschemadev/arktype - ArkType adapter

  client/         # @xschemadev/client - runtime validation client

example/          # example project using generated code
```

## Code Style

### Imports

External packages first, then relative (blank line between):

```typescript
import type { ConvertInput, ConvertResult, JSONSchema } from "@xschemadev/core";
import { parse } from "@xschemadev/core";

import { render } from "./renderer.js";
import type { RenderContext } from "./types.js";
```

**Always use `.js` extension** in relative imports (ESM requirement).

### Type Imports

Prefer `import type` for type-only imports:

```typescript
import type { SchemaNode, NodeKind } from "./ir/index.js";
import { parse, normalizeSchema } from "./parser/index.js";
```

### Exports

Use named exports. Re-export types with `export type`:

```typescript
// Re-export types
export type { JSONSchema, JSONSchemaVersion } from "./schema/index.js";

// Re-export values
export { parse, normalizeSchema } from "./parser/index.js";
```

### Naming

- Types/Interfaces: `PascalCase` (e.g., `SchemaNode`, `ConvertResult`)
- Functions/Variables: `camelCase` (e.g., `parseSchema`, `renderNode`)
- Constants: `SCREAMING_SNAKE_CASE` for true constants, `camelCase` otherwise

### Types vs Interfaces

- Use `interface` for object shapes that may be extended
- Use `type` for unions, aliases, and computed types

```typescript
// Interface for object shapes
interface ConvertInput {
  namespace: string;
  id: string;
  schema: unknown;
}

// Type for unions/aliases
type NodeKind = "string" | "number" | "boolean" | "null";
type SchemaNode = StringNode | NumberNode | BooleanNode;
```

## Adapter Development

### Adapter Protocol

Adapters implement `convert(input: ConvertInput): ConvertResult`:

```typescript
import type { ConvertInput, ConvertResult, JSONSchema } from "@xschemadev/core";
import { parse } from "@xschemadev/core";

export function convert(input: ConvertInput): ConvertResult {
  const { namespace, id, schema } = input;
  const ir = parse(schema as JSONSchema);
  const schemaCode = render(ir);
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

**Import format:** Full import statements without trailing semicolons. The generator merges and deduplicates imports.

**validate field:** Optional function for compliance testing. Takes `data`, returns `boolean`, can reference `schema` variable. Empty = type-only adapter.

### CLI Entry Point

Use `createAdapterCLI` from core:

```typescript
import { createAdapterCLI } from "@xschemadev/core";
import { convert } from "./index.js";

createAdapterCLI(convert);
```

### Package.json

```json
{
  "type": "module",
  "exports": {
    ".": {
      "types": "./dist/index.d.ts",
      "import": "./dist/index.js"
    }
  },
  "bin": {
    "xschema-zod": "./dist/cli.js"
  },
  "dependencies": {
    "@xschemadev/core": "workspace:*"
  },
  "peerDependencies": {
    "zod": "^3.0.0"
  }
}
```

## IR Types

Key types in `@xschemadev/core`:

```typescript
type NodeKind =
  | "string" | "number" | "boolean" | "null"
  | "object" | "array" | "tuple"
  | "union" | "intersection" | "oneOf"
  | "literal" | "enum" | "any" | "never"
  | "ref" | "conditional" | "nullable";

interface StringNode {
  kind: "string";
  constraints?: StringConstraints;
}

interface ObjectNode {
  kind: "object";
  properties: PropertyDef[];
  additionalProperties?: SchemaNode | boolean;
  required: string[];
}
```

## Utilities

### Code Builder

```typescript
import { chain, buildUnion, buildLiteral } from "@xschemadev/core";

const code = chain("z.string()", ".min(1)", ".max(100)");
const union = buildUnion(["z.string()", "z.number()"]);
const literal = buildLiteral("hello");  // '"hello"'
```

### JSON Pointer

```typescript
import { resolveJsonPointer, getRefName } from "@xschemadev/core";

const name = getRefName("#/definitions/User");  // "User"
```

## Gotchas

- Build core before other packages: `bun run build` handles this
- Always use `.js` extension in imports
- `workspace:*` for internal dependencies
- Adapter CLI reads from stdin, writes to stdout
- IR parser normalizes JSON Schema versions (draft-04 to 2020-12)
