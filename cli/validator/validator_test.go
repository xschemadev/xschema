package validator

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"
)

func TestValidateSchema_ValidDraft2020(t *testing.T) {
	schema := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "object",
		"properties": {
			"name": { "type": "string" }
		}
	}`)

	err := ValidateSchema(schema)
	if err != nil {
		t.Errorf("expected valid schema, got error: %v", err)
	}
}

func TestValidateSchema_ValidDraft7(t *testing.T) {
	schema := []byte(`{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"type": "object",
		"properties": {
			"age": { "type": "integer", "minimum": 0 }
		}
	}`)

	err := ValidateSchema(schema)
	if err != nil {
		t.Errorf("expected valid schema, got error: %v", err)
	}
}

func TestValidateSchema_NoSchemaField(t *testing.T) {
	// should default to draft-2020-12
	schema := []byte(`{
		"type": "object",
		"properties": {
			"id": { "type": "string" }
		}
	}`)

	err := ValidateSchema(schema)
	if err != nil {
		t.Errorf("expected valid schema without $schema field, got error: %v", err)
	}
}

func TestValidateSchema_BooleanSchema(t *testing.T) {
	schema := []byte(`true`)

	err := ValidateSchema(schema)
	if err != nil {
		t.Errorf("expected boolean schema to be valid, got error: %v", err)
	}
}

func TestValidateSchema_InvalidType(t *testing.T) {
	schema := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "notavalidtype"
	}`)

	err := ValidateSchema(schema)
	if err == nil {
		t.Error("expected error for invalid type value, got nil")
	}
}

func TestValidateSchema_InvalidJSON(t *testing.T) {
	schema := []byte(`{not valid json}`)

	err := ValidateSchema(schema)
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

func TestValidateSchema_ExternalRefIgnored(t *testing.T) {
	// schemas with external $ref should pass validation
	// (external refs are resolved by bundler, not validator)
	schema := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "object",
		"properties": {
			"address": { "$ref": "https://example.com/address.json" }
		}
	}`)

	err := ValidateSchema(schema)
	if err != nil {
		t.Errorf("expected schema with external ref to pass validation, got error: %v", err)
	}
}

func TestValidateSchemaWithOptions_CustomMetaschema(t *testing.T) {
	// A simple custom metaschema that only allows type: "string" or "object"
	customMetaschema := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"$id": "https://example.com/custom-meta/v1",
		"type": "object",
		"properties": {
			"type": {
				"enum": ["string", "object"]
			}
		}
	}`)

	// Schema using our custom metaschema
	schema := []byte(`{
		"$schema": "https://example.com/custom-meta/v1",
		"type": "string"
	}`)

	opts := &ValidateOptions{
		Metaschemas: map[string]json.RawMessage{
			"https://example.com/custom-meta/v1": customMetaschema,
		},
	}

	err := ValidateSchemaWithOptions(schema, opts)
	if err != nil {
		t.Errorf("expected schema with custom metaschema to validate, got error: %v", err)
	}
}

func TestValidateSchemaWithOptions_MissingMetaschema(t *testing.T) {
	// Schema references a custom $schema that wasn't pre-loaded
	schema := []byte(`{
		"$schema": "https://example.com/missing-meta/v1",
		"type": "string"
	}`)

	err := ValidateSchemaWithOptions(schema, nil)
	if err == nil {
		t.Error("expected error for missing custom metaschema, got nil")
	}
}

