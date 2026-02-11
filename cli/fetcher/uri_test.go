package fetcher

import "testing"

func TestResolveURI(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		base    string
		want    string
		wantErr bool
	}{
		// Absolute URL stays absolute
		{
			name: "absolute url with different base",
			ref:  "http://example.com/schema.json",
			base: "http://other.com/base.json",
			want: "http://example.com/schema.json",
		},
		// Relative URL resolved against base
		{
			name: "relative url resolved against base",
			ref:  "other.json",
			base: "http://example.com/schemas/base.json",
			want: "http://example.com/schemas/other.json",
		},
		// File paths
		{
			name: "relative file path resolved against absolute base",
			ref:  "other.json",
			base: "/home/user/schemas/base.json",
			want: "/home/user/schemas/other.json",
		},
		// No base - ref returned as-is
		{
			name: "no base returns ref as-is",
			ref:  "schema.json",
			base: "",
			want: "schema.json",
		},
		// URL with path traversal
		{
			name: "url with path traversal",
			ref:  "../other.json",
			base: "http://example.com/schemas/v1/base.json",
			want: "http://example.com/schemas/other.json",
		},
		// Absolute file path overrides base
		{
			name: "absolute file ref ignores base",
			ref:  "/absolute/path.json",
			base: "/some/other/base.json",
			want: "/absolute/path.json",
		},
		// Fragment-only ref
		{
			name: "fragment-only ref appends to base",
			ref:  "#/definitions/Foo",
			base: "http://example.com/schema.json",
			want: "http://example.com/schema.json#/definitions/Foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveURI(tt.ref, tt.base)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("ResolveURI(%q, %q) = %q, want %q", tt.ref, tt.base, got, tt.want)
			}
		})
	}
}
