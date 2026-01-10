package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/xschemadev/xschema/compliance"
	"github.com/xschemadev/xschema/language"
	"github.com/xschemadev/xschema/ui"
)

var (
	complianceDraft       string
	complianceKeyword     string
	complianceLang        string
	complianceDevReport   bool
	complianceVerbose     bool
	complianceAdapterPath string
	complianceProfile     bool
	complianceConcurrency int
)

var complianceCmd = &cobra.Command{
	Use:   "compliance",
	Short: "Run JSON Schema Test Suite compliance tests for adapters",
	Long: `Run compliance tests against the official JSON Schema Test Suite.

This command must be run from within an adapter package directory.

Examples:
  # Run from within an adapter package (prints results)
  xschema compliance

  # Run all drafts and write results to compliance/results/
  xschema compliance --dev-report

  # Run specific draft only (prints results, no file output)
  xschema compliance --draft draft2020-12

  # Run specific keyword across all drafts
  xschema compliance --keyword additionalProperties

  # Run specific keyword for a specific draft
  xschema compliance --draft draft2020-12 --keyword additionalProperties`,
	RunE: runCompliance,
	Args: cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(complianceCmd)

	complianceCmd.Flags().StringVarP(&complianceDraft, "draft", "d", "", "specific draft to test (e.g., draft2020-12)")
	complianceCmd.Flags().StringVarP(&complianceKeyword, "keyword", "k", "", "specific keyword to test (e.g., additionalProperties)")
	complianceCmd.Flags().StringVarP(&complianceLang, "lang", "l", "typescript", "language (typescript, python)")
	complianceCmd.Flags().BoolVar(&complianceDevReport, "dev-report", false, "write results to compliance/results/ (for adapter developers)")
	complianceCmd.Flags().BoolVarP(&complianceVerbose, "verbose", "v", false, "show verbose output")
	complianceCmd.Flags().BoolVar(&complianceProfile, "profile", false, "print timing breakdown summary")
	complianceCmd.Flags().IntVarP(&complianceConcurrency, "concurrency", "c", min(runtime.NumCPU(), 8), "number of parallel jobs (keywords processed concurrently)")
	complianceCmd.Flags().StringVarP(&complianceAdapterPath, "adapter-path", "p", "", "path to adapter directory (defaults to current directory)")
}

