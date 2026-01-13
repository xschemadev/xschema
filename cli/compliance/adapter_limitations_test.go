package compliance

import "testing"

func TestMatchesUnsupportedPattern(t *testing.T) {
	tests := []struct {
		name         string
		adapterName  string
		description  string
		wantMatches  bool
		wantContains string // substring expected in reason (if matches)
	}{
		{
			name:         "valibot JS property names - properties",
			adapterName:  "@xschemadev/valibot",
			description:  "properties whose names are Javascript object property names",
			wantMatches:  true,
			wantContains: "valibot library limitation",
		},
		{
			name:         "valibot JS property names - required",
			adapterName:  "@xschemadev/valibot",
			description:  "required properties whose names are Javascript object property names",
			wantMatches:  true,
			wantContains: "JS prototype",
		},
		{
			name:        "valibot normal properties",
			adapterName: "@xschemadev/valibot",
			description: "object properties validation",
			wantMatches: false,
		},
		{
			name:        "zod not affected",
			adapterName: "@xschemadev/zod",
			description: "properties whose names are Javascript object property names",
			wantMatches: false,
		},
		{
			name:        "arktype not affected",
			adapterName: "@xschemadev/arktype",
			description: "properties whose names are Javascript object property names",
			wantMatches: false,
		},
		{
			name:        "effect not affected",
			adapterName: "@xschemadev/effect",
			description: "properties whose names are Javascript object property names",
			wantMatches: false,
		},
		{
			name:        "unknown adapter",
			adapterName: "unknown-adapter",
			description: "properties whose names are Javascript object property names",
			wantMatches: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches, reason := MatchesUnsupportedPattern(tt.adapterName, tt.description)
			if matches != tt.wantMatches {
				t.Errorf("MatchesUnsupportedPattern(%q, %q) matches = %v, want %v", tt.adapterName, tt.description, matches, tt.wantMatches)
			}
			if tt.wantMatches && tt.wantContains != "" {
				if reason == "" {
					t.Errorf("MatchesUnsupportedPattern(%q, %q) reason is empty, want contains %q", tt.adapterName, tt.description, tt.wantContains)
				} else if !containsString(reason, tt.wantContains) {
					t.Errorf("MatchesUnsupportedPattern(%q, %q) reason = %q, want contains %q", tt.adapterName, tt.description, reason, tt.wantContains)
				}
			}
		})
	}
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestGetAdapterLimitation(t *testing.T) {
	tests := []struct {
		name        string
		adapterName string
		wantNil     bool
	}{
		{
			name:        "valibot has limitations",
			adapterName: "@xschemadev/valibot",
			wantNil:     false,
		},
		{
			name:        "zod has no limitations",
			adapterName: "@xschemadev/zod",
			wantNil:     true,
		},
		{
			name:        "unknown adapter",
			adapterName: "unknown",
			wantNil:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lim := GetAdapterLimitation(tt.adapterName)
			isNil := lim == nil
			if isNil != tt.wantNil {
				t.Errorf("GetAdapterLimitation(%q) isNil = %v, want %v", tt.adapterName, isNil, tt.wantNil)
			}
		})
	}
}
