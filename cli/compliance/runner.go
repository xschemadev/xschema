package compliance

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/xschemadev/xschema/adapter"
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
	Language       *language.Language              // language configuration
	Verbose        bool
	Timing         *TimingSummary
	Jobs           int                     // number of parallel jobs (default 1)
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
		Adapter: opts.AdapterName,
		Drafts:  []DraftResult{},
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

		jobs := opts.Jobs
		if jobs < 1 {
			jobs = 1
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
			language:     opts.Language,
			workDir:      opts.AdapterPath,
			verbose:      opts.Verbose,
			jobs:         jobs,
			outputFunc:   opts.OutputFunc,
			progressFunc: opts.ProgressFunc,
			timing:       opts.Timing,
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
	language     *language.Language
	workDir      string // directory to run harness from (for dependency resolution)
	verbose      bool
	jobs         int // parallel keyword processing (default 1)
	outputFunc   func(string)
	progressFunc func(ProgressUpdate)
	timing       *TimingSummary
}

func runDraft(ctx context.Context, opts runDraftOptions) (*DraftResult, error) {
	// Load test suite for this draft
	loadStart := time.Now()
	suite, err := LoadTestSuite(opts.suitePath, opts.draft, opts.keyword)
	if opts.timing != nil {
		opts.timing.addSuiteLoad(time.Since(loadStart))
	}
	if err != nil {
		return nil, err
	}

	// Sort keywords for consistent output
	keywords := make([]string, 0, len(suite))
	for keyword := range suite {
		keywords = append(keywords, keyword)
	}
	sort.Strings(keywords)

	// Filter to specific keyword if requested
	if opts.keyword != "" {
		found := slices.Contains(keywords, opts.keyword)
		if !found {
			return nil, fmt.Errorf("keyword %q not found in %s (available: %d keywords)", opts.keyword, opts.draft, len(keywords))
		}
		keywords = []string{opts.keyword}
	}

	totalKeywords := len(keywords)

	// Sequential execution (jobs == 1) - original behavior
	if opts.jobs <= 1 {
		return runDraftSequential(ctx, opts, suite, keywords, totalKeywords)
	}

	// Parallel execution (jobs > 1)
	return runDraftParallel(ctx, opts, suite, keywords, totalKeywords)
}

// runDraftSequential processes keywords one at a time
func runDraftSequential(ctx context.Context, opts runDraftOptions, suite map[string][]TestGroup, keywords []string, totalKeywords int) (*DraftResult, error) {
	result := DraftResult{
		Draft:    opts.draft,
		Keywords: []KeywordResult{},
		Summary:  DraftSummary{},
	}

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

		if err := processKeyword(ctx, opts, groups, &keywordResult, &result.Summary); err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return &result, err
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

// runDraftParallel processes keywords concurrently with bounded parallelism
func runDraftParallel(ctx context.Context, opts runDraftOptions, suite map[string][]TestGroup, keywords []string, totalKeywords int) (*DraftResult, error) {
	result := DraftResult{
		Draft:    opts.draft,
		Keywords: make([]KeywordResult, totalKeywords),
		Summary:  DraftSummary{},
	}

	// Use a semaphore channel for bounded concurrency
	sem := make(chan struct{}, opts.jobs)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstErr error

	// Use a separate mutex for progress reporting to avoid interleaving
	var progressMu sync.Mutex
	progressCounter := 0

	for i, keyword := range keywords {
		select {
		case <-ctx.Done():
			result.Keywords = result.Keywords[:i]
			return &result, ctx.Err()
		default:
		}

		i, keyword := i, keyword // capture loop variables
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Acquire semaphore
			select {
			case sem <- struct{}{}:
				defer func() { <-sem }()
			case <-ctx.Done():
				return
			}

			// Report progress (with lock to prevent interleaving)
			if opts.progressFunc != nil {
				progressMu.Lock()
				progressCounter++
				opts.progressFunc(ProgressUpdate{
					Draft:        opts.draft,
					DraftNum:     opts.draftNum,
					DraftTotal:   opts.draftTotal,
					Keyword:      keyword,
					KeywordNum:   progressCounter,
					KeywordTotal: totalKeywords,
				})
				progressMu.Unlock()
			}

			groups := suite[keyword]

			keywordResult := KeywordResult{
				Keyword:  keyword,
				Failures: []TestResult{},
			}

			// Create a local summary to avoid concurrent writes
			var localSummary DraftSummary

			if err := processKeyword(ctx, opts, groups, &keywordResult, &localSummary); err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
					mu.Lock()
					if firstErr == nil {
						firstErr = err
					}
					mu.Unlock()
					return
				}
			}

			// Store result at the correct index for deterministic order
			mu.Lock()
			result.Keywords[i] = keywordResult
			result.Summary.Passed += localSummary.Passed
			result.Summary.Failed += localSummary.Failed
			result.Summary.Skipped += localSummary.Skipped
			result.Summary.Total += localSummary.Total
			mu.Unlock()
		}()
	}

	wg.Wait()

	if firstErr != nil {
		return &result, firstErr
	}

	// Calculate percentage
	if result.Summary.Total > 0 {
		result.Summary.Percentage = float64(result.Summary.Passed) / float64(result.Summary.Total) * 100
	}

	return &result, nil
}

