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

// Check if T has actual schema keys (not just empty Record<string, unknown>)
type HasSchemas<T> = Record<string, unknown> extends T ? false : true;

type PleaseRunXSchemaGenerate = {
  readonly __error: 'Run `xschema generate`';
};

// Configuration - parsed by CLI from createXSchemaClient call
export type XSchemaConfig<T extends Record<string, unknown> = RegisteredSchemas> = {
  schemas?: T;
  outputDir?: string;          // Output directory (default: .xschema)
  maxParallelFetches?: number; // Max concurrent HTTP requests (default: 10)
  requestTimeoutMs?: number;   // HTTP request timeout in ms (default: 30000)
  maxFetchRetries?: number;    // Max retry attempts for fetching schemas (default: 3)
};

type SchemaResult<T, N extends string> = HasSchemas<T> extends true
  ? N extends keyof T ? T[N] : PleaseRunXSchemaGenerate
  : PleaseRunXSchemaGenerate;

export function createXSchemaClient<T extends Record<string, unknown> = RegisteredSchemas>(
  config: XSchemaConfig<T> = {}
) {
  const schemas = config.schemas ?? ({} as T);

  function lookup<N extends string>(name: N): SchemaResult<T, N> {
    if (!(name in schemas)) {
      throw new Error(`Run xschema generate - unknown schema: ${name}`);
    }
    return schemas[name as keyof T] as SchemaResult<T, N>;
  }

  return {
    fromURL: <N extends string>(name: N, _url: string, _adapter: XSchemaAdapter) => lookup(name),
    fromFile: <N extends string>(name: N, _path: string, _adapter: XSchemaAdapter) => lookup(name),
    ...schemas,
  } as {
    fromURL: <N extends string>(name: N, url: string, adapter: XSchemaAdapter) => SchemaResult<T, N>;
    fromFile: <N extends string>(name: N, path: string, adapter: XSchemaAdapter) => SchemaResult<T, N>;
  } & T;
}
