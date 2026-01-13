
import { safeParse } from "valibot"
import * as v from "valibot"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = v.pipe(v.any(), v.check((val) => {
        if (Array.isArray(val)) {
          const result = v.safeParse(v.pipe(v.array(v.any()), v.maxLength(2)), val);
          if (!result.success) return false;
        }
        return true;
      }, "Type-guarded validation failed"));
    return { isTypeOnly: false, validate: (data: unknown) => safeParse(schema, data).success };
  })(),
  "group_1": (() => {
    const schema = v.pipe(v.any(), v.check((val) => {
        if (Array.isArray(val)) {
          const result = v.safeParse(v.pipe(v.array(v.any()), v.maxLength(2)), val);
          if (!result.success) return false;
        }
        return true;
      }, "Type-guarded validation failed"));
    return { isTypeOnly: false, validate: (data: unknown) => safeParse(schema, data).success };
  })(),
};

const testData: Array<{
  groupId: string;
  tests: Array<{ data: unknown; valid: boolean }>;
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"shorter is valid\",\"data\":[1],\"valid\":true},{\"description\":\"exact length is valid\",\"data\":[1,2],\"valid\":true},{\"description\":\"too long is invalid\",\"data\":[1,2,3],\"valid\":false},{\"description\":\"ignores non-arrays\",\"data\":\"foobar\",\"valid\":true}]},{\"groupId\":\"group_1\",\"tests\":[{\"description\":\"shorter is valid\",\"data\":[1],\"valid\":true},{\"description\":\"too long is invalid\",\"data\":[1,2,3],\"valid\":false}]}]");

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
