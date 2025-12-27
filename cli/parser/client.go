package parser

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

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
	Output      string // output directory, default: ".xschema"
	Concurrency int    // max concurrent HTTP requests, default: 10
	HTTPTimeout int    // HTTP timeout in ms, default: 30000
	Retries     int    // max retries, default: 3
}

// DefaultConfig returns default configuration values
func DefaultConfig() ClientConfig {
	return ClientConfig{
		Output:      ".xschema",
		Concurrency: 10,
		HTTPTimeout: 30000,
		Retries:     3,
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

	logger.Debug("parsed config", "output", config.Output, "concurrency", config.Concurrency)

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

// applyConfig sets config field based on key/value
func applyConfig(config *ClientConfig, key, value string) {
	switch key {
	case "output":
		config.Output = value
	case "concurrency":
		if v, err := parseIntValue(value); err == nil {
			config.Concurrency = v
		}
	case "httpTimeout", "http_timeout":
		if v, err := parseIntValue(value); err == nil {
			config.HTTPTimeout = v
		}
	case "retries":
		if v, err := parseIntValue(value); err == nil {
			config.Retries = v
		}
	}
}

func parseIntValue(s string) (int, error) {
	var v int
	_, err := fmt.Sscanf(s, "%d", &v)
	return v, err
}
