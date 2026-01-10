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

	sb.WriteString(fmt.Sprintf("# %s Compliance Report\n\n", sortedReport.Adapter))

	// Summary table
	sb.WriteString("## Summary\n\n")
	sb.WriteString("| Draft | Passed | Failed | Skipped | Coverage |\n")
	sb.WriteString("| ----- | ------ | ------ | ------- | -------- |\n")

	for _, draft := range sortedReport.Drafts {
		sb.WriteString(fmt.Sprintf("| %s | %d | %d | %d | %.1f%% |\n",
			draft.Draft,
			draft.Summary.Passed,
			draft.Summary.Failed,
			draft.Summary.Skipped,
			draft.Summary.Percentage,
		))
	}
	sb.WriteString("\n")

	// Badges
	sb.WriteString("## Badges\n\n")
	for _, draft := range sortedReport.Drafts {
		badgeURL := generateBadgeURL(draft.Draft, draft.Summary.Percentage)
		sb.WriteString(fmt.Sprintf("![%s](%s)\n", draft.Draft, badgeURL))
	}
	sb.WriteString("\n")

	// Per-draft details
	for _, draft := range sortedReport.Drafts {
		sb.WriteString(fmt.Sprintf("## %s\n\n", draft.Draft))
		sb.WriteString("| Keyword | Status | Pass/Total |\n")
		sb.WriteString("| ------- | ------ | ---------- |\n")

		for _, keyword := range draft.Keywords {
			status := getStatusEmoji(keyword)
			sb.WriteString(fmt.Sprintf("| %s | %s | %d/%d |\n",
				keyword.Keyword, status, keyword.Passed, keyword.Total))
		}
		sb.WriteString("\n")

		// Failures details
		failedKeywords := getFailedKeywords(draft.Keywords)
		if len(failedKeywords) > 0 {
			sb.WriteString("### Failures\n\n")
			for _, keyword := range failedKeywords {
				sb.WriteString(fmt.Sprintf("<details>\n<summary>%s - %d failure%s</summary>\n\n",
					keyword.Keyword, len(keyword.Failures), pluralize(len(keyword.Failures))))

				for _, failure := range keyword.Failures {
					expected := "invalid"
					if failure.Expected {
						expected = "valid"
					}
					got := failure.Actual
					if failure.Error != "" {
						got = fmt.Sprintf("error: %s", failure.Error)
					}

					sb.WriteString(fmt.Sprintf("- **%s**\n", failure.Group))
					sb.WriteString(fmt.Sprintf("  - Test: %s\n", failure.Test))
					sb.WriteString(fmt.Sprintf("  - Expected: `%s`, Got: `%s`\n", expected, got))
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
