#!/bin/bash
set -euo pipefail

cd typescript
bun install
bun run build
