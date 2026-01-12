// Package processor handles the crawl-fetch-validate-bundle pipeline.
// It sits between Retriever and Generator, iteratively discovering and fetching
// external refs before bundling schemas into self-contained units.
package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/xschemadev/xschema/bundler"
	"github.com/xschemadev/xschema/metaschema"
	"github.com/xschemadev/xschema/retriever"
	"github.com/xschemadev/xschema/ui"
	"github.com/xschemadev/xschema/validator"
)

// ProcessedSchema contains a fully processed schema ready for code generation.
type ProcessedSchema struct {
	Namespace  string          // namespace from config
	ID         string          // schema ID from config
	Schema     json.RawMessage // bundled schema (self-contained)
	Adapter    string          // adapter package ref
	SourceURI  string          // original source URI
	Vocabulary map[string]bool // $vocabulary from custom metaschema (nil = all enabled)
}

// Key returns the full namespaced key like "namespace:id"
func (p ProcessedSchema) Key() string {
	return p.Namespace + ":" + p.ID
}

// Options configures processing behavior.
type Options struct {
	Fetcher   Fetcher           // fetcher for external refs (required)
	OnVerbose func(msg string)  // callback for verbose logging (optional)
}

// Fetcher retrieves schemas by URI.
type Fetcher interface {
	Fetch(ctx context.Context, uri string) (json.RawMessage, error)
}

// FetchFunc adapts a function to the Fetcher interface.
type FetchFunc func(ctx context.Context, uri string) (json.RawMessage, error)

// Fetch implements Fetcher.
func (f FetchFunc) Fetch(ctx context.Context, uri string) (json.RawMessage, error) {
	return f(ctx, uri)
}

// Process runs the full processing pipeline on retrieved schemas:
// 1. crawlAndFetch - iteratively discover and fetch external refs
// 2. validateAll - validate all schemas (declared + external) [future US-011]
// 3. bundleAll - bundle each declared schema using cache [future US-012]
func Process(ctx context.Context, schemas []retriever.RetrievedSchema, opts Options) ([]ProcessedSchema, error) {
	if opts.Fetcher == nil {
		return nil, fmt.Errorf("processor: Fetcher is required")
	}

	verbose := func(msg string) {
		if opts.OnVerbose != nil {
			opts.OnVerbose(msg)
		}
	}

	verbose(fmt.Sprintf("processor: starting with %d schemas", len(schemas)))

	// Phase 1: Crawl and fetch all external refs
	cache, err := crawlAndFetch(ctx, schemas, opts.Fetcher, verbose)
	if err != nil {
		return nil, err
	}

	verbose(fmt.Sprintf("processor: crawl complete, cache has %d external schemas", len(cache)))

	// Phase 2: Validate all schemas (declared + external)
	if err := validateAll(schemas, cache, verbose); err != nil {
		return nil, err
	}

	verbose("processor: validation complete")

	// Phase 3: Bundle all schemas using the cache
	result, err := bundleAll(ctx, schemas, cache, verbose)
	if err != nil {
		return nil, err
	}

	verbose(fmt.Sprintf("processor: bundling complete, produced %d schemas", len(result)))

	return result, nil
}

// externalRefCache maps normalized URI → raw schema bytes
type externalRefCache map[string]json.RawMessage

// validateAll validates all declared schemas and external schemas from cache.
// Runs after crawlAndFetch to ensure all schemas are valid before bundling.
func validateAll(schemas []retriever.RetrievedSchema, cache externalRefCache, verbose func(string)) error {
	// Validate declared schemas
	for _, s := range schemas {
		if err := validator.ValidateSchema(s.Schema); err != nil {
			return fmt.Errorf("validation failed for %s: %w", s.SourceURI, err)
		}
	}

	verbose(fmt.Sprintf("processor: validated %d declared schemas", len(schemas)))

	// Validate external schemas from cache
	for uri, data := range cache {
		if err := validator.ValidateSchema(data); err != nil {
			return fmt.Errorf("validation failed for external schema %s: %w", uri, err)
		}
	}

	verbose(fmt.Sprintf("processor: validated %d external schemas", len(cache)))

	return nil
}

