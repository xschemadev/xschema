package language

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var windowsDrivePrefix = regexp.MustCompile(`^[A-Za-z]:`) // keep deterministic on non-windows

// NormalizeRelativePath validates that a path is safe and relative, then normalizes it.
// It rejects absolute paths, any '.' or '..' segments, and empty segments.
// Returned paths always use forward slashes.
func NormalizeRelativePath(path string) (string, error) {
	trimmed := strings.TrimSpace(path)
	if trimmed == "" {
		return "", fmt.Errorf("path is required")
	}

	if strings.HasPrefix(trimmed, "/") || strings.HasPrefix(trimmed, "\\") {
		return "", fmt.Errorf("path must be relative: %q", path)
	}
	if filepath.IsAbs(trimmed) {
		return "", fmt.Errorf("path must be relative: %q", path)
	}
	if windowsDrivePrefix.MatchString(trimmed) {
		return "", fmt.Errorf("path must be relative (windows drive paths are not allowed): %q", path)
	}

	normalized := strings.ReplaceAll(trimmed, "\\", "/")
	parts := strings.SplitSeq(normalized, "/")
	for part := range parts {
		if part == "" {
			return "", fmt.Errorf("path contains empty segment: %q", path)
		}
		if part == "." || part == ".." {
			return "", fmt.Errorf("path contains disallowed segment %q: %q", part, path)
		}
	}

	cleaned := filepath.Clean(filepath.FromSlash(normalized))
	if cleaned == "." {
		return "", fmt.Errorf("path is invalid: %q", path)
	}
	if strings.HasPrefix(cleaned, ".."+string(filepath.Separator)) || cleaned == ".." {
		return "", fmt.Errorf("path escapes output root: %q", path)
	}

	return filepath.ToSlash(cleaned), nil
}
