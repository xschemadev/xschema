package python

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xschemadev/xschema/language"
)

const adapterBinPrefix = "xschema-"

func init() {
	if err := language.Register(Language()); err != nil {
		panic(fmt.Sprintf("language python register failed: %v", err))
	}
}

func Language() language.Language {
	return language.Language{
		Name:                 "python",
		Extensions:           []string{".py"},
		SchemaURL:            language.XSchemaBaseURL + "python.jsonc",
		SchemaExt:            "python.jsonc",
		AdapterBinPrefix:     adapterBinPrefix,
		DetectRunner:         detectRunner,
		AdapterInvoker:       adapterInvoker{},
		BuildSchemasImport:   buildSchemasImport,
		ImportPattern:        `(?m)^(?:from\s+\S+\s+import\s+.*|import\s+.*)$`,
		InjectSchemasKey:     injectSchemasKeyDict,
		ClientFactoryPattern: `create_xschema_client\s*\(\s*(\{[^}]*\})\s*\)`,
		OutputDir:            "_xschema",
		OutputFile:           "__init__.py",
		Template:             outputTemplate,
		MergeImports:         MergeImports,
		BuildVarName:         buildVarName,
		IgnoreDirs:           []string{"__pycache__", ".venv", "venv", ".tox", ".eggs", "site-packages", ".mypy_cache", ".pytest_cache"},
		DetectHarnessRunner:  detectHarnessRunner,
		GetPackageName:       getPackageName,
		AdapterCLIPath:       getAdapterCLIPath,
		HarnessExtension:     ".py",
		HarnessTemplate:      harnessTemplate,
	}
}

type adapterInvoker struct{}

func (adapterInvoker) BuildAdapterCommand(ctx context.Context, input language.AdapterCommandInput) (language.CommandSpec, error) {
	_ = ctx

	projectRoot := strings.TrimSpace(input.ProjectRoot)
	if projectRoot == "" {
		return language.CommandSpec{}, fmt.Errorf("project root is required")
	}

	adapterRef := strings.TrimSpace(input.AdapterRef)
	if adapterRef == "" {
		return language.CommandSpec{}, fmt.Errorf("adapter ref is required")
	}

	// Validate adapter ref format: xschemadev/<adapter>
	// Migration help for legacy names like "pydantic"
	if !strings.HasPrefix(adapterRef, "xschemadev/") {
		if !strings.Contains(adapterRef, "/") {
			return language.CommandSpec{}, fmt.Errorf(
				"invalid python adapter ref %q: expected PyPI package ref like %q (migration: change adapter to %q)",
				adapterRef,
				"xschemadev/pydantic",
				"xschemadev/"+adapterRef,
			)
		}
		return language.CommandSpec{}, fmt.Errorf(
			"invalid python adapter ref %q: expected %q scope (example: %q)",
			adapterRef,
			"xschemadev",
			"xschemadev/pydantic",
		)
	}

	parts := strings.Split(adapterRef, "/")
	if len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
		return language.CommandSpec{}, fmt.Errorf(
			"invalid python adapter ref %q: expected format %q",
			adapterRef,
			"xschemadev/<adapter>",
		)
	}

	pkgName := parts[1]
	binName := adapterBinPrefix + pkgName

	runner, runnerArgs, err := detectRunnerInDir(projectRoot)
	if err != nil {
		return language.CommandSpec{}, fmt.Errorf("failed to detect python runner: %w", err)
	}

	return language.CommandSpec{
		Cmd:  runner,
		Args: append(runnerArgs, binName),
		Dir:  projectRoot,
	}, nil
}

func buildSchemasImport(importPath string) string {
	return "from " + importPath + " import schemas"
}

func getPackageName(dir string) string {
	// Try pyproject.toml first
	pyprojectPath := filepath.Join(dir, "pyproject.toml")
	if data, err := os.ReadFile(pyprojectPath); err == nil {
		if name := extractPyProjectName(string(data)); name != "" {
			return name
		}
	}

	// Try setup.cfg
	setupCfgPath := filepath.Join(dir, "setup.cfg")
	if data, err := os.ReadFile(setupCfgPath); err == nil {
		if name := extractSetupCfgName(string(data)); name != "" {
			return name
		}
	}

	// Fallback to directory name
	return filepath.Base(dir)
}

func extractPyProjectName(content string) string {
	// Look for name = "..." in [project] section
	inProject := false
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "[project]" {
			inProject = true
			continue
		}
		if strings.HasPrefix(line, "[") && line != "[project]" {
			inProject = false
			continue
		}
		if inProject && strings.HasPrefix(line, "name") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[1])
				name = strings.Trim(name, `"'`)
				if name != "" {
					return name
				}
			}
		}
	}
	return ""
}

func extractSetupCfgName(content string) string {
	// Look for name = ... in [metadata] section
	inMetadata := false
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "[metadata]" {
			inMetadata = true
			continue
		}
		if strings.HasPrefix(line, "[") && line != "[metadata]" {
			inMetadata = false
			continue
		}
		if inMetadata && strings.HasPrefix(line, "name") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[1])
				if name != "" {
					return name
				}
			}
		}
	}
	return ""
}

func getAdapterCLIPath(adapterPath string) string {
	// Python adapters use __main__.py for CLI entry point
	// The runner must set the working directory to adapterPath for imports to resolve
	return filepath.Join(adapterPath, "src", "xschema_pydantic", "__main__.py")
}

