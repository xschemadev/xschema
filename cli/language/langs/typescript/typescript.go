package typescript

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
		panic(fmt.Sprintf("language typescript register failed: %v", err))
	}
}

func Language() language.Language {
	return language.Language{
		Name:                 "typescript",
		Extensions:           []string{".ts", ".tsx", ".js", ".jsx"},
		SchemaURL:            language.XSchemaBaseURL + "typescript.jsonc",
		SchemaExt:            "typescript.jsonc",
		AdapterBinPrefix:     adapterBinPrefix,
		DetectRunner:         detectRunner,
		AdapterInvoker:       adapterInvoker{},
		BuildSchemasImport:   buildSchemasImport,
		ImportPattern:        `(?m)^import\s+.*$`,
		InjectSchemasKey:     injectSchemasKeyBrace,
		ClientFactoryPattern: `createXSchemaClient\s*\(\s*(\{[^}]*\})\s*\)`,
		OutputDir:            ".xschema",
		OutputFiles: []language.OutputFileSpec{{
			Path:     "xschema.gen.ts",
			Template: outputTemplate,
		}},
		MergeImports:        MergeImports,
		BuildVarName:        buildVarName,
		IgnoreDirs:          []string{"node_modules", "dist", "build"},
		DetectHarnessRunner: detectHarnessRunner,
		GetPackageName:      getPackageName,
		AdapterCLIPath:      getAdapterCLIPath,
		HarnessExtension:    ".ts",
		HarnessTemplate:     harnessTemplate,
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

	// Migration help for legacy adapter names like "zod".
	if !strings.HasPrefix(adapterRef, "@") && !strings.Contains(adapterRef, "/") {
		return language.CommandSpec{}, fmt.Errorf(
			"invalid typescript adapter ref %q: expected scoped npm package ref like %q (migration: change adapter to %q)",
			adapterRef,
			"@xschemadev/zod",
			"@xschemadev/"+adapterRef,
		)
	}

	if !strings.HasPrefix(adapterRef, "@xschemadev/") {
		return language.CommandSpec{}, fmt.Errorf(
			"invalid typescript adapter ref %q: expected %q scope (example: %q)",
			adapterRef,
			"@xschemadev",
			"@xschemadev/zod",
		)
	}

	parts := strings.Split(adapterRef, "/")
	if len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
		return language.CommandSpec{}, fmt.Errorf("invalid typescript adapter ref %q: expected format %q", adapterRef, "@xschemadev/<adapter>")
	}

	pkgName := parts[1]
	binName := adapterBinPrefix + pkgName

	runner, runnerArgs, err := detectRunnerInDir(projectRoot)
	if err != nil {
		return language.CommandSpec{}, fmt.Errorf("failed to detect typescript runner: %w", err)
	}

	return language.CommandSpec{
		Cmd:  runner,
		Args: append(runnerArgs, binName),
		Dir:  projectRoot,
	}, nil
}

func buildSchemasImport(importPath string) string {
	return `import { schemas } from "` + importPath + `";`
}

func getPackageName(dir string) string {
	pkgPath := filepath.Join(dir, "package.json")
	data, err := os.ReadFile(pkgPath)
	if err != nil {
		return filepath.Base(dir)
	}

	var pkg struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return filepath.Base(dir)
	}

	if pkg.Name != "" {
		return pkg.Name
	}
	return filepath.Base(dir)
}

func getAdapterCLIPath(adapterPath string) string {
	return filepath.Join(adapterPath, "dist", "cli.js")
}

