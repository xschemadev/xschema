# @xschemadev/arktype

ArkType adapter for xschema - converts JSON Schema to ArkType validators.

## JSON Schema Compliance

See [compliance/results/REPORT.md](./compliance/results/REPORT.md) for detailed test results.

## Installation

```bash
bun add @xschemadev/arktype
```

## Usage

This adapter is used by xschema CLI to convert JSON Schema to ArkType validators.

```jsonc
// user.ts.jsonc
{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  "schemas": [
    {
      "id": "User",
      "adapter": "@xschemadev/arktype",
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

Then use the generated schemas with the [xschema client](/docs/typescript/client).

## ArkType Features Used

The adapter leverages ArkType's powerful features:

- **Fluent API**: `type.string`, `type.number`, `type.boolean`
- **Range syntax**: `type("number >= 0")`, `type("0 < number <= 100")`
- **Object types**: `type({ key: schema, "optionalKey?": schema })`
- **Array types**: `schema.array()`
- **Union/Intersection**: `.or()` and `.and()` methods
- **Enumerated values**: `type.enumerated(val1, val2, ...)`
- **Custom validation**: `.narrow()` for complex constraints
- **Format keywords**: `string.email`, `string.uuid`, `string.url`, etc.
