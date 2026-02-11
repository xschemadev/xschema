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
		id: "proto",
		varName: "test_proto",
		schema: jsonSchema,
	});
	if (!result.schema) throw new Error("No schema generated");
	return evalSchema(result.schema);
}

describe("prototype-property object validation", () => {
	// must use JSON.parse for __proto__ — JS object literals set prototype instead of creating a property
	const schema = makeSchema(
		JSON.parse(
			'{"type":"object","properties":{"__proto__":{"type":"number"},"constructor":{"type":"string"}},"required":["__proto__"]}',
		),
	);

	test("rejects string input", () => {
		expect(schema.safeParse("hello").success).toBe(false);
	});

	test("rejects number input", () => {
		expect(schema.safeParse(42).success).toBe(false);
	});

	test("rejects boolean input", () => {
		expect(schema.safeParse(true).success).toBe(false);
	});

	test("rejects null input", () => {
		expect(schema.safeParse(null).success).toBe(false);
	});

	test("rejects array input", () => {
		expect(schema.safeParse([1, 2]).success).toBe(false);
	});

	test("rejects missing required proto property", () => {
		expect(schema.safeParse({}).success).toBe(false);
	});

	test("accepts valid object with proto properties", () => {
		const obj = JSON.parse('{"__proto__": 42, "constructor": "foo"}');
		expect(schema.safeParse(obj).success).toBe(true);
	});

	test("accepts object with only required proto property", () => {
		const obj = JSON.parse('{"__proto__": 42}');
		expect(schema.safeParse(obj).success).toBe(true);
	});

	test("rejects invalid proto property value", () => {
		const obj = JSON.parse('{"__proto__": "not a number"}');
		expect(schema.safeParse(obj).success).toBe(false);
	});

	test("validates constructor property type when present", () => {
		const obj = JSON.parse('{"__proto__": 42, "constructor": 123}');
		expect(schema.safeParse(obj).success).toBe(false);
	});
});

describe("prototype-property with additionalProperties: false", () => {
	const schema = makeSchema({
		type: "object",
		properties: {
			toString: { type: "string" },
		},
		additionalProperties: false,
	});

	test("rejects primitive input", () => {
		expect(schema.safeParse("hello").success).toBe(false);
	});

	test("accepts valid object", () => {
		expect(schema.safeParse({ toString: "foo" }).success).toBe(true);
	});

	test("rejects extra properties", () => {
		expect(
			schema.safeParse({ toString: "foo", extra: 1 }).success,
		).toBe(false);
	});
});
