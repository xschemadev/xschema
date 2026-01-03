package injector

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/xschema/cli/generator"
	"github.com/xschema/cli/language"
	"github.com/xschema/cli/ui"
)

type InjectInput struct {
	Language string                     `json:"language"` // typescript, python, go
	Outputs  []generator.GenerateOutput `json:"outputs"`
	OutDir   string                     `json:"outDir"` // default .xschema
}

// TemplateData is passed to the language template
type TemplateData struct {
	Imports string                 // merged imports
	Schemas []language.SchemaEntry // individual schema entries
	Header  string                 // language-specific header (e.g., Go package decl)
	Footer  string                 // language-specific footer
}

// Inject writes generated code to output directory
func Inject(input InjectInput) error {
	lang := language.ByName(input.Language)
	if lang == nil {
		ui.Verbosef("unsupported language: %s", input.Language)
		return fmt.Errorf("unsupported language: %s", input.Language)
	}

	if lang.Template == "" {
		ui.Verbosef("no template defined for language: %s", input.Language)
		return fmt.Errorf("no template defined for language: %s", input.Language)
	}

	ui.Verbosef("injecting schemas: language=%s, outputs=%d, outDir=%s", input.Language, len(input.Outputs), input.OutDir)

	// Build template data
	data := buildTemplateData(input, lang)

	ui.Verbosef("template data: imports=%d, schemas=%d", len(data.Imports), len(data.Schemas))

	// Parse and execute template
	tmpl, err := template.New("inject").Parse(lang.Template)
	if err != nil {
		ui.Verbosef("failed to parse template for language: %s", input.Language)
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		ui.Verbosef("failed to execute template for language: %s", input.Language)
		return fmt.Errorf("failed to execute template: %w", err)
	}

	ui.Verbosef("template execution successful: %d bytes", buf.Len())

	// Ensure output directory exists
	if err := os.MkdirAll(input.OutDir, 0755); err != nil {
		ui.Verbosef("failed to create output directory: %s", input.OutDir)
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write output file
	outPath := filepath.Join(input.OutDir, lang.OutputFile)
	if err := os.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
		ui.Verbosef("failed to write output file: %s", outPath)
		return fmt.Errorf("failed to write output file: %w", err)
	}

	ui.Verbosef("successfully injected schemas: path=%s, bytes=%d", outPath, buf.Len())
	return nil
}

func buildTemplateData(input InjectInput, lang *language.Language) TemplateData {
	// Collect all imports
	var allImports []string
	for _, out := range input.Outputs {
		allImports = append(allImports, out.Imports...)
	}

	// Merge imports using language-specific function
	var mergedImports string
	if lang.MergeImports != nil {
		mergedImports = lang.MergeImports(allImports)
	}

	// Build schema entries
	schemas := make([]language.SchemaEntry, len(input.Outputs))
	for i, out := range input.Outputs {
		schemas[i] = language.SchemaEntry{
			Name: out.Name,
			Code: out.Schema,
			Type: out.Type,
		}
	}

	// Build header/footer
	var header, footer string
	if lang.BuildHeader != nil {
		header = lang.BuildHeader(input.OutDir, schemas)
	}
	if lang.BuildFooter != nil {
		footer = lang.BuildFooter(input.OutDir, schemas)
	}

	return TemplateData{
		Imports: mergedImports,
		Schemas: schemas,
		Header:  header,
		Footer:  footer,
	}
}

// InjectClientInput holds info needed to inject schemas import into client file
type InjectClientInput struct {
	ClientFile string             // path to client file
	Language   *language.Language // language config
	OutDir     string             // output directory (e.g., ".xschema")
}

