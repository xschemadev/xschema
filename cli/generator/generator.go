package generator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/xschemadev/xschema/adapter"
	"github.com/xschemadev/xschema/bundler"
	"github.com/xschemadev/xschema/language"
	"github.com/xschemadev/xschema/retriever"
	"github.com/xschemadev/xschema/ui"
)

// validateOutputs checks that each output has at least schema or type
func validateOutputs(outputs []adapter.ConvertResult, adapterName string) error {
	for _, output := range outputs {
		if output.Schema == "" && output.Type == "" {
			return fmt.Errorf("adapter %s returned neither schema nor type for %s", adapterName, output.Key())
		}
	}
	return nil
}

// GenerateBatchInput groups schemas by adapter for batch processing
type GenerateBatchInput struct {
	AdapterRef  string // adapter ref e.g., "@xschemadev/zod"
	Language    string // language name e.g., "typescript"
	ProjectRoot string // project root directory
	Schemas     []retriever.RetrievedSchema
}

// Generate calls the adapter to convert schemas to native code
func Generate(ctx context.Context, input GenerateBatchInput) ([]adapter.ConvertResult, error) {
	lang := language.ByName(input.Language)
	if lang == nil {
		return nil, fmt.Errorf("unsupported language: %s", input.Language)
	}

	if lang.AdapterInvoker == nil {
		return nil, fmt.Errorf("language %s does not have adapter invocation configured", input.Language)
	}

	cmdSpec, err := lang.AdapterInvoker.BuildAdapterCommand(ctx, language.AdapterCommandInput{
		ProjectRoot: input.ProjectRoot,
		AdapterRef:  input.AdapterRef,
	})
	if err != nil {
		return nil, err
	}
	if err := cmdSpec.Validate(); err != nil {
		return nil, fmt.Errorf("invalid adapter command: %w", err)
	}

	ui.Verbosef("running adapter: %s (language: %s, cmd: %s, schemas: %d)", input.AdapterRef, input.Language, cmdSpec.Cmd, len(input.Schemas))

	// Check command exists
	if _, err := exec.LookPath(cmdSpec.Cmd); err != nil {
		ui.Verbosef("command not found: %s", cmdSpec.Cmd)
		return nil, fmt.Errorf("%s not found: %w", cmdSpec.Cmd, err)
	}

	// Bundle schemas to resolve external $refs
	adapterInput := make([]adapter.ConvertInput, len(input.Schemas))
	for i, s := range input.Schemas {
		schema := s.Schema
		if s.SourceURI != "" {
			bundled, err := bundler.Bundle(ctx, s.Schema, bundler.Options{
				BaseURI: s.SourceURI,
			})
			if err != nil {
				ui.Verbosef("failed to bundle schema %s: %v", s.Key(), err)
				return nil, fmt.Errorf("failed to bundle schema %s: %w", s.Key(), err)
			}
			schema = bundled
		}
		adapterInput[i] = adapter.ConvertInput{
			Namespace: s.Namespace,
			ID:        s.ID,
			Schema:    schema,
		}
	}

	cmd := exec.CommandContext(ctx, cmdSpec.Cmd, cmdSpec.Args...)
	if cmdSpec.Dir != "" {
		cmd.Dir = cmdSpec.Dir
	}
	if len(cmdSpec.Env) > 0 {
		cmd.Env = append(os.Environ(), cmdSpec.Env...)
	}

	// Pipe schemas to stdin
	stdinData, err := json.Marshal(adapterInput)
	if err != nil {
		ui.Verbosef("failed to marshal schemas for adapter %s", input.AdapterRef)
		return nil, fmt.Errorf("failed to marshal schemas: %w", err)
	}
	cmd.Stdin = bytes.NewReader(stdinData)

	ui.Verbosef("executing adapter command: %s %v", cmdSpec.Cmd, cmdSpec.Args)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		ui.Verbosef("adapter execution failed: %s - %s", input.AdapterRef, stderr.String())
		return nil, fmt.Errorf("adapter %s failed: %w\n%s", input.AdapterRef, err, stderr.String())
	}

	var outputs []adapter.ConvertResult
	if err := json.Unmarshal(stdout.Bytes(), &outputs); err != nil {
		ui.Verbosef("invalid adapter output from %s: %s", input.AdapterRef, stdout.String())
		return nil, fmt.Errorf("invalid output from %s: %w\noutput: %s", input.AdapterRef, err, stdout.String())
	}

	if err := validateOutputs(outputs, input.AdapterRef); err != nil {
		return nil, err
	}

	ui.Verbosef("adapter execution successful: %s (outputs: %d)", input.AdapterRef, len(outputs))
	return outputs, nil
}

// GenerateAll runs generation for all adapter groups and returns all outputs
func GenerateAll(ctx context.Context, schemas []retriever.RetrievedSchema, langName string, projectRoot string) ([]adapter.ConvertResult, error) {
	groups := retriever.GroupByAdapter(schemas)
	adapters := retriever.SortedAdapters(groups)

	var allOutputs []adapter.ConvertResult

	for _, adapterRef := range adapters {
		batch := GenerateBatchInput{
			AdapterRef:  adapterRef,
			Language:    langName,
			ProjectRoot: projectRoot,
			Schemas:     groups[adapterRef],
		}

		outputs, err := Generate(ctx, batch)
		if err != nil {
			return nil, err
		}

		allOutputs = append(allOutputs, outputs...)
	}

	return allOutputs, nil
}
