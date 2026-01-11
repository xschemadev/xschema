/**
 * Effect/Schema Renderer
 * Converts SchemaNode IR to Effect Schema code strings
 */

import type {
	SchemaNode,
	StringNode,
	NumberNode,
	ObjectNode,
	ArrayNode,
	TupleNode,
	UnionNode,
	IntersectionNode,
	OneOfNode,
	NotNode,
	LiteralNode,
	EnumNode,
	RefNode,
	ConditionalNode,
	TypeGuardedNode,
	NullableNode,
	PropertyDef,
} from "@xschemadev/core";
import { escapeString, isPrimitive, sortedStringify } from "@xschemadev/core";

// JS reserved words that require computed property syntax
const RESERVED_WORDS = new Set([
	"break", "case", "catch", "class", "const", "continue", "debugger", "default",
	"delete", "do", "else", "enum", "export", "extends", "false", "finally", "for",
	"function", "if", "implements", "import", "in", "instanceof", "interface", "let",
	"new", "null", "package", "private", "protected", "public", "return", "static",
	"super", "switch", "this", "throw", "true", "try", "typeof", "var", "void",
	"while", "with", "yield"
]);

// Valid JS identifier pattern
const VALID_IDENTIFIER = /^[a-zA-Z_$][a-zA-Z0-9_$]*$/;

function formatPropertyKey(key: string): string {
	if (VALID_IDENTIFIER.test(key) && !RESERVED_WORDS.has(key)) {
		return key;
	}
	return `[${escapeString(key)}]`;
}

/**
 * Render a SchemaNode to Effect Schema code
 */
export function render(node: SchemaNode): string {
	switch (node.kind) {
		case "string":
			return renderString(node);
		case "number":
			return renderNumber(node);
		case "boolean":
			return "S.Boolean";
		case "null":
			return "S.Null";
		case "object":
			return renderObject(node);
		case "array":
			return renderArray(node);
		case "tuple":
			return renderTuple(node);
		case "union":
			return renderUnion(node);
		case "intersection":
			return renderIntersection(node);
		case "oneOf":
			return renderOneOf(node);
		case "not":
			return renderNot(node);
		case "literal":
			return renderLiteral(node);
		case "enum":
			return renderEnum(node);
		case "ref":
			return renderRef(node);
		case "conditional":
			return renderConditional(node);
		case "typeGuarded":
			return renderTypeGuarded(node);
		case "nullable":
			return renderNullable(node);
		case "any":
			return "S.Unknown";
		case "never":
			return "S.Never";
		default: {
			const _exhaustive: never = node;
			throw new Error(`Unknown node kind: ${(node as SchemaNode).kind}`);
		}
	}
}

function renderString(node: StringNode): string {
	let result = "S.String";

	// Format validations
	if (node.format) {
		switch (node.format) {
			case "email":
				result = "S.String.pipe(S.email())";
				break;
			case "uuid":
				result = "S.String.pipe(S.uuid())";
				break;
			case "date-time":
				result = "S.String.pipe(S.isoDateTime())";
				break;
			case "date":
				result = "S.String.pipe(S.date())";
				break;
			case "time":
				result = "S.String.pipe(S.time())";
				break;
			case "ipv4":
				result = "S.String.pipe(S.ipv4())";
				break;
			case "ipv6":
				result = "S.String.pipe(S.ipv6())";
				break;
			case "hostname":
			case "idn-hostname":
				result = `S.String.pipe(S.pattern(/^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/))`;
				break;
			default:
				// Unknown format - ignore per JSON Schema spec
				result = "S.String";
		}
	}

	// Constraints - use grapheme cluster counting per JSON Schema spec
	const { minLength, maxLength, pattern } = node.constraints;
	const hasConstraints = minLength !== undefined || maxLength !== undefined || pattern !== undefined;

	if (hasConstraints) {
		const filters: string[] = [];

		if (minLength !== undefined) {
			filters.push(`S.filter((val) => [...new Intl.Segmenter().segment(val)].length >= ${minLength}, { message: () => "String must have at least ${minLength} character(s)" })`);
		}
		if (maxLength !== undefined) {
			filters.push(`S.filter((val) => [...new Intl.Segmenter().segment(val)].length <= ${maxLength}, { message: () => "String must have at most ${maxLength} character(s)" })`);
		}
		if (pattern) {
			filters.push(`S.pattern(new RegExp(${escapeString(pattern)}))`);
		}

		// If we already have format pipe, extend it
		if (node.format && node.format !== "hostname" && node.format !== "idn-hostname") {
			// Remove the closing paren and add filters
			result = result.slice(0, -1) + ", " + filters.join(", ") + ")";
		} else {
			// Create new pipe
			result = `S.String.pipe(${filters.join(", ")})`;
		}
	}

	return result;
}

