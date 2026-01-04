package retriever

import (
	"context"
	"encoding/json"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/xschemadev/xschema/parser"
)

func testdataPath(name string) string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "testdata", name)
}

func TestRetrieveFromFile(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc") // Use as config path for relative resolution

	tests := []struct {
		name     string
		file     string
		wantType string
		wantErr  bool
	}{
		{"user schema", "user.json", "object", false},
		{"post schema", "post.json", "object", false},
		{"config schema", "config.json", "object", false},
		{"invalid json", "invalid.txt", "", true},
		{"not found", "nonexistent.json", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := retrieveFromFile(ctx, tt.file, configPath)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			var parsed map[string]any
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("invalid JSON: %v", err)
			}
			if parsed["type"] != tt.wantType {
				t.Errorf("expected type=%s, got %v", tt.wantType, parsed["type"])
			}
		})
	}
}

func TestRetrieveFromURL(t *testing.T) {
	ctx := context.Background()
	opts := DefaultOptions()

	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{"eslint schema", "https://json.schemastore.org/eslintrc.json", false},
		{"opencode config", "https://opencode.ai/config.json", false},
		{"not found", "https://httpstat.us/404", true},
		{"server error", "https://httpstat.us/500", true},
		{"invalid json", "https://httpstat.us/200", true}, // returns text, not JSON
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := retrieveFromURL(ctx, tt.url, opts)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !json.Valid(result) {
				t.Error("result is not valid JSON")
			}
		})
	}
}

