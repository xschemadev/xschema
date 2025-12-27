package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/xschema/cli/generator"
	"github.com/xschema/cli/injector"
	"github.com/xschema/cli/logger"
	"github.com/xschema/cli/parser"
	"github.com/xschema/cli/retriever"
)

var (
	clientFile string
	verbose    bool
	dryRun     bool
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Parse codebase, convert schemas, output native validators",
	RunE:  runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&clientFile, "client", "c", "", "path to client file (required)")
	generateCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	generateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "show what would be generated without writing")

	generateCmd.MarkFlagRequired("client")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	logger.SetLogger(logger.New(verbose))

	ctx := cmd.Context()

	// 1. Parse client file
	logger.Info("parsing client", "file", clientFile)
	client, err := parser.ParseClient(ctx, clientFile)
	if err != nil {
		logger.Error("failed to parse client", "error", err)
		return fmt.Errorf("parse client: %w", err)
	}

	logger.Info("found client", "name", client.ClientName, "language", client.Language.Name)
	logger.Debug("client config", "output", client.Config.Output, "concurrency", client.Config.Concurrency)

	// 2. Parse codebase for declarations
	logger.Info("parsing codebase", "language", client.Language.Name)
	decls, err := parser.Parse(ctx, ".", client)
	if err != nil {
		logger.Error("parse failed", "error", err)
		return fmt.Errorf("parse: %w", err)
	}

	logger.Info("found declarations", "count", len(decls))
	if len(decls) == 0 {
		logger.Warn("no xschema declarations found")
		return nil
	}

	// 3. Retrieve schemas
	retrieverOpts := retriever.Options{
		Concurrency: client.Config.Concurrency,
		HTTPTimeout: time.Duration(client.Config.HTTPTimeout) * time.Millisecond,
		Retries:     client.Config.Retries,
	}
	batches, err := retriever.Retrieve(ctx, decls, retrieverOpts)
	if err != nil {
		logger.Error("retrieve failed", "error", err)
		return fmt.Errorf("retrieve: %w", err)
	}

	// 4. Generate
	outputsByLang := make(map[string][]generator.GenerateOutput)
	for _, batch := range batches {
		if dryRun {
			logger.Info("dry run", "adapter", batch.Adapter, "language", batch.Language)
			printDryRun(batch)
			continue
		}

		outputs, err := generator.Generate(ctx, batch)
		if err != nil {
			logger.Error("generate failed", "adapter", batch.Adapter, "error", err)
			return fmt.Errorf("generate (%s): %w", batch.Adapter, err)
		}
		outputsByLang[batch.Language] = append(outputsByLang[batch.Language], outputs...)
	}

	if dryRun {
		return nil
	}

	// 5. Inject
	for lang, outputs := range outputsByLang {
		err := injector.Inject(injector.InjectInput{
			Language: lang,
			Outputs:  outputs,
			OutDir:   client.Config.Output,
		})
		if err != nil {
			logger.Error("inject failed", "language", lang, "error", err)
			return fmt.Errorf("inject (%s): %w", lang, err)
		}
	}

	logger.Info("complete", "output", client.Config.Output)
	return nil
}

func printDryRun(batch generator.GenerateBatchInput) {
	logger.Info("adapter batch", "adapter", batch.Adapter, "language", batch.Language, "schemas", len(batch.Schemas))
	for _, s := range batch.Schemas {
		var schema map[string]any
		json.Unmarshal(s.Schema, &schema)
		logger.Info("  - schema", "name", s.Name, "type", schema["type"])
	}
}
