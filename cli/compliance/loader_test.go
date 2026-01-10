package compliance

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadTestSuite(t *testing.T) {
	// Create test suite structure
	tmpDir := t.TempDir()
	testsDir := filepath.Join(tmpDir, "tests", "draft2020-12")
	if err := os.MkdirAll(testsDir, 0755); err != nil {
		t.Fatalf("failed to create tests dir: %v", err)
	}

	// Create test files
	typeJSON := `[
		{
			"description": "integer type matches integers",
			"schema": {"type": "integer"},
			"tests": [
				{"description": "an integer is an integer", "data": 1, "valid": true},
				{"description": "a string is not an integer", "data": "foo", "valid": false}
			]
		}
	]`
	if err := os.WriteFile(filepath.Join(testsDir, "type.json"), []byte(typeJSON), 0644); err != nil {
		t.Fatalf("failed to write type.json: %v", err)
	}

	constJSON := `[
		{
			"description": "const validation",
			"schema": {"const": 2},
			"tests": [
				{"description": "same value is valid", "data": 2, "valid": true}
			]
		}
	]`
	if err := os.WriteFile(filepath.Join(testsDir, "const.json"), []byte(constJSON), 0644); err != nil {
		t.Fatalf("failed to write const.json: %v", err)
	}

	// Load suite
	suite, err := LoadTestSuite(tmpDir, "draft2020-12", "")
	if err != nil {
		t.Fatalf("LoadTestSuite() error = %v", err)
	}

	// Check type keyword
	if _, ok := suite["type"]; !ok {
		t.Error("suite missing 'type' keyword")
	}
	if len(suite["type"]) != 1 {
		t.Errorf("suite['type'] has %d groups, want 1", len(suite["type"]))
	}
	if len(suite["type"][0].Tests) != 2 {
		t.Errorf("suite['type'][0].Tests has %d tests, want 2", len(suite["type"][0].Tests))
	}

	// Check const keyword
	if _, ok := suite["const"]; !ok {
		t.Error("suite missing 'const' keyword")
	}
}

func TestLoadTestSuite_WithKeyword(t *testing.T) {
	tmpDir := t.TempDir()
	testsDir := filepath.Join(tmpDir, "tests", "draft2020-12")
	if err := os.MkdirAll(testsDir, 0755); err != nil {
		t.Fatalf("failed to create tests dir: %v", err)
	}

	validJSON := `[{"description": "test", "schema": {}, "tests": []}]`
	if err := os.WriteFile(filepath.Join(testsDir, "type.json"), []byte(validJSON), 0644); err != nil {
		t.Fatalf("failed to write type.json: %v", err)
	}
	if err := os.WriteFile(filepath.Join(testsDir, "const.json"), []byte("not json"), 0644); err != nil {
		t.Fatalf("failed to write const.json: %v", err)
	}

	suite, err := LoadTestSuite(tmpDir, "draft2020-12", "type")
	if err != nil {
		t.Fatalf("LoadTestSuite() error = %v", err)
	}
	if len(suite) != 1 {
		t.Errorf("suite has %d keywords, want 1", len(suite))
	}
	if _, ok := suite["type"]; !ok {
		t.Error("suite missing 'type' keyword")
	}
}

func TestLoadTestSuite_MissingKeyword(t *testing.T) {
	tmpDir := t.TempDir()
	testsDir := filepath.Join(tmpDir, "tests", "draft2020-12")
	if err := os.MkdirAll(testsDir, 0755); err != nil {
		t.Fatalf("failed to create tests dir: %v", err)
	}

	validJSON := `[{"description": "test", "schema": {}, "tests": []}]`
	if err := os.WriteFile(filepath.Join(testsDir, "type.json"), []byte(validJSON), 0644); err != nil {
		t.Fatalf("failed to write type.json: %v", err)
	}

	_, err := LoadTestSuite(tmpDir, "draft2020-12", "missing")
	if err == nil {
		t.Fatal("LoadTestSuite() expected error for missing keyword, got nil")
	}
	if !strings.Contains(err.Error(), "available") || !strings.Contains(err.Error(), "type") {
		t.Errorf("LoadTestSuite() error = %v, want available keywords", err)
	}
}

func TestLoadTestSuite_NotFound(t *testing.T) {
	tmpDir := t.TempDir()

	_, err := LoadTestSuite(tmpDir, "draft2020-12", "")
	if err == nil {
		t.Error("LoadTestSuite() expected error for missing dir, got nil")
	}
}

func TestLoadTestSuite_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	testsDir := filepath.Join(tmpDir, "tests", "draft2020-12")
	if err := os.MkdirAll(testsDir, 0755); err != nil {
		t.Fatalf("failed to create tests dir: %v", err)
	}

	// Write invalid JSON
	if err := os.WriteFile(filepath.Join(testsDir, "invalid.json"), []byte("not json"), 0644); err != nil {
		t.Fatalf("failed to write invalid.json: %v", err)
	}

	_, err := LoadTestSuite(tmpDir, "draft2020-12", "")
	if err == nil {
		t.Error("LoadTestSuite() expected error for invalid JSON, got nil")
	}
}

