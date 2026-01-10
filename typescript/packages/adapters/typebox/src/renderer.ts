/**
 * TypeBox Renderer
 * Converts SchemaNode IR to TypeBox code strings
 *
 * TypeBox uses Type.* constructors that generate JSON Schema objects
 * with TypeScript inference via Static<typeof schema>
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
} from "@xschemadev/core";
import { escapeString, isPrimitive, sortedStringify } from "@xschemadev/core";

/**
 * Render a SchemaNode to TypeBox code
 */
export function render(node: SchemaNode): string {
	switch (node.kind) {
		case "string":
			return renderString(node);
		case "number":
			return renderNumber(node);
		case "boolean":
			return "Type.Boolean()";
		case "null":
			return "Type.Null()";
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
			return "Type.Any()";
		case "never":
			return "Type.Never()";
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
	const options: string[] = [];

	// Format - TypeBox supports format directly in options
	if (node.format) {
		switch (node.format) {
			case "email":
			case "uri":
			case "uri-reference":
			case "uuid":
			case "date-time":
			case "date":
			case "time":
			case "ipv4":
			case "ipv6":
			case "hostname":
			case "idn-hostname":
				options.push(`format: ${escapeString(node.format)}`);
				break;
			// Unknown formats ignored per JSON Schema spec
		}
	}

	// Constraints
	const { minLength, maxLength, pattern } = node.constraints;

	if (minLength !== undefined) {
		options.push(`minLength: ${minLength}`);
	}
	if (maxLength !== undefined) {
		options.push(`maxLength: ${maxLength}`);
	}
	if (pattern) {
		options.push(`pattern: ${escapeString(pattern)}`);
	}

	if (options.length === 0) {
		return "Type.String()";
	}

	return `Type.String({ ${options.join(", ")} })`;
}

function renderNumber(node: NumberNode): string {
	const options: string[] = [];
	const {
		minimum,
		maximum,
		exclusiveMinimum,
		exclusiveMaximum,
		multipleOf,
	} = node.constraints;

	// Range constraints - TypeBox supports these directly in options
	if (minimum !== undefined) {
		options.push(`minimum: ${minimum}`);
	}
	if (exclusiveMinimum !== undefined) {
		options.push(`exclusiveMinimum: ${exclusiveMinimum}`);
	}
	if (maximum !== undefined) {
		options.push(`maximum: ${maximum}`);
	}
	if (exclusiveMaximum !== undefined) {
		options.push(`exclusiveMaximum: ${exclusiveMaximum}`);
	}

	// Multiple of
	if (multipleOf !== undefined) {
		options.push(`multipleOf: ${multipleOf}`);
	}

	// Use Type.Integer() for integer constraint, Type.Number() otherwise
	const baseType = node.integer ? "Type.Integer" : "Type.Number";

	if (options.length === 0) {
		return `${baseType}()`;
	}

	return `${baseType}({ ${options.join(", ")} })`;
}

function renderObject(node: ObjectNode): string {
	const propKeys = Array.from(node.properties.keys());
	const options: string[] = [];

	// If using record-style (no properties, just additionalProperties schema)
	if (
		propKeys.length === 0 &&
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const valueSchema = render(node.additionalProperties);
		// Add constraints to record options
		const recordOpts = collectObjectOptions(node);
		if (recordOpts.length > 0) {
			return `Type.Record(Type.String(), ${valueSchema}, { ${recordOpts.join(", ")} })`;
		}
		return `Type.Record(Type.String(), ${valueSchema})`;
	}

	// Build shape
	const shape = propKeys.map((key) => {
		const prop = node.properties.get(key)!;
		let propCode = render(prop.schema as SchemaNode);
		if (!prop.required) {
			propCode = `Type.Optional(${propCode})`;
		}
		return `[${escapeString(key)}]: ${propCode}`;
	});

	// Collect all options - TypeBox supports these natively
	if (node.additionalProperties === false) {
		options.push("additionalProperties: false");
	} else if (
		typeof node.additionalProperties === "object" &&
		node.additionalProperties.kind !== "any"
	) {
		const additionalSchema = render(node.additionalProperties);
		options.push(`additionalProperties: ${additionalSchema}`);
	}

	// minProperties/maxProperties - native support
	if (node.minProperties !== undefined) {
		options.push(`minProperties: ${node.minProperties}`);
	}
	if (node.maxProperties !== undefined) {
		options.push(`maxProperties: ${node.maxProperties}`);
	}

	// patternProperties - native support
	if (node.patternProperties.length > 0) {
		const patterns = node.patternProperties.map((p) => {
			const patternSchema = render(p.schema as SchemaNode);
			return `[${escapeString(p.pattern)}]: ${patternSchema}`;
		});
		options.push(`patternProperties: { ${patterns.join(", ")} }`);
	}

	// propertyNames - native support
	if (node.propertyNames) {
		const keySchema = render(node.propertyNames);
		options.push(`propertyNames: ${keySchema}`);
	}

	// dependencies - native support (use dependentRequired/dependentSchemas for modern JSON Schema)
	if (node.dependencies.size > 0) {
		const depRequired: string[] = [];
		const depSchemas: string[] = [];

		for (const [prop, dep] of node.dependencies) {
			if (dep.kind === "property" && dep.requiredProperties.length > 0) {
				depRequired.push(`[${escapeString(prop)}]: ${JSON.stringify(dep.requiredProperties)}`);
			} else if (dep.kind === "schema") {
				const depCode = render(dep.schema as SchemaNode);
				depSchemas.push(`[${escapeString(prop)}]: ${depCode}`);
			}
		}

		if (depRequired.length > 0) {
			options.push(`dependentRequired: { ${depRequired.join(", ")} }`);
		}
		if (depSchemas.length > 0) {
			options.push(`dependentSchemas: { ${depSchemas.join(", ")} }`);
		}
	}

	// Build result
	const shapeCode = shape.length > 0 ? `{ ${shape.join(", ")} }` : "{}";
	const optsCode = options.length > 0 ? `, { ${options.join(", ")} }` : "";

	return `Type.Object(${shapeCode}${optsCode})`;
}

function collectObjectOptions(node: ObjectNode): string[] {
	const options: string[] = [];

	if (node.minProperties !== undefined) {
		options.push(`minProperties: ${node.minProperties}`);
	}
	if (node.maxProperties !== undefined) {
		options.push(`maxProperties: ${node.maxProperties}`);
	}

	return options;
}

function renderArray(node: ArrayNode): string {
	const itemSchema = render(node.items);
	const options = collectArrayOptions(node.constraints);

	if (options.length > 0) {
		return `Type.Array(${itemSchema}, { ${options.join(", ")} })`;
	}
	return `Type.Array(${itemSchema})`;
}

function collectArrayOptions(constraints: ArrayNode["constraints"]): string[] {
	const options: string[] = [];

	// TypeBox supports these natively
	if (constraints.minItems !== undefined) {
		options.push(`minItems: ${constraints.minItems}`);
	}
	if (constraints.maxItems !== undefined) {
		options.push(`maxItems: ${constraints.maxItems}`);
	}
	if (constraints.uniqueItems) {
		options.push("uniqueItems: true");
	}

	// contains - TypeBox supports this natively
	if (constraints.contains) {
		const containsSchema = render(constraints.contains.schema as SchemaNode);
		options.push(`contains: ${containsSchema}`);

		if (constraints.contains.minContains !== 1) {
			options.push(`minContains: ${constraints.contains.minContains}`);
		}
		if (constraints.contains.maxContains !== undefined) {
			options.push(`maxContains: ${constraints.contains.maxContains}`);
		}
	}

	return options;
}

function renderTuple(node: TupleNode): string {
	const tupleSchemas = node.prefixItems.map((item) => render(item));

	// Collect array constraints
	const arrayOpts = collectArrayOptions(node.constraints);

	// Always use Type.Unsafe with prefixItems for draft-2020-12 compatibility
	// TypeBox's Type.Tuple generates draft-07 style items:[] which Ajv2020 rejects
	const opts: string[] = [];

	// prefixItems for positional schemas
	if (tupleSchemas.length > 0) {
		opts.push(`prefixItems: [${tupleSchemas.join(", ")}]`);
	}

	if (node.restItems === false) {
		// No additional items allowed
		opts.push(`items: false`);
	} else {
		// Rest items schema
		const restSchema = render(node.restItems);
		opts.push(`items: ${restSchema}`);
	}

	// Add array constraints
	opts.push(...arrayOpts);

	return `Type.Unsafe<unknown[]>({ type: "array", ${opts.join(", ")} })`;
}

function renderUnion(node: UnionNode): string {
	if (node.variants.length === 0) return "Type.Never()";

	const filtered = node.variants.filter((v) => v.kind !== "never");
	if (filtered.length === 0) return "Type.Never()";
	if (filtered.length === 1) return render(filtered[0]!);

	const schemas = filtered.map((v) => render(v));
	return `Type.Union([${schemas.join(", ")}])`;
}

function renderIntersection(node: IntersectionNode): string {
	if (node.schemas.length === 0) return "Type.Any()";

	if (node.schemas.some((s) => s.kind === "never")) {
		return "Type.Never()";
	}

	const filtered = node.schemas.filter((s) => s.kind !== "any");
	if (filtered.length === 0) return "Type.Any()";
	if (filtered.length === 1) return render(filtered[0]!);

	const schemas = filtered.map((s) => render(s));
	return `Type.Intersect([${schemas.join(", ")}])`;
}

function renderOneOf(node: OneOfNode): string {
	if (node.schemas.length === 0) return "Type.Never()";
	if (node.schemas.length === 1) return render(node.schemas[0]!);

	const filtered = node.schemas.filter((s) => s.kind !== "never");
	if (filtered.length === 0) return "Type.Never()";
	if (filtered.length === 1) return render(filtered[0]!);

	// Count how many "any" schemas - if > 1, always matches multiple
	const anyCount = filtered.filter((s) => s.kind === "any").length;
	if (anyCount > 1) {
		return `Type.Never()`;
	}

	// TypeBox supports oneOf natively via Type.Unsafe
	const schemas = filtered.map((s) => render(s));
	return `Type.Unsafe<unknown>({ oneOf: [${schemas.join(", ")}] })`;
}

function renderNot(node: NotNode): string {
	const schema = render(node.schema);
	// TypeBox has native Type.Not support
	return `Type.Not(${schema})`;
}

function renderLiteral(node: LiteralNode): string {
	// null needs special handling - Type.Literal(null) incorrectly adds type:"object"
	if (node.value === null) {
		return `Type.Unsafe<null>({ const: null })`;
	}

	// TypeBox Type.Literal only supports primitives (string, number, boolean)
	if (isPrimitive(node.value)) {
		return `Type.Literal(${JSON.stringify(node.value)})`;
	}

	// For complex values (objects/arrays), use Type.Unsafe with const
	const sorted = sortedStringify(node.value);
	return `Type.Unsafe<unknown>({ const: ${sorted} })`;
}

function renderEnum(node: EnumNode): string {
	if (node.values.length === 0) return "Type.Never()";
	if (node.values.length === 1) {
		return renderLiteral({ kind: "literal", value: node.values[0] });
	}

	// Check if any value is null or complex (non-primitive)
	// Type.Literal(null) has a bug (adds type:"object"), so use Type.Unsafe for enums with null
	const hasNull = node.values.some((v) => v === null);
	const hasComplex = node.values.some((v) => !isPrimitive(v));

	if (hasNull || hasComplex) {
		// Use Type.Unsafe with enum keyword for null or complex values
		const enumValues = node.values.map((v) => {
			if (isPrimitive(v)) {
				return JSON.stringify(v);
			}
			return sortedStringify(v);
		});
		return `Type.Unsafe<unknown>({ enum: [${enumValues.join(", ")}] })`;
	}

	// All primitives without null - use union of literals
	const literals = node.values.map((v) => `Type.Literal(${JSON.stringify(v)})`);
	return `Type.Union([${literals.join(", ")}])`;
}

function renderRef(node: RefNode): string {
	return render(node.resolved);
}

function renderConditional(node: ConditionalNode): string {
	const ifSchema = render(node.if);
	const thenSchema = node.then ? render(node.then) : null;
	const elseSchema = node.else ? render(node.else) : null;

	// TypeBox supports if/then/else natively via Type.Unsafe
	const parts: string[] = [`if: ${ifSchema}`];

	if (thenSchema) {
		parts.push(`then: ${thenSchema}`);
	}
	if (elseSchema) {
		parts.push(`else: ${elseSchema}`);
	}

	if (!thenSchema && !elseSchema) {
		// if without then/else has no effect
		return "Type.Any()";
	}

	return `Type.Unsafe<unknown>({ ${parts.join(", ")} })`;
}

function renderTypeGuarded(node: TypeGuardedNode): string {
	if (node.guards.length === 0) return "Type.Any()";

	// TypeGuarded nodes only apply validation when the value matches the type
	// This handles JSON Schema without explicit "type" - e.g., { minimum: 1 } ignores non-numbers
	// Build a conditional schema using if/then for each guard
	const conditions: string[] = [];

	for (const guard of node.guards) {
		const schema = render(guard.schema);
		let typeSchema: string;

		switch (guard.check) {
			case "object":
				typeSchema = "Type.Object({})";
				break;
			case "array":
				typeSchema = "Type.Array(Type.Any())";
				break;
			case "string":
				typeSchema = "Type.String()";
				break;
			case "number":
				typeSchema = "Type.Number()";
				break;
		}

		// For each guard, create: if value is type, then validate against schema
		conditions.push(`{ if: ${typeSchema}, then: ${schema} }`);
	}

	if (conditions.length === 1) {
		// Single guard - extract if/then
		const guard = node.guards[0]!;
		const schema = render(guard.schema);
		let typeSchema: string;

		switch (guard.check) {
			case "object":
				typeSchema = "Type.Object({})";
				break;
			case "array":
				typeSchema = "Type.Array(Type.Any())";
				break;
			case "string":
				typeSchema = "Type.String()";
				break;
			case "number":
				typeSchema = "Type.Number()";
				break;
		}

		return `Type.Unsafe<unknown>({ if: ${typeSchema}, then: ${schema} })`;
	}

	// Multiple guards - use allOf of conditionals
	return `Type.Unsafe<unknown>({ allOf: [${conditions.join(", ")}] })`;
}

function renderNullable(node: NullableNode): string {
	const inner = render(node.inner);
	return `Type.Union([${inner}, Type.Null()])`;
}
