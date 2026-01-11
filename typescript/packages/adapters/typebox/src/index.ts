import { parse, type ConvertInput, type ConvertResult, type JSONSchema } from "@xschemadev/core";

import { render } from "./renderer.js";

export function convert(input: ConvertInput): ConvertResult {
	const { namespace, id, varName, schema } = input;
	const { node: ir } = parse(schema as JSONSchema);
	const schemaCode = render(ir);

	return {
		namespace,
		id,
		varName,
		imports: [
			'import { Type, type Static } from "@sinclair/typebox"',
			'import { Value } from "@sinclair/typebox/value"',
		],
		schema: schemaCode,
		type: `Static<typeof ${varName}>`,
		validate: "(data: unknown) => Value.Check(schema, data)",
	};
}
