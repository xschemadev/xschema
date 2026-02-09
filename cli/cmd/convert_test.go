package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/xschemadev/xschema/adapter"
	_ "github.com/xschemadev/xschema/language/langs"
)

// adapterCLIPath returns the absolute path to the zod adapter CLI.
// Returns empty string if the adapter isn't built.
func adapterCLIPath() string {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	// cli/cmd/convert_test.go -> project root is ../../
	cliDir := filepath.Dir(filepath.Dir(thisFile))
	p := filepath.Join(cliDir, "..", "typescript", "packages", "adapters", "zod", "dist", "cli.js")
	if _, err := os.Stat(p); err != nil {
		return ""
	}
	abs, _ := filepath.Abs(p)
	return abs
}

// setupTestProject creates a temp dir that looks like a bun project with the
// adapter binary symlinked into node_modules/.bin. Returns the project dir.
func setupTestProject(t *testing.T) string {
	t.Helper()

	cliPath := adapterCLIPath()
	if cliPath == "" {
		t.Skip("zod adapter CLI not built; run bun run build from typescript/")
	}
	if _, err := exec.LookPath("bunx"); err != nil {
		t.Skip("bunx not found in PATH")
	}

	dir := t.TempDir()

	// Minimal package.json so runner detection works
	if err := os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{}`), 0644); err != nil {
		t.Fatalf("write package.json: %v", err)
	}
	// bun.lock so detectRunnerInDir picks bunx
	if err := os.WriteFile(filepath.Join(dir, "bun.lock"), []byte(`{}`), 0644); err != nil {
		t.Fatalf("write bun.lock: %v", err)
	}
	// Symlink adapter binary
	binDir := filepath.Join(dir, "node_modules", ".bin")
	if err := os.MkdirAll(binDir, 0755); err != nil {
		t.Fatalf("mkdir bin: %v", err)
	}
	if err := os.Symlink(cliPath, filepath.Join(binDir, "xschema-zod")); err != nil {
		t.Fatalf("symlink adapter: %v", err)
	}

	return dir
}

// executeConvert runs the convert command with the given args and stdin, returning stdout, stderr, and error.
func executeConvert(t *testing.T, args []string, stdin string) (stdout, stderr string, err error) {
	t.Helper()

	var outBuf, errBuf bytes.Buffer
	rootCmd.SetOut(&outBuf)
	rootCmd.SetErr(&errBuf)
	rootCmd.SetIn(strings.NewReader(stdin))
	rootCmd.SetArgs(args)

	err = rootCmd.ExecuteContext(context.Background())

	// Reset io for other tests
	rootCmd.SetOut(nil)
	rootCmd.SetErr(nil)
	rootCmd.SetIn(nil)

	return outBuf.String(), errBuf.String(), err
}

func TestConvert_ValidSingleSchema(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	projectDir := setupTestProject(t)

	input := `[{"namespace":"test","id":"User","schema":{"type":"object","properties":{"name":{"type":"string"}}}}]`
	args := []string{"convert", "--adapter", "@xschemadev/zod", "--project", projectDir}

	stdout, _, err := executeConvert(t, args, input)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	var results []adapter.ConvertResult
	if err := json.Unmarshal([]byte(stdout), &results); err != nil {
		t.Fatalf("failed to parse stdout as JSON: %v\nstdout: %s", err, stdout)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	r := results[0]
	if r.Namespace != "test" {
		t.Errorf("namespace: got %q, want %q", r.Namespace, "test")
	}
	if r.ID != "User" {
		t.Errorf("id: got %q, want %q", r.ID, "User")
	}
	if r.Schema == "" {
		t.Error("schema should not be empty")
	}
	if r.Type == "" {
		t.Error("type should not be empty")
	}
	if r.VarName == "" {
		t.Error("varName should not be empty")
	}
}

func TestConvert_MultipleSchemas(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	projectDir := setupTestProject(t)

	input := `[
		{"namespace":"users","id":"User","schema":{"type":"object","properties":{"name":{"type":"string"}}}},
		{"namespace":"posts","id":"Post","schema":{"type":"object","properties":{"title":{"type":"string"}}}}
	]`
	args := []string{"convert", "--adapter", "@xschemadev/zod", "--project", projectDir}

	stdout, _, err := executeConvert(t, args, input)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	var results []adapter.ConvertResult
	if err := json.Unmarshal([]byte(stdout), &results); err != nil {
		t.Fatalf("failed to parse stdout as JSON: %v\nstdout: %s", err, stdout)
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	keys := map[string]bool{}
	for _, r := range results {
		keys[r.Key()] = true
		if r.Schema == "" {
			t.Errorf("schema empty for %s", r.Key())
		}
		if r.Type == "" {
			t.Errorf("type empty for %s", r.Key())
		}
	}
	if !keys["users:User"] {
		t.Error("missing users:User in results")
	}
	if !keys["posts:Post"] {
		t.Error("missing posts:Post in results")
	}
}

func TestConvert_MissingAdapterFlag(t *testing.T) {
	args := []string{"convert"}
	_, _, err := executeConvert(t, args, `[{"namespace":"t","id":"T","schema":{"type":"string"}}]`)
	if err == nil {
		t.Fatal("expected error for missing --adapter flag")
	}
	if !strings.Contains(err.Error(), "adapter") {
		t.Errorf("error should mention adapter flag, got: %v", err)
	}
}

func TestConvert_InvalidJSONStdin(t *testing.T) {
	args := []string{"convert", "--adapter", "@xschemadev/zod"}
	_, stderr, err := executeConvert(t, args, `not valid json`)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
	if !strings.Contains(err.Error(), "invalid JSON") {
		t.Errorf("error should mention invalid JSON, got: %v", err)
	}
	// stderr should contain JSON error object
	if !strings.Contains(stderr, "invalid JSON") {
		t.Errorf("stderr should contain JSON error, got: %s", stderr)
	}
}

func TestConvert_EmptyArrayStdin(t *testing.T) {
	args := []string{"convert", "--adapter", "@xschemadev/zod"}
	stdout, _, err := executeConvert(t, args, `[]`)
	if err != nil {
		t.Fatalf("expected no error for empty array, got: %v", err)
	}
	trimmed := strings.TrimSpace(stdout)
	if trimmed != "[]" {
		t.Errorf("expected empty array output, got: %q", trimmed)
	}
}

func TestConvert_ExternalRefWithoutAllowFetch(t *testing.T) {
	// Schema with an external $ref should fail without --allow-fetch
	// and the error message should mention --allow-fetch
	input := `[{"namespace":"test","id":"Ref","schema":{"$ref":"https://example.com/schema.json"}}]`
	args := []string{"convert", "--adapter", "@xschemadev/zod"}

	_, stderr, err := executeConvert(t, args, input)
	if err == nil {
		t.Fatal("expected error for external $ref without --allow-fetch")
	}
	if !strings.Contains(err.Error(), "--allow-fetch") {
		t.Errorf("error should mention --allow-fetch, got: %v", err)
	}
	if !strings.Contains(stderr, "--allow-fetch") {
		t.Errorf("stderr should mention --allow-fetch, got: %s", stderr)
	}
}
