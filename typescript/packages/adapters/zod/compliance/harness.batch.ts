import { z, type ZodTypeAny } from "zod";

const schemas: Record<string, ZodTypeAny> = {{BATCH_SCHEMAS}};

const batchData: Array<{
  groupId: string;
  tests: Array<{ data: unknown; valid: boolean }>;
}> = JSON.parse({{BATCH_TEST_DATA}});

const results = batchData.flatMap(({ groupId, tests }) => {
  const schema = schemas[groupId];
  if (!schema) {
    return tests.map((tc, index) => ({
      groupId,
      index,
      expected: tc.valid,
      actual: "error" as const,
      error: `schema not found for group ${groupId}`,
    }));
  }

  return tests.map((tc, index) => {
    try {
      const result = schema.safeParse(tc.data);
      return {
        groupId,
        index,
        expected: tc.valid,
        actual: result.success ? ("true" as const) : ("false" as const),
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
