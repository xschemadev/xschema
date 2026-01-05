/**
 * JSON Schema to Zod code converter
 * Based on Zod 4's fromJSONSchema approach but generates code strings
 * Also edited it to get to 100% compliance with json schema test suite
 */

export interface JSONSchema {
	$schema?: string;
	$ref?: string;
	$defs?: Record<string, JSONSchema>;
	definitions?: Record<string, JSONSchema>;
	$id?: string;
	$comment?: string;

	// Type
	type?: string | string[];
	enum?: unknown[];
	const?: unknown;

	// Composition
	anyOf?: JSONSchema[];
	oneOf?: JSONSchema[];
	allOf?: JSONSchema[];
	not?: JSONSchema;

	// Object
	properties?: Record<string, JSONSchema>;
	required?: string[];
	additionalProperties?: boolean | JSONSchema;
	patternProperties?: Record<string, JSONSchema>;
	propertyNames?: JSONSchema;
	minProperties?: number;
	maxProperties?: number;

	// Array
	items?: JSONSchema | JSONSchema[];
	prefixItems?: JSONSchema[];
	additionalItems?: boolean | JSONSchema;
	minItems?: number;
	maxItems?: number;
	uniqueItems?: boolean;
	contains?: JSONSchema;
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
	exclusiveMinimum?: number | boolean;
	exclusiveMaximum?: number | boolean;
	multipleOf?: number;

	// Metadata
	description?: string;
	default?: unknown;
	title?: string;

	// OpenAPI
	nullable?: boolean;
	readOnly?: boolean;

	// Conditional (will throw)
	if?: JSONSchema;
	then?: JSONSchema;
	else?: JSONSchema;
	dependentSchemas?: Record<string, JSONSchema>;
	dependentRequired?: Record<string, string[]>;

	// Legacy (draft3-7) - combined into dependentRequired/dependentSchemas in 2019-09+
	dependencies?: Record<string, string[] | JSONSchema>;

	// Legacy draft3
	disallow?: string | string[] | JSONSchema[];  // inverse of type
	extends?: JSONSchema | JSONSchema[];  // like allOf
	divisibleBy?: number;  // alias for multipleOf

	// Unevaluated (partially supported - simple cases only)
	unevaluatedItems?: JSONSchema | boolean;
	unevaluatedProperties?: JSONSchema | boolean;

	[key: string]: unknown;
}

type JSONSchemaVersion = "draft-2020-12" | "draft-7" | "draft-4" | "openapi-3.0";

interface ConversionContext {
	version: JSONSchemaVersion;
	defs: Record<string, JSONSchema>;
	refs: Map<string, string>;
	processing: Set<string>;
	rootSchema: JSONSchema;
}

function detectVersion(schema: JSONSchema): JSONSchemaVersion {
	const $schema = schema.$schema;

	if ($schema === "https://json-schema.org/draft/2020-12/schema") {
		return "draft-2020-12";
	}
	if ($schema?.includes("draft-07")) {
		return "draft-7";
	}
	if ($schema?.includes("draft-04")) {
		return "draft-4";
	}

	return "draft-2020-12";
}

function escapeString(str: string): string {
	return JSON.stringify(str);
}

/**
 * Helper to create deep equality check for const/enum with objects/arrays
 * Uses JSON.stringify with sorted keys for property order independence
 */
function createDeepEqualityCheck(value: unknown): string {
	const sortedValue = JSON.stringify(value, Object.keys(value as object).sort());
	return `JSON.stringify(val, Object.keys(val as object).sort()) === ${JSON.stringify(sortedValue)}`;
}

/**
 * Check if a value is a primitive (string, number, boolean, null)
 */
function isPrimitive(value: unknown): boolean {
	return value === null ||
		typeof value === 'string' ||
		typeof value === 'number' ||
		typeof value === 'boolean';
}

function resolveRef(ref: string, ctx: ConversionContext): JSONSchema {
	if (!ref.startsWith("#")) {
		throw new Error(`External $ref is not supported: ${ref}`);
	}

	const path = ref.slice(1).split("/").filter(Boolean);

	// Handle root reference "#"
	if (path.length === 0) {
		return ctx.rootSchema;
	}

	// Traverse the full JSON pointer path from root
	let current: unknown = ctx.rootSchema;
	for (const segment of path) {
		const decoded = segment.replace(/~1/g, "/").replace(/~0/g, "~");
		if (current && typeof current === "object" && decoded in current) {
			current = (current as Record<string, unknown>)[decoded];
		} else {
			throw new Error(`Reference not found: ${ref}`);
		}
	}

	return current as JSONSchema;
}

