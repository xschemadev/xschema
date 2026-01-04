# Publishing Go CLI Binaries to npm

This document explains how to distribute the xschema Go CLI via npm, following the patterns used by esbuild, turbo, and biome.

## Overview

The pattern involves publishing:
1. **Platform-specific packages** - contain the actual binary for each OS/arch combination
2. **Main wrapper package** - references platform packages as `optionalDependencies` and includes a bin script to execute the correct one

```
@xschema/cli              # wrapper - what users install
@xschema/cli-darwin-arm64 # macOS ARM64 binary
@xschema/cli-darwin-x64   # macOS Intel binary
@xschema/cli-linux-x64    # Linux x64 binary
@xschema/cli-linux-arm64  # Linux ARM64 binary
@xschema/cli-win32-x64    # Windows x64 binary
@xschema/cli-win32-arm64  # Windows ARM64 binary
```

---

## 1. Package Structure

### Main Wrapper Package (`@xschema/cli`)

```
npm/@xschema/cli/
  bin/
    xschema           # JavaScript entrypoint script
  package.json
  README.md
```

**package.json:**

```json
{
  "name": "@xschema/cli",
  "version": "1.0.0",
  "description": "JSON Schema to native validators with full type safety",
  "repository": {
    "type": "git",
    "url": "git+https://github.com/xschemadev/xschema.git"
  },
  "license": "MIT",
  "bin": {
    "xschema": "bin/xschema"
  },
  "files": ["bin/xschema", "README.md"],
  "engines": {
    "node": ">=18"
  },
  "optionalDependencies": {
    "@xschema/cli-darwin-arm64": "1.0.0",
    "@xschema/cli-darwin-x64": "1.0.0",
    "@xschema/cli-linux-x64": "1.0.0",
    "@xschema/cli-linux-arm64": "1.0.0",
    "@xschema/cli-win32-x64": "1.0.0",
    "@xschema/cli-win32-arm64": "1.0.0"
  }
}
```

### Platform Package (e.g., `@xschema/cli-darwin-arm64`)

```
npm/@xschema/cli-darwin-arm64/
  xschema            # the actual binary
  package.json
```

**package.json:**

```json
{
  "name": "@xschema/cli-darwin-arm64",
  "version": "1.0.0",
  "description": "xschema CLI binary for macOS ARM64",
  "repository": {
    "type": "git",
    "url": "git+https://github.com/xschemadev/xschema.git"
  },
  "license": "MIT",
  "engines": {
    "node": ">=18"
  },
  "os": ["darwin"],
  "cpu": ["arm64"]
}
```

The `os` and `cpu` fields tell npm to only install this package on matching platforms.

---

## 2. Binary Selection Script

The main package needs a script to locate and execute the correct platform binary.

**bin/xschema:**

```javascript
#!/usr/bin/env node
const { platform, arch } = process;
const { execFileSync } = require('child_process');
const path = require('path');

const PLATFORMS = {
  win32: {
    x64: '@xschema/cli-win32-x64/xschema.exe',
    arm64: '@xschema/cli-win32-arm64/xschema.exe',
  },
  darwin: {
    x64: '@xschema/cli-darwin-x64/xschema',
    arm64: '@xschema/cli-darwin-arm64/xschema',
  },
  linux: {
    x64: '@xschema/cli-linux-x64/xschema',
    arm64: '@xschema/cli-linux-arm64/xschema',
  },
};

function getBinaryPath() {
  // Allow override via environment variable
  if (process.env.XSCHEMA_BINARY_PATH) {
    return process.env.XSCHEMA_BINARY_PATH;
  }

  const platformBinaries = PLATFORMS[platform];
  if (!platformBinaries) {
    console.error(`Unsupported platform: ${platform}`);
    process.exit(1);
  }

  const binaryPath = platformBinaries[arch];
  if (!binaryPath) {
    console.error(`Unsupported architecture: ${arch} on ${platform}`);
    process.exit(1);
  }

  try {
    return require.resolve(binaryPath);
  } catch (e) {
    console.error(`Could not find xschema binary for ${platform}-${arch}`);
    console.error('This usually happens when installing with --no-optional flag.');
    console.error(`Expected package: @xschema/cli-${platform}-${arch === 'x64' ? 'x64' : arch}`);
    process.exit(1);
  }
}

try {
  execFileSync(getBinaryPath(), process.argv.slice(2), { stdio: 'inherit' });
} catch (e) {
  if (e.status) process.exit(e.status);
  throw e;
}
```

