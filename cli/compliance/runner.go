package compliance

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/xschemadev/xschema/bundler"
	"github.com/xschemadev/xschema/language"
)

// ProgressUpdate contains info about current test progress
type ProgressUpdate struct {
	Draft        string // current draft being tested
	DraftNum     int    // current draft number (1-based)
	DraftTotal   int    // total number of drafts
	Keyword      string // current keyword being tested
	KeywordNum   int    // current keyword number (1-based)
	KeywordTotal int    // total keywords in this draft
}

// RunOptions configures the compliance test run
type RunOptions struct {
	AdapterPath    string                          // path to adapter package
	AdapterName    string                          // adapter name for display
	AdapterCLIPath func(adapterPath string) string // function to get adapter CLI path
	Drafts         []string                        // drafts to test (empty = all)
	Keyword        string                          // specific keyword to test (empty = all)
	SuitePath      string                          // path to JSON Schema Test Suite
	Runner         string                          // e.g., "bun", "bunx"
	RunnerArgs     []string                        // e.g., ["run"]
	QualityCheck   QualityCheckConfig              // quality check configuration
	Language       *language.Language              // language configuration
	Verbose        bool
	OutputFunc     func(string)            // for simple progress output (deprecated, use ProgressFunc)
	ProgressFunc   func(ProgressUpdate)    // for live progress updates
	DraftDoneFunc  func(draft DraftResult) // called when a draft completes
}

// Run executes compliance tests for an adapter
func Run(ctx context.Context, opts RunOptions) (*ComplianceReport, error) {
	// Determine which drafts to test
	drafts := opts.Drafts
	if len(drafts) == 0 {
		drafts = Drafts
	}

	// Validate language config
	if opts.Language == nil {
		return nil, fmt.Errorf("Language not configured")
	}
	if opts.Language.HarnessTemplate == "" {
		return nil, fmt.Errorf("harness template not configured for language %s", opts.Language.Name)
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

	for i, draft := range drafts {
		select {
		case <-ctx.Done():
			return &report, ctx.Err()
		default:
		}

		// Legacy output func support
		if opts.OutputFunc != nil && opts.ProgressFunc == nil {
			opts.OutputFunc(fmt.Sprintf("Testing %s...", draft))
		}

		draftResult, err := runDraft(ctx, runDraftOptions{
			draft:        draft,
			draftNum:     i + 1,
			draftTotal:   len(drafts),
			keyword:      opts.Keyword,
			suitePath:    opts.SuitePath,
			adapterBin:   adapterBin,
			runner:       opts.Runner,
			runnerArgs:   opts.RunnerArgs,
			qualityCheck: opts.QualityCheck,
			language:     opts.Language,
			workDir:      opts.AdapterPath,
			verbose:      opts.Verbose,
			outputFunc:   opts.OutputFunc,
			progressFunc: opts.ProgressFunc,
		})

		if draftResult != nil {
			report.Drafts = append(report.Drafts, *draftResult)
			// Notify caller that draft is complete
			if opts.DraftDoneFunc != nil {
				opts.DraftDoneFunc(*draftResult)
			}
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
	draft        string
	draftNum     int // 1-based
	draftTotal   int
	keyword      string // filter to specific keyword (empty = all)
	suitePath    string
	adapterBin   string
	runner       string
	runnerArgs   []string
	qualityCheck QualityCheckConfig
	language     *language.Language
	workDir      string // directory to run harness from (for dependency resolution)
	verbose      bool
	outputFunc   func(string)
	progressFunc func(ProgressUpdate)
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

	// Filter to specific keyword if requested
	if opts.keyword != "" {
		found := false
		for _, k := range keywords {
			if k == opts.keyword {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("keyword %q not found in %s (available: %d keywords)", opts.keyword, opts.draft, len(keywords))
		}
		keywords = []string{opts.keyword}
	}

	totalKeywords := len(keywords)

	for i, keyword := range keywords {
		select {
		case <-ctx.Done():
			return &result, ctx.Err()
		default:
		}

		// Report progress
		if opts.progressFunc != nil {
			opts.progressFunc(ProgressUpdate{
				Draft:        opts.draft,
				DraftNum:     opts.draftNum,
				DraftTotal:   opts.draftTotal,
				Keyword:      keyword,
				KeywordNum:   i + 1,
				KeywordTotal: totalKeywords,
			})
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

	tempHarness, err := GenerateTempHarness(opts.language, adapterOutput, group.Tests, opts.workDir)
	if err != nil {
		markAllFailed(keywordResult, summary, group, fmt.Sprintf("harness generation error: %v", err))
		return nil
	}
	defer os.Remove(tempHarness)

	// Run quality check on the generated harness file
	qualityCheckResult, err := QualityCheckHarness(ctx, tempHarness, opts.qualityCheck, opts.workDir)
	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		markAllFailed(keywordResult, summary, group, fmt.Sprintf("quality check error: %v", err))
		return nil
	}
	if !qualityCheckResult.Success {
		markAllFailed(keywordResult, summary, group, fmt.Sprintf("quality check failed: %s", qualityCheckResult.Output))
		return nil
	}

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
	// Normalize error message to remove machine-specific paths
	errorMsg = normalizeErrorPath(errorMsg)

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

// normalizeErrorPath replaces machine-specific cache paths with a placeholder
// to ensure compliance reports are deterministic across different machines
func normalizeErrorPath(errorMsg string) string {
	cacheDir, err := GetCacheDir()
	if err == nil && cacheDir != "" {
		errorMsg = strings.ReplaceAll(errorMsg, cacheDir, "$CACHE")
	}
	return errorMsg
}

func processResults(harnessResults []HarnessResult, group TestGroup, keywordResult *KeywordResult, summary *DraftSummary) {
	for i, hr := range harnessResults {
		if i >= len(group.Tests) {
			break
		}
		tc := group.Tests[i]

		keywordResult.Total++
		summary.Total++

		// Handle skipped results (from type-only adapters)
		if hr.Actual == "skipped" {
			keywordResult.Skipped++
			summary.Skipped++
			continue
		}

		passed := (hr.Actual == "true" && tc.Valid) || (hr.Actual == "false" && !tc.Valid)

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
				Error:    normalizeErrorPath(hr.Error),
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
