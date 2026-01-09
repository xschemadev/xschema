import { Type, type Static } from "@sinclair/typebox";
import { Value } from "@sinclair/typebox/value";
import Ajv from "ajv";
import addFormats from "ajv-formats";

// Use Ajv for validation - supports all JSON Schema keywords including if/then/else, oneOf
const ajv = new Ajv({ strict: false, allErrors: true });
addFormats(ajv);

const schema = {{GENERATED_CODE}};
// Use JSON.parse to ensure __proto__ and other special property names are preserved as own properties
// (directly embedding as JS object literals would interpret __proto__ as prototype setter)
const testCases: Array<{ data: unknown; valid: boolean }> = JSON.parse({{TEST_CASES_STRING}});

// Compile the schema once
const validate = ajv.compile(schema);

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

console.log(JSON.stringify(results));