---

## 3. GoReleaser Integration

### Option A: GoReleaser Pro NPM Publisher (Recommended if using Pro)

GoReleaser Pro v2.8+ has built-in npm support:

```yaml
# cli/.goreleaser.yaml
npms:
  - name: "@xschema/cli"
    ids:
      - xschema
    description: "JSON Schema to native validators with full type safety"
    license: MIT
    homepage: "https://xschema.dev"
    repository: "https://github.com/xschemadev/xschema"
    keywords:
      - json-schema
      - typescript
      - zod
      - validators
    access: public
```

This auto-generates packages and handles publishing. However, this downloads binaries at install time (not the optionalDependencies pattern).

### Option B: Separate Workflow Step (Recommended for OSS GoReleaser)

For the true optionalDependencies pattern (like esbuild/biome), you need a separate workflow step after goreleaser builds the binaries.

**Updated `.goreleaser.yaml`:**

```yaml
version: 2

project_name: xschema

before:
  hooks:
    - go mod tidy

builds:
  - main: .
    binary: xschema
    env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
      - windows
    goarch:
      - amd64
      - arm64
    ldflags:
      - -s -w
      - -X github.com/xschemadev/xschema/cmd.version={{.Version}}

archives:
  - format: binary
    name_template: "xschema-{{ .Os }}-{{ .Arch }}"

checksum:
  name_template: "checksums.txt"

release:
  prerelease: auto
  name_template: "CLI v{{.Version}}"
```

Key change: `format: binary` outputs raw binaries without archiving.

---

## 4. GitHub Actions Workflow

### Complete Release Workflow

