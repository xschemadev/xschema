/**
 * Parser context for tracking state during parsing
 */

import type { JSONSchema, JSONSchemaVersion } from "../schema/json-schema.js";
import type { SchemaNode } from "../ir/nodes.js";

export interface ParseContext {
	/** Detected JSON Schema version */
	version: JSONSchemaVersion;

	/** Root schema (for $ref resolution) */
	rootSchema: JSONSchema;

	/** Cache of resolved refs */
	refs: Map<string, SchemaNode>;

	/** Currently processing refs (for circular detection) */
	processing: Set<string>;
}

export function createContext(
	rootSchema: JSONSchema,
	version: JSONSchemaVersion,
): ParseContext {
	return {
		version,
		rootSchema,
		refs: new Map(),
		processing: new Set(),
	};
}
