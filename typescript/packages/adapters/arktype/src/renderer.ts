/**
 * ArkType Renderer
 * Converts SchemaNode IR to ArkType code strings using fluent API
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
	PatternPropertyDef,
} from "@xschemadev/core";
import {
	escapeString,
	isPrimitive,
	sortedStringify,
	hasPrototypeProperties,
	DEEP_SORTED_STRINGIFY_RUNTIME,
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

function formatPropertyKey(key: string, optional: boolean): string {
	if (optional) {
		return canUseDirectSyntax(key) ? `"${key}?"` : `${escapeString(key + "?")}`;
	}
	return canUseDirectSyntax(key) ? key : escapeString(key);
}

/**
 * Render a SchemaNode to ArkType code
 */
export function render(node: SchemaNode): string {
	switch (node.kind) {
		case "string":
			return renderString(node);
		case "number":
			return renderNumber(node);
		case "boolean":
			return "type.boolean";
		case "null":
			return "type.null";
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
			return "type.unknown";
		case "never":
			return "type.never";

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
	const { minLength, maxLength, pattern } = node.constraints;
	const parts: string[] = [];

	// Start with format-specific type or base string
	let base = "type.string";
	if (node.format) {
		switch (node.format) {
			case "email":
				base = 'type("string.email")';
				break;
			case "uri":
			case "uri-reference":
				base = 'type("string.url")';
				break;
			case "uuid":
				base = 'type("string.uuid")';
				break;
			case "date-time":
				base = 'type("string.date.iso")';
				break;
			case "date":
				base = 'type("string.date")';
				break;
			case "ipv4":
				base = 'type("string.ip.v4")';
				break;
			case "ipv6":
				base = 'type("string.ip.v6")';
				break;
			case "hostname":
			case "idn-hostname":
				base = `type(/^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/)`;
				break;
			case "time":
				// No native arktype time format, use regex
				base = `type(/^\\d{2}:\\d{2}:\\d{2}(?:\\.\\d+)?(?:Z|[+-]\\d{2}:\\d{2})?$/)`;
				break;
		}
	}

	// Pattern constraint - use regex type
	if (pattern) {
		if (base === "type.string") {
			base = `type(new RegExp(${escapeString(pattern)}))`;
		} else {
			// Need to intersect pattern with format
			parts.push(`.and(type(new RegExp(${escapeString(pattern)})))`);
		}
	}

	// Length constraints - use grapheme cluster counting per JSON Schema spec
	if (minLength !== undefined || maxLength !== undefined) {
		const checks: string[] = [];
		if (minLength !== undefined) {
			checks.push(
				`[...new Intl.Segmenter().segment(s)].length >= ${minLength}`,
			);
		}
		if (maxLength !== undefined) {
			checks.push(
				`[...new Intl.Segmenter().segment(s)].length <= ${maxLength}`,
			);
		}
		parts.push(`.narrow((s, ctx) => (${checks.join(" && ")}) || ctx.mustBe("a string with valid length"))`);
	}

	return base + parts.join("");
}

function renderNumber(node: NumberNode): string {
	const { minimum, maximum, exclusiveMinimum, exclusiveMaximum, multipleOf } =
		node.constraints;

	// Build range expression if we have bounds
	const hasMinInclusive = minimum !== undefined;
	const hasMaxInclusive = maximum !== undefined;
	const hasMinExclusive = exclusiveMinimum !== undefined;
	const hasMaxExclusive = exclusiveMaximum !== undefined;

	// arktype v2 uses "number % 1" for integers (divisible by 1)
	const baseType = "number";
	let result: string;

	// Try to build a range expression
	if (
		(hasMinInclusive || hasMinExclusive) &&
		(hasMaxInclusive || hasMaxExclusive)
	) {
		// Both bounds - use range syntax
		const minVal = hasMinExclusive ? exclusiveMinimum : minimum;
		const maxVal = hasMaxExclusive ? exclusiveMaximum : maximum;
		const minOp = hasMinExclusive ? "<" : "<=";
		const maxOp = hasMaxExclusive ? "<" : "<=";
		result = `type("${minVal} ${minOp} ${baseType} ${maxOp} ${maxVal}")`;
	} else if (hasMinInclusive) {
		result = `type("${baseType} >= ${minimum}")`;
	} else if (hasMinExclusive) {
		result = `type("${baseType} > ${exclusiveMinimum}")`;
	} else if (hasMaxInclusive) {
		result = `type("${baseType} <= ${maximum}")`;
	} else if (hasMaxExclusive) {
		result = `type("${baseType} < ${exclusiveMaximum}")`;
	} else {
		result = "type.number";
	}

	// Add integer constraint using divisor (number % 1)
	if (node.integer) {
		if (result === "type.number") {
			result = 'type("number % 1")';
		} else {
			result += '.and(type("number % 1"))';
		}
	}

	// multipleOf constraint
	if (multipleOf !== undefined) {
		if (Number.isInteger(multipleOf) && multipleOf > 0) {
			// Use arktype's divisor syntax for integers
			result = result === "type.number" ? `type("number % ${multipleOf}")` : `${result}.and(type("number % ${multipleOf}"))`;
		} else {
			// For non-integer multipleOf, use narrow with epsilon comparison
			result += `.narrow((n, ctx) => Math.abs(n - Math.round(n / ${multipleOf}) * ${multipleOf}) < 1e-10 || ctx.mustBe("a multiple of ${multipleOf}"))`;
		}
	}

	return result;
}

function renderObject(node: ObjectNode): string {
	const propKeys: string[] = Array.from(node.properties.keys());
	const hasPatternProps = node.patternProperties.length > 0;
	const hasProtoProps =
		hasPrototypeProperties(propKeys) ||
		hasPrototypeProperties(
			propKeys.filter((k) => node.properties.get(k)!.required),
		);

	// If we have prototype properties, we need special handling
	if (hasProtoProps) {
		return renderObjectWithProtoProps(node);
	}

	// Determine strictness from additionalProperties and unevaluatedProperties
	// Key insight: if additionalProperties is explicitly set (true, schema, or false), it handles extra props.
	// unevaluatedProperties only takes effect when additionalProperties is NOT explicitly set (undefined).
	const additionalPropsExplicit = node.additionalProperties !== undefined;

	// isStrict: reject extra properties
	// - additionalProperties: false → strict
	// - additionalProperties not set AND unevaluatedProperties: false → strict
	const isStrict =
		node.additionalProperties === false ||
		(!additionalPropsExplicit && node.unevaluatedProperties === false);

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

	let result: string;

	if (propKeys.length === 0 && !hasPatternProps) {
		// No properties defined
		if (isStrict) {
			result = 'type({ "+": "reject" })';
		} else if (additionalSchema) {
			const valueSchema = render(additionalSchema);
			result = `type({ "[string]": ${valueSchema} })`;
		} else {
			result = "type.object";
		}
	} else {
		// Build object shape
		const shape: string[] = [];

		for (const key of propKeys) {
			const prop = node.properties.get(key)!;
			const propCode = render(prop.schema as SchemaNode);
			const keyStr = formatPropertyKey(key, !prop.required);
			shape.push(`${keyStr}: ${propCode}`);
		}

		// Handle strict mode (additionalProperties: false or unevaluatedProperties: false)
		if (isStrict && !hasPatternProps) {
			shape.push('"+": "reject"');
		}
		// Note: additionalProperties/unevaluatedProperties with schema is handled below via narrow
		// to validate only additional keys (not the defined properties)

		result = `type({ ${shape.join(", ")} })`;
	}

	// Handle additionalProperties/unevaluatedProperties with schema - validate only additional keys
	if (additionalSchema && !hasPatternProps && propKeys.length > 0) {
		const valueSchema = render(additionalSchema);
		const definedPropsJson = JSON.stringify(propKeys);
		result += `.narrow((obj, ctx) => {
      const definedProps = new Set(${definedPropsJson});
      for (const [key, value] of Object.entries(obj)) {
        if (!definedProps.has(key)) {
          if (!${valueSchema}.allows(value)) return ctx.mustBe("valid additional properties");
        }
      }
      return true;
    })`;
	}

	// Pattern properties need narrow
	if (hasPatternProps) {
		result += renderPatternPropsNarrow(node, propKeys);
	}

	// Property names validation
	if (node.propertyNames) {
		const keySchema = render(node.propertyNames);
		result += `.narrow((obj, ctx) => {
      for (const key of Object.keys(obj)) {
        if (!${keySchema}.allows(key)) {
          return ctx.mustBe("an object with valid property names");
        }
      }
      return true;
    })`;
	}

	// Min/max properties
	if (node.minProperties !== undefined) {
		result += `.narrow((obj, ctx) => Object.keys(obj).length >= ${node.minProperties} || ctx.mustBe("an object with at least ${node.minProperties} properties"))`;
	}
	if (node.maxProperties !== undefined) {
		result += `.narrow((obj, ctx) => Object.keys(obj).length <= ${node.maxProperties} || ctx.mustBe("an object with at most ${node.maxProperties} properties"))`;
	}

	// Dependencies
	result += renderDependencies(node);

	// Exclude arrays - JSON Schema type: object does not include arrays
	result += `.narrow((val, ctx) => !Array.isArray(val) || ctx.mustBe("an object, not an array"))`;

	return result;
}

function renderObjectWithProtoProps(node: ObjectNode): string {
	const propKeys: string[] = Array.from(node.properties.keys());

	// Build validation using narrow on unknown object
	const validators: string[] = [];

	for (const key of propKeys) {
		const prop = node.properties.get(key)!;
		const propCode = render(prop.schema as SchemaNode);
		const keyExpr = escapeString(key);

		if (prop.required) {
			validators.push(`
        if (Object.hasOwn(obj, ${keyExpr})) {
          if (!${propCode}.allows(obj[${keyExpr}])) return ctx.mustBe("valid at property ${key}");
        } else {
          return ctx.mustBe("an object with required property ${key}");
        }`);
		} else {
			validators.push(`
        if (Object.hasOwn(obj, ${keyExpr})) {
          if (!${propCode}.allows(obj[${keyExpr}])) return ctx.mustBe("valid at property ${key}");
        }`);
		}
	}

	// Additional properties check
	if (node.additionalProperties === false) {
		const definedPropsJson = JSON.stringify(propKeys);
		validators.push(`
        const definedProps = new Set(${definedPropsJson});
        for (const key of Object.keys(obj)) {
          if (!definedProps.has(key)) return ctx.mustBe("an object without additional properties");
        }`);
	} else if (
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const additionalSchema = render(node.additionalProperties);
		const definedPropsJson = JSON.stringify(propKeys);
		validators.push(`
        const definedProps = new Set(${definedPropsJson});
        for (const [key, value] of Object.entries(obj)) {
          if (!definedProps.has(key)) {
            if (!${additionalSchema}.allows(value)) return ctx.mustBe("valid additional properties");
          }
        }`);
	}

	return `type.object.narrow((obj, ctx) => {
      if (Array.isArray(obj)) return ctx.mustBe("an object, not an array");${validators.join("")}
      return true;
    })`;
}

function renderPatternPropsNarrow(
	node: ObjectNode,
	propKeys: string[],
): string {
	const patterns = node.patternProperties;
	const definedPropsJson = JSON.stringify(propKeys);

	// Check if additionalProperties is explicitly set
	const additionalPropsExplicit = node.additionalProperties !== undefined;
	// If additionalProperties is not set and unevaluatedProperties: false, reject non-matching props
	const shouldRejectNonMatching =
		node.additionalProperties === false ||
		(!additionalPropsExplicit && node.unevaluatedProperties === false);

	let body = `
      const definedProps = new Set(${definedPropsJson});
      const patterns = [${patterns.map((p) => `new RegExp(${escapeString(p.pattern)})`).join(", ")}];
    `;

	// Validate pattern properties
	patterns.forEach((p, i) => {
		const patternCode = render(p.schema as SchemaNode);
		body += `
      for (const [key, value] of Object.entries(obj)) {
        if (patterns[${i}].test(key)) {
          if (!${patternCode}.allows(value)) return ctx.mustBe("valid for pattern ${p.pattern}");
        }
      }`;
	});

	// Additional/unevaluated properties validation
	if (shouldRejectNonMatching) {
		body += `
      for (const key of Object.keys(obj)) {
        if (definedProps.has(key)) continue;
        const matchesPattern = patterns.some(p => p.test(key));
        if (!matchesPattern) return ctx.mustBe("an object without additional properties");
      }`;
	} else if (
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const additionalSchema = render(node.additionalProperties);
		body += `
      for (const [key, value] of Object.entries(obj)) {
        if (definedProps.has(key)) continue;
        const matchesPattern = patterns.some(p => p.test(key));
        if (!matchesPattern) {
          if (!${additionalSchema}.allows(value)) return ctx.mustBe("valid additional properties");
        }
      }`;
	}

	return `.narrow((obj, ctx) => {${body}
      return true;
    })`;
}

function renderDependencies(node: ObjectNode): string {
	let result = "";

	for (const [prop, dep] of node.dependencies) {
		if (dep.kind === "property") {
			if (dep.requiredProperties.length > 0) {
				const deps = dep.requiredProperties
					.map((d) => `Object.hasOwn(obj, ${escapeString(d)})`)
					.join(" && ");
				const escapedProp = JSON.stringify(prop).slice(1, -1);
				const escapedDeps = dep.requiredProperties.map((d) => JSON.stringify(d).slice(1, -1)).join(", ");
				result += `.narrow((obj, ctx) => {
          if (Object.hasOwn(obj, ${escapeString(prop)})) {
            return (${deps}) || ctx.mustBe("an object where ${escapedProp} requires ${escapedDeps}");
          }
          return true;
        })`;
			}
		} else {
			const depCode = render(dep.schema as SchemaNode);
			result += `.narrow((obj, ctx) => {
        if (Object.hasOwn(obj, ${escapeString(prop)})) {
          return ${depCode}.allows(obj) || ctx.mustBe("valid for dependency schema");
        }
        return true;
      })`;
		}
	}

	return result;
}

function renderArray(node: ArrayNode): string {
	// Check if this is a schema with only unevaluatedItems (no actual items schema)
	// In that case, all items are "unevaluated" and subject to unevaluatedItems constraint
	const hasRealItems = node.items.kind !== "any";

	if (!hasRealItems && node.unevaluatedItems === false) {
		// Schema like { "unevaluatedItems": false } - empty array only
		let result = "type.unknown.array().narrow((arr, ctx) => arr.length === 0 || ctx.mustBe(\"an empty array\"))";
		result += renderArrayConstraints(node.constraints);
		return result;
	}

	if (!hasRealItems && node.unevaluatedItems !== undefined && node.unevaluatedItems !== false) {
		// Schema like { "unevaluatedItems": { "type": "string" } } - all items must match schema
		const unevalSchema = render(node.unevaluatedItems);
		let result = `${unevalSchema}.array()`;
		result += renderArrayConstraints(node.constraints);
		return result;
	}

	// Normal array with items schema
	const itemSchema = render(node.items);
	let result = `${itemSchema}.array()`;

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
		result += `.narrow((arr, ctx) => {
      const schema = ${unevalSchema};
      for (let i = 0; i < arr.length; i++) {
        if (!schema.allows(arr[i])) return ctx.mustBe("valid unevaluated items");
      }
      return true;
    })`;
	}

	result += renderArrayConstraints(node.constraints);
	return result;
}

