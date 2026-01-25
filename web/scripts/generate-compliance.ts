#!/usr/bin/env bun
/**
 * Compliance MDX Generation Script
 *
 * Discovers adapters from the docs structure and generates compliance MDX pages
 * from their compliance test results.
 *
 * Logic:
 * 1. Scan content/docs/{lang}/adapters/{adapter}/ for existing adapter docs
 * 2. For each adapter doc found, require compliance results to exist
 * 3. If results missing → ERROR (fail the build)
 * 4. If results exist → generate compliance.mdx
 *
 * Output: content/docs/{lang}/adapters/{adapter}/compliance.mdx
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

const REPO_ROOT = join(import.meta.dirname, "..", "..");
const CONTENT_DIR = join(import.meta.dirname, "..", "content", "docs");

interface KeywordResult {
  keyword: string;
  passed: number;
  failed: number;
  skipped: number;
  total: number;
  failures?: {
    group: string;
    test: string;
    expected: boolean;
    actual: string;
    passed: boolean;
  }[];
}

interface DraftResult {
  draft: string;
  keywords: KeywordResult[];
  summary: {
    passed: number;
    failed: number;
    skipped: number;
    total: number;
    percentage: number;
    unsupportedFeatures?: {
      count: number;
    };
  };
}

interface AdapterInfo {
  lang: string;
  adapter: string;
  drafts: DraftResult[];
}

interface AdapterDoc {
  lang: string;
  adapter: string;
}

/**
 * Discover adapters from the docs structure.
 * Looks for content/docs/{lang}/adapters/{adapter}/index.mdx
 */
function discoverAdaptersFromDocs(): AdapterDoc[] {
  const adapters: AdapterDoc[] = [];

  // List all directories in content/docs/
  if (!existsSync(CONTENT_DIR)) {
    return adapters;
  }

  const langEntries = readdirSync(CONTENT_DIR, { withFileTypes: true });

  for (const langEntry of langEntries) {
    if (!langEntry.isDirectory()) continue;
    if (langEntry.name.startsWith(".")) continue;

    const adaptersDocsPath = join(CONTENT_DIR, langEntry.name, "adapters");

    if (
      !existsSync(adaptersDocsPath) ||
      !statSync(adaptersDocsPath).isDirectory()
    ) {
      continue;
    }

    const adapterEntries = readdirSync(adaptersDocsPath, {
      withFileTypes: true,
    });

    for (const adapterEntry of adapterEntries) {
      if (!adapterEntry.isDirectory()) continue;
      if (adapterEntry.name.startsWith(".")) continue;

      // Check if index.mdx exists (adapter has docs)
      const indexPath = join(adaptersDocsPath, adapterEntry.name, "index.mdx");
      if (existsSync(indexPath)) {
        adapters.push({
          lang: langEntry.name,
          adapter: adapterEntry.name,
        });
      }
    }
  }

  return adapters;
}

/**
 * Get compliance results for an adapter.
 * Throws an error if results don't exist.
 */
function getComplianceResults(lang: string, adapter: string): DraftResult[] {
  const compliancePath = join(
    REPO_ROOT,
    lang,
    "packages",
    "adapters",
    adapter,
    "compliance",
    "results",
  );

  if (!existsSync(compliancePath) || !statSync(compliancePath).isDirectory()) {
    throw new Error(
      `Missing compliance results for ${lang}/${adapter}.\n` +
        `Expected results at: ${compliancePath}\n` +
        `Run compliance tests for this adapter before building docs.`,
    );
  }

  const drafts = readComplianceResults(compliancePath);

  if (drafts.length === 0) {
    throw new Error(
      `No compliance JSON files found for ${lang}/${adapter}.\n` +
        `Expected JSON files at: ${compliancePath}\n` +
        `Run compliance tests for this adapter before building docs.`,
    );
  }

  return drafts;
}

/**
 * Read all compliance JSON files from a results directory
 */
