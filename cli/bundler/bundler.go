// Package bundler resolves external $ref in JSON Schemas, producing self-contained schemas.
package bundler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/xschemadev/xschema/fetcher"
	"github.com/xschemadev/xschema/ui"
	"github.com/xschemadev/xschema/unsupported"
)

// BundleInput contains all inputs needed to bundle a schema.
type BundleInput struct {
	Schema    json.RawMessage // the schema to bundle
	SourceURI string          // base URI for resolving relative refs (empty if schema has no source)
	Fetcher   fetcher.Fetcher // fetcher for external refs (nil means external refs will error)
	Draft     string          // JSON Schema draft for $schema injection if missing
}

// keyInvalidChars matches characters not allowed in $defs keys
var keyInvalidChars = regexp.MustCompile(`[^a-zA-Z0-9_]`)

// bundleContext holds state during bundling
type bundleContext struct {
	sourceURI  string            // base URI for resolving relative refs
	fetcher    fetcher.Fetcher   // fetcher for external schemas (nil = no external refs allowed)
	cache      map[string]any    // normalized URI → parsed schema (to avoid refetching)
	defs       map[string]any    // key → embedded schema (collected $defs)
	processing map[string]bool   // URIs currently being processed (circular ref detection)
	localIDs   map[string]string // normalized $id URI → JSON pointer path
	anchors    map[string]string // anchor name → JSON pointer path (e.g., "foo" → "/$defs/A")
	ctx        context.Context
	draft      string // JSON Schema draft for normalizing fetched schemas
}

// draftToSchemaURI maps draft names to their canonical $schema URIs
var draftToSchemaURI = map[string]string{
	"draft3":       "http://json-schema.org/draft-03/schema#",
	"draft4":       "http://json-schema.org/draft-04/schema#",
	"draft6":       "http://json-schema.org/draft-06/schema#",
	"draft7":       "http://json-schema.org/draft-07/schema#",
	"draft2019-09": "https://json-schema.org/draft/2019-09/schema",
	"draft2020-12": "https://json-schema.org/draft/2020-12/schema",
}

// Bundle resolves all external $refs in a schema, returning a self-contained schema.
// External refs are fetched via the Fetcher, bundled recursively, and embedded in $defs.
func Bundle(ctx context.Context, input BundleInput) (json.RawMessage, error) {
	var parsed any
	if err := json.Unmarshal(input.Schema, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse schema: %w", err)
	}

	// Inject $schema if missing and draft is specified
	if input.Draft != "" {
		if obj, ok := parsed.(map[string]any); ok {
			if _, hasSchema := obj["$schema"]; !hasSchema {
				if schemaURI, known := draftToSchemaURI[input.Draft]; known {
					obj["$schema"] = schemaURI
				}
			}
		}
	}

	// Normalize legacy syntax to draft 2020-12
	// Detect draft from schema or input, apply normalization if needed
	draft := input.Draft
	if draft == "" {
		if obj, ok := parsed.(map[string]any); ok {
			draft = detectDraftFromSchema(obj)
		}
	}
	if needsNormalization(draft) {
		parsed = normalizeLegacySyntax(parsed)
		// Update $schema to 2020-12 after normalization
		if obj, ok := parsed.(map[string]any); ok {
			obj["$schema"] = draftToSchemaURI["draft2020-12"]
		}
		ui.Verbosef("bundler: normalized %s syntax to draft2020-12", draft)
	}

	bctx := &bundleContext{
		sourceURI:  input.SourceURI,
		fetcher:    input.Fetcher,
		cache:      make(map[string]any),
		defs:       make(map[string]any),
		processing: make(map[string]bool),
		localIDs:   make(map[string]string),
		anchors:    make(map[string]string),
		ctx:        ctx,
		draft:      draft,
	}

	// First pass: collect all $id and $anchor declarations in the schema
	bctx.collectIDsAndAnchors(parsed, input.SourceURI, "", "")

	result, err := bctx.processNode(parsed, input.SourceURI, "")
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

	// Extract $vocabulary from custom metaschema if present
	if obj, ok := result.(map[string]any); ok {
		if err := bctx.extractMetaschemaVocabulary(obj); err != nil {
			return nil, err
		}
	}

	// Validate all internal refs point to valid JSON pointer targets
	if err := validateInternalRefs(result); err != nil {
		return nil, err
	}

	return json.Marshal(result)
}

