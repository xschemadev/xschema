/**
 * Parsers for composition types: allOf, anyOf, oneOf, not
 */

import type { JSONSchema } from "../schema/json-schema.js";
import type {
	IntersectionNode,
	UnionNode,
	OneOfNode,
	NotNode,
	ConditionalNode,
	SchemaNode,
} from "../ir/nodes.js";
import type { ParseContext } from "./context.js";

type ParseSchemaFn = (
	schema: JSONSchema | boolean,
	ctx: ParseContext,
) => SchemaNode;

export function parseAllOf(
	schemas: (JSONSchema | boolean)[],
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): IntersectionNode {
	return {
		kind: "intersection",
		schemas: schemas.map((s) => parseSchema(s, ctx)),
	};
}

export function parseAnyOf(
	schemas: (JSONSchema | boolean)[],
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): UnionNode {
	return {
		kind: "union",
		variants: schemas.map((s) => parseSchema(s, ctx)),
	};
}

export function parseOneOf(
	schemas: (JSONSchema | boolean)[],
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): OneOfNode {
	return {
		kind: "oneOf",
		schemas: schemas.map((s) => parseSchema(s, ctx)),
	};
}

export function parseNot(
	schema: JSONSchema | boolean,
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): NotNode {
	return {
		kind: "not",
		schema: parseSchema(schema, ctx),
	};
}

export function parseConditional(
	schema: JSONSchema,
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): ConditionalNode {
	return {
		kind: "conditional",
		if: parseSchema(schema.if!, ctx),
		then:
			schema.then !== undefined ? parseSchema(schema.then, ctx) : undefined,
		else:
			schema.else !== undefined ? parseSchema(schema.else, ctx) : undefined,
	};
}
