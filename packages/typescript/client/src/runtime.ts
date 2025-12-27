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

type PleaseRunXSchemaGenerate = {
  readonly __error: 'Run `xschema generate`';
};

export function createXSchemaClient<T extends Record<string, unknown>>(schemas: T) {
  function lookup<N extends string>(name: N): N extends keyof T ? T[N] : PleaseRunXSchemaGenerate {
    if (!(name in schemas)) {
      throw new Error(`Run xschema generate - unknown schema: ${name}`);
    }
    return schemas[name as keyof T] as any;
  }

  return {
    fromURL: <N extends string>(name: N, _url: string, _adapter: XSchemaAdapter) => lookup(name),
    fromFile: <N extends string>(name: N, _path: string, _adapter: XSchemaAdapter) => lookup(name),
    ...schemas,
  } as const;
}
