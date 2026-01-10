package language

// TSHarnessTemplate is a Go text/template for TypeScript compliance harness
const TSHarnessTemplate = `{{- if .Imports}}
{{.Imports}}
{{- end}}

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
{{- range $i, $s := .Schemas}}
{{- if $s.IsTypeOnly}}
  "{{$s.GroupID}}": (() => {
    type GeneratedType = {{$s.Type}};
    return { isTypeOnly: true, validate: () => true };
  })(),
{{- else}}
  "{{$s.GroupID}}": (() => {
    const schema = {{$s.Schema}};
    return { isTypeOnly: false, validate: {{$s.Validate}} };
  })(),
{{- end}}
{{- end}}
};

const testData: Array<{
  groupId: string;
  tests: Array<{ data: unknown; valid: boolean }>;
}> = JSON.parse({{.TestDataString}});

const results = testData.flatMap(({ groupId, tests }) => {
  const entry = schemas[groupId];
  if (!entry) {
    return tests.map((tc, index) => ({
      groupId,
      index,
      expected: tc.valid,
      actual: "error" as const,
      error: "schema not found for group " + groupId,
    }));
  }

  if (entry.isTypeOnly) {
    return tests.map((tc, index) => ({
      groupId,
      index,
      expected: tc.valid,
      actual: "skipped" as const,
    }));
  }

  return tests.map((tc, index) => {
    try {
      const isValid = entry.validate(tc.data);
      return {
        groupId,
        index,
        expected: tc.valid,
        actual: isValid ? ("true" as const) : ("false" as const),
      };
    } catch (e) {
      return {
        groupId,
        index,
        expected: tc.valid,
        actual: "error" as const,
        error: e instanceof Error ? e.message : String(e),
      };
    }
  });
});

console.log(JSON.stringify(results));
`