// bundleAll bundles each declared schema using the CacheFetcher.
// No I/O occurs during bundling - all refs resolve from the pre-populated cache.
func bundleAll(ctx context.Context, schemas []retriever.RetrievedSchema, cache externalRefCache, verbose func(string)) ([]ProcessedSchema, error) {
	// Add declared schemas to cache so bundler can resolve circular refs back to them
	for _, s := range schemas {
		if s.SourceURI != "" {
			normalized := normalizeURI(s.SourceURI)
			cache[normalized] = s.Schema
		}
	}

	fetcher := NewCacheFetcher(cache)
	result := make([]ProcessedSchema, len(schemas))

	for i, s := range schemas {
		verbose(fmt.Sprintf("processor: bundling %s", s.SourceURI))

		bundled, err := bundler.Bundle(ctx, bundler.BundleInput{
			Schema:    s.Schema,
			SourceURI: s.SourceURI,
			Fetcher:   fetcher,
			Draft:     "", // let bundler detect draft from $schema
		})
		if err != nil {
			return nil, fmt.Errorf("bundling failed for %s: %w", s.SourceURI, err)
		}

		// Extract vocabulary from custom metaschema if present
		vocab, err := extractVocabulary(ctx, bundled)
		if err != nil {
			// Non-fatal: log and continue without vocabulary
			ui.Verbosef("processor: could not extract vocabulary for %s: %v", s.SourceURI, err)
		}

		result[i] = ProcessedSchema{
			Namespace:  s.Namespace,
			ID:         s.ID,
			Schema:     bundled,
			Adapter:    s.Adapter,
			SourceURI:  s.SourceURI,
			Vocabulary: vocab,
		}
	}

	return result, nil
}

// extractVocabulary extracts $vocabulary from a bundled schema's custom metaschema.
// Returns nil if schema uses standard draft or metaschema can't be fetched.
func extractVocabulary(ctx context.Context, schema json.RawMessage) (map[string]bool, error) {
	var parsed map[string]any
	if err := json.Unmarshal(schema, &parsed); err != nil {
		return nil, nil // not an object schema, no vocabulary
	}

	schemaURI, ok := parsed["$schema"].(string)
	if !ok || schemaURI == "" {
		return nil, nil // no $schema, use defaults
	}

	// Skip standard drafts - they have well-known vocabularies
	if metaschema.IsStandardDraft(schemaURI) {
		return nil, nil
	}

	// Check if vocabulary already embedded in schema (by bundler)
	if vocab, ok := parsed["$vocabulary"].(map[string]any); ok {
		result := make(map[string]bool, len(vocab))
		for uri, val := range vocab {
			if required, ok := val.(bool); ok {
				result[uri] = required
			}
		}
		return result, nil
	}

	// Fetch custom metaschema and extract vocabulary
	meta, err := metaschema.Get(ctx, schemaURI)
	if err != nil {
		return nil, err
	}

	return meta.Vocabulary, nil
}

// crawlAndFetch iteratively discovers external $refs in schemas and fetches them.
// It continues until no new URIs are found, building a complete cache of external schemas.
// Returns error immediately on any fetch failure (fail fast).
func crawlAndFetch(ctx context.Context, schemas []retriever.RetrievedSchema, fetcher Fetcher, verbose func(string)) (externalRefCache, error) {
	cache := make(externalRefCache)
	visited := make(map[string]bool) // tracks URIs we've processed (including declared schemas)

	// Mark declared schema source URIs as visited (they're already fetched)
	for _, s := range schemas {
		if s.SourceURI != "" {
			normalized := normalizeURI(s.SourceURI)
			visited[normalized] = true
		}
	}

	// Initial frontier: refs from declared schemas
	frontier := make(map[string]string) // URI → base URI for resolution
	for _, s := range schemas {
		refs := extractExternalRefs(s.Schema, s.SourceURI)
		for _, ref := range refs {
			normalized := normalizeURI(ref)
			if !visited[normalized] {
				frontier[ref] = s.SourceURI
			}
		}
	}

	iteration := 0
	for len(frontier) > 0 {
		iteration++

		// Check context
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		verbose(fmt.Sprintf("processor: crawl iteration %d, %d new refs to fetch", iteration, len(frontier)))

		// Fetch all URIs in current frontier
		newFrontier := make(map[string]string)
		for uri, baseURI := range frontier {
			// Resolve relative URI against base
			resolved, err := resolveURI(uri, baseURI)
			if err != nil {
				return nil, fmt.Errorf("processor: failed to resolve %q against %q: %w", uri, baseURI, err)
			}

			normalized := normalizeURI(resolved)
			if visited[normalized] {
				continue
			}
			visited[normalized] = true

			ui.Verbosef("processor: fetching %s", resolved)

			raw, err := fetcher.Fetch(ctx, resolved)
			if err != nil {
				return nil, fmt.Errorf("processor: failed to fetch %q: %w", resolved, err)
			}

			cache[normalized] = raw

			// Extract refs from this schema and add to next frontier
			newRefs := extractExternalRefs(raw, resolved)
			for _, ref := range newRefs {
				refNorm := normalizeURI(ref)
				if !visited[refNorm] {
					newFrontier[ref] = resolved
				}
			}
		}

		frontier = newFrontier
	}

	verbose(fmt.Sprintf("processor: crawl finished in %d iterations", iteration))
	return cache, nil
}

