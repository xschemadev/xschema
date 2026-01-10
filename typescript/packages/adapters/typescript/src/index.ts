import type { ConvertInput, ConvertResult, JSONSchema } from "@xschemadev/core";
import { parse } from "@xschemadev/core";

import { render } from "./renderer.js";

/**
 * Converts a JSON Schema to a TypeScript type definition.
 *
 * This is a TYPE-ONLY adapter - it generates TypeScript type expressions,
 * NOT runtime validation code. The `schema` field in the result is empty.
 */
export function convert(input: ConvertInput): ConvertResult {
  const { namespace, id, schema } = input;
  const ir = parse(schema as JSONSchema);
  const typeExpr = render(ir);

  return {
    namespace,
    id,
    imports: [], // No imports needed for pure TS types
    schema: "", // No runtime code - type-only adapter
    type: typeExpr,
    validate: "",
    validateImports: [],
  };
}
