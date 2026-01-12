package metaschema

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
)

func TestExtractVocabulary_MissingVocabulary(t *testing.T) {
	schema := map[string]any{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type":    "object",
	}

	result := ExtractVocabulary(schema)
	if result != nil {
		t.Errorf("expected nil for missing $vocabulary, got %v", result)
	}
}

func TestExtractVocabulary_EmptyVocabulary(t *testing.T) {
	schema := map[string]any{
		"$vocabulary": map[string]any{},
	}

	result := ExtractVocabulary(schema)
	if result == nil {
		t.Fatal("expected non-nil map for empty $vocabulary")
	}
	if len(result) != 0 {
		t.Errorf("expected empty map, got %v", result)
	}
}

func TestExtractVocabulary_PartialVocabulary(t *testing.T) {
	schema := map[string]any{
		"$vocabulary": map[string]any{
			"https://json-schema.org/draft/2020-12/vocab/core":       true,
			"https://json-schema.org/draft/2020-12/vocab/validation": false,
		},
	}

	result := ExtractVocabulary(schema)
	if result == nil {
		t.Fatal("expected non-nil map")
	}
	if len(result) != 2 {
		t.Errorf("expected 2 entries, got %d", len(result))
	}
	if !result["https://json-schema.org/draft/2020-12/vocab/core"] {
		t.Error("expected core vocab to be true (required)")
	}
	if result["https://json-schema.org/draft/2020-12/vocab/validation"] {
		t.Error("expected validation vocab to be false (optional)")
	}
}

func TestExtractVocabulary_NonObjectSchema(t *testing.T) {
	// schema can be a boolean in JSON Schema
	result := ExtractVocabulary(true)
	if result != nil {
		t.Errorf("expected nil for boolean schema, got %v", result)
	}
}

func TestExtractVocabulary_NonBooleanValues(t *testing.T) {
	schema := map[string]any{
		"$vocabulary": map[string]any{
			"https://example.com/valid":   true,
			"https://example.com/invalid": "not-a-bool",
			"https://example.com/number":  123,
		},
	}

	result := ExtractVocabulary(schema)
	if result == nil {
		t.Fatal("expected non-nil map")
	}
	// only boolean values should be included
	if len(result) != 1 {
		t.Errorf("expected 1 entry (only valid bool), got %d: %v", len(result), result)
	}
	if _, ok := result["https://example.com/valid"]; !ok {
		t.Error("expected valid vocab to be present")
	}
}

func TestIsStandardDraft_JsonSchemaOrg(t *testing.T) {
	tests := []struct {
		uri      string
		expected bool
	}{
		{"https://json-schema.org/draft/2020-12/schema", true},
		{"https://json-schema.org/draft/2019-09/schema", true},
		{"http://json-schema.org/draft-07/schema#", true},
		{"http://json-schema.org/draft-04/schema#", true},
		{"https://json-schema.org/draft/2020-12/meta/core", true},
	}

	for _, tc := range tests {
		t.Run(tc.uri, func(t *testing.T) {
			got := IsStandardDraft(tc.uri)
			if got != tc.expected {
				t.Errorf("IsStandardDraft(%q) = %v, want %v", tc.uri, got, tc.expected)
			}
		})
	}
}

func TestIsStandardDraft_CustomURLs(t *testing.T) {
	tests := []string{
		"https://example.com/my-schema",
		"https://my-company.com/schemas/v1/schema",
		"http://localhost:8080/schema.json",
		"file:///path/to/schema.json",
		"",
	}

	for _, uri := range tests {
		t.Run(uri, func(t *testing.T) {
			if IsStandardDraft(uri) {
				t.Errorf("IsStandardDraft(%q) = true, want false", uri)
			}
		})
	}
}

// mockRoundTripper allows mocking HTTP responses
type mockRoundTripper struct {
	response *http.Response
	err      error
	calls    atomic.Int32
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.calls.Add(1)
	return m.response, m.err
}

