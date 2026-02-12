# issue-10: properties / properties whose names are Javascript object property names

## Error Signature

- **Keyword**: `properties`
- **Case**: `properties whose names are Javascript object property names`
- **Normalized got**: `error: this.entries[key]._run is not a function. (In 'this.entries[key]._run({ typed: !1, value: value2 }, config2)', 'this.entries[key]._run' is undefined)`

## Root Cause

Valibot's object schemas store property entries as a plain JS object. When property names like `__proto__`, `constructor`, or `toString` are used as keys in the object literal passed to `v.looseObject({ __proto__: v.number() })`, the lookup `this.entries[key]._run(...)` hits Object.prototype methods instead of valibot schema entries. The key `__proto__` in an object literal sets the prototype rather than creating an own property.

## Baseline (25 failures)

| # | Adapter | Draft | Test | Expected | Got | Report Path |
|---|---------|-------|------|----------|-----|-------------|
| 1 | typescript valibot | draft4 | __proto__ not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 2 | typescript valibot | draft4 | all present and valid | true | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 3 | typescript valibot | draft4 | constructor not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 4 | typescript valibot | draft4 | none of the properties mentioned | true | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 5 | typescript valibot | draft4 | toString not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 6 | typescript valibot | draft6 | __proto__ not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 7 | typescript valibot | draft6 | all present and valid | true | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 8 | typescript valibot | draft6 | constructor not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 9 | typescript valibot | draft6 | none of the properties mentioned | true | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 10 | typescript valibot | draft6 | toString not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 11 | typescript valibot | draft7 | __proto__ not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 12 | typescript valibot | draft7 | all present and valid | true | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 13 | typescript valibot | draft7 | constructor not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 14 | typescript valibot | draft7 | none of the properties mentioned | true | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 15 | typescript valibot | draft7 | toString not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 16 | typescript valibot | draft2019-09 | __proto__ not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 17 | typescript valibot | draft2019-09 | all present and valid | true | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 18 | typescript valibot | draft2019-09 | constructor not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 19 | typescript valibot | draft2019-09 | none of the properties mentioned | true | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 20 | typescript valibot | draft2019-09 | toString not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 21 | typescript valibot | draft2020-12 | __proto__ not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 22 | typescript valibot | draft2020-12 | all present and valid | true | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 23 | typescript valibot | draft2020-12 | constructor not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 24 | typescript valibot | draft2020-12 | none of the properties mentioned | true | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
| 25 | typescript valibot | draft2020-12 | toString not valid | false | error: this.entries[key]._run is not a function | typescript/packages/adapters/valibot/compliance/results/REPORT.md |
