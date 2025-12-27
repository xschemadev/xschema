import { jsonSchemaToZod } from "json-schema-to-zod";

export interface XSchemaAdapter {
  readonly __brand: "xschema-adapter";
  readonly name: string;
  readonly language: string;
}

export interface ConvertInput {
  name: string;
  schema: object;
}

export interface ConvertResult {
  name: string;
  imports: string[];
  schema: string;
  type: string;
}

export const zodAdapter: XSchemaAdapter = {
  __brand: "xschema-adapter",
  name: "@xschema/zod",
  language: "typescript",
};

export function convert(name: string, schema: object): ConvertResult {
  const schemaCode = jsonSchemaToZod(schema);

  return {
    name,
    imports: ['import { z } from "zod"'],
    schema: schemaCode,
    type: `z.infer<typeof ${name}>`,
  };
}
