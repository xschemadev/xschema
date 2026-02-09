import type { ConvertResult, XSchemaAdapter } from "@xschemadev/core";
import { spawn } from "node:child_process";

export type { ConvertResult, XSchemaAdapter } from "@xschemadev/core";

export interface SchemaInput {
	namespace: string;
	id: string;
	schema: object | boolean;
}

export interface ConvertOptions {
	adapter: XSchemaAdapter;
	cwd?: string;
	executablePath?: string;
	allowFetch?: boolean;
}

export class ConvertError extends Error {
	readonly stderr: string;

	constructor(message: string, stderr: string) {
		super(message);
		this.name = "ConvertError";
		this.stderr = stderr;
	}
}

export async function convert(
	schemas: SchemaInput | SchemaInput[],
	options: ConvertOptions,
): Promise<ConvertResult[]> {
	const input = Array.isArray(schemas) ? schemas : [schemas];
	const bin = options.executablePath ?? "xschema";

	const args = ["convert", "--adapter", options.adapter.id];
	if (options.allowFetch) args.push("--allow-fetch");
	if (options.cwd) args.push("--project", options.cwd);

	return new Promise<ConvertResult[]>((resolve, reject) => {
		const child = spawn(bin, args, { stdio: ["pipe", "pipe", "pipe"] });

		let stdout = "";
		let stderr = "";

		child.stdout.on("data", (chunk: Buffer) => {
			stdout += chunk.toString();
		});

		child.stderr.on("data", (chunk: Buffer) => {
			stderr += chunk.toString();
		});

		child.on("error", (err) => {
			reject(
				new ConvertError(
					`failed to spawn xschema: ${err.message}`,
					err.message,
				),
			);
		});

		child.on("close", (code) => {
			if (code !== 0) {
				const msg = stderr.trim() || `xschema exited with code ${code}`;
				reject(new ConvertError(msg, stderr));
				return;
			}

			try {
				const results: ConvertResult[] = JSON.parse(stdout);
				resolve(results);
			} catch {
				reject(
					new ConvertError(
						`failed to parse xschema output: ${stdout}`,
						stderr,
					),
				);
			}
		});

		child.stdin.write(JSON.stringify(input));
		child.stdin.end();
	});
}
