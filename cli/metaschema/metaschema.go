// Package metaschema provides fetching, caching, and vocabulary extraction for JSON Schema metaschemas.
package metaschema

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/xschemadev/xschema/fetcher"
)

const (
	defaultHTTPTimeout = 30 * time.Second
	userAgent          = "xschema-cli/1.0"
)

// Metaschema represents a fetched and parsed metaschema
type Metaschema struct {
	URI        string
	Raw        json.RawMessage
	Vocabulary map[string]bool // $vocabulary mapping: URI → required
}

// cache holds fetched metaschemas
var (
	cacheMu sync.RWMutex
	cache   = make(map[string]*Metaschema)
)

// HTTPClient allows mocking HTTP requests in tests
var HTTPClient = &http.Client{Timeout: defaultHTTPTimeout}

// Get fetches and caches a metaschema by URI.
// Returns cached result on subsequent calls for the same URI.
func Get(ctx context.Context, uri string) (*Metaschema, error) {
	return GetWithCache(ctx, uri, nil)
}

// GetWithCache fetches and caches a metaschema by URI, using the shared cache if provided.
// The shared cache is checked first for raw schema data, avoiding duplicate HTTP requests
// for metaschemas that were already fetched by retriever or processor.
func GetWithCache(ctx context.Context, uri string, sharedCache *fetcher.SharedCache) (*Metaschema, error) {
	// Check internal parsed cache first
	cacheMu.RLock()
	if m, ok := cache[uri]; ok {
		cacheMu.RUnlock()
		return m, nil
	}
	cacheMu.RUnlock()

	var raw json.RawMessage
	var err error

	// Check shared cache for raw schema data
	if sharedCache != nil {
		if cached, ok := sharedCache.Get(uri); ok {
			raw = cached
		}
	}

	// Fetch from network if not in shared cache
	if raw == nil {
		raw, err = fetchMetaschema(ctx, uri)
		if err != nil {
			return nil, err
		}

		// Store in shared cache if provided
		if sharedCache != nil {
			sharedCache.Set(uri, raw)
		}
	}

	// Parse to extract $vocabulary
	var parsed map[string]any
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return nil, fmt.Errorf("metaschema %s: invalid JSON: %w", uri, err)
	}

	vocab := ExtractVocabulary(parsed)

	m := &Metaschema{
		URI:        uri,
		Raw:        raw,
		Vocabulary: vocab,
	}

	// Store in internal parsed cache
	cacheMu.Lock()
	cache[uri] = m
	cacheMu.Unlock()

	return m, nil
}

// fetchMetaschema fetches a metaschema from a URL
func fetchMetaschema(ctx context.Context, uri string) (json.RawMessage, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for %s: %w", uri, err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch metaschema %s: %w", uri, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch metaschema %s: status %d", uri, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read metaschema %s: %w", uri, err)
	}

	if !json.Valid(data) {
		return nil, fmt.Errorf("metaschema %s: invalid JSON", uri)
	}

	return json.RawMessage(data), nil
}

// ExtractVocabulary extracts $vocabulary from a metaschema.
// Returns:
// - nil if no $vocabulary field (all keywords enabled by default)
// - map with all values false if $vocabulary is empty object {}
// - map with specified vocabulary URIs and their required status otherwise
func ExtractVocabulary(schema any) map[string]bool {
	obj, ok := schema.(map[string]any)
	if !ok {
		return nil
	}

	vocabRaw, hasVocab := obj["$vocabulary"]
	if !hasVocab {
		return nil // no $vocabulary = all enabled
	}

	vocabMap, ok := vocabRaw.(map[string]any)
	if !ok {
		return nil
	}

	result := make(map[string]bool, len(vocabMap))
	for uri, val := range vocabMap {
		// $vocabulary values should be booleans indicating if the vocabulary is required
		if required, ok := val.(bool); ok {
			result[uri] = required
		}
	}
	return result
}

// IsStandardDraft checks if a URI is a standard JSON Schema draft metaschema.
// Returns true for any json-schema.org URL.
func IsStandardDraft(uri string) bool {
	return strings.HasPrefix(uri, "http://json-schema.org/") ||
		strings.HasPrefix(uri, "https://json-schema.org/")
}

// ClearCache clears the metaschema cache (useful for testing)
func ClearCache() {
	cacheMu.Lock()
	cache = make(map[string]*Metaschema)
	cacheMu.Unlock()
}

// HasCustomMetaschema returns true if the schema uses a custom (non-standard) $schema URI.
func HasCustomMetaschema(data json.RawMessage) bool {
	var parsed struct {
		Schema string `json:"$schema"`
	}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return false
	}
	if parsed.Schema == "" {
		return false
	}
	return !IsStandardDraft(parsed.Schema)
}

// BuildMetaschemasMap extracts custom metaschemas from cache for validation.
// Returns a map of URI → schema data for non-standard $schema references found in schemas.
func BuildMetaschemasMap(schemas []json.RawMessage, cache fetcher.Cache) map[string]json.RawMessage {
	result := make(map[string]json.RawMessage)

	// Helper to extract and add custom $schema from a schema
	addMetaschema := func(data json.RawMessage) {
		var parsed map[string]any
		if err := json.Unmarshal(data, &parsed); err != nil {
			return
		}
		schemaURI, ok := parsed["$schema"].(string)
		if !ok || schemaURI == "" {
			return
		}
		if IsStandardDraft(schemaURI) {
			return
		}
		// Look up in cache
		normalized := fetcher.NormalizeURI(schemaURI)
		if metaData, ok := cache[normalized]; ok {
			result[schemaURI] = metaData
		}
	}

	// Check declared schemas
	for _, schema := range schemas {
		addMetaschema(schema)
	}

	// Check cached schemas (in case they reference custom metaschemas too)
	for _, data := range cache {
		addMetaschema(data)
	}

	return result
}
