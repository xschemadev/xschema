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
			exclusiveMinimum: schema.exclusiveMinimum, // already normalized to number by Go CLI
			exclusiveMaximum: schema.exclusiveMaximum, // already normalized to number by Go CLI
			multipleOf: schema.multipleOf,
		},
	};
}
