import type { ConvertInput, ConvertResult, JSONSchema } from "@xschemadev/core";
import { parse } from "@xschemadev/core";
import { render } from "./renderer.js";

export function convert(input: ConvertInput): ConvertResult {
	const { namespace, id, schema } = input;
	const ir = parse(schema as JSONSchema);
	const schemaCode = render(ir);
	const varName = `${namespace}_${id}`;

	return {
		namespace,
		id,
		imports: ['import { z } from "zod"'],
		schema: schemaCode,
		type: `z.infer<typeof ${varName}>`,
	};
}

// Re-export the old function for backwards compatibility during transition
export { jsonSchemaToZodCode } from "./converter.js";
