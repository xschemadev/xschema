/**
 * Extracts inferred types from the probe fixture using TypeScript compiler API.
 * Run: bun typescript/packages/adapters/arktype/type-probe/extract-types.ts
 */
import ts from "typescript";
import path from "path";

const configPath = path.resolve(import.meta.dirname, "tsconfig.json");
const configFile = ts.readConfigFile(configPath, ts.sys.readFile);
const parsed = ts.parseJsonConfigFileContent(configFile.config, ts.sys, path.dirname(configPath));

const program = ts.createProgram(parsed.fileNames, parsed.options);
const checker = program.getTypeChecker();

const sourceFile = program.getSourceFile(path.resolve(import.meta.dirname, "probe-fixture.ts"));
if (!sourceFile) {
	console.error("Could not find probe-fixture.ts");
	process.exit(1);
}

const results: { name: string; type: string; hasUnknown: boolean }[] = [];

ts.forEachChild(sourceFile, (node) => {
	if (ts.isTypeAliasDeclaration(node) && node.name.text.startsWith("Probe_")) {
		const type = checker.getTypeAtLocation(node);
		const typeStr = checker.typeToString(type, undefined, ts.TypeFormatFlags.NoTruncation);
		const hasUnknown = typeStr === "unknown" || typeStr.includes("unknown[]") || typeStr.includes("unknown,");
		results.push({
			name: node.name.text,
			type: typeStr,
			hasUnknown,
		});
	}
});

console.log("# ArkType Type Probe Results");
console.log("");
console.log("| Probe | Inferred Type | Has unknown? |");
console.log("| ----- | ------------- | ------------ |");
for (const r of results) {
	console.log(`| ${r.name} | \`${r.type}\` | ${r.hasUnknown ? "YES" : "no"} |`);
}

const unknownCount = results.filter((r) => r.hasUnknown).length;
console.log("");
console.log(`Total probes: ${results.length}`);
console.log(`Probes with unknown: ${unknownCount}`);
console.log(`Probes without unknown: ${results.length - unknownCount}`);
