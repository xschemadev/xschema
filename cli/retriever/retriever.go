package retriever

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/xschema/cli/parser"
	"github.com/xschema/cli/ui"
	"golang.org/x/sync/errgroup"
)

const (
	defaultRetries     = 3
	defaultConcurrency = 10
	defaultHTTPTimeout = 30 * time.Second
	retryBaseDelay     = 500 * time.Millisecond
	userAgent          = "xschema-cli/1.0"
)

// Options configures retrieval behavior
type Options struct {
	Concurrency int
	HTTPTimeout time.Duration
	Retries     int
	NoCache     bool
}

// DefaultOptions returns sensible defaults
func DefaultOptions() Options {
	return Options{
		Concurrency: defaultConcurrency,
		HTTPTimeout: defaultHTTPTimeout,
		Retries:     defaultRetries,
		NoCache:     false,
	}
}

// RetrievedSchema contains a fetched schema with its metadata
type RetrievedSchema struct {
	Namespace string
	ID        string
	Schema    json.RawMessage
	Adapter   string
}

// Key returns the full namespaced key like "namespace:id"
func (r RetrievedSchema) Key() string {
	return r.Namespace + ":" + r.ID
}

// schemaCache caches retrieved schemas
type schemaCache struct {
	mu    sync.RWMutex
	items map[string]json.RawMessage
}

func newSchemaCache() *schemaCache {
	return &schemaCache{items: make(map[string]json.RawMessage)}
}

func (c *schemaCache) get(key string) (json.RawMessage, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.items[key]
	return v, ok
}

func (c *schemaCache) set(key string, val json.RawMessage) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = val
}

// retrieveFromURL fetches a JSON schema from a URL with retry
func retrieveFromURL(ctx context.Context, url string, opts Options) (json.RawMessage, error) {
	client := &http.Client{Timeout: opts.HTTPTimeout}
	var lastErr error

	maxAttempts := opts.Retries
	if maxAttempts < 1 {
		maxAttempts = 1
	}

	ui.Verbosef("fetching from URL: %s (max_attempts: %d)", url, maxAttempts)

	for attempt := range maxAttempts {
		if attempt > 0 {
			ui.Verbosef("retrying request: url=%s, attempt=%d/%d", url, attempt+1, maxAttempts)
			delay := retryBaseDelay * time.Duration(1<<(attempt-1))
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(delay):
			}
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request for %s: %w", url, err)
		}
		req.Header.Set("User-Agent", userAgent)

		resp, err := client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("failed to fetch %s: %w", url, err)
			ui.Verbosef("HTTP request failed: url=%s, error=%v", url, err)
			continue
		}

		data, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			lastErr = fmt.Errorf("failed to read response from %s: %w", url, err)
			ui.Verbosef("failed to read response: url=%s, error=%v", url, err)
			continue
		}

		if resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("server error fetching %s: status %d", url, resp.StatusCode)
			ui.Verbosef("server error: url=%s, status=%d", url, resp.StatusCode)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch %s: status %d", url, resp.StatusCode)
		}

		if !json.Valid(data) {
			return nil, fmt.Errorf("invalid JSON from %s", url)
		}

		ui.Verbosef("successfully fetched from URL: url=%s, status=%d, bytes=%d", url, resp.StatusCode, len(data))
		return json.RawMessage(data), nil
	}

	return nil, lastErr
}

// retrieveFromFile reads a JSON schema from a file relative to the config file
func retrieveFromFile(ctx context.Context, filePath string, configPath string) (json.RawMessage, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Resolve path relative to config file's directory
	configDir := filepath.Dir(configPath)
	fullPath := filepath.Join(configDir, filePath)

	ui.Verbosef("reading file: %s (relative to %s)", fullPath, configDir)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		ui.Verbosef("failed to read file: path=%s, error=%v", fullPath, err)
		return nil, fmt.Errorf("failed to read %s: %w", fullPath, err)
	}

	if !json.Valid(data) {
		ui.Verbosef("invalid JSON in file: %s", fullPath)
		return nil, fmt.Errorf("invalid JSON in %s", fullPath)
	}

	ui.Verbosef("successfully read file: path=%s, bytes=%d", fullPath, len(data))
	return json.RawMessage(data), nil
}

