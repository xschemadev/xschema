
import { Schema as S } from "effect"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
    const result = S.decodeUnknownEither(S.Number.pipe(S.int()))(val);
    return result._tag === "Left";
  }, { message: () => "Value must not match schema" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_1": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
    const result = S.decodeUnknownEither(S.Union(S.Number.pipe(S.int()), S.Boolean))(val);
    return result._tag === "Left";
  }, { message: () => "Value must not match schema" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_2": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
    const result = S.decodeUnknownEither(S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;
        if (Object.hasOwn(val, "foo")) {
          const result = S.decodeUnknownEither(S.String)(val.foo);
          if (result._tag === "Left") return false;
        }
      return true;
    }, { message: () => "Object validation failed" })))(val);
    return result._tag === "Left";
  }, { message: () => "Value must not match schema" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_3": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = S.decodeUnknownEither(S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;
        if (Object.hasOwn(val, "foo")) {
          const result = S.decodeUnknownEither(S.Never)(val.foo);
          if (result._tag === "Left") return false;
        }
      return true;
    }, { message: () => "Object validation failed" })))(val);
          if (result._tag === "Left") return false;
        }
        return true;
      }, { message: () => "Type-guarded validation failed" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_4": (() => {
    const schema = S.Never;
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_5": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
    const result = S.decodeUnknownEither(S.Unknown)(val);
    return result._tag === "Left";
  }, { message: () => "Value must not match schema" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_6": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
    const result = S.decodeUnknownEither(S.Never)(val);
    return result._tag === "Left";
  }, { message: () => "Value must not match schema" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_7": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
    const result = S.decodeUnknownEither(S.Never)(val);
    return result._tag === "Left";
  }, { message: () => "Value must not match schema" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
};

const testData: Array<{
  groupId: string;
  tests: Array<{ data: unknown; valid: boolean }>;
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"allowed\",\"data\":\"foo\",\"valid\":true},{\"description\":\"disallowed\",\"data\":1,\"valid\":false}]},{\"groupId\":\"group_1\",\"tests\":[{\"description\":\"valid\",\"data\":\"foo\",\"valid\":true},{\"description\":\"mismatch\",\"data\":1,\"valid\":false},{\"description\":\"other mismatch\",\"data\":true,\"valid\":false}]},{\"groupId\":\"group_2\",\"tests\":[{\"description\":\"match\",\"data\":1,\"valid\":true},{\"description\":\"other match\",\"data\":{\"foo\":1},\"valid\":true},{\"description\":\"mismatch\",\"data\":{\"foo\":\"bar\"},\"valid\":false}]},{\"groupId\":\"group_3\",\"tests\":[{\"description\":\"property present\",\"data\":{\"bar\":2,\"foo\":1},\"valid\":false},{\"description\":\"property absent\",\"data\":{\"bar\":1,\"baz\":2},\"valid\":true}]},{\"groupId\":\"group_4\",\"tests\":[{\"description\":\"number is invalid\",\"data\":1,\"valid\":false},{\"description\":\"string is invalid\",\"data\":\"foo\",\"valid\":false},{\"description\":\"boolean true is invalid\",\"data\":true,\"valid\":false},{\"description\":\"boolean false is invalid\",\"data\":false,\"valid\":false},{\"description\":\"null is invalid\",\"data\":null,\"valid\":false},{\"description\":\"object is invalid\",\"data\":{\"foo\":\"bar\"},\"valid\":false},{\"description\":\"empty object is invalid\",\"data\":{},\"valid\":false},{\"description\":\"array is invalid\",\"data\":[\"foo\"],\"valid\":false},{\"description\":\"empty array is invalid\",\"data\":[],\"valid\":false}]},{\"groupId\":\"group_5\",\"tests\":[{\"description\":\"number is invalid\",\"data\":1,\"valid\":false},{\"description\":\"string is invalid\",\"data\":\"foo\",\"valid\":false},{\"description\":\"boolean true is invalid\",\"data\":true,\"valid\":false},{\"description\":\"boolean false is invalid\",\"data\":false,\"valid\":false},{\"description\":\"null is invalid\",\"data\":null,\"valid\":false},{\"description\":\"object is invalid\",\"data\":{\"foo\":\"bar\"},\"valid\":false},{\"description\":\"empty object is invalid\",\"data\":{},\"valid\":false},{\"description\":\"array is invalid\",\"data\":[\"foo\"],\"valid\":false},{\"description\":\"empty array is invalid\",\"data\":[],\"valid\":false}]},{\"groupId\":\"group_6\",\"tests\":[{\"description\":\"number is valid\",\"data\":1,\"valid\":true},{\"description\":\"string is valid\",\"data\":\"foo\",\"valid\":true},{\"description\":\"boolean true is valid\",\"data\":true,\"valid\":true},{\"description\":\"boolean false is valid\",\"data\":false,\"valid\":true},{\"description\":\"null is valid\",\"data\":null,\"valid\":true},{\"description\":\"object is valid\",\"data\":{\"foo\":\"bar\"},\"valid\":true},{\"description\":\"empty object is valid\",\"data\":{},\"valid\":true},{\"description\":\"array is valid\",\"data\":[\"foo\"],\"valid\":true},{\"description\":\"empty array is valid\",\"data\":[],\"valid\":true}]},{\"groupId\":\"group_7\",\"tests\":[{\"description\":\"any value is valid\",\"data\":\"foo\",\"valid\":true}]}]");

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
