import { parse, type ConvertInput, type ConvertResult, type JSONSchema, type XSchemaAdapter } from "@xschemadev/core";

import { render } from "./renderer.js";

export const adapter = {
	__xschema: true,
	id: "@xschemadev/effect",
	name: "effect",
} as const satisfies XSchemaAdapter;

export function convert(input: ConvertInput): ConvertResult {
	const { namespace, id, varName, schema } = input;
	const ir = parse(schema as JSONSchema);
	const schemaCode = render(ir, varName);

	return {
		namespace,
		id,
		varName,
		imports: ['import { Schema as S } from "effect"'],
		schema: schemaCode,
		type: `typeof ${varName}.Type`,
		validate: '(data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right"',
	};
}
