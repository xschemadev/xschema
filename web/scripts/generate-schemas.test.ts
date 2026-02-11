import { mkdirSync, rmSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { afterAll, beforeAll, describe, expect, test } from "bun:test";

import {
  type LanguageInfo,
  discoverAdapters,
  discoverLanguages,
  generateSchema,
  getLanguageDisplayName,
} from "./generate-schemas";

const SCRIPT_DIR = dirname(fileURLToPath(import.meta.url));
const TEST_DIR = join(SCRIPT_DIR, ".test-fixtures");

function writeAdapterMetadata(adapterDir: string, contents: string): void {
  mkdirSync(adapterDir, { recursive: true });
  writeFileSync(join(adapterDir, "xschema.adapter.json"), contents);
}

beforeAll(() => {
  mkdirSync(TEST_DIR, { recursive: true });
});

afterAll(() => {
  rmSync(TEST_DIR, { recursive: true, force: true });
});

describe("discoverAdapters", () => {
  test("reads adapter refs from xschema.adapter.json", () => {
    const adaptersDir = join(TEST_DIR, "adapters1");
    writeAdapterMetadata(
      join(adaptersDir, "zod"),
      JSON.stringify({ ref: "@xschemadev/zod" }),
    );
    writeAdapterMetadata(
      join(adaptersDir, "yup"),
      JSON.stringify({ ref: "@xschemadev/yup" }),
    );

    const adapters = discoverAdapters(adaptersDir);

    expect(adapters).toEqual(["@xschemadev/yup", "@xschemadev/zod"]);
  });

  test("ignores files", () => {
    const adaptersDir = join(TEST_DIR, "adapters2");
    mkdirSync(adaptersDir, { recursive: true });
    writeAdapterMetadata(
      join(adaptersDir, "zod"),
      JSON.stringify({ ref: "@xschemadev/zod" }),
    );
    writeFileSync(join(adaptersDir, "README.md"), "# Adapters");

    const adapters = discoverAdapters(adaptersDir);

    expect(adapters).toEqual(["@xschemadev/zod"]);
  });

  test("ignores hidden directories", () => {
    const adaptersDir = join(TEST_DIR, "adapters3");
    writeAdapterMetadata(
      join(adaptersDir, "zod"),
      JSON.stringify({ ref: "@xschemadev/zod" }),
    );
    writeAdapterMetadata(
      join(adaptersDir, ".hidden"),
      JSON.stringify({ ref: "@xschemadev/hidden" }),
    );

    const adapters = discoverAdapters(adaptersDir);

    expect(adapters).toEqual(["@xschemadev/zod"]);
  });

  test("errors on missing xschema.adapter.json", () => {
    const adaptersDir = join(TEST_DIR, "adapters-missing");
    mkdirSync(join(adaptersDir, "zod"), { recursive: true });

    expect(() => discoverAdapters(adaptersDir)).toThrow(
      "missing xschema.adapter.json",
    );
  });

  test("errors on invalid json", () => {
    const adaptersDir = join(TEST_DIR, "adapters-invalid-json");
    writeAdapterMetadata(join(adaptersDir, "zod"), "{");

    expect(() => discoverAdapters(adaptersDir)).toThrow(
      "invalid json in adapter metadata",
    );
  });

  test("errors on missing ref", () => {
    const adaptersDir = join(TEST_DIR, "adapters-missing-ref");
    writeAdapterMetadata(join(adaptersDir, "zod"), JSON.stringify({}));

    expect(() => discoverAdapters(adaptersDir)).toThrow(
      'adapter metadata missing string "ref"',
    );
  });

  test("errors on duplicate ref", () => {
    const adaptersDir = join(TEST_DIR, "adapters-duplicate-ref");
    writeAdapterMetadata(
      join(adaptersDir, "zod"),
      JSON.stringify({ ref: "@xschemadev/zod" }),
    );
    writeAdapterMetadata(
      join(adaptersDir, "zod2"),
      JSON.stringify({ ref: "@xschemadev/zod" }),
    );

    expect(() => discoverAdapters(adaptersDir)).toThrow(
      "duplicate adapter ref",
    );
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
    writeAdapterMetadata(
      join(repoDir, "typescript", "packages", "adapters", "zod"),
      JSON.stringify({ ref: "@xschemadev/zod" }),
    );

    const languages = discoverLanguages(repoDir);

    expect(languages).toHaveLength(1);
    expect(languages[0].name).toBe("typescript");
    expect(languages[0].adapters).toEqual(["@xschemadev/zod"]);
  });

  test("finds multiple languages", () => {
    const repoDir = join(TEST_DIR, "repo2");
    writeAdapterMetadata(
      join(repoDir, "typescript", "packages", "adapters", "zod"),
      JSON.stringify({ ref: "@xschemadev/zod" }),
    );
    writeAdapterMetadata(
      join(repoDir, "python", "packages", "adapters", "pydantic"),
      JSON.stringify({ ref: "@xschemadev/pydantic" }),
    );

    const languages = discoverLanguages(repoDir);

    expect(languages).toHaveLength(2);
    expect(languages.map((l) => l.name)).toEqual(["python", "typescript"]);
  });

  test("errors on duplicate ref across languages", () => {
    const repoDir = join(TEST_DIR, "repo-duplicate-ref");
    writeAdapterMetadata(
      join(repoDir, "typescript", "packages", "adapters", "zod"),
      JSON.stringify({ ref: "@xschemadev/zod" }),
    );
    writeAdapterMetadata(
      join(repoDir, "python", "packages", "adapters", "also-zod"),
      JSON.stringify({ ref: "@xschemadev/zod" }),
    );

    expect(() => discoverLanguages(repoDir)).toThrow("duplicate adapter ref");
  });
});

describe("getLanguageDisplayName", () => {
  test("returns TypeScript for typescript", () => {
    expect(getLanguageDisplayName("typescript")).toBe("TypeScript");
  });

  test("returns TypeScript for ts", () => {
    expect(getLanguageDisplayName("ts")).toBe("TypeScript");
  });

  test("returns Python for python", () => {
    expect(getLanguageDisplayName("python")).toBe("Python");
  });

  test("returns input for unknown language", () => {
    expect(getLanguageDisplayName("rust")).toBe("rust");
  });
});

describe("generateSchema", () => {
  const tsLang: LanguageInfo = {
    name: "typescript",
    adapters: ["@xschemadev/zod"],
  };

  test("includes JSON Schema draft 2020-12", () => {
    const schema = generateSchema(tsLang);

    expect(schema.$schema).toBe("https://json-schema.org/draft/2020-12/schema");
  });

  test("has correct $id with version", () => {
    const schema = generateSchema(tsLang);

    expect(schema.$id).toBe("https://xschema.dev/schemas/v1/typescript.jsonc");
  });

  test("uses custom base URL", () => {
    const schema = generateSchema(tsLang, "https://example.com/schemas");

    expect(schema.$id).toBe("https://example.com/schemas/v1/typescript.jsonc");
  });

  test("has title with display name", () => {
    const schema = generateSchema(tsLang);

    expect(schema.title).toBe("XSchema TypeScript Configuration");
  });

  test("$schema property const matches unversioned URL", () => {
    const schema = generateSchema(tsLang);
    const properties = schema.properties as Record<string, unknown>;
    const schemaProperty = properties.$schema as Record<string, unknown>;

    expect(schemaProperty.const).toBe(
      "https://xschema.dev/schemas/typescript.jsonc",
    );
  });

  test("adapter enum contains all adapters (sorted)", () => {
    const lang: LanguageInfo = {
      name: "typescript",
      adapters: [
        "@xschemadev/zod",
        "@xschemadev/valibot",
        "@xschemadev/arktype",
      ],
    };
    const schema = generateSchema(lang);

    const properties = schema.properties as Record<string, unknown>;
    const schemas = properties.schemas as Record<string, unknown>;
    const items = schemas.items as Record<string, unknown>;
    const itemProps = items.properties as Record<string, unknown>;
    const adapter = itemProps.adapter as Record<string, unknown>;

    expect(adapter.enum).toEqual([
      "@xschemadev/arktype",
      "@xschemadev/valibot",
      "@xschemadev/zod",
    ]);
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
