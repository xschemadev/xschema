package language

import "testing"

func TestNormalizeRelativePath(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{name: "file", in: "xschema.gen.ts", want: "xschema.gen.ts"},
		{name: "nested", in: "schemas/xschema.gen.ts", want: "schemas/xschema.gen.ts"},
		{name: "backslashes", in: "schemas\\xschema.gen.ts", want: "schemas/xschema.gen.ts"},
		{name: "empty", in: "", wantErr: true},
		{name: "dot", in: ".", wantErr: true},
		{name: "dotdot", in: "..", wantErr: true},
		{name: "escapes", in: "../a.ts", wantErr: true},
		{name: "contains dot segment", in: "a/./b.ts", wantErr: true},
		{name: "contains dotdot segment", in: "a/../b.ts", wantErr: true},
		{name: "empty segment", in: "a//b.ts", wantErr: true},
		{name: "absolute", in: "/tmp/a.ts", wantErr: true},
		{name: "windows drive", in: "C:\\tmp\\a.ts", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeRelativePath(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NormalizeRelativePath() err=%v wantErr=%v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Fatalf("NormalizeRelativePath()=%q want=%q", got, tt.want)
			}
		})
	}
}
