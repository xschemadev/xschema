package compliance

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	// GitHub repo for JSON Schema Test Suite
	testSuiteRepo = "json-schema-org/JSON-Schema-Test-Suite"
	// Bump by updating this constant and regenerating compliance results.
	testSuiteVersion  = "Test-JSON-Schema-Acceptance-1.035"
	testSuiteURL      = "https://github.com/" + testSuiteRepo + "/archive/refs/tags/" + testSuiteVersion + ".tar.gz"
	testSuiteCacheDir = "json-schema-test-suite-" + testSuiteVersion
)

// GetCacheDir returns the cache directory for xschema
func GetCacheDir() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		// Fallback to temp dir
		cacheDir = os.TempDir()
	}
	return filepath.Join(cacheDir, "xschema"), nil
}

// FetchTestSuite downloads the JSON Schema Test Suite if not cached
// Returns the path to the test suite directory
func FetchTestSuite(ctx context.Context) (string, error) {
	cacheDir, err := GetCacheDir()
	if err != nil {
		return "", fmt.Errorf("failed to get cache dir: %w", err)
	}

	suiteDir := filepath.Join(cacheDir, testSuiteCacheDir)
	testsDir := filepath.Join(suiteDir, "tests")

	// Check if already cached (has tests/draft2020-12 directory)
	if _, err := os.Stat(filepath.Join(testsDir, "draft2020-12")); err == nil {
		return suiteDir, nil
	}

	// Download and extract
	if err := downloadAndExtract(ctx, testSuiteURL, suiteDir); err != nil {
		return "", fmt.Errorf("failed to download test suite: %w", err)
	}

	return suiteDir, nil
}

func downloadAndExtract(ctx context.Context, url, destDir string) error {
	// Create parent directory
	if err := os.MkdirAll(filepath.Dir(destDir), 0755); err != nil {
		return fmt.Errorf("failed to create cache dir: %w", err)
	}

	// Download tarball
	downloadCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(downloadCtx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{Timeout: 2 * time.Minute}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status: %s", resp.Status)
	}

	// Extract gzip
	gzReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzReader.Close()

	// Extract tar
	tarReader := tar.NewReader(gzReader)

	// The tarball has a root directory like "JSON-Schema-Test-Suite-main/"
	// We want to extract its contents directly to destDir
	var rootDir string

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar: %w", err)
		}

		// Skip PAX headers (used by GitHub for metadata)
		if header.Typeflag == tar.TypeXGlobalHeader || header.Typeflag == tar.TypeXHeader {
			continue
		}

		// Get the root directory name from first real entry
		if rootDir == "" {
			parts := strings.SplitN(header.Name, "/", 2)
			if len(parts) > 0 {
				rootDir = parts[0]
			}
		}

		// Strip the root directory prefix
		relPath := strings.TrimPrefix(header.Name, rootDir+"/")
		if relPath == "" {
			continue // Skip root dir itself
		}

		targetPath := filepath.Join(destDir, relPath)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return fmt.Errorf("failed to create dir %s: %w", targetPath, err)
			}

		case tar.TypeReg:
			// Ensure parent directory exists
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return fmt.Errorf("failed to create parent dir for %s: %w", targetPath, err)
			}

			outFile, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", targetPath, err)
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return fmt.Errorf("failed to write file %s: %w", targetPath, err)
			}
			outFile.Close()
		}
	}

	return nil
}
