/**
 * JSON Pointer utilities for $ref resolution
 */

import type { JSONSchema } from "../schema/json-schema.js";

/**
 * Resolve a JSON Pointer reference within a schema
 */
export function resolveJsonPointer(
	ref: string,
	rootSchema: JSONSchema,
): JSONSchema {
	if (!ref.startsWith("#")) {
		throw new Error(`External $ref is not supported: ${ref}`);
	}

	const path = ref.slice(1).split("/").filter(Boolean);

	// Handle root reference "#"
	if (path.length === 0) {
		return rootSchema;
	}

	// Traverse the full JSON pointer path from root
	let current: unknown = rootSchema;
	for (const segment of path) {
		// Decode JSON Pointer escape sequences
		const decoded = segment.replace(/~1/g, "/").replace(/~0/g, "~");

		if (current && typeof current === "object" && decoded in current) {
			current = (current as Record<string, unknown>)[decoded];
		} else {
			throw new Error(`Reference not found: ${ref}`);
		}
	}

	return current as JSONSchema;
}

/**
 * Extract definition name from a $ref path
 * e.g., "#/$defs/User" -> "User"
 */
export function getRefName(ref: string): string | null {
	const match = ref.match(/#\/(?:\$defs|definitions)\/(.+)$/);
	return match ? match[1] : null;
}
