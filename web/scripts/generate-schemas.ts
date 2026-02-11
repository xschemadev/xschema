/**
 * Schema Generation Script
 *
 * Generates JSON Schema files that IDEs can use for xschema config validation.
 *
 * Languages are directories at the repo root that contain packages/adapters/.
 * Adapters are discovered from {lang}/packages/adapters/<adapter>/xschema.adapter.json.
 */

import {
  existsSync,
  mkdirSync,
  readdirSync,
  readFileSync,
  statSync,
  writeFileSync,
} from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const SCRIPT_DIR = dirname(fileURLToPath(import.meta.url));
const REPO_ROOT = join(SCRIPT_DIR, "..", "..");
const PUBLIC_DIR = join(SCRIPT_DIR, "..", "public");
const BASE_URL = "https://xschema.dev/schemas";

export interface LanguageInfo {
  name: string; // directory name, e.g., "typescript"
  adapters: string[]; // adapter refs, e.g., ["@xschemadev/zod"]
}

interface AdapterMetadata {
  ref: string;
}

function parseAdapterMetadata(metadataPath: string): AdapterMetadata {
  let raw: string;
  try {
    raw = readFileSync(metadataPath, "utf8");
  } catch (error) {
    throw new Error(
      `[generate-schemas] failed reading adapter metadata: ${metadataPath}: ${String(error)}`,
    );
  }

  let parsed: unknown;
  try {
    parsed = JSON.parse(raw);
  } catch (error) {
    throw new Error(
      `[generate-schemas] invalid json in adapter metadata: ${metadataPath}: ${String(error)}`,
    );
  }

  if (!parsed || typeof parsed !== "object") {
    throw new Error(
      `[generate-schemas] adapter metadata must be an object: ${metadataPath}`,
    );
  }

  const ref = (parsed as Record<string, unknown>).ref;
  if (typeof ref !== "string" || ref.trim().length === 0) {
    throw new Error(
      `[generate-schemas] adapter metadata missing string "ref": ${metadataPath}`,
    );
  }

  return { ref };
}

function readAdapterRef(adapterDir: string): string {
  const metadataPath = join(adapterDir, "xschema.adapter.json");
  if (!existsSync(metadataPath)) {
    throw new Error(
      `[generate-schemas] missing xschema.adapter.json for adapter: ${adapterDir}`,
    );
  }

  return parseAdapterMetadata(metadataPath).ref;
}

/**
 * Discover adapters by listing subdirectories and reading `xschema.adapter.json`.
 */
export function discoverAdapters(
  adaptersPath: string,
  globalRefToAdapterDir?: Map<string, string>,
): string[] {
  const refs: string[] = [];
  const refsInLanguage = new Map<string, string>();

  const entries = readdirSync(adaptersPath, { withFileTypes: true });

  for (const entry of entries) {
    if (!entry.isDirectory()) continue;
    if (entry.name.startsWith(".")) continue;

    const adapterDir = join(adaptersPath, entry.name);
    const ref = readAdapterRef(adapterDir);

    const existingInLanguage = refsInLanguage.get(ref);
    if (existingInLanguage) {
      throw new Error(
        `[generate-schemas] duplicate adapter ref "${ref}": ${existingInLanguage} and ${adapterDir}`,
      );
    }
    refsInLanguage.set(ref, adapterDir);

    if (globalRefToAdapterDir) {
      const existingGlobal = globalRefToAdapterDir.get(ref);
      if (existingGlobal) {
        throw new Error(
          `[generate-schemas] duplicate adapter ref "${ref}": ${existingGlobal} and ${adapterDir}`,
        );
      }
      globalRefToAdapterDir.set(ref, adapterDir);
    }

    refs.push(ref);
  }

  refs.sort();
  return refs;
}

/**
 * Discover languages by finding directories with `packages/adapters/` subdirectory.
 */
export function discoverLanguages(
  repoRoot: string = REPO_ROOT,
): LanguageInfo[] {
  const languages: LanguageInfo[] = [];
  const globalRefToAdapterDir = new Map<string, string>();

  const entries = readdirSync(repoRoot, { withFileTypes: true });

  for (const entry of entries) {
    if (!entry.isDirectory()) continue;
    if (entry.name.startsWith(".") || entry.name === "node_modules") continue;

    const adaptersPath = join(repoRoot, entry.name, "packages", "adapters");

    if (existsSync(adaptersPath) && statSync(adaptersPath).isDirectory()) {
      const adapters = discoverAdapters(adaptersPath, globalRefToAdapterDir);
      if (adapters.length > 0) {
        languages.push({
          name: entry.name,
          adapters,
        });
      }
    }
  }

  languages.sort((a, b) => a.name.localeCompare(b.name));
  return languages;
}

/**
 * Get display name for a language.
 */
export function getLanguageDisplayName(lang: string): string {
  const names: Record<string, string> = {
    ts: "TypeScript",
    typescript: "TypeScript",
    py: "Python",
    python: "Python",
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
 * Generate JSON Schema for a language.
 */
export function generateSchema(
  lang: LanguageInfo,
  baseUrl: string = BASE_URL,
): Record<string, unknown> {
  const displayName = getLanguageDisplayName(lang.name);
  const adapters = [...lang.adapters].sort();

  // Build conditional validation rules for source field
  const urlCondition = createConditional(
    {
      properties: {
        sourceType: { const: "url" },
      },
    },
    {
      properties: {
        source: {
          type: "string",
          description: "URL to fetch the JSON Schema from.",
        },
        headers: {
          type: "object",
          description:
            "HTTP headers to include when fetching. Values support ${ENV_VAR} syntax for environment variable substitution.",
          additionalProperties: { type: "string" },
        },
      },
    },
  );

  const fileCondition = createConditional(
    {
      properties: {
        sourceType: { const: "file" },
      },
    },
    {
      properties: {
        source: {
          type: "string",
          description:
            "File path to the JSON Schema. Relative paths are resolved from this config file.",
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
              enum: adapters,
              description: `The validation library adapter to use for code generation. Available adapters: ${adapters.join(", ")}.`,
            },
          },
          required: ["id", "sourceType", "source", "adapter"],
          allOf: [urlCondition, fileCondition, jsonCondition],
        },
      },
    },
    required: ["$schema", "schemas"],
  };
}

/**
 * Write schema to file.
 */
export function writeSchema(path: string, schema: object): void {
  const dir = dirname(path);
  if (!existsSync(dir)) {
    mkdirSync(dir, { recursive: true });
  }
  writeFileSync(path, JSON.stringify(schema, null, 2) + "\n");
}

/**
 * Main.
 */
function main(): void {
  console.log("discovering languages and adapters...");

  const languages = discoverLanguages();

  if (languages.length === 0) {
    console.log("no languages found with adapters");
    return;
  }

  console.log(`found ${languages.length} language(s):`);
  for (const lang of languages) {
    console.log(`  - ${lang.name}: ${lang.adapters.join(", ")}`);
  }

  console.log("\ngenerating schemas...");

  for (const lang of languages) {
    const schema = generateSchema(lang);

    // Write versioned schema
    const v1Path = join(PUBLIC_DIR, "schemas", "v1", `${lang.name}.jsonc`);
    writeSchema(v1Path, schema);
    console.log(`  created: ${v1Path}`);

    // Write unversioned schema (copy of latest)
    const latestPath = join(PUBLIC_DIR, "schemas", `${lang.name}.jsonc`);
    writeSchema(latestPath, schema);
    console.log(`  created: ${latestPath}`);
  }

  console.log("\nschema generation complete");
}

// Only run main when executed directly
if (import.meta.main) {
  main();
}
