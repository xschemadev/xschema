package typescript

import (
	"regexp"
	"slices"
	"sort"
	"strings"
)

// MergeImports dedupes and formats TypeScript imports.
func MergeImports(imports []string) string {
	if len(imports) == 0 {
		return ""
	}

	sourceToNames := make(map[string][]string)
	sourceToTypeNames := make(map[string][]string)
	defaultImports := make(map[string]string)
	namespaceImports := make(map[string]string)

	typeOnlyNamedRe := regexp.MustCompile(`import\s+type\s*\{([^}]+)\}\s*from\s*['"]([^'"]+)['"]`)
	namedRe := regexp.MustCompile(`import\s*\{([^}]+)\}\s*from\s*['"]([^'"]+)['"]`)
	namespaceRe := regexp.MustCompile(`import\s*\*\s*as\s+(\w+)\s+from\s*['"]([^'"]+)['"]`)
	defaultRe := regexp.MustCompile(`import\s+(\w+)\s+from\s*['"]([^'"]+)['"]`)
	sideEffectRe := regexp.MustCompile(`import\s*['"]([^'"]+)['"]`)

	var sideEffects []string

	for _, imp := range imports {
		imp = strings.TrimSpace(imp)
		if imp == "" {
			continue
		}

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

		if matches := namespaceRe.FindStringSubmatch(imp); matches != nil {
			namespaceImports[matches[2]] = matches[1]
			continue
		}

		if matches := namedRe.FindStringSubmatch(imp); matches != nil {
			names := strings.Split(matches[1], ",")
			source := matches[2]
			for _, name := range names {
				name = strings.TrimSpace(name)
				if name == "" {
					continue
				}
				if after, ok := strings.CutPrefix(name, "type "); ok {
					typeName := strings.TrimSpace(after)
					if typeName != "" {
						sourceToTypeNames[source] = append(sourceToTypeNames[source], typeName)
					}
				} else {
					sourceToNames[source] = append(sourceToNames[source], name)
				}
			}
			continue
		}

		if matches := defaultRe.FindStringSubmatch(imp); matches != nil {
			defaultImports[matches[2]] = matches[1]
			continue
		}

		if matches := sideEffectRe.FindStringSubmatch(imp); matches != nil {
			sideEffects = append(sideEffects, matches[1])
			continue
		}
	}

	for source, names := range sourceToNames {
		slices.Sort(names)
		sourceToNames[source] = slices.Compact(names)
	}
	for source, names := range sourceToTypeNames {
		slices.Sort(names)
		sourceToTypeNames[source] = slices.Compact(names)
	}

	var result []string

	slices.Sort(sideEffects)
	sideEffects = slices.Compact(sideEffects)
	for _, source := range sideEffects {
		result = append(result, `import "`+source+`"`)
	}

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

	for _, source := range sources {
		var parts []string

		if def, ok := defaultImports[source]; ok {
			parts = append(parts, def)
		}

		var namedParts []string
		if names, ok := sourceToNames[source]; ok {
			namedParts = append(namedParts, names...)
		}
		if typeNames, ok := sourceToTypeNames[source]; ok {
			for _, typeName := range typeNames {
				namedParts = append(namedParts, "type "+typeName)
			}
		}
		if len(namedParts) > 0 {
			sort.Strings(namedParts)
			parts = append(parts, "{ "+strings.Join(namedParts, ", ")+" }")
		}

		if len(parts) > 0 {
			result = append(result, "import "+strings.Join(parts, ", ")+" from \""+source+"\"")
		}
	}

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
