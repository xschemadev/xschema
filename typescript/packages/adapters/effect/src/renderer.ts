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
