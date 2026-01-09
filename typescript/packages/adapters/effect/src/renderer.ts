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
			case "uri":
			case "uri-reference":
				result = "S.String.pipe(S.url())";
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

function renderObject(node: ObjectNode): string {
	const propKeys = Array.from(node.properties.keys());

	// If using record-style (no properties, just additionalProperties schema)
	if (
		propKeys.length === 0 &&
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const valueSchema = render(node.additionalProperties);
		let result = `S.Record({ key: S.String, value: ${valueSchema} })`;
		const filters = renderObjectConstraints(node);
		if (filters.length > 0) {
			result = `${result}.pipe(${filters.join(", ")})`;
		}
		return result;
	}

	// Build shape
	const shape = propKeys.map((key) => {
		const prop = node.properties.get(key)!;
		let propCode = render(prop.schema as SchemaNode);
		if (!prop.required) {
			propCode = `S.optional(${propCode})`;
		}
		return `${escapeString(key)}: ${propCode}`;
	});

	let result = "";

	// Choose object type based on additionalProperties
	if (node.additionalProperties === false) {
		// Strict mode - no additional properties allowed
		result =
			shape.length > 0
				? `S.Struct({ ${shape.join(", ")} })`
				: "S.Struct({})";
	} else if (
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		// Validate additional properties against schema using catchall
		const restSchema = render(node.additionalProperties);
		result =
			shape.length > 0
				? `S.Struct({ ${shape.join(", ")} }, S.Record({ key: S.String, value: ${restSchema} }))`
				: `S.Record({ key: S.String, value: ${restSchema} })`;
	} else {
		// Allow additional properties (loose mode)
		result =
			shape.length > 0
				? `S.Struct({ ${shape.join(", ")} })`
				: "S.Struct({})";
	}

	// Collect all validation filters
	const filters: string[] = [];

	// Handle additionalProperties: false by rejecting unknown keys
	if (node.additionalProperties === false) {
		const definedKeys = JSON.stringify(propKeys);
		const hasPatternProps = node.patternProperties.length > 0;
		
		if (hasPatternProps) {
			const patterns = node.patternProperties.map(p => `new RegExp(${escapeString(p.pattern)})`).join(", ");
			filters.push(`S.filter((val) => {
        const definedProps = new Set(${definedKeys});
        const patterns = [${patterns}];
        for (const key of Object.keys(val)) {
          if (definedProps.has(key)) continue;
          const matchesPattern = patterns.some(p => p.test(key));
          if (!matchesPattern) return false;
        }
        return true;
      }, { message: () => "Additional properties not allowed" })`);
		} else {
			filters.push(`S.filter((val) => {
        const definedProps = new Set(${definedKeys});
        for (const key of Object.keys(val)) {
          if (!definedProps.has(key)) return false;
        }
        return true;
      }, { message: () => "Additional properties not allowed" })`);
		}
	}

	// Pattern properties validation
	if (node.patternProperties.length > 0) {
		filters.push(...renderPatternPropsFilters(node));
	}

	// Property names validation
	if (node.propertyNames) {
		const keySchema = render(node.propertyNames);
		filters.push(`S.filter((val) => {
      for (const key of Object.keys(val)) {
        const result = S.decodeUnknownEither(${keySchema})(key);
        if (result._tag === "Left") return false;
      }
      return true;
    }, { message: () => "Invalid property name" })`);
	}

	// Min/max properties
	filters.push(...renderObjectConstraints(node));

	// Dependencies
	filters.push(...renderDependenciesFilters(node));

	// Apply all filters in pipe
	if (filters.length > 0) {
		result = `${result}.pipe(${filters.join(", ")})`;
	}

	return result;
}

function renderPatternPropsFilters(node: ObjectNode): string[] {
	const patterns = node.patternProperties;

	const checks: string[] = [];

	// Validate pattern properties
	patterns.forEach((p) => {
		const patternCode = render(p.schema as SchemaNode);
		const patternStr = escapeString(p.pattern);
		checks.push(`
      for (const [key, value] of Object.entries(val)) {
        if (new RegExp(${patternStr}).test(key)) {
          const result = S.decodeUnknownEither(${patternCode})(value);
          if (result._tag === "Left") return false;
        }
      }`);
	});

	return [`S.filter((val) => {${checks.join("")}
      return true;
    }, { message: () => "Pattern property validation failed" })`];
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

	let result = "";

	// Rest items handling
	if (node.restItems === false) {
		// Strict tuple - no additional items (schemas as individual arguments)
		if (tupleSchemas.length === 0) {
			result = "S.Tuple()";
		} else {
			result = `S.Tuple(${tupleSchemas.join(", ")})`;
		}
	} else if (node.restItems.kind !== "any") {
		// Tuple with rest items - first arg is array of schemas, second is rest schema
		const restSchema = render(node.restItems);
		if (tupleSchemas.length === 0) {
			result = `S.Array(${restSchema})`;
		} else {
			result = `S.Tuple([${tupleSchemas.join(", ")}], ${restSchema})`;
		}
	} else {
		// Tuple allowing any rest items
		if (tupleSchemas.length === 0) {
			result = "S.Array(S.Unknown)";
		} else {
			result = `S.Tuple([${tupleSchemas.join(", ")}], S.Unknown)`;
		}
	}

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

	// Check for complex values
	const hasComplexValues = values.some((v) => !isPrimitive(v));

	if (hasComplexValues) {
		const sortedValues = values.map((v) =>
			JSON.stringify(
				v,
				v != null && typeof v === "object"
					? Object.keys(v as object).sort()
					: undefined,
			),
		);
		return `S.Unknown.pipe(S.filter((val) => {
      const sorted = JSON.stringify(val, val != null && typeof val === "object" ? Object.keys(val as object).sort() : undefined);
      return [${sortedValues.join(", ")}].includes(sorted);
    }, { message: () => "Value must match one of the enum values" }))`;
	}

	// Mixed primitives - use union of literals
	const literals = values.map((v) => `S.Literal(${JSON.stringify(v)})`);
	return `S.Union(${literals.join(", ")})`;
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
	// For now, render as union of all guarded schemas
	if (node.guards.length === 0) return "S.Unknown";
	if (node.guards.length === 1) return render(node.guards[0].schema);
	return `S.Union(${node.guards.map((g) => render(g.schema)).join(", ")})`;
}

function renderNullable(node: NullableNode): string {
	const innerSchema = render(node.inner);
	return `S.NullOr(${innerSchema})`;
}
