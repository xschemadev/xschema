import { parse, type ConvertInput, type ConvertResult, type JSONSchema, type XSchemaAdapter } from "@xschemadev/core";

import { render } from "./renderer.js";

export const adapter = {
	__xschema: true,
	id: "@xschemadev/valibot",
	name: "valibot",
} as const satisfies XSchemaAdapter;

export function convert(input: ConvertInput): ConvertResult {
	const { namespace, id, varName, schema } = input;
	const ir = parse(schema as JSONSchema);
	const schemaCode = render(ir, varName);

	return {
		namespace,
		id,
		varName,
		imports: ['import * as v from "valibot"'],
		schema: schemaCode,
		type: `v.InferOutput<typeof ${varName}>`,
		validate: "(data: unknown) => safeParse(schema, data).success",
		validationImports: ['import { safeParse } from "valibot"'],
	};
}
