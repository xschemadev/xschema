package processor

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

var nonSchemaKeywords = map[string]bool{
	"const":    true,
	"default":  true,
	"enum":     true,
	"example":  true,
	"examples": true,
}

// propertyMapKeywords are keywords whose values are objects where
// keys are arbitrary names (property names, definition names, etc.),
// not JSON Schema keywords. The values under these keys ARE subschemas,
// but the containing object itself is not a schema — so we must not
// interpret its keys (like "$ref") as JSON Schema keywords.
var propertyMapKeywords = map[string]bool{
	"properties":        true,
	"patternProperties": true,
	"$defs":             true,
	"definitions":       true,
	"dependentSchemas":  true,
	"dependencies":      true,
}

var legacyRefSiblingDrafts = map[string]bool{
	"draft3": true,
	"draft4": true,
	"draft6": true,
	"draft7": true,
}

type localRefResolver struct {
	root              any
	resolving         map[string]bool
	ignoreRefSiblings bool
	scopes            []any // stack of sub-schemas with $id (nearest first, root last)
}

func resolveInternalRefs(schema json.RawMessage, draft string) (json.RawMessage, error) {
	var root any
	if err := json.Unmarshal(schema, &root); err != nil {
		return nil, fmt.Errorf("failed to parse bundled schema: %w", err)
	}

	resolver := &localRefResolver{
		root:              root,
		resolving:         make(map[string]bool),
		ignoreRefSiblings: shouldIgnoreRefSiblings(draft, root),
		scopes:            []any{root},
	}

	resolved, err := resolver.resolveNode(root, false, false)
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(resolved)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal schema with resolved refs: %w", err)
	}

	return result, nil
}

func (r *localRefResolver) resolveNode(node any, inNonSchemaContext bool, inPropertyMap bool) (any, error) {
	if inNonSchemaContext {
		return cloneNode(node), nil
	}

	switch v := node.(type) {
	case map[string]any:
		// track $id scope changes — sub-schemas with $id create new resolution scopes
		popScope := false
		if !inPropertyMap {
			if _, hasID := v["$id"].(string); hasID {
				r.scopes = append([]any{v}, r.scopes...)
				popScope = true
			} else if _, hasLegacyID := v["id"].(string); hasLegacyID {
				r.scopes = append([]any{v}, r.scopes...)
				popScope = true
			}
		}
		if popScope {
			defer func() { r.scopes = r.scopes[1:] }()
		}

		// only check for $ref keyword when this object is a schema,
		// not when it's a property map (e.g. the value of "properties")
		// where keys are arbitrary names like "$ref"
		if !inPropertyMap {
			if rawRef, ok := v["$ref"]; ok {
				ref, isString := rawRef.(string)
				if !isString {
					return nil, fmt.Errorf("$ref value must be a string")
				}
				return r.resolveRefObject(v, ref)
			}
		}

		result := make(map[string]any, len(v))
		for key, val := range v {
			childIsNonSchema := nonSchemaKeywords[key]
			childIsPropertyMap := propertyMapKeywords[key]
			resolvedChild, err := r.resolveNode(val, childIsNonSchema, childIsPropertyMap)
			if err != nil {
				return nil, err
			}
			result[key] = resolvedChild
		}
		return result, nil

	case []any:
		result := make([]any, len(v))
		for i, item := range v {
			resolvedChild, err := r.resolveNode(item, false, false)
			if err != nil {
				return nil, err
			}
			result[i] = resolvedChild
		}
		return result, nil

	default:
		return node, nil
	}
}

func (r *localRefResolver) resolveRefObject(obj map[string]any, ref string) (any, error) {
	if !strings.HasPrefix(ref, "#") {
		return nil, fmt.Errorf("encountered unresolved non-local $ref %q", ref)
	}

	if r.resolving[ref] {
		// recursive ref — keep $ref intact for the adapter to handle via lazy/suspend
		return obj, nil
	}

	// $ref: "#" always creates a cycle (root references itself) — keep intact
	// to avoid one level of useless inlining that confuses typeless-object detection
	if ref == "#" {
		return obj, nil
	}

	r.resolving[ref] = true
	defer delete(r.resolving, ref)

	// try each scope in order (nearest $id first, root last)
	var target any
	var lastErr error
	for _, scope := range r.scopes {
		t, err := resolveJSONPointer(ref, scope)
		if err == nil {
			target = t
			lastErr = nil
			break
		}
		lastErr = err
	}
	if lastErr != nil {
		return nil, fmt.Errorf("failed to resolve local $ref %q: %w", ref, lastErr)
	}

	resolvedTarget, err := r.resolveNode(cloneNode(target), false, false)
	if err != nil {
		return nil, err
	}

	siblings := make(map[string]any, len(obj))
	for key, val := range obj {
		if key == "$ref" {
			continue
		}
		siblings[key] = val
	}

	if len(siblings) == 0 || r.ignoreRefSiblings {
		return resolvedTarget, nil
	}

	resolvedSiblings, err := r.resolveNode(siblings, false, false)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"allOf": []any{resolvedTarget, resolvedSiblings},
	}, nil
}

func resolveJSONPointer(ref string, root any) (any, error) {
	if ref == "#" {
		return root, nil
	}

	pointer := strings.TrimPrefix(ref, "#")
	if !strings.HasPrefix(pointer, "/") {
		return nil, fmt.Errorf("invalid JSON pointer %q", ref)
	}

	segments := strings.Split(pointer[1:], "/")
	current := root

	for _, segment := range segments {
		decoded, err := url.PathUnescape(segment)
		if err != nil {
			return nil, fmt.Errorf("invalid URI encoding in %q: %w", ref, err)
		}
		segment = strings.ReplaceAll(decoded, "~1", "/")
		segment = strings.ReplaceAll(segment, "~0", "~")

		switch value := current.(type) {
		case map[string]any:
			next, ok := value[segment]
			if !ok {
				return nil, fmt.Errorf("key %q not found", segment)
			}
			current = next
		case []any:
			idx, err := strconv.Atoi(segment)
			if err != nil {
				return nil, fmt.Errorf("%q is not a valid array index", segment)
			}
			if idx < 0 || idx >= len(value) {
				return nil, fmt.Errorf("array index %d out of bounds", idx)
			}
			current = value[idx]
		default:
			return nil, fmt.Errorf("cannot traverse into %T", current)
		}
	}

	return current, nil
}

func cloneNode(node any) any {
	switch v := node.(type) {
	case map[string]any:
		result := make(map[string]any, len(v))
		for key, val := range v {
			result[key] = cloneNode(val)
		}
		return result
	case []any:
		result := make([]any, len(v))
		for i, val := range v {
			result[i] = cloneNode(val)
		}
		return result
	default:
		return node
	}
}

func shouldIgnoreRefSiblings(draft string, root any) bool {
	if draft != "" {
		return legacyRefSiblingDrafts[strings.ToLower(draft)]
	}

	if obj, ok := root.(map[string]any); ok {
		if schemaURI, ok := obj["$schema"].(string); ok {
			if strings.Contains(schemaURI, "draft-03") || strings.Contains(schemaURI, "draft-04") || strings.Contains(schemaURI, "draft-06") || strings.Contains(schemaURI, "draft-07") {
				return true
			}
		}
	}

	return false
}
