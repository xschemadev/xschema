package processor

import (
	"context"
	"encoding/json"
	"errors"
	"sync/atomic"
	"testing"

	"github.com/xschemadev/xschema/fetcher"
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
	fetcher := fetcher.FetchFunc(func(ctx context.Context, uri string) (json.RawMessage, error) {
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
	fn := fetcher.FetchFunc(func(ctx context.Context, uri string) (json.RawMessage, error) {
		called = true
		return json.RawMessage(`{"type":"string"}`), nil
	})

	result, err := fn.Fetch(context.Background(), "http://example.com/test.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !called {
		t.Error("fetcher.FetchFunc was not called")
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
	cache := fetcher.Cache{
		"http://example.com/schema.json": json.RawMessage(`{"type":"string"}`),
	}
	f := fetcher.NewCacheFetcher(cache)

	result, err := f.Fetch(context.Background(), "http://example.com/schema.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(result) != `{"type":"string"}` {
		t.Errorf("expected schema bytes, got: %s", result)
	}
}

func TestCacheFetcher_CacheMiss(t *testing.T) {
	cache := fetcher.Cache{
		"http://example.com/exists.json": json.RawMessage(`{"type":"string"}`),
	}
	f := fetcher.NewCacheFetcher(cache)

	_, err := f.Fetch(context.Background(), "http://example.com/missing.json")
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
}

func TestCacheFetcher_EmptyCache(t *testing.T) {
	cache := fetcher.Cache{}
	f := fetcher.NewCacheFetcher(cache)

	_, err := f.Fetch(context.Background(), "http://example.com/any.json")
	if err == nil {
		t.Fatal("expected error for empty cache, got nil")
	}
	if !contains(err.Error(), "cache miss") {
		t.Errorf("error should mention 'cache miss', got: %s", err)
	}
}

func TestCacheFetcher_CaseSensitivity(t *testing.T) {
	// URIs should be matched exactly - different case = different URI
	cache := fetcher.Cache{
		"http://Example.com/Schema.json": json.RawMessage(`{"type":"string"}`),
	}
	f := fetcher.NewCacheFetcher(cache)

	// Exact match should work
	result, err := f.Fetch(context.Background(), "http://Example.com/Schema.json")
	if err != nil {
		t.Fatalf("unexpected error for exact match: %v", err)
	}
	if string(result) != `{"type":"string"}` {
		t.Errorf("expected schema bytes, got: %s", result)
	}

	// Different case should NOT match (case sensitive)
	_, err = f.Fetch(context.Background(), "http://example.com/schema.json")
	if err == nil {
		t.Fatal("expected error for case-different URI, got nil (URIs should be case-sensitive)")
	}

	// Another case variation
	_, err = f.Fetch(context.Background(), "http://EXAMPLE.COM/SCHEMA.JSON")
	if err == nil {
		t.Fatal("expected error for uppercase URI, got nil (URIs should be case-sensitive)")
	}
}

func TestCacheFetcher_FragmentStripped(t *testing.T) {
	// CacheFetcher normalizes URI (strips fragment) before lookup
	cache := fetcher.Cache{
		"http://example.com/defs.json": json.RawMessage(`{"$defs":{"A":{"type":"string"}}}`),
	}
	f := fetcher.NewCacheFetcher(cache)

	// Fetch with fragment should find the base URI in cache
	result, err := f.Fetch(context.Background(), "http://example.com/defs.json#/$defs/A")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(result) != `{"$defs":{"A":{"type":"string"}}}` {
		t.Errorf("expected full schema (fragment stripped), got: %s", result)
	}
}

// --- Process() Full Pipeline Tests (US-013) ---

func TestProcess_NoRefs_ReturnsBundledUnchanged(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"name": { "type": "string" },
			"age": { "type": "integer" }
		},
		"required": ["name"]
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "simple",
			Schema:    schema,
			Adapter:   "zod",
			SourceURI: "http://test.com/simple.json",
		},
	}

	result, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	p := result[0]
	if p.Namespace != "test" {
		t.Errorf("expected namespace 'test', got %q", p.Namespace)
	}
	if p.ID != "simple" {
		t.Errorf("expected ID 'simple', got %q", p.ID)
	}
	if p.Adapter != "zod" {
		t.Errorf("expected adapter 'zod', got %q", p.Adapter)
	}
	if p.SourceURI != "http://test.com/simple.json" {
		t.Errorf("expected SourceURI 'http://test.com/simple.json', got %q", p.SourceURI)
	}

	// Schema should be essentially unchanged (bundled but same structure)
	var bundled map[string]any
	if err := json.Unmarshal(p.Schema, &bundled); err != nil {
		t.Fatalf("failed to parse bundled schema: %v", err)
	}
	if bundled["type"] != "object" {
		t.Errorf("bundled schema should preserve type:object")
	}
	props := bundled["properties"].(map[string]any)
	if props["name"] == nil || props["age"] == nil {
		t.Errorf("bundled schema should preserve properties")
	}

	// No fetches should have occurred
	if len(fetcher.fetchCalls) != 0 {
		t.Errorf("expected 0 fetches for schema with no refs, got %d", len(fetcher.fetchCalls))
	}
}

