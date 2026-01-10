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
		imports: ['* as v from "valibot"'],
		schema: schemaCode,
		type: `v.InferOutput<typeof ${varName}>`,
		validate: "(data: unknown) => safeParse(schema, data).success",
		validateImports: ['{ safeParse } from "valibot"'],
	};
}