function renderTuple(node: TupleNode): string {
	const prefixSchemas = node.prefixItems.map((item) => render(item));

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

	// If no extra items allowed - allow 0 to prefixItems.length items
	// JSON Schema allows fewer items, just not more
	if (disallowExtraItems) {
		let result = `type.unknown.array().narrow((arr, ctx) => {
      const schemas = [${prefixSchemas.join(", ")}];
      if (arr.length > schemas.length) return ctx.mustBe("an array with at most " + schemas.length + " items");
      for (let i = 0; i < arr.length; i++) {
        if (!schemas[i].allows(arr[i])) return ctx.mustBe("valid tuple items");
      }
      return true;
    })`;
		result += renderArrayConstraints(node.constraints);
		return result;
	}

	// With rest/unevaluated items of "any" type - validate prefix then allow anything else
	if (!extraItemsSchema) {
		let result = `type.unknown.array().narrow((arr, ctx) => {
      const schemas = [${prefixSchemas.join(", ")}];
      for (let i = 0; i < Math.min(arr.length, schemas.length); i++) {
        if (!schemas[i].allows(arr[i])) return ctx.mustBe("valid tuple items");
      }
      return true;
    })`;
		result += renderArrayConstraints(node.constraints);
		return result;
	}

	// Tuple with typed rest/unevaluated items
	const restSchema = render(extraItemsSchema);
	let result = `type.unknown.array().narrow((arr, ctx) => {
      const schemas = [${prefixSchemas.join(", ")}];
      for (let i = 0; i < Math.min(arr.length, schemas.length); i++) {
        if (!schemas[i].allows(arr[i])) return ctx.mustBe("valid tuple items");
      }
      const restSchema = ${restSchema};
      for (let i = schemas.length; i < arr.length; i++) {
        if (!restSchema.allows(arr[i])) return ctx.mustBe("valid rest items");
      }
      return true;
    })`;
	result += renderArrayConstraints(node.constraints);
	return result;
}

