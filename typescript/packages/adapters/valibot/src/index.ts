import { parse, type ConvertInput, type ConvertResult, type JSONSchema } from "@xschemadev/core";

import { render } from "./renderer.js";

export function convert(input: ConvertInput): ConvertResult {
	const { namespace, id, varName, schema, vocabulary } = input;
	const ir = parse(schema as JSONSchema, { vocabulary });
	const schemaCode = render(ir);

	return {
		namespace,
		id,
		varName,
		imports: ['import * as v from "valibot"'],
		schema: schemaCode,
		type: `v.InferOutput<typeof ${varName}>`,
		validate: "(data: unknown) => safeParse(schema, data).success",
		validateImports: ['import { safeParse } from "valibot"'],
	};
}