function convertSchema(schema: JSONSchema | boolean, ctx: ConversionContext): string {
	// Handle boolean schemas
	if (typeof schema === "boolean") {
		return schema ? "z.any()" : "z.never()";
	}

	// Handle unsupported features - throw explicit errors
	if (schema.unevaluatedItems !== undefined) {
		throw new Error("unevaluatedItems is not supported");
	}
	// unevaluatedProperties with applicators requires runtime tracking - not supported
	// But simple case (no applicators) is equivalent to additionalProperties
	if (schema.unevaluatedProperties !== undefined) {
		const hasApplicators = schema.allOf || schema.anyOf || schema.oneOf ||
			schema.if || schema.$ref || schema.dependentSchemas || schema.not;
		if (hasApplicators) {
			throw new Error("unevaluatedProperties with applicators is not supported");
		}
		// If additionalProperties is already present, it "evaluates" all additional properties
		// So unevaluatedProperties has nothing left to check - just ignore it
		if (schema.additionalProperties !== undefined) {
			const converted = { ...schema };
			delete converted.unevaluatedProperties;
			return convertSchema(converted, ctx);
		}
		// Simple case: treat as additionalProperties
		// Convert by moving unevaluatedProperties to additionalProperties
		const converted = { ...schema };
		converted.additionalProperties = schema.unevaluatedProperties;
		delete converted.unevaluatedProperties;
		return convertSchema(converted, ctx);
	}

	// Handle draft3 legacy keywords
	// divisibleBy is alias for multipleOf
	if (schema.divisibleBy !== undefined && schema.multipleOf === undefined) {
		const converted = { ...schema, multipleOf: schema.divisibleBy };
		delete converted.divisibleBy;
		return convertSchema(converted, ctx);
	}

	// extends is like allOf - merge with existing allOf if present
	if (schema.extends !== undefined) {
		const converted = { ...schema };
		const extendsSchemas = Array.isArray(schema.extends) ? schema.extends : [schema.extends];
		converted.allOf = [...(schema.allOf || []), ...extendsSchemas];
		delete converted.extends;
		return convertSchema(converted, ctx);
	}

	// disallow is inverse of type - use "not" with type
	if (schema.disallow !== undefined) {
		const disallowValue = schema.disallow;
		const converted = { ...schema };
		delete converted.disallow;

		// disallow can be a string, array of strings, array of schemas, or mixed
		let notSchema: JSONSchema;
		if (typeof disallowValue === 'string') {
			notSchema = { type: disallowValue };
		} else if (Array.isArray(disallowValue)) {
			// Check if it's an array of strings only
			const allStrings = disallowValue.every(v => typeof v === 'string');
			if (allStrings) {
				notSchema = { type: disallowValue as string[] };
			} else {
				// Mixed array or array of schemas - normalize strings to {type: x}
				const schemas = disallowValue.map(v =>
					typeof v === 'string' ? { type: v } : v
				) as JSONSchema[];
				notSchema = { anyOf: schemas };
			}
		} else {
			notSchema = disallowValue as JSONSchema;
		}

		// Combine with existing schema using intersection
		// The result must match converted AND not match notSchema
		const notCode = convertSchema({ not: notSchema }, ctx);
		const baseCode = convertSchema(converted, ctx);

		// If base is just z.any(), return the not constraint alone
		if (baseCode === "z.any()") {
			return notCode;
		}

		return `z.intersection(${baseCode}, ${notCode})`;
	}

	// Handle $ref
	if (schema.$ref) {
		const refPath = schema.$ref;

		let refCode: string;
		if (ctx.refs.has(refPath)) {
			refCode = ctx.refs.get(refPath)!;
		} else if (ctx.processing.has(refPath)) {
			// Circular reference - use lazy
			// For now, return z.any() for circular refs
			refCode = "z.lazy(() => z.any())";
		} else {
			ctx.processing.add(refPath);
			const resolved = resolveRef(refPath, ctx);
			refCode = convertSchema(resolved, ctx);
			ctx.refs.set(refPath, refCode);
			ctx.processing.delete(refPath);
		}

		// In JSON Schema 2019-09+, $ref can have sibling keywords that must also apply
		// In draft-07 and earlier, $ref overrides all siblings (return ref only)
		const supportsRefSiblings = ctx.version === "draft-2020-12";

		if (supportsRefSiblings) {
			// Check for sibling keywords (excluding metadata/structural keywords)
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

			const siblingKeys = Object.keys(siblingSchema);
			if (siblingKeys.length > 0) {
				// Has sibling keywords - intersect ref with sibling constraints
				const siblingCode = convertSchema(siblingSchema, ctx);
				return `z.intersection(${refCode}, ${siblingCode})`;
			}
		}

		return refCode;
	}

	// Handle not - limited support
	if (schema.not !== undefined) {
		// { not: {} } represents never
		if (typeof schema.not === "object" && Object.keys(schema.not).length === 0) {
			return "z.never()";
		}
		// For other "not" schemas, we can use refine
		const notSchema = convertSchema(schema.not, ctx);
		return `z.any().refine((val) => !${notSchema}.safeParse(val).success, { message: "Value must not match schema" })`;
	}

	// Handle if/then/else - conditional schema validation
	if (schema.if !== undefined) {
		const ifSchema = convertSchema(schema.if, ctx);
		const thenSchema = schema.then !== undefined ? convertSchema(schema.then, ctx) : null;
		const elseSchema = schema.else !== undefined ? convertSchema(schema.else, ctx) : null;

		if (thenSchema && elseSchema) {
			return `z.any().superRefine((val, ctx) => {
        const ifResult = ${ifSchema}.safeParse(val);
        if (ifResult.success) {
          const thenResult = ${thenSchema}.safeParse(val);
          if (!thenResult.success) {
            thenResult.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        } else {
          const elseResult = ${elseSchema}.safeParse(val);
          if (!elseResult.success) {
            elseResult.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      })`;
		} else if (thenSchema) {
			return `z.any().superRefine((val, ctx) => {
        const ifResult = ${ifSchema}.safeParse(val);
        if (ifResult.success) {
          const thenResult = ${thenSchema}.safeParse(val);
          if (!thenResult.success) {
            thenResult.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      })`;
		} else if (elseSchema) {
			return `z.any().superRefine((val, ctx) => {
        const ifResult = ${ifSchema}.safeParse(val);
        if (!ifResult.success) {
          const elseResult = ${elseSchema}.safeParse(val);
          if (!elseResult.success) {
            elseResult.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      })`;
		}
		// if without then/else - just return any (if alone has no effect)
		return "z.any()";
	}

	// Handle composition - but check for base schema properties too
	const hasBaseSchema = schema.type !== undefined ||
		schema.properties !== undefined ||
		schema.items !== undefined ||
		schema.minimum !== undefined ||
		schema.maximum !== undefined ||
		schema.minLength !== undefined ||
		schema.maxLength !== undefined ||
		schema.pattern !== undefined ||
		schema.enum !== undefined ||
		schema.const !== undefined;

	// Handle composition operators - they can appear together and must all be satisfied
	const hasAllOf = schema.allOf && schema.allOf.length > 0;
	const hasAnyOf = schema.anyOf && schema.anyOf.length > 0;
	const hasOneOf = schema.oneOf && schema.oneOf.length > 0;

	if (hasAllOf || hasAnyOf || hasOneOf) {
		const parts: string[] = [];

		// Add base schema constraints if present
		if (hasBaseSchema) {
			const baseSchema = { ...schema };
			delete baseSchema.allOf;
			delete baseSchema.anyOf;
			delete baseSchema.oneOf;
			parts.push(convertSchema(baseSchema, ctx));
		}

		// Add allOf (intersection of all)
		if (hasAllOf) {
			parts.push(convertAllOf(schema.allOf!, ctx));
		}

		// Add anyOf (union - at least one must match)
		if (hasAnyOf) {
			parts.push(convertAnyOf(schema.anyOf!, ctx));
		}

		// Add oneOf (exactly one must match)
		if (hasOneOf) {
			parts.push(convertOneOf(schema.oneOf!, ctx));
		}

		// Combine all parts with intersection
		if (parts.length === 1) {
			return parts[0]!;
		}

		let result = parts[0]!;
		for (let i = 1; i < parts.length; i++) {
			result = `z.intersection(${result}, ${parts[i]})`;
		}
		return result;
	}

	// Handle enum
	if (schema.enum !== undefined) {
		return convertEnum(schema.enum, schema, ctx);
	}

	// Handle const
	if (schema.const !== undefined) {
		const constValue = schema.const;

		// For primitives, use z.literal
		if (isPrimitive(constValue)) {
			return `z.literal(${JSON.stringify(constValue)})`;
		}

		// For objects/arrays, use refine with deep equality
		const isArray = Array.isArray(constValue);
		const baseType = isArray ? 'z.array(z.any())' : 'z.object({}).passthrough()';
		const sortedJson = JSON.stringify(constValue, Object.keys(constValue as object).sort());

		return `${baseType}.refine((val) => JSON.stringify(val, Object.keys(val as object).sort()) === ${JSON.stringify(sortedJson)}, { message: "Value must equal the const value" })`;
	}

	// Handle type
	const type = schema.type;

	if (Array.isArray(type)) {
		// Multiple types - create union
		const typeSchemas = type.map((t) => {
			const typeSchema: JSONSchema = { ...schema, type: t };
			delete typeSchema.enum;
			return convertSchema(typeSchema, ctx);
		});

		if (typeSchemas.length === 0) {
			return "z.never()";
		}
		if (typeSchemas.length === 1) {
			return typeSchemas[0]!;
		}
		return `z.union([${typeSchemas.join(", ")}])`;
	}

	if (!type) {
		// Infer type from keywords if possible
		// In JSON Schema, type-specific keywords only apply when value matches that type
		// Non-matching types pass validation
		const hasObjectKeywords = schema.properties !== undefined ||
			schema.additionalProperties !== undefined ||
			schema.patternProperties !== undefined ||
			schema.required !== undefined ||
			schema.propertyNames !== undefined ||
			schema.minProperties !== undefined ||
			schema.maxProperties !== undefined ||
			schema.dependentRequired !== undefined ||
			schema.dependentSchemas !== undefined ||
			schema.dependencies !== undefined;

		const hasArrayKeywords = schema.items !== undefined ||
			schema.prefixItems !== undefined ||
			schema.additionalItems !== undefined ||
			schema.minItems !== undefined ||
			schema.maxItems !== undefined ||
			schema.uniqueItems !== undefined ||
			schema.contains !== undefined;

		if (hasObjectKeywords && hasArrayKeywords) {
			// Both object and array keywords - validate based on actual type
			const objSchema = convertObject(schema, ctx);
			const arrSchema = convertArray(schema, ctx);
			return `z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = ${objSchema}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        } else if (Array.isArray(val)) {
          const result = ${arrSchema}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      })`;
		}
		if (hasObjectKeywords) {
			// Object keywords only - pass non-objects, validate objects
			const objSchema = convertObject(schema, ctx);
			return `z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = ${objSchema}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      })`;
		}
		if (hasArrayKeywords) {
			// Array keywords only - pass non-arrays, validate arrays
			const arrSchema = convertArray(schema, ctx);
			return `z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = ${arrSchema}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      })`;
		}

		// Check for numeric keywords
		const hasNumericKeywords = schema.minimum !== undefined ||
			schema.maximum !== undefined ||
			schema.exclusiveMinimum !== undefined ||
			schema.exclusiveMaximum !== undefined ||
			schema.multipleOf !== undefined;

		if (hasNumericKeywords) {
			// Numeric keywords only apply when value is a number
			const numSchema = convertNumber(schema, false);
			return `z.any().superRefine((val, ctx) => {
        if (typeof val === "number") {
          const result = ${numSchema}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      })`;
		}

		// Check for string keywords
		const hasStringKeywords = schema.minLength !== undefined ||
			schema.maxLength !== undefined ||
			schema.pattern !== undefined;

		if (hasStringKeywords) {
			// String keywords only apply when value is a string
			const strSchema = convertString(schema);
			return `z.any().superRefine((val, ctx) => {
        if (typeof val === "string") {
          const result = ${strSchema}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      })`;
		}

		// No type specified and no type-specific keywords - any
		return "z.any()";
	}

	let result: string;

	switch (type) {
		case "string":
			result = convertString(schema);
			break;
		case "number":
		case "integer":
			result = convertNumber(schema, type === "integer");
			break;
		case "boolean":
			result = "z.boolean()";
			break;
		case "null":
			result = "z.null()";
			break;
		case "object":
			result = convertObject(schema, ctx);
			break;
		case "array":
			result = convertArray(schema, ctx);
			break;
		default:
			result = "z.any()";
	}

	// Handle nullable (OpenAPI 3.0)
	if (schema.nullable === true) {
		result = `${result}.nullable()`;
	}

	return result;
}

