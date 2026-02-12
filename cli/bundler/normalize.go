package bundler

import (
	"encoding/json"
	"strings"
)

// nonSchemaKeywords lists JSON Schema keywords whose values are literal data,
// not subschemas. Normalization must not recurse into these — transforming
// their contents would corrupt data (e.g. rewriting $ref strings inside enum).
var nonSchemaKeywords = map[string]bool{
	"const":    true,
	"default":  true,
	"enum":     true,
	"example":  true,
	"examples": true,
}

// normalizeLegacySyntax transforms legacy JSON Schema syntax to draft 2020-12 equivalents.
// This enables compatibility with older drafts (3, 4, 6, 7) by converting deprecated
// keywords to their modern counterparts before bundling.
//
// Transformations applied:
// - items (array) → prefixItems
// - additionalItems → items (when items was array)
// - exclusiveMaximum: true + maximum → exclusiveMaximum: number
// - exclusiveMinimum: true + minimum → exclusiveMinimum: number
// - extends → allOf
// - required: true (property-level) → required: [] (object-level)
// - type: "any" → {} (remove type)
// - divisibleBy → multipleOf
// - id → $id
// - definitions → $defs
func normalizeLegacySyntax(node any) any {
	return normalizeNode(node, nil)
}

// normalizeNode recursively normalizes a schema node.
// parentObj is the parent object (for property-level required collection).
func normalizeNode(node any, parentObj map[string]any) any {
	switch v := node.(type) {
	case map[string]any:
		return normalizeObject(v)
	case []any:
		return normalizeArray(v)
	default:
		return node
	}
}

