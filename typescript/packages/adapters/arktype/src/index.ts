import { parse, type ConvertInput, type ConvertResult, type JSONSchema, type XSchemaAdapter } from "@xschemadev/core";

import { render } from "./renderer.js";

export const adapter = {
	__xschema: true,
	id: "@xschemadev/arktype",
	name: "arktype",
} as const satisfies XSchemaAdapter;

export function convert(input: ConvertInput): ConvertResult {
	const { namespace, id, varName, schema } = input;
	const ir = parse(schema as JSONSchema);
	const schemaCode = render(ir, varName);

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
