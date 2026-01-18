#!/usr/bin/env bun
/**
 * Unsupported Features MDX Generation Script
 *
 * Generates the unsupported-features.mdx page from cli/unsupported/unsupported-features.json.
 * Each limitation group is displayed as an accordion with:
 * - Header: name and keywords that trigger the limitation
 * - Body: reason explanation and detailed explanation
 *
 * Output: content/docs/compliance/unsupported-features.mdx
 */

import { existsSync, mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";

const REPO_ROOT = join(import.meta.dirname, "..", "..");
const CONTENT_DIR = join(import.meta.dirname, "..", "content", "docs");
const UNSUPPORTED_FEATURES_PATH = join(REPO_ROOT, "cli", "unsupported", "unsupported-features.json");

interface UnsupportedFeatureGroup {
  name: string;
  reason: string;
  explanation?: string;
  keywords: string[];
}

function loadUnsupportedFeatures(): UnsupportedFeatureGroup[] {
  if (!existsSync(UNSUPPORTED_FEATURES_PATH)) {
    throw new Error(
      `Unsupported features file not found at: ${UNSUPPORTED_FEATURES_PATH}\n` +
        "Ensure cli/unsupported/unsupported-features.json exists.",
    );
  }

  const content = readFileSync(UNSUPPORTED_FEATURES_PATH, "utf-8");
  return JSON.parse(content) as UnsupportedFeatureGroup[];
}

/**
 * Escape curly braces for MDX (they're interpreted as JSX expressions)
 */
function escapeMdx(str: string): string {
  return str.replace(/\{/g, "\\{").replace(/\}/g, "\\}");
}

function generateUnsupportedFeaturesMdx(groups: UnsupportedFeatureGroup[]): string {
  const lines: string[] = [];

  // Frontmatter
  lines.push("---");
  lines.push("title: Unsupported Features");
  lines.push(
    "description: Fundamental limitations of static JSON Schema code generation",
  );
  lines.push("icon: Ban");
  lines.push("---");
  lines.push("");
  lines.push("{/* This file is auto-generated. Do not edit manually. */}");
  lines.push("");

  // Intro section
  lines.push(
    "XSchema generates validation code at build time from JSON Schema definitions. " +
      "This static approach provides type safety, performance, and smaller bundles. " +
      "However, some JSON Schema features require runtime evaluation that cannot be " +
      "pre-compiled.",
  );
  lines.push("");
  lines.push(
    "The following limitations are fundamental to any static code generation approach, " +
      "not specific to XSchema or any particular adapter.",
  );
  lines.push("");

  // Count total keywords
  const totalKeywords = groups.reduce((sum, g) => sum + g.keywords.length, 0);
  lines.push(
    `**${groups.length} limitation categories** covering **${totalKeywords} keywords** that cannot be statically compiled.`,
  );
  lines.push("");

  // Accordions for each group
  lines.push("<Accordions>");
  lines.push("");

  for (const group of groups) {
    const keywordDisplay =
      group.keywords.length > 0
        ? group.keywords.map((k) => `\`${k}\``).join(", ")
        : "context-dependent";
    const title = `${formatGroupName(group.name)}`;
    lines.push(`<Accordion title="${title}">`);
    lines.push("");
    lines.push(`**Keywords:** ${keywordDisplay}`);
    lines.push("");
    lines.push(group.reason);
    lines.push("");

    // Add detailed explanation if present
    if (group.explanation) {
      lines.push("### Why This Cannot Be Supported");
      lines.push("");
      // Split explanation by double newlines to preserve paragraph structure
      const paragraphs = group.explanation.split("\n\n");
      for (const paragraph of paragraphs) {
        lines.push(escapeMdx(paragraph));
        lines.push("");
      }
    }

    lines.push("</Accordion>");
    lines.push("");
  }

  lines.push("</Accordions>");
  lines.push("");

  return lines.join("\n");
}

/**
 * Convert kebab-case name to Title Case for display
 */
function formatGroupName(name: string): string {
  return name
    .split("-")
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
    .join(" ");
}

function writeMdx(path: string, content: string): void {
  const dir = dirname(path);
  if (!existsSync(dir)) {
    mkdirSync(dir, { recursive: true });
  }
  writeFileSync(path, content);
}

function main(): void {
  console.log("Loading unsupported features from CLI...");
  const groups = loadUnsupportedFeatures();
  console.log(`  Found ${groups.length} limitation group(s)`);

  console.log("\nGenerating unsupported-features MDX...");
  const mdxContent = generateUnsupportedFeaturesMdx(groups);
  const outputPath = join(CONTENT_DIR, "compliance", "unsupported-features.mdx");
  writeMdx(outputPath, mdxContent);
  console.log(`  Created: ${outputPath}`);

  console.log("\nUnsupported features MDX generation complete!");
}

main();
