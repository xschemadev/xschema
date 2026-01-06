/**
 * Constraint types for SchemaNode IR
 * These represent the validation constraints extracted from JSON Schema
 */

/** String validation constraints */
export interface StringConstraints {
	minLength?: number;
	maxLength?: number;
	pattern?: string;
}

/** Number validation constraints */
export interface NumberConstraints {
	minimum?: number;
	maximum?: number;
	exclusiveMinimum?: number;
	exclusiveMaximum?: number;
	multipleOf?: number;
}

/** Array validation constraints */
export interface ArrayConstraints {
	minItems?: number;
	maxItems?: number;
	uniqueItems?: boolean;
	contains?: ContainsConstraint;
}

/** Contains constraint for arrays */
export interface ContainsConstraint {
	schema: unknown; // Will be SchemaNode, forward reference
	minContains: number;
	maxContains?: number;
}

/** Property definition in an object */
export interface PropertyDef {
	schema: unknown; // Will be SchemaNode, forward reference
	required: boolean;
}

/** Pattern property definition */
export interface PatternPropertyDef {
	pattern: string;
	schema: unknown; // Will be SchemaNode, forward reference
}

/** Dependency types */
export interface PropertyDependency {
	kind: "property";
	requiredProperties: string[];
}

export interface SchemaDependency {
	kind: "schema";
	schema: unknown; // Will be SchemaNode, forward reference
}

export type Dependency = PropertyDependency | SchemaDependency;
