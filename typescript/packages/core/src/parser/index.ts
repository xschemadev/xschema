/**
 * JSON Schema Parser
 * Converts JSON Schema to SchemaNode IR
 */

import type { JSONSchema } from "../schema/json-schema.js";
import type { SchemaNode, TypeGuard } from "../ir/nodes.js";
import { detectVersion, supportsRefSiblings } from "../schema/version.js";
import { normalizeDeep } from "../schema/normalizer.js";
import { resolveJsonPointer } from "../utils/json-pointer.js";
import { isEmptyObject } from "../utils/primitives.js";
import { createContext, type ParseContext } from "./context.js";
import { parseString, parseNumber } from "./primitives.js";
import { parseObject, parseArray } from "./collections.js";
import {
	parseAllOf,
	parseAnyOf,
	parseOneOf,
	parseNot,
	parseConditional,
} from "./composition.js";
import { parseLiteral, parseEnum } from "./values.js";

export type { ParseContext } from "./context.js";

/**
 * Parse a JSON Schema into SchemaNode IR
 * @param schema The JSON Schema to parse
 */
export function parse(schema: JSONSchema | boolean): SchemaNode {
	// Handle boolean schemas at root level
	if (typeof schema === "boolean") {
		return schema ? { kind: "any" } : { kind: "never" };
	}

	const version = detectVersion(schema);
	const normalized = normalizeDeep(schema, version);
	const ctx = createContext(normalized, version);
	return parseSchema(normalized, ctx);
}

/**
 * Internal schema parser - recursive
 */
export function parseSchema(
	schema: JSONSchema | boolean,
	ctx: ParseContext,
): SchemaNode {
	// Handle boolean schemas
	if (typeof schema === "boolean") {
		return schema ? { kind: "any" } : { kind: "never" };
	}

	// Fail on unsupported keywords that require evaluation semantics
	if (schema.unevaluatedItems !== undefined) {
		throw new Error("unevaluatedItems is not supported");
	}
	if (schema.unevaluatedProperties !== undefined) {
		const hasApplicators =
			schema.allOf ||
			schema.anyOf ||
			schema.oneOf ||
			schema.if ||
			schema.$ref ||
			schema.dependentSchemas ||
			schema.not;
		if (hasApplicators) {
			throw new Error("unevaluatedProperties with applicators is not supported");
		}
	}
	// Fail on dynamic/recursive ref keywords
	if (schema.$dynamicRef !== undefined) {
		throw new Error("$dynamicRef is not supported");
	}
	if (schema.$dynamicAnchor !== undefined) {
		throw new Error("$dynamicAnchor is not supported");
	}
	if (schema.$recursiveRef !== undefined) {
		throw new Error("$recursiveRef is not supported");
	}
	if (schema.$recursiveAnchor !== undefined) {
		throw new Error("$recursiveAnchor is not supported");
	}

	// Handle $ref
	if (schema.$ref) {
		return parseRef(schema, ctx);
	}

	// Handle not - check before composition since { not: {} } = never
	if (schema.not !== undefined) {
		if (typeof schema.not === "object" && isEmptyObject(schema.not)) {
			return { kind: "never" };
		}
		return parseNot(schema.not, ctx, parseSchema);
	}

	// Handle if/then/else
	if (schema.if !== undefined) {
		return parseConditional(schema, ctx, parseSchema);
	}

	// Check for composition operators
	const hasAllOf = schema.allOf && schema.allOf.length > 0;
	const hasAnyOf = schema.anyOf && schema.anyOf.length > 0;
	const hasOneOf = schema.oneOf && schema.oneOf.length > 0;

	// Check for base schema keywords alongside composition operators.
	// If we have base constraints (like additionalProperties/dependencies/etc), they must
	// be intersected with the composition result.
	const hasBaseSchema =
		schema.type !== undefined ||
		// object keywords
		schema.properties !== undefined ||
		schema.additionalProperties !== undefined ||
		schema.patternProperties !== undefined ||
		schema.required !== undefined ||
		schema.propertyNames !== undefined ||
		schema.minProperties !== undefined ||
		schema.maxProperties !== undefined ||
		schema.dependentRequired !== undefined ||
		schema.dependentSchemas !== undefined ||
		schema.dependencies !== undefined ||
		// array keywords
		schema.items !== undefined ||
		schema.prefixItems !== undefined ||
		schema.additionalItems !== undefined ||
		schema.minItems !== undefined ||
		schema.maxItems !== undefined ||
		schema.uniqueItems !== undefined ||
		schema.contains !== undefined ||
		// primitive keywords
		schema.minimum !== undefined ||
		schema.maximum !== undefined ||
		schema.exclusiveMinimum !== undefined ||
		schema.exclusiveMaximum !== undefined ||
		schema.multipleOf !== undefined ||
		schema.minLength !== undefined ||
		schema.maxLength !== undefined ||
		schema.pattern !== undefined ||
		schema.format !== undefined ||
		// value keywords
		schema.enum !== undefined ||
		schema.const !== undefined;

	if (hasAllOf || hasAnyOf || hasOneOf) {
		return parseComposition(schema, ctx, hasBaseSchema);
	}

	// Handle enum
	if (schema.enum !== undefined) {
		return parseEnum(schema);
	}

	// Handle const
	if (schema.const !== undefined) {
		return parseLiteral(schema);
	}

	// Handle type
	const type = schema.type;

	if (Array.isArray(type)) {
		// Multiple types - create union
		const variants = type.map((t) => {
			// Draft3: type array can contain inline schemas
			if (typeof t === "object" && t !== null) {
				return parseSchema(t as JSONSchema, ctx);
			}
			const typeSchema: JSONSchema = { ...schema, type: t };
			delete typeSchema.enum;
			return parseSchema(typeSchema, ctx);
		});

		if (variants.length === 0) return { kind: "never" };
		if (variants.length === 1) return variants[0]!;
		return { kind: "union", variants };
	}

	if (!type) {
		// No type - infer from keywords or return typeGuarded
		return parseTypeless(schema, ctx);
	}

	// Parse by type
	let result: SchemaNode;

	switch (type) {
		case "string":
			result = parseString(schema);
			break;
		case "number":
			result = parseNumber(schema, false);
			break;
		case "integer":
			result = parseNumber(schema, true);
			break;
		case "boolean":
			result = { kind: "boolean" };
			break;
		case "null":
			result = { kind: "null" };
			break;
		case "object":
			result = parseObject(schema, ctx, parseSchema);
			break;
		case "array":
			result = parseArray(schema, ctx, parseSchema);
			break;
		default:
			result = { kind: "any" };
	}

	// Handle nullable (OpenAPI 3.0)
	if (schema.nullable === true) {
		return { kind: "nullable", inner: result };
	}

	return result;
}