func TestRetrieveAggregation(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	decls := []parser.Declaration{
		// Zod schemas
		{Namespace: "user", ID: "User", SourceType: parser.SourceFile, Source: json.RawMessage(`"user.json"`), Adapter: "zod", ConfigPath: configPath},
		{Namespace: "user", ID: "Post", SourceType: parser.SourceFile, Source: json.RawMessage(`"post.json"`), Adapter: "zod", ConfigPath: configPath},
		// Pydantic schemas
		{Namespace: "config", ID: "Config", SourceType: parser.SourceFile, Source: json.RawMessage(`"config.json"`), Adapter: "@xschema/pydantic", ConfigPath: configPath},
		{Namespace: "config", ID: "UserPy", SourceType: parser.SourceFile, Source: json.RawMessage(`"user.json"`), Adapter: "@xschema/pydantic", ConfigPath: configPath},
	}

	schemas, err := Retrieve(ctx, decls, DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(schemas) != 4 {
		t.Fatalf("expected 4 schemas, got %d", len(schemas))
	}

	// Group by adapter
	groups := GroupByAdapter(schemas)
	if len(groups) != 2 {
		t.Fatalf("expected 2 adapter groups, got %d", len(groups))
	}

	pydanticSchemas := groups["@xschema/pydantic"]
	zodSchemas := groups["zod"]

	if len(pydanticSchemas) != 2 {
		t.Errorf("expected 2 pydantic schemas, got %d", len(pydanticSchemas))
	}
	if len(zodSchemas) != 2 {
		t.Errorf("expected 2 zod schemas, got %d", len(zodSchemas))
	}
}

func TestRetrieveConcurrency(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	// Create 15 declarations to test concurrency (limit is 10)
	var decls []parser.Declaration
	files := []string{"user.json", "post.json", "config.json"}
	for i := 0; i < 15; i++ {
		file := files[i%len(files)]
		adapter := "zod"
		if i%2 == 0 {
			adapter = "@xschema/pydantic"
		}
		decls = append(decls, parser.Declaration{
			Namespace:  "test",
			ID:         file[:len(file)-5] + string(rune('A'+i)), // userA, postB, etc
			SourceType: parser.SourceFile,
			Source:     json.RawMessage(`"` + file + `"`),
			Adapter:    adapter,
			ConfigPath: configPath,
		})
	}

	schemas, err := Retrieve(ctx, decls, DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(schemas) != 15 {
		t.Errorf("expected 15 schemas, got %d", len(schemas))
	}

	groups := GroupByAdapter(schemas)
	if len(groups) != 2 {
		t.Fatalf("expected 2 adapter groups, got %d", len(groups))
	}
}

func TestRetrieveErrors(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	tests := []struct {
		name  string
		decls []parser.Declaration
	}{
		{
			name: "file not found",
			decls: []parser.Declaration{
				{Namespace: "test", ID: "Missing", SourceType: parser.SourceFile, Source: json.RawMessage(`"nonexistent.json"`), Adapter: "zod", ConfigPath: configPath},
			},
		},
		{
			name: "invalid json file",
			decls: []parser.Declaration{
				{Namespace: "test", ID: "Invalid", SourceType: parser.SourceFile, Source: json.RawMessage(`"invalid.txt"`), Adapter: "zod", ConfigPath: configPath},
			},
		},
		{
			name: "url not found",
			decls: []parser.Declaration{
				{Namespace: "test", ID: "NotFound", SourceType: parser.SourceURL, Source: json.RawMessage(`"https://httpstat.us/404"`), Adapter: "zod", ConfigPath: configPath},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Retrieve(ctx, tt.decls, DefaultOptions())
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}

func TestRetrieveContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	configPath := testdataPath("fake.jsonc")
	decls := []parser.Declaration{
		{Namespace: "test", ID: "User", SourceType: parser.SourceURL, Source: json.RawMessage(`"https://json.schemastore.org/eslintrc.json"`), Adapter: "zod", ConfigPath: configPath},
	}

	_, err := Retrieve(ctx, decls, DefaultOptions())
	if err == nil {
		t.Error("expected context cancellation error")
	}
}

func TestRetrieveNoCache(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	// Same file referenced twice with different names
	decls := []parser.Declaration{
		{Namespace: "test", ID: "User1", SourceType: parser.SourceFile, Source: json.RawMessage(`"user.json"`), Adapter: "zod", ConfigPath: configPath},
		{Namespace: "test", ID: "User2", SourceType: parser.SourceFile, Source: json.RawMessage(`"user.json"`), Adapter: "zod", ConfigPath: configPath},
	}

	// With cache (default) - should work
	opts := DefaultOptions()
	schemas, err := Retrieve(ctx, decls, opts)
	if err != nil {
		t.Fatalf("with cache: %v", err)
	}
	if len(schemas) != 2 {
		t.Errorf("expected 2 schemas, got %d", len(schemas))
	}

	// Without cache - should also work (just fetches twice)
	opts.NoCache = true
	schemas, err = Retrieve(ctx, decls, opts)
	if err != nil {
		t.Fatalf("without cache: %v", err)
	}
	if len(schemas) != 2 {
		t.Errorf("expected 2 schemas, got %d", len(schemas))
	}
}

func TestRetrieveCustomConcurrency(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	decls := []parser.Declaration{
		{Namespace: "test", ID: "User", SourceType: parser.SourceFile, Source: json.RawMessage(`"user.json"`), Adapter: "zod", ConfigPath: configPath},
		{Namespace: "test", ID: "Post", SourceType: parser.SourceFile, Source: json.RawMessage(`"post.json"`), Adapter: "zod", ConfigPath: configPath},
		{Namespace: "test", ID: "Config", SourceType: parser.SourceFile, Source: json.RawMessage(`"config.json"`), Adapter: "zod", ConfigPath: configPath},
	}

	// Concurrency = 1 (sequential)
	opts := Options{
		Concurrency: 1,
		HTTPTimeout: DefaultOptions().HTTPTimeout,
		Retries:     DefaultOptions().Retries,
	}
	schemas, err := Retrieve(ctx, decls, opts)
	if err != nil {
		t.Fatalf("concurrency=1: %v", err)
	}
	if len(schemas) != 3 {
		t.Errorf("expected 3 schemas, got %d", len(schemas))
	}
}

func TestRetrieveCustomTimeout(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	// Very short timeout for a slow endpoint
	opts := Options{
		Concurrency: 1,
		HTTPTimeout: 1 * time.Millisecond, // impossibly short
		Retries:     1,
	}

	decls := []parser.Declaration{
		{Namespace: "test", ID: "Slow", SourceType: parser.SourceURL, Source: json.RawMessage(`"https://httpstat.us/200?sleep=5000"`), Adapter: "zod", ConfigPath: configPath},
	}

	_, err := Retrieve(ctx, decls, opts)
	if err == nil {
		t.Error("expected timeout error")
	}
}

func TestRetrieveSingleAttempt(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	// Single attempt (Retries=1) - should fail on 500 without retrying
	opts := Options{
		Concurrency: 1,
		HTTPTimeout: 5 * time.Second,
		Retries:     1,
	}

	decls := []parser.Declaration{
		{Namespace: "test", ID: "ServerError", SourceType: parser.SourceURL, Source: json.RawMessage(`"https://httpstat.us/500"`), Adapter: "zod", ConfigPath: configPath},
	}

	_, err := Retrieve(ctx, decls, opts)
	if err == nil {
		t.Error("expected error with single attempt on 500")
	}
}

func TestRetrieveInlineJSON(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	inlineSchema := `{"type": "object", "properties": {"name": {"type": "string"}}}`
	decls := []parser.Declaration{
		{Namespace: "test", ID: "Inline", SourceType: parser.SourceJSON, Source: json.RawMessage(inlineSchema), Adapter: "zod", ConfigPath: configPath},
	}

	schemas, err := Retrieve(ctx, decls, DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(schemas) != 1 {
		t.Fatalf("expected 1 schema, got %d", len(schemas))
	}

	// The schema should be the inline JSON itself
	if string(schemas[0].Schema) != inlineSchema {
		t.Errorf("expected inline schema to be passed through, got %s", schemas[0].Schema)
	}
}

func TestRetrievedSchemaKey(t *testing.T) {
	s := RetrievedSchema{
		Namespace: "user",
		ID:        "TestUrl",
	}

	if s.Key() != "user:TestUrl" {
		t.Errorf("expected key 'user:TestUrl', got %q", s.Key())
	}
}

func TestRetrieveEmptyDeclarations(t *testing.T) {
	ctx := context.Background()
	schemas, err := Retrieve(ctx, []parser.Declaration{}, DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if schemas != nil && len(schemas) != 0 {
		t.Errorf("expected nil or empty slice, got %d schemas", len(schemas))
	}
}

func TestRetrieveUnknownSourceType(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	decls := []parser.Declaration{
		{Namespace: "test", ID: "Unknown", SourceType: "invalid", Source: json.RawMessage(`"test"`), Adapter: "zod", ConfigPath: configPath},
	}

	_, err := Retrieve(ctx, decls, DefaultOptions())
	if err == nil {
		t.Error("expected error for unknown source type")
	}
}

func TestRetrieveInvalidURLSource(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	decls := []parser.Declaration{
		{Namespace: "test", ID: "BadURL", SourceType: parser.SourceURL, Source: json.RawMessage(`123`), Adapter: "zod", ConfigPath: configPath},
	}

	_, err := Retrieve(ctx, decls, DefaultOptions())
	if err == nil {
		t.Error("expected error for invalid URL source JSON")
	}
}

func TestRetrieveInvalidFileSource(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	decls := []parser.Declaration{
		{Namespace: "test", ID: "BadFile", SourceType: parser.SourceFile, Source: json.RawMessage(`123`), Adapter: "zod", ConfigPath: configPath},
	}

	_, err := Retrieve(ctx, decls, DefaultOptions())
	if err == nil {
		t.Error("expected error for invalid file source JSON")
	}
}

func TestGroupByAdapterEmpty(t *testing.T) {
	groups := GroupByAdapter([]RetrievedSchema{})
	if len(groups) != 0 {
		t.Errorf("expected 0 groups, got %d", len(groups))
	}
}

func TestGroupByAdapterSingle(t *testing.T) {
	schemas := []RetrievedSchema{
		{Namespace: "user", ID: "User", Adapter: "zod"},
	}
	groups := GroupByAdapter(schemas)

	if len(groups) != 1 {
		t.Errorf("expected 1 group, got %d", len(groups))
	}
	if len(groups["zod"]) != 1 {
		t.Errorf("expected 1 schema in zod group, got %d", len(groups["zod"]))
	}
}

func TestSortedAdapters(t *testing.T) {
	groups := map[string][]RetrievedSchema{
		"zod":      {},
		"@xschema/pydantic": {},
		"@xschema/ajv":      {},
	}

	sorted := SortedAdapters(groups)

	expected := []string{"@xschema/ajv", "@xschema/pydantic", "zod"}
	for i, v := range expected {
		if sorted[i] != v {
			t.Errorf("expected sorted[%d]=%s, got %s", i, v, sorted[i])
		}
	}
}

func TestSortedAdaptersEmpty(t *testing.T) {
	groups := map[string][]RetrievedSchema{}
	sorted := SortedAdapters(groups)

	if len(sorted) != 0 {
		t.Errorf("expected 0 sorted adapters, got %d", len(sorted))
	}
}

func TestRetrieveCacheHit(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	// Same source, different IDs - second should hit cache
	decls := []parser.Declaration{
		{Namespace: "test", ID: "User1", SourceType: parser.SourceFile, Source: json.RawMessage(`"user.json"`), Adapter: "zod", ConfigPath: configPath},
		{Namespace: "test", ID: "User2", SourceType: parser.SourceFile, Source: json.RawMessage(`"user.json"`), Adapter: "zod", ConfigPath: configPath},
	}

	opts := DefaultOptions()
	opts.Concurrency = 1 // Force sequential to test cache

	schemas, err := Retrieve(ctx, decls, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(schemas) != 2 {
		t.Errorf("expected 2 schemas, got %d", len(schemas))
	}

	// Both should have same schema content
	if string(schemas[0].Schema) != string(schemas[1].Schema) {
		t.Error("cache should return same content for same source")
	}
}

func TestRetrieveMixedSourceTypes(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	inlineSchema := `{"type": "object", "properties": {"inline": {"type": "boolean"}}}`
	decls := []parser.Declaration{
		{Namespace: "test", ID: "FromFile", SourceType: parser.SourceFile, Source: json.RawMessage(`"user.json"`), Adapter: "zod", ConfigPath: configPath},
		{Namespace: "test", ID: "Inline", SourceType: parser.SourceJSON, Source: json.RawMessage(inlineSchema), Adapter: "zod", ConfigPath: configPath},
	}

	schemas, err := Retrieve(ctx, decls, DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(schemas) != 2 {
		t.Fatalf("expected 2 schemas, got %d", len(schemas))
	}

	// Verify inline schema is passed through as-is
	if string(schemas[1].Schema) != inlineSchema {
		t.Errorf("inline schema mismatch: got %s", schemas[1].Schema)
	}
}

func TestRetrievePreservesMetadata(t *testing.T) {
	ctx := context.Background()
	configPath := testdataPath("fake.jsonc")

	decls := []parser.Declaration{
		{Namespace: "myns", ID: "MyID", SourceType: parser.SourceFile, Source: json.RawMessage(`"user.json"`), Adapter: "@xschema/custom", ConfigPath: configPath},
	}

	schemas, err := Retrieve(ctx, decls, DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if schemas[0].Namespace != "myns" {
		t.Errorf("expected namespace 'myns', got %s", schemas[0].Namespace)
	}
	if schemas[0].ID != "MyID" {
		t.Errorf("expected ID 'MyID', got %s", schemas[0].ID)
	}
	if schemas[0].Adapter != "@xschema/custom" {
		t.Errorf("expected adapter '@xschema/custom', got %s", schemas[0].Adapter)
	}
}

func TestDefaultOptions(t *testing.T) {
	opts := DefaultOptions()

	if opts.Concurrency != 10 {
		t.Errorf("expected concurrency 10, got %d", opts.Concurrency)
	}
	if opts.HTTPTimeout != 30*time.Second {
		t.Errorf("expected timeout 30s, got %v", opts.HTTPTimeout)
	}
	if opts.Retries != 3 {
		t.Errorf("expected retries 3, got %d", opts.Retries)
	}
	if opts.NoCache != false {
		t.Error("expected NoCache false by default")
	}
}
