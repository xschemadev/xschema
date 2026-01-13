
import { Schema as S } from "effect"

const schemas: Record<string, { isTypeOnly: boolean; validate: (data: unknown) => boolean }> = {
  "group_0": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
        if (typeof val === "string") {
          const result = S.decodeUnknownEither(S.String.pipe(S.filter((val) => [...new Intl.Segmenter().segment(val)].length <= 2, { message: () => "String must have at most 2 character(s)" })))(val);
          if (result._tag === "Left") return false;
        }
        return true;
      }, { message: () => "Type-guarded validation failed" }));
    return { isTypeOnly: false, validate: (data: unknown) => S.decodeUnknownEither(schema)(data)._tag === "Right" };
  })(),
  "group_1": (() => {
    const schema = S.Unknown.pipe(S.filter((val) => {
        if (typeof val === "string") {
          const result = S.decodeUnknownEither(S.String.pipe(S.filter((val) => [...new Intl.Segmenter().segment(val)].length <= 2, { message: () => "String must have at most 2 character(s)" })))(val);
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
}> = JSON.parse("[{\"groupId\":\"group_0\",\"tests\":[{\"description\":\"shorter is valid\",\"data\":\"f\",\"valid\":true},{\"description\":\"exact length is valid\",\"data\":\"fo\",\"valid\":true},{\"description\":\"too long is invalid\",\"data\":\"foo\",\"valid\":false},{\"description\":\"ignores non-strings\",\"data\":100,\"valid\":true},{\"description\":\"two graphemes is long enough\",\"data\":\"💩💩\",\"valid\":true}]},{\"groupId\":\"group_1\",\"tests\":[{\"description\":\"shorter is valid\",\"data\":\"f\",\"valid\":true},{\"description\":\"too long is invalid\",\"data\":\"foo\",\"valid\":false}]}]");

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
