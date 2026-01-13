package unsupported

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	features := Load()
	if len(features) == 0 {
		t.Fatal("expected unsupported features to be loaded")
	}

	// Check that known features exist
	found := false
	for _, f := range features {
		if f.Name == "dynamic-refs" {
			found = true
			if len(f.Keywords) == 0 {
				t.Error("dynamic-refs should have keywords")
			}
			// Tests array is now empty - keyword detection handles this
		}
	}
	if !found {
		t.Error("expected dynamic-refs feature to exist")
	}
}

func TestKeywords(t *testing.T) {
	keywords := Keywords()
	if len(keywords) == 0 {
		t.Fatal("expected unsupported keywords")
	}

	// Check known unsupported keywords
	expected := []string{"$dynamicRef", "$dynamicAnchor", "$recursiveRef", "$recursiveAnchor"}
	for _, kw := range expected {
		if _, ok := keywords[kw]; !ok {
			t.Errorf("expected %s to be unsupported", kw)
		}
	}
}

func TestValidateKeywords(t *testing.T) {
	tests := []struct {
		name    string
		schema  any
		wantErr bool
	}{
		{
			name:    "valid schema",
			schema:  map[string]any{"type": "string"},
			wantErr: false,
		},
		{
			name:    "schema with $dynamicRef",
			schema:  map[string]any{"$dynamicRef": "#foo"},
			wantErr: true,
		},
		{
			name:    "schema with $dynamicAnchor",
			schema:  map[string]any{"$dynamicAnchor": "foo"},
			wantErr: true,
		},
		{
			name:    "schema with $recursiveRef",
			schema:  map[string]any{"$recursiveRef": "#"},
			wantErr: true,
		},
		{
			name:    "nested unsupported keyword",
			schema:  map[string]any{"properties": map[string]any{"foo": map[string]any{"$dynamicRef": "#bar"}}},
			wantErr: true,
		},
		{
			name:    "deeply nested valid",
			schema:  map[string]any{"properties": map[string]any{"foo": map[string]any{"type": "string"}}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateKeywords(tt.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateKeywords() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContainsTest(t *testing.T) {
	features := Load()

	// Tests array is now empty - keyword detection handles unsupported schemas
	// ContainsTest always returns false with empty tests arrays
	found, _ := features.ContainsTest("draft2020-12/dynamicRef/$dynamicRef points to a boolean schema/follow $dynamicRef to a false schema")
	if found {
		t.Error("expected test NOT to be found (tests arrays are empty)")
	}

	// Test path that should be supported - still returns false
	found, _ = features.ContainsTest("draft2020-12/type/type is string")
	if found {
		t.Error("expected test to NOT be in unsupported list")
	}
}

func TestTestCount(t *testing.T) {
	features := Load()
	count := features.TestCount()
	// Tests arrays are now empty - keyword detection handles unsupported schemas
	if count != 0 {
		t.Errorf("expected zero test count (tests arrays are empty), got %d", count)
	}
}

func TestContextAwareUnevaluatedProperties(t *testing.T) {
	tests := []struct {
		name    string
		schema  map[string]any
		wantErr bool
	}{
		{
			name:    "standalone unevaluatedProperties is allowed",
			schema:  map[string]any{"type": "object", "unevaluatedProperties": false},
			wantErr: false,
		},
		{
			name:    "unevaluatedProperties + allOf is unsupported",
			schema:  map[string]any{"unevaluatedProperties": false, "allOf": []any{map[string]any{"type": "object"}}},
			wantErr: true,
		},
		{
			name:    "unevaluatedProperties + anyOf is unsupported",
			schema:  map[string]any{"unevaluatedProperties": false, "anyOf": []any{map[string]any{"type": "object"}}},
			wantErr: true,
		},
		{
			name:    "unevaluatedProperties + oneOf is unsupported",
			schema:  map[string]any{"unevaluatedProperties": false, "oneOf": []any{map[string]any{"type": "object"}}},
			wantErr: true,
		},
		{
			name:    "unevaluatedProperties + if is unsupported",
			schema:  map[string]any{"unevaluatedProperties": false, "if": map[string]any{"type": "object"}},
			wantErr: true,
		},
		{
			name:    "unevaluatedProperties + $ref is unsupported",
			schema:  map[string]any{"unevaluatedProperties": false, "$ref": "#/$defs/foo"},
			wantErr: true,
		},
		{
			name:    "unevaluatedProperties + dependentSchemas is unsupported",
			schema:  map[string]any{"unevaluatedProperties": false, "dependentSchemas": map[string]any{"foo": map[string]any{}}},
			wantErr: true,
		},
		{
			name:    "unevaluatedProperties + not is unsupported",
			schema:  map[string]any{"unevaluatedProperties": false, "not": map[string]any{"type": "null"}},
			wantErr: true,
		},
		{
			name: "unevaluatedProperties nested with applicator is unsupported",
			schema: map[string]any{
				"properties": map[string]any{
					"nested": map[string]any{
						"unevaluatedProperties": false,
						"allOf":                 []any{map[string]any{"type": "object"}},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateKeywords(tt.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateKeywords() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContextAwareUnevaluatedItems(t *testing.T) {
	tests := []struct {
		name    string
		schema  map[string]any
		wantErr bool
	}{
		{
			name:    "standalone unevaluatedItems is allowed",
			schema:  map[string]any{"type": "array", "unevaluatedItems": false},
			wantErr: false,
		},
		{
			name:    "unevaluatedItems + prefixItems is unsupported",
			schema:  map[string]any{"unevaluatedItems": false, "prefixItems": []any{map[string]any{"type": "string"}}},
			wantErr: true,
		},
		{
			name:    "unevaluatedItems + contains is unsupported",
			schema:  map[string]any{"unevaluatedItems": false, "contains": map[string]any{"type": "string"}},
			wantErr: true,
		},
		{
			name:    "unevaluatedItems + allOf is unsupported",
			schema:  map[string]any{"unevaluatedItems": false, "allOf": []any{map[string]any{"type": "array"}}},
			wantErr: true,
		},
		{
			name:    "unevaluatedItems + anyOf is unsupported",
			schema:  map[string]any{"unevaluatedItems": false, "anyOf": []any{map[string]any{"type": "array"}}},
			wantErr: true,
		},
		{
			name:    "unevaluatedItems + oneOf is unsupported",
			schema:  map[string]any{"unevaluatedItems": false, "oneOf": []any{map[string]any{"type": "array"}}},
			wantErr: true,
		},
		{
			name:    "unevaluatedItems + if is unsupported",
			schema:  map[string]any{"unevaluatedItems": false, "if": map[string]any{"type": "array"}},
			wantErr: true,
		},
		{
			name: "unevaluatedItems nested with applicator is unsupported",
			schema: map[string]any{
				"items": map[string]any{
					"unevaluatedItems": false,
					"prefixItems":      []any{map[string]any{"type": "string"}},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateKeywords(tt.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateKeywords() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// testSuiteSchema represents a test group from the JSON Schema Test Suite
type testSuiteSchema struct {
	Description string `json:"description"`
	Schema      any    `json:"schema"`
}

// TestValidateKeywords_RealTestSuiteSchemas loads actual schemas from the JSON Schema Test Suite
// and verifies that ValidateKeywords correctly detects unsupported keywords.
// This is the "opposite" compliance test - ensuring our detection catches real unsupported schemas.
func TestValidateKeywords_RealTestSuiteSchemas(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("cannot get home directory:", err)
	}

	testSuiteDir := filepath.Join(homeDir, ".cache", "xschema", "json-schema-test-suite", "tests")
	if _, err := os.Stat(testSuiteDir); os.IsNotExist(err) {
		t.Skip("JSON Schema Test Suite not found at:", testSuiteDir)
	}

	// Test cases: each contains a test file path and the minimum number of schemas
	// we expect to detect as unsupported. This is more robust than expecting exact counts
	// since test suite files may have schemas that reference external URLs (which we don't validate inline).
	tests := []struct {
		name       string
		draft      string
		file       string
		minErrors  int // minimum number of schemas that should error
		minAllowed int // minimum number of schemas allowed (0 means we don't care)
	}{
		// $dynamicRef schemas - most use $dynamicRef/$dynamicAnchor directly,
		// but one (#18: "$ref to $dynamicRef finds detached $dynamicAnchor") only has $ref
		// to an external URL, so it passes (the remote schema would be validated separately)
		{
			name:       "dynamicRef schemas should be detected",
			draft:      "draft2020-12",
			file:       "dynamicRef.json",
			minErrors:  20, // 20 out of 21 have $dynamicRef or $dynamicAnchor
			minAllowed: 0,  // at least 1 is just $ref to external URL
		},
		// $recursiveRef schemas - all use $recursiveRef/$recursiveAnchor
		{
			name:       "recursiveRef schemas should all be detected",
			draft:      "draft2019-09",
			file:       "recursiveRef.json",
			minErrors:  8, // all 8 schemas use $recursiveRef or $recursiveAnchor
			minAllowed: 0,
		},
		// unevaluatedProperties - some are standalone (allowed), some have applicators (error)
		{
			name:       "unevaluatedProperties with applicators detected",
			draft:      "draft2020-12",
			file:       "unevaluatedProperties.json",
			minErrors:  20, // many have applicators like allOf, $ref, etc.
			minAllowed: 10, // some are standalone (allowed)
		},
		// unevaluatedItems - some are standalone (allowed), some have applicators (error)
		{
			name:       "unevaluatedItems with applicators detected",
			draft:      "draft2020-12",
			file:       "unevaluatedItems.json",
			minErrors:  15, // many have prefixItems, contains, allOf, etc.
			minAllowed: 5,  // some are standalone (allowed)
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			filePath := filepath.Join(testSuiteDir, tc.draft, tc.file)
			data, err := os.ReadFile(filePath)
			if err != nil {
				t.Fatalf("failed to read %s: %v", filePath, err)
			}

			var testGroups []testSuiteSchema
			if err := json.Unmarshal(data, &testGroups); err != nil {
				t.Fatalf("failed to parse %s: %v", filePath, err)
			}

			if len(testGroups) == 0 {
				t.Fatal("no test groups found in", filePath)
			}

			var errorCount, successCount int
			for _, group := range testGroups {
				unsupErr := ValidateKeywords(group.Schema)
				if unsupErr != nil {
					errorCount++
					// ValidateKeywords returns *UnsupportedKeywordError directly
					// so we just verify it has the expected fields
					if unsupErr.Keyword == "" {
						t.Errorf("expected Keyword to be set, got empty string")
					}
				} else {
					successCount++
				}
			}

			// Verify minimum error count
			if errorCount < tc.minErrors {
				t.Errorf("expected at least %d schemas to error, but only %d did", tc.minErrors, errorCount)
			}
			// Verify minimum allowed count (if specified)
			if tc.minAllowed > 0 && successCount < tc.minAllowed {
				t.Errorf("expected at least %d schemas to be allowed, but only %d were", tc.minAllowed, successCount)
			}
			// Log the counts for debugging
			t.Logf("%s: %d errors, %d allowed (total: %d)", tc.file, errorCount, successCount, len(testGroups))
		})
	}
}

// TestValidateKeywords_SpecificUnsupportedSchemas tests specific schemas that MUST be detected
func TestValidateKeywords_SpecificUnsupportedSchemas(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("cannot get home directory:", err)
	}

	testSuiteDir := filepath.Join(homeDir, ".cache", "xschema", "json-schema-test-suite", "tests")
	if _, err := os.Stat(testSuiteDir); os.IsNotExist(err) {
		t.Skip("JSON Schema Test Suite not found at:", testSuiteDir)
	}

	// Load a specific schema that we know uses unevaluatedProperties + allOf
	filePath := filepath.Join(testSuiteDir, "draft2020-12", "unevaluatedProperties.json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read %s: %v", filePath, err)
	}

	var testGroups []testSuiteSchema
	if err := json.Unmarshal(data, &testGroups); err != nil {
		t.Fatalf("failed to parse %s: %v", filePath, err)
	}

	// Find the specific test "unevaluatedProperties with nested properties" which has allOf
	var foundAndTested bool
	for _, group := range testGroups {
		if group.Description == "unevaluatedProperties with nested properties" {
			unsupErr := ValidateKeywords(group.Schema)
			if unsupErr == nil {
				t.Error("expected error for 'unevaluatedProperties with nested properties' but got nil")
			} else if unsupErr.Keyword != "unevaluatedProperties" {
				t.Errorf("expected keyword 'unevaluatedProperties', got %q", unsupErr.Keyword)
			}
			foundAndTested = true
			break
		}
	}

	if !foundAndTested {
		t.Error("could not find test 'unevaluatedProperties with nested properties'")
	}
}
