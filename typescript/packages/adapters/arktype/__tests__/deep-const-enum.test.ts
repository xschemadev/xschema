import { describe, test, expect } from "bun:test";
import { type } from "arktype";
import { convert } from "../src/index.js";

function evalSchema(schemaCode: string) {
	const fn = new Function("type", `return ${schemaCode}`);
	return fn(type);
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

function validates(schema: ReturnType<typeof makeSchema>, data: unknown): boolean {
	return !(schema(data) instanceof type.errors);
}

describe("const with nested objects", () => {
	const schema = makeSchema({
		const: { outer: { z: 1, a: 2 }, b: "hello" },
	});

	test("accepts matching value regardless of key order", () => {
		expect(validates(schema, { b: "hello", outer: { a: 2, z: 1 } })).toBe(true);
	});

	test("accepts matching value with original key order", () => {
		expect(validates(schema, { outer: { z: 1, a: 2 }, b: "hello" })).toBe(true);
	});

	test("rejects value with different nested value", () => {
		expect(validates(schema, { outer: { z: 999, a: 2 }, b: "hello" })).toBe(false);
	});

	test("rejects value with missing nested key", () => {
		expect(validates(schema, { outer: { z: 1 }, b: "hello" })).toBe(false);
	});
});

describe("const with nested arrays containing objects", () => {
	const schema = makeSchema({
		const: [{ z: 1, a: 2 }, { y: 3, b: 4 }],
	});

	test("accepts matching array with reordered inner keys", () => {
		expect(validates(schema, [{ a: 2, z: 1 }, { b: 4, y: 3 }])).toBe(true);
	});

	test("rejects array with wrong element order", () => {
		expect(validates(schema, [{ y: 3, b: 4 }, { z: 1, a: 2 }])).toBe(false);
	});

	test("rejects array with extra element", () => {
		expect(validates(schema, [{ z: 1, a: 2 }, { y: 3, b: 4 }, "extra"])).toBe(false);
	});
});

describe("const with deeply nested objects (3+ levels)", () => {
	const schema = makeSchema({
		const: { level1: { level2: { c: 3, a: 1, b: 2 } } },
	});

	test("accepts matching value with reordered keys at all levels", () => {
		expect(validates(schema, { level1: { level2: { a: 1, b: 2, c: 3 } } })).toBe(true);
	});

	test("rejects value with wrong deeply nested value", () => {
		expect(validates(schema, { level1: { level2: { a: 1, b: 2, c: 999 } } })).toBe(false);
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
		expect(validates(schema, { name: "alice", meta: { a: 2, z: 1 } })).toBe(true);
	});

	test("accepts second enum value with reordered nested keys", () => {
		expect(validates(schema, { name: "bob", meta: { b: 4, y: 3 } })).toBe(true);
	});

	test("accepts primitive enum value", () => {
		expect(validates(schema, "simple")).toBe(true);
	});

	test("rejects non-matching value", () => {
		expect(validates(schema, { name: "charlie" })).toBe(false);
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
		expect(validates(schema, [{ a: 2, z: 1 }])).toBe(true);
	});

	test("accepts second enum value with reordered inner keys", () => {
		expect(validates(schema, [{ b: 4, y: 3 }])).toBe(true);
	});

	test("rejects non-matching array", () => {
		expect(validates(schema, [{ x: 5 }])).toBe(false);
	});
});

describe("primitive const/enum unchanged", () => {
	test("const string works", () => {
		const schema = makeSchema({ const: "hello" });
		expect(validates(schema, "hello")).toBe(true);
		expect(validates(schema, "world")).toBe(false);
	});

	test("const number works", () => {
		const schema = makeSchema({ const: 42 });
		expect(validates(schema, 42)).toBe(true);
		expect(validates(schema, 43)).toBe(false);
	});

	test("const null works", () => {
		const schema = makeSchema({ const: null });
		expect(validates(schema, null)).toBe(true);
		expect(validates(schema, 0)).toBe(false);
	});

	test("const boolean works", () => {
		const schema = makeSchema({ const: true });
		expect(validates(schema, true)).toBe(true);
		expect(validates(schema, false)).toBe(false);
	});

	test("enum all primitives works", () => {
		const schema = makeSchema({ enum: [1, "two", true, null] });
		expect(validates(schema, 1)).toBe(true);
		expect(validates(schema, "two")).toBe(true);
		expect(validates(schema, true)).toBe(true);
		expect(validates(schema, null)).toBe(true);
		expect(validates(schema, false)).toBe(false);
	});

	test("const flat object works", () => {
		const schema = makeSchema({ const: { a: 1, b: 2 } });
		expect(validates(schema, { a: 1, b: 2 })).toBe(true);
		expect(validates(schema, { b: 2, a: 1 })).toBe(true);
		expect(validates(schema, { a: 1 })).toBe(false);
	});
});
