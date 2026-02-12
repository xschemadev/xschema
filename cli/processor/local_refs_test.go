package processor

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestResolveInternalRefs_ScopedByID(t *testing.T) {
	// Schema where a sub-schema has $id and its own $defs,
	// and a $ref inside that scope points to #/$defs/inner.
	// The root does NOT have $defs, so resolving against root would fail.
	schema := `{
		"$id": "http://example.com/root.json",
		"properties": {
			"foo": {
				"$id": "http://example.com/nested.json",
				"$defs": {
					"inner": {
						"type": "string"
					}
				},
				"$ref": "#/$defs/inner"
			}
		}
	}`

	result, err := resolveInternalRefs(json.RawMessage(schema), "draft2020-12")
	if err != nil {
		t.Fatalf("resolveInternalRefs failed: %v", err)
	}

	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// The $ref should be resolved — the "foo" property should contain
	// the inlined content of $defs/inner (type: string), not a $ref
	props := parsed["properties"].(map[string]any)
	foo := props["foo"].(map[string]any)

	if _, hasRef := foo["$ref"]; hasRef {
		t.Errorf("expected $ref to be resolved, but it still exists")
	}

	// In draft2020-12, $ref allows siblings, so we expect an allOf merge
	// since the original object had both $ref and $defs/$id siblings.
	// The resolved target should have type: "string"
	if allOf, ok := foo["allOf"]; ok {
		items := allOf.([]any)
		found := false
		for _, item := range items {
			obj := item.(map[string]any)
			if obj["type"] == "string" {
				found = true
			}
		}
		if !found {
			t.Errorf("expected resolved ref to contain type:string, got %v", foo)
		}
	} else if foo["type"] != "string" {
		t.Errorf("expected resolved ref to have type:string, got %v", foo)
	}
}

func TestResolveInternalRefs_ScopedByLegacyID(t *testing.T) {
	// Same test but with legacy "id" instead of "$id"
	schema := `{
		"id": "http://example.com/root.json",
		"properties": {
			"foo": {
				"id": "http://example.com/nested.json",
				"$defs": {
					"inner": {
						"type": "number"
					}
				},
				"$ref": "#/$defs/inner"
			}
		}
	}`

	result, err := resolveInternalRefs(json.RawMessage(schema), "draft4")
	if err != nil {
		t.Fatalf("resolveInternalRefs failed: %v", err)
	}

	// In legacy drafts, $ref siblings are ignored, so we should get
	// just the resolved target (type: number)
	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	props := parsed["properties"].(map[string]any)
	foo := props["foo"].(map[string]any)

	if foo["type"] != "number" {
		t.Errorf("expected resolved ref to have type:number, got %v", foo)
	}
}

func TestResolveInternalRefs_FallbackToRoot(t *testing.T) {
	// A $ref at the root level should still resolve against root
	schema := `{
		"$defs": {
			"myType": {
				"type": "boolean"
			}
		},
		"$ref": "#/$defs/myType"
	}`

	result, err := resolveInternalRefs(json.RawMessage(schema), "draft2020-12")
	if err != nil {
		t.Fatalf("resolveInternalRefs failed: %v", err)
	}

	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// allOf merge since root had siblings ($defs) alongside $ref
	if allOf, ok := parsed["allOf"]; ok {
		items := allOf.([]any)
		found := false
		for _, item := range items {
			obj := item.(map[string]any)
			if obj["type"] == "boolean" {
				found = true
			}
		}
		if !found {
			t.Errorf("expected resolved ref to contain type:boolean, got %v", parsed)
		}
	} else if parsed["type"] != "boolean" {
		t.Errorf("expected resolved ref to have type:boolean, got %v", parsed)
	}
}

func TestResolveInternalRefs_NestedScopeDoesNotBreakRootRef(t *testing.T) {
	// A nested $id scope should not prevent root-level refs from resolving
	schema := `{
		"$defs": {
			"rootType": { "type": "integer" },
			"nested": {
				"$id": "http://example.com/nested",
				"$defs": {
					"innerType": { "type": "string" }
				}
			}
		},
		"properties": {
			"a": { "$ref": "#/$defs/rootType" }
		}
	}`

	result, err := resolveInternalRefs(json.RawMessage(schema), "draft2020-12")
	if err != nil {
		t.Fatalf("resolveInternalRefs failed: %v", err)
	}

	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	props := parsed["properties"].(map[string]any)
	a := props["a"].(map[string]any)

	if a["type"] != "integer" {
		t.Errorf("expected resolved ref to have type:integer, got %v", a)
	}
}

func TestResolveInternalRefs_RecursiveRootRef(t *testing.T) {
	// $ref: "#" is always kept intact — it always cycles back to root
	schema := `{
		"properties": {
			"foo": { "$ref": "#" }
		},
		"additionalProperties": false
	}`

	result, err := resolveInternalRefs(json.RawMessage(schema), "draft2020-12")
	if err != nil {
		t.Fatalf("resolveInternalRefs failed: %v", err)
	}

	var parsed map[string]any
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// foo should keep $ref: "#" directly (no inlining for root refs)
	props := parsed["properties"].(map[string]any)
	foo := props["foo"].(map[string]any)

	ref, ok := foo["$ref"].(string)
	if !ok || ref != "#" {
		t.Errorf("expected foo to have $ref: \"#\", got %v", foo)
	}
}

func TestResolveInternalRefs_RecursiveDefsRef(t *testing.T) {
	// mutual recursion via $defs — $ref to a def that refs back
	schema := `{
		"type": "object",
		"properties": {
			"value": { "type": "number" },
			"child": { "$ref": "#/$defs/node" }
		},
		"$defs": {
			"node": {
				"type": "object",
				"properties": {
					"value": { "type": "number" },
					"child": { "$ref": "#/$defs/node" }
				}
			}
		}
	}`

	result, err := resolveInternalRefs(json.RawMessage(schema), "draft2020-12")
	if err != nil {
		t.Fatalf("resolveInternalRefs failed: %v", err)
	}

	// The output should contain a $ref for the recursive cycle
	if !strings.Contains(string(result), `"$ref"`) {
		t.Errorf("expected output to contain $ref for recursive cycle, got %s", string(result))
	}
}

func TestResolveInternalRefs_ScopeNotFoundFallsToRoot(t *testing.T) {
	// If a ref can't be found in the nearest scope, fall back to root
	schema := `{
		"$defs": {
			"shared": { "type": "string" }
		},
		"properties": {
			"foo": {
				"$id": "http://example.com/nested.json",
				"$ref": "#/$defs/shared"
			}
		}
	}`

	result, err := resolveInternalRefs(json.RawMessage(schema), "draft2020-12")
	if err != nil {
		t.Fatalf("resolveInternalRefs failed: %v", err)
	}

	// Should resolve against root (fallback) since nested scope doesn't have $defs/shared
	if !strings.Contains(string(result), `"string"`) {
		t.Errorf("expected resolved ref to contain string type, got %s", string(result))
	}
}