// bundledGroup holds a test group with its pre-bundled schema
type bundledGroup struct {
	group         TestGroup
	bundledSchema RawSchema
	bundleErr     error
}

// processKeyword batches adapter invocation and harness execution for all groups in a keyword
func processKeyword(ctx context.Context, opts runDraftOptions, groups []TestGroup, keywordResult *KeywordResult, summary *DraftSummary) error {
	if len(groups) == 0 {
		return nil
	}

	// Phase 1: Bundle all schemas
	bundled := make([]bundledGroup, len(groups))
	for i, group := range groups {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		bundleStart := time.Now()
		bundledSchema, err := bundleSchema(ctx, group.Schema, opts.suitePath, opts.draft)
		opts.timing.addSchemaBundling(time.Since(bundleStart))

		bundled[i] = bundledGroup{
			group:         group,
			bundledSchema: bundledSchema,
			bundleErr:     err,
		}
	}

	// Mark bundle errors as failed first
	for i, bg := range bundled {
		if bg.bundleErr != nil {
			markAllFailed(keywordResult, summary, bg.group, fmt.Sprintf("bundling error: %v", bg.bundleErr))
			bundled[i].bundleErr = bg.bundleErr // keep marked
		}
	}

	// Phase 2: Batch adapter call for groups that bundled successfully
	var adapterInputs []adapter.ConvertInput
	var inputIndexes []int // maps adapter input index to bundled group index
	for i, bg := range bundled {
		if bg.bundleErr == nil {
			groupID := fmt.Sprintf("group_%d", i)
			adapterInputs = append(adapterInputs, adapter.ConvertInput{
				Namespace: "compliance",
				ID:        groupID,
				VarName:   groupID,
				Schema:    bg.bundledSchema.Raw(),
			})
			inputIndexes = append(inputIndexes, i)
		}
	}

	if len(adapterInputs) == 0 {
		return nil // all groups failed bundling
	}

	adapterStart := time.Now()
	adapterOutputs, err := CallAdapterBatch(ctx, opts.adapterBin, opts.runner, opts.runnerArgs, adapterInputs)
	opts.timing.addAdapterInvocation(time.Since(adapterStart))
	if err != nil {
		return fmt.Errorf("adapter call failed: %w", err)
	}

	// Build map of group index -> adapter output
	adapterOutputByGroup := make(map[int]*adapter.ConvertResult)
	for i, outIdx := range inputIndexes {
		adapterOutputByGroup[outIdx] = &adapterOutputs[i]
	}

	// Phase 3: Build harness items and execute single batch harness
	var harnessItems []HarnessItem
	groupByID := make(map[string]*TestGroup) // for result processing

	for i, bg := range bundled {
		if bg.bundleErr != nil {
			continue // already marked as failed
		}

		adapterOutput := adapterOutputByGroup[i]
		if adapterOutput == nil {
			markAllFailed(keywordResult, summary, bg.group, "adapter did not return output for this group")
			continue
		}

		groupID := fmt.Sprintf("group_%d", i)
		harnessItems = append(harnessItems, HarnessItem{
			GroupID:       groupID,
			AdapterOutput: adapterOutput,
			Tests:         bg.group.Tests,
		})
		groupByID[groupID] = &bg.group
	}

	if len(harnessItems) == 0 {
		return nil // all groups failed
	}

	// Generate single batch harness
	harnessStart := time.Now()
	tempHarness, err := GenerateHarness(opts.language, harnessItems, opts.workDir)
	opts.timing.addHarnessGeneration(time.Since(harnessStart))
	if err != nil {
		// Mark all groups as failed
		for _, item := range harnessItems {
			if group := groupByID[item.GroupID]; group != nil {
				markAllFailed(keywordResult, summary, *group, fmt.Sprintf("harness generation error: %v", err))
			}
		}
		return nil
	}
	defer os.Remove(tempHarness)

	// Execute batch harness
	execStart := time.Now()
	harnessResults, err := ExecuteHarness(ctx, tempHarness, opts.runner, opts.runnerArgs, opts.workDir)
	opts.timing.addHarnessExecution(time.Since(execStart))

	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		// Mark all groups as failed
		for _, item := range harnessItems {
			if group := groupByID[item.GroupID]; group != nil {
				markAllFailed(keywordResult, summary, *group, fmt.Sprintf("harness execution error: %v", err))
			}
		}
		return nil
	}

	// Process results by groupId
	processHarnessResults(harnessResults, groupByID, keywordResult, summary)
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

func processHarnessResults(harnessResults []HarnessResult, groupByID map[string]*TestGroup, keywordResult *KeywordResult, summary *DraftSummary) {
	for _, hr := range harnessResults {
		group := groupByID[hr.GroupID]
		if group == nil {
			continue // unknown group, skip
		}

		if hr.Index >= len(group.Tests) {
			continue // invalid index, skip
		}
		tc := group.Tests[hr.Index]

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

func bundleSchema(ctx context.Context, schema RawSchema, suitePath, draft string) (RawSchema, error) {
	schemaJSON, err := json.Marshal(schema)
	if err != nil {
		return RawSchema{}, fmt.Errorf("failed to marshal schema: %w", err)
	}

	bundleOpts := bundler.Options{
		RemotesPath: filepath.Join(suitePath, "remotes"),
		Draft:       draft,
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