function renderArrayConstraints(
	constraints: ArrayNode["constraints"],
): string {
	let result = "";

	// Use arktype's array length syntax where possible
	const { minItems, maxItems, uniqueItems, contains } = constraints;

	if (minItems !== undefined && maxItems !== undefined) {
		// Can't easily express this in arktype's DSL, use narrow
		result += `.narrow((arr, ctx) => arr.length >= ${minItems} && arr.length <= ${maxItems} || ctx.mustBe("an array with ${minItems}-${maxItems} items"))`;
	} else if (minItems !== undefined) {
		result += `.narrow((arr, ctx) => arr.length >= ${minItems} || ctx.mustBe("an array with at least ${minItems} items"))`;
	} else if (maxItems !== undefined) {
		result += `.narrow((arr, ctx) => arr.length <= ${maxItems} || ctx.mustBe("an array with at most ${maxItems} items"))`;
	}

	if (uniqueItems) {
		result += `.narrow((arr, ctx) => {
      const seen = new Set();
      for (const item of arr) {
        const key = JSON.stringify(item);
        if (seen.has(key)) return ctx.mustBe("an array with unique items");
        seen.add(key);
      }
      return true;
    })`;
	}

	if (contains) {
		const containsSchema = render(contains.schema as SchemaNode);
		const minContains = contains.minContains;
		const maxContains = contains.maxContains;

		if (maxContains !== undefined) {
			result += `.narrow((arr, ctx) => {
        let count = 0;
        for (const item of arr) {
          if (${containsSchema}.allows(item)) count++;
        }
        return (count >= ${minContains} && count <= ${maxContains}) || ctx.mustBe("an array containing ${minContains}-${maxContains} matching items");
      })`;
		} else {
			result += `.narrow((arr, ctx) => {
        let count = 0;
        for (const item of arr) {
          if (${containsSchema}.allows(item)) count++;
        }
        return count >= ${minContains} || ctx.mustBe("an array containing at least ${minContains} matching items");
      })`;
		}
	}

	return result;
}

