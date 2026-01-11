package generator

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/xschemadev/xschema/adapter"
	_ "github.com/xschemadev/xschema/language/langs"
	"github.com/xschemadev/xschema/retriever"
)

func TestConvertResultKey(t *testing.T) {
	o := adapter.ConvertResult{
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

	adapterRef := "@xschemadev/nonexistent-adapter"
	input := GenerateBatchInput{
		Schemas: []retriever.RetrievedSchema{
			{Namespace: "test", ID: "Test", Schema: json.RawMessage(`{"type": "string"}`), Adapter: adapterRef},
		},
		AdapterRef:  adapterRef,
		Language:    "typescript",
		ProjectRoot: ".",
	}

	_, err := Generate(context.Background(), input)
	if err == nil {
		t.Error("expected error for non-existent adapter")
	}
}

func TestGenerateLegacyAdapterRefMigrationGuidance(t *testing.T) {
	input := GenerateBatchInput{
		Schemas: []retriever.RetrievedSchema{
			{Namespace: "test", ID: "Test", Schema: json.RawMessage(`{"type": "string"}`), Adapter: "zod"},
		},
		AdapterRef:  "zod",
		Language:    "typescript",
		ProjectRoot: ".",
	}

	_, err := Generate(context.Background(), input)
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "migration") || !strings.Contains(err.Error(), "@xschemadev/zod") {
		t.Fatalf("expected migration guidance for legacy adapter ref, got: %v", err)
	}
}

func TestGenerateUnsupportedLanguage(t *testing.T) {
	input := GenerateBatchInput{
		Schemas: []retriever.RetrievedSchema{
			{Namespace: "test", ID: "Test", Schema: json.RawMessage(`{"type": "string"}`), Adapter: "zod"},
		},
		AdapterRef:  "zod",
		Language:    "unsupported-lang",
		ProjectRoot: ".",
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

	adapterRef := "@xschemadev/zod"
	input := GenerateBatchInput{
		Schemas: []retriever.RetrievedSchema{
			{Namespace: "test", ID: "Test", Schema: json.RawMessage(`{"type": "string"}`), Adapter: adapterRef},
		},
		AdapterRef:  adapterRef,
		Language:    "typescript",
		ProjectRoot: ".",
	}

	_, err := Generate(ctx, input)
	if err == nil {
		t.Error("expected context cancellation error")
	}
}

func TestGenerateAllEmptySchemas(t *testing.T) {
	outputs, err := GenerateAll(context.Background(), []retriever.RetrievedSchema{}, "typescript", ".")
	if err != nil {
		t.Fatalf("GenerateAll failed: %v", err)
	}

	if len(outputs) != 0 {
		t.Errorf("expected 0 outputs, got %d", len(outputs))
	}
}

func TestConvertInputJSON(t *testing.T) {
	input := adapter.ConvertInput{
		Namespace: "user",
		ID:        "Test",
		Schema:    json.RawMessage(`{"type": "string"}`),
	}

	data, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded adapter.ConvertInput
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.Namespace != "user" || decoded.ID != "Test" {
		t.Errorf("round-trip failed: %+v", decoded)
	}
}

func TestValidateOutputs(t *testing.T) {
	tests := []struct {
		name    string
		outputs []adapter.ConvertResult
		wantErr bool
	}{
		{
			name:    "empty outputs",
			outputs: []adapter.ConvertResult{},
			wantErr: false,
		},
		{
			name: "valid with schema only",
			outputs: []adapter.ConvertResult{
				{Namespace: "test", ID: "Test", Schema: "z.string()", Type: ""},
			},
			wantErr: false,
		},
		{
			name: "valid with type only",
			outputs: []adapter.ConvertResult{
				{Namespace: "test", ID: "Test", Schema: "", Type: "string"},
			},
			wantErr: false,
		},
		{
			name: "valid with both",
			outputs: []adapter.ConvertResult{
				{Namespace: "test", ID: "Test", Schema: "z.string()", Type: "string"},
			},
			wantErr: false,
		},
		{
			name: "invalid with neither",
			outputs: []adapter.ConvertResult{
				{Namespace: "test", ID: "Test", Schema: "", Type: ""},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateOutputs(tt.outputs, "test-adapter")
			if (err != nil) != tt.wantErr {
				t.Errorf("validateOutputs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
