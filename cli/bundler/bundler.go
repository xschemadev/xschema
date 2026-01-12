// Package bundler resolves external $ref in JSON Schemas, producing self-contained schemas.
package bundler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"sort"
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
	Draft       string        // JSON Schema draft (e.g., "draft4", "draft7") for $schema injection
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
	cache      map[string]any    // normalized URI → parsed schema (to avoid refetching)
	defs       map[string]any    // key → embedded schema (collected $defs)
	processing map[string]bool   // URIs currently being processed (circular ref detection)
	localIDs   map[string]bool   // $id values declared in the schema (to skip local refs)
	anchors    map[string]string // anchor name → JSON pointer path (e.g., "foo" → "/$defs/A")
	ctx        context.Context
}

// draftToSchemaURI maps draft names to their canonical $schema URIs
var draftToSchemaURI = map[string]string{
	"draft3":      "http://json-schema.org/draft-03/schema#",
	"draft4":      "http://json-schema.org/draft-04/schema#",
	"draft6":      "http://json-schema.org/draft-06/schema#",
	"draft7":      "http://json-schema.org/draft-07/schema#",
	"draft2019-09": "https://json-schema.org/draft/2019-09/schema",
	"draft2020-12": "https://json-schema.org/draft/2020-12/schema",
}

// Bundle resolves all external $refs in a schema, returning a self-contained schema.
// External refs are fetched, bundled recursively, and embedded in $defs.
func Bundle(ctx context.Context, schema json.RawMessage, opts Options) (json.RawMessage, error) {
	var parsed any
	if err := json.Unmarshal(schema, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse schema: %w", err)
	}

	// Inject $schema if missing and draft is specified
	if opts.Draft != "" {
		if obj, ok := parsed.(map[string]any); ok {
			if _, hasSchema := obj["$schema"]; !hasSchema {
				if schemaURI, known := draftToSchemaURI[opts.Draft]; known {
					obj["$schema"] = schemaURI
				}
			}
		}
	}

	bctx := &bundleContext{
		opts:       opts,
		cache:      make(map[string]any),
		defs:       make(map[string]any),
		processing: make(map[string]bool),
		localIDs:   make(map[string]bool),
		anchors:    make(map[string]string),
		ctx:        ctx,
	}

	// First pass: collect all $id and $anchor declarations in the schema
	bctx.collectIDsAndAnchors(parsed, opts.BaseURI, "")

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

	// Validate all internal refs point to valid JSON pointer targets
	if err := validateInternalRefs(result); err != nil {
		return nil, err
	}

	return json.Marshal(result)
}

