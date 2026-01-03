package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/xschema/cli/generator"
	"github.com/xschema/cli/injector"
	"github.com/xschema/cli/parser"
	"github.com/xschema/cli/retriever"
	"github.com/xschema/cli/ui"
)

var (
	projectDir string
	outputDir  string
	langFilter string
	verbose    bool
	dryRun     bool
	//TODO
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

	// Make output directory absolute relative to project root
	outDir := outputDir
	if !filepath.IsAbs(outDir) {
		outDir = filepath.Join(root, outDir)
	}

	// Step 1: Parse config files
	ui.Step(1, 4, "Scanning for xschema config files")
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
	ui.Step(2, 4, "Fetching schemas")
	retrieverOpts := retriever.DefaultOptions()

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

	// Step 3: Generate (with spinner per adapter)
	ui.Step(3, 4, "Generating validators")
	var outputs []generator.GenerateOutput
	err = ui.RunWithSpinner("Running adapters...", func() error {
		var genErr error
		outputs, genErr = generator.GenerateAll(ctx, schemas, result.Language.Name)
		return genErr
	})
	if err != nil {
		ui.ErrorMsg("Generation failed", err, "Make sure the adapter is installed")
		return err
	}

	// Step 4: Inject
	ui.Step(4, 4, "Writing output files")
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
	printSummary(schemas, outDir, generatedFile, time.Since(start))

	return nil
}

func printSummary(schemas []retriever.RetrievedSchema, outDir string, generatedFile string, duration time.Duration) {
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
