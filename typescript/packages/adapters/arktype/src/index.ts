import { parse, type ConvertInput, type ConvertResult, type JSONSchema } from "@xschemadev/core";

import { render } from "./renderer.js";

export function convert(input: ConvertInput): ConvertResult {
	const { namespace, id, varName, schema } = input;
	const ir = parse(schema as JSONSchema);
	const schemaCode = render(ir);

	return {
		namespace,
		id,
		varName,
		imports: ['import { type } from "arktype"'],
		schema: schemaCode,
		type: `typeof ${varName}.infer`,
		validate: "(data: unknown) => !(schema(data) instanceof type.errors)",
	};
}
