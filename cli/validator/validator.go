package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// noopLoader returns an error for all URL loads.
// External refs are handled by the bundler, not the validator.
type noopLoader struct{}

func (noopLoader) Load(url string) (any, error) {
	return nil, fmt.Errorf("external ref %s (will be resolved by bundler)", url)
}

// draftURLs maps $schema URLs to jsonschema draft constants
var draftURLs = map[string]*jsonschema.Draft{
	"https://json-schema.org/draft/2020-12/schema": jsonschema.Draft2020,
	"http://json-schema.org/draft/2020-12/schema":  jsonschema.Draft2020,
	"https://json-schema.org/draft/2019-09/schema": jsonschema.Draft2019,
	"http://json-schema.org/draft/2019-09/schema":  jsonschema.Draft2019,
	"http://json-schema.org/draft-07/schema#":      jsonschema.Draft7,
	"https://json-schema.org/draft-07/schema#":     jsonschema.Draft7,
	"http://json-schema.org/draft-06/schema#":      jsonschema.Draft6,
	"https://json-schema.org/draft-06/schema#":     jsonschema.Draft6,
	"http://json-schema.org/draft-04/schema#":      jsonschema.Draft4,
	"https://json-schema.org/draft-04/schema#":     jsonschema.Draft4,
}

// detectDraft detects the JSON Schema draft from the $schema field
func detectDraft(data []byte) *jsonschema.Draft {
	var schema struct {
		Schema string `json:"$schema"`
	}
	if err := json.Unmarshal(data, &schema); err != nil || schema.Schema == "" {
		return jsonschema.Draft2020 // default
	}

	// exact match
	if draft, ok := draftURLs[schema.Schema]; ok {
		return draft
	}

	// pattern match for variations
	s := schema.Schema
	if strings.Contains(s, "draft/2020-12") || strings.Contains(s, "draft-2020-12") {
		return jsonschema.Draft2020
	}
	if strings.Contains(s, "draft/2019-09") || strings.Contains(s, "draft-2019-09") {
		return jsonschema.Draft2019
	}
	if strings.Contains(s, "draft-07") {
		return jsonschema.Draft7
	}
	if strings.Contains(s, "draft-06") {
		return jsonschema.Draft6
	}
	if strings.Contains(s, "draft-04") {
		return jsonschema.Draft4
	}

	return jsonschema.Draft2020 // default to latest
}

// ValidateSchema validates that the given JSON bytes represent a valid JSON Schema.
// Returns nil if valid, error with details if invalid.
// External $ref loading errors are ignored (handled by bundler later).
func ValidateSchema(data []byte) error {
	draft := detectDraft(data)

	// Parse JSON first using library's unmarshaler
	doc, err := jsonschema.UnmarshalJSON(strings.NewReader(string(data)))
	if err != nil {
		return fmt.Errorf("invalid JSON Schema: %w", err)
	}

	// The jsonschema library validates against the meta-schema during compilation
	compiler := jsonschema.NewCompiler()
	compiler.DefaultDraft(draft)
	// Use noop loader - external refs are handled by bundler
	compiler.UseLoader(noopLoader{})

	if err := compiler.AddResource("schema.json", doc); err != nil {
		return fmt.Errorf("invalid JSON Schema: %w", err)
	}

	_, err = compiler.Compile("schema.json")
	if err != nil {
		// Ignore external ref loading errors - bundler handles these
		var loadErr *jsonschema.LoadURLError
		if errors.As(err, &loadErr) {
			return nil
		}
		return fmt.Errorf("invalid JSON Schema: %w", err)
	}

	return nil
}
