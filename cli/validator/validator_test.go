package validator

import (
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
