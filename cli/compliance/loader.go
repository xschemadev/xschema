package compliance

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
func LoadTestSuite(suitePath, draft string) (map[string][]TestGroup, error) {
	testsDir := filepath.Join(suitePath, "tests", draft)

	if _, err := os.Stat(testsDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("test suite not found at %s", testsDir)
	}

	result := make(map[string][]TestGroup)

	entries, err := os.ReadDir(testsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read test directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		keyword := strings.TrimSuffix(entry.Name(), ".json")
		filePath := filepath.Join(testsDir, entry.Name())

		groups, err := loadTestFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s: %w", entry.Name(), err)
		}

		result[keyword] = groups
	}

	return result, nil
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
