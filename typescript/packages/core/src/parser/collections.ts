/**
 * Parsers for collection types: object, array, tuple
 */

import type { JSONSchema } from "../schema/json-schema.js";
import type {
	ObjectNode,
	ArrayNode,
	TupleNode,
	SchemaNode,
	PropertyDef,
	PatternPropertyDef,
	Dependency,
	ContainsConstraint,
} from "../ir/nodes.js";
import { type ParseContext } from "./context.js";

type ParseSchemaFn = (
	schema: JSONSchema | boolean,
	ctx: ParseContext,
) => SchemaNode;

export function parseObject(
	schema: JSONSchema,
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): ObjectNode {
	const properties = new Map<string, PropertyDef>();
	const requiredSet = new Set<string>();

	// Collect required properties
	if (Array.isArray(schema.required)) {
		for (const key of schema.required) {
			requiredSet.add(key);
		}
	}

	// Draft3 style: check for required: true on individual properties
	const schemaProps = schema.properties || {};
	for (const [key, propSchema] of Object.entries(schemaProps)) {
		if (
			propSchema &&
			typeof propSchema === "object" &&
			(propSchema as Record<string, unknown>).required === true
		) {
			requiredSet.add(key);
		}
	}

	// Parse properties
	for (const key of Object.getOwnPropertyNames(schemaProps)) {
		const propSchema = schemaProps[key];
		if (propSchema === undefined) continue;

		properties.set(key, {
			schema: parseSchema(propSchema, ctx),
			required: requiredSet.has(key),
		});
	}

	// Add required properties that don't have schemas (must exist, any type)
	for (const key of requiredSet) {
		if (!properties.has(key)) {
			properties.set(key, {
				schema: { kind: "any" },
				required: true,
			});
		}
	}

	// Parse pattern properties
	const patternProperties: PatternPropertyDef[] = [];
	if (schema.patternProperties) {
		for (const [pattern, patternSchema] of Object.entries(
			schema.patternProperties,
		)) {
			patternProperties.push({
				pattern,
				schema: parseSchema(patternSchema, ctx),
			});
		}
	}

	// Parse additionalProperties
	// Keep track of whether it was explicitly set to true vs not specified
	let additionalProperties: SchemaNode | boolean | undefined;
	if (schema.additionalProperties === false) {
		additionalProperties = false;
	} else if (schema.additionalProperties === true) {
		additionalProperties = true; // Explicitly true
	} else if (
		schema.additionalProperties &&
		typeof schema.additionalProperties === "object"
	) {
		additionalProperties = parseSchema(schema.additionalProperties, ctx);
	}
	// If not specified, leave as undefined (not defaulting to true)

	// Parse propertyNames
	const propertyNames =
		schema.propertyNames !== undefined
			? parseSchema(schema.propertyNames, ctx)
			: undefined;

	// Parse dependencies
	const dependencies = new Map<string, Dependency>();

	// Handle dependentRequired
	if (schema.dependentRequired) {
		for (const [prop, deps] of Object.entries(schema.dependentRequired)) {
			dependencies.set(prop, {
				kind: "property",
				requiredProperties: deps,
			});
		}
	}

	// Handle dependentSchemas
	if (schema.dependentSchemas) {
		for (const [prop, depSchema] of Object.entries(schema.dependentSchemas)) {
			dependencies.set(prop, {
				kind: "schema",
				schema: parseSchema(depSchema, ctx),
			});
		}
	}

	// Note: legacy dependencies should be normalized by the Go CLI to draft2020-12 format
	// (dependentRequired for property deps, dependentSchemas for schema deps).
	// If we encounter the legacy 'dependencies' keyword, it indicates incomplete bundling.
	if (schema.dependencies) {
		throw new Error(
			`Encountered legacy 'dependencies' keyword - schemas must be normalized to draft2020-12 by the Go CLI. ` +
			`Run the schema through xschema generate to normalize the format.`
		);
	}

	// Parse unevaluatedProperties (only valid for standalone use - Go CLI blocks applicator combinations)
	let unevaluatedProperties: SchemaNode | false | undefined;
	if (schema.unevaluatedProperties === false) {
		unevaluatedProperties = false;
	} else if (
		schema.unevaluatedProperties !== undefined &&
		typeof schema.unevaluatedProperties === "object"
	) {
		unevaluatedProperties = parseSchema(schema.unevaluatedProperties, ctx);
	}

	return {
		kind: "object",
		properties,
		additionalProperties,
		patternProperties,
		propertyNames,
		minProperties: schema.minProperties,
		maxProperties: schema.maxProperties,
		dependencies,
		unevaluatedProperties,
	};
}

