/**
 * JSON Schema Normalizer
 * Transforms legacy (draft3/4) keywords to modern equivalents before parsing
 */

import type { JSONSchema, JSONSchemaVersion } from "./json-schema.js";
import { usesBooleanExclusive } from "./version.js";

/**
 * Normalize a JSON Schema to use modern keywords
 * This simplifies the parser by handling legacy keywords in one place
 */
export function normalizeSchema(
	schema: JSONSchema,
	version: JSONSchemaVersion,
): JSONSchema {
	const result = { ...schema };

	// divisibleBy → multipleOf (draft3)
	if (result.divisibleBy !== undefined && result.multipleOf === undefined) {
		result.multipleOf = result.divisibleBy;
		delete result.divisibleBy;
	}

	// extends → allOf (draft3)
	if (result.extends !== undefined) {
		const extendsArr = Array.isArray(result.extends)
			? result.extends
			: [result.extends];
		result.allOf = [...(result.allOf || []), ...extendsArr];
		delete result.extends;
	}

	// disallow → not + anyOf (draft3)
	if (result.disallow !== undefined) {
		const disallowValue = result.disallow;
		delete result.disallow;

		let notSchema: JSONSchema;
		if (typeof disallowValue === "string") {
			notSchema = { type: disallowValue };
		} else if (Array.isArray(disallowValue)) {
			const allStrings = disallowValue.every((v) => typeof v === "string");
			if (allStrings) {
				notSchema = { type: disallowValue as string[] };
			} else {
				const schemas = disallowValue.map((v) =>
					typeof v === "string" ? { type: v } : v,
				) as JSONSchema[];
				notSchema = { anyOf: schemas };
			}
		} else {
			notSchema = disallowValue as JSONSchema;
		}

		// Combine with existing not using allOf
		if (result.not !== undefined) {
			result.allOf = [
				...(result.allOf || []),
				{ not: result.not },
				{ not: notSchema },
			];
			delete result.not;
		} else {
			result.not = notSchema;
		}
	}

	// exclusiveMinimum: boolean → exclusiveMinimum: number (draft4 → draft6+)
	if (usesBooleanExclusive(version)) {
		if (
			typeof result.exclusiveMinimum === "boolean" &&
			result.minimum !== undefined
		) {
			if (result.exclusiveMinimum) {
				result.exclusiveMinimum = result.minimum;
				delete result.minimum;
			} else {
				delete result.exclusiveMinimum;
			}
		}

		if (
			typeof result.exclusiveMaximum === "boolean" &&
			result.maximum !== undefined
		) {
			if (result.exclusiveMaximum) {
				result.exclusiveMaximum = result.maximum;
				delete result.maximum;
			} else {
				delete result.exclusiveMaximum;
			}
		}
	}

	// Handle unevaluatedProperties without applicators (treat as additionalProperties)
	if (result.unevaluatedProperties !== undefined) {
		const hasApplicators =
			result.allOf ||
			result.anyOf ||
			result.oneOf ||
			result.if ||
			result.$ref ||
			result.dependentSchemas ||
			result.not;

		if (hasApplicators) {
			// Cannot support with applicators - will be handled by parser throwing
		} else if (result.additionalProperties === undefined) {
			// Simple case: treat as additionalProperties
			result.additionalProperties = result.unevaluatedProperties;
			delete result.unevaluatedProperties;
		} else {
			// additionalProperties already set - unevaluatedProperties has nothing to check
			delete result.unevaluatedProperties;
		}
	}

	// dependencies → dependentRequired/dependentSchemas (draft7 → 2019-09+)
	// Note: We keep dependencies as-is and handle in parser since both forms are valid

	return result;
}

/**
 * Helper to normalize a schema that may be a boolean
 */
function normalizeNested(
	schema: JSONSchema | boolean,
	version: JSONSchemaVersion,
): JSONSchema | boolean {
	return typeof schema === "boolean" ? schema : normalizeDeep(schema, version);
}

/**
 * Recursively normalize all nested schemas
 */
export function normalizeDeep(
	schema: JSONSchema,
	version: JSONSchemaVersion,
): JSONSchema {
	const normalized = normalizeSchema(schema, version);

	// Recursively normalize nested schemas
	const result = { ...normalized };

	if (result.properties) {
		result.properties = Object.fromEntries(
			Object.entries(result.properties).map(([k, v]) => [
				k,
				normalizeNested(v, version),
			]),
		);
	}

	if (result.additionalProperties && typeof result.additionalProperties === "object") {
		result.additionalProperties = normalizeDeep(result.additionalProperties, version);
	}

	if (result.patternProperties) {
		result.patternProperties = Object.fromEntries(
			Object.entries(result.patternProperties).map(([k, v]) => [
				k,
				normalizeNested(v, version),
			]),
		);
	}

	if (result.items !== undefined && typeof result.items !== "boolean") {
		if (Array.isArray(result.items)) {
			result.items = result.items.map((i) => normalizeNested(i, version));
		} else {
			result.items = normalizeDeep(result.items, version);
		}
	}

	if (result.prefixItems) {
		result.prefixItems = result.prefixItems.map((i) => normalizeNested(i, version));
	}

	if (result.additionalItems && typeof result.additionalItems === "object") {
		result.additionalItems = normalizeDeep(result.additionalItems, version);
	}

	if (result.allOf) {
		result.allOf = result.allOf.map((s) => normalizeNested(s, version));
	}

	if (result.anyOf) {
		result.anyOf = result.anyOf.map((s) => normalizeNested(s, version));
	}

	if (result.oneOf) {
		result.oneOf = result.oneOf.map((s) => normalizeNested(s, version));
	}

	if (result.not && typeof result.not !== "boolean") {
		result.not = normalizeDeep(result.not, version);
	}

	if (result.if && typeof result.if !== "boolean") {
		result.if = normalizeDeep(result.if, version);
	}

	if (result.then && typeof result.then !== "boolean") {
		result.then = normalizeDeep(result.then, version);
	}

	if (result.else && typeof result.else !== "boolean") {
		result.else = normalizeDeep(result.else, version);
	}

	if (result.contains && typeof result.contains !== "boolean") {
		result.contains = normalizeDeep(result.contains, version);
	}

	if (result.propertyNames && typeof result.propertyNames !== "boolean") {
		result.propertyNames = normalizeDeep(result.propertyNames, version);
	}

	if (result.$defs) {
		result.$defs = Object.fromEntries(
			Object.entries(result.$defs).map(([k, v]) => [
				k,
				normalizeNested(v, version),
			]),
		);
	}

	if (result.definitions) {
		result.definitions = Object.fromEntries(
			Object.entries(result.definitions).map(([k, v]) => [
				k,
				normalizeNested(v, version),
			]),
		);
	}

	if (result.dependentSchemas) {
		result.dependentSchemas = Object.fromEntries(
			Object.entries(result.dependentSchemas).map(([k, v]) => [
				k,
				normalizeNested(v, version),
			]),
		);
	}

	if (result.dependencies) {
		result.dependencies = Object.fromEntries(
			Object.entries(result.dependencies).map(([k, v]) => {
				// draft3 allows string shorthand: { "bar": "foo" }
				if (typeof v === "string") {
					return [k, [v]];
				}
				if (Array.isArray(v)) {
					return [k, v];
				}
				return [k, normalizeNested(v as JSONSchema | boolean, version)];
			}),
		);
	}

	return result;
}
