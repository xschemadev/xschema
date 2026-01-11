import { parse, type ConvertInput, type ConvertResult, type JSONSchema } from "@xschemadev/core";

import { render } from "./renderer.js";

/**
 * Converts a JSON Schema to a TypeScript type definition.
 *
 * This is a TYPE-ONLY adapter - it generates TypeScript type expressions,
 * NOT runtime validation code. The `schema` field in the result is empty.
 */
export function convert(input: ConvertInput): ConvertResult {
  const { namespace, id, varName, schema } = input;
  const { node: ir } = parse(schema as JSONSchema);
  const typeExpr = render(ir);

  return {
    namespace,
    id,
    varName,
    imports: [],
    type: typeExpr,
  };
}
