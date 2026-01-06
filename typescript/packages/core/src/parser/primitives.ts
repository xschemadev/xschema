/**
 * Parsers for primitive types: string, number, boolean, null
 */

import type { JSONSchema } from "../schema/json-schema.js";
import type { StringNode, NumberNode } from "../ir/nodes.js";

export function parseString(schema: JSONSchema): StringNode {
	return {
		kind: "string",
		constraints: {
			minLength: schema.minLength,
			maxLength: schema.maxLength,
			pattern: schema.pattern,
		},
		format: schema.format,
	};
}

export function parseNumber(schema: JSONSchema, integer: boolean): NumberNode {
	return {
		kind: "number",
		integer,
		constraints: {
			minimum: schema.minimum,
			maximum: schema.maximum,
			exclusiveMinimum:
				typeof schema.exclusiveMinimum === "number"
					? schema.exclusiveMinimum
					: undefined,
			exclusiveMaximum:
				typeof schema.exclusiveMaximum === "number"
					? schema.exclusiveMaximum
					: undefined,
			multipleOf: schema.multipleOf,
		},
	};
}
