import { parse, type ConvertInput, type ConvertResult, type JSONSchema } from "@xschemadev/core";

import { render } from "./renderer.js";

export function convert(input: ConvertInput): ConvertResult {
	const { namespace, id, varName, schema } = input;
	const ir = parse(schema as JSONSchema);
	const schemaCode = render(ir, varName);

	return {
		namespace,
		id,
		varName,
		imports: ['import { z } from "zod"'],
		schema: schemaCode,
		type: `z.infer<typeof ${varName}>`,
		validate: "(data: unknown) => schema.safeParse(data).success",
	};
}
