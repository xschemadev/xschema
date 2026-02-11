# @xschemadev/zod

Zod adapter for xschema - converts JSON Schema to Zod validators.

## JSON Schema Compliance

See [compliance/results/REPORT.md](./compliance/results/REPORT.md) for detailed test results.

## Installation

```bash
bun add @xschemadev/zod
```

## Usage

This adapter is used by xschema CLI to convert JSON Schema to Zod validators.

```jsonc
// user.ts.jsonc
{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
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

Then use the generated schemas with the [xschema client](/docs/typescript/client).

## Local Verification

Run from the adapter directory (`typescript/packages/adapters/zod/`):

```bash
# runtime compliance (requires Go CLI built at cli/)
bun run compliance

# unit tests
bun test

# typecheck
bun run typecheck

# type-fidelity harness (checks for any-leakage regressions)
bun run type-fidelity
```

## Fallback Typing Guardrails

The Zod adapter uses `z.unknown()` as a safe fallback for constructs where the validated domain is unbounded. It uses `z.any()` only where Zod's type system cannot represent the constraint statically. The goal: minimize `any` in inferred TypeScript types.

### Allowed `z.any()` fallbacks

These are known, accepted degradations where `z.any()` is required.

| Construct | Renderer function | Reason |
|-----------|-------------------|--------|
| IR `any` node | `render` (case "any") | empty schema `{}` or `true` — no constraints to express |
| empty intersection | `renderIntersection` | zero schemas = no constraints |
| all-any intersection | `renderIntersection` | all sub-schemas are `any` |
| tuple base | `renderTuple` | `z.array(z.any()).superRefine(...)` — Zod's `z.tuple()` doesn't support open tuples with rest items; positional validation is done in superRefine |
| complex const array | `renderLiteral` | `z.array(z.any()).refine(...)` — base ensures array shape, deep equality validated via `DEEP_SORTED_STRINGIFY_RUNTIME` |
| complex enum values | `renderEnum` | `z.any().refine(...)` — enum values can span multiple types (objects, arrays, primitives) |
| prototype-property objects | `renderObjectWithProtoProps` | `z.any().superRefine(...)` — `z.object()` cannot safely handle keys like `__proto__`, `constructor`, `toString` |

### Allowed `z.unknown()` usage (NOT type-degrading)

`z.unknown()` infers TypeScript `unknown`, which is safe — it forces callers to narrow before use. These are structural, not degradations.

| Construct | Renderer function | Reason |
|-----------|-------------------|--------|
| `not` | `renderNot` | negation ("everything except X") is semantically unbounded |
| `oneOf` | `renderOneOf` | runtime validates exactly-one match via superRefine |
| `conditional` (all branches) | `renderConditional` | if/then/else dispatches at runtime; single-branch conditionals have unconstrained passthrough |
| `typeGuarded` | `renderTypeGuarded` | dispatches to per-type validators; unmatched types pass through |
| bare `if` (no then/else) | `renderConditional` | no validation effect per JSON Schema spec |
| empty guards | `renderTypeGuarded` | no type dispatch = anything passes |

### Narrowed constructs that must NOT use `any`

These constructs previously used `z.any()` and were narrowed. Regressions are caught by the type-fidelity harness.

| Construct | Expected Zod output | How it narrows |
|-----------|---------------------|----------------|
| `oneOf` (2+ schemas) | `z.unknown().superRefine(...)` | `z.unknown()` → infers `unknown`, not `any` |
| `not` | `z.unknown().refine(...)` | same — `z.unknown()` instead of `z.any()` |
| `conditional` | `z.unknown().superRefine(...)` | same |
| `typeGuarded` | `z.unknown().superRefine(...)` | same |
| `const` (object) | `z.object({}).passthrough().refine(...)` | passthrough object base → not `any` |

### Disallowed patterns

| Pattern | Required behavior |
|---------|-------------------|
| unknown IR node kind | renderer `render()` must throw `Error` via exhaustive switch — never silently produce `z.any()` |
| new renderer branch producing `z.any()` without documented reason | must be added to "allowed" table above or changed to `z.unknown()` |
| `any`-regression on a probe with `expectAny: false` | type-fidelity harness exits non-zero — must be fixed before merge |

### Required test coverage for new renderer branches

- semantic wrappers (`oneOf`, `not`, `conditional`, `typeGuarded`): must have tests verifying the generated code uses `z.unknown()`, not `z.any()`
- tuple generation: must test that size constraints (`.min()`, `.max()`) are applied before `.superRefine()` — calling array methods on `ZodEffects` is a compile error
- prototype-property objects: must test that non-object values are rejected (the superRefine guard must add an issue, not silently pass)
- complex const/enum: must test nested object/array deep equality via `DEEP_SORTED_STRINGIFY_RUNTIME`

## Troubleshooting

### type-fidelity harness shows `IMPROVED`

A probe expected `any` but got a narrower type. This means the code improved beyond current expectations. Update `expectAny` from `true` to `false` in `type-probe/type-fidelity.ts` to lock in the improvement. Don't ignore it — unlocked improvements can silently regress.

### type-fidelity harness shows `FAIL`

A probe expected no-any but inferred `any`. Either:
1. a renderer change regressed the type output — fix the regression
2. the probe's schema changed — verify the new schema still warrants `expectAny: false`

### `ZodEffects` method errors

If generated code fails to compile with errors like "Property 'min' does not exist on type 'ZodEffects'", array size constraints are being applied after a `.superRefine()` or `.refine()` call. Size constraints must come before effect wrappers. See `renderArraySizeConstraints` vs `renderArrayRefinementConstraints`.

### compliance failures after renderer changes

Run compliance from the adapter directory: `bun run compliance`. Compare against the baseline in `tasks/type-fidelity-baseline/zod/compliance-summary.md`. Any regression in pass counts indicates a runtime behavior change — the renderer must preserve validation semantics.
