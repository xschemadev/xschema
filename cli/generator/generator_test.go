package generator

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/xschema/cli/retriever"
)

// Note: Tests that call Generate() use @xschema/zod as a real adapter to test
// the generator pipeline. They are NOT testing zod-specific behavior, just that
// the generator correctly invokes adapters and processes their output.

func TestGenerateTypescript(t *testing.T) {
	// Change to testdata dir where node_modules exists
	originalDir, _ := os.Getwd()
	os.Chdir("testdata/typescript")
	defer os.Chdir(originalDir)

	input := GenerateBatchInput{
		Schemas: []retriever.RetrievedSchema{
			{Namespace: "user", ID: "User", Schema: json.RawMessage(`{"type": "string"}`), Adapter: "@xschema/zod"},
			{Namespace: "user", ID: "Post", Schema: json.RawMessage(`{"type": "number"}`), Adapter: "@xschema/zod"},
		},
		Adapter:  "@xschema/zod",
		Language: "typescript",
	}

	outputs, err := Generate(context.Background(), input)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	if len(outputs) != 2 {
		t.Fatalf("expected 2 outputs, got %d", len(outputs))
	}

	// Check first output
	if outputs[0].ID != "User" {
		t.Errorf("expected ID 'User', got '%s'", outputs[0].ID)
	}
	if outputs[0].Namespace != "user" {
		t.Errorf("expected namespace 'user', got '%s'", outputs[0].Namespace)
	}
	if outputs[0].Schema != "z.string()" {
		t.Errorf("expected schema 'z.string()', got '%s'", outputs[0].Schema)
	}
	if outputs[0].Type != "z.infer<typeof user_User>" {
		t.Errorf("unexpected type: %s", outputs[0].Type)
	}

	// Check second output
	if outputs[1].ID != "Post" {
		t.Errorf("expected ID 'Post', got '%s'", outputs[1].ID)
	}
	if outputs[1].Schema != "z.number()" {
		t.Errorf("expected schema 'z.number()', got '%s'", outputs[1].Schema)
	}
}

func TestGenerateOutputKey(t *testing.T) {
	o := GenerateOutput{
		Namespace: "user",
		ID:        "TestUrl",
	}

	if o.Key() != "user:TestUrl" {
		t.Errorf("expected key 'user:TestUrl', got %q", o.Key())
	}
}

func TestGenerateMultipleSchemas(t *testing.T) {
	originalDir, _ := os.Getwd()
	os.Chdir("testdata/typescript")
	defer os.Chdir(originalDir)

	input := GenerateBatchInput{
		Schemas: []retriever.RetrievedSchema{
			{Namespace: "ns1", ID: "Schema1", Schema: json.RawMessage(`{"type": "string"}`), Adapter: "@xschema/zod"},
			{Namespace: "ns1", ID: "Schema2", Schema: json.RawMessage(`{"type": "number"}`), Adapter: "@xschema/zod"},
			{Namespace: "ns2", ID: "Schema3", Schema: json.RawMessage(`{"type": "boolean"}`), Adapter: "@xschema/zod"},
		},
		Adapter:  "@xschema/zod",
		Language: "typescript",
	}

	outputs, err := Generate(context.Background(), input)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	if len(outputs) != 3 {
		t.Fatalf("expected 3 outputs, got %d", len(outputs))
	}

	// Verify namespace preservation
	if outputs[0].Namespace != "ns1" || outputs[0].ID != "Schema1" {
		t.Errorf("first output mismatch: ns=%s id=%s", outputs[0].Namespace, outputs[0].ID)
	}
	if outputs[2].Namespace != "ns2" || outputs[2].ID != "Schema3" {
		t.Errorf("third output mismatch: ns=%s id=%s", outputs[2].Namespace, outputs[2].ID)
	}
}

func TestGenerateComplexSchema(t *testing.T) {
	originalDir, _ := os.Getwd()
	os.Chdir("testdata/typescript")
	defer os.Chdir(originalDir)

	complexSchema := `{
		"type": "object",
		"properties": {
			"id": {"type": "string"},
			"name": {"type": "string"},
			"age": {"type": "number", "minimum": 0},
			"email": {"type": "string", "format": "email"},
			"tags": {"type": "array", "items": {"type": "string"}}
		},
		"required": ["id", "name"]
	}`

	input := GenerateBatchInput{
		Schemas: []retriever.RetrievedSchema{
			{Namespace: "user", ID: "ComplexUser", Schema: json.RawMessage(complexSchema), Adapter: "@xschema/zod"},
		},
		Adapter:  "@xschema/zod",
		Language: "typescript",
	}

	outputs, err := Generate(context.Background(), input)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	if len(outputs) != 1 {
		t.Fatalf("expected 1 output, got %d", len(outputs))
	}

	// Should contain z.object - just verifying adapter produced output
	if !strings.Contains(outputs[0].Schema, "z.object") {
		t.Errorf("expected schema to contain z.object, got: %s", outputs[0].Schema)
	}
}

