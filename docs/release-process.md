# Release Process

xschema uses **release-please** for automated versioning and releases in a monorepo setup. This ensures consistent versioning, changelog generation, and multi-platform distribution.

## Overview

**Release-please** is a tool that automates releases by:

1. Analyzing commit messages (Conventional Commits)
2. Creating/updating pull requests with version bumps and changelogs
3. Merging the PR → automatically creates GitHub releases
4. Triggering build and publish workflows

**Trigger:** Push commits to `master` branch

**Workflow:** `.github/workflows/release-please.yml`

## Release Configuration

### Main Configuration

**File:** `release-please-config.json`

```json
{
  "$schema": "https://raw.githubusercontent.com/googleapis/release-please/main/schemas/config.json",
  "release-type": "node",
  "bump-minor-pre-major": true,
  "bump-patch-for-minor-pre-major": true,
  "packages": {
    "cli": {
      "release-type": "simple",
      "component": "cli",
      "extra-files": [
        "cli/cmd/root.go",
        {
          "type": "json",
          "path": "cli/npm/cli/package.json",
          "jsonpath": "$.version"
        }
      ]
    },
    "typescript/packages/core": {
      "release-type": "node",
      "component": "core"
    },
    "typescript/packages/client": {
      "release-type": "node",
      "component": "client"
    },
    "typescript/packages/zod": {
      "release-type": "node",
      "component": "zod"
    }
  },
  "separate-pull-requests": true,
  "tag-separator": "-",
  "include-component-in-tag": true,
  "include-v-in-tag": true
}
```

### Version Manifest

**File:** `.release-please-manifest.json`

Tracks current versions for each package:

```json
{
  "cli": "0.0.3",
  "typescript/packages/core": "0.0.2",
  "typescript/packages/client": "0.0.2",
  "typescript/packages/zod": "0.0.1"
}
```

This file is automatically updated by release-please.

## Release Flow

### 1. Write Commits

Write commits using **Conventional Commits** format:

```bash
# Feature (minor bump)
git commit -m "feat: add new adapter support"

# Bug fix (patch bump)
git commit -m "fix: handle edge case in parser"

# Breaking change (major bump)
git commit -m "feat!: change adapter protocol
BREAKING CHANGE: old format no longer supported"
```

**Conventional Commit Types:**

| Type | Semver | Example |
|------|--------|---------|
| `feat` | Minor | "feat: add new feature" |
| `fix` | Patch | "fix: resolve bug" |
| `feat!` or `BREAKING CHANGE` | Major | "feat!: breaking change" |
| `docs` | None | "docs: update README" |
| `style` | None | "style: format code" |
| `refactor` | None | "refactor: improve code" |
| `perf` | Patch | "perf: optimize" |
| `test` | None | "test: add tests" |
| `chore` | None | "chore: update deps" |

### 2. Push to Master

```bash
git push origin master
```

### 3. Release-Please Creates PR(s)

The `release-please` action runs automatically and creates one PR per package that needs releasing:

**Example PR created:**
- Title: `chore(release): release v0.0.4` (for CLI)
- Body includes:
  - Version bump: 0.0.3 → 0.0.4
  - Changelog with all commits
  - List of files changed

**Multiple PRs for multi-package releases:**
```
PR 1: chore(release): @xschemadev/core: release v0.0.3
PR 2: chore(release): @xschemadev/client: release v0.0.3
PR 3: chore(release): @xschemadev/zod: release v0.0.2
PR 4: chore(release): cli: release v0.0.4
```

### 4. Review and Merge PR

Review the generated PR to verify:
- ✅ Version bumps are correct
- ✅ Changelog is accurate
- ✅ Files being updated are expected

Then merge the PR.

### 5. GitHub Release & Tags Created

When the PR is merged, release-please:

1. Creates **Git tags** for each released component:
   - `cli-v0.0.4`
   - `core-v0.0.3`
   - `client-v0.0.3`
   - `zod-v0.0.2`

