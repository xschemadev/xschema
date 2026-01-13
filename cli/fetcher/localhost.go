package fetcher

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// LocalhostFetcher implements Fetcher for compliance testing.
// It maps http://localhost:1234/* URLs to local files in a remotes/ directory.
//
// The JSON Schema Test Suite has schemas that reference URLs in several ways:
//
//  1. localhost:1234/* URLs - These map to files in the remotes/ directory.
//     Example: http://localhost:1234/integer.json → remotes/integer.json
//
//  2. example.com, urn:, etc URLs - These appear in test schemas as $ids that
//     "claim" those URLs locally. The schema defines the content inline.
//     Example: A schema with $id:"http://example.com/foo.json" and a nested
//     $ref:"http://example.com/foo.json" should resolve to itself, not fetch.
//
//  3. localhost URLs without files - Some test schemas define $id like
//     "http://localhost:1234/draft2020-12/tree" but no tree file exists.
//     The schema claims that URL with a local $id, so the bundler won't fetch.
//
// The processor's crawl phase extracts ALL external-looking refs and tries to
// fetch them upfront. It doesn't understand local $ids - that's the bundler's
// job. So crawl may try to fetch URLs that the bundler would skip.
//
// Our solution: return stub schemas for unfetchable URLs. The bundler checks
// local $ids first, so stubs for locally-defined refs are never used. For
// truly missing external refs, the bundler will fail when it can't resolve.
type LocalhostFetcher struct {
	RemotesPath string // path to the remotes directory containing schema files
}

// NewLocalhostFetcher creates a LocalhostFetcher for the given remotes directory.
func NewLocalhostFetcher(remotesPath string) *LocalhostFetcher {
	return &LocalhostFetcher{RemotesPath: remotesPath}
}

const localhostPrefix = "http://localhost:1234/"

// Fetch implements Fetcher for compliance testing.
func (f *LocalhostFetcher) Fetch(_ context.Context, uri string) (json.RawMessage, error) {
	// Non-localhost URLs (example.com, urn:, etc) are never real external refs
	// in the test suite - they're always claimed by local $ids.
	if !strings.HasPrefix(uri, localhostPrefix) {
		return json.RawMessage(`{}`), nil
	}

	// Try to read from remotes directory
	path := strings.TrimPrefix(uri, localhostPrefix)
	if idx := strings.Index(path, "?"); idx != -1 {
		path = path[:idx]
	}

	localPath := filepath.Join(f.RemotesPath, path)
	data, err := os.ReadFile(localPath)
	if err != nil {
		// File doesn't exist or is a directory - probably a localhost URL claimed
		// by local $id. Return stub; bundler will resolve to local $id and never use this.
		return json.RawMessage(`{}`), nil
	}

	return json.RawMessage(data), nil
}
