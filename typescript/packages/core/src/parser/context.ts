/**
 * Parser context for tracking state during parsing
 */

import type { JSONSchema } from "../schema/json-schema.js";

export interface ParseContext {
	/** Root schema - kept for potential future use */
	rootSchema: JSONSchema;
}

export function createContext(rootSchema: JSONSchema): ParseContext {
	return {
		rootSchema,
	};
}