// normalizeObject applies all transformations to an object schema.
func normalizeObject(obj map[string]any) map[string]any {
	result := make(map[string]any, len(obj))

	// Copy all keys first (we'll modify as needed)
	for k, v := range obj {
		result[k] = v
	}

	// Transform id → $id (draft4 and earlier)
	if id, ok := result["id"]; ok {
		if _, hasNewID := result["$id"]; !hasNewID {
			result["$id"] = id
		}
		delete(result, "id")
	}

	// Transform fragment-only $id (e.g., "#foo") to $anchor (draft4-7 used $id for anchors)
	if id, ok := result["$id"].(string); ok && strings.HasPrefix(id, "#") {
		anchor := id[1:] // strip the #
		if _, hasAnchor := result["$anchor"]; !hasAnchor && anchor != "" {
			result["$anchor"] = anchor
		}
		delete(result, "$id")
	}

	// Transform definitions → $defs
	if defs, ok := result["definitions"]; ok {
		if _, hasNewDefs := result["$defs"]; !hasNewDefs {
			result["$defs"] = defs
		}
		delete(result, "definitions")
	}

	// Rewrite $ref paths: definitions → $defs, items/N → prefixItems/N
	// Must handle ALL occurrences in path (for nested definitions)
	if ref, ok := result["$ref"].(string); ok {
		newRef := ref
		// Replace all /definitions/ with /$defs/ (handles nested definitions)
		newRef = strings.ReplaceAll(newRef, "/definitions/", "/$defs/")
		// Also handle trailing /definitions (no slash after)
		if strings.HasSuffix(newRef, "/definitions") {
			newRef = newRef[:len(newRef)-len("/definitions")] + "/$defs"
		}
		// Replace /items/ with /prefixItems/ (for array-form items refs)
		newRef = strings.ReplaceAll(newRef, "/items/", "/prefixItems/")
		if newRef != ref {
			result["$ref"] = newRef
		}
	}

	// Transform array-form items → prefixItems
	if items, ok := result["items"].([]any); ok {
		// If items array is empty, don't create prefixItems (2020-12 requires minItems: 1)
		// Just set items to additionalItems value if present
		if len(items) == 0 {
			if additionalItems, ok := result["additionalItems"]; ok {
				result["items"] = additionalItems
				delete(result, "additionalItems")
			} else {
				// Empty items with no additionalItems = everything allowed
				delete(result, "items")
			}
		} else {
			result["prefixItems"] = items

			// Transform additionalItems → items
			if additionalItems, ok := result["additionalItems"]; ok {
				result["items"] = additionalItems
				delete(result, "additionalItems")
			} else {
				// No additionalItems means additional items are allowed (default behavior)
				// In 2020-12, omitting "items" when prefixItems exists means additional items allowed
				delete(result, "items")
			}
		}
	}

	// Transform exclusiveMaximum: true + maximum → exclusiveMaximum: number
	if exclusive, ok := result["exclusiveMaximum"].(bool); ok && exclusive {
		if max, ok := result["maximum"]; ok {
			result["exclusiveMaximum"] = max
			delete(result, "maximum")
		}
	} else if exclusive, ok := result["exclusiveMaximum"].(bool); ok && !exclusive {
		// exclusiveMaximum: false just means use maximum normally
		delete(result, "exclusiveMaximum")
	}

	// Transform exclusiveMinimum: true + minimum → exclusiveMinimum: number
	if exclusive, ok := result["exclusiveMinimum"].(bool); ok && exclusive {
		if min, ok := result["minimum"]; ok {
			result["exclusiveMinimum"] = min
			delete(result, "minimum")
		}
	} else if exclusive, ok := result["exclusiveMinimum"].(bool); ok && !exclusive {
		// exclusiveMinimum: false just means use minimum normally
		delete(result, "exclusiveMinimum")
	}

	// Transform divisibleBy → multipleOf (draft3)
	if divisibleBy, ok := result["divisibleBy"]; ok {
		if _, hasMultipleOf := result["multipleOf"]; !hasMultipleOf {
			result["multipleOf"] = divisibleBy
		}
		delete(result, "divisibleBy")
	}

	// Transform type: "any" → remove type constraint (draft3)
	if t, ok := result["type"].(string); ok && t == "any" {
		delete(result, "type")
	}

	// Transform disallow (draft3) to not + type/anyOf
	// disallow: "string" → not: {type: "string"}
	// disallow: ["string", "number"] → not: {anyOf: [{type: "string"}, {type: "number"}]}
	if disallow, ok := result["disallow"]; ok {
		var notSchema map[string]any

		switch d := disallow.(type) {
		case string:
			notSchema = map[string]any{"type": d}
		case []any:
			if len(d) == 1 {
				if s, ok := d[0].(string); ok {
					notSchema = map[string]any{"type": s}
				}
			} else {
				anyOfItems := make([]any, 0, len(d))
				for _, t := range d {
					if s, ok := t.(string); ok {
						anyOfItems = append(anyOfItems, map[string]any{"type": s})
					} else if obj, ok := t.(map[string]any); ok {
						anyOfItems = append(anyOfItems, obj)
					}
				}
				notSchema = map[string]any{"anyOf": anyOfItems}
			}
		}

		if notSchema != nil {
			if existingNot, ok := result["not"].(map[string]any); ok {
				// Combine with allOf if not already exists
				result["not"] = map[string]any{
					"allOf": []any{existingNot, notSchema},
				}
			} else {
				result["not"] = notSchema
			}
		}
		delete(result, "disallow")
	}

	// Transform type array with "any" → remove type constraint
	if types, ok := result["type"].([]any); ok {
		for _, t := range types {
			if ts, ok := t.(string); ok && ts == "any" {
				delete(result, "type")
				break
			}
		}
	}

	// Transform extends → allOf (draft3)
	if extends, ok := result["extends"]; ok {
		var extendsSchemas []any

		switch e := extends.(type) {
		case map[string]any:
			extendsSchemas = []any{e}
		case []any:
			extendsSchemas = e
		}

		if len(extendsSchemas) > 0 {
			if existingAllOf, ok := result["allOf"].([]any); ok {
				// Merge with existing allOf
				result["allOf"] = append(extendsSchemas, existingAllOf...)
			} else {
				result["allOf"] = extendsSchemas
			}
		}
		delete(result, "extends")
	}

	// Transform dependencies (draft3/4) to dependentRequired/dependentSchemas (2020-12)
	// In older drafts:
	// - dependencies: {"bar": "foo"} or {"bar": ["foo"]} → dependentRequired
	// - dependencies: {"bar": {schema}} → dependentSchemas
	if deps, ok := result["dependencies"].(map[string]any); ok {
		dependentRequired := make(map[string]any)
		dependentSchemas := make(map[string]any)

		for prop, dep := range deps {
			switch d := dep.(type) {
			case string:
				// Single string dependency: "bar": "foo" → ["foo"]
				dependentRequired[prop] = []any{d}
			case []any:
				// Array of strings: "bar": ["foo", "baz"]
				dependentRequired[prop] = d
			case map[string]any:
				// Schema dependency
				dependentSchemas[prop] = d
			case bool:
				// Boolean schema dependency: true = always valid, false = always invalid
				dependentSchemas[prop] = d
			}
		}

		if len(dependentRequired) > 0 {
			if existing, ok := result["dependentRequired"].(map[string]any); ok {
				for k, v := range dependentRequired {
					existing[k] = v
				}
			} else {
				result["dependentRequired"] = dependentRequired
			}
		}
		if len(dependentSchemas) > 0 {
			if existing, ok := result["dependentSchemas"].(map[string]any); ok {
				for k, v := range dependentSchemas {
					existing[k] = v
				}
			} else {
				result["dependentSchemas"] = dependentSchemas
			}
		}
		delete(result, "dependencies")
	}

	// Transform type array with inline schemas (draft3) to anyOf
	// e.g., type: ["integer", {properties: {...}}] → anyOf: [{type: "integer"}, {properties: {...}}]
	if types, ok := result["type"].([]any); ok {
		hasSchemas := false
		for _, t := range types {
			if _, isObj := t.(map[string]any); isObj {
				hasSchemas = true
				break
			}
		}

		if hasSchemas {
			anyOfItems := make([]any, 0, len(types))
			for _, t := range types {
				switch v := t.(type) {
				case string:
					anyOfItems = append(anyOfItems, map[string]any{"type": v})
				case map[string]any:
					anyOfItems = append(anyOfItems, v)
				}
			}
			if existingAnyOf, ok := result["anyOf"].([]any); ok {
				result["anyOf"] = append(existingAnyOf, anyOfItems...)
			} else {
				result["anyOf"] = anyOfItems
			}
			delete(result, "type")
		}
	}

	// Transform property-level required: true/false → object-level required array (draft3)
	// This needs special handling: collect required from properties and move to parent
	if props, ok := result["properties"].(map[string]any); ok {
		var requiredProps []string

		for propName, propSchema := range props {
			if propObj, ok := propSchema.(map[string]any); ok {
				if req, ok := propObj["required"].(bool); ok {
					if req {
						requiredProps = append(requiredProps, propName)
					}
					// Remove required: true/false from property (it's not valid in 2020-12)
					delete(propObj, "required")
				}
			}
		}

		if len(requiredProps) > 0 {
			// Merge with existing required array
			if existingRequired, ok := result["required"].([]any); ok {
				for _, prop := range requiredProps {
					existingRequired = append(existingRequired, prop)
				}
				result["required"] = existingRequired
			} else {
				reqAny := make([]any, len(requiredProps))
				for i, p := range requiredProps {
					reqAny[i] = p
				}
				result["required"] = reqAny
			}
		}
	}

	// Recursively normalize all nested schemas, but skip data-only keywords
	// whose values are literal data, not subschemas (e.g. enum, const, default).
	// Normalizing inside these would corrupt data values — for instance, rewriting
	// a literal {"$ref": "#/definitions/a_string"} inside an enum array.
	for k, v := range result {
		if nonSchemaKeywords[k] {
			continue
		}
		result[k] = normalizeNode(v, result)
	}

	return result
}

