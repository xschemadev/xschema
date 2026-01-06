/**
 * Zod Renderer
 * Converts SchemaNode IR to Zod code strings
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
import {
	escapeString,
	isPrimitive,
	sortedStringify,
	hasPrototypeProperties,
	buildIntersection,
} from "@xschemadev/core";

/**
 * Render a SchemaNode to Zod code
 */
export function render(node: SchemaNode): string {
	switch (node.kind) {
		case "string":
			return renderString(node);
		case "number":
			return renderNumber(node);
		case "boolean":
			return "z.boolean()";
		case "null":
			return "z.null()";
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
			return "z.any()";
		case "never":
			return "z.never()";
		case "ref":
			return renderRef(node);
		case "conditional":
			return renderConditional(node);
		case "typeGuarded":
			return renderTypeGuarded(node);
		case "nullable":
			return renderNullable(node);
	}
}

function renderString(node: StringNode): string {
	let result = "z.string()";

	// Format
	if (node.format) {
		switch (node.format) {
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
				result += `.regex(/^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/)`;
				break;
		}
	}

	// Constraints - use grapheme cluster counting per JSON Schema spec
	const { minLength, maxLength, pattern } = node.constraints;

	if (minLength !== undefined) {
		result += `.refine((val) => [...new Intl.Segmenter().segment(val)].length >= ${minLength}, { message: "String must have at least ${minLength} character(s)" })`;
	}
	if (maxLength !== undefined) {
		result += `.refine((val) => [...new Intl.Segmenter().segment(val)].length <= ${maxLength}, { message: "String must have at most ${maxLength} character(s)" })`;
	}
	if (pattern) {
		result += `.regex(new RegExp(${escapeString(pattern)}))`;
	}

	return result;
}

function renderNumber(node: NumberNode): string {
	let result = node.integer ? "z.number().int()" : "z.number()";
	const {
		minimum,
		maximum,
		exclusiveMinimum,
		exclusiveMaximum,
		multipleOf,
	} = node.constraints;

	if (minimum !== undefined) {
		result += `.gte(${minimum})`;
	}
	if (exclusiveMinimum !== undefined) {
		result += `.gt(${exclusiveMinimum})`;
	}
	if (maximum !== undefined) {
		result += `.lte(${maximum})`;
	}
	if (exclusiveMaximum !== undefined) {
		result += `.lt(${exclusiveMaximum})`;
	}

	if (multipleOf !== undefined) {
		// For small values, use epsilon-based comparison for float precision
		if (multipleOf < 1 && multipleOf > 0) {
			result += `.refine((val) => Math.abs(val - Math.round(val / ${multipleOf}) * ${multipleOf}) < 1e-10, { message: "Number must be a multiple of ${multipleOf}" })`;
		} else {
			result += `.multipleOf(${multipleOf})`;
		}
	}

	return result;
}

