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
	"github.com/xschemadev/xschema/fetcher"
	"github.com/xschemadev/xschema/metaschema"
	"github.com/xschemadev/xschema/retriever"
	"github.com/xschemadev/xschema/ui"
	"github.com/xschemadev/xschema/validator"
	"github.com/xschemadev/xschema/vocabulary"
)

// ProcessedSchema contains a fully processed schema ready for code generation.
type ProcessedSchema struct {
	Namespace string          // namespace from config
	ID        string          // schema ID from config
	Schema    json.RawMessage // bundled schema (self-contained, filtered by vocabulary)
	Adapter   string          // adapter package ref
	SourceURI string          // original source URI
}

// Key returns the full namespaced key like "namespace:id"
func (p ProcessedSchema) Key() string {
	return p.Namespace + ":" + p.ID
}

// Options configures processing behavior.
type Options struct {
	Fetcher   fetcher.Fetcher       // fetcher for external refs (required)
	OnVerbose func(msg string)      // callback for verbose logging (optional)
	Cache     *fetcher.SharedCache  // shared cache (optional, creates internal if nil)
}

// Process runs the full processing pipeline on retrieved schemas:
// 1. validateDeclared - validate declared schemas before any fetching (fail fast)
// 2. crawlAndFetch - iteratively discover and fetch external refs
// 3. validateExternal - validate external schemas after crawl
// 4. bundleAll - bundle each declared schema using cache
func Process(ctx context.Context, schemas []retriever.RetrievedSchema, opts Options) ([]ProcessedSchema, error) {
	if opts.Fetcher == nil {
		return nil, fmt.Errorf("processor: Fetcher is required")
	}

	verbose := func(msg string) {
		if opts.OnVerbose != nil {
			opts.OnVerbose(msg)
		}
	}

	// Use provided cache or create a new one
	cache := opts.Cache
	if cache == nil {
		cache = fetcher.NewSharedCache()
	}

	verbose(fmt.Sprintf("processor: starting with %d schemas", len(schemas)))

	// Phase 1: Validate declared schemas early (fail fast before fetching)
	if err := validateDeclared(schemas, verbose); err != nil {
		return nil, err
	}

	verbose(fmt.Sprintf("processor: declared schemas validated (%d schemas)", len(schemas)))

	// Phase 2: Crawl and fetch all external refs
	if err := crawlAndFetch(ctx, schemas, opts.Fetcher, cache, verbose); err != nil {
		return nil, err
	}

	verbose(fmt.Sprintf("processor: crawl complete, cache has %d external schemas", cache.Len()))

	// Phase 3: Validate external schemas
	if err := validateExternal(schemas, cache, verbose); err != nil {
		return nil, err
	}

	verbose("processor: validation complete")

	// Phase 4: Bundle all schemas using the cache
	result, err := bundleAll(ctx, schemas, cache, verbose)
	if err != nil {
		return nil, err
	}

	verbose(fmt.Sprintf("processor: bundling complete, produced %d schemas", len(result)))

	return result, nil
}

// validateDeclared validates declared schemas early, before any external fetching.
// This enables fail-fast behavior: invalid schemas are caught before wasting time
// fetching external refs. Schemas with custom metaschemas ($schema pointing to
// non-standard URI) are skipped here - they'll be validated after crawl when the
// custom metaschema is available.
func validateDeclared(schemas []retriever.RetrievedSchema, verbose func(string)) error {
	for _, s := range schemas {
		// Skip schemas with custom metaschemas (they need the metaschema fetched first)
		if hasCustomMetaschema(s.Schema) {
			continue
		}
		if err := validator.ValidateSchema(s.Schema); err != nil {
			return fmt.Errorf("validation failed for %s: %w", s.SourceURI, err)
		}
	}
	return nil
}

// hasCustomMetaschema returns true if the schema uses a custom (non-standard) $schema URI.
func hasCustomMetaschema(data json.RawMessage) bool {
	var parsed struct {
		Schema string `json:"$schema"`
	}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return false
	}
	if parsed.Schema == "" {
		return false
	}
	return !metaschema.IsStandardDraft(parsed.Schema)
}