var tsReservedWords = map[string]bool{
	"break":       true,
	"case":        true,
	"catch":       true,
	"class":       true,
	"const":       true,
	"continue":    true,
	"debugger":    true,
	"default":     true,
	"delete":      true,
	"do":          true,
	"else":        true,
	"enum":        true,
	"export":      true,
	"extends":     true,
	"false":       true,
	"finally":     true,
	"for":         true,
	"function":    true,
	"if":          true,
	"import":      true,
	"in":          true,
	"instanceof":  true,
	"new":         true,
	"null":        true,
	"return":      true,
	"super":       true,
	"switch":      true,
	"this":        true,
	"throw":       true,
	"true":        true,
	"try":         true,
	"typeof":      true,
	"var":         true,
	"void":        true,
	"while":       true,
	"with":        true,
	"as":          true,
	"implements":  true,
	"interface":   true,
	"let":         true,
	"package":     true,
	"private":     true,
	"protected":   true,
	"public":      true,
	"static":      true,
	"yield":       true,
	"any":         true,
	"boolean":     true,
	"constructor": true,
	"declare":     true,
	"get":         true,
	"module":      true,
	"require":     true,
	"number":      true,
	"set":         true,
	"string":      true,
	"symbol":      true,
	"type":        true,
	"from":        true,
	"of":          true,
}

func buildVarName(namespace, id string) string {
	ns := toLowerCamelTSIdent(namespace)
	if ns == "" {
		ns = "schema"
	}

	name := ns + "_" + toPascalTSIdent(id)
	if name == "_" {
		name = "schema"
	}

	if tsReservedWords[strings.ToLower(name)] {
		name = "_" + name
	}

	if name != "" {
		first := name[0]
		if !(first == '_' || first == '$' || (first >= 'A' && first <= 'Z') || (first >= 'a' && first <= 'z')) {
			name = "_" + name
		}
	}
	return name
}

func toLowerCamelTSIdent(s string) string {
	tokens := splitIdentifierTokens(s)
	if len(tokens) == 0 {
		return ""
	}

	first := normalizeToken(tokens[0], true)
	if first == "" {
		return ""
	}

	var b strings.Builder
	b.WriteString(strings.ToLower(first))

	for _, t := range tokens[1:] {
		part := normalizeToken(t, false)
		if part == "" {
			continue
		}
		b.WriteString(part)
	}
	return b.String()
}

func toPascalTSIdent(s string) string {
	tokens := splitIdentifierTokens(s)
	if len(tokens) == 0 {
		return "Schema"
	}
	var b strings.Builder
	for _, t := range tokens {
		part := normalizeToken(t, false)
		if part == "" {
			continue
		}
		b.WriteString(part)
	}
	out := b.String()
	if out == "" {
		return "Schema"
	}
	return out
}

func splitIdentifierTokens(s string) []string {
	// Split on any non-ASCII letter/digit.
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

	// Further split raw tokens on camelCase + digit boundaries.
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

func normalizeToken(token string, lower bool) string {
	if token == "" {
		return ""
	}

	allDigits := true
	allUpper := true
	for i := 0; i < len(token); i++ {
		c := token[i]
		if c < '0' || c > '9' {
			allDigits = false
		}
		if c >= 'a' && c <= 'z' {
			allUpper = false
		}
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			continue
		}
		return ""
	}

	if allDigits {
		return token
	}

	if allUpper {
		if lower {
			return strings.ToLower(token)
		}
		return token
	}

	if lower {
		return strings.ToLower(token[:1]) + token[1:]
	}

	return strings.ToUpper(token[:1]) + strings.ToLower(token[1:])
}

func injectSchemasKeyBrace(configContent string) string {
	openIdx := strings.Index(configContent, "{")
	if openIdx == -1 {
		return configContent
	}

	if len(configContent) < openIdx+2 {
		return "{ schemas }"
	}

	inner := configContent[openIdx+1 : len(configContent)-1]
	innerTrimmed := strings.TrimSpace(inner)

	if strings.HasPrefix(innerTrimmed, "schemas") && (len(innerTrimmed) == 7 || strings.HasPrefix(innerTrimmed[7:], ",") || strings.HasPrefix(innerTrimmed[7:], "}")) {
		return configContent
	}

	if strings.HasPrefix(innerTrimmed, "schemas:") {
		return configContent
	}

	if strings.HasPrefix(innerTrimmed, `"schemas":`) || strings.HasPrefix(innerTrimmed, `'schemas':`) {
		return configContent
	}

	if innerTrimmed == "" {
		return "{ schemas }"
	}
	return "{ schemas, " + inner + " }"
}
