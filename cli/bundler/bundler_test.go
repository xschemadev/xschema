package bundler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/xschemadev/xschema/fetcher"
	"github.com/xschemadev/xschema/unsupported"
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

// testFileFetcher creates a fetcher that reads files relative to the testdata directory
// and also maps localhost:1234 URLs to the remotes/ subdirectory
func testFileFetcher() fetcher.Fetcher {
	return fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		// Handle localhost:1234 URLs (for JSON Schema Test Suite compatibility)
		if strings.HasPrefix(uri, "http://localhost:1234/") {
			localPath := filepath.Join(testdataPath("remotes"), strings.TrimPrefix(uri, "http://localhost:1234/"))
			data, err := os.ReadFile(localPath)
			if err != nil {
				return nil, err
			}
			return json.RawMessage(data), nil
		}
		// Otherwise read from filesystem
		data, err := os.ReadFile(uri)
		if err != nil {
			return nil, err
		}
		return json.RawMessage(data), nil
	})
}

func TestBundleSimpleSchema(t *testing.T) {
	ctx := context.Background()
	schema := readTestFile(t, "simple.json")

	bundled, err := Bundle(ctx, BundleInput{Schema: schema})
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

	bundled, err := Bundle(ctx, BundleInput{Schema: schema})
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

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    schema,
		SourceURI: testdataPath("with-file-ref.json"),
		Fetcher:   testFileFetcher(),
	})
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

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    schema,
		SourceURI: testdataPath("with-fragment-ref.json"),
		Fetcher:   testFileFetcher(),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The ref should point to the flattened definition
	props := result["properties"].(map[string]any)
	addressProp := props["address"].(map[string]any)
	ref := addressProp["$ref"].(string)

	// Definitions are flattened: #/$defs/key__Address (not #/$defs/key/definitions/Address)
	if !strings.Contains(ref, "__Address") {
		t.Errorf("expected ref to contain flattened __Address, got %s", ref)
	}

	// Verify the flattened definition exists in $defs
	defs := result["$defs"].(map[string]any)
	var foundFlattenedAddress bool
	for key := range defs {
		if strings.HasSuffix(key, "__Address") {
			foundFlattenedAddress = true
			break
		}
	}
	if !foundFlattenedAddress {
		t.Error("expected flattened Address definition in $defs")
	}
}

func TestBundleLocalhostMapping(t *testing.T) {
	ctx := context.Background()
	schema := readTestFile(t, "with-localhost-ref.json")

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    schema,
		SourceURI: testdataPath("with-localhost-ref.json"),
		Fetcher:   testFileFetcher(),
	})
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

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schemaWithInternalRefs),
		SourceURI: testdataPath("test.json"),
		Fetcher:   testFileFetcher(),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The internal refs in the embedded schema should be rewritten to flattened location
	defs := result["$defs"].(map[string]any)

	// With flattening, the Address definition should be at $defs/key__Address
	// and the ref in the user schema should point there
	var foundUserSchema bool
	for key, def := range defs {
		defObj := def.(map[string]any)
		if props, ok := defObj["properties"].(map[string]any); ok {
			if addr, ok := props["address"].(map[string]any); ok {
				foundUserSchema = true
				ref := addr["$ref"].(string)
				// The internal ref #/definitions/Address should be rewritten
				// to #/$defs/key__Address (flattened)
				if !strings.HasPrefix(ref, "#/$defs/") {
					t.Errorf("internal ref should be rewritten, got %s", ref)
				}
				if !strings.Contains(ref, "__Address") {
					t.Errorf("internal ref should point to flattened Address, got %s", ref)
				}
				// The flattened key should exist
				expectedDefKey := key + "__Address"
				if _, ok := defs[expectedDefKey]; !ok {
					// The ref might use a different key format, just check it exists
					refKey := strings.TrimPrefix(ref, "#/$defs/")
					if _, ok := defs[refKey]; !ok {
						t.Errorf("flattened Address def should exist, ref points to %s", refKey)
					}
				}
			}
		}
	}
	if !foundUserSchema {
		t.Error("embedded user schema not found in $defs")
	}
}

func TestBundleMissingFile(t *testing.T) {
	ctx := context.Background()
	schema := `{"$ref": "nonexistent.json"}`

	_, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: testdataPath("test.json"),
		Fetcher:   testFileFetcher(),
	})
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestBundleInvalidJSON(t *testing.T) {
	ctx := context.Background()
	schema := `{not valid json}`

	_, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestBundleContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	schema := `{"$ref": "http://example.com/schema.json"}`

	// Create a fetcher that checks for context cancellation
	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		return nil, ctx.Err()
	})

	_, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})
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
			got, err := fetcher.ResolveURI(tt.ref, tt.base)
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

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: testdataPath("test.json"),
		Fetcher:   testFileFetcher(),
	})
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

	_, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
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

	_, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
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

	_, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
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

	_, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
	if err == nil {
		t.Error("expected error for $recursiveRef")
	}
	var ukErr *unsupported.UnsupportedKeywordError
	if !errors.As(err, &ukErr) {
		t.Errorf("expected UnsupportedKeywordError, got: %T", err)
	} else if ukErr.Keyword != "$recursiveRef" {
		t.Errorf("expected keyword '$recursiveRef', got: %s", ukErr.Keyword)
	}
}

