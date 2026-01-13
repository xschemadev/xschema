package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/xschemadev/xschema/adapter"
	"github.com/xschemadev/xschema/fetcher"
	"github.com/xschemadev/xschema/generator"
	"github.com/xschemadev/xschema/injector"
	"github.com/xschemadev/xschema/parser"
	"github.com/xschemadev/xschema/processor"
	"github.com/xschemadev/xschema/retriever"
	"github.com/xschemadev/xschema/ui"
)

var (
	projectDir  string
	outputDir   string
	langFilter  string
	verbose     bool
	dryRun      bool
	concurrency int
	envFile     string
	// TODO: implement watch mode
	watch bool
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Parse config files, convert schemas, output native validators",
	RunE:  runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&projectDir, "project", "p", "", "project root directory (default: current directory)")
	generateCmd.Flags().StringVarP(&outputDir, "output", "o", ".xschema", "output directory for generated files")
	generateCmd.Flags().StringVar(&langFilter, "lang", "", "filter to specific language if multiple detected")
	generateCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show verbose output")
	generateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "show what would be generated without writing")
	generateCmd.Flags().IntVarP(&concurrency, "concurrency", "c", min(runtime.NumCPU(), 8), "number of parallel schema fetches")
	generateCmd.Flags().StringVar(&envFile, "env-file", "", "path to env file for header variable substitution")
	generateCmd.Flags().BoolVarP(&watch, "watch", "w", false, "watch for changes and regenerate")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	start := time.Now()

	// Setup verbose mode
	ui.SetVerbose(verbose)

	ctx := cmd.Context()

	// Determine project root
	root := projectDir
	if root == "" {
		var err error
		root, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %w", err)
		}
	}

	// Load env file for header variable substitution
	if envFile != "" {
		ui.Verbosef("loading env file: %s", envFile)
		if err := godotenv.Load(envFile); err != nil {
			return fmt.Errorf("failed to load env file %s: %w", envFile, err)
		}
	} else {
		// Try to load .env from cwd if it exists
		defaultEnvPath := filepath.Join(root, ".env")
		if _, err := os.Stat(defaultEnvPath); err == nil {
			ui.Verbosef("loading default .env file: %s", defaultEnvPath)
			_ = godotenv.Load(defaultEnvPath) // ignore error if .env doesn't exist or is invalid
		}
	}

	// Make output directory absolute relative to project root
	outDir := outputDir
	if !filepath.IsAbs(outDir) {
		outDir = filepath.Join(root, outDir)
	}

	// Step 1: Parse config files
	ui.Step(1, 5, "Scanning for xschema config files")
	result, err := parser.Parse(ctx, root, langFilter)
	if err != nil {
		ui.ErrorMsg("Failed to parse config files", err)
		return err
	}
	ui.Detail(fmt.Sprintf("Found %d config files, %d schemas (%s)",
		len(result.Configs), len(result.Declarations), result.Language.Name))

	if len(result.Declarations) == 0 {
		ui.WarnMsg("No schema declarations found")
		return nil
	}

	// Step 2: Fetch schemas (with spinner)
	ui.Step(2, 5, "Fetching schemas")

	// Create shared cache for all schema fetching (retriever, processor, metaschema)
	sharedCache := fetcher.NewSharedCache()

	retrieverOpts := retriever.DefaultOptions()
	retrieverOpts.Concurrency = concurrency
	retrieverOpts.Cache = sharedCache

	var schemas []retriever.RetrievedSchema
	err = ui.RunWithSpinner("Fetching schemas...", func() error {
		var fetchErr error
		schemas, fetchErr = retriever.Retrieve(ctx, result.Declarations, retrieverOpts)
		return fetchErr
	})
	if err != nil {
		ui.ErrorMsg("Failed to fetch schemas", err)
		return err
	}

	// Show what we fetched
	for _, s := range schemas {
		ui.Detail(fmt.Sprintf("%s from %s", ui.Primary.Render(s.Key()), s.Adapter))
	}
	ui.SuccessMsg(fmt.Sprintf("Fetched %d schemas", len(schemas)))

	// Handle dry-run mode
	if dryRun {
		ui.Println()
		ui.Println(ui.Bold.Render("Dry run mode - no files will be written"))
		ui.Println()
		printDryRunSchemas(schemas)
		return nil
	}

	// Step 3: Process schemas (crawl external refs, validate, bundle)
	ui.Step(3, 5, "Processing schemas")
	var processed []processor.ProcessedSchema
	err = ui.RunWithSpinner("Processing schemas...", func() error {
		var procErr error
		processed, procErr = processor.Process(ctx, schemas, processor.Options{
			Fetcher:   newRetrieverFetcher(ctx, retrieverOpts),
			OnVerbose: verboseCallback(),
			Cache:     sharedCache, // reuse cache from retriever
		})
		return procErr
	})
	if err != nil {
		ui.ErrorMsg("Processing failed", err)
		return err
	}
	ui.SuccessMsg(fmt.Sprintf("Processed %d schemas", len(processed)))

	// Step 4: Generate (with spinner per adapter)
	ui.Step(4, 5, "Generating validators")
	var outputs []adapter.ConvertResult
	err = ui.RunWithSpinner("Running adapters...", func() error {
		var genErr error
		outputs, genErr = generator.GenerateAll(ctx, processed, result.Language.Name, root)
		return genErr
	})
	if err != nil {
		ui.ErrorMsg("Generation failed", err, "Make sure the adapter is installed")
		return err
	}

	// Step 5: Inject
	ui.Step(5, 5, "Writing output files")
	err = injector.Inject(injector.InjectInput{
		Language: result.Language.Name,
		Outputs:  outputs,
		OutDir:   outDir,
	})
	if err != nil {
		ui.ErrorMsg("Failed to write output", err)
		return err
	}

	// Summary
	generatedFile := filepath.Join(outDir, result.Language.OutputFile)
	printSummary(schemas, generatedFile, time.Since(start))

	return nil
}

