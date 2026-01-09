import { Type, type Static } from "@sinclair/typebox";
import { Value } from "@sinclair/typebox/value";

const schema = {{GENERATED_CODE}};
// Use JSON.parse to ensure __proto__ and other special property names are preserved as own properties
// (directly embedding as JS object literals would interpret __proto__ as prototype setter)
const testCases: Array<{ data: unknown; valid: boolean }> = JSON.parse({{TEST_CASES_STRING}});

const results = testCases.map((tc, index) => {
  // Value.Check handles validation; try-catch guards against malformed generated code
  try {
    const isValid = Value.Check(schema, tc.data);
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

console.log(JSON.stringify(results));
