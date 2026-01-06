#!/usr/bin/env bun
/**
 * Schema Generation Script
 *
 * Discovers languages and adapters from the repository structure and generates
 * JSON Schema files that IDEs can use for config file validation.
 *
 * Languages are directories at the repo root that contain `packages/adapters/`.
 * Adapters are subdirectories of `{lang}/packages/adapters/`.
 */

import {
	existsSync,
	mkdirSync,
	readdirSync,
	statSync,
	writeFileSync,
} from "node:fs";
import { dirname, join } from "node:path";

const REPO_ROOT = join(import.meta.dirname, "..", "..");
const PUBLIC_DIR = join(import.meta.dirname, "..", "public");
const BASE_URL = "https://xschema.dev/schemas";

export interface LanguageInfo {
	name: string; // directory name, e.g., "ts"
	adapters: string[]; // adapter names, e.g., ["zod"]
}

/**
 * Discover adapters by listing subdirectories
 */
export function discoverAdapters(adaptersPath: string): string[] {
	const adapters: string[] = [];

	const entries = readdirSync(adaptersPath, { withFileTypes: true });

	for (const entry of entries) {
		if (!entry.isDirectory()) continue;
		if (entry.name.startsWith(".")) continue;

		adapters.push(entry.name);
	}

	return adapters;
}

/**
 * Discover languages by finding directories with packages/adapters/ subdirectory
 */
export function discoverLanguages(
	repoRoot: string = REPO_ROOT,
): LanguageInfo[] {
	const languages: LanguageInfo[] = [];

	const entries = readdirSync(repoRoot, { withFileTypes: true });

	for (const entry of entries) {
		if (!entry.isDirectory()) continue;
		if (entry.name.startsWith(".") || entry.name === "node_modules") continue;

		const adaptersPath = join(repoRoot, entry.name, "packages", "adapters");

		if (existsSync(adaptersPath) && statSync(adaptersPath).isDirectory()) {
			const adapters = discoverAdapters(adaptersPath);
			if (adapters.length > 0) {
				languages.push({
					name: entry.name,
					adapters,
				});
			}
		}
	}

	return languages;
}

/**
 * Get display name for a language
 */
export function getLanguageDisplayName(lang: string): string {
	const names: Record<string, string> = {
		ts: "TypeScript",
		py: "Python",
	};
	return names[lang] || lang;
}

// Needed to not mess up with TS compiler
// JSON Schema keyword that conflicts with TS/biome
const THEN = "then";

/**
 * Create a conditional schema rule for JSON Schema if/then
 */
function createConditional(
	condition: object,
	result: object,
): Record<string, unknown> {
	const obj: Record<string, unknown> = { if: condition };
	obj[THEN] = result;
	return obj;
}

/**
 * Generate JSON Schema for a language
 */
export function generateSchema(
	lang: LanguageInfo,
	baseUrl: string = BASE_URL,
): Record<string, unknown> {
	const displayName = getLanguageDisplayName(lang.name);

	// Build conditional validation rules for source field
	const urlFileCondition = createConditional(
		{
			properties: {
				sourceType: { enum: ["url", "file"] },
			},
		},
		{
			properties: {
				source: {
					type: "string",
					description:
						"URL or file path to the JSON Schema. File paths are relative to this config file.",
				},
			},
		},
	);

	const jsonCondition = createConditional(
		{
			properties: {
				sourceType: { const: "json" },
			},
		},
		{
			properties: {
				source: {
					type: "object",
					description: "Inline JSON Schema object.",
				},
			},
		},
	);

	return {
		$schema: "https://json-schema.org/draft/2020-12/schema",
		$id: `${baseUrl}/v1/${lang.name}.jsonc`,
		title: `XSchema ${displayName} Configuration`,
		description: `Configuration file for XSchema ${displayName} projects. Defines schemas to be converted to native validation libraries.`,
		type: "object",
		properties: {
			$schema: {
				type: "string",
				const: `${baseUrl}/${lang.name}.jsonc`,
				description: "JSON Schema URL for IDE validation and autocomplete.",
			},
			namespace: {
				type: "string",
				description:
					"Optional namespace override. Defaults to the config filename without extension. Used to group related schemas.",
				pattern: "^[a-zA-Z_][a-zA-Z0-9_]*$",
			},
			schemas: {
				type: "array",
				description: "Array of schema definitions to convert.",
				items: {
					type: "object",
					description:
						"A schema definition specifying source and target adapter.",
					properties: {
						id: {
							type: "string",
							description:
								"Unique identifier for this schema within its namespace. Used as the variable name in generated code.",
							pattern: "^[a-zA-Z_][a-zA-Z0-9_]*$",
						},
						sourceType: {
							type: "string",
							enum: ["url", "file", "json"],
							description:
								"How to retrieve the JSON Schema. 'url' fetches from a remote URL, 'file' reads from a local file path (relative to this config), 'json' uses an inline schema object.",
						},
						source: {
							description:
								"The schema source. For 'url': a full URL. For 'file': a relative path. For 'json': an inline JSON Schema object.",
						},
						adapter: {
							type: "string",
							enum: lang.adapters,
							description: `The validation library adapter to use for code generation. Available adapters: ${lang.adapters.join(", ")}.`,
						},
					},
					required: ["id", "sourceType", "source", "adapter"],
					allOf: [urlFileCondition, jsonCondition],
				},
			},
		},
		required: ["$schema", "schemas"],
	};
}

/**
 * Write schema to file
 */
export function writeSchema(path: string, schema: object): void {
	const dir = dirname(path);
	if (!existsSync(dir)) {
		mkdirSync(dir, { recursive: true });
	}
	writeFileSync(path, JSON.stringify(schema, null, 2) + "\n");
}

/**
 * Main
 */
function main(): void {
	console.log("Discovering languages and adapters...");

	const languages = discoverLanguages();

	if (languages.length === 0) {
		console.log("No languages found with adapters.");
		return;
	}

	console.log(`Found ${languages.length} language(s):`);
	for (const lang of languages) {
		console.log(`  - ${lang.name}: ${lang.adapters.join(", ")}`);
	}

	console.log("\nGenerating schemas...");

	for (const lang of languages) {
		const schema = generateSchema(lang);

		// Write versioned schema
		const v1Path = join(PUBLIC_DIR, "schemas", "v1", `${lang.name}.jsonc`);
		writeSchema(v1Path, schema);
		console.log(`  Created: ${v1Path}`);

		// Write unversioned schema (copy of latest)
		const latestPath = join(PUBLIC_DIR, "schemas", `${lang.name}.jsonc`);
		writeSchema(latestPath, schema);
		console.log(`  Created: ${latestPath}`);
	}

	console.log("\nSchema generation complete!");
}

// Only run main when executed directly
if (import.meta.main) {
	main();
}
