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
	// TODO: Implement string constraints
	return "S.String";
}

function renderNumber(node: NumberNode): string {
	// TODO: Implement number constraints
	return "S.Number";
}

function renderObject(node: ObjectNode): string {
	// TODO: Implement object rendering
	return "S.Struct({})";
}

function renderArray(node: ArrayNode): string {
	// TODO: Implement array rendering
	const itemsSchema = render(node.items);
	return `S.Array(${itemsSchema})`;
}

function renderTuple(node: TupleNode): string {
	// TODO: Implement tuple rendering
	return "S.Tuple()";
}

function renderUnion(node: UnionNode): string {
	// TODO: Implement union rendering
	return "S.Union()";
}

function renderIntersection(node: IntersectionNode): string {
	// TODO: Implement intersection rendering
	return "S.Struct({})";
}

function renderOneOf(node: OneOfNode): string {
	// TODO: Implement oneOf rendering
	return "S.Union()";
}

function renderNot(node: NotNode): string {
	// TODO: Implement not rendering
	return "S.Unknown";
}

function renderLiteral(node: LiteralNode): string {
	// TODO: Implement literal rendering
	if (typeof node.value === "string") {
		return `S.Literal(${escapeString(node.value)})`;
	}
	return `S.Literal(${JSON.stringify(node.value)})`;
}

function renderEnum(node: EnumNode): string {
	// TODO: Implement enum rendering
	return "S.Union()";
}

function renderRef(node: RefNode): string {
	return render(node.resolved);
}

function renderConditional(node: ConditionalNode): string {
	// TODO: Implement conditional rendering
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
