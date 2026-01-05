import { z } from "zod";

const schema = {{GENERATED_CODE}};
const testCases: Array<{ data: unknown; valid: boolean }> = {{TEST_CASES}};

const results = testCases.map((tc, index) => {
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
