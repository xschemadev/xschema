# xschema-pydantic

XSchema adapter for generating Pydantic v2 models from JSON Schema.

## Local Verification

All commands run from the adapter directory (`python/packages/adapters/pydantic/`).

### Compliance (runtime validation)

```bash
# from repo root cli/ directory — builds CLI then runs full compliance
cd ../../cli && go build -o xschema . && ./xschema compliance --lang python --adapter-path ../python/packages/adapters/pydantic
```

### Unit tests

```bash
uv run pytest
```

### Type checking (static)

```bash
uv run pyright src/
```

## Fallback Typing Guardrails

### Allowed `Any` fallbacks

These constructs produce `Annotated[Any, BeforeValidator(...)]` by design. The static type is `Any` because the semantic domain is unbounded — no Python type expression can represent the true set of valid values.

| Construct | Reason | Example |
|-----------|--------|---------|
| `not` | Negation is "everything except X" — no union can express this | `not: { type: "string" }` accepts any non-string |
| `conditional` (if/then only, if/else only) | When only one branch exists, values not matching `if` pass through unconstrained | `if: {...}, then: {...}` with no `else` |
| `typeGuarded` (heterogeneous dispatch) | Unmatched types pass through the guard — domain is unbounded | Type guard with object/array branches, but string input passes |
| Open tuple (no `items: false`) | Extra elements beyond prefix items are untyped | `prefixItems: [string, int]` without `items: false` |
| Recursive refs | Self-referencing `$ref` (cycle detected, path starts with `#`) | `$ref: "#"` or `$ref: "#/$defs/TreeNode"` |

### Narrowed constructs (must NOT use `Any`)

These constructs must produce a narrower type than `Any`.

| Construct | Expected type | How it narrows |
|-----------|--------------|----------------|
| `oneOf` | `T1 \| T2 \| ...` | `_union_base_type()` computes union from sub-schema types |
| `conditional` (if/then/else) | `ThenType \| ElseType` | `_union_base_type([then_type, else_type])` |
| Closed tuple (`items: false`) | `tuple[T1 \| T2, ...]` | `_union_base_type(item_types)` on prefix items |
| Tuple with rest schema | `tuple[T1 \| ... \| TRest, ...]` | `_union_base_type(item_types + [rest_type])` |
| `const` (object) | `dict` | `_json_value_type(value)` |
| `const` (array) | `list` | `_json_value_type(value)` |
| `const` (primitive) | `bool`, `int`, `float`, `str`, `None` | `_json_value_type(value)` |
| `enum` (complex values) | `dict \| str`, `list \| int`, etc. | `_json_values_union_type(values)` |
| `allOf` (object merges) | BaseModel subclass | Intersection merges fields into a single model |
| Type-guarded object | BaseModel subclass | Guard dispatches to object renderer |

### Disallowed patterns

These patterns are errors. The renderer must raise `ConversionError` instead:

- **Unknown IR node kind**: the catch-all `case _:` in `render()` raises `ConversionError`
- **Unresolved external refs**: `render_ref` raises `ConversionError` when `resolved` is None and path doesn't start with `#`
- **Silent `Any` fallback**: new renderer branches must NOT return `Any` without an explicit, documented reason. If a construct can be narrowed, it must be.

### Required test coverage for new renderer branches

When adding or modifying a renderer function:

1. **If the branch produces `Any`**: add a unit test proving it (e.g. `test_render_not_stays_any`) and document the reason in the "Allowed `Any` fallbacks" table above
2. **If the branch produces a narrower type**: add a unit test proving the narrow type (e.g. `test_render_oneof_narrows_primitives`)
3. **If the branch raises `ConversionError`**: add a unit test proving the error (e.g. `test_render_unknown_ir_kind_raises`)
4. **Compliance must not regress**: run full compliance after any renderer change

## Troubleshooting

### pyright errors on `src/`

The adapter source must pass `uv run pyright src/` with zero errors. Common issues:

- `reveal_type()` left in source code — remove before committing
- Missing type stubs — add to pyproject.toml `[tool.pyright]` config
- `Any` in type annotations — use explicit types or `object` where possible

### Compliance failures after renderer changes

Run the full compliance suite and compare against baseline. Common causes:

- Changed runtime validation behavior (not just types) — the validator lambda must be functionally identical
- Changed intersection mutation semantics — intersection validators must NOT reassign `v`
- Type guard check too narrow — `isinstance(v, (list, tuple))` not just `isinstance(v, list)` for array checks
