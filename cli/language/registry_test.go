package language_test

import (
	"slices"
	"testing"

	"github.com/xschemadev/xschema/language"
	_ "github.com/xschemadev/xschema/language/langs"
	"github.com/xschemadev/xschema/language/langs/typescript"
)

func TestRegistry_BuiltinsAndLookups(t *testing.T) {
	langs := language.SupportedLanguages()
	if !slices.Contains(langs, "typescript") {
		t.Fatalf("expected built-in language typescript, got %v", langs)
	}
	if !slices.IsSorted(langs) {
		t.Fatalf("expected SupportedLanguages sorted, got %v", langs)
	}

	if language.ByName("typescript") == nil {
		t.Fatal("expected ByName(typescript) to return a language")
	}
	if language.BySchemaURL(language.XSchemaBaseURL+"typescript.jsonc") == nil {
		t.Fatal("expected BySchemaURL(typescript.jsonc) to return a language")
	}
	if language.BySchemaURL("https://example.com/schemas/typescript.jsonc") != nil {
		t.Fatal("expected BySchemaURL(non-xschema) to return nil")
	}

	ignore := language.AllIgnoreDirs()
	if !ignore["node_modules"] {
		t.Fatalf("expected node_modules to be ignored, got %v", ignore)
	}
}

func TestRegistry_RegisterValidationAndReset(t *testing.T) {
	language.ResetForTests()
	t.Cleanup(language.ResetForTests)

	if err := language.Register(typescript.Language()); err != nil {
		t.Fatalf("expected Register(typescript) to succeed, got %v", err)
	}

	custom := language.Language{
		Name:       "test",
		SchemaURL:  language.XSchemaBaseURL + "test.jsonc",
		Extensions: []string{".test"},
		IgnoreDirs: []string{".test-cache"},
	}
	if err := language.Register(custom); err != nil {
		t.Fatalf("expected Register to succeed, got %v", err)
	}

	got := language.ByName("test")
	if got == nil {
		t.Fatal("expected ByName(test) to return a language")
	}
	if got.SchemaExt != "test.jsonc" {
		t.Fatalf("expected schemaExt to be derived from schema url, got %q", got.SchemaExt)
	}
	if language.BySchemaURL(language.XSchemaBaseURL+"test.jsonc") == nil {
		t.Fatal("expected BySchemaURL(test.jsonc) to return a language")
	}
	if !language.AllIgnoreDirs()[".test-cache"] {
		t.Fatal("expected .test-cache to be included in ignore dirs")
	}

	if err := language.Register(language.Language{Name: "typescript", SchemaURL: language.XSchemaBaseURL + "another.jsonc"}); err == nil {
		t.Fatal("expected duplicate language name error")
	}
	if err := language.Register(language.Language{Name: "another", SchemaURL: language.XSchemaBaseURL + "typescript.jsonc"}); err == nil {
		t.Fatal("expected duplicate schema url mapping error")
	}
	if err := language.Register(language.Language{Name: "bad", SchemaURL: "https://example.com/bad.jsonc"}); err == nil {
		t.Fatal("expected non-xschema schema url error")
	}

	language.ResetForTests()
	if language.ByName("test") != nil {
		t.Fatal("expected ResetForTests to remove registered test language")
	}
}
