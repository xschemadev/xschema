// Package unsupported defines JSON Schema features that cannot be converted to
// static validator code. Used by both validator (to reject schemas) and
// compliance (to skip tests).
package unsupported

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"
)

//go:embed unsupported-features.json
var featuresData []byte

// FeatureGroup represents a group of unsupported features with a common reason
type FeatureGroup struct {
	Name     string   `json:"name"`
	Reason   string   `json:"reason"`
	Keywords []string `json:"keywords"`
	Tests    []string `json:"tests"`
}

// Features is a list of unsupported feature groups
type Features []FeatureGroup

// cached data (loaded once)
var (
	loadOnce sync.Once
	features Features
	keywords map[string]string
)

// load parses the embedded JSON once
func load() {
	loadOnce.Do(func() {
		if err := json.Unmarshal(featuresData, &features); err != nil {
			// Should never happen with embedded data
			features = Features{}
			keywords = make(map[string]string)
			return
		}
		keywords = make(map[string]string)
		for _, group := range features {
			for _, kw := range group.Keywords {
				keywords[kw] = group.Reason
			}
		}
	})
}

// Load returns all unsupported feature groups
func Load() Features {
	load()
	return features
}

// Keywords returns a map of keyword -> reason for all unsupported keywords
func Keywords() map[string]string {
	load()
	return keywords
}

// ContainsTest checks if a test path is in the unsupported features list
func (f Features) ContainsTest(testPath string) (bool, string) {
	for _, group := range f {
		for _, test := range group.Tests {
			if test == testPath {
				return true, group.Reason
			}
		}
	}
	return false, ""
}

// TestPaths returns all test paths as a flat list
func (f Features) TestPaths() []string {
	var paths []string
	for _, group := range f {
		paths = append(paths, group.Tests...)
	}
	return paths
}

// TestCount returns the total number of unsupported feature tests
func (f Features) TestCount() int {
	count := 0
	for _, group := range f {
		count += len(group.Tests)
	}
	return count
}

// ValidateKeywords checks a parsed schema for unsupported keywords.
// Returns an error if any unsupported keyword is found.
func ValidateKeywords(node any) error {
	load()
	return validateNode(node, "")
}

func validateNode(node any, path string) error {
	switch v := node.(type) {
	case map[string]any:
		return validateObject(v, path)
	case []any:
		for i, item := range v {
			if err := validateNode(item, fmt.Sprintf("%s/%d", path, i)); err != nil {
				return err
			}
		}
	}
	return nil
}

func validateObject(obj map[string]any, path string) error {
	for keyword, reason := range keywords {
		if _, ok := obj[keyword]; ok {
			loc := "root"
			if path != "" {
				loc = path
			}
			return fmt.Errorf("%s is not supported: %s (at %s)", keyword, reason, loc)
		}
	}

	for k, v := range obj {
		if err := validateNode(v, path+"/"+k); err != nil {
			return err
		}
	}
	return nil
}