function renderNumber(node: NumberNode): string {
	const filters: string[] = [];
	const {
		minimum,
		maximum,
		exclusiveMinimum,
		exclusiveMaximum,
		multipleOf,
	} = node.constraints;

	// Integer check
	if (node.integer) {
		filters.push("S.int()");
	}

	// Range constraints
	if (minimum !== undefined) {
		filters.push(`S.greaterThanOrEqualTo(${minimum})`);
	}
	if (exclusiveMinimum !== undefined) {
		filters.push(`S.greaterThan(${exclusiveMinimum})`);
	}
	if (maximum !== undefined) {
		filters.push(`S.lessThanOrEqualTo(${maximum})`);
	}
	if (exclusiveMaximum !== undefined) {
		filters.push(`S.lessThan(${exclusiveMaximum})`);
	}

	// Multiple of
	if (multipleOf !== undefined) {
		// For small values, use epsilon-based comparison for float precision
		if (multipleOf < 1 && multipleOf > 0) {
			filters.push(`S.filter((val) => Math.abs(val - Math.round(val / ${multipleOf}) * ${multipleOf}) < 1e-10, { message: () => "Number must be a multiple of ${multipleOf}" })`);
		} else {
			filters.push(`S.multipleOf(${multipleOf})`);
		}
	}

	if (filters.length === 0) {
		return "S.Number";
	}

	return `S.Number.pipe(${filters.join(", ")})`;
}

// Check for JS prototype property names that cause runtime issues
function hasProtoProps(node: ObjectNode): boolean {
	const protoNames = new Set(["__proto__", "constructor", "toString"]);
	for (const key of node.properties.keys()) {
		if (protoNames.has(key)) return true;
	}
	for (const dep of node.dependencies.keys()) {
		if (protoNames.has(dep)) return true;
	}
	return false;
}

function renderObject(node: ObjectNode): string {
	const propKeys = Array.from(node.properties.keys());
	const requiredKeys = propKeys.filter((k) => node.properties.get(k)!.required);

	// For objects with JS prototype property names, use full custom validation
	if (hasProtoProps(node)) {
		return renderObjectWithProtoProps(node);
	}

	const hasPatternProps = node.patternProperties.length > 0;
	const needsCustomValidation = hasPatternProps || node.propertyNames !== undefined ||
		node.additionalProperties === false ||
		(typeof node.additionalProperties === "object" && node.additionalProperties.kind !== "any");

	// If using record-style (no properties, just additionalProperties schema)
	if (
		propKeys.length === 0 &&
		!hasPatternProps &&
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const valueSchema = render(node.additionalProperties);
		let result = `S.Unknown.pipe(S.filter((val) => typeof val === "object" && val !== null && !Array.isArray(val), { message: () => "Expected object" }), S.filter((val) => {
      for (const [key, value] of Object.entries(val)) {
        const result = S.decodeUnknownEither(${valueSchema})(value);
        if (result._tag === "Left") return false;
      }
      return true;
    }, { message: () => "Invalid property value" }))`;
		const filters = renderObjectConstraints(node);
		if (filters.length > 0) {
			result = `${result}.pipe(${filters.join(", ")})`;
		}
		return result;
	}

	// For objects that need custom validation (additionalProperties, patternProperties, propertyNames)
	// OR simple objects - always use custom validation to properly reject arrays
	return renderObjectWithPatternProps(node);
}