func TestProcess_ExternalRefs_BundledWithRefsEmbedded(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	// External schema to be fetched
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
			"name": { "type": "string" },
			"address": { "$ref": "http://example.com/address.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "person",
			Schema:    schema,
			Adapter:   "zod",
			SourceURI: "http://test.com/person.json",
		},
	}

	result, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	// Bundled schema should have the external ref embedded in $defs
	var bundled map[string]any
	if err := json.Unmarshal(result[0].Schema, &bundled); err != nil {
		t.Fatalf("failed to parse bundled schema: %v", err)
	}

	// Check that $defs was created with the embedded external schema
	defs, hasDefs := bundled["$defs"].(map[string]any)
	if !hasDefs {
		t.Fatal("bundled schema should have $defs with embedded external schema")
	}

	// There should be at least one embedded definition
	if len(defs) == 0 {
		t.Error("$defs should contain the embedded address schema")
	}

	// Check that the $ref was rewritten to point to $defs
	props := bundled["properties"].(map[string]any)
	addressProp := props["address"].(map[string]any)
	ref, hasRef := addressProp["$ref"].(string)
	if !hasRef {
		t.Fatal("address property should have $ref")
	}
	if !contains(ref, "#/$defs/") {
		t.Errorf("$ref should be rewritten to local $defs path, got %q", ref)
	}

	// Fetcher should have been called for address.json
	if len(fetcher.fetchCalls) != 1 {
		t.Errorf("expected 1 fetch for address.json, got %d: %v", len(fetcher.fetchCalls), fetcher.fetchCalls)
	}
}

