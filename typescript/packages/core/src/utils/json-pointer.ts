/**
 * JSON Pointer utilities
 * Note: All $ref resolution is handled by the Go CLI bundler.
 * TypeScript receives pre-bundled schemas with all refs resolved.
 */

/**
 * Extract definition name from a $ref path
 * e.g., "#/$defs/User" -> "User"
 * 
 * @deprecated This function may be removed in future versions as refs
 * should be resolved by the Go bundler before reaching TypeScript.
 */
export function getRefName(ref: string): string | null {
	const match = ref.match(/#\/(?:\$defs|definitions)\/(.+)$/);
	return match ? match[1] : null;
}
