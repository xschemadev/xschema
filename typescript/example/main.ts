/**
 * XSchema Example - TypeScript
 *
 * This example shows how to use xschema to:
 * 1. Define JSON Schema sources in config files (*.jsonc)
 * 2. Generate type-safe validators (Zod, ArkType, etc.)
 * 3. Use the generated schemas with full type inference
 *
 * Run `bun run generate` to regenerate schemas, then `bun run start` to run this file.
 */

import { type } from "arktype";

import { schemas } from "./.xschema/xschema.gen";
import { createXSchemaClient, XSchemaType } from "@xschemadev/client";

// Create the xschema client with generated schemas
// defaultNamespace allows shorthand lookups: xschema("Calendar") instead of xschema("user:Calendar")
export const xschema = createXSchemaClient({
  schemas,
  defaultNamespace: "user",
});

// ============================================
// Type extraction using XSchemaType helper
// ============================================

// "namespace:id"
export type CalendarType = XSchemaType<"user:Calendar">;
export type TSConfigType = XSchemaType<"another:TSConfig">;
export type UserType = XSchemaType<"user:User">;

// Native ts converter
export type TSConfigNativeTS = XSchemaType<"another:TSConfigNative">;

// ============================================
// Zod schema
// ============================================

// Use full "namespace:id" to get schemas from any namespace
const tsConfigSchema = xschema("another:TSConfig");
const calendarSchema = xschema("user:Calendar");

// When defaultNamespace is set, you can omit it for that namespace
const calendar = xschema("Calendar"); // Same as xschema("user:Calendar")
const validCalendar = calendar.parse({
  dtstart: "2024-01-01",
  summary: "New Year's Day",
});
const result = calendar.safeParse({ invalid: "data" });
type InferredCalendar = typeof validCalendar;

// ============================================
// ArkType schema
// ============================================

// The User schema uses ArkType adapter (see user.jsonc)
const user = xschema("User");
const validUser = user({
  id: "123",
  email: "alice@example.com",
  name: "Alice",
  age: 30,
});
