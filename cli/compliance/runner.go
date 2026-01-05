package compliance

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/xschemadev/xschema/bundler"
)

// RunOptions configures the compliance test run
type RunOptions struct {
	AdapterPath    string                          // path to adapter package
	AdapterName    string                          // adapter name for display
	AdapterCLIPath func(adapterPath string) string // function to get adapter CLI path
	Drafts         []string                        // drafts to test (empty = all)
	SuitePath      string                          // path to JSON Schema Test Suite
	Runner         string                          // e.g., "bun", "bunx"
	RunnerArgs     []string                        // e.g., ["run"]
	Verbose        bool
	OutputFunc     func(string) // for progress output
}

// Run executes compliance tests for an adapter
func Run(ctx context.Context, opts RunOptions) (*ComplianceReport, error) {
	// Find harness template
	harnessFile, err := FindHarness(opts.AdapterPath)
	if err != nil {
		return nil, err
	}

	// Determine which drafts to test
	drafts := opts.Drafts
	if len(drafts) == 0 {
		drafts = Drafts
	}

	// Determine adapter CLI path
	if opts.AdapterCLIPath == nil {
		return nil, fmt.Errorf("AdapterCLIPath function not configured")
	}
	adapterBin := opts.AdapterCLIPath(opts.AdapterPath)
	if _, err := os.Stat(adapterBin); os.IsNotExist(err) {
		return nil, fmt.Errorf("adapter CLI not found at %s\nMake sure the adapter is built", adapterBin)
	}

	report := ComplianceReport{
		Adapter:     opts.AdapterName,
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Drafts:      []DraftResult{},
	}

	for _, draft := range drafts {
		select {
		case <-ctx.Done():
			return &report, ctx.Err()
		default:
		}

		if opts.OutputFunc != nil {
			opts.OutputFunc(fmt.Sprintf("Testing %s...", draft))
		}

		draftResult, err := runDraft(ctx, runDraftOptions{
			draft:       draft,
			suitePath:   opts.SuitePath,
			harnessFile: harnessFile,
			adapterBin:  adapterBin,
			runner:      opts.Runner,
			runnerArgs:  opts.RunnerArgs,
			workDir:     opts.AdapterPath,
			verbose:     opts.Verbose,
			outputFunc:  opts.OutputFunc,
		})

		if draftResult != nil {
			report.Drafts = append(report.Drafts, *draftResult)
		}

		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return &report, err
			}
			return nil, fmt.Errorf("failed to run %s: %w", draft, err)
		}
	}

	return &report, nil
}

type runDraftOptions struct {
	draft       string
	suitePath   string
	harnessFile string
	adapterBin  string
	runner      string
	runnerArgs  []string
	workDir     string // directory to run harness from (for dependency resolution)
	verbose     bool
	outputFunc  func(string)
}

func runDraft(ctx context.Context, opts runDraftOptions) (*DraftResult, error) {
	// Load test suite for this draft
	suite, err := LoadTestSuite(opts.suitePath, opts.draft)
	if err != nil {
		return nil, err
	}

	result := DraftResult{
		Draft:    opts.draft,
		Keywords: []KeywordResult{},
		Summary:  DraftSummary{},
	}

	// Sort keywords for consistent output
	keywords := make([]string, 0, len(suite))
	for keyword := range suite {
		keywords = append(keywords, keyword)
	}
	sort.Strings(keywords)

	for _, keyword := range keywords {
		select {
		case <-ctx.Done():
			return &result, ctx.Err()
		default:
		}

		groups := suite[keyword]

		keywordResult := KeywordResult{
			Keyword:  keyword,
			Failures: []TestResult{},
		}

		for _, group := range groups {
			if err := processGroup(ctx, opts, group, &keywordResult, &result.Summary); err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
					return &result, err
				}
			}
		}

		result.Keywords = append(result.Keywords, keywordResult)
	}

	// Calculate percentage
	if result.Summary.Total > 0 {
		result.Summary.Percentage = float64(result.Summary.Passed) / float64(result.Summary.Total) * 100
	}

	return &result, nil
}

