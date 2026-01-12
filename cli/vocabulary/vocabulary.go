// Package vocabulary handles JSON Schema vocabulary-based filtering.
// It strips keywords from schemas based on which vocabularies are enabled
// in the metaschema's $vocabulary field.
package vocabulary

import (
	"encoding/json"
)

// Standard JSON Schema vocabulary URIs
const (
	Validation2020 = "https://json-schema.org/draft/2020-12/vocab/validation"
	Validation2019 = "https://json-schema.org/draft/2019-09/vocab/validation"
	Applicator2020 = "https://json-schema.org/draft/2020-12/vocab/applicator"
	Applicator2019 = "https://json-schema.org/draft/2019-09/vocab/applicator"
	Format2020     = "https://json-schema.org/draft/2020-12/vocab/format-annotation"
	Format2019     = "https://json-schema.org/draft/2019-09/vocab/format"
)

// validationKeywords are stripped when validation vocabulary is disabled.
// NOTE: "type" is NOT included - it's essential for IR generation.
var validationKeywords = map[string]bool{
	// String validation
	"minLength": true,
	"maxLength": true,
	"pattern":   true,
	// Numeric validation
	"minimum":          true,
	"maximum":          true,
	"exclusiveMinimum": true,
	"exclusiveMaximum": true,
	"multipleOf":       true,
	// Array validation
	"minItems":    true,
	"maxItems":    true,
	"uniqueItems": true,
	"minContains": true,
	"maxContains": true,
	// Object validation
	"minProperties":    true,
	"maxProperties":    true,
	"required":         true,
	"dependentRequired": true,
	// Enum/const (but NOT type)
	"enum":  true,
	"const": true,
}

// formatKeywords are stripped when format vocabulary is disabled.
var formatKeywords = map[string]bool{
	"format": true,
}

// FilterSchema strips keywords from disabled vocabularies.
// If vocab is nil, returns schema unchanged (all vocabularies enabled by default).
// If vocab is empty map, strips all vocabulary-controlled keywords.
func FilterSchema(schema json.RawMessage, vocab map[string]bool) (json.RawMessage, error) {
	if vocab == nil {
		return schema, nil
	}

	var parsed any
	if err := json.Unmarshal(schema, &parsed); err != nil {
		return schema, nil // not valid JSON, return as-is
	}

	validationEnabled := isVocabEnabled(vocab, Validation2020, Validation2019)
	formatEnabled := isVocabEnabled(vocab, Format2020, Format2019)

	filtered := filterNode(parsed, validationEnabled, formatEnabled)

	return json.Marshal(filtered)
}

// isVocabEnabled checks if any of the given vocabulary URIs is present in vocab.
// Per JSON Schema spec, presence in $vocabulary (regardless of true/false value) means enabled.
func isVocabEnabled(vocab map[string]bool, uris ...string) bool {
	for _, uri := range uris {
		if _, ok := vocab[uri]; ok {
			return true
		}
	}
	return false
}

// filterNode recursively filters a schema node.
func filterNode(node any, validationEnabled, formatEnabled bool) any {
	switch v := node.(type) {
	case map[string]any:
		return filterObject(v, validationEnabled, formatEnabled)
	case []any:
		return filterArray(v, validationEnabled, formatEnabled)
	default:
		return node
	}
}

// filterObject filters an object schema node, removing disabled keywords.
func filterObject(obj map[string]any, validationEnabled, formatEnabled bool) map[string]any {
	result := make(map[string]any, len(obj))

	for key, val := range obj {
		// Check if keyword should be stripped
		if !validationEnabled && validationKeywords[key] {
			continue
		}
		if !formatEnabled && formatKeywords[key] {
			continue
		}

		// Recursively filter nested schemas
		result[key] = filterNode(val, validationEnabled, formatEnabled)
	}

	return result
}

// filterArray filters an array, recursively filtering each element.
func filterArray(arr []any, validationEnabled, formatEnabled bool) []any {
	result := make([]any, len(arr))
	for i, elem := range arr {
		result[i] = filterNode(elem, validationEnabled, formatEnabled)
	}
	return result
}
