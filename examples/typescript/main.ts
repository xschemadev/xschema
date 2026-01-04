/**
 * XSchema Example - TypeScript
 * 
 * This example shows how to use xschema to:
 * 1. Define JSON Schema sources in config files (*.jsonc)
 * 2. Generate type-safe Zod validators
 * 3. Use the generated schemas with full type inference
 * 
 * Run `bun run generate` to regenerate schemas, then `bun run start` to run this file.
 */

import { schemas } from "./.xschema/xschema.gen";
import { createXSchemaClient, XSchemaType } from "@xschema/client";

// Create the xschema client with generated schemas
// defaultNamespace allows shorthand lookups: xschema("Calendar") instead of xschema("user:Calendar")
export const xschema = createXSchemaClient({ schemas, defaultNamespace: "user" });

// ============================================
// Type extraction using XSchemaType helper
// ============================================

// XSchemaType requires full "namespace:id" keys (no shorthand)
export type CalendarType = XSchemaType<"user:Calendar">;
export type ProfileType = XSchemaType<"user:Profile">;
export type TSConfigType = XSchemaType<"another:TSConfig">;

// ============================================
// Schema lookup - explicit namespace
// ============================================

// Use full "namespace:id" to get schemas from any namespace
const tsConfigSchema = xschema("another:TSConfig");
const calendarSchema = xschema("user:Calendar");

// ============================================
// Schema lookup - with default namespace
// ============================================

// When defaultNamespace is set, you can omit it for that namespace
const calendar = xschema("Calendar");  // Same as xschema("user:Calendar")
const profile = xschema("Profile");    // Same as xschema("user:Profile")

// ============================================
// Using the schemas - full Zod API available
// ============================================

// Parse data (throws on invalid)
const validCalendar = calendar.parse({
	dtstart: "2024-01-01",
	summary: "New Year's Day",
});

// Safe parse (returns result object)
const result = calendar.safeParse({ invalid: "data" });
if (!result.success) {
	console.log("Validation errors:", result.error.issues);
}

// Type inference works automatically
type InferredCalendar = typeof validCalendar;
//   ^? { startDate?: string, endDate?: string, summary: string, ... }

console.log("XSchema example running successfully!");
console.log("Available schemas:", Object.keys(schemas));