// validateExternal validates external schemas and declared schemas with custom metaschemas.
// Runs after crawlAndFetch when custom metaschemas are available in the cache.
// Custom metaschemas from cache are passed to the validator so schemas using
// custom $schema URIs can be validated.
func validateExternal(schemas []retriever.RetrievedSchema, cache *fetcher.SharedCache, verbose func(string)) error {
	// Get snapshot of cache for iteration
	cacheSnapshot := cache.ToCache()

	// Build metaschemas map from cache for custom $schema validation
	metaschemas := buildMetaschemasMap(schemas, cacheSnapshot)

	opts := &validator.ValidateOptions{
		Metaschemas: metaschemas,
	}

	// Validate declared schemas with custom metaschemas (skipped in early validation)
	for _, s := range schemas {
		if hasCustomMetaschema(s.Schema) {
			if err := validator.ValidateSchemaWithOptions(s.Schema, opts); err != nil {
				return fmt.Errorf("validation failed for %s: %w", s.SourceURI, err)
			}
		}
	}

	// Validate external schemas from cache
	for uri, data := range cacheSnapshot {
		if err := validator.ValidateSchemaWithOptions(data, opts); err != nil {
			return fmt.Errorf("validation failed for external schema %s: %w", uri, err)
		}
	}

	verbose(fmt.Sprintf("processor: validated %d external schemas", len(cacheSnapshot)))

	return nil
}

// buildMetaschemasMap extracts custom metaschemas from cache for validation.
// Returns a map of URI → schema data for non-standard $schema references.
func buildMetaschemasMap(schemas []retriever.RetrievedSchema, cache fetcher.Cache) map[string]json.RawMessage {
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
		if metaschema.IsStandardDraft(schemaURI) {
			return
		}
		// Look up in cache
		normalized := fetcher.NormalizeURI(schemaURI)
		if metaData, ok := cache[normalized]; ok {
			result[schemaURI] = metaData
		}
	}

	// Check declared schemas
	for _, s := range schemas {
		addMetaschema(s.Schema)
	}

	// Check cached schemas (in case they reference custom metaschemas too)
	for _, data := range cache {
		addMetaschema(data)
	}

	return result
}

// bundleAll bundles each declared schema using a cache fetcher.
// No I/O occurs during bundling - all refs resolve from the pre-populated cache.
func bundleAll(ctx context.Context, schemas []retriever.RetrievedSchema, cache *fetcher.SharedCache, verbose func(string)) ([]ProcessedSchema, error) {
	// Add declared schemas to cache so bundler can resolve circular refs back to them
	for _, s := range schemas {
		if s.SourceURI != "" {
			cache.Set(s.SourceURI, s.Schema)
		}
	}

	// Create CacheFetcher from snapshot for bundling
	cacheFetcher := fetcher.NewCacheFetcher(cache.ToCache())
	result := make([]ProcessedSchema, len(schemas))

	for i, s := range schemas {
		verbose(fmt.Sprintf("processor: bundling %s", s.SourceURI))

		bundled, err := bundler.Bundle(ctx, bundler.BundleInput{
			Schema:    s.Schema,
			SourceURI: s.SourceURI,
			Fetcher:   cacheFetcher,
			Draft:     "", // let bundler detect draft from $schema
		})
		if err != nil {
			return nil, fmt.Errorf("bundling failed for %s: %w", s.SourceURI, err)
		}

		// Extract vocabulary from custom metaschema if present
		vocab, err := extractVocabulary(ctx, bundled, cache)
		if err != nil {
			// Non-fatal: log and continue without vocabulary filtering
			ui.Verbosef("processor: could not extract vocabulary for %s: %v", s.SourceURI, err)
		}

		// Filter schema by vocabulary (strips disabled keywords)
		if vocab != nil {
			bundled, err = vocabulary.FilterSchema(bundled, vocab)
			if err != nil {
				return nil, fmt.Errorf("filtering by vocabulary failed for %s: %w", s.SourceURI, err)
			}
		}

		result[i] = ProcessedSchema{
			Namespace: s.Namespace,
			ID:        s.ID,
			Schema:    bundled,
			Adapter:   s.Adapter,
			SourceURI: s.SourceURI,
		}
	}

	return result, nil
}

