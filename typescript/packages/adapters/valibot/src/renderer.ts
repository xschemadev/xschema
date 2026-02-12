/**
 * Valibot Renderer
 * Converts SchemaNode IR to Valibot code strings
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
	ConditionalNode,
	TypeGuardedNode,
	NullableNode,
	PropertyDef,
} from "@xschemadev/core";
import {
	DEEP_SORTED_STRINGIFY_RUNTIME,
	escapeString,
	hasPrototypeProperties,
	isPrimitive,
	sortedStringify,
} from "@xschemadev/core";

// JS identifier regex - keys matching this can use direct property syntax
const IDENTIFIER_RE = /^[a-zA-Z_$][a-zA-Z0-9_$]*$/;

// Reserved words that can't be used as unquoted property names
const RESERVED_WORDS = new Set([
	"break", "case", "catch", "continue", "debugger", "default", "delete",
	"do", "else", "finally", "for", "function", "if", "in", "instanceof",
	"new", "return", "switch", "this", "throw", "try", "typeof", "var",
	"void", "while", "with", "class", "const", "enum", "export", "extends",
	"import", "super", "implements", "interface", "let", "package", "private",
	"protected", "public", "static", "yield"
]);

function canUseDirectSyntax(key: string): boolean {
	return IDENTIFIER_RE.test(key) && !RESERVED_WORDS.has(key);
}

function formatPropertyKey(key: string): string {
	return canUseDirectSyntax(key) ? key : `[${escapeString(key)}]`;
}

/**
 * Render a SchemaNode to Valibot code
 */
export function render(node: SchemaNode): string {
	switch (node.kind) {
		case "string":
			return renderString(node);
		case "number":
			return renderNumber(node);
		case "boolean":
			return "v.boolean()";
		case "null":
			return "v.null_()";
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
		case "any":
			return "v.any()";
		case "never":
			return "v.never()";

		case "conditional":
			return renderConditional(node);
		case "typeGuarded":
			return renderTypeGuarded(node);
		case "nullable":
			return renderNullable(node);
		default:
			// Exhaustive check - should never reach here
			const _exhaustive: never = node;
			throw new Error(`Unhandled node kind: ${(node as any).kind}`);
	}
}

function renderString(node: StringNode): string {
	let result = "v.string()";

	// Format validations
	if (node.format) {
		switch (node.format) {
			case "email":
				result = "v.pipe(v.string(), v.email())";
				break;
			case "uri":
			case "uri-reference":
				result = "v.pipe(v.string(), v.url())";
				break;
			case "uuid":
				result = "v.pipe(v.string(), v.uuid())";
				break;
			case "date-time":
				result = "v.pipe(v.string(), v.isoDateTime())";
				break;
			case "date":
				result = "v.pipe(v.string(), v.isoDate())";
				break;
			case "time":
				result = "v.pipe(v.string(), v.isoTime())";
				break;
			case "ipv4":
				result = "v.pipe(v.string(), v.ipv4())";
				break;
			case "ipv6":
				result = "v.pipe(v.string(), v.ipv6())";
				break;
			case "hostname":
			case "idn-hostname":
				result = `v.pipe(v.string(), v.regex(/^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/))`;
				break;
			default:
				// Unknown format - ignore per JSON Schema spec
				result = "v.string()";
		}
	}

	// Constraints - use grapheme cluster counting per JSON Schema spec
	const { minLength, maxLength, pattern } = node.constraints;
	const hasConstraints = minLength !== undefined || maxLength !== undefined || pattern !== undefined;

	if (hasConstraints) {
		const actions: string[] = [];
		
		if (minLength !== undefined) {
			actions.push(`v.check((val) => [...new Intl.Segmenter().segment(val)].length >= ${minLength}, "String must have at least ${minLength} character(s)")`);
		}
		if (maxLength !== undefined) {
			actions.push(`v.check((val) => [...new Intl.Segmenter().segment(val)].length <= ${maxLength}, "String must have at most ${maxLength} character(s)")`);
		}
		if (pattern) {
			actions.push(`v.regex(new RegExp(${escapeString(pattern)}))`);
		}

		// If we already have format pipe, extend it
		if (node.format && node.format !== "hostname" && node.format !== "idn-hostname") {
			// Remove the closing paren and add actions
			result = result.slice(0, -1) + ", " + actions.join(", ") + ")";
		} else {
			// Create new pipe
			result = `v.pipe(v.string(), ${actions.join(", ")})`;
		}
	}

	return result;
}

