package injector

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/xschema/cli/generator"
	"github.com/xschema/cli/language"
	"github.com/xschema/cli/logger"
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
		logger.Error("unsupported language", "language", input.Language)
		return fmt.Errorf("unsupported language: %s", input.Language)
	}

	if lang.Template == "" {
		logger.Error("no template defined", "language", input.Language)
		return fmt.Errorf("no template defined for language: %s", input.Language)
	}

	logger.Info("injecting schemas", "language", input.Language, "outputs", len(input.Outputs), "outDir", input.OutDir)

	// Build template data
	data := buildTemplateData(input, lang)

	logger.Debug("template data", "imports", len(data.Imports), "schemas", len(data.Schemas))

	// Parse and execute template
	tmpl, err := template.New("inject").Parse(lang.Template)
	if err != nil {
		logger.Error("failed to parse template", "language", input.Language, "error", err)
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		logger.Error("failed to execute template", "language", input.Language, "error", err)
		return fmt.Errorf("failed to execute template: %w", err)
	}

	logger.Debug("template execution successful", "bytes", buf.Len())

	// Ensure output directory exists
	if err := os.MkdirAll(input.OutDir, 0755); err != nil {
		logger.Error("failed to create output directory", "outDir", input.OutDir, "error", err)
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write output file
	outPath := filepath.Join(input.OutDir, lang.OutputFile)
	if err := os.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
		logger.Error("failed to write output file", "path", outPath, "error", err)
		return fmt.Errorf("failed to write output file: %w", err)
	}

	logger.Info("successfully injected schemas", "path", outPath, "bytes", buf.Len())
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
