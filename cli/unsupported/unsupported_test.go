package unsupported

import (
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
			if len(f.Tests) == 0 {
				t.Error("dynamic-refs should have tests")
			}
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

	// Test path that should be unsupported
	found, reason := features.ContainsTest("draft2020-12/dynamicRef/$dynamicRef points to a boolean schema/follow $dynamicRef to a false schema")
	if !found {
		t.Error("expected test to be found in unsupported list")
	}
	if reason == "" {
		t.Error("expected reason to be non-empty")
	}

	// Test path that should be supported
	found, _ = features.ContainsTest("draft2020-12/type/type is string")
	if found {
		t.Error("expected test to NOT be in unsupported list")
	}
}

func TestTestCount(t *testing.T) {
	features := Load()
	count := features.TestCount()
	if count == 0 {
		t.Error("expected non-zero test count")
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