func TestBundleRejectsRecursiveAnchor(t *testing.T) {
	ctx := context.Background()
	schema := `{
		"type": "object",
		"$recursiveAnchor": true
	}`

	_, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
	if err == nil {
		t.Error("expected error for $recursiveAnchor")
	}
	var ukErr *unsupported.UnsupportedKeywordError
	if !errors.As(err, &ukErr) {
		t.Errorf("expected UnsupportedKeywordError, got: %T", err)
	} else if ukErr.Keyword != "$recursiveAnchor" {
		t.Errorf("expected keyword '$recursiveAnchor', got: %s", ukErr.Keyword)
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

	// No SourceURI and no Fetcher - external ref should error
	_, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
	if err == nil {
		t.Error("expected error for external ref without base URI")
	}
	if !strings.Contains(err.Error(), "requires a base URI") {
		t.Errorf("expected 'requires a base URI' in error, got: %v", err)
	}
}

func TestBundleAllowsAbsoluteExternalRefWithoutBaseURI(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	schema := `{
		"type": "object",
		"properties": {
			"user": { "$ref": "http://example.com/schema.json" }
		}
	}`

	// No SourceURI but ref is absolute - needs a fetcher though
	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		return nil, ctx.Err()
	})
	_, err := Bundle(ctx, BundleInput{
		Schema:  json.RawMessage(schema),
		Fetcher: fetcher,
	})
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

	// Note: Legacy drafts (3, 4, 6, 7) are normalized to 2020-12 syntax,
	// so they all get the 2020-12 $schema URI after bundling.
	tests := []struct {
		draft      string
		wantSchema string
	}{
		// Legacy drafts get normalized to 2020-12
		{"draft3", "https://json-schema.org/draft/2020-12/schema"},
		{"draft4", "https://json-schema.org/draft/2020-12/schema"},
		{"draft6", "https://json-schema.org/draft/2020-12/schema"},
		{"draft7", "https://json-schema.org/draft/2020-12/schema"},
		// Modern drafts are not normalized
		{"draft2019-09", "https://json-schema.org/draft/2019-09/schema"},
		{"draft2020-12", "https://json-schema.org/draft/2020-12/schema"},
	}

	for _, tt := range tests {
		t.Run(tt.draft, func(t *testing.T) {
			schema := `{"type": "string"}`

			bundled, err := Bundle(ctx, BundleInput{
				Schema: json.RawMessage(schema),
				Draft:  tt.draft,
			})
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

	// Schema already has $schema for 2020-12 - should be preserved (no normalization needed)
	schema := `{"$schema": "https://json-schema.org/draft/2020-12/schema", "type": "string"}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema: json.RawMessage(schema),
		Draft:  "draft4", // Draft hint ignored when $schema is present
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	got := result["$schema"].(string)
	// Should preserve the 2020-12 schema (no normalization needed)
	if got != "https://json-schema.org/draft/2020-12/schema" {
		t.Errorf("existing $schema should be preserved, got %q", got)
	}
}

func TestBundleLegacySchemaGetsNormalized(t *testing.T) {
	ctx := context.Background()

	// Schema has legacy draft7 $schema - gets normalized to 2020-12
	schema := `{"$schema": "http://json-schema.org/draft-07/schema#", "type": "string"}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema: json.RawMessage(schema),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	got := result["$schema"].(string)
	// Legacy draft gets normalized to 2020-12
	if got != "https://json-schema.org/draft/2020-12/schema" {
		t.Errorf("legacy $schema should be normalized to 2020-12, got %q", got)
	}
}

func TestBundleNoDraftNoInjection(t *testing.T) {
	ctx := context.Background()

	// No draft specified - should not inject $schema
	schema := `{"type": "string"}`

	bundled, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
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
			// Keys with literal JSON pointer escape sequences (edge cases)
			"literal~1key": map[string]any{"type": "string"},  // key is literally "~1"
			"literal~0key": map[string]any{"type": "number"},  // key is literally "~0"
			"uri%2Fslash":  map[string]any{"type": "boolean"}, // key contains literal "%2F"
			"combo~/field": map[string]any{"type": "integer"}, // key has both ~ and /
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

		// RFC 6901 order test cases: URI decode first, then JSON pointer unescape
		// Key "literal~1key" - the ~1 is literal, so we need ~01 (escape the ~)
		{"literal tilde-one in key", "#/$defs/literal~01key", false},
		// Key "literal~0key" - the ~0 is literal, so we need ~00 (escape the ~)
		{"literal tilde-zero in key", "#/$defs/literal~00key", false},
		// Key "uri%2Fslash" - contains literal %2F, URI encode the % as %25
		{"literal percent-2F in key", "#/$defs/uri%252Fslash", false},
		// Key "combo~/field" - has ~ (needs ~0) and / (needs ~1)
		{"combo tilde and slash", "#/$defs/combo~0~1field", false},
		// Wrong: using %2F for a key that literally contains / would fail
		{"wrong slash encoding for literal percent2F", "#/$defs/uri%2Fslash", true},
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

	bundled, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
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

	bundled, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
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

	_, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
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

	bundled, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
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

	bundled, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
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

	bundled, err := Bundle(ctx, BundleInput{Schema: json.RawMessage(schema)})
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

func TestBundleRemoteSchemaWithAnchor(t *testing.T) {
	ctx := context.Background()

	// Schema that refs a remote schema, and the remote schema uses anchors internally
	schema := `{
		"type": "object",
		"properties": {
			"remote": { "$ref": "schema-with-anchor.json" }
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: testdataPath("test.json"),
		Fetcher:   testFileFetcher(),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The remote schema is embedded in $defs. Its internal anchor ref (#myAnchor)
	// should be rewritten to point to the correct flattened location.
	defs := result["$defs"].(map[string]any)

	// Find the embedded schema (should have properties.value.$ref)
	var embeddedRef string
	for _, def := range defs {
		defObj, ok := def.(map[string]any)
		if !ok {
			continue
		}
		props, ok := defObj["properties"].(map[string]any)
		if !ok {
			continue
		}
		valueProp, ok := props["value"].(map[string]any)
		if !ok {
			continue
		}
		if ref, ok := valueProp["$ref"].(string); ok {
			embeddedRef = ref
			break
		}
	}

	if embeddedRef == "" {
		t.Fatal("could not find embedded schema with properties.value.$ref")
	}

	// The anchor should be resolved to point to flattened location
	// #myAnchor → #/$defs/schema_with_anchor_json__Target (flattened)
	if !strings.HasPrefix(embeddedRef, "#/$defs/") {
		t.Errorf("expected anchor ref to be rewritten to #/$defs/..., got %s", embeddedRef)
	}
	if !strings.Contains(embeddedRef, "__Target") {
		t.Errorf("expected anchor ref to point to flattened Target def, got %s", embeddedRef)
	}
}

func TestBundleExternalAnchorRef(t *testing.T) {
	ctx := context.Background()

	// Schema that refs an anchor in a remote schema directly: remote.json#anchor
	schema := `{
		"type": "object",
		"properties": {
			"value": { "$ref": "schema-with-anchor.json#myAnchor" }
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: testdataPath("test.json"),
		Fetcher:   testFileFetcher(),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The ref should point directly to the flattened Target def
	props := result["properties"].(map[string]any)
	valueProp := props["value"].(map[string]any)
	ref := valueProp["$ref"].(string)

	// Should be #/$defs/key__Target (flattened location)
	if !strings.HasPrefix(ref, "#/$defs/") {
		t.Errorf("expected ref to be #/$defs/..., got %s", ref)
	}
	if !strings.Contains(ref, "__Target") {
		t.Errorf("expected ref to point to flattened Target def, got %s", ref)
	}

	// The flattened Target def should exist
	defs := result["$defs"].(map[string]any)
	refKey := strings.TrimPrefix(ref, "#/$defs/")
	if _, ok := defs[refKey]; !ok {
		t.Errorf("flattened Target def %s not found in $defs", refKey)
	}
}

// =============================================================================
// Fetcher Interface Tests (US-004)
// =============================================================================

func TestFetcherCalledForExternalRefs(t *testing.T) {
	ctx := context.Background()

	var fetchedURIs []string
	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		fetchedURIs = append(fetchedURIs, uri)
		return json.RawMessage(`{"type": "string"}`), nil
	})

	schema := `{
		"type": "object",
		"properties": {
			"external": { "$ref": "http://example.com/schema.json" }
		}
	}`

	_, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(fetchedURIs) != 1 {
		t.Errorf("expected Fetcher.Fetch() called once, got %d calls", len(fetchedURIs))
	}
	if len(fetchedURIs) > 0 && fetchedURIs[0] != "http://example.com/schema.json" {
		t.Errorf("expected fetch URI http://example.com/schema.json, got %s", fetchedURIs[0])
	}
}

func TestFetcherCalledForRelativeRefs(t *testing.T) {
	ctx := context.Background()

	var fetchedURIs []string
	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		fetchedURIs = append(fetchedURIs, uri)
		return json.RawMessage(`{"type": "integer"}`), nil
	})

	schema := `{
		"type": "object",
		"properties": {
			"value": { "$ref": "other.json" }
		}
	}`

	_, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/schemas/base.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(fetchedURIs) != 1 {
		t.Errorf("expected Fetcher.Fetch() called once, got %d calls", len(fetchedURIs))
	}
	// Relative ref should be resolved against base URI
	expected := "http://example.com/schemas/other.json"
	if len(fetchedURIs) > 0 && fetchedURIs[0] != expected {
		t.Errorf("expected fetch URI %s, got %s", expected, fetchedURIs[0])
	}
}

func TestFetcherErrorPropagated(t *testing.T) {
	ctx := context.Background()

	fetchErr := fmt.Errorf("network error: connection refused")
	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		return nil, fetchErr
	})

	schema := `{
		"type": "object",
		"properties": {
			"external": { "$ref": "http://example.com/schema.json" }
		}
	}`

	_, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})

	if err == nil {
		t.Fatal("expected error from Fetcher to be propagated")
	}
	if !strings.Contains(err.Error(), "network error") {
		t.Errorf("expected error to contain 'network error', got: %v", err)
	}
}

