package typescript

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
			expected:  "user_User",
		},
		{
			name:      "dotted namespace",
			namespace: "user.xschema",
			id:        "User",
			expected:  "userXschema_User",
		},
		{
			name:      "namespace with digits",
			namespace: "billing.v2",
			id:        "invoice-item",
			expected:  "billingV2_InvoiceItem",
		},
		{
			name:      "leading digit",
			namespace: "123",
			id:        "User",
			expected:  "_123_User",
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
