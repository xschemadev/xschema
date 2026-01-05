package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/xschemadev/xschema/compliance"
	"github.com/xschemadev/xschema/language"
	"github.com/xschemadev/xschema/ui"
)

var (
	complianceDraft     string
	complianceLang      string
	complianceDevReport bool
	complianceVerbose   bool
)

var complianceCmd = &cobra.Command{
	Use:   "compliance",
	Short: "Run JSON Schema Test Suite compliance tests for adapters",
	Long: `Run compliance tests against the official JSON Schema Test Suite.

This command must be run from within an adapter package directory
(a directory containing compliance/harness.*).

Examples:
  # Run from within an adapter package (prints results)
  xschema compliance

  # Run all drafts and write results to compliance/results/
  xschema compliance --dev-report

  # Run specific draft only (prints results, no file output)
  xschema compliance --draft draft2020-12`,
	RunE: runCompliance,
	Args: cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(complianceCmd)

	complianceCmd.Flags().StringVarP(&complianceDraft, "draft", "d", "", "specific draft to test (e.g., draft2020-12)")
	complianceCmd.Flags().StringVarP(&complianceLang, "lang", "l", "typescript", "language (typescript, python)")
	complianceCmd.Flags().BoolVar(&complianceDevReport, "dev-report", false, "write results to compliance/results/ (for adapter developers)")
	complianceCmd.Flags().BoolVarP(&complianceVerbose, "verbose", "v", false, "show verbose output")
}

func runCompliance(cmd *cobra.Command, args []string) error {
	start := time.Now()
	ctx := cmd.Context()

	ui.SetVerbose(complianceVerbose)

	// --dev-report and --draft are mutually exclusive
	if complianceDevReport && complianceDraft != "" {
		return fmt.Errorf("--dev-report and --draft cannot be used together\n\n--dev-report runs all drafts and writes the full report")
	}

	// Get language config
	lang := language.ByName(complianceLang)
	if lang == nil {
		return fmt.Errorf("unknown language: %s", complianceLang)
	}

	// Must be run from within an adapter package (has compliance/harness.*)
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	harnessFile, err := compliance.FindHarness(cwd)
	if err != nil {
		return fmt.Errorf("not in an adapter package (no compliance/harness.* found)\n\nRun this command from within an adapter package directory")
	}

	adapterPath := cwd
	var adapterName string
	if lang.GetPackageName != nil {
		adapterName = lang.GetPackageName(cwd)
	} else {
		adapterName = filepath.Base(cwd)
	}
	ui.Verbosef("found harness at %s", harnessFile)

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
	}

	ui.Println(ui.Bold.Render("JSON Schema Compliance Testing"))
	ui.Printf("Adapter: %s\n", ui.Primary.Render(adapterName))
	ui.Println()

	opts := compliance.RunOptions{
		AdapterPath:    adapterPath,
		AdapterName:    adapterName,
		AdapterCLIPath: lang.AdapterCLIPath,
		Drafts:         drafts,
		SuitePath:      suitePath,
		Runner:         runner,
		RunnerArgs:     runnerArgs,
		Verbose:        complianceVerbose,
		OutputFunc: func(msg string) {
			ui.Detail(msg)
		},
	}

	report, err := compliance.Run(ctx, opts)
	if err != nil && (errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)) {
		ui.Println()
		ui.WarnMsg("Aborted by user")
		return nil
	}

	if err != nil {
		return fmt.Errorf("compliance tests failed: %w", err)
	}

	// Print summary
	printComplianceSummary(report)

	// Write results if --dev-report
	if complianceDevReport {
		if err := compliance.WriteResults(adapterPath, report); err != nil {
			return fmt.Errorf("failed to write results: %w", err)
		}
		resultsPath := filepath.Join(adapterPath, "compliance", "results", "REPORT.md")
		ui.Printf("\nResults written to: %s\n", ui.Primary.Render(resultsPath))
	}

	ui.Println()
	ui.SuccessMsg(fmt.Sprintf("Compliance testing complete (%s)", ui.FormatDuration(time.Since(start))))

	return nil
}

func printComplianceSummary(report *compliance.ComplianceReport) {
	ui.Println("Results:")
	for _, draft := range report.Drafts {
		status := "✅"
		if draft.Summary.Percentage < 80 {
			status = "❌"
		} else if draft.Summary.Percentage < 95 {
			status = "⚠️"
		}

		ui.Printf("  %s %s: %d/%d (%.1f%%)\n",
			status,
			draft.Draft,
			draft.Summary.Passed,
			draft.Summary.Total,
			draft.Summary.Percentage,
		)
	}
}
