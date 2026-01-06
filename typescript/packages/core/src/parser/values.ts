/**
 * Parsers for value types: literal, enum
 */

import type { JSONSchema } from "../schema/json-schema.js";
import type { LiteralNode, EnumNode } from "../ir/nodes.js";

export function parseLiteral(schema: JSONSchema): LiteralNode {
	return {
		kind: "literal",
		value: schema.const,
	};
}

export function parseEnum(schema: JSONSchema): EnumNode {
	return {
		kind: "enum",
		values: schema.enum!,
	};
}
