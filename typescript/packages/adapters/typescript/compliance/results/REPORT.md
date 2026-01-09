# @xschemadev/typescript Compliance Report

Generated: 2026-01-09T19:00:00Z

## Type-Only Adapter

This adapter generates **TypeScript type definitions only** - no runtime validation code is produced.

Standard JSON Schema compliance tests cannot be run against this adapter because:
1. Compliance tests require runtime validation (checking if data matches a schema)
2. TypeScript types exist only at compile-time and are erased at runtime
3. There is no runtime code to execute test cases against

## Structural Features Supported

The following JSON Schema structural information is converted to TypeScript types:

| JSON Schema Feature | TypeScript Output | Status |
|---------------------|-------------------|--------|
| `type: "string"` | `string` | ✅ |
| `type: "number"` | `number` | ✅ |
| `type: "integer"` | `number` | ✅ |
| `type: "boolean"` | `boolean` | ✅ |
| `type: "null"` | `null` | ✅ |
| `type: "object"` with `properties` | `{ prop: T; optProp?: T }` | ✅ |
| `type: "array"` with `items` | `T[]` | ✅ |
| `prefixItems` (tuples) | `[T1, T2]` or `[T1, ...T2[]]` | ✅ |
| `additionalProperties: false` | No index signature | ✅ |
| `additionalProperties: true` | `[key: string]: unknown` | ✅ |
| `additionalProperties: schema` | `[key: string]: T` | ✅ |
| `anyOf` | `T1 \| T2` (union) | ✅ |
| `allOf` | `T1 & T2` (intersection) | ✅ |
| `oneOf` | `T1 \| T2` (union) | ✅ |
| `const` (primitives) | Literal type (`"value"`, `42`) | ✅ |
| `const` (objects/arrays) | `readonly` literal type | ✅ |
| `enum` | Union of literals | ✅ |
| `$ref` (resolved) | Rendered type | ✅ |
| `nullable` | `T \| null` | ✅ |

## Runtime Constraints Not Expressible

TypeScript's type system cannot express runtime validation constraints. These JSON Schema features are intentionally ignored:

| JSON Schema Feature | Reason |
|---------------------|--------|
| `minLength`, `maxLength` | String length is a runtime property |
| `pattern` | Regex matching is a runtime operation |
| `format` | Format validation (email, uri, etc.) is runtime |
| `minimum`, `maximum` | Number bounds are runtime constraints |
| `exclusiveMinimum`, `exclusiveMaximum` | Number bounds are runtime constraints |
| `multipleOf` | Number divisibility is runtime |
| `minItems`, `maxItems` | Array length is runtime |
| `uniqueItems` | Uniqueness check is runtime |
| `contains`, `minContains`, `maxContains` | Array content validation is runtime |
| `minProperties`, `maxProperties` | Object key count is runtime |
| `patternProperties` | Dynamic property matching is runtime |
| `propertyNames` | Key validation is runtime |
| `dependencies`, `dependentRequired`, `dependentSchemas` | Property dependency is runtime |
| `not` | Negation requires runtime checking (renders as `unknown`) |
| `if`/`then`/`else` | Conditional validation is runtime (renders as union of then/else) |

## Validation Approach

Since this adapter produces type-only output, validation is performed through:

1. **TypeScript Compiler**: Types are checked at compile-time via `bun run typecheck`
2. **Syntax Validation**: Generated type expressions are syntactically valid TypeScript
3. **Structural Correctness**: JSON Schema structure maps correctly to TypeScript types

## When to Use This Adapter

Use `@xschemadev/typescript` when you:
- Only need TypeScript types for IDE support and documentation
- Want zero runtime overhead (no validation library bundled)
- Will validate data through other means (API layer, database constraints)

For runtime validation, use:
- `@xschemadev/zod` (98.9% compliance)
- `@xschemadev/typebox` (98.3% compliance)
- `@xschemadev/effect` (94.3% compliance)
- `@xschemadev/valibot` (89.9% compliance)
