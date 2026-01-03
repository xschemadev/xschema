package language

import (
	"os"
	"os/exec"
	"path/filepath"
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
	Name         string
	Extensions   []string // file extensions for source files (for injector)
	SchemaURL    string   // e.g., "https://xschema.dev/schemas/ts.jsonc"
	SchemaExt    string   // e.g., "ts.jsonc" - extracted from SchemaURL
	DetectRunner func() (cmd string, args []string, err error)

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
}

var Languages = []Language{
	{
		Name:                 "typescript",
		Extensions:           []string{".ts", ".tsx", ".js", ".jsx"},
		SchemaURL:            XSchemaBaseURL + "ts.jsonc",
		SchemaExt:            "ts.jsonc",
		DetectRunner:         detectTSRunner,
		BuildSchemasImport:   buildTSSchemasImport,
		ImportPattern:        `(?m)^import\s+.*$`,
		InjectSchemasKey:     injectSchemasKeyBrace,
		ClientFactoryPattern: `createXSchemaClient\s*\(\s*(\{[^}]*\})\s*\)`,
		OutputFile:           "xschema.gen.ts",
		Template:             TSTemplate,
		MergeImports:         MergeTSImports,
		BuildVarName:         buildVarNameUnderscore,
		IgnoreDirs:           []string{"node_modules", "dist", "build"},
	},
	{
		Name:                 "python",
		Extensions:           []string{".py"},
		SchemaURL:            XSchemaBaseURL + "py.jsonc",
		SchemaExt:            "py.jsonc",
		DetectRunner:         detectPythonRunner,
		BuildSchemasImport:   buildPySchemasImport,
		ImportPattern:        `(?m)^(?:import\s+|from\s+).*$`,
		InjectSchemasKey:     injectSchemasKeyBrace,
		ClientFactoryPattern: `create_xschema_client\s*\(\s*(\{[^}]*\})\s*\)`,
		OutputFile:           "__init__.py",
		Template:             PyTemplate,
		MergeImports:         MergePyImports,
		BuildFooter:          BuildPythonFooter,
		BuildVarName:         buildVarNameUnderscore,
		IgnoreDirs:           []string{"__pycache__", ".venv", "venv"},
	},
}

// languageBySchemaExt maps schema extensions to languages
var languageBySchemaExt map[string]*Language

func init() {
	languageBySchemaExt = make(map[string]*Language)
	for i := range Languages {
		languageBySchemaExt[Languages[i].SchemaExt] = &Languages[i]
	}
}

// BySchemaURL returns the language for a $schema URL like "https://xschema.dev/schemas/ts.jsonc"
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

func detectTSRunner() (string, []string, error) {
	checkCmd := func(cmd string) bool {
		_, err := exec.LookPath(cmd)
		return err == nil
	}

	if _, err := os.Stat(filepath.Join(".", "package.json")); err == nil {
		content, err := os.ReadFile(filepath.Join(".", "package.json"))
		if err == nil {
			pm := detectPackageManager(string(content))
			if pm != "" && checkCmd(pm) {
				switch pm {
				case "bun":
					return "bunx", nil, nil
				case "pnpm":
					return "pnpm", []string{"exec"}, nil
				case "yarn":
					return "yarn", nil, nil
				case "npm":
					return "npx", nil, nil
				}
			}
		}
	}

	lockfileCmds := map[string][]string{
		"bun.lock":          {"bunx"},
		"bun.lockb":         {"bunx"},
		"pnpm-lock.yaml":    {"pnpm", "exec"},
		"yarn.lock":         {"yarn"},
		"package-lock.json": {"npx"},
	}

	for lf, cmd := range lockfileCmds {
		if _, err := os.Stat(filepath.Join(".", lf)); err == nil {
			if checkCmd(cmd[0]) {
				return cmd[0], cmd[1:], nil
			}
		}
	}

	for _, cmd := range []string{"bunx", "pnpm", "yarn", "npx"} {
		if checkCmd(cmd) {
			if cmd == "pnpm" {
				return cmd, []string{"exec"}, nil
			}
			return cmd, nil, nil
		}
	}

	return "npx", nil, nil
}

func detectPackageManager(content string) string {
	lines := strings.SplitSeq(content, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, `"packageManager"`) {
			if strings.Contains(line, `"bun@`) {
				return "bun"
			}
			if strings.Contains(line, `"pnpm@`) {
				return "pnpm"
			}
			if strings.Contains(line, `"yarn@`) {
				return "yarn"
			}
			if strings.Contains(line, `"npm@`) {
				return "npm"
			}
		}
	}
	return ""
}

func detectPythonRunner() (string, []string, error) {
	checkCmd := func(cmd string) bool {
		_, err := exec.LookPath(cmd)
		return err == nil
	}

	lockfileCmds := map[string][]string{
		"uv.lock":     {"uv", "run"},
		"poetry.lock": {"poetry", "run"},
		"Pipfile":     {"pipenv", "run"},
	}

	for lf, cmd := range lockfileCmds {
		if _, err := os.Stat(filepath.Join(".", lf)); err == nil {
			if checkCmd(cmd[0]) {
				return cmd[0], cmd[1:], nil
			}
		}
	}

	if _, err := os.Stat(filepath.Join(".", "pyproject.toml")); err == nil {
		content, err := os.ReadFile(filepath.Join(".", "pyproject.toml"))
		if err == nil {
			buildSystem := detectBuildSystem(string(content))
			if buildSystem != "" && checkCmd(buildSystem) {
				return buildSystem, []string{"run"}, nil
			}
		}
	}

	return "python", []string{"-m"}, nil
}

func detectBuildSystem(content string) string {
	lines := strings.SplitSeq(content, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "build-backend") && strings.Contains(line, "uv") {
			return "uv"
		}
		if strings.Contains(line, "requires") && strings.Contains(line, "poetry-core") {
			return "poetry"
		}
		if strings.Contains(line, "requires") && strings.Contains(line, "flit") {
			return "flit"
		}
		if strings.Contains(line, "requires") && strings.Contains(line, "setuptools") {
			return ""
		}
	}
	return ""
}

// buildTSSchemasImport builds TypeScript import for schemas
func buildTSSchemasImport(importPath string) string {
	return `import { schemas } from "` + importPath + `";`
}

// buildPySchemasImport builds Python import for schemas
func buildPySchemasImport(importPath string) string {
	// Convert path to module notation: ./.xschema/xschema -> .xschema.xschema
	module := strings.ReplaceAll(importPath, "/", ".")
	module = strings.TrimPrefix(module, ".")
	return "from " + module + " import schemas"
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
