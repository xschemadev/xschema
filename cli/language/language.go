package language

import (
	"os/exec"
	"strings"
)

const (
	// XSchemaBaseURL is the base URL for xschema.dev schema files
	XSchemaBaseURL = "https://xschema.dev/schemas/"
)

// SchemaEntry represents a generated schema for template data
type SchemaEntry struct {
	Namespace string // e.g., "user"
	ID        string // e.g., "TestUrl"
	VarName   string // e.g., "user_TestUrl" (safe variable name)
	Code      string // generated schema code
	Type      string // type expression
}

// Key returns the full namespaced key like "user:TestUrl"
func (s SchemaEntry) Key() string {
	return s.Namespace + ":" + s.ID
}

type Language struct {
	Name             string
	Extensions       []string // file extensions for source files (for injector)
	SchemaURL        string   // e.g., "https://xschema.dev/schemas/typescript.jsonc"
	SchemaExt        string   // e.g., "ts.jsonc" - extracted from SchemaURL
	AdapterBinPrefix string   // e.g., "xschema-" - prefix for adapter binaries
	DetectRunner     func() (cmd string, args []string, err error)

	// Client injection (after generation)
	BuildSchemasImport   func(importPath string) string    // build import statement for schemas
	ImportPattern        string                            // regex to find import lines
	InjectSchemasKey     func(configContent string) string // inject "schemas" into config object
	ClientFactoryPattern string                            // regex to find client factory calls e.g. createXSchemaClient({ ... })

	// Output generation
	OutputFile   string                                            // e.g. "xschema.gen.ts", "__init__.py"
	Template     string                                            // Go text/template for output
	MergeImports func(imports []string) string                     // dedupe/format imports
	BuildHeader  func(outDir string, schemas []SchemaEntry) string // inserted at top
	BuildFooter  func(outDir string, schemas []SchemaEntry) string // inserted at bottom
	BuildVarName func(namespace, id string) string                 // build variable name from namespace and id

	// Parser (fallback when git not available)
	IgnoreDirs []string // directories to skip when walking

	// Compliance testing
	DetectHarnessRunner func(dir string) (cmd string, args []string, err error) // detect runner for local script files
	GetPackageName      func(string) string                                     // function to get package name from a directory
	AdapterCLIPath      func(adapterPath string) string                         // returns path to adapter CLI binary (e.g., "dist/cli.js" for TS)
}

var Languages = []Language{
	typescript,
	python,
}

// languageBySchemaExt maps schema extensions to languages
var languageBySchemaExt map[string]*Language

func init() {
	languageBySchemaExt = make(map[string]*Language)
	for i := range Languages {
		languageBySchemaExt[Languages[i].SchemaExt] = &Languages[i]
	}
}

// BySchemaURL returns the language for a $schema URL like "https://xschema.dev/schemas/typescript.jsonc"
// Returns nil if URL doesn't match xschema.dev pattern
func BySchemaURL(url string) *Language {
	if !strings.HasPrefix(url, XSchemaBaseURL) {
		return nil
	}
	ext := strings.TrimPrefix(url, XSchemaBaseURL)
	return languageBySchemaExt[ext]
}

// ByName returns the language config by name
func ByName(name string) *Language {
	for i, lang := range Languages {
		if lang.Name == name {
			return &Languages[i]
		}
	}
	return nil
}

// AllIgnoreDirs returns a combined set of all ignore dirs from all languages
// Used when walking directories before language detection
func AllIgnoreDirs() map[string]bool {
	dirs := make(map[string]bool)
	for _, lang := range Languages {
		for _, dir := range lang.IgnoreDirs {
			dirs[dir] = true
		}
	}
	return dirs
}

// IsXSchemaURL checks if a URL is an xschema.dev schema URL
func IsXSchemaURL(url string) bool {
	return strings.HasPrefix(url, XSchemaBaseURL)
}

// commandExists checks if a command is available in PATH
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// buildVarNameUnderscore builds a variable name using underscore separator: namespace_id
func buildVarNameUnderscore(namespace, id string) string {
	return namespace + "_" + id
}

// injectSchemasKeyBrace injects "schemas" into a brace-delimited config (JS/TS/Python dict)
func injectSchemasKeyBrace(configContent string) string {
	// Find first { and insert after it
	openIdx := strings.Index(configContent, "{")
	if openIdx == -1 {
		return configContent
	}

	// Check length to avoid panic
	if len(configContent) < openIdx+2 {
		return "{ schemas }"
	}

	// Check if schemas already exists (shorthand {schemas} or pair {schemas: schemas})
	inner := configContent[openIdx+1 : len(configContent)-1]
	innerTrimmed := strings.TrimSpace(inner)

	// Check for shorthand: {schemas, ...}
	if strings.HasPrefix(innerTrimmed, "schemas") && (len(innerTrimmed) == 7 || strings.HasPrefix(innerTrimmed[7:], ",") || strings.HasPrefix(innerTrimmed[7:], "}")) {
		return configContent
	}

	// Check for pair: {schemas: schemas, ...}
	if strings.HasPrefix(innerTrimmed, "schemas:") {
		return configContent
	}

	// Check for quoted pair: {"schemas": schemas, ...}
	if strings.HasPrefix(innerTrimmed, `"schemas":`) || strings.HasPrefix(innerTrimmed, `'schemas':`) {
		return configContent
	}

	if innerTrimmed == "" {
		return "{ schemas }"
	}
	return "{ schemas, " + inner + " }"
}