func TestGet_Caching(t *testing.T) {
	ClearCache()
	defer ClearCache()

	// save original client and restore after test
	origClient := HTTPClient
	defer func() { HTTPClient = origClient }()

	// mock HTTP response
	responseBody := `{"$id": "https://example.com/test", "$vocabulary": {"https://example.com/vocab": true}}`
	mock := &mockRoundTripper{
		response: &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(responseBody)),
		},
	}
	HTTPClient = &http.Client{Transport: mock}

	ctx := context.Background()
	uri := "https://example.com/test"

	// first call should fetch
	m1, err := Get(ctx, uri)
	if err != nil {
		t.Fatalf("first Get() failed: %v", err)
	}
	if m1 == nil {
		t.Fatal("first Get() returned nil")
	}
	if mock.calls.Load() != 1 {
		t.Errorf("expected 1 HTTP call, got %d", mock.calls.Load())
	}

	// second call should use cache (no additional HTTP call)
	// need to reset the response body since it was consumed
	mock.response.Body = io.NopCloser(strings.NewReader(responseBody))
	m2, err := Get(ctx, uri)
	if err != nil {
		t.Fatalf("second Get() failed: %v", err)
	}
	if m2 == nil {
		t.Fatal("second Get() returned nil")
	}
	if mock.calls.Load() != 1 {
		t.Errorf("expected still 1 HTTP call (cached), got %d", mock.calls.Load())
	}

	// verify it's the same cached instance
	if m1 != m2 {
		t.Error("expected same instance from cache")
	}

	// verify vocabulary was extracted
	if m1.Vocabulary == nil {
		t.Fatal("expected vocabulary to be extracted")
	}
	if !m1.Vocabulary["https://example.com/vocab"] {
		t.Error("expected vocab to be true")
	}
}

func TestGet_FetchError(t *testing.T) {
	ClearCache()
	defer ClearCache()

	origClient := HTTPClient
	defer func() { HTTPClient = origClient }()

	mock := &mockRoundTripper{
		response: &http.Response{
			StatusCode: 404,
			Body:       io.NopCloser(strings.NewReader("not found")),
		},
	}
	HTTPClient = &http.Client{Transport: mock}

	ctx := context.Background()
	_, err := Get(ctx, "https://example.com/not-found")
	if err == nil {
		t.Error("expected error for 404 response")
	}
}

func TestGet_InvalidJSON(t *testing.T) {
	ClearCache()
	defer ClearCache()

	origClient := HTTPClient
	defer func() { HTTPClient = origClient }()

	mock := &mockRoundTripper{
		response: &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader("not valid json {")),
		},
	}
	HTTPClient = &http.Client{Transport: mock}

	ctx := context.Background()
	_, err := Get(ctx, "https://example.com/invalid")
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestGet_ExtractsVocabulary(t *testing.T) {
	ClearCache()
	defer ClearCache()

	origClient := HTTPClient
	defer func() { HTTPClient = origClient }()

	responseBody := `{
		"$id": "https://example.com/metaschema",
		"$vocabulary": {
			"https://json-schema.org/draft/2020-12/vocab/core": true,
			"https://json-schema.org/draft/2020-12/vocab/applicator": true,
			"https://json-schema.org/draft/2020-12/vocab/validation": false
		}
	}`
	mock := &mockRoundTripper{
		response: &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(responseBody)),
		},
	}
	HTTPClient = &http.Client{Transport: mock}

	ctx := context.Background()
	m, err := Get(ctx, "https://example.com/metaschema")
	if err != nil {
		t.Fatalf("Get() failed: %v", err)
	}

	if m.Vocabulary == nil {
		t.Fatal("expected vocabulary to be extracted")
	}
	if len(m.Vocabulary) != 3 {
		t.Errorf("expected 3 vocabulary entries, got %d", len(m.Vocabulary))
	}
	if !m.Vocabulary["https://json-schema.org/draft/2020-12/vocab/core"] {
		t.Error("expected core vocab to be true")
	}
	if !m.Vocabulary["https://json-schema.org/draft/2020-12/vocab/applicator"] {
		t.Error("expected applicator vocab to be true")
	}
	if m.Vocabulary["https://json-schema.org/draft/2020-12/vocab/validation"] {
		t.Error("expected validation vocab to be false")
	}
}

func TestClearCache(t *testing.T) {
	ClearCache()
	defer ClearCache()

	origClient := HTTPClient
	defer func() { HTTPClient = origClient }()

	responseBody := `{"$id": "https://example.com/test"}`
	mock := &mockRoundTripper{
		response: &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(responseBody)),
		},
	}
	HTTPClient = &http.Client{Transport: mock}

	ctx := context.Background()
	uri := "https://example.com/test"

	// first call
	_, err := Get(ctx, uri)
	if err != nil {
		t.Fatalf("first Get() failed: %v", err)
	}
	if mock.calls.Load() != 1 {
		t.Errorf("expected 1 HTTP call, got %d", mock.calls.Load())
	}

	// clear cache
	ClearCache()

	// second call should fetch again
	mock.response.Body = io.NopCloser(strings.NewReader(responseBody))
	_, err = Get(ctx, uri)
	if err != nil {
		t.Fatalf("second Get() after clear failed: %v", err)
	}
	if mock.calls.Load() != 2 {
		t.Errorf("expected 2 HTTP calls after cache clear, got %d", mock.calls.Load())
	}
}
