package parser

import (
	"encoding/json"

	"github.com/xschema/cli/language"
)

// SourceType represents how to retrieve a schema
type SourceType string

const (
	SourceURL  SourceType = "url"
	SourceFile SourceType = "file"
	SourceJSON SourceType = "json"
)

// ConfigFileRaw is the raw JSON structure of an xschema config file
type ConfigFileRaw struct {
	Schema    string           `json:"$schema"`
	Namespace string           `json:"namespace,omitempty"` // optional namespace override
	Schemas   []SchemaEntryRaw `json:"schemas"`
}

// SchemaEntryRaw represents one schema entry in a config file
type SchemaEntryRaw struct {
	ID         string          `json:"id"`
	SourceType SourceType      `json:"sourceType"` // "url", "file", "json"
	Source     json.RawMessage `json:"source"`     // string for url/file, object for json
	Adapter    string          `json:"adapter"`    // full package name e.g., "@xschema/zod"
}

// ConfigFile represents a parsed xschema config file
type ConfigFile struct {
	Path      string             // absolute path to config file
	Namespace string             // derived from filename or explicit
	Language  *language.Language // detected from $schema URL
	Schemas   []SchemaEntryRaw   // raw schema entries
}

// Declaration represents a schema declaration ready for retrieval
type Declaration struct {
	Namespace  string          // e.g., "user"
	ID         string          // e.g., "TestUrl"
	SourceType SourceType      // "url", "file", "json"
	Source     json.RawMessage // URL string, file path string, or inline JSON object
	Adapter    string          // full adapter package e.g., "@xschema/zod"
	ConfigPath string          // path to config file (for relative file resolution)
}

// Key returns the full namespaced key like "user:TestUrl"
func (d Declaration) Key() string {
	return d.Namespace + ":" + d.ID
}

// ParseResult contains all parsed config files and declarations
type ParseResult struct {
	Language     *language.Language // detected language (error if multiple)
	Configs      []ConfigFile       // all parsed config files
	Declarations []Declaration      // flattened declarations from all configs
}

// DeclarationsByNamespace groups declarations by namespace
func (r *ParseResult) DeclarationsByNamespace() map[string][]Declaration {
	result := make(map[string][]Declaration)
	for _, d := range r.Declarations {
		result[d.Namespace] = append(result[d.Namespace], d)
	}
	return result
}

// DeclarationsByAdapter groups declarations by adapter
func (r *ParseResult) DeclarationsByAdapter() map[string][]Declaration {
	result := make(map[string][]Declaration)
	for _, d := range r.Declarations {
		result[d.Adapter] = append(result[d.Adapter], d)
	}
	return result
}