func TestFetcherInvalidJSONError(t *testing.T) {
	ctx := context.Background()

	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		return json.RawMessage(`{not valid json`), nil
	})

	schema := `{
		"type": "object",
		"properties": {
			"external": { "$ref": "http://example.com/schema.json" }
		}
	}`

	_, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})

	if err == nil {
		t.Fatal("expected error for invalid JSON from Fetcher")
	}
	if !strings.Contains(err.Error(), "parse") {
		t.Errorf("expected error to mention parsing, got: %v", err)
	}
}

func TestNilFetcherExternalRefError(t *testing.T) {
	ctx := context.Background()

	schema := `{
		"type": "object",
		"properties": {
			"external": { "$ref": "http://example.com/schema.json" }
		}
	}`

	_, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   nil, // explicitly nil
	})

	if err == nil {
		t.Fatal("expected error for external ref with nil Fetcher")
	}
	if !strings.Contains(err.Error(), "no fetcher") {
		t.Errorf("expected error to mention 'no fetcher', got: %v", err)
	}
}

func TestLocalRefsDoNotCallFetcher(t *testing.T) {
	ctx := context.Background()

	fetcherCalled := false
	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		fetcherCalled = true
		return nil, fmt.Errorf("fetcher should not be called")
	})

	// Schema with only local refs
	schema := `{
		"type": "object",
		"$defs": {
			"User": { "type": "string" }
		},
		"properties": {
			"name": { "$ref": "#/$defs/User" }
		}
	}`

	_, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fetcherCalled {
		t.Error("Fetcher should not be called for local refs (#/...)")
	}
}

