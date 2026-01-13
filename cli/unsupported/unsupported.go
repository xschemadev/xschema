// Package unsupported defines JSON Schema features that cannot be converted to
// static validator code. Used by both validator (to reject schemas) and
// compliance (to skip tests).
package unsupported

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

//go:embed unsupported-features.json
var featuresData []byte

// FeatureGroup represents a group of unsupported features with a common reason
type FeatureGroup struct {
	Name     string   `json:"name"`
	Reason   string   `json:"reason"`
	Keywords []string `json:"keywords"`
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

// Property applicators that require annotation tracking when combined with unevaluatedProperties
var propertyApplicators = []string{"allOf", "anyOf", "oneOf", "if", "$ref", "dependentSchemas", "not"}

// Item applicators that require annotation tracking when combined with unevaluatedItems
var itemApplicators = []string{"prefixItems", "contains", "allOf", "anyOf", "oneOf", "if", "$ref", "items", "additionalItems", "not"}

// hasPropertyApplicators checks if an object has any keyword that applies schemas to properties.
// When unevaluatedProperties is combined with these, it needs annotation tracking.
func hasPropertyApplicators(obj map[string]any) bool {
	for _, kw := range propertyApplicators {
		if _, ok := obj[kw]; ok {
			return true
		}
	}
	return false
}

// hasItemApplicators checks if an object has any keyword that applies schemas to items.
// When unevaluatedItems is combined with these, it needs annotation tracking.
func hasItemApplicators(obj map[string]any) bool {
	for _, kw := range itemApplicators {
		if _, ok := obj[kw]; ok {
			return true
		}
	}
	return false
}

// ValidateKeywords checks a parsed schema for unsupported keywords.
// Returns *UnsupportedKeywordError if any unsupported keyword is found, nil otherwise.
func ValidateKeywords(node any) *UnsupportedKeywordError {
	load()
	if err := validateNode(node, "", false); err != nil {
		return err
	}

	// Check for cyclic $ref combined with unevaluatedProperties at root
	if obj, ok := node.(map[string]any); ok {
		if _, hasUnevalProps := obj["unevaluatedProperties"]; hasUnevalProps {
			if hasCyclicRef(obj) {
				return &UnsupportedKeywordError{
					Keyword: "unevaluatedProperties",
					Reason:  "cyclic $ref with unevaluatedProperties requires recursive annotation tracking",
					Path:    "",
				}
			}
		}
	}

	return nil
}

// hasCyclicRef checks if a schema contains a $ref that cycles back to the root.
// This includes direct "#" refs and refs through $defs that eventually reference root.
func hasCyclicRef(root map[string]any) bool {
	// Track which $defs we're currently visiting to detect cycles through $defs
	visiting := make(map[string]bool)
	return checkForCyclicRef(root, root, visiting)
}

// checkForCyclicRef recursively checks for cyclic refs
func checkForCyclicRef(node any, root map[string]any, visiting map[string]bool) bool {
	switch v := node.(type) {
	case map[string]any:
		// Check if this object has a $ref
		if ref, ok := v["$ref"].(string); ok {
			if isCyclicRef(ref, root, visiting) {
				return true
			}
		}
		// Recurse into all nested objects
		for _, child := range v {
			if checkForCyclicRef(child, root, visiting) {
				return true
			}
		}
	case []any:
		for _, item := range v {
			if checkForCyclicRef(item, root, visiting) {
				return true
			}
		}
	}
	return false
}

// isCyclicRef checks if a $ref value creates a cycle
func isCyclicRef(ref string, root map[string]any, visiting map[string]bool) bool {
	// Direct root reference
	if ref == "#" {
		return true
	}

	// Check refs to $defs - these might eventually cycle back
	if len(ref) > 8 && ref[:8] == "#/$defs/" {
		defName := ref[8:]
		// Check if there's a slash (nested path) and extract just the def name
		for i := 0; i < len(defName); i++ {
			if defName[i] == '/' {
				defName = defName[:i]
				break
			}
		}

		// If we're already visiting this def, we found a cycle
		if visiting[defName] {
			return true
		}

		// Look up the definition and check it for cycles
		if defs, ok := root["$defs"].(map[string]any); ok {
			if defSchema, ok := defs[defName]; ok {
				visiting[defName] = true
				result := checkForCyclicRef(defSchema, root, visiting)
				visiting[defName] = false
				return result
			}
		}
	}

	// Also check legacy definitions keyword
	if len(ref) > 14 && ref[:14] == "#/definitions/" {
		defName := ref[14:]
		for i := 0; i < len(defName); i++ {
			if defName[i] == '/' {
				defName = defName[:i]
				break
			}
		}

		if visiting[defName] {
			return true
		}

		if defs, ok := root["definitions"].(map[string]any); ok {
			if defSchema, ok := defs[defName]; ok {
				visiting[defName] = true
				result := checkForCyclicRef(defSchema, root, visiting)
				visiting[defName] = false
				return result
			}
		}
	}

	return false
}

func validateNode(node any, path string, insideApplicator bool) *UnsupportedKeywordError {
	switch v := node.(type) {
	case map[string]any:
		return validateObject(v, path, insideApplicator)
	case []any:
		for i, item := range v {
			if err := validateNode(item, fmt.Sprintf("%s/%d", path, i), insideApplicator); err != nil {
				return err
			}
		}
	}
	return nil
}

// applicatorKeywords are keywords that contain subschemas where unevaluated* can't see sibling annotations
var applicatorKeywords = map[string]bool{
	"allOf": true, "anyOf": true, "oneOf": true, "if": true, "then": true, "else": true,
}

func validateObject(obj map[string]any, path string, insideApplicator bool) *UnsupportedKeywordError {
	// Check for absolutely unsupported keywords (always error)
	// Sort keywords for deterministic error reporting (same keyword detected on every run)
	sortedKeywords := make([]string, 0, len(keywords))
	for kw := range keywords {
		sortedKeywords = append(sortedKeywords, kw)
	}
	sort.Strings(sortedKeywords)

	for _, keyword := range sortedKeywords {
		if _, ok := obj[keyword]; ok {
			return &UnsupportedKeywordError{
				Keyword: keyword,
				Reason:  keywords[keyword],
				Path:    path,
			}
		}
	}

	// Check for unevaluated keywords inside applicator subschemas (cousins problem)
	// When unevaluatedItems/unevaluatedProperties appears inside an allOf/anyOf/oneOf/if subschema,
	// it can't see annotations from sibling subschemas
	if insideApplicator {
		if _, hasUnevalProps := obj["unevaluatedProperties"]; hasUnevalProps {
			return &UnsupportedKeywordError{
				Keyword: "unevaluatedProperties",
				Reason:  "inside applicator subschema cannot see sibling annotations (cousins problem)",
				Path:    path,
			}
		}
		if _, hasUnevalItems := obj["unevaluatedItems"]; hasUnevalItems {
			return &UnsupportedKeywordError{
				Keyword: "unevaluatedItems",
				Reason:  "inside applicator subschema cannot see sibling annotations (cousins problem)",
				Path:    path,
			}
		}
	}

	// Check for context-aware unevaluated detection at same level
	// unevaluatedProperties is only unsupported when combined with property applicators
	if _, hasUnevalProps := obj["unevaluatedProperties"]; hasUnevalProps && hasPropertyApplicators(obj) {
		return &UnsupportedKeywordError{
			Keyword: "unevaluatedProperties",
			Reason:  "requires annotation tracking when combined with applicators (allOf, anyOf, oneOf, if, $ref, dependentSchemas, not)",
			Path:    path,
		}
	}

	// unevaluatedItems is only unsupported when combined with item applicators
	if _, hasUnevalItems := obj["unevaluatedItems"]; hasUnevalItems && hasItemApplicators(obj) {
		return &UnsupportedKeywordError{
			Keyword: "unevaluatedItems",
			Reason:  "requires annotation tracking when combined with applicators (prefixItems, contains, allOf, anyOf, oneOf, if)",
			Path:    path,
		}
	}

	// Recurse into nested schemas in sorted key order for deterministic traversal
	sortedKeys := make([]string, 0, len(obj))
	for k := range obj {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, k := range sortedKeys {
		// Set insideApplicator=true when entering applicator keyword children
		childInsideApplicator := insideApplicator || applicatorKeywords[k]
		if err := validateNode(obj[k], path+"/"+k, childInsideApplicator); err != nil {
			return err
		}
	}
	return nil
}
