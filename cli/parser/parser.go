package parser

import "encoding/json"

type SourceType string

const (
	SourceURL    SourceType = "url"
	SourceFile   SourceType = "file"
	SourceSchema SourceType = "schema"
)

type AdapterRef struct {
	Name     string `json:"name"`
	Package  string `json:"package"`
	Language string `json:"language"`
}

type Declaration struct {
	Name     string          `json:"name"`
	Source   SourceType      `json:"source"`
	Location string          `json:"location,omitempty"` // URL or file path
	Schema   json.RawMessage `json:"schema,omitempty"`   // inline schema
	Adapter  AdapterRef      `json:"adapter"`
	File     string          `json:"file"`
	Line     int             `json:"line"`
}

// Parse finds all xschema declarations in the given directory
func Parse(dir string) ([]Declaration, error) {
	// TODO: implement
	return nil, nil
}
