package vocabulary

import (
	"encoding/json"
	"testing"
)

func TestFilterSchema_NilVocab(t *testing.T) {
	schema := json.RawMessage(`{"type":"string","minLength":1,"format":"email"}`)

	result, err := FilterSchema(schema, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should return unchanged
	var got map[string]any
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	if got["minLength"] != float64(1) {
		t.Errorf("minLength should be preserved, got %v", got["minLength"])
	}
	if got["format"] != "email" {
		t.Errorf("format should be preserved, got %v", got["format"])
	}
}

func TestFilterSchema_ValidationDisabled(t *testing.T) {
	schema := json.RawMessage(`{
		"type": "string",
		"minLength": 1,
		"maxLength": 100,
		"pattern": "^[a-z]+$",
		"format": "email"
	}`)

	// Only format vocab enabled
	vocab := map[string]bool{
		Format2020: true,
	}

	result, err := FilterSchema(schema, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	// type should be preserved (not filtered)
	if got["type"] != "string" {
		t.Errorf("type should be preserved, got %v", got["type"])
	}

	// validation keywords should be stripped
	if _, ok := got["minLength"]; ok {
		t.Error("minLength should be stripped")
	}
	if _, ok := got["maxLength"]; ok {
		t.Error("maxLength should be stripped")
	}
	if _, ok := got["pattern"]; ok {
		t.Error("pattern should be stripped")
	}

	// format should be preserved
	if got["format"] != "email" {
		t.Errorf("format should be preserved, got %v", got["format"])
	}
}

func TestFilterSchema_FormatDisabled(t *testing.T) {
	schema := json.RawMessage(`{
		"type": "string",
		"minLength": 1,
		"format": "email"
	}`)

	// Only validation vocab enabled
	vocab := map[string]bool{
		Validation2020: true,
	}

	result, err := FilterSchema(schema, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	// validation keywords should be preserved
	if got["minLength"] != float64(1) {
		t.Errorf("minLength should be preserved, got %v", got["minLength"])
	}

	// format should be stripped
	if _, ok := got["format"]; ok {
		t.Error("format should be stripped")
	}
}

func TestFilterSchema_EmptyVocab(t *testing.T) {
	schema := json.RawMessage(`{
		"type": "string",
		"minLength": 1,
		"format": "email",
		"required": ["name"]
	}`)

	// Empty vocab = no vocabularies enabled
	vocab := map[string]bool{}

	result, err := FilterSchema(schema, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	// type should be preserved (not a vocabulary keyword)
	if got["type"] != "string" {
		t.Errorf("type should be preserved, got %v", got["type"])
	}

	// all vocabulary keywords should be stripped
	if _, ok := got["minLength"]; ok {
		t.Error("minLength should be stripped")
	}
	if _, ok := got["format"]; ok {
		t.Error("format should be stripped")
	}
	if _, ok := got["required"]; ok {
		t.Error("required should be stripped")
	}
}

func TestFilterSchema_NestedSchemas(t *testing.T) {
	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"name": {
				"type": "string",
				"minLength": 1
			},
			"age": {
				"type": "integer",
				"minimum": 0
			}
		}
	}`)

	// Only applicator vocab enabled (not in our filter list)
	vocab := map[string]bool{
		Applicator2020: true,
	}

	result, err := FilterSchema(schema, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	// properties should be preserved (applicator)
	props, ok := got["properties"].(map[string]any)
	if !ok {
		t.Fatal("properties should be preserved")
	}

	// nested validation keywords should be stripped
	name := props["name"].(map[string]any)
	if _, ok := name["minLength"]; ok {
		t.Error("nested minLength should be stripped")
	}

	age := props["age"].(map[string]any)
	if _, ok := age["minimum"]; ok {
		t.Error("nested minimum should be stripped")
	}
}

func TestFilterSchema_Defs(t *testing.T) {
	schema := json.RawMessage(`{
		"$defs": {
			"PositiveInt": {
				"type": "integer",
				"minimum": 1
			}
		},
		"$ref": "#/$defs/PositiveInt"
	}`)

	vocab := map[string]bool{
		Applicator2020: true,
	}

	result, err := FilterSchema(schema, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	// $defs should be preserved
	defs, ok := got["$defs"].(map[string]any)
	if !ok {
		t.Fatal("$defs should be preserved")
	}

	// nested validation keywords should be stripped
	posInt := defs["PositiveInt"].(map[string]any)
	if _, ok := posInt["minimum"]; ok {
		t.Error("minimum in $defs should be stripped")
	}
	if posInt["type"] != "integer" {
		t.Error("type in $defs should be preserved")
	}
}

func TestFilterSchema_BooleanSchema(t *testing.T) {
	// Boolean schemas should pass through unchanged
	schemaTrue := json.RawMessage(`true`)
	schemaFalse := json.RawMessage(`false`)

	vocab := map[string]bool{}

	resultTrue, err := FilterSchema(schemaTrue, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(resultTrue) != "true" {
		t.Errorf("true schema should be unchanged, got %s", resultTrue)
	}

	resultFalse, err := FilterSchema(schemaFalse, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(resultFalse) != "false" {
		t.Errorf("false schema should be unchanged, got %s", resultFalse)
	}
}

func TestFilterSchema_AllValidationKeywords(t *testing.T) {
	schema := json.RawMessage(`{
		"minLength": 1,
		"maxLength": 100,
		"pattern": "^[a-z]+$",
		"minimum": 0,
		"maximum": 100,
		"exclusiveMinimum": 0,
		"exclusiveMaximum": 100,
		"multipleOf": 5,
		"minItems": 1,
		"maxItems": 10,
		"uniqueItems": true,
		"minContains": 1,
		"maxContains": 5,
		"minProperties": 1,
		"maxProperties": 10,
		"required": ["name"],
		"dependentRequired": {"foo": ["bar"]},
		"enum": ["a", "b"],
		"const": "fixed"
	}`)

	vocab := map[string]bool{} // empty = nothing enabled

	result, err := FilterSchema(schema, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	// All validation keywords should be stripped
	if len(got) != 0 {
		t.Errorf("all validation keywords should be stripped, got %v", got)
	}
}

func TestFilterSchema_Draft2019Vocab(t *testing.T) {
	schema := json.RawMessage(`{"type":"string","minLength":1}`)

	// Using 2019 vocabulary URIs
	vocab := map[string]bool{
		Validation2019: true,
	}

	result, err := FilterSchema(schema, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	// Should preserve minLength when 2019 validation vocab is enabled
	if got["minLength"] != float64(1) {
		t.Errorf("minLength should be preserved with 2019 vocab, got %v", got["minLength"])
	}
}

func TestFilterSchema_ArrayItems(t *testing.T) {
	schema := json.RawMessage(`{
		"type": "array",
		"items": {
			"type": "string",
			"minLength": 1
		},
		"minItems": 1
	}`)

	vocab := map[string]bool{
		Applicator2020: true,
	}

	result, err := FilterSchema(schema, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	// items should be preserved (applicator)
	items, ok := got["items"].(map[string]any)
	if !ok {
		t.Fatal("items should be preserved")
	}

	// nested minLength should be stripped
	if _, ok := items["minLength"]; ok {
		t.Error("nested minLength should be stripped")
	}

	// minItems should be stripped (validation)
	if _, ok := got["minItems"]; ok {
		t.Error("minItems should be stripped")
	}
}

func TestFilterSchema_AllOf(t *testing.T) {
	schema := json.RawMessage(`{
		"allOf": [
			{"type": "string", "minLength": 1},
			{"maxLength": 100}
		]
	}`)

	vocab := map[string]bool{
		Applicator2020: true,
	}

	result, err := FilterSchema(schema, vocab)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	// allOf should be preserved
	allOf, ok := got["allOf"].([]any)
	if !ok {
		t.Fatal("allOf should be preserved")
	}

	// nested validation keywords should be stripped
	first := allOf[0].(map[string]any)
	if _, ok := first["minLength"]; ok {
		t.Error("minLength in allOf should be stripped")
	}

	second := allOf[1].(map[string]any)
	if _, ok := second["maxLength"]; ok {
		t.Error("maxLength in allOf should be stripped")
	}
}