// InjectClient adds schemas import and schemas key to createXSchemaClient call
func InjectClient(ctx context.Context, input InjectClientInput) error {
	content, err := os.ReadFile(input.ClientFile)
	if err != nil {
		return fmt.Errorf("failed to read client file: %w", err)
	}

	lang := input.Language

	// Parse with tree-sitter to find createXSchemaClient call position
	parser := sitter.NewParser()
	parser.SetLanguage(lang.GetSitterLang())

	tree, err := parser.ParseCtx(ctx, nil, content)
	if err != nil {
		return fmt.Errorf("failed to parse client file: %w", err)
	}

	// Find the createXSchemaClient call and config object
	callInfo, err := findClientCall(tree, content, lang)
	if err != nil {
		return err
	}
	if callInfo == nil {
		return fmt.Errorf("no %s call found", lang.ClientFactory)
	}

	// Build import path: use base of OutDir for relative import
	// Import needs ./ prefix for relative imports
	relOutDir := filepath.Base(input.OutDir)
	importPath := "./" + relOutDir + "/" + strings.TrimSuffix(lang.OutputFile, filepath.Ext(lang.OutputFile))

	// Modify content
	modified := string(content)

	// 1. Add/update schemas in config object
	modified, err = injectSchemasKey(modified, callInfo, lang)
	if err != nil {
		return err
	}

	// 2. Add import if not present
	modified = injectSchemasImport(modified, importPath, lang)

	// Write back
	if err := os.WriteFile(input.ClientFile, []byte(modified), 0644); err != nil {
		return fmt.Errorf("failed to write client file: %w", err)
	}

	ui.Verbosef("injected schemas into client: %s", input.ClientFile)
	return nil
}

type clientCallInfo struct {
	configStart uint32 // byte offset of config object start
	configEnd   uint32 // byte offset of config object end
	hasSchemas  bool   // whether schemas key already exists
}

func findClientCall(tree *sitter.Tree, content []byte, lang *language.Language) (*clientCallInfo, error) {
	// Query for createXSchemaClient call with config object
	queryStr := lang.ClientCallQuery
	if queryStr == "" {
		return nil, nil
	}

	q, err := sitter.NewQuery([]byte(queryStr), lang.GetSitterLang())
	if err != nil {
		return nil, fmt.Errorf("failed to compile query: %w", err)
	}

	qc := sitter.NewQueryCursor()
	qc.Exec(q, tree.RootNode())

	for {
		match, ok := qc.NextMatch()
		if !ok {
			break
		}

		match = qc.FilterPredicates(match, content)
		if len(match.Captures) == 0 {
			continue
		}

		info := &clientCallInfo{}
		for _, cap := range match.Captures {
			capName := q.CaptureNameForId(cap.Index)
			switch capName {
			case "config":
				info.configStart = cap.Node.StartByte()
				info.configEnd = cap.Node.EndByte()
			case "schemas_key":
				info.hasSchemas = true
			}
		}

		if info.configStart > 0 {
			return info, nil
		}
	}

	return nil, nil
}

func injectSchemasKey(content string, info *clientCallInfo, lang *language.Language) (string, error) {
	if info.hasSchemas {
		// Already has schemas key, assume it's correct
		return content, nil
	}

	if lang.InjectSchemasKey == nil {
		return content, fmt.Errorf("no InjectSchemasKey defined for language: %s", lang.Name)
	}

	configContent := content[info.configStart:info.configEnd]
	newConfig := lang.InjectSchemasKey(configContent)

	return content[:info.configStart] + newConfig + content[info.configEnd:], nil
}

func injectSchemasImport(content, importPath string, lang *language.Language) string {
	if lang.BuildSchemasImport == nil {
		return content
	}

	// Check if import already exists (with or without ./ prefix)
	normalizedPath := strings.TrimPrefix(importPath, "./")
	if strings.Contains(content, importPath) || strings.Contains(content, normalizedPath) {
		return content
	}

	importLine := lang.BuildSchemasImport(importPath)
	if importLine == "" {
		return content
	}

	// Find last import statement using language-specific pattern
	pattern := lang.ImportPattern
	if pattern == "" {
		// No pattern, add at top
		return importLine + "\n" + content
	}

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringIndex(content, -1)

	if len(matches) == 0 {
		// No imports, add at top
		return importLine + "\n" + content
	}

	// Insert after last import
	lastMatch := matches[len(matches)-1]
	insertPos := lastMatch[1]

	return content[:insertPos] + "\n" + importLine + content[insertPos:]
}
