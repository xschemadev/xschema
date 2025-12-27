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

	"github.com/xschema/cli/generator"
	"github.com/xschema/cli/logger"
	"github.com/xschema/cli/parser"
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

// adapterGroup groups schemas by adapter
type adapterGroup struct {
	adapter  string
	language string
	schemas  []generator.GenerateInput
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

	logger.Debug("fetching from URL", "url", url, "max_attempts", maxAttempts)

	for attempt := range maxAttempts {
		if attempt > 0 {
			logger.Debug("retrying request", "url", url, "attempt", attempt+1, "max", maxAttempts)
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
			logger.Warn("HTTP request failed", "url", url, "error", err)
			continue
		}

		data, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			lastErr = fmt.Errorf("failed to read response from %s: %w", url, err)
			logger.Warn("failed to read response", "url", url, "error", err)
			continue
		}

		if resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("server error fetching %s: status %d", url, resp.StatusCode)
			logger.Warn("server error", "url", url, "status", resp.StatusCode)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch %s: status %d", url, resp.StatusCode)
		}

		if !json.Valid(data) {
			return nil, fmt.Errorf("invalid JSON from %s", url)
		}

		logger.Debug("successfully fetched from URL", "url", url, "status", resp.StatusCode, "bytes", len(data))
		return json.RawMessage(data), nil
	}

	return nil, lastErr
}

// retrieveFromFile reads a JSON schema from a file relative to the declaration file
func retrieveFromFile(ctx context.Context, file string, declPath string) (json.RawMessage, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	dir := filepath.Dir(declPath)
	fullPath := filepath.Join(dir, file)

	logger.Debug("reading file", "path", fullPath)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		logger.Error("failed to read file", "path", fullPath, "error", err)
		return nil, fmt.Errorf("failed to read %s: %w", fullPath, err)
	}

	if !json.Valid(data) {
		logger.Error("invalid JSON in file", "path", fullPath)
		return nil, fmt.Errorf("invalid JSON in %s", fullPath)
	}

	logger.Debug("successfully read file", "path", fullPath, "bytes", len(data))
	return json.RawMessage(data), nil
}

// Retrieve fetches all schemas from declarations and groups them by adapter
func Retrieve(ctx context.Context, decls []parser.Declaration, opts Options) ([]generator.GenerateBatchInput, error) {
	if len(decls) == 0 {
		return nil, nil
	}

	var cache *schemaCache
	if !opts.NoCache {
		cache = newSchemaCache()
	}
	results := make([]json.RawMessage, len(decls))

	logger.Info("retrieving schemas", "count", len(decls), "concurrency", opts.Concurrency, "cache_enabled", cache != nil)

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(opts.Concurrency)

	for i, decl := range decls {
		idx, d := i, decl

		// Build cache key
		cacheKey := string(d.Source) + ":" + d.Location
		if d.Source == "file" {
			cacheKey = "file:" + filepath.Join(filepath.Dir(d.File), d.Location)
		}

		// Check cache first (if enabled)
		if cache != nil {
			if cached, ok := cache.get(cacheKey); ok {
				logger.Debug("cache hit", "schema", d.Name, "key", cacheKey)
				results[idx] = cached
				continue
			}
			logger.Debug("cache miss", "schema", d.Name, "key", cacheKey)
		}

		g.Go(func() error {
			var schema json.RawMessage
			var err error

			switch d.Source {
			case "url":
				schema, err = retrieveFromURL(ctx, d.Location, opts)
			case "file":
				schema, err = retrieveFromFile(ctx, d.Location, d.File)
			default:
				err = fmt.Errorf("unknown source type: %s", d.Source)
			}

			if err != nil {
				logger.Error("failed to retrieve schema", "schema", d.Name, "source", d.Source, "location", d.Location, "error", err)
				return fmt.Errorf("failed to retrieve schema %s: %w", d.Name, err)
			}

			if cache != nil {
				cache.set(cacheKey, schema)
			}
			results[idx] = schema
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Group by adapter
	groups := make(map[string]*adapterGroup)
	for i, decl := range decls {
		adapterKey := decl.Adapter.Package
		if adapterKey == "" {
			adapterKey = decl.Adapter.Name
		}

		group, ok := groups[adapterKey]
		if !ok {
			group = &adapterGroup{
				adapter:  adapterKey,
				language: decl.Adapter.Language,
			}
			groups[adapterKey] = group
		}
		group.schemas = append(group.schemas, generator.GenerateInput{
			Name:   decl.Name,
			Schema: results[i],
		})
	}

	// Sort keys for deterministic output
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	batches := make([]generator.GenerateBatchInput, 0, len(groups))
	for _, k := range keys {
		g := groups[k]
		batches = append(batches, generator.GenerateBatchInput{
			Schemas:  g.schemas,
			Adapter:  g.adapter,
			Language: g.language,
		})
		logger.Info("grouped schemas by adapter", "adapter", g.adapter, "language", g.language, "schemas", len(g.schemas))
	}

	logger.Info("retrieval complete", "batches", len(batches))
	return batches, nil
}