function convertString(schema: JSONSchema): string {
	let result = "z.string()";

	// Apply format
	if (schema.format) {
		const format = schema.format;
		switch (format) {
			case "email":
				result += ".email()";
				break;
			case "uri":
			case "uri-reference":
				result += ".url()";
				break;
			case "uuid":
				result += ".uuid()";
				break;
			case "date-time":
				result += ".datetime()";
				break;
			case "date":
				result += ".date()";
				break;
			case "time":
				result += ".time()";
				break;
			case "ipv4":
				result += ".ip({ version: 'v4' })";
				break;
			case "ipv6":
				result += ".ip({ version: 'v6' })";
				break;
			case "hostname":
			case "idn-hostname":
				// No direct Zod equivalent, use regex
				result += `.regex(/^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/)`;
				break;
			// Other formats are ignored (JSON Schema treats unknown formats as annotations)
		}
	}

	// Apply constraints - use grapheme cluster counting per JSON Schema spec
	// JSON Schema counts grapheme clusters, not code units (e.g., "👨‍👩‍👦" = 1 grapheme, not 8)
	if (typeof schema.minLength === "number") {
		result += `.refine((val) => [...new Intl.Segmenter().segment(val)].length >= ${schema.minLength}, { message: "String must have at least ${schema.minLength} character(s)" })`;
	}
	if (typeof schema.maxLength === "number") {
		result += `.refine((val) => [...new Intl.Segmenter().segment(val)].length <= ${schema.maxLength}, { message: "String must have at most ${schema.maxLength} character(s)" })`;
	}
	if (schema.pattern) {
		// JSON Schema patterns are not implicitly anchored
		result += `.regex(new RegExp(${escapeString(schema.pattern)}))`;
	}

	return result;
}

