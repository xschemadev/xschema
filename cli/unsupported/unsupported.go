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

// UnsupportedKeywordError is returned when a schema contains an unsupported keyword
type UnsupportedKeywordError struct {
	Keyword string // the unsupported keyword (e.g., "$dynamicRef")
	Reason  string // why this keyword is unsupported
	Path    string // JSON pointer path to the keyword (e.g., "/properties/foo")
}

// Error implements the error interface
func (e *UnsupportedKeywordError) Error() string {
	loc := "root"
	if e.Path != "" {
		loc = e.Path
	}
	return fmt.Sprintf("%s is not supported: %s (at %s)", e.Keyword, e.Reason, loc)
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
// Returns *UnsupportedKeywordError if any unsupported keyword is found, nil otherwise.
func ValidateKeywords(node any) *UnsupportedKeywordError {
	load()
	return validateNode(node, "")
}

func validateNode(node any, path string) *UnsupportedKeywordError {
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

func validateObject(obj map[string]any, path string) *UnsupportedKeywordError {
	for keyword, reason := range keywords {
		if _, ok := obj[keyword]; ok {
			return &UnsupportedKeywordError{
				Keyword: keyword,
				Reason:  reason,
				Path:    path,
			}
		}
	}

	for k, v := range obj {
		if err := validateNode(v, path+"/"+k); err != nil {
			return err
		}
	}
	return nil
}