2. Creates **GitHub releases** for each tag

3. Triggers **build and publish workflows** (see next section)

## Build and Publish Workflows

The GitHub Actions workflow (`.github/workflows/release-please.yml`) has multiple jobs:

### 1. Release-Please Job

**Runs:** On every push to master

**Creates:** Release PRs and tags

**Outputs:**
- `releases_created`: true if any releases were created
- `cli--release_created`: true if CLI was released
- `core--release_created`: true if core was released
- etc. (one for each component)

### 2. CLI Release Job

**Runs:** Only if `cli--release_created == 'true'`

**Steps:**

#### a. Build Binaries for All Platforms

Builds Go binary for 6 target platforms:

```bash
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/xschema-linux-x64/xschema
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o dist/xschema-linux-arm64/xschema

# macOS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/xschema-darwin-x64/xschema
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o dist/xschema-darwin-arm64/xschema

# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/xschema-win32-x64/xschema.exe
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o dist/xschema-win32-arm64/xschema.exe
```

**Build Flags:**
```
-ldflags "-s -w -X github.com/xschemadev/xschema/cmd.version=0.0.4 -X github.com/xschemadev/xschema/cmd.commit=abc1234 -X github.com/xschemadev/xschema/cmd.date=2024-01-04T..."
```

Sets version info at compile time.

#### b. Create Archives

```bash
# Compress each platform directory
tar -czf xschema-linux-x64.tar.gz -C dist/xschema-linux-x64 .
tar -czf xschema-darwin-x64.tar.gz -C dist/xschema-darwin-x64 .
# ... etc for all platforms
```

#### c. Upload to GitHub Release

```bash
gh release upload cli-v0.0.4 *.tar.gz --clobber
```

Assets uploaded:
- `xschema-linux-x64.tar.gz`
- `xschema-linux-arm64.tar.gz`
- `xschema-darwin-x64.tar.gz`
- `xschema-darwin-arm64.tar.gz`
- `xschema-win32-x64.tar.gz`
- `xschema-win32-arm64.tar.gz`

#### d. Distribute to npm Platform Packages

Copy binaries to npm package directories:

```bash
# Structure for npm
cli/npm/
├── cli/                          # Main package wrapper
├── cli-linux-x64/                # Platform packages
│   ├── package.json
│   └── xschema                   # Actual binary
├── cli-linux-arm64/
├── cli-darwin-x64/
├── cli-darwin-arm64/
├── cli-win32-x64/
└── cli-win32-arm64/
```

#### e. Update Package Versions

Updates `package.json` files with new version:

```bash
# Update platform packages
jq --arg v "0.0.4" '.version = $v' cli/npm/cli-linux-x64/package.json

# Update main package and optionalDependencies
jq --arg v "0.0.4" '
  .version = $v |
  .optionalDependencies = (.optionalDependencies | with_entries(.value = $v))
' cli/npm/cli/package.json
```

#### f. Publish Platform Packages to npm

```bash
# Publish each platform package first
npm publish cli-linux-x64/     # @xschemadev/cli-linux-x64@0.0.4
npm publish cli-linux-arm64/
npm publish cli-darwin-x64/
npm publish cli-darwin-arm64/
npm publish cli-win32-x64/
npm publish cli-win32-arm64/

# Then publish main package that depends on them
npm publish cli/                # @xschemadev/cli@0.0.4
```

The main `@xschemadev/cli` package has `optionalDependencies` on platform packages. npm will automatically select and install the correct binary for the user's platform.

### 3. TypeScript Release Job

**Runs:** If any TypeScript package was released

**Steps:**

#### a. Install and Build

```bash
cd typescript
bun install
bun run build
```

#### b. Publish Each Package

For each released package:

```bash
npm publish packages/core/      # @xschemadev/core@0.0.3
npm publish packages/client/    # @xschemadev/client@0.0.3
npm publish packages/zod/       # @xschemadev/zod@0.0.2
```

