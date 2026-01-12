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
import { isValidationEnabled, type ParseContext } from "./context.js";

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
	const validation = isValidationEnabled(ctx);

	// Collect required properties (only if validation enabled)
	if (validation && Array.isArray(schema.required)) {
		for (const key of schema.required) {
			requiredSet.add(key);
		}
	}

	// Draft3 style: check for required: true on individual properties (only if validation enabled)
	const schemaProps = schema.properties || {};
	if (validation) {
		for (const [key, propSchema] of Object.entries(schemaProps)) {
			if (
				propSchema &&
				typeof propSchema === "object" &&
				(propSchema as Record<string, unknown>).required === true
			) {
				requiredSet.add(key);
			}
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
	// Only if validation enabled
	if (validation) {
		for (const key of requiredSet) {
			if (!properties.has(key)) {
				properties.set(key, {
					schema: { kind: "any" },
					required: true,
				});
			}
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
	let additionalProperties: SchemaNode | boolean = true;
	if (schema.additionalProperties === false) {
		additionalProperties = false;
	} else if (
		schema.additionalProperties &&
		typeof schema.additionalProperties === "object"
	) {
		additionalProperties = parseSchema(schema.additionalProperties, ctx);
	}

	// Parse propertyNames
	const propertyNames =
		schema.propertyNames !== undefined
			? parseSchema(schema.propertyNames, ctx)
			: undefined;

	// Parse dependencies
	const dependencies = new Map<string, Dependency>();

	// Handle dependentRequired (validation keyword - only if enabled)
	if (validation && schema.dependentRequired) {
		for (const [prop, deps] of Object.entries(schema.dependentRequired)) {
			dependencies.set(prop, {
				kind: "property",
				requiredProperties: deps,
			});
		}
	}

	// Handle dependentSchemas (applicator - always enabled)
	if (schema.dependentSchemas) {
		for (const [prop, depSchema] of Object.entries(schema.dependentSchemas)) {
			dependencies.set(prop, {
				kind: "schema",
				schema: parseSchema(depSchema, ctx),
			});
		}
	}

	// Handle legacy dependencies (draft3-7)
	if (schema.dependencies) {
		for (const [prop, dep] of Object.entries(schema.dependencies)) {
			// draft3 allows a single string as a shorthand for a 1-element dependency list
			if (typeof dep === "string") {
				// Property dependencies are validation - only if enabled
				if (validation) {
					dependencies.set(prop, {
						kind: "property",
						requiredProperties: [dep],
					});
				}
				continue;
			}

			if (Array.isArray(dep)) {
				// Property dependencies are validation - only if enabled
				if (validation) {
					dependencies.set(prop, {
						kind: "property",
						requiredProperties: dep,
					});
				}
				continue;
			}

			// Schema dependencies are applicators - always enabled
			dependencies.set(prop, {
				kind: "schema",
				schema: parseSchema(dep, ctx),
			});
		}
	}

	return {
		kind: "object",
		properties,
		additionalProperties,
		patternProperties,
		propertyNames,
		minProperties: validation ? schema.minProperties : undefined,
		maxProperties: validation ? schema.maxProperties : undefined,
		dependencies,
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
		};
	}

	return {
		kind: "array",
		items: itemSchema,
		constraints: buildArrayConstraints(schema, ctx, parseSchema),
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
	};
}

function buildArrayConstraints(
	schema: JSONSchema,
	ctx: ParseContext,
	parseSchema: ParseSchemaFn,
): ArrayNode["constraints"] {
	const validation = isValidationEnabled(ctx);
	let contains: ContainsConstraint | undefined;

	// contains is an applicator, but minContains/maxContains are validation
	if (schema.contains !== undefined) {
		contains = {
			schema: parseSchema(schema.contains, ctx),
			minContains: validation ? (schema.minContains ?? 1) : 1,
			maxContains: validation ? schema.maxContains : undefined,
		};
	}

	return {
		minItems: validation ? schema.minItems : undefined,
		maxItems: validation ? schema.maxItems : undefined,
		uniqueItems: validation ? schema.uniqueItems : undefined,
		contains,
	};
}