function readComplianceResults(resultsPath: string): DraftResult[] {
  const drafts: DraftResult[] = [];
  const files = readdirSync(resultsPath);

  // Define preferred draft order
  const draftOrder = [
    "draft2020-12",
    "draft2019-09",
    "draft7",
    "draft6",
    "draft4",
    "draft3",
    "v1",
  ];

  for (const file of files) {
    if (!file.endsWith(".json")) continue;

    const filePath = join(resultsPath, file);
    try {
      const content = readFileSync(filePath, "utf-8");
      const data = JSON.parse(content) as DraftResult;
      drafts.push(data);
    } catch (e) {
      console.warn(`  Warning: Failed to parse ${filePath}`);
    }
  }

  // Sort by preferred order
  drafts.sort((a, b) => {
    const aIndex = draftOrder.indexOf(a.draft);
    const bIndex = draftOrder.indexOf(b.draft);
    if (aIndex === -1 && bIndex === -1) return a.draft.localeCompare(b.draft);
    if (aIndex === -1) return 1;
    if (bIndex === -1) return -1;
    return aIndex - bIndex;
  });

  return drafts;
}

/**
 * Escape curly braces for MDX (they're interpreted as JSX expressions)
 */
function escapeMdx(str: string): string {
  return str.replace(/\{/g, "\\{").replace(/\}/g, "\\}");
}

/**
 * Get display name for adapter
 */
function getAdapterDisplayName(adapter: string): string {
  const names: Record<string, string> = {
    zod: "Zod",
    yup: "Yup",
    valibot: "Valibot",
  };
  return names[adapter] || adapter.charAt(0).toUpperCase() + adapter.slice(1);
}

/**
 * Format percentage with color indicator
 */
function formatPercentage(percentage: number): string {
  return `${percentage.toFixed(1)}%`;
}

/**
 * Get status emoji based on pass rate
 */
function getStatusEmoji(passed: number, total: number): string {
  if (total === 0) return "➖";
  const rate = passed / total;
  if (rate === 1) return "✅";
  if (rate >= 0.9) return "🟡";
  return "❌";
}

/**
 * Group failures by keyword across all drafts
 */
interface KeywordIssue {
  keyword: string;
  affectedDrafts: {
    draft: string;
    passed: number;
    total: number;
    failures: NonNullable<KeywordResult["failures"]>;
  }[];
  totalFailures: number;
}

function groupIssuesByKeyword(drafts: DraftResult[]): KeywordIssue[] {
  const keywordMap = new Map<string, KeywordIssue>();

  for (const draft of drafts) {
    for (const kw of draft.keywords) {
      if (kw.failed > 0 && kw.failures && kw.failures.length > 0) {
        if (!keywordMap.has(kw.keyword)) {
          keywordMap.set(kw.keyword, {
            keyword: kw.keyword,
            affectedDrafts: [],
            totalFailures: 0,
          });
        }
        const issue = keywordMap.get(kw.keyword)!;
        issue.affectedDrafts.push({
          draft: draft.draft,
          passed: kw.passed,
          total: kw.total,
          failures: kw.failures,
        });
        issue.totalFailures += kw.failures.length;
      }
    }
  }

  // Sort by total failures descending
  return Array.from(keywordMap.values()).sort(
    (a, b) => b.totalFailures - a.totalFailures,
  );
}

/**
 * Check if adapter is type-only (no runtime validation).
 * Type-only adapters have no passed tests and at least some skipped tests.
 */
function isTypeOnlyAdapter(drafts: DraftResult[]): boolean {
  if (drafts.length === 0) return false;

  let totalPassed = 0;
  let totalSkipped = 0;

  for (const draft of drafts) {
    totalPassed += draft.summary.passed;
    totalSkipped += draft.summary.skipped;
  }

  return totalPassed === 0 && totalSkipped > 0;
}

/**
 * Generate MDX content for an adapter's compliance page
 */