function convertNumber(schema: JSONSchema, isInteger: boolean): string {
	let result = isInteger ? "z.number().int()" : "z.number()";

	// Handle minimum
	if (typeof schema.minimum === "number") {
		if (schema.exclusiveMinimum === true) {
			// Draft 4 style
			result += `.gt(${schema.minimum})`;
		} else {
			result += `.gte(${schema.minimum})`;
		}
	}

	// Handle exclusiveMinimum (Draft 6+ style)
	if (typeof schema.exclusiveMinimum === "number") {
		result += `.gt(${schema.exclusiveMinimum})`;
	}

	// Handle maximum
	if (typeof schema.maximum === "number") {
		if (schema.exclusiveMaximum === true) {
			// Draft 4 style
			result += `.lt(${schema.maximum})`;
		} else {
			result += `.lte(${schema.maximum})`;
		}
	}

	// Handle exclusiveMaximum (Draft 6+ style)
	if (typeof schema.exclusiveMaximum === "number") {
		result += `.lt(${schema.exclusiveMaximum})`;
	}

	// Handle multipleOf
	if (typeof schema.multipleOf === "number") {
		const mult = schema.multipleOf;
		// For small multipleOf values, use epsilon-based comparison to avoid float precision issues
		if (mult < 1 && mult > 0) {
			result += `.refine((val) => Math.abs(val - Math.round(val / ${mult}) * ${mult}) < 1e-10, { message: "Number must be a multiple of ${mult}" })`;
		} else {
			result += `.multipleOf(${mult})`;
		}
	}

	return result;
}

// Property names that exist on Object.prototype and need special handling
const PROTOTYPE_PROPERTY_NAMES = new Set([
	'__proto__', 'constructor', 'toString', 'valueOf', 'hasOwnProperty',
	'isPrototypeOf', 'propertyIsEnumerable', 'toLocaleString', '__defineGetter__',
	'__defineSetter__', '__lookupGetter__', '__lookupSetter__'
]);

