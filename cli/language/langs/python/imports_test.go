package python

import "testing"

func TestMergeImports(t *testing.T) {
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
			name: "typing imports grouped",
			imports: []string{
				`from typing import Optional`,
				`from typing import List`,
				`from typing import Dict`,
			},
			expected: `from typing import Dict, List, Optional`,
		},
		{
			name: "stdlib before third party",
			imports: []string{
				`from pydantic import BaseModel`,
				`from typing import Optional`,
			},
			expected: "from typing import Optional\n\nfrom pydantic import BaseModel",
		},
		{
			name: "module import",
			imports: []string{
				`import json`,
			},
			expected: `import json`,
		},
		{
			name: "module import with alias",
			imports: []string{
				`import numpy as np`,
			},
			expected: `import numpy as np`,
		},
		{
			name: "from import with alias preserved",
			imports: []string{
				`from pydantic import Field as F`,
			},
			expected: `from pydantic import Field as F`,
		},
		{
			name: "local imports last",
			imports: []string{
				`from pydantic import BaseModel`,
				`from .models import User`,
			},
			expected: "from pydantic import BaseModel\n\nfrom .models import User",
		},
		{
			name: "real world pydantic case",
			imports: []string{
				`from pydantic import BaseModel`,
				`from pydantic import Field`,
				`from pydantic import ConfigDict`,
				`from typing import Optional`,
				`from typing import List`,
			},
			expected: "from typing import List, Optional\n\nfrom pydantic import BaseModel, ConfigDict, Field",
		},
		{
			name: "multiple import styles mixed",
			imports: []string{
				`import json`,
				`from typing import Any`,
				`from pydantic import BaseModel`,
				`import numpy as np`,
			},
			expected: "import json\nfrom typing import Any\n\nimport numpy as np\nfrom pydantic import BaseModel",
		},
		{
			name: "stdlib module imports before from imports",
			imports: []string{
				`from json import loads`,
				`import json`,
			},
			expected: "import json\nfrom json import loads",
		},
		{
			name: "relative imports at end",
			imports: []string{
				`from typing import Dict`,
				`from . import utils`,
				`from pydantic import BaseModel`,
			},
			expected: "from typing import Dict\n\nfrom pydantic import BaseModel\n\nfrom . import utils",
		},
		{
			name: "dedupe names within same module",
			imports: []string{
				`from pydantic import BaseModel`,
				`from pydantic import Field`,
				`from pydantic import BaseModel`,
				`from pydantic import Field`,
			},
			expected: `from pydantic import BaseModel, Field`,
		},
		{
			name: "submodule imports categorized correctly",
			imports: []string{
				`from typing import Optional`,
				`from collections.abc import Sequence`,
				`from pydantic.fields import FieldInfo`,
			},
			expected: "from collections.abc import Sequence\nfrom typing import Optional\n\nfrom pydantic.fields import FieldInfo",
		},
		{
			name: "empty strings ignored",
			imports: []string{
				``,
				`from typing import Any`,
				``,
			},
			expected: `from typing import Any`,
		},
		{
			name: "whitespace trimmed",
			imports: []string{
				`  from typing import Any  `,
				`	from pydantic import BaseModel	`,
			},
			expected: "from typing import Any\n\nfrom pydantic import BaseModel",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeImports(tt.imports)
			if got != tt.expected {
				t.Errorf("MergeImports() =\n%q\nwant\n%q", got, tt.expected)
			}
		})
	}
}

func TestIsStdlibModule(t *testing.T) {
	tests := []struct {
		module   string
		expected bool
	}{
		{"typing", true},
		{"json", true},
		{"os", true},
		{"collections", true},
		{"collections.abc", true},
		{"typing_extensions", false},
		{"pydantic", false},
		{"numpy", false},
		{"fastapi", false},
		{"datetime", true},
		{"pathlib", true},
		{"dataclasses", true},
	}

	for _, tt := range tests {
		t.Run(tt.module, func(t *testing.T) {
			got := isStdlibModule(tt.module)
			if got != tt.expected {
				t.Errorf("isStdlibModule(%q) = %v; want %v", tt.module, got, tt.expected)
			}
		})
	}
}
