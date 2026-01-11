# TypeBox Adapter Discovery

## Baseline Results

| Draft | Passed | Total | Coverage |
|-------|--------|-------|----------|
| draft2020-12 | 356 | 920 | 38.7% |
| draft2019-09 | 337 | 915 | 36.8% |
| draft7 | 355 | 844 | 42.1% |
| draft6 | 299 | 768 | 38.9% |
| draft4 | 203 | 568 | 35.7% |
| draft3 | 161 | 407 | 39.6% |

## Root Cause Analysis

### Primary Issue: Value.Check vs Type.Unsafe Incompatibility

The typebox adapter generates JSON Schema constructs via `Type.Unsafe<T>({...})` for features TypeBox doesn't express natively. However, `Value.Check()` **cannot validate arbitrary JSON Schema** - it only processes TypeBox's native type kinds.

When `Value.Check()` encounters a `Type.Unsafe` schema with JSON Schema constructs like `if/then/else`, `oneOf`, `contains`, it throws:
```
error: Unknown type
```

This explains why ~60% of tests fail with the exact same error.

### What Value.Check Supports

**Fully validated** (native TypeBox types):
- `Type.String()` with `minLength`, `maxLength`, `pattern`
- `Type.Number()` / `Type.Integer()` with `minimum`, `maximum`, `exclusiveMinimum`, `exclusiveMaximum`, `multipleOf`
- `Type.Array()` with `minItems`, `maxItems`, `uniqueItems`, `contains`
- `Type.Object()` with `minProperties`, `maxProperties`, `additionalProperties`
- `Type.Tuple()`, `Type.Union()`, `Type.Intersect()`, `Type.Not()`, `Type.Literal()`, `Type.Null()`, `Type.Never()`
- `Type.Record()` for patternProperties-like behavior

**Generates JSON Schema but NOT validated by Value.Check**:
- `if/then/else` (conditional schemas)
- `oneOf` (exactly one must match)
- `dependentRequired`, `dependentSchemas`
- `propertyNames`
- `patternProperties` (as object option, different from Record)

**Format validation**: Requires manual registration in `FormatRegistry`. Unregistered formats cause validation to reject all values.

## Expressible vs Runtime-Enforced Mismatch

TypeBox has a unique situation: many JSON Schema features can be **expressed** (the schema generates correct JSON Schema output), but **not enforced at runtime** by `Value.Check()`.

| Feature | Expressible | Runtime Enforced |
|---------|-------------|------------------|
| type constraints | Yes | Yes |
| string constraints | Yes | Yes |
| number constraints | Yes | Yes |
| array length constraints | Yes | Yes |
| uniqueItems | Yes | Yes |
| contains/minContains/maxContains | Yes | Yes |
| additionalProperties | Yes | Yes |
| required | Yes | Yes |
| format | Yes | No (needs FormatRegistry) |
| oneOf | Via Unsafe | No |
| if/then/else | Via Unsafe | No |
| dependentRequired | Via Unsafe | No |
| dependentSchemas | Via Unsafe | No |
| propertyNames | Via option | No (always passes) |
| patternProperties | Via option | No (always passes) |

## Failing Tests Breakdown

### Category 1: Value.Check Doesn't Support Type.Unsafe Constructs (~500 tests)

All tests showing `error: Unknown type` fall here. This includes:
- All tests for schemas without explicit `type` (generates `Type.Unsafe` with `if/then`)
- All `oneOf` tests
- All `if/then/else` tests
- All `dependentRequired`/`dependentSchemas` tests
- All `additionalProperties` tests (when combined with other features)
- All `properties` tests
- All `required` tests
- All `pattern` tests (when no explicit type)
- etc.

### Category 2: Bundler Limitations (~40 tests)

Same as zod/valibot adapters:
- `$recursiveRef`, `$recursiveAnchor` - intentionally unsupported
- `$dynamicRef`, `$dynamicAnchor` - intentionally unsupported
- Some complex $ref/$id scoping issues

### Category 3: Features Not Enforced by Value.Check (~20 tests)

- `propertyNames` - option exists but validation always passes
- `patternProperties` - option exists but validation always passes
- `format` - requires FormatRegistry setup

## Recommended Fixes

### Fix 1: Switch Validation to Ajv (HIGH IMPACT - ~500 tests)

Replace `Value.Check(schema, data)` with Ajv-based validation:

```typescript
validate: `(() => {
  const ajv = new Ajv({ allErrors: true });
  addFormats(ajv);
  const validate = ajv.compile(schema);
  return (data: unknown) => validate(data);
})()`
```

This is the **single most impactful fix** - it would enable validation of all JSON Schema constructs that TypeBox expresses via `Type.Unsafe`.

**Pros:**
- Fixes ~500+ tests in one change
- Ajv is already a devDependency
- TypeBox schemas ARE valid JSON Schema - Ajv can compile them

