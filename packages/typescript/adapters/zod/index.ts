import { jsonSchemaToZod } from "json-schema-to-zod";

export interface XSchemaAdapter {
  readonly __brand: "xschema-adapter";
  readonly name: string;
  readonly language: string;
}

export interface ConvertInput {
  namespace: string;
  id: string;
  schema: object;
}

export interface ConvertResult {
  namespace: string;
  id: string;
  imports: string[];
  schema: string;
  type: string;
}

export const zodAdapter: XSchemaAdapter = {
  __brand: "xschema-adapter",
  name: "@xschema/zod",
  language: "typescript",
};

export function convert(namespace: string, id: string, schema: object): ConvertResult {
  const schemaCode = jsonSchemaToZod(schema);
  const varName = `${namespace}_${id}`;

  return {
    namespace,
    id,
    imports: ['import { z } from "zod"'],
    schema: schemaCode,
    type: `z.infer<typeof ${varName}>`,
  };
}
