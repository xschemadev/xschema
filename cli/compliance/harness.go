package compliance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/xschemadev/xschema/adapter"
	"github.com/xschemadev/xschema/language"
)

// HarnessItem represents a single group's data for harness generation
type HarnessItem struct {
	GroupID       string
	AdapterOutput *adapter.ConvertResult
	Tests         []TestCase
}

// GenerateHarness creates a temporary harness file using Go templates
// The file is created in targetDir so that package resolution works correctly
func GenerateHarness(lang *language.Language, items []HarnessItem, targetDir string) (string, error) {
	if lang.HarnessTemplate == "" {
		return "", fmt.Errorf("no harness template configured for language %s", lang.Name)
	}

	if len(items) == 0 {
		return "", fmt.Errorf("no items to generate harness for")
	}

	tmpl, err := template.New("harness").Parse(lang.HarnessTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse harness template: %w", err)
	}

	// Collect all imports and build schema entries
	var allImports []string
	schemas := make([]HarnessSchemaEntry, len(items))
	testData := make([]BatchTestData, len(items))

	for i, item := range items {
		// Collect imports
		allImports = append(allImports, item.AdapterOutput.Imports...)
		allImports = append(allImports, item.AdapterOutput.ValidateImports...)

		// Build schema entry
		schemas[i] = HarnessSchemaEntry{
			GroupID:    item.GroupID,
			Schema:     item.AdapterOutput.Schema,
			Type:       item.AdapterOutput.Type,
			Validate:   item.AdapterOutput.Validate,
			IsTypeOnly: item.AdapterOutput.Validate == "",
		}

		// Build test data entry
		testData[i] = BatchTestData{
			GroupID: item.GroupID,
			Tests:   item.Tests,
		}
	}

	// Merge and format imports
	formattedImports := lang.MergeImports(allImports)

	// Serialize test data to JSON string
	testDataJSON, err := json.Marshal(testData)
	if err != nil {
		return "", fmt.Errorf("failed to serialize test data: %w", err)
	}
	testDataString, err := json.Marshal(string(testDataJSON))
	if err != nil {
		return "", fmt.Errorf("failed to serialize test data string: %w", err)
	}

	// Build template data
	data := HarnessTemplateData{
		Imports:        formattedImports,
		Schemas:        schemas,
		TestDataString: string(testDataString),
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute harness template: %w", err)
	}

	// Create temp file in target directory (so package resolution works)
	tmpFile, err := os.CreateTemp(targetDir, "xschema-harness-*"+lang.HarnessExtension)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}

	if _, err := tmpFile.Write(buf.Bytes()); err != nil {
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
	args := append(runnerArgs, harnessFile)
	cmd := exec.CommandContext(ctx, runner, args...)

	if workDir != "" {
		cmd.Dir = workDir
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("harness execution failed: %w\nstderr: %s", err, stderr.String())
	}

	var results []HarnessResult
	if err := json.Unmarshal(stdout.Bytes(), &results); err != nil {
		return nil, fmt.Errorf("failed to parse harness output: %w\nstdout: %s", err, stdout.String())
	}

	return results, nil
}

// CallAdapter calls the adapter to convert a schema to code
func CallAdapter(ctx context.Context, adapterBin string, runner string, runnerArgs []string, schema RawSchema) (*adapter.ConvertResult, error) {
	outputs, err := CallAdapterBatch(ctx, adapterBin, runner, runnerArgs, []adapter.ConvertInput{
		{Namespace: "compliance", ID: "Test", Schema: schema.Raw()},
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
func CallAdapterBatch(ctx context.Context, adapterBin string, runner string, runnerArgs []string, inputs []adapter.ConvertInput) ([]adapter.ConvertResult, error) {
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

	var outputs []adapter.ConvertResult
	if err := json.Unmarshal(stdout.Bytes(), &outputs); err != nil {
		return nil, fmt.Errorf("failed to parse adapter output: %w\nstdout: %s", err, stdout.String())
	}

	// Build map by ID for lookup
	outputByID := make(map[string]*adapter.ConvertResult, len(outputs))
	for i := range outputs {
		outputByID[outputs[i].ID] = &outputs[i]
	}

	// Return outputs in same order as inputs
	result := make([]adapter.ConvertResult, len(inputs))
	for i, inp := range inputs {
		out, ok := outputByID[inp.ID]
		if !ok {
			return nil, fmt.Errorf("adapter did not return output for id %q", inp.ID)
		}
		result[i] = *out
	}

	return result, nil
}
