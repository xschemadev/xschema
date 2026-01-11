package bundler

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func testdataPath(name string) string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "testdata", name)
}

func readTestFile(t *testing.T, name string) json.RawMessage {
	t.Helper()
	path := testdataPath(name)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read test file %s: %v", name, err)
	}
	return json.RawMessage(data)
}

func TestBundleSimpleSchema(t *testing.T) {
	ctx := context.Background()
	schema := readTestFile(t, "simple.json")

	bundled, err := Bundle(ctx, schema, DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should be unchanged (no external refs)
	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if result["type"] != "object" {
		t.Errorf("expected type=object, got %v", result["type"])
	}

	// Should not have $defs added (no external refs)
	if _, ok := result["$defs"]; ok {
		t.Error("should not have $defs for schema without external refs")
	}
}

func TestBundleLocalRef(t *testing.T) {
	ctx := context.Background()
	schema := readTestFile(t, "local-ref.json")

	bundled, err := Bundle(ctx, schema, DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// Local ref should be unchanged
	props := result["properties"].(map[string]any)
	userProp := props["user"].(map[string]any)
	if userProp["$ref"] != "#/$defs/User" {
		t.Errorf("local ref should be unchanged, got %v", userProp["$ref"])
	}
}

func TestBundleFileRef(t *testing.T) {
	ctx := context.Background()
	schema := readTestFile(t, "with-file-ref.json")

	opts := Options{
		BaseURI: testdataPath("with-file-ref.json"),
	}

	bundled, err := Bundle(ctx, schema, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// Should have $defs with the embedded schema
	defs, ok := result["$defs"].(map[string]any)
	if !ok {
		t.Fatal("expected $defs to be present")
	}

	if len(defs) != 1 {
		t.Errorf("expected 1 def, got %d", len(defs))
	}

	// The ref should be rewritten to local
	props := result["properties"].(map[string]any)
	addressProp := props["address"].(map[string]any)
	ref := addressProp["$ref"].(string)
	if !strings.HasPrefix(ref, "#/$defs/") {
		t.Errorf("expected ref to be rewritten to #/$defs/..., got %s", ref)
	}

	// Verify the embedded schema has the right structure
	var foundAddress bool
	for _, def := range defs {
		defObj := def.(map[string]any)
		if defObj["type"] == "object" {
			defProps, ok := defObj["properties"].(map[string]any)
			if ok && defProps["street"] != nil && defProps["city"] != nil {
				foundAddress = true
				break
			}
		}
	}
	if !foundAddress {
		t.Error("embedded address schema not found in $defs")
	}
}

func TestBundleFragmentRef(t *testing.T) {
	ctx := context.Background()
	schema := readTestFile(t, "with-fragment-ref.json")

	opts := Options{
		BaseURI: testdataPath("with-fragment-ref.json"),
	}

	bundled, err := Bundle(ctx, schema, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The ref should include the fragment
	props := result["properties"].(map[string]any)
	addressProp := props["address"].(map[string]any)
	ref := addressProp["$ref"].(string)

	// Should be something like #/$defs/key/definitions/Address
	if !strings.Contains(ref, "/definitions/Address") {
		t.Errorf("expected ref to contain /definitions/Address, got %s", ref)
	}
}

func TestBundleLocalhostMapping(t *testing.T) {
	ctx := context.Background()
	schema := readTestFile(t, "with-localhost-ref.json")

	opts := Options{
		BaseURI:     testdataPath("with-localhost-ref.json"),
		RemotesPath: testdataPath("remotes"),
	}

	bundled, err := Bundle(ctx, schema, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// Should have $defs with the integer schema
	defs, ok := result["$defs"].(map[string]any)
	if !ok {
		t.Fatal("expected $defs to be present")
	}

	// Verify the embedded schema is the integer type
	var foundInteger bool
	for _, def := range defs {
		defObj := def.(map[string]any)
		if defObj["type"] == "integer" {
			foundInteger = true
			break
		}
	}
	if !foundInteger {
		t.Error("embedded integer schema not found in $defs")
	}

	// The ref should be rewritten to local
	props := result["properties"].(map[string]any)
	countProp := props["count"].(map[string]any)
	ref := countProp["$ref"].(string)
	if !strings.HasPrefix(ref, "#/$defs/") {
		t.Errorf("expected ref to be rewritten to #/$defs/..., got %s", ref)
	}
}

func TestBundleInternalRefRewriting(t *testing.T) {
	ctx := context.Background()

	// Schema that references user-with-internal-refs.json (which has internal refs)
	schemaWithInternalRefs := `{
		"type": "object",
		"properties": {
			"user": { "$ref": "user-with-internal-refs.json" }
		}
	}`

	opts := Options{
		BaseURI: testdataPath("test.json"),
	}

	bundled, err := Bundle(ctx, json.RawMessage(schemaWithInternalRefs), opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The internal refs in the embedded schema should be rewritten
	defs := result["$defs"].(map[string]any)
	for _, def := range defs {
		defObj := def.(map[string]any)
		if props, ok := defObj["properties"].(map[string]any); ok {
			if addr, ok := props["address"].(map[string]any); ok {
				ref := addr["$ref"].(string)
				// The internal ref #/definitions/Address should be rewritten
				// to #/$defs/key/definitions/Address
				if !strings.HasPrefix(ref, "#/$defs/") {
					t.Errorf("internal ref should be rewritten, got %s", ref)
				}
				if !strings.Contains(ref, "/definitions/Address") {
					t.Errorf("internal ref should preserve path, got %s", ref)
				}
			}
		}
	}
}

func TestBundleMissingFile(t *testing.T) {
	ctx := context.Background()
	schema := `{"$ref": "nonexistent.json"}`

	opts := Options{
		BaseURI: testdataPath("test.json"),
	}

	_, err := Bundle(ctx, json.RawMessage(schema), opts)
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestBundleInvalidJSON(t *testing.T) {
	ctx := context.Background()
	schema := `{not valid json}`

	_, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestBundleContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	schema := `{"$ref": "http://example.com/schema.json"}`
	opts := Options{
		BaseURI: "http://example.com/base.json",
	}

	_, err := Bundle(ctx, json.RawMessage(schema), opts)
	if err == nil {
		t.Error("expected error for cancelled context")
	}
}

func TestSplitFragment(t *testing.T) {
	tests := []struct {
		uri          string
		wantBase     string
		wantFragment string
	}{
		{"foo.json", "foo.json", ""},
		{"foo.json#", "foo.json", ""},
		{"foo.json#/definitions/Bar", "foo.json", "/definitions/Bar"},
		{"http://example.com/schema.json#/defs/X", "http://example.com/schema.json", "/defs/X"},
		{"#/local", "", "/local"},
	}

	for _, tt := range tests {
		t.Run(tt.uri, func(t *testing.T) {
			base, fragment := splitFragment(tt.uri)
			if base != tt.wantBase {
				t.Errorf("base: got %q, want %q", base, tt.wantBase)
			}
			if fragment != tt.wantFragment {
				t.Errorf("fragment: got %q, want %q", fragment, tt.wantFragment)
			}
		})
	}
}

func TestResolveURI(t *testing.T) {
	tests := []struct {
		ref     string
		base    string
		want    string
		wantErr bool
	}{
		// Absolute URL stays absolute
		{"http://example.com/schema.json", "http://other.com/base.json", "http://example.com/schema.json", false},
		// Relative URL resolved against base
		{"other.json", "http://example.com/schemas/base.json", "http://example.com/schemas/other.json", false},
		// File paths
		{"other.json", "/home/user/schemas/base.json", "/home/user/schemas/other.json", false},
		// No base
		{"schema.json", "", "schema.json", false},
	}

	for _, tt := range tests {
		t.Run(tt.ref+"_"+tt.base, func(t *testing.T) {
			got, err := resolveURI(tt.ref, tt.base)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestURIToKey(t *testing.T) {
	b := &bundleContext{}

	tests := []struct {
		uri  string
		want string
	}{
		{"http://example.com/schema.json", "example_com_schema_json"},
		{"https://example.com/foo/bar.json", "example_com_foo_bar_json"},
		{"/home/user/schema.json", "_home_user_schema_json"},
		{"123start.json", "_123start_json"},
	}

	for _, tt := range tests {
		t.Run(tt.uri, func(t *testing.T) {
			got := b.uriToKey(tt.uri)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBundlePreservesExistingDefs(t *testing.T) {
	ctx := context.Background()
	// Schema with both existing $defs and external ref
	schema := `{
		"type": "object",
		"$defs": {
			"Existing": { "type": "string" }
		},
		"properties": {
			"existing": { "$ref": "#/$defs/Existing" },
			"external": { "$ref": "address.json" }
		}
	}`

	opts := Options{
		BaseURI: testdataPath("test.json"),
	}

	bundled, err := Bundle(ctx, json.RawMessage(schema), opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	defs := result["$defs"].(map[string]any)

	// Should have both existing and new def
	if _, ok := defs["Existing"]; !ok {
		t.Error("existing $def should be preserved")
	}

	if len(defs) < 2 {
		t.Errorf("expected at least 2 defs, got %d", len(defs))
	}
}

func TestBundleRejectsMissingInternalRef(t *testing.T) {
	ctx := context.Background()
	schema := `{
		"type": "object",
		"properties": {
			"user": { "$ref": "#/$defs/NonExistent" }
		}
	}`

	_, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err == nil {
		t.Error("expected error for missing internal ref")
	}
	if !strings.Contains(err.Error(), "missing target") {
		t.Errorf("expected 'missing target' in error, got: %v", err)
	}
}

func TestBundleRejectsDynamicRef(t *testing.T) {
	ctx := context.Background()
	schema := `{
		"type": "object",
		"$dynamicRef": "#foo"
	}`

	_, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err == nil {
		t.Error("expected error for $dynamicRef")
	}
	if !strings.Contains(err.Error(), "$dynamicRef") {
		t.Errorf("expected '$dynamicRef' in error, got: %v", err)
	}
}

func TestBundleRejectsDynamicAnchor(t *testing.T) {
	ctx := context.Background()
	schema := `{
		"type": "object",
		"$dynamicAnchor": "foo"
	}`

	_, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err == nil {
		t.Error("expected error for $dynamicAnchor")
	}
	if !strings.Contains(err.Error(), "$dynamicAnchor") {
		t.Errorf("expected '$dynamicAnchor' in error, got: %v", err)
	}
}

func TestBundleRejectsRecursiveRef(t *testing.T) {
	ctx := context.Background()
	schema := `{
		"type": "object",
		"$recursiveRef": "#"
	}`

	_, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err == nil {
		t.Error("expected error for $recursiveRef")
	}
	if !strings.Contains(err.Error(), "$recursiveRef") {
		t.Errorf("expected '$recursiveRef' in error, got: %v", err)
	}
}

func TestBundleRejectsRecursiveAnchor(t *testing.T) {
	ctx := context.Background()
	schema := `{
		"type": "object",
		"$recursiveAnchor": true
	}`

	_, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err == nil {
		t.Error("expected error for $recursiveAnchor")
	}
	if !strings.Contains(err.Error(), "$recursiveAnchor") {
		t.Errorf("expected '$recursiveAnchor' in error, got: %v", err)
	}
}

func TestBundleRejectsExternalRefWithoutBaseURI(t *testing.T) {
	ctx := context.Background()
	schema := `{
		"type": "object",
		"properties": {
			"user": { "$ref": "other.json" }
		}
	}`

	// No BaseURI set
	_, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err == nil {
		t.Error("expected error for external ref without base URI")
	}
	if !strings.Contains(err.Error(), "requires a base URI") {
		t.Errorf("expected 'requires a base URI' in error, got: %v", err)
	}
}

func TestBundleAllowsAbsoluteExternalRefWithoutBaseURI(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately so we don't actually fetch

	schema := `{
		"type": "object",
		"properties": {
			"user": { "$ref": "http://example.com/schema.json" }
		}
	}`

	// No BaseURI set but ref is absolute
	_, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	// Should fail with fetch error (context cancelled), not base URI error
	if err == nil {
		t.Error("expected error (context cancelled)")
	}
	if strings.Contains(err.Error(), "requires a base URI") {
		t.Errorf("should not require base URI for absolute refs, got: %v", err)
	}
}

func TestValidateJSONPointer(t *testing.T) {
	root := map[string]any{
		"type": "object",
		"$defs": map[string]any{
			"User": map[string]any{
				"type": "object",
			},
		},
		"properties": map[string]any{
			"name": map[string]any{"type": "string"},
		},
	}

	tests := []struct {
		ref     string
		wantErr bool
	}{
		{"#", false},                        // root ref
		{"#/$defs/User", false},             // valid path
		{"#/properties/name", false},        // valid nested path
		{"#/$defs/NonExistent", true},       // missing key
		{"#/properties/name/foo", true},     // traverse into primitive
		{"#/$defs/User/properties/x", true}, // path doesn't exist
	}

	for _, tt := range tests {
		t.Run(tt.ref, func(t *testing.T) {
			err := validateJSONPointer(tt.ref, root)
			if tt.wantErr && err == nil {
				t.Error("expected error")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestBundleInjectsSchemaForDraft(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		draft      string
		wantSchema string
	}{
		{"draft3", "http://json-schema.org/draft-03/schema#"},
		{"draft4", "http://json-schema.org/draft-04/schema#"},
		{"draft6", "http://json-schema.org/draft-06/schema#"},
		{"draft7", "http://json-schema.org/draft-07/schema#"},
		{"draft2019-09", "https://json-schema.org/draft/2019-09/schema"},
		{"draft2020-12", "https://json-schema.org/draft/2020-12/schema"},
	}

	for _, tt := range tests {
		t.Run(tt.draft, func(t *testing.T) {
			schema := `{"type": "string"}`
			opts := Options{Draft: tt.draft}

			bundled, err := Bundle(ctx, json.RawMessage(schema), opts)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			var result map[string]any
			if err := json.Unmarshal(bundled, &result); err != nil {
				t.Fatalf("failed to parse result: %v", err)
			}

			got, ok := result["$schema"].(string)
			if !ok {
				t.Fatalf("$schema not found in bundled schema")
			}
			if got != tt.wantSchema {
				t.Errorf("$schema: got %q, want %q", got, tt.wantSchema)
			}
		})
	}
}

func TestBundlePreservesExistingSchema(t *testing.T) {
	ctx := context.Background()

	// Schema already has $schema - should NOT be overwritten
	schema := `{"$schema": "http://json-schema.org/draft-07/schema#", "type": "string"}`
	opts := Options{Draft: "draft4"} // trying to inject draft4

	bundled, err := Bundle(ctx, json.RawMessage(schema), opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	got := result["$schema"].(string)
	// Should preserve the original draft7 schema, not inject draft4
	if got != "http://json-schema.org/draft-07/schema#" {
		t.Errorf("existing $schema should be preserved, got %q", got)
	}
}

func TestBundleNoDraftNoInjection(t *testing.T) {
	ctx := context.Background()

	// No draft specified - should not inject $schema
	schema := `{"type": "string"}`
	opts := DefaultOptions() // no Draft set

	bundled, err := Bundle(ctx, json.RawMessage(schema), opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if _, hasSchema := result["$schema"]; hasSchema {
		t.Error("should not inject $schema when no draft is specified")
	}
}

func TestValidateJSONPointerURIEncoded(t *testing.T) {
	// Schema with keys that need URI encoding in refs
	root := map[string]any{
		"$defs": map[string]any{
			"percent%field": map[string]any{"type": "string"},
			"slash/field":   map[string]any{"type": "number"},
			"tilde~field":   map[string]any{"type": "boolean"},
			"quote\"field":  map[string]any{"type": "integer"},
			"space field":   map[string]any{"type": "array"},
		},
	}

	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		// URI-encoded percent: %25 → %
		{"percent encoded", "#/$defs/percent%25field", false},
		// URI-encoded slash: %2F → / (also needs ~1 for JSON pointer)
		{"slash encoded", "#/$defs/slash~1field", false},
		// URI-encoded tilde: %7E → ~ (but JSON pointer uses ~0)
		{"tilde via json pointer", "#/$defs/tilde~0field", false},
		// URI-encoded quote: %22 → "
		{"quote encoded", "#/$defs/quote%22field", false},
		// URI-encoded space: %20 → space
		{"space encoded", "#/$defs/space%20field", false},
		// Direct keys (not encoded) should also work
		{"direct percent - should fail", "#/$defs/percent%field", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateJSONPointer(tt.ref, root)
			if tt.wantErr && err == nil {
				t.Error("expected error")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestBundleAnchorRef(t *testing.T) {
	ctx := context.Background()

	// Schema with $anchor and $ref to that anchor
	schema := `{
		"$ref": "#foo",
		"$defs": {
			"A": {
				"$anchor": "foo",
				"type": "integer"
			}
		}
	}`

	bundled, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The anchor ref should be rewritten to a JSON pointer
	ref := result["$ref"].(string)
	if ref != "#/$defs/A" {
		t.Errorf("expected $ref to be rewritten to #/$defs/A, got %s", ref)
	}
}

func TestBundleNestedAnchorRef(t *testing.T) {
	ctx := context.Background()

	// Schema with nested $anchor
	schema := `{
		"type": "object",
		"properties": {
			"value": { "$ref": "#bar" }
		},
		"$defs": {
			"outer": {
				"$defs": {
					"inner": {
						"$anchor": "bar",
						"type": "string"
					}
				}
			}
		}
	}`

	bundled, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The anchor ref should be rewritten to the nested JSON pointer
	props := result["properties"].(map[string]any)
	valueProp := props["value"].(map[string]any)
	ref := valueProp["$ref"].(string)
	if ref != "#/$defs/outer/$defs/inner" {
		t.Errorf("expected $ref to be rewritten to #/$defs/outer/$defs/inner, got %s", ref)
	}
}

func TestBundleMissingAnchorRef(t *testing.T) {
	ctx := context.Background()

	// Schema with $ref to non-existent anchor
	schema := `{
		"$ref": "#nonexistent"
	}`

	_, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err == nil {
		t.Error("expected error for missing anchor ref")
	}
	// Should fail validation (anchor not found, so ref stays as #nonexistent which isn't a valid JSON pointer)
}

func TestBundleMultipleAnchors(t *testing.T) {
	ctx := context.Background()

	// Schema with multiple anchors
	schema := `{
		"oneOf": [
			{ "$ref": "#first" },
			{ "$ref": "#second" }
		],
		"$defs": {
			"A": {
				"$anchor": "first",
				"type": "string"
			},
			"B": {
				"$anchor": "second",
				"type": "number"
			}
		}
	}`

	bundled, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	oneOf := result["oneOf"].([]any)
	firstRef := oneOf[0].(map[string]any)["$ref"].(string)
	secondRef := oneOf[1].(map[string]any)["$ref"].(string)

	if firstRef != "#/$defs/A" {
		t.Errorf("expected first ref to be #/$defs/A, got %s", firstRef)
	}
	if secondRef != "#/$defs/B" {
		t.Errorf("expected second ref to be #/$defs/B, got %s", secondRef)
	}
}

func TestBundleFragmentIDAsAnchor(t *testing.T) {
	ctx := context.Background()

	// Schema using draft4/6/7 style fragment $id as location-independent identifier
	schema := `{
		"allOf": [{ "$ref": "#foo" }],
		"definitions": {
			"A": {
				"$id": "#foo",
				"type": "integer"
			}
		}
	}`

	bundled, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The anchor ref should be rewritten to a JSON pointer
	allOf := result["allOf"].([]any)
	ref := allOf[0].(map[string]any)["$ref"].(string)
	if ref != "#/definitions/A" {
		t.Errorf("expected $ref to be rewritten to #/definitions/A, got %s", ref)
	}
}

func TestBundleLegacyIDAsAnchor(t *testing.T) {
	ctx := context.Background()

	// Schema using draft4 style "id" (not "$id") with fragment
	schema := `{
		"allOf": [{ "$ref": "#bar" }],
		"definitions": {
			"B": {
				"id": "#bar",
				"type": "string"
			}
		}
	}`

	bundled, err := Bundle(ctx, json.RawMessage(schema), DefaultOptions())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The anchor ref should be rewritten to a JSON pointer
	allOf := result["allOf"].([]any)
	ref := allOf[0].(map[string]any)["$ref"].(string)
	if ref != "#/definitions/B" {
		t.Errorf("expected $ref to be rewritten to #/definitions/B, got %s", ref)
	}
}
