package generator

import (
	"context"
	"encoding/json"
	"os"
	"testing"
)

func TestGenerateTypescript(t *testing.T) {
	// Change to testdata dir where node_modules exists
	originalDir, _ := os.Getwd()
	os.Chdir("testdata/typescript")
	defer os.Chdir(originalDir)

	input := GenerateBatchInput{
		Schemas: []GenerateInput{
			{Name: "User", Schema: json.RawMessage(`{"type": "string"}`)},
			{Name: "Post", Schema: json.RawMessage(`{"type": "number"}`)},
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
	if outputs[0].Name != "User" {
		t.Errorf("expected name 'User', got '%s'", outputs[0].Name)
	}
	if outputs[0].Schema != "z.string()" {
		t.Errorf("expected schema 'z.string()', got '%s'", outputs[0].Schema)
	}
	if outputs[0].Type != "z.infer<typeof User>" {
		t.Errorf("unexpected type: %s", outputs[0].Type)
	}

	// Check second output
	if outputs[1].Name != "Post" {
		t.Errorf("expected name 'Post', got '%s'", outputs[1].Name)
	}
	if outputs[1].Schema != "z.number()" {
		t.Errorf("expected schema 'z.number()', got '%s'", outputs[1].Schema)
	}
	if outputs[1].Type != "z.infer<typeof Post>" {
		t.Errorf("unexpected type: %s", outputs[1].Type)
	}
}
