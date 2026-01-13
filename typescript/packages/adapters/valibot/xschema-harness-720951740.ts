
import { safeParse } from "valibot"
import * as v from "valibot"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = v.intersect([v.pipe(v.any(), v.check((val) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = v.safeParse(v.pipe(v.unknown(), v.rawCheck((ctx) => { if (Array.isArray(ctx.dataset.value)) ctx.addIssue({ message: "Expected object, not array" }); }), v.looseObject({ foo: v.optional(v.pipe(v.number(), v.integer())) })), val);
          if (!result.success) return false;
        }
        return true;
      }, "Type-guarded validation failed")), v.pipe(v.any(), v.check((val) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = v.safeParse(v.pipe(v.unknown(), v.rawCheck((ctx) => { if (Array.isArray(ctx.dataset.value)) ctx.addIssue({ message: "Expected object, not array" }); }), v.record(v.string(), v.pipe(v.number(), v.integer()))), val);
          if (!result.success) return false;
        }
        return true;
      }, "Type-guarded validation failed"))]);
    return { isTypeOnly: false, validate: (data: unknown) => safeParse(schema, data).success };
  })(),
};

const testData: Array<{
  groupId: string;
  tests: Array<{ data: unknown; valid: boolean }>;
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"passing case\",\"data\":{\"foo\":1},\"valid\":true},{\"description\":\"failing case\",\"data\":{\"foo\":\"a string\"},\"valid\":false}]}]");

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
