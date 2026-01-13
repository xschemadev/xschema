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
	"github.com/xschemadev/xschema/fetcher"
	"github.com/xschemadev/xschema/language"
	"github.com/xschemadev/xschema/processor"
	"github.com/xschemadev/xschema/retriever"
	"github.com/xschemadev/xschema/unsupported"
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
			adapterName:  opts.AdapterName,
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
	adapterName  string // adapter name for limitation checking
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
			result.Summary.UnsupportedFeatures.Count += localSummary.UnsupportedFeatures.Count
			result.Summary.UnsupportedFeatures.Items = append(result.Summary.UnsupportedFeatures.Items, localSummary.UnsupportedFeatures.Items...)
			mu.Unlock()
		}()
	}

	wg.Wait()

	if firstErr != nil {
		return &result, firstErr
	}

	// Sort unsupported features for deterministic output
	sort.Slice(result.Summary.UnsupportedFeatures.Items, func(i, j int) bool {
		return result.Summary.UnsupportedFeatures.Items[i].Path < result.Summary.UnsupportedFeatures.Items[j].Path
	})

	// Calculate percentage
	if result.Summary.Total > 0 {
		result.Summary.Percentage = float64(result.Summary.Passed) / float64(result.Summary.Total) * 100
	}

	return &result, nil
}

// bundledGroup holds a test group with its pre-bundled schema
type bundledGroup struct {
	group          TestGroup
	bundledSchema  RawSchema
	bundleErr      error
	unsupportedErr *unsupported.UnsupportedKeywordError
}

// groupFilter holds test information for a group
type groupFilter struct {
	allKnown           bool
	adapterLimitation  bool   // true if group matches an adapter-specific unsupported pattern
	adapterLimitReason string // reason for adapter limitation
	filteredTests      []TestCase
	filteredIndices    []int
}

// processKeyword batches adapter invocation and harness execution for all groups in a keyword.
// Pipeline: filter → bundle → adapt → harness → results
func processKeyword(ctx context.Context, opts runDraftOptions, groups []TestGroup, keywordResult *KeywordResult, summary *DraftSummary) error {
	if len(groups) == 0 {
		return nil
	}

	// Phase 1: Initialize group filters (check adapter limitations, unsupported keyword detection happens in bundling)
	groupFilters := initGroupFilters(groups, opts.adapterName)

	// Phase 1b: Mark adapter limitations as unsupported (before bundling)
	markAdapterLimitations(groups, groupFilters, keywordResult, summary, opts.draft, keywordResult.Keyword)

	// Phase 2: Bundle schemas
	bundled, err := bundleSchemas(ctx, groups, groupFilters, opts, keywordResult.Keyword)
	if err != nil {
		return err
	}
	markBundleErrors(bundled, groupFilters, keywordResult, summary, opts.draft, keywordResult.Keyword)

	// Phase 3: Call adapter
	adapterOutputByGroup, err := callAdapter(ctx, bundled, opts)
	if err != nil {
		return err
	}
	if adapterOutputByGroup == nil {
		return nil // all groups failed bundling
	}

	// Phase 4: Build harness items
	harnessItems, groupByID, filteredTestsByGroup := buildHarnessItems(
		bundled, groupFilters, adapterOutputByGroup, keywordResult, summary,
	)
	if len(harnessItems) == 0 {
		return nil // all groups failed
	}

	// Phase 5: Execute harness
	harnessResults, err := executeHarness(ctx, harnessItems, groupByID, opts, keywordResult, summary)
	if err != nil {
		return err
	}
	if harnessResults == nil {
		return nil // harness failed, errors already recorded
	}

	// Phase 6: Process results
	processHarnessResults(harnessResults, groupByID, filteredTestsByGroup, keywordResult, summary)
	return nil
}

// initGroupFilters creates groupFilter entries for each test group.
// Checks for adapter-specific limitations. Unsupported keyword detection happens during bundling.
func initGroupFilters(groups []TestGroup, adapterName string) []groupFilter {
	groupFilters := make([]groupFilter, len(groups))

	for i, group := range groups {
		filteredIndices := make([]int, len(group.Tests))
		for j := range group.Tests {
			filteredIndices[j] = j
		}

		gf := groupFilter{
			allKnown:        false,
			filteredTests:   group.Tests,
			filteredIndices: filteredIndices,
		}

		// Check for adapter-specific limitations
		if matches, reason := MatchesUnsupportedPattern(adapterName, group.Description); matches {
			gf.adapterLimitation = true
			gf.adapterLimitReason = reason
		}

		groupFilters[i] = gf
	}

	return groupFilters
}

