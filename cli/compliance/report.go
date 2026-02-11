package compliance

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// GenerateMarkdownReport creates a markdown report from compliance results
func GenerateMarkdownReport(report ComplianceReport) string {
	sortedReport := sortComplianceReport(report)
	var sb strings.Builder

	fmt.Fprintf(&sb, "# %s Compliance Report\n\n", sortedReport.Adapter)

	// Detect type-only adapter (all tests skipped)
	isTypeOnly := isTypeOnlyAdapter(sortedReport)
	if isTypeOnly {
		sb.WriteString("## Type-Only Adapter\n\n")
		sb.WriteString("This adapter generates **type definitions only** - no runtime validation code is produced.\n\n")
		sb.WriteString("Runtime validation tests are **skipped** because:\n")
		sb.WriteString("1. Compliance tests require runtime validation (checking if data matches a schema)\n")
		sb.WriteString("2. Types exist only at compile-time and are erased at runtime\n")
		sb.WriteString("3. Static type correctness should be verified separately via the language's type checker\n\n")
		sb.WriteString("**Note:** Any failures shown below indicate compilation errors in the generated types, not runtime validation failures.\n\n")
	}

	// Summary table
	sb.WriteString("## Summary\n\n")
	sb.WriteString("| Draft | Passed | Failed | Skipped | Unsupported | Coverage |\n")
	sb.WriteString("| ----- | ------ | ------ | ------- | ----------- | -------- |\n")

	for _, draft := range sortedReport.Drafts {
		coverage := "N/A (type-only)"
		if !isTypeOnly {
			coverage = fmt.Sprintf("%.1f%%", draft.Summary.Percentage)
		}
		fmt.Fprintf(&sb, "| %s | %d | %d | %d | %d | %s |\n",
			draft.Draft,
			draft.Summary.Passed,
			draft.Summary.Failed,
			draft.Summary.Skipped,
			draft.Summary.UnsupportedFeatures.Count,
			coverage)
	}
	sb.WriteString("\n")

	// Badges (skip for type-only adapters - no meaningful percentage)
	if !isTypeOnly {
		sb.WriteString("## Badges\n\n")
		for _, draft := range sortedReport.Drafts {
			badgeURL := generateBadgeURL(draft.Draft, draft.Summary.Percentage)
			fmt.Fprintf(&sb, "![%s](%s)\n", draft.Draft, badgeURL)
		}
		sb.WriteString("\n")
	}

	// Per-draft details
	for _, draft := range sortedReport.Drafts {
		fmt.Fprintf(&sb, "## %s\n\n", draft.Draft)
		sb.WriteString("| Keyword | Status | Pass/Total |\n")
		sb.WriteString("| ------- | ------ | ---------- |\n")

		for _, keyword := range draft.Keywords {
			status := getStatusEmoji(keyword)
			fmt.Fprintf(&sb, "| %s | %s | %d/%d |\n",
				keyword.Keyword, status, keyword.Passed, keyword.Total)
		}
		sb.WriteString("\n")

		// Unsupported Features section (grouped by reason)
		if draft.Summary.UnsupportedFeatures.Count > 0 {
			sb.WriteString("### Unsupported Features\n\n")
			sb.WriteString("These tests are intentionally excluded due to documented limitations.\n\n")

			byReason := groupUnsupportedFeaturesByReason(draft.Summary.UnsupportedFeatures.Items)
			for _, group := range byReason {
				fmt.Fprintf(&sb, "<details>\n<summary>%s (%d test%s)</summary>\n\n",
					group.Reason, len(group.Tests), pluralize(len(group.Tests)))
				for _, path := range group.Tests {
					fmt.Fprintf(&sb, "- `%s`\n", path)
				}
				sb.WriteString("\n</details>\n\n")
			}
		}

		// Unexpected Failures section
		failedKeywords := getFailedKeywords(draft.Keywords)
		if len(failedKeywords) > 0 {
			sb.WriteString("### Unexpected Failures\n\n")
			for _, keyword := range failedKeywords {
				fmt.Fprintf(&sb, "<details>\n<summary>%s - %d failure%s</summary>\n\n",
					keyword.Keyword, len(keyword.Failures), pluralize(len(keyword.Failures)))

				for _, failure := range keyword.Failures {
					expected := "invalid"
					if failure.Expected {
						expected = "valid"
					}
					got := failure.Actual
					if failure.Error != "" {
						got = fmt.Sprintf("error: %s", failure.Error)
					}

					fmt.Fprintf(&sb, "- **%s**\n", failure.Group)
					fmt.Fprintf(&sb, "  - Test: %s\n", failure.Test)
					fmt.Fprintf(&sb, "  - Expected: `%s`, Got: `%s`\n", expected, got)
				}
				sb.WriteString("\n</details>\n\n")
			}
		}
	}

	return sb.String()
}

