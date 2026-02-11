# Zod Type Probe Results - Baseline

Captured: 2026-02-11

## Inferred Types

| Probe | Inferred Type | Has any? |
| ----- | ------------- | -------- |
| Probe_oneOfStringOrNumber | `any` | YES |
| Probe_oneOfObjects | `any` | YES |
| Probe_notString | `any` | YES |
| Probe_notBoolean | `any` | YES |
| Probe_conditionalIfThenElse | `any` | YES |
| Probe_conditionalIfThen | `any` | YES |
| Probe_typeGuardedObject | `any` | YES |
| Probe_typeGuardedArrayMinItems | `any` | YES |
| Probe_tupleStringNumber | `any[]` | YES |
| Probe_tupleStringNumberClosed | `any[]` | YES |
| Probe_tupleMixedWithRest | `any[]` | YES |
| Probe_constObject | `objectOutputType<{}, ZodTypeAny, "passthrough">` | no |
| Probe_constArray | `any[]` | YES |
| Probe_enumWithObjects | `any` | YES |
| Probe_enumWithArrays | `any` | YES |

## Summary

- **Total probes**: 15
- **Probes with any/any[]**: 14
- **Probes without any**: 1

## Analysis by Construct

### oneOf (2 probes → all `any`)
- `z.any().superRefine(...)` pattern makes all oneOf schemas infer `any`
- Even simple `oneOf: [string, number]` degrades to `any`

### not (2 probes → all `any`)
- `z.any().refine(...)` pattern causes `any`
- Best possible: `unknown` (not can't narrow positively)

### conditional/if-then-else (2 probes → all `any`)
- `z.any().superRefine(...)` pattern causes `any`
- Best possible: `unknown` (conditional can't narrow positively)

### typeGuarded (2 probes → all `any`)
- `z.any().superRefine(...)` pattern causes `any`
- Best possible: `unknown` (type guard is conditional)

### tuple (3 probes → all `any[]`)
- `z.array(z.any()).superRefine(...)` pattern causes `any[]`
- Best possible: `[string, number]` for closed, `[string, number, ...any[]]` for open
- This is a significant degradation - tuples should have positional type info

### complex const (2 probes → 1 `any[]`, 1 broad object)
- Array const: `z.array(z.any()).refine(...)` → `any[]`
- Object const: `z.object({}).passthrough().refine(...)` → `objectOutputType<{}, ZodTypeAny, "passthrough">`
- Object is technically not `any` but still overly broad (allows any keys/values)

### complex enum (2 probes → all `any`)
- `z.any().refine(...)` pattern for enums with non-primitive values → `any`
- Best possible: union of literal types

## Compile-Time Issues

The probe fixture compiles **without errors** under `tsc --noEmit --strict`.
The `z.any()` pattern is type-safe in that it doesn't cause compilation failures.
The problem is purely about type inference quality - `any` propagates unsafely
through downstream code that uses these schemas.
