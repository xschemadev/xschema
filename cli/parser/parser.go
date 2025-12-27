package parser

import (
	"context"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/xschema/cli/language"
)

// Options configures parser behavior
type Options struct {
	Include *regexp.Regexp // files matching this are included
	Exclude *regexp.Regexp // files matching this are excluded
}

// queryCache caches compiled queries per language
var (
	queryCache         = make(map[string]*sitter.Query)
	queryCacheMu       sync.RWMutex
	importQueryCache   = make(map[string]*sitter.Query)
	importQueryCacheMu sync.RWMutex
)

func getQuery(lang *language.Language) (*sitter.Query, error) {
	queryCacheMu.RLock()
	if q, ok := queryCache[lang.Name]; ok {
		queryCacheMu.RUnlock()
		return q, nil
	}
	queryCacheMu.RUnlock()

	queryCacheMu.Lock()
	defer queryCacheMu.Unlock()

	// Double-check after acquiring write lock
	if q, ok := queryCache[lang.Name]; ok {
		return q, nil
	}

	q, err := sitter.NewQuery([]byte(lang.Query), lang.GetSitterLang())
	if err != nil {
		return nil, err
	}
	queryCache[lang.Name] = q
	return q, nil
}

func getImportQuery(lang *language.Language) (*sitter.Query, error) {
	if lang.ImportQuery == "" {
		return nil, nil
	}

	importQueryCacheMu.RLock()
	if q, ok := importQueryCache[lang.Name]; ok {
		importQueryCacheMu.RUnlock()
		return q, nil
	}
	importQueryCacheMu.RUnlock()

	importQueryCacheMu.Lock()
	defer importQueryCacheMu.Unlock()

	// Double-check after acquiring write lock
	if q, ok := importQueryCache[lang.Name]; ok {
		return q, nil
	}

	q, err := sitter.NewQuery([]byte(lang.ImportQuery), lang.GetSitterLang())
	if err != nil {
		return nil, err
	}
	importQueryCache[lang.Name] = q
	return q, nil
}

type AdapterRef struct {
	Name     string `json:"name"`
	Package  string `json:"package"`
	Language string `json:"language"`
}

type Declaration struct {
	Name     string              `json:"name"`
	Source   language.SourceType `json:"source"`
	Location string              `json:"location,omitempty"` // URL or file path
	Adapter  AdapterRef          `json:"adapter"`
	File     string              `json:"file"`
	Line     int                 `json:"line"`
}

// Parse finds all xschema declarations in the given directory
func Parse(ctx context.Context, dir string, opts Options) ([]Declaration, error) {
	files, err := getSourceFiles(ctx, dir)
	if err != nil {
		return nil, err
	}

	var decls []Declaration
	for _, path := range files {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// Apply include/exclude filters
		if opts.Exclude != nil && opts.Exclude.MatchString(path) {
			continue
		}
		if opts.Include != nil && !opts.Include.MatchString(path) {
			continue
		}

		ext := filepath.Ext(path)
		lang := language.ByExtension(ext)
		if lang == nil {
			continue
		}

		fileDecls, err := parseFile(ctx, path, lang)
		if err != nil {
			return nil, err
		}
		decls = append(decls, fileDecls...)
	}

	return decls, nil
}

// getSourceFiles returns source files using git ls-files if in a git repo,
// otherwise falls back to walking the directory
func getSourceFiles(ctx context.Context, dir string) ([]string, error) {
	args := append([]string{"ls-files", "--cached", "--others", "--exclude-standard"}, language.ExtensionGlobs()...)
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		// Not a git repo or git not available - fallback
		return walkDirFallback(ctx, dir)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return nil, nil
	}

	files := make([]string, 0, len(lines))
	for _, line := range lines {
		if line != "" {
			files = append(files, filepath.Join(dir, line))
		}
	}
	return files, nil
}

// walkDirFallback walks directory manually when git is not available
func walkDirFallback(ctx context.Context, dir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			// Not exaustive but fallback if not git ->
			// if not git probably doesn't have much boilerplate
			if name == "node_modules" || name == ".git" || name == "__pycache__" || name == ".venv" || name == "venv" {
				return filepath.SkipDir
			}
			return nil
		}

		if language.ByExtension(filepath.Ext(path)) != nil {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func parseFile(ctx context.Context, path string, lang *language.Language) ([]Declaration, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	parser := sitter.NewParser()
	parser.SetLanguage(lang.GetSitterLang())

	tree, err := parser.ParseCtx(ctx, nil, content)
	if err != nil {
		return nil, err
	}

	importMap := parseImports(tree, content, lang)

	q, err := getQuery(lang)
	if err != nil {
		return nil, err
	}

	qc := sitter.NewQueryCursor()
	qc.Exec(q, tree.RootNode())

	return extractDeclarations(qc, q, content, path, lang, importMap)
}

func parseImports(tree *sitter.Tree, content []byte, lang *language.Language) map[string]string {
	q, err := getImportQuery(lang)
	if err != nil || q == nil {
		return nil
	}

	importMap := make(map[string]string)
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

		var importSource string
		var importedName string

		for _, cap := range match.Captures {
			capName := q.CaptureNameForId(cap.Index)
			text := cap.Node.Content(content)

			switch capName {
			case "package":
				importSource = unquoteString(text)
			case "imported_name":
				importedName = text
			}
		}

		if importSource != "" && importedName != "" {
			importMap[importedName] = importSource
		}
	}

	return importMap
}

func extractDeclarations(qc *sitter.QueryCursor, q *sitter.Query, content []byte, path string, lang *language.Language, importMap map[string]string) ([]Declaration, error) {
	var decls []Declaration

	for {
		match, ok := qc.NextMatch()
		if !ok {
			break
		}

		match = qc.FilterPredicates(match, content)
		if len(match.Captures) == 0 {
			continue
		}

		var method, name, source, adapter string
		var sourceLine int

		for _, cap := range match.Captures {
			capName := q.CaptureNameForId(cap.Index)
			text := cap.Node.Content(content)

			switch capName {
			case "method":
				method = text
				sourceLine = int(cap.Node.StartPoint().Row) + 1
			case "name":
				name = unquoteString(text)
			case "source":
				source = text
			case "adapter":
				adapter = text
			}
		}

		// Look up source type from language's method mapping
		sourceType, ok := lang.MethodMapping[method]
		if !ok {
			continue
		}

		decl := Declaration{
			Name:     name,
			Source:   sourceType,
			Location: unquoteString(source),
			Adapter: AdapterRef{
				Name:     adapter,
				Package:  importMap[adapter],
				Language: lang.Name,
			},
			File: path,
			Line: sourceLine,
		}

		decls = append(decls, decl)
	}

	return decls, nil
}

// unquoteString removes surrounding quotes from a string literal
// Handles: "...", '...', `...`, """...""", ”'...”', r"...", r'...'
func unquoteString(s string) string {
	if len(s) < 2 {
		return s
	}

	// Strip raw string prefix (r"..." or r'...')
	if len(s) >= 3 && (s[0] == 'r' || s[0] == 'R') && (s[1] == '"' || s[1] == '\'') {
		s = s[1:]
	}

	// Handle triple quotes first (""" or ''')
	if len(s) >= 6 {
		if (s[:3] == `"""` && s[len(s)-3:] == `"""`) ||
			(s[:3] == `'''` && s[len(s)-3:] == `'''`) {
			return s[3 : len(s)-3]
		}
	}

	// Handle single, double, or backtick quotes
	if (s[0] == '"' || s[0] == '\'' || s[0] == '`') && s[0] == s[len(s)-1] {
		return s[1 : len(s)-1]
	}
	return s
}