// pythonReservedWords contains Python keywords and common builtins to avoid
var pythonReservedWords = map[string]bool{
	// Keywords (Python 3.12)
	"False":    true,
	"None":     true,
	"True":     true,
	"and":      true,
	"as":       true,
	"assert":   true,
	"async":    true,
	"await":    true,
	"break":    true,
	"class":    true,
	"continue": true,
	"def":      true,
	"del":      true,
	"elif":     true,
	"else":     true,
	"except":   true,
	"finally":  true,
	"for":      true,
	"from":     true,
	"global":   true,
	"if":       true,
	"import":   true,
	"in":       true,
	"is":       true,
	"lambda":   true,
	"nonlocal": true,
	"not":      true,
	"or":       true,
	"pass":     true,
	"raise":    true,
	"return":   true,
	"try":      true,
	"while":    true,
	"with":     true,
	"yield":    true,
	// Soft keywords (Python 3.10+)
	"match": true,
	"case":  true,
	"type":  true,
	// Common builtins to avoid shadowing
	"id":        true,
	"list":      true,
	"dict":      true,
	"set":       true,
	"str":       true,
	"int":       true,
	"float":     true,
	"bool":      true,
	"bytes":     true,
	"object":    true,
	"property":  true,
	"super":     true,
	"self":      true,
	"cls":       true,
	"input":     true,
	"print":     true,
	"open":      true,
	"file":      true,
	"range":     true,
	"len":       true,
	"map":       true,
	"filter":    true,
	"zip":       true,
	"all":       true,
	"any":       true,
	"sum":       true,
	"min":       true,
	"max":       true,
	"abs":       true,
	"round":     true,
	"sorted":    true,
	"reversed":  true,
	"enumerate": true,
	"format":    true,
	"hash":      true,
	"help":      true,
	"hex":       true,
	"iter":      true,
	"next":      true,
	"repr":      true,
	"slice":     true,
	"vars":      true,
}

func buildVarName(namespace, id string) string {
	ns := toSnakeCase(namespace)
	idSnake := toSnakeCase(id)

	// Handle empty cases
	if ns == "" && idSnake == "" {
		return "schema"
	}
	if ns == "" {
		ns = "schema"
	}
	if idSnake == "" {
		idSnake = "schema"
	}

	name := ns + "_" + idSnake

	// Handle reserved words by appending underscore
	if pythonReservedWords[name] {
		name = name + "_"
	}

	// Ensure valid Python identifier (cannot start with digit)
	if name != "" && name[0] >= '0' && name[0] <= '9' {
		name = "_" + name
	}

	return name
}

func toSnakeCase(s string) string {
	tokens := splitIdentifierTokens(s)
	if len(tokens) == 0 {
		return ""
	}

	var parts []string
	for _, t := range tokens {
		normalized := strings.ToLower(t)
		if normalized != "" {
			parts = append(parts, normalized)
		}
	}

	return strings.Join(parts, "_")
}

func splitIdentifierTokens(s string) []string {
	// Split on any non-ASCII letter/digit
	var raw []string
	var cur strings.Builder
	flush := func() {
		if cur.Len() == 0 {
			return
		}
		raw = append(raw, cur.String())
		cur.Reset()
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		isLetter := (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
		isDigit := c >= '0' && c <= '9'
		if isLetter || isDigit {
			cur.WriteByte(c)
			continue
		}
		flush()
	}
	flush()

	// Further split raw tokens on camelCase + digit boundaries
	var tokens []string
	for _, tok := range raw {
		tokens = append(tokens, splitCamelAndDigits(tok)...)
	}
	return tokens
}

func splitCamelAndDigits(s string) []string {
	if s == "" {
		return nil
	}

	var out []string
	start := 0
	isUpper := func(c byte) bool { return c >= 'A' && c <= 'Z' }
	isLower := func(c byte) bool { return c >= 'a' && c <= 'z' }
	isDigit := func(c byte) bool { return c >= '0' && c <= '9' }

	for i := 1; i < len(s); i++ {
		prev := s[i-1]
		cur := s[i]

		// letter<->digit boundary
		if (isDigit(cur) && !isDigit(prev)) || (!isDigit(cur) && isDigit(prev)) {
			out = append(out, s[start:i])
			start = i
			continue
		}

		// lower->upper boundary: userName
		if isLower(prev) && isUpper(cur) {
			out = append(out, s[start:i])
			start = i
			continue
		}

		// acronym boundary: URLValue -> URL + Value
		if i >= 2 {
			prevPrev := s[i-2]
			if isUpper(prevPrev) && isUpper(prev) && isLower(cur) {
				out = append(out, s[start:i-1])
				start = i - 1
				continue
			}
		}
	}

	out = append(out, s[start:])
	return out
}

func injectSchemasKeyDict(configContent string) string {
	openIdx := strings.Index(configContent, "{")
	if openIdx == -1 {
		return configContent
	}

	if len(configContent) < openIdx+2 {
		return "{'schemas': schemas}"
	}

	inner := configContent[openIdx+1 : len(configContent)-1]
	innerTrimmed := strings.TrimSpace(inner)

	// Check if schemas already present
	if strings.HasPrefix(innerTrimmed, "'schemas'") || strings.HasPrefix(innerTrimmed, `"schemas"`) {
		return configContent
	}

	if innerTrimmed == "" {
		return "{'schemas': schemas}"
	}
	return "{'schemas': schemas, " + inner + "}"
}

// ClientConfig represents Python client configuration for schema injection
type ClientConfig struct {
	Schemas json.RawMessage `json:"schemas,omitempty"`
}