// extractVocabulary extracts $vocabulary from a bundled schema.
// Checks embedded $vocabulary first, then fetches from custom metaschema if needed.
// Returns nil if schema uses standard draft without embedded vocabulary.
func extractVocabulary(ctx context.Context, schema json.RawMessage, cache *fetcher.SharedCache) (map[string]bool, error) {
	var parsed map[string]any
	if err := json.Unmarshal(schema, &parsed); err != nil {
		return nil, nil // not an object schema, no vocabulary
	}

	// Check if vocabulary already embedded in schema (by bundler or test)
	// This takes precedence over fetching from metaschema
	if vocab, ok := parsed["$vocabulary"].(map[string]any); ok {
		result := make(map[string]bool, len(vocab))
		for uri, val := range vocab {
			if required, ok := val.(bool); ok {
				result[uri] = required
			}
		}
		return result, nil
	}

	schemaURI, ok := parsed["$schema"].(string)
	if !ok || schemaURI == "" {
		return nil, nil // no $schema, use defaults
	}

	// Skip standard drafts - they have well-known vocabularies
	if metaschema.IsStandardDraft(schemaURI) {
		return nil, nil
	}

	// Fetch custom metaschema and extract vocabulary (using shared cache)
	meta, err := metaschema.GetWithCache(ctx, schemaURI, cache)
	if err != nil {
		return nil, err
	}

	return meta.Vocabulary, nil
}

// crawlAndFetch iteratively discovers external $refs in schemas and fetches them.
// It continues until no new URIs are found, populating the shared cache.
// Returns error immediately on any fetch failure (fail fast).
// The cache is checked before fetching to avoid duplicate requests for URIs
// already fetched by retriever or metaschema.
func crawlAndFetch(ctx context.Context, schemas []retriever.RetrievedSchema, f fetcher.Fetcher, cache *fetcher.SharedCache, verbose func(string)) error {
	visited := make(map[string]bool) // tracks URIs we've processed (including declared schemas)

	// Mark declared schema source URIs as visited (they're already fetched)
	for _, s := range schemas {
		if s.SourceURI != "" {
			normalized := fetcher.NormalizeURI(s.SourceURI)
			visited[normalized] = true
		}
	}

	// Initial frontier: refs from declared schemas
	frontier := make(map[string]string) // URI → base URI for resolution
	for _, s := range schemas {
		refs := extractExternalRefs(s.Schema, s.SourceURI)
		for _, ref := range refs {
			normalized := fetcher.NormalizeURI(ref)
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
			return ctx.Err()
		default:
		}

		verbose(fmt.Sprintf("processor: crawl iteration %d, %d new refs to fetch", iteration, len(frontier)))

		// Fetch all URIs in current frontier
		newFrontier := make(map[string]string)
		for uri, baseURI := range frontier {
			// Resolve relative URI against base
			resolved, err := resolveURI(uri, baseURI)
			if err != nil {
				return fmt.Errorf("processor: failed to resolve %q against %q: %w", uri, baseURI, err)
			}

			normalized := fetcher.NormalizeURI(resolved)
			if visited[normalized] {
				continue
			}
			visited[normalized] = true

			// Check shared cache first (may have been fetched by retriever or metaschema)
			if cached, ok := cache.Get(resolved); ok {
				ui.Verbosef("processor: cache hit for %s", resolved)
				// Still need to extract refs from cached schema
				newRefs := extractExternalRefs(cached, resolved)
				for _, ref := range newRefs {
					refNorm := fetcher.NormalizeURI(ref)
					if !visited[refNorm] {
						newFrontier[ref] = resolved
					}
				}
				continue
			}

			ui.Verbosef("processor: fetching %s", resolved)

			raw, err := f.Fetch(ctx, resolved)
			if err != nil {
				return fmt.Errorf("processor: failed to fetch %q: %w", resolved, err)
			}

			cache.Set(resolved, raw)

			// Extract refs from this schema and add to next frontier
			newRefs := extractExternalRefs(raw, resolved)
			for _, ref := range newRefs {
				refNorm := fetcher.NormalizeURI(ref)
				if !visited[refNorm] {
					newFrontier[ref] = resolved
				}
			}
		}

		frontier = newFrontier
	}

	verbose(fmt.Sprintf("processor: crawl finished in %d iterations", iteration))
	return nil
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

		// Check for $schema (custom metaschema needs fetching for validation)
		if schemaURI, ok := v["$schema"].(string); ok {
			if isExternalRef(schemaURI) && !metaschema.IsStandardDraft(schemaURI) {
				resolved, err := resolveURI(schemaURI, baseURI)
				if err == nil {
					if idx := strings.Index(resolved, "#"); idx >= 0 {
						resolved = resolved[:idx]
					}
					if resolved != "" {
						refs = append(refs, resolved)
					}
				}
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