// isStandardMetaschema checks if a $schema URI is a standard JSON Schema draft
func isStandardMetaschema(schemaURI string) bool {
	for _, uri := range draftToSchemaURI {
		if schemaURI == uri {
			return true
		}
	}
	// Also check common variations
	standardPrefixes := []string{
		"http://json-schema.org/",
		"https://json-schema.org/",
	}
	for _, prefix := range standardPrefixes {
		if strings.HasPrefix(schemaURI, prefix) {
			return true
		}
	}
	return false
}

// extractMetaschemaVocabulary fetches a custom metaschema and extracts its $vocabulary
// into the schema being bundled, so the TypeScript parser can respect vocabulary settings.
func (b *bundleContext) extractMetaschemaVocabulary(schema map[string]any) error {
	schemaURI, ok := schema["$schema"].(string)
	if !ok || schemaURI == "" {
		return nil
	}

	// Skip standard metaschemas - they have well-known vocabularies
	if isStandardMetaschema(schemaURI) {
		return nil
	}

	// Don't override if schema already has $vocabulary
	if _, hasVocab := schema["$vocabulary"]; hasVocab {
		return nil
	}

	ui.Verbosef("bundler: fetching custom metaschema: %s", schemaURI)

	// Fetch the metaschema
	raw, err := b.fetch(schemaURI)
	if err != nil {
		// Non-fatal: if we can't fetch the metaschema, continue without vocabulary
		ui.Verbosef("bundler: could not fetch metaschema %s: %v", schemaURI, err)
		return nil
	}

	var metaschema map[string]any
	if err := json.Unmarshal(raw, &metaschema); err != nil {
		ui.Verbosef("bundler: could not parse metaschema %s: %v", schemaURI, err)
		return nil
	}

	// Extract $vocabulary from metaschema
	vocab, ok := metaschema["$vocabulary"].(map[string]any)
	if !ok || len(vocab) == 0 {
		return nil
	}

	ui.Verbosef("bundler: extracted $vocabulary from metaschema: %v", vocab)
	schema["$vocabulary"] = vocab
	return nil
}

