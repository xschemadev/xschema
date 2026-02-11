# Pydantic Type Probe Results - Baseline

Captured: 2026-02-11

## Inferred Types (pyright)

| Probe | Inferred Type | Has Any? |
| ----- | ------------- | -------- |
| probe_allOfMixed | `TypeAdapter[ProbeAllofmixed]` | no (but `name` field is `Annotated[Any, ...]` internally) |
| probe_allOfObjectAndAdditional | `TypeAdapter[ProbeAllofobjectandadditional]` | no |
| probe_oneOfStringOrNumber | `TypeAdapter[Unknown]` | YES |
| probe_oneOfObjects | `TypeAdapter[Unknown]` | YES |
| probe_notString | `TypeAdapter[Unknown]` | YES |
| probe_notBoolean | `TypeAdapter[Unknown]` | YES |
| probe_conditionalIfThenElse | `TypeAdapter[Unknown]` | YES |
| probe_conditionalIfThen | `TypeAdapter[Unknown]` | YES |
| probe_typeGuardedObject | `TypeAdapter[ProbeTypeguardedobject]` | no |
| probe_typeGuardedArrayMinItems | `TypeAdapter[Unknown]` | YES |
| probe_tupleStringNumber | `TypeAdapter[Unknown]` | YES |
| probe_tupleStringNumberClosed | `TypeAdapter[Unknown]` | YES |
| probe_tupleMixedWithRest | `TypeAdapter[Unknown]` | YES |
| probe_constObject | `TypeAdapter[Unknown]` | YES |
| probe_constArray | `TypeAdapter[Unknown]` | YES |
| probe_enumWithObjects | `TypeAdapter[Unknown]` | YES |
| probe_enumWithArrays | `TypeAdapter[Unknown]` | YES |

## Summary

- **Total probes**: 17
- **Probes with TypeAdapter[Unknown]**: 14
- **Probes with concrete types**: 3

## Analysis by Construct

### mixed allOf (2 probes → both preserve outer type)
- When allOf contains only objects, the result is a BaseModel class → `TypeAdapter[Model]`
- However, mixed allOf fields (e.g. `name` with intersection constraint) internally use `Annotated[Any, BeforeValidator(...)]`
- The outer type is preserved but field-level type info is lost for intersected fields

### oneOf (2 probes → all `TypeAdapter[Unknown]`)
- `Annotated[Any, BeforeValidator(validator)]` pattern makes all oneOf schemas degrade
- Even simple `oneOf: [string, number]` degrades to `TypeAdapter[Unknown]`
- Best possible: `TypeAdapter[str | float]` or `TypeAdapter[Model1 | Model2]`

### not (2 probes → all `TypeAdapter[Unknown]`)
- `Annotated[Any, BeforeValidator(validator)]` pattern causes degradation
- Best possible: `TypeAdapter[Any]` (negation truly can't narrow positively)

### conditional/if-then-else (2 probes → all `TypeAdapter[Unknown]`)
- `Annotated[Any, BeforeValidator(validator)]` pattern causes degradation
- Best possible: Union of then/else types, e.g. `TypeAdapter[ModelThen | ModelElse]`

### typeGuarded (2 probes → 1 concrete, 1 Unknown)
- Simple object typeGuarded → `TypeAdapter[ProbeTypeguardedobject]` (BaseModel subclass)
- Array constraint typeGuarded → `TypeAdapter[Unknown]` (uses `Annotated[list[Any], Field(...)]`)

### tuple (3 probes → all `TypeAdapter[Unknown]`)
- `Annotated[tuple[Any, ...], BeforeValidator(validator)]` erases all positional type info
- Best possible: `TypeAdapter[tuple[str, float, ...]]` for open, `TypeAdapter[tuple[str, float]]` for closed

### complex const (2 probes → all `TypeAdapter[Unknown]`)
- `Annotated[Any, BeforeValidator(_make_const_validator(...))]` erases type
- Best possible: Could narrow to `TypeAdapter[dict[str, str | int]]` or `TypeAdapter[list[int]]`

### complex enum (2 probes → all `TypeAdapter[Unknown]`)
- `Annotated[Any, BeforeValidator(_make_enum_validator(...))]` erases type
- Best possible: Union of value types

## Notes

- pyright reports `TypeAdapter[Unknown]` rather than `TypeAdapter[Any]` for `Annotated[Any, ...]` patterns.
  This is because pyright's inference for pydantic's TypeAdapter generic parameter doesn't resolve through
  `Annotated` wrappers with `Any` base cleanly. The effect is the same: complete loss of type information.
- The `reportRedeclaration` errors in the probe fixture are expected artifacts of concatenating per-schema
  outputs (each schema embeds its own helper functions). These don't affect type inference results.
- The 3 probes that preserve types (allOfMixed, allOfObjectAndAdditional, typeGuardedObject) all render
  as simple BaseModel subclasses, which pydantic's TypeAdapter can parameterize correctly.
