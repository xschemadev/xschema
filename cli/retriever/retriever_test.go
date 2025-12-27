package retriever

import (
	"context"
	"encoding/json"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/xschema/cli/parser"
)

func testdataPath(name string) string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "testdata", name)
}

func TestRetrieveFromFile(t *testing.T) {
	ctx := context.Background()
	declPath := testdataPath("fake.ts")

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
			result, err := retrieveFromFile(ctx, tt.file, declPath)
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
	declPath := testdataPath("fake.ts")

	decls := []parser.Declaration{
		// Zod schemas
		{Name: "User", Source: "file", Location: "user.json", Adapter: parser.AdapterRef{Package: "@xschema/zod", Language: "typescript"}, File: declPath},
		{Name: "Post", Source: "file", Location: "post.json", Adapter: parser.AdapterRef{Package: "@xschema/zod", Language: "typescript"}, File: declPath},
		// Pydantic schemas
		{Name: "Config", Source: "file", Location: "config.json", Adapter: parser.AdapterRef{Package: "@xschema/pydantic", Language: "python"}, File: declPath},
		{Name: "UserPy", Source: "file", Location: "user.json", Adapter: parser.AdapterRef{Package: "@xschema/pydantic", Language: "python"}, File: declPath},
	}

	batches, err := Retrieve(ctx, decls, DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(batches) != 2 {
		t.Fatalf("expected 2 batches, got %d", len(batches))
	}

	// Results are sorted by adapter key
	pydanticBatch := batches[0] // @xschema/pydantic comes before @xschema/zod
	zodBatch := batches[1]

	if pydanticBatch.Adapter != "@xschema/pydantic" {
		t.Errorf("expected first batch adapter=@xschema/pydantic, got %s", pydanticBatch.Adapter)
	}
	if pydanticBatch.Language != "python" {
		t.Errorf("expected language=python, got %s", pydanticBatch.Language)
	}
	if len(pydanticBatch.Schemas) != 2 {
		t.Errorf("expected 2 pydantic schemas, got %d", len(pydanticBatch.Schemas))
	}

	if zodBatch.Adapter != "@xschema/zod" {
		t.Errorf("expected second batch adapter=@xschema/zod, got %s", zodBatch.Adapter)
	}
	if zodBatch.Language != "typescript" {
		t.Errorf("expected language=typescript, got %s", zodBatch.Language)
	}
	if len(zodBatch.Schemas) != 2 {
		t.Errorf("expected 2 zod schemas, got %d", len(zodBatch.Schemas))
	}
}

func TestRetrieveConcurrency(t *testing.T) {
	ctx := context.Background()
	declPath := testdataPath("fake.ts")

	// Create 15 declarations to test concurrency (limit is 10)
	var decls []parser.Declaration
	files := []string{"user.json", "post.json", "config.json"}
	for i := 0; i < 15; i++ {
		file := files[i%len(files)]
		adapter := "@xschema/zod"
		lang := "typescript"
		if i%2 == 0 {
			adapter = "@xschema/pydantic"
			lang = "python"
		}
		decls = append(decls, parser.Declaration{
			Name:     file[:len(file)-5] + string(rune('A'+i)), // userA, postB, etc
			Source:   "file",
			Location: file,
			Adapter:  parser.AdapterRef{Package: adapter, Language: lang},
			File:     declPath,
		})
	}

	batches, err := Retrieve(ctx, decls, DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(batches) != 2 {
		t.Fatalf("expected 2 batches, got %d", len(batches))
	}

	total := 0
	for _, b := range batches {
		total += len(b.Schemas)
	}
	if total != 15 {
		t.Errorf("expected 15 total schemas, got %d", total)
	}
}

func TestRetrieveErrors(t *testing.T) {
	ctx := context.Background()
	declPath := testdataPath("fake.ts")

	tests := []struct {
		name  string
		decls []parser.Declaration
	}{
		{
			name: "file not found",
			decls: []parser.Declaration{
				{Name: "Missing", Source: "file", Location: "nonexistent.json", Adapter: parser.AdapterRef{Package: "@xschema/zod"}, File: declPath},
			},
		},
		{
			name: "invalid json file",
			decls: []parser.Declaration{
				{Name: "Invalid", Source: "file", Location: "invalid.txt", Adapter: parser.AdapterRef{Package: "@xschema/zod"}, File: declPath},
			},
		},
		{
			name: "url not found",
			decls: []parser.Declaration{
				{Name: "NotFound", Source: "url", Location: "https://httpstat.us/404", Adapter: parser.AdapterRef{Package: "@xschema/zod"}, File: declPath},
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

	declPath := testdataPath("fake.ts")
	decls := []parser.Declaration{
		{Name: "User", Source: "url", Location: "https://json.schemastore.org/eslintrc.json", Adapter: parser.AdapterRef{Package: "@xschema/zod"}, File: declPath},
	}

	_, err := Retrieve(ctx, decls, DefaultOptions())
	if err == nil {
		t.Error("expected context cancellation error")
	}
}

func TestRetrieveNoCache(t *testing.T) {
	ctx := context.Background()
	declPath := testdataPath("fake.ts")

	// Same file referenced twice with different names
	decls := []parser.Declaration{
		{Name: "User1", Source: "file", Location: "user.json", Adapter: parser.AdapterRef{Package: "@xschema/zod", Language: "typescript"}, File: declPath},
		{Name: "User2", Source: "file", Location: "user.json", Adapter: parser.AdapterRef{Package: "@xschema/zod", Language: "typescript"}, File: declPath},
	}

	// With cache (default) - should work
	opts := DefaultOptions()
	batches, err := Retrieve(ctx, decls, opts)
	if err != nil {
		t.Fatalf("with cache: %v", err)
	}
	if len(batches[0].Schemas) != 2 {
		t.Errorf("expected 2 schemas, got %d", len(batches[0].Schemas))
	}

	// Without cache - should also work (just fetches twice)
	opts.NoCache = true
	batches, err = Retrieve(ctx, decls, opts)
	if err != nil {
		t.Fatalf("without cache: %v", err)
	}
	if len(batches[0].Schemas) != 2 {
		t.Errorf("expected 2 schemas, got %d", len(batches[0].Schemas))
	}
}

func TestRetrieveCustomConcurrency(t *testing.T) {
	ctx := context.Background()
	declPath := testdataPath("fake.ts")

	decls := []parser.Declaration{
		{Name: "User", Source: "file", Location: "user.json", Adapter: parser.AdapterRef{Package: "@xschema/zod", Language: "typescript"}, File: declPath},
		{Name: "Post", Source: "file", Location: "post.json", Adapter: parser.AdapterRef{Package: "@xschema/zod", Language: "typescript"}, File: declPath},
		{Name: "Config", Source: "file", Location: "config.json", Adapter: parser.AdapterRef{Package: "@xschema/zod", Language: "typescript"}, File: declPath},
	}

	// Concurrency = 1 (sequential)
	opts := Options{
		Concurrency: 1,
		HTTPTimeout: DefaultOptions().HTTPTimeout,
		Retries:     DefaultOptions().Retries,
	}
	batches, err := Retrieve(ctx, decls, opts)
	if err != nil {
		t.Fatalf("concurrency=1: %v", err)
	}
	if len(batches[0].Schemas) != 3 {
		t.Errorf("expected 3 schemas, got %d", len(batches[0].Schemas))
	}
}

func TestRetrieveCustomTimeout(t *testing.T) {
	ctx := context.Background()
	declPath := testdataPath("fake.ts")

	// Very short timeout for a slow endpoint
	opts := Options{
		Concurrency: 1,
		HTTPTimeout: 1 * time.Millisecond, // impossibly short
		Retries:     1,
	}

	decls := []parser.Declaration{
		{Name: "Slow", Source: "url", Location: "https://httpstat.us/200?sleep=5000", Adapter: parser.AdapterRef{Package: "@xschema/zod"}, File: declPath},
	}

	_, err := Retrieve(ctx, decls, opts)
	if err == nil {
		t.Error("expected timeout error")
	}
}

func TestRetrieveSingleAttempt(t *testing.T) {
	ctx := context.Background()
	declPath := testdataPath("fake.ts")

	// Single attempt (Retries=1) - should fail on 500 without retrying
	opts := Options{
		Concurrency: 1,
		HTTPTimeout: 5 * time.Second,
		Retries:     1,
	}

	decls := []parser.Declaration{
		{Name: "ServerError", Source: "url", Location: "https://httpstat.us/500", Adapter: parser.AdapterRef{Package: "@xschema/zod"}, File: declPath},
	}

	_, err := Retrieve(ctx, decls, opts)
	if err == nil {
		t.Error("expected error with single attempt on 500")
	}
}
