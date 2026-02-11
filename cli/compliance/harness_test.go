package compliance

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/xschemadev/xschema/adapter"
	"github.com/xschemadev/xschema/language"
	_ "github.com/xschemadev/xschema/language/langs"
)

func TestGenerateHarness(t *testing.T) {
	tmpDir := t.TempDir()

	items := []HarnessItem{
		{
			GroupID: "group_0",
			AdapterOutput: &adapter.ConvertResult{
				VarName:  "group_0",
				Schema:   `z.object({ name: z.string() })`,
				Imports:  []string{`import { z } from "zod"`},
				Validate: `(data) => schema.safeParse(data).success`,
			},
			Tests: []TestCase{
				{Description: "valid object", Data: map[string]any{"name": "test"}, Valid: true},
				{Description: "invalid object", Data: "not an object", Valid: false},
			},
		},
	}

	harnessPath, err := GenerateHarness(language.ByName("typescript"), items, tmpDir)
	if err != nil {
		t.Fatalf("GenerateHarness() error = %v", err)
	}
	defer os.Remove(harnessPath)

	// Verify file exists in target directory
	if !strings.HasPrefix(harnessPath, tmpDir) {
		t.Errorf("harness created in wrong directory: %s", harnessPath)
	}

	// Verify extension
	if !strings.HasSuffix(harnessPath, ".ts") {
		t.Errorf("harness has wrong extension: %s", harnessPath)
	}

	// Read and verify content
	content, err := os.ReadFile(harnessPath)
	if err != nil {
		t.Fatalf("failed to read harness: %v", err)
	}

	// Check schema was inserted
	if !strings.Contains(string(content), items[0].AdapterOutput.Schema) {
		t.Error("harness missing schema")
	}

	// Check import was inserted
	if !strings.Contains(string(content), `import { z } from "zod"`) {
		t.Error("harness missing import")
	}

	// Check validate function was inserted
	if !strings.Contains(string(content), items[0].AdapterOutput.Validate) {
		t.Error("harness missing validate function")
	}

	// Check test cases were inserted
	if !strings.Contains(string(content), "valid object") {
		t.Error("harness missing test case descriptions")
	}
}

func TestGenerateHarness_MultipleGroups(t *testing.T) {
	tmpDir := t.TempDir()

	items := []HarnessItem{
		{
			GroupID: "group_0",
			AdapterOutput: &adapter.ConvertResult{
				VarName:  "group_0",
				Schema:   `z.string()`,
				Imports:  []string{`import { z } from "zod"`},
				Validate: `(data) => schema.safeParse(data).success`,
			},
			Tests: []TestCase{{Description: "test 1", Data: "hello", Valid: true}},
		},
		{
			GroupID: "group_1",
			AdapterOutput: &adapter.ConvertResult{
				VarName:  "group_1",
				Schema:   `z.number()`,
				Imports:  []string{`import { z } from "zod"`},
				Validate: `(data) => schema.safeParse(data).success`,
			},
			Tests: []TestCase{{Description: "test 2", Data: 42, Valid: true}},
		},
	}

	harnessPath, err := GenerateHarness(language.ByName("typescript"), items, tmpDir)
	if err != nil {
		t.Fatalf("GenerateHarness() error = %v", err)
	}
	defer os.Remove(harnessPath)

	content, _ := os.ReadFile(harnessPath)
	contentStr := string(content)

	// Should have both schemas
	if !strings.Contains(contentStr, `z.string()`) {
		t.Error("harness missing first schema")
	}
	if !strings.Contains(contentStr, `z.number()`) {
		t.Error("harness missing second schema")
	}

	// Should have both group IDs
	if !strings.Contains(contentStr, `"group_0"`) {
		t.Error("harness missing group_0")
	}
	if !strings.Contains(contentStr, `"group_1"`) {
		t.Error("harness missing group_1")
	}
}

