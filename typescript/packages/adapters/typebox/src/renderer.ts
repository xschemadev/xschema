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
import { escapeString } from "@xschemadev/core";

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
	// TODO: implement string constraints (minLength, maxLength, pattern, format)
	return "Type.String()";
}

function renderNumber(node: NumberNode): string {
	// TODO: implement number constraints (minimum, maximum, exclusiveMinimum, exclusiveMaximum, multipleOf, integer)
	return "Type.Number()";
}

function renderObject(node: ObjectNode): string {
	// TODO: implement object properties, required, additionalProperties, patternProperties, etc.
	return "Type.Object({})";
}

function renderArray(node: ArrayNode): string {
	// TODO: implement array items and constraints (minItems, maxItems, uniqueItems, contains)
	const itemSchema = render(node.items);
	return `Type.Array(${itemSchema})`;
}

function renderTuple(node: TupleNode): string {
	// TODO: implement tuple prefixItems and restItems
	const tupleSchemas = node.prefixItems.map((item) => render(item));
	return `Type.Tuple([${tupleSchemas.join(", ")}])`;
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
	// TODO: implement oneOf (exactly one match validation)
	if (node.schemas.length === 0) return "Type.Never()";
	if (node.schemas.length === 1) return render(node.schemas[0]!);

	const filtered = node.schemas.filter((s) => s.kind !== "never");
	if (filtered.length === 0) return "Type.Never()";
	if (filtered.length === 1) return render(filtered[0]!);

	const schemas = filtered.map((s) => render(s));
	return `Type.Union([${schemas.join(", ")}])`;
}

function renderNot(node: NotNode): string {
	// TODO: implement not (negation)
	const schema = render(node.schema);
	return `Type.Not(${schema})`;
}

function renderLiteral(node: LiteralNode): string {
	// TODO: handle complex values (objects/arrays)
	return `Type.Literal(${JSON.stringify(node.value)})`;
}

function renderEnum(node: EnumNode): string {
	if (node.values.length === 0) return "Type.Never()";
	if (node.values.length === 1) {
		return renderLiteral({ kind: "literal", value: node.values[0] });
	}

	// All primitives - use union of literals
	const literals = node.values.map((v) => `Type.Literal(${JSON.stringify(v)})`);
	return `Type.Union([${literals.join(", ")}])`;
}

function renderRef(node: RefNode): string {
	return render(node.resolved);
}

function renderConditional(node: ConditionalNode): string {
	// TODO: implement if/then/else
	return "Type.Any()";
}

function renderTypeGuarded(node: TypeGuardedNode): string {
	if (node.guards.length === 0) return "Type.Any()";

	// TypeGuarded nodes need runtime type checking
	// Use Type.Any() with transform/refinement
	const schemas = node.guards.map((g) => render(g.schema));
	if (schemas.length === 1) return schemas[0]!;

	return `Type.Union([${schemas.join(", ")}])`;
}

function renderNullable(node: NullableNode): string {
	const inner = render(node.inner);
	return `Type.Union([${inner}, Type.Null()])`;
}
