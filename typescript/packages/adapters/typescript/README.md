# @xschemadev/typescript

TypeScript **type-only** adapter for xschema - converts JSON Schema to TypeScript type definitions.

**Important:** This adapter generates TypeScript types ONLY. No runtime validation code is produced. For runtime validation, use `@xschemadev/zod`, `@xschemadev/valibot`, `@xschemadev/effect`, or `@xschemadev/typebox`.

## What This Adapter Does

Converts JSON Schema structural information into TypeScript types:

| JSON Schema | TypeScript Type |
|-------------|-----------------|
| `{ "type": "string" }` | `string` |
| `{ "type": "number" }` | `number` |
| `{ "type": "boolean" }` | `boolean` |
| `{ "type": "null" }` | `null` |
| `{ "type": "object", "properties": {...} }` | `{ prop: T; optProp?: T }` |
| `{ "type": "array", "items": {...} }` | `T[]` |
| `{ "prefixItems": [...] }` | `[T1, T2]` or `[T1, ...T2[]]` |
| `{ "anyOf": [...] }` | `T1 \| T2` |
| `{ "allOf": [...] }` | `T1 & T2` |
| `{ "oneOf": [...] }` | `T1 \| T2` |
| `{ "const": "value" }` | `"value"` (literal type) |
| `{ "enum": ["a", "b"] }` | `"a" \| "b"` |

## What Cannot Be Expressed

TypeScript types are structural - they cannot enforce runtime constraints:

| JSON Schema | TypeScript Limitation |
|-------------|----------------------|
| `minLength`, `maxLength`, `pattern` | String constraints - just `string` |
| `minimum`, `maximum`, `multipleOf` | Number constraints - just `number` |
| `minItems`, `maxItems`, `uniqueItems` | Array constraints - just `T[]` |
| `minProperties`, `maxProperties` | Object constraints - just `{ ... }` |
| `if`/`then`/`else` | Runtime conditional - renders as union of then/else types |
| `not` | Negation types - renders as `unknown` |
| `contains`, `minContains`, `maxContains` | Array content constraints - ignored |
| `patternProperties`, `propertyNames` | Dynamic property constraints - ignored |
| `format` | String format validation - just `string` |

## Output Format

The adapter returns:
- `schema: ""` - Empty string (no runtime code)
- `type: "<TypeScript type>"` - The type expression
- `imports: []` - Empty array (pure TS types need no imports)

## Installation

```bash
bun add @xschemadev/typescript
```

## Usage

This adapter is used by xschema CLI to generate TypeScript type definitions.

```jsonc
// user.ts.jsonc
{
  "$schema": "https://xschema.dev/schemas/typescript.jsonc",
  "schemas": [
    {
      "id": "User",
      "adapter": "@xschemadev/typescript",
      "source": {
        "type": "object",
        "properties": {
          "name": { "type": "string" },
          "age": { "type": "integer" },
          "email": { "type": "string", "format": "email" }
        },
        "required": ["name"],
        "additionalProperties": false
      }
    }
  ]
}
```

Then run:

```bash
xschema generate
```

Generated output:

```typescript
export type User = { name: string; age?: number; email?: string };
```

Note:
- `integer` becomes `number` (TS has no integer type)
- `format: "email"` is ignored (runtime constraint)
- `additionalProperties: false` means no index signature (strict object)

## Comparison with Other Adapters

| Feature | @xschemadev/typescript | @xschemadev/zod |
|---------|----------------------|-----------------|
| Runtime validation | No | Yes |
| Type inference | Direct type | `z.infer<typeof schema>` |
| Bundle size impact | Zero | Zod library |
| Constraint enforcement | None | Full |

Use this adapter when you:
- Only need TypeScript types for documentation/IDE support
- Want zero runtime overhead
- Will validate data elsewhere (API layer, database)

Use a validation adapter (zod/valibot/effect/typebox) when you:
- Need runtime validation
- Want type-safe parsing with error messages
- Need to validate user input or external data

## Type Rendering Details

### Objects

```typescript
// { "properties": { "a": {...} }, "required": ["a"] }
{ a: T }

// { "properties": { "a": {...} } }  (not required)
{ a?: T }

// { "additionalProperties": true }
{ [key: string]: unknown }

// { "additionalProperties": { "type": "string" } }
{ [key: string]: string }

// { "additionalProperties": false }
// (no index signature - strict object)
```

### Tuples

```typescript
// { "prefixItems": [...], "items": false }
[T1, T2]  // strict tuple

// { "prefixItems": [...], "items": {...} }
[T1, T2, ...T3[]]  // variadic tuple
```

### Literals and Enums

```typescript
// { "const": "hello" }
"hello"

// { "const": [1, 2] }
readonly [1, 2]

// { "const": { "a": 1 } }
{ readonly a: 1 }

// { "enum": ["a", "b", 1] }
"a" | "b" | 1
```

### Conditionals and Negation

```typescript
// { "if": {...}, "then": {...}, "else": {...} }
ThenType | ElseType  // union of possible outcomes

// { "not": {...} }
unknown  // negation cannot be expressed
```
