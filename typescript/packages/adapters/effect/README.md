# @xschemadev/effect

Effect/Schema adapter for xschema - converts JSON Schema to Effect Schema validators.

## JSON Schema Compliance

See [compliance/results/REPORT.md](./compliance/results/REPORT.md) for detailed test results.

### Known Limitations

These features are intentionally not supported due to architectural constraints:

#### `$dynamicRef` / `$recursiveRef` (draft2019-09, draft2020-12)

Dynamic references resolve based on the evaluation path at runtime, requiring scope tracking through the entire validation tree. This is fundamentally incompatible with Effect Schema's static code generation approach. Implementing this would require:
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

When a schema contains circular references (e.g., `$ref: "#"`), the recursive part is handled with `S.Unknown`. This means constraints like `unevaluatedProperties: false` won't be enforced on deeply nested levels of recursive structures.

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

#### Effect/Schema-Specific Considerations

Effect/Schema's functional API differs from other validation libraries in several ways:

- **Schema namespace**: All Effect/Schema types are accessed via `S.` prefix (e.g., `S.String`, `S.Struct`, `S.Array`). The import is `import { Schema as S } from "effect"`.

- **Type extraction**: Effect/Schema uses `.Type` property for compile-time type inference (e.g., `typeof User.Type`), which differs from Zod's `z.infer<>` or Valibot's `v.InferOutput<>`.

- **Pipe syntax**: Validations are chained using `.pipe()` method (e.g., `S.String.pipe(S.minLength(1), S.maxLength(100))`).

- **Filter validation**: Custom validations use `S.filter()` with a predicate function. Error messages are provided via `{ message: () => "..." }` syntax.

- **Either-based validation**: `S.decodeUnknownEither(schema)(data)` returns an `Either` type with `_tag: "Right"` for success and `_tag: "Left"` for failure, following functional programming conventions.

- **Object types**: Uses `S.Struct({...})` for objects. Additional properties are handled via `S.Record({ key: S.String, value: schema })` as a second argument to `S.Struct`.

- **Union types**: `S.Union(schema1, schema2, ...)` takes schemas as individual arguments rather than an array.

- **Intersection types**: `S.extend(obj1, obj2)` for merging object schemas. General intersections use `S.filter()` to validate all schemas match.

- **Optional properties**: `S.optional(schema)` wraps property schemas for optional fields.

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
