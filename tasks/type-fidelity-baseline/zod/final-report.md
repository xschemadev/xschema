# Zod Adapter - Final Non-Regression Report

## Runtime Compliance: No Regressions

| Draft | Baseline | Final | Delta |
| ----- | -------- | ----- | ----- |
| draft3 | 397 passed (100.0%) | 397 passed (100.0%) | unchanged |
| draft4 | 526 passed (100.0%) | 526 passed (100.0%) | unchanged |
| draft6 | 707 passed (100.0%) | 707 passed (100.0%) | unchanged |
| draft7 | 783 passed (100.0%) | 783 passed (100.0%) | unchanged |
| draft2019-09 | 906 passed, 178 unsupported (100.0%) | 906 passed, 178 unsupported (100.0%) | unchanged |
| draft2020-12 | 919 passed, 201 unsupported (100.0%) | 919 passed, 201 unsupported (100.0%) | unchanged |

All 6 drafts remain at 100% runtime compliance. Zero regressions.

## Type-Fidelity: 8 Probes Improved

| Probe | Baseline Type | Final Type | Improved? |
| ----- | ------------- | ---------- | --------- |
| oneOfStringOrNumber | `any` | `unknown` | YES |
| oneOfObjects | `any` | `unknown` | YES |
| notString | `any` | `unknown` | YES |
| notBoolean | `any` | `unknown` | YES |
| conditionalIfThenElse | `any` | `unknown` | YES |
| conditionalIfThen | `any` | `unknown` | YES |
| typeGuardedObject | `any` | `unknown` | YES |
| typeGuardedArrayMinItems | `any` | `unknown` | YES |
| tupleStringNumber | `any[]` | `any[]` | no |
| tupleStringNumberClosed | `any[]` | `any[]` | no |
| tupleMixedWithRest | `any[]` | `any[]` | no |
| constObject | broad object | broad object | no (already any-free) |
| constArray | `any[]` | `any[]` | no |
| enumWithObjects | `any` | `any` | no |
| enumWithArrays | `any` | `any` | no |

**Summary**: 9 any-free (was 1), 6 any (was 14). 8 probes improved from `any` to `unknown`.

## Changes That Drove Improvements

### Task 11: z.any() → z.unknown() in semantic wrappers
- `renderOneOf`, `renderNot`, `renderConditional`, `renderTypeGuarded` all switched from `z.any()` to `z.unknown()`
- Runtime behavior identical (both accept any value), but `z.unknown()` infers `unknown` instead of `any`
- This is the single change responsible for all 8 improved probes

### Task 12: Tuple/array constraint chaining fix
- Split `renderArrayConstraints` so `.min()`/`.max()` go before `.superRefine()` wrappers
- Fixed TypeScript compile errors when tuple/array has size constraints + effect wrappers
- No type-fidelity change (tuples still use `z.array(z.any())`)

### Task 13: Prototype-property object validation
- `renderObjectWithProtoProps` now rejects non-objects for `type: "object"` schemas
- Correctness fix, not a type-fidelity change

### Task 14: Deep equality for complex const/enum
- Recursive key normalization for nested objects in const/enum validation
- Correctness fix, not a type-fidelity change

## Accepted `any` Remaining (6 probes)

| Probe | Type | Reason |
| ----- | ---- | ------ |
| tupleStringNumber | `any[]` | `z.array(z.any()).superRefine(...)` — positional types can't be expressed via ZodArray |
| tupleStringNumberClosed | `any[]` | same as above |
| tupleMixedWithRest | `any[]` | same as above |
| constArray | `any[]` | `z.array(z.any()).refine(...)` — array literal types can't be expressed |
| enumWithObjects | `any` | `z.any().refine(...)` — heterogeneous enum values need runtime deep comparison |
| enumWithArrays | `any` | `z.any().refine(...)` — same reason |

Tuple `any[]` is a Zod limitation: `ZodTuple` has a different API from `ZodArray` and switching would require rearchitecting tuple rendering. The runtime validation is correct via `superRefine`.

Complex enum/const `any` is also a Zod limitation: there's no Zod type that represents "one of these specific complex values" with correct inference. The runtime validation via deep-equality `refine` is correct.

## Type-Fidelity Harness

All 15 probes pass with correct expectations:
- 9 probes: `expectAny: false` — verified no `any` in inferred type
- 6 probes: `expectAny: true` — accepted `any` documented above

Run: `bun run type-fidelity` from `typescript/packages/adapters/zod/`