function renderObjectWithProtoProps(node: ObjectNode): string {
	const propKeys = Array.from(node.properties.keys());
	const validators: string[] = [];

	// Property validators
	for (const key of propKeys) {
		const prop = node.properties.get(key)!;
		const propCode = render(prop.schema as SchemaNode);
		const keyExpr = escapeString(key);

		if (prop.required) {
			validators.push(`
        if (Object.hasOwn(val, ${keyExpr})) {
          const result = S.decodeUnknownEither(${propCode})(val[${keyExpr}]);
          if (result._tag === "Left") return false;
        } else {
          return false;
        }`);
		} else {
			validators.push(`
        if (Object.hasOwn(val, ${keyExpr})) {
          const result = S.decodeUnknownEither(${propCode})(val[${keyExpr}]);
          if (result._tag === "Left") return false;
        }`);
		}
	}

	// Additional properties check
	let additionalCheck = "";
	if (node.additionalProperties === false) {
		const definedPropsJson = JSON.stringify(propKeys);
		additionalCheck = `
        const definedProps = new Set(${definedPropsJson});
        for (const key of Object.keys(val)) {
          if (!definedProps.has(key)) return false;
        }`;
	} else if (
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const additionalSchema = render(node.additionalProperties);
		const definedPropsJson = JSON.stringify(propKeys);
		additionalCheck = `
        const definedProps = new Set(${definedPropsJson});
        for (const [key, value] of Object.entries(val)) {
          if (!definedProps.has(key)) {
            const result = S.decodeUnknownEither(${additionalSchema})(value);
            if (result._tag === "Left") return false;
          }
        }`;
	}

	// Dependencies check
	let dependenciesCheck = "";
	for (const [prop, dep] of node.dependencies) {
		if (dep.kind === "property") {
			if (dep.requiredProperties.length > 0) {
				dependenciesCheck += `
        if (Object.hasOwn(val, ${escapeString(prop)})) {
          if (!(${dep.requiredProperties.map((d) => `Object.hasOwn(val, ${escapeString(d)})`).join(" && ")})) return false;
        }`;
			}
		} else {
			const depCode = render(dep.schema as SchemaNode);
			dependenciesCheck += `
        if (Object.hasOwn(val, ${escapeString(prop)})) {
          const result = S.decodeUnknownEither(${depCode})(val);
          if (result._tag === "Left") return false;
        }`;
		}
	}

	return `S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;${validators.join("")}${additionalCheck}${dependenciesCheck}
      return true;
    }, { message: () => "Object validation failed" }))`;
}

