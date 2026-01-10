package compliance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// FindHarness locates the harness file in an adapter's compliance directory
func FindHarness(adapterPath string) (string, error) {
	complianceDir := filepath.Join(adapterPath, "compliance")

	entries, err := os.ReadDir(complianceDir)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("compliance directory not found at %s", complianceDir)
		}
		return "", fmt.Errorf("failed to read compliance directory: %w", err)
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "harness.") && !strings.Contains(entry.Name(), ".batch.") && !entry.IsDir() {
			return filepath.Join(complianceDir, entry.Name()), nil
		}
	}

	return "", fmt.Errorf("no harness file found in %s (expected harness.*)", complianceDir)
}

// FindBatchHarness locates the batch harness file in an adapter's compliance directory
func FindBatchHarness(adapterPath string) (string, error) {
	complianceDir := filepath.Join(adapterPath, "compliance")

	entries, err := os.ReadDir(complianceDir)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("compliance directory not found at %s", complianceDir)
		}
		return "", fmt.Errorf("failed to read compliance directory: %w", err)
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "harness.batch.") && !entry.IsDir() {
			return filepath.Join(complianceDir, entry.Name()), nil
		}
	}

	return "", fmt.Errorf("no batch harness file found in %s (expected harness.batch.*)", complianceDir)
}

// GenerateTempHarness creates a temporary harness file with injected code and test cases
// The file is created in targetDir so that package resolution works correctly
func GenerateTempHarness(harnessTemplate, generatedCode string, testCases []TestCase, targetDir string) (string, error) {
	// Read template
	templateBytes, err := os.ReadFile(harnessTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to read harness template: %w", err)
	}
	template := string(templateBytes)

	// Serialize test cases to JSON
	testCasesJSON, err := json.Marshal(testCases)
	if err != nil {
		return "", fmt.Errorf("failed to serialize test cases: %w", err)
	}

	// Create a JS string literal containing the JSON
	// This ensures __proto__ and other special property names are preserved
	// (directly embedding as JS object literals would interpret __proto__ as prototype setter)
	testCasesString, err := json.Marshal(string(testCasesJSON))
	if err != nil {
		return "", fmt.Errorf("failed to serialize test cases string: %w", err)
	}

	// Replace placeholders
	content := template
	content = strings.ReplaceAll(content, "{{GENERATED_CODE}}", generatedCode)
	content = strings.ReplaceAll(content, "{{TEST_CASES_STRING}}", string(testCasesString))
	// Keep legacy support for {{TEST_CASES}} in case other adapters use it
	content = strings.ReplaceAll(content, "{{TEST_CASES}}", string(testCasesJSON))

	// Get extension from template
	ext := filepath.Ext(harnessTemplate)

	// Create temp file in target directory (so package resolution works)
	tmpFile, err := os.CreateTemp(targetDir, "xschema-harness-*"+ext)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}

	if _, err := tmpFile.WriteString(content); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to write temp file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), nil
}

// ExecuteHarness runs the harness file and returns the results
func ExecuteHarness(ctx context.Context, harnessFile string, runner string, runnerArgs []string, workDir string) ([]HarnessResult, error) {
	// Build command: e.g., "bun harness.ts"
	args := append(runnerArgs, harnessFile)
	cmd := exec.CommandContext(ctx, runner, args...)

	// Set working directory so bun can find dependencies
	if workDir != "" {
		cmd.Dir = workDir
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("harness execution failed: %w\nstderr: %s", err, stderr.String())
	}

	// Parse JSON output
	var results []HarnessResult
	if err := json.Unmarshal(stdout.Bytes(), &results); err != nil {
		return nil, fmt.Errorf("failed to parse harness output: %w\nstdout: %s", err, stdout.String())
	}

	return results, nil
}

// AdapterInput represents a single schema input for the adapter
type AdapterInput struct {
	Namespace string    `json:"namespace"`
	ID        string    `json:"id"`
	Schema    RawSchema `json:"schema"`
}

// CallAdapter calls the adapter to convert a schema to code
func CallAdapter(ctx context.Context, adapterBin string, runner string, runnerArgs []string, schema RawSchema) (*AdapterOutput, error) {
	outputs, err := CallAdapterBatch(ctx, adapterBin, runner, runnerArgs, []AdapterInput{
		{Namespace: "compliance", ID: "Test", Schema: schema},
	})
	if err != nil {
		return nil, err
	}
	if len(outputs) == 0 {
		return nil, fmt.Errorf("adapter returned no output")
	}
	return &outputs[0], nil
}

