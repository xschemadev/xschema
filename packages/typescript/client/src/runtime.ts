export type XSchemaAdapter = {
  readonly name: string;
  readonly __brand: 'xschema-adapter';
};

// Declaration merging interface - extended by generated code
export interface Register {
  // Populated by generated code via declare module
}

// Type helper to get registered schemas
export type RegisteredSchemas = Register extends { schemas: infer S } ? S : Record<string, unknown>;

// Get all registered schema keys as a union type (e.g., "user:Profile" | "another:TSConfig")
type SchemaKeys = Register extends { schemas: infer S } ? keyof S & string : never;

// Extract the ID part from "namespace:id" keys for a specific namespace
// e.g., ExtractID<"user:Profile" | "user:Calendar", "user"> = "Profile" | "Calendar"
type ExtractID<Keys extends string, NS extends string> = 
  Keys extends `${NS}:${infer ID}` ? ID : never;

// Valid keys when using a default namespace - includes full keys AND shorthand IDs
// e.g., if defaultNamespace is "user" and schemas has "user:Profile", "another:TSConfig"
// then valid keys are: "user:Profile" | "another:TSConfig" | "Profile"
type ValidKeys<T, DefaultNS extends string | undefined> = 
  DefaultNS extends string
    ? (keyof T & string) | ExtractID<keyof T & string, DefaultNS>
    : keyof T & string;

// Resolve a shorthand key to its full "namespace:id" form
type ResolveKey<K extends string, DefaultNS extends string | undefined> = 
  K extends `${string}:${string}` 
    ? K  // Already has namespace
    : DefaultNS extends string 
      ? `${DefaultNS}:${K}`  // Prepend default namespace
      : K;  // No default namespace

// Type helper to extract schema types by name
// Only accepts full "namespace:id" keys for explicitness
// Use the xschema client for shorthand ID lookups
export type XSchemaType<N extends SchemaKeys> = 
  Register extends { schemaTypes: infer T }
    ? N extends keyof T 
      ? T[N]
      : never
    : never;

// Configuration for createXSchemaClient
export type XSchemaConfig<T extends Record<string, unknown> = RegisteredSchemas> = {
  schemas?: T;
  defaultNamespace?: string;
};

/**
 * Creates an xschema client for looking up schemas by namespace:id
 * 
 * Provides full TypeScript autocompletion and compile-time errors for invalid keys.
 * 
 * @example
 * ```ts
 * import { schemas } from "./.xschema/xschema.gen";
 * import { createXSchemaClient } from "@xschema/client";
 * 
 * const xschema = createXSchemaClient({ schemas, defaultNamespace: "user" });
 * 
 * // Full autocompletion for all schema keys
 * const userSchema = xschema("user:Profile");  // ✓
 * const tsConfig = xschema("another:TSConfig"); // ✓
 * 
 * // With defaultNamespace, can omit namespace for that namespace
 * const profile = xschema("Profile");  // ✓ resolves to "user:Profile"
 * 
 * // TypeScript error for invalid keys
 * const invalid = xschema("nonexistent");  // ✗ Type error!
 * ```
 */
export function createXSchemaClient<
  const T extends Record<string, unknown> = RegisteredSchemas,
  const DefaultNS extends string | undefined = undefined
>(
  config: XSchemaConfig<T> & { defaultNamespace?: DefaultNS }
) {
  const schemas = config.schemas ?? ({} as T);
  const defaultNs = config.defaultNamespace;

  /**
   * Look up a schema by key.
   * @param key - Schema key in "namespace:id" format, or just "id" if defaultNamespace is set
   * @returns The schema validator (e.g., Zod schema)
   */
  function lookup<K extends ValidKeys<T, DefaultNS>>(
    key: K
  ): T[ResolveKey<K, DefaultNS> & keyof T] {
    // If key includes ":", use as-is; otherwise prepend defaultNamespace
    const resolvedKey = (key as string).includes(':') 
      ? key 
      : defaultNs 
        ? `${defaultNs}:${key}` 
        : key;
    
    if (!(resolvedKey in schemas)) {
      throw new Error(`Unknown schema: ${resolvedKey}. Run \`xschema generate\`.`);
    }
    
    return schemas[resolvedKey as keyof T] as T[ResolveKey<K, DefaultNS> & keyof T];
  }

  return lookup;
}
