// Type-only adapter harness
// This harness validates that the generated TypeScript types are syntactically valid
// but cannot run runtime validation tests since types are compile-time only

// The generated type (checked at compile-time by TypeScript compiler)
type GeneratedType = {{GENERATED_CODE}};

// Test cases from JSON Schema Test Suite - parsed but not validated at runtime
const testCases: Array<{ data: unknown; valid: boolean }> = JSON.parse({{TEST_CASES_STRING}});

// Type-only adapters cannot validate data at runtime
// All tests are marked as "skipped" to indicate type-only mode
const results = testCases.map((tc, index) => ({
  index,
  expected: tc.valid,
  actual: "skipped",
}));

console.log(JSON.stringify(results));
