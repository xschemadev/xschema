/**
 * Utility functions for adapters
 */

export {
	isPrimitive,
	escapeString,
	isEmptyObject,
	getOwnProperty,
	PROTOTYPE_PROPERTY_NAMES,
	hasPrototypeProperties,
} from "./primitives.js";

export { resolveJsonPointer, getRefName } from "./json-pointer.js";

export {
	chain,
	buildUnion,
	buildIntersection,
	buildSuperRefine,
	buildRefine,
	buildLiteral,
	buildSafeParseCheck,
	buildPropertyCheck,
	sortedStringify,
} from "./code-builder.js";
