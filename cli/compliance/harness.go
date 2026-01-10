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
)

// TypecheckConfig configures how typecheck runs
type TypecheckConfig struct {
	Enabled    bool     // whether typecheck is enabled
	Runner     string   // command to run tsc (e.g., "bun", "npx")
	RunnerArgs []string // args to pass before tsc (e.g., ["run"] for bun, ["tsc"] for npx)
}

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
		if strings.HasPrefix(entry.Name(), "harness.") && !entry.IsDir() {
			return filepath.Join(complianceDir, entry.Name()), nil
		}
	}

	return "", fmt.Errorf("no harness file found in %s (expected harness.*)", complianceDir)
}

// Language represents a target language for harness generation
type Language string

const (
	LanguageTypeScript Language = "typescript"
	LanguagePython     Language = "python"
)

// GetHarnessTemplate returns the harness template for the given language
func GetHarnessTemplate(lang Language) (string, error) {
	switch lang {
	case LanguageTypeScript:
		return TSHarnessTemplate, nil
	default:
		return "", fmt.Errorf("unsupported language: %s", lang)
	}
}

// GetHarnessExtension returns the file extension for the given language
func GetHarnessExtension(lang Language) string {
	switch lang {
	case LanguageTypeScript:
		return ".ts"
	case LanguagePython:
		return ".py"
	default:
		return ".txt"
	}
}

// GenerateTempHarness creates a temporary harness file using Go templates
// The file is created in targetDir so that package resolution works correctly
func GenerateTempHarness(lang Language, adapterOutput *AdapterOutput, testCases []TestCase, targetDir string) (string, error) {
	// Get template for language
	templateStr, err := GetHarnessTemplate(lang)
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("harness").Parse(templateStr)
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

	// Build template data
	data := HarnessTemplateData{
		Schema:          adapterOutput.Schema,
		Imports:         adapterOutput.Imports,
		Validate:        adapterOutput.Validate,
		ValidateImports: adapterOutput.ValidateImports,
		TestCasesString: string(testCasesString),
		IsTypeOnly:      adapterOutput.Validate == "",
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute harness template: %w", err)
	}

	// Create temp file in target directory (so package resolution works)
	ext := GetHarnessExtension(lang)
	tmpFile, err := os.CreateTemp(targetDir, "xschema-harness-*"+ext)
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
