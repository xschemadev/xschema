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

## Local Verification

Run from the adapter directory (`typescript/packages/adapters/arktype/`):

```bash
# runtime compliance (requires Go CLI built at cli/)
bun run compliance

# unit tests
bun test

# typecheck
bun run typecheck

# type-fidelity harness (checks for unknown-leakage regressions)
bun run type-fidelity
```

## Fallback Typing Guardrails

The ArkType adapter uses `type.unknown` as a safe fallback for constructs where the validated domain is unbounded. It uses typed bases (`.or()` chains, `type.object`, `type.number.array()`, etc.) wherever the schema provides enough positive type information. The goal: minimize `unknown` in inferred TypeScript types.

### Allowed `type.unknown` fallbacks

These are known, accepted degradations where `type.unknown` is required.

| Construct | Renderer function | Reason |
|-----------|-------------------|--------|
| IR `any` node | `render` (case "any") | empty schema `{}` or `true` — no constraints to express |
| empty intersection | `renderIntersection` | zero schemas = no constraints |
| all-any intersection | `renderIntersection` | all sub-schemas are `any` |
| `not` | `renderNot` | negation ("everything except X") is semantically unbounded — no positive type info |
| `conditional` (if/then only) | `renderConditional` | when condition doesn't match, value passes through unconstrained |
| `conditional` (if/else only) | `renderConditional` | same — unconstrained passthrough on condition match |
| `conditional` (bare if) | `renderConditional` | no then/else = no validation effect per JSON Schema spec |
| `typeGuarded` | `renderTypeGuarded` | dispatches to per-type validators; unmatched types pass through |
| empty guards | `renderTypeGuarded` | no type dispatch = anything passes |
| tuple (all variants) | `renderTuple` | `type.unknown.array().narrow(...)` — ArkType's `.array()` validates elements at runtime, so a typed base like `type.string.or(type.number).array()` rejects valid heterogeneous items that the `.narrow()` would accept |
| `jsonValueBaseType` fallback | `jsonValueBaseType` helper | when array elements contain nested objects/arrays or the array is empty, can't build a flat type union |
| `enumBaseType` fallback | `enumBaseType` helper | when no primitive types are detected across enum values |

### Narrowed constructs that must NOT use `type.unknown`

These constructs previously used `type.unknown` and were narrowed to typed bases. Regressions are caught by the type-fidelity harness.

| Construct | Expected ArkType output | How it narrows |
|-----------|-------------------------|----------------|
| `oneOf` (2+ schemas) | `schema1.or(schema2).narrow(...)` | `.or()` chain builds typed union base |
| `conditional` (if/then/else) | `thenSchema.or(elseSchema).narrow(...)` | union of both branches as typed base |
| `const` (object) | `type.object.narrow(...)` | `type.object` instead of `type.unknown` |
| `const` (array) | `jsonValueBaseType(arr).array().narrow(...)` | inspects element types to build narrow base |
| `enum` (complex values) | `enumBaseType(values).narrow(...)` | computes union from JS types of enum values |

### Disallowed patterns

| Pattern | Required behavior |
|---------|-------------------|
| unknown IR node kind | `render()` has exhaustive switch with `never` default — never silently produces `type.unknown` |
| new renderer branch producing `type.unknown` without documented reason | must be added to "allowed" table above or changed to a typed base |
| `unknown`-regression on a probe with `expectUnknown: false` | type-fidelity harness exits non-zero — must be fixed before merge |

### Required test coverage for new renderer branches

- **combinators** (`oneOf`, `not`, `conditional`, `typeGuarded`): must test that generated code uses typed `.or()` bases where applicable, and `type.unknown` only where documented
- **tuple generation**: must test that `.narrow()` validates positional items correctly — ArkType's `.array()` enforces element types at runtime, so tuple bases must stay `type.unknown.array()`
- **complex const/enum**: must test nested object/array deep equality via `DEEP_SORTED_STRINGIFY_RUNTIME` — key order must not affect equality
- **prototype-property objects**: must test that `type.object.narrow()` with `Object.hasOwn` checks rejects non-objects and validates prototype-named keys correctly

## Troubleshooting

### type-fidelity harness shows `IMPROVED`

A probe expected `unknown` but got a narrower type. This means the code improved beyond current expectations. Update `expectUnknown` from `true` to `false` in `type-probe/type-fidelity.ts` to lock in the improvement. Don't ignore it — unlocked improvements can silently regress.

### type-fidelity harness shows `FAIL`

A probe expected no-unknown but inferred `unknown`. Either:
1. a renderer change regressed the type output — fix the regression
2. the probe's schema changed — verify the new schema still warrants `expectUnknown: false`

### ArkType `.array()` validates elements at runtime

Unlike Zod's `z.array(z.any())` which accepts anything, ArkType's `type.string.array()` rejects non-string elements. This is why tuples use `type.unknown.array()` as the base — a typed base would reject valid heterogeneous items. If you attempt to narrow the tuple base, expect compliance failures in `items`, `prefixItems`, and `additionalItems` keywords.

### compliance failures after renderer changes

Run compliance from the adapter directory: `bun run compliance`. Compare against the baseline in `tasks/type-fidelity-baseline/arktype/compliance-summary.md`. Any regression in pass counts indicates a runtime behavior change — the renderer must preserve validation semantics.