function renderNumber(node: NumberNode): string {
	const actions: string[] = [];
	const {
		minimum,
		maximum,
		exclusiveMinimum,
		exclusiveMaximum,
		multipleOf,
	} = node.constraints;

	// Integer check
	if (node.integer) {
		actions.push("v.integer()");
	}

	// Range constraints - valibot uses separate functions for exclusive bounds
	if (minimum !== undefined) {
		actions.push(`v.minValue(${minimum})`);
	}
	if (exclusiveMinimum !== undefined) {
		// Use custom check for exclusive minimum since valibot doesn't have v.gt in all versions
		actions.push(`v.check((val) => val > ${exclusiveMinimum}, "Number must be greater than ${exclusiveMinimum}")`);
	}
	if (maximum !== undefined) {
		actions.push(`v.maxValue(${maximum})`);
	}
	if (exclusiveMaximum !== undefined) {
		// Use custom check for exclusive maximum since valibot doesn't have v.lt in all versions
		actions.push(`v.check((val) => val < ${exclusiveMaximum}, "Number must be less than ${exclusiveMaximum}")`);
	}

	// Multiple of
	if (multipleOf !== undefined) {
		// For small values, use epsilon-based comparison for float precision
		if (multipleOf < 1 && multipleOf > 0) {
			actions.push(`v.check((val) => Math.abs(val - Math.round(val / ${multipleOf}) * ${multipleOf}) < 1e-10, "Number must be a multiple of ${multipleOf}")`);
		} else {
			actions.push(`v.multipleOf(${multipleOf})`);
		}
	}

	if (actions.length === 0) {
		return "v.number()";
	}

	return `v.pipe(v.number(), ${actions.join(", ")})`;
}

// JSON Schema: type:object excludes arrays (unlike JS where typeof [] === 'object')
// Must use rawCheck BEFORE the object schema, since valibot's object schemas coerce arrays
const ARRAY_REJECTION_CHECK = `v.rawCheck((ctx) => { if (Array.isArray(ctx.dataset.value)) ctx.addIssue({ message: "Expected object, not array" }); })`;

