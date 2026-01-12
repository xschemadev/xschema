package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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
