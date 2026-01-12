/**
 * Tests for vocabulary-aware keyword skipping
 *
 * When vocabulary is disabled, validation keywords (minLength, maxLength, etc.)
 * should be omitted from the IR, while applicator keywords (properties, items, etc.)
 * should still work.
 */

import { describe, test, expect } from "bun:test";
import { parse, type ParseOptions } from "../parser/index.js";
import { VOCAB } from "./context.js";
import type {
	StringNode,
	NumberNode,
	ObjectNode,
	ArrayNode,
	SchemaNode,
} from "../ir/nodes.js";

/**
 * Creates vocabulary with validation disabled (both 2019 and 2020 drafts)
 */
function validationDisabled(): ParseOptions {
	return {
		vocabulary: {
			[VOCAB.VALIDATION_2020]: false,
			[VOCAB.VALIDATION_2019]: false,
		},
	};
}

/**
 * Creates vocabulary with format disabled
 */
function formatDisabled(): ParseOptions {
	return {
		vocabulary: {
			[VOCAB.FORMAT_ANNOTATION_2020]: false,
			[VOCAB.FORMAT_2019]: false,
		},
	};
}

/**
 * Creates vocabulary with only applicators enabled (validation and format disabled)
 */
function onlyApplicatorsEnabled(): ParseOptions {
	return {
		vocabulary: {
			[VOCAB.VALIDATION_2020]: false,
			[VOCAB.VALIDATION_2019]: false,
			[VOCAB.FORMAT_ANNOTATION_2020]: false,
			[VOCAB.FORMAT_2019]: false,
		},
	};
}

