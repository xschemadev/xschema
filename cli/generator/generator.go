package generator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/xschemadev/xschema/language"
	"github.com/xschemadev/xschema/retriever"
	"github.com/xschemadev/xschema/ui"
)

// GenerateInput is sent to the adapter CLI
type GenerateInput struct {
	Namespace string          `json:"namespace"`
	ID        string          `json:"id"`
	Schema    json.RawMessage `json:"schema"`
}

// GenerateOutput is received from the adapter CLI
type GenerateOutput struct {
	Namespace string   `json:"namespace"`
	ID        string   `json:"id"`
	Schema    string   `json:"schema"`  // generated code expression
	Type      string   `json:"type"`    // type expression
	Imports   []string `json:"imports"` // required imports
}

// Key returns the full namespaced key like "namespace:id"
func (o GenerateOutput) Key() string {
	return o.Namespace + ":" + o.ID
}

// GenerateBatchInput groups schemas by adapter for batch processing
type GenerateBatchInput struct {
	Adapter  string // adapter package e.g., "zod"
	Language string // language name e.g., "typescript"
	Schemas  []retriever.RetrievedSchema
}

// Generate calls the adapter to convert schemas to native code
func Generate(ctx context.Context, input GenerateBatchInput) ([]GenerateOutput, error) {
	lang := language.ByName(input.Language)
	if lang == nil {
		return nil, fmt.Errorf("unsupported language: %s", input.Language)
	}

	runner, args, err := lang.DetectRunner()
	if err != nil {
		return nil, err
	}

	// Construct bin name: "zod" -> "xschema-zod"
	binName := lang.AdapterBinPrefix + input.Adapter

	ui.Verbosef("running adapter: %s (language: %s, runner: %s, schemas: %d)", binName, input.Language, runner, len(input.Schemas))

	// Check runner exists
	if _, err := exec.LookPath(runner); err != nil {
		ui.Verbosef("runner not found: %s", runner)
		return nil, fmt.Errorf("%s not found: %w", runner, err)
	}

	// Build input for adapter
	adapterInput := make([]GenerateInput, len(input.Schemas))
	for i, s := range input.Schemas {
		adapterInput[i] = GenerateInput{
			Namespace: s.Namespace,
			ID:        s.ID,
			Schema:    s.Schema,
		}
	}

	cmdArgs := append(args, binName)
	cmd := exec.CommandContext(ctx, runner, cmdArgs...)

	// Pipe schemas to stdin
	stdinData, err := json.Marshal(adapterInput)
	if err != nil {
		ui.Verbosef("failed to marshal schemas for adapter %s", binName)
		return nil, fmt.Errorf("failed to marshal schemas: %w", err)
	}
	cmd.Stdin = bytes.NewReader(stdinData)

	ui.Verbosef("executing adapter command: %s %v", runner, cmdArgs)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		ui.Verbosef("adapter execution failed: %s - %s", binName, stderr.String())
		return nil, fmt.Errorf("adapter %s failed: %w\n%s", binName, err, stderr.String())
	}

	var outputs []GenerateOutput
	if err := json.Unmarshal(stdout.Bytes(), &outputs); err != nil {
		ui.Verbosef("invalid adapter output from %s: %s", binName, stdout.String())
		return nil, fmt.Errorf("invalid output from %s: %w\noutput: %s", binName, err, stdout.String())
	}

	ui.Verbosef("adapter execution successful: %s (outputs: %d)", binName, len(outputs))
	return outputs, nil
}

// GenerateAll runs generation for all adapter groups and returns all outputs
func GenerateAll(ctx context.Context, schemas []retriever.RetrievedSchema, langName string) ([]GenerateOutput, error) {
	groups := retriever.GroupByAdapter(schemas)
	adapters := retriever.SortedAdapters(groups)

	var allOutputs []GenerateOutput

	for _, adapter := range adapters {
		batch := GenerateBatchInput{
			Adapter:  adapter,
			Language: langName,
			Schemas:  groups[adapter],
		}

		outputs, err := Generate(ctx, batch)
		if err != nil {
			return nil, err
		}

		allOutputs = append(allOutputs, outputs...)
	}

	return allOutputs, nil
}
