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
		imports: [
			'import { Type, type Static } from "@sinclair/typebox"',
			'import { Value } from "@sinclair/typebox/value"',
		],
		schema: schemaCode,
		type: `Static<typeof ${varName}>`,
		validate: "(data: unknown) => Value.Check(schema, data)",
	};
}