function renderUnion(node: UnionNode): string {
	if (node.variants.length === 0) return "type.never";

	const filtered = node.variants.filter((v) => v.kind !== "never");
	if (filtered.length === 0) return "type.never";
	if (filtered.length === 1) return render(filtered[0]!);

	const schemas = filtered.map((v) => render(v));
	return schemas.reduce((acc: string, s: string) => `${acc}.or(${s})`);
}

function renderIntersection(node: IntersectionNode): string {
	if (node.schemas.length === 0) return "type.unknown";

	if (node.schemas.some((s) => s.kind === "never")) {
		return "type.never";
	}

	const filtered = node.schemas.filter((s) => s.kind !== "any");
	if (filtered.length === 0) return "type.unknown";
	if (filtered.length === 1) return render(filtered[0]!);

	const schemas = filtered.map((s) => render(s));
	return schemas.reduce((acc: string, s: string) => `${acc}.and(${s})`);
}

function renderOneOf(node: OneOfNode): string {
	if (node.schemas.length === 0) return "type.never";
	if (node.schemas.length === 1) return render(node.schemas[0]!);

	const filtered = node.schemas.filter((s) => s.kind !== "never");
	if (filtered.length === 0) return "type.never";
	if (filtered.length === 1) return render(filtered[0]!);

	const anyCount = filtered.filter((s) => s.kind === "any").length;
	if (anyCount > 1) {
		return `type.unknown.narrow(() => false)`;
	}

	const schemas = filtered.map((s) => render(s));
	return `type.unknown.narrow((val, ctx) => {
    const schemas = [${schemas.join(", ")}];
    const validCount = schemas.filter(s => s.allows(val)).length;
    if (validCount === 0) return ctx.mustBe("matching exactly one schema (matched none)");
    if (validCount > 1) return ctx.mustBe("matching exactly one schema (matched multiple)");
    return true;
  })`;
}

