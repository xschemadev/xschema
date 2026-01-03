package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/xschema/cli/generator"
	"github.com/xschema/cli/injector"
	"github.com/xschema/cli/language"
	"github.com/xschema/cli/parser"
	"github.com/xschema/cli/retriever"
	"github.com/xschema/cli/ui"
)

var (
	clientFile string
	verbose    bool
	dryRun     bool
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Parse codebase, convert schemas, output native validators",
	RunE:  runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&clientFile, "client", "c", "", "path to client file (required)")
	generateCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show verbose output")
	generateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "show what would be generated without writing")

	generateCmd.MarkFlagRequired("client")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	start := time.Now()

	// Setup verbose mode
	ui.SetVerbose(verbose)

	ctx := cmd.Context()

	// Step 1: Parse client file
	ui.Step(1, 5, "Parsing client file")
	client, err := parser.ParseClient(ctx, clientFile)
	if err != nil {
		ui.ErrorMsg("Failed to parse client", err)
		return err
	}
	ui.Detail(fmt.Sprintf("Found %s (%s)", ui.Primary.Render(client.ClientName), client.Language.Name))
	ui.Verbosef("Config: outputDir=%s, maxParallelFetches=%d", client.Config.OutputDir, client.Config.MaxParallelFetches)

	// Step 2: Scan codebase for declarations
	ui.Step(2, 5, "Scanning for declarations")
	clientDir := filepath.Dir(clientFile)
	decls, err := parser.Parse(ctx, clientDir, client)
	if err != nil {
		ui.ErrorMsg("Failed to scan codebase", err)
		return err
	}
	ui.Detail(fmt.Sprintf("Found %d declarations", len(decls)))

	if len(decls) == 0 {
		ui.WarnMsg("No xschema declarations found")
		return nil
	}

	// Step 3: Fetch schemas (with spinner)
	var batches []generator.GenerateBatchInput
	retrieverOpts := retriever.Options{
		Concurrency: client.Config.MaxParallelFetches,
		HTTPTimeout: time.Duration(client.Config.RequestTimeoutMs) * time.Millisecond,
		Retries:     client.Config.MaxFetchRetries,
	}

	err = ui.RunWithSpinner("Fetching schemas...", func() error {
		var fetchErr error
		batches, fetchErr = retriever.Retrieve(ctx, decls, retrieverOpts)
		return fetchErr
	})
	if err != nil {
		ui.ErrorMsg("Failed to fetch schemas", err)
		return err
	}

	// Show what we fetched
	for _, d := range decls {
		ui.Detail(fmt.Sprintf("%s from %s", ui.Primary.Render(d.Name), d.Location))
	}
	ui.SuccessMsg(fmt.Sprintf("Fetched %d schemas", len(decls)))

	// Handle dry-run mode
	if dryRun {
		ui.Println()
		ui.Println(ui.Bold.Render("Dry run mode - no files will be written"))
		ui.Println()
		for _, batch := range batches {
			printDryRun(batch)
		}
		return nil
	}

	// Step 4: Generate (with spinner per adapter)
	ui.Step(4, 5, "Generating validators")
	outputsByLang := make(map[string][]generator.GenerateOutput)
	for _, batch := range batches {
		var outputs []generator.GenerateOutput
		err = ui.RunWithSpinner(fmt.Sprintf("Running %s adapter...", ui.Primary.Render(batch.Adapter)), func() error {
			var genErr error
			outputs, genErr = generator.Generate(ctx, batch)
			return genErr
		})
		if err != nil {
			ui.ErrorMsg("Generation failed", err, "Make sure the adapter is installed")
			return err
		}
		outputsByLang[batch.Language] = append(outputsByLang[batch.Language], outputs...)
	}

	// Step 6: Inject
	ui.Step(5, 5, "Writing output files")
	var generatedFiles []string
	for lang, outputs := range outputsByLang {
		err := injector.Inject(injector.InjectInput{
			Language: lang,
			Outputs:  outputs,
			OutDir:   client.Config.OutputDir,
		})
		if err != nil {
			ui.ErrorMsg("Failed to write output", err)
			return err
		}
		// Track generated file
		l := language.ByName(lang)
		if l != nil {
			generatedFiles = append(generatedFiles, filepath.Join(client.Config.OutputDir, l.OutputFile))
		}
	}

	// Inject schemas import into client file
	if err := injector.InjectClient(ctx, injector.InjectClientInput{
		ClientFile: clientFile,
		Language:   client.Language,
		OutDir:     client.Config.OutputDir,
	}); err != nil {
		ui.ErrorMsg("Failed to inject client import", err)
		return err
	}

	// Summary
	printSummary(batches, client.Config.OutputDir, generatedFiles, time.Since(start))

	return nil
}

func printSummary(batches []generator.GenerateBatchInput, outDir string, files []string, duration time.Duration) {
	ui.Println()
	ui.SuccessMsg(fmt.Sprintf("Generation complete (%s)", ui.FormatDuration(duration)))
	ui.Println()

	ui.Println("  Schemas generated:")
	for _, b := range batches {
		ui.Printf("    %s\n", ui.Primary.Render(b.Adapter))
		for _, s := range b.Schemas {
			ui.Printf("      %s %s\n", ui.Dim.Render("•"), s.Name)
		}
	}
	ui.Println()

	if len(files) > 0 {
		ui.Printf("  Output: %s\n", ui.Primary.Render(files[0]))
		for _, f := range files[1:] {
			ui.Printf("          %s\n", ui.Primary.Render(f))
		}
	} else {
		ui.Printf("  Output: %s\n", ui.Primary.Render(outDir))
	}
	ui.Println()

	ui.Printf("  %s Check the generated file to verify the output\n", ui.Dim.Render("Tip:"))
}

func printDryRun(batch generator.GenerateBatchInput) {
	ui.Printf("  %s (%s)\n", ui.Primary.Render(batch.Adapter), batch.Language)
	for _, s := range batch.Schemas {
		var schema map[string]any
		json.Unmarshal(s.Schema, &schema)
		schemaType := "object"
		if t, ok := schema["type"].(string); ok {
			schemaType = t
		}
		ui.Printf("    %s %s %s\n", ui.Dim.Render("•"), s.Name, ui.Dim.Render(fmt.Sprintf("(%s)", schemaType)))
	}
}