function parseRef(schema: JSONSchema, ctx: ParseContext): SchemaNode {
	const refPath = schema.$ref!;

	// Check cache
	const cached = ctx.refs.get(refPath);
	if (cached) {
		return {
			kind: "ref",
			path: refPath,
			resolved: cached,
		};
	}

	// Check for circular reference
	if (ctx.processing.has(refPath)) {
		// Circular ref - return a lazy any
		return {
			kind: "ref",
			path: refPath,
			resolved: { kind: "any" },
		};
	}

	// Resolve and parse
	ctx.processing.add(refPath);

	const resolved = resolveJsonPointer(refPath, ctx.rootSchema);
	const parsed = parseSchema(resolved, ctx);
	ctx.refs.set(refPath, parsed);
	ctx.processing.delete(refPath);

	// Handle sibling keywords in draft-2019-09+
	if (supportsRefSiblings(ctx.version)) {
		const siblingSchema = { ...schema };
		delete siblingSchema.$ref;
		delete siblingSchema.$id;
		delete siblingSchema.$anchor;
		delete siblingSchema.$defs;
		delete siblingSchema.definitions;
		delete siblingSchema.$schema;
		delete siblingSchema.$comment;
		delete siblingSchema.$dynamicRef;
		delete siblingSchema.$dynamicAnchor;

		if (Object.keys(siblingSchema).length > 0) {
			const siblingNode = parseSchema(siblingSchema, ctx);
			return {
				kind: "intersection",
				schemas: [
					{ kind: "ref", path: refPath, resolved: parsed },
					siblingNode,
				],
			};
		}
	}

	return {
		kind: "ref",
		path: refPath,
		resolved: parsed,
	};
}

