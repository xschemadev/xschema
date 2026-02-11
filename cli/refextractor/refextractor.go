// Package refextractor discovers external $ref URIs in JSON Schema documents.
// It walks the schema tree, tracks $id base URI changes, and returns all
// external refs that need to be fetched.
package refextractor

import (
	"encoding/json"
	"strings"

	"github.com/xschemadev/xschema/fetcher"
	"github.com/xschemadev/xschema/metaschema"
)

// ExtractExternalRefs finds all external $ref URIs in a schema.
// Returns only external refs (not fragment-only refs like #/$defs/Foo).
// Also extracts custom $schema URIs that need to be fetched.
func ExtractExternalRefs(data json.RawMessage, baseURI string) []string {
	var parsed any
	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil
	}
	return extractFromNode(parsed, baseURI)
}

// extractFromNode recursively extracts external refs from a parsed schema.
func extractFromNode(node any, baseURI string) []string {
	var refs []string

	switch v := node.(type) {
	case map[string]any:
		// Check for $id which changes base URI
		if id, ok := v["$id"].(string); ok {
			if resolved, err := fetcher.ResolveURI(id, baseURI); err == nil {
				baseURI = resolved
			}
		} else if id, ok := v["id"].(string); ok {
			if resolved, err := fetcher.ResolveURI(id, baseURI); err == nil {
				baseURI = resolved
			}
		}

		// Check for $schema (custom metaschema needs fetching for validation)
		if schemaURI, ok := v["$schema"].(string); ok {
			if IsExternal(schemaURI) && !metaschema.IsStandardDraft(schemaURI) {
				resolved, err := fetcher.ResolveURI(schemaURI, baseURI)
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
			if IsExternal(ref) {
				// Resolve relative ref against base URI
				resolved, err := fetcher.ResolveURI(ref, baseURI)
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
			refs = append(refs, extractFromNode(val, baseURI)...)
		}

	case []any:
		for _, val := range v {
			refs = append(refs, extractFromNode(val, baseURI)...)
		}
	}

	return refs
}

// IsExternal returns true if the ref points to an external schema (not fragment-only).
func IsExternal(ref string) bool {
	// Fragment-only refs are local
	if strings.HasPrefix(ref, "#") {
		return false
	}
	return true
}
