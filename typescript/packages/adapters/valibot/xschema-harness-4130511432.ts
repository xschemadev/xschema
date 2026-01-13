
import { safeParse } from "valibot"
import * as v from "valibot"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = v.any();
    return { isTypeOnly: false, validate: (data: unknown) => safeParse(schema, data).success };
  })(),
  "group_1": (() => {
    const schema = v.pipe(v.any(), v.check((val) => {
        if (Array.isArray(val)) {
          const result = v.safeParse(v.pipe(v.array(v.any()), v.check((arr) => {
        let count = 0;
        for (const item of arr) {
          if (v.safeParse(v.literal(1), item).success) count++;
        }
        return count >= 1 && count <= 1;
      }, "Array must contain between 1 and 1 items matching schema")), val);
          if (!result.success) return false;
        }
        return true;
      }, "Type-guarded validation failed"));
    return { isTypeOnly: false, validate: (data: unknown) => safeParse(schema, data).success };
  })(),
  "group_2": (() => {
    const schema = v.pipe(v.any(), v.check((val) => {
        if (Array.isArray(val)) {
          const result = v.safeParse(v.pipe(v.array(v.any()), v.check((arr) => {
        let count = 0;
        for (const item of arr) {
          if (v.safeParse(v.literal(1), item).success) count++;
        }
        return count >= 1 && count <= 1;
      }, "Array must contain between 1 and 1 items matching schema")), val);
          if (!result.success) return false;
        }
        return true;
      }, "Type-guarded validation failed"));
    return { isTypeOnly: false, validate: (data: unknown) => safeParse(schema, data).success };
  })(),
  "group_3": (() => {
    const schema = v.pipe(v.any(), v.check((val) => {
        if (Array.isArray(val)) {
          const result = v.safeParse(v.pipe(v.array(v.any()), v.check((arr) => {
        let count = 0;
        for (const item of arr) {
          if (v.safeParse(v.literal(1), item).success) count++;
        }
        return count >= 1 && count <= 3;
      }, "Array must contain between 1 and 3 items matching schema")), val);
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
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"one item valid against lone maxContains\",\"data\":[1],\"valid\":true},{\"description\":\"two items still valid against lone maxContains\",\"data\":[1,2],\"valid\":true}]},{\"groupId\":\"group_1\",\"tests\":[{\"description\":\"empty data\",\"data\":[],\"valid\":false},{\"description\":\"all elements match, valid maxContains\",\"data\":[1],\"valid\":true},{\"description\":\"all elements match, invalid maxContains\",\"data\":[1,1],\"valid\":false},{\"description\":\"some elements match, valid maxContains\",\"data\":[1,2],\"valid\":true},{\"description\":\"some elements match, invalid maxContains\",\"data\":[1,2,1],\"valid\":false}]},{\"groupId\":\"group_2\",\"tests\":[{\"description\":\"one element matches, valid maxContains\",\"data\":[1],\"valid\":true},{\"description\":\"too many elements match, invalid maxContains\",\"data\":[1,1],\"valid\":false}]},{\"groupId\":\"group_3\",\"tests\":[{\"description\":\"actual \\u003c minContains \\u003c maxContains\",\"data\":[],\"valid\":false},{\"description\":\"minContains \\u003c actual \\u003c maxContains\",\"data\":[1,1],\"valid\":true},{\"description\":\"minContains \\u003c maxContains \\u003c actual\",\"data\":[1,1,1,1],\"valid\":false}]}]");

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