**Cons:**
- Different validation semantics (Ajv is strict JSON Schema, Value.Check has TypeBox-specific behavior)
- Increases bundle size if Ajv not already present
- Need to add ajv-formats for format validation

### Fix 2: Use Native TypeBox Types Where Possible (MEDIUM IMPACT - ~50 tests)

The renderer uses `Type.Unsafe` for some cases where native TypeBox types would work:

1. **TypeGuarded nodes**: Currently generates `if/then` via Unsafe. Could use `Type.Union` with type checks.
2. **Record with constraints**: Currently uses `Type.Object` + options. Native `Type.Record` might work better.

This requires renderer changes but would make some schemas work with `Value.Check`.

### Fix 3: Register Format Validators (LOW IMPACT - ~20 tests)

Add format validators to FormatRegistry in the harness:
```typescript
FormatRegistry.Set('email', (v) => /^[^@]+@[^@]+$/.test(v));
FormatRegistry.Set('uri', (v) => /^https?:\/\//.test(v));
// etc.
```

## Expected Regressions

None expected - current baseline is 37-42% so any fix would be an improvement.

## Recommended Implementation Order

1. **US-NEW-1**: Switch validation from Value.Check to Ajv (~500 tests)
2. **US-NEW-2**: Optimize renderer to use native types where possible (~50 tests)
3. **US-NEW-3**: Register standard format validators (~20 tests)

Total expected improvement: ~37% → ~90% (similar to zod/valibot after fixes)

## Full Failing Tests List by Draft

### draft2020-12 (564 failures)

| Keyword | Failures | Category |
|---------|----------|----------|
| additionalProperties | 21 | Value.Check |
| allOf | 19 | Value.Check |
| anyOf | 8 | Value.Check |
| const | 21 | Value.Check |
| contains | 21 | Value.Check |
| default | 4 | Value.Check |
| defs | 2 | Bundler |
| dependentRequired | 20 | Value.Check |
| dependentSchemas | 20 | Value.Check |
| dynamicRef | 39 | Bundler |
| enum | 20 | Value.Check |
| exclusiveMaximum | 4 | Value.Check |
| exclusiveMinimum | 4 | Value.Check |
| if-then-else | 18 | Value.Check |
| infinite-loop-detection | 2 | Value.Check |
| items | 25 | Value.Check |
| maxContains | 10 | Value.Check |
| maxItems | 6 | Value.Check |
| maxLength | 7 | Value.Check |
| maxProperties | 10 | Value.Check |
| maximum | 8 | Value.Check |
| minContains | 26 | Value.Check |
| minItems | 6 | Value.Check |
| minLength | 7 | Value.Check |
| minProperties | 8 | Value.Check |
| minimum | 11 | Value.Check |
| multipleOf | 9 | Value.Check |
| oneOf | 20 | Value.Check |
| pattern | 9 | Value.Check |
| patternProperties | 23 | Value.Check |
| prefixItems | 16 | Value.Check |
| properties | 28 | Value.Check |
| propertyNames | 20 | Value.Check |
| ref | 8 | Bundler |
| refRemote | 4 | Bundler |
| required | 16 | Value.Check |
| unevaluatedItems | 2 | Bundler |
| unevaluatedProperties | 2 | Bundler |
| uniqueItems | 69 | Value.Check |

### draft2019-09 (578 failures)

| Keyword | Failures | Category |
|---------|----------|----------|
| additionalItems | 19 | Value.Check |
| additionalProperties | 21 | Value.Check |
| allOf | 19 | Value.Check |
| anyOf | 8 | Value.Check |
| const | 21 | Value.Check |
| contains | 21 | Value.Check |
| default | 4 | Value.Check |
| defs | 2 | Bundler |
| dependentRequired | 20 | Value.Check |
| dependentSchemas | 20 | Value.Check |
| enum | 20 | Value.Check |
| exclusiveMaximum | 4 | Value.Check |
| exclusiveMinimum | 4 | Value.Check |
| if-then-else | 18 | Value.Check |
| infinite-loop-detection | 2 | Value.Check |
| items | 25 | Value.Check |
| maxContains | 10 | Value.Check |
| maxItems | 6 | Value.Check |
| maxLength | 7 | Value.Check |
| maxProperties | 10 | Value.Check |
| maximum | 8 | Value.Check |
| minContains | 26 | Value.Check |
| minItems | 6 | Value.Check |
| minLength | 7 | Value.Check |
| minProperties | 8 | Value.Check |
| minimum | 11 | Value.Check |
| multipleOf | 9 | Value.Check |
| oneOf | 20 | Value.Check |
| pattern | 9 | Value.Check |
| patternProperties | 23 | Value.Check |
| properties | 28 | Value.Check |
| propertyNames | 20 | Value.Check |
| recursiveRef | 34 | Bundler |
| ref | 12 | Bundler |
| refRemote | 4 | Bundler |
| required | 16 | Value.Check |
| unevaluatedItems | 2 | Bundler |
| unevaluatedProperties | 2 | Bundler |
| uniqueItems | 69 | Value.Check |
| vocabulary | 3 | Bundler |

