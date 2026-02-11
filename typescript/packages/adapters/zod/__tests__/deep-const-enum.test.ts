import { describe, test, expect } from "bun:test";
import { z } from "zod";
import { convert } from "../src/index.js";

function evalSchema(schemaCode: string) {
	const fn = new Function("z", `return ${schemaCode}`);
	return fn(z);
}

function makeSchema(jsonSchema: Record<string, unknown>) {
	const result = convert({
		namespace: "test",
		id: "constEnum",
		varName: "test_constEnum",
		schema: jsonSchema,
	});
	if (!result.schema) throw new Error("No schema generated");
	return evalSchema(result.schema);
}

describe("const with nested objects", () => {
	const schema = makeSchema({
		const: { outer: { z: 1, a: 2 }, b: "hello" },
	});

	test("accepts matching value regardless of key order", () => {
		expect(schema.safeParse({ b: "hello", outer: { a: 2, z: 1 } }).success).toBe(true);
	});

	test("accepts matching value with original key order", () => {
		expect(schema.safeParse({ outer: { z: 1, a: 2 }, b: "hello" }).success).toBe(true);
	});

	test("rejects value with different nested value", () => {
		expect(schema.safeParse({ outer: { z: 999, a: 2 }, b: "hello" }).success).toBe(false);
	});

	test("rejects value with missing nested key", () => {
		expect(schema.safeParse({ outer: { z: 1 }, b: "hello" }).success).toBe(false);
	});
});

describe("const with nested arrays containing objects", () => {
	const schema = makeSchema({
		const: [{ z: 1, a: 2 }, { y: 3, b: 4 }],
	});

	test("accepts matching array with reordered inner keys", () => {
		expect(schema.safeParse([{ a: 2, z: 1 }, { b: 4, y: 3 }]).success).toBe(true);
	});

	test("rejects array with wrong element order", () => {
		expect(schema.safeParse([{ y: 3, b: 4 }, { z: 1, a: 2 }]).success).toBe(false);
	});

	test("rejects array with extra element", () => {
		expect(schema.safeParse([{ z: 1, a: 2 }, { y: 3, b: 4 }, "extra"]).success).toBe(false);
	});
});

describe("const with deeply nested objects (3+ levels)", () => {
	const schema = makeSchema({
		const: { level1: { level2: { c: 3, a: 1, b: 2 } } },
	});

	test("accepts matching value with reordered keys at all levels", () => {
		expect(schema.safeParse({ level1: { level2: { a: 1, b: 2, c: 3 } } }).success).toBe(true);
	});

	test("rejects value with wrong deeply nested value", () => {
		expect(schema.safeParse({ level1: { level2: { a: 1, b: 2, c: 999 } } }).success).toBe(false);
	});
});

describe("enum with nested objects", () => {
	const schema = makeSchema({
		enum: [
			{ name: "alice", meta: { z: 1, a: 2 } },
			{ name: "bob", meta: { y: 3, b: 4 } },
			"simple",
		],
	});

	test("accepts first enum value with reordered nested keys", () => {
		expect(schema.safeParse({ name: "alice", meta: { a: 2, z: 1 } }).success).toBe(true);
	});

	test("accepts second enum value with reordered nested keys", () => {
		expect(schema.safeParse({ name: "bob", meta: { b: 4, y: 3 } }).success).toBe(true);
	});

	test("accepts primitive enum value", () => {
		expect(schema.safeParse("simple").success).toBe(true);
	});

	test("rejects non-matching value", () => {
		expect(schema.safeParse({ name: "charlie" }).success).toBe(false);
	});
});

describe("enum with arrays containing nested objects", () => {
	const schema = makeSchema({
		enum: [
			[{ z: 1, a: 2 }],
			[{ y: 3, b: 4 }],
		],
	});

	test("accepts first enum value with reordered inner keys", () => {
		expect(schema.safeParse([{ a: 2, z: 1 }]).success).toBe(true);
	});

	test("accepts second enum value with reordered inner keys", () => {
		expect(schema.safeParse([{ b: 4, y: 3 }]).success).toBe(true);
	});

	test("rejects non-matching array", () => {
		expect(schema.safeParse([{ x: 5 }]).success).toBe(false);
	});
});

describe("primitive const/enum unchanged", () => {
	test("const string works", () => {
		const schema = makeSchema({ const: "hello" });
		expect(schema.safeParse("hello").success).toBe(true);
		expect(schema.safeParse("world").success).toBe(false);
	});

	test("const number works", () => {
		const schema = makeSchema({ const: 42 });
		expect(schema.safeParse(42).success).toBe(true);
		expect(schema.safeParse(43).success).toBe(false);
	});

	test("const null works", () => {
		const schema = makeSchema({ const: null });
		expect(schema.safeParse(null).success).toBe(true);
		expect(schema.safeParse(0).success).toBe(false);
	});

	test("const boolean works", () => {
		const schema = makeSchema({ const: true });
		expect(schema.safeParse(true).success).toBe(true);
		expect(schema.safeParse(false).success).toBe(false);
	});

	test("enum all strings works", () => {
		const schema = makeSchema({ enum: ["a", "b", "c"] });
		expect(schema.safeParse("a").success).toBe(true);
		expect(schema.safeParse("d").success).toBe(false);
	});

	test("enum all primitives works", () => {
		const schema = makeSchema({ enum: [1, "two", true, null] });
		expect(schema.safeParse(1).success).toBe(true);
		expect(schema.safeParse("two").success).toBe(true);
		expect(schema.safeParse(true).success).toBe(true);
		expect(schema.safeParse(null).success).toBe(true);
		expect(schema.safeParse(false).success).toBe(false);
	});

	test("const flat object works", () => {
		const schema = makeSchema({ const: { a: 1, b: 2 } });
		expect(schema.safeParse({ a: 1, b: 2 }).success).toBe(true);
		expect(schema.safeParse({ b: 2, a: 1 }).success).toBe(true);
		expect(schema.safeParse({ a: 1 }).success).toBe(false);
	});
});
