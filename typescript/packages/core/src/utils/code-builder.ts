/**
 * Code builder utilities for generating validation code
 * These helpers make code generation safer and more readable
 */

import { escapeString } from "./primitives.js";

/**
 * Chain method calls onto a base expression
 * Filters out null/undefined values
 */
export function chain(
	base: string,
	...methods: (string | null | undefined)[]
): string {
	return methods
		.filter((m): m is string => m != null && m !== "")
		.reduce((acc, m) => `${acc}.${m}`, base);
}

/**
 * Build a Zod union from an array of schema strings
 */
export function buildUnion(schemas: string[]): string {
	if (schemas.length === 0) return "z.never()";
	if (schemas.length === 1) return schemas[0]!;
	return `z.union([${schemas.join(", ")}])`;
}

/**
 * Build a Zod intersection from an array of schema strings
 */
export function buildIntersection(schemas: string[]): string {
	if (schemas.length === 0) return "z.any()";
	if (schemas.length === 1) return schemas[0]!;

	let result = schemas[0]!;
	for (let i = 1; i < schemas.length; i++) {
		result = `z.intersection(${result}, ${schemas[i]})`;
	}
	return result;
}

/**
 * Build a superRefine call
 */
export function buildSuperRefine(body: string): string {
	return `.superRefine((val, ctx) => {${body}})`;
}

/**
 * Build a refine call with a predicate and message
 */
export function buildRefine(predicate: string, message: string): string {
	return `.refine(${predicate}, { message: ${escapeString(message)} })`;
}

/**
 * Build a literal schema
 */
export function buildLiteral(value: unknown): string {
	return `z.literal(${JSON.stringify(value)})`;
}

/**
 * Build validation code that checks if a safeParse failed and adds issues
 */
export function buildSafeParseCheck(
	schemaCode: string,
	valueExpr: string = "val",
	pathPrefix: string[] = [],
): string {
	const pathArray =
		pathPrefix.length > 0
			? `[${pathPrefix.map((p) => escapeString(p)).join(", ")}, ...issue.path]`
			: "issue.path";

	return `{
      const result = ${schemaCode}.safeParse(${valueExpr});
      if (!result.success) {
        result.error.issues.forEach(issue => ctx.addIssue({ ...issue, path: ${pathArray} }));
      }
    }`;
}

/**
 * Build validation code that checks safeParse at a specific path
 */
export function buildPropertyCheck(
	schemaCode: string,
	propKey: string,
	required: boolean,
): string {
	const keyExpr = escapeString(propKey);
	const check = buildSafeParseCheck(schemaCode, `val[${keyExpr}]`, [propKey]);

	if (required) {
		return `
        if (Object.hasOwn(val, ${keyExpr})) ${check} else {
          ctx.addIssue({ code: z.ZodIssueCode.custom, path: [${keyExpr}], message: "Required" });
        }`;
	}
	return `
        if (Object.hasOwn(val, ${keyExpr})) ${check}`;
}

/**
 * Recursively normalize a JSON value so that object keys are sorted at every depth,
 * then stringify. Produces deterministic output regardless of key insertion order.
 */
export function sortedStringify(value: unknown): string {
	return JSON.stringify(deepSortKeys(value));
}

/**
 * Runtime code snippet for deep-sort-stringify.
 * Adapters embed this in generated refine/narrow callbacks for complex const/enum comparison.
 * Call as: `${DEEP_SORTED_STRINGIFY_RUNTIME}(val)` — returns a JSON string with recursively sorted keys.
 */
export const DEEP_SORTED_STRINGIFY_RUNTIME = `((v) => { const s = (v) => { if (v === null || typeof v !== 'object') return v; if (Array.isArray(v)) return v.map(s); const o = {}; for (const k of Object.keys(v).sort()) o[k] = s(v[k]); return o; }; return JSON.stringify(s(v)); })`;

function deepSortKeys(value: unknown): unknown {
	if (value === null || typeof value !== "object") {
		return value;
	}
	if (Array.isArray(value)) {
		return value.map(deepSortKeys);
	}
	const sorted: Record<string, unknown> = {};
	for (const key of Object.keys(value as Record<string, unknown>).sort()) {
		sorted[key] = deepSortKeys((value as Record<string, unknown>)[key]);
	}
	return sorted;
}