func runCompliance(cmd *cobra.Command, args []string) error {
	start := time.Now()
	ctx := cmd.Context()

	ui.SetVerbose(complianceVerbose)

	if complianceDevReport && (complianceDraft != "" || complianceKeyword != "") {
		return fmt.Errorf("--dev-report cannot be used with --draft or --keyword")
	}

	if complianceConcurrency < 1 {
		return fmt.Errorf("--concurrency must be >= 1")
	}

	// Get language config
	lang := language.ByName(complianceLang)
	if lang == nil {
		return fmt.Errorf("unknown language: %s", complianceLang)
	}

	// Determine adapter path (from flag or current directory)
	var adapterPath string
	if complianceAdapterPath != "" {
		absPath, err := filepath.Abs(complianceAdapterPath)
		if err != nil {
			return fmt.Errorf("failed to resolve adapter path: %w", err)
		}
		adapterPath = absPath
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %w", err)
		}
		adapterPath = cwd
	}
	var adapterName string
	if lang.GetPackageName != nil {
		adapterName = lang.GetPackageName(adapterPath)
	} else {
		adapterName = filepath.Base(adapterPath)
	}

	// Fetch test suite (downloads from GitHub if not cached)
	var suitePath string
	if err := ui.RunWithSpinner("Fetching JSON Schema Test Suite...", func() error {
		var fetchErr error
		suitePath, fetchErr = compliance.FetchTestSuite(ctx)
		return fetchErr
	}); err != nil {
		return fmt.Errorf("failed to fetch test suite: %w", err)
	}

	// Detect harness runner for this language/directory
	if lang.DetectHarnessRunner == nil {
		return fmt.Errorf("language %s does not have harness runner detection configured", complianceLang)
	}
	runner, runnerArgs, err := lang.DetectHarnessRunner(adapterPath)
	if err != nil {
		return fmt.Errorf("failed to detect harness runner: %w", err)
	}

	// Determine drafts to test
	var drafts []string
	if complianceDraft != "" {
		drafts = []string{complianceDraft}
	} else {
		drafts = compliance.Drafts
	}

	// Header
	ui.Println(ui.Bold.Render("JSON Schema Compliance Testing"))
	ui.Printf("Adapter: %s\n", ui.Primary.Render(adapterName))
	ui.Println()

	// Show what we're testing
	ui.Printf("Testing %d drafts: %s\n", len(drafts), ui.Dim.Render(formatDraftList(drafts)))
	if complianceKeyword != "" {
		ui.Printf("Keyword filter: %s\n", ui.Primary.Render(complianceKeyword))
	}
	ui.Println()

	// Create progress spinner for live updates
	spinner := ui.NewProgressSpinner()

	var timing *compliance.TimingSummary
	if complianceProfile {
		timing = &compliance.TimingSummary{}
	}

	opts := compliance.RunOptions{
		AdapterPath:    adapterPath,
		AdapterName:    adapterName,
		AdapterCLIPath: lang.AdapterCLIPath,
		Drafts:         drafts,
		Keyword:        complianceKeyword,
		SuitePath:      suitePath,
		Runner:         runner,
		RunnerArgs:     runnerArgs,
		Language:       lang,
		Verbose:        complianceVerbose,
		Timing:         timing,
		Jobs:           complianceConcurrency,
		ProgressFunc: func(p compliance.ProgressUpdate) {
			msg := ui.FormatDraftProgress(p.Draft, p.KeywordNum, p.KeywordTotal, p.Keyword)
			spinner.Update(msg)
		},
		DraftDoneFunc: func(draft compliance.DraftResult) {
			status, coverage := formatDraftResult(draft.Summary)
			msg := fmt.Sprintf("%s %s: %s", status, ui.Bold.Render(draft.Draft), coverage)
			spinner.PrintAboveSpinner(msg)
		},
	}

	// Start spinner and run tests
	spinner.Start(ui.FormatDraftProgress(drafts[0], 0, 0, "starting..."))
	report, err := compliance.Run(ctx, opts)
	spinner.Stop()

	if err != nil && (errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)) {
		ui.Println()
		ui.WarnMsg("Aborted by user")
		return nil
	}

	if err != nil {
		return fmt.Errorf("compliance tests failed: %w", err)
	}

	// Write results if --dev-report
	if complianceDevReport {
		if err := compliance.WriteResults(adapterPath, report); err != nil {
			return fmt.Errorf("failed to write results: %w", err)
		}
		resultsPath := filepath.Join(adapterPath, "compliance", "results", "REPORT.md")
		ui.Printf("\nResults written to: %s\n", ui.Primary.Render(resultsPath))
	}

	if complianceProfile {
		printTimingSummary(timing)
	}

	ui.Println()
	ui.SuccessMsg(fmt.Sprintf("Compliance testing complete (%s)", ui.FormatDuration(time.Since(start))))

	return nil
}

func formatDraftResult(summary compliance.DraftSummary) (status, coverage string) {
	isTypeOnly := summary.Skipped == summary.Total && summary.Total > 0
	if isTypeOnly {
		return ui.Bold.Render("○"), "type-only (skipped)"
	}
	status = ui.Success.Render("✓")
	if summary.Percentage < 80 {
		status = ui.Error.Render("✗")
	} else if summary.Percentage < 95 {
		status = ui.Warning.Render("!")
	}
	coverage = fmt.Sprintf("%d/%d (%.1f%%)", summary.Passed, summary.Total, summary.Percentage)
	return status, coverage
}

func printTimingSummary(timing *compliance.TimingSummary) {
	if timing == nil {
		return
	}

	ui.Println()
	ui.Println("Timing summary:")
	ui.Printf("  Suite load: %s\n", ui.FormatDuration(timing.SuiteLoad))
	ui.Printf("  Schema bundling: %s\n", ui.FormatDuration(timing.SchemaBundling))
	ui.Printf("  Adapter invocation: %s\n", ui.FormatDuration(timing.AdapterInvocation))
	ui.Printf("  Harness generation: %s\n", ui.FormatDuration(timing.HarnessGeneration))
	ui.Printf("  Harness execution: %s\n", ui.FormatDuration(timing.HarnessExecution))
}

func formatDraftList(drafts []string) string {
	if len(drafts) <= 3 {
		return strings.Join(drafts, ", ")
	}
	return fmt.Sprintf("%s, %s, ... +%d more", drafts[0], drafts[1], len(drafts)-2)
}