### draft7 (489 failures)

| Keyword | Failures | Category |
|---------|----------|----------|
| additionalItems | 19 | Value.Check |
| additionalProperties | 19 | Value.Check |
| allOf | 19 | Value.Check |
| anyOf | 8 | Value.Check |
| const | 15 | Value.Check |
| contains | 12 | Value.Check |
| default | 4 | Value.Check |
| dependencies | 41 | Value.Check |
| enum | 20 | Value.Check |
| exclusiveMaximum | 4 | Value.Check |
| exclusiveMinimum | 4 | Value.Check |
| if-then-else | 18 | Value.Check |
| infinite-loop-detection | 2 | Value.Check |
| items | 25 | Value.Check |
| maxItems | 6 | Value.Check |
| maxLength | 7 | Value.Check |
| maxProperties | 10 | Value.Check |
| maximum | 8 | Value.Check |
| minItems | 6 | Value.Check |
| minLength | 7 | Value.Check |
| minProperties | 8 | Value.Check |
| minimum | 11 | Value.Check |
| multipleOf | 9 | Value.Check |
| oneOf | 20 | Value.Check |
| pattern | 9 | Value.Check |
| patternProperties | 19 | Value.Check |
| properties | 28 | Value.Check |
| propertyNames | 13 | Value.Check |
| ref | 10 | Value.Check / Bundler |
| refRemote | 6 | Bundler |
| required | 16 | Value.Check |
| uniqueItems | 69 | Value.Check |

### draft6 (469 failures)

| Keyword | Failures | Category |
|---------|----------|----------|
| additionalItems | 19 | Value.Check |
| additionalProperties | 19 | Value.Check |
| allOf | 19 | Value.Check |
| anyOf | 8 | Value.Check |
| const | 15 | Value.Check |
| contains | 12 | Value.Check |
| default | 4 | Value.Check |
| dependencies | 41 | Value.Check |
| enum | 20 | Value.Check |
| exclusiveMaximum | 4 | Value.Check |
| exclusiveMinimum | 4 | Value.Check |
| items | 25 | Value.Check |
| maxItems | 6 | Value.Check |
| maxLength | 7 | Value.Check |
| maxProperties | 10 | Value.Check |
| maximum | 8 | Value.Check |
| minItems | 6 | Value.Check |
| minLength | 7 | Value.Check |
| minProperties | 8 | Value.Check |
| minimum | 11 | Value.Check |
| multipleOf | 9 | Value.Check |
| oneOf | 20 | Value.Check |
| pattern | 9 | Value.Check |
| patternProperties | 19 | Value.Check |
| properties | 28 | Value.Check |
| propertyNames | 13 | Value.Check |
| ref | 10 | Value.Check / Bundler |
| refRemote | 6 | Bundler |
| required | 16 | Value.Check |
| uniqueItems | 69 | Value.Check |

### draft4 (365 failures)

| Keyword | Failures | Category |
|---------|----------|----------|
| additionalItems | 19 | Value.Check |
| additionalProperties | 17 | Value.Check |
| allOf | 19 | Value.Check |
| anyOf | 8 | Value.Check |
| dependencies | 41 | Value.Check |
| enum | 22 | Value.Check |
| items | 25 | Value.Check |
| maxItems | 6 | Value.Check |
| maxLength | 7 | Value.Check |
| maxProperties | 10 | Value.Check |
| maximum | 10 | Value.Check |
| minItems | 6 | Value.Check |
| minLength | 7 | Value.Check |
| minProperties | 6 | Value.Check |
| minimum | 11 | Value.Check |
| multipleOf | 9 | Value.Check |
| oneOf | 20 | Value.Check |
| pattern | 9 | Value.Check |
| patternProperties | 19 | Value.Check |
| properties | 22 | Value.Check |
| ref | 8 | Value.Check / Bundler |
| refRemote | 4 | Bundler |
| required | 14 | Value.Check |
| uniqueItems | 46 | Value.Check |

### draft3 (246 failures)

| Keyword | Failures | Category |
|---------|----------|----------|
| additionalItems | 13 | Value.Check |
| additionalProperties | 15 | Value.Check |
| dependencies | 18 | Value.Check |
| disallow | 24 | Value.Check |
| divisibleBy | 9 | Value.Check |
| enum | 18 | Value.Check |
| extends | 22 | Value.Check |
| items | 16 | Value.Check |
| maxItems | 4 | Value.Check |
| maxLength | 5 | Value.Check |
| maximum | 10 | Value.Check |
| minItems | 4 | Value.Check |
| minLength | 5 | Value.Check |
| minimum | 11 | Value.Check |
| pattern | 7 | Value.Check |
| patternProperties | 13 | Value.Check |
| properties | 14 | Value.Check |
| ref | 6 | Value.Check / Bundler |
| refRemote | 4 | Bundler |
| required | 4 | Value.Check |
| uniqueItems | 30 | Value.Check |