func printSummary(schemas []retriever.RetrievedSchema, generatedFile string, duration time.Duration) {
	ui.Println()
	ui.SuccessMsg(fmt.Sprintf("Generation complete (%s)", ui.FormatDuration(duration)))
	ui.Println()

	// Group by namespace for display
	byNamespace := make(map[string][]retriever.RetrievedSchema)
	for _, s := range schemas {
		byNamespace[s.Namespace] = append(byNamespace[s.Namespace], s)
	}

	ui.Println("  Schemas generated:")
	for ns, nsSchemas := range byNamespace {
		ui.Printf("    %s\n", ui.Primary.Render(ns))
		for _, s := range nsSchemas {
			ui.Printf("      %s %s\n", ui.Dim.Render("•"), s.ID)
		}
	}
	ui.Println()

	ui.Printf("  Output: %s\n", ui.Primary.Render(generatedFile))
	ui.Println()

	ui.Printf("  %s Check the generated file to verify the output\n", ui.Dim.Render("Tip:"))
}

func printDryRunSchemas(schemas []retriever.RetrievedSchema) {
	// Group by adapter
	byAdapter := retriever.GroupByAdapter(schemas)
	adapters := retriever.SortedAdapters(byAdapter)

	for _, adapter := range adapters {
		adapterSchemas := byAdapter[adapter]
		ui.Printf("  %s\n", ui.Primary.Render(adapter))
		for _, s := range adapterSchemas {
			ui.Printf("    %s %s\n", ui.Dim.Render("•"), s.Key())
		}
	}
}

// retrieverFetcher wraps retriever package to implement fetcher.Fetcher.
type retrieverFetcher struct {
	ctx  context.Context
	opts retriever.Options
}

func newRetrieverFetcher(ctx context.Context, opts retriever.Options) *retrieverFetcher {
	return &retrieverFetcher{ctx: ctx, opts: opts}
}

func (f *retrieverFetcher) Fetch(ctx context.Context, uri string) (json.RawMessage, error) {
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		return retriever.RetrieveFromURL(ctx, uri, f.opts)
	}
	return retriever.RetrieveFromFilePath(ctx, uri)
}

// verboseCallback returns a callback for processor verbose output when verbose mode is enabled.
func verboseCallback() func(string) {
	if !verbose {
		return nil
	}
	return func(msg string) {
		ui.Verbosef("%s", msg)
	}
}
