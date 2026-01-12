#!/usr/bin/env bun
/**
 * Known Issues MDX Generation Script
 *
 * Generates the known-issues.mdx page from cli/compliance/known-issues.json.
 * Each limitation group is displayed as an accordion with:
 * - Header: name and count of affected tests
 * - Body: reason explanation and list of test names
 *
 * Output: content/docs/compliance/known-issues.mdx
 */

import { existsSync, mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";

const REPO_ROOT = join(import.meta.dirname, "..", "..");
const CONTENT_DIR = join(import.meta.dirname, "..", "content", "docs");
const KNOWN_ISSUES_PATH = join(REPO_ROOT, "cli", "compliance", "known-issues.json");

interface KnownIssueGroup {
  name: string;
  reason: string;
  tests: string[];
}

function loadKnownIssues(): KnownIssueGroup[] {
  if (!existsSync(KNOWN_ISSUES_PATH)) {
    throw new Error(
      `Known issues file not found at: ${KNOWN_ISSUES_PATH}\n` +
        "Ensure cli/compliance/known-issues.json exists.",
    );
  }

  const content = readFileSync(KNOWN_ISSUES_PATH, "utf-8");
  return JSON.parse(content) as KnownIssueGroup[];
}

/**
 * Escape curly braces for MDX (they're interpreted as JSX expressions)
 */
function escapeMdx(str: string): string {
  return str.replace(/\{/g, "\\{").replace(/\}/g, "\\}");
}

function generateKnownIssuesMdx(groups: KnownIssueGroup[]): string {
  const lines: string[] = [];

  // Frontmatter
  lines.push("---");
  lines.push("title: Known Issues");
  lines.push(
    "description: Fundamental limitations of static JSON Schema code generation",
  );
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

  // Total count
  const totalTests = groups.reduce((sum, g) => sum + g.tests.length, 0);
  lines.push(
    `**${groups.length} limitation categories** affecting **${totalTests} tests** in the JSON Schema Test Suite.`,
  );
  lines.push("");

  // Accordions for each group
  lines.push("<Accordions>");
  lines.push("");

  for (const group of groups) {
    const title = `${formatGroupName(group.name)} (${group.tests.length} tests)`;
    lines.push(`<Accordion title="${title}">`);
    lines.push("");
    lines.push(group.reason);
    lines.push("");
    lines.push("**Affected tests:**");
    lines.push("");
    for (const test of group.tests) {
      lines.push(`- ${escapeMdx(test)}`);
    }
    lines.push("");
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
  console.log("Loading known issues from CLI...");
  const groups = loadKnownIssues();
  console.log(`  Found ${groups.length} limitation group(s)`);

  console.log("\nGenerating known-issues MDX...");
  const mdxContent = generateKnownIssuesMdx(groups);
  const outputPath = join(CONTENT_DIR, "compliance", "known-issues.mdx");
  writeMdx(outputPath, mdxContent);
  console.log(`  Created: ${outputPath}`);

  console.log("\nKnown issues MDX generation complete!");
}

main();