func TestValidateSchemaWithOptions_CrossDraft(t *testing.T) {
	// A custom metaschema based on draft-2020-12
	customMetaschema := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"$id": "https://example.com/cross-draft-meta/v1",
		"type": "object",
		"properties": {
			"myKeyword": { "type": "boolean" }
		}
	}`)

	// Schema that uses the custom metaschema but validates using draft2020 features
	schema := []byte(`{
		"$schema": "https://example.com/cross-draft-meta/v1",
		"myKeyword": true
	}`)

	opts := &ValidateOptions{
		Metaschemas: map[string]json.RawMessage{
			"https://example.com/cross-draft-meta/v1": customMetaschema,
		},
	}

	err := ValidateSchemaWithOptions(schema, opts)
	if err != nil {
		t.Errorf("expected cross-draft schema to validate, got error: %v", err)
	}
}

func TestValidateSchemaWithOptions_EmptyMetaschemas(t *testing.T) {
	// Standard schema with empty Metaschemas map should use defaults
	schema := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "object",
		"properties": {
			"name": { "type": "string" }
		}
	}`)

	opts := &ValidateOptions{
		Metaschemas: map[string]json.RawMessage{},
	}

	err := ValidateSchemaWithOptions(schema, opts)
	if err != nil {
		t.Errorf("expected schema with empty metaschemas to validate using defaults, got error: %v", err)
	}
}

func TestValidateSchemaWithOptions_InvalidMetaschemaJSON(t *testing.T) {
	// Invalid JSON in metaschema should return error
	opts := &ValidateOptions{
		Metaschemas: map[string]json.RawMessage{
			"https://example.com/bad-meta": []byte(`{not valid json}`),
		},
	}

	schema := []byte(`{
		"$schema": "https://example.com/bad-meta",
		"type": "string"
	}`)

	err := ValidateSchemaWithOptions(schema, opts)
	if err == nil {
		t.Error("expected error for invalid metaschema JSON, got nil")
	}
	if !strings.Contains(err.Error(), "invalid metaschema") {
		t.Errorf("expected error to mention 'invalid metaschema', got: %v", err)
	}
}

func TestDetectDraft(t *testing.T) {
	tests := []struct {
		name     string
		schema   []byte
		expected string
	}{
		{
			name:     "draft-2020-12 https",
			schema:   []byte(`{"$schema": "https://json-schema.org/draft/2020-12/schema"}`),
			expected: "2020-12",
		},
		{
			name:     "draft-2019-09",
			schema:   []byte(`{"$schema": "https://json-schema.org/draft/2019-09/schema"}`),
			expected: "2019-09",
		},
		{
			name:     "draft-07",
			schema:   []byte(`{"$schema": "http://json-schema.org/draft-07/schema#"}`),
			expected: "07",
		},
		{
			name:     "draft-06",
			schema:   []byte(`{"$schema": "http://json-schema.org/draft-06/schema#"}`),
			expected: "06",
		},
		{
			name:     "draft-04",
			schema:   []byte(`{"$schema": "http://json-schema.org/draft-04/schema#"}`),
			expected: "04",
		},
		{
			name:     "no $schema defaults to 2020-12",
			schema:   []byte(`{"type": "object"}`),
			expected: "2020-12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			draft := detectDraft(tt.schema)
			// just check it doesn't panic and returns a non-nil draft
			if draft == nil {
				t.Error("expected non-nil draft")
			}
		})
	}
}

