#!/usr/bin/env node
const { execFileSync } = require("child_process");
const { platform, arch } = process;

const PLATFORMS = {
  darwin: {
    arm64: "@xschemadev/cli-darwin-arm64",
    x64: "@xschemadev/cli-darwin-x64",
  },
  linux: {
    arm64: "@xschemadev/cli-linux-arm64",
    x64: "@xschemadev/cli-linux-x64",
  },
  win32: {
    arm64: "@xschemadev/cli-win32-arm64",
    x64: "@xschemadev/cli-win32-x64",
  },
};

const pkg = PLATFORMS[platform]?.[arch];
if (!pkg) {
  console.error(`xschema: unsupported platform ${platform}-${arch}`);
  console.error("Please open an issue at https://github.com/xschemadev/xschema/issues");
  process.exit(1);
}

const bin = platform === "win32" ? "xschema.exe" : "xschema";

let binPath;
try {
  binPath = require.resolve(`${pkg}/${bin}`);
} catch (e) {
  console.error(`xschema: could not find binary for ${platform}-${arch}`);
  console.error(`Expected package: ${pkg}`);
  console.error("");
  console.error("This usually means npm/bun was run with --ignore-optional.");
  console.error("Try reinstalling: npm install @xschemadev/cli");
  process.exit(1);
}

try {
  execFileSync(binPath, process.argv.slice(2), { stdio: "inherit" });
} catch (e) {
  if (e.status !== undefined) {
    process.exit(e.status);
  }
  throw e;
}