func TestLocalAnchorRefsDoNotCallFetcher(t *testing.T) {
	ctx := context.Background()

	fetcherCalled := false
	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		fetcherCalled = true
		return nil, fmt.Errorf("fetcher should not be called")
	})

	// Schema with local anchor refs
	schema := `{
		"$ref": "#foo",
		"$defs": {
			"A": {
				"$anchor": "foo",
				"type": "integer"
			}
		}
	}`

	_, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fetcherCalled {
		t.Error("Fetcher should not be called for local anchor refs (#anchor)")
	}
}

func TestFetcherCalledOncePerURI(t *testing.T) {
	ctx := context.Background()

	fetchCounts := make(map[string]int)
	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		fetchCounts[uri]++
		return json.RawMessage(`{"type": "string"}`), nil
	})

	// Schema that refs the same external schema multiple times
	schema := `{
		"type": "object",
		"properties": {
			"a": { "$ref": "http://example.com/schema.json" },
			"b": { "$ref": "http://example.com/schema.json" },
			"c": { "$ref": "http://example.com/schema.json" }
		}
	}`

	_, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should only fetch once due to caching
	uri := "http://example.com/schema.json"
	if fetchCounts[uri] != 1 {
		t.Errorf("expected Fetcher.Fetch() called once for %s, got %d calls", uri, fetchCounts[uri])
	}
}

func TestFetchFuncAdapter(t *testing.T) {
	// Test that FetchFunc correctly implements Fetcher interface
	fn := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		if uri == "test://ok" {
			return json.RawMessage(`{"type": "number"}`), nil
		}
		return nil, fmt.Errorf("unknown uri: %s", uri)
	})

	// Test successful fetch
	result, err := fn.Fetch(context.Background(), "test://ok")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(result) != `{"type": "number"}` {
		t.Errorf("expected `{\"type\": \"number\"}`, got %s", string(result))
	}

	// Test error case
	_, err = fn.Fetch(context.Background(), "test://fail")
	if err == nil {
		t.Error("expected error for unknown uri")
	}
	if !strings.Contains(err.Error(), "unknown uri") {
		t.Errorf("expected 'unknown uri' in error, got: %v", err)
	}
}

// =============================================================================
// Nested $defs Flattening Tests (US-005, US-006)
// =============================================================================