// normalizeArray normalizes an array of schemas.
func normalizeArray(arr []any) []any {
	result := make([]any, len(arr))
	for i, v := range arr {
		result[i] = normalizeNode(v, nil)
	}
	return result
}

// detectDraftFromSchema detects the JSON Schema draft from $schema field.
// Returns empty string if no draft can be detected.
func detectDraftFromSchema(schema map[string]any) string {
	schemaURI, ok := schema["$schema"].(string)
	if !ok {
		return ""
	}

	if strings.Contains(schemaURI, "draft-03") {
		return "draft3"
	}
	if strings.Contains(schemaURI, "draft-04") {
		return "draft4"
	}
	if strings.Contains(schemaURI, "draft-06") {
		return "draft6"
	}
	if strings.Contains(schemaURI, "draft-07") {
		return "draft7"
	}
	if strings.Contains(schemaURI, "draft/2019-09") || strings.Contains(schemaURI, "draft-2019-09") {
		return "draft2019-09"
	}
	if strings.Contains(schemaURI, "draft/2020-12") || strings.Contains(schemaURI, "draft-2020-12") {
		return "draft2020-12"
	}

	return ""
}

// needsNormalization returns true if the schema draft requires legacy syntax normalization.
func needsNormalization(draft string) bool {
	switch draft {
	case "draft3", "draft4", "draft6", "draft7":
		return true
	default:
		return false
	}
}

// NeedsNormalization is the exported version for use by the processor.
func NeedsNormalization(draft string) bool {
	return needsNormalization(draft)
}

// NormalizeSchema normalizes a legacy draft schema to draft 2020-12.
// Returns the normalized schema as JSON bytes.
func NormalizeSchema(schema []byte, draft string) ([]byte, error) {
	var parsed any
	if err := json.Unmarshal(schema, &parsed); err != nil {
		return nil, err
	}

	// Inject $schema if missing
	if obj, ok := parsed.(map[string]any); ok {
		if _, hasSchema := obj["$schema"]; !hasSchema {
			if schemaURI, known := draftToSchemaURI[draft]; known {
				obj["$schema"] = schemaURI
			}
		}
	}

	// Normalize the schema
	normalized := normalizeLegacySyntax(parsed)

	// Update $schema to 2020-12
	if obj, ok := normalized.(map[string]any); ok {
		obj["$schema"] = draftToSchemaURI["draft2020-12"]
	}

	return json.Marshal(normalized)
}
