package injector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/xschemadev/xschema/adapter"
	"github.com/xschemadev/xschema/language"
	"github.com/xschemadev/xschema/ui"
)

type InjectInput struct {
	Language string                  `json:"language"` // typescript, python, go
	Outputs  []adapter.ConvertResult `json:"outputs"`
	OutDir   string                  `json:"outDir"` // default .xschema
}

// TemplateData is passed to the language template
type TemplateData struct {
	Imports string                 // merged imports
	Schemas []language.SchemaEntry // individual schema entries
	Header  string                 // language-specific header (e.g., Go package decl)
	Footer  string                 // language-specific footer
}

const manifestFileName = "xschema.manifest.json"

type manifest struct {
	Files []string `json:"files"`
}

// WriteGeneratedFiles writes generated files under the output root.
// Each GeneratedFile.Path is interpreted as a relative path.
//
// It also writes a manifest (xschema.manifest.json) and removes stale files from the previous manifest.
func WriteGeneratedFiles(outDir string, files []language.GeneratedFile) error {
	if strings.TrimSpace(outDir) == "" {
		return fmt.Errorf("outDir is required")
	}
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	normalized := make([]language.GeneratedFile, len(files))
	for i, file := range files {
		path, err := language.NormalizeRelativePath(file.Path)
		if err != nil {
			return fmt.Errorf("invalid generated file path %q: %w", file.Path, err)
		}
		normalized[i] = language.GeneratedFile{Path: path, Contents: file.Contents}
	}

	language.SortGeneratedFiles(normalized)
	if err := language.ValidateGeneratedFiles(normalized); err != nil {
		return fmt.Errorf("invalid generated files: %w", err)
	}

	newPathsOrdered := make([]string, len(normalized))
	newPathSet := make(map[string]struct{}, len(normalized))
	for i, file := range normalized {
		newPathsOrdered[i] = file.Path
		newPathSet[file.Path] = struct{}{}
	}

	manifestPath := filepath.Join(outDir, manifestFileName)
	previousPaths, err := loadManifest(manifestPath)
	if err != nil {
		return err
	}

	var stalePaths []string
	for _, previousPath := range previousPaths {
		if _, ok := newPathSet[previousPath]; ok {
			continue
		}
		stalePaths = append(stalePaths, previousPath)
	}

	if err := deleteStaleFiles(outDir, stalePaths); err != nil {
		return err
	}

	for _, file := range normalized {
		outPath := filepath.Join(outDir, filepath.FromSlash(file.Path))
		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
		if err := os.WriteFile(outPath, []byte(file.Contents), 0644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
	}

	if err := writeManifestAtomic(outDir, newPathsOrdered); err != nil {
		return err
	}

	return nil
}

func loadManifest(manifestPath string) ([]string, error) {
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read manifest %q: %w", manifestPath, err)
	}

	var m manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to parse manifest %q: %w", manifestPath, err)
	}

	seen := make(map[string]struct{}, len(m.Files))
	paths := make([]string, 0, len(m.Files))
	for _, path := range m.Files {
		if strings.TrimSpace(path) == "" {
			return nil, fmt.Errorf("manifest %q contains empty path", manifestPath)
		}
		if path == manifestFileName {
			continue
		}
		normalized, err := language.NormalizeRelativePath(path)
		if err != nil {
			return nil, fmt.Errorf("manifest %q contains invalid path %q: %w", manifestPath, path, err)
		}
		if _, ok := seen[normalized]; ok {
			continue
		}
		seen[normalized] = struct{}{}
		paths = append(paths, normalized)
	}

	sort.Strings(paths)
	return paths, nil
}

func writeManifestAtomic(outDir string, paths []string) error {
	manifestPath := filepath.Join(outDir, manifestFileName)
	m := manifest{Files: paths}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal manifest: %w", err)
	}
	data = append(data, '\n')

	tmp, err := os.CreateTemp(outDir, manifestFileName+".tmp-*")
	if err != nil {
		return fmt.Errorf("failed to create manifest temp file: %w", err)
	}
	tmpName := tmp.Name()

	writeErr := func() error {
		if _, err := tmp.Write(data); err != nil {
			return fmt.Errorf("failed to write manifest temp file: %w", err)
		}
		if err := tmp.Close(); err != nil {
			return fmt.Errorf("failed to close manifest temp file: %w", err)
		}
		if err := os.Chmod(tmpName, 0644); err != nil {
			return fmt.Errorf("failed to set manifest permissions: %w", err)
		}
		if err := os.Rename(tmpName, manifestPath); err != nil {
			return fmt.Errorf("failed to write manifest %q: %w", manifestPath, err)
		}
		return nil
	}()
	if writeErr != nil {
		_ = tmp.Close()
		_ = os.Remove(tmpName)
		return writeErr
	}

	return nil
}

