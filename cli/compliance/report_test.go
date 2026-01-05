package compliance

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestGenerateMarkdownReport(t *testing.T) {
	report := ComplianceReport{
		Adapter:     "@xschemadev/zod",
		GeneratedAt: "2025-01-05T12:00:00Z",
		Drafts: []DraftResult{
			{
				Draft: "draft2020-12",
				Keywords: []KeywordResult{
					{Keyword: "type", Passed: 10, Failed: 0, Skipped: 0, Total: 10},
					{Keyword: "const", Passed: 5, Failed: 2, Skipped: 0, Total: 7, Failures: []TestResult{
						{Group: "const validation", Test: "null is valid", Expected: true, Actual: "false", Passed: false},
						{Group: "const validation", Test: "another test", Expected: false, Actual: "error", Passed: false, Error: "parse error"},
					}},
				},
				Summary: DraftSummary{Passed: 15, Failed: 2, Skipped: 0, Total: 17, Percentage: 88.2},
			},
		},
	}

	md := GenerateMarkdownReport(report)

	// Check header
	if !strings.Contains(md, "# @xschemadev/zod Compliance Report") {
		t.Error("markdown missing report header")
	}

	// Check generated date
	if !strings.Contains(md, "2025-01-05T12:00:00Z") {
		t.Error("markdown missing generated date")
	}

	// Check summary table
	if !strings.Contains(md, "## Summary") {
		t.Error("markdown missing summary section")
	}
	if !strings.Contains(md, "| draft2020-12 | 15 | 2 | 0 | 88.2% |") {
		t.Error("markdown missing/incorrect summary row")
	}

	// Check badges
	if !strings.Contains(md, "## Badges") {
		t.Error("markdown missing badges section")
	}
	if !strings.Contains(md, "img.shields.io/badge") {
		t.Error("markdown missing badge URL")
	}

	// Check keyword table
	if !strings.Contains(md, "| type | ✅ | 10/10 |") {
		t.Error("markdown missing type keyword row")
	}
	if !strings.Contains(md, "| const | ⚠️ | 5/7 |") {
		t.Error("markdown missing const keyword row")
	}

	// Check failures section
	if !strings.Contains(md, "### Failures") {
		t.Error("markdown missing failures section")
	}
	if !strings.Contains(md, "<details>") {
		t.Error("markdown missing details element")
	}
	if !strings.Contains(md, "const - 2 failures") {
		t.Error("markdown missing failure count")
	}
	if !strings.Contains(md, "null is valid") {
		t.Error("markdown missing failure test name")
	}
	if !strings.Contains(md, "parse error") {
		t.Error("markdown missing error message")
	}
}

func TestGenerateMarkdownReport_NoFailures(t *testing.T) {
	report := ComplianceReport{
		Adapter:     "test-adapter",
		GeneratedAt: "2025-01-05T12:00:00Z",
		Drafts: []DraftResult{
			{
				Draft: "draft7",
				Keywords: []KeywordResult{
					{Keyword: "type", Passed: 10, Failed: 0, Total: 10},
				},
				Summary: DraftSummary{Passed: 10, Failed: 0, Total: 10, Percentage: 100.0},
			},
		},
	}

	md := GenerateMarkdownReport(report)

	// Should NOT have failures section
	if strings.Contains(md, "### Failures") {
		t.Error("markdown should not have failures section when no failures")
	}
}

func TestGenerateMarkdownReport_MultipleDrafts(t *testing.T) {
	report := ComplianceReport{
		Adapter:     "test-adapter",
		GeneratedAt: "2025-01-05T12:00:00Z",
		Drafts: []DraftResult{
			{
				Draft:    "draft2020-12",
				Keywords: []KeywordResult{{Keyword: "type", Passed: 10, Total: 10}},
				Summary:  DraftSummary{Passed: 10, Total: 10, Percentage: 100.0},
			},
			{
				Draft:    "draft7",
				Keywords: []KeywordResult{{Keyword: "type", Passed: 8, Total: 10}},
				Summary:  DraftSummary{Passed: 8, Failed: 2, Total: 10, Percentage: 80.0},
			},
		},
	}

	md := GenerateMarkdownReport(report)

	// Check both drafts in summary
	if !strings.Contains(md, "| draft2020-12 |") {
		t.Error("markdown missing draft2020-12 in summary")
	}
	if !strings.Contains(md, "| draft7 |") {
		t.Error("markdown missing draft7 in summary")
	}

	// Check both draft sections
	if !strings.Contains(md, "## draft2020-12") {
		t.Error("markdown missing draft2020-12 section")
	}
	if !strings.Contains(md, "## draft7") {
		t.Error("markdown missing draft7 section")
	}
}

