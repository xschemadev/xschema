package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/xschemadev/xschema/parser"
	"github.com/xschemadev/xschema/retriever"
)

// TestIntegration_MultipleConfigs tests parsing multiple config files
func TestIntegration_MultipleConfigs(t *testing.T) {
	tmpDir := t.TempDir()

	// Create two config files with different namespaces
	config1 := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "User", "sourceType": "json", "source": {"type": "string"}, "adapter": "zod"}
		]
	}`
	config2 := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "Post", "sourceType": "json", "source": {"type": "number"}, "adapter": "zod"}
		]
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "users.jsonc"), []byte(config1), 0644); err != nil {
		t.Fatalf("failed to write config1: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "posts.jsonc"), []byte(config2), 0644); err != nil {
		t.Fatalf("failed to write config2: %v", err)
	}

	ctx := context.Background()
	result, err := parser.Parse(ctx, tmpDir, "")
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// Should have 2 configs with different namespaces (derived from filenames)
	if len(result.Configs) != 2 {
		t.Errorf("expected 2 configs, got %d", len(result.Configs))
	}

	byNs := result.DeclarationsByNamespace()
	if len(byNs) != 2 {
		t.Errorf("expected 2 namespaces, got %d", len(byNs))
	}
	if _, ok := byNs["users"]; !ok {
		t.Error("expected 'users' namespace")
	}
	if _, ok := byNs["posts"]; !ok {
		t.Error("expected 'posts' namespace")
	}
}

// TestIntegration_InlineJSONPassthrough tests inline JSON schemas
func TestIntegration_InlineJSONPassthrough(t *testing.T) {
	tmpDir := t.TempDir()

	inlineSchema := `{"type": "object", "properties": {"test": {"type": "boolean"}}}`
	config := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "Inline", "sourceType": "json", "source": ` + inlineSchema + `, "adapter": "zod"}
		]
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "test.jsonc"), []byte(config), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	ctx := context.Background()
	result, err := parser.Parse(ctx, tmpDir, "")
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	schemas, err := retriever.Retrieve(ctx, result.Declarations, retriever.DefaultOptions())
	if err != nil {
		t.Fatalf("Retrieve failed: %v", err)
	}

	// Inline schema should be passed through as-is
	if string(schemas[0].Schema) != inlineSchema {
		t.Errorf("inline schema not passed through correctly: %s", schemas[0].Schema)
	}
}

// TestIntegration_AdapterGrouping tests that schemas are grouped by adapter
func TestIntegration_AdapterGrouping(t *testing.T) {
	tmpDir := t.TempDir()

	config := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "Zod1", "sourceType": "json", "source": {"type": "string"}, "adapter": "zod"},
			{"id": "Zod2", "sourceType": "json", "source": {"type": "number"}, "adapter": "zod"},
			{"id": "Other1", "sourceType": "json", "source": {"type": "boolean"}, "adapter": "@xschema/other"},
			{"id": "Other2", "sourceType": "json", "source": {"type": "null"}, "adapter": "@xschema/other"}
		]
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "test.jsonc"), []byte(config), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	ctx := context.Background()
	result, err := parser.Parse(ctx, tmpDir, "")
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	schemas, err := retriever.Retrieve(ctx, result.Declarations, retriever.DefaultOptions())
	if err != nil {
		t.Fatalf("Retrieve failed: %v", err)
	}

	groups := retriever.GroupByAdapter(schemas)
	if len(groups) != 2 {
		t.Errorf("expected 2 adapter groups, got %d", len(groups))
	}
	if len(groups["zod"]) != 2 {
		t.Errorf("expected 2 zod schemas, got %d", len(groups["zod"]))
	}
	if len(groups["@xschema/other"]) != 2 {
		t.Errorf("expected 2 other schemas, got %d", len(groups["@xschema/other"]))
	}

	// Sorted adapters should be deterministic
	sorted := retriever.SortedAdapters(groups)
	if sorted[0] != "@xschema/other" || sorted[1] != "zod" {
		t.Errorf("unexpected sort order: %v", sorted)
	}
}
