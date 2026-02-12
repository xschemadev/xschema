// Package processor handles the crawl-fetch-validate-bundle pipeline.
// It sits between Retriever and Generator, iteratively discovering and fetching
// external refs before bundling schemas into self-contained units.
package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"sync"

	"github.com/xschemadev/xschema/bundler"
	"github.com/xschemadev/xschema/fetcher"
	"github.com/xschemadev/xschema/metaschema"
	"github.com/xschemadev/xschema/refextractor"
	"github.com/xschemadev/xschema/retriever"
	"github.com/xschemadev/xschema/ui"
	"github.com/xschemadev/xschema/unsupported"
	"github.com/xschemadev/xschema/validator"
	"github.com/xschemadev/xschema/vocabulary"
	"golang.org/x/sync/errgroup"
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
	Fetcher     fetcher.Fetcher      // fetcher for external refs (required)
	OnVerbose   func(msg string)     // callback for verbose logging (optional)
	Cache       *fetcher.SharedCache // shared cache (optional, creates internal if nil)
	Concurrency int                  // max concurrent fetches (0 = default: min(NumCPU, 8))
	Draft       string               // JSON Schema draft hint (e.g., "draft7") for schemas without $schema
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

	// Default concurrency: min(NumCPU, 8)
	concurrency := opts.Concurrency
	if concurrency <= 0 {
		concurrency = min(runtime.NumCPU(), 8)
	}

	verbose(fmt.Sprintf("processor: starting with %d schemas", len(schemas)))

	// Phase 0: Normalize legacy draft schemas to draft 2020-12
	schemas = normalizeDeclared(schemas, opts.Draft, verbose)

	// Phase 1: Validate declared schemas early (fail fast before fetching)
	if err := validateDeclared(schemas, verbose); err != nil {
		return nil, err
	}

	verbose(fmt.Sprintf("processor: declared schemas validated (%d schemas)", len(schemas)))

	// Phase 2: Crawl and fetch all external refs (parallel)
	if err := crawlAndFetch(ctx, schemas, opts.Fetcher, cache, concurrency, verbose); err != nil {
		return nil, err
	}

	verbose(fmt.Sprintf("processor: crawl complete, cache has %d external schemas", cache.Len()))

	// Phase 2.5: Normalize external schemas if needed (before validation)
	normalizeExternal(cache, opts.Draft, verbose)

	// Phase 3: Validate external schemas
	if err := validateExternal(schemas, cache, verbose); err != nil {
		return nil, err
	}

	verbose("processor: validation complete")

	// Phase 4: Bundle all schemas using the cache
	result, err := bundleAll(ctx, schemas, cache, verbose, opts.Draft)
	if err != nil {
		return nil, err
	}

	verbose(fmt.Sprintf("processor: bundling complete, produced %d schemas", len(result)))

	return result, nil
}

// normalizeDeclared normalizes legacy draft schemas to draft 2020-12.
// Returns a new slice with normalized schemas (original schemas are not modified).
func normalizeDeclared(schemas []retriever.RetrievedSchema, draft string, verbose func(string)) []retriever.RetrievedSchema {
	if !bundler.NeedsNormalization(draft) {
		return schemas
	}

	result := make([]retriever.RetrievedSchema, len(schemas))
	for i, s := range schemas {
		normalized, err := bundler.NormalizeSchema(s.Schema, draft)
		if err != nil {
			// If normalization fails, keep original (validation will catch errors)
			result[i] = s
			continue
		}
		result[i] = retriever.RetrievedSchema{
			Namespace: s.Namespace,
			ID:        s.ID,
			Schema:    normalized,
			Adapter:   s.Adapter,
			SourceURI: s.SourceURI,
		}
	}
	verbose(fmt.Sprintf("processor: normalized %d schemas from %s to draft2020-12", len(schemas), draft))
	return result
}

// validateDeclared validates declared schemas early, before any external fetching.
// This enables fail-fast behavior: invalid schemas are caught before wasting time
// fetching external refs. Schemas with custom metaschemas ($schema pointing to
// non-standard URI) are skipped here - they'll be validated after crawl when the
// custom metaschema is available.
func validateDeclared(schemas []retriever.RetrievedSchema, verbose func(string)) error {
	for _, s := range schemas {
		// Skip schemas with custom metaschemas (they need the metaschema fetched first)
		if metaschema.HasCustomMetaschema(s.Schema) {
			continue
		}
		if err := validator.ValidateSchema(s.Schema); err != nil {
			return fmt.Errorf("validation failed for %s: %w", s.SourceURI, err)
		}
	}
	return nil
}

// normalizeExternal normalizes legacy draft schemas in the cache to draft 2020-12.
// This must happen before validation since the validator uses 2020-12 rules.
func normalizeExternal(cache *fetcher.SharedCache, draft string, verbose func(string)) {
	if !bundler.NeedsNormalization(draft) {
		return
	}

	cacheSnapshot := cache.ToCache()
	normalized := 0

	for uri, data := range cacheSnapshot {
		result, err := bundler.NormalizeSchema(data, draft)
		if err != nil {
			continue // Skip schemas that fail to normalize
		}
		cache.Set(uri, result)
		normalized++
	}

	if normalized > 0 {
		verbose(fmt.Sprintf("processor: normalized %d external schemas from %s to draft2020-12", normalized, draft))
	}
}

