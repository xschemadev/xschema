#!/usr/bin/env bun

import { z } from "zod";

// Test const with object - property order independence
const constObjSchema = z.object({}).passthrough().refine((val) => JSON.stringify(val, Object.keys(val as object).sort()) === "{\"level\":5,\"type\":\"admin\"}", { message: "Value must equal the const value" });

console.log("=== Test const with object ===");
console.log("Same object (same order):", constObjSchema.safeParse({ type: "admin", level: 5 }).success); // should be true
console.log("Same object (different order):", constObjSchema.safeParse({ level: 5, type: "admin" }).success); // should be true
console.log("Different object:", constObjSchema.safeParse({ type: "user", level: 1 }).success); // should be false

// Test const with array
const constArrSchema = z.array(z.any()).refine((val) => JSON.stringify(val, Object.keys(val as object).sort()) === "[1,2,3]", { message: "Value must equal the const value" });

console.log("\n=== Test const with array ===");
console.log("Same array:", constArrSchema.safeParse([1, 2, 3]).success); // should be true
console.log("Different array:", constArrSchema.safeParse([3, 2, 1]).success); // should be false

// Test heterogeneous enum
const heterogeneousEnumSchema = z.any().refine((val) => {
  const normalized = JSON.stringify(val, Object.keys(val as object).sort());
  const validValues = ["\"admin\"", "{\"level\":1,\"type\":\"user\"}", "42"];
  return validValues.includes(normalized);
}, { message: "Value must be one of the enum values" });

console.log("\n=== Test heterogeneous enum ===");
console.log("String value:", heterogeneousEnumSchema.safeParse("admin").success); // should be true
console.log("Number value:", heterogeneousEnumSchema.safeParse(42).success); // should be true
console.log("Object value:", heterogeneousEnumSchema.safeParse({ type: "user", level: 1 }).success); // should be true
console.log("Object value (different order):", heterogeneousEnumSchema.safeParse({ level: 1, type: "user" }).success); // should be true
console.log("Invalid value:", heterogeneousEnumSchema.safeParse("invalid").success); // should be false

// Test enum with objects
const objectEnumSchema = z.any().refine((val) => {
  const normalized = JSON.stringify(val, Object.keys(val as object).sort());
  const validValues = ["{\"level\":5,\"type\":\"admin\"}", "{\"level\":1,\"type\":\"user\"}"];
  return validValues.includes(normalized);
}, { message: "Value must be one of the enum values" });

console.log("\n=== Test enum with objects ===");
console.log("First object:", objectEnumSchema.safeParse({ type: "admin", level: 5 }).success); // should be true
console.log("Second object:", objectEnumSchema.safeParse({ type: "user", level: 1 }).success); // should be true
console.log("Invalid object:", objectEnumSchema.safeParse({ type: "guest", level: 0 }).success); // should be false
