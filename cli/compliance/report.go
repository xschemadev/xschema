package compliance

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// GenerateMarkdownReport creates a markdown report from compliance results
func GenerateMarkdownReport(report ComplianceReport) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s Compliance Report\n\n", report.Adapter))
	sb.WriteString(fmt.Sprintf("Generated: %s\n\n", report.GeneratedAt))

	// Summary table
	sb.WriteString("## Summary\n\n")
	sb.WriteString("| Draft | Passed | Failed | Skipped | Coverage |\n")
	sb.WriteString("| ----- | ------ | ------ | ------- | -------- |\n")

	for _, draft := range report.Drafts {
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
	for _, draft := range report.Drafts {
		badgeURL := generateBadgeURL(draft.Draft, draft.Summary.Percentage)
		sb.WriteString(fmt.Sprintf("![%s](%s)\n", draft.Draft, badgeURL))
	}
	sb.WriteString("\n")

	// Per-draft details
	for _, draft := range report.Drafts {
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
	data, err := json.MarshalIndent(draftResult, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
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
