package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// Fetcher retrieves schemas by URI.
type Fetcher interface {
	Fetch(ctx context.Context, uri string) (json.RawMessage, error)
}

// FetchFunc adapts a function to the Fetcher interface.
type FetchFunc func(ctx context.Context, uri string) (json.RawMessage, error)

func (f FetchFunc) Fetch(ctx context.Context, uri string) (json.RawMessage, error) {
	return f(ctx, uri)
}

// Cache is a URI → schema cache used by CacheFetcher.
type Cache map[string]json.RawMessage

// CacheFetcher implements Fetcher by looking up schemas in a pre-populated cache.
type CacheFetcher struct {
	cache Cache
}

// NewCacheFetcher creates a CacheFetcher from a cache.
func NewCacheFetcher(cache Cache) *CacheFetcher {
	return &CacheFetcher{cache: cache}
}

// Fetch returns the cached schema for the given URI.
func (f *CacheFetcher) Fetch(ctx context.Context, uri string) (json.RawMessage, error) {
	normalized := NormalizeURI(uri)
	if data, ok := f.cache[normalized]; ok {
		return data, nil
	}
	return nil, fmt.Errorf("cache miss for %q: URI was not pre-fetched", uri)
}

// NormalizeURI normalizes a URI for cache lookups (strips fragment, normalizes scheme).
func NormalizeURI(uri string) string {
	if idx := strings.Index(uri, "#"); idx >= 0 {
		uri = uri[:idx]
	}
	return uri
}

// SharedCache is a thread-safe cache for schema data shared across retriever, processor, and metaschema.
// It ensures no URI is fetched more than once regardless of how it's referenced.
type SharedCache struct {
	mu    sync.RWMutex
	items map[string]json.RawMessage
}

// NewSharedCache creates a new empty SharedCache.
func NewSharedCache() *SharedCache {
	return &SharedCache{items: make(map[string]json.RawMessage)}
}

// Get retrieves a schema from cache by URI. Returns nil, false if not found.
func (c *SharedCache) Get(uri string) (json.RawMessage, bool) {
	normalized := NormalizeURI(uri)
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.items[normalized]
	return v, ok
}

// Set stores a schema in cache by URI.
func (c *SharedCache) Set(uri string, data json.RawMessage) {
	normalized := NormalizeURI(uri)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[normalized] = data
}

// Has checks if a URI exists in cache.
func (c *SharedCache) Has(uri string) bool {
	normalized := NormalizeURI(uri)
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.items[normalized]
	return ok
}

// ToCache returns a snapshot of the cache as a plain Cache map.
// Used when passing to CacheFetcher for bundling phase.
func (c *SharedCache) ToCache() Cache {
	c.mu.RLock()
	defer c.mu.RUnlock()
	result := make(Cache, len(c.items))
	for k, v := range c.items {
		result[k] = v
	}
	return result
}

// Len returns the number of items in cache.
func (c *SharedCache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}