func TestGenerateJSONReport(t *testing.T) {
	draftResult := DraftResult{
		Draft: "draft2020-12",
		Keywords: []KeywordResult{
			{
				Keyword: "type",
				Passed:  10,
				Failed:  2,
				Skipped: 0,
				Total:   12,
				Failures: []TestResult{
					{
						Group:    "test group",
						Test:     "test case",
						Expected: true,
						Actual:   "false",
						Passed:   false,
					},
				},
			},
		},
		Summary: DraftSummary{
			Passed:     10,
			Failed:     2,
			Total:      12,
			Percentage: 83.3,
		},
	}

	jsonStr, err := GenerateJSONReport(draftResult)
	if err != nil {
		t.Fatalf("GenerateJSONReport() error = %v", err)
	}

	// Verify it's valid JSON
	var parsed DraftResult
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		t.Errorf("GenerateJSONReport() produced invalid JSON: %v", err)
	}

	// Verify content
	if parsed.Draft != "draft2020-12" {
		t.Errorf("parsed.Draft = %s, want draft2020-12", parsed.Draft)
	}
	if len(parsed.Keywords) != 1 {
		t.Errorf("parsed.Keywords length = %d, want 1", len(parsed.Keywords))
	}
	if parsed.Summary.Percentage != 83.3 {
		t.Errorf("parsed.Summary.Percentage = %f, want 83.3", parsed.Summary.Percentage)
	}
}

func TestGenerateJSONReport_Formatting(t *testing.T) {
	draftResult := DraftResult{
		Draft:    "draft7",
		Keywords: []KeywordResult{},
		Summary:  DraftSummary{},
	}

	jsonStr, err := GenerateJSONReport(draftResult)
	if err != nil {
		t.Fatalf("GenerateJSONReport() error = %v", err)
	}

	// Should be indented (pretty printed)
	if !strings.Contains(jsonStr, "\n") {
		t.Error("JSON should be pretty printed with newlines")
	}
	if !strings.Contains(jsonStr, "  ") {
		t.Error("JSON should be indented with spaces")
	}
}

func TestGenerateBadgeURL(t *testing.T) {
	tests := []struct {
		name       string
		draft      string
		percentage float64
		wantColor  string
	}{
		{
			name:       "high percentage (>=95) is green",
			draft:      "draft2020-12",
			percentage: 95.0,
			wantColor:  "brightgreen",
		},
		{
			name:       "100% is green",
			draft:      "draft7",
			percentage: 100.0,
			wantColor:  "brightgreen",
		},
		{
			name:       "medium percentage (>=80) is yellow",
			draft:      "draft6",
			percentage: 80.0,
			wantColor:  "yellow",
		},
		{
			name:       "94.9% is yellow",
			draft:      "draft4",
			percentage: 94.9,
			wantColor:  "yellow",
		},
		{
			name:       "low percentage (<80) is red",
			draft:      "draft3",
			percentage: 79.9,
			wantColor:  "red",
		},
		{
			name:       "0% is red",
			draft:      "v1",
			percentage: 0.0,
			wantColor:  "red",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := generateBadgeURL(tt.draft, tt.percentage)

			if !strings.Contains(url, "img.shields.io/badge") {
				t.Error("URL should contain shields.io domain")
			}
			if !strings.HasSuffix(url, tt.wantColor) {
				t.Errorf("URL = %s, should end with color %s", url, tt.wantColor)
			}
			if !strings.Contains(url, tt.draft) {
				t.Errorf("URL should contain draft name %s", tt.draft)
			}
		})
	}
}

