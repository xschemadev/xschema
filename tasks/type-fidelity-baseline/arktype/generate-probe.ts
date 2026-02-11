/**
 * Generates the type probe fixture by running the arktype adapter's convert function
 * against representative JSON schemas for each targeted construct.
 *
 * Run: bun run tasks/type-fidelity-baseline/arktype/generate-probe.ts
 */
import { convert } from "../../../typescript/packages/adapters/arktype/src/index.js";

interface ProbeCase {
	name: string;
	description: string;
	schema: unknown;
}

const probeCases: ProbeCase[] = [
	// --- oneOf ---
	{
		name: "oneOfStringOrNumber",
		description: "oneOf with string | number - should ideally be string | number, currently unknown",
		schema: {
			oneOf: [{ type: "string" }, { type: "number" }],
		},
	},
	{
		name: "oneOfObjects",
		description: "oneOf with discriminated objects - currently unknown",
		schema: {
			oneOf: [
				{
					type: "object",
					properties: { kind: { const: "a" }, value: { type: "string" } },
					required: ["kind", "value"],
				},
				{
					type: "object",
					properties: { kind: { const: "b" }, value: { type: "number" } },
					required: ["kind", "value"],
				},
			],
		},
	},

	// --- not ---
	{
		name: "notString",
		description: "not string - should be unknown, currently unknown",
		schema: {
			not: { type: "string" },
		},
	},
	{
		name: "notBoolean",
		description: "not boolean - should be unknown, currently unknown",
		schema: {
			not: { type: "boolean" },
		},
	},

	// --- conditional (if/then/else) ---
	{
		name: "conditionalIfThenElse",
		description: "if/then/else - should be unknown, currently unknown",
		schema: {
			if: { type: "object", properties: { kind: { const: "a" } }, required: ["kind"] },
			then: { type: "object", properties: { kind: { const: "a" }, value: { type: "string" } }, required: ["kind", "value"] },
			else: { type: "object", properties: { kind: { const: "b" }, value: { type: "number" } }, required: ["kind", "value"] },
		},
	},
	{
		name: "conditionalIfThen",
		description: "if/then only - should be unknown, currently unknown",
		schema: {
			if: { type: "string", minLength: 1 },
			then: { type: "string", maxLength: 10 },
		},
	},

	// --- typeGuarded ---
	{
		name: "typeGuardedObject",
		description: "type-guarded object properties without explicit type - currently unknown",
		schema: {
			properties: { name: { type: "string" } },
			required: ["name"],
		},
	},
	{
		name: "typeGuardedArrayMinItems",
		description: "type-guarded array constraint without explicit type - currently unknown",
		schema: {
			minItems: 1,
		},
	},

	// --- tuple ---
	{
		name: "tupleStringNumber",
		description: "tuple [string, number] - should be [string, number, ...unknown[]], currently unknown[]",
		schema: {
			type: "array",
			prefixItems: [{ type: "string" }, { type: "number" }],
		},
	},
	{
		name: "tupleStringNumberClosed",
		description: "closed tuple [string, number] - should be [string, number], currently unknown[]",
		schema: {
			type: "array",
			prefixItems: [{ type: "string" }, { type: "number" }],
			items: false,
		},
	},
	{
		name: "tupleMixedWithRest",
		description: "tuple [string, number, ...boolean[]] - currently unknown[]",
		schema: {
			type: "array",
			prefixItems: [{ type: "string" }, { type: "number" }],
			items: { type: "boolean" },
		},
	},

	// --- complex const ---
	{
		name: "constObject",
		description: "const object - should be narrow type, currently unknown",
		schema: {
			const: { name: "alice", age: 30 },
		},
	},
	{
		name: "constArray",
		description: "const array - should be narrow type, currently unknown",
		schema: {
			const: [1, 2, 3],
		},
	},

	// --- complex enum ---
	{
		name: "enumWithObjects",
		description: "enum with object values - currently unknown",
		schema: {
			enum: [{ a: 1 }, { b: 2 }, "simple"],
		},
	},
	{
		name: "enumWithArrays",
		description: "enum with array values - currently unknown",
		schema: {
			enum: [[1, 2], [3, 4], null],
		},
	},
];

// Generate the probe fixture
let output = `/**
 * ArkType Type Probe Fixture - AUTO-GENERATED
 *
 * This file contains generated ArkType code for representative schemas of each
 * targeted construct. The inferred types reveal where type.unknown degrades the
 * TypeScript type information.
 *
 * Generated: ${new Date().toISOString().split("T")[0]}
 * Do NOT edit manually - regenerate with: bun run tasks/type-fidelity-baseline/arktype/generate-probe.ts
 */
import { type } from "arktype";

`;

for (const probe of probeCases) {
	const result = convert({
		namespace: "probe",
		id: probe.name,
		varName: `probe_${probe.name}`,
		schema: probe.schema as object,
	});

	output += `// ${probe.description}\n`;
	output += `const probe_${probe.name} = ${result.schema};\n`;
	output += `type Probe_${probe.name} = ${result.type};\n`;
	output += `\n`;
}

// Write the exports for type inspection
output += `// Export all probe types for external inspection\n`;
output += `export type {\n`;
for (const probe of probeCases) {
	output += `  Probe_${probe.name},\n`;
}
output += `};\n`;

const outPath = new URL("./probe-fixture.ts", import.meta.url).pathname;
await Bun.write(outPath, output);
console.log(`Wrote probe fixture to: ${outPath}`);
console.log(`Probe cases: ${probeCases.length}`);