function renderNot(node: NotNode): string {
	const schema = render(node.schema);
	return `type.unknown.narrow((val, ctx) => !${schema}.allows(val) || ctx.mustBe("not matching the excluded schema"))`;
}

function renderLiteral(node: LiteralNode): string {
	if (node.value === null) {
		return "type.null";
	}

	if (isPrimitive(node.value)) {
		return `type.unit(${JSON.stringify(node.value)})`;
	}

	// Objects/arrays need deep equality with recursive key normalization
	const sorted = sortedStringify(node.value);
	return `type.unknown.narrow((val, ctx) => ${DEEP_SORTED_STRINGIFY_RUNTIME}(val) === ${JSON.stringify(sorted)} || ctx.mustBe("equal to the const value"))`;
}

function renderEnum(node: EnumNode): string {
	const values = node.values;

	if (values.length === 0) return "type.never";
	if (values.length === 1) {
		return renderLiteral({ kind: "literal", value: values[0] });
	}

	const hasComplexValues = values.some((v) => !isPrimitive(v));

	if (hasComplexValues) {
		const sortedValues = values.map((v) => sortedStringify(v));
		const valuesArrayCode = `[${sortedValues.map((v) => JSON.stringify(v)).join(", ")}]`;

		return `type.unknown.narrow((val, ctx) => {
      const validValues = ${valuesArrayCode};
      return validValues.includes(${DEEP_SORTED_STRINGIFY_RUNTIME}(val)) || ctx.mustBe("one of the enum values");
    })`;
	}

	return `type.enumerated(${values.map((v) => JSON.stringify(v)).join(", ")})`;
}