describe("vocabulary-aware keyword skipping", () => {
	describe("string validation keywords", () => {
		test("validation disabled: minLength, maxLength, pattern ignored", () => {
			const schema = {
				type: "string" as const,
				minLength: 5,
				maxLength: 10,
				pattern: "^[a-z]+$",
			};

			const result = parse(schema, validationDisabled()) as StringNode;

			expect(result.kind).toBe("string");
			expect(result.constraints.minLength).toBeUndefined();
			expect(result.constraints.maxLength).toBeUndefined();
			expect(result.constraints.pattern).toBeUndefined();
		});

		test("validation enabled (default): minLength, maxLength, pattern preserved", () => {
			const schema = {
				type: "string" as const,
				minLength: 5,
				maxLength: 10,
				pattern: "^[a-z]+$",
			};

			const result = parse(schema) as StringNode;

			expect(result.kind).toBe("string");
			expect(result.constraints.minLength).toBe(5);
			expect(result.constraints.maxLength).toBe(10);
			expect(result.constraints.pattern).toBe("^[a-z]+$");
		});

		test("format disabled: format ignored but type correct", () => {
			const schema = {
				type: "string" as const,
				format: "email",
			};

			const result = parse(schema, formatDisabled()) as StringNode;

			expect(result.kind).toBe("string");
			expect(result.format).toBeUndefined();
		});

		test("format enabled (default): format preserved", () => {
			const schema = {
				type: "string" as const,
				format: "email",
			};

			const result = parse(schema) as StringNode;

			expect(result.kind).toBe("string");
			expect(result.format).toBe("email");
		});
	});

	describe("number validation keywords", () => {
		test("validation disabled: minimum, maximum, multipleOf ignored", () => {
			const schema = {
				type: "number" as const,
				minimum: 0,
				maximum: 100,
				exclusiveMinimum: 0,
				exclusiveMaximum: 100,
				multipleOf: 5,
			};

			const result = parse(schema, validationDisabled()) as NumberNode;

			expect(result.kind).toBe("number");
			expect(result.constraints.minimum).toBeUndefined();
			expect(result.constraints.maximum).toBeUndefined();
			expect(result.constraints.exclusiveMinimum).toBeUndefined();
			expect(result.constraints.exclusiveMaximum).toBeUndefined();
			expect(result.constraints.multipleOf).toBeUndefined();
		});

		test("validation enabled (default): numeric constraints preserved", () => {
			const schema = {
				type: "number" as const,
				minimum: 0,
				maximum: 100,
				exclusiveMinimum: 0.5,
				exclusiveMaximum: 99.5,
				multipleOf: 5,
			};

			const result = parse(schema) as NumberNode;

			expect(result.kind).toBe("number");
			expect(result.constraints.minimum).toBe(0);
			expect(result.constraints.maximum).toBe(100);
			expect(result.constraints.exclusiveMinimum).toBe(0.5);
			expect(result.constraints.exclusiveMaximum).toBe(99.5);
			expect(result.constraints.multipleOf).toBe(5);
		});

		test("validation disabled: integer type still produces correct base type", () => {
			const schema = {
				type: "integer" as const,
				minimum: 0,
			};

			const result = parse(schema, validationDisabled()) as NumberNode;

			expect(result.kind).toBe("number");
			expect(result.integer).toBe(true);
			expect(result.constraints.minimum).toBeUndefined();
		});
	});

	describe("object validation keywords", () => {
		test("validation disabled: required, minProperties, maxProperties ignored", () => {
			const schema = {
				type: "object" as const,
				properties: {
					name: { type: "string" as const },
					age: { type: "number" as const },
				},
				required: ["name"],
				minProperties: 1,
				maxProperties: 5,
			};

			const result = parse(schema, validationDisabled()) as ObjectNode;

			expect(result.kind).toBe("object");
			// Properties are still parsed (applicator)
			expect(result.properties.size).toBe(2);
			// But required flag should be false since validation is disabled
			expect(result.properties.get("name")?.required).toBe(false);
			expect(result.properties.get("age")?.required).toBe(false);
			expect(result.minProperties).toBeUndefined();
			expect(result.maxProperties).toBeUndefined();
		});

		test("validation enabled (default): required, minProperties, maxProperties preserved", () => {
			const schema = {
				type: "object" as const,
				properties: {
					name: { type: "string" as const },
					age: { type: "number" as const },
				},
				required: ["name"],
				minProperties: 1,
				maxProperties: 5,
			};

			const result = parse(schema) as ObjectNode;

			expect(result.kind).toBe("object");
			expect(result.properties.get("name")?.required).toBe(true);
			expect(result.properties.get("age")?.required).toBe(false);
			expect(result.minProperties).toBe(1);
			expect(result.maxProperties).toBe(5);
		});

		test("validation disabled: dependentRequired ignored but dependentSchemas works", () => {
			const schema = {
				type: "object" as const,
				properties: {
					creditCard: { type: "string" as const },
					billingAddress: { type: "string" as const },
				},
				dependentRequired: {
					creditCard: ["billingAddress"],
				},
				dependentSchemas: {
					creditCard: {
						properties: {
							cvv: { type: "string" as const },
						},
					},
				},
			};

			const result = parse(schema, validationDisabled()) as ObjectNode;

			expect(result.kind).toBe("object");
			// dependentRequired is validation - should be missing
			expect(result.dependencies.has("creditCard")).toBe(true);
			// But dependentSchemas is applicator - should work
			const dep = result.dependencies.get("creditCard");
			expect(dep?.kind).toBe("schema");
		});
	});

	describe("array validation keywords", () => {
		test("validation disabled: minItems, maxItems, uniqueItems ignored", () => {
			const schema = {
				type: "array" as const,
				items: { type: "string" as const },
				minItems: 1,
				maxItems: 10,
				uniqueItems: true,
			};

			const result = parse(schema, validationDisabled()) as ArrayNode;

			expect(result.kind).toBe("array");
			// items is applicator - should work
			expect(result.items.kind).toBe("string");
			// Constraints are validation - should be undefined
			expect(result.constraints.minItems).toBeUndefined();
			expect(result.constraints.maxItems).toBeUndefined();
			expect(result.constraints.uniqueItems).toBeUndefined();
		});

		test("validation enabled (default): array constraints preserved", () => {
			const schema = {
				type: "array" as const,
				items: { type: "string" as const },
				minItems: 1,
				maxItems: 10,
				uniqueItems: true,
			};

			const result = parse(schema) as ArrayNode;

			expect(result.kind).toBe("array");
			expect(result.constraints.minItems).toBe(1);
			expect(result.constraints.maxItems).toBe(10);
			expect(result.constraints.uniqueItems).toBe(true);
		});

		test("validation disabled: contains applicator works, minContains/maxContains ignored", () => {
			const schema = {
				type: "array" as const,
				items: { type: "number" as const },
				contains: { minimum: 5 },
				minContains: 2,
				maxContains: 5,
			};

			const result = parse(schema, validationDisabled()) as ArrayNode;

			expect(result.kind).toBe("array");
			// contains is applicator - should work
			expect(result.constraints.contains).toBeDefined();
			// minContains/maxContains are validation - minContains defaults to 1, maxContains undefined
			expect(result.constraints.contains?.minContains).toBe(1);
			expect(result.constraints.contains?.maxContains).toBeUndefined();
		});
	});

	describe("applicator vocabulary (always enabled)", () => {
		test("properties, additionalProperties, patternProperties work regardless of validation", () => {
			const schema = {
				type: "object" as const,
				properties: {
					name: { type: "string" as const },
				},
				additionalProperties: { type: "number" as const },
				patternProperties: {
					"^x-": { type: "boolean" as const },
				},
			};

			const result = parse(schema, validationDisabled()) as ObjectNode;

			expect(result.kind).toBe("object");
			expect(result.properties.has("name")).toBe(true);
			expect((result.additionalProperties as SchemaNode).kind).toBe("number");
			expect(result.patternProperties.length).toBe(1);
			expect(result.patternProperties[0]?.pattern).toBe("^x-");
		});

		test("items, prefixItems work regardless of validation", () => {
			const schema = {
				type: "array" as const,
				prefixItems: [{ type: "string" as const }, { type: "number" as const }],
				items: { type: "boolean" as const },
			};

			const result = parse(schema, validationDisabled());

			expect(result.kind).toBe("tuple");
			if (result.kind === "tuple") {
				expect(result.prefixItems.length).toBe(2);
				expect(result.prefixItems[0]?.kind).toBe("string");
				expect(result.prefixItems[1]?.kind).toBe("number");
				expect((result.restItems as SchemaNode).kind).toBe("boolean");
			}
		});

		test("allOf, anyOf, oneOf work regardless of validation", () => {
			const schema = {
				allOf: [
					{ type: "object" as const },
					{ type: "object" as const },
				],
			};

			const result = parse(schema, validationDisabled());

			expect(result.kind).toBe("intersection");
			if (result.kind === "intersection") {
				expect(result.schemas.length).toBe(2);
			}
		});
	});

	describe("empty $vocabulary", () => {
		test("empty vocabulary object produces minimal IR with only structure", () => {
			const schema = {
				type: "string" as const,
				minLength: 5,
				maxLength: 10,
				pattern: "^[a-z]+$",
				format: "email",
			};

			// Empty vocabulary = all vocabularies explicitly set to false
			const emptyVocab: ParseOptions = {
				vocabulary: {
					[VOCAB.VALIDATION_2020]: false,
					[VOCAB.VALIDATION_2019]: false,
					[VOCAB.FORMAT_ANNOTATION_2020]: false,
					[VOCAB.FORMAT_2019]: false,
				},
			};

			const result = parse(schema, emptyVocab) as StringNode;

			expect(result.kind).toBe("string");
			// All validation and format constraints should be undefined
			expect(result.constraints.minLength).toBeUndefined();
			expect(result.constraints.maxLength).toBeUndefined();
			expect(result.constraints.pattern).toBeUndefined();
			expect(result.format).toBeUndefined();
		});
	});

	describe("partial vocabulary", () => {
		test("only specified vocabularies affect keyword processing", () => {
			const schema = {
				type: "string" as const,
				minLength: 5,
				format: "email",
			};

			// Only disable validation, keep format enabled
			const partialVocab: ParseOptions = {
				vocabulary: {
					[VOCAB.VALIDATION_2020]: false,
					[VOCAB.VALIDATION_2019]: false,
					// Format not disabled - should still work
				},
			};

			const result = parse(schema, partialVocab) as StringNode;

			expect(result.kind).toBe("string");
			// Validation disabled
			expect(result.constraints.minLength).toBeUndefined();
			// Format still enabled (not in disabled list)
			expect(result.format).toBe("email");
		});
	});

	describe("missing $vocabulary (undefined)", () => {
		test("undefined vocabulary: all keywords processed (default enabled)", () => {
			const schema = {
				type: "object" as const,
				properties: {
					name: { type: "string" as const, minLength: 1 },
				},
				required: ["name"],
				minProperties: 1,
			};

			// No vocabulary option = undefined = all enabled
			const result = parse(schema) as ObjectNode;

			expect(result.kind).toBe("object");
			expect(result.properties.get("name")?.required).toBe(true);
			expect(result.minProperties).toBe(1);

			const nameNode = result.properties.get("name")?.schema as StringNode;
			expect(nameNode.constraints.minLength).toBe(1);
		});
	});

	describe("type constraint with validation disabled", () => {
		test("type: string still produces StringNode with no constraints", () => {
			const schema = {
				type: "string" as const,
				minLength: 5,
			};

			const result = parse(schema, validationDisabled()) as StringNode;

			expect(result.kind).toBe("string");
			expect(result.constraints.minLength).toBeUndefined();
		});

		test("type: number still produces NumberNode with no constraints", () => {
			const schema = {
				type: "number" as const,
				minimum: 0,
			};

			const result = parse(schema, validationDisabled()) as NumberNode;

			expect(result.kind).toBe("number");
			expect(result.integer).toBe(false);
			expect(result.constraints.minimum).toBeUndefined();
		});
	});

	describe("typeless schemas with validation disabled", () => {
		test("typeless with only validation keywords becomes any when validation disabled", () => {
			const schema = {
				minLength: 5,
				maxLength: 10,
			};

			const result = parse(schema, validationDisabled());

			// With validation disabled, these keywords don't trigger type detection
			expect(result.kind).toBe("any");
		});

		test("typeless with applicator keywords still creates typeGuarded when validation disabled", () => {
			const schema = {
				properties: {
					name: { type: "string" as const },
				},
			};

			const result = parse(schema, validationDisabled());

			// properties is applicator - should still trigger object detection
			expect(result.kind).toBe("typeGuarded");
			if (result.kind === "typeGuarded") {
				expect(result.guards.length).toBe(1);
				expect(result.guards[0]?.check).toBe("object");
			}
		});
	});
});
