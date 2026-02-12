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

export { getRefName } from "./json-pointer.js";

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
	DEEP_SORTED_STRINGIFY_RUNTIME,
} from "./code-builder.js";