export function parseArray(
	schema: JSONSchema,
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): ArrayNode | TupleNode {
	const prefixItems = schema.prefixItems;
	const items = schema.items;

	// Check for tuple (prefixItems or items as array)
	if (prefixItems && Array.isArray(prefixItems)) {
		return parseTuple(schema, ctx, parseSchema);
	}

	if (Array.isArray(items)) {
		return parseTupleFromItems(schema, ctx, parseSchema);
	}

	// Regular array
	let itemSchema: SchemaNode = { kind: "any" };
	if (items && typeof items === "object") {
		itemSchema = parseSchema(items, ctx);
	} else if (items === false) {
		// items: false means empty tuple only
		return {
			kind: "tuple",
			prefixItems: [],
			restItems: false,
			constraints: buildArrayConstraints(schema, ctx, parseSchema),
			unevaluatedItems: parseUnevaluatedItems(schema, ctx, parseSchema),
		};
	}

	return {
		kind: "array",
		items: itemSchema,
		constraints: buildArrayConstraints(schema, ctx, parseSchema),
		unevaluatedItems: parseUnevaluatedItems(schema, ctx, parseSchema),
	};
}

function parseTuple(
	schema: JSONSchema,
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): TupleNode {
	const prefixItems = schema.prefixItems!;
	const items = schema.items;
	const additionalItems = schema.additionalItems;

	const parsedPrefixItems = prefixItems.map((item) => parseSchema(item, ctx));

	// Determine rest items
	let restItems: SchemaNode | false = { kind: "any" };

	if (items === false || additionalItems === false) {
		restItems = false;
	} else if (items && typeof items === "object" && !Array.isArray(items)) {
		restItems = parseSchema(items, ctx);
	} else if (additionalItems && typeof additionalItems === "object") {
		restItems = parseSchema(additionalItems, ctx);
	}

	return {
		kind: "tuple",
		prefixItems: parsedPrefixItems,
		restItems,
		constraints: buildArrayConstraints(schema, ctx, parseSchema),
		unevaluatedItems: parseUnevaluatedItems(schema, ctx, parseSchema),
	};
}

function parseTupleFromItems(
	schema: JSONSchema,
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): TupleNode {
	const items = schema.items as JSONSchema[];
	const additionalItems = schema.additionalItems;

	const parsedItems = items.map((item) => parseSchema(item, ctx));

	// Determine rest items
	let restItems: SchemaNode | false = { kind: "any" };

	if (additionalItems === false) {
		restItems = false;
	} else if (additionalItems && typeof additionalItems === "object") {
		restItems = parseSchema(additionalItems, ctx);
	}

	return {
		kind: "tuple",
		prefixItems: parsedItems,
		restItems,
		constraints: buildArrayConstraints(schema, ctx, parseSchema),
		unevaluatedItems: parseUnevaluatedItems(schema, ctx, parseSchema),
	};
}

function buildArrayConstraints(
	schema: JSONSchema,
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): ArrayNode["constraints"] {
	let contains: ContainsConstraint | undefined;

	if (schema.contains !== undefined) {
		contains = {
			schema: parseSchema(schema.contains, ctx),
			minContains: schema.minContains ?? 1,
			maxContains: schema.maxContains,
		};
	}

	return {
		minItems: schema.minItems,
		maxItems: schema.maxItems,
		uniqueItems: schema.uniqueItems,
		contains,
	};
}

function parseUnevaluatedItems(
	schema: JSONSchema,
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): SchemaNode | false | undefined {
	if (schema.unevaluatedItems === false) {
		return false;
	}
	if (
		schema.unevaluatedItems !== undefined &&
		typeof schema.unevaluatedItems === "object"
	) {
		return parseSchema(schema.unevaluatedItems, ctx);
	}
	return undefined;
}