function renderObjectWithPatternProps(node: ObjectNode): string {
	const propKeys = Array.from(node.properties.keys());
	const requiredKeys = propKeys.filter((k) => node.properties.get(k)!.required);
	const validators: string[] = [];

	// Property validators
	for (const key of propKeys) {
		const prop = node.properties.get(key)!;
		const propCode = render(prop.schema as SchemaNode);
		const keyExpr = escapeString(key);

		if (prop.required && (prop.schema as SchemaNode).kind === "any") {
			// Required property with any schema - need explicit hasOwn check
			validators.push(`
        if (!Object.hasOwn(val, ${keyExpr})) return false;`);
		} else if (prop.required) {
			validators.push(`
        if (Object.hasOwn(val, ${keyExpr})) {
          const result = S.decodeUnknownEither(${propCode})(val[${keyExpr}]);
          if (result._tag === "Left") return false;
        } else {
          return false;
        }`);
		} else {
			validators.push(`
        if (Object.hasOwn(val, ${keyExpr})) {
          const result = S.decodeUnknownEither(${propCode})(val[${keyExpr}]);
          if (result._tag === "Left") return false;
        }`);
		}
	}

	// Pattern properties validation
	const patterns = node.patternProperties;
	const patternValidators: string[] = [];
	if (patterns.length > 0) {
		const definedProps = JSON.stringify(propKeys);
		for (const p of patterns) {
			const patternCode = render(p.schema as SchemaNode);
			const patternStr = escapeString(p.pattern);
			patternValidators.push(`
        for (const [key, value] of Object.entries(val)) {
          if (new RegExp(${patternStr}).test(key)) {
            const result = S.decodeUnknownEither(${patternCode})(value);
            if (result._tag === "Left") return false;
          }
        }`);
		}
	}

	// Additional properties handling
	let additionalCheck = "";
	if (node.additionalProperties === false) {
		const definedPropsJson = JSON.stringify(propKeys);
		const patternChecks = patterns.map(p => `new RegExp(${escapeString(p.pattern)})`).join(", ");
		if (patterns.length > 0) {
			additionalCheck = `
        const definedProps = new Set(${definedPropsJson});
        const patterns = [${patternChecks}];
        for (const key of Object.keys(val)) {
          if (definedProps.has(key)) continue;
          const matchesPattern = patterns.some(p => p.test(key));
          if (!matchesPattern) return false;
        }`;
		} else {
			additionalCheck = `
        const definedProps = new Set(${definedPropsJson});
        for (const key of Object.keys(val)) {
          if (!definedProps.has(key)) return false;
        }`;
		}
	} else if (
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const additionalSchema = render(node.additionalProperties);
		const definedPropsJson = JSON.stringify(propKeys);
		const patternChecks = patterns.map(p => `new RegExp(${escapeString(p.pattern)})`).join(", ");
		if (patterns.length > 0) {
			additionalCheck = `
        const definedProps = new Set(${definedPropsJson});
        const patterns = [${patternChecks}];
        for (const [key, value] of Object.entries(val)) {
          if (definedProps.has(key)) continue;
          const matchesPattern = patterns.some(p => p.test(key));
          if (!matchesPattern) {
            const result = S.decodeUnknownEither(${additionalSchema})(value);
            if (result._tag === "Left") return false;
          }
        }`;
		} else {
			additionalCheck = `
        const definedProps = new Set(${definedPropsJson});
        for (const [key, value] of Object.entries(val)) {
          if (!definedProps.has(key)) {
            const result = S.decodeUnknownEither(${additionalSchema})(value);
            if (result._tag === "Left") return false;
          }
        }`;
		}
	}

	// Property names validation
	let propertyNamesCheck = "";
	if (node.propertyNames) {
		const keySchema = render(node.propertyNames);
		propertyNamesCheck = `
        for (const key of Object.keys(val)) {
          const result = S.decodeUnknownEither(${keySchema})(key);
          if (result._tag === "Left") return false;
        }`;
	}

	// Dependencies check
	let dependenciesCheck = "";
	for (const [prop, dep] of node.dependencies) {
		if (dep.kind === "property") {
			if (dep.requiredProperties.length > 0) {
				dependenciesCheck += `
        if (Object.hasOwn(val, ${escapeString(prop)})) {
          if (!(${dep.requiredProperties.map((d) => `Object.hasOwn(val, ${escapeString(d)})`).join(" && ")})) return false;
        }`;
			}
		} else {
			const depCode = render(dep.schema as SchemaNode);
			dependenciesCheck += `
        if (Object.hasOwn(val, ${escapeString(prop)})) {
          const result = S.decodeUnknownEither(${depCode})(val);
          if (result._tag === "Left") return false;
        }`;
		}
	}

	// Min/max properties
	let sizeChecks = "";
	if (node.minProperties !== undefined) {
		sizeChecks += `
        if (Object.keys(val).length < ${node.minProperties}) return false;`;
	}
	if (node.maxProperties !== undefined) {
		sizeChecks += `
        if (Object.keys(val).length > ${node.maxProperties}) return false;`;
	}

	return `S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;${validators.join("")}${patternValidators.join("")}${additionalCheck}${propertyNamesCheck}${dependenciesCheck}${sizeChecks}
      return true;
    }, { message: () => "Object validation failed" }))`;
}

function renderObjectConstraints(node: ObjectNode): string[] {
	const filters: string[] = [];

	if (node.minProperties !== undefined) {
		filters.push(`S.filter((val) => Object.keys(val).length >= ${node.minProperties}, { message: () => "Object must have at least ${node.minProperties} properties" })`);
	}
	if (node.maxProperties !== undefined) {
		filters.push(`S.filter((val) => Object.keys(val).length <= ${node.maxProperties}, { message: () => "Object must have at most ${node.maxProperties} properties" })`);
	}

	return filters;
}

function renderDependenciesFilters(node: ObjectNode): string[] {
	const filters: string[] = [];

	for (const [prop, dep] of node.dependencies) {
		if (dep.kind === "property") {
			if (dep.requiredProperties.length > 0) {
				const message = escapeString(
					`Property ${prop} requires ${dep.requiredProperties.join(", ")}`,
				);
				filters.push(`S.filter((val) => {
          if (Object.hasOwn(val, ${escapeString(prop)})) {
            return ${dep.requiredProperties.map((d) => `Object.hasOwn(val, ${escapeString(d)})`).join(" && ")};
          }
          return true;
        }, { message: () => ${message} })`);
			}
		} else {
			const depCode = render(dep.schema as SchemaNode);
			filters.push(`S.filter((val) => {
        if (Object.hasOwn(val, ${escapeString(prop)})) {
          const result = S.decodeUnknownEither(${depCode})(val);
          return result._tag === "Right";
        }
        return true;
      }, { message: () => "Schema dependency validation failed" })`);
		}
	}

	return filters;
}

