# XSchema Adapters

Adapters convert JSON Schema to native validation code for a specific library.

## Adapter Contract

### Input (stdin)

The CLI sends a JSON array of schemas:

```json
[
  {
    "name": "User",
    "schema": {
      "type": "object",
      "properties": {
        "id": { "type": "string", "format": "uuid" },
        "name": { "type": "string", "minLength": 1 }
      },
      "required": ["id", "name"]
    }
  },
  {
    "name": "Post",
    "schema": { ... }
  }
]
```

| Field | Type | Description |
|-------|------|-------------|
| `name` | `string` | Schema name (used for export name) |
| `schema` | `object` | JSON Schema object |

### Output (stdout)

The adapter must output a JSON array:

```json
[
  {
    "name": "User",
    "schema": "z.object({ id: z.string().uuid(), name: z.string().min(1) })",
    "type": "z.infer<typeof User>",
    "imports": ["import { z } from \"zod\""]
  }
]
```

| Field | Type | Description |
|-------|------|-------------|
| `name` | `string` | Schema name (passed through) |
| `schema` | `string` | Generated validation code (runtime). Empty if type-only adapter. |
| `type` | `string` | Generated type expression (compile-time). Empty if schema-only adapter. |
| `imports` | `string[]` | Required import statements |

### Output Modes

Adapters can return:

1. **Both** - `schema` + `type` (e.g., Zod)
2. **Schema only** - `schema` only, empty `type` (e.g., Ajv)
3. **Type only** - `type` only, empty `schema` (e.g., pure type generators)

## Adapter Structure

```
adapters/
├── typescript/
│   └── zod/
│       ├── index.ts      # Exports adapter identifier + convert function
│       ├── cli.ts        # CLI entry point (reads stdin, writes stdout)
│       └── package.json  # bin: { "@xschema/zod": "./cli.ts" }
│
└── python/
    └── pydantic/
        ├── __init__.py
        ├── cli.py
        └── pyproject.toml
```

## TypeScript Adapter Example

```typescript
// index.ts
export interface ConvertInput {
  name: string;
  schema: object;
}

export interface ConvertResult {
  name: string;
  schema: string;
  type: string;
  imports: string[];
}

export function convert(name: string, schema: object): ConvertResult {
  // Convert JSON Schema to native code
  return {
    name,
    schema: "z.object({ ... })",
    type: "z.infer<typeof " + name + ">",
    imports: ['import { z } from "zod"'],
  };
}
```

```typescript
// cli.ts
#!/usr/bin/env node
import { convert, type ConvertInput } from "./index.ts";

const chunks: string[] = [];
process.stdin.on("data", (chunk) => chunks.push(String(chunk)));
process.stdin.on("end", () => {
  const inputs: ConvertInput[] = JSON.parse(chunks.join(""));
  const outputs = inputs.map(({ name, schema }) => convert(name, schema));
  console.log(JSON.stringify(outputs));
});
```

## Python Adapter Example

```python
# cli.py
import json
import sys

def convert(name: str, schema: dict) -> dict:
    # Convert JSON Schema to Pydantic
    return {
        "name": name,
        "schema": f"class {name}(BaseModel): ...",
        "type": name,
        "imports": ["from pydantic import BaseModel"],
    }

if __name__ == "__main__":
    inputs = json.load(sys.stdin)
    outputs = [convert(i["name"], i["schema"]) for i in inputs]
    json.dump(outputs, sys.stdout)
```
