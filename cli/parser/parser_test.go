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
				"adapter": "zod"
			},
			{
				"id": "Post",
				"sourceType": "file",
				"source": "./post.json",
				"adapter": "zod"
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
				"adapter": "zod"
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
				"adapter": "zod"
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
			{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "zod"}
		]
	}`
	config2 := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "Post", "sourceType": "url", "source": "https://example.com/post.json", "adapter": "zod"}
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
			{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "zod"}
		]
	}`
	config2 := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"namespace": "shared",
		"schemas": [
			{"id": "User", "sourceType": "url", "source": "https://example.com/user2.json", "adapter": "zod"}
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
			{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "zod"}
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
			{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "zod"}
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

func TestParseConfigFileEmptySchemas(t *testing.T) {
	tmpDir := t.TempDir()

	configContent := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": []
	}`

	configPath := filepath.Join(tmpDir, "empty.jsonc")
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
	if len(config.Schemas) != 0 {
		t.Errorf("expected 0 schemas, got %d", len(config.Schemas))
	}
}

func TestParseConfigFileMalformedJSON(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name    string
		content string
	}{
		{"truncated", `{"$schema": "https://xschema.dev/schemas/ts.jsonc"`},
		{"invalid syntax", `{"$schema": "https://xschema.dev/schemas/ts.jsonc", schemas: []}`},
		{"empty file", ``},
		{"not json", `this is not json`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configPath := filepath.Join(tmpDir, tt.name+".jsonc")
			if err := os.WriteFile(configPath, []byte(tt.content), 0644); err != nil {
				t.Fatalf("failed to write config: %v", err)
			}

			_, err := parseConfigFile(configPath)
			if err == nil {
				t.Error("expected error for malformed JSON")
			}
		})
	}
}

func TestParseConfigFileMissingSchemaURL(t *testing.T) {
	tmpDir := t.TempDir()

	// Valid JSON but no $schema field
	configContent := `{
		"schemas": [
			{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "zod"}
		]
	}`

	configPath := filepath.Join(tmpDir, "no-schema.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	config, err := parseConfigFile(configPath)
	if err != nil {
		t.Fatalf("parseConfigFile: %v", err)
	}

	// Should return nil (not an xschema config)
	if config != nil {
		t.Error("expected nil for config without $schema")
	}
}

func TestParseConfigFileUnknownSchemaLanguage(t *testing.T) {
	tmpDir := t.TempDir()

	// xschema.dev URL but unknown language
	configContent := `{
		"$schema": "https://xschema.dev/schemas/unknown.jsonc",
		"schemas": []
	}`

	configPath := filepath.Join(tmpDir, "unknown-lang.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	_, err := parseConfigFile(configPath)
	if err == nil {
		t.Error("expected error for unknown language in $schema")
	}
}

func TestParseNoConfigFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Empty directory - no config files
	ctx := context.Background()
	_, err := Parse(ctx, tmpDir, "")
	if err == nil {
		t.Error("expected error when no config files found")
	}
}

func TestParseConfigFileInSubdirectory(t *testing.T) {
	tmpDir := t.TempDir()

	// Create subdirectory with config
	subDir := filepath.Join(tmpDir, "schemas", "user")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("failed to create subdir: %v", err)
	}

	configContent := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "zod"}
		]
	}`

	configPath := filepath.Join(subDir, "user.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	ctx := context.Background()
	result, err := Parse(ctx, tmpDir, "")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}

	if len(result.Configs) != 1 {
		t.Errorf("expected 1 config from subdirectory, got %d", len(result.Configs))
	}
}

func TestParseConfigFileWithAllSourceTypes(t *testing.T) {
	tmpDir := t.TempDir()

	configContent := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [
			{"id": "FromURL", "sourceType": "url", "source": "https://example.com/schema.json", "adapter": "zod"},
			{"id": "FromFile", "sourceType": "file", "source": "./local.json", "adapter": "zod"},
			{"id": "Inline", "sourceType": "json", "source": {"type": "string"}, "adapter": "zod"}
		]
	}`

	configPath := filepath.Join(tmpDir, "all-sources.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	config, err := parseConfigFile(configPath)
	if err != nil {
		t.Fatalf("parseConfigFile: %v", err)
	}

	if len(config.Schemas) != 3 {
		t.Fatalf("expected 3 schemas, got %d", len(config.Schemas))
	}

	if config.Schemas[0].SourceType != SourceURL {
		t.Errorf("expected first schema sourceType 'url', got %q", config.Schemas[0].SourceType)
	}
	if config.Schemas[1].SourceType != SourceFile {
		t.Errorf("expected second schema sourceType 'file', got %q", config.Schemas[1].SourceType)
	}
	if config.Schemas[2].SourceType != SourceJSON {
		t.Errorf("expected third schema sourceType 'json', got %q", config.Schemas[2].SourceType)
	}
}

func TestParseResultDeclarationsByNamespace(t *testing.T) {
	result := &ParseResult{
		Declarations: []Declaration{
			{Namespace: "user", ID: "User"},
			{Namespace: "user", ID: "Profile"},
			{Namespace: "post", ID: "Post"},
		},
	}

	byNs := result.DeclarationsByNamespace()

	if len(byNs) != 2 {
		t.Errorf("expected 2 namespaces, got %d", len(byNs))
	}
	if len(byNs["user"]) != 2 {
		t.Errorf("expected 2 user declarations, got %d", len(byNs["user"]))
	}
	if len(byNs["post"]) != 1 {
		t.Errorf("expected 1 post declaration, got %d", len(byNs["post"]))
	}
}

func TestParseResultDeclarationsByAdapter(t *testing.T) {
	result := &ParseResult{
		Declarations: []Declaration{
			{Namespace: "user", ID: "User", Adapter: "zod"},
			{Namespace: "user", ID: "Profile", Adapter: "zod"},
			{Namespace: "post", ID: "Post", Adapter: "@xschema/pydantic"},
		},
	}

	byAdapter := result.DeclarationsByAdapter()

	if len(byAdapter) != 2 {
		t.Errorf("expected 2 adapters, got %d", len(byAdapter))
	}
	if len(byAdapter["zod"]) != 2 {
		t.Errorf("expected 2 zod declarations, got %d", len(byAdapter["zod"]))
	}
	if len(byAdapter["@xschema/pydantic"]) != 1 {
		t.Errorf("expected 1 pydantic declaration, got %d", len(byAdapter["@xschema/pydantic"]))
	}
}

func TestParseContextCancellation(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a config file
	configContent := `{
		"$schema": "https://xschema.dev/schemas/ts.jsonc",
		"schemas": [{"id": "User", "sourceType": "url", "source": "https://example.com/user.json", "adapter": "zod"}]
	}`
	if err := os.WriteFile(filepath.Join(tmpDir, "test.jsonc"), []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err := Parse(ctx, tmpDir, "")
	if err == nil {
		t.Error("expected context cancellation error")
	}
}

func TestParseConfigFileNonExistent(t *testing.T) {
	_, err := parseConfigFile("/nonexistent/path/config.jsonc")
	if err == nil {
		t.Error("expected error for non-existent file")
	}
}
