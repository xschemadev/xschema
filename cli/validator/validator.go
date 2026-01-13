package validator

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// CompiledCache caches compiled schemas by content hash to avoid recompilation.
// Thread-safe for concurrent access.
type CompiledCache struct {
	mu    sync.RWMutex
	cache map[string]*jsonschema.Schema
	stats CacheStats
}

// CacheStats tracks cache statistics for testing/debugging
type CacheStats struct {
	Hits     int
	Misses   int
	Compiles int
}

// NewCompiledCache creates a new compiled schema cache.
func NewCompiledCache() *CompiledCache {
	return &CompiledCache{
		cache: make(map[string]*jsonschema.Schema),
	}
}

// contentHash computes a SHA256 hash of the schema content combined with
// draft hint and metaschema URIs for cache key uniqueness.
func contentHash(data []byte, draftHint string, metaschemaURIs []string) string {
	h := sha256.New()
	h.Write(data)
	h.Write([]byte(draftHint))
	for _, uri := range metaschemaURIs {
		h.Write([]byte(uri))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Get retrieves a compiled schema from cache by content hash.
// Returns (schema, true) if cached, (nil, false) if not.
func (c *CompiledCache) Get(key string) (*jsonschema.Schema, bool) {
	c.mu.RLock()
	schema, ok := c.cache[key]
	c.mu.RUnlock()

	if ok {
		c.mu.Lock()
		c.stats.Hits++
		c.mu.Unlock()
	}
	return schema, ok
}

// Set stores a compiled schema in the cache.
func (c *CompiledCache) Set(key string, schema *jsonschema.Schema) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = schema
}

// Stats returns a copy of the current cache statistics.
func (c *CompiledCache) Stats() CacheStats {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.stats
}

// recordMiss increments the miss counter.
func (c *CompiledCache) recordMiss() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.stats.Misses++
}

// recordCompile increments the compile counter.
func (c *CompiledCache) recordCompile() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.stats.Compiles++
}

// ValidateOptions configures schema validation behavior
type ValidateOptions struct {
	// Metaschemas maps URI to raw JSON for custom metaschemas.
	// Pre-loaded metaschemas allow validation of schemas using custom $schema URIs.
	Metaschemas map[string]json.RawMessage

	// Cache is an optional compiled schema cache for reusing compiled schemas.
	// When provided, schemas are cached by content hash and reused on subsequent validations.
	Cache *CompiledCache
}

// noopLoader returns an error for all URL loads.
// External refs are handled by the bundler, not the validator.
type noopLoader struct{}

func (noopLoader) Load(url string) (any, error) {
	return nil, fmt.Errorf("external ref %s (will be resolved by bundler)", url)
}

// draftNames maps draft string names to jsonschema draft constants
var draftNames = map[string]*jsonschema.Draft{
	"draft4":       jsonschema.Draft4,
	"draft6":       jsonschema.Draft6,
	"draft7":       jsonschema.Draft7,
	"draft2019-09": jsonschema.Draft2019,
	"draft2020-12": jsonschema.Draft2020,
}

// draftURLs maps $schema URLs to jsonschema draft constants
var draftURLs = map[string]*jsonschema.Draft{
	"https://json-schema.org/draft/2020-12/schema": jsonschema.Draft2020,
	"http://json-schema.org/draft/2020-12/schema":  jsonschema.Draft2020,
	"https://json-schema.org/draft/2019-09/schema": jsonschema.Draft2019,
	"http://json-schema.org/draft/2019-09/schema":  jsonschema.Draft2019,
	"http://json-schema.org/draft-07/schema#":      jsonschema.Draft7,
	"https://json-schema.org/draft-07/schema#":     jsonschema.Draft7,
	"http://json-schema.org/draft-06/schema#":      jsonschema.Draft6,
	"https://json-schema.org/draft-06/schema#":     jsonschema.Draft6,
	"http://json-schema.org/draft-04/schema#":      jsonschema.Draft4,
	"https://json-schema.org/draft-04/schema#":     jsonschema.Draft4,
}

// detectDraft detects the JSON Schema draft from the $schema field
func detectDraft(data []byte) *jsonschema.Draft {
	var schema struct {
		Schema string `json:"$schema"`
	}
	if err := json.Unmarshal(data, &schema); err != nil || schema.Schema == "" {
		return jsonschema.Draft2020 // default
	}

	// exact match
	if draft, ok := draftURLs[schema.Schema]; ok {
		return draft
	}

	// pattern match for variations
	s := schema.Schema
	if strings.Contains(s, "draft/2020-12") || strings.Contains(s, "draft-2020-12") {
		return jsonschema.Draft2020
	}
	if strings.Contains(s, "draft/2019-09") || strings.Contains(s, "draft-2019-09") {
		return jsonschema.Draft2019
	}
	if strings.Contains(s, "draft-07") {
		return jsonschema.Draft7
	}
	if strings.Contains(s, "draft-06") {
		return jsonschema.Draft6
	}
	if strings.Contains(s, "draft-04") {
		return jsonschema.Draft4
	}

	return jsonschema.Draft2020 // default to latest
}

