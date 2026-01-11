/**
 * Parser context for tracking state during parsing
 */

import type { SchemaNode } from "../ir/nodes.js";
import type { JSONSchema, JSONSchemaVersion } from "../schema/json-schema.js";

export interface ParseIssue {
	code: "ref_resolution_failed";
	message: string;
	ref: string;
}

export interface ParseContext {
	/** Detected JSON Schema version */
	version: JSONSchemaVersion;

	/** Root schema (for $ref resolution) */
	rootSchema: JSONSchema;

	/** Cache of resolved refs */
	refs: Map<string, SchemaNode>;

	/** Currently processing refs (for circular detection) */
	processing: Set<string>;

	/** Non-fatal parsing issues */
	issues: ParseIssue[];

	addIssue: (issue: ParseIssue) => void;
}

export function createContext(
	rootSchema: JSONSchema,
	version: JSONSchemaVersion,
): ParseContext {
	const issues: ParseIssue[] = [];

	return {
		version,
		rootSchema,
		refs: new Map(),
		processing: new Set(),
		issues,
		addIssue: (issue) => {
			issues.push(issue);
		},
	};
}
