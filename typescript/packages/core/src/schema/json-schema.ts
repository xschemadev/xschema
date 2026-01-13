/**
 * JSON Schema type definition
 * Covers draft-3, draft-4, draft-6, draft-7, draft-2019-09, draft-2020-12, and OpenAPI 3.0
 */

export interface JSONSchema {
	// Meta
	$schema?: string;
	$id?: string;
	$ref?: string;
	$defs?: Record<string, JSONSchema | boolean>;
	definitions?: Record<string, JSONSchema | boolean>;
	$comment?: string;
	$anchor?: string;
	$dynamicRef?: string;
	$dynamicAnchor?: string;
	$vocabulary?: Record<string, boolean>;

	// Type
	type?: string | string[];
	enum?: unknown[];
	const?: unknown;

	// Composition (boolean schemas allowed in draft6+)
	anyOf?: (JSONSchema | boolean)[];
	oneOf?: (JSONSchema | boolean)[];
	allOf?: (JSONSchema | boolean)[];
	not?: JSONSchema | boolean;

	// Object (boolean schemas allowed in draft6+)
	properties?: Record<string, JSONSchema | boolean>;
	required?: string[] | boolean; // boolean for draft3 property-level required
	additionalProperties?: boolean | JSONSchema;
	patternProperties?: Record<string, JSONSchema | boolean>;
	propertyNames?: JSONSchema | boolean;
	minProperties?: number;
	maxProperties?: number;

	// Array (boolean schemas allowed in draft6+)
	items?: JSONSchema | (JSONSchema | boolean)[] | boolean;
	prefixItems?: (JSONSchema | boolean)[];
	additionalItems?: boolean | JSONSchema;
	minItems?: number;
	maxItems?: number;
	uniqueItems?: boolean;
	contains?: JSONSchema | boolean;
	minContains?: number;
	maxContains?: number;

	// String
	minLength?: number;
	maxLength?: number;
	pattern?: string;
	format?: string;

	// Number
	minimum?: number;
	maximum?: number;
	exclusiveMinimum?: number | boolean; // boolean for draft4
	exclusiveMaximum?: number | boolean; // boolean for draft4
	multipleOf?: number;

	// Metadata
	description?: string;
	default?: unknown;
	title?: string;

	// OpenAPI 3.0
	nullable?: boolean;
	readOnly?: boolean;

	// Conditional (boolean schemas allowed in draft6+)
	if?: JSONSchema | boolean;
	then?: JSONSchema | boolean;
	else?: JSONSchema | boolean;
	dependentSchemas?: Record<string, JSONSchema | boolean>;
	dependentRequired?: Record<string, string[]>;

	// Legacy (draft3-7) - combined into dependentRequired/dependentSchemas in 2019-09+
	dependencies?: Record<string, string[] | JSONSchema | boolean>;

	// Legacy draft3
	disallow?: string | string[] | JSONSchema[];
	extends?: JSONSchema | JSONSchema[];
	divisibleBy?: number;

	// Unevaluated
	unevaluatedItems?: JSONSchema | boolean;
	unevaluatedProperties?: JSONSchema | boolean;

	// Allow additional properties for extensibility
	[key: string]: unknown;
}
