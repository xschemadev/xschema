package python

import "testing"

func TestBuildVarName(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		id        string
		expected  string
	}{
		{
			name:      "simple",
			namespace: "user",
			id:        "User",
			expected:  "user_user",
		},
		{
			name:      "camelCase to snake_case",
			namespace: "userProfile",
			id:        "CreateUser",
			expected:  "user_profile_create_user",
		},
		{
			name:      "dotted namespace",
			namespace: "user.xschema",
			id:        "User",
			expected:  "user_xschema_user",
		},
		{
			name:      "namespace with digits",
			namespace: "billing.v2",
			id:        "invoice-item",
			expected:  "billing_v_2_invoice_item",
		},
		{
			name:      "leading digit in result",
			namespace: "123",
			id:        "User",
			expected:  "_123_user",
		},
		{
			name:      "reserved word in id only",
			namespace: "auth",
			id:        "class",
			expected:  "auth_class",
		},
		{
			name:      "reserved word in id only type",
			namespace: "schema",
			id:        "type",
			expected:  "schema_type",
		},
		{
			name:      "PascalCase id",
			namespace: "api",
			id:        "UserResponse",
			expected:  "api_user_response",
		},
		{
			name:      "kebab-case",
			namespace: "api-v2",
			id:        "user-create",
			expected:  "api_v_2_user_create",
		},
		{
			name:      "all uppercase acronym",
			namespace: "api",
			id:        "URLValue",
			expected:  "api_url_value",
		},
		{
			name:      "multiple acronyms",
			namespace: "http",
			id:        "HTTPSURLParser",
			expected:  "http_httpsurl_parser",
		},
		{
			name:      "snake_case input",
			namespace: "user_data",
			id:        "create_user",
			expected:  "user_data_create_user",
		},
		{
			name:      "mixed separators",
			namespace: "user-data.v1",
			id:        "Create_User-Request",
			expected:  "user_data_v_1_create_user_request",
		},
		{
			name:      "empty namespace fallback",
			namespace: "",
			id:        "User",
			expected:  "schema_user",
		},
		{
			name:      "empty id fallback",
			namespace: "user",
			id:        "",
			expected:  "user_schema",
		},
		{
			name:      "both empty fallback",
			namespace: "",
			id:        "",
			expected:  "schema",
		},
		{
			name:      "digits only namespace",
			namespace: "123",
			id:        "456",
			expected:  "_123_456",
		},
		{
			name:      "reserved word import in id",
			namespace: "module",
			id:        "import",
			expected:  "module_import",
		},
		{
			name:      "builtin list in id",
			namespace: "data",
			id:        "list",
			expected:  "data_list",
		},
		{
			name:      "underscore prefix not needed",
			namespace: "user",
			id:        "profile",
			expected:  "user_profile",
		},
		{
			name:      "special characters stripped",
			namespace: "user@domain",
			id:        "email!check",
			expected:  "user_domain_email_check",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildVarName(tt.namespace, tt.id)
			if got != tt.expected {
				t.Fatalf("buildVarName(%q, %q) = %q; want %q", tt.namespace, tt.id, got, tt.expected)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"User", "user"},
		{"UserProfile", "user_profile"},
		{"userProfile", "user_profile"},
		{"user_profile", "user_profile"},
		{"user-profile", "user_profile"},
		{"user.profile", "user_profile"},
		{"URLParser", "url_parser"},
		{"HTTPSConnection", "https_connection"},
		{"getHTTPSURL", "get_httpsurl"},
		{"API", "api"},
		{"APIv2", "ap_iv_2"},
		{"v2API", "v_2_api"},
		{"123", "123"},
		{"user123", "user_123"},
		{"123user", "123_user"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toSnakeCase(tt.input)
			if got != tt.expected {
				t.Errorf("toSnakeCase(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestSplitIdentifierTokens(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"user", []string{"user"}},
		{"userName", []string{"user", "Name"}},
		{"user_name", []string{"user", "name"}},
		{"user-name", []string{"user", "name"}},
		{"user.name", []string{"user", "name"}},
		{"URLValue", []string{"URL", "Value"}},
		{"getURL", []string{"get", "URL"}},
		{"user123", []string{"user", "123"}},
		{"123user", []string{"123", "user"}},
		{"API", []string{"API"}},
		{"", nil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := splitIdentifierTokens(tt.input)
			if len(got) != len(tt.expected) {
				t.Fatalf("splitIdentifierTokens(%q) = %v; want %v", tt.input, got, tt.expected)
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Fatalf("splitIdentifierTokens(%q) = %v; want %v", tt.input, got, tt.expected)
				}
			}
		})
	}
}

func TestPythonReservedWords(t *testing.T) {
	// Test that common Python keywords are in the reserved words map
	keywords := []string{
		"False", "None", "True", "and", "as", "assert", "async", "await",
		"break", "class", "continue", "def", "del", "elif", "else", "except",
		"finally", "for", "from", "global", "if", "import", "in", "is",
		"lambda", "nonlocal", "not", "or", "pass", "raise", "return", "try",
		"while", "with", "yield", "match", "case", "type",
	}

	for _, kw := range keywords {
		if !pythonReservedWords[kw] {
			t.Errorf("expected %q to be a Python reserved word", kw)
		}
	}

	// Test common builtins
	builtins := []string{
		"id", "list", "dict", "set", "str", "int", "float", "bool",
		"print", "len", "range", "map", "filter",
	}

	for _, b := range builtins {
		if !pythonReservedWords[b] {
			t.Errorf("expected %q to be in Python reserved words (builtin)", b)
		}
	}
}
