# PRD: Twoslash Integration for Documentation

## Introduction

Add TypeScript Twoslash to the xschema documentation to show inline type information and type hovers. Examples import from real generated code, proving the type-safety benefits of xschema with actual types (not mocks).

## Goals

- **Autocomplete**: Show schema key suggestions so users see what's available
- **Type inference**: Show what `xschema()` returns so users know how to use it
- **Error prevention**: Show compile-time errors for invalid schema keys
- **Extracted types**: Display `XSchemaType<>` for using types in signatures
- **Workflow clarity**: Make "run `xschema generate`" obvious in the flow
- **Real types**: Use actual generated code, not mocks, proving it works

## User Stories

### Create examples directory structure

**Description:** As a developer, I need a structured examples directory so each adapter has its own generated code that Twoslash can import, supporting both simple and multi-namespace usage.

**Acceptance Criteria:**

- [ ] Create `web/examples/schemas/user.schema.json` (id, name?, email)
- [ ] Create `web/examples/schemas/app.schema.json` (theme enum, language, debug?)
- [ ] Create adapter configs, each with both schemas in different namespaces (user:Profile, app:Config):
  - [ ] `web/examples/typescript/zod/xschema.jsonc`
  - [ ] `web/examples/typescript/effect/xschema.jsonc`
  - [ ] `web/examples/typescript/arktype/xschema.jsonc`
  - [ ] `web/examples/typescript/valibot/xschema.jsonc`
- [ ] Run `xschema generate` in each adapter directory
- [ ] Commit generated `.xschema/xschema.gen.ts` files
- [ ] Typecheck passes

### Configure Twoslash in Fumadocs

**Description:** As a developer, I need Twoslash configured so code blocks can show type information from real imports.

**Acceptance Criteria:**

- [ ] Install `fumadocs-twoslash` and `twoslash` packages
- [ ] Add `transformerTwoslash()` to `source.config.ts` rehypeCodeOptions
- [ ] Add Twoslash CSS import to `src/styles/app.css`
- [ ] Add Twoslash UI components to MDX components in `src/routes/docs/$.tsx`
- [ ] Configure Vite to externalize `typescript` and `twoslash` for SSR
- [ ] Configure Twoslash `compilerOptions` to resolve paths to examples directory
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill - hover over a `^?` annotation shows type popup

### Add Twoslash to typescript.mdx client setup

**Description:** As a user reading the TypeScript guide, I want to see real inferred types when using `createXSchemaClient`.

**Acceptance Criteria:**

- [ ] "Use the client" example shows type hover on `profileSchema` variable
- [ ] Type popup displays the actual Zod schema type from generated code
- [ ] Imports from `examples/typescript/zod/.xschema/xschema.gen`
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### Add Twoslash to typescript.mdx schema lookup section

**Description:** As a user, I want to see return types for schema lookups.

**Acceptance Criteria:**

- [ ] "Schema lookup" section shows type hover on lookup result
- [ ] Shows `defaultNamespace` shorthand working with correct types
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### Add Twoslash to typescript.mdx XSchemaType section

**Description:** As a user, I want to see what types `XSchemaType<>` extracts.

**Acceptance Criteria:**

- [ ] XSchemaType example shows extracted type on hover
- [ ] Type displays the real inferred object shape from generated code
- [ ] "Using in function signatures" example shows typed property access
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### Add autocomplete and error examples to typescript.mdx

**Description:** As a user, I want to see that I get autocomplete for schema keys and errors for invalid keys, so I understand the DX benefits.

**Acceptance Criteria:**

- [ ] Add "Developer Experience" or similar section to typescript.mdx
- [ ] Show autocomplete example using `^|` annotation - displays available schema keys
- [ ] Show error example using `@errors` annotation - invalid key like `xschema("wrong:Key")` shows compile error
- [ ] Add callout reminding users to run `xschema generate` after schema changes
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill - autocomplete popup and error highlight visible

### Add Twoslash to Quick Start index.mdx

**Description:** As a new user on the Quick Start page, I want to see types in the "Use the generated schemas" example.

**Acceptance Criteria:**

- [ ] "Use the generated schemas" step shows type hover on `profileSchema`
- [ ] Imports from real Zod generated code
- [ ] Keeps the example concise - just one hover annotation
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### Add Twoslash to Zod adapter usage

**Description:** As a user viewing the Zod adapter page, I want to see the typed parse result.

**Acceptance Criteria:**

- [ ] Usage example shows type hover on `user` variable after `.parse()`
- [ ] Type displays the actual parsed object shape from Zod generated code
- [ ] Imports from `examples/typescript/zod/.xschema/xschema.gen`
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### Add Twoslash to Effect adapter usage

**Description:** As a user viewing the Effect adapter page, I want to see the typed decode result.

**Acceptance Criteria:**