func TestBundle3LevelNestedDefs(t *testing.T) {
	ctx := context.Background()

	// Remote schema with 3-level nesting: $defs/A/$defs/B/$defs/C
	remoteSchema := `{
		"type": "object",
		"$defs": {
			"A": {
				"type": "object",
				"$defs": {
					"B": {
						"type": "object",
						"$defs": {
							"C": {
								"type": "string"
							}
						},
						"properties": {
							"c": { "$ref": "#/$defs/A/$defs/B/$defs/C" }
						}
					}
				},
				"properties": {
					"b": { "$ref": "#/$defs/A/$defs/B" }
				}
			}
		},
		"properties": {
			"a": { "$ref": "#/$defs/A" }
		}
	}`

	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		return json.RawMessage(remoteSchema), nil
	})

	schema := `{
		"type": "object",
		"properties": {
			"remote": { "$ref": "http://example.com/nested.json" }
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	defs := result["$defs"].(map[string]any)

	// Should have flattened keys with __ separators
	// Expected keys: example_com_nested_json, example_com_nested_json__A,
	//               example_com_nested_json__A__B, example_com_nested_json__A__B__C
	var foundKeys []string
	for key := range defs {
		foundKeys = append(foundKeys, key)
	}

	// Check all nested defs are flattened to root level
	var hasA, hasB, hasC bool
	for key := range defs {
		if strings.Contains(key, "__A__B__C") {
			hasC = true
		} else if strings.Contains(key, "__A__B") && !strings.Contains(key, "__C") {
			hasB = true
		} else if strings.Contains(key, "__A") && !strings.Contains(key, "__B") {
			hasA = true
		}
	}

	if !hasA {
		t.Errorf("expected flattened key containing __A, found: %v", foundKeys)
	}
	if !hasB {
		t.Errorf("expected flattened key containing __A__B, found: %v", foundKeys)
	}
	if !hasC {
		t.Errorf("expected flattened key containing __A__B__C, found: %v", foundKeys)
	}
}

func TestBundleMixedDefsAndDefinitions(t *testing.T) {
	ctx := context.Background()

	// Remote schema with both $defs and definitions (draft-04 style)
	remoteSchema := `{
		"type": "object",
		"$defs": {
			"Modern": {
				"type": "string"
			}
		},
		"definitions": {
			"Legacy": {
				"type": "integer"
			}
		},
		"properties": {
			"modern": { "$ref": "#/$defs/Modern" },
			"legacy": { "$ref": "#/definitions/Legacy" }
		}
	}`

	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		return json.RawMessage(remoteSchema), nil
	})

	schema := `{
		"type": "object",
		"properties": {
			"remote": { "$ref": "http://example.com/mixed.json" }
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	defs := result["$defs"].(map[string]any)

	// Both Modern and Legacy should be flattened to root $defs
	var foundModern, foundLegacy bool
	for key := range defs {
		if strings.HasSuffix(key, "__Modern") {
			foundModern = true
		}
		if strings.HasSuffix(key, "__Legacy") {
			foundLegacy = true
		}
	}

	if !foundModern {
		t.Error("expected flattened Modern definition in $defs")
	}
	if !foundLegacy {
		t.Error("expected flattened Legacy definition in $defs")
	}
}

func TestBundleKeyCollisionWithCounter(t *testing.T) {
	ctx := context.Background()

	// Two remote schemas with the same def name "Shared"
	schemaA := `{
		"type": "object",
		"$defs": {
			"Shared": {
				"type": "string",
				"description": "from A"
			}
		}
	}`

	schemaB := `{
		"type": "object",
		"$defs": {
			"Shared": {
				"type": "integer",
				"description": "from B"
			}
		}
	}`

	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		if strings.Contains(uri, "a.json") {
			return json.RawMessage(schemaA), nil
		}
		return json.RawMessage(schemaB), nil
	})

	schema := `{
		"type": "object",
		"properties": {
			"a": { "$ref": "http://example.com/a.json" },
			"b": { "$ref": "http://example.com/b.json" }
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	defs := result["$defs"].(map[string]any)

	// Both Shared defs should exist with unique keys (one may have counter suffix)
	var sharedCount int
	for key := range defs {
		if strings.Contains(key, "Shared") {
			sharedCount++
		}
	}

	if sharedCount < 2 {
		t.Errorf("expected 2 Shared definitions (with unique keys), got %d", sharedCount)
	}

	// Verify they have different types (string vs integer)
	var foundString, foundInteger bool
	for key, def := range defs {
		if strings.Contains(key, "Shared") {
			defObj := def.(map[string]any)
			if defObj["type"] == "string" {
				foundString = true
			}
			if defObj["type"] == "integer" {
				foundInteger = true
			}
		}
	}

	if !foundString || !foundInteger {
		t.Error("expected both string and integer Shared definitions to be preserved")
	}
}

func TestBundleRefsInFlattenedDefsRewritten(t *testing.T) {
	ctx := context.Background()

	// Remote schema where nested def references another nested def
	remoteSchema := `{
		"type": "object",
		"$defs": {
			"Outer": {
				"$defs": {
					"Inner": {
						"type": "string"
					}
				},
				"properties": {
					"value": { "$ref": "#/$defs/Outer/$defs/Inner" }
				}
			}
		}
	}`

	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		return json.RawMessage(remoteSchema), nil
	})

	schema := `{
		"type": "object",
		"properties": {
			"remote": { "$ref": "http://example.com/refs.json" }
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	defs := result["$defs"].(map[string]any)

	// Find the flattened Outer def and check its internal ref was rewritten
	var outerKey string
	for key := range defs {
		if strings.HasSuffix(key, "__Outer") {
			outerKey = key
			break
		}
	}

	if outerKey == "" {
		t.Fatal("expected flattened Outer definition")
	}

	outerDef := defs[outerKey].(map[string]any)
	props := outerDef["properties"].(map[string]any)
	valueRef := props["value"].(map[string]any)["$ref"].(string)

	// The ref should point to the flattened Inner location
	expectedRef := "#/$defs/" + outerKey + "__Inner"
	if valueRef != expectedRef {
		t.Errorf("expected ref to be rewritten to %s, got %s", expectedRef, valueRef)
	}

	// Verify the flattened Inner def exists
	innerKey := outerKey + "__Inner"
	if _, ok := defs[innerKey]; !ok {
		t.Errorf("expected flattened Inner definition at key %s", innerKey)
	}
}

func TestBundleDeeplyNestedRefChain(t *testing.T) {
	ctx := context.Background()

	// Schema A references B/$defs/X, which references C/$defs/Y
	schemaC := `{
		"type": "object",
		"$defs": {
			"Y": {
				"type": "string",
				"minLength": 1
			}
		}
	}`

	schemaB := `{
		"type": "object",
		"$defs": {
			"X": {
				"$ref": "http://example.com/c.json#/$defs/Y"
			}
		}
	}`

	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		if strings.Contains(uri, "b.json") {
			return json.RawMessage(schemaB), nil
		}
		if strings.Contains(uri, "c.json") {
			return json.RawMessage(schemaC), nil
		}
		return nil, fmt.Errorf("unknown URI: %s", uri)
	})

	schema := `{
		"type": "object",
		"properties": {
			"value": { "$ref": "http://example.com/b.json#/$defs/X" }
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/a.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	defs := result["$defs"].(map[string]any)

	// Should have both B's X and C's Y flattened
	var foundX, foundY bool
	for key := range defs {
		if strings.Contains(key, "__X") || strings.HasSuffix(key, "_X") {
			foundX = true
		}
		if strings.Contains(key, "__Y") || strings.HasSuffix(key, "_Y") {
			foundY = true
		}
	}

	if !foundX {
		t.Error("expected flattened X definition from b.json")
	}
	if !foundY {
		t.Error("expected flattened Y definition from c.json")
	}

	// Verify the ref chain resolves correctly
	props := result["properties"].(map[string]any)
	valueRef := props["value"].(map[string]any)["$ref"].(string)
	if !strings.HasPrefix(valueRef, "#/$defs/") {
		t.Errorf("expected ref to be rewritten to #/$defs/..., got %s", valueRef)
	}
}

// TestBundleSiblingRefAfterID verifies that $id in one property doesn't affect
// refs in sibling properties. Per JSON Schema spec, $id creates a new scope only
// for its descendants, not for siblings at the same level.
func TestBundleSiblingRefAfterID(t *testing.T) {
	ctx := context.Background()

	// Schema where "nested" has $id but "sibling" should still resolve refs
	// relative to the root, not relative to nested's scope
	schema := `{
		"type": "object",
		"$defs": {
			"Target": { "type": "string" }
		},
		"properties": {
			"nested": {
				"$id": "http://example.com/nested",
				"type": "object"
			},
			"sibling": {
				"$ref": "#/$defs/Target"
			}
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/root.json",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The sibling's $ref should remain unchanged - it should still point to
	// #/$defs/Target (root scope), not be affected by nested's $id
	props := result["properties"].(map[string]any)
	sibling := props["sibling"].(map[string]any)
	siblingRef := sibling["$ref"].(string)

	if siblingRef != "#/$defs/Target" {
		t.Errorf("sibling ref should remain #/$defs/Target, got %s", siblingRef)
	}
}

// TestBundle5LevelNestedDefs verifies that schemas with 5+ levels of nested $defs
// are correctly flattened with all refs rewritten properly.
func TestBundle5LevelNestedDefs(t *testing.T) {
	ctx := context.Background()

	// Remote schema with 5 levels of nesting: $defs/L1/$defs/L2/$defs/L3/$defs/L4/$defs/L5
	remoteSchema := `{
		"type": "object",
		"$defs": {
			"L1": {
				"type": "object",
				"$defs": {
					"L2": {
						"type": "object",
						"$defs": {
							"L3": {
								"type": "object",
								"$defs": {
									"L4": {
										"type": "object",
										"$defs": {
											"L5": {
												"type": "string",
												"description": "deepest level"
											}
										},
										"properties": {
											"l5ref": { "$ref": "#/$defs/L1/$defs/L2/$defs/L3/$defs/L4/$defs/L5" }
										}
									}
								},
								"properties": {
									"l4ref": { "$ref": "#/$defs/L1/$defs/L2/$defs/L3/$defs/L4" }
								}
							}
						},
						"properties": {
							"l3ref": { "$ref": "#/$defs/L1/$defs/L2/$defs/L3" }
						}
					}
				},
				"properties": {
					"l2ref": { "$ref": "#/$defs/L1/$defs/L2" }
				}
			}
		},
		"properties": {
			"l1ref": { "$ref": "#/$defs/L1" }
		}
	}`

	fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
		return json.RawMessage(remoteSchema), nil
	})

	schema := `{
		"type": "object",
		"properties": {
			"remote": { "$ref": "http://example.com/deep.json" }
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/base.json",
		Fetcher:   fetcher,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	defs := result["$defs"].(map[string]any)

	// Should have flattened keys with __ separators for all 5 levels
	var foundL1, foundL2, foundL3, foundL4, foundL5 bool
	for key := range defs {
		switch {
		case strings.Contains(key, "__L1__L2__L3__L4__L5"):
			foundL5 = true
		case strings.Contains(key, "__L1__L2__L3__L4") && !strings.Contains(key, "__L5"):
			foundL4 = true
		case strings.Contains(key, "__L1__L2__L3") && !strings.Contains(key, "__L4"):
			foundL3 = true
		case strings.Contains(key, "__L1__L2") && !strings.Contains(key, "__L3"):
			foundL2 = true
		case strings.Contains(key, "__L1") && !strings.Contains(key, "__L2"):
			foundL1 = true
		}
	}

	if !foundL1 {
		t.Error("expected flattened L1 definition")
	}
	if !foundL2 {
		t.Error("expected flattened L2 definition")
	}
	if !foundL3 {
		t.Error("expected flattened L3 definition")
	}
	if !foundL4 {
		t.Error("expected flattened L4 definition")
	}
	if !foundL5 {
		t.Error("expected flattened L5 definition")
	}

	// Verify refs were rewritten correctly - find L4 and check its l5ref
	for key, def := range defs {
		if strings.Contains(key, "__L1__L2__L3__L4") && !strings.Contains(key, "__L5") {
			defObj := def.(map[string]any)
			props, ok := defObj["properties"].(map[string]any)
			if !ok {
				continue
			}
			l5ref, ok := props["l5ref"].(map[string]any)
			if !ok {
				continue
			}
			ref := l5ref["$ref"].(string)
			// Should point to flattened L5, not the nested path
			if !strings.HasPrefix(ref, "#/$defs/") || !strings.Contains(ref, "__L5") {
				t.Errorf("L4's l5ref should be rewritten to flattened L5, got %s", ref)
			}
		}
	}
}

// TestBundleSiblingRefToParentPath tests that refs in sibling properties can
// reference paths that would be invalid if resolved against a sibling's $id scope.
func TestBundleSiblingRefToParentPath(t *testing.T) {
	ctx := context.Background()

	// Schema with nested $id that establishes new scope, but sibling refs
	// should resolve against the parent (root) scope, not the $id's scope
	schema := `{
		"type": "object",
		"$defs": {
			"Local": { "type": "string" }
		},
		"properties": {
			"nested": {
				"$id": "http://example.com/nested.json",
				"type": "object",
				"properties": {
					"innerProp": { "type": "number" }
				}
			},
			"sibling": {
				"$ref": "#/$defs/Local"
			}
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: "http://example.com/root.json",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The sibling's $ref should remain #/$defs/Local, unchanged
	props := result["properties"].(map[string]any)
	sibling := props["sibling"].(map[string]any)
	siblingRef := sibling["$ref"].(string)

	if siblingRef != "#/$defs/Local" {
		t.Errorf("sibling ref should remain #/$defs/Local, got %s", siblingRef)
	}

	// The root $defs should still contain Local
	defs := result["$defs"].(map[string]any)
	if _, ok := defs["Local"]; !ok {
		t.Error("root $defs should still contain Local definition")
	}
}

// =============================================================================
// Benchmark Tests for Linear Scaling (US-006)
// =============================================================================

// generateNestedDefsSchema generates a schema with N levels of nested $defs
// where each level has refs to the next level down.
func generateNestedDefsSchema(levels int) string {
	if levels <= 0 {
		return `{"type": "string"}`
	}

	// Build from innermost to outermost
	inner := `{"type": "string", "description": "leaf"}`

	for i := levels; i >= 1; i-- {
		defName := fmt.Sprintf("L%d", i)
		var refPath string
		if i == levels {
			// Innermost level has no ref
			refPath = ""
		} else {
			// Build the ref path like #/$defs/L1/$defs/L2/.../$defs/L(i+1)
			parts := []string{"#"}
			for j := 1; j <= i+1; j++ {
				parts = append(parts, fmt.Sprintf("$defs/L%d", j))
			}
			refPath = strings.Join(parts, "/")
		}

		var propsSection string
		if refPath != "" {
			propsSection = fmt.Sprintf(`,"properties":{"ref":{"$ref":"%s"}}`, refPath)
		}

		inner = fmt.Sprintf(`{
			"type": "object",
			"$defs": {
				"%s": %s
			}%s
		}`, defName, inner, propsSection)
	}

	return inner
}

// BenchmarkFlattenDefs measures flattening performance with increasing nesting depth.
// We verify linear scaling by checking that time grows linearly with depth.
func BenchmarkFlattenDefs(b *testing.B) {
	depths := []int{2, 4, 6, 8, 10}

	for _, depth := range depths {
		b.Run(fmt.Sprintf("depth_%d", depth), func(b *testing.B) {
			schema := generateNestedDefsSchema(depth)
			fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
				return json.RawMessage(schema), nil
			})

			rootSchema := `{
				"type": "object",
				"properties": {
					"remote": { "$ref": "http://example.com/nested.json" }
				}
			}`

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := Bundle(context.Background(), BundleInput{
					Schema:    json.RawMessage(rootSchema),
					SourceURI: "http://example.com/base.json",
					Fetcher:   fetcher,
				})
				if err != nil {
					b.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

// TestFlattenDefsLinearScaling verifies that flattening time scales linearly
// with nesting depth (not quadratically).
func TestFlattenDefsLinearScaling(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping linear scaling test in short mode")
	}

	ctx := context.Background()

	// Test with depths 5 and 10 - if quadratic, 10 would be ~4x slower than 5
	// With linear scaling, 10 should be ~2x slower than 5
	depths := []int{5, 10}
	iterations := 100
	times := make(map[int]int64)

	for _, depth := range depths {
		schema := generateNestedDefsSchema(depth)
		fetcher := fetcher.FetchFunc(func(_ context.Context, uri string) (json.RawMessage, error) {
			return json.RawMessage(schema), nil
		})

		rootSchema := `{
			"type": "object",
			"properties": {
				"remote": { "$ref": "http://example.com/nested.json" }
			}
		}`

		// Warm up
		for i := 0; i < 5; i++ {
			_, _ = Bundle(ctx, BundleInput{
				Schema:    json.RawMessage(rootSchema),
				SourceURI: "http://example.com/base.json",
				Fetcher:   fetcher,
			})
		}

		// Time the iterations
		start := time.Now()
		for i := 0; i < iterations; i++ {
			_, err := Bundle(ctx, BundleInput{
				Schema:    json.RawMessage(rootSchema),
				SourceURI: "http://example.com/base.json",
				Fetcher:   fetcher,
			})
			if err != nil {
				t.Fatalf("depth %d: unexpected error: %v", depth, err)
			}
		}
		times[depth] = time.Since(start).Nanoseconds()
	}

	// With linear scaling: time(10) / time(5) should be approximately 2
	// With quadratic scaling: time(10) / time(5) would be approximately 4
	// Allow some margin for variance
	ratio := float64(times[10]) / float64(times[5])
	t.Logf("depth 5: %dns, depth 10: %dns, ratio: %.2f", times[5], times[10], ratio)

	// If ratio > 3, it's likely quadratic
	if ratio > 3.0 {
		t.Errorf("scaling appears quadratic: ratio %.2f (expected ~2 for linear)", ratio)
	}
}

// =============================================================================
// Anchor Rewriting After Flattening (US-008)
// =============================================================================

func TestBundleDeeplyNestedAnchorRef(t *testing.T) {
	ctx := context.Background()

	// Schema that refs an anchor in a deeply nested $def (3+ levels)
	// The external schema has: $defs/Level1/$defs/Level2/$defs/Target with $anchor: "deepAnchor"
	schema := `{
		"type": "object",
		"properties": {
			"value": { "$ref": "schema-deep-anchor.json#deepAnchor" }
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: testdataPath("test.json"),
		Fetcher:   testFileFetcher(),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The ref should point to the flattened location of the deeply nested Target
	props := result["properties"].(map[string]any)
	valueProp := props["value"].(map[string]any)
	ref := valueProp["$ref"].(string)

	// Should be #/$defs/key__Level1__Level2__Target (fully flattened)
	if !strings.HasPrefix(ref, "#/$defs/") {
		t.Errorf("expected ref to be #/$defs/..., got %s", ref)
	}
	if !strings.Contains(ref, "Target") {
		t.Errorf("expected ref to contain 'Target', got %s", ref)
	}

	// The flattened def must exist
	defs := result["$defs"].(map[string]any)
	refKey := strings.TrimPrefix(ref, "#/$defs/")
	if _, ok := defs[refKey]; !ok {
		t.Errorf("flattened def %s not found in $defs. Available keys: %v", refKey, getMapKeys(defs))
	}
}

func TestBundleExternalSchemaWithDeeplyNestedInternalAnchorRef(t *testing.T) {
	ctx := context.Background()

	// Schema that refs a remote schema. The remote schema has a deeply nested anchor
	// and an internal $ref to that anchor.
	schema := `{
		"type": "object",
		"properties": {
			"remote": { "$ref": "schema-deep-anchor.json" }
		}
	}`

	bundled, err := Bundle(ctx, BundleInput{
		Schema:    json.RawMessage(schema),
		SourceURI: testdataPath("test.json"),
		Fetcher:   testFileFetcher(),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(bundled, &result); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	defs := result["$defs"].(map[string]any)

	// Find the embedded schema and its internal anchor ref
	var embeddedRef string
	for _, def := range defs {
		defObj, ok := def.(map[string]any)
		if !ok {
			continue
		}
		props, ok := defObj["properties"].(map[string]any)
		if !ok {
			continue
		}
		valueProp, ok := props["value"].(map[string]any)
		if !ok {
			continue
		}
		if ref, ok := valueProp["$ref"].(string); ok {
			embeddedRef = ref
			break
		}
	}

	if embeddedRef == "" {
		t.Fatal("could not find embedded schema with properties.value.$ref")
	}

	// The internal anchor ref should be rewritten to point to the flattened location
	if !strings.HasPrefix(embeddedRef, "#/$defs/") {
		t.Errorf("expected anchor ref to be rewritten to #/$defs/..., got %s", embeddedRef)
	}

	// The flattened def must exist
	refKey := strings.TrimPrefix(embeddedRef, "#/$defs/")
	if _, ok := defs[refKey]; !ok {
		t.Errorf("flattened def %s not found in $defs. Available keys: %v", refKey, getMapKeys(defs))
	}
}

func getMapKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
