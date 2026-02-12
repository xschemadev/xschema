/**
 * Parser context for tracking state during parsing
 */

import type { JSONSchema } from "../schema/json-schema.js";

export interface ParseContext {
	/** Root schema for $ref resolution */
	rootSchema: JSONSchema;
	/** Tracks $ref paths currently being resolved (cycle detection) */
	resolving: Set<string>;
}

export function createContext(rootSchema: JSONSchema): ParseContext {
	return {
		rootSchema,
		resolving: new Set(),
	};
}
