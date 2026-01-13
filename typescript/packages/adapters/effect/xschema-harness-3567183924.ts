
import { Schema as S } from "effect"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = S.decodeUnknownEither(S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;
        if (Object.hasOwn(val, "bar")) {
          const result = S.decodeUnknownEither(S.Unknown)(val.bar);
          if (result._tag === "Left") return false;
        }
        if (!Object.hasOwn(val, "foo")) return false;
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
        if (Object.hasOwn(val, "foo")) {
          const result = S.decodeUnknownEither(S.Unknown)(val.foo);
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
  "group_2": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = S.decodeUnknownEither(S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;
        if (Object.hasOwn(val, "foo")) {
          const result = S.decodeUnknownEither(S.Unknown)(val.foo);
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
  "group_3": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = S.decodeUnknownEither(S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;
        if (!Object.hasOwn(val, "foo\nbar")) return false;
        if (!Object.hasOwn(val, "foo\"bar")) return false;
        if (!Object.hasOwn(val, "foo\\bar")) return false;
        if (!Object.hasOwn(val, "foo\rbar")) return false;
        if (!Object.hasOwn(val, "foo\tbar")) return false;
        if (!Object.hasOwn(val, "foo\fbar")) return false;
      return true;
    }, { message: () => "Object validation failed" })))(val);
          if (result._tag === "Left") return false;
        }
        return true;
      }, { message: () => "Type-guarded validation failed" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_4": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
        if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          const result = S.decodeUnknownEither(S.Unknown.pipe(S.filter((val) => {
      if (typeof val !== "object" || val === null || Array.isArray(val)) return false;
        if (Object.hasOwn(val, "__proto__")) {
          const result = S.decodeUnknownEither(S.Unknown)(val["__proto__"]);
          if (result._tag === "Left") return false;
        } else {
          return false;
        }
        if (Object.hasOwn(val, "toString")) {
          const result = S.decodeUnknownEither(S.Unknown)(val["toString"]);
          if (result._tag === "Left") return false;
        } else {
          return false;
        }
        if (Object.hasOwn(val, "constructor")) {
          const result = S.decodeUnknownEither(S.Unknown)(val["constructor"]);
          if (result._tag === "Left") return false;
        } else {
          return false;
        }
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
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"present required property is valid\",\"data\":{\"foo\":1},\"valid\":true},{\"description\":\"non-present required property is invalid\",\"data\":{\"bar\":1},\"valid\":false},{\"description\":\"ignores arrays\",\"data\":[],\"valid\":true},{\"description\":\"ignores strings\",\"data\":\"\",\"valid\":true},{\"description\":\"ignores other non-objects\",\"data\":12,\"valid\":true}]},{\"groupId\":\"group_1\",\"tests\":[{\"description\":\"not required by default\",\"data\":{},\"valid\":true}]},{\"groupId\":\"group_2\",\"tests\":[{\"description\":\"property not required\",\"data\":{},\"valid\":true}]},{\"groupId\":\"group_3\",\"tests\":[{\"description\":\"object with all properties present is valid\",\"data\":{\"foo\\tbar\":1,\"foo\\nbar\":1,\"foo\\fbar\":1,\"foo\\rbar\":1,\"foo\\\"bar\":1,\"foo\\\\bar\":1},\"valid\":true},{\"description\":\"object with some properties missing is invalid\",\"data\":{\"foo\\nbar\":\"1\",\"foo\\\"bar\":\"1\"},\"valid\":false}]},{\"groupId\":\"group_4\",\"tests\":[{\"description\":\"ignores arrays\",\"data\":[],\"valid\":true},{\"description\":\"ignores other non-objects\",\"data\":12,\"valid\":true},{\"description\":\"none of the properties mentioned\",\"data\":{},\"valid\":false},{\"description\":\"__proto__ present\",\"data\":{\"__proto__\":\"foo\"},\"valid\":false},{\"description\":\"toString present\",\"data\":{\"toString\":{\"length\":37}},\"valid\":false},{\"description\":\"constructor present\",\"data\":{\"constructor\":{\"length\":37}},\"valid\":false},{\"description\":\"all present\",\"data\":{\"__proto__\":12,\"constructor\":37,\"toString\":{\"length\":\"foo\"}},\"valid\":true}]}]");

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
