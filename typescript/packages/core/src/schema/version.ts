/**
 * JSON Schema version detection
 */

import type { JSONSchema, JSONSchemaVersion } from "./json-schema.js";

/**
 * Detect JSON Schema version from $schema keyword
 */
export function detectVersion(schema: JSONSchema): JSONSchemaVersion {
	const $schema = schema.$schema;

	if (!$schema) {
		// Default to latest when not specified
		return "draft-2020-12";
	}

	// Draft 2020-12
	if (
		$schema === "https://json-schema.org/draft/2020-12/schema" ||
		$schema.includes("draft/2020-12")
	) {
		return "draft-2020-12";
	}

	// Draft 2019-09
	if (
		$schema === "https://json-schema.org/draft/2019-09/schema" ||
		$schema.includes("draft/2019-09")
	) {
		return "draft-2019-09";
	}

	// Draft 7
	if ($schema.includes("draft-07")) {
		return "draft-7";
	}

	// Draft 6
	if ($schema.includes("draft-06")) {
		return "draft-6";
	}

	// Draft 4
	if ($schema.includes("draft-04")) {
		return "draft-4";
	}

	// Draft 3
	if ($schema.includes("draft-03")) {
		return "draft-3";
	}

	// OpenAPI 3.0
	if ($schema.includes("openapi") && $schema.includes("3.0")) {
		return "openapi-3.0";
	}

	// Default to latest
	return "draft-2020-12";
}

/**
 * Check if version supports $ref siblings (draft-2019-09+)
 */
export function supportsRefSiblings(version: JSONSchemaVersion): boolean {
	return version === "draft-2020-12" || version === "draft-2019-09";
}

/**
 * Check if version uses boolean exclusiveMinimum/Maximum (draft-4)
 */
export function usesBooleanExclusive(version: JSONSchemaVersion): boolean {
	return version === "draft-4" || version === "draft-3";
}
