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
	"text/template"

	"github.com/xschemadev/xschema/language"
)

// TypecheckConfig configures how typecheck runs
type TypecheckConfig struct {
	Enabled    bool     // whether typecheck is enabled
	Runner     string   // command to run tsc (e.g., "bun", "npx")
	RunnerArgs []string // args to pass before tsc (e.g., ["run"] for bun, ["tsc"] for npx)
}

// GenerateTempHarness creates a temporary harness file using Go templates
// The file is created in targetDir so that package resolution works correctly
func GenerateTempHarness(lang *language.Language, adapterOutput *AdapterOutput, testCases []TestCase, targetDir string) (string, error) {
	if lang.HarnessTemplate == "" {
		return "", fmt.Errorf("no harness template configured for language %s", lang.Name)
	}

	tmpl, err := template.New("harness").Parse(lang.HarnessTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse harness template: %w", err)
	}

	// Serialize test cases to JSON
	testCasesJSON, err := json.Marshal(testCases)
	if err != nil {
		return "", fmt.Errorf("failed to serialize test cases: %w", err)
	}

	// Create a JS string literal containing the JSON
	// This ensures __proto__ and other special property names are preserved
	testCasesString, err := json.Marshal(string(testCasesJSON))
	if err != nil {
		return "", fmt.Errorf("failed to serialize test cases string: %w", err)
	}

	mergedImports := append([]string{}, adapterOutput.Imports...)
	mergedImports = append(mergedImports, adapterOutput.ValidateImports...)

	formattedImports := ""
	if len(mergedImports) > 0 && lang.MergeImports != nil {
		formattedImports = lang.MergeImports(mergedImports)
	} else if len(mergedImports) > 0 {
		formattedImports = strings.Join(mergedImports, "\n")
	}

	// Build template data
	data := HarnessTemplateData{
		Schema:          adapterOutput.Schema,
		Type:            adapterOutput.Type,
		Imports:         formattedImports,
		Validate:        adapterOutput.Validate,
		TestCasesString: string(testCasesString),
		IsTypeOnly:      adapterOutput.Validate == "",
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
	// Build command: e.g., "bun run harness.ts"
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

// CallAdapter calls the adapter to convert a schema to code
func CallAdapter(ctx context.Context, adapterBin string, runner string, runnerArgs []string, schema RawSchema) (*AdapterOutput, error) {
	// Build input
	input := []map[string]interface{}{
		{
			"namespace": "compliance",
			"id":        "Test",
			"schema":    schema,
		},
	}

	inputJSON, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize input: %w", err)
	}

	// Build command: e.g., "bunx @xschemadev/zod"
	args := append(runnerArgs, adapterBin)
	cmd := exec.CommandContext(ctx, runner, args...)
	cmd.Stdin = bytes.NewReader(inputJSON)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("adapter call failed: %w\nstderr: %s", err, stderr.String())
	}

	// Parse output
	var outputs []AdapterOutput
	if err := json.Unmarshal(stdout.Bytes(), &outputs); err != nil {
		return nil, fmt.Errorf("failed to parse adapter output: %w\nstdout: %s", err, stdout.String())
	}

	if len(outputs) == 0 {
		return nil, fmt.Errorf("adapter returned no output")
	}

	return &outputs[0], nil
}

// TypecheckResult contains the result of typechecking a harness file
type TypecheckResult struct {
	Success bool
	Output  string // stderr from tsc (contains error messages)
}

// TypecheckHarness runs the TypeScript compiler on a harness file to catch type errors.
// The harness file must be a .ts file.
// workDir is used for dependency resolution (finding tsconfig.json and node_modules).
func TypecheckHarness(ctx context.Context, harnessFile string, config TypecheckConfig, workDir string) (*TypecheckResult, error) {
	if !config.Enabled {
		return &TypecheckResult{Success: true}, nil
	}

	// Only typecheck TypeScript files
	ext := filepath.Ext(harnessFile)
	if ext != ".ts" && ext != ".tsx" {
		return &TypecheckResult{Success: true}, nil
	}

	// Build command: e.g., "bun run tsc --noEmit --skipLibCheck <file>"
	// or "npx tsc --noEmit --skipLibCheck <file>"
	args := append([]string{}, config.RunnerArgs...)
	args = append(args, "--noEmit", "--skipLibCheck", harnessFile)

	cmd := exec.CommandContext(ctx, config.Runner, args...)

	// Set working directory for dependency resolution
	if workDir != "" {
		cmd.Dir = workDir
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	// Combine stdout and stderr (tsc outputs errors to stdout, not stderr)
	output := stdout.String()
	if stderr.Len() > 0 {
		if output != "" {
			output += "\n"
		}
		output += stderr.String()
	}

	if err != nil {
		// Type error or other tsc error - return result with output
		return &TypecheckResult{
			Success: false,
			Output:  strings.TrimSpace(output),
		}, nil
	}

	return &TypecheckResult{Success: true}, nil
}
