package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xschema/cli/generator"
	"github.com/xschema/cli/injector"
	"github.com/xschema/cli/language"
	"github.com/xschema/cli/parser"
	"github.com/xschema/cli/retriever"
)

var cfg Config

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Parse codebase, convert schemas, output native validators",
	RunE:  runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Directories
	generateCmd.Flags().StringVarP(&cfg.InputDir, "input", "i", ".", "input directory to parse")
	generateCmd.Flags().StringVarP(&cfg.OutputDir, "output", "o", ".xschema", "output directory")

	// HTTP/Retriever
	generateCmd.Flags().IntVarP(&cfg.Concurrency, "concurrency", "c", 10, "max concurrent HTTP requests")
	generateCmd.Flags().DurationVar(&cfg.HTTPTimeout, "http-timeout", retriever.DefaultOptions().HTTPTimeout, "HTTP request timeout")
	generateCmd.Flags().IntVar(&cfg.Retries, "retries", 3, "max retry attempts for failed requests")
	generateCmd.Flags().BoolVar(&cfg.NoCache, "no-cache", false, "disable schema caching")

	// Filtering
	generateCmd.Flags().StringVar(&cfg.Include, "include", "", "regex pattern for files to include")
	generateCmd.Flags().StringVar(&cfg.Exclude, "exclude", "", "regex pattern for files to exclude")
	generateCmd.Flags().StringVar(&cfg.Adapter, "adapter", "", "only process specific adapter")

	// Output behavior
	generateCmd.Flags().BoolVar(&cfg.DryRun, "dry-run", false, "show what would be generated without writing")
	generateCmd.Flags().BoolVar(&cfg.Force, "force", false, "overwrite existing files without prompt")
	generateCmd.Flags().BoolVarP(&cfg.Verbose, "verbose", "v", false, "verbose output")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// Compile include/exclude regexes
	parserOpts, err := buildParserOpts()
	if err != nil {
		return err
	}

	// 1. Parse
	if cfg.Verbose {
		fmt.Fprintf(os.Stderr, "parsing %s...\n", cfg.InputDir)
	}
	decls, err := parser.Parse(ctx, cfg.InputDir, parserOpts)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	if cfg.Verbose {
		fmt.Fprintf(os.Stderr, "found %d declarations\n", len(decls))
	}

	// If no declarations found, generate stub file
	if len(decls) == 0 {
		return generateStub()
	}

	// 2. Retrieve
	retrieverOpts := retriever.Options{
		Concurrency: cfg.Concurrency,
		HTTPTimeout: cfg.HTTPTimeout,
		Retries:     cfg.Retries,
		NoCache:     cfg.NoCache,
	}
	if cfg.Verbose {
		fmt.Fprintf(os.Stderr, "retrieving schemas (concurrency=%d)...\n", cfg.Concurrency)
	}
	batches, err := retriever.Retrieve(ctx, decls, retrieverOpts)
	if err != nil {
		return fmt.Errorf("retrieve: %w", err)
	}

	// Filter by adapter if specified
	if cfg.Adapter != "" {
		batches = filterBatchesByAdapter(batches, cfg.Adapter)
		if len(batches) == 0 {
			return fmt.Errorf("no schemas found for adapter %q", cfg.Adapter)
		}
	}

	// 3. Generate
	outputsByLang := make(map[string][]generator.GenerateOutput)
	for _, batch := range batches {
		if cfg.Verbose {
			fmt.Fprintf(os.Stderr, "generating %s schemas via %s...\n", batch.Language, batch.Adapter)
		}

		if cfg.DryRun {
			printDryRun(batch)
			continue
		}

		outputs, err := generator.Generate(ctx, batch)
		if err != nil {
			return fmt.Errorf("generate (%s): %w", batch.Adapter, err)
		}
		outputsByLang[batch.Language] = append(outputsByLang[batch.Language], outputs...)
	}

	if cfg.DryRun {
		return nil
	}

	// 4. Inject
	for lang, outputs := range outputsByLang {
		if cfg.Verbose {
			fmt.Fprintf(os.Stderr, "injecting %d %s schemas to %s...\n", len(outputs), lang, cfg.OutputDir)
		}

		err := injector.Inject(injector.InjectInput{
			Language: lang,
			Outputs:  outputs,
			OutDir:   cfg.OutputDir,
		})
		if err != nil {
			return fmt.Errorf("inject (%s): %w", lang, err)
		}
	}

	if cfg.Verbose {
		fmt.Fprintln(os.Stderr, "done")
	}
	return nil
}

// generateStub creates an empty schema file for the target language
func generateStub() error {
	lang := detectLanguage()
	if lang == "" {
		fmt.Fprintln(os.Stderr, "no xschema declarations found")
		return nil
	}

	if cfg.DryRun {
		fmt.Printf("would generate empty %s stub in %s\n", lang, cfg.OutputDir)
		return nil
	}

	if cfg.Verbose {
		fmt.Fprintf(os.Stderr, "generating empty %s stub...\n", lang)
	}

	err := injector.Inject(injector.InjectInput{
		Language: lang,
		Outputs:  []generator.GenerateOutput{}, // empty
		OutDir:   cfg.OutputDir,
	})
	if err != nil {
		return fmt.Errorf("inject stub (%s): %w", lang, err)
	}

	fmt.Fprintf(os.Stderr, "created %s/%s\n", cfg.OutputDir, language.ByName(lang).OutputFile)
	return nil
}

// detectLanguage tries to detect the project language from common files
func detectLanguage() string {
	// Check for TypeScript/JavaScript
	for _, f := range []string{"package.json", "tsconfig.json", "bun.lockb", "package-lock.json"} {
		if _, err := os.Stat(f); err == nil {
			return "typescript"
		}
	}
	// Check for Python
	for _, f := range []string{"pyproject.toml", "setup.py", "requirements.txt", "Pipfile"} {
		if _, err := os.Stat(f); err == nil {
			return "python"
		}
	}
	// Check for Go
	if _, err := os.Stat("go.mod"); err == nil {
		return "go"
	}
	return ""
}

func buildParserOpts() (parser.Options, error) {
	var opts parser.Options
	var err error

	if cfg.Include != "" {
		opts.Include, err = regexp.Compile(cfg.Include)
		if err != nil {
			return opts, fmt.Errorf("invalid --include regex: %w", err)
		}
	}
	if cfg.Exclude != "" {
		opts.Exclude, err = regexp.Compile(cfg.Exclude)
		if err != nil {
			return opts, fmt.Errorf("invalid --exclude regex: %w", err)
		}
	}
	return opts, nil
}

func filterBatchesByAdapter(batches []generator.GenerateBatchInput, adapter string) []generator.GenerateBatchInput {
	var filtered []generator.GenerateBatchInput
	for _, b := range batches {
		if strings.Contains(b.Adapter, adapter) {
			filtered = append(filtered, b)
		}
	}
	return filtered
}

func printDryRun(batch generator.GenerateBatchInput) {
	fmt.Printf("adapter: %s (%s)\n", batch.Adapter, batch.Language)
	for _, s := range batch.Schemas {
		var schema map[string]any
		json.Unmarshal(s.Schema, &schema)
		fmt.Printf("  - %s: %v\n", s.Name, schema["type"])
	}
}