```yaml
# .github/workflows/release-please.yml
name: Release

on:
  push:
    branches:
      - master

permissions:
  contents: write
  pull-requests: write

jobs:
  release-please:
    runs-on: ubuntu-latest
    outputs:
      releases_created: ${{ steps.release.outputs.releases_created }}
      cli--release_created: ${{ steps.release.outputs['cli--release_created'] }}
      cli--tag_name: ${{ steps.release.outputs['cli--tag_name'] }}
      cli--version: ${{ steps.release.outputs['cli--version'] }}
    steps:
      - uses: googleapis/release-please-action@v4
        id: release
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
          config-file: release-please-config.json
          manifest-file: .release-please-manifest.json

  # Build binaries for all platforms
  build-binaries:
    needs: release-please
    if: ${{ needs.release-please.outputs.cli--release_created == 'true' }}
    strategy:
      matrix:
        include:
          - os: ubuntu-latest
            goos: linux
            goarch: amd64
            npm_platform: linux-x64
          - os: ubuntu-latest
            goos: linux
            goarch: arm64
            npm_platform: linux-arm64
          - os: macos-latest
            goos: darwin
            goarch: amd64
            npm_platform: darwin-x64
          - os: macos-latest
            goos: darwin
            goarch: arm64
            npm_platform: darwin-arm64
          - os: windows-latest
            goos: windows
            goarch: amd64
            npm_platform: win32-x64
          - os: windows-latest
            goos: windows
            goarch: arm64
            npm_platform: win32-arm64
    runs-on: ${{ matrix.os }}
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: "1.22"

      - name: Build binary
        working-directory: cli
        env:
          GOOS: ${{ matrix.goos }}
          GOARCH: ${{ matrix.goarch }}
          CGO_ENABLED: 0
        run: |
          VERSION=${{ needs.release-please.outputs.cli--version }}
          EXT=""
          if [ "${{ matrix.goos }}" = "windows" ]; then EXT=".exe"; fi
          go build -ldflags="-s -w -X github.com/xschemadev/xschema/cmd.version=${VERSION}" \
            -o xschema${EXT} .
        shell: bash

      - name: Upload binary artifact
        uses: actions/upload-artifact@v4
        with:
          name: binary-${{ matrix.npm_platform }}
          path: cli/xschema*

  # Run goreleaser for GitHub release
  release-github:
    needs: [release-please, build-binaries]
    if: ${{ needs.release-please.outputs.cli--release_created == 'true' }}
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: "1.22"

      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v6
        with:
          distribution: goreleaser
          version: "~> v2"
          args: release --clean
          workdir: cli
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

  # Publish to npm
  publish-npm:
    needs: [release-please, build-binaries]
    if: ${{ needs.release-please.outputs.cli--release_created == 'true' }}
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-node@v4
        with:
          node-version: "20"
          registry-url: "https://registry.npmjs.org"

      - name: Download all binary artifacts
        uses: actions/download-artifact@v4
        with:
          path: binaries

      - name: Generate and publish npm packages
        env:
          NODE_AUTH_TOKEN: ${{ secrets.NPM_TOKEN }}
          VERSION: ${{ needs.release-please.outputs.cli--version }}
        run: |
          # Create npm packages directory
          mkdir -p npm-packages

          # Platform packages
          declare -A PLATFORMS=(
            ["linux-x64"]="linux x64"
            ["linux-arm64"]="linux arm64"
            ["darwin-x64"]="darwin x64"
            ["darwin-arm64"]="darwin arm64"
            ["win32-x64"]="win32 x64"
            ["win32-arm64"]="win32 arm64"
          )

          for platform in "${!PLATFORMS[@]}"; do
            read -r os cpu <<< "${PLATFORMS[$platform]}"
            pkg_dir="npm-packages/@xschema/cli-${platform}"
            mkdir -p "$pkg_dir"
            
            # Find and copy binary
            if [ "$os" = "win32" ]; then
              cp "binaries/binary-${platform}/xschema.exe" "$pkg_dir/"
            else
              cp "binaries/binary-${platform}/xschema" "$pkg_dir/"
              chmod +x "$pkg_dir/xschema"
            fi
            
            # Create package.json
            cat > "$pkg_dir/package.json" << EOF
          {
            "name": "@xschema/cli-${platform}",
            "version": "${VERSION}",
            "description": "xschema CLI binary for ${os} ${cpu}",
            "repository": {
              "type": "git",
              "url": "git+https://github.com/xschemadev/xschema.git"
            },
            "license": "MIT",
            "os": ["${os}"],
            "cpu": ["${cpu}"]
          }
          EOF
            
            # Publish platform package
            cd "$pkg_dir"
            npm publish --access public
            cd -
          done

          # Main wrapper package
          wrapper_dir="npm-packages/@xschema/cli"
          mkdir -p "$wrapper_dir/bin"
          
          # Copy bin script (from your repo's npm/ directory)
          cp npm/@xschema/cli/bin/xschema "$wrapper_dir/bin/"
          chmod +x "$wrapper_dir/bin/xschema"
          
          # Create wrapper package.json
          cat > "$wrapper_dir/package.json" << EOF
          {
            "name": "@xschema/cli",
            "version": "${VERSION}",
            "description": "JSON Schema to native validators with full type safety",
            "repository": {
              "type": "git",
              "url": "git+https://github.com/xschemadev/xschema.git"
            },
            "license": "MIT",
            "bin": {
              "xschema": "bin/xschema"
            },
            "engines": {
              "node": ">=18"
            },
            "optionalDependencies": {
              "@xschema/cli-darwin-arm64": "${VERSION}",
              "@xschema/cli-darwin-x64": "${VERSION}",
              "@xschema/cli-linux-x64": "${VERSION}",
              "@xschema/cli-linux-arm64": "${VERSION}",
              "@xschema/cli-win32-x64": "${VERSION}",
              "@xschema/cli-win32-arm64": "${VERSION}"
            }
          }
          EOF
          
          # Copy README
          cp README.md "$wrapper_dir/"
          
          # Publish wrapper package
          cd "$wrapper_dir"
          npm publish --access public
```

