// Package bundler resolves external $ref in JSON Schemas, producing self-contained schemas.
package bundler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/xschemadev/xschema/retriever"
	"github.com/xschemadev/xschema/ui"
)

// keyInvalidChars matches characters not allowed in $defs keys
var keyInvalidChars = regexp.MustCompile(`[^a-zA-Z0-9_]`)

// Options configures bundling behavior
type Options struct {
	BaseURI     string        // base URI for resolving relative refs
	RemotesPath string        // path to remotes/ folder for localhost:1234 mapping
	HTTPTimeout time.Duration // timeout for HTTP requests
	Retries     int           // number of retries for HTTP requests
}

// DefaultOptions returns sensible defaults
func DefaultOptions() Options {
	return Options{
		HTTPTimeout: 30 * time.Second,
		Retries:     3,
	}
}

// bundleContext holds state during bundling
type bundleContext struct {
	opts       Options
	cache      map[string]any  // normalized URI → parsed schema (to avoid refetching)
	defs       map[string]any  // key → embedded schema (collected $defs)
	processing map[string]bool // URIs currently being processed (circular ref detection)
	localIDs   map[string]bool // $id values declared in the schema (to skip local refs)
	ctx        context.Context
}

// Bundle resolves all external $refs in a schema, returning a self-contained schema.
// External refs are fetched, bundled recursively, and embedded in $defs.
func Bundle(ctx context.Context, schema json.RawMessage, opts Options) (json.RawMessage, error) {
	var parsed any
	if err := json.Unmarshal(schema, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse schema: %w", err)
	}

	bctx := &bundleContext{
		opts:       opts,
		cache:      make(map[string]any),
		defs:       make(map[string]any),
		processing: make(map[string]bool),
		localIDs:   make(map[string]bool),
		ctx:        ctx,
	}

	// First pass: collect all $id declarations in the schema (with resolved URIs)
	bctx.collectIDs(parsed, opts.BaseURI)

	result, err := bctx.processNode(parsed, opts.BaseURI)
	if err != nil {
		return nil, err
	}

	// If we collected any external defs, merge them into the schema
	if len(bctx.defs) > 0 {
		if obj, ok := result.(map[string]any); ok {
			// Merge into existing $defs or create new
			existingDefs, _ := obj["$defs"].(map[string]any)
			if existingDefs == nil {
				existingDefs = make(map[string]any)
			}
			for k, v := range bctx.defs {
				existingDefs[k] = v
			}
			obj["$defs"] = existingDefs
			result = obj
		}
	}

	return json.Marshal(result)
}

// collectIDs recursively collects all $id declarations in the schema,
// resolving them against the current base URI
func (b *bundleContext) collectIDs(node any, baseURI string) {
	switch v := node.(type) {
	case map[string]any:
		// Check for $id and resolve it against current base
		if id, ok := v["$id"].(string); ok {
			// Resolve relative $id against base URI
			resolved, err := resolveURI(id, baseURI)
			if err == nil && resolved != "" {
				// Store without fragment
				resolvedBase, _ := splitFragment(resolved)
				if resolvedBase != "" {
					b.localIDs[resolvedBase] = true
					ui.Verbosef("bundler: found local $id: %s (resolved from %s)", resolvedBase, id)
					// Update base URI for children
					baseURI = resolvedBase
				}
			}
		}
		// Recurse into all values with updated base URI
		for _, val := range v {
			b.collectIDs(val, baseURI)
		}
	case []any:
		for _, val := range v {
			b.collectIDs(val, baseURI)
		}
	}
}

// processNode recursively processes a schema node, resolving external refs
func (b *bundleContext) processNode(node any, baseURI string) (any, error) {
	switch v := node.(type) {
	case map[string]any:
		return b.processObject(v, baseURI)
	case []any:
		return b.processArray(v, baseURI)
	default:
		// Primitives pass through unchanged
		return node, nil
	}
}

// processObject handles object nodes, looking for $ref
func (b *bundleContext) processObject(obj map[string]any, baseURI string) (any, error) {
	// Check for $id and update baseURI for this scope
	if id, ok := obj["$id"].(string); ok {
		newBase, err := resolveURI(id, baseURI)
		if err == nil && newBase != "" {
			ui.Verbosef("bundler: $id changes base URI: %q → %q", baseURI, newBase)
			baseURI = newBase
		}
	}

	// Check for $ref
	if ref, ok := obj["$ref"].(string); ok {
		return b.processRef(obj, ref, baseURI)
	}

	// Recursively process all properties
	result := make(map[string]any, len(obj))
	for k, v := range obj {
		processed, err := b.processNode(v, baseURI)
		if err != nil {
			return nil, err
		}
		result[k] = processed
	}
	return result, nil
}

// processArray handles array nodes
func (b *bundleContext) processArray(arr []any, baseURI string) (any, error) {
	result := make([]any, len(arr))
	for i, v := range arr {
		processed, err := b.processNode(v, baseURI)
		if err != nil {
			return nil, err
		}
		result[i] = processed
	}
	return result, nil
}

