# ArkType Adapter - Final Non-Regression Report

## Runtime Compliance

| Draft | Baseline | Final | Status |
| ----- | -------- | ----- | ------ |
| draft3 | 397 (100.0%) | 397 (100.0%) | no change |
| draft4 | 526 (100.0%) | 526 (100.0%) | no change |
| draft6 | 707 (100.0%) | 707 (100.0%) | no change |
| draft7 | 783 (100.0%) | 783 (100.0%) | no change |
| draft2019-09 | 906 (100.0%) | 906 (100.0%) | no change |
| draft2020-12 | 919 (100.0%) | 919 (100.0%) | no change |

**Zero regressions.** Runtime behavior identical to baseline.

## Type Fidelity

| Probe | Baseline | Final | Status |
| ----- | -------- | ----- | ------ |
| oneOfStringOrNumber | `unknown` | `string \| number` | IMPROVED |
| oneOfObjects | `unknown` | `{ kind: "a"; value: string; } \| { kind: "b"; value: number; }` | IMPROVED |
| notString | `unknown` | `unknown` | unchanged |
| notBoolean | `unknown` | `unknown` | unchanged |
| conditionalIfThenElse | `unknown` | `{ kind: "a"; value: string; } \| { kind: "b"; value: number; }` | IMPROVED |
| conditionalIfThen | `unknown` | `unknown` | unchanged |
| typeGuardedObject | `unknown` | `unknown` | unchanged |
| typeGuardedArrayMinItems | `unknown` | `unknown` | unchanged |
| tupleStringNumber | `unknown[]` | `unknown[]` | unchanged |
| tupleStringNumberClosed | `unknown[]` | `unknown[]` | unchanged |
| tupleMixedWithRest | `unknown[]` | `unknown[]` | unchanged |
| constObject | `unknown` | `object` | IMPROVED |
| constArray | `unknown` | `number[]` | IMPROVED |
| enumWithObjects | `unknown` | `string \| object` | IMPROVED |
| enumWithArrays | `unknown` | `number[] \| null` | IMPROVED |

### Summary

- **Baseline**: 0 unknown-free, 15 unknown
- **Final**: 7 unknown-free, 8 unknown
- **Improved probes**: 7

### Improvements by construct

| Construct | Change | How |
| --------- | ------ | --- |
| oneOf (2) | `unknown` → typed union | `.or()` chain builds typed union base instead of `type.unknown` |
| conditional if/then/else (1) | `unknown` → typed union | `then.or(else)` as base type |
| constObject (1) | `unknown` → `object` | `type.object` for const objects |
| constArray (1) | `unknown` → `number[]` | `jsonValueBaseType(arr).array()` for const arrays |
| enumWithObjects (1) | `unknown` → `string \| object` | `enumBaseType()` computes union from JS types |
| enumWithArrays (1) | `unknown` → `number[] \| null` | `enumBaseType()` computes union from JS types |

### Accepted unknown (8)

| Probe | Reason |
| ----- | ------ |
| notString | negation is semantically unbounded |
| notBoolean | negation is semantically unbounded |
| conditionalIfThen | if/then without else has unconstrained passthrough |
| typeGuardedObject | unmatched types pass through |
| typeGuardedArrayMinItems | element types unknown at type level |
| tupleStringNumber | ArkType `.array()` validates at runtime; typed base rejects valid heterogeneous items |
| tupleStringNumberClosed | same as above |
| tupleMixedWithRest | same as above |

## Verification

- Runtime compliance: 100% all 6 drafts (no regressions)
- Type-fidelity harness: 15/15 pass, 7 unknown-free, 8 unknown
- TypeScript typecheck: all 6 packages pass
- TypeScript build: all 6 packages compile