func TestGenerateEmptySchemas(t *testing.T) {
	originalDir, _ := os.Getwd()
	os.Chdir("testdata/typescript")
	defer os.Chdir(originalDir)

	input := GenerateBatchInput{
		Schemas:  []retriever.RetrievedSchema{},
		Adapter:  "@xschema/zod",
		Language: "typescript",
	}

	outputs, err := Generate(context.Background(), input)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	if len(outputs) != 0 {
		t.Errorf("expected 0 outputs for empty input, got %d", len(outputs))
	}
}

func TestGenerateAdapterNotFound(t *testing.T) {
	originalDir, _ := os.Getwd()
	os.Chdir("testdata/typescript")
	defer os.Chdir(originalDir)

	input := GenerateBatchInput{
		Schemas: []retriever.RetrievedSchema{
			{Namespace: "test", ID: "Test", Schema: json.RawMessage(`{"type": "string"}`), Adapter: "@xschema/nonexistent-adapter"},
		},
		Adapter:  "@xschema/nonexistent-adapter",
		Language: "typescript",
	}

	_, err := Generate(context.Background(), input)
	if err == nil {
		t.Error("expected error for non-existent adapter")
	}
}

func TestGenerateUnsupportedLanguage(t *testing.T) {
	input := GenerateBatchInput{
		Schemas: []retriever.RetrievedSchema{
			{Namespace: "test", ID: "Test", Schema: json.RawMessage(`{"type": "string"}`), Adapter: "@xschema/zod"},
		},
		Adapter:  "@xschema/zod",
		Language: "unsupported-lang",
	}

	_, err := Generate(context.Background(), input)
	if err == nil {
		t.Error("expected error for unsupported language")
	}
}

func TestGenerateContextCancellation(t *testing.T) {
	originalDir, _ := os.Getwd()
	os.Chdir("testdata/typescript")
	defer os.Chdir(originalDir)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	input := GenerateBatchInput{
		Schemas: []retriever.RetrievedSchema{
			{Namespace: "test", ID: "Test", Schema: json.RawMessage(`{"type": "string"}`), Adapter: "@xschema/zod"},
		},
		Adapter:  "@xschema/zod",
		Language: "typescript",
	}

	_, err := Generate(ctx, input)
	if err == nil {
		t.Error("expected context cancellation error")
	}
}

func TestGenerateAllGroupsByAdapter(t *testing.T) {
	originalDir, _ := os.Getwd()
	os.Chdir("testdata/typescript")
	defer os.Chdir(originalDir)

	schemas := []retriever.RetrievedSchema{
		{Namespace: "user", ID: "User", Schema: json.RawMessage(`{"type": "string"}`), Adapter: "@xschema/zod"},
		{Namespace: "user", ID: "Post", Schema: json.RawMessage(`{"type": "number"}`), Adapter: "@xschema/zod"},
	}

	outputs, err := GenerateAll(context.Background(), schemas, "typescript")
	if err != nil {
		t.Fatalf("GenerateAll failed: %v", err)
	}

	if len(outputs) != 2 {
		t.Errorf("expected 2 outputs, got %d", len(outputs))
	}
}

func TestGenerateAllEmptySchemas(t *testing.T) {
	outputs, err := GenerateAll(context.Background(), []retriever.RetrievedSchema{}, "typescript")
	if err != nil {
		t.Fatalf("GenerateAll failed: %v", err)
	}

	if len(outputs) != 0 {
		t.Errorf("expected 0 outputs, got %d", len(outputs))
	}
}

func TestGenerateInputJSON(t *testing.T) {
	input := GenerateInput{
		Namespace: "user",
		ID:        "Test",
		Schema:    json.RawMessage(`{"type": "string"}`),
	}

	data, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded GenerateInput
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.Namespace != "user" || decoded.ID != "Test" {
		t.Errorf("round-trip failed: %+v", decoded)
	}
}
