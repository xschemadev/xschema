package generator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/xschema/cli/language"
	"github.com/xschema/cli/logger"
)

type GenerateInput struct {
	Name   string          `json:"name"`
	Schema json.RawMessage `json:"schema"`
}

type GenerateBatchInput struct {
	Schemas  []GenerateInput `json:"schemas"`
	Adapter  string          `json:"adapter"`  // e.g. "@xschema/zod"
	Language string          `json:"language"` // e.g. "typescript", "python"
}

type GenerateOutput struct {
	Name    string   `json:"name"`
	Schema  string   `json:"schema"`
	Type    string   `json:"type"`
	Imports []string `json:"imports"`
}

// Generate calls the adapter to convert schemas to native code
func Generate(ctx context.Context, input GenerateBatchInput) ([]GenerateOutput, error) {
	runner, args, err := getRunner(input.Language)
	if err != nil {
		return nil, err
	}

	logger.Info("running adapter", "adapter", input.Adapter, "language", input.Language, "runner", runner, "schemas", len(input.Schemas))

	// Check runner exists
	if _, err := exec.LookPath(runner); err != nil {
		logger.Error("runner not found", "runner", runner, "error", err)
		return nil, fmt.Errorf("%s not found: %w", runner, err)
	}

	cmdArgs := append(args, input.Adapter)
	cmd := exec.CommandContext(ctx, runner, cmdArgs...)

	// Pipe schemas to stdin
	stdinData, err := json.Marshal(input.Schemas)
	if err != nil {
		logger.Error("failed to marshal schemas", "adapter", input.Adapter, "error", err)
		return nil, fmt.Errorf("failed to marshal schemas: %w", err)
	}
	cmd.Stdin = bytes.NewReader(stdinData)

	logger.Debug("executing adapter command", "adapter", input.Adapter, "args", cmdArgs)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		logger.Error("adapter execution failed", "adapter", input.Adapter, "error", err, "stderr", stderr.String())
		return nil, fmt.Errorf("adapter %s failed: %w\n%s", input.Adapter, err, stderr.String())
	}

	var outputs []GenerateOutput
	if err := json.Unmarshal(stdout.Bytes(), &outputs); err != nil {
		logger.Error("invalid adapter output", "adapter", input.Adapter, "error", err, "output", stdout.String())
		return nil, fmt.Errorf("invalid output from %s: %w\noutput: %s", input.Adapter, err, stdout.String())
	}

	logger.Info("adapter execution successful", "adapter", input.Adapter, "outputs", len(outputs))
	return outputs, nil
}

// getRunner returns the command and args to run an adapter based on language
func getRunner(langName string) (string, []string, error) {
	lang := language.ByName(langName)
	if lang == nil {
		return "", nil, fmt.Errorf("unsupported language: %s", langName)
	}
	return lang.DetectRunner()
}
