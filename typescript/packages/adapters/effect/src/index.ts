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
		imports: ['import { Schema as S } from "effect"'],
		schema: schemaCode,
		type: `typeof ${varName}.Type`,
		validate: '(data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right"',
	};
}