function convertObject(schema: JSONSchema, ctx: ConversionContext): string {
	const properties = schema.properties || {};
	// Start with top-level required array
	const requiredArr = [...(schema.required || [])];

	// Draft3 style: check for required: true on individual properties
	for (const [key, propSchema] of Object.entries(properties)) {
		if (propSchema && typeof propSchema === 'object' && (propSchema as Record<string, unknown>).required === true) {
			if (!requiredArr.includes(key)) {
				requiredArr.push(key);
			}
		}
	}

	const requiredSet = new Set(requiredArr);
	// Use Object.getOwnPropertyNames to include __proto__ and other special names
	// that Object.keys might miss if the object was created with them as prototype
	const propKeys = Object.getOwnPropertyNames(properties);
	const hasPatternProperties = schema.patternProperties && Object.keys(schema.patternProperties).length > 0;

	// Check if any property names shadow Object.prototype properties
	// Must check both propKeys and requiredArr since required props get added to shape
	const hasPrototypeProps = propKeys.some(k => PROTOTYPE_PROPERTY_NAMES.has(k)) ||
		requiredArr.some(k => PROTOTYPE_PROPERTY_NAMES.has(k));

	// Build shape - include both defined properties and required properties not in properties
	const shape: string[] = [];
	const addedKeys = new Set<string>();

	// First add all defined properties
	for (const key of propKeys) {
		// Use bracket notation to safely access properties with special names like __proto__
		const propSchema = (properties as Record<string, JSONSchema>)[key];
		if (propSchema === undefined) continue;
		let propCode = convertSchema(propSchema, ctx);

		// If not required, make optional
		if (!requiredSet.has(key)) {
			propCode += ".optional()";
		}

		// Use computed property syntax [key]: to avoid __proto__ being treated as prototype
		shape.push(`[${escapeString(key)}]: ${propCode}`);
		addedKeys.add(key);
	}

	// Add required properties that aren't in properties (they must exist but can be any type)
	for (const key of requiredArr) {
		if (!addedKeys.has(key)) {
			shape.push(`[${escapeString(key)}]: z.any()`);
			addedKeys.add(key);
		}
	}

	let result: string;

	// When we have patternProperties, we need passthrough + superRefine for validation
	const needsPassthrough = hasPatternProperties || schema.propertyNames;

	// Special handling for schemas with Object.prototype property names
	// z.object() would incorrectly validate inherited properties, so we use manual validation
	if (hasPrototypeProps) {
		// Build validators for each property
		const propValidators: string[] = [];
		const handledKeys = new Set<string>();

		// First handle properties with schemas
		for (const key of propKeys) {
			const propSchema = (properties as Record<string, JSONSchema>)[key];
			if (propSchema === undefined) continue;
			const propCode = convertSchema(propSchema, ctx);
			const isRequired = requiredSet.has(key);
			handledKeys.add(key);

			propValidators.push(`
        if (Object.hasOwn(val, ${escapeString(key)})) {
          const result = ${propCode}.safeParse(val[${escapeString(key)}]);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue({ ...issue, path: [${escapeString(key)}, ...issue.path] }));
          }
        }${isRequired ? ` else {
          ctx.addIssue({ code: z.ZodIssueCode.custom, path: [${escapeString(key)}], message: "Required" });
        }` : ''}`);
		}

		// Then handle required properties that don't have schemas (they're z.any())
		for (const key of requiredArr) {
			if (handledKeys.has(key)) continue;
			handledKeys.add(key);

			propValidators.push(`
        if (!Object.hasOwn(val, ${escapeString(key)})) {
          ctx.addIssue({ code: z.ZodIssueCode.custom, path: [${escapeString(key)}], message: "Required" });
        }`);
		}

		// Handle additionalProperties
		const allDefinedKeys = [...handledKeys];
		let additionalCheck = '';
		if (schema.additionalProperties === false) {
			const definedPropsJson = JSON.stringify(allDefinedKeys);
			additionalCheck = `
        const definedProps = new Set(${definedPropsJson});
        for (const key of Object.keys(val)) {
          if (!definedProps.has(key)) {
            ctx.addIssue({ code: z.ZodIssueCode.custom, path: [key], message: "Additional property not allowed" });
          }
        }`;
		} else if (typeof schema.additionalProperties === 'object') {
			const additionalSchema = convertSchema(schema.additionalProperties, ctx);
			const definedPropsJson = JSON.stringify(allDefinedKeys);
			additionalCheck = `
        const definedProps = new Set(${definedPropsJson});
        for (const [key, value] of Object.entries(val)) {
          if (!definedProps.has(key)) {
            const result = ${additionalSchema}.safeParse(value);
            if (!result.success) {
              result.error.issues.forEach(issue => ctx.addIssue({ ...issue, path: [key, ...issue.path] }));
            }
          }
        }`;
		}

		// Use z.any() to avoid Zod transforming the object (which strips __proto__)
		// We validate the raw input directly
		result = `z.any().superRefine((val, ctx) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return;${propValidators.join('')}${additionalCheck}
    })`;
	} else if (shape.length === 0 && !needsPassthrough) {
		// No properties defined, no pattern properties
		if (schema.additionalProperties === false) {
			result = "z.object({}).strict()";
		} else if (typeof schema.additionalProperties === "object") {
			const valueSchema = convertSchema(schema.additionalProperties, ctx);
			result = `z.record(z.string(), ${valueSchema})`;
		} else {
			result = "z.object({}).passthrough()";
		}
	} else {
		result = shape.length > 0
			? `z.object({ ${shape.join(", ")} })`
			: "z.object({})";

		// When we have patternProperties, always use passthrough and validate in superRefine
		if (needsPassthrough) {
			result += ".passthrough()";
		} else if (schema.additionalProperties === false) {
			result += ".strict()";
		} else if (typeof schema.additionalProperties === "object") {
			const valueSchema = convertSchema(schema.additionalProperties, ctx);
			result += `.catchall(${valueSchema})`;
		} else {
			result += ".passthrough()";
		}
	}

	// Handle patternProperties and/or additionalProperties validation with superRefine
	// This is needed when:
	// - We have patternProperties (need to validate matching keys)
	// - We have additionalProperties: false with passthrough (need to reject extra keys)
	// - We have additionalProperties schema with passthrough (need to validate extra values)
	const needsAdditionalPropsValidation = hasPatternProperties ||
		(needsPassthrough && schema.additionalProperties === false) ||
		(needsPassthrough && typeof schema.additionalProperties === "object");

	if (needsAdditionalPropsValidation) {
		const patterns = Object.entries(schema.patternProperties || {});
		const definedProps = JSON.stringify(propKeys);

		let superRefineBody = `
      const definedProps = new Set(${definedProps});
      const patterns = [${patterns.map(([p]) => `new RegExp(${escapeString(p)})`).join(", ")}];
    `;

		// Validate pattern properties
		patterns.forEach(([pattern, patternSchema], i) => {
			const patternCode = convertSchema(patternSchema, ctx);
			superRefineBody += `
      for (const [key, value] of Object.entries(val)) {
        if (patterns[${i}].test(key)) {
          const result = ${patternCode}.safeParse(value);
          if (!result.success) {
            ctx.addIssue({
              code: z.ZodIssueCode.custom,
              path: [key],
              message: "Property matching pattern ${pattern} is invalid",
            });
          }
        }
      }`;
		});

		// If additionalProperties is false, check for unknown properties
		if (schema.additionalProperties === false) {
			superRefineBody += `
      for (const key of Object.keys(val)) {
        if (definedProps.has(key)) continue;
        const matchesPattern = patterns.some(p => p.test(key));
        if (!matchesPattern) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            path: [key],
            message: "Additional property not allowed",
          });
        }
      }`;
		} else if (typeof schema.additionalProperties === "object") {
			// Validate additional properties that don't match patterns
			const additionalSchema = convertSchema(schema.additionalProperties, ctx);
			superRefineBody += `
      for (const [key, value] of Object.entries(val)) {
        if (definedProps.has(key)) continue;
        const matchesPattern = patterns.some(p => p.test(key));
        if (!matchesPattern) {
          const result = ${additionalSchema}.safeParse(value);
          if (!result.success) {
            ctx.addIssue({
              code: z.ZodIssueCode.custom,
              path: [key],
              message: "Additional property is invalid",
            });
          }
        }
      }`;
		}

		result += `.superRefine((val, ctx) => {${superRefineBody}})`;
	}

	// Handle propertyNames (check !== undefined to handle boolean false)
	if (schema.propertyNames !== undefined) {
		const keySchema = convertSchema(schema.propertyNames, ctx);
		result += `.superRefine((val, ctx) => {
      for (const key of Object.keys(val)) {
        const result = ${keySchema}.safeParse(key);
        if (!result.success) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            path: [key],
            message: "Invalid property name",
          });
        }
      }
    })`;
	}

	// Handle required properties explicitly - z.any() accepts undefined so we need to check key presence
	// Skip if we already handled it in the hasPrototypeProps path
	if (requiredArr.length > 0 && !hasPrototypeProps) {
		const requiredJson = JSON.stringify(requiredArr);
		result += `.superRefine((val, ctx) => {
      const required = ${requiredJson};
      for (const key of required) {
        if (!Object.hasOwn(val, key)) {
          ctx.addIssue({ code: z.ZodIssueCode.custom, path: [key], message: "Required" });
        }
      }
    })`;
	}

	// Handle minProperties/maxProperties
	if (typeof schema.minProperties === "number") {
		result += `.refine((val) => Object.keys(val).length >= ${schema.minProperties}, { message: "Object must have at least ${schema.minProperties} properties" })`;
	}
	if (typeof schema.maxProperties === "number") {
		result += `.refine((val) => Object.keys(val).length <= ${schema.maxProperties}, { message: "Object must have at most ${schema.maxProperties} properties" })`;
	}

	// Handle dependentRequired - if prop X exists, props Y,Z must also exist
	if (schema.dependentRequired) {
		for (const [prop, deps] of Object.entries(schema.dependentRequired)) {
			const depsArr = deps as string[];
			if (depsArr.length > 0) {
				const message = escapeString(`Property ${prop} requires ${depsArr.join(", ")}`);
				result += `.refine((val) => {
          if (Object.hasOwn(val, ${escapeString(prop)})) {
            return ${depsArr.map(d => `Object.hasOwn(val, ${escapeString(d)})`).join(" && ")};
          }
          return true;
        }, { message: ${message} })`;
			}
		}
	}

	// Handle dependentSchemas - if prop X exists, validate against additional schema
	if (schema.dependentSchemas) {
		for (const [prop, depSchema] of Object.entries(schema.dependentSchemas)) {
			const depCode = convertSchema(depSchema as JSONSchema, ctx);
			result += `.superRefine((val, ctx) => {
        if (Object.hasOwn(val, ${escapeString(prop)})) {
          const result = ${depCode}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      })`;
		}
	}

	// Handle dependencies (draft3-7) - combines dependentRequired and dependentSchemas
	if (schema.dependencies) {
		for (const [prop, dep] of Object.entries(schema.dependencies)) {
			if (Array.isArray(dep)) {
				// Property dependency - like dependentRequired
				if (dep.length > 0) {
					const message = escapeString(`Property ${prop} requires ${dep.join(", ")}`);
					result += `.refine((val) => {
            if (Object.hasOwn(val, ${escapeString(prop)})) {
              return ${dep.map(d => `Object.hasOwn(val, ${escapeString(d)})`).join(" && ")};
            }
            return true;
          }, { message: ${message} })`;
				}
			} else {
				// Schema dependency - like dependentSchemas
				const depCode = convertSchema(dep as JSONSchema, ctx);
				result += `.superRefine((val, ctx) => {
          if (Object.hasOwn(val, ${escapeString(prop)})) {
            const result = ${depCode}.safeParse(val);
            if (!result.success) {
              result.error.issues.forEach(issue => ctx.addIssue(issue));
            }
          }
        })`;
			}
		}
	}

	return result;
}