// collectIDsAndAnchors recursively collects all $id and $anchor declarations in the schema,
// resolving them against the current base URI and tracking their JSON pointer paths
func (b *bundleContext) collectIDsAndAnchors(node any, baseURI string, path string) {
	switch v := node.(type) {
	case map[string]any:
		// Check for $id (draft6+) or id (draft4/draft3)
		id, ok := v["$id"].(string)
		if !ok {
			id, ok = v["id"].(string)
		}
		if ok {
			// Check if $id is fragment-only (e.g., "#foo") - this is a location-independent identifier
			// In draft4/6/7, $id: "#foo" serves the same purpose as $anchor: "foo" in draft2019-09+
			if strings.HasPrefix(id, "#") && len(id) > 1 {
				anchor := id[1:] // strip the #
				b.anchors[anchor] = path
				ui.Verbosef("bundler: found fragment $id %q as anchor at path %q", id, path)
			} else {
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
		}

		// Check for $anchor (draft2019-09+)
		if anchor, ok := v["$anchor"].(string); ok {
			b.anchors[anchor] = path
			ui.Verbosef("bundler: found $anchor %q at path %q", anchor, path)
		}

		// Recurse into all values with updated base URI and paths
		for key, val := range v {
			childPath := path + "/" + escapeJSONPointer(key)
			b.collectIDsAndAnchors(val, baseURI, childPath)
		}
	case []any:
		for i, val := range v {
			childPath := fmt.Sprintf("%s/%d", path, i)
			b.collectIDsAndAnchors(val, baseURI, childPath)
		}
	}
}

// escapeJSONPointer escapes a key for use in a JSON pointer (RFC 6901)
func escapeJSONPointer(key string) string {
	// ~ must be escaped first (as ~0), then / (as ~1)
	key = strings.ReplaceAll(key, "~", "~0")
	key = strings.ReplaceAll(key, "/", "~1")
	return key
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

// forbiddenKeywords are keywords that require evaluation semantics we can't provide
var forbiddenKeywords = []string{"$dynamicRef", "$dynamicAnchor", "$recursiveRef", "$recursiveAnchor"}

// isMetaschema checks if a URI is an official JSON Schema metaschema URL.
// These should not be fetched/bundled - they reference complex recursive schemas
// that use $recursiveAnchor and other features adapters can't handle.
func isMetaschema(uri string) bool {
	return strings.HasPrefix(uri, "http://json-schema.org/draft-") ||
		strings.HasPrefix(uri, "https://json-schema.org/draft/")
}

// processObject handles object nodes, looking for $ref
func (b *bundleContext) processObject(obj map[string]any, baseURI string) (any, error) {
	// Reject forbidden keywords (dynamic/recursive refs)
	for _, kw := range forbiddenKeywords {
		if _, exists := obj[kw]; exists {
			return nil, fmt.Errorf("unsupported keyword %q: dynamic and recursive references are not supported", kw)
		}
	}

	// Check for $id (draft6+) or id (draft4/draft3) and update baseURI for this scope
	id, ok := obj["$id"].(string)
	if !ok {
		id, ok = obj["id"].(string)
	}
	if ok {
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

	// Recursively process all properties in sorted order for deterministic error messages
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := make(map[string]any, len(obj))
	for _, k := range keys {
		processed, err := b.processNode(obj[k], baseURI)
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

// processRef handles a $ref, resolving external refs and anchor refs
func (b *bundleContext) processRef(obj map[string]any, ref string, baseURI string) (any, error) {
	// Local ref - check if it's an anchor ref that needs rewriting
	if strings.HasPrefix(ref, "#") {
		fragment := ref[1:] // strip the #

		// If fragment is empty or starts with /, it's a JSON pointer - pass through
		if fragment == "" || strings.HasPrefix(fragment, "/") {
			return obj, nil
		}

		// Fragment is an anchor reference (e.g., "#foo")
		// Look up the anchor and rewrite to its JSON pointer path
		if path, ok := b.anchors[fragment]; ok {
			result := copyObject(obj)
			result["$ref"] = "#" + path
			ui.Verbosef("bundler: rewriting anchor ref #%s → #%s", fragment, path)
			return result, nil
		}

		// Anchor not found - will fail validation later
		return obj, nil
	}

	// Parse the ref URI
	refURI, fragment := splitFragment(ref)

	// Check if this is an external ref with no usable base URI
	// A relative file ref without a base URI would resolve against cwd, which is not allowed
	refURL, _ := url.Parse(refURI)
	isRelative := refURL == nil || !refURL.IsAbs()
	if isRelative && baseURI == "" {
		return nil, fmt.Errorf("external $ref %q requires a base URI: relative file refs cannot resolve against cwd", ref)
	}

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

	// Check if this is a metaschema ref - don't try to fetch/bundle these
	// They use $recursiveAnchor and other features that would fail
	if isMetaschema(resolvedURI) {
		ui.Verbosef("bundler: ref %q is a metaschema, skipping bundling", ref)
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

		// Collect $id and $anchor declarations from the fetched schema
		b.collectIDsAndAnchors(parsed, resolvedURI, "")

		// Recursively bundle the fetched schema (it may have its own external refs)
		bundled, err := b.processNode(parsed, resolvedURI)
		if err != nil {
			return nil, fmt.Errorf("failed to bundle schema from %q: %w", resolvedURI, err)
		}

		// Flatten nested $defs/definitions and rewrite internal refs
		key := b.uriToKey(resolvedURI)
		bundled = b.flattenDefs(bundled, key)

		// Cache and add to defs
		b.cache[resolvedURI] = bundled
		b.defs[key] = bundled
	}

	// Rewrite the ref to point to $defs
	key := b.uriToKey(resolvedURI)
	localRef := "#/$defs/" + key

	// If fragment points to $defs/X or definitions/X, rewrite to flattened location
	if fragment != "" {
		if defName, ok := parseDefFragment(fragment); ok {
			// Fragment points to a definition - rewrite to flattened key
			localRef = "#/$defs/" + key + "__" + defName
			ui.Verbosef("bundler: rewriting def fragment %s → %s", fragment, localRef)
		} else {
			// Not a definition fragment - append the path to the embedded schema
			localRef += fragment
		}
	}

	result := copyObject(obj)
	result["$ref"] = localRef
	return result, nil
}

// parseDefFragment checks if a fragment points to a definition and returns the def name.
// It handles patterns like /$defs/Name, /definitions/Name, etc.
func parseDefFragment(fragment string) (string, bool) {
	// Try /$defs/Name pattern
	if strings.HasPrefix(fragment, "/$defs/") {
		rest := fragment[len("/$defs/"):]
		// Extract just the definition name (stop at next / if any)
		if idx := strings.Index(rest, "/"); idx >= 0 {
			return rest[:idx], true
		}
		return rest, true
	}
	// Try /definitions/Name pattern
	if strings.HasPrefix(fragment, "/definitions/") {
		rest := fragment[len("/definitions/"):]
		if idx := strings.Index(rest, "/"); idx >= 0 {
			return rest[:idx], true
		}
		return rest, true
	}
	return "", false
}

// flattenDefs extracts $defs/definitions from an embedded schema and adds them
// to bctx.defs with prefixed keys. It rewrites refs in both the embedded schema
// and extracted defs to point to the flattened locations.
func (b *bundleContext) flattenDefs(node any, parentKey string) any {
	obj, ok := node.(map[string]any)
	if !ok {
		return node
	}

	// Look for $defs or definitions in the embedded schema
	var defsKey string
	var defs map[string]any
	if d, ok := obj["$defs"].(map[string]any); ok {
		defsKey = "$defs"
		defs = d
	} else if d, ok := obj["definitions"].(map[string]any); ok {
		defsKey = "definitions"
		defs = d
	}

	if defs == nil {
		// No nested defs, still need to rewrite refs for the parent key
		return b.rewriteInternalRefs(node, parentKey)
	}

	// Build the ref rewrite map: #/$defs/X or #/definitions/X → #/$defs/parentKey__X
	refRewrites := make(map[string]string)
	for defName := range defs {
		oldRef := "#/" + defsKey + "/" + defName
		newKey := parentKey + "__" + defName
		newRef := "#/$defs/" + newKey
		refRewrites[oldRef] = newRef
	}

	// Extract and flatten each definition
	for defName, defSchema := range defs {
		newKey := parentKey + "__" + defName
		// Recursively flatten (in case nested defs have their own $defs)
		flattened := b.flattenDefs(defSchema, newKey)
		// Rewrite refs in the flattened schema using our rewrite map
		rewritten := b.rewriteNestedRefs(flattened, refRewrites, parentKey)
		b.defs[newKey] = rewritten
		ui.Verbosef("bundler: flattened %s/%s → $defs/%s", defsKey, defName, newKey)
	}

	// Remove $defs/definitions from the embedded schema
	result := make(map[string]any, len(obj)-1)
	for k, v := range obj {
		if k == "$defs" || k == "definitions" {
			continue
		}
		result[k] = v
	}

	// Rewrite refs in the embedded schema (now without $defs/definitions)
	// First apply nested ref rewrites, then the parent key rewrites
	rewritten := b.rewriteNestedRefs(result, refRewrites, parentKey)
	return rewritten
}

// rewriteNestedRefs rewrites refs according to a rewrite map, and also handles
// root-relative refs that should point to the parent def
func (b *bundleContext) rewriteNestedRefs(node any, rewrites map[string]string, parentKey string) any {
	switch v := node.(type) {
	case map[string]any:
		result := make(map[string]any, len(v))
		for k, val := range v {
			if k == "$ref" {
				if ref, ok := val.(string); ok {
					// Check for exact matches in rewrite map
					if newRef, ok := rewrites[ref]; ok {
						result[k] = newRef
						continue
					}
					// Check for refs with sub-paths (e.g., #/$defs/X/properties/foo)
					for oldPrefix, newPrefix := range rewrites {
						if strings.HasPrefix(ref, oldPrefix+"/") {
							suffix := ref[len(oldPrefix):]
							result[k] = newPrefix + suffix
							continue
						}
					}
					// Handle root refs
					if ref == "#" {
						result[k] = "#/$defs/" + parentKey
						continue
					}
					// Handle other root-relative refs (not in rewrites)
					if strings.HasPrefix(ref, "#/") {
						// Check if this is a ref to $defs or definitions that we already handled
						if !strings.HasPrefix(ref, "#/$defs/") && !strings.HasPrefix(ref, "#/definitions/") {
							result[k] = "#/$defs/" + parentKey + ref[1:]
							continue
						}
					}
				}
			}
			result[k] = b.rewriteNestedRefs(val, rewrites, parentKey)
		}
		return result
	case []any:
		result := make([]any, len(v))
		for i, val := range v {
			result[i] = b.rewriteNestedRefs(val, rewrites, parentKey)
		}
		return result
	default:
		return node
	}
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
		return retriever.RetrieveFromFilePath(b.ctx, localPath, b.opts.Draft)
	}

	// Check if it's a file path or URL
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		opts := retriever.Options{
			HTTPTimeout: b.opts.HTTPTimeout,
			Retries:     b.opts.Retries,
		}
		return retriever.RetrieveFromURL(b.ctx, uri, opts, b.opts.Draft)
	}

	// Assume it's a file path
	return retriever.RetrieveFromFilePath(b.ctx, uri, b.opts.Draft)
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

// validateInternalRefs validates that all internal $refs (starting with #) point to valid targets
func validateInternalRefs(root any) error {
	return validateInternalRefsRecursive(root, root)
}

func validateInternalRefsRecursive(node, root any) error {
	switch v := node.(type) {
	case map[string]any:
		if ref, ok := v["$ref"].(string); ok && strings.HasPrefix(ref, "#") {
			if err := validateJSONPointer(ref, root); err != nil {
				return err
			}
		}
		for _, val := range v {
			if err := validateInternalRefsRecursive(val, root); err != nil {
				return err
			}
		}
	case []any:
		for _, val := range v {
			if err := validateInternalRefsRecursive(val, root); err != nil {
				return err
			}
		}
	}
	return nil
}

// validateJSONPointer checks that a JSON pointer ref resolves to a valid target in the schema
func validateJSONPointer(ref string, root any) error {
	// Strip the leading #
	pointer := strings.TrimPrefix(ref, "#")

	// Empty pointer (#) refers to root, always valid
	if pointer == "" {
		return nil
	}

	// Must start with /
	if !strings.HasPrefix(pointer, "/") {
		return fmt.Errorf("invalid JSON pointer in $ref %q: must start with /", ref)
	}

	// Split into path segments
	segments := strings.Split(pointer[1:], "/")
	current := root

	for _, segment := range segments {
		// URI-decode first (RFC 6901 §6: URI fragment encoding)
		decoded, err := url.PathUnescape(segment)
		if err != nil {
			return fmt.Errorf("invalid URI encoding in $ref %q: %w", ref, err)
		}
		segment = decoded

		// Then unescape JSON pointer escapes: ~1 → /, ~0 → ~
		segment = strings.ReplaceAll(segment, "~1", "/")
		segment = strings.ReplaceAll(segment, "~0", "~")

		switch v := current.(type) {
		case map[string]any:
			next, exists := v[segment]
			if !exists {
				return fmt.Errorf("$ref %q points to missing target: key %q not found", ref, segment)
			}
			current = next
		case []any:
			// Parse array index
			var idx int
			if _, err := fmt.Sscanf(segment, "%d", &idx); err != nil {
				return fmt.Errorf("$ref %q points to missing target: %q is not a valid array index", ref, segment)
			}
			if idx < 0 || idx >= len(v) {
				return fmt.Errorf("$ref %q points to missing target: array index %d out of bounds", ref, idx)
			}
			current = v[idx]
		default:
			return fmt.Errorf("$ref %q points to missing target: cannot traverse into %T", ref, current)
		}
	}

	return nil
}
