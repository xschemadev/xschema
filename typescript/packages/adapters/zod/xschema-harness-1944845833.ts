
import { z } from "zod"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({ bar: z.string().optional(), foo: z.number().int().optional() }).passthrough().safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_1": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({ bar: z.array(z.any()).optional(), foo: z.array(z.any()).max(3).optional() }).passthrough().superRefine((val, ctx) => {
      const definedProps = new Set(["bar","foo"]);
      const patterns = [new RegExp("f.o")];
    
      for (const [key, value] of Object.entries(val)) {
        if (patterns[0].test(key)) {
          const result = z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).min(2).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      }).safeParse(value);
          if (!result.success) {
            ctx.addIssue({
              code: z.ZodIssueCode.custom,
              path: [key],
              message: "Property matching pattern f.o is invalid",
            });
          }
        }
      }
      for (const [key, value] of Object.entries(val)) {
        if (definedProps.has(key)) continue;
        const matchesPattern = patterns.some(p => p.test(key));
        if (!matchesPattern) {
          const result = z.number().int().safeParse(value);
          if (!result.success) {
            ctx.addIssue({
              code: z.ZodIssueCode.custom,
              path: [key],
              message: "Additional property is invalid",
            });
          }
        }
      }}).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_2": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({ bar: z.never().optional(), foo: z.any().optional() }).passthrough().safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_3": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({ ["foo\tbar"]: z.number().optional(), ["foo\nbar"]: z.number().optional(), ["foo\fbar"]: z.number().optional(), ["foo\rbar"]: z.number().optional(), ["foo\"bar"]: z.number().optional(), ["foo\\bar"]: z.number().optional() }).passthrough().safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_4": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({ foo: z.null().optional() }).passthrough().safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_5": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.any().superRefine((val, ctx) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return;
        if (Object.hasOwn(val, "__proto__")) {
          const result = z.number().safeParse(val["__proto__"]);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue({ ...issue, path: ["__proto__", ...issue.path] }));
          }
        }
        if (Object.hasOwn(val, "constructor")) {
          const result = z.number().safeParse(val["constructor"]);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue({ ...issue, path: ["constructor", ...issue.path] }));
          }
        }
        if (Object.hasOwn(val, "toString")) {
          const result = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({ length: z.string().optional() }).passthrough().safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      }).safeParse(val["toString"]);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue({ ...issue, path: ["toString", ...issue.path] }));
          }
        }
    }).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
};

const testData: Array<{
  groupId: string;
  tests: Array<{ data: unknown; valid: boolean }>;
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"both properties present and valid is valid\",\"data\":{\"bar\":\"baz\",\"foo\":1},\"valid\":true},{\"description\":\"one property invalid is invalid\",\"data\":{\"bar\":{},\"foo\":1},\"valid\":false},{\"description\":\"both properties invalid is invalid\",\"data\":{\"bar\":{},\"foo\":[]},\"valid\":false},{\"description\":\"doesn't invalidate other properties\",\"data\":{\"quux\":[]},\"valid\":true},{\"description\":\"ignores arrays\",\"data\":[],\"valid\":true},{\"description\":\"ignores other non-objects\",\"data\":12,\"valid\":true}]},{\"groupId\":\"group_1\",\"tests\":[{\"description\":\"property validates property\",\"data\":{\"foo\":[1,2]},\"valid\":true},{\"description\":\"property invalidates property\",\"data\":{\"foo\":[1,2,3,4]},\"valid\":false},{\"description\":\"patternProperty invalidates property\",\"data\":{\"foo\":[]},\"valid\":false},{\"description\":\"patternProperty validates nonproperty\",\"data\":{\"fxo\":[1,2]},\"valid\":true},{\"description\":\"patternProperty invalidates nonproperty\",\"data\":{\"fxo\":[]},\"valid\":false},{\"description\":\"additionalProperty ignores property\",\"data\":{\"bar\":[]},\"valid\":true},{\"description\":\"additionalProperty validates others\",\"data\":{\"quux\":3},\"valid\":true},{\"description\":\"additionalProperty invalidates others\",\"data\":{\"quux\":\"foo\"},\"valid\":false}]},{\"groupId\":\"group_2\",\"tests\":[{\"description\":\"no property present is valid\",\"data\":{},\"valid\":true},{\"description\":\"only 'true' property present is valid\",\"data\":{\"foo\":1},\"valid\":true},{\"description\":\"only 'false' property present is invalid\",\"data\":{\"bar\":2},\"valid\":false},{\"description\":\"both properties present is invalid\",\"data\":{\"bar\":2,\"foo\":1},\"valid\":false}]},{\"groupId\":\"group_3\",\"tests\":[{\"description\":\"object with all numbers is valid\",\"data\":{\"foo\\tbar\":1,\"foo\\nbar\":1,\"foo\\fbar\":1,\"foo\\rbar\":1,\"foo\\\"bar\":1,\"foo\\\\bar\":1},\"valid\":true},{\"description\":\"object with strings is invalid\",\"data\":{\"foo\\tbar\":\"1\",\"foo\\nbar\":\"1\",\"foo\\fbar\":\"1\",\"foo\\rbar\":\"1\",\"foo\\\"bar\":\"1\",\"foo\\\\bar\":\"1\"},\"valid\":false}]},{\"groupId\":\"group_4\",\"tests\":[{\"description\":\"allows null values\",\"data\":{\"foo\":null},\"valid\":true}]},{\"groupId\":\"group_5\",\"tests\":[{\"description\":\"ignores arrays\",\"data\":[],\"valid\":true},{\"description\":\"ignores other non-objects\",\"data\":12,\"valid\":true},{\"description\":\"none of the properties mentioned\",\"data\":{},\"valid\":true},{\"description\":\"__proto__ not valid\",\"data\":{\"__proto__\":\"foo\"},\"valid\":false},{\"description\":\"toString not valid\",\"data\":{\"toString\":{\"length\":37}},\"valid\":false},{\"description\":\"constructor not valid\",\"data\":{\"constructor\":{\"length\":37}},\"valid\":false},{\"description\":\"all present and valid\",\"data\":{\"__proto__\":12,\"constructor\":37,\"toString\":{\"length\":\"foo\"}},\"valid\":true}]}]");

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
