package language

import (
	"regexp"
	"slices"
	"sort"
	"strings"
)

// MergeTSImports dedupes and formats TypeScript imports
// Input: ["import { z } from \"zod\"", "import { z } from \"zod\"", "import { foo } from \"bar\""]
// Output: "import { z } from \"zod\"\nimport { foo } from \"bar\""
//
// Handles all real-world TS import forms:
// - Named imports: import { A, B } from "x"
// - Type-only named imports: import type { A } from "x"
// - Mixed type imports: import { A, type B } from "x"
// - Namespace imports: import * as v from "x" (kept separate from named)
// - Default imports: import React from "react"
// - Side-effect imports: import "x"
func MergeTSImports(imports []string) string {
	if len(imports) == 0 {
		return ""
	}

	// Separate tracking for different import kinds
	sourceToNames := make(map[string][]string)     // source -> regular named imports
	sourceToTypeNames := make(map[string][]string) // source -> type-only named imports
	defaultImports := make(map[string]string)      // source -> default import name
	namespaceImports := make(map[string]string)    // source -> namespace alias (e.g. "v" for * as v)

	// Regex patterns
	// Type-only import: import type { A, B } from "x"
	typeOnlyNamedRe := regexp.MustCompile(`import\s+type\s*\{([^}]+)\}\s*from\s*['"]([^'"]+)['"]`)
	// Named imports: import { A, type B } from "x"
	namedRe := regexp.MustCompile(`import\s*\{([^}]+)\}\s*from\s*['"]([^'"]+)['"]`)
	// Namespace import: import * as v from "x"
	namespaceRe := regexp.MustCompile(`import\s*\*\s*as\s+(\w+)\s+from\s*['"]([^'"]+)['"]`)
	// Default import: import x from "source"
	defaultRe := regexp.MustCompile(`import\s+(\w+)\s+from\s*['"]([^'"]+)['"]`)
	// Side-effect import: import "source"
	sideEffectRe := regexp.MustCompile(`import\s*['"]([^'"]+)['"]`)

	var sideEffects []string

	for _, imp := range imports {
		imp = strings.TrimSpace(imp)
		if imp == "" {
			continue
		}

		// Try type-only named imports: import type { A, B } from "source"
		if matches := typeOnlyNamedRe.FindStringSubmatch(imp); matches != nil {
			names := strings.Split(matches[1], ",")
			source := matches[2]
			for _, name := range names {
				name = strings.TrimSpace(name)
				if name != "" {
					sourceToTypeNames[source] = append(sourceToTypeNames[source], name)
				}
			}
			continue
		}

		// Try namespace import: import * as v from "source"
		if matches := namespaceRe.FindStringSubmatch(imp); matches != nil {
			namespaceImports[matches[2]] = matches[1]
			continue
		}

		// Try named imports: import { x, y, type z } from "source"
		if matches := namedRe.FindStringSubmatch(imp); matches != nil {
			names := strings.Split(matches[1], ",")
			source := matches[2]
			for _, name := range names {
				name = strings.TrimSpace(name)
				if name == "" {
					continue
				}
				// Check for "type X" prefix within named imports
				if after, ok := strings.CutPrefix(name, "type "); ok {
					typeName := after
					typeName = strings.TrimSpace(typeName)
					if typeName != "" {
						sourceToTypeNames[source] = append(sourceToTypeNames[source], typeName)
					}
				} else {
					sourceToNames[source] = append(sourceToNames[source], name)
				}
			}
			continue
		}

		// Try default import: import x from "source"
		if matches := defaultRe.FindStringSubmatch(imp); matches != nil {
			defaultImports[matches[2]] = matches[1]
			continue
		}

		// Try side-effect import: import "source"
		if matches := sideEffectRe.FindStringSubmatch(imp); matches != nil {
			sideEffects = append(sideEffects, matches[1])
			continue
		}
	}

	// Dedupe named imports per source
	for source, names := range sourceToNames {
		slices.Sort(names)
		sourceToNames[source] = slices.Compact(names)
	}
	for source, names := range sourceToTypeNames {
		slices.Sort(names)
		sourceToTypeNames[source] = slices.Compact(names)
	}

	// Build output
	var result []string

	// Side-effect imports first
	slices.Sort(sideEffects)
	sideEffects = slices.Compact(sideEffects)
	for _, source := range sideEffects {
		result = append(result, `import "`+source+`"`)
	}

	// Collect all sources that have named/default/type-only imports
	sourceSet := make(map[string]bool)
	for source := range sourceToNames {
		sourceSet[source] = true
	}
	for source := range sourceToTypeNames {
		sourceSet[source] = true
	}
	for source := range defaultImports {
		sourceSet[source] = true
	}
	var sources []string
	for source := range sourceSet {
		sources = append(sources, source)
	}
	sort.Strings(sources)

	// Build import statements for default + named imports
	for _, source := range sources {
		var parts []string

		// Default import
		if def, ok := defaultImports[source]; ok {
			parts = append(parts, def)
		}

		// Named imports (regular + type prefixed)
		var namedParts []string
		if names, ok := sourceToNames[source]; ok {
			namedParts = append(namedParts, names...)
		}
		if typeNames, ok := sourceToTypeNames[source]; ok {
			for _, tn := range typeNames {
				namedParts = append(namedParts, "type "+tn)
			}
		}
		if len(namedParts) > 0 {
			// Sort the combined named parts for determinism
			sort.Strings(namedParts)
			parts = append(parts, "{ "+strings.Join(namedParts, ", ")+" }")
		}

		if len(parts) > 0 {
			result = append(result, "import "+strings.Join(parts, ", ")+" from \""+source+"\"")
		}
	}

	// Namespace imports - separate statements, cannot be merged with named
	var nsSources []string
	for source := range namespaceImports {
		nsSources = append(nsSources, source)
	}
	sort.Strings(nsSources)
	for _, source := range nsSources {
		alias := namespaceImports[source]
		result = append(result, "import * as "+alias+" from \""+source+"\"")
	}

	return strings.Join(result, "\n")
}
