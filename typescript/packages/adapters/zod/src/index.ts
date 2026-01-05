import type { ConvertInput, ConvertResult } from "@xschemadev/core";
import { jsonSchemaToZodCode, type JSONSchema } from "./converter.js";

export function convert(input: ConvertInput): ConvertResult {
  const { namespace, id, schema } = input;
  const schemaCode = jsonSchemaToZodCode(schema as JSONSchema);
  const varName = `${namespace}_${id}`;

  return {
    namespace,
    id,
    imports: ['import { z } from "zod"'],
    schema: schemaCode,
    type: `z.infer<typeof ${varName}>`,
  };
}
