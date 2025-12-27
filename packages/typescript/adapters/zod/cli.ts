#!/usr/bin/env node
import { convert, type ConvertInput } from "./index.ts";

const chunks: string[] = [];
process.stdin.on("data", (chunk) => chunks.push(String(chunk)));
process.stdin.on("end", () => {
  try {
    const inputs: ConvertInput[] = JSON.parse(chunks.join(""));
    const outputs = inputs.map(({ name, schema }) => convert(name, schema));
    console.log(JSON.stringify(outputs));
  } catch (err) {
    console.error(err instanceof Error ? err.message : err);
    process.exit(1);
  }
});
