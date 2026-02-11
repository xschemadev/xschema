# ArkType Type Probe Results - Baseline

Captured: 2026-02-11

## Inferred Types

| Probe | Inferred Type | Has unknown? |
| ----- | ------------- | ------------ |
| Probe_oneOfStringOrNumber | `unknown` | YES |
| Probe_oneOfObjects | `unknown` | YES |
| Probe_notString | `unknown` | YES |
| Probe_notBoolean | `unknown` | YES |
| Probe_conditionalIfThenElse | `unknown` | YES |
| Probe_conditionalIfThen | `unknown` | YES |
| Probe_typeGuardedObject | `unknown` | YES |
| Probe_typeGuardedArrayMinItems | `unknown` | YES |
| Probe_tupleStringNumber | `unknown[]` | YES |
| Probe_tupleStringNumberClosed | `unknown[]` | YES |
| Probe_tupleMixedWithRest | `unknown[]` | YES |
| Probe_constObject | `unknown` | YES |
| Probe_constArray | `unknown` | YES |
| Probe_enumWithObjects | `unknown` | YES |
| Probe_enumWithArrays | `unknown` | YES |

## Summary

- **Total probes**: 15
- **Probes with unknown/unknown[]**: 15
- **Probes without unknown**: 0

## Analysis by Construct

### oneOf (2 probes -> all `unknown`)
- `type.unknown.narrow(...)` pattern makes all oneOf schemas infer `unknown`
- Even simple `oneOf: [string, number]` degrades to `unknown`
- Best possible: `string | number` union type

### not (2 probes -> all `unknown`)
- `type.unknown.narrow(...)` pattern causes `unknown`
- Best possible: `unknown` (not can't narrow positively)
- This is actually acceptable - `unknown` is correct for not schemas

### conditional/if-then-else (2 probes -> all `unknown`)
- `type.unknown.narrow(...)` pattern causes `unknown`
- Best possible: `unknown` (conditional can't narrow positively in general)

### typeGuarded (2 probes -> all `unknown`)
- `type.unknown.narrow(...)` pattern causes `unknown`
- Best possible: `unknown` (type guard is conditional)

### tuple (3 probes -> all `unknown[]`)
- `type.unknown.array().narrow(...)` pattern causes `unknown[]`
- Best possible: `[string, number]` for closed, `[string, number, ...unknown[]]` for open
- This is a significant degradation - tuples should have positional type info

### complex const (2 probes -> all `unknown`)
- `type.unknown.narrow(...)` for deep equality check -> `unknown`
- Primitive const uses `type.unit()` which preserves type (not tested here as it works)
- Best possible: narrow literal type for the const value

### complex enum (2 probes -> all `unknown`)
- `type.unknown.narrow(...)` for enums with non-primitive values -> `unknown`
- Primitive-only enums use `type.enumerated()` which preserves type (not tested here)
- Best possible: union of literal types

## Compile-Time Issues

The probe fixture compiles **without errors** under `tsc --noEmit --strict`.
The `type.unknown` pattern is type-safe in that it doesn't cause compilation failures.
The problem is purely about type inference quality - `unknown` is safe but overly
conservative, requiring explicit type narrowing in downstream code.

## Comparison with Zod Baseline

| Construct | Zod Baseline | ArkType Baseline | Notes |
| --------- | ------------ | ---------------- | ----- |
| oneOf | `any` | `unknown` | ArkType is safer (unknown vs any) |
| not | `any` | `unknown` | ArkType is safer |
| conditional | `any` | `unknown` | ArkType is safer |
| typeGuarded | `any` | `unknown` | ArkType is safer |
| tuple | `any[]` | `unknown[]` | ArkType is safer |
| constObject | broad object | `unknown` | Zod slightly better (preserves object shape) |
| constArray | `any[]` | `unknown` | ArkType loses array nature |
| complexEnum | `any` | `unknown` | ArkType is safer |

ArkType consistently uses `unknown` where Zod uses `any`. This is inherently safer
since `unknown` requires explicit narrowing while `any` silently bypasses type checking.
However, both adapters lose the same type precision for these constructs.