## Published Packages

### npm Packages

**1. @xschemadev/core**
- **Type:** TypeScript library
- **Version:** 0.0.2
- **Purpose:** Adapter protocol types and CLI helper
- **Exports:** `ConvertInput`, `ConvertResult`, `createAdapterCLI`
- **Used by:** Adapter authors
- **Depends on:** @types/node (devDep)

**2. @xschemadev/client**
- **Type:** TypeScript library
- **Version:** 0.0.2
- **Purpose:** Runtime client with TypeScript type inference
- **Exports:** `createXSchemaClient`, type helpers
- **Used by:** End users in their projects
- **No dependencies**

**3. @xschemadev/zod**
- **Type:** TypeScript adapter
- **Version:** 0.0.1
- **Purpose:** Converts JSON Schema to Zod validators
- **Binary:** `xschema-zod`
- **Used by:** End users via CLI (installed automatically)
- **Depends on:** @xschemadev/core, json-schema-to-zod

**4. @xschemadev/cli** (Wrapper)
- **Type:** npm package (wrapper)
- **Version:** 0.0.4
- **Purpose:** Main CLI package entry point
- **Installs:** Platform-specific binary via optional dependencies
- **Optional Dependencies:**
  - `@xschemadev/cli-linux-x64`
  - `@xschemadev/cli-linux-arm64`
  - `@xschemadev/cli-darwin-x64`
  - `@xschemadev/cli-darwin-arm64`
  - `@xschemadev/cli-win32-x64`
  - `@xschemadev/cli-win32-arm64`

**5. @xschemadev/cli-{platform}** (Platform Packages)
- **Type:** Binary distribution
- **Versions:** All match main CLI (0.0.4)
- **Contains:** Pre-built Go binary for specific platform
- **One per platform:** linux-x64, linux-arm64, darwin-x64, darwin-arm64, win32-x64, win32-arm64

### GitHub Releases

**Location:** https://github.com/xschemadev/xschema/releases

**Content for each tag:**
- Release notes (auto-generated from commits)
- Tarball assets for all 6 platforms
  - `xschema-linux-x64.tar.gz`
  - `xschema-linux-arm64.tar.gz`
  - `xschema-darwin-x64.tar.gz`
  - `xschema-darwin-arm64.tar.gz`
  - `xschema-win32-x64.tar.gz`
  - `xschema-win32-arm64.tar.gz`

**Example tag:**
```
cli-v0.0.4
├── Release notes
├── Assets
│   ├── xschema-linux-x64.tar.gz
│   ├── xschema-linux-arm64.tar.gz
│   ├── xschema-darwin-x64.tar.gz
│   ├── xschema-darwin-arm64.tar.gz
│   ├── xschema-win32-x64.tar.gz
│   └── xschema-win32-arm64.tar.gz
```

## Installation

### From npm

```bash
# Install CLI with automatic platform detection
npm install -g @xschemadev/cli

# Or in your project
npm install @xschemadev/cli @xschemadev/client @xschemadev/zod
```

When installed, npm automatically selects the correct binary for the user's platform.

### From GitHub Releases

For platform-specific binaries:

```bash
# Download tarball for your platform
curl -L https://github.com/xschemadev/xschema/releases/download/cli-v0.0.4/xschema-linux-x64.tar.gz
tar -xzf xschema-linux-x64.tar.gz
./xschema generate
```

## How to Trigger a Release

### Quick Summary

1. Write commits with Conventional Commits format
2. Push to `master`
3. Review release-please PR
4. Merge PR
5. Workflows automatically build and publish

### Detailed Steps

**Step 1: Create commits**

```bash
# Feature for CLI
git commit -m "feat(cli): support multiple languages"

# Bug fix for adapter
git commit -m "fix: handle null values in schema conversion"

# Breaking change
git commit -m "feat!: change adapter protocol

BREAKING CHANGE: old JSON format no longer supported"
```

**Step 2: Push to master**

