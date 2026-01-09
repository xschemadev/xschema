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
import { escapeString } from "@xschemadev/core";

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
