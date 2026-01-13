package fetcher

import (
	"net/url"
	"path/filepath"
)

// ResolveURI resolves a potentially relative URI against a base URI.
// Handles both URL-based refs (http://...) and file path refs.
func ResolveURI(ref, base string) (string, error) {
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
