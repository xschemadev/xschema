# @xschemadev/zod

Zod adapter for xschema - converts JSON Schema to Zod validators.

## JSON Schema Compliance

![draft2020-12](https://img.shields.io/badge/draft2020--12-70.9%25-red)
![draft2019-09](https://img.shields.io/badge/draft2019--09-71.3%25-red)
![draft7](https://img.shields.io/badge/draft7-75.1%25-red)
![draft6](https://img.shields.io/badge/draft6-74.0%25-red)
![draft4](https://img.shields.io/badge/draft4-73.4%25-red)
![draft3](https://img.shields.io/badge/draft3-76.0%25-red)

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
