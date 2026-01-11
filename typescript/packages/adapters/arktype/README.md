# @xschemadev/arktype

ArkType adapter for xschema - converts JSON Schema to ArkType validators.

## JSON Schema Compliance

See [compliance/results/REPORT.md](./compliance/results/REPORT.md) for detailed test results.

### Known Limitations

These features are intentionally not supported due to architectural constraints:

#### `$dynamicRef` / `$recursiveRef` (draft2019-09, draft2020-12)

Dynamic references resolve based on the evaluation path at runtime, requiring scope tracking through the entire validation tree. This is fundamentally incompatible with ArkType's static code generation approach. Implementing this would require:
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

When a schema contains circular references (e.g., `$ref: "#"`), the recursive part is handled with `type.unknown`. This means constraints like `unevaluatedProperties: false` won't be enforced on deeply nested levels of recursive structures.

```json
{
  "properties": { "x": { "$ref": "#" } },
  "unevaluatedProperties": false
}
// Deeply nested unevaluated properties may not be caught
```

**Workaround:** Limit recursion depth in your application logic, or validate recursive structures separately.

#### Relative URI `$id` Scoping

When a schema uses nested `$id` to create a new base URI scope, refs within that scope may fail to resolve correctly. The bundler currently resolves all refs against the root schema.

```json
{
  "$id": "http://example.com/root",
  "$defs": {
    "nested": {
      "$id": "nested/",
      "properties": {
        "x": { "$ref": "#/$defs/inner" }
      }
    }
  }
}
// The $ref should resolve relative to "nested/" but resolves against root
```

**Workaround:** Use absolute refs or avoid nested `$id` scoping.

#### URN Identifiers

Schemas using URN-style identifiers (e.g., `urn:uuid:...`) with nested pointer refs are not fully supported.

#### Metaschema Validation

Validating schemas against the JSON Schema metaschema is not performed at code generation time. The `definitions`/`$defs` keyword tests that expect metaschema validation will not catch invalid definition schemas.

#### Draft Version Detection

The adapter detects the JSON Schema draft version from the `$schema` keyword. If `$schema` is not present, it defaults to **draft2020-12 behavior**. This affects:
- `$ref` sibling handling (draft4-7: siblings ignored; draft2020-12: siblings combined)

**Recommendation:** Always include `$schema` in your schemas for consistent behavior.

#### draft3-specific Features

The following draft3-only features have limited support:
- `disallow` - Supported (converted to `not` + `type`)
- `extends` - Supported (converted to `allOf`)
- `divisibleBy` - Supported (alias for `multipleOf`)
- `type` with inline schemas - Not supported (legacy quirk)
- `required: true` on properties - Supported (draft3 style)

**Recommendation:** Upgrade schemas to draft4+ for best compatibility.

## Installation

```bash
bun add @xschemadev/arktype
```

## Usage

This adapter is used by xschema CLI to convert JSON Schema to ArkType validators.

```jsonc
// user.ts.jsonc
{
  "$schema": "https://xschema.dev/schemas/ts.jsonc",
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
