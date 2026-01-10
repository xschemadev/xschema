package compliance

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Available draft versions in the test suite
// Note: v1 excluded until spec is released (metaschema URL returns 404)
var Drafts = []string{
	"draft2020-12",
	"draft2019-09",
	"draft7",
	"draft6",
	"draft4",
	"draft3",
}

// LoadTestSuite loads the JSON Schema Test Suite for a specific draft
// Returns a map of keyword name to test groups
func LoadTestSuite(suitePath, draft, keyword string) (map[string][]TestGroup, error) {
	testsDir := filepath.Join(suitePath, "tests", draft)

	if _, err := os.Stat(testsDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("test suite not found at %s", testsDir)
	}

	if keyword != "" {
		return loadKeywordSuite(testsDir, keyword)
	}

	return loadAllTestFiles(testsDir)
}

func loadAllTestFiles(testsDir string) (map[string][]TestGroup, error) {
	keywords, err := listKeywords(testsDir)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]TestGroup)
	for _, keyword := range keywords {
		filePath := filepath.Join(testsDir, keyword+".json")
		groups, err := loadTestFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s.json: %w", keyword, err)
		}

		result[keyword] = groups
	}

	return result, nil
}

func loadKeywordSuite(testsDir, keyword string) (map[string][]TestGroup, error) {
	filePath := filepath.Join(testsDir, keyword+".json")
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			available, listErr := listKeywords(testsDir)
			if listErr != nil {
				return nil, fmt.Errorf("keyword %q not found in %s", keyword, testsDir)
			}
			if len(available) == 0 {
				return nil, fmt.Errorf("keyword %q not found in %s (available: 0 keywords)", keyword, testsDir)
			}
			return nil, fmt.Errorf("keyword %q not found in %s (available: %s)", keyword, testsDir, strings.Join(available, ", "))
		}
		return nil, err
	}

	groups, err := loadTestFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load %s.json: %w", keyword, err)
	}

	return map[string][]TestGroup{keyword: groups}, nil
}

func listKeywords(testsDir string) ([]string, error) {
	entries, err := os.ReadDir(testsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read test directory: %w", err)
	}

	keywords := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		keywords = append(keywords, strings.TrimSuffix(entry.Name(), ".json"))
	}

	sort.Strings(keywords)

	return keywords, nil
}

func loadTestFile(path string) ([]TestGroup, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var groups []TestGroup
	if err := json.Unmarshal(data, &groups); err != nil {
		return nil, err
	}

	return groups, nil
}
