package compliance

import "encoding/json"

// TestCase represents a single test case from the JSON Schema Test Suite
type TestCase struct {
	Description string `json:"description"`
	Data        any    `json:"data"`
	Valid       bool   `json:"valid"`
}

// TestGroup represents a group of test cases sharing the same schema
type TestGroup struct {
	Description string     `json:"description"`
	Schema      RawSchema  `json:"schema"`
	Tests       []TestCase `json:"tests"`
}

// RawSchema is a raw JSON value representing a JSON Schema
// Can be a boolean (true/false) or an object
type RawSchema struct {
	value any
}

func (s *RawSchema) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.value)
}

func (s RawSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

func (s RawSchema) Value() any {
	return s.value
}

// TestResult represents the result of running a single test case
type TestResult struct {
	Group    string `json:"group"`
	Test     string `json:"test"`
	Expected bool   `json:"expected"`
	Actual   string `json:"actual"` // "true", "false", or "error"
	Passed   bool   `json:"passed"`
	Error    string `json:"error,omitempty"`
}

// KeywordResult contains results for a single keyword
type KeywordResult struct {
	Keyword  string       `json:"keyword"`
	Passed   int          `json:"passed"`
	Failed   int          `json:"failed"`
	Skipped  int          `json:"skipped"`
	Total    int          `json:"total"`
	Failures []TestResult `json:"failures,omitempty"`
}

// DraftResult contains results for a single draft version
type DraftResult struct {
	Draft    string          `json:"draft"`
	Keywords []KeywordResult `json:"keywords"`
	Summary  DraftSummary    `json:"summary"`
}

// DraftSummary contains aggregate statistics for a draft
type DraftSummary struct {
	Passed     int     `json:"passed"`
	Failed     int     `json:"failed"`
	Skipped    int     `json:"skipped"`
	Total      int     `json:"total"`
	Percentage float64 `json:"percentage"`
}

// ComplianceReport is the complete report for an adapter
type ComplianceReport struct {
	Adapter     string        `json:"adapter"`
	GeneratedAt string        `json:"generatedAt"`
	Drafts      []DraftResult `json:"drafts"`
}

// HarnessResult is the JSON output from executing a harness file
type HarnessResult struct {
	Index    int    `json:"index"`
	Expected bool   `json:"expected"`
	Actual   string `json:"actual"` // "true", "false", or "error"
	Error    string `json:"error,omitempty"`
}

// AdapterOutput is the response from calling an adapter's convert function
type AdapterOutput struct {
	Namespace       string   `json:"namespace"`
	ID              string   `json:"id"`
	Schema          string   `json:"schema"`
	Type            string   `json:"type"`
	Imports         []string `json:"imports"`
	Validate        string   `json:"validate,omitempty"`        // validation function in target language (empty = type-only adapter)
	ValidateImports []string `json:"validateImports,omitempty"` // imports needed by validate function
}
