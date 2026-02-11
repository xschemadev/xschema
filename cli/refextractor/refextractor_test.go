package refextractor

import (
	"encoding/json"
	"slices"
	"sort"
	"testing"
)

func TestExtractExternalRefs(t *testing.T) {
	tests := []struct {
		name     string
		schema   string
		baseURI  string
		wantRefs []string
	}{
		{
			name:     "simple $ref",
			schema:   `{"$ref": "http://example.com/other.json"}`,
			baseURI:  "",
			wantRefs: []string{"http://example.com/other.json"},
		},
		{
			name:     "relative $ref with base",
			schema:   `{"$ref": "other.json"}`,
			baseURI:  "http://example.com/schemas/base.json",
			wantRefs: []string{"http://example.com/schemas/other.json"},
		},
		{
			name:     "fragment-only $ref is local",
			schema:   `{"$ref": "#/$defs/Foo"}`,
			baseURI:  "http://example.com/base.json",
			wantRefs: nil,
		},
		{
			name:     "$ref with fragment strips fragment",
			schema:   `{"$ref": "http://example.com/other.json#/definitions/Bar"}`,
			baseURI:  "",
			wantRefs: []string{"http://example.com/other.json"},
		},
		{
			name: "nested $refs",
			schema: `{
				"properties": {
					"foo": {"$ref": "http://example.com/foo.json"},
					"bar": {"$ref": "http://example.com/bar.json"}
				}
			}`,
			baseURI:  "",
			wantRefs: []string{"http://example.com/bar.json", "http://example.com/foo.json"},
		},
		{
			name: "$id changes base URI",
			schema: `{
				"$id": "http://example.com/schemas/",
				"properties": {
					"foo": {"$ref": "foo.json"}
				}
			}`,
			baseURI:  "http://other.com/base.json",
			wantRefs: []string{"http://example.com/schemas/foo.json"},
		},
		{
			name:     "custom $schema is extracted",
			schema:   `{"$schema": "http://example.com/my-meta.json", "type": "string"}`,
			baseURI:  "",
			wantRefs: []string{"http://example.com/my-meta.json"},
		},
		{
			name:     "standard $schema is not extracted",
			schema:   `{"$schema": "https://json-schema.org/draft/2020-12/schema", "type": "string"}`,
			baseURI:  "",
			wantRefs: nil,
		},
		{
			name:     "invalid JSON returns nil",
			schema:   `{not valid json`,
			baseURI:  "",
			wantRefs: nil,
		},
		{
			name:     "boolean schema returns nil",
			schema:   `true`,
			baseURI:  "",
			wantRefs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractExternalRefs(json.RawMessage(tt.schema), tt.baseURI)
			sort.Strings(got)
			sort.Strings(tt.wantRefs)

			if !slices.Equal(got, tt.wantRefs) {
				t.Errorf("ExtractExternalRefs() = %v, want %v", got, tt.wantRefs)
			}
		})
	}
}

func TestIsExternal(t *testing.T) {
	tests := []struct {
		ref  string
		want bool
	}{
		{"#/definitions/Foo", false},
		{"#", false},
		{"http://example.com/schema.json", true},
		{"other.json", true},
		{"./relative.json", true},
		{"../parent.json", true},
	}

	for _, tt := range tests {
		t.Run(tt.ref, func(t *testing.T) {
			if got := IsExternal(tt.ref); got != tt.want {
				t.Errorf("IsExternal(%q) = %v, want %v", tt.ref, got, tt.want)
			}
		})
	}
}