```bash
git push origin master
```

**Step 3: Check GitHub Actions**

The `release-please` workflow will run and create PRs within 1-2 minutes.

Visit: https://github.com/xschemadev/xschema/pulls?q=is%3Apr+is%3Aopen+release-please

**Step 4: Review the PR**

Example PR:
```
chore(release): release v0.0.4

CHANGELOG:

### Features
- **cli**: support multiple languages ([commit abc123](link))

### Bug Fixes
- handle null values in schema conversion ([commit def456](link))

### Files Updated
- cli/cmd/root.go (version)
- cli/npm/cli/package.json (version)
- CHANGELOG.md (changelog)
```

Verify:
- ✅ Version bumps are correct
- ✅ Changelog entries are accurate
- ✅ Only expected files are modified

**Step 5: Merge the PR**

Use GitHub's merge button or command line:

```bash
# Option 1: Use GitHub web interface
# Click "Merge pull request"

# Option 2: Use gh CLI
gh pr merge --auto
```

**Step 6: Workflows Run Automatically**

After merge, GitHub Actions:
1. Creates Git tags: `cli-v0.0.4`, `core-v0.0.3`, etc.
2. Builds Go binaries for all platforms
3. Uploads to GitHub releases
4. Publishes to npm

This takes about 5-10 minutes.

**Step 7: Verify Release**

Check:
- ✅ GitHub releases page has new entries
- ✅ npm has new versions: `npm view @xschemadev/cli`

## Version Management

### Independent Versioning

Each package has its own version:
- `cli` and `typescript/packages/core` can be at different versions
- Example: core@0.0.2, zod@0.0.1, cli@0.0.3

### Version Bumping Rules

Released-please analyzes commits and bumps versions:

```
Old version → New version (based on commits)

0.0.3 → 0.0.4  (if feat: or fix:)
0.0.3 → 0.1.0  (if feat! or BREAKING CHANGE:)
0.0.3 → (no bump)  (if only docs:, chore:, etc.)
```

### Pre-Major Version

Special rules for versions before 1.0.0:

```json
{
  "bump-minor-pre-major": true,
  "bump-patch-for-minor-pre-major": true
}
```

This means:
- `feat:` bumps 0.0.3 → 0.1.0 (not 0.0.4)
- `feat!:` bumps 0.1.0 → 1.0.0 (not 0.2.0)

### Extra Files

The CLI version is synced to npm wrapper package:

```json
{
  "packages": {
    "cli": {
      "extra-files": [
        "cli/cmd/root.go",
        {
          "type": "json",
          "path": "cli/npm/cli/package.json",
          "jsonpath": "$.version"
        }
      ]
    }
  }
}
```

When CLI is released to 0.0.4:
1. `cli/cmd/root.go` version variable updated
2. `cli/npm/cli/package.json` .version field updated
3. All platform package versions also updated

## Troubleshooting

### No Release PR Created

**Issue:** release-please didn't create a PR

**Causes:**
- No conventional commits since last release
- Commits only have `docs:`, `chore:`, or `style:`

**Solution:**
Write a proper `feat:` or `fix:` commit

### Wrong Version Bumped

**Issue:** Version bump is incorrect (expected 0.0.4, got 0.1.0)

**Cause:** Pre-major version bumping rules (see above)

**Solution:**
- If you want patch bump but got minor: Wait until 1.0.0
- If you want minor bump but got major: You used `feat!:` or BREAKING CHANGE

### Package Failed to Publish

**Issue:** Workflow failed at "Publish to npm"

**Causes:**
- npm token expired or invalid
- Package already exists at that version

**Solution:**
- Check GitHub Actions logs for error
- Verify npm credentials in GitHub secrets
- If version already published, check manifest and force bump

## Related Files

- `.github/workflows/release-please.yml` - Full workflow definition
- `release-please-config.json` - Release configuration
- `.release-please-manifest.json` - Current versions
- `.commitlintrc.js` - Validates commit format