function renderObject(node: ObjectNode): string {
	const propKeys = Array.from(node.properties.keys());
	const requiredKeys = propKeys.filter(
		(k) => node.properties.get(k)!.required,
	);

	// prototype property names break valibot's internal object entry lookup
	if (hasPrototypeProperties(propKeys) || hasPrototypeProperties(requiredKeys)) {
		return renderObjectWithProtoProps(node);
	}

	const hasPatternProps = node.patternProperties.length > 0;
	const needsPassthrough = hasPatternProps || node.propertyNames !== undefined;

	// Determine if additionalProperties is explicitly set (true, schema, or false)
	// unevaluatedProperties only takes effect when additionalProperties is NOT explicitly set (undefined)
	const additionalPropsExplicit = node.additionalProperties !== undefined;

	// If only pattern properties or property names validation, render superRefine
	if (propKeys.length === 0 && needsPassthrough) {
		return renderObjectWithPatternProps(node);
	}

	// Schema for validating extra properties
	// - additionalProperties as schema takes precedence
	// - Then unevaluatedProperties as schema (only if additionalProperties not set)
	const additionalSchema =
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
			? node.additionalProperties
			: !additionalPropsExplicit &&
				  typeof node.unevaluatedProperties === "object" &&
				  node.unevaluatedProperties.kind !== "any"
				? node.unevaluatedProperties
				: null;

	// isStrict: reject extra properties
	// - additionalProperties: false → strict
	// - additionalProperties not set AND unevaluatedProperties: false → strict
	const isStrict =
		node.additionalProperties === false ||
		(!additionalPropsExplicit && node.unevaluatedProperties === false);

	// If using record-style (no properties, just additionalProperties/unevaluatedProperties schema)
	if (propKeys.length === 0 && additionalSchema) {
		const valueSchema = render(additionalSchema);
		const record = `v.record(v.string(), ${valueSchema})`;
		const actions = renderObjectConstraintsActions(node);
		// Always wrap with array rejection check
		return `v.pipe(v.unknown(), ${ARRAY_REJECTION_CHECK}, ${record}${actions.length > 0 ? ", " + actions.join(", ") : ""})`;
	}

	// Build shape
	const shape = propKeys.map((key) => {
		const prop = node.properties.get(key)!;
		let propCode = render(prop.schema as SchemaNode);
		if (!prop.required) {
			propCode = `v.optional(${propCode})`;
		}
		return `${formatPropertyKey(key)}: ${propCode}`;
	});

	let result = "";

	// Choose object type based on additionalProperties/unevaluatedProperties
	// When patternProperties exist, we can't use strictObject (it would reject pattern-matched keys)
	if (isStrict && !hasPatternProps) {
		// Strict mode - no additional properties (only when no pattern props)
		result =
			shape.length > 0
				? `v.strictObject({ ${shape.join(", ")} })`
				: "v.strictObject({})";
	} else if (additionalSchema && !hasPatternProps) {
		// Validate additional properties against schema (only when no pattern props)
		const restSchema = render(additionalSchema);
		result =
			shape.length > 0
				? `v.objectWithRest({ ${shape.join(", ")} }, ${restSchema})`
				: `v.record(v.string(), ${restSchema})`;
	} else {
		// Allow additional properties (loose mode) - pattern props handling is done separately
		result =
			shape.length > 0
				? `v.looseObject({ ${shape.join(", ")} })`
				: "v.looseObject({})";
	}

	// Collect all validation actions (run AFTER the object schema)
	const postActions: string[] = [];

	// Pattern properties validation + additionalProperties/unevaluatedProperties check when both are present
	if (hasPatternProps) {
		postActions.push(...renderPatternPropsActionsWithAdditional(node, propKeys));
	}

	// If propertyNames is set without patternProperties, and we have unevaluatedProperties,
	// we need to validate unevaluated property values
	// (propertyNames validates keys but doesn't evaluate values)
	if (
		node.propertyNames &&
		!hasPatternProps &&
		!additionalPropsExplicit &&
		typeof node.unevaluatedProperties === "object" &&
		node.unevaluatedProperties.kind !== "any"
	) {
		const unevalSchema = render(node.unevaluatedProperties);
		const definedPropsJson = JSON.stringify(propKeys);
		postActions.push(`v.check((val) => {
      const definedProps = new Set(${definedPropsJson});
      for (const [key, value] of Object.entries(val)) {
        if (definedProps.has(key)) continue;
        const result = v.safeParse(${unevalSchema}, value);
        if (!result.success) return false;
      }
      return true;
    }, "Unevaluated property validation failed")`);
	}

	// Property names validation
	if (node.propertyNames) {
		const keySchema = render(node.propertyNames);
		postActions.push(`v.check((val) => {
      for (const key of Object.keys(val)) {
        const result = v.safeParse(${keySchema}, key);
        if (!result.success) return false;
      }
      return true;
    }, "Invalid property name")`);
	}

	// Required properties validation - only needed when v.any() props exist (which accept undefined)
	const requiredAnyProps: string[] = [];
	for (const key of propKeys) {
		const prop = node.properties.get(key)!;
		if (prop.required && (prop.schema as SchemaNode).kind === "any") {
			requiredAnyProps.push(key);
		}
	}
	if (requiredAnyProps.length > 0) {
		const requiredJson = JSON.stringify(requiredAnyProps);
		postActions.push(`v.check((val) => {
      for (const key of ${requiredJson}) {
        if (!Object.hasOwn(val, key)) return false;
      }
      return true;
    }, "Required property missing")`);
	}

	// Min/max properties
	postActions.push(...renderObjectConstraintsActions(node));

	// Dependencies
	postActions.push(...renderDependenciesActions(node));

	// Always wrap with array rejection check BEFORE the object schema
	const parts = ["v.unknown()", ARRAY_REJECTION_CHECK, result, ...postActions];
	return `v.pipe(${parts.join(", ")})`;
}

