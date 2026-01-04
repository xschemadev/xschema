import type { ConvertInput, ConvertResult } from "../../index.js";
import { jsonSchemaToZod } from "json-schema-to-zod";

export function convert(input: ConvertInput): ConvertResult {
  const { namespace, id, schema } = input;
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