function convertArray(schema: JSONSchema, ctx: ConversionContext): string {
	const prefixItems = schema.prefixItems;
	const items = schema.items as JSONSchema | JSONSchema[] | boolean | undefined;

	let result: string;

	// Check for tuple (prefixItems or items as array)
	if (prefixItems && Array.isArray(prefixItems)) {
		// Draft 2020-12 tuple - JSON Schema allows incomplete tuples
		// Use superRefine to validate items at each position IF present
		const tupleSchemas = prefixItems.map((item) => convertSchema(item, ctx));
		const schemasArray = `[${tupleSchemas.join(", ")}]`;

		result = `z.array(z.any()).superRefine((val, ctx) => {
      const schemas = ${schemasArray};
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = schemas[i].safeParse(val[i]);
        if (!itemResult.success) {
          itemResult.error.issues.forEach(issue => {
            ctx.addIssue({ ...issue, path: [i, ...issue.path] });
          });
        }
      }
    })`;

		// Handle items after prefixItems
		if (items === false) {
			// items: false means no items allowed beyond prefixItems length
			result += `.refine((val) => val.length <= ${prefixItems.length}, { message: "Array must not have more than ${prefixItems.length} items" })`;
		} else if (items && typeof items === "object" && !Array.isArray(items)) {
			const restSchema = convertSchema(items, ctx);
			result += `.superRefine((val, ctx) => {
        const restSchema = ${restSchema};
        for (let i = ${prefixItems.length}; i < val.length; i++) {
          const itemResult = restSchema.safeParse(val[i]);
          if (!itemResult.success) {
            itemResult.error.issues.forEach(issue => {
              ctx.addIssue({ ...issue, path: [i, ...issue.path] });
            });
          }
        }
      })`;
		} else if (schema.additionalItems === false) {
			// No items allowed beyond prefixItems length
			result += `.refine((val) => val.length <= ${prefixItems.length}, { message: "Array must not have more than ${prefixItems.length} items" })`;
		} else if (schema.additionalItems && typeof schema.additionalItems === "object") {
			const restSchema = convertSchema(schema.additionalItems, ctx);
			result += `.superRefine((val, ctx) => {
        const restSchema = ${restSchema};
        for (let i = ${prefixItems.length}; i < val.length; i++) {
          const itemResult = restSchema.safeParse(val[i]);
          if (!itemResult.success) {
            itemResult.error.issues.forEach(issue => {
              ctx.addIssue({ ...issue, path: [i, ...issue.path] });
            });
          }
        }
      })`;
		}
		// else: no items constraint - any additional items allowed
	} else if (Array.isArray(items)) {
		// Draft 7 tuple (items as array) - JSON Schema allows incomplete tuples
		const tupleSchemas = items.map((item) => convertSchema(item, ctx));
		const schemasArray = `[${tupleSchemas.join(", ")}]`;

		result = `z.array(z.any()).superRefine((val, ctx) => {
      const schemas = ${schemasArray};
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = schemas[i].safeParse(val[i]);
        if (!itemResult.success) {
          itemResult.error.issues.forEach(issue => {
            ctx.addIssue({ ...issue, path: [i, ...issue.path] });
          });
        }
      }
    })`;

		// Handle additionalItems
		if (schema.additionalItems && typeof schema.additionalItems === "object") {
			const restSchema = convertSchema(schema.additionalItems, ctx);
			result += `.superRefine((val, ctx) => {
        const restSchema = ${restSchema};
        for (let i = ${items.length}; i < val.length; i++) {
          const itemResult = restSchema.safeParse(val[i]);
          if (!itemResult.success) {
            itemResult.error.issues.forEach(issue => {
              ctx.addIssue({ ...issue, path: [i, ...issue.path] });
            });
          }
        }
      })`;
		} else if (schema.additionalItems === false) {
			// No items allowed beyond tuple length
			result += `.refine((val) => val.length <= ${items.length}, { message: "Array must not have more than ${items.length} items" })`;
		}
		// else: additionalItems defaults to true (any items allowed)
	} else if (items && typeof items === "object") {
		// Regular array with items schema
		const itemSchema = convertSchema(items, ctx);
		result = `z.array(${itemSchema})`;
	} else if (items === false) {
		// No items allowed
		result = "z.tuple([])";
	} else {
		// No items constraint
		result = "z.array(z.any())";
	}

	// Apply constraints
	if (typeof schema.minItems === "number") {
		result += `.min(${schema.minItems})`;
	}
	if (typeof schema.maxItems === "number") {
		result += `.max(${schema.maxItems})`;
	}

	// Handle uniqueItems
	if (schema.uniqueItems === true) {
		result += `.refine((arr) => {
      const seen = new Set();
      for (const item of arr) {
        const key = JSON.stringify(item);
        if (seen.has(key)) return false;
        seen.add(key);
      }
      return true;
    }, { message: "Array items must be unique" })`;
	}

	// Handle contains (check !== undefined to handle boolean false)
	if (schema.contains !== undefined) {
		const containsSchema = convertSchema(schema.contains, ctx);
		const minContains = schema.minContains ?? 1;
		const maxContains = schema.maxContains;

		if (maxContains !== undefined) {
			result += `.refine((arr) => {
        let count = 0;
        for (const item of arr) {
          if (${containsSchema}.safeParse(item).success) count++;
        }
        return count >= ${minContains} && count <= ${maxContains};
      }, { message: "Array must contain between ${minContains} and ${maxContains} items matching schema" })`;
		} else {
			result += `.refine((arr) => {
        let count = 0;
        for (const item of arr) {
          if (${containsSchema}.safeParse(item).success) count++;
        }
        return count >= ${minContains};
      }, { message: "Array must contain at least ${minContains} item(s) matching schema" })`;
		}
	}

	return result;
}

