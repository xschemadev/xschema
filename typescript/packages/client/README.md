# @xschemadev/client

runtime client for xschema with full typescript type inference.

## installation

```bash
npm install @xschemadev/client
# or
bun add @xschemadev/client
```

## usage

### basic client setup

```typescript
import { createXSchemaClient } from "@xschemadev/client";
import { schemas } from "./.xschema/xschema.gen.js";

const xschema = createXSchemaClient({ schemas });

// lookup a schema by full namespace:id key
const userSchema = xschema("user:Profile");
const tsConfigSchema = xschema("another:TSConfig");

// use the schema (e.g. with zod)
const result = userSchema.parse({ id: "123", name: "john" });
```

### default namespace

set a default namespace to omit the namespace prefix for that namespace:

```typescript
const xschema = createXSchemaClient({ 
  schemas, 
  defaultNamespace: "user" 
});

// shorthand for "user:Profile"
const profile = xschema("Profile");

// still need full key for other namespaces
const tsConfig = xschema("another:TSConfig");
```

### type helpers

use `XSchemaType` to extract typescript types from your schemas:

```typescript
import type { XSchemaType } from "@xschemadev/client";

// extract type from any schema
type UserProfile = XSchemaType<"user:Profile">;
type TSConfig = XSchemaType<"another:TSConfig">;

// use in function signatures
function validateUser(data: unknown): UserProfile {
  const schema = xschema("user:Profile");
  return schema.parse(data);
}
```

## schema-only vs type-only vs both

xschema adapters can generate three kinds of outputs:

### 1. schema + type (most common)

adapters like zod and arktype generate both a runtime validator and a type:

```typescript
// generated code
const user_User = z.object({ id: z.string(), name: z.string() });
type user_UserType = z.infer<typeof user_User>;

// in schemas object (runtime lookup available)
export const schemas = {
  "user:User": user_User,
} as const;

// in SchemaTypes (type extraction available)
export type SchemaTypes = {
  "user:User": user_UserType;
};
```

both `xschema("user:User")` and `XSchemaType<"user:User">` work.

### 2. type-only

future adapters may generate only types (e.g. pure typescript type generator):

```typescript
// generated code
export type user_User = { id: string; name: string };

// NOT in schemas object (no runtime validator)
export const schemas = {} as const;

// in SchemaTypes (type extraction available)
export type SchemaTypes = {
  "user:User": user_User;
};
```

only `XSchemaType<"user:User">` works. attempting `xschema("user:User")` causes a compile error because there's no runtime validator.

### 3. schema-only

future adapters may generate only validators without type inference:

```typescript
// generated code
const user_User = customValidator({ id: "string", name: "string" });

// in schemas object (runtime lookup available)
export const schemas = {
  "user:User": user_User,
} as const;

// in SchemaTypes using typeof (type extraction available)
export type SchemaTypes = {
  "user:User": typeof user_User;
};
```

both `xschema("user:User")` and `XSchemaType<"user:User">` work, but the type is `typeof user_User` instead of an explicit interface.

## key behaviors

- **runtime lookup**: only entries with a runtime schema (`.Code` in generated output) appear in the `schemas` object
- **type extraction**: all entries appear in `SchemaTypes`, regardless of whether they have runtime validators
- **compile-time safety**: typescript prevents using type-only keys with `xschema()` and schema-only/type-only keys where not applicable

## examples

### using with zod

```typescript
import { createXSchemaClient } from "@xschemadev/client";
import type { XSchemaType } from "@xschemadev/client";
import { schemas } from "./.xschema/xschema.gen.js";

const xschema = createXSchemaClient({ schemas, defaultNamespace: "user" });

// validate data
const userSchema = xschema("Profile");
const validatedUser = userSchema.parse(unknownData);

// extract types
type User = XSchemaType<"user:Profile">;

function createUser(data: User): void {
  // ...
}
```

### using with arktype

```typescript
import { createXSchemaClient } from "@xschemadev/client";
import type { XSchemaType } from "@xschemadev/client";
import { schemas } from "./.xschema/xschema.gen.js";

const xschema = createXSchemaClient({ schemas });

// validate data
const configSchema = xschema("app:Config");
const result = configSchema(unknownData);

if (result.problems) {
  console.error(result.problems);
} else {
  console.log(result.data);
}

// extract types
type Config = XSchemaType<"app:Config">;
```

### type-only adapter (future)

```typescript
import type { XSchemaType } from "@xschemadev/client";

// NO runtime client needed - types only
type User = XSchemaType<"user:Profile">;
type Config = XSchemaType<"app:Config">;

// compile error: type-only entries cannot be looked up at runtime
// const schema = xschema("user:Profile"); // ❌
```

## error handling

if you try to look up a schema that doesn't exist:

```typescript
const schema = xschema("nonexistent:Key");
// Error: Unknown schema: nonexistent:Key. Run `xschema generate`.
```

typescript will catch this at compile time with full autocomplete support.

## typescript features

- full autocomplete for all registered schema keys
- compile-time errors for invalid keys
- type inference for schema validators
- separates type-only from runtime-available schemas

## learn more

- [xschema documentation](https://xschema.dev/docs)
- [github repository](https://github.com/xschemadev/xschema)
