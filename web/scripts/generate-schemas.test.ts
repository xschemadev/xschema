import { mkdirSync, rmSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { afterAll, beforeAll, describe, expect, test } from "bun:test";

import {
  type LanguageInfo,
  discoverAdapters,
  discoverLanguages,
  generateSchema,
  getLanguageDisplayName,
} from "./generate-schemas";

const TEST_DIR = join(import.meta.dirname, ".test-fixtures");

beforeAll(() => {
  // Create test fixture directory structure
  mkdirSync(TEST_DIR, { recursive: true });
});

afterAll(() => {
  // Clean up test fixtures
  rmSync(TEST_DIR, { recursive: true, force: true });
});

describe("discoverAdapters", () => {
  test("finds adapter directories", () => {
    const adaptersDir = join(TEST_DIR, "adapters1");
    mkdirSync(join(adaptersDir, "zod"), { recursive: true });
    mkdirSync(join(adaptersDir, "yup"), { recursive: true });

    const adapters = discoverAdapters(adaptersDir);

    expect(adapters).toContain("zod");
    expect(adapters).toContain("yup");
    expect(adapters).toHaveLength(2);
  });

  test("ignores files", () => {
    const adaptersDir = join(TEST_DIR, "adapters2");
    mkdirSync(adaptersDir, { recursive: true });
    mkdirSync(join(adaptersDir, "zod"), { recursive: true });
    writeFileSync(join(adaptersDir, "README.md"), "# Adapters");

    const adapters = discoverAdapters(adaptersDir);

    expect(adapters).toEqual(["zod"]);
  });

  test("ignores hidden directories", () => {
    const adaptersDir = join(TEST_DIR, "adapters3");
    mkdirSync(join(adaptersDir, "zod"), { recursive: true });
    mkdirSync(join(adaptersDir, ".hidden"), { recursive: true });

    const adapters = discoverAdapters(adaptersDir);

    expect(adapters).toEqual(["zod"]);
  });

  test("returns empty array for empty directory", () => {
    const adaptersDir = join(TEST_DIR, "adapters-empty");
    mkdirSync(adaptersDir, { recursive: true });

    const adapters = discoverAdapters(adaptersDir);

    expect(adapters).toEqual([]);
  });
});

describe("discoverLanguages", () => {
  test("finds language with packages/adapters structure", () => {
    const repoDir = join(TEST_DIR, "repo1");
    mkdirSync(join(repoDir, "ts", "packages", "adapters", "zod"), {
      recursive: true,
    });

    const languages = discoverLanguages(repoDir);

    expect(languages).toHaveLength(1);
    expect(languages[0].name).toBe("ts");
    expect(languages[0].adapters).toEqual(["zod"]);
  });

  test("finds multiple languages", () => {
    const repoDir = join(TEST_DIR, "repo2");
    mkdirSync(join(repoDir, "ts", "packages", "adapters", "zod"), {
      recursive: true,
    });
    mkdirSync(join(repoDir, "py", "packages", "adapters", "pydantic"), {
      recursive: true,
    });

    const languages = discoverLanguages(repoDir);

    expect(languages).toHaveLength(2);
    const names = languages.map((l) => l.name).sort();
    expect(names).toEqual(["py", "ts"]);
  });

  test("ignores directories without packages/adapters", () => {
    const repoDir = join(TEST_DIR, "repo3");
    mkdirSync(join(repoDir, "ts", "packages", "adapters", "zod"), {
      recursive: true,
    });
    mkdirSync(join(repoDir, "docs"), { recursive: true });
    mkdirSync(join(repoDir, "cli", "cmd"), { recursive: true });

    const languages = discoverLanguages(repoDir);

    expect(languages).toHaveLength(1);
    expect(languages[0].name).toBe("ts");
  });

  test("ignores hidden directories", () => {
    const repoDir = join(TEST_DIR, "repo4");
    mkdirSync(join(repoDir, "ts", "packages", "adapters", "zod"), {
      recursive: true,
    });
    mkdirSync(join(repoDir, ".git", "packages", "adapters", "fake"), {
      recursive: true,
    });

    const languages = discoverLanguages(repoDir);

    expect(languages).toHaveLength(1);
    expect(languages[0].name).toBe("ts");
  });

  test("ignores node_modules", () => {
    const repoDir = join(TEST_DIR, "repo5");
    mkdirSync(join(repoDir, "ts", "packages", "adapters", "zod"), {
      recursive: true,
    });
    mkdirSync(join(repoDir, "node_modules", "packages", "adapters", "fake"), {
      recursive: true,
    });

    const languages = discoverLanguages(repoDir);

    expect(languages).toHaveLength(1);
    expect(languages[0].name).toBe("ts");
  });

  test("ignores language with empty adapters directory", () => {
    const repoDir = join(TEST_DIR, "repo6");
    mkdirSync(join(repoDir, "ts", "packages", "adapters"), { recursive: true });

    const languages = discoverLanguages(repoDir);

    expect(languages).toHaveLength(0);
  });
});

describe("getLanguageDisplayName", () => {
  test("returns TypeScript for ts", () => {
    expect(getLanguageDisplayName("ts")).toBe("TypeScript");
  });

  test("returns Python for py", () => {
    expect(getLanguageDisplayName("py")).toBe("Python");
  });

  test("returns input for unknown language", () => {
    expect(getLanguageDisplayName("rust")).toBe("rust");
  });
});

describe("generateSchema", () => {
  const tsLang: LanguageInfo = { name: "ts", adapters: ["zod"] };

  test("includes JSON Schema draft 2020-12", () => {
    const schema = generateSchema(tsLang);

    expect(schema.$schema).toBe("https://json-schema.org/draft/2020-12/schema");
  });

  test("has correct $id with version", () => {
    const schema = generateSchema(tsLang);

    expect(schema.$id).toBe("https://xschema.dev/schemas/v1/ts.jsonc");
  });

  test("uses custom base URL", () => {
    const schema = generateSchema(tsLang, "https://example.com/schemas");

    expect(schema.$id).toBe("https://example.com/schemas/v1/ts.jsonc");
  });

  test("has title with display name", () => {
    const schema = generateSchema(tsLang);

    expect(schema.title).toBe("XSchema TypeScript Configuration");
  });

  test("$schema property const matches unversioned URL", () => {
    const schema = generateSchema(tsLang);
    const properties = schema.properties as Record<string, unknown>;
    const schemaProperty = properties.$schema as Record<string, unknown>;

    expect(schemaProperty.const).toBe("https://xschema.dev/schemas/ts.jsonc");
  });

  test("adapter enum contains all adapters", () => {
    const lang: LanguageInfo = { name: "ts", adapters: ["zod", "yup", "valibot"] };
    const schema = generateSchema(lang);

    const properties = schema.properties as Record<string, unknown>;
    const schemas = properties.schemas as Record<string, unknown>;
    const items = schemas.items as Record<string, unknown>;
    const itemProps = items.properties as Record<string, unknown>;
    const adapter = itemProps.adapter as Record<string, unknown>;

    expect(adapter.enum).toEqual(["zod", "yup", "valibot"]);
  });

  test("has required fields at root level", () => {
    const schema = generateSchema(tsLang);

    expect(schema.required).toEqual(["$schema", "schemas"]);
  });

  test("schema items have required fields", () => {
    const schema = generateSchema(tsLang);

    const properties = schema.properties as Record<string, unknown>;
    const schemas = properties.schemas as Record<string, unknown>;
    const items = schemas.items as Record<string, unknown>;

    expect(items.required).toEqual(["id", "sourceType", "source", "adapter"]);
  });

  test("has conditional validation for source field", () => {
    const schema = generateSchema(tsLang);

    const properties = schema.properties as Record<string, unknown>;
    const schemas = properties.schemas as Record<string, unknown>;
    const items = schemas.items as Record<string, unknown>;
    const allOf = items.allOf as Array<Record<string, unknown>>;

    expect(allOf).toHaveLength(2);

    // First condition: url/file -> source is string
    expect(allOf[0].if).toEqual({
      properties: { sourceType: { enum: ["url", "file"] } },
    });
    expect(allOf[0]["then"]).toEqual({
      properties: {
        source: {
          type: "string",
          description:
            "URL or file path to the JSON Schema. File paths are relative to this config file.",
        },
      },
    });

    // Second condition: json -> source is object
    expect(allOf[1].if).toEqual({
      properties: { sourceType: { const: "json" } },
    });
    expect(allOf[1]["then"]).toEqual({
      properties: {
        source: { type: "object", description: "Inline JSON Schema object." },
      },
    });
  });

  test("sourceType has correct enum values", () => {
    const schema = generateSchema(tsLang);

    const properties = schema.properties as Record<string, unknown>;
    const schemas = properties.schemas as Record<string, unknown>;
    const items = schemas.items as Record<string, unknown>;
    const itemProps = items.properties as Record<string, unknown>;
    const sourceType = itemProps.sourceType as Record<string, unknown>;

    expect(sourceType.enum).toEqual(["url", "file", "json"]);
  });
});
