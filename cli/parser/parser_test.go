package parser

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestParseConfigFile(t *testing.T) {
	// Create a temp directory with a test config file
	tmpDir := t.TempDir()

	// Create a valid xschema config file
	configContent := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{
				"id": "User",
				"sourceType": "url",
				"source": "https://example.com/user.json",
				"adapter": "@xschema/zod"
			},
			{
				"id": "Post",
				"sourceType": "file",
				"source": "./post.json",
				"adapter": "@xschema/zod"
			}
		]
	}`

	configPath := filepath.Join(tmpDir, "user.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	// Parse the config file
	config, err := parseConfigFile(configPath)
	if err != nil {
		t.Fatalf("parseConfigFile: %v", err)
	}

	// Assertions
	if config == nil {
		t.Fatal("expected config, got nil")
	}
	if config.Namespace != "user" {
		t.Errorf("expected namespace 'user', got %q", config.Namespace)
	}
	if config.Language.Name != "typescript" {
		t.Errorf("expected language 'typescript', got %q", config.Language.Name)
	}
	if len(config.Schemas) != 2 {
		t.Errorf("expected 2 schemas, got %d", len(config.Schemas))
	}

	// Check first schema
	if config.Schemas[0].ID != "User" {
		t.Errorf("expected first schema ID 'User', got %q", config.Schemas[0].ID)
	}
	if config.Schemas[0].SourceType != SourceURL {
		t.Errorf("expected sourceType 'url', got %q", config.Schemas[0].SourceType)
	}
}

func TestParseConfigFileWithNamespaceOverride(t *testing.T) {
	tmpDir := t.TempDir()

	configContent := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"namespace": "custom",
		"schemas": [
			{
				"id": "Test",
				"sourceType": "url",
				"source": "https://example.com/test.json",
				"adapter": "@xschema/zod"
			}
		]
	}`

	configPath := filepath.Join(tmpDir, "user.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	config, err := parseConfigFile(configPath)
	if err != nil {
		t.Fatalf("parseConfigFile: %v", err)
	}

	if config.Namespace != "custom" {
		t.Errorf("expected namespace 'custom' (override), got %q", config.Namespace)
	}
}

func TestParseConfigFileNotXSchema(t *testing.T) {
	tmpDir := t.TempDir()

	// A regular JSON Schema file, not an xschema config
	configContent := `{
		"$schema": "https://json-schema.org/draft-07/schema#",
		"type": "object"
	}`

	configPath := filepath.Join(tmpDir, "regular.json")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	config, err := parseConfigFile(configPath)
	if err != nil {
		t.Fatalf("parseConfigFile: %v", err)
	}

	// Should return nil for non-xschema config
	if config != nil {
		t.Error("expected nil for non-xschema config file")
	}
}

func TestParseConfigFileWithJSONC(t *testing.T) {
	tmpDir := t.TempDir()

	// JSONC with comments
	configContent := `{
		// This is a comment
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{
				"id": "Test",
				"sourceType": "url",
				"source": "https://example.com/test.json",
				"adapter": "@xschema/zod"
			}
		]
	}`

	configPath := filepath.Join(tmpDir, "test.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	config, err := parseConfigFile(configPath)
	if err != nil {
		t.Fatalf("parseConfigFile: %v", err)
	}

	if config == nil {
		t.Fatal("expected config, got nil")
	}
	if len(config.Schemas) != 1 {
		t.Errorf("expected 1 schema, got %d", len(config.Schemas))
	}
}

func TestParse(t *testing.T) {
	tmpDir := t.TempDir()

	// Create two config files
	config1 := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "@xschema/zod"}
		]
	}`
	config2 := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "Post", "sourceType": "url", "source": "https://example.com/post.json", "adapter": "@xschema/zod"}
		]
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "user.jsonc"), []byte(config1), 0644); err != nil {
		t.Fatalf("failed to write config1: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "post.jsonc"), []byte(config2), 0644); err != nil {
		t.Fatalf("failed to write config2: %v", err)
	}

	ctx := context.Background()
	result, err := Parse(ctx, tmpDir, "")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}

	if result.Language.Name != "typescript" {
		t.Errorf("expected language 'typescript', got %q", result.Language.Name)
	}
	if len(result.Configs) != 2 {
		t.Errorf("expected 2 configs, got %d", len(result.Configs))
	}
	if len(result.Declarations) != 2 {
		t.Errorf("expected 2 declarations, got %d", len(result.Declarations))
	}
}

func TestParseDuplicateIDError(t *testing.T) {
	tmpDir := t.TempDir()

	// Two config files with same namespace (same filename) but shouldn't happen
	// Actually - two different files with same ID in same namespace
	config1 := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"namespace": "shared",
		"schemas": [
			{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "@xschema/zod"}
		]
	}`
	config2 := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"namespace": "shared",
		"schemas": [
			{"id": "User", "sourceType": "url", "source": "https://example.com/user2.json", "adapter": "@xschema/zod"}
		]
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "a.jsonc"), []byte(config1), 0644); err != nil {
		t.Fatalf("failed to write config1: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "b.jsonc"), []byte(config2), 0644); err != nil {
		t.Fatalf("failed to write config2: %v", err)
	}

	ctx := context.Background()
	_, err := Parse(ctx, tmpDir, "")
	if err == nil {
		t.Error("expected error for duplicate ID in same namespace")
	}
}

func TestParseMultipleLanguagesError(t *testing.T) {
	tmpDir := t.TempDir()

	// One TypeScript, one Python config
	tsConfig := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "@xschema/zod"}
		]
	}`
	pyConfig := `{
		"$schema": "https://xschema.dev/schemas/py.jsonc",
		"schemas": [
			{"id": "Post", "sourceType": "url", "source": "https://example.com/post.json", "adapter": "xschema-pydantic"}
		]
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "ts.jsonc"), []byte(tsConfig), 0644); err != nil {
		t.Fatalf("failed to write ts config: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "py.jsonc"), []byte(pyConfig), 0644); err != nil {
		t.Fatalf("failed to write py config: %v", err)
	}

	ctx := context.Background()
	_, err := Parse(ctx, tmpDir, "")
	if err == nil {
		t.Error("expected error for multiple languages without --lang filter")
	}
}

func TestParseWithLanguageFilter(t *testing.T) {
	tmpDir := t.TempDir()

	// One TypeScript, one Python config
	tsConfig := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "@xschema/zod"}
		]
	}`
	pyConfig := `{
		"$schema": "https://xschema.dev/schemas/py.jsonc",
		"schemas": [
			{"id": "Post", "sourceType": "url", "source": "https://example.com/post.json", "adapter": "xschema-pydantic"}
		]
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "ts.jsonc"), []byte(tsConfig), 0644); err != nil {
		t.Fatalf("failed to write ts config: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "py.jsonc"), []byte(pyConfig), 0644); err != nil {
		t.Fatalf("failed to write py config: %v", err)
	}

	ctx := context.Background()
	result, err := Parse(ctx, tmpDir, "typescript")
	if err != nil {
		t.Fatalf("Parse with filter: %v", err)
	}

	if result.Language.Name != "typescript" {
		t.Errorf("expected language 'typescript', got %q", result.Language.Name)
	}
	if len(result.Declarations) != 1 {
		t.Errorf("expected 1 declaration (filtered), got %d", len(result.Declarations))
	}
}

func TestDeclarationKey(t *testing.T) {
	d := Declaration{
		Namespace: "user",
		ID:        "TestUrl",
	}

	if d.Key() != "user:TestUrl" {
		t.Errorf("expected key 'user:TestUrl', got %q", d.Key())
	}
}
