package language

import (
	"slices"
	"testing"
)

func TestRegistry_BuiltinsAndLookups(t *testing.T) {
	ResetForTests()
	t.Cleanup(ResetForTests)

	langs := SupportedLanguages()
	if !slices.Contains(langs, "typescript") {
		t.Fatalf("expected built-in language typescript, got %v", langs)
	}
	if !slices.Contains(langs, "py") {
		t.Fatalf("expected built-in language py, got %v", langs)
	}
	if !slices.IsSorted(langs) {
		t.Fatalf("expected SupportedLanguages sorted, got %v", langs)
	}

	if ByName("typescript") == nil {
		t.Fatal("expected ByName(typescript) to return a language")
	}
	if BySchemaURL(XSchemaBaseURL+"typescript.jsonc") == nil {
		t.Fatal("expected BySchemaURL(typescript.jsonc) to return a language")
	}
	if BySchemaURL("https://example.com/schemas/typescript.jsonc") != nil {
		t.Fatal("expected BySchemaURL(non-xschema) to return nil")
	}

	ignore := AllIgnoreDirs()
	if !ignore["node_modules"] {
		t.Fatalf("expected node_modules to be ignored, got %v", ignore)
	}
	if !ignore["__pycache__"] {
		t.Fatalf("expected __pycache__ to be ignored, got %v", ignore)
	}
}

func TestRegistry_RegisterValidationAndReset(t *testing.T) {
	ResetForTests()
	t.Cleanup(ResetForTests)

	custom := Language{
		Name:       "test",
		SchemaURL:  XSchemaBaseURL + "test.jsonc",
		Extensions: []string{".test"},
		IgnoreDirs: []string{".test-cache"},
	}
	if err := Register(custom); err != nil {
		t.Fatalf("expected Register to succeed, got %v", err)
	}

	got := ByName("test")
	if got == nil {
		t.Fatal("expected ByName(test) to return a language")
	}
	if got.SchemaExt != "test.jsonc" {
		t.Fatalf("expected schemaExt to be derived from schema url, got %q", got.SchemaExt)
	}
	if BySchemaURL(XSchemaBaseURL+"test.jsonc") == nil {
		t.Fatal("expected BySchemaURL(test.jsonc) to return a language")
	}
	if !AllIgnoreDirs()[".test-cache"] {
		t.Fatal("expected .test-cache to be included in ignore dirs")
	}

	if err := Register(Language{Name: "typescript", SchemaURL: XSchemaBaseURL + "another.jsonc"}); err == nil {
		t.Fatal("expected duplicate language name error")
	}
	if err := Register(Language{Name: "another", SchemaURL: XSchemaBaseURL + "typescript.jsonc"}); err == nil {
		t.Fatal("expected duplicate schema url mapping error")
	}
	if err := Register(Language{Name: "bad", SchemaURL: "https://example.com/bad.jsonc"}); err == nil {
		t.Fatal("expected non-xschema schema url error")
	}

	ResetForTests()
	if ByName("test") != nil {
		t.Fatal("expected ResetForTests to remove registered test language")
	}
}
