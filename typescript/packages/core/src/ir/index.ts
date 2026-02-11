/**
 * Intermediate Representation (IR) for JSON Schema
 *
 * This module exports the SchemaNode discriminated union and related types.
 * Adapters use these types to implement their renderers.
 */

export type {
	// Main node type
	SchemaNode,
	NodeKind,
	NodeOfKind,

	// Individual node types
	StringNode,
	NumberNode,
	BooleanNode,
	NullNode,
	ObjectNode,
	ArrayNode,
	TupleNode,
	UnionNode,
	IntersectionNode,
	OneOfNode,
	NotNode,
	LiteralNode,
	EnumNode,
	AnyNode,
	NeverNode,
	ConditionalNode,
	TypeGuardedNode,
	NullableNode,
	TypeGuard,
	TypeCheck,

	// Constraint types (re-exported from nodes.ts)
	StringConstraints,
	NumberConstraints,
	ArrayConstraints,
	ContainsConstraint,
	PropertyDef,
	PatternPropertyDef,
	Dependency,
} from "./nodes.js";

export type { PropertyDependency, SchemaDependency } from "./constraints.js";
