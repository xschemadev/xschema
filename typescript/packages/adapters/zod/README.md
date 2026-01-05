# @xschemadev/zod

Zod adapter for xschema - converts JSON Schema to Zod validators.

## JSON Schema Compliance

See [compliance/results/REPORT.md](./compliance/results/REPORT.md) for detailed results.

## Installation

```bash
bun add @xschemadev/zod
```

## Usage

This adapter is used by xschema CLI to convert JSON Schema to Zod validators.

```jsonc
// user.ts.jsonc
{
  "$schema": "https://xschema.dev/schemas/ts.jsonc",
  "schemas": [
    {
      "id": "User",
      "adapter": "@xschemadev/zod",
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
