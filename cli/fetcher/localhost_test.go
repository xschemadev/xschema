package fetcher

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestLocalhostFetcher_FetchesLocalFile(t *testing.T) {
	// Create temp remotes directory with a schema
	remotesDir := t.TempDir()
	schemaContent := `{"type": "integer"}`
	if err := os.WriteFile(filepath.Join(remotesDir, "integer.json"), []byte(schemaContent), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	fetcher := NewLocalhostFetcher(remotesDir)
	data, err := fetcher.Fetch(context.Background(), "http://localhost:1234/integer.json")
	if err != nil {
		t.Fatalf("Fetch() error = %v", err)
	}

	if string(data) != schemaContent {
		t.Errorf("Fetch() = %s, want %s", string(data), schemaContent)
	}
}

func TestLocalhostFetcher_FetchesSubdirectoryFile(t *testing.T) {
	// Create temp remotes directory with subdirectory
	remotesDir := t.TempDir()
	subDir := filepath.Join(remotesDir, "subdir")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("failed to create subdir: %v", err)
	}

	schemaContent := `{"type": "string"}`
	if err := os.WriteFile(filepath.Join(subDir, "schema.json"), []byte(schemaContent), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	fetcher := NewLocalhostFetcher(remotesDir)
	data, err := fetcher.Fetch(context.Background(), "http://localhost:1234/subdir/schema.json")
	if err != nil {
		t.Fatalf("Fetch() error = %v", err)
	}

	if string(data) != schemaContent {
		t.Errorf("Fetch() = %s, want %s", string(data), schemaContent)
	}
}

func TestLocalhostFetcher_ReturnsStubForNonLocalhostURL(t *testing.T) {
	// Non-localhost URLs return stub schemas instead of errors.
	// This is intentional - the processor's crawl phase extracts all external-looking
	// refs, but the bundler will check local $ids first. Stubs for unfetchable URLs
	// are never used if the URL is claimed by a local $id.
	remotesDir := t.TempDir()
	fetcher := NewLocalhostFetcher(remotesDir)

	tests := []string{
		"http://example.com/schema.json",
		"https://json-schema.org/draft/2020-12/schema",
		"http://localhost:5000/schema.json",  // different port
		"https://localhost:1234/schema.json", // https not http
	}

	for _, url := range tests {
		data, err := fetcher.Fetch(context.Background(), url)
		if err != nil {
			t.Errorf("Fetch(%q) unexpected error: %v", url, err)
		}
		if string(data) != "{}" {
			t.Errorf("Fetch(%q) = %q, want stub schema {}", url, string(data))
		}
	}
}

func TestLocalhostFetcher_ReturnsStubForMissingFile(t *testing.T) {
	// Missing localhost files return stub schemas instead of errors.
	// This handles test schemas that define $id claiming a localhost URL without
	// a corresponding file - the bundler will resolve via local $id, not the stub.
	remotesDir := t.TempDir()
	fetcher := NewLocalhostFetcher(remotesDir)

	data, err := fetcher.Fetch(context.Background(), "http://localhost:1234/nonexistent.json")
	if err != nil {
		t.Fatalf("Fetch() unexpected error: %v", err)
	}

	if string(data) != "{}" {
		t.Errorf("Fetch() = %q, want stub schema {}", string(data))
	}
}

func TestLocalhostFetcher_StripsQueryParams(t *testing.T) {
	// Create temp remotes directory with a schema
	remotesDir := t.TempDir()
	schemaContent := `{"type": "boolean"}`
	if err := os.WriteFile(filepath.Join(remotesDir, "schema.json"), []byte(schemaContent), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	fetcher := NewLocalhostFetcher(remotesDir)

	// Fetch with query params - should strip them and find the file
	data, err := fetcher.Fetch(context.Background(), "http://localhost:1234/schema.json?foo=bar&baz=qux")
	if err != nil {
		t.Fatalf("Fetch() error = %v", err)
	}

	if string(data) != schemaContent {
		t.Errorf("Fetch() = %s, want %s", string(data), schemaContent)
	}
}