// extractExternalRefs finds all external $ref URIs in a schema.
// Returns only external refs (not fragment-only refs like #/$defs/Foo).
func extractExternalRefs(data json.RawMessage, baseURI string) []string {
	var parsed any
	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil
	}
	return extractRefsFromNode(parsed, baseURI)
}

// extractRefsFromNode recursively extracts external refs from a parsed schema.
func extractRefsFromNode(node any, baseURI string) []string {
	var refs []string

	switch v := node.(type) {
	case map[string]any:
		// Check for $id which changes base URI
		if id, ok := v["$id"].(string); ok {
			if resolved, err := resolveURI(id, baseURI); err == nil {
				baseURI = resolved
			}
		} else if id, ok := v["id"].(string); ok {
			if resolved, err := resolveURI(id, baseURI); err == nil {
				baseURI = resolved
			}
		}

		// Check for $ref
		if ref, ok := v["$ref"].(string); ok {
			if isExternalRef(ref) {
				// Resolve relative ref against base URI
				resolved, err := resolveURI(ref, baseURI)
				if err == nil {
					// Strip fragment for fetching
					if idx := strings.Index(resolved, "#"); idx >= 0 {
						resolved = resolved[:idx]
					}
					if resolved != "" {
						refs = append(refs, resolved)
					}
				}
			}
		}

		// Recurse into all values
		for _, val := range v {
			refs = append(refs, extractRefsFromNode(val, baseURI)...)
		}

	case []any:
		for _, val := range v {
			refs = append(refs, extractRefsFromNode(val, baseURI)...)
		}
	}

	return refs
}

// isExternalRef returns true if the ref points to an external schema (not fragment-only).
func isExternalRef(ref string) bool {
	// Fragment-only refs are local
	if strings.HasPrefix(ref, "#") {
		return false
	}
	return true
}

// resolveURI resolves a potentially relative URI against a base URI.
func resolveURI(ref, base string) (string, error) {
	if base == "" {
		return ref, nil
	}

	// Parse base
	baseURL, err := url.Parse(base)
	if err != nil {
		// Base might be a file path
		if filepath.IsAbs(ref) {
			return ref, nil
		}
		baseDir := filepath.Dir(base)
		return filepath.Join(baseDir, ref), nil
	}

	// Parse ref
	refURL, err := url.Parse(ref)
	if err != nil {
		return "", err
	}

	// If ref is absolute, use it directly
	if refURL.IsAbs() {
		return ref, nil
	}

	// Resolve relative URL
	resolved := baseURL.ResolveReference(refURL)
	return resolved.String(), nil
}

// normalizeURI normalizes a URI for use as cache key.
func normalizeURI(uri string) string {
	// Strip fragment
	if idx := strings.Index(uri, "#"); idx >= 0 {
		uri = uri[:idx]
	}
	return uri
}

// CacheFetcher implements bundler.Fetcher by looking up schemas in a pre-populated cache.
// It's used during the bundle phase to resolve external refs without network I/O.
type CacheFetcher struct {
	cache externalRefCache
}

// NewCacheFetcher creates a CacheFetcher from an externalRefCache.
func NewCacheFetcher(cache externalRefCache) *CacheFetcher {
	return &CacheFetcher{cache: cache}
}

// Fetch implements bundler.Fetcher. Returns the cached schema for the given URI.
// Returns an error if URI is not in cache (indicates crawler bug - all refs should be pre-fetched).
func (f *CacheFetcher) Fetch(uri string) (json.RawMessage, error) {
	normalized := normalizeURI(uri)
	if data, ok := f.cache[normalized]; ok {
		return data, nil
	}
	return nil, fmt.Errorf("cache miss for %q: URI was not pre-fetched by crawler (this is a bug)", uri)
}