function convertEnum(values: unknown[], schema: JSONSchema, ctx: ConversionContext): string {
	// Handle nullable enum in OpenAPI 3.0
	if (ctx.version === "openapi-3.0" && schema.nullable === true && values.length === 1 && values[0] === null) {
		return "z.null()";
	}

	if (values.length === 0) {
		return "z.never()";
	}

	if (values.length === 1) {
		const singleValue = values[0];

		// For primitives, use z.literal
		if (isPrimitive(singleValue)) {
			return `z.literal(${JSON.stringify(singleValue)})`;
		}

		// For objects/arrays, use refine with deep equality
		const isArray = Array.isArray(singleValue);
		const baseType = isArray ? 'z.array(z.any())' : 'z.object({}).passthrough()';
		const sortedJson = JSON.stringify(singleValue, Object.keys(singleValue as object).sort());

		return `${baseType}.refine((val) => JSON.stringify(val, Object.keys(val as object).sort()) === ${JSON.stringify(sortedJson)}, { message: "Value must equal the const value" })`;
	}

	// Check if all values are strings
	if (values.every((v) => typeof v === "string")) {
		return `z.enum([${values.map((v) => JSON.stringify(v)).join(", ")}])`;
	}

	// Check if there are any complex values (objects/arrays)
	const hasComplexValues = values.some((v) => !isPrimitive(v));

	if (hasComplexValues) {
		// Use refine with deep equality check against all enum values
		// Handle null safely - null doesn't have Object.keys()
		const sortedValues = values.map(v =>
			JSON.stringify(v, v != null && typeof v === 'object' ? Object.keys(v).sort() : undefined)
		);
		const valuesArrayCode = `[${sortedValues.map(v => JSON.stringify(v)).join(", ")}]`;

		return `z.any().refine((val) => {
      const normalized = JSON.stringify(val, val != null && typeof val === 'object' ? Object.keys(val).sort() : undefined);
      const validValues = ${valuesArrayCode};
      return validValues.includes(normalized);
    }, { message: "Value must be one of the enum values" })`;
	}

	// All primitives (mixed types) - use union of literals
	const literals = values.map((v) => `z.literal(${JSON.stringify(v)})`);
	return `z.union([${literals.join(", ")}])`;
}

