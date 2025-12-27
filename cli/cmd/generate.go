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
	"github.com/xschema/cli/logger"
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
	logger.SetLogger(logger.New(cfg.Verbose))

	ctx := cmd.Context()

	// Compile include/exclude regexes
	parserOpts, err := buildParserOpts()
	if err != nil {
		return err
	}

	// 1. Parse
	logger.Info("parsing directory", "input", cfg.InputDir)
	decls, err := parser.Parse(ctx, cfg.InputDir, parserOpts)
	if err != nil {
		logger.Error("parse failed", "error", err)
		return fmt.Errorf("parse: %w", err)
	}

	logger.Info("found declarations", "count", len(decls))
	if len(decls) == 0 {
		logger.Warn("no xschema declarations found")
		return generateStub()
	}

	// 2. Retrieve
	retrieverOpts := retriever.Options{
		Concurrency: cfg.Concurrency,
		HTTPTimeout: cfg.HTTPTimeout,
		Retries:     cfg.Retries,
		NoCache:     cfg.NoCache,
	}
	batches, err := retriever.Retrieve(ctx, decls, retrieverOpts)
	if err != nil {
		logger.Error("retrieve failed", "error", err)
		return fmt.Errorf("retrieve: %w", err)
	}

	// Filter by adapter if specified
	if cfg.Adapter != "" {
		logger.Info("filtering by adapter", "adapter", cfg.Adapter)
		batches = filterBatchesByAdapter(batches, cfg.Adapter)
		if len(batches) == 0 {
			logger.Error("no schemas found for adapter", "adapter", cfg.Adapter)
			return fmt.Errorf("no schemas found for adapter %q", cfg.Adapter)
		}
	}

	// 3. Generate
	outputsByLang := make(map[string][]generator.GenerateOutput)
	for _, batch := range batches {
		if cfg.DryRun {
			logger.Info("dry run", "adapter", batch.Adapter, "language", batch.Language)
			printDryRun(batch)
			continue
		}

		outputs, err := generator.Generate(ctx, batch)
		if err != nil {
			logger.Error("generate failed", "adapter", batch.Adapter, "error", err)
			return fmt.Errorf("generate (%s): %w", batch.Adapter, err)
		}
		outputsByLang[batch.Language] = append(outputsByLang[batch.Language], outputs...)
	}

	if cfg.DryRun {
		return nil
	}

	// 4. Inject
	for lang, outputs := range outputsByLang {
		err := injector.Inject(injector.InjectInput{
			Language: lang,
			Outputs:  outputs,
			OutDir:   cfg.OutputDir,
		})
		if err != nil {
			logger.Error("inject failed", "language", lang, "error", err)
			return fmt.Errorf("inject (%s): %w", lang, err)
		}
	}

	logger.Info("complete")
	return nil
}

// generateStub creates an empty schema file for the target language
func generateStub() error {
	lang := detectLanguage()
	if lang == "" {
		logger.Warn("no xschema declarations found")
		return nil
	}

	if cfg.DryRun {
		logger.Info("would generate stub", "language", lang, "output_dir", cfg.OutputDir)
		return nil
	}

	logger.Info("generating stub", "language", lang)

	err := injector.Inject(injector.InjectInput{
		Language: lang,
		Outputs:  []generator.GenerateOutput{}, // empty
		OutDir:   cfg.OutputDir,
	})
	if err != nil {
		return fmt.Errorf("inject stub (%s): %w", lang, err)
	}

	logger.Info("created stub", "path", cfg.OutputDir+"/"+language.ByName(lang).OutputFile)
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
	logger.Info("adapter batch", "adapter", batch.Adapter, "language", batch.Language, "schemas", len(batch.Schemas))
	for _, s := range batch.Schemas {
		var schema map[string]any
		json.Unmarshal(s.Schema, &schema)
		logger.Info("  - schema", "name", s.Name, "type", schema["type"])
	}
}
