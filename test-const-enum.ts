#!/usr/bin/env bun

import { jsonSchemaToZodCode } from "./ts/packages/adapters/zod/src/converter";

// Test 1: const with object
const constObjectSchema = {
  const: { type: "admin", level: 5 }
};

console.log("=== Test 1: const with object ===");
const constObjCode = jsonSchemaToZodCode(constObjectSchema);
console.log("Generated:", constObjCode);

// Test 2: const with array
const constArraySchema = {
  const: [1, 2, 3]
};

console.log("\n=== Test 2: const with array ===");
const constArrCode = jsonSchemaToZodCode(constArraySchema);
console.log("Generated:", constArrCode);

// Test 3: heterogeneous enum with objects
const heterogeneousEnumSchema = {
  enum: ["admin", { type: "user", level: 1 }, 42]
};

console.log("\n=== Test 3: heterogeneous enum ===");
const heterogeneousEnumCode = jsonSchemaToZodCode(heterogeneousEnumSchema);
console.log("Generated:", heterogeneousEnumCode);

// Test 4: enum with only objects
const objectEnumSchema = {
  enum: [
    { type: "admin", level: 5 },
    { type: "user", level: 1 }
  ]
};

console.log("\n=== Test 4: enum with objects ===");
const objectEnumCode = jsonSchemaToZodCode(objectEnumSchema);
console.log("Generated:", objectEnumCode);

// Test 5: const with primitive (should still work)
const constPrimitiveSchema = {
  const: "admin"
};

console.log("\n=== Test 5: const with primitive ===");
const constPrimCode = jsonSchemaToZodCode(constPrimitiveSchema);
console.log("Generated:", constPrimCode);