func TestGenerateHarness_MergedImports(t *testing.T) {
	tmpDir := t.TempDir()

	items := []HarnessItem{
		{
			GroupID: "group_0",
			AdapterOutput: &adapter.ConvertResult{
				VarName:         "group_0",
				Schema:          `z.object({ name: z.string() })`,
				Imports:         []string{`import { z } from "zod"`, `import { helper } from "helper"`},
				Validate:        `(data) => schema.safeParse(data).success`,
				ValidationImports: []string{`import { ZodError } from "zod"`, `import { helper } from "helper"`},
			},
			Tests: []TestCase{{Description: "valid", Data: map[string]any{"name": "test"}, Valid: true}},
		},
	}

	harnessPath, err := GenerateHarness(language.ByName("typescript"), items, tmpDir)
	if err != nil {
		t.Fatalf("GenerateHarness() error = %v", err)
	}
	defer os.Remove(harnessPath)

	content, err := os.ReadFile(harnessPath)
	if err != nil {
		t.Fatalf("failed to read harness: %v", err)
	}

	mergedImports := "import { helper } from \"helper\"\nimport { ZodError, z } from \"zod\""
	if !strings.Contains(string(content), mergedImports) {
		t.Errorf("harness missing merged imports: %s", mergedImports)
	}
}

func TestGenerateHarness_TypeOnly(t *testing.T) {
	tmpDir := t.TempDir()

	items := []HarnessItem{
		{
			GroupID: "group_0",
			AdapterOutput: &adapter.ConvertResult{
				VarName:  "group_0",
				Schema:   "",
				Type:     `{ name: string }`,
				Imports:  nil,
				Validate: "",
			},
			Tests: []TestCase{
				{Description: "test 1", Data: map[string]any{"name": "test"}, Valid: true},
			},
		},
	}

	harnessPath, err := GenerateHarness(language.ByName("typescript"), items, tmpDir)
	if err != nil {
		t.Fatalf("GenerateHarness() error = %v", err)
	}
	defer os.Remove(harnessPath)

	content, _ := os.ReadFile(harnessPath)
	contentStr := string(content)

	// Should have type declaration (type-only mode)
	if !strings.Contains(contentStr, "type GeneratedType") {
		t.Error("type-only harness should have type declaration")
	}

	if !strings.Contains(contentStr, items[0].AdapterOutput.Type) {
		t.Error("type-only harness should include type expression")
	}

	// Should mark results as skipped via isTypeOnly check
	if !strings.Contains(contentStr, `isTypeOnly: true`) {
		t.Error("type-only harness should have isTypeOnly: true")
	}
}

func TestGenerateHarness_MissingTemplate(t *testing.T) {
	tmpDir := t.TempDir()

	items := []HarnessItem{
		{
			GroupID: "group_0",
			AdapterOutput: &adapter.ConvertResult{
				VarName:  "group_0",
				Schema:   "some_schema",
				Validate: "validate_fn",
			},
			Tests: []TestCase{},
		},
	}

	// Language with no HarnessTemplate should fail
	lang := &language.Language{Name: "nolang", HarnessExtension: ".txt"}
	_, err := GenerateHarness(lang, items, tmpDir)
	if err == nil {
		t.Error("GenerateHarness() expected error for language without template")
	}
}

func TestGenerateHarness_EmptyItems(t *testing.T) {
	tmpDir := t.TempDir()

	_, err := GenerateHarness(language.ByName("typescript"), []HarnessItem{}, tmpDir)
	if err == nil {
		t.Error("GenerateHarness() expected error for empty items")
	}
}

