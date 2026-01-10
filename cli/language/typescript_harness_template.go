package language

// TSHarnessTemplate is a Go text/template for TypeScript compliance harness
const TSHarnessTemplate = `{{- if .Imports}}
{{.Imports}}
{{- end}}
{{- if .IsTypeOnly}}
// Type-only adapter harness
// Validates that generated TypeScript types are syntactically valid
// but cannot run runtime validation tests since types are compile-time only

type GeneratedType = {{.Type}};

const testCases: Array<{ data: unknown; valid: boolean }> = JSON.parse({{.TestCasesString}});

const results = testCases.map((tc, index) => ({
  index,
  expected: tc.valid,
  actual: "skipped",
}));
{{- else}}
const schema = {{.Schema}};

const testCases: Array<{ data: unknown; valid: boolean }> = JSON.parse({{.TestCasesString}});

const validate = {{.Validate}};

const results = testCases.map((tc, index) => {
  try {
    const isValid = validate(tc.data);
    return {
      index,
      expected: tc.valid,
      actual: isValid ? "true" : "false",
    };
  } catch (e) {
    return {
      index,
      expected: tc.valid,
      actual: "error",
      error: e instanceof Error ? e.message : String(e),
    };
  }
});
{{- end}}

console.log(JSON.stringify(results));
`