function renderObjectWithProtoProps(node: ObjectNode): string {
	const propKeys = Array.from(node.properties.keys());
	const validators: string[] = [];

	for (const key of propKeys) {
		const prop = node.properties.get(key)!;
		const propCode = render(prop.schema as SchemaNode);
		const keyExpr = escapeString(key);

		if (prop.required) {
			validators.push(`
      if (Object.hasOwn(val, ${keyExpr})) {
        if (!v.safeParse(${propCode}, val[${keyExpr}]).success) return false;
      } else {
        return false;
      }`);
		} else {
			validators.push(`
      if (Object.hasOwn(val, ${keyExpr})) {
        if (!v.safeParse(${propCode}, val[${keyExpr}]).success) return false;
      }`);
		}
	}

	// additional properties check
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
          if (!v.safeParse(${additionalSchema}, value).success) return false;
        }
      }`;
	}

	const postActions: string[] = [];
	postActions.push(...renderObjectConstraintsActions(node));
	postActions.push(...renderDependenciesActions(node));

	const parts = [
		"v.unknown()",
		ARRAY_REJECTION_CHECK,
		`v.check((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;${validators.join("")}${additionalCheck}
      return true;
    }, "Object validation failed")`,
		...postActions,
	];

	return `v.pipe(${parts.join(", ")})`;
}

function renderObjectWithPatternProps(node: ObjectNode): string {
	const patterns = node.patternProperties;
	const definedProps = JSON.stringify(Array.from(node.properties.keys()));

	// Check if additionalProperties is explicitly set
	const additionalPropsExplicit = node.additionalProperties !== undefined;
	// If additionalProperties is not set and unevaluatedProperties: false, reject non-matching props
	const shouldRejectNonMatching =
		node.additionalProperties === false ||
		(!additionalPropsExplicit && node.unevaluatedProperties === false);

	let checks: string[] = [];

	// Validate pattern properties
	patterns.forEach((p) => {
		const patternCode = render(p.schema as SchemaNode);
		const patternStr = escapeString(p.pattern);
		checks.push(`
      for (const [key, value] of Object.entries(val)) {
        if (new RegExp(${patternStr}).test(key)) {
          const result = v.safeParse(${patternCode}, value);
          if (!result.success) return false;
        }
      }`);
	});

	// Additional/unevaluated properties validation
	if (shouldRejectNonMatching) {
		checks.push(`
      const definedProps = new Set(${definedProps});
      const patterns = [${patterns.map((p) => `new RegExp(${escapeString(p.pattern)})`).join(", ")}];
      for (const key of Object.keys(val)) {
        if (definedProps.has(key)) continue;
        const matchesPattern = patterns.some(p => p.test(key));
        if (!matchesPattern) return false;
      }`);
	} else if (
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const additionalSchema = render(node.additionalProperties);
		checks.push(`
      const definedProps = new Set(${definedProps});
      const patterns = [${patterns.map((p) => `new RegExp(${escapeString(p.pattern)})`).join(", ")}];
      for (const [key, value] of Object.entries(val)) {
        if (definedProps.has(key)) continue;
        const matchesPattern = patterns.some(p => p.test(key));
        if (!matchesPattern) {
          const result = v.safeParse(${additionalSchema}, value);
          if (!result.success) return false;
        }
      }`);
	} else if (
		!additionalPropsExplicit &&
		typeof node.unevaluatedProperties === "object" &&
		node.unevaluatedProperties.kind !== "any"
	) {
		// unevaluatedProperties schema applies when additionalProperties is not set
		// This handles cases like { propertyNames: {...}, unevaluatedProperties: {...} }
		// where propertyNames validates keys but unevaluatedProperties validates values
		const unevalSchema = render(node.unevaluatedProperties);
		checks.push(`
      const definedProps = new Set(${definedProps});
      const patterns = [${patterns.map((p) => `new RegExp(${escapeString(p.pattern)})`).join(", ")}];
      for (const [key, value] of Object.entries(val)) {
        if (definedProps.has(key)) continue;
        const matchesPattern = patterns.some(p => p.test(key));
        if (!matchesPattern) {
          const result = v.safeParse(${unevalSchema}, value);
          if (!result.success) return false;
        }
      }`);
	}

	const allActions: string[] = [];
	allActions.push(`v.check((val) => {${checks.join("")}
      return true;
    }, "Object validation failed")`);

	// Property names validation
	if (node.propertyNames) {
		const keySchema = render(node.propertyNames);
		allActions.push(`v.check((val) => {
      for (const key of Object.keys(val)) {
        const result = v.safeParse(${keySchema}, key);
        if (!result.success) return false;
      }
      return true;
    }, "Invalid property name")`);
	}

	allActions.push(...renderObjectConstraintsActions(node));
	allActions.push(...renderDependenciesActions(node));

	// Wrap with array rejection check BEFORE the object schema
	return `v.pipe(v.unknown(), ${ARRAY_REJECTION_CHECK}, v.looseObject({}), ${allActions.join(", ")})`;
}

function renderPatternPropsActionsWithAdditional(
	node: ObjectNode,
	propKeys: string[],
): string[] {
	const patterns = node.patternProperties;
	const definedProps = JSON.stringify(propKeys);

	let checks: string[] = [];

	// Validate pattern properties
	patterns.forEach((p) => {
		const patternCode = render(p.schema as SchemaNode);
		const patternStr = escapeString(p.pattern);
		checks.push(`
      for (const [key, value] of Object.entries(val)) {
        if (new RegExp(${patternStr}).test(key)) {
          const result = v.safeParse(${patternCode}, value);
          if (!result.success) return false;
        }
      }`);
	});

	// Check if additionalProperties is explicitly set
	const additionalPropsExplicit = node.additionalProperties !== undefined;
	// If additionalProperties is not set and unevaluatedProperties: false, reject non-matching props
	const shouldRejectNonMatching =
		node.additionalProperties === false ||
		(!additionalPropsExplicit && node.unevaluatedProperties === false);

	// Handle additionalProperties/unevaluatedProperties when pattern props are present
	if (shouldRejectNonMatching) {
		checks.push(`
      const definedProps = new Set(${definedProps});
      const patterns = [${patterns.map((p) => `new RegExp(${escapeString(p.pattern)})`).join(", ")}];
      for (const key of Object.keys(val)) {
        if (definedProps.has(key)) continue;
        const matchesPattern = patterns.some(p => p.test(key));
        if (!matchesPattern) return false;
      }`);
	} else if (
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const additionalSchema = render(node.additionalProperties);
		checks.push(`
      const definedProps = new Set(${definedProps});
      const patterns = [${patterns.map((p) => `new RegExp(${escapeString(p.pattern)})`).join(", ")}];
      for (const [key, value] of Object.entries(val)) {
        if (definedProps.has(key)) continue;
        const matchesPattern = patterns.some(p => p.test(key));
        if (!matchesPattern) {
          const result = v.safeParse(${additionalSchema}, value);
          if (!result.success) return false;
        }
      }`);
	}

	return [`v.check((val) => {${checks.join("")}
      return true;
    }, "Object validation failed")`];
}

function renderObjectConstraintsActions(node: ObjectNode): string[] {
	const actions: string[] = [];

	if (node.minProperties !== undefined) {
		actions.push(`v.check((val) => Object.keys(val).length >= ${node.minProperties}, "Object must have at least ${node.minProperties} properties")`);
	}
	if (node.maxProperties !== undefined) {
		actions.push(`v.check((val) => Object.keys(val).length <= ${node.maxProperties}, "Object must have at most ${node.maxProperties} properties")`);
	}

	return actions;
}

function renderDependenciesActions(node: ObjectNode): string[] {
	const actions: string[] = [];

	for (const [prop, dep] of node.dependencies) {
		if (dep.kind === "property") {
			if (dep.requiredProperties.length > 0) {
				const message = escapeString(
					`Property ${prop} requires ${dep.requiredProperties.join(", ")}`,
				);
				actions.push(`v.check((val) => {
          if (Object.hasOwn(val, ${escapeString(prop)})) {
            return ${dep.requiredProperties.map((d) => `Object.hasOwn(val, ${escapeString(d)})`).join(" && ")};
          }
          return true;
        }, ${message})`);
			}
		} else {
			const depCode = render(dep.schema as SchemaNode);
			actions.push(`v.check((val) => {
        if (Object.hasOwn(val, ${escapeString(prop)})) {
          const result = v.safeParse(${depCode}, val);
          return result.success;
        }
        return true;
      }, "Schema dependency validation failed")`);
		}
	}

	return actions;
}

function renderArray(node: ArrayNode): string {
	// Check if this is a schema with only unevaluatedItems (no actual items schema)
	// In that case, all items are "unevaluated" and subject to unevaluatedItems constraint
	const hasRealItems = node.items.kind !== "any";

	if (!hasRealItems && node.unevaluatedItems === false) {
		// Schema like { "unevaluatedItems": false } - empty array only
		const constraints = renderArrayConstraints(node.constraints);
		if (constraints) {
			return `v.pipe(v.array(v.never())${constraints})`;
		}
		return `v.pipe(v.array(v.never()), v.maxLength(0))`;
	}

	if (!hasRealItems && node.unevaluatedItems !== undefined && node.unevaluatedItems !== false) {
		// Schema like { "unevaluatedItems": { "type": "string" } } - all items must match schema
		const unevalSchema = render(node.unevaluatedItems);
		const constraints = renderArrayConstraints(node.constraints);
		if (constraints) {
			return `v.pipe(v.array(${unevalSchema})${constraints})`;
		}
		return `v.array(${unevalSchema})`;
	}

	// Normal array with items schema
	const itemSchema = render(node.items);
	let result = `v.array(${itemSchema})`;

	// If items is defined AND unevaluatedItems is also defined, we need both validations
	// But typically items covers all items, so unevaluatedItems wouldn't have effect
	// Just in case, add the refinement
	if (
		hasRealItems &&
		node.unevaluatedItems !== undefined &&
		node.unevaluatedItems !== false &&
		node.unevaluatedItems.kind !== "any"
	) {
		const unevalSchema = render(node.unevaluatedItems);
		result = `v.pipe(${result}, v.check((arr) => {
      const schema = ${unevalSchema};
      for (const item of arr) {
        if (!v.safeParse(schema, item).success) return false;
      }
      return true;
    }, "Unevaluated item validation failed"))`;
	}

	const constraints = renderArrayConstraints(node.constraints);
	if (constraints) {
		if (result.startsWith("v.pipe(")) {
			// Already a pipe, extend it
			return result.replace(/\)$/, constraints + ")");
		}
		return `v.pipe(${result}${constraints})`;
	}
	return result;
}

function renderTuple(node: TupleNode): string {
	// JSON Schema semantics: prefixItems only validates items at positions IF they exist
	// Empty arrays and incomplete arrays (fewer items than prefixItems) are valid
	// valibot's v.tuple/v.tupleWithRest enforce minimum length, so we use v.array with custom validation
	const tupleSchemas = node.prefixItems.map((item) => render(item));
	const schemasArray = `[${tupleSchemas.join(", ")}]`;

	// Determine if extra items are allowed and what schema to use
	// unevaluatedItems takes precedence for standalone schemas (no applicators)
	const disallowExtraItems =
		node.restItems === false || node.unevaluatedItems === false;
	const extraItemsSchema =
		node.restItems !== false && node.restItems.kind !== "any"
			? node.restItems
			: node.unevaluatedItems !== undefined &&
				  node.unevaluatedItems !== false &&
				  node.unevaluatedItems.kind !== "any"
				? node.unevaluatedItems
				: null;

	let result: string;

	if (disallowExtraItems) {
		// No additional items allowed beyond prefix
		result = `v.pipe(v.array(v.any()), v.check((val) => {
      const schemas = ${schemasArray};
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = v.safeParse(schemas[i], val[i]);
        if (!itemResult.success) return false;
      }
      return val.length <= schemas.length;
    }, "Tuple validation failed"))`;
	} else if (extraItemsSchema) {
		// Rest/unevaluated items must match schema
		const restSchema = render(extraItemsSchema);
		result = `v.pipe(v.array(v.any()), v.check((val) => {
      const schemas = ${schemasArray};
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = v.safeParse(schemas[i], val[i]);
        if (!itemResult.success) return false;
      }
      const restSchema = ${restSchema};
      for (let i = schemas.length; i < val.length; i++) {
        const itemResult = v.safeParse(restSchema, val[i]);
        if (!itemResult.success) return false;
      }
      return true;
    }, "Tuple validation failed"))`;
	} else {
		// Allow any extra items
		result = `v.pipe(v.array(v.any()), v.check((val) => {
      const schemas = ${schemasArray};
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = v.safeParse(schemas[i], val[i]);
        if (!itemResult.success) return false;
      }
      return true;
    }, "Tuple items validation failed"))`;
	}

	const constraints = renderArrayConstraints(node.constraints);
	if (constraints) {
		// Remove the outer v.pipe and add constraints to it
		result = result.replace(/^v\.pipe\(/, "v.pipe(").replace(/\)$/, constraints + ")");
	}
	return result;
}

function renderArrayConstraints(
	constraints: ArrayNode["constraints"],
): string {
	const actions: string[] = [];

	if (constraints.minItems !== undefined) {
		actions.push(`v.minLength(${constraints.minItems})`);
	}
	if (constraints.maxItems !== undefined) {
		actions.push(`v.maxLength(${constraints.maxItems})`);
	}

	if (constraints.uniqueItems) {
		actions.push(`v.check((arr) => {
      const seen = new Set();
      for (const item of arr) {
        const key = JSON.stringify(item);
        if (seen.has(key)) return false;
        seen.add(key);
      }
      return true;
    }, "Array items must be unique")`);
	}

	if (constraints.contains) {
		const containsSchema = render(constraints.contains.schema as SchemaNode);
		const minContains = constraints.contains.minContains;
		const maxContains = constraints.contains.maxContains;

		if (maxContains !== undefined) {
			actions.push(`v.check((arr) => {
        let count = 0;
        for (const item of arr) {
          if (v.safeParse(${containsSchema}, item).success) count++;
        }
        return count >= ${minContains} && count <= ${maxContains};
      }, "Array must contain between ${minContains} and ${maxContains} items matching schema")`);
		} else {
			actions.push(`v.check((arr) => {
        let count = 0;
        for (const item of arr) {
          if (v.safeParse(${containsSchema}, item).success) count++;
        }
        return count >= ${minContains};
      }, "Array must contain at least ${minContains} item(s) matching schema")`);
		}
	}

	if (actions.length === 0) {
		return "";
	}

	return `, ${actions.join(", ")}`;
}

function renderUnion(node: UnionNode): string {
	if (node.variants.length === 0) return "v.never()";

	// Filter out "never" variants since they can't match anything
	const filtered = node.variants.filter((v) => v.kind !== "never");
	if (filtered.length === 0) return "v.never()";
	if (filtered.length === 1) return render(filtered[0]!);

	// Detect nullable pattern: union of [T, null] -> v.nullable(T)
	if (filtered.length === 2) {
		const nullIndex = filtered.findIndex((v) => v.kind === "null");
		if (nullIndex !== -1) {
			const otherIndex = nullIndex === 0 ? 1 : 0;
			return `v.nullable(${render(filtered[otherIndex]!)})`;
		}
	}

	const schemas = filtered.map((v) => render(v));
	return `v.union([${schemas.join(", ")}])`;
}

function renderIntersection(node: IntersectionNode): string {
	if (node.schemas.length === 0) return "v.any()";

	// Short-circuit: if ANY schema is never, intersection is never
	if (node.schemas.some((s) => s.kind === "never")) {
		return "v.never()";
	}

	// Filter out "any" schemas since they don't constrain the intersection
	const filtered = node.schemas.filter((s) => s.kind !== "any");
	if (filtered.length === 0) return "v.any()";
	if (filtered.length === 1) return render(filtered[0]!);

	const schemas = filtered.map((s) => render(s));
	return `v.intersect([${schemas.join(", ")}])`;
}

function renderOneOf(node: OneOfNode): string {
	if (node.schemas.length === 0) return "v.never()";
	if (node.schemas.length === 1) return render(node.schemas[0]!);

	// Filter out "never" schemas since they can never match
	const filtered = node.schemas.filter((s) => s.kind !== "never");

	// If all schemas are never, nothing can match exactly one
	if (filtered.length === 0) return "v.never()";

	// If exactly one schema remains after filtering never, it must match
	if (filtered.length === 1) return render(filtered[0]!);

	// Count how many "any" schemas there are - if > 1, always matches multiple
	const anyCount = filtered.filter((s) => s.kind === "any").length;
	if (anyCount > 1) {
		// Multiple "any" schemas means any value matches multiple
		return `v.pipe(v.any(), v.check(() => false, "oneOf has multiple 'true' schemas - impossible to match exactly one"))`;
	}

	const schemas = filtered.map((s) => render(s));
	return `v.pipe(v.any(), v.check((val) => {
    const schemas = [${schemas.join(", ")}];
    const results = schemas.map(s => v.safeParse(s, val));
    const validCount = results.filter(r => r.success).length;
    return validCount === 1;
  }, "Value must match exactly one schema in oneOf"))`;
}

function renderNot(node: NotNode): string {
	const schema = render(node.schema);
	return `v.pipe(v.any(), v.check((val) => !v.safeParse(${schema}, val).success, "Value must not match schema"))`;
}

function renderLiteral(node: LiteralNode): string {
	if (isPrimitive(node.value)) {
		return `v.literal(${JSON.stringify(node.value)})`;
	}

	// Objects/arrays need deep equality check
	const isArray = Array.isArray(node.value);
	const baseType = isArray ? "v.array(v.any())" : "v.looseObject({})";
	const sorted = sortedStringify(node.value);

	return `v.pipe(${baseType}, v.check((val) => ${DEEP_SORTED_STRINGIFY_RUNTIME}(val) === ${JSON.stringify(sorted)}, "Value must equal the const value"))`;
}

function renderEnum(node: EnumNode): string {
	const values = node.values;

	if (values.length === 0) return "v.never()";

	if (values.length === 1) {
		return renderLiteral({ kind: "literal", value: values[0] });
	}

	// All strings - use v.picklist
	if (values.every((v) => typeof v === "string")) {
		return `v.picklist([${values.map((v) => JSON.stringify(v)).join(", ")}])`;
	}

	// Check for complex values (objects/arrays)
	const hasComplexValues = values.some((v) => !isPrimitive(v));

	if (hasComplexValues) {
		// Stringify all values for comparison (sorted keys for objects)
		const sortedValues = values.map((v) =>
			JSON.stringify(
				JSON.stringify(
					v,
					v != null && typeof v === "object"
						? Object.keys(v as object).sort()
						: undefined,
				),
			),
		);
		return `v.pipe(v.any(), v.check((val) => {
      const sorted = JSON.stringify(val, val != null && typeof val === "object" ? Object.keys(val as object).sort() : undefined);
      return [${sortedValues.join(", ")}].includes(sorted);
    }, "Value must match one of the enum values"))`;
	}

	// Mixed primitives - use union of literals
	const literals = values.map((v) => `v.literal(${JSON.stringify(v)})`);
	return `v.union([${literals.join(", ")}])`;
}



function renderConditional(node: ConditionalNode): string {
	const ifSchema = render(node.if);
	const thenSchema = node.then ? render(node.then) : null;
	const elseSchema = node.else ? render(node.else) : null;

	if (thenSchema && elseSchema) {
		// Both then and else present: validate with then if condition matches, else otherwise
		return `v.pipe(v.any(), v.check((val) => {
      const ifResult = v.safeParse(${ifSchema}, val);
      if (ifResult.success) {
        const thenResult = v.safeParse(${thenSchema}, val);
        return thenResult.success;
      } else {
        const elseResult = v.safeParse(${elseSchema}, val);
        return elseResult.success;
      }
    }, "Conditional validation failed"))`;
	} else if (thenSchema) {
		// Only then present: validate with then if condition matches
		return `v.pipe(v.any(), v.check((val) => {
      const ifResult = v.safeParse(${ifSchema}, val);
      if (ifResult.success) {
        const thenResult = v.safeParse(${thenSchema}, val);
        return thenResult.success;
      }
      return true;
    }, "Conditional 'then' validation failed"))`;
	} else if (elseSchema) {
		// Only else present: validate with else if condition doesn't match
		return `v.pipe(v.any(), v.check((val) => {
      const ifResult = v.safeParse(${ifSchema}, val);
      if (!ifResult.success) {
        const elseResult = v.safeParse(${elseSchema}, val);
        return elseResult.success;
      }
      return true;
    }, "Conditional 'else' validation failed"))`;
	}

	// if without then/else has no effect
	return "v.any()";
}

function renderTypeGuarded(node: TypeGuardedNode): string {
	if (node.guards.length === 0) return "v.any()";

	const checks: string[] = [];

	for (const guard of node.guards) {
		const schema = render(guard.schema);
		switch (guard.check) {
			case "object":
				checks.push(`if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = v.safeParse(${schema}, val);
          if (!result.success) return false;
        }`);
				break;
			case "array":
				checks.push(`if (Array.isArray(val)) {
          const result = v.safeParse(${schema}, val);
          if (!result.success) return false;
        }`);
				break;
			case "string":
				checks.push(`if (typeof val === "string") {
          const result = v.safeParse(${schema}, val);
          if (!result.success) return false;
        }`);
				break;
			case "number":
				checks.push(`if (typeof val === "number") {
          const result = v.safeParse(${schema}, val);
          if (!result.success) return false;
        }`);
				break;
		}
	}

	return `v.pipe(v.any(), v.check((val) => {
        ${checks.join("\n        ")}
        return true;
      }, "Type-guarded validation failed"))`;
}

function renderNullable(node: NullableNode): string {
	return `v.nullable(${render(node.inner)})`;
}