func TestGetStatusEmoji(t *testing.T) {
	tests := []struct {
		name    string
		keyword KeywordResult
		want    string
	}{
		{
			name:    "all passed",
			keyword: KeywordResult{Passed: 10, Failed: 0, Skipped: 0, Total: 10},
			want:    "✅",
		},
		{
			name:    "some failed",
			keyword: KeywordResult{Passed: 8, Failed: 2, Skipped: 0, Total: 10},
			want:    "⚠️",
		},
		{
			name:    "some skipped",
			keyword: KeywordResult{Passed: 8, Failed: 0, Skipped: 2, Total: 10},
			want:    "⚠️",
		},
		{
			name:    "all failed",
			keyword: KeywordResult{Passed: 0, Failed: 10, Skipped: 0, Total: 10},
			want:    "❌",
		},
		{
			name:    "none passed with skipped",
			keyword: KeywordResult{Passed: 0, Failed: 5, Skipped: 5, Total: 10},
			want:    "❌",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getStatusEmoji(tt.keyword)
			if got != tt.want {
				t.Errorf("getStatusEmoji() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestGetFailedKeywords(t *testing.T) {
	keywords := []KeywordResult{
		{Keyword: "type", Failures: nil},
		{Keyword: "const", Failures: []TestResult{{}}},
		{Keyword: "enum", Failures: nil},
		{Keyword: "format", Failures: []TestResult{{}, {}}},
	}

	failed := getFailedKeywords(keywords)

	if len(failed) != 2 {
		t.Errorf("getFailedKeywords() returned %d keywords, want 2", len(failed))
	}

	// Check correct keywords returned
	foundConst := false
	foundFormat := false
	for _, k := range failed {
		if k.Keyword == "const" {
			foundConst = true
		}
		if k.Keyword == "format" {
			foundFormat = true
		}
	}

	if !foundConst {
		t.Error("getFailedKeywords() missing 'const'")
	}
	if !foundFormat {
		t.Error("getFailedKeywords() missing 'format'")
	}
}

func TestGetFailedKeywords_Empty(t *testing.T) {
	keywords := []KeywordResult{
		{Keyword: "type", Failures: nil},
		{Keyword: "const", Failures: []TestResult{}},
	}

	failed := getFailedKeywords(keywords)

	if len(failed) != 0 {
		t.Errorf("getFailedKeywords() returned %d keywords, want 0", len(failed))
	}
}

func TestPluralize(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{0, "s"},
		{1, ""},
		{2, "s"},
		{10, "s"},
	}

	for _, tt := range tests {
		got := pluralize(tt.n)
		if got != tt.want {
			t.Errorf("pluralize(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestGenerateMarkdownReport_FailureDetails(t *testing.T) {
	report := ComplianceReport{
		Adapter:     "test",
		GeneratedAt: "2025-01-05",
		Drafts: []DraftResult{
			{
				Draft: "draft7",
				Keywords: []KeywordResult{
					{
						Keyword: "const",
						Passed:  1,
						Failed:  1,
						Total:   2,
						Failures: []TestResult{
							{
								Group:    "const validation",
								Test:     "wrong type of value",
								Expected: true,
								Actual:   "false",
								Passed:   false,
								Error:    "",
							},
						},
					},
				},
				Summary: DraftSummary{Passed: 1, Failed: 1, Total: 2, Percentage: 50.0},
			},
		},
	}

	md := GenerateMarkdownReport(report)

	// Check failure format
	if !strings.Contains(md, "Expected: `valid`, Got: `false`") {
		t.Error("markdown should show 'valid' for Expected: true")
	}
}

func TestGenerateMarkdownReport_FailureWithError(t *testing.T) {
	report := ComplianceReport{
		Adapter:     "test",
		GeneratedAt: "2025-01-05",
		Drafts: []DraftResult{
			{
				Draft: "draft7",
				Keywords: []KeywordResult{
					{
						Keyword: "type",
						Passed:  0,
						Failed:  1,
						Total:   1,
						Failures: []TestResult{
							{
								Group:    "type group",
								Test:     "error test",
								Expected: false,
								Actual:   "error",
								Error:    "schema parse failed",
							},
						},
					},
				},
				Summary: DraftSummary{Passed: 0, Failed: 1, Total: 1, Percentage: 0},
			},
		},
	}

	md := GenerateMarkdownReport(report)

	// Check error format
	if !strings.Contains(md, "Expected: `invalid`, Got: `error: schema parse failed`") {
		t.Error("markdown should show 'invalid' for Expected: false and include error")
	}
}

func TestGenerateMarkdownReport_SingleFailure(t *testing.T) {
	report := ComplianceReport{
		Adapter:     "test",
		GeneratedAt: "2025-01-05",
		Drafts: []DraftResult{
			{
				Draft: "draft7",
				Keywords: []KeywordResult{
					{
						Keyword:  "type",
						Passed:   9,
						Failed:   1,
						Total:    10,
						Failures: []TestResult{{}},
					},
				},
				Summary: DraftSummary{Passed: 9, Failed: 1, Total: 10, Percentage: 90.0},
			},
		},
	}

	md := GenerateMarkdownReport(report)

	// Should use singular "failure" not "failures"
	if !strings.Contains(md, "type - 1 failure</summary>") {
		t.Error("markdown should use singular 'failure' for count of 1")
	}
}
