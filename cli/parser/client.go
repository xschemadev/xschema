package parser

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/xschema/cli/language"
	"github.com/xschema/cli/logger"
)

// ClientInfo holds parsed client information
type ClientInfo struct {
	File       string // path to client file
	Language   *language.Language
	ClientName string // variable name (e.g., "xschema")
	Config     ClientConfig
}

// ClientConfig holds configuration from createXSchemaClient call
type ClientConfig struct {
	OutputDir          string // output directory, default: ".xschema"
	MaxParallelFetches int    // max concurrent HTTP requests, default: 10
	RequestTimeoutMs   int    // HTTP timeout in ms, default: 30000
	MaxFetchRetries    int    // max retries for fetching schemas, default: 3
}

// DefaultConfig returns default configuration values
func DefaultConfig() ClientConfig {
	return ClientConfig{
		OutputDir:          ".xschema",
		MaxParallelFetches: 10,
		RequestTimeoutMs:   30000,
		MaxFetchRetries:    3,
	}
}

// ParseClient parses the client file and extracts client info + config
func ParseClient(ctx context.Context, file string) (*ClientInfo, error) {
	// Determine language from extension
	ext := filepath.Ext(file)
	lang := language.ByExtension(ext)
	if lang == nil {
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}

	logger.Debug("parsing client file", "file", file, "language", lang.Name)

	// Read file
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read client file: %w", err)
	}

	// Parse with tree-sitter
	parser := sitter.NewParser()
	parser.SetLanguage(lang.GetSitterLang())

	tree, err := parser.ParseCtx(ctx, nil, content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse client file: %w", err)
	}

	// Find client variable name
	clientName, err := findClientName(tree, content, lang)
	if err != nil {
		return nil, err
	}
	if clientName == "" {
		return nil, fmt.Errorf("no %s call found in %s", lang.ClientFactory, file)
	}

	logger.Debug("found client", "name", clientName)

	// Parse config
	config := DefaultConfig()
	if err := parseClientConfig(tree, content, lang, &config); err != nil {
		return nil, err
	}

	// Resolve outputDir relative to client file's directory
	clientDir := filepath.Dir(file)
	config.OutputDir = filepath.Join(clientDir, config.OutputDir)

	logger.Debug("parsed config", "outputDir", config.OutputDir, "maxParallelFetches", config.MaxParallelFetches)

	return &ClientInfo{
		File:       file,
		Language:   lang,
		ClientName: clientName,
		Config:     config,
	}, nil
}

// findClientName extracts the variable name assigned to createXSchemaClient
func findClientName(tree *sitter.Tree, content []byte, lang *language.Language) (string, error) {
	if lang.ClientQuery == "" {
		return "", nil
	}

	q, err := sitter.NewQuery([]byte(lang.ClientQuery), lang.GetSitterLang())
	if err != nil {
		return "", fmt.Errorf("failed to compile client query: %w", err)
	}

	qc := sitter.NewQueryCursor()
	qc.Exec(q, tree.RootNode())

	for {
		match, ok := qc.NextMatch()
		if !ok {
			break
		}

		match = qc.FilterPredicates(match, content)
		if len(match.Captures) == 0 {
			continue
		}

		for _, cap := range match.Captures {
			capName := q.CaptureNameForId(cap.Index)
			if capName == "client_name" {
				return cap.Node.Content(content), nil
			}
		}
	}

	return "", nil
}

// parseClientConfig extracts config values from createXSchemaClient call
func parseClientConfig(tree *sitter.Tree, content []byte, lang *language.Language, config *ClientConfig) error {
	if lang.ConfigQuery == "" {
		return nil
	}

	q, err := sitter.NewQuery([]byte(lang.ConfigQuery), lang.GetSitterLang())
	if err != nil {
		return fmt.Errorf("failed to compile config query: %w", err)
	}

	qc := sitter.NewQueryCursor()
	qc.Exec(q, tree.RootNode())

	for {
		match, ok := qc.NextMatch()
		if !ok {
			break
		}

		match = qc.FilterPredicates(match, content)
		if len(match.Captures) == 0 {
			continue
		}

		var key, value string
		for _, cap := range match.Captures {
			capName := q.CaptureNameForId(cap.Index)
			text := cap.Node.Content(content)

			switch capName {
			case "config_key":
				key = unquoteString(text)
			case "config_value":
				value = unquoteString(text)
			}
		}

		if key != "" && value != "" {
			applyConfig(config, key, value)
		}
	}

	return nil
}

// normalizeKey converts any case convention to lowercase without separators
// e.g. "outputDir", "output_dir", "OutputDir" -> "outputdir"
func normalizeKey(key string) string {
	return strings.ToLower(strings.ReplaceAll(key, "_", ""))
}

// applyConfig sets config field based on key/value
func applyConfig(config *ClientConfig, key, value string) {
	switch normalizeKey(key) {
	case "outputdir":
		config.OutputDir = value
	case "maxparallelfetches":
		if v, err := parseIntValue(value); err == nil {
			config.MaxParallelFetches = v
		}
	case "requesttimeoutms":
		if v, err := parseIntValue(value); err == nil {
			config.RequestTimeoutMs = v
		}
	case "maxfetchretries":
		if v, err := parseIntValue(value); err == nil {
			config.MaxFetchRetries = v
		}
	}
}

func parseIntValue(s string) (int, error) {
	var v int
	_, err := fmt.Sscanf(s, "%d", &v)
	return v, err
}
