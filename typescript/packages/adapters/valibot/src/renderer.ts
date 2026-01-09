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
	RefNode,
	ConditionalNode,
	TypeGuardedNode,
	NullableNode,
	PropertyDef,
} from "@xschemadev/core";

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

function renderString(_node: StringNode): string {
	return "v.string()";
}

function renderNumber(_node: NumberNode): string {
	return "v.number()";
}

function renderObject(_node: ObjectNode): string {
	return "v.object({})";
}

function renderArray(_node: ArrayNode): string {
	return "v.array(v.any())";
}

function renderTuple(_node: TupleNode): string {
	return "v.tuple([])";
}

function renderUnion(_node: UnionNode): string {
	return "v.union([v.any()])";
}

function renderIntersection(_node: IntersectionNode): string {
	return "v.intersect([v.any()])";
}

function renderOneOf(_node: OneOfNode): string {
	return "v.union([v.any()])";
}

function renderNot(_node: NotNode): string {
	return "v.never()";
}

function renderLiteral(_node: LiteralNode): string {
	return "v.literal(null)";
}

function renderEnum(_node: EnumNode): string {
	return "v.union([v.literal(null)])";
}

function renderRef(node: RefNode): string {
	return render(node.resolved);
}

function renderConditional(_node: ConditionalNode): string {
	return "v.any()";
}

function renderTypeGuarded(node: TypeGuardedNode): string {
	// For now, render as union of all guarded schemas
	if (node.guards.length === 0) return "v.any()";
	if (node.guards.length === 1) return render(node.guards[0].schema);
	return `v.union([${node.guards.map((g) => render(g.schema)).join(", ")}])`;
}

function renderNullable(node: NullableNode): string {
	return `v.nullable(${render(node.inner)})`;
}
