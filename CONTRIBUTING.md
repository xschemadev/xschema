# Contributing to xschema

Thank you for your interest in contributing to xschema!

## Getting Started

### Prerequisites

- [Go](https://golang.org/) 1.22+
- [Bun](https://bun.sh/) 1.0+
- [Git](https://git-scm.com/)

### Setup

```bash
# Clone the repository
git clone https://github.com/xschema/xschema.git
cd xschema

# Install root dependencies (husky, commitlint)
bun install

# Install TypeScript package dependencies
cd packages/typescript && bun install && cd ../..

# Build TypeScript packages
cd packages/typescript && bun run build && cd ../..

# Run Go tests
cd cli && go test ./... && cd ..
```

## Development Workflow

### Making Changes

1. Create a branch for your changes:
   ```bash
   git checkout -b feat/my-feature
   ```

2. Make your changes

3. Run tests:
   ```bash
   # Go
   cd cli && go test ./...
   
   # TypeScript
   cd packages/typescript && bun run build
   ```

4. Commit using conventional commit format (see below)

5. Push and create a PR

### Project Structure

```
xschema/
├── cli/                          # Go CLI
│   ├── cmd/                      # Cobra commands
│   ├── parser/                   # JSON/JSONC config parser
│   ├── retriever/                # Schema fetching
│   ├── generator/                # Adapter invocation
│   ├── injector/                 # Code injection
│   └── language/                 # Language-specific logic
├── packages/
│   └── typescript/               # TypeScript packages
│       ├── core/                 # @xschema/core - shared types
│       ├── client/               # @xschema/client - runtime
│       ├── adapters/
│       │   └── zod/              # @xschema/zod - Zod adapter
│       └── example/              # Example project
└── docs/                         # Documentation
```

## Commit Messages

We use [Conventional Commits](https://www.conventionalcommits.org/). All commits are validated via husky + commitlint.

### Format

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

### Types

| Type       | Description                      |
| ---------- | -------------------------------- |
| `feat`     | New feature                      |
| `fix`      | Bug fix                          |
| `docs`     | Documentation only               |
| `style`    | Formatting, no code change       |
| `refactor` | Code change, no feature/fix      |
| `perf`     | Performance improvement          |
| `test`     | Adding/updating tests            |
| `chore`    | Build process, dependencies      |

### Scopes

| Scope    | Description         |
| -------- | ------------------- |
| `cli`    | Go CLI              |
| `core`   | @xschema/core       |
| `client` | @xschema/client     |
| `zod`    | @xschema/zod        |
| `deps`   | Dependency updates  |

### Examples

```bash
# Feature
git commit -m "feat(cli): add --dry-run flag"

# Bug fix
git commit -m "fix(zod): handle array types correctly"

# Breaking change
git commit -m "feat(core)!: rename ConvertResult to GenerateResult"

# Docs (no scope needed)
git commit -m "docs: add installation instructions"
```

### Breaking Changes

For breaking changes, either:

1. Add `!` after the type/scope:
   ```bash
   git commit -m "feat(cli)!: change config file format"
   ```

2. Or add a `BREAKING CHANGE:` footer:
   ```bash
   git commit -m "refactor(core): simplify adapter interface

   BREAKING CHANGE: Adapters must now implement the new interface"
   ```

## Pull Requests

### Before Submitting

- [ ] Tests pass (`go test ./...` in cli/, `bun run build` in packages/typescript/)
- [ ] Commits follow conventional commit format
- [ ] PR description explains the changes

### PR Title

Use conventional commit format for PR titles too:

```
feat(cli): add schema validation
fix(zod): handle nullable arrays
docs: update README
```

### Review Process

1. CI must pass (tests, linting, commitlint)
2. At least one maintainer approval
3. Squash-merge to master

## Releases

Releases are automated via [release-please](https://github.com/googleapis/release-please). When your PR is merged:

1. release-please analyzes commits
2. Creates/updates a Release PR with changelog
3. When Release PR is merged, packages are published

You don't need to manually update versions or changelogs.

## Questions?

- Open an issue for bugs or feature requests
- Start a discussion for questions