function generateComplianceMdx(info: AdapterInfo): string {
  const adapterName = getAdapterDisplayName(info.adapter);
  const lines: string[] = [];

  // Frontmatter
  lines.push("---");
  lines.push(`title: ${adapterName} Compliance`);
  lines.push(
    `description: JSON Schema compliance report for @xschemadev/${info.adapter}`,
  );
  lines.push("icon: ClipboardCheck");
  lines.push("---");
  lines.push("");
  lines.push("{/* This file is auto-generated. Do not edit manually. */}");
  lines.push("");

  // Intro
  lines.push(
    `This report shows how well the ${adapterName} adapter handles JSON Schema validation. ` +
      "Tests are run against the official [JSON Schema Test Suite](https://github.com/json-schema-org/JSON-Schema-Test-Suite).",
  );
  lines.push("");

  // Type-only adapter callout - nothing else to show
  const typeOnly = isTypeOnlyAdapter(info.drafts);
  if (typeOnly) {
    lines.push('<Callout type="info" title="Type-Only Adapter">');
    lines.push(
      "This adapter generates TypeScript types only — no runtime validation. " +
        "Compliance tests are skipped since there's no validate function.",
    );
    lines.push("</Callout>");
    lines.push("");
    return lines.join("\n");
  }

  // Summary section
  lines.push("## Summary");
  lines.push("");

  // Check if any draft has unsupported features to decide on table format
  const hasUnsupportedFeatures = info.drafts.some(
    (d) =>
      d.summary.unsupportedFeatures && d.summary.unsupportedFeatures.count > 0,
  );

  if (hasUnsupportedFeatures) {
    lines.push("| Draft | Passed | Failed | Unsupported | Coverage |");
    lines.push("| ----- | ------ | ------ | ----------- | -------- |");
  } else {
    lines.push("| Draft | Passed | Failed | Coverage |");
    lines.push("| ----- | ------ | ------ | -------- |");
  }

  for (const draft of info.drafts) {
    const { passed, failed, percentage, unsupportedFeatures } = draft.summary;
    const unsupportedCount = unsupportedFeatures?.count ?? 0;
    const coverageEmoji =
      percentage >= 99 ? "🟢" : percentage >= 95 ? "🟡" : "🔴";
    if (hasUnsupportedFeatures) {
      lines.push(
        `| ${draft.draft} | ${passed} | ${failed} | ${unsupportedCount} | ${coverageEmoji} ${formatPercentage(percentage)} |`,
      );
    } else {
      lines.push(
        `| ${draft.draft} | ${passed} | ${failed} | ${coverageEmoji} ${formatPercentage(percentage)} |`,
      );
    }
  }
  lines.push("");

  // Coverage by Draft - using tabs
  lines.push("## Coverage by Draft");
  lines.push("");
  lines.push(
    "<Tabs items={[" +
      info.drafts.map((d) => `"${d.draft}"`).join(", ") +
      "]}>",
  );
  lines.push("");

  for (const draft of info.drafts) {
    lines.push(`<Tab value="${draft.draft}">`);
    lines.push("");
    lines.push(
      `**${draft.summary.passed}/${draft.summary.total}** tests passing (${formatPercentage(draft.summary.percentage)})`,
    );
    lines.push("");

    // Only show keywords with issues, or a success message
    const failingKeywords = draft.keywords.filter((kw) => kw.failed > 0);
    const passingKeywords = draft.keywords.filter(
      (kw) => kw.total > 0 && kw.failed === 0,
    );

    if (failingKeywords.length > 0) {
      lines.push("**Keywords with issues:**");
      lines.push("");
      lines.push("| Keyword | Status | Pass Rate |");
      lines.push("| ------- | ------ | --------- |");
      for (const kw of failingKeywords) {
        const rate = formatPercentage((kw.passed / kw.total) * 100);
        lines.push(
          `| \`${kw.keyword}\` | ${getStatusEmoji(kw.passed, kw.total)} | ${kw.passed}/${kw.total} (${rate}) |`,
        );
      }
      lines.push("");
    }

    if (passingKeywords.length > 0) {
      lines.push("<Accordions>");
      lines.push(
        `<Accordion title="✅ ${passingKeywords.length} fully supported keywords">`,
      );
      lines.push("");
      lines.push(passingKeywords.map((kw) => `\`${kw.keyword}\``).join(", "));
      lines.push("");
      lines.push("</Accordion>");
      lines.push("</Accordions>");
      lines.push("");
    }

    lines.push("</Tab>");
    lines.push("");
  }

  lines.push("</Tabs>");
  lines.push("");

  // Unexpected Failures section - actual failures grouped by keyword
  const issues = groupIssuesByKeyword(info.drafts);

  if (issues.length > 0) {
    lines.push("## Unexpected Failures");
    lines.push("");
    lines.push(
      "The following JSON Schema keywords have incomplete support. " +
        "These are edge cases that may not affect typical usage.",
    );
    lines.push("");
    lines.push("<Accordions>");

    for (const issue of issues) {
      const draftsAffected = issue.affectedDrafts
        .map((d) => d.draft)
        .join(", ");
      const title = `${issue.keyword} — ${issue.totalFailures} failing test(s)`;
      lines.push(`<Accordion title="${title}">`);
      lines.push("");
      lines.push(`**Affected drafts:** ${draftsAffected}`);
      lines.push("");

      // Show failures grouped by draft
      for (const draftInfo of issue.affectedDrafts) {
        const rate = formatPercentage(
          (draftInfo.passed / draftInfo.total) * 100,
        );
        lines.push(
          `**${draftInfo.draft}** (${draftInfo.passed}/${draftInfo.total} passing, ${rate}):`,
        );
        lines.push("");
        for (const failure of draftInfo.failures) {
          lines.push(`- ${escapeMdx(failure.group)}`);
          lines.push(`  - *${escapeMdx(failure.test)}*`);
        }
        lines.push("");
      }

      lines.push("</Accordion>");
    }

    lines.push("</Accordions>");
    lines.push("");
  }

  // Link to global unsupported features
  lines.push("## Unsupported Features");
  lines.push("");
  lines.push(
    "Some tests are excluded from compliance percentages due to fundamental limitations " +
      "of static code generation. See [Unsupported Features](/docs/compliance/unsupported-features) for details.",
  );
  lines.push("");

  return lines.join("\n");
}