---

## 5. Required npm Setup

### Create npm Organization

1. Go to [npmjs.com](https://www.npmjs.com)
2. Sign in or create an account
3. Click your profile > Add Organization
4. Create `@xschema` organization (or use existing)

### Create Access Token

1. Go to npmjs.com > Account Settings > Access Tokens
2. Click "Generate New Token" > "Granular Access Token"
3. Name: `github-actions-publish`
4. Expiration: set appropriate expiry
5. Packages and scopes: Select packages or "Read and write"
6. Organizations: Select `@xschema`
7. Copy the token

### Add Token to GitHub Secrets

1. Go to your GitHub repo > Settings > Secrets and variables > Actions
2. Click "New repository secret"
3. Name: `NPM_TOKEN`
4. Value: paste the token

---

## 6. Integration with release-please

The workflow above integrates with release-please via outputs:

```yaml
outputs:
  cli--release_created: ${{ steps.release.outputs['cli--release_created'] }}
  cli--version: ${{ steps.release.outputs['cli--version'] }}
```

This triggers npm publish only when a CLI release is created.

### release-please-config.json

Ensure your config includes the CLI:

```json
{
  "packages": {
    "cli": {
      "release-type": "go",
      "component": "cli",
      "extra-files": ["version.txt"]
    }
  }
}
```

---

## 7. Directory Structure

Add these files to your repo:

```
npm/
  @xschema/
    cli/
      bin/
        xschema       # JavaScript bin script
      package.json    # Template (version gets replaced)
      README.md
    cli-darwin-arm64/
      package.json    # Template
    cli-darwin-x64/
      package.json
    cli-linux-x64/
      package.json
    cli-linux-arm64/
      package.json
    cli-win32-x64/
      package.json
    cli-win32-arm64/
      package.json
```

Or generate them dynamically in CI (as shown in the workflow).

---

## 8. Alternative: Simpler "xschema" Package Name

If you prefer `xschema` instead of `@xschema/cli`:

```json
{
  "name": "xschema",
  "optionalDependencies": {
    "xschema-darwin-arm64": "1.0.0",
    "xschema-darwin-x64": "1.0.0",
    ...
  }
}
```

This is simpler but requires owning the `xschema` package name on npm.

---

## 9. Comparison with Reference Projects

| Project | Pattern | Binary Resolution |
|---------|---------|-------------------|
| **esbuild** | `@esbuild/<platform>` | postinstall script downloads/validates |
| **turbo** | `turbo-<platform>-<arch>` | bin script with `require.resolve` |
| **biome** | `@biomejs/cli-<platform>` | bin script with `require.resolve` |
| **rollup** | `@rollup/rollup-<platform>` | optionalDependencies |

All use `optionalDependencies` with `os` and `cpu` fields for platform filtering.

---

## 10. Testing Locally

Before publishing:

```bash
# Build for your platform
cd cli && go build -o xschema .

# Create test package structure
mkdir -p /tmp/xschema-test/@xschema/cli-linux-x64
cp xschema /tmp/xschema-test/@xschema/cli-linux-x64/

# Create package.json files and test
cd /tmp/xschema-test/@xschema/cli
npm pack  # Creates tarball without publishing

# Install locally to test
npm install ./xschema-cli-1.0.0.tgz
npx xschema --version
```

---

## Summary

1. **Create npm org** `@xschema`
2. **Add npm packages** to repo under `npm/`
3. **Update workflow** to build binaries per-platform and publish to npm
4. **Add `NPM_TOKEN` secret** to GitHub
5. **Release** via release-please triggers full pipeline

The key insight is that npm's `optionalDependencies` combined with `os`/`cpu` fields automatically installs only the matching platform package, and the wrapper's bin script uses `require.resolve` to find and execute it.
