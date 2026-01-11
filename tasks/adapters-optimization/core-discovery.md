# core discovery: lossless parsing + version semantics

## baseline observations

- entrypoint is `parse(schema: JSONSchema | boolean): SchemaNode` in `typescript/packages/core/src/parser/index.ts`
- version detection is string-based on `$schema` in `typescript/packages/core/src/schema/version.ts`:
  - missing/unknown `$schema` defaults to `draft-2020-12`
  - openapi 3.0 detected via substring match (`openapi` + `3.0`)
- schemas are normalized before parsing via `normalizeDeep(schema, version)` in `typescript/packages/core/src/schema/normalizer.ts`:
  - legacy transforms: `divisibleBy -> multipleOf`, `extends -> allOf`, `disallow -> not` (draft3)
  - boolean exclusive min/max (draft3/4) are rewritten to numeric exclusive keywords when possible
  - `unevaluatedProperties` is opportunistically rewritten to `additionalProperties` only when there are no applicators and `additionalProperties` is unset
- parsing produces a `SchemaNode` union in `typescript/packages/core/src/ir/nodes.ts`
  - ir is focused on validation semantics: it does not carry schema metadata (e.g. `title`, `description`, `default`, `$comment`, `readOnly`, custom keywords)
  - unknown keywords are effectively dropped: they’re present on `JSONSchema` via `[key: string]: unknown`, but the parser never stores them in the ir
- `$ref` handling in `parseRef`:
  - only supports internal json pointers (`#...`)
  - resolves via `resolveJsonPointer` (throws on missing paths)
  - caches ref results in `ctx.refs` and breaks circular refs by returning `ref -> any`
  - for draft-2019-09+ siblings, it intersects the `$ref` with a parsed schema made from the remaining sibling keywords

## keyword-based throws

current core parsing is not total: it throws hard `Error`s for some keywords/conditions.

### in `typescript/packages/core/src/parser/index.ts`

- `unevaluatedItems` (any presence)
  - throws: `"unevaluatedItems is not supported"`
- `unevaluatedProperties` + applicators (anyOf/allOf/oneOf/if/$ref/dependentSchemas/not)
  - throws: `"unevaluatedProperties with applicators is not supported"`

note: `unevaluatedProperties` without applicators is often normalized into `additionalProperties` (so it won’t reach the throw), but only in the simple case.

### in `typescript/packages/core/src/utils/json-pointer.ts`

- external `$ref` (`$ref` not starting with `#`)
  - throws: `External $ref is not supported: ...`
- missing pointer target
  - throws: `Reference not found: ...`

## preservation proposal

goal: make parsing total (no keyword throws) and preserve unknown/unsupported keywords so adapters can choose behavior.

a small, incremental plan that keeps adapters stable:

1. **add an optional “source/unknown keywords” payload on every ir node**
   - new field (name bikeshed): `source?: { version: JSONSchemaVersion; unknown: Record<string, unknown> }`
   - `unknown` contains *only* keywords not consumed by the parser + normalizer
   - preserve raw keyword values as-is (including objects/arrays), no re-marshalling

2. **track “draft context” via a schema path**
   - extend `ParseContext` to carry a current path (json-pointer-ish), passed down during recursion
   - store `path` in `source` or in a separate `issues` list

3. **replace throws with captured issues + best-effort parsing**
   - introduce `ParseIssue` (e.g. `{ code, keyword, path, message }`)
   - expose a new api `parseWithIssues(schema): { node: SchemaNode; issues: ParseIssue[] }`
   - keep existing `parse()` as a compatibility wrapper (either:
     - returns `node` and drops `issues`, or
     - returns `node` and only throws on truly unrecoverable cases behind a flag)
   - for previously-throwing keywords:
     - `unevaluatedItems`: keep semantics best-effort (likely ignore for most adapters) but capture issue + preserve keyword in `unknown`
     - `unevaluatedProperties` + applicators: capture issue; do *not* throw; preserve keyword in `unknown`
     - external/missing `$ref`: capture issue; return a `ref` that resolves to `any` (or a dedicated `unknownRef` kind if we accept adapter updates)

4. **document a stable “known keyword set”**
   - implement a single source of truth list for “consumed” keywords per parser module so “unknown” is deterministic
   - normalize step should also register which legacy keys were rewritten/deleted, so we can optionally preserve them (useful for debugging)

## migration notes

- if we only add optional fields (`source`, `unknown`) on existing node interfaces, adapters remain source-compatible (they can ignore it).
- introducing new node kinds (e.g. `unsupported`, `unknownRef`) would require updating every adapter renderer switch to stay exhaustive.
  - prefer optional fields + issues first, then consider new kinds once adapter updates are planned.
- if we add `parseWithIssues`, keep `parse()` unchanged for now; migrate adapters/cli later when we want surfaced warnings.
