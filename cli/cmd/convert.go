package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/xschemadev/xschema/fetcher"
	"github.com/xschemadev/xschema/generator"
	"github.com/xschemadev/xschema/processor"
	"github.com/xschemadev/xschema/retriever"
	"github.com/xschemadev/xschema/ui"
)

var (
	convertAdapter    string
	convertLang       string
	convertProject    string
	convertVerbose    bool
	convertAllowFetch bool
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert JSON schemas from stdin to native validators via stdout",
	Long: `Read JSON schema array from stdin, process and convert via an adapter, output JSON array to stdout.

All status/error output goes to stderr. Only the JSON result goes to stdout.

Examples:
  echo '[{"namespace":"test","id":"User","schema":{"type":"object","properties":{"name":{"type":"string"}}}}]' | xschema convert --adapter @xschemadev/zod

  cat schemas.json | xschema convert --adapter @xschemadev/zod --allow-fetch`,
	RunE: runConvert,
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVar(&convertAdapter, "adapter", "", "adapter package ref (e.g., @xschemadev/zod)")
	convertCmd.Flags().StringVar(&convertLang, "lang", "typescript", "target language")
	convertCmd.Flags().StringVar(&convertProject, "project", "", "project root directory (default: current directory)")
	convertCmd.Flags().BoolVarP(&convertVerbose, "verbose", "v", false, "show verbose output on stderr")
	convertCmd.Flags().BoolVar(&convertAllowFetch, "allow-fetch", false, "allow fetching external $ref URIs")

	_ = convertCmd.MarkFlagRequired("adapter")
}

// ConvertSchemaInput is the expected shape of each item in the stdin JSON array.
type ConvertSchemaInput struct {
	Namespace string          `json:"namespace"`
	ID        string          `json:"id"`
	Schema    json.RawMessage `json:"schema"`
}

func runConvert(cmd *cobra.Command, args []string) error {
	ui.SetVerbose(convertVerbose)
	ctx := cmd.Context()

	// Determine project root
	root := convertProject
	if root == "" {
		var err error
		root, err = os.Getwd()
		if err != nil {
			return writeJSONError(fmt.Errorf("failed to get current directory: %w", err))
		}
	}

	// Read stdin
	stdinData, err := io.ReadAll(os.Stdin)
	if err != nil {
		return writeJSONError(fmt.Errorf("failed to read stdin: %w", err))
	}
	if len(stdinData) == 0 {
		return writeJSONError(fmt.Errorf("no input provided on stdin"))
	}

	// Parse input
	var inputs []ConvertSchemaInput
	if err := json.Unmarshal(stdinData, &inputs); err != nil {
		return writeJSONError(fmt.Errorf("invalid JSON input: %w", err))
	}

	if len(inputs) == 0 {
		// Empty array -> empty array output
		fmt.Fprintln(os.Stdout, "[]")
		return nil
	}

	// Convert to retriever.RetrievedSchema with inline source type
	schemas := make([]retriever.RetrievedSchema, len(inputs))
	for i, in := range inputs {
		schemas[i] = retriever.RetrievedSchema{
			Namespace: in.Namespace,
			ID:        in.ID,
			Schema:    in.Schema,
			Adapter:   convertAdapter,
		}
	}

	// Process schemas (validate, crawl refs, bundle)
	sharedCache := fetcher.NewSharedCache()
	var f fetcher.Fetcher
	if convertAllowFetch {
		retOpts := retriever.DefaultOptions()
		retOpts.Cache = sharedCache
		f = newRetrieverFetcher(ctx, retOpts)
	} else {
		f = fetcher.FetchFunc(noFetchFetcher)
	}

	processed, err := processor.Process(ctx, schemas, processor.Options{
		Fetcher:   f,
		OnVerbose: convertVerboseCallback(),
		Cache:     sharedCache,
	})
	if err != nil {
		return writeJSONError(fmt.Errorf("processing failed: %w", err))
	}

	// Generate via adapter
	outputs, err := generator.GenerateAll(ctx, processed, convertLang, root)
	if err != nil {
		return writeJSONError(fmt.Errorf("generation failed: %w", err))
	}

	// Write JSON result to stdout
	result, err := json.Marshal(outputs)
	if err != nil {
		return writeJSONError(fmt.Errorf("failed to marshal output: %w", err))
	}
	fmt.Fprintln(os.Stdout, string(result))

	return nil
}

// noFetchFetcher returns an error explaining that external $refs require --allow-fetch.
func noFetchFetcher(_ context.Context, uri string) (json.RawMessage, error) {
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		return nil, fmt.Errorf("external $ref %q requires --allow-fetch flag", uri)
	}
	// File refs also blocked without --allow-fetch for convert (schemas should be self-contained)
	return nil, fmt.Errorf("external $ref %q requires --allow-fetch flag", uri)
}

// writeJSONError writes a JSON error object to stderr and returns the error for cobra exit code handling.
func writeJSONError(err error) error {
	errObj := map[string]string{"error": err.Error()}
	data, _ := json.Marshal(errObj)
	fmt.Fprintln(os.Stderr, string(data))
	return err
}

func convertVerboseCallback() func(string) {
	if !convertVerbose {
		return nil
	}
	return func(msg string) {
		ui.Verbosef("%s", msg)
	}
}
