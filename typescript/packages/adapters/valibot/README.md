# @xschemadev/valibot

Valibot adapter for xschema - converts JSON Schema to Valibot validators.

## JSON Schema Compliance

See [compliance/results/REPORT.md](./compliance/results/REPORT.md) for detailed test results.

### Known Limitations

#### JavaScript Prototype Property Names (Valibot Bug)

Properties named `__proto__`, `constructor`, or `toString` cause runtime errors in Valibot's object validation. This is a known bug in Valibot's internal property lookup mechanism.

```json
{
  "properties": {
    "__proto__": { "type": "string" },
    "constructor": { "type": "string" }
  }
}
// Error: this.entries[key]._run is not a function
```

**Workaround:** Avoid using JS prototype property names in schemas, or use `additionalProperties` with pattern validation instead.

## Installation

```bash
bun add @xschemadev/valibot
```

## Usage

This adapter is used by xschema CLI to convert JSON Schema to Valibot validators.

```jsonc
// user.ts.jsonc
{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  "schemas": [
    {
      "id": "User",
      "adapter": "@xschemadev/valibot",
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
import * as v from "valibot";

export const User = v.object({
  name: v.string(),
  age: v.optional(v.pipe(v.number(), v.integer())),
});

export type User = v.InferOutput<typeof User>;
```