func TestLoadTestSuite_SkipsDirectories(t *testing.T) {
	tmpDir := t.TempDir()
	testsDir := filepath.Join(tmpDir, "tests", "draft2020-12")
	if err := os.MkdirAll(testsDir, 0755); err != nil {
		t.Fatalf("failed to create tests dir: %v", err)
	}

	// Create a subdirectory (should be skipped)
	if err := os.MkdirAll(filepath.Join(testsDir, "optional"), 0755); err != nil {
		t.Fatalf("failed to create optional dir: %v", err)
	}

	// Create a valid test file
	validJSON := `[{"description": "test", "schema": {}, "tests": []}]`
	if err := os.WriteFile(filepath.Join(testsDir, "type.json"), []byte(validJSON), 0644); err != nil {
		t.Fatalf("failed to write type.json: %v", err)
	}

	suite, err := LoadTestSuite(tmpDir, "draft2020-12", "")
	if err != nil {
		t.Fatalf("LoadTestSuite() error = %v", err)
	}

	// Should only have "type", not "optional"
	if len(suite) != 1 {
		t.Errorf("suite has %d keywords, want 1", len(suite))
	}
	if _, ok := suite["type"]; !ok {
		t.Error("suite missing 'type' keyword")
	}
}

func TestLoadTestSuite_SkipsNonJSON(t *testing.T) {
	tmpDir := t.TempDir()
	testsDir := filepath.Join(tmpDir, "tests", "draft2020-12")
	if err := os.MkdirAll(testsDir, 0755); err != nil {
		t.Fatalf("failed to create tests dir: %v", err)
	}

	// Create non-JSON file (should be skipped)
	if err := os.WriteFile(filepath.Join(testsDir, "readme.txt"), []byte("hello"), 0644); err != nil {
		t.Fatalf("failed to write readme.txt: %v", err)
	}

	// Create valid test file
	validJSON := `[{"description": "test", "schema": {}, "tests": []}]`
	if err := os.WriteFile(filepath.Join(testsDir, "type.json"), []byte(validJSON), 0644); err != nil {
		t.Fatalf("failed to write type.json: %v", err)
	}

	suite, err := LoadTestSuite(tmpDir, "draft2020-12", "")
	if err != nil {
		t.Fatalf("LoadTestSuite() error = %v", err)
	}

	if len(suite) != 1 {
		t.Errorf("suite has %d keywords, want 1", len(suite))
	}
}

func TestLoadTestFile(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name       string
		content    string
		wantGroups int
		wantErr    bool
	}{
		{
			name:       "valid single group",
			content:    `[{"description": "test", "schema": {}, "tests": []}]`,
			wantGroups: 1,
			wantErr:    false,
		},
		{
			name:       "valid multiple groups",
			content:    `[{"description": "a", "schema": {}, "tests": []}, {"description": "b", "schema": {}, "tests": []}]`,
			wantGroups: 2,
			wantErr:    false,
		},
		{
			name:       "empty array",
			content:    `[]`,
			wantGroups: 0,
			wantErr:    false,
		},
		{
			name:    "invalid JSON",
			content: `not json`,
			wantErr: true,
		},
		{
			name:    "wrong JSON structure",
			content: `{"not": "an array"}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := filepath.Join(tmpDir, tt.name+".json")
			if err := os.WriteFile(filePath, []byte(tt.content), 0644); err != nil {
				t.Fatalf("failed to write test file: %v", err)
			}

			groups, err := loadTestFile(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadTestFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(groups) != tt.wantGroups {
				t.Errorf("loadTestFile() returned %d groups, want %d", len(groups), tt.wantGroups)
			}
		})
	}
}

func TestLoadTestFile_BooleanSchema(t *testing.T) {
	tmpDir := t.TempDir()

	// JSON Schema allows boolean schemas (true = accept all, false = reject all)
	content := `[
		{
			"description": "boolean schema true",
			"schema": true,
			"tests": [
				{"description": "any value is valid", "data": {}, "valid": true}
			]
		},
		{
			"description": "boolean schema false",
			"schema": false,
			"tests": [
				{"description": "any value is invalid", "data": {}, "valid": false}
			]
		}
	]`

	filePath := filepath.Join(tmpDir, "boolean.json")
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	groups, err := loadTestFile(filePath)
	if err != nil {
		t.Fatalf("loadTestFile() error = %v", err)
	}

	if len(groups) != 2 {
		t.Errorf("loadTestFile() returned %d groups, want 2", len(groups))
	}

	// Verify boolean schemas are preserved
	schema0, _ := json.Marshal(groups[0].Schema.Value())
	if string(schema0) != "true" {
		t.Errorf("first schema = %s, want true", string(schema0))
	}

	schema1, _ := json.Marshal(groups[1].Schema.Value())
	if string(schema1) != "false" {
		t.Errorf("second schema = %s, want false", string(schema1))
	}
}

func TestDrafts(t *testing.T) {
	// Ensure Drafts contains expected versions (v1 excluded until spec released)
	expected := []string{"draft2020-12", "draft2019-09", "draft7", "draft6", "draft4", "draft3"}

	if len(Drafts) != len(expected) {
		t.Errorf("Drafts has %d entries, want %d", len(Drafts), len(expected))
	}

	for i, draft := range expected {
		if i < len(Drafts) && Drafts[i] != draft {
			t.Errorf("Drafts[%d] = %s, want %s", i, Drafts[i], draft)
		}
	}
}
