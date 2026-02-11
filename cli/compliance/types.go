package compliance

import (
	"encoding/json"
	"time"
)

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

// Raw returns the schema as json.RawMessage for use with adapter.ConvertInput
func (s RawSchema) Raw() json.RawMessage {
	data, _ := json.Marshal(s.value)
	return data
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

// UnsupportedFeatureItem represents a single unsupported feature test with its reason
type UnsupportedFeatureItem struct {
	Path   string `json:"path"`
	Reason string `json:"reason"`
}

// UnsupportedFeaturesSummary contains unsupported features data for JSON output
type UnsupportedFeaturesSummary struct {
	Count int                      `json:"count"`
	Items []UnsupportedFeatureItem `json:"items"`
}

// DraftSummary contains aggregate statistics for a draft
type DraftSummary struct {
	Passed              int                        `json:"passed"`
	Failed              int                        `json:"failed"`
	Skipped             int                        `json:"skipped"`
	Total               int                        `json:"total"`
	Percentage          float64                    `json:"percentage"`
	UnsupportedFeatures UnsupportedFeaturesSummary `json:"unsupportedFeatures"`
}

// ComplianceReport is the complete report for an adapter
type ComplianceReport struct {
	Adapter string        `json:"adapter"`
	Drafts  []DraftResult `json:"drafts"`
}

// TimingSummary holds aggregate timings for a compliance run.
type TimingSummary struct {
	SuiteLoad         time.Duration
	SchemaBundling    time.Duration
	AdapterInvocation time.Duration
	HarnessGeneration time.Duration
	HarnessExecution  time.Duration
}

func (t *TimingSummary) addSuiteLoad(duration time.Duration) {
	if t == nil {
		return
	}
	t.SuiteLoad += duration
}

func (t *TimingSummary) addSchemaBundling(duration time.Duration) {
	if t == nil {
		return
	}
	t.SchemaBundling += duration
}

func (t *TimingSummary) addAdapterInvocation(duration time.Duration) {
	if t == nil {
		return
	}
	t.AdapterInvocation += duration
}

func (t *TimingSummary) addHarnessGeneration(duration time.Duration) {
	if t == nil {
		return
	}
	t.HarnessGeneration += duration
}

func (t *TimingSummary) addHarnessExecution(duration time.Duration) {
	if t == nil {
		return
	}
	t.HarnessExecution += duration
}

// HarnessResult is the JSON output from executing a harness file
type HarnessResult struct {
	GroupID  string `json:"groupId"`
	Index    int    `json:"index"`
	Expected bool   `json:"expected"`
	Actual   string `json:"actual"` // "true", "false", "skipped", or "error"
	Error    string `json:"error,omitempty"`
}

// BatchTestData is the test data structure passed to harness templates
type BatchTestData struct {
	GroupID string     `json:"groupId"`
	Tests   []TestCase `json:"tests"`
}
