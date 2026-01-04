package generator

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/xschemadev/xschema/retriever"
)

func TestGenerateOutputKey(t *testing.T) {
	o := GenerateOutput{
		Namespace: "user",
		ID:        "TestUrl",
	}

	if o.Key() != "user:TestUrl" {
		t.Errorf("expected key 'user:TestUrl', got %q", o.Key())
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
			{Namespace: "test", ID: "Test", Schema: json.RawMessage(`{"type": "string"}`), Adapter: "zod"},
		},
		Adapter:  "zod",
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
			{Namespace: "test", ID: "Test", Schema: json.RawMessage(`{"type": "string"}`), Adapter: "zod"},
		},
		Adapter:  "zod",
		Language: "typescript",
	}

	_, err := Generate(ctx, input)
	if err == nil {
		t.Error("expected context cancellation error")
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