// GenerateJSONReport creates a JSON report for a single draft
func GenerateJSONReport(draftResult DraftResult) (string, error) {
	sortedDraft := sortDraftResult(draftResult)
	data, err := json.MarshalIndent(sortedDraft, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func sortComplianceReport(report ComplianceReport) ComplianceReport {
	sorted := ComplianceReport{
		Adapter: report.Adapter,
		Drafts:  make([]DraftResult, len(report.Drafts)),
	}

	for i, draft := range report.Drafts {
		sorted.Drafts[i] = sortDraftResult(draft)
	}

	sort.Slice(sorted.Drafts, func(i, j int) bool {
		return sorted.Drafts[i].Draft < sorted.Drafts[j].Draft
	})

	return sorted
}

func sortDraftResult(draft DraftResult) DraftResult {
	sorted := DraftResult{
		Draft:    draft.Draft,
		Summary:  draft.Summary,
		Keywords: make([]KeywordResult, len(draft.Keywords)),
	}

	for i, keyword := range draft.Keywords {
		sortedKeyword := keyword
		if len(keyword.Failures) > 0 {
			failures := make([]TestResult, len(keyword.Failures))
			copy(failures, keyword.Failures)
			sort.Slice(failures, func(i, j int) bool {
				return lessTestResult(failures[i], failures[j])
			})
			sortedKeyword.Failures = failures
		}
		sorted.Keywords[i] = sortedKeyword
	}

	sort.Slice(sorted.Keywords, func(i, j int) bool {
		return sorted.Keywords[i].Keyword < sorted.Keywords[j].Keyword
	})

	return sorted
}

func lessTestResult(left, right TestResult) bool {
	if left.Group != right.Group {
		return left.Group < right.Group
	}
	if left.Test != right.Test {
		return left.Test < right.Test
	}
	if left.Expected != right.Expected {
		return !left.Expected && right.Expected
	}
	if left.Actual != right.Actual {
		return left.Actual < right.Actual
	}
	return left.Error < right.Error
}

func generateBadgeURL(draft string, percentage float64) string {
	color := "red"
	if percentage >= 95 {
		color = "brightgreen"
	} else if percentage >= 80 {
		color = "yellow"
	}

	label := url.PathEscape(draft + " compliance")
	message := url.PathEscape(fmt.Sprintf("%.1f%%", percentage))

	return fmt.Sprintf("https://img.shields.io/badge/%s-%s-%s", label, message, color)
}

func getStatusEmoji(keyword KeywordResult) string {
	if keyword.Failed == 0 && keyword.Skipped == 0 {
		return "✅"
	}
	if keyword.Passed == 0 {
		return "❌"
	}
	return "⚠️"
}

func getFailedKeywords(keywords []KeywordResult) []KeywordResult {
	var failed []KeywordResult
	for _, k := range keywords {
		if len(k.Failures) > 0 {
			failed = append(failed, k)
		}
	}
	return failed
}

func pluralize(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

// reasonGroup holds unsupported feature paths grouped by reason
type reasonGroup struct {
	Reason string
	Tests  []string
}

// groupUnsupportedFeaturesByReason groups unsupported feature items by their reason
func groupUnsupportedFeaturesByReason(items []UnsupportedFeatureItem) []reasonGroup {
	byReason := make(map[string][]string)
	var order []string

	for _, item := range items {
		if _, exists := byReason[item.Reason]; !exists {
			order = append(order, item.Reason)
		}
		byReason[item.Reason] = append(byReason[item.Reason], item.Path)
	}

	// Sort reasons alphabetically for deterministic output
	sort.Strings(order)

	var result []reasonGroup
	for _, reason := range order {
		paths := byReason[reason]
		sort.Strings(paths)
		result = append(result, reasonGroup{Reason: reason, Tests: paths})
	}
	return result
}

// isTypeOnlyAdapter returns true if the adapter produced type-only output.
// Type-only adapters have no passed tests (no runtime validation) and at least some skipped tests.
// Failures may occur due to typecheck errors in the generated types.
func isTypeOnlyAdapter(report ComplianceReport) bool {
	if len(report.Drafts) == 0 {
		return false
	}

	totalPassed := 0
	totalSkipped := 0

	for _, draft := range report.Drafts {
		totalPassed += draft.Summary.Passed
		totalSkipped += draft.Summary.Skipped
	}

	// Type-only: no runtime tests passed, but some were skipped
	return totalPassed == 0 && totalSkipped > 0
}