func processGroup(ctx context.Context, opts runDraftOptions, group TestGroup, keywordResult *KeywordResult, summary *DraftSummary) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	bundledSchema, err := bundleSchema(ctx, group.Schema, opts.suitePath)
	if err != nil {
		markAllFailed(keywordResult, summary, group, fmt.Sprintf("bundling error: %v", err))
		return nil
	}

	adapterOutput, err := CallAdapter(ctx, opts.adapterBin, opts.runner, opts.runnerArgs, bundledSchema)
	if err != nil {
		return fmt.Errorf("adapter call failed: %w", err)
	}

	tempHarness, err := GenerateTempHarness(opts.harnessFile, adapterOutput.Schema, group.Tests, opts.workDir)
	if err != nil {
		markAllFailed(keywordResult, summary, group, fmt.Sprintf("harness generation error: %v", err))
		return nil
	}
	defer os.Remove(tempHarness)

	harnessResults, err := ExecuteHarness(ctx, tempHarness, opts.runner, opts.runnerArgs, opts.workDir)
	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		markAllFailed(keywordResult, summary, group, fmt.Sprintf("harness execution error: %v", err))
		return nil
	}

	processResults(harnessResults, group, keywordResult, summary)
	return nil
}

func markAllFailed(keywordResult *KeywordResult, summary *DraftSummary, group TestGroup, errorMsg string) {
	for _, tc := range group.Tests {
		keywordResult.Failed++
		keywordResult.Total++
		summary.Failed++
		summary.Total++

		keywordResult.Failures = append(keywordResult.Failures, TestResult{
			Group:    group.Description,
			Test:     tc.Description,
			Expected: tc.Valid,
			Actual:   "error",
			Passed:   false,
			Error:    errorMsg,
		})
	}
}

func processResults(harnessResults []HarnessResult, group TestGroup, keywordResult *KeywordResult, summary *DraftSummary) {
	for i, hr := range harnessResults {
		if i >= len(group.Tests) {
			break
		}
		tc := group.Tests[i]

		passed := (hr.Actual == "true" && tc.Valid) || (hr.Actual == "false" && !tc.Valid)

		keywordResult.Total++
		summary.Total++

		if passed {
			keywordResult.Passed++
			summary.Passed++
		} else {
			keywordResult.Failed++
			summary.Failed++

			keywordResult.Failures = append(keywordResult.Failures, TestResult{
				Group:    group.Description,
				Test:     tc.Description,
				Expected: tc.Valid,
				Actual:   hr.Actual,
				Passed:   false,
				Error:    hr.Error,
			})
		}
	}
}

// WriteResults writes the compliance results to the adapter's results directory
func WriteResults(adapterPath string, report *ComplianceReport) error {
	resultsDir := filepath.Join(adapterPath, "compliance", "results")

	// Create results directory if it doesn't exist
	if err := os.MkdirAll(resultsDir, 0755); err != nil {
		return fmt.Errorf("failed to create results directory: %w", err)
	}

	// Write individual draft JSON files
	for _, draft := range report.Drafts {
		jsonContent, err := GenerateJSONReport(draft)
		if err != nil {
			return fmt.Errorf("failed to generate JSON for %s: %w", draft.Draft, err)
		}

		jsonPath := filepath.Join(resultsDir, draft.Draft+".json")
		if err := os.WriteFile(jsonPath, []byte(jsonContent), 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", jsonPath, err)
		}
	}

	// Write markdown report
	markdownContent := GenerateMarkdownReport(*report)
	markdownPath := filepath.Join(resultsDir, "REPORT.md")
	if err := os.WriteFile(markdownPath, []byte(markdownContent), 0644); err != nil {
		return fmt.Errorf("failed to write REPORT.md: %w", err)
	}

	return nil
}

func bundleSchema(ctx context.Context, schema RawSchema, suitePath string) (RawSchema, error) {
	schemaJSON, err := json.Marshal(schema)
	if err != nil {
		return RawSchema{}, fmt.Errorf("failed to marshal schema: %w", err)
	}

	bundleOpts := bundler.Options{
		RemotesPath: filepath.Join(suitePath, "remotes"),
	}

	bundled, err := bundler.Bundle(ctx, schemaJSON, bundleOpts)
	if err != nil {
		return RawSchema{}, err
	}

	var result RawSchema
	if err := json.Unmarshal(bundled, &result); err != nil {
		return RawSchema{}, fmt.Errorf("failed to unmarshal bundled schema: %w", err)
	}

	return result, nil
}
