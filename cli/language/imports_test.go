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
		{
			name: "namespace import",
			imports: []string{
				`import * as v from "valibot"`,
			},
			expected: `import * as v from "valibot"`,
		},
		{
			name: "namespace import with named from same source",
			imports: []string{
				`import * as v from "valibot"`,
				`import { safeParse } from "valibot"`,
			},
			// namespace and named stay separate
			expected: "import { safeParse } from \"valibot\"\nimport * as v from \"valibot\"",
		},
		{
			name: "type-only named import",
			imports: []string{
				`import type { Static } from "@sinclair/typebox"`,
			},
			expected: `import { type Static } from "@sinclair/typebox"`,
		},
		{
			name: "mixed regular and type named imports",
			imports: []string{
				`import { Type, type Static } from "@sinclair/typebox"`,
			},
			expected: `import { Type, type Static } from "@sinclair/typebox"`,
		},
		{
			name: "merge type-only with regular from same source",
			imports: []string{
				`import { Type } from "@sinclair/typebox"`,
				`import type { Static } from "@sinclair/typebox"`,
			},
			expected: `import { Type, type Static } from "@sinclair/typebox"`,
		},
		{
			name: "dedupe type imports",
			imports: []string{
				`import type { Static } from "@sinclair/typebox"`,
				`import type { Static } from "@sinclair/typebox"`,
			},
			expected: `import { type Static } from "@sinclair/typebox"`,
		},
		{
			name: "side-effect import",
			imports: []string{
				`import "reflect-metadata"`,
			},
			expected: `import "reflect-metadata"`,
		},
		{
			name: "side-effect with named",
			imports: []string{
				`import "reflect-metadata"`,
				`import { z } from "zod"`,
			},
			expected: "import \"reflect-metadata\"\nimport { z } from \"zod\"",
		},
		{
			name: "real-world valibot case",
			imports: []string{
				`import * as v from "valibot"`,
				`import { safeParse } from "valibot"`,
				`import * as v from "valibot"`, // duplicate
			},
			expected: "import { safeParse } from \"valibot\"\nimport * as v from \"valibot\"",
		},
		{
			name: "real-world typebox case",
			imports: []string{
				`import { Type, type Static } from "@sinclair/typebox"`,
				`import { Value } from "@sinclair/typebox/value"`,
			},
			expected: "import { Type, type Static } from \"@sinclair/typebox\"\nimport { Value } from \"@sinclair/typebox/value\"",
		},
		{
			name: "arktype type named export",
			imports: []string{
				`import { type } from "arktype"`,
			},
			expected: `import { type } from "arktype"`,
		},
		{
			name: "effect schema alias",
			imports: []string{
				`import { Schema as S } from "effect"`,
			},
			expected: `import { Schema as S } from "effect"`,
		},
		{
			name: "multiple namespace imports from different sources",
			imports: []string{
				`import * as v from "valibot"`,
				`import * as z from "zod"`,
			},
			expected: "import * as v from \"valibot\"\nimport * as z from \"zod\"",
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

func TestBuildPythonFooter(t *testing.T) {
	schemas := []SchemaEntry{
		{Namespace: "user", ID: "User", VarName: "user_User", Code: "class user_User(BaseModel): pass", Type: "user_User"},
		{Namespace: "user", ID: "Post", VarName: "user_Post", Code: "class user_Post(BaseModel): pass", Type: "user_Post"},
	}

	footer := BuildPythonFooter("", schemas)

	// Check that it contains overloads for both schemas with namespace:id format
	if !strings.Contains(footer, `Literal["user:User"]`) {
		t.Error("expected user:User literal in footer")
	}
	if !strings.Contains(footer, `Literal["user:Post"]`) {
		t.Error("expected user:Post literal in footer")
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
