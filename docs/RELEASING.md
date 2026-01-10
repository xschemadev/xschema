# Releasing xschema

This document explains how releases work for xschema.

## Overview

xschema uses [release-please](https://github.com/googleapis/release-please) for automated releases. When commits are merged to `master`, release-please:

1. Analyzes commit messages (conventional commits)
2. Creates/updates Release PRs with changelogs
3. When Release PRs are merged, triggers package publishing

## Components

| Component          | Package                  | Published To                       |
| ------------------ | ------------------------ | ---------------------------------- |
| CLI                | `@xschemadev/cli`        | npm + GitHub Releases              |
| Core               | `@xschemadev/core`       | npm                                |
| Client             | `@xschemadev/client`     | npm                                |
| Zod Adapter        | `@xschemadev/zod`        | npm                                |
| ArkType Adapter    | `@xschemadev/arktype`    | npm                                |
| Effect Adapter     | `@xschemadev/effect`     | npm                                |
| TypeBox Adapter    | `@xschemadev/typebox`    | npm                                |
| Valibot Adapter    | `@xschemadev/valibot`    | npm                                |
| TypeScript Adapter | `@xschemadev/typescript` | npm                                |

## How Versioning Works

Each component is versioned independently based on commits that touch its path:

| Path                                    | Component  |
| --------------------------------------- | ---------- |
| `cli/**`                                | CLI        |
| `typescript/packages/core/**`           | core       |
| `typescript/packages/client/**`         | client     |
| `typescript/packages/adapters/zod/**`   | zod        |
| `typescript/packages/adapters/arktype/**` | arktype  |
| `typescript/packages/adapters/effect/**`  | effect   |
| `typescript/packages/adapters/typebox/**` | typebox  |
| `typescript/packages/adapters/valibot/**` | valibot  |
| `typescript/packages/adapters/typescript/**` | typescript |

### Version Bump Rules

| Commit Type                      | Version Bump  |
| -------------------------------- | ------------- |
| `fix:`                           | Patch (0.0.X) |
| `feat:`                          | Minor (0.X.0) |
| `feat!:` or `BREAKING CHANGE:`   | Major (X.0.0) |

> Note: While version < 1.0.0, breaking changes bump minor instead of major (`bump-minor-pre-major`).

## Release Process

### 1. Make Changes

```bash
git checkout -b feat/my-feature
# ... make changes ...
git commit -m "feat(zod): add support for oneOf"
git push -u origin feat/my-feature
```

### 2. Create PR

- PR title should follow conventional commits
- CI runs tests and commitlint
- Get review and merge

### 3. Release PR Created

After merge to master, release-please automatically:
- Detects which packages changed based on commit scopes and paths
- Creates a Release PR with:
  - Version bumps in package.json
  - CHANGELOG.md updates
  - Release notes

Each component gets its own Release PR (configured via `separate-pull-requests: true`).

### 4. Merge Release PR

When you're ready to release:
1. Review the Release PR changelog
2. Merge it

### 5. Packages Published

After Release PR merge, GitHub Actions:
- **TypeScript packages**: Publishes to npm automatically
- **CLI**: goreleaser builds binaries and creates GitHub Release

## Tags

Release tags follow the pattern: `{component}-v{version}`

Examples: `cli-v1.2.3`, `zod-v0.2.1`, `effect-v0.1.0`

## Manual Release (Emergency)

If automation fails, you can publish manually:

### TypeScript Packages

```bash
cd typescript
bun install
bun run build

# Publish core first (dependency), then others
cd packages/core && bun publish --access public && cd ../..
cd packages/client && bun publish --access public && cd ../..
cd packages/adapters/zod && bun publish --access public && cd ../../..
# repeat for other adapters as needed
```

### CLI

```bash
cd cli
goreleaser release --clean
```

## Secrets Required

| Secret         | Purpose                    | Where to set          |
| -------------- | -------------------------- | --------------------- |
| `GITHUB_TOKEN` | release-please, goreleaser | Automatic             |
| `NPM_TOKEN`    | npm publish                | Repository secrets    |

### Setting up NPM_TOKEN

1. Go to [npmjs.com](https://www.npmjs.com/) → Access Tokens
2. Generate a new "Automation" token
3. Add to GitHub repo: Settings → Secrets → Actions → New repository secret
4. Name: `NPM_TOKEN`, Value: your token

## Troubleshooting

### Release PR not created

- Check that commits follow conventional commit format
- Ensure commits touched a monitored path (cli/, typescript/packages/*)
- Check GitHub Actions logs for release-please errors

### npm publish fails

- Verify NPM_TOKEN is set and valid
- Check you have publish permissions on @xschemadev org
- Ensure package.json version matches what release-please bumped

### goreleaser fails

- Check Go version in workflow matches go.mod
- Verify .goreleaser.yaml syntax
- Check GitHub token has release permissions

## CLI npm Distribution

The CLI is distributed via npm using platform-specific packages:

| Package                        | Platform        |
| ------------------------------ | --------------- |
| `@xschemadev/cli`              | Main (shim)     |
| `@xschemadev/cli-darwin-arm64` | macOS ARM64     |
| `@xschemadev/cli-darwin-x64`   | macOS Intel     |
| `@xschemadev/cli-linux-arm64`  | Linux ARM64     |
| `@xschemadev/cli-linux-x64`    | Linux x64       |
| `@xschemadev/cli-win32-arm64`  | Windows ARM64   |
| `@xschemadev/cli-win32-x64`    | Windows x64     |

Users install via:
```bash
npm install @xschemadev/cli
# or
bun add @xschemadev/cli
```

The main package uses `optionalDependencies` - npm/bun automatically installs only the package matching the user's platform.

## Configuration Files

| File                                   | Purpose                         |
| -------------------------------------- | ------------------------------- |
| `release-please-config.json`           | Package paths and release types |
| `.release-please-manifest.json`        | Current versions                |
| `.github/workflows/release-please.yml` | Release automation              |
| `cli/npm/`                             | CLI npm package templates       |

## Adding a New Adapter

When creating a new adapter, update these 3 files:

### 1. `release-please-config.json`

Add to `packages`:

```json
"typescript/packages/adapters/{name}": {
  "release-type": "node",
  "component": "{name}"
}
```

### 2. `.release-please-manifest.json`

Add with initial version:

```json
"typescript/packages/adapters/{name}": "0.1.0"
```

### 3. `.github/workflows/release-please.yml`

**a) Add output** (in `release-please` job outputs):

```yaml
{name}--release_created: ${{ steps.release.outputs['typescript/packages/adapters/{name}--release_created'] }}
```

**b) Add to condition** (in `release-ts` job `if`):

```yaml
needs.release-please.outputs.{name}--release_created == 'true' ||
```

**c) Add publish step** (in `release-ts` job steps):

```yaml
- name: Publish @xschemadev/{name}
  if: needs.release-please.outputs.{name}--release_created == 'true'
  run: bun publish --access public
  working-directory: typescript/packages/adapters/{name}
  env:
    NPM_CONFIG_TOKEN: ${{ secrets.NPM_TOKEN }}
```