/**
 * Write MDX file
 */
function writeMdx(path: string, content: string): void {
  const dir = dirname(path);
  if (!existsSync(dir)) {
    mkdirSync(dir, { recursive: true });
  }
  writeFileSync(path, content);
}

/**
 * Main
 */
function main(): void {
  console.log("Discovering adapters from docs...");

  const adapterDocs = discoverAdaptersFromDocs();

  if (adapterDocs.length === 0) {
    console.log("No adapter docs found. Nothing to generate.");
    return;
  }

  console.log(`Found ${adapterDocs.length} adapter doc(s):`);
  for (const doc of adapterDocs) {
    console.log(`  - ${doc.lang}/${doc.adapter}`);
  }

  console.log("\nLoading compliance results and generating MDX files...");

  const adapters: AdapterInfo[] = [];

  for (const doc of adapterDocs) {
    // This will throw if compliance results are missing
    const drafts = getComplianceResults(doc.lang, doc.adapter);
    adapters.push({
      lang: doc.lang,
      adapter: doc.adapter,
      drafts,
    });
    console.log(
      `  Loaded: ${doc.lang}/${doc.adapter} (${drafts.length} drafts)`,
    );
  }

  console.log("\nGenerating compliance MDX files...");

  for (const adapter of adapters) {
    const mdxContent = generateComplianceMdx(adapter);
    const outputPath = join(
      CONTENT_DIR,
      adapter.lang,
      "adapters",
      adapter.adapter,
      "compliance.mdx",
    );
    writeMdx(outputPath, mdxContent);
    console.log(`  Created: ${outputPath}`);
  }

  console.log("\nCompliance MDX generation complete!");
}

main();
