# @xschemadev/typebox

TypeBox adapter for xschema - converts JSON Schema to TypeBox validators.

## JSON Schema Compliance

See [compliance/results/REPORT.md](./compliance/results/REPORT.md) for detailed test results.

### Known Limitations

These features are intentionally not supported due to architectural constraints:

#### `$dynamicRef` / `$recursiveRef` (draft2019-09, draft2020-12)

Dynamic references resolve based on the evaluation path at runtime, requiring scope tracking through the entire validation tree. This is fundamentally incompatible with TypeBox's static code generation approach. Implementing this would require:
- Runtime reference resolution instead of compile-time
- Tracking the dynamic scope during validation
- Significant performance overhead

**Workaround:** Flatten recursive schemas or use explicit `$ref` where possible.

#### `unevaluatedProperties` / `unevaluatedItems` with Applicators (draft2019-09+)

Simple cases (no applicators) are supported - `unevaluatedProperties` is treated as `additionalProperties`. However, when combined with `allOf`, `anyOf`, `oneOf`, `if/then/else`, or `$ref`, full tracking of "evaluated" properties across branches is not supported.

This would require:
- Passing evaluation state through all schema branches
- Collecting evaluated paths from each applicator
- Computing the complement set at runtime

**Workaround:** Use `additionalProperties: false` at the appropriate schema level, or restructure schemas to avoid needing cross-applicator tracking.

#### `additionalProperties` with Applicators (all drafts)

Per JSON Schema spec, `additionalProperties` only considers properties defined at the same schema level - it does NOT look into `allOf`/`anyOf`/`oneOf` branches. This is **correct behavior per spec** (applicators evaluate independently).

If you need properties from applicators to be considered "known", use `unevaluatedProperties` instead (simple cases without nested applicators are supported).

#### `vocabulary` (draft2019-09+)

Custom vocabularies that disable the validation vocabulary are not supported. We always apply validation keywords.

#### JavaScript Object Property Names

Properties with names that collide with JavaScript Object prototype methods (`__proto__`, `toString`, `constructor`) may not validate correctly due to edge cases in property enumeration.

**Workaround:** Avoid using JavaScript reserved property names in your schemas.

#### Circular `$ref` with Constraints

When a schema contains circular references (e.g., `$ref: "#"`), the recursive part is handled with `Type.Any()`. This means constraints like `unevaluatedProperties: false` won't be enforced on deeply nested levels of recursive structures.

```json
{
  "properties": { "x": { "$ref": "#" } },
  "unevaluatedProperties": false
}
// Deeply nested unevaluated properties may not be caught
```

**Workaround:** Limit recursion depth in your application logic, or validate recursive structures separately.

#### Draft Version Detection

The adapter detects the JSON Schema draft version from the `$schema` keyword. If `$schema` is not present, it defaults to **draft2020-12 behavior**. This affects:
- `$ref` sibling handling (draft4-7: siblings ignored; draft2020-12: siblings combined)

**Recommendation:** Always include `$schema` in your schemas for consistent behavior.

#### draft3-specific Features

The following draft3-only features have limited support:
- `disallow` - ✅ Supported (converted to `not` + `type`)
- `extends` - ✅ Supported (converted to `allOf`)
- `divisibleBy` - ✅ Supported (alias for `multipleOf`)
- `type` with inline schemas - ❌ Not supported (legacy quirk)
- `required: true` on properties - ✅ Supported (draft3 style)

**Recommendation:** Upgrade schemas to draft4+ for best compatibility.

#### TypeBox-Specific Considerations

TypeBox was designed with JSON Schema compatibility in mind, making it exceptionally well-suited for xschema:

- **Native JSON Schema output**: TypeBox's `Type.*` constructors generate valid JSON Schema objects directly. The generated code validates using Ajv for full JSON Schema compliance.

- **Type extraction**: TypeBox uses `Static<typeof schema>` for compile-time type inference, providing full TypeScript type safety from your JSON Schemas.

- **Constraint options**: Most JSON Schema constraints map directly to TypeBox options (e.g., `Type.String({ minLength: 1, maxLength: 100, pattern: "^[a-z]+$" })`).

- **Integer type**: Uses `Type.Integer()` instead of `Type.Number()` when the `integer` keyword is present.

- **Type.Unsafe for advanced features**: Complex JSON Schema features like `oneOf`, `if/then/else`, and `prefixItems` use `Type.Unsafe<T>({...})` to generate raw JSON Schema that Ajv validates correctly.

- **Validation runtime**: Uses Ajv (`ajv/dist/2020`) for validation to support all draft-2020-12 features. TypeBox's built-in `Value.Check()` doesn't support all JSON Schema keywords.

## Installation

```bash
bun add @xschemadev/typebox
```

## Usage

This adapter is used by xschema CLI to convert JSON Schema to TypeBox validators.

```jsonc
// user.ts.jsonc
{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  "schemas": [
    {
      "id": "User",
      "adapter": "@xschemadev/typebox",
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
import { Type, type Static } from "@sinclair/typebox";
import { Value } from "@sinclair/typebox/value";

export const User = Type.Object({
  name: Type.String(),
  age: Type.Optional(Type.Integer()),
});

export type User = Static<typeof User>;
```