func deleteStaleFiles(outDir string, paths []string) error {
	if len(paths) == 0 {
		return nil
	}

	absOutDir, err := filepath.Abs(outDir)
	if err != nil {
		return fmt.Errorf("failed to resolve output directory: %w", err)
	}

	for _, path := range paths {
		normalized, err := language.NormalizeRelativePath(path)
		if err != nil {
			return fmt.Errorf("invalid stale path %q: %w", path, err)
		}

		absPath := filepath.Join(absOutDir, filepath.FromSlash(normalized))
		rel, err := filepath.Rel(absOutDir, absPath)
		if err != nil {
			return fmt.Errorf("failed to resolve stale path %q: %w", normalized, err)
		}
		if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			return fmt.Errorf("refusing to delete stale path outside output directory: %q", normalized)
		}

		if err := os.Remove(absPath); err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return fmt.Errorf("failed to remove stale file %q: %w", normalized, err)
		}

		cleanupEmptyParents(absOutDir, filepath.Dir(absPath))
	}

	return nil
}

func cleanupEmptyParents(outRoot, dir string) {
	for {
		if dir == outRoot {
			return
		}
		if len(dir) < len(outRoot) {
			return
		}

		err := os.Remove(dir)
		if err != nil {
			if os.IsNotExist(err) {
				dir = filepath.Dir(dir)
				continue
			}
			return
		}
		dir = filepath.Dir(dir)
	}
}

// Inject writes generated code to output directory.
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

	files := []language.GeneratedFile{{
		Path:     lang.OutputFile,
		Contents: buf.String(),
	}}

	if err := WriteGeneratedFiles(input.OutDir, files); err != nil {
		ui.Verbosef("failed to write generated files: outDir=%s", input.OutDir)
		return err
	}

	ui.Verbosef("successfully injected schemas: files=%d, outDir=%s", len(files), input.OutDir)
	return nil
}

func buildTemplateData(input InjectInput, lang *language.Language) TemplateData {
	// Collect all imports
	var allImports []string
	for _, out := range input.Outputs {
		allImports = append(allImports, out.Imports...)
	}

	mergedImports := lang.MergeImports(allImports)

	// Build schema entries
	schemas := make([]language.SchemaEntry, len(input.Outputs))
	for i, out := range input.Outputs {
		varName := out.Namespace + "_" + out.ID // default
		if lang.BuildVarName != nil {
			varName = lang.BuildVarName(out.Namespace, out.ID)
		}
		schemas[i] = language.SchemaEntry{
			Namespace: out.Namespace,
			ID:        out.ID,
			VarName:   varName,
			Code:      out.Schema,
			Type:      out.Type,
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

// InjectClient adds schemas import to a client file
// This is a simplified version that doesn't use tree-sitter
// It looks for createXSchemaClient({ and injects schemas
func InjectClient(input InjectClientInput) error {
	content, err := os.ReadFile(input.ClientFile)
	if err != nil {
		return fmt.Errorf("failed to read client file: %w", err)
	}

	lang := input.Language
	modified := string(content)

	// Build import path: use base of OutDir for relative import
	relOutDir := filepath.Base(input.OutDir)
	importPath := "./" + relOutDir + "/" + strings.TrimSuffix(lang.OutputFile, filepath.Ext(lang.OutputFile))

	// 1. Try to inject schemas into config object using regex
	// Match: createXSchemaClient({ ... })
	// This is a simplified approach that works for common cases
	modified, injected := injectSchemasIntoConfig(modified, lang)

	if !injected {
		ui.Verbosef("could not find config object to inject schemas - manual injection may be needed")
	}

	// 2. Add import if not present
	modified = injectSchemasImport(modified, importPath, lang)

	// Write back only if modified
	if modified != string(content) {
		if err := os.WriteFile(input.ClientFile, []byte(modified), 0644); err != nil {
			return fmt.Errorf("failed to write client file: %w", err)
		}
		ui.Verbosef("injected schemas into client: %s", input.ClientFile)
	}

	return nil
}

// injectSchemasIntoConfig tries to inject "schemas" into the config object
// Returns the modified content and whether injection was successful
func injectSchemasIntoConfig(content string, lang *language.Language) (string, bool) {
	if lang.InjectSchemasKey == nil || lang.ClientFactoryPattern == "" {
		return content, false
	}

	re := regexp.MustCompile(lang.ClientFactoryPattern)
	matches := re.FindStringSubmatchIndex(content)
	if matches == nil {
		return content, false
	}

	// Extract config object
	configStart := matches[2]
	configEnd := matches[3]
	configContent := content[configStart:configEnd]

	// Inject schemas key
	newConfig := lang.InjectSchemasKey(configContent)
	if newConfig == configContent {
		// Already has schemas or couldn't inject
		return content, true
	}

	// Replace in content
	return content[:configStart] + newConfig + content[configEnd:], true
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
