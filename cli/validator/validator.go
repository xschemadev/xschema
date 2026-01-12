package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// ValidateOptions configures schema validation behavior
type ValidateOptions struct {
	// Metaschemas maps URI to raw JSON for custom metaschemas.
	// Pre-loaded metaschemas allow validation of schemas using custom $schema URIs.
	Metaschemas map[string]json.RawMessage
}

// noopLoader returns an error for all URL loads.
// External refs are handled by the bundler, not the validator.
type noopLoader struct{}

func (noopLoader) Load(url string) (any, error) {
	return nil, fmt.Errorf("external ref %s (will be resolved by bundler)", url)
}

// draftNames maps draft string names to jsonschema draft constants
var draftNames = map[string]*jsonschema.Draft{
	"draft4":       jsonschema.Draft4,
	"draft6":       jsonschema.Draft6,
	"draft7":       jsonschema.Draft7,
	"draft2019-09": jsonschema.Draft2019,
	"draft2020-12": jsonschema.Draft2020,
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
// If draftHint is provided and schema has no $schema field, it will be used instead of defaulting to draft2020-12.
func ValidateSchema(data []byte, draftHint ...string) error {
	return ValidateSchemaWithOptions(data, nil, draftHint...)
}

// ValidateSchemaWithOptions validates a schema with configurable options.
// Use this to pre-load custom metaschemas for schemas with custom $schema URIs.
func ValidateSchemaWithOptions(data []byte, opts *ValidateOptions, draftHint ...string) error {
	var draft *jsonschema.Draft
	if len(draftHint) > 0 && draftHint[0] != "" {
		if d, ok := draftNames[draftHint[0]]; ok {
			draft = d
		} else {
			draft = detectDraft(data)
		}
	} else {
		draft = detectDraft(data)
	}

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

	// Pre-load custom metaschemas via AddResource
	// Each metaschema is added by its URI so schemas using $schema can reference it
	if opts != nil && opts.Metaschemas != nil {
		for uri, raw := range opts.Metaschemas {
			metaDoc, err := jsonschema.UnmarshalJSON(strings.NewReader(string(raw)))
			if err != nil {
				return fmt.Errorf("invalid metaschema %s: %w", uri, err)
			}
			if err := compiler.AddResource(uri, metaDoc); err != nil {
				return fmt.Errorf("failed to add metaschema %s: %w", uri, err)
			}
		}
	}

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