// CallAdapterBatch calls the adapter to convert multiple schemas in a single process invocation.
// Each input must have a unique ID; outputs are returned in the same order as inputs.
func CallAdapterBatch(ctx context.Context, adapterBin string, runner string, runnerArgs []string, inputs []AdapterInput) ([]AdapterOutput, error) {
	if len(inputs) == 0 {
		return nil, nil
	}

	inputJSON, err := json.Marshal(inputs)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize input: %w", err)
	}

	args := append(runnerArgs, adapterBin)
	cmd := exec.CommandContext(ctx, runner, args...)
	cmd.Stdin = bytes.NewReader(inputJSON)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("adapter call failed: %w\nstderr: %s", err, stderr.String())
	}

	var outputs []AdapterOutput
	if err := json.Unmarshal(stdout.Bytes(), &outputs); err != nil {
		return nil, fmt.Errorf("failed to parse adapter output: %w\nstdout: %s", err, stdout.String())
	}

	// Build map by ID for lookup
	outputByID := make(map[string]*AdapterOutput, len(outputs))
	for i := range outputs {
		outputByID[outputs[i].ID] = &outputs[i]
	}

	// Return outputs in same order as inputs
	result := make([]AdapterOutput, len(inputs))
	for i, inp := range inputs {
		out, ok := outputByID[inp.ID]
		if !ok {
			return nil, fmt.Errorf("adapter did not return output for id %q", inp.ID)
		}
		result[i] = *out
	}

	return result, nil
}

// BatchHarnessItem represents a single group's data for batch harness generation
type BatchHarnessItem struct {
	GroupID string
	Code    string
	Tests   []TestCase
}

// GenerateBatchHarness creates a temporary batch harness file with all groups' schemas and tests
func GenerateBatchHarness(harnessTemplate string, items []BatchHarnessItem, targetDir string) (string, error) {
	templateBytes, err := os.ReadFile(harnessTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to read batch harness template: %w", err)
	}
	template := string(templateBytes)

	// Build schemas object: { "group_0": <code>, "group_1": <code>, ... }
	var schemasBuilder strings.Builder
	schemasBuilder.WriteString("{\n")
	for i, item := range items {
		if i > 0 {
			schemasBuilder.WriteString(",\n")
		}
		// Escape the groupID for use as object key
		keyJSON, _ := json.Marshal(item.GroupID)
		schemasBuilder.WriteString("  ")
		schemasBuilder.Write(keyJSON)
		schemasBuilder.WriteString(": ")
		schemasBuilder.WriteString(item.Code)
	}
	schemasBuilder.WriteString("\n}")

	// Build test data array: [{groupId, tests}, ...]
	testDataItems := make([]BatchTestData, len(items))
	for i, item := range items {
		testDataItems[i] = BatchTestData{
			GroupID: item.GroupID,
			Tests:   item.Tests,
		}
	}
	testDataJSON, err := json.Marshal(testDataItems)
	if err != nil {
		return "", fmt.Errorf("failed to serialize test data: %w", err)
	}
	// Wrap in JSON string for JS parsing
	testDataString, err := json.Marshal(string(testDataJSON))
	if err != nil {
		return "", fmt.Errorf("failed to serialize test data string: %w", err)
	}

	// Replace placeholders
	content := template
	content = strings.ReplaceAll(content, "{{BATCH_SCHEMAS}}", schemasBuilder.String())
	content = strings.ReplaceAll(content, "{{BATCH_TEST_DATA}}", string(testDataString))

	// Get extension from template
	ext := filepath.Ext(harnessTemplate)

	// Create temp file in target directory
	tmpFile, err := os.CreateTemp(targetDir, "xschema-batch-harness-*"+ext)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}

	if _, err := tmpFile.WriteString(content); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to write temp file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to close temp file: %w", err)
	}

	return tmpFile.Name(), nil
}

// ExecuteBatchHarness runs the batch harness file and returns the results
func ExecuteBatchHarness(ctx context.Context, harnessFile string, runner string, runnerArgs []string, workDir string) ([]BatchHarnessResult, error) {
	args := append(runnerArgs, harnessFile)
	cmd := exec.CommandContext(ctx, runner, args...)

	if workDir != "" {
		cmd.Dir = workDir
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("batch harness execution failed: %w\nstderr: %s", err, stderr.String())
	}

	var results []BatchHarnessResult
	if err := json.Unmarshal(stdout.Bytes(), &results); err != nil {
		return nil, fmt.Errorf("failed to parse batch harness output: %w\nstdout: %s", err, stdout.String())
	}

	return results, nil
}