function renderArray(node: ArrayNode): string {
	const itemSchema = render(node.items);
	const filters = renderArrayConstraints(node.constraints);
	
	if (filters.length > 0) {
		return `S.Array(${itemSchema}).pipe(${filters.join(", ")})`;
	}
	return `S.Array(${itemSchema})`;
}

function renderArrayConstraints(constraints: ArrayNode["constraints"]): string[] {
	const filters: string[] = [];

	if (constraints.minItems !== undefined) {
		filters.push(`S.minItems(${constraints.minItems})`);
	}
	if (constraints.maxItems !== undefined) {
		filters.push(`S.maxItems(${constraints.maxItems})`);
	}

	if (constraints.uniqueItems) {
		filters.push(`S.filter((arr) => {
      const seen = new Set();
      for (const item of arr) {
        const key = JSON.stringify(item);
        if (seen.has(key)) return false;
        seen.add(key);
      }
      return true;
    }, { message: () => "Array items must be unique" })`);
	}

	if (constraints.contains) {
		const containsSchema = render(constraints.contains.schema as SchemaNode);
		const minContains = constraints.contains.minContains;
		const maxContains = constraints.contains.maxContains;

		if (maxContains !== undefined) {
			filters.push(`S.filter((arr) => {
        let count = 0;
        for (const item of arr) {
          const result = S.decodeUnknownEither(${containsSchema})(item);
          if (result._tag === "Right") count++;
        }
        return count >= ${minContains} && count <= ${maxContains};
      }, { message: () => "Array must contain between ${minContains} and ${maxContains} items matching schema" })`);
		} else {
			filters.push(`S.filter((arr) => {
        let count = 0;
        for (const item of arr) {
          const result = S.decodeUnknownEither(${containsSchema})(item);
          if (result._tag === "Right") count++;
        }
        return count >= ${minContains};
      }, { message: () => "Array must contain at least ${minContains} item(s) matching schema" })`);
		}
	}

	return filters;
}

function renderTuple(node: TupleNode): string {
	const tupleSchemas = node.prefixItems.map((item) => render(item));
	const schemasArray = `[${tupleSchemas.join(", ")}]`;

	// Use custom validation to allow arrays shorter than prefixItems length
	// JSON Schema allows [1] to be valid for prefixItems: [number, string] with items: false
	let result = `S.Array(S.Unknown).pipe(S.filter((val) => {
      const schemas = ${schemasArray};
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const result = S.decodeUnknownEither(schemas[i])(val[i]);
        if (result._tag === "Left") return false;
      }
      return true;
    }, { message: () => "Tuple item validation failed" }))`;

	// Rest items handling
	if (node.restItems === false) {
		// No additional items beyond prefixItems length
		result += `.pipe(S.filter((val) => val.length <= ${node.prefixItems.length}, { message: () => "Array must not have more than ${node.prefixItems.length} items" }))`;
	} else if (node.restItems.kind !== "any") {
		// Validate rest items against schema
		const restSchema = render(node.restItems);
		result += `.pipe(S.filter((val) => {
        const restSchema = ${restSchema};
        for (let i = ${node.prefixItems.length}; i < val.length; i++) {
          const result = S.decodeUnknownEither(restSchema)(val[i]);
          if (result._tag === "Left") return false;
        }
        return true;
      }, { message: () => "Rest items validation failed" }))`;
	}
	// If restItems is any, no additional validation needed

	const filters = renderArrayConstraints(node.constraints);
	if (filters.length > 0) {
		result = `${result}.pipe(${filters.join(", ")})`;
	}
	return result;
}

