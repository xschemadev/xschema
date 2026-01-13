
import { z } from "zod"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).superRefine((val, ctx) => {
      const schemas = [z.any()];
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = schemas[i].safeParse(val[i]);
        if (!itemResult.success) {
          itemResult.error.issues.forEach(issue => {
            ctx.addIssue({ ...issue, path: [i, ...issue.path] });
          });
        }
      }
    }).superRefine((val, ctx) => {
        const restSchema = z.number().int();
        for (let i = 1; i < val.length; i++) {
          const itemResult = restSchema.safeParse(val[i]);
          if (!itemResult.success) {
            itemResult.error.issues.forEach(issue => {
              ctx.addIssue({ ...issue, path: [i, ...issue.path] });
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
        if (Array.isArray(val)) {
          const result = z.array(z.number().int()).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_2": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_3": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).superRefine((val, ctx) => {
      const schemas = [z.any(), z.any(), z.any()];
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = schemas[i].safeParse(val[i]);
        if (!itemResult.success) {
          itemResult.error.issues.forEach(issue => {
            ctx.addIssue({ ...issue, path: [i, ...issue.path] });
          });
        }
      }
    }).refine((val) => val.length <= 3, { message: "Array must not have more than 3 items" }).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_4": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_5": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).superRefine((val, ctx) => {
      const schemas = [z.number().int()];
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = schemas[i].safeParse(val[i]);
        if (!itemResult.success) {
          itemResult.error.issues.forEach(issue => {
            ctx.addIssue({ ...issue, path: [i, ...issue.path] });
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
  "group_6": (() => {
    const schema = z.intersection(z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).superRefine((val, ctx) => {
      const schemas = [z.number().int()];
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = schemas[i].safeParse(val[i]);
        if (!itemResult.success) {
          itemResult.error.issues.forEach(issue => {
            ctx.addIssue({ ...issue, path: [i, ...issue.path] });
          });
        }
      }
    }).superRefine((val, ctx) => {
        const restSchema = z.boolean();
        for (let i = 1; i < val.length; i++) {
          const itemResult = restSchema.safeParse(val[i]);
          if (!itemResult.success) {
            itemResult.error.issues.forEach(issue => {
              ctx.addIssue({ ...issue, path: [i, ...issue.path] });
            });
          }
        }
      }).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      }), z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).superRefine((val, ctx) => {
      const schemas = [z.number().int(), z.string()];
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = schemas[i].safeParse(val[i]);
        if (!itemResult.success) {
          itemResult.error.issues.forEach(issue => {
            ctx.addIssue({ ...issue, path: [i, ...issue.path] });
          });
        }
      }
    }).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      }));
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_7": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).superRefine((val, ctx) => {
      const schemas = [z.string()];
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = schemas[i].safeParse(val[i]);
        if (!itemResult.success) {
          itemResult.error.issues.forEach(issue => {
            ctx.addIssue({ ...issue, path: [i, ...issue.path] });
          });
        }
      }
    }).superRefine((val, ctx) => {
        const restSchema = z.number().int();
        for (let i = 1; i < val.length; i++) {
          const itemResult = restSchema.safeParse(val[i]);
          if (!itemResult.success) {
            itemResult.error.issues.forEach(issue => {
              ctx.addIssue({ ...issue, path: [i, ...issue.path] });
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
  "group_8": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).superRefine((val, ctx) => {
      const schemas = [z.any()];
      for (let i = 0; i < Math.min(val.length, schemas.length); i++) {
        const itemResult = schemas[i].safeParse(val[i]);
        if (!itemResult.success) {
          itemResult.error.issues.forEach(issue => {
            ctx.addIssue({ ...issue, path: [i, ...issue.path] });
          });
        }
      }
    }).refine((val) => val.length <= 1, { message: "Array must not have more than 1 items" }).safeParse(val);
          if (!result.success) {
            result.error.issues.forEach(issue => ctx.addIssue(issue));
          }
        }
      });
    return { isTypeOnly: false, validate: (data: unknown) => schema.safeParse(data).success };
  })(),
  "group_9": (() => {
    const schema = z.any().superRefine((val, ctx) => {
        if (Array.isArray(val)) {
          const result = z.array(z.any()).safeParse(val);
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
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"additional items match schema\",\"data\":[null,2,3,4],\"valid\":true},{\"description\":\"additional items do not match schema\",\"data\":[null,2,3,\"foo\"],\"valid\":false}]},{\"groupId\":\"group_1\",\"tests\":[{\"description\":\"valid with a array of type integers\",\"data\":[1,2,3],\"valid\":true},{\"description\":\"invalid with a array of mixed types\",\"data\":[1,\"2\",\"3\"],\"valid\":false}]},{\"groupId\":\"group_2\",\"tests\":[{\"description\":\"all items match schema\",\"data\":[1,2,3,4,5],\"valid\":true}]},{\"groupId\":\"group_3\",\"tests\":[{\"description\":\"empty array\",\"data\":[],\"valid\":true},{\"description\":\"fewer number of items present (1)\",\"data\":[1],\"valid\":true},{\"description\":\"fewer number of items present (2)\",\"data\":[1,2],\"valid\":true},{\"description\":\"equal number of items present\",\"data\":[1,2,3],\"valid\":true},{\"description\":\"additional items are not permitted\",\"data\":[1,2,3,4],\"valid\":false}]},{\"groupId\":\"group_4\",\"tests\":[{\"description\":\"items defaults to empty schema so everything is valid\",\"data\":[1,2,3,4,5],\"valid\":true},{\"description\":\"ignores non-arrays\",\"data\":{\"foo\":\"bar\"},\"valid\":true}]},{\"groupId\":\"group_5\",\"tests\":[{\"description\":\"only the first item is validated\",\"data\":[1,\"foo\",false],\"valid\":true}]},{\"groupId\":\"group_6\",\"tests\":[{\"description\":\"items defined in allOf are not examined\",\"data\":[1,\"hello\"],\"valid\":false}]},{\"groupId\":\"group_7\",\"tests\":[{\"description\":\"valid items\",\"data\":[\"x\",2,3],\"valid\":true},{\"description\":\"wrong type of second item\",\"data\":[\"x\",\"y\"],\"valid\":false}]},{\"groupId\":\"group_8\",\"tests\":[{\"description\":\"heterogeneous invalid instance\",\"data\":[\"foo\",\"bar\",37],\"valid\":false},{\"description\":\"valid instance\",\"data\":[null],\"valid\":true}]},{\"groupId\":\"group_9\",\"tests\":[{\"description\":\"allows null elements\",\"data\":[null],\"valid\":true}]}]");

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