// validateExternal validates external schemas and declared schemas with custom metaschemas.
// Runs after crawlAndFetch when custom metaschemas are available in the cache.
// Custom metaschemas from cache are passed to the validator so schemas using
// custom $schema URIs can be validated.
func validateExternal(schemas []retriever.RetrievedSchema, cache *fetcher.SharedCache, verbose func(string)) error {
	// Get snapshot of cache for iteration
	cacheSnapshot := cache.ToCache()

	// Build metaschemas map from cache for custom $schema validation
	schemaData := make([]json.RawMessage, len(schemas))
	for i, s := range schemas {
		schemaData[i] = s.Schema
	}
	metaschemas := metaschema.BuildMetaschemasMap(schemaData, cacheSnapshot)

	opts := &validator.ValidateOptions{
		Metaschemas: metaschemas,
	}

	// Validate declared schemas with custom metaschemas (skipped in early validation)
	for _, s := range schemas {
		if metaschema.HasCustomMetaschema(s.Schema) {
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

// bundleAll bundles each declared schema using a cache fetcher.
// No I/O occurs during bundling - all refs resolve from the pre-populated cache.
func bundleAll(ctx context.Context, schemas []retriever.RetrievedSchema, cache *fetcher.SharedCache, verbose func(string), draft string) ([]ProcessedSchema, error) {
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
			Draft:     draft,
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

		// Check for unsupported keywords (skip for boolean schemas which have no keywords)
		var parsed any
		if err := json.Unmarshal(bundled, &parsed); err != nil {
			return nil, fmt.Errorf("failed to parse bundled schema for %s: %w", s.SourceURI, err)
		}
		if ukErr := unsupported.ValidateKeywords(parsed); ukErr != nil {
			return nil, fmt.Errorf("schema %s contains unsupported keyword: %w", s.SourceURI, ukErr)
		}

		bundled, err = resolveInternalRefs(bundled, draft)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve internal refs for %s: %w", s.SourceURI, err)
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
// Uses wave-based parallel fetching: collect frontier -> parallel fetch -> extract refs -> repeat.
func crawlAndFetch(ctx context.Context, schemas []retriever.RetrievedSchema, f fetcher.Fetcher, cache *fetcher.SharedCache, concurrency int, verbose func(string)) error {
	var visitedMu sync.Mutex
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
		refs := refextractor.ExtractExternalRefs(s.Schema, s.SourceURI)
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

		// Build list of URIs to fetch in this wave (resolve and dedupe)
		type fetchItem struct {
			resolved string
			baseURI  string
		}
		var toFetch []fetchItem

		for uri, baseURI := range frontier {
			// Resolve relative URI against base
			resolved, err := fetcher.ResolveURI(uri, baseURI)
			if err != nil {
				return fmt.Errorf("processor: failed to resolve %q against %q: %w", uri, baseURI, err)
			}

			normalized := fetcher.NormalizeURI(resolved)
			if visited[normalized] {
				continue
			}
			visited[normalized] = true

			toFetch = append(toFetch, fetchItem{resolved: resolved, baseURI: resolved})
		}

		if len(toFetch) == 0 {
			break
		}

		// Parallel fetch all URIs in this wave
		var newFrontierMu sync.Mutex
		newFrontier := make(map[string]string)

		g, gctx := errgroup.WithContext(ctx)
		g.SetLimit(concurrency)

		for _, item := range toFetch {
			item := item // capture for closure
			g.Go(func() error {
				var raw json.RawMessage

				// Check cache again (concurrent fetch may have populated it)
				if cached, ok := cache.Get(item.resolved); ok {
					ui.Verbosef("processor: cache hit for %s", item.resolved)
					raw = cached
				} else {
					ui.Verbosef("processor: fetching %s", item.resolved)
					var err error
					raw, err = f.Fetch(gctx, item.resolved)
					if err != nil {
						return fmt.Errorf("processor: failed to fetch %q: %w", item.resolved, err)
					}
					cache.Set(item.resolved, raw)
				}

				// Extract refs from this schema and add to next frontier
				newRefs := refextractor.ExtractExternalRefs(raw, item.resolved)
				newFrontierMu.Lock()
				for _, ref := range newRefs {
					refNorm := fetcher.NormalizeURI(ref)
					visitedMu.Lock()
					alreadyVisited := visited[refNorm]
					visitedMu.Unlock()
					if !alreadyVisited {
						newFrontier[ref] = item.resolved
					}
				}
				newFrontierMu.Unlock()
				return nil
			})
		}

		if err := g.Wait(); err != nil {
			return err
		}

		frontier = newFrontier
	}

	verbose(fmt.Sprintf("processor: crawl finished in %d iterations", iteration))
	return nil
}
