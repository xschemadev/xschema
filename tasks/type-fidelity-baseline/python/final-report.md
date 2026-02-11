# Pydantic Adapter - Final Type Fidelity Report

## Runtime Compliance: Before vs After

| Draft | Before | After | Diff |
| ----- | ------ | ----- | ---- |
| draft3 | 407 passed (100.0%) | 407 passed (100.0%) | unchanged |
| draft4 | 566 passed (100.0%) | 566 passed (100.0%) | unchanged |
| draft6 | 760 passed (100.0%) | 760 passed (100.0%) | unchanged |
| draft7 | 836 passed (100.0%) | 836 passed (100.0%) | unchanged |
| draft2019-09 | 936 passed, 178 unsupported (100.0%) | 936 passed, 178 unsupported (100.0%) | unchanged |
| draft2020-12 | 950 passed, 201 unsupported (100.0%) | 950 passed, 201 unsupported (100.0%) | unchanged |

**Zero runtime regressions.** All 6 drafts remain at 100% compliance.

## Type Fidelity: Before vs After

| Probe | Before (type) | Before (any?) | After (type) | After (any?) | Status |
| ----- | ------------- | ------------- | ------------ | ------------ | ------ |
| allOfMixed | `TypeAdapter[ProbeAllofmixed]` | no | `ProbeAllofmixed` | no | unchanged |
| allOfObjectAndAdditional | `TypeAdapter[ProbeAllofobjectandadditional]` | no | `ProbeAllofobjectandadditional` | no | unchanged |
| oneOfStringOrNumber | `TypeAdapter[Unknown]` | YES | `Annotated[StrictStr \| StrictFloat, ...]` | no | **improved** |
| oneOfObjects | `TypeAdapter[Unknown]` | YES | `Annotated[ProbeOneofobjectsOption0 \| ProbeOneofobjectsOption1, ...]` | no | **improved** |
| notString | `TypeAdapter[Unknown]` | YES | `Annotated[Any, ...]` | YES | accepted |
| notBoolean | `TypeAdapter[Unknown]` | YES | `Annotated[Any, ...]` | YES | accepted |
| conditionalIfThenElse | `TypeAdapter[Unknown]` | YES | `Annotated[...Then \| ...Else, ...]` | no | **improved** |
| conditionalIfThen | `TypeAdapter[Unknown]` | YES | `Annotated[Any, ...]` | YES | accepted |
| typeGuardedObject | `TypeAdapter[ProbeTypeguardedobject]` | no | `ProbeTypeguardedobject` | no | unchanged |
| typeGuardedArrayMinItems | `TypeAdapter[Unknown]` | YES | `Annotated[list[Any], Field(min_length=1)]` | YES | accepted |
| tupleStringNumber | `TypeAdapter[Unknown]` | YES | `Annotated[tuple[Any, ...], ...]` | YES | accepted |
| tupleStringNumberClosed | `TypeAdapter[Unknown]` | YES | `Annotated[tuple[StrictStr \| StrictFloat, ...], ...]` | no | **improved** |
| tupleMixedWithRest | `TypeAdapter[Unknown]` | YES | `Annotated[tuple[StrictStr \| StrictFloat \| StrictBool, ...], ...]` | no | **improved** |
| constObject | `TypeAdapter[Unknown]` | YES | `Annotated[dict, ...]` | no | **improved** |
| constArray | `TypeAdapter[Unknown]` | YES | `Annotated[list, ...]` | no | **improved** |
| enumWithObjects | `TypeAdapter[Unknown]` | YES | `Annotated[dict \| str, ...]` | no | **improved** |
| enumWithArrays | `TypeAdapter[Unknown]` | YES | `Annotated[list \| None, ...]` | no | **improved** |

## Summary

| Metric | Before | After | Change |
| ------ | ------ | ----- | ------ |
| Runtime compliance | 100% (all 6 drafts) | 100% (all 6 drafts) | unchanged |
| Probes any-free | 3 / 17 | 12 / 17 | **+9 improved** |
| Probes with Any | 14 / 17 | 5 / 17 | **-9 reduced** |

### Improvements by construct

- **oneOf**: both probes narrowed from `Any` to typed unions (`StrictStr | StrictFloat`, model unions)
- **conditional (if/then/else)**: narrowed to `Then | Else` union when both branches present
- **tuple (closed)**: narrowed from `tuple[Any, ...]` to `tuple[StrictStr | StrictFloat, ...]`
- **tuple (with rest)**: narrowed from `tuple[Any, ...]` to typed homogeneous union
- **const (object/array)**: narrowed from `Any` to `dict` / `list`
- **enum (complex)**: narrowed from `Any` to union of value types (`dict | str`, `list | None`)

### Accepted Any (5 probes — semantically correct)

- **not** (2): negation is unbounded — "anything except X" can't be represented as a union
- **conditional (if/then only)**: when only `then` exists, unmatched inputs pass through unconstrained
- **typeGuarded (array)**: type guard with array constraint uses `list[Any]` — element types unknown
- **tuple (open)**: open tuple with heterogeneous prefix items can't safely narrow element union
