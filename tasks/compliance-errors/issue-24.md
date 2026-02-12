# issue-24: const / const with array

**Error signature**: keyword `const`, case `const with array`, test `same array is valid`, normalized got `false`

## Baseline (8 failures)

| # | Adapter | Draft | Test | Expected | Got | Report Path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript effect | draft6 | same array is valid | true | false | typescript/packages/adapters/effect/compliance/results/draft6.json |
| 2 | typescript effect | draft7 | same array is valid | true | false | typescript/packages/adapters/effect/compliance/results/draft7.json |
| 3 | typescript effect | draft2019-09 | same array is valid | true | false | typescript/packages/adapters/effect/compliance/results/draft2019-09.json |
| 4 | typescript effect | draft2020-12 | same array is valid | true | false | typescript/packages/adapters/effect/compliance/results/draft2020-12.json |
| 5 | typescript valibot | draft6 | same array is valid | true | false | typescript/packages/adapters/valibot/compliance/results/draft6.json |
| 6 | typescript valibot | draft7 | same array is valid | true | false | typescript/packages/adapters/valibot/compliance/results/draft7.json |
| 7 | typescript valibot | draft2019-09 | same array is valid | true | false | typescript/packages/adapters/valibot/compliance/results/draft2019-09.json |
| 8 | typescript valibot | draft2020-12 | same array is valid | true | false | typescript/packages/adapters/valibot/compliance/results/draft2020-12.json |

## Root Cause

Both effect and valibot adapters used a shallow key-sort for deep equality comparison in `renderLiteral`:
```javascript
JSON.stringify(val, Object.keys(val as object).sort())
```

`JSON.stringify`'s replacer array only sorts top-level keys. For arrays containing objects (like `[{"foo": "bar"}]`), nested object keys are not sorted. The build-time `sortedStringify()` uses recursive `deepSortKeys()`, creating a mismatch.

## Fix

Replaced the shallow sort with `DEEP_SORTED_STRINGIFY_RUNTIME` from `@xschemadev/core` — the same recursive sort-then-stringify IIFE that the zod adapter uses. This recursively walks arrays and sorts object keys at every depth before stringifying.

Files changed:
- `typescript/packages/adapters/effect/src/renderer.ts`
- `typescript/packages/adapters/valibot/src/renderer.ts`