// processRef handles a $ref, resolving external refs
func (b *bundleContext) processRef(obj map[string]any, ref string, baseURI string) (any, error) {
	// Local ref - nothing to do
	if strings.HasPrefix(ref, "#") {
		return obj, nil
	}

	// Parse the ref URI
	refURI, fragment := splitFragment(ref)

	// Resolve relative URI against base
	resolvedURI, err := resolveURI(refURI, baseURI)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve ref %q against base %q: %w", ref, baseURI, err)
	}

	// Check if this ref points to a local $id - if so, skip bundling
	if b.localIDs[resolvedURI] {
		ui.Verbosef("bundler: ref %q matches local $id, skipping", ref)
		return obj, nil
	}

	ui.Verbosef("bundler: resolving external ref %q → %q", ref, resolvedURI)

	// Check for circular reference
	if b.processing[resolvedURI] {
		ui.Verbosef("bundler: circular ref detected for %q, using lazy reference", resolvedURI)
		// Return the ref as-is with a local pointer to where it will be
		key := b.uriToKey(resolvedURI)
		localRef := "#/$defs/" + key
		if fragment != "" {
			localRef += fragment
		}
		result := copyObject(obj)
		result["$ref"] = localRef
		return result, nil
	}

	// Check cache
	if _, ok := b.cache[resolvedURI]; !ok {
		// Mark as processing (for circular ref detection)
		b.processing[resolvedURI] = true
		defer delete(b.processing, resolvedURI)

		// Fetch the external schema
		fetched, err := b.fetch(resolvedURI)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch %q: %w", resolvedURI, err)
		}

		var parsed any
		if err := json.Unmarshal(fetched, &parsed); err != nil {
			return nil, fmt.Errorf("failed to parse schema from %q: %w", resolvedURI, err)
		}

		// Collect $id declarations from the fetched schema (so we don't try to fetch local refs)
		b.collectIDs(parsed, resolvedURI)

		// Recursively bundle the fetched schema (it may have its own external refs)
		bundled, err := b.processNode(parsed, resolvedURI)
		if err != nil {
			return nil, fmt.Errorf("failed to bundle schema from %q: %w", resolvedURI, err)
		}

		// Rewrite internal refs in the bundled schema
		key := b.uriToKey(resolvedURI)
		bundled = b.rewriteInternalRefs(bundled, key)

		// Cache and add to defs
		b.cache[resolvedURI] = bundled
		b.defs[key] = bundled
	}

	// Rewrite the ref to point to $defs
	key := b.uriToKey(resolvedURI)
	localRef := "#/$defs/" + key
	if fragment != "" {
		localRef += fragment
	}

	result := copyObject(obj)
	result["$ref"] = localRef
	return result, nil
}

// rewriteInternalRefs rewrites root-relative refs (e.g., #/definitions/Foo)
// to point into the embedded location (e.g., #/$defs/key/definitions/Foo)
func (b *bundleContext) rewriteInternalRefs(node any, defsKey string) any {
	switch v := node.(type) {
	case map[string]any:
		result := make(map[string]any, len(v))
		for k, val := range v {
			if k == "$ref" {
				if ref, ok := val.(string); ok {
					if ref == "#" {
						// Bare root ref: # → #/$defs/key
						result[k] = "#/$defs/" + defsKey
						continue
					}
					if strings.HasPrefix(ref, "#/") {
						// Path ref: #/foo/bar → #/$defs/key/foo/bar
						result[k] = "#/$defs/" + defsKey + ref[1:]
						continue
					}
				}
			}
			result[k] = b.rewriteInternalRefs(val, defsKey)
		}
		return result
	case []any:
		result := make([]any, len(v))
		for i, val := range v {
			result[i] = b.rewriteInternalRefs(val, defsKey)
		}
		return result
	default:
		return node
	}
}

// fetch retrieves a schema from a URI, handling localhost mapping
func (b *bundleContext) fetch(uri string) (json.RawMessage, error) {
	// Handle localhost:1234 mapping for JSON Schema Test Suite
	if b.opts.RemotesPath != "" && strings.HasPrefix(uri, "http://localhost:1234/") {
		localPath := filepath.Join(b.opts.RemotesPath, strings.TrimPrefix(uri, "http://localhost:1234/"))
		ui.Verbosef("bundler: mapping localhost URL to local file: %s → %s", uri, localPath)
		return retriever.RetrieveFromFilePath(b.ctx, localPath)
	}

	// Check if it's a file path or URL
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		opts := retriever.Options{
			HTTPTimeout: b.opts.HTTPTimeout,
			Retries:     b.opts.Retries,
		}
		return retriever.RetrieveFromURL(b.ctx, uri, opts)
	}

	// Assume it's a file path
	return retriever.RetrieveFromFilePath(b.ctx, uri)
}

// uriToKey converts a URI to a valid $defs key
func (b *bundleContext) uriToKey(uri string) string {
	// Remove scheme
	key := uri
	key = strings.TrimPrefix(key, "http://")
	key = strings.TrimPrefix(key, "https://")
	key = strings.TrimPrefix(key, "file://")

	// Replace invalid characters with underscores
	key = keyInvalidChars.ReplaceAllString(key, "_")

	// Ensure it starts with a letter or underscore
	if len(key) > 0 && key[0] >= '0' && key[0] <= '9' {
		key = "_" + key
	}

	// Truncate if too long
	if len(key) > 100 {
		key = key[:100]
	}

	return key
}

// splitFragment splits a URI into base and fragment parts
// e.g., "foo.json#/definitions/Bar" → ("foo.json", "/definitions/Bar")
func splitFragment(uri string) (base, fragment string) {
	if idx := strings.Index(uri, "#"); idx >= 0 {
		return uri[:idx], uri[idx+1:]
	}
	return uri, ""
}

// resolveURI resolves a potentially relative URI against a base URI
func resolveURI(ref, base string) (string, error) {
	if base == "" {
		// No base, ref must be absolute or we treat it as-is
		return ref, nil
	}

	// Parse base
	baseURL, err := url.Parse(base)
	if err != nil {
		// Base might be a file path
		if filepath.IsAbs(ref) {
			return ref, nil
		}
		// Resolve as file path
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

// copyObject creates a shallow copy of a map
func copyObject(obj map[string]any) map[string]any {
	result := make(map[string]any, len(obj))
	for k, v := range obj {
		result[k] = v
	}
	return result
}
