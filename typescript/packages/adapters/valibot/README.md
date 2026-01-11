# @xschemadev/valibot

Valibot adapter for xschema - converts JSON Schema to Valibot validators.

## JSON Schema Compliance

See [compliance/results/REPORT.md](./compliance/results/REPORT.md) for detailed test results.

### Known Limitations

These features are intentionally not supported due to architectural constraints:

#### `$dynamicRef` / `$recursiveRef` (draft2019-09, draft2020-12)

Dynamic references resolve based on the evaluation path at runtime, requiring scope tracking through the entire validation tree. This is fundamentally incompatible with Valibot's static code generation approach. Implementing this would require:
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

#### Circular `$ref` with Constraints

When a schema contains circular references (e.g., `$ref: "#"`), the recursive part is handled with `v.lazy(() => v.any())`. This means constraints like `unevaluatedProperties: false` won't be enforced on deeply nested levels of recursive structures.

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

#### Valibot-Specific Considerations

Valibot's functional API differs from Zod's chaining approach in several ways:

- **Object validation modes**: Valibot uses different functions for different additionalProperties modes (`v.strictObject`, `v.looseObject`, `v.objectWithRest`) instead of chaining methods. This provides clearer semantics but can be more verbose.

- **Pipe validation order**: All validation actions must be collected and applied in a single `v.pipe()` call. Chaining multiple `.pipe()` calls doesn't work as expected.

- **Exclusive bounds**: Valibot v0.42 doesn't have native support for `exclusiveMinimum`/`exclusiveMaximum` options, so these are implemented using `v.check()` with custom validation logic.

- **Pattern properties**: Valibot has no native support for pattern properties, so these are implemented using custom `v.check()` validation that tests keys against regex patterns.

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
