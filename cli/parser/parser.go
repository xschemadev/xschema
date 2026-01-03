package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tailscale/hujson"
	"github.com/xschema/cli/language"
	"github.com/xschema/cli/ui"
)

// Parse finds all xschema config files in the project and returns merged declarations
// langFilter can be empty (auto-detect) or a language name to filter by
func Parse(ctx context.Context, projectRoot string, langFilter string) (*ParseResult, error) {
	ui.Verbosef("parsing project: root=%s, langFilter=%s", projectRoot, langFilter)

	// Find all JSON/JSONC files
	files, err := getConfigFiles(ctx, projectRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to find config files: %w", err)
	}

	ui.Verbosef("found potential config files: count=%d", len(files))

	// Parse each file, filter by xschema.dev $schema
	var configs []ConfigFile
	var detectedLang *language.Language
	languageConflict := false

	for _, path := range files {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		config, err := parseConfigFile(path)
		if err != nil {
			ui.Verbosef("skipping file (parse error): path=%s, error=%v", path, err)
			continue
		}
		if config == nil {
			// Not an xschema config file
			continue
		}

		ui.Verbosef("found xschema config: path=%s, namespace=%s, language=%s, schemas=%d",
			path, config.Namespace, config.Language.Name, len(config.Schemas))

		// Check language consistency
		if detectedLang == nil {
			detectedLang = config.Language
		} else if detectedLang.Name != config.Language.Name {
			languageConflict = true
		}

		configs = append(configs, *config)
	}

	if len(configs) == 0 {
		return nil, fmt.Errorf("no xschema config files found in %s", projectRoot)
	}

	// Handle language filter/conflict
	if languageConflict {
		if langFilter == "" {
			// List detected languages
			langs := make(map[string]bool)
			for _, c := range configs {
				langs[c.Language.Name] = true
			}
			var langList []string
			for l := range langs {
				langList = append(langList, l)
			}
			return nil, fmt.Errorf("multiple languages detected (%s). Use --lang to specify which one to use",
				strings.Join(langList, ", "))
		}
		// Filter configs by language
		var filtered []ConfigFile
		for _, c := range configs {
			if c.Language.Name == langFilter {
				filtered = append(filtered, c)
			}
		}
		configs = filtered
		detectedLang = language.ByName(langFilter)
		if detectedLang == nil {
			return nil, fmt.Errorf("unknown language: %s", langFilter)
		}
	}

	// Merge declarations, checking for conflicts
	declarations, err := mergeDeclarations(configs)
	if err != nil {
		return nil, err
	}

	ui.Verbosef("parsed %d configs, %d declarations", len(configs), len(declarations))

	return &ParseResult{
		Language:     detectedLang,
		Configs:      configs,
		Declarations: declarations,
	}, nil
}

// getConfigFiles returns all JSON/JSONC files in the project
func getConfigFiles(ctx context.Context, projectRoot string) ([]string, error) {
	// Try git ls-files first
	ui.Verbosef("getting config files using git in %s", projectRoot)
	args := []string{"ls-files", "--cached", "--others", "--exclude-standard", "*.json", "*.jsonc"}
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = projectRoot
	output, err := cmd.Output()
	if err != nil {
		ui.Verbose("git not available, using directory walk")
		return walkDirForConfigs(ctx, projectRoot)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		ui.Verbosef("no files found via git in %s", projectRoot)
		return nil, nil
	}

	files := make([]string, 0, len(lines))
	for _, line := range lines {
		if line != "" {
			files = append(files, filepath.Join(projectRoot, line))
		}
	}
	ui.Verbosef("found files via git: count=%d", len(files))
	return files, nil
}

// walkDirForConfigs walks directory manually when git is not available
func walkDirForConfigs(ctx context.Context, projectRoot string) ([]string, error) {
	ui.Verbosef("walking directory for configs: %s", projectRoot)

	var files []string
	err := filepath.WalkDir(projectRoot, func(path string, d fs.DirEntry, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			// Skip common non-project directories
			if name == "node_modules" || name == ".git" || name == "__pycache__" ||
				name == ".venv" || name == "venv" || name == "vendor" ||
				name == "dist" || name == "build" {
				ui.Verbosef("skipping directory: %s", path)
				return filepath.SkipDir
			}
			return nil
		}

		ext := filepath.Ext(path)
		if ext == ".json" || ext == ".jsonc" {
			files = append(files, path)
		}
		return nil
	})

	ui.Verbosef("directory walk complete: files=%d", len(files))
	return files, err
}

// parseConfigFile parses a single config file
// Returns nil if file is not an xschema config (no matching $schema)
func parseConfigFile(path string) (*ConfigFile, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Standardize JSONC to JSON using hujson
	standardized, err := hujson.Standardize(content)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON/JSONC: %w", err)
	}

	// Parse JSON
	var raw ConfigFileRaw
	if err := json.Unmarshal(standardized, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Check if this is an xschema config file
	if !language.IsXSchemaURL(raw.Schema) {
		return nil, nil
	}

	// Detect language from $schema URL
	lang := language.BySchemaURL(raw.Schema)
	if lang == nil {
		return nil, fmt.Errorf("unknown xschema language in $schema: %s", raw.Schema)
	}

	// Derive namespace from filename or use explicit override
	namespace := raw.Namespace
	if namespace == "" {
		// Use filename without extension
		base := filepath.Base(path)
		ext := filepath.Ext(base)
		namespace = strings.TrimSuffix(base, ext)
	}

	return &ConfigFile{
		Path:      path,
		Namespace: namespace,
		Language:  lang,
		Schemas:   raw.Schemas,
	}, nil
}

// mergeDeclarations merges all config files into a flat list of declarations
// Same namespace from different files is merged; duplicate IDs within namespace are an error
func mergeDeclarations(configs []ConfigFile) ([]Declaration, error) {
	// Track seen IDs per namespace for duplicate detection
	seenIDs := make(map[string]map[string]string) // namespace -> id -> config path

	var declarations []Declaration

	for _, config := range configs {
		if seenIDs[config.Namespace] == nil {
			seenIDs[config.Namespace] = make(map[string]string)
		}

		for _, schema := range config.Schemas {
			// Check for duplicate ID in this namespace
			if existingPath, exists := seenIDs[config.Namespace][schema.ID]; exists {
				return nil, fmt.Errorf("duplicate schema ID %q in namespace %q: defined in both %s and %s",
					schema.ID, config.Namespace, existingPath, config.Path)
			}
			seenIDs[config.Namespace][schema.ID] = config.Path

			declarations = append(declarations, Declaration{
				Namespace:  config.Namespace,
				ID:         schema.ID,
				SourceType: schema.SourceType,
				Source:     schema.Source,
				Adapter:    schema.Adapter,
				ConfigPath: config.Path,
			})
		}
	}

	return declarations, nil
}
