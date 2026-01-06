/**
 * JSON Schema types and utilities
 */

export type { JSONSchema, JSONSchemaVersion } from "./json-schema.js";
export {
	detectVersion,
	supportsRefSiblings,
	usesBooleanExclusive,
} from "./version.js";
export { normalizeSchema, normalizeDeep } from "./normalizer.js";