function renderUnion(node: UnionNode): string {
	if (node.variants.length === 0) return "S.Never";

	// Filter out "never" variants since they can't match anything
	const filtered = node.variants.filter((v) => v.kind !== "never");
	if (filtered.length === 0) return "S.Never";
	if (filtered.length === 1) return render(filtered[0]!);

	const schemas = filtered.map((v) => render(v));
	return `S.Union(${schemas.join(", ")})`;
}

function renderIntersection(node: IntersectionNode): string {
	if (node.schemas.length === 0) return "S.Unknown";

	// Short-circuit: if ANY schema is never, intersection is never
	if (node.schemas.some((s) => s.kind === "never")) {
		return "S.Never";
	}

	// Filter out "any" schemas since they don't constrain the intersection
	const filtered = node.schemas.filter((s) => s.kind !== "any");
	if (filtered.length === 0) return "S.Unknown";
	if (filtered.length === 1) return render(filtered[0]!);

	// Effect/Schema doesn't have a generic intersection type like valibot
	// Use S.extend for objects, or filter validation for general case
	const allObjects = filtered.every(
		(s) => s.kind === "object" || (s.kind === "ref" && s.resolved.kind === "object")
	);

	if (allObjects) {
		// Use S.extend for object intersection (merges object types)
		const schemas = filtered.map((s) => render(s));
		return schemas.reduce((acc, schema) => `S.extend(${acc}, ${schema})`);
	}

	// For non-object intersections, validate all schemas
	const schemas = filtered.map((s) => render(s));
	return `S.Unknown.pipe(S.filter((val) => {
      const results = [${schemas.map(s => `S.decodeUnknownEither(${s})(val)`).join(", ")}];
      return results.every(r => r._tag === "Right");
    }, { message: () => "Value must match all schemas in allOf" }))`;
}

function renderOneOf(node: OneOfNode): string {
	if (node.schemas.length === 0) return "S.Never";
	if (node.schemas.length === 1) return render(node.schemas[0]!);

	// Filter out "never" schemas since they can never match
	const filtered = node.schemas.filter((s) => s.kind !== "never");

	// If all schemas are never, nothing can match exactly one
	if (filtered.length === 0) return "S.Never";

	// If exactly one schema remains after filtering never, it must match
	if (filtered.length === 1) return render(filtered[0]!);

	// Count how many "any" schemas there are - if > 1, always matches multiple
	const anyCount = filtered.filter((s) => s.kind === "any").length;
	if (anyCount > 1) {
		// Multiple "any" schemas means any value matches multiple
		return `S.Unknown.pipe(S.filter(() => false, { message: () => "oneOf has multiple 'true' schemas - impossible to match exactly one" }))`;
	}

	const schemas = filtered.map((s) => render(s));
	return `S.Unknown.pipe(S.filter((val) => {
    const schemas = [${schemas.join(", ")}];
    const results = schemas.map(s => S.decodeUnknownEither(s)(val));
    const validCount = results.filter(r => r._tag === "Right").length;
    return validCount === 1;
  }, { message: () => "Value must match exactly one schema in oneOf" }))`;
}

function renderNot(node: NotNode): string {
	const schema = render(node.schema);
	return `S.Unknown.pipe(S.filter((val) => {
    const result = S.decodeUnknownEither(${schema})(val);
    return result._tag === "Left";
  }, { message: () => "Value must not match schema" }))`;
}

function renderLiteral(node: LiteralNode): string {
	// Primitive values use S.Literal
	if (isPrimitive(node.value)) {
		return `S.Literal(${JSON.stringify(node.value)})`;
	}

	// Objects/arrays need deep equality check
	const isArray = Array.isArray(node.value);
	const baseType = isArray ? "S.Array(S.Unknown)" : "S.Struct({})";
	const sorted = sortedStringify(node.value);

	return `${baseType}.pipe(S.filter((val) => JSON.stringify(val, Object.keys(val as object).sort()) === ${JSON.stringify(sorted)}, { message: () => "Value must equal the const value" }))`;
}

