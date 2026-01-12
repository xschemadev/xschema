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

	"github.com/xschemadev/xschema/retriever"
	"github.com/xschemadev/xschema/ui"
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

	// Phase 2: validateAll (placeholder - implemented in US-011)
	// Phase 3: bundleAll (placeholder - implemented in US-012)

	// For now, return schemas as ProcessedSchema without bundling
	// (bundling happens in US-012)
	result := make([]ProcessedSchema, len(schemas))
	for i, s := range schemas {
		result[i] = ProcessedSchema{
			Namespace:  s.Namespace,
			ID:         s.ID,
			Schema:     s.Schema,
			Adapter:    s.Adapter,
			SourceURI:  s.SourceURI,
			Vocabulary: nil, // extracted in bundleAll phase
		}
	}

	return result, nil
}

// externalRefCache maps normalized URI → raw schema bytes
type externalRefCache map[string]json.RawMessage

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
