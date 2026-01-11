package language

import "testing"

func TestSortGeneratedFiles(t *testing.T) {
	files := []GeneratedFile{
		{Path: "b.ts", Contents: "b"},
		{Path: "a.ts", Contents: "a"},
		{Path: "c.ts", Contents: "c"},
	}

	SortGeneratedFiles(files)

	got := []string{files[0].Path, files[1].Path, files[2].Path}
	want := []string{"a.ts", "b.ts", "c.ts"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("sorted[%d]=%q want=%q", i, got[i], want[i])
		}
	}
}

func TestValidateGeneratedFiles(t *testing.T) {
	if err := ValidateGeneratedFiles([]GeneratedFile{{Path: "a.ts"}, {Path: "a.ts"}}); err == nil {
		t.Fatalf("expected duplicate path error")
	}
}