// ValidateSchema validates that the given JSON bytes represent a valid JSON Schema.
// Returns nil if valid, error with details if invalid.
// External $ref loading errors are ignored (handled by bundler later).
// If draftHint is provided and schema has no $schema field, it will be used instead of defaulting to draft2020-12.
func ValidateSchema(data []byte, draftHint ...string) error {
	return ValidateSchemaWithOptions(data, nil, draftHint...)
}

// extractSchemaURL extracts the $schema URL from schema bytes
func extractSchemaURL(data []byte) string {
	var schema struct {
		Schema string `json:"$schema"`
	}
	if err := json.Unmarshal(data, &schema); err != nil {
		return ""
	}
	return schema.Schema
}

// isStandardMetaschema returns true if the URL is a known standard JSON Schema metaschema
func isStandardMetaschema(url string) bool {
	_, ok := draftURLs[url]
	return ok
}

// ValidateSchemaWithOptions validates a schema with configurable options.
// Use this to pre-load custom metaschemas for schemas with custom $schema URIs.
// If opts.Cache is provided, compiled schemas are cached by content hash.
func ValidateSchemaWithOptions(data []byte, opts *ValidateOptions, draftHint ...string) error {
	hint := ""
	if len(draftHint) > 0 {
		hint = draftHint[0]
	}

	// Build cache key from content, draft hint, and metaschema URIs
	var metaURIs []string
	if opts != nil && opts.Metaschemas != nil {
		metaURIs = make([]string, 0, len(opts.Metaschemas))
		for uri := range opts.Metaschemas {
			metaURIs = append(metaURIs, uri)
		}
		// Sort for deterministic hash
		sortStrings(metaURIs)
	}
	cacheKey := contentHash(data, hint, metaURIs)

	// Check cache
	if opts != nil && opts.Cache != nil {
		if _, ok := opts.Cache.Get(cacheKey); ok {
			return nil // Cache hit - schema already validated
		}
		opts.Cache.recordMiss()
	}

	var draft *jsonschema.Draft
	if hint != "" {
		if d, ok := draftNames[hint]; ok {
			draft = d
		} else {
			draft = detectDraft(data)
		}
	} else {
		draft = detectDraft(data)
	}

	// Parse JSON first using library's unmarshaler
	doc, err := jsonschema.UnmarshalJSON(strings.NewReader(string(data)))
	if err != nil {
		return fmt.Errorf("invalid JSON Schema: %w", err)
	}

	// The jsonschema library validates against the meta-schema during compilation
	compiler := jsonschema.NewCompiler()
	compiler.DefaultDraft(draft)
	// Use noop loader - external refs are handled by bundler
	compiler.UseLoader(noopLoader{})

	// Pre-load custom metaschemas via AddResource
	// Each metaschema is added by its URI so schemas using $schema can reference it
	if opts != nil && opts.Metaschemas != nil {
		for uri, raw := range opts.Metaschemas {
			metaDoc, err := jsonschema.UnmarshalJSON(strings.NewReader(string(raw)))
			if err != nil {
				return fmt.Errorf("invalid metaschema %s: %w", uri, err)
			}
			if err := compiler.AddResource(uri, metaDoc); err != nil {
				return fmt.Errorf("failed to add metaschema %s: %w", uri, err)
			}
		}
	}

	if err := compiler.AddResource("schema.json", doc); err != nil {
		return fmt.Errorf("invalid JSON Schema: %w", err)
	}

	schemaURL := extractSchemaURL(data)

	// Track compilation
	if opts != nil && opts.Cache != nil {
		opts.Cache.recordCompile()
	}

	compiled, err := compiler.Compile("schema.json")
	if err != nil {
		var loadErr *jsonschema.LoadURLError
		if errors.As(err, &loadErr) {
			// If error is for $schema URL and it's not a standard draft, fail
			// Custom metaschemas must be pre-loaded
			if loadErr.URL == schemaURL && !isStandardMetaschema(schemaURL) {
				return fmt.Errorf("custom metaschema %s not pre-loaded: %w", schemaURL, err)
			}
			// Ignore external ref loading errors - bundler handles these
			// Still cache as valid since the schema itself is fine
			if opts != nil && opts.Cache != nil {
				opts.Cache.Set(cacheKey, nil) // nil schema means "valid but had external refs"
			}
			return nil
		}
		return fmt.Errorf("invalid JSON Schema: %w", err)
	}

	// Cache successful compilation
	if opts != nil && opts.Cache != nil {
		opts.Cache.Set(cacheKey, compiled)
	}

	return nil
}

// sortStrings sorts a slice of strings in place
func sortStrings(s []string) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}