- [ ] Usage example shows type hover on `user` variable
- [ ] Shows Effect Schema decode pattern with actual types
- [ ] Imports from `examples/typescript/effect/.xschema/xschema.gen`
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### Add Twoslash to ArkType adapter usage

**Description:** As a user viewing the ArkType adapter page, I want to see the typed validation result.

**Acceptance Criteria:**

- [ ] Usage example shows type hover on `user` variable
- [ ] Shows ArkType's function-call pattern with actual types
- [ ] Imports from `examples/typescript/arktype/.xschema/xschema.gen`
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

### Add Twoslash to Valibot adapter usage

**Description:** As a user viewing the Valibot adapter page, I want to see the typed parse result.

**Acceptance Criteria:**

- [ ] Usage example shows type hover on `user` variable
- [ ] Shows Valibot's `v.parse()` pattern with actual types
- [ ] Imports from `examples/typescript/valibot/.xschema/xschema.gen`
- [ ] Typecheck passes
- [ ] Verify in browser using dev-browser skill

## Functional Requirements

- FR-1: Create examples directory structure at `web/examples/` organized by language then adapter
- FR-2: Share two source schemas (`schemas/user.schema.json`, `schemas/app.schema.json`) across all adapter examples to support multi-namespace demos
- FR-3: Each adapter has its own xschema config and generated output
- FR-4: Generated code is committed to the repo (not generated at build time)
- FR-5: Twoslash code blocks import from real generated files using relative paths
- FR-6: Use `^?` annotation to show type hovers on key variables
- FR-7: Use `^|` annotation to show autocomplete for schema keys
- FR-8: Use `@errors` annotation to show compile-time errors for invalid keys
- FR-9: Include prominent "run `xschema generate`" reminder in workflow sections
- FR-10: Limit Twoslash usage to TypeScript usage examples - don't add to JSON/bash blocks

## Non-Goals

- No Twoslash on JSON Schema examples (not TypeScript)
- No Twoslash on bash/shell commands
- No inline type mocking with `// @filename` - use real generated code
- No Twoslash on configuration reference pages
- No build-time generation - commit generated code for reliability

## Technical Considerations

### Directory Structure

```
web/
  examples/
    schemas/
      user.schema.json           # Profile schema (id, name?, email)
      app.schema.json            # Config schema (theme, language, debug?)
    
    typescript/
      zod/
        xschema.jsonc            # both schemas: user:Profile, app:Config
        .xschema/
          xschema.gen.ts         # committed generated code
      effect/
        xschema.jsonc
        .xschema/
          xschema.gen.ts
      arktype/
        xschema.jsonc
        .xschema/
          xschema.gen.ts
      valibot/
        xschema.jsonc
        .xschema/
          xschema.gen.ts
    
    python/                      # future
      pydantic/
        xschema.jsonc
        .xschema/
          xschema.gen.py
```

Each adapter config includes both schemas with different namespaces to support:
- Simple usage: `xschema("user:Profile")`
- Multi-namespace: `xschema("app:Config")`  
- Default namespace: `defaultNamespace: "user"` with shorthand `xschema("Profile")`

### Twoslash Configuration

- **Vite SSR**: Externalize `typescript` and `twoslash` packages
- **Compiler options**: Configure `paths` to resolve imports to examples directory
- **Performance**: Enable filesystem cache with `createFileSystemTypesCache()`
- **Languages**: Configure Shiki with `js`, `jsx`, `ts`, `tsx`

### Regenerating Examples

When adapters change, regenerate with:
```bash
cd web/examples/typescript/zod && bunx xschema generate
cd web/examples/typescript/effect && bunx xschema generate
# ... etc
```

Consider adding a script: `bun run regenerate:examples`

## Design Considerations

Twoslash should demonstrate three key DX benefits:

### 1. Autocomplete - "What schemas do I have?"
```ts twoslash
const schema = xschema("")
//                      ^|
// Shows: "user:Profile", "app:Config"
```

### 2. Type inference - "What does this return?"
```ts twoslash
const profile = xschema("user:Profile")
//    ^? ZodObject<{ id: ZodString; name: ZodOptional<ZodString>; email: ZodString }>
```

### 3. Error prevention - "Catch mistakes early"
```ts twoslash
// @errors: 2345
const wrong = xschema("typo:Schema")
// Error: Argument of type '"typo:Schema"' is not assignable...
```

### Workflow reminder
Add a Callout in relevant sections:
> **Remember:** Run `xschema generate` after changing your JSON Schemas to update the generated types.

## Success Metrics

- Users immediately understand they get autocomplete for schema keys
- Users see that invalid keys cause compile-time errors (not runtime surprises)
- Users can hover over any annotated variable and see its real type
- "Run xschema generate" workflow is clear and prominent
- No build errors or Twoslash failures
- Examples stay in sync when regenerated