function renderObject(node: ObjectNode): string {
	const propKeys = Array.from(node.properties.keys());
	const requiredKeys = propKeys.filter(
		(k) => node.properties.get(k)!.required,
	);

	// Check for prototype properties that need special handling
	const hasProtoProps =
		hasPrototypeProperties(propKeys) || hasPrototypeProperties(requiredKeys);

	if (hasProtoProps) {
		return renderObjectWithProtoProps(node);
	}

	const hasPatternProps = node.patternProperties.length > 0;
	const needsPassthrough = hasPatternProps || node.propertyNames !== undefined;

	let result: string;

	if (propKeys.length === 0 && !needsPassthrough) {
		// No properties
		if (node.additionalProperties === false) {
			result = "z.object({}).strict()";
		} else if (
			typeof node.additionalProperties === "object" &&
			node.additionalProperties.kind !== "any"
		) {
			const valueSchema = render(node.additionalProperties);
			result = `z.record(z.string(), ${valueSchema})`;
		} else {
			result = "z.object({}).passthrough()";
		}
	} else {
		// Build shape
		const shape = propKeys.map((key) => {
			const prop = node.properties.get(key)!;
			let propCode = render(prop.schema as SchemaNode);
			if (!prop.required) {
				propCode += ".optional()";
			}
			return `[${escapeString(key)}]: ${propCode}`;
		});

		result =
			shape.length > 0
				? `z.object({ ${shape.join(", ")} })`
				: "z.object({})";

		// Handle additional properties mode
		if (needsPassthrough) {
			result += ".passthrough()";
		} else if (node.additionalProperties === false) {
			result += ".strict()";
		} else if (
			typeof node.additionalProperties === "object" &&
			node.additionalProperties.kind !== "any"
		) {
			const valueSchema = render(node.additionalProperties);
			result += `.catchall(${valueSchema})`;
		} else {
			result += ".passthrough()";
		}
	}

	// Pattern properties validation
	if (hasPatternProps || needsPassthrough) {
		result += renderPatternPropsValidation(node, propKeys);
	}

	// Property names validation
	if (node.propertyNames) {
		const keySchema = render(node.propertyNames);
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

	// Required properties validation (for z.any() props)
	if (requiredKeys.length > 0) {
		const requiredJson = JSON.stringify(requiredKeys);
		result += `.superRefine((val, ctx) => {
      const required = ${requiredJson};
      for (const key of required) {
        if (!Object.hasOwn(val, key)) {
          ctx.addIssue({ code: z.ZodIssueCode.custom, path: [key], message: "Required" });
        }
      }
    })`;
	}

	// Min/max properties
	if (node.minProperties !== undefined) {
		result += `.refine((val) => Object.keys(val).length >= ${node.minProperties}, { message: "Object must have at least ${node.minProperties} properties" })`;
	}
	if (node.maxProperties !== undefined) {
		result += `.refine((val) => Object.keys(val).length <= ${node.maxProperties}, { message: "Object must have at most ${node.maxProperties} properties" })`;
	}

	// Dependencies
	result += renderDependencies(node);

	return result;
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
          const result = ${propCode}.safeParse(val[${keyExpr}]);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue({ ...issue, path: [${keyExpr}, ...issue.path] }));
          }
        } else {
          ctx.addIssue({ code: z.ZodIssueCode.custom, path: [${keyExpr}], message: "Required" });
        }`);
		} else {
			validators.push(`
        if (Object.hasOwn(val, ${keyExpr})) {
          const result = ${propCode}.safeParse(val[${keyExpr}]);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue({ ...issue, path: [${keyExpr}, ...issue.path] }));
          }
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
          if (!definedProps.has(key)) {
            ctx.addIssue({ code: z.ZodIssueCode.custom, path: [key], message: "Additional property not allowed" });
          }
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
            const result = ${additionalSchema}.safeParse(value);
            if (!result.success) {
              result.error.issues.forEach(issue => ctx.addIssue({ ...issue, path: [key, ...issue.path] }));
            }
          }
        }`;
	}

	return `z.any().superRefine((val, ctx) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return;${validators.join("")}${additionalCheck}
    })`;
}

function renderPatternPropsValidation(
	node: ObjectNode,
	propKeys: string[],
): string {
	const patterns = node.patternProperties;
	const definedProps = JSON.stringify(propKeys);

	let body = `
      const definedProps = new Set(${definedProps});
      const patterns = [${patterns.map((p) => `new RegExp(${escapeString(p.pattern)})`).join(", ")}];
    `;

	// Validate pattern properties
	patterns.forEach((p, i) => {
		const patternCode = render(p.schema as SchemaNode);
		body += `
      for (const [key, value] of Object.entries(val)) {
        if (patterns[${i}].test(key)) {
          const result = ${patternCode}.safeParse(value);
          if (!result.success) {
            ctx.addIssue({
              code: z.ZodIssueCode.custom,
              path: [key],
              message: "Property matching pattern ${p.pattern} is invalid",
            });
          }
        }
      }`;
	});

	// Additional properties validation
	if (node.additionalProperties === false) {
		body += `
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
	} else if (
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const additionalSchema = render(node.additionalProperties);
		body += `
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

	return `.superRefine((val, ctx) => {${body}})`;
}

function renderDependencies(node: ObjectNode): string {
	let result = "";

	for (const [prop, dep] of node.dependencies) {
		if (dep.kind === "property") {
			if (dep.requiredProperties.length > 0) {
				const message = escapeString(
					`Property ${prop} requires ${dep.requiredProperties.join(", ")}`,
				);
				result += `.refine((val) => {
          if (Object.hasOwn(val, ${escapeString(prop)})) {
            return ${dep.requiredProperties.map((d) => `Object.hasOwn(val, ${escapeString(d)})`).join(" && ")};
          }
          return true;
        }, { message: ${message} })`;
			}
		} else {
			const depCode = render(dep.schema as SchemaNode);
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

	return result;
}

function renderArray(node: ArrayNode): string {
	const itemSchema = render(node.items);
	let result = `z.array(${itemSchema})`;
	result += renderArrayConstraints(node.constraints);
	return result;
}

function renderTuple(node: TupleNode): string {
	const tupleSchemas = node.prefixItems.map((item) => render(item));
	const schemasArray = `[${tupleSchemas.join(", ")}]`;

	let result = `z.array(z.any()).superRefine((val, ctx) => {
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

	// Rest items
	if (node.restItems === false) {
		result += `.refine((val) => val.length <= ${node.prefixItems.length}, { message: "Array must not have more than ${node.prefixItems.length} items" })`;
	} else if (node.restItems.kind !== "any") {
		const restSchema = render(node.restItems);
		result += `.superRefine((val, ctx) => {
        const restSchema = ${restSchema};
        for (let i = ${node.prefixItems.length}; i < val.length; i++) {
          const itemResult = restSchema.safeParse(val[i]);
          if (!itemResult.success) {
            itemResult.error.issues.forEach(issue => {
              ctx.addIssue({ ...issue, path: [i, ...issue.path] });
            });
          }
        }
      })`;
	}

	result += renderArrayConstraints(node.constraints);
	return result;
}

function renderArrayConstraints(
	constraints: ArrayNode["constraints"],
): string {
	let result = "";

	if (constraints.minItems !== undefined) {
		result += `.min(${constraints.minItems})`;
	}
	if (constraints.maxItems !== undefined) {
		result += `.max(${constraints.maxItems})`;
	}

	if (constraints.uniqueItems) {
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

	if (constraints.contains) {
		const containsSchema = render(constraints.contains.schema as SchemaNode);
		const minContains = constraints.contains.minContains;
		const maxContains = constraints.contains.maxContains;

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

function renderUnion(node: UnionNode): string {
	if (node.variants.length === 0) return "z.never()";

	// Filter out "never" variants since they can't match anything
	const filtered = node.variants.filter((v) => v.kind !== "never");
	if (filtered.length === 0) return "z.never()";
	if (filtered.length === 1) return render(filtered[0]!);

	const schemas = filtered.map((v) => render(v));
	return `z.union([${schemas.join(", ")}])`;
}

function renderIntersection(node: IntersectionNode): string {
	if (node.schemas.length === 0) return "z.any()";

	// Short-circuit: if ANY schema is never, intersection is never
	if (node.schemas.some((s) => s.kind === "never")) {
		return "z.never()";
	}

	// Filter out "any" schemas since they don't constrain the intersection
	const filtered = node.schemas.filter((s) => s.kind !== "any");
	if (filtered.length === 0) return "z.any()";
	if (filtered.length === 1) return render(filtered[0]!);

	const schemas = filtered.map((s) => render(s));
	return buildIntersection(schemas);
}

function renderOneOf(node: OneOfNode): string {
	if (node.schemas.length === 0) return "z.never()";
	if (node.schemas.length === 1) return render(node.schemas[0]!);

	// Filter out "never" schemas since they can never match
	const filtered = node.schemas.filter((s) => s.kind !== "never");

	// If all schemas are never, nothing can match exactly one
	if (filtered.length === 0) return "z.never()";

	// If exactly one schema remains after filtering never, it must match
	if (filtered.length === 1) return render(filtered[0]!);

	// Count how many "any" schemas there are - if > 1, always matches multiple
	const anyCount = filtered.filter((s) => s.kind === "any").length;
	if (anyCount > 1) {
		// Multiple "any" schemas means any value matches multiple
		return `z.any().refine(() => false, { message: "oneOf has multiple 'true' schemas - impossible to match exactly one" })`;
	}

	const schemas = filtered.map((s) => render(s));
	return `z.any().superRefine((val, ctx) => {
    const schemas = [${schemas.join(", ")}];
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

function renderNot(node: NotNode): string {
	const schema = render(node.schema);
	return `z.any().refine((val) => !${schema}.safeParse(val).success, { message: "Value must not match schema" })`;
}

function renderLiteral(node: LiteralNode): string {
	if (isPrimitive(node.value)) {
		return `z.literal(${JSON.stringify(node.value)})`;
	}

	// Objects/arrays need deep equality
	const isArray = Array.isArray(node.value);
	const baseType = isArray ? "z.array(z.any())" : "z.object({}).passthrough()";
	const sorted = sortedStringify(node.value);

	return `${baseType}.refine((val) => JSON.stringify(val, Object.keys(val as object).sort()) === ${JSON.stringify(sorted)}, { message: "Value must equal the const value" })`;
}

function renderEnum(node: EnumNode): string {
	const values = node.values;

	if (values.length === 0) return "z.never()";

	if (values.length === 1) {
		return renderLiteral({ kind: "literal", value: values[0] });
	}

	// All strings - use z.enum
	if (values.every((v) => typeof v === "string")) {
		return `z.enum([${values.map((v) => JSON.stringify(v)).join(", ")}])`;
	}

	// Check for complex values
	const hasComplexValues = values.some((v) => !isPrimitive(v));

	if (hasComplexValues) {
		const sortedValues = values.map((v) =>
			JSON.stringify(
				v,
				v != null && typeof v === "object"
					? Object.keys(v).sort()
					: undefined,
			),
		);
		const valuesArrayCode = `[${sortedValues.map((v) => JSON.stringify(v)).join(", ")}]`;

		return `z.any().refine((val) => {
      const normalized = JSON.stringify(val, val != null && typeof val === 'object' ? Object.keys(val).sort() : undefined);
      const validValues = ${valuesArrayCode};
      return validValues.includes(normalized);
    }, { message: "Value must be one of the enum values" })`;
	}

	// All primitives - use union of literals
	const literals = values.map((v) => `z.literal(${JSON.stringify(v)})`);
	return `z.union([${literals.join(", ")}])`;
}

function renderRef(node: RefNode): string {
	// The resolved schema is already parsed, just render it
	return render(node.resolved);
}

function renderConditional(node: ConditionalNode): string {
	const ifSchema = render(node.if);
	const thenSchema = node.then ? render(node.then) : null;
	const elseSchema = node.else ? render(node.else) : null;

	if (thenSchema && elseSchema) {
		// If then is never, matching if means invalid
		// If else is never, not matching if means invalid
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

	// if without then/else has no effect
	return "z.any()";
}

function renderTypeGuarded(node: TypeGuardedNode): string {
	if (node.guards.length === 0) return "z.any()";

	const checks: string[] = [];

	for (const guard of node.guards) {
		const schema = render(guard.schema);
		switch (guard.check) {
			case "object":
				checks.push(`if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = ${schema}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }`);
				break;
			case "array":
				checks.push(`if (Array.isArray(val)) {
          const result = ${schema}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }`);
				break;
			case "string":
				checks.push(`if (typeof val === "string") {
          const result = ${schema}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }`);
				break;
			case "number":
				checks.push(`if (typeof val === "number") {
          const result = ${schema}.safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }`);
				break;
		}
	}

	return `z.any().superRefine((val, ctx) => {
        ${checks.join("\n        ")}
      })`;
}

function renderNullable(node: NullableNode): string {
	const inner = render(node.inner);
	return `${inner}.nullable()`;
}
