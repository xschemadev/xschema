package adapter

import "encoding/json"

// ConvertInput is sent to the adapter CLI
type ConvertInput struct {
	Namespace  string          `json:"namespace"`
	ID         string          `json:"id"`
	VarName    string          `json:"varName"`
	Schema     json.RawMessage `json:"schema"`
	Vocabulary map[string]bool `json:"vocabulary,omitempty"` // $vocabulary from custom metaschema (nil = all enabled)
}

// ConvertResult is received from the adapter CLI
type ConvertResult struct {
	Namespace       string   `json:"namespace"`
	ID              string   `json:"id"`
	VarName         string   `json:"varName"`
	Imports         []string `json:"imports"`
	Schema          string   `json:"schema,omitempty"`
	Type            string   `json:"type,omitempty"`
	Validate        string   `json:"validate,omitempty"`
	ValidateImports []string `json:"validateImports,omitempty"`
}

// Key returns the full namespaced key like "namespace:id"
func (r ConvertResult) Key() string {
	return r.Namespace + ":" + r.ID
}
