package language

import (
	"strings"
	"testing"
)

func TestMergeTSImports(t *testing.T) {
	tests := []struct {
		name     string
		imports  []string
		expected string
	}{
		{
			name:     "empty",
			imports:  []string{},
			expected: "",
		},
		{
			name: "dedupe same import",
			imports: []string{
				`import { z } from "zod"`,
				`import { z } from "zod"`,
			},
			expected: `import { z } from "zod"`,
		},
		{
			name: "merge named imports from same source",
			imports: []string{
				`import { z } from "zod"`,
				`import { ZodError } from "zod"`,
			},
			expected: `import { ZodError, z } from "zod"`,
		},
		{
			name: "multiple sources",
			imports: []string{
				`import { z } from "zod"`,
				`import { foo } from "bar"`,
			},
			expected: "import { foo } from \"bar\"\nimport { z } from \"zod\"",
		},
		{
			name: "default import",
			imports: []string{
				`import React from "react"`,
			},
			expected: `import React from "react"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeTSImports(tt.imports)
			if got != tt.expected {
				t.Errorf("MergeTSImports() =\n%q\nwant\n%q", got, tt.expected)
			}
		})
	}
}

func TestMergePyImports(t *testing.T) {
	tests := []struct {
		name     string
		imports  []string
		expected string
	}{
		{
			name:     "empty",
			imports:  []string{},
			expected: "",
		},
		{
			name: "dedupe same import",
			imports: []string{
				`from pydantic import BaseModel`,
				`from pydantic import BaseModel`,
			},
			expected: `from pydantic import BaseModel`,
		},
		{
			name: "merge from same module",
			imports: []string{
				`from pydantic import BaseModel`,
				`from pydantic import Field`,
			},
			expected: `from pydantic import BaseModel, Field`,
		},
		{
			name: "multiple modules",
			imports: []string{
				`from pydantic import BaseModel`,
				`from uuid import UUID`,
			},
			expected: "from pydantic import BaseModel\nfrom uuid import UUID",
		},
		{
			name: "direct import",
			imports: []string{
				`import json`,
			},
			expected: `import json`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergePyImports(tt.imports)
			if got != tt.expected {
				t.Errorf("MergePyImports() =\n%q\nwant\n%q", got, tt.expected)
			}
		})
	}
}

func TestMergeGoImports(t *testing.T) {
	tests := []struct {
		name     string
		imports  []string
		expected string
	}{
		{
			name:     "empty",
			imports:  []string{},
			expected: "",
		},
		{
			name: "single import",
			imports: []string{
				`import "fmt"`,
			},
			expected: `import "fmt"`,
		},
		{
			name: "dedupe same import",
			imports: []string{
				`import "fmt"`,
				`import "fmt"`,
			},
			expected: `import "fmt"`,
		},
		{
			name: "multiple imports",
			imports: []string{
				`import "fmt"`,
				`import "strings"`,
			},
			expected: "import (\n\t\"fmt\"\n\t\"strings\"\n)",
		},
		{
			name: "aliased import",
			imports: []string{
				`import v "github.com/go-playground/validator/v10"`,
			},
			expected: `import v "github.com/go-playground/validator/v10"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeGoImports(tt.imports)
			if got != tt.expected {
				t.Errorf("MergeGoImports() =\n%q\nwant\n%q", got, tt.expected)
			}
		})
	}
}

func TestBuildGoHeader(t *testing.T) {
	tests := []struct {
		outDir   string
		expected string
	}{
		{"xschema", "package xschema"},
		{"/home/user/project/xschema", "package xschema"},
		{"schemas", "package schemas"},
	}

	for _, tt := range tests {
		t.Run(tt.outDir, func(t *testing.T) {
			got := BuildGoHeader(tt.outDir, nil)
			if got != tt.expected {
				t.Errorf("BuildGoHeader(%q) = %q, want %q", tt.outDir, got, tt.expected)
			}
		})
	}
}

func TestBuildPythonFooter(t *testing.T) {
	schemas := []SchemaEntry{
		{Name: "User", Code: "class User(BaseModel): pass", Type: "User"},
		{Name: "Post", Code: "class Post(BaseModel): pass", Type: "Post"},
	}

	footer := BuildPythonFooter("", schemas)

	// Check that it contains overloads for both schemas
	if !strings.Contains(footer, `Literal["User"]`) {
		t.Error("expected User literal in footer")
	}
	if !strings.Contains(footer, `Literal["Post"]`) {
		t.Error("expected Post literal in footer")
	}
	if !strings.Contains(footer, "from_url") {
		t.Error("expected from_url in footer")
	}
	if !strings.Contains(footer, "from_file") {
		t.Error("expected from_file in footer")
	}
}

func TestBuildPythonFooterEmpty(t *testing.T) {
	footer := BuildPythonFooter("", nil)
	if footer != "" {
		t.Errorf("expected empty footer for no schemas, got %q", footer)
	}
}
