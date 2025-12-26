package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
func Generate(input GenerateBatchInput) ([]GenerateOutput, error) {
	runner, args, err := getRunner(input.Language)
	if err != nil {
		return nil, err
	}

	// Check runner exists
	if _, err := exec.LookPath(runner); err != nil {
		return nil, fmt.Errorf("%s not found: %w", runner, err)
	}

	cmdArgs := append(args, input.Adapter)
	cmd := exec.Command(runner, cmdArgs...)

	// Pipe schemas to stdin
	stdinData, err := json.Marshal(input.Schemas)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal schemas: %w", err)
	}
	cmd.Stdin = bytes.NewReader(stdinData)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("adapter %s failed: %w\n%s", input.Adapter, err, stderr.String())
	}

	var outputs []GenerateOutput
	if err := json.Unmarshal(stdout.Bytes(), &outputs); err != nil {
		return nil, fmt.Errorf("invalid output from %s: %w\noutput: %s", input.Adapter, err, stdout.String())
	}

	return outputs, nil
}

// getRunner returns the command and args to run an adapter based on language
func getRunner(language string) (string, []string, error) {
	switch language {
	case "typescript":
		return detectTSRunner()
	default:
		return "", nil, fmt.Errorf("unsupported language: %s", language)
	}
}

// detectTSRunner checks lockfiles to determine which runner to use
func detectTSRunner() (string, []string, error) {
	lockfiles := []struct {
		file string
		cmd  []string
	}{
		{"bun.lock", []string{"bunx"}},
		{"bun.lockb", []string{"bunx"}},
		{"pnpm-lock.yaml", []string{"pnpm", "exec"}},
		{"yarn.lock", []string{"yarn"}},
		{"package-lock.json", []string{"npx"}},
	}

	for _, lf := range lockfiles {
		if _, err := os.Stat(filepath.Join(".", lf.file)); err == nil {
			return lf.cmd[0], lf.cmd[1:], nil
		}
	}

	return "", nil, fmt.Errorf("no lockfile found (bun.lock, yarn.lock, pnpm-lock.yaml, or package-lock.json)")
}
