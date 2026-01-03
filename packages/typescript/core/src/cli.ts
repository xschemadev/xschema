import type { ConvertInput, ConvertResult } from "./types";

/**
 * Creates a CLI handler for xschema adapters.
 * Reads JSON array of ConvertInput from stdin, calls convert for each, outputs JSON array of ConvertResult.
 *
 * @example
 * ```ts
 * #!/usr/bin/env node
 * import { createAdapterCLI } from "@xschema/core";
 * import { convert } from "./index";
 *
 * createAdapterCLI(convert);
 * ```
 */
export function createAdapterCLI(
  convert: (input: ConvertInput) => ConvertResult
): void {
  const chunks: string[] = [];
  process.stdin.on("data", (chunk) => chunks.push(String(chunk)));
  process.stdin.on("end", () => {
    try {
      const inputs: ConvertInput[] = JSON.parse(chunks.join(""));
      const outputs = inputs.map(convert);
      console.log(JSON.stringify(outputs));
    } catch (err) {
      console.error(err instanceof Error ? err.message : err);
      process.exit(1);
    }
  });
}
