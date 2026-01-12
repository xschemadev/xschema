# @xschemadev/effect

Effect/Schema adapter for xschema - converts JSON Schema to Effect Schema validators.

## JSON Schema Compliance

See [compliance/results/REPORT.md](./compliance/results/REPORT.md) for detailed test results.

## Installation

```bash
bun add @xschemadev/effect
```

## Usage

This adapter is used by xschema CLI to convert JSON Schema to Effect Schema validators.

```jsonc
// user.ts.jsonc
{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  "schemas": [
    {
      "id": "User",
      "adapter": "@xschemadev/effect",
      "source": {
        "type": "object",
        "properties": {
          "name": { "type": "string" },
          "age": { "type": "integer" }
        },
        "required": ["name"]
      }
    }
  ]
}
```

Then run:

```bash
xschema generate
```

This will generate:

```typescript
import { Schema as S } from "effect";

export const User = S.Struct({
  name: S.String,
  age: S.optional(S.Number.pipe(S.int())),
});

export type User = typeof User.Type;
```
