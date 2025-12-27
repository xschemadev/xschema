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
func MergeTSImports(imports []string) string {
	if len(imports) == 0 {
		return ""
	}

	// Parse imports into source -> named imports map
	// e.g. "zod" -> ["z"], "bar" -> ["foo"]
	sourceToNames := make(map[string][]string)
	defaultImports := make(map[string]string) // source -> default import name

	// Regex patterns
	namedRe := regexp.MustCompile(`import\s*\{([^}]+)\}\s*from\s*['"]([^'"]+)['"]`)
	defaultRe := regexp.MustCompile(`import\s+(\w+)\s+from\s*['"]([^'"]+)['"]`)
	sideEffectRe := regexp.MustCompile(`import\s*['"]([^'"]+)['"]`)

	var sideEffects []string

	for _, imp := range imports {
		imp = strings.TrimSpace(imp)
		if imp == "" {
			continue
		}

		// Try named imports: import { x, y } from "source"
		if matches := namedRe.FindStringSubmatch(imp); matches != nil {
			names := strings.Split(matches[1], ",")
			source := matches[2]
			for _, name := range names {
				name = strings.TrimSpace(name)
				if name != "" {
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

	// Build output
	var result []string

	// Side-effect imports first
	slices.Sort(sideEffects)
	sideEffects = slices.Compact(sideEffects)
	for _, source := range sideEffects {
		result = append(result, `import "`+source+`"`)
	}

	// Collect all sources and sort
	var sources []string
	for source := range sourceToNames {
		sources = append(sources, source)
	}
	for source := range defaultImports {
		if _, exists := sourceToNames[source]; !exists {
			sources = append(sources, source)
		}
	}
	sort.Strings(sources)

	// Build import statements
	for _, source := range sources {
		var parts []string

		// Default import
		if def, ok := defaultImports[source]; ok {
			parts = append(parts, def)
		}

		// Named imports
		if names, ok := sourceToNames[source]; ok && len(names) > 0 {
			parts = append(parts, "{ "+strings.Join(names, ", ")+" }")
		}

		if len(parts) > 0 {
			result = append(result, "import "+strings.Join(parts, ", ")+" from \""+source+"\"")
		}
	}

	return strings.Join(result, "\n")
}

// MergePyImports dedupes and formats Python imports
// Input: ["from pydantic import BaseModel", "from pydantic import Field", "from uuid import UUID"]
// Output: "from pydantic import BaseModel, Field\nfrom uuid import UUID"
func MergePyImports(imports []string) string {
	if len(imports) == 0 {
		return ""
	}

	// Parse: "from X import Y" or "import X"
	fromImports := make(map[string][]string) // module -> names
	directImports := []string{}

	fromRe := regexp.MustCompile(`from\s+(\S+)\s+import\s+(.+)`)
	importRe := regexp.MustCompile(`^import\s+(\S+)`)

	for _, imp := range imports {
		imp = strings.TrimSpace(imp)
		if imp == "" {
			continue
		}

		if matches := fromRe.FindStringSubmatch(imp); matches != nil {
			module := matches[1]
			names := strings.SplitSeq(matches[2], ",")
			for name := range names {
				name = strings.TrimSpace(name)
				if name != "" {
					fromImports[module] = append(fromImports[module], name)
				}
			}
			continue
		}

		if matches := importRe.FindStringSubmatch(imp); matches != nil {
			directImports = append(directImports, matches[1])
		}
	}

	// Dedupe
	for module, names := range fromImports {
		slices.Sort(names)
		fromImports[module] = slices.Compact(names)
	}
	slices.Sort(directImports)
	directImports = slices.Compact(directImports)

	// Build output
	var result []string

	// Direct imports first
	for _, mod := range directImports {
		result = append(result, "import "+mod)
	}

	// From imports, sorted by module
	var modules []string
	for mod := range fromImports {
		modules = append(modules, mod)
	}
	sort.Strings(modules)

	for _, mod := range modules {
		names := fromImports[mod]
		result = append(result, "from "+mod+" import "+strings.Join(names, ", "))
	}

	return strings.Join(result, "\n")
}

// BuildPythonFooter generates Python overload stubs for type safety
func BuildPythonFooter(_ string, schemas []SchemaEntry) string {
	if len(schemas) == 0 {
		return ""
	}

	var lines []string

	// from_url overloads
	for _, s := range schemas {
		lines = append(lines, `    @staticmethod
    @overload
    def from_url(name: Literal["`+s.Name+`"], url: str, adapter: XSchemaAdapter) -> type[`+s.Name+`]: ...`)
	}
	lines = append(lines, `    @staticmethod
    @overload
    def from_url(name: str, url: str, adapter: XSchemaAdapter) -> type: ...`)
	lines = append(lines, `
    @staticmethod
    def from_url(name: str, url: str, adapter: XSchemaAdapter) -> type:
        _ = url, adapter
        return _schemas[name]`)

	lines = append(lines, "")

	// from_file overloads
	for _, s := range schemas {
		lines = append(lines, `    @staticmethod
    @overload
    def from_file(name: Literal["`+s.Name+`"], path: str, adapter: XSchemaAdapter) -> type[`+s.Name+`]: ...`)
	}
	lines = append(lines, `    @staticmethod
    @overload
    def from_file(name: str, path: str, adapter: XSchemaAdapter) -> type: ...`)
	lines = append(lines, `
    @staticmethod
    def from_file(name: str, path: str, adapter: XSchemaAdapter) -> type:
        _ = path, adapter
        return _schemas[name]`)

	return strings.Join(lines, "\n")
}