// Retrieve fetches all schemas from declarations
func Retrieve(ctx context.Context, decls []parser.Declaration, opts Options) ([]RetrievedSchema, error) {
	if len(decls) == 0 {
		return nil, nil
	}

	var cache *schemaCache
	if !opts.NoCache {
		cache = newSchemaCache()
	}

	results := make([]RetrievedSchema, len(decls))

	ui.Verbosef("retrieving schemas: count=%d, concurrency=%d, cache_enabled=%v", len(decls), opts.Concurrency, cache != nil)

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(opts.Concurrency)

	for i, decl := range decls {
		idx, d := i, decl

		// Build cache key based on source type
		var cacheKey string
		switch d.SourceType {
		case parser.SourceURL:
			var url string
			json.Unmarshal(d.Source, &url)
			cacheKey = "url:" + url
		case parser.SourceFile:
			var filePath string
			json.Unmarshal(d.Source, &filePath)
			cacheKey = "file:" + filepath.Join(filepath.Dir(d.ConfigPath), filePath)
		case parser.SourceJSON:
			// Inline JSON - use the declaration key as cache key
			cacheKey = "json:" + d.Key()
		}

		// Check cache first (if enabled)
		if cache != nil {
			if cached, ok := cache.get(cacheKey); ok {
				ui.Verbosef("cache hit: schema=%s, key=%s", d.Key(), cacheKey)
				results[idx] = RetrievedSchema{
					Namespace: d.Namespace,
					ID:        d.ID,
					Schema:    cached,
					Adapter:   d.Adapter,
				}
				continue
			}
			ui.Verbosef("cache miss: schema=%s, key=%s", d.Key(), cacheKey)
		}

		g.Go(func() error {
			var schema json.RawMessage
			var err error

			switch d.SourceType {
			case parser.SourceURL:
				var url string
				if err := json.Unmarshal(d.Source, &url); err != nil {
					return fmt.Errorf("invalid URL source for %s: %w", d.Key(), err)
				}
				schema, err = retrieveFromURL(ctx, url, opts)
			case parser.SourceFile:
				var filePath string
				if err := json.Unmarshal(d.Source, &filePath); err != nil {
					return fmt.Errorf("invalid file source for %s: %w", d.Key(), err)
				}
				schema, err = retrieveFromFile(ctx, filePath, d.ConfigPath)
			case parser.SourceJSON:
				// Inline JSON - source is already the schema
				schema = d.Source
			default:
				err = fmt.Errorf("unknown source type: %s", d.SourceType)
			}

			if err != nil {
				ui.Verbosef("failed to retrieve schema: key=%s, source=%s, error=%v", d.Key(), d.SourceType, err)
				return fmt.Errorf("failed to retrieve schema %s: %w", d.Key(), err)
			}

			if cache != nil {
				cache.set(cacheKey, schema)
			}

			results[idx] = RetrievedSchema{
				Namespace: d.Namespace,
				ID:        d.ID,
				Schema:    schema,
				Adapter:   d.Adapter,
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	ui.Verbosef("retrieval complete: schemas=%d", len(results))
	return results, nil
}

// GroupByAdapter groups retrieved schemas by adapter package
func GroupByAdapter(schemas []RetrievedSchema) map[string][]RetrievedSchema {
	groups := make(map[string][]RetrievedSchema)
	for _, s := range schemas {
		groups[s.Adapter] = append(groups[s.Adapter], s)
	}
	return groups
}

// SortedAdapters returns adapter keys in sorted order for deterministic output
func SortedAdapters(groups map[string][]RetrievedSchema) []string {
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
