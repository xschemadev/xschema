package generator

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/xschema/cli/retriever"
)

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

	if o.VarName() != "user_TestUrl" {
		t.Errorf("expected varName 'user_TestUrl', got %q", o.VarName())
	}
}
