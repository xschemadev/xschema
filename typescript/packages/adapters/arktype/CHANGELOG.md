# Changelog

## [0.2.0](https://github.com/xschemadev/xschema/compare/ts-arktype-v0.1.3...ts-arktype-v0.2.0) (2026-02-12)


### ⚠ BREAKING CHANGES

* redesign internal cli flow ([#58](https://github.com/xschemadev/xschema/issues/58))
* rename validateImports to validationIMmports

### Features

* adapter compliance ([#48](https://github.com/xschemadev/xschema/issues/48)) ([f25debb](https://github.com/xschemadev/xschema/commit/f25debbf90a503d7bc1e8bdc29a4f57ecd06576a))
* build var name in cli ([0a6026e](https://github.com/xschemadev/xschema/commit/0a6026e545ef006f50e902d2e8a3dc5eb65d237d))
* internal language cli refactor ([#55](https://github.com/xschemadev/xschema/issues/55)) ([e5ac06b](https://github.com/xschemadev/xschema/commit/e5ac06b7b0a8be852c2518fcc0e0105d81b56956))
* perfect adapters ([53dbfac](https://github.com/xschemadev/xschema/commit/53dbfacc9f86fca3943968757ace7a66fe131400))
* python ([#71](https://github.com/xschemadev/xschema/issues/71)) ([0d3a0a7](https://github.com/xschemadev/xschema/commit/0d3a0a7d2d24852a4b454bf24e094a02ba0f3313))
* redesign internal cli flow ([#58](https://github.com/xschemadev/xschema/issues/58)) ([ce34e5c](https://github.com/xschemadev/xschema/commit/ce34e5c75363186f27350702fca8d2831f2d6be6))
* rename validateImports to validationIMmports ([da05d77](https://github.com/xschemadev/xschema/commit/da05d7756f63da12c707597fe986081e13b11f95))
* **ts:** arktype implementation of xschema ([#37](https://github.com/xschemadev/xschema/issues/37)) ([1b5ca3c](https://github.com/xschemadev/xschema/commit/1b5ca3cb4e661ebb8799cc6bf6a36300fae1e59d))
* **ts:** re-run arktype compliance + remove useless files from zod adapter ([abe0dbd](https://github.com/xschemadev/xschema/commit/abe0dbd1a9e54a2d1fa355d5906d2ca3c649332d))


### Bug Fixes

* **cli:** metaschema fetching ([d44eeab](https://github.com/xschemadev/xschema/commit/d44eeabe0c3a9f369a312b91b780d5ea8a5ab5b2))

## [0.1.3](https://github.com/xschemadev/xschema/compare/arktype-v0.1.2...arktype-v0.1.3) (2026-01-11)


### Features

* adapter compliance ([#48](https://github.com/xschemadev/xschema/issues/48)) ([f25debb](https://github.com/xschemadev/xschema/commit/f25debbf90a503d7bc1e8bdc29a4f57ecd06576a))
* build var name in cli ([0a6026e](https://github.com/xschemadev/xschema/commit/0a6026e545ef006f50e902d2e8a3dc5eb65d237d))
* internal language cli refactor ([#55](https://github.com/xschemadev/xschema/issues/55)) ([e5ac06b](https://github.com/xschemadev/xschema/commit/e5ac06b7b0a8be852c2518fcc0e0105d81b56956))

## [0.1.2](https://github.com/xschemadev/xschema/compare/arktype-v0.1.1...arktype-v0.1.2) (2026-01-06)


### Features

* **ts:** re-run arktype compliance + remove useless files from zod adapter ([abe0dbd](https://github.com/xschemadev/xschema/commit/abe0dbd1a9e54a2d1fa355d5906d2ca3c649332d))

## [0.1.1](https://github.com/xschemadev/xschema/compare/arktype-v0.1.0...arktype-v0.1.1) (2026-01-06)


### Features

* **ts:** arktype implementation of xschema ([#37](https://github.com/xschemadev/xschema/issues/37)) ([1b5ca3c](https://github.com/xschemadev/xschema/commit/1b5ca3cb4e661ebb8799cc6bf6a36300fae1e59d))

## [0.1.0](https://github.com/xschemadev/xschema/releases/tag/arktype-v0.1.0) (2026-01-06)

### Features

* initial arktype adapter with full JSON Schema compliance
  * 98.9% compliance on draft2020-12
  * 98.8% compliance on draft2019-09
  * 99.7% compliance on draft7
  * 99.6% compliance on draft6
  * 96.3% compliance on draft4
  * 96.0% compliance on draft3
  * 98.6% compliance on v1
