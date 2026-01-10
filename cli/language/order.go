package language

import (
	"fmt"
	"sort"
)

// SortGeneratedFiles sorts by normalized Path for deterministic output.
func SortGeneratedFiles(files []GeneratedFile) {
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Path < files[j].Path
	})
}

// ValidateGeneratedFiles checks file paths are safe and non-colliding.
// It expects caller-visible paths to already be normalized.
func ValidateGeneratedFiles(files []GeneratedFile) error {
	seen := make(map[string]bool, len(files))
	for _, file := range files {
		if file.Path == "" {
			return fmt.Errorf("generated file path is required")
		}
		if seen[file.Path] {
			return fmt.Errorf("duplicate generated file path: %s", file.Path)
		}
		seen[file.Path] = true
	}
	return nil
}
