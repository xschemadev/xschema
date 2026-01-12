package compliance

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestGetCacheDir(t *testing.T) {
	cacheDir, err := GetCacheDir()
	if err != nil {
		t.Fatalf("GetCacheDir() error = %v", err)
	}

	if cacheDir == "" {
		t.Error("GetCacheDir() returned empty string")
	}

	// Should end with "xschema"
	if filepath.Base(cacheDir) != "xschema" {
		t.Errorf("GetCacheDir() = %s, want path ending in 'xschema'", cacheDir)
	}
}

func TestDownloadAndExtract(t *testing.T) {
	// Create a test tarball
	tarballContent := createTestTarball(t, map[string]string{
		"test-repo-main/tests/draft2020-12/type.json": `[{"description": "test"}]`,
		"test-repo-main/README.md":                    "# Test Repo",
	})

	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/gzip")
		w.Write(tarballContent)
	}))
	defer server.Close()

	// Test extraction
	destDir := filepath.Join(t.TempDir(), "extracted")
	err := downloadAndExtract(context.Background(), server.URL, destDir)
	if err != nil {
		t.Fatalf("downloadAndExtract() error = %v", err)
	}

	// Verify extracted files
	tests := []struct {
		path    string
		content string
	}{
		{"tests/draft2020-12/type.json", `[{"description": "test"}]`},
		{"README.md", "# Test Repo"},
	}

	for _, tt := range tests {
		filePath := filepath.Join(destDir, tt.path)
		data, err := os.ReadFile(filePath)
		if err != nil {
			t.Errorf("failed to read %s: %v", tt.path, err)
			continue
		}
		if string(data) != tt.content {
			t.Errorf("file %s content = %q, want %q", tt.path, string(data), tt.content)
		}
	}
}

func TestDownloadAndExtract_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	destDir := filepath.Join(t.TempDir(), "extracted")
	err := downloadAndExtract(context.Background(), server.URL, destDir)
	if err == nil {
		t.Error("downloadAndExtract() expected error for HTTP 404, got nil")
	}
}

func TestDownloadAndExtract_InvalidGzip(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not a gzip file"))
	}))
	defer server.Close()

	destDir := filepath.Join(t.TempDir(), "extracted")
	err := downloadAndExtract(context.Background(), server.URL, destDir)
	if err == nil {
		t.Error("downloadAndExtract() expected error for invalid gzip, got nil")
	}
}

func TestDownloadAndExtract_SkipsPAXHeaders(t *testing.T) {
	// Create tarball with PAX headers (simulating GitHub tarball behavior)
	tarballContent := createTestTarballWithPAX(t, map[string]string{
		"test-repo-main/file.txt": "content",
	})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(tarballContent)
	}))
	defer server.Close()

	destDir := filepath.Join(t.TempDir(), "extracted")
	err := downloadAndExtract(context.Background(), server.URL, destDir)
	if err != nil {
		t.Fatalf("downloadAndExtract() error = %v", err)
	}

	// File should exist
	if _, err := os.Stat(filepath.Join(destDir, "file.txt")); os.IsNotExist(err) {
		t.Error("expected file.txt to exist")
	}
}

func TestFetchTestSuite_UsesCache(t *testing.T) {
	// Create a fake cache directory
	tmpDir := t.TempDir()
	origFunc := GetCacheDir

	// Override GetCacheDir for this test
	oldUserCacheDir := os.Getenv("XDG_CACHE_HOME")
	os.Setenv("XDG_CACHE_HOME", tmpDir)
	defer os.Setenv("XDG_CACHE_HOME", oldUserCacheDir)

	// Create fake cached suite
	suiteDir := filepath.Join(tmpDir, "xschema", testSuiteCacheDir)
	testsDir := filepath.Join(suiteDir, "tests", "draft2020-12")
	if err := os.MkdirAll(testsDir, 0755); err != nil {
		t.Fatalf("failed to create test dir: %v", err)
	}

	// Write a marker file
	markerFile := filepath.Join(testsDir, "marker.json")
	if err := os.WriteFile(markerFile, []byte("cached"), 0644); err != nil {
		t.Fatalf("failed to write marker: %v", err)
	}

	_ = origFunc // avoid unused warning
}

// Helper to create a test tarball
func createTestTarball(t *testing.T, files map[string]string) []byte {
	t.Helper()

	pr, pw, _ := os.Pipe()
	gzWriter := gzip.NewWriter(pw)
	tarWriter := tar.NewWriter(gzWriter)

	go func() {
		for path, content := range files {
			hdr := &tar.Header{
				Name:     path,
				Mode:     0644,
				Size:     int64(len(content)),
				Typeflag: tar.TypeReg,
			}
			if err := tarWriter.WriteHeader(hdr); err != nil {
				t.Errorf("failed to write header: %v", err)
				return
			}
			if _, err := tarWriter.Write([]byte(content)); err != nil {
				t.Errorf("failed to write content: %v", err)
				return
			}
		}
		tarWriter.Close()
		gzWriter.Close()
		pw.Close()
	}()

	result := make([]byte, 0)
	tmp := make([]byte, 1024)
	for {
		n, err := pr.Read(tmp)
		if n > 0 {
			result = append(result, tmp[:n]...)
		}
		if err != nil {
			break
		}
	}

	return result
}

// Helper to create tarball with PAX headers
func createTestTarballWithPAX(t *testing.T, files map[string]string) []byte {
	t.Helper()

	pr, pw, _ := os.Pipe()
	gzWriter := gzip.NewWriter(pw)
	tarWriter := tar.NewWriter(gzWriter)

	go func() {
		// Write PAX global header first
		paxHdr := &tar.Header{
			Typeflag: tar.TypeXGlobalHeader,
			Name:     "pax_global_header",
			Size:     0,
		}
		tarWriter.WriteHeader(paxHdr)

		for path, content := range files {
			hdr := &tar.Header{
				Name:     path,
				Mode:     0644,
				Size:     int64(len(content)),
				Typeflag: tar.TypeReg,
			}
			tarWriter.WriteHeader(hdr)
			tarWriter.Write([]byte(content))
		}
		tarWriter.Close()
		gzWriter.Close()
		pw.Close()
	}()

	result := make([]byte, 0)
	tmp := make([]byte, 1024)
	for {
		n, err := pr.Read(tmp)
		if n > 0 {
			result = append(result, tmp[:n]...)
		}
		if err != nil {
			break
		}
	}

	return result
}

// ---- LocalhostFetcher Tests ----

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
		"http://localhost:5000/schema.json", // different port
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