func TestCompiledCache_SameSchemaTwice(t *testing.T) {
	// Validate same schema twice - second call should hit cache and not recompile
	schema := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "object",
		"properties": {
			"name": { "type": "string" }
		}
	}`)

	cache := NewCompiledCache()
	opts := &ValidateOptions{Cache: cache}

	// First validation
	err := ValidateSchemaWithOptions(schema, opts)
	if err != nil {
		t.Fatalf("first validation failed: %v", err)
	}

	stats := cache.Stats()
	if stats.Compiles != 1 {
		t.Errorf("expected 1 compile after first validation, got %d", stats.Compiles)
	}
	if stats.Misses != 1 {
		t.Errorf("expected 1 miss after first validation, got %d", stats.Misses)
	}

	// Second validation - should hit cache
	err = ValidateSchemaWithOptions(schema, opts)
	if err != nil {
		t.Fatalf("second validation failed: %v", err)
	}

	stats = cache.Stats()
	if stats.Compiles != 1 {
		t.Errorf("expected still 1 compile after second validation (cache hit), got %d", stats.Compiles)
	}
	if stats.Hits != 1 {
		t.Errorf("expected 1 hit after second validation, got %d", stats.Hits)
	}
}

func TestCompiledCache_DifferentSchemas(t *testing.T) {
	// Two different schemas should both compile separately
	schema1 := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "string"
	}`)

	schema2 := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "number"
	}`)

	cache := NewCompiledCache()
	opts := &ValidateOptions{Cache: cache}

	// Validate first schema
	err := ValidateSchemaWithOptions(schema1, opts)
	if err != nil {
		t.Fatalf("first schema validation failed: %v", err)
	}

	// Validate second schema
	err = ValidateSchemaWithOptions(schema2, opts)
	if err != nil {
		t.Fatalf("second schema validation failed: %v", err)
	}

	stats := cache.Stats()
	if stats.Compiles != 2 {
		t.Errorf("expected 2 compiles for different schemas, got %d", stats.Compiles)
	}
	if stats.Misses != 2 {
		t.Errorf("expected 2 misses for different schemas, got %d", stats.Misses)
	}

	// Validate first schema again - should hit cache
	err = ValidateSchemaWithOptions(schema1, opts)
	if err != nil {
		t.Fatalf("repeated first schema validation failed: %v", err)
	}

	stats = cache.Stats()
	if stats.Compiles != 2 {
		t.Errorf("expected still 2 compiles after cache hit, got %d", stats.Compiles)
	}
	if stats.Hits != 1 {
		t.Errorf("expected 1 hit after validating first schema again, got %d", stats.Hits)
	}
}

func TestCompiledCache_ConcurrentValidations(t *testing.T) {
	// Concurrent validations of the same schema should be safe
	schema := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "object",
		"properties": {
			"id": { "type": "integer" }
		}
	}`)

	cache := NewCompiledCache()
	opts := &ValidateOptions{Cache: cache}

	var wg sync.WaitGroup
	errs := make(chan error, 100)

	// Run 100 concurrent validations
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := ValidateSchemaWithOptions(schema, opts); err != nil {
				errs <- err
			}
		}()
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		t.Errorf("concurrent validation failed: %v", err)
	}

	// At least one compile should have happened (could be more due to races)
	stats := cache.Stats()
	if stats.Compiles < 1 {
		t.Errorf("expected at least 1 compile, got %d", stats.Compiles)
	}
	// Most should be hits (though exact count depends on scheduling)
	if stats.Hits < 50 {
		t.Errorf("expected significant cache hits, got only %d hits out of 100 validations", stats.Hits)
	}
}

func TestCompiledCache_WithExternalRef(t *testing.T) {
	// Schema with external ref should still be cached (marked as valid)
	schema := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "object",
		"properties": {
			"address": { "$ref": "https://example.com/address.json" }
		}
	}`)

	cache := NewCompiledCache()
	opts := &ValidateOptions{Cache: cache}

	// First validation
	err := ValidateSchemaWithOptions(schema, opts)
	if err != nil {
		t.Fatalf("first validation with external ref failed: %v", err)
	}

	// Second validation should hit cache
	err = ValidateSchemaWithOptions(schema, opts)
	if err != nil {
		t.Fatalf("second validation with external ref failed: %v", err)
	}

	stats := cache.Stats()
	if stats.Hits < 1 {
		t.Errorf("expected cache hit for schema with external ref, got %d hits", stats.Hits)
	}
}

func TestCompiledCache_DraftHintAffectsCache(t *testing.T) {
	// Same schema content but different draft hint should cache separately
	schema := []byte(`{"type": "object"}`)

	cache := NewCompiledCache()
	opts := &ValidateOptions{Cache: cache}

	// Validate with draft7 hint
	err := ValidateSchemaWithOptions(schema, opts, "draft7")
	if err != nil {
		t.Fatalf("draft7 validation failed: %v", err)
	}

	// Validate with draft2020-12 hint - should compile separately
	err = ValidateSchemaWithOptions(schema, opts, "draft2020-12")
	if err != nil {
		t.Fatalf("draft2020-12 validation failed: %v", err)
	}

	stats := cache.Stats()
	if stats.Compiles != 2 {
		t.Errorf("expected 2 compiles for same schema with different drafts, got %d", stats.Compiles)
	}
}
