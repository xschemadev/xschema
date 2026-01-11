# Adapters Optimization Discovery Workflow

## Directory Structure

Discovery docs live in this folder (`tasks/adapters-optimization/`):

- `core-discovery.md` - core/ir parser behavior analysis
- `{adapter}-discovery.md` - per-adapter compliance analysis (zod, valibot, typebox, effect, arktype, typescript)

## Required Discovery Sections

Each discovery doc must include:

### For Core Discovery (`core-discovery.md`)
1. **Baseline observations** - current parser/ir behavior
2. **Keyword-based throws** - list of keywords that cause exceptions
3. **Preservation proposal** - how to preserve unknown/unsupported keywords (name + raw value + draft context)
4. **Migration notes** - minimal notes for adapter authors

### For Adapter Discovery (`{adapter}-discovery.md`)
1. **Baseline per draft** - pass % for each JSON Schema draft
2. **Failing tests list** - full list as `(keyword, group, test)` per draft
3. **Categorized failures**:
   - `core-missing` - blocked by core parser limitations
   - `adapter-native` - fixable using adapter's native features
   - `forced-emulation` - requires non-idiomatic workarounds
   - `not-supported` - intentionally out of scope
4. **Expected regressions** - changes that might break existing behavior

## The Rule

After completing a discovery story:

1. Add 1-3 new **one-iteration implementation stories** into `prd.jsonc`
2. Insert them **directly after** the discovery story
3. Order by **dependencies first**
4. **Renumber all subsequent story IDs** to remain sequential (US-XXX format)

This keeps the PRD as the single source of truth and ensures implementation stories are scoped to one iteration.