func TestProcess_InvalidSchema_ValidationError(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	// Invalid schema: type is an array but should be string or array of strings
	// Actually, JSON Schema allows type to be string or array of strings.
	// Let's use a truly invalid schema - an invalid $schema reference
	schema := json.RawMessage(`{
		"type": 12345
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "invalid",
			Schema:    schema,
			Adapter:   "zod",
			SourceURI: "http://test.com/invalid.json",
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err == nil {
		t.Fatal("expected validation error for invalid schema, got nil")
	}

	// Error should mention validation failed
	errStr := err.Error()
	if !contains(errStr, "validation") || !contains(errStr, "failed") {
		t.Errorf("error should mention validation failed, got: %s", errStr)
	}
}

func TestProcess_InvalidExternalSchema_ValidationError(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	// External schema with invalid type
	fetcher.addResponse("http://example.com/bad.json", `{
		"type": ["not", "valid", "type", "list", 123]
	}`)

	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"bad": { "$ref": "http://example.com/bad.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "hasbad",
			Schema:    schema,
			Adapter:   "zod",
			SourceURI: "http://test.com/hasbad.json",
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err == nil {
		t.Fatal("expected validation error for invalid external schema, got nil")
	}

	// Error should mention external schema and validation
	errStr := err.Error()
	if !contains(errStr, "validation") && !contains(errStr, "external") {
		t.Errorf("error should mention validation of external schema, got: %s", errStr)
	}
}

func TestProcess_CustomMetaschema_FiltersVocabulary(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	// Schema with embedded $vocabulary that disables validation vocab.
	// Validation keywords like minLength should be stripped from output.
	schema := json.RawMessage(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"$vocabulary": {
			"https://json-schema.org/draft/2020-12/vocab/core": true,
			"https://json-schema.org/draft/2020-12/vocab/applicator": true
		},
		"type": "string",
		"minLength": 5,
		"format": "email"
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "customvocab",
			Schema:    schema,
			Adapter:   "zod",
			SourceURI: "http://test.com/customvocab.json",
		},
	}

	result, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	// Parse the output schema
	var outputSchema map[string]any
	if err := json.Unmarshal(result[0].Schema, &outputSchema); err != nil {
		t.Fatalf("failed to parse output schema: %v", err)
	}

	// type should be preserved (not a validation keyword in our filter)
	if outputSchema["type"] != "string" {
		t.Errorf("type should be preserved, got %v", outputSchema["type"])
	}

	// minLength should be stripped (validation vocab not in $vocabulary)
	if _, ok := outputSchema["minLength"]; ok {
		t.Error("minLength should be stripped when validation vocab is disabled")
	}

	// format should be stripped (format vocab not in $vocabulary)
	if _, ok := outputSchema["format"]; ok {
		t.Error("format should be stripped when format vocab is disabled")
	}

	// $vocabulary and $schema should still be present
	if outputSchema["$schema"] == nil {
		t.Error("$schema should be preserved")
	}
	if outputSchema["$vocabulary"] == nil {
		t.Error("$vocabulary should be preserved")
	}
}

func TestProcess_MultipleSchemas_AllProcessed(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	// External schema used by schema2
	fetcher.addResponse("http://example.com/shared.json", `{"type": "integer"}`)

	schema1 := json.RawMessage(`{
		"type": "string",
		"minLength": 1
	}`)
	schema2 := json.RawMessage(`{
		"type": "object",
		"properties": {
			"count": { "$ref": "http://example.com/shared.json" }
		}
	}`)
	schema3 := json.RawMessage(`{
		"type": "array",
		"items": { "type": "boolean" }
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "ns1",
			ID:        "schema1",
			Schema:    schema1,
			Adapter:   "zod",
			SourceURI: "http://test.com/schema1.json",
		},
		{
			Namespace: "ns2",
			ID:        "schema2",
			Schema:    schema2,
			Adapter:   "effect",
			SourceURI: "http://test.com/schema2.json",
		},
		{
			Namespace: "ns3",
			ID:        "schema3",
			Schema:    schema3,
			Adapter:   "arktype",
			SourceURI: "http://test.com/schema3.json",
		},
	}

	result, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 3 {
		t.Fatalf("expected 3 results, got %d", len(result))
	}

	// Check each result
	expected := []struct {
		namespace string
		id        string
		adapter   string
	}{
		{"ns1", "schema1", "zod"},
		{"ns2", "schema2", "effect"},
		{"ns3", "schema3", "arktype"},
	}

	for i, exp := range expected {
		if result[i].Namespace != exp.namespace {
			t.Errorf("result[%d].Namespace = %q, want %q", i, result[i].Namespace, exp.namespace)
		}
		if result[i].ID != exp.id {
			t.Errorf("result[%d].ID = %q, want %q", i, result[i].ID, exp.id)
		}
		if result[i].Adapter != exp.adapter {
			t.Errorf("result[%d].Adapter = %q, want %q", i, result[i].Adapter, exp.adapter)
		}
	}

	// shared.json should be fetched exactly once (for schema2)
	if len(fetcher.fetchCalls) != 1 {
		t.Errorf("expected 1 fetch for shared.json, got %d: %v", len(fetcher.fetchCalls), fetcher.fetchCalls)
	}
}

