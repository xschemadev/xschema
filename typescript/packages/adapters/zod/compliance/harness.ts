import { z } from "zod";

const schema = {{GENERATED_CODE}};
// Use JSON.parse to ensure __proto__ and other special property names are preserved as own properties
// (directly embedding as JS object literals would interpret __proto__ as prototype setter)
const testCases: Array<{ data: unknown; valid: boolean }> = JSON.parse({{TEST_CASES_STRING}});

const results = testCases.map((tc, index) => {
  // safeParse handles validation errors; try-catch guards against malformed generated code
  try {
    const result = schema.safeParse(tc.data);
    return {
      index,
      expected: tc.valid,
      actual: result.success ? "true" : "false",
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

console.log(JSON.stringify(results));
