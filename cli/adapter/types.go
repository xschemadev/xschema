package adapter

import "encoding/json"

// ConvertInput is sent to the adapter CLI
type ConvertInput struct {
	Namespace string          `json:"namespace"`
	ID        string          `json:"id"`
	Schema    json.RawMessage `json:"schema"`
}

// ConvertResult is received from the adapter CLI
type ConvertResult struct {
	Namespace       string   `json:"namespace"`
	ID              string   `json:"id"`
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