function renderConditional(node: ConditionalNode): string {
	const ifSchema = render(node.if);
	const thenSchema = node.then ? render(node.then) : null;
	const elseSchema = node.else ? render(node.else) : null;

	if (thenSchema && elseSchema) {
		return `type.unknown.narrow((val, ctx) => {
      if (${ifSchema}.allows(val)) {
        return ${thenSchema}.allows(val) || ctx.mustBe("valid for then branch");
      } else {
        return ${elseSchema}.allows(val) || ctx.mustBe("valid for else branch");
      }
    })`;
	} else if (thenSchema) {
		return `type.unknown.narrow((val, ctx) => {
      if (${ifSchema}.allows(val)) {
        return ${thenSchema}.allows(val) || ctx.mustBe("valid for then branch");
      }
      return true;
    })`;
	} else if (elseSchema) {
		return `type.unknown.narrow((val, ctx) => {
      if (!${ifSchema}.allows(val)) {
        return ${elseSchema}.allows(val) || ctx.mustBe("valid for else branch");
      }
      return true;
    })`;
	}

	// if without then/else has no effect
	return "type.unknown";
}

function renderTypeGuarded(node: TypeGuardedNode): string {
	if (node.guards.length === 0) return "type.unknown";

	const checks: string[] = [];

	for (const guard of node.guards) {
		const schema = render(guard.schema);
		switch (guard.check) {
			case "object":
				checks.push(`if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          if (!${schema}.allows(val)) return ctx.mustBe("a valid object");
        }`);
				break;
			case "array":
				checks.push(`if (Array.isArray(val)) {
          if (!${schema}.allows(val)) return ctx.mustBe("a valid array");
        }`);
				break;
			case "string":
				checks.push(`if (typeof val === "string") {
          if (!${schema}.allows(val)) return ctx.mustBe("a valid string");
        }`);
				break;
			case "number":
				checks.push(`if (typeof val === "number") {
          if (!${schema}.allows(val)) return ctx.mustBe("a valid number");
        }`);
				break;
		}
	}

	return `type.unknown.narrow((val, ctx) => {
    ${checks.join("\n    ")}
    return true;
  })`;
}

function renderNullable(node: NullableNode): string {
	const inner = render(node.inner);
	return `${inner}.or(type.null)`;
}
