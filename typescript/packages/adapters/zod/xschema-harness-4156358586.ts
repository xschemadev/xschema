
import { z } from "zod"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({}).passthrough().superRefine((val, ctx) => {
      const definedProps = new Set([]);
      const patterns = [];
    }).superRefine((val, ctx) => {
      for (const key of Object.keys(val)) {
        const result = z.any().superRefine((val, ctx) => {
        if (typeof val === "string") {
          const result = z.string().refine((val) => [...new Intl.Segmenter().segment(val)].length <= 3, { message: "String must have at most 3 character(s)" }).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      }).safeParse(key);
        if (!result.success) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            path: [key],
            message: "Invalid property name",
          });
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
  "group_1": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({}).passthrough().superRefine((val, ctx) => {
      const definedProps = new Set([]);
      const patterns = [];
    }).superRefine((val, ctx) => {
      for (const key of Object.keys(val)) {
        const result = z.any().superRefine((val, ctx) => {
        if (typeof val === "string") {
          const result = z.string().regex(new RegExp("^a+$")).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      }).safeParse(key);
        if (!result.success) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            path: [key],
            message: "Invalid property name",
          });
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
  "group_2": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({}).passthrough().superRefine((val, ctx) => {
      const definedProps = new Set([]);
      const patterns = [];
    }).superRefine((val, ctx) => {
      for (const key of Object.keys(val)) {
        const result = z.any().safeParse(key);
        if (!result.success) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            path: [key],
            message: "Invalid property name",
          });
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
  "group_3": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({}).passthrough().superRefine((val, ctx) => {
      const definedProps = new Set([]);
      const patterns = [];
    }).superRefine((val, ctx) => {
      for (const key of Object.keys(val)) {
        const result = z.never().safeParse(key);
        if (!result.success) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            path: [key],
            message: "Invalid property name",
          });
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
  "group_4": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({}).passthrough().superRefine((val, ctx) => {
      const definedProps = new Set([]);
      const patterns = [];
    }).superRefine((val, ctx) => {
      for (const key of Object.keys(val)) {
        const result = z.literal("foo").safeParse(key);
        if (!result.success) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            path: [key],
            message: "Invalid property name",
          });
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
  "group_5": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = z.object({}).passthrough().superRefine((val, ctx) => {
      const definedProps = new Set([]);
      const patterns = [];
    }).superRefine((val, ctx) => {
      for (const key of Object.keys(val)) {
        const result = z.enum(["foo", "bar"]).safeParse(key);
        if (!result.success) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            path: [key],
            message: "Invalid property name",
          });
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
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"all property names valid\",\"data\":{\"f\":{},\"foo\":{}},\"valid\":true},{\"description\":\"some property names invalid\",\"data\":{\"foo\":{},\"foobar\":{}},\"valid\":false},{\"description\":\"object without properties is valid\",\"data\":{},\"valid\":true},{\"description\":\"ignores arrays\",\"data\":[1,2,3,4],\"valid\":true},{\"description\":\"ignores strings\",\"data\":\"foobar\",\"valid\":true},{\"description\":\"ignores other non-objects\",\"data\":12,\"valid\":true}]},{\"groupId\":\"group_1\",\"tests\":[{\"description\":\"matching property names valid\",\"data\":{\"a\":{},\"aa\":{},\"aaa\":{}},\"valid\":true},{\"description\":\"non-matching property name is invalid\",\"data\":{\"aaA\":{}},\"valid\":false},{\"description\":\"object without properties is valid\",\"data\":{},\"valid\":true}]},{\"groupId\":\"group_2\",\"tests\":[{\"description\":\"object with any properties is valid\",\"data\":{\"foo\":1},\"valid\":true},{\"description\":\"empty object is valid\",\"data\":{},\"valid\":true}]},{\"groupId\":\"group_3\",\"tests\":[{\"description\":\"object with any properties is invalid\",\"data\":{\"foo\":1},\"valid\":false},{\"description\":\"empty object is valid\",\"data\":{},\"valid\":true}]},{\"groupId\":\"group_4\",\"tests\":[{\"description\":\"object with property foo is valid\",\"data\":{\"foo\":1},\"valid\":true},{\"description\":\"object with any other property is invalid\",\"data\":{\"bar\":1},\"valid\":false},{\"description\":\"empty object is valid\",\"data\":{},\"valid\":true}]},{\"groupId\":\"group_5\",\"tests\":[{\"description\":\"object with property foo is valid\",\"data\":{\"foo\":1},\"valid\":true},{\"description\":\"object with property foo and bar is valid\",\"data\":{\"bar\":1,\"foo\":1},\"valid\":true},{\"description\":\"object with any other property is invalid\",\"data\":{\"baz\":1},\"valid\":false},{\"description\":\"empty object is valid\",\"data\":{},\"valid\":true}]}]");

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