func TestGenerateHarness_ComplexTestData(t *testing.T) {
	tmpDir := t.TempDir()

	items := []HarnessItem{
		{
			GroupID: "group_0",
			AdapterOutput: &adapter.ConvertResult{
				VarName:  "group_0",
				Schema:   `z.any()`,
				Imports:  []string{`import { z } from "zod"`},
				Validate: `(data) => true`,
			},
			Tests: []TestCase{
				{
					Description: "nested object",
					Data: map[string]any{
						"nested": map[string]any{
							"array": []any{1, 2, 3},
							"null":  nil,
							"bool":  true,
						},
					},
					Valid: true,
				},
				{
					Description: "special chars",
					Data:        "hello\nworld\t\"quoted\"",
					Valid:       false,
				},
			},
		},
	}

	harnessPath, err := GenerateHarness(language.ByName("typescript"), items, tmpDir)
	if err != nil {
		t.Fatalf("GenerateHarness() error = %v", err)
	}
	defer os.Remove(harnessPath)

	content, _ := os.ReadFile(harnessPath)
	contentStr := string(content)

	// Should contain JSON.parse for test data
	if !strings.Contains(contentStr, "JSON.parse(") {
		t.Error("harness should use JSON.parse for test data")
	}
}

func TestGenerateHarness_FilenamePattern(t *testing.T) {
	tmpDir := t.TempDir()

	items := []HarnessItem{
		{
			GroupID: "group_0",
			AdapterOutput: &adapter.ConvertResult{
				VarName:  "group_0",
				Schema:   `z.string()`,
				Imports:  []string{`import { z } from "zod"`},
				Validate: `(data) => schema.safeParse(data).success`,
			},
			Tests: []TestCase{},
		},
	}

	harnessPath, err := GenerateHarness(language.ByName("typescript"), items, tmpDir)
	if err != nil {
		t.Fatalf("GenerateHarness() error = %v", err)
	}
	defer os.Remove(harnessPath)

	// Filename should match pattern "xschema-harness-*"
	basename := filepath.Base(harnessPath)
	if !strings.HasPrefix(basename, "xschema-harness-") {
		t.Errorf("harness filename = %s, want prefix 'xschema-harness-'", basename)
	}
}

func TestTSHarnessTemplate_RuntimeValidation(t *testing.T) {
	lang := language.ByName("typescript")
	if lang == nil {
		t.Fatal("expected typescript language to be registered")
	}

	tmpl, err := template.New("harness").Parse(lang.HarnessTemplate)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	testData := []BatchTestData{
		{GroupID: "group_0", Tests: []TestCase{{Data: map[string]any{"name": "test"}, Valid: true}}},
	}
	testDataJSON, _ := json.Marshal(testData)
	testDataString, _ := json.Marshal(string(testDataJSON))

	data := HarnessTemplateData{
		Imports: `import { z } from "zod"`,
		Schemas: []HarnessSchemaEntry{
			{
				GroupID:    "group_0",
				Schema:     `z.object({ name: z.string() })`,
				Validate:   `(data) => schema.safeParse(data).success`,
				IsTypeOnly: false,
			},
		},
		TestDataString: string(testDataString),
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		t.Fatalf("failed to execute template: %v", err)
	}

	result := buf.String()

	// Should have import
	if !strings.Contains(result, `import { z } from "zod"`) {
		t.Error("missing zod import")
	}

	// Should have schema in IIFE
	if !strings.Contains(result, `const schema = z.object({ name: z.string() });`) {
		t.Error("missing schema declaration")
	}

	// Should have validate function
	if !strings.Contains(result, `validate: (data) => schema.safeParse(data).success`) {
		t.Error("missing validate function")
	}

	// Should have try-catch for runtime validation
	if !strings.Contains(result, "try {") {
		t.Error("missing try-catch block")
	}

	// Should output JSON
	if !strings.Contains(result, "console.log(JSON.stringify(results));") {
		t.Error("missing JSON output")
	}
}