func TestProcess_VerboseCallback_ReceivesProgressUpdates(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	fetcher.addResponse("http://example.com/ext.json", `{"type": "string"}`)

	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"ext": { "$ref": "http://example.com/ext.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "verbose",
			Schema:    schema,
			Adapter:   "zod",
			SourceURI: "http://test.com/verbose.json",
		},
	}

	var messages []string
	_, err := Process(ctx, schemas, Options{
		Fetcher: fetcher,
		OnVerbose: func(msg string) {
			messages = append(messages, msg)
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should receive various progress updates
	if len(messages) == 0 {
		t.Fatal("expected verbose callback to receive messages")
	}

	// Check for expected message patterns
	hasStarting := false
	hasCrawlIteration := false
	hasCrawlComplete := false
	hasValidation := false
	hasBundling := false

	for _, msg := range messages {
		if contains(msg, "starting") {
			hasStarting = true
		}
		if contains(msg, "crawl iteration") {
			hasCrawlIteration = true
		}
		if contains(msg, "crawl complete") || contains(msg, "crawl finished") {
			hasCrawlComplete = true
		}
		if contains(msg, "validated") || contains(msg, "validation") {
			hasValidation = true
		}
		if contains(msg, "bundling") {
			hasBundling = true
		}
	}

	if !hasStarting {
		t.Error("verbose should include starting message")
	}
	if !hasCrawlIteration {
		t.Error("verbose should include crawl iteration message")
	}
	if !hasCrawlComplete {
		t.Error("verbose should include crawl complete message")
	}
	if !hasValidation {
		t.Error("verbose should include validation message")
	}
	if !hasBundling {
		t.Error("verbose should include bundling message")
	}
}

func TestProcess_CustomMetaschema_FetchedViaDollarSchema(t *testing.T) {
	ctx := context.Background()
	fetcher := newMockFetcher()

	// Custom metaschema (self-contained, no external refs)
	fetcher.addResponse("http://example.com/custom-meta.json", `{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"$id": "http://example.com/custom-meta.json",
		"$vocabulary": {
			"https://json-schema.org/draft/2020-12/vocab/core": true,
			"https://json-schema.org/draft/2020-12/vocab/applicator": true
		}
	}`)

	// Schema using custom metaschema - the $schema URL should be fetched
	schema := json.RawMessage(`{
		"$schema": "http://example.com/custom-meta.json",
		"type": "string",
		"minLength": 5
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "custom-meta-test",
			Schema:    schema,
			Adapter:   "zod",
			SourceURI: "http://test.com/custom-meta-test.json",
		},
	}

	result, err := Process(ctx, schemas, Options{Fetcher: fetcher})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	// Crawl should have fetched the custom metaschema via $schema URL
	foundMetaschemaFetch := false
	for _, call := range fetcher.fetchCalls {
		if call == "http://example.com/custom-meta.json" {
			foundMetaschemaFetch = true
			break
		}
	}
	if !foundMetaschemaFetch {
		t.Errorf("expected fetch for custom metaschema via $schema, got calls: %v", fetcher.fetchCalls)
	}

	// Verify vocabulary filtering was applied (minLength stripped)
	var outputSchema map[string]any
	if err := json.Unmarshal(result[0].Schema, &outputSchema); err != nil {
		t.Fatalf("failed to parse output schema: %v", err)
	}

	if _, ok := outputSchema["minLength"]; ok {
		t.Error("minLength should be stripped when validation vocab is disabled")
	}
}

// TestSharedCache_DeclaredSchemaAlsoReferencedViaRef tests that when a declared schema
// is also referenced via $ref from another schema, it's only fetched once
// (from retriever via shared cache, not refetched by processor).
func TestSharedCache_DeclaredSchemaAlsoReferencedViaRef(t *testing.T) {
	ctx := context.Background()
	mockFetch := newMockFetcher()

	// This is the address schema - it will be "declared" (pre-fetched by retriever)
	// and also referenced via $ref in person schema
	addressSchema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"street": { "type": "string" },
			"city": { "type": "string" }
		}
	}`)

	// Add to mock fetcher in case processor tries to fetch it (which it shouldn't)
	mockFetch.addResponse("http://example.com/address.json", string(addressSchema))

	// Person schema references address via $ref
	personSchema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"name": { "type": "string" },
			"address": { "$ref": "http://example.com/address.json" }
		}
	}`)

	// Simulate retriever having already fetched both schemas
	// and populated the shared cache
	sharedCache := fetcher.NewSharedCache()
	sharedCache.Set("http://example.com/address.json", addressSchema)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "address",
			Schema:    addressSchema,
			SourceURI: "http://example.com/address.json",
		},
		{
			Namespace: "test",
			ID:        "person",
			Schema:    personSchema,
			SourceURI: "http://example.com/person.json",
		},
	}

	_, err := Process(ctx, schemas, Options{
		Fetcher: mockFetch,
		Cache:   sharedCache,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// The address schema should NOT be fetched again via the Fetcher
	// because it's already in the shared cache
	if mockFetch.callCount != 0 {
		t.Errorf("expected 0 fetches (address should come from shared cache), got %d: %v",
			mockFetch.callCount, mockFetch.fetchCalls)
	}
}

// TestEarlyValidation_InvalidDeclaredSchemaFailsBeforeFetch tests US-003:
// Invalid declared schema should fail fast before any external refs are fetched.
func TestEarlyValidation_InvalidDeclaredSchemaFailsBeforeFetch(t *testing.T) {
	ctx := context.Background()
	mockFetch := newMockFetcher()

	// Add external schema that should NOT be fetched (we expect early failure)
	mockFetch.addResponse("http://example.com/external.json", `{"type": "string"}`)

	// Invalid schema: type must be string or array of strings, not integer
	invalidSchema := json.RawMessage(`{
		"type": 12345,
		"properties": {
			"external": { "$ref": "http://example.com/external.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "invalid",
			Schema:    invalidSchema,
			SourceURI: "http://test.com/invalid.json",
		},
	}

	_, err := Process(ctx, schemas, Options{Fetcher: mockFetch})
	if err == nil {
		t.Fatal("expected validation error for invalid declared schema")
	}

	// Error should mention validation failed
	if !contains(err.Error(), "validation") || !contains(err.Error(), "failed") {
		t.Errorf("error should mention validation failed, got: %s", err.Error())
	}

	// CRITICAL: No external refs should have been fetched
	if len(mockFetch.fetchCalls) != 0 {
		t.Errorf("expected 0 fetches (validation should fail before crawl), got %d: %v",
			len(mockFetch.fetchCalls), mockFetch.fetchCalls)
	}
}

// TestEarlyValidation_ValidSchemaWithExternalRefsProceedsToCrawl tests US-003:
// Valid schema with external refs should pass early validation and proceed to crawl.
func TestEarlyValidation_ValidSchemaWithExternalRefsProceedsToCrawl(t *testing.T) {
	ctx := context.Background()
	mockFetch := newMockFetcher()

	// External schema that should be fetched after early validation passes
	mockFetch.addResponse("http://example.com/address.json", `{
		"type": "object",
		"properties": {
			"street": { "type": "string" }
		}
	}`)

	// Valid schema with external ref
	validSchema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"address": { "$ref": "http://example.com/address.json" }
		}
	}`)

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "person",
			Schema:    validSchema,
			SourceURI: "http://test.com/person.json",
		},
	}

	result, err := Process(ctx, schemas, Options{Fetcher: mockFetch})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should have produced result
	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	// External ref should have been fetched (crawl phase executed)
	if len(mockFetch.fetchCalls) != 1 {
		t.Errorf("expected 1 fetch for external ref, got %d: %v",
			len(mockFetch.fetchCalls), mockFetch.fetchCalls)
	}
	if len(mockFetch.fetchCalls) > 0 && mockFetch.fetchCalls[0] != "http://example.com/address.json" {
		t.Errorf("expected fetch for address.json, got %s", mockFetch.fetchCalls[0])
	}
}

// TestSharedCache_CustomMetaschemaFetchedOnce tests that when multiple schemas
// use the same custom $schema URI, it's only fetched once via shared cache.
func TestSharedCache_CustomMetaschemaFetchedOnce(t *testing.T) {
	ctx := context.Background()
	mockFetch := newMockFetcher()

	// Custom metaschema
	customMeta := `{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"$vocabulary": {
			"https://json-schema.org/draft/2020-12/vocab/core": true
		}
	}`
	mockFetch.addResponse("http://example.com/custom-meta.json", customMeta)

	// Two schemas using the same custom metaschema
	schema1 := json.RawMessage(`{
		"$schema": "http://example.com/custom-meta.json",
		"type": "string"
	}`)
	schema2 := json.RawMessage(`{
		"$schema": "http://example.com/custom-meta.json",
		"type": "integer"
	}`)

	sharedCache := fetcher.NewSharedCache()

	schemas := []retriever.RetrievedSchema{
		{
			Namespace: "test",
			ID:        "schema1",
			Schema:    schema1,
			SourceURI: "http://test.com/schema1.json",
		},
		{
			Namespace: "test",
			ID:        "schema2",
			Schema:    schema2,
			SourceURI: "http://test.com/schema2.json",
		},
	}

	_, err := Process(ctx, schemas, Options{
		Fetcher: mockFetch,
		Cache:   sharedCache,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Count how many times the custom metaschema was fetched
	metaFetchCount := 0
	for _, call := range mockFetch.fetchCalls {
		if call == "http://example.com/custom-meta.json" {
			metaFetchCount++
		}
	}

	// Should be fetched exactly once (first schema triggers fetch, second uses cache)
	if metaFetchCount != 1 {
		t.Errorf("expected custom metaschema to be fetched exactly once, got %d times. All calls: %v",
			metaFetchCount, mockFetch.fetchCalls)
	}
}
