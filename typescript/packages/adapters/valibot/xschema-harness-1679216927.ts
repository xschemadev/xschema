
import { safeParse } from "valibot"
import * as v from "valibot"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = v.pipe(v.any(), v.check((val) => {
        if (typeof val === "number") {
          const result = v.safeParse(v.pipe(v.number(), v.minValue(1.1)), val);
          if (!result.success) return false;
        }
        return true;
      }, "Type-guarded validation failed"));
    return { isTypeOnly: false, validate: (data: unknown) => safeParse(schema, data).success };
  })(),
  "group_1": (() => {
    const schema = v.pipe(v.any(), v.check((val) => {
        if (typeof val === "number") {
          const result = v.safeParse(v.pipe(v.number(), v.minValue(-2)), val);
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
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"above the minimum is valid\",\"data\":2.6,\"valid\":true},{\"description\":\"boundary point is valid\",\"data\":1.1,\"valid\":true},{\"description\":\"below the minimum is invalid\",\"data\":0.6,\"valid\":false},{\"description\":\"ignores non-numbers\",\"data\":\"x\",\"valid\":true}]},{\"groupId\":\"group_1\",\"tests\":[{\"description\":\"negative above the minimum is valid\",\"data\":-1,\"valid\":true},{\"description\":\"positive above the minimum is valid\",\"data\":0,\"valid\":true},{\"description\":\"boundary point is valid\",\"data\":-2,\"valid\":true},{\"description\":\"boundary point with float is valid\",\"data\":-2,\"valid\":true},{\"description\":\"float below the minimum is invalid\",\"data\":-2.0001,\"valid\":false},{\"description\":\"int below the minimum is invalid\",\"data\":-3,\"valid\":false},{\"description\":\"ignores non-numbers\",\"data\":\"x\",\"valid\":true}]}]");

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