// markAdapterLimitations marks test groups that match adapter-specific unsupported patterns as unsupported.
func markAdapterLimitations(groups []TestGroup, groupFilters []groupFilter, keywordResult *KeywordResult, summary *DraftSummary, draft, keyword string) {
	for i, gf := range groupFilters {
		if !gf.adapterLimitation {
			continue
		}

		group := groups[i]
		for _, tc := range group.Tests {
			testPath := fmt.Sprintf("%s/%s/%s/%s", draft, keyword, group.Description, tc.Description)
			summary.UnsupportedFeatures.Count++
			summary.UnsupportedFeatures.Items = append(summary.UnsupportedFeatures.Items, UnsupportedFeatureItem{
				Path:   testPath,
				Reason: gf.adapterLimitReason,
			})
		}
	}
}

// bundleSchemas processes each schema through the processor pipeline.
func bundleSchemas(ctx context.Context, groups []TestGroup, groupFilters []groupFilter, opts runDraftOptions, keyword string) ([]bundledGroup, error) {
	bundled := make([]bundledGroup, len(groups))
	remotesPath := filepath.Join(opts.suitePath, "remotes")
	localhostFetcher := fetcher.NewLocalhostFetcher(remotesPath)

	for i, group := range groups {
		// Skip groups where all tests are known issues
		if groupFilters[i].allKnown {
			bundled[i] = bundledGroup{
				group:     group,
				bundleErr: fmt.Errorf("skipped: all tests are known issues"),
			}
			continue
		}

		// Skip groups with adapter-specific limitations (already marked as unsupported)
		if groupFilters[i].adapterLimitation {
			bundled[i] = bundledGroup{
				group:     group,
				bundleErr: fmt.Errorf("skipped: adapter limitation"),
			}
			continue
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		schemaJSON, err := json.Marshal(group.Schema)
		if err != nil {
			bundled[i] = bundledGroup{
				group:     group,
				bundleErr: fmt.Errorf("failed to marshal schema: %w", err),
			}
			continue
		}

		bundleStart := time.Now()
		toProcess := []retriever.RetrievedSchema{{
			Namespace: "compliance",
			ID:        fmt.Sprintf("group_%d", i),
			Schema:    schemaJSON,
			Adapter:   "compliance",
			SourceURI: fmt.Sprintf("compliance://%s/%s/group_%d", opts.draft, keyword, i),
		}}

		processed, err := processor.Process(ctx, toProcess, processor.Options{
			Fetcher: localhostFetcher,
			Draft:   opts.draft,
		})
		opts.timing.addSchemaBundling(time.Since(bundleStart))

		if err != nil {
			// Check if this is an UnsupportedKeywordError - these should be skipped, not failed
			var unsupportedErr *unsupported.UnsupportedKeywordError
			if errors.As(err, &unsupportedErr) {
				bundled[i] = bundledGroup{group: group, unsupportedErr: unsupportedErr}
			} else {
				bundled[i] = bundledGroup{group: group, bundleErr: err}
			}
			continue
		}

		if len(processed) == 0 {
			bundled[i] = bundledGroup{group: group, bundleErr: fmt.Errorf("processor returned no results")}
			continue
		}

		var schema RawSchema
		if jsonErr := json.Unmarshal(processed[0].Schema, &schema); jsonErr != nil {
			bundled[i] = bundledGroup{group: group, bundleErr: fmt.Errorf("failed to unmarshal processed schema: %w", jsonErr)}
			continue
		}

		bundled[i] = bundledGroup{group: group, bundledSchema: schema}
	}

	return bundled, nil
}

// markBundleErrors marks bundle failures as test failures and unsupported keyword errors as unsupported features.
func markBundleErrors(bundled []bundledGroup, groupFilters []groupFilter, keywordResult *KeywordResult, summary *DraftSummary, draft, keyword string) {
	for i, bg := range bundled {
		// Skip groups where all tests were already marked as known issues
		if groupFilters[i].allKnown {
			continue
		}

		// Skip groups with adapter-specific limitations (already marked as unsupported in markAdapterLimitations)
		if groupFilters[i].adapterLimitation {
			continue
		}

		// Handle unsupported keyword errors - mark as unsupported features, not failures
		if bg.unsupportedErr != nil {
			for _, tc := range bg.group.Tests {
				testPath := fmt.Sprintf("%s/%s/%s/%s", draft, keyword, bg.group.Description, tc.Description)
				summary.UnsupportedFeatures.Count++
				summary.UnsupportedFeatures.Items = append(summary.UnsupportedFeatures.Items, UnsupportedFeatureItem{
					Path:   testPath,
					Reason: bg.unsupportedErr.Error(),
				})
			}
			continue
		}

		// Handle regular bundle errors - mark as failures
		if bg.bundleErr != nil {
			markAllFailed(keywordResult, summary, bg.group, fmt.Sprintf("bundling error: %v", bg.bundleErr))
		}
	}
}

// callAdapter calls the adapter for all successfully bundled schemas.
// Returns nil map if all groups failed bundling.
func callAdapter(ctx context.Context, bundled []bundledGroup, opts runDraftOptions) (map[int]*adapter.ConvertResult, error) {
	var adapterInputs []adapter.ConvertInput
	var inputIndexes []int

	for i, bg := range bundled {
		// Skip groups with bundle errors or unsupported keyword errors
		if bg.bundleErr == nil && bg.unsupportedErr == nil {
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
		return nil, nil
	}

	adapterStart := time.Now()
	adapterOutputs, err := CallAdapterBatch(ctx, opts.adapterBin, opts.runner, opts.runnerArgs, adapterInputs)
	opts.timing.addAdapterInvocation(time.Since(adapterStart))
	if err != nil {
		return nil, fmt.Errorf("adapter call failed: %w", err)
	}

	result := make(map[int]*adapter.ConvertResult)
	for i, outIdx := range inputIndexes {
		result[outIdx] = &adapterOutputs[i]
	}
	return result, nil
}

// buildHarnessItems creates harness items from adapter outputs.
func buildHarnessItems(
	bundled []bundledGroup,
	groupFilters []groupFilter,
	adapterOutputByGroup map[int]*adapter.ConvertResult,
	keywordResult *KeywordResult,
	summary *DraftSummary,
) ([]HarnessItem, map[string]*TestGroup, map[string][]int) {
	var harnessItems []HarnessItem
	groupByID := make(map[string]*TestGroup)
	filteredTestsByGroup := make(map[string][]int)

	for i, bg := range bundled {
		// Skip groups with bundle errors or unsupported keyword errors
		if bg.bundleErr != nil || bg.unsupportedErr != nil {
			continue
		}

		adapterOutput := adapterOutputByGroup[i]
		if adapterOutput == nil {
			markAllFailed(keywordResult, summary, bg.group, "adapter did not return output for this group")
			continue
		}

		gf := groupFilters[i]

		if len(gf.filteredTests) == 0 {
			continue
		}

		groupID := fmt.Sprintf("group_%d", i)
		harnessItems = append(harnessItems, HarnessItem{
			GroupID:       groupID,
			AdapterOutput: adapterOutput,
			Tests:         gf.filteredTests,
		})
		groupByID[groupID] = &bg.group
		filteredTestsByGroup[groupID] = gf.filteredIndices
	}

	return harnessItems, groupByID, filteredTestsByGroup
}

// executeHarness generates and runs the test harness.
// Returns nil results if harness failed (errors already recorded).
func executeHarness(
	ctx context.Context,
	harnessItems []HarnessItem,
	groupByID map[string]*TestGroup,
	opts runDraftOptions,
	keywordResult *KeywordResult,
	summary *DraftSummary,
) ([]HarnessResult, error) {
	harnessStart := time.Now()
	tempHarness, err := GenerateHarness(opts.language, harnessItems, opts.workDir)
	opts.timing.addHarnessGeneration(time.Since(harnessStart))
	if err != nil {
		for _, item := range harnessItems {
			if group := groupByID[item.GroupID]; group != nil {
				markAllFailed(keywordResult, summary, *group, fmt.Sprintf("harness generation error: %v", err))
			}
		}
		return nil, nil
	}
	defer os.Remove(tempHarness)

	execStart := time.Now()
	harnessResults, err := ExecuteHarness(ctx, tempHarness, opts.runner, opts.runnerArgs, opts.workDir)
	opts.timing.addHarnessExecution(time.Since(execStart))

	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return nil, err
		}
		for _, item := range harnessItems {
			if group := groupByID[item.GroupID]; group != nil {
				markAllFailed(keywordResult, summary, *group, fmt.Sprintf("harness execution error: %v", err))
			}
		}
		return nil, nil
	}

	return harnessResults, nil
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

func processHarnessResults(harnessResults []HarnessResult, groupByID map[string]*TestGroup, filteredTestsByGroup map[string][]int, keywordResult *KeywordResult, summary *DraftSummary) {
	for _, hr := range harnessResults {
		group := groupByID[hr.GroupID]
		if group == nil {
			continue // unknown group, skip
		}

		// Map filtered index back to original index
		filteredIndices := filteredTestsByGroup[hr.GroupID]
		if hr.Index >= len(filteredIndices) {
			continue // invalid index, skip
		}
		originalIdx := filteredIndices[hr.Index]
		if originalIdx >= len(group.Tests) {
			continue // invalid original index, skip
		}
		tc := group.Tests[originalIdx]

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
