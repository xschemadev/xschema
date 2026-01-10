package compliance

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
)

func TestGenerateTempHarness(t *testing.T) {
	tmpDir := t.TempDir()

	adapterOutput := &AdapterOutput{
		Schema:   `z.object({ name: z.string() })`,
		Imports:  []string{`import { z } from "zod"`},
		Validate: `(data) => schema.safeParse(data).success`,
	}
	testCases := []TestCase{
		{Description: "valid object", Data: map[string]interface{}{"name": "test"}, Valid: true},
		{Description: "invalid object", Data: "not an object", Valid: false},
	}

	harnessPath, err := GenerateTempHarness(LanguageTypeScript, adapterOutput, testCases, tmpDir)
	if err != nil {
		t.Fatalf("GenerateTempHarness() error = %v", err)
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
	if !strings.Contains(string(content), adapterOutput.Schema) {
		t.Error("harness missing schema")
	}

	// Check import was inserted
	if !strings.Contains(string(content), `import { z } from "zod"`) {
		t.Error("harness missing import")
	}

	// Check validate function was inserted
	if !strings.Contains(string(content), adapterOutput.Validate) {
		t.Error("harness missing validate function")
	}

	// Check test cases were inserted
	if !strings.Contains(string(content), "valid object") {
		t.Error("harness missing test case descriptions")
	}
}

func TestGenerateTempHarness_MergedImports(t *testing.T) {
	tmpDir := t.TempDir()

	adapterOutput := &AdapterOutput{
		Schema:          `z.object({ name: z.string() })`,
		Imports:         []string{`import { z } from "zod"`, `import { helper } from "helper"`},
		Validate:        `(data) => schema.safeParse(data).success`,
		ValidateImports: []string{`import { ZodError } from "zod"`, `import { helper } from "helper"`},
	}
	testCases := []TestCase{{Description: "valid", Data: map[string]interface{}{"name": "test"}, Valid: true}}

	harnessPath, err := GenerateTempHarness(LanguageTypeScript, adapterOutput, testCases, tmpDir)
	if err != nil {
		t.Fatalf("GenerateTempHarness() error = %v", err)
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

func TestGenerateTempHarness_TypeOnly(t *testing.T) {
	tmpDir := t.TempDir()

	// Type-only adapter has empty Validate
	adapterOutput := &AdapterOutput{
		Schema:   "",
		Type:     `{ name: string }`,
		Imports:  nil,
		Validate: "",
	}
	testCases := []TestCase{
		{Description: "test 1", Data: map[string]interface{}{"name": "test"}, Valid: true},
	}

	harnessPath, err := GenerateTempHarness(LanguageTypeScript, adapterOutput, testCases, tmpDir)
	if err != nil {
		t.Fatalf("GenerateTempHarness() error = %v", err)
	}
	defer os.Remove(harnessPath)

	content, _ := os.ReadFile(harnessPath)
	contentStr := string(content)

	// Should have type declaration (type-only mode)
	if !strings.Contains(contentStr, "type GeneratedType") {
		t.Error("type-only harness should have type declaration")
	}

	if !strings.Contains(contentStr, adapterOutput.Type) {
		t.Error("type-only harness should include type expression")
	}

	// Should mark results as skipped
	if !strings.Contains(contentStr, `actual: "skipped"`) {
		t.Error("type-only harness should mark results as skipped")
	}

	// Should NOT have const schema = (runtime mode)
	if strings.Contains(contentStr, "const schema =") {
		t.Error("type-only harness should not have const schema")
	}
}

func TestGenerateTempHarness_UnsupportedLanguage(t *testing.T) {
	tmpDir := t.TempDir()

	adapterOutput := &AdapterOutput{
		Schema:   "some_schema",
		Validate: "validate_fn",
	}

	_, err := GenerateTempHarness(Language("unsupported"), adapterOutput, []TestCase{}, tmpDir)
	if err == nil {
		t.Error("GenerateTempHarness() expected error for unsupported language")
	}
}

func TestGenerateTempHarness_ComplexTestData(t *testing.T) {
	tmpDir := t.TempDir()

	adapterOutput := &AdapterOutput{
		Schema:   `z.any()`,
		Imports:  []string{`import { z } from "zod"`},
		Validate: `(data) => true`,
	}

	// Test with complex nested data
	testCases := []TestCase{
		{
			Description: "nested object",
			Data: map[string]interface{}{
				"nested": map[string]interface{}{
					"array": []interface{}{1, 2, 3},
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
	}

	harnessPath, err := GenerateTempHarness(LanguageTypeScript, adapterOutput, testCases, tmpDir)
	if err != nil {
		t.Fatalf("GenerateTempHarness() error = %v", err)
	}
	defer os.Remove(harnessPath)

	content, _ := os.ReadFile(harnessPath)
	contentStr := string(content)

	// Should contain JSON.parse for test cases
	if !strings.Contains(contentStr, "JSON.parse(") {
		t.Error("harness should use JSON.parse for test cases")
	}
}

func TestGenerateTempHarness_EmptyTestCases(t *testing.T) {
	tmpDir := t.TempDir()

	adapterOutput := &AdapterOutput{
		Schema:   `z.string()`,
		Imports:  []string{`import { z } from "zod"`},
		Validate: `(data) => schema.safeParse(data).success`,
	}

	harnessPath, err := GenerateTempHarness(LanguageTypeScript, adapterOutput, []TestCase{}, tmpDir)
	if err != nil {
		t.Fatalf("GenerateTempHarness() error = %v", err)
	}
	defer os.Remove(harnessPath)

	content, _ := os.ReadFile(harnessPath)

	// Should still have empty JSON array
	if !strings.Contains(string(content), `"[]"`) {
		t.Error("empty test cases should produce escaped empty array")
	}
}

func TestGenerateTempHarness_FilenamePattern(t *testing.T) {
	tmpDir := t.TempDir()

	adapterOutput := &AdapterOutput{
		Schema:   `z.string()`,
		Imports:  []string{`import { z } from "zod"`},
		Validate: `(data) => schema.safeParse(data).success`,
	}

	harnessPath, err := GenerateTempHarness(LanguageTypeScript, adapterOutput, []TestCase{}, tmpDir)
	if err != nil {
		t.Fatalf("GenerateTempHarness() error = %v", err)
	}
	defer os.Remove(harnessPath)

	// Filename should match pattern "xschema-harness-*"
	basename := filepath.Base(harnessPath)
	if !strings.HasPrefix(basename, "xschema-harness-") {
		t.Errorf("harness filename = %s, want prefix 'xschema-harness-'", basename)
	}
}

func TestTSHarnessTemplate_RuntimeValidation(t *testing.T) {
	tmpl, err := template.New("harness").Parse(TSHarnessTemplate)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	data := HarnessTemplateData{
		Schema:          `z.object({ name: z.string() })`,
		Imports:         `import { z } from "zod"`,
		Validate:        `(data) => schema.safeParse(data).success`,
		TestCasesString: `"[{\"data\":{\"name\":\"test\"},\"valid\":true}]"`,
		IsTypeOnly:      false,
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

	// Should have schema
	if !strings.Contains(result, `const schema = z.object({ name: z.string() });`) {
		t.Error("missing schema declaration")
	}

	// Should have validate function
	if !strings.Contains(result, `const validate = (data) => schema.safeParse(data).success;`) {
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

	// Should NOT have type-only code
	if strings.Contains(result, "type GeneratedType") {
		t.Error("should not have type-only code for runtime adapter")
	}
}

func TestTSHarnessTemplate_TypeOnly(t *testing.T) {
	tmpl, err := template.New("harness").Parse(TSHarnessTemplate)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	data := HarnessTemplateData{
		Schema:          "",
		Type:            `{ name: string }`,
		Imports:         "",
		Validate:        "",
		TestCasesString: `"[{\"data\":{\"name\":\"test\"},\"valid\":true}]"`,
		IsTypeOnly:      true,
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

	// Should mark results as skipped
	if !strings.Contains(result, `actual: "skipped"`) {
		t.Error("missing skipped actual value")
	}

	// Should NOT have runtime validation code
	if strings.Contains(result, "const schema =") {
		t.Error("should not have schema const for type-only adapter")
	}
	if strings.Contains(result, "const validate =") {
		t.Error("should not have validate function for type-only adapter")
	}
	if strings.Contains(result, "try {") {
		t.Error("should not have try-catch for type-only adapter")
	}

	// Should output JSON
	if !strings.Contains(result, "console.log(JSON.stringify(results));") {
		t.Error("missing JSON output")
	}
}

func TestTSHarnessTemplate_WithValidateImports(t *testing.T) {
	tmpl, err := template.New("harness").Parse(TSHarnessTemplate)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	data := HarnessTemplateData{
		Schema:          `v.object({ name: v.string() })`,
		Imports:         "import * as v from \"valibot\"\nimport { safeParse } from \"valibot\"",
		Validate:        `(data) => safeParse(schema, data).success`,
		TestCasesString: `"[]"`,
		IsTypeOnly:      false,
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		t.Fatalf("failed to execute template: %v", err)
	}

	result := buf.String()

	// Should have both imports
	if !strings.Contains(result, `import * as v from "valibot"`) {
		t.Error("missing schema import")
	}
	if !strings.Contains(result, `import { safeParse } from "valibot"`) {
		t.Error("missing validateImports")
	}

	if strings.Contains(result, "import import ") {
		t.Error("should not prefix import statements")
	}
}

func TestTSHarnessTemplate_ValidateImportsOnly(t *testing.T) {
	tmpl, err := template.New("harness").Parse(TSHarnessTemplate)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	data := HarnessTemplateData{
		Schema:          `v.object({ name: v.string() })`,
		Imports:         "import * as v from \"valibot\"\nimport { safeParse } from \"valibot\"",
		Validate:        `(data) => safeParse(schema, data).success`,
		TestCasesString: `"[]"`,
		IsTypeOnly:      false,
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		t.Fatalf("failed to execute template: %v", err)
	}

	result := buf.String()

	if !strings.Contains(result, `import { safeParse } from "valibot"`) {
		t.Error("missing validateImports safeParse")
	}
	if !strings.Contains(result, `import * as v from "valibot"`) {
		t.Error("missing validateImports namespace import")
	}
}

func TestTSHarnessTemplate_MultipleImports(t *testing.T) {
	tmpl, err := template.New("harness").Parse(TSHarnessTemplate)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	data := HarnessTemplateData{
		Schema:          `Schema.Struct({ name: Schema.String })`,
		Imports:         "import { Schema } from \"effect\"\nimport { pipe } from \"effect/Function\"",
		Validate:        `(data) => Schema.decodeUnknownSync(schema)(data) !== undefined`,
		TestCasesString: `"[]"`,
		IsTypeOnly:      false,
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

func TestTSHarnessTemplate_EscapedTestCases(t *testing.T) {
	tmpl, err := template.New("harness").Parse(TSHarnessTemplate)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	// Test with properly escaped JSON string (as it would come from json.Marshal)
	testCases := []TestCase{
		{Description: "test with \"quotes\"", Data: "hello\nworld", Valid: true},
	}
	testCasesJSON, _ := json.Marshal(testCases)
	testCasesString, _ := json.Marshal(string(testCasesJSON))

	data := HarnessTemplateData{
		Schema:          `z.string()`,
		Imports:         `import { z } from "zod"`,
		Validate:        `(data) => schema.safeParse(data).success`,
		TestCasesString: string(testCasesString),
		IsTypeOnly:      false,
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
		t.Error("missing escaped quotes in test cases")
	}
}
