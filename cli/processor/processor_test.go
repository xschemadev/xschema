package processor

import (
	"context"
	"encoding/json"
	"errors"
	"sync/atomic"
	"testing"

	"github.com/xschemadev/xschema/retriever"
)

// mockFetcher tracks fetch calls and returns predefined responses
type mockFetcher struct {
	responses  map[string]json.RawMessage
	errors     map[string]error
	fetchCalls []string
	callCount  int32
}

func newMockFetcher() *mockFetcher {
	return &mockFetcher{
		responses: make(map[string]json.RawMessage),
		errors:    make(map[string]error),
	}
}

func (m *mockFetcher) Fetch(ctx context.Context, uri string) (json.RawMessage, error) {
	atomic.AddInt32(&m.callCount, 1)
	m.fetchCalls = append(m.fetchCalls, uri)

	if err, ok := m.errors[uri]; ok {
		return nil, err
	}
	if data, ok := m.responses[uri]; ok {
		return data, nil
	}
	return nil, errors.New("no response configured for: " + uri)
}

func (m *mockFetcher) addResponse(uri string, schema string) {
	m.responses[uri] = json.RawMessage(schema)
}

func (m *mockFetcher) addError(uri string, err error) {
	m.errors[uri] = err
}

func TestCrawlAndFetch_NoExternalRefs(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"name": { "type": "string" },
			"count": { "$ref": "#/$defs/Counter" }
		},
		"$defs": {
			"Counter": { "type": "integer" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "schema1",
			Schema:    schema,
			SourceURI: "file:///test/schema1.json",
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(fetcher.fetchCalls) != 0 {
		t.Errorf("expected 0 fetches, got %d: %v", len(fetcher.fetchCalls), fetcher.fetchCalls)
	}
}

func TestCrawlAndFetch_SingleExternalRef(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	fetcher.addResponse("http://example.com/address.json", `{
		"type": "object",
		"properties": {
			"street": { "type": "string" },
			"city": { "type": "string" }
		}
	}`)

	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"address": { "$ref": "http://example.com/address.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "schema1",
			Schema:    schema,
			SourceURI: "http://test.com/schema1.json",
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(fetcher.fetchCalls) != 1 {
		t.Errorf("expected 1 fetch, got %d: %v", len(fetcher.fetchCalls), fetcher.fetchCalls)
	}
	if fetcher.fetchCalls[0] != "http://example.com/address.json" {
		t.Errorf("expected fetch for address.json, got %s", fetcher.fetchCalls[0])
	}
}

func TestCrawlAndFetch_ChainedRefs(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	// A refs B, B refs C
	fetcher.addResponse("http://example.com/b.json", `{
		"type": "object",
		"properties": {
			"c": { "$ref": "http://example.com/c.json" }
		}
	}`)
	fetcher.addResponse("http://example.com/c.json", `{
		"type": "string"
	}`)

	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"b": { "$ref": "http://example.com/b.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "a",
			Schema:    schema,
			SourceURI: "http://test.com/a.json",
		},
	}

	var iterations int
	_, err := Process(ctx, schemas, Options{
		Fetcher: fetcher,
		OnVerbose: func(msg string) {
			if len(msg) > 20 && msg[:20] == "processor: crawl ite" {
				iterations++
			}
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(fetcher.fetchCalls) != 2 {
		t.Errorf("expected 2 fetches (B and C), got %d: %v", len(fetcher.fetchCalls), fetcher.fetchCalls)
	}

	// Should have 2 iterations: first for B, second for C
	if iterations != 2 {
		t.Errorf("expected 2 crawl iterations, got %d", iterations)
	}
}

func TestCrawlAndFetch_CircularRefs(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	// A refs B, B refs A (circular)
	fetcher.addResponse("http://example.com/b.json", `{
		"type": "object",
		"properties": {
			"a": { "$ref": "http://example.com/a.json" }
		}
	}`)
	fetcher.addResponse("http://example.com/a.json", `{
		"type": "object",
		"properties": {
			"b": { "$ref": "http://example.com/b.json" }
		}
	}`)

	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"external": { "$ref": "http://example.com/b.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "main",
			Schema:    schema,
			SourceURI: "http://example.com/a.json", // declared at a.json location
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should only fetch b.json once, not enter infinite loop
	// a.json is already "visited" since it's the declared schema's source URI
	if len(fetcher.fetchCalls) != 1 {
		t.Errorf("expected 1 fetch (only B, A is declared), got %d: %v", len(fetcher.fetchCalls), fetcher.fetchCalls)
	}
	if fetcher.fetchCalls[0] != "http://example.com/b.json" {
		t.Errorf("expected fetch for b.json, got %s", fetcher.fetchCalls[0])
	}
}

func TestCrawlAndFetch_FragmentOnlyRefs(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"user": { "$ref": "#/$defs/User" },
			"role": { "$ref": "#Role" }
		},
		"$defs": {
			"User": {
				"$anchor": "User",
				"type": "object"
			},
			"Role": {
				"$anchor": "Role",
				"type": "string"
			}
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "schema1",
			Schema:    schema,
			SourceURI: "http://test.com/schema1.json",
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(fetcher.fetchCalls) != 0 {
		t.Errorf("fragment-only refs should not trigger fetches, got %d: %v", len(fetcher.fetchCalls), fetcher.fetchCalls)
	}
}

func TestCrawlAndFetch_RelativeRefWithBaseURIChange(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	// The schema changes base URI via $id, so relative ref should resolve against new base
	fetcher.addResponse("http://other.com/schemas/nested.json", `{"type": "string"}`)

	schema := json.RawMessage(`{
		"$id": "http://other.com/schemas/base.json",
		"type": "object",
		"properties": {
			"nested": { "$ref": "nested.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "schema1",
			Schema:    schema,
			SourceURI: "file:///local/schema1.json", // different from $id
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(fetcher.fetchCalls) != 1 {
		t.Errorf("expected 1 fetch, got %d: %v", len(fetcher.fetchCalls), fetcher.fetchCalls)
	}
	// Relative "nested.json" should resolve against $id "http://other.com/schemas/base.json"
	if fetcher.fetchCalls[0] != "http://other.com/schemas/nested.json" {
		t.Errorf("expected fetch for http://other.com/schemas/nested.json, got %s", fetcher.fetchCalls[0])
	}
}

func TestCrawlAndFetch_FetchError_FailFast(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	fetcher.addError("http://example.com/failing.json", errors.New("connection refused"))

	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"bad": { "$ref": "http://example.com/failing.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "schema1",
			Schema:    schema,
			SourceURI: "http://test.com/schema1.json",
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, errors.New("")) && err.Error() == "" {
		t.Errorf("expected non-empty error message, got empty")
	}
	// Error should mention the failing URI
	errStr := err.Error()
	if !contains(errStr, "failing.json") && !contains(errStr, "connection refused") {
		t.Errorf("error should mention failing URI or cause, got: %s", errStr)
	}
}

func TestCrawlAndFetch_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// cancellingFetcher cancels context after first fetch
	cancelOnce := false
	fetcher := FetchFunc(func(ctx context.Context, uri string) (json.RawMessage, error) {
		if cancelOnce {
			return json.RawMessage(`{"type": "string"}`), nil
		}
		cancelOnce = true
		cancel() // cancel after first fetch
		return json.RawMessage(`{
			"properties": {
				"next": { "$ref": "http://example.com/second.json" }
			}
		}`), nil
	})

	schema := json.RawMessage(`{
		"properties": {
			"first": { "$ref": "http://example.com/first.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "schema1",
			Schema:    schema,
			SourceURI: "http://test.com/schema1.json",
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

func TestCrawlAndFetch_RefWithFragment_FetchesBaseOnly(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	fetcher.addResponse("http://example.com/definitions.json", `{
		"$defs": {
			"Address": { "type": "object" },
			"User": { "type": "object" }
		}
	}`)

	// Schema refs same URL with different fragments - should only fetch once
	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"address": { "$ref": "http://example.com/definitions.json#/$defs/Address" },
			"user": { "$ref": "http://example.com/definitions.json#/$defs/User" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "schema1",
			Schema:    schema,
			SourceURI: "http://test.com/schema1.json",
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should fetch definitions.json only once (fragment stripped)
	if len(fetcher.fetchCalls) != 1 {
		t.Errorf("expected 1 fetch (fragments stripped), got %d: %v", len(fetcher.fetchCalls), fetcher.fetchCalls)
	}
	if fetcher.fetchCalls[0] != "http://example.com/definitions.json" {
		t.Errorf("expected fetch for definitions.json without fragment, got %s", fetcher.fetchCalls[0])
	}
}

func TestProcess_NilFetcher_ReturnsError(t *testing.T) {
	ctx := context.Background()

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "schema1",
			Schema:    json.RawMessage(`{"type": "string"}`),
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: nil})
	if err == nil {
		t.Fatal("expected error for nil Fetcher, got nil")
	}
	if !contains(err.Error(), "Fetcher is required") {
		t.Errorf("expected error about missing Fetcher, got: %v", err)
	}
}

func TestFetchFunc_ImplementsFetcher(t *testing.T) {
	var called bool
	fn := FetchFunc(func(ctx context.Context, uri string) (json.RawMessage, error) {
		called = true
		return json.RawMessage(`{"type":"string"}`), nil
	})

	result, err := fn.Fetch(context.Background(), "http://example.com/test.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !called {
		t.Error("FetchFunc was not called")
	}
	if string(result) != `{"type":"string"}` {
		t.Errorf("unexpected result: %s", result)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// --- CacheFetcher Tests (US-010) ---

func TestCacheFetcher_CacheHit(t *testing.T) {
	cache := externalRefCache{
		"http://example.com/schema.json": json.RawMessage(`{"type":"string"}`),
	}
	fetcher := NewCacheFetcher(cache)

	result, err := fetcher.Fetch("http://example.com/schema.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(result) != `{"type":"string"}` {
		t.Errorf("expected schema bytes, got: %s", result)
	}
}

func TestCacheFetcher_CacheMiss(t *testing.T) {
	cache := externalRefCache{
		"http://example.com/exists.json": json.RawMessage(`{"type":"string"}`),
	}
	fetcher := NewCacheFetcher(cache)

	_, err := fetcher.Fetch("http://example.com/missing.json")
	if err == nil {
		t.Fatal("expected error for cache miss, got nil")
	}
	errStr := err.Error()
	if !contains(errStr, "cache miss") {
		t.Errorf("error should mention 'cache miss', got: %s", errStr)
	}
	if !contains(errStr, "missing.json") {
		t.Errorf("error should mention the missing URI, got: %s", errStr)
	}
	if !contains(errStr, "bug") {
		t.Errorf("error should mention this indicates a bug, got: %s", errStr)
	}
}

func TestCacheFetcher_EmptyCache(t *testing.T) {
	cache := externalRefCache{}
	fetcher := NewCacheFetcher(cache)

	_, err := fetcher.Fetch("http://example.com/any.json")
	if err == nil {
		t.Fatal("expected error for empty cache, got nil")
	}
	if !contains(err.Error(), "cache miss") {
		t.Errorf("error should mention 'cache miss', got: %s", err)
	}
}

func TestCacheFetcher_CaseSensitivity(t *testing.T) {
	// URIs should be matched exactly - different case = different URI
	cache := externalRefCache{
		"http://Example.com/Schema.json": json.RawMessage(`{"type":"string"}`),
	}
	fetcher := NewCacheFetcher(cache)

	// Exact match should work
	result, err := fetcher.Fetch("http://Example.com/Schema.json")
	if err != nil {
		t.Fatalf("unexpected error for exact match: %v", err)
	}
	if string(result) != `{"type":"string"}` {
		t.Errorf("expected schema bytes, got: %s", result)
	}

	// Different case should NOT match (case sensitive)
	_, err = fetcher.Fetch("http://example.com/schema.json")
	if err == nil {
		t.Fatal("expected error for case-different URI, got nil (URIs should be case-sensitive)")
	}

	// Another case variation
	_, err = fetcher.Fetch("http://EXAMPLE.COM/SCHEMA.JSON")
	if err == nil {
		t.Fatal("expected error for uppercase URI, got nil (URIs should be case-sensitive)")
	}
}

func TestCacheFetcher_FragmentStripped(t *testing.T) {
	// CacheFetcher normalizes URI (strips fragment) before lookup
	cache := externalRefCache{
		"http://example.com/defs.json": json.RawMessage(`{"$defs":{"A":{"type":"string"}}}`),
	}
	fetcher := NewCacheFetcher(cache)

	// Fetch with fragment should find the base URI in cache
	result, err := fetcher.Fetch("http://example.com/defs.json#/$defs/A")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(result) != `{"$defs":{"A":{"type":"string"}}}` {
		t.Errorf("expected full schema (fragment stripped), got: %s", result)
	}
}
