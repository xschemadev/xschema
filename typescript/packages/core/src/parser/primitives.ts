/**
 * Parsers for primitive types: string, number, boolean, null
 */

import type { JSONSchema } from "../schema/json-schema.js";
import type { StringNode, NumberNode } from "../ir/nodes.js";
import {
	isValidationEnabled,
	isFormatEnabled,
	type ParseContext,
} from "./context.js";

export function parseString(schema: JSONSchema, ctx: ParseContext): StringNode {
	const validation = isValidationEnabled(ctx);
	const formatEnabled = isFormatEnabled(ctx);

	return {
		kind: "string",
		constraints: {
			minLength: validation ? schema.minLength : undefined,
			maxLength: validation ? schema.maxLength : undefined,
			pattern: validation ? schema.pattern : undefined,
		},
		format: formatEnabled ? schema.format : undefined,
	};
}

export function parseNumber(
	schema: JSONSchema,
	integer: boolean,
	ctx: ParseContext,
): NumberNode {
	const validation = isValidationEnabled(ctx);

	return {
		kind: "number",
		integer,
		constraints: {
			minimum: validation ? schema.minimum : undefined,
			maximum: validation ? schema.maximum : undefined,
			exclusiveMinimum: validation
				? typeof schema.exclusiveMinimum === "number"
					? schema.exclusiveMinimum
					: undefined
				: undefined,
			exclusiveMaximum: validation
				? typeof schema.exclusiveMaximum === "number"
					? schema.exclusiveMaximum
					: undefined
				: undefined,
			multipleOf: validation ? schema.multipleOf : undefined,
		},
	};
}