function parseComposition(
	schema: JSONSchema,
	ctx: ParseContext,
	hasBaseSchema: boolean,
): SchemaNode {
	const parts: SchemaNode[] = [];

	// Add base schema constraints if present
	if (hasBaseSchema) {
		const baseSchema = { ...schema };
		delete baseSchema.allOf;
		delete baseSchema.anyOf;
		delete baseSchema.oneOf;
		parts.push(parseSchema(baseSchema, ctx));
	}

	// Add allOf (intersection)
	if (schema.allOf && schema.allOf.length > 0) {
		const allOfNode = parseAllOf(schema.allOf, ctx, parseSchema);
		// Flatten single-item intersection
		if (allOfNode.schemas.length === 1) {
			parts.push(allOfNode.schemas[0]!);
		} else {
			parts.push(allOfNode);
		}
	}

	// Add anyOf (union)
	if (schema.anyOf && schema.anyOf.length > 0) {
		const anyOfNode = parseAnyOf(schema.anyOf, ctx, parseSchema);
		// Flatten single-item union
		if (anyOfNode.variants.length === 1) {
			parts.push(anyOfNode.variants[0]!);
		} else {
			parts.push(anyOfNode);
		}
	}

	// Add oneOf (exactly one)
	if (schema.oneOf && schema.oneOf.length > 0) {
		const oneOfNode = parseOneOf(schema.oneOf, ctx, parseSchema);
		// Flatten single-item oneOf
		if (oneOfNode.schemas.length === 1) {
			parts.push(oneOfNode.schemas[0]!);
		} else {
			parts.push(oneOfNode);
		}
	}

	// Combine parts with intersection
	if (parts.length === 0) return { kind: "any" };
	if (parts.length === 1) return parts[0]!;

	return {
		kind: "intersection",
		schemas: parts,
	};
}

function parseTypeless(schema: JSONSchema, ctx: ParseContext): SchemaNode {
	const guards: TypeGuard[] = [];

	// Check for object keywords
	// Note: required as boolean (draft3 property-level) should NOT trigger object detection
	// Only array-style required (draft4+) indicates schema-level object constraints
	const hasObjectKeywords =
		schema.properties !== undefined ||
		schema.additionalProperties !== undefined ||
		schema.patternProperties !== undefined ||
		(Array.isArray(schema.required) && schema.required.length > 0) ||
		schema.propertyNames !== undefined ||
		schema.minProperties !== undefined ||
		schema.maxProperties !== undefined ||
		schema.dependentRequired !== undefined ||
		schema.dependentSchemas !== undefined ||
		schema.dependencies !== undefined;

	// Check for array keywords
	const hasArrayKeywords =
		schema.items !== undefined ||
		schema.prefixItems !== undefined ||
		schema.additionalItems !== undefined ||
		schema.minItems !== undefined ||
		schema.maxItems !== undefined ||
		schema.uniqueItems !== undefined ||
		schema.contains !== undefined;

	// Check for numeric keywords
	const hasNumericKeywords =
		schema.minimum !== undefined ||
		schema.maximum !== undefined ||
		schema.exclusiveMinimum !== undefined ||
		schema.exclusiveMaximum !== undefined ||
		schema.multipleOf !== undefined;

	// Check for string keywords
	const hasStringKeywords =
		schema.minLength !== undefined ||
		schema.maxLength !== undefined ||
		schema.pattern !== undefined;

	if (hasObjectKeywords) {
		guards.push({
			check: "object",
			schema: parseObject(schema, ctx, parseSchema),
		});
	}

	if (hasArrayKeywords) {
		guards.push({
			check: "array",
			schema: parseArray(schema, ctx, parseSchema),
		});
	}

	if (hasNumericKeywords) {
		guards.push({
			check: "number",
			schema: parseNumber(schema, false),
		});
	}

	if (hasStringKeywords) {
		guards.push({
			check: "string",
			schema: parseString(schema),
		});
	}

	if (guards.length === 0) {
		return { kind: "any" };
	}

	return {
		kind: "typeGuarded",
		guards,
	};
}