function renderEnum(node: EnumNode): string {
	const values = node.values;

	if (values.length === 0) return "S.Never";

	if (values.length === 1) {
		return renderLiteral({ kind: "literal", value: values[0] });
	}

	// All strings - use S.Literal union
	if (values.every((v) => typeof v === "string")) {
		const literals = values.map((v) => `S.Literal(${JSON.stringify(v)})`);
		return `S.Union(${literals.join(", ")})`;
	}

	// All same-type primitives - use S.Literal union
	if (values.every((v) => isPrimitive(v))) {
		const allSameType = values.every((v) => typeof v === typeof values[0]);
		if (allSameType) {
			const literals = values.map((v) => `S.Literal(${JSON.stringify(v)})`);
			return `S.Union(${literals.join(", ")})`;
		}
	}

	// Heterogeneous enum (mixed types or complex values) - stringify for comparison
	// Single-stringify each value (with sorted keys for objects), then JSON.stringify to embed as string literal
	const sortedValues = values.map((v) => {
		const stringified = JSON.stringify(
			v,
			v != null && typeof v === "object"
				? Object.keys(v as object).sort()
				: undefined,
		);
		// Wrap in JSON.stringify to produce a valid JS string literal for code gen
		return JSON.stringify(stringified);
	});
	return `S.Unknown.pipe(S.filter((val) => {
      const sorted = JSON.stringify(val, val != null && typeof val === "object" ? Object.keys(val as object).sort() : undefined);
      return [${sortedValues.join(", ")}].includes(sorted);
    }, { message: () => "Value must match one of the enum values" }))`;
}

function renderRef(node: RefNode): string {
	return render(node.resolved);
}

function renderConditional(node: ConditionalNode): string {
	const ifSchema = render(node.if);
	const thenSchema = node.then ? render(node.then) : null;
	const elseSchema = node.else ? render(node.else) : null;

	if (thenSchema && elseSchema) {
		// Both then and else present: validate with then if condition matches, else otherwise
		return `S.Unknown.pipe(S.filter((val) => {
      const ifResult = S.decodeUnknownEither(${ifSchema})(val);
      if (ifResult._tag === "Right") {
        const thenResult = S.decodeUnknownEither(${thenSchema})(val);
        return thenResult._tag === "Right";
      } else {
        const elseResult = S.decodeUnknownEither(${elseSchema})(val);
        return elseResult._tag === "Right";
      }
    }, { message: () => "Conditional validation failed" }))`;
	} else if (thenSchema) {
		// Only then present: validate with then if condition matches
		return `S.Unknown.pipe(S.filter((val) => {
      const ifResult = S.decodeUnknownEither(${ifSchema})(val);
      if (ifResult._tag === "Right") {
        const thenResult = S.decodeUnknownEither(${thenSchema})(val);
        return thenResult._tag === "Right";
      }
      return true;
    }, { message: () => "Conditional 'then' validation failed" }))`;
	} else if (elseSchema) {
		// Only else present: validate with else if condition doesn't match
		return `S.Unknown.pipe(S.filter((val) => {
      const ifResult = S.decodeUnknownEither(${ifSchema})(val);
      if (ifResult._tag === "Left") {
        const elseResult = S.decodeUnknownEither(${elseSchema})(val);
        return elseResult._tag === "Right";
      }
      return true;
    }, { message: () => "Conditional 'else' validation failed" }))`;
	}

	// if without then/else has no effect
	return "S.Unknown";
}

function renderTypeGuarded(node: TypeGuardedNode): string {
	if (node.guards.length === 0) return "S.Unknown";

	const checks: string[] = [];

	for (const guard of node.guards) {
		const schema = render(guard.schema);
		switch (guard.check) {
			case "object":
				checks.push(`if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = S.decodeUnknownEither(${schema})(val);
          if (result._tag === "Left") return false;
        }`);
				break;
			case "array":
				checks.push(`if (Array.isArray(val)) {
          const result = S.decodeUnknownEither(${schema})(val);
          if (result._tag === "Left") return false;
        }`);
				break;
			case "string":
				checks.push(`if (typeof val === "string") {
          const result = S.decodeUnknownEither(${schema})(val);
          if (result._tag === "Left") return false;
        }`);
				break;
			case "number":
				checks.push(`if (typeof val === "number") {
          const result = S.decodeUnknownEither(${schema})(val);
          if (result._tag === "Left") return false;
        }`);
				break;
		}
	}

	return `S.Unknown.pipe(S.filter((val) => {
        ${checks.join("\n        ")}
        return true;
      }, { message: () => "Type-guarded validation failed" }))`;
}

function renderNullable(node: NullableNode): string {
	const innerSchema = render(node.inner);
	return `S.NullOr(${innerSchema})`;
}
