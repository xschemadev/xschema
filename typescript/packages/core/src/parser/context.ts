/**
 * Parser context for tracking state during parsing
 */

import type { SchemaNode } from "../ir/nodes.js";
import type { JSONSchema, JSONSchemaVersion } from "../schema/json-schema.js";

/**
 * Vocabulary URI constants for standard JSON Schema vocabularies
 */
export const VOCAB = {
	// Draft 2020-12
	VALIDATION_2020: "https://json-schema.org/draft/2020-12/vocab/validation",
	APPLICATOR_2020: "https://json-schema.org/draft/2020-12/vocab/applicator",
	FORMAT_ANNOTATION_2020:
		"https://json-schema.org/draft/2020-12/vocab/format-annotation",
	// Draft 2019-09
	VALIDATION_2019: "https://json-schema.org/draft/2019-09/vocab/validation",
	APPLICATOR_2019: "https://json-schema.org/draft/2019-09/vocab/applicator",
	FORMAT_2019: "https://json-schema.org/draft/2019-09/vocab/format",
} as const;

/**
 * Keywords controlled by the validation vocabulary
 */
export const VALIDATION_KEYWORDS = new Set([
	// String validation
	"minLength",
	"maxLength",
	"pattern",
	// Numeric validation
	"minimum",
	"maximum",
	"exclusiveMinimum",
	"exclusiveMaximum",
	"multipleOf",
	// Array validation
	"minItems",
	"maxItems",
	"uniqueItems",
	"minContains",
	"maxContains",
	// Object validation
	"minProperties",
	"maxProperties",
	"required",
	"dependentRequired",
	// Type & enum/const
	"type",
	"enum",
	"const",
]);

/**
 * Keywords controlled by the applicator vocabulary
 */
export const APPLICATOR_KEYWORDS = new Set([
	// Object applicators
	"properties",
	"patternProperties",
	"additionalProperties",
	"propertyNames",
	"dependentSchemas",
	// Array applicators
	"items",
	"prefixItems",
	"additionalItems",
	"contains",
	// Composition
	"allOf",
	"anyOf",
	"oneOf",
	"not",
	"if",
	"then",
	"else",
]);

/**
 * Keywords controlled by the format vocabulary
 */
export const FORMAT_KEYWORDS = new Set(["format"]);

export interface ParseContext {
	/** Detected JSON Schema version */
	version: JSONSchemaVersion;

	/** Root schema (for $ref resolution) */
	rootSchema: JSONSchema;

	/** Cache of resolved refs */
	refs: Map<string, SchemaNode>;

	/** Currently processing refs (for circular detection) */
	processing: Set<string>;

	/** Vocabulary settings from metaschema (undefined = all enabled) */
	vocabulary?: Record<string, boolean>;
}

export interface ParseOptions {
	/** Vocabulary from metaschema $vocabulary (undefined = all enabled) */
	vocabulary?: Record<string, boolean>;
}

export function createContext(
	rootSchema: JSONSchema,
	version: JSONSchemaVersion,
	options?: ParseOptions,
): ParseContext {
	return {
		version,
		rootSchema,
		refs: new Map(),
		processing: new Set(),
		vocabulary: options?.vocabulary,
	};
}

/**
 * Check if a vocabulary is enabled.
 * Returns true if vocabulary is undefined (all enabled) or the vocab URI is not explicitly false.
 */
export function isVocabEnabled(
	ctx: ParseContext,
	...vocabUris: string[]
): boolean {
	if (!ctx.vocabulary) return true;
	// Check if any of the vocab URIs is enabled (present and not false)
	for (const uri of vocabUris) {
		if (ctx.vocabulary[uri] !== false) return true;
	}
	return false;
}

/**
 * Check if validation vocabulary is enabled
 */
export function isValidationEnabled(ctx: ParseContext): boolean {
	return isVocabEnabled(ctx, VOCAB.VALIDATION_2020, VOCAB.VALIDATION_2019);
}

/**
 * Check if applicator vocabulary is enabled
 */
export function isApplicatorEnabled(ctx: ParseContext): boolean {
	return isVocabEnabled(ctx, VOCAB.APPLICATOR_2020, VOCAB.APPLICATOR_2019);
}

/**
 * Check if format vocabulary is enabled
 */
export function isFormatEnabled(ctx: ParseContext): boolean {
	return isVocabEnabled(
		ctx,
		VOCAB.FORMAT_ANNOTATION_2020,
		VOCAB.FORMAT_2019,
	);
}