// collectIDsAndAnchors recursively collects all $id and $anchor declarations in the schema,
// resolving them against the current base URI and tracking their JSON pointer paths.
// pathPrefix is prepended to all collected anchor paths (used when embedding remote schemas).
func (b *bundleContext) collectIDsAndAnchors(node any, baseURI string, path string, pathPrefix string) {
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
				fullPath := pathPrefix + path
				b.anchors[anchor] = fullPath
				ui.Verbosef("bundler: found fragment $id %q as anchor at path %q", id, fullPath)
			} else {
				// Resolve relative $id against base URI
				resolved, err := fetcher.ResolveURI(id, baseURI)
				if err == nil && resolved != "" {
					// Store without fragment
					resolvedBase, _ := splitFragment(resolved)
					if resolvedBase != "" {
						fullPath := pathPrefix + path
						normalizedBase := fetcher.NormalizeURI(resolvedBase)
						b.localIDs[normalizedBase] = fullPath
						ui.Verbosef("bundler: found local $id: %s at path %s (resolved from %s)", normalizedBase, fullPath, id)
						// Update base URI for children
						baseURI = normalizedBase
					}
				}
			}
		}

		// Check for $anchor (draft2019-09+)
		if anchor, ok := v["$anchor"].(string); ok {
			fullPath := pathPrefix + path
			b.anchors[anchor] = fullPath
			ui.Verbosef("bundler: found $anchor %q at path %q", anchor, fullPath)
		}

		// Recurse into all values with updated base URI and paths,
		// but skip data-only keywords (enum, const, default, etc.)
		for key, val := range v {
			if nonSchemaKeywords[key] {
				continue
			}
			childPath := path + "/" + escapeJSONPointer(key)
			b.collectIDsAndAnchors(val, baseURI, childPath, pathPrefix)
		}
	case []any:
		for i, val := range v {
			childPath := fmt.Sprintf("%s/%d", path, i)
			b.collectIDsAndAnchors(val, baseURI, childPath, pathPrefix)
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
// scopePath is the JSON pointer path of the nearest ancestor with $id (for rewriting scoped refs)
func (b *bundleContext) processNode(node any, baseURI string, scopePath string) (any, error) {
	switch v := node.(type) {
	case map[string]any:
		return b.processObject(v, baseURI, scopePath)
	case []any:
		return b.processArray(v, baseURI, scopePath)
	default:
		// Primitives pass through unchanged
		return node, nil
	}
}

// isMetaschema checks if a URI is an official JSON Schema metaschema URL.
// IsMetaschema checks if a URI is an official JSON Schema metaschema URL.
// These should not be fetched/bundled - they reference complex recursive schemas
// that use $recursiveAnchor and other features adapters can't handle.
func IsMetaschema(uri string) bool {
	return strings.HasPrefix(uri, "http://json-schema.org/draft-") ||
		strings.HasPrefix(uri, "https://json-schema.org/draft/")
}

// processObject handles object nodes, looking for $ref
// scopePath is the JSON pointer path of the nearest ancestor with $id (for rewriting scoped refs)
func (b *bundleContext) processObject(obj map[string]any, baseURI string, scopePath string) (any, error) {
	// Reject forbidden keywords (dynamic/recursive refs) with proper UnsupportedKeywordError
	keywords := unsupported.Keywords()
	for kw, reason := range keywords {
		if _, exists := obj[kw]; exists {
			return nil, &unsupported.UnsupportedKeywordError{
				Keyword: kw,
				Reason:  reason,
				Path:    scopePath,
			}
		}
	}

	// Track current path for scope tracking
	currentScopePath := scopePath

	// In pre-2019-09 drafts, $ref consumes the entire object — sibling $id/$anchor
	// are ignored. When $ref and $id appear on the same node, the $id must NOT
	// change the base URI used to resolve the $ref.
	_, hasRef := obj["$ref"].(string)
	refIgnoresSiblings := hasRef && needsNormalization(b.draft)

	// Check for $id (draft6+) or id (draft4/draft3) and update baseURI for this scope
	id, ok := obj["$id"].(string)
	if !ok {
		id, ok = obj["id"].(string)
	}
	if ok && !refIgnoresSiblings {
		newBase, err := fetcher.ResolveURI(id, baseURI)
		if err == nil && newBase != "" {
			ui.Verbosef("bundler: $id changes base URI: %q → %q", baseURI, newBase)
			baseURI = newBase
			// This object establishes a new scope - clear the scope path
			// (refs inside this scope are relative to this object)
			currentScopePath = ""
		}
	}

	// Check for $ref - handle scoped refs
	if ref, ok := obj["$ref"].(string); ok {
		refResult, err := b.processRef(obj, ref, baseURI)
		if err != nil {
			return nil, err
		}

		// processRef returns a shallow copy with rewritten $ref but unprocessed siblings.
		// For draft2019-09+, $ref doesn't consume the entire object — sibling keys
		// (like $defs containing other schemas) must still be processed.
		if refObj, ok := refResult.(map[string]any); ok {
			keys := make([]string, 0, len(refObj))
			for k := range refObj {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			result := make(map[string]any, len(refObj))
			for _, k := range keys {
				if k == "$ref" || nonSchemaKeywords[k] {
					result[k] = refObj[k]
					continue
				}
				var childScopePath string
				if currentScopePath == "" {
					childScopePath = "/" + escapeJSONPointer(k)
				} else {
					childScopePath = currentScopePath + "/" + escapeJSONPointer(k)
				}
				processed, err := b.processNode(refObj[k], baseURI, childScopePath)
				if err != nil {
					return nil, err
				}
				result[k] = processed
			}
			return result, nil
		}
		return refResult, nil
	}

	// Recursively process all properties in sorted order for deterministic error messages
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := make(map[string]any, len(obj))
	for _, k := range keys {
		// Skip data-only keywords — their values are literal data, not subschemas.
		// Processing them would incorrectly resolve $ref strings in enum values etc.
		if nonSchemaKeywords[k] {
			result[k] = obj[k]
			continue
		}
		// Build the child path (for scope tracking)
		var childScopePath string
		if currentScopePath == "" {
			childScopePath = "/" + escapeJSONPointer(k)
		} else {
			childScopePath = currentScopePath + "/" + escapeJSONPointer(k)
		}
		processed, err := b.processNode(obj[k], baseURI, childScopePath)
		if err != nil {
			return nil, err
		}
		result[k] = processed
	}
	return result, nil
}

// processArray handles array nodes
func (b *bundleContext) processArray(arr []any, baseURI string, scopePath string) (any, error) {
	result := make([]any, len(arr))
	for i, v := range arr {
		childPath := fmt.Sprintf("%s/%d", scopePath, i)
		processed, err := b.processNode(v, baseURI, childPath)
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
	resolvedURI, err := fetcher.ResolveURI(refURI, baseURI)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve ref %q against base %q: %w", ref, baseURI, err)
	}

	// Check if this ref points to a local $id - rewrite to local JSON pointer
	normalizedResolved := fetcher.NormalizeURI(resolvedURI)
	if localPath, ok := b.localIDs[normalizedResolved]; ok {
		localRef := "#"
		if localPath != "" {
			localRef += localPath
		}

		if fragment != "" {
			if strings.HasPrefix(fragment, "/") {
				localRef += fragment
			} else if anchorPath, found := b.anchors[fragment]; found {
				localRef = "#" + anchorPath
			} else {
				localRef = "#" + fragment
			}
		}

		result := copyObject(obj)
		result["$ref"] = localRef
		ui.Verbosef("bundler: ref %q matches local $id, rewriting to %s", ref, localRef)
		return result, nil
	}

	// Check if this is a metaschema ref - don't try to fetch/bundle these.
	// They use $recursiveAnchor and other features that would fail.
	// Replace with empty schema {} ("accept anything") since no adapter
	// can generate a "validate this is a valid JSON Schema" check.
	if IsMetaschema(resolvedURI) {
		ui.Verbosef("bundler: ref %q is a metaschema, replacing with unconstrained schema", ref)
		return map[string]any{}, nil
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

		// Normalize legacy syntax in fetched schema
		if obj, ok := parsed.(map[string]any); ok {
			fetchedDraft := detectDraftFromSchema(obj)
			// If fetched schema has no $schema, use parent's draft
			if fetchedDraft == "" {
				fetchedDraft = b.draft
			}
			if needsNormalization(fetchedDraft) {
				parsed = normalizeLegacySyntax(parsed)
				if obj, ok := parsed.(map[string]any); ok {
					obj["$schema"] = draftToSchemaURI["draft2020-12"]
				}
				ui.Verbosef("bundler: normalized fetched schema %s from %s to draft2020-12", resolvedURI, fetchedDraft)
			}
		}

		// Flatten nested $defs/definitions and rewrite internal refs
		key := b.uriToKey(resolvedURI)

		// Collect $id and $anchor declarations from the fetched schema.
		// We collect without prefix for internal ref resolution during processNode.
		// External anchor refs (remote.json#anchor) are resolved in the else-if branch below.
		b.collectIDsAndAnchors(parsed, resolvedURI, "", "")

		// Recursively bundle the fetched schema (it may have its own external refs)
		bundled, err := b.processNode(parsed, resolvedURI, "")
		if err != nil {
			return nil, fmt.Errorf("failed to bundle schema from %q: %w", resolvedURI, err)
		}
		bundled = b.flattenDefs(bundled, key)

		// Cache and add to defs
		b.cache[resolvedURI] = bundled
		b.defs[key] = bundled
	}

	// Rewrite the ref to point to $defs
	key := b.uriToKey(resolvedURI)
	localRef := "#/$defs/" + key

	// Handle fragment: could be JSON pointer (/path/to/thing), anchor (anchorName), or definition (/$defs/X)
	if fragment != "" {
		if defName, ok := parseDefFragment(fragment); ok {
			// Fragment points to a definition - rewrite to flattened key
			localRef = "#/$defs/" + key + "__" + defName
			ui.Verbosef("bundler: rewriting def fragment %s → %s", fragment, localRef)
		} else if !strings.HasPrefix(fragment, "/") {
			// Fragment is an anchor reference (e.g., "myAnchor" from remote.json#myAnchor)
			// Look up the anchor - it was collected relative to the remote schema root.
			if anchorPath, ok := b.anchors[fragment]; ok {
				// The anchor path is relative to the remote schema. We need to:
				// 1. Apply the /$defs/key prefix since that's where the schema is embedded
				// 2. Check if the location got flattened (/$defs/key/$defs/X → /$defs/key__X)
				fullPath := "/$defs/" + key + anchorPath
				localRef = "#" + b.rewriteAnchorPath(fullPath, key)
				ui.Verbosef("bundler: resolved anchor %s (path %s) → %s", fragment, anchorPath, localRef)
			} else {
				// Anchor not found - append as-is (will fail validation later)
				localRef = "#" + fragment
				ui.Verbosef("bundler: anchor %s not found, using #%s", fragment, fragment)
			}
		} else {
			// JSON pointer fragment - append the path to the embedded schema
			localRef += fragment
		}
	}

	result := copyObject(obj)
	result["$ref"] = localRef
	return result, nil
}

// rewriteAnchorPath rewrites an anchor path that may point to a location that got flattened.
// Handles deeply nested paths like /$defs/key/$defs/L1/$defs/L2/$defs/Target → /$defs/key__L1__L2__Target
func (b *bundleContext) rewriteAnchorPath(anchorPath string, embeddedKey string) string {
	prefix := "/$defs/" + embeddedKey

	// If path doesn't start with the expected prefix, return as-is
	if !strings.HasPrefix(anchorPath, prefix) {
		return anchorPath
	}

	// Get the part after /$defs/key
	rest := anchorPath[len(prefix):]
	if rest == "" {
		return anchorPath
	}

	// Extract all nested def names from the path
	// rest could be: /$defs/L1/$defs/L2/$defs/Target or /$defs/Target/properties/foo
	var defNames []string
	remaining := rest

	for remaining != "" {
		// Check for $defs or definitions prefix
		var defPrefix string
		if strings.HasPrefix(remaining, "/$defs/") {
			defPrefix = "/$defs/"
		} else if strings.HasPrefix(remaining, "/definitions/") {
			defPrefix = "/definitions/"
		} else {
			// No more def prefixes - any remaining path is preserved as suffix
			break
		}

		// Extract the def name
		afterPrefix := remaining[len(defPrefix):]
		nextSlash := strings.Index(afterPrefix, "/")
		var defName string
		if nextSlash >= 0 {
			defName = afterPrefix[:nextSlash]
			remaining = afterPrefix[nextSlash:]
		} else {
			defName = afterPrefix
			remaining = ""
		}
		defNames = append(defNames, defName)
	}

	// If no def names extracted, return original
	if len(defNames) == 0 {
		return anchorPath
	}

	// Build the flattened path: key__L1__L2__Target
	flattenedKey := embeddedKey
	for _, name := range defNames {
		flattenedKey += "__" + name
	}

	// Append any remaining path (e.g., /properties/foo)
	return "/$defs/" + flattenedKey + remaining
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

// defEntry holds information about a definition to be flattened
type defEntry struct {
	newKey   string   // flattened key (e.g., "parent__A__B")
	schema   any      // the schema content (with nested $defs stripped)
	oldPaths []string // all old paths that should rewrite to this (e.g., "#/$defs/A", "#/$defs/A/$defs/B")
}

// flattenDefs extracts $defs/definitions from an embedded schema and adds them
// to bctx.defs with prefixed keys. Uses O(n) single-pass algorithm:
// 1. Single DFS to collect all defs and build complete rewrite map
// 2. Single pass to rewrite all refs
func (b *bundleContext) flattenDefs(node any, parentKey string) any {
	obj, ok := node.(map[string]any)
	if !ok {
		return node
	}

	// Check if there are any defs to flatten
	hasDefs := false
	if d, ok := obj["$defs"].(map[string]any); ok && len(d) > 0 {
		hasDefs = true
	}
	if d, ok := obj["definitions"].(map[string]any); ok && len(d) > 0 {
		hasDefs = true
	}
	if !hasDefs {
		return node
	}

	// Phase 1: Single DFS pass to collect all defs and build rewrite map
	var entries []defEntry
	refRewrites := make(map[string]string)
	b.collectDefsRecursive(obj, parentKey, &entries, refRewrites)

	// Phase 2: Add all collected defs to b.defs and rewrite their refs
	for _, entry := range entries {
		rewritten := b.rewriteRefs(entry.schema, refRewrites, parentKey)
		b.defs[entry.newKey] = rewritten
		ui.Verbosef("bundler: flattened → $defs/%s", entry.newKey)
	}

	// Strip $defs/definitions from the root object
	result := make(map[string]any, len(obj))
	for k, v := range obj {
		if k == "$defs" || k == "definitions" {
			continue
		}
		result[k] = v
	}

	// Rewrite refs in the root object
	return b.rewriteRefs(result, refRewrites, parentKey)
}

// collectDefsRecursive performs a single DFS to collect all nested definitions.
// It populates entries with all defs found and builds the complete refRewrites map.
// Returns a slice of newly added rewrite keys (for parent to build nested paths).
func (b *bundleContext) collectDefsRecursive(
	node any,
	parentKey string,
	entries *[]defEntry,
	refRewrites map[string]string,
) []string {
	obj, ok := node.(map[string]any)
	if !ok {
		return nil
	}

	var newRewrites []string

	// Process both $defs and definitions
	for _, defsKey := range []string{"$defs", "definitions"} {
		defs, ok := obj[defsKey].(map[string]any)
		if !ok || len(defs) == 0 {
			continue
		}

		for defName, defSchema := range defs {
			// Build the old path (relative to this node)
			oldRef := "#/" + defsKey + "/" + defName

			// Generate unique flattened key
			newKey := b.uniqueDefKey(parentKey + "__" + defName)
			newRef := "#/$defs/" + newKey

			// Add to rewrite map
			refRewrites[oldRef] = newRef
			newRewrites = append(newRewrites, oldRef)

			// Recursively collect nested defs (returns keys added by children)
			childRewrites := b.collectDefsRecursive(defSchema, newKey, entries, refRewrites)

			// Add nested path rewrites for refs from this level down to children
			// e.g., #/$defs/L1/$defs/L2 needs to map to #/$defs/key__L1__L2
			for _, childOld := range childRewrites {
				childPath := childOld[1:] // strip # to get /$defs/...
				fullOldRef := oldRef + childPath
				if childNew, ok := refRewrites[childOld]; ok {
					refRewrites[fullOldRef] = childNew
					newRewrites = append(newRewrites, fullOldRef)
				}
			}

			// Strip nested $defs/definitions from the schema before storing
			strippedSchema := b.stripDefs(defSchema)

			*entries = append(*entries, defEntry{
				newKey:   newKey,
				schema:   strippedSchema,
				oldPaths: []string{oldRef},
			})
		}
	}

	return newRewrites
}

// stripDefs removes $defs and definitions from an object (shallow, non-recursive)
func (b *bundleContext) stripDefs(node any) any {
	obj, ok := node.(map[string]any)
	if !ok {
		return node
	}

	// Check if stripping is needed
	_, hasDefs := obj["$defs"]
	_, hasDefinitions := obj["definitions"]
	if !hasDefs && !hasDefinitions {
		return node
	}

	result := make(map[string]any, len(obj))
	for k, v := range obj {
		if k == "$defs" || k == "definitions" {
			continue
		}
		result[k] = v
	}
	return result
}

// rewriteRefs rewrites all $ref values according to the rewrite map.
// Also handles root-relative refs that should point to the parent def.
func (b *bundleContext) rewriteRefs(node any, rewrites map[string]string, parentKey string) any {
	switch v := node.(type) {
	case map[string]any:
		result := make(map[string]any, len(v))
		for k, val := range v {
			// Skip data-only keywords — don't rewrite $ref inside literal data
			if nonSchemaKeywords[k] {
				result[k] = val
				continue
			}
			if k == "$ref" {
				if ref, ok := val.(string); ok {
					if newRef := b.lookupRewrite(ref, rewrites, parentKey); newRef != "" {
						result[k] = newRef
						continue
					}
				}
			}
			result[k] = b.rewriteRefs(val, rewrites, parentKey)
		}
		return result
	case []any:
		result := make([]any, len(v))
		for i, val := range v {
			result[i] = b.rewriteRefs(val, rewrites, parentKey)
		}
		return result
	default:
		return node
	}
}

// lookupRewrite finds the appropriate rewritten ref, handling exact matches and prefix matches
func (b *bundleContext) lookupRewrite(ref string, rewrites map[string]string, parentKey string) string {
	// Exact match
	if newRef, ok := rewrites[ref]; ok {
		return newRef
	}

	// Check for refs with sub-paths (e.g., #/$defs/X/properties/foo)
	// We need to find the longest matching prefix
	var bestMatch string
	var bestMatchLen int
	for oldPrefix, newPrefix := range rewrites {
		if strings.HasPrefix(ref, oldPrefix+"/") && len(oldPrefix) > bestMatchLen {
			bestMatch = newPrefix + ref[len(oldPrefix):]
			bestMatchLen = len(oldPrefix)
		}
	}
	if bestMatch != "" {
		return bestMatch
	}

	// Handle root ref (#)
	if ref == "#" {
		return "#/$defs/" + parentKey
	}

	// Handle other root-relative refs not in rewrites
	if strings.HasPrefix(ref, "#/") {
		// Don't rewrite if it's already pointing to $defs or definitions
		if !strings.HasPrefix(ref, "#/$defs/") && !strings.HasPrefix(ref, "#/definitions/") {
			return "#/$defs/" + parentKey + ref[1:]
		}
	}

	return ""
}

// uniqueDefKey returns a unique key for a definition, appending a counter if needed
func (b *bundleContext) uniqueDefKey(baseKey string) string {
	if _, exists := b.defs[baseKey]; !exists {
		return baseKey
	}
	// Key collision - find a unique suffix
	for i := 2; ; i++ {
		candidate := fmt.Sprintf("%s__%d", baseKey, i)
		if _, exists := b.defs[candidate]; !exists {
			return candidate
		}
	}
}

// fetch retrieves a schema from a URI using the configured Fetcher
func (b *bundleContext) fetch(uri string) (json.RawMessage, error) {
	if b.fetcher == nil {
		return nil, fmt.Errorf("no fetcher configured: cannot fetch %q", uri)
	}
	return b.fetcher.Fetch(b.ctx, uri)
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
	// scopes tracks all subschemas with $id - refs inside should be validated against their scope
	scopes := []any{root}
	return validateInternalRefsRecursive(root, root, scopes)
}

// validateInternalRefsRecursive validates refs, handling scope changes from $id.
// node is the current node being processed
// root is the schema root
// scopes is the stack of enclosing schemas with $id (for validating scoped refs)
func validateInternalRefsRecursive(node, root any, scopes []any) error {
	switch v := node.(type) {
	case map[string]any:
		// Check for $id that creates a new scope
		newScopes := scopes
		if _, ok := v["$id"].(string); ok {
			newScopes = append([]any{v}, scopes...)
		} else if _, ok := v["id"].(string); ok {
			newScopes = append([]any{v}, scopes...)
		}

		if ref, ok := v["$ref"].(string); ok && strings.HasPrefix(ref, "#") {
			// Try validating against each scope in order (nearest first)
			var lastErr error
			for _, scope := range newScopes {
				err := validateJSONPointer(ref, scope)
				if err == nil {
					lastErr = nil
					break
				}
				lastErr = err
			}
			if lastErr != nil {
				return lastErr
			}
		}
		for key, val := range v {
			if nonSchemaKeywords[key] {
				continue
			}
			if err := validateInternalRefsRecursive(val, root, newScopes); err != nil {
				return err
			}
		}
	case []any:
		for _, val := range v {
			if err := validateInternalRefsRecursive(val, root, scopes); err != nil {
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
