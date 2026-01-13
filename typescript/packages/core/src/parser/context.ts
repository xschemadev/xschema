/**
 * Parser context for tracking state during parsing
 */

import type { SchemaNode } from "../ir/nodes.js";
import type { JSONSchema } from "../schema/json-schema.js";

export interface ParseContext {
	/** Root schema (for $ref resolution) */
	rootSchema: JSONSchema;

	/** Cache of resolved refs */
	refs: Map<string, SchemaNode>;

	/** Currently processing refs (for circular detection) */
	processing: Set<string>;
}

export function createContext(rootSchema: JSONSchema): ParseContext {
	return {
		rootSchema,
		refs: new Map(),
		processing: new Set(),
	};
}
