package parser

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/xschema/cli/language"
)

type expectedDecl struct {
	Name     string `json:"name"`
	Source   string `json:"source"`
	Location string `json:"location,omitempty"`
}

// commonTests are run for all languages
var commonTests = []string{"basic", "edge_cases", "invalid"}

func TestCommon(t *testing.T) {
	for _, lang := range language.Languages {
		for _, testName := range commonTests {
			t.Run(lang.Name+"/"+testName, func(t *testing.T) {
				// Load expected from common/
				expectedPath := filepath.Join("testdata", "common", testName+".json")
				expectedData, err := os.ReadFile(expectedPath)
				if err != nil {
					t.Fatalf("read expected: %v", err)
				}
				var expected []expectedDecl
				if err := json.Unmarshal(expectedData, &expected); err != nil {
					t.Fatalf("parse expected: %v", err)
				}

				// Find source file for this language
				var sourceFile string
				for _, ext := range lang.Extensions {
					candidate := filepath.Join("testdata", lang.Name, testName+ext)
					if _, err := os.Stat(candidate); err == nil {
						sourceFile = candidate
						break
					}
				}
				if sourceFile == "" {
					t.Skipf("no source file for %s/%s", lang.Name, testName)
				}

				// Parse
				decls, err := parseFile(context.Background(), sourceFile, &lang)
				if err != nil {
					t.Fatalf("parseFile: %v", err)
				}

				// Assert
				assertDecls(t, decls, expected)
			})
		}
	}
}

// Language-specific string tests
func TestTypeScriptStrings(t *testing.T) {
	lang := language.ByExtension(".ts")
	decls, err := parseFile(context.Background(), "testdata/typescript/strings.ts", lang)
	if err != nil {
		t.Fatalf("parseFile: %v", err)
	}

	expected := []expectedDecl{
		{"DoubleQuote", "url", "https://example.com/a.json"},
		{"SingleQuote", "url", "https://example.com/b.json"},
		{"TemplateLit", "url", "https://example.com/c.json"},
	}
	assertDecls(t, decls, expected)

	// Verify interpolated ones are NOT found
	for _, d := range decls {
		if d.Name == "Interpolated" || d.Name == "Schema${version}" {
			t.Errorf("should NOT find %q (template interpolation)", d.Name)
		}
	}
}

func TestPythonStrings(t *testing.T) {
	lang := language.ByExtension(".py")
	decls, err := parseFile(context.Background(), "testdata/python/strings.py", lang)
	if err != nil {
		t.Fatalf("parseFile: %v", err)
	}

	expected := []expectedDecl{
		{"DoubleQuote", "url", "https://example.com/a.json"},
		{"SingleQuote", "url", "https://example.com/b.json"},
		{"TripleDouble", "url", "https://example.com/c.json"},
		{"TripleSingle", "url", "https://example.com/d.json"},
		{"RawString", "file", "./schemas/raw.json"},
	}
	assertDecls(t, decls, expected)
}

// Metadata tests
func TestLineNumbers(t *testing.T) {
	lang := language.ByExtension(".ts")
	decls, err := parseFile(context.Background(), "testdata/typescript/basic.ts", lang)
	if err != nil {
		t.Fatalf("parseFile: %v", err)
	}

	prevLine := 0
	for _, d := range decls {
		if d.Line <= 0 {
			t.Errorf("%q: line should be > 0, got %d", d.Name, d.Line)
		}
		if d.Line < prevLine {
			t.Errorf("%q: line %d should be >= prev line %d", d.Name, d.Line, prevLine)
		}
		prevLine = d.Line
	}
}

func TestAdapterCapture(t *testing.T) {
	lang := language.ByExtension(".ts")
	decls, err := parseFile(context.Background(), "testdata/typescript/basic.ts", lang)
	if err != nil {
		t.Fatalf("parseFile: %v", err)
	}

	for _, d := range decls {
		if d.Adapter.Name != "zodAdapter" {
			t.Errorf("%q: expected adapter=zodAdapter, got %q", d.Name, d.Adapter.Name)
		}
		if d.Adapter.Language != "typescript" {
			t.Errorf("%q: expected language=typescript, got %q", d.Name, d.Adapter.Language)
		}
	}
}

func TestFilePath(t *testing.T) {
	lang := language.ByExtension(".ts")
	decls, err := parseFile(context.Background(), "testdata/typescript/basic.ts", lang)
	if err != nil {
		t.Fatalf("parseFile: %v", err)
	}

	for _, d := range decls {
		if d.File != "testdata/typescript/basic.ts" {
			t.Errorf("%q: expected file path, got %q", d.Name, d.File)
		}
	}
}

func TestParseDirectory(t *testing.T) {
	decls, err := Parse(context.Background(), "testdata")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}

	langCounts := make(map[string]int)
	for _, d := range decls {
		langCounts[d.Adapter.Language]++
	}

	for _, lang := range language.Languages {
		if langCounts[lang.Name] == 0 {
			t.Errorf("expected %s declarations", lang.Name)
		}
	}
	t.Logf("Found declarations: %v", langCounts)
}

// assertDecls checks decls match expected (order matters)
func assertDecls(t *testing.T, decls []Declaration, expected []expectedDecl) {
	t.Helper()

	if len(decls) != len(expected) {
		t.Fatalf("expected %d decls, got %d", len(expected), len(decls))
	}

	for i, exp := range expected {
		d := decls[i]
		if d.Name != exp.Name {
			t.Errorf("decl[%d]: expected name %q, got %q", i, exp.Name, d.Name)
		}
		if string(d.Source) != exp.Source {
			t.Errorf("decl[%d] %q: expected source %q, got %q", i, d.Name, exp.Source, d.Source)
		}
		if exp.Location != "" && d.Location != exp.Location {
			t.Errorf("decl[%d] %q: expected location %q, got %q", i, d.Name, exp.Location, d.Location)
		}
	}
}