function convertAnyOf(schemas: JSONSchema[], ctx: ConversionContext): string {
	if (schemas.length === 0) {
		return "z.never()";
	}

	if (schemas.length === 1) {
		return convertSchema(schemas[0]!, ctx);
	}

	const converted = schemas.map((s) => convertSchema(s, ctx));
	return `z.union([${converted.join(", ")}])`;
}

function convertOneOf(schemas: JSONSchema[], ctx: ConversionContext): string {
	if (schemas.length === 0) {
		return "z.never()";
	}

	if (schemas.length === 1) {
		return convertSchema(schemas[0]!, ctx);
	}

	// oneOf means exactly one must match
	// Zod's discriminatedUnion is ideal when there's a discriminator
	// For general case, use union with refine to ensure exactly one matches
	const converted = schemas.map((s) => convertSchema(s, ctx));

	// Use union but add refinement to ensure exactly one matches
	return `z.any().superRefine((val, ctx) => {
    const schemas = [${converted.join(", ")}];
    const results = schemas.map(s => s.safeParse(val));
    const validCount = results.filter(r => r.success).length;
    if (validCount === 0) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: "Value must match exactly one schema in oneOf, but matched none",
      });
    } else if (validCount > 1) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: "Value must match exactly one schema in oneOf, but matched multiple",
      });
    }
  })`;
}

function convertAllOf(schemas: JSONSchema[], ctx: ConversionContext): string {
	if (schemas.length === 0) {
		return "z.any()";
	}

	if (schemas.length === 1) {
		return convertSchema(schemas[0]!, ctx);
	}

	// allOf - all schemas must match
	// Use intersection for pairs, chain for more
	const converted = schemas.map((s) => convertSchema(s, ctx));

	if (converted.length === 2) {
		return `z.intersection(${converted[0]}, ${converted[1]})`;
	}

	// Chain intersections: ((a & b) & c) & d
	let result = `z.intersection(${converted[0]}, ${converted[1]})`;
	for (let i = 2; i < converted.length; i++) {
		result = `z.intersection(${result}, ${converted[i]})`;
	}

	return result;
}

/**
 * Convert a JSON Schema to Zod code string
 */
export function jsonSchemaToZodCode(schema: JSONSchema): string {
	const version = detectVersion(schema);

	// Collect $defs/definitions
	const defs: Record<string, JSONSchema> = {
		...(schema.$defs || {}),
		...(schema.definitions || {}),
	};

	const ctx: ConversionContext = {
		version,
		defs,
		refs: new Map(),
		processing: new Set(),
		rootSchema: schema,
	};

	return convertSchema(schema, ctx);
}
