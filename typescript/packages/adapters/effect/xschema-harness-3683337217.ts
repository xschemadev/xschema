
import { Schema as S } from "effect"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = S.decodeUnknownEither(S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;
        if (Object.keys(val).length > 2) return false;
      return true;
    }, { message: () => "Object validation failed" })))(val);
          if (result._tag === "Left") return false;
        }
        return true;
      }, { message: () => "Type-guarded validation failed" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_1": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = S.decodeUnknownEither(S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;
        if (Object.keys(val).length > 2) return false;
      return true;
    }, { message: () => "Object validation failed" })))(val);
          if (result._tag === "Left") return false;
        }
        return true;
      }, { message: () => "Type-guarded validation failed" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_2": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = S.decodeUnknownEither(S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;
        if (Object.keys(val).length > 0) return false;
      return true;
    }, { message: () => "Object validation failed" })))(val);
          if (result._tag === "Left") return false;
        }
        return true;
      }, { message: () => "Type-guarded validation failed" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
};

const testData: Array<{
  groupId: string;
  tests: Array<{ data: unknown; valid: boolean }>;
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"shorter is valid\",\"data\":{\"foo\":1},\"valid\":true},{\"description\":\"exact length is valid\",\"data\":{\"bar\":2,\"foo\":1},\"valid\":true},{\"description\":\"too long is invalid\",\"data\":{\"bar\":2,\"baz\":3,\"foo\":1},\"valid\":false},{\"description\":\"ignores arrays\",\"data\":[1,2,3],\"valid\":true},{\"description\":\"ignores strings\",\"data\":\"foobar\",\"valid\":true},{\"description\":\"ignores other non-objects\",\"data\":12,\"valid\":true}]},{\"groupId\":\"group_1\",\"tests\":[{\"description\":\"shorter is valid\",\"data\":{\"foo\":1},\"valid\":true},{\"description\":\"too long is invalid\",\"data\":{\"bar\":2,\"baz\":3,\"foo\":1},\"valid\":false}]},{\"groupId\":\"group_2\",\"tests\":[{\"description\":\"no properties is valid\",\"data\":{},\"valid\":true},{\"description\":\"one property is invalid\",\"data\":{\"foo\":1},\"valid\":false}]}]");

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
