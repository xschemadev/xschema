package compliance

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestRunDraftSequential_KeywordErrorsAreMarkedFailed(t *testing.T) {
	suite := map[string][]TestGroup{
		"ref": {
			{
				Description: "group",
				Schema:      mustRawSchema(t, `{"type":"string"}`),
				Tests: []TestCase{
					{Description: "valid", Data: "ok", Valid: true},
				},
			},
		},
	}

	keywords := []string{"ref"}
	opts := runDraftOptions{
		draft:      "draft2020-12",
		draftNum:   1,
		draftTotal: 1,
		suitePath:  t.TempDir(),
		adapterBin: "does-not-matter.js",
		runner:     "definitely-not-a-runner",
	}

	result, err := runDraftSequential(context.Background(), opts, suite, keywords, len(keywords))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assertKeywordFailureResult(t, result)
}

func TestRunDraftParallel_KeywordErrorsAreMarkedFailed(t *testing.T) {
	suite := map[string][]TestGroup{
		"ref": {
			{
				Description: "group",
				Schema:      mustRawSchema(t, `{"type":"string"}`),
				Tests: []TestCase{
					{Description: "valid", Data: "ok", Valid: true},
				},
			},
		},
	}

	keywords := []string{"ref"}
	opts := runDraftOptions{
		draft:      "draft2020-12",
		draftNum:   1,
		draftTotal: 1,
		suitePath:  t.TempDir(),
		adapterBin: "does-not-matter.js",
		runner:     "definitely-not-a-runner",
		jobs:       2,
	}

	result, err := runDraftParallel(context.Background(), opts, suite, keywords, len(keywords))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assertKeywordFailureResult(t, result)
}

func mustRawSchema(t *testing.T, schema string) RawSchema {
	t.Helper()

	var raw RawSchema
	if err := json.Unmarshal([]byte(schema), &raw); err != nil {
		t.Fatalf("failed to build RawSchema: %v", err)
	}

	return raw
}

func assertKeywordFailureResult(t *testing.T, result *DraftResult) {
	t.Helper()

	if result == nil {
		t.Fatal("expected draft result, got nil")
	}
	if len(result.Keywords) != 1 {
		t.Fatalf("expected 1 keyword result, got %d", len(result.Keywords))
	}

	keyword := result.Keywords[0]
	if keyword.Total != 1 {
		t.Fatalf("expected total=1, got %d", keyword.Total)
	}
	if keyword.Failed != 1 {
		t.Fatalf("expected failed=1, got %d", keyword.Failed)
	}
	if keyword.Passed != 0 {
		t.Fatalf("expected passed=0, got %d", keyword.Passed)
	}
	if len(keyword.Failures) != 1 {
		t.Fatalf("expected 1 failure item, got %d", len(keyword.Failures))
	}
	if !strings.Contains(keyword.Failures[0].Error, "adapter call failed") {
		t.Fatalf("expected adapter call failure message, got: %s", keyword.Failures[0].Error)
	}

	if result.Summary.Total != 1 {
		t.Fatalf("expected summary total=1, got %d", result.Summary.Total)
	}
	if result.Summary.Failed != 1 {
		t.Fatalf("expected summary failed=1, got %d", result.Summary.Failed)
	}
}