func TestTSHarnessTemplate_TypeOnly(t *testing.T) {
	lang := language.ByName("typescript")
	if lang == nil {
		t.Fatal("expected typescript language to be registered")
	}

	tmpl, err := template.New("harness").Parse(lang.HarnessTemplate)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	testData := []BatchTestData{
		{GroupID: "group_0", Tests: []TestCase{{Data: map[string]any{"name": "test"}, Valid: true}}},
	}
	testDataJSON, _ := json.Marshal(testData)
	testDataString, _ := json.Marshal(string(testDataJSON))

	data := HarnessTemplateData{
		Imports: "",
		Schemas: []HarnessSchemaEntry{
			{
				GroupID:    "group_0",
				Schema:     "",
				Type:       `{ name: string }`,
				Validate:   "",
				IsTypeOnly: true,
			},
		},
		TestDataString: string(testDataString),
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		t.Fatalf("failed to execute template: %v", err)
	}

	result := buf.String()

	// Should have type declaration
	if !strings.Contains(result, `type GeneratedType = { name: string };`) {
		t.Error("missing type declaration")
	}

	// Should have isTypeOnly: true
	if !strings.Contains(result, `isTypeOnly: true`) {
		t.Error("missing isTypeOnly: true")
	}

	// Should output JSON
	if !strings.Contains(result, "console.log(JSON.stringify(results));") {
		t.Error("missing JSON output")
	}
}

func TestTSHarnessTemplate_MultipleImports(t *testing.T) {
	lang := language.ByName("typescript")
	if lang == nil {
		t.Fatal("expected typescript language to be registered")
	}

	tmpl, err := template.New("harness").Parse(lang.HarnessTemplate)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	testData := []BatchTestData{
		{GroupID: "group_0", Tests: []TestCase{}},
	}
	testDataJSON, _ := json.Marshal(testData)
	testDataString, _ := json.Marshal(string(testDataJSON))

	data := HarnessTemplateData{
		Imports: "import { Schema } from \"effect\"\nimport { pipe } from \"effect/Function\"",
		Schemas: []HarnessSchemaEntry{
			{
				GroupID:    "group_0",
				Schema:     `Schema.Struct({ name: Schema.String })`,
				Validate:   `(data) => Schema.decodeUnknownSync(schema)(data) !== undefined`,
				IsTypeOnly: false,
			},
		},
		TestDataString: string(testDataString),
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		t.Fatalf("failed to execute template: %v", err)
	}

	result := buf.String()

	// Should have both imports
	if !strings.Contains(result, `import { Schema } from "effect"`) {
		t.Error("missing first import")
	}
	if !strings.Contains(result, `import { pipe } from "effect/Function"`) {
		t.Error("missing second import")
	}
}

func TestTSHarnessTemplate_EscapedTestData(t *testing.T) {
	lang := language.ByName("typescript")
	if lang == nil {
		t.Fatal("expected typescript language to be registered")
	}

	tmpl, err := template.New("harness").Parse(lang.HarnessTemplate)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	// Test with properly escaped JSON string (as it would come from json.Marshal)
	testData := []BatchTestData{
		{GroupID: "group_0", Tests: []TestCase{
			{Description: "test with \"quotes\"", Data: "hello\nworld", Valid: true},
		}},
	}
	testDataJSON, _ := json.Marshal(testData)
	testDataString, _ := json.Marshal(string(testDataJSON))

	data := HarnessTemplateData{
		Imports: `import { z } from "zod"`,
		Schemas: []HarnessSchemaEntry{
			{
				GroupID:    "group_0",
				Schema:     `z.string()`,
				Validate:   `(data) => schema.safeParse(data).success`,
				IsTypeOnly: false,
			},
		},
		TestDataString: string(testDataString),
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		t.Fatalf("failed to execute template: %v", err)
	}

	result := buf.String()

	// Should contain JSON.parse with escaped string
	if !strings.Contains(result, "JSON.parse(") {
		t.Error("missing JSON.parse")
	}

	// The escaped string should be present (double-escaped in the JS string literal)
	if !strings.Contains(result, `\\\"quotes\\\"`) {
		t.Error("missing escaped quotes in test data")
	}
}
