/**
 * Primitive utility functions
 */

/**
 * Check if a value is a primitive (string, number, boolean, null)
 */
export function isPrimitive(value: unknown): boolean {
	return (
		value === null ||
		typeof value === "string" ||
		typeof value === "number" ||
		typeof value === "boolean"
	);
}

/**
 * Escape a string for use in generated code (JSON.stringify)
 */
export function escapeString(str: string): string {
	return JSON.stringify(str);
}

/**
 * Check if an object is empty
 */
export function isEmptyObject(obj: object): boolean {
	return Object.keys(obj).length === 0;
}

/**
 * Get own property safely (handles __proto__ etc.)
 */
export function getOwnProperty<T>(
	obj: Record<string, T>,
	key: string,
): T | undefined {
	if (Object.prototype.hasOwnProperty.call(obj, key)) {
		return obj[key];
	}
	return undefined;
}

/**
 * Property names that exist on Object.prototype and need special handling
 */
export const PROTOTYPE_PROPERTY_NAMES = new Set([
	"__proto__",
	"constructor",
	"toString",
	"valueOf",
	"hasOwnProperty",
	"isPrototypeOf",
	"propertyIsEnumerable",
	"toLocaleString",
	"__defineGetter__",
	"__defineSetter__",
	"__lookupGetter__",
	"__lookupSetter__",
]);

/**
 * Check if any keys are prototype property names
 */
export function hasPrototypeProperties(keys: string[]): boolean {
	return keys.some((k) => PROTOTYPE_PROPERTY_NAMES.has(k));
}
