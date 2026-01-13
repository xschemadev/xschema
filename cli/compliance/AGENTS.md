# Compliance Package

Tests adapters against the official JSON Schema Test Suite to verify correctness.

## Overview

The compliance runner:
1. Loads test cases from the JSON Schema Test Suite
2. Bundles schemas through the processor pipeline
3. Calls the adapter to generate validator code
4. Executes a harness to validate test data against generated validators
5. Compares results and generates reports

## Key Files

- `runner.go` - main test execution logic, supports parallel keyword processing
- `harness.go` - generates and executes test harness code
- `types.go` - data structures for test cases, results, reports
- `report.go` - generates markdown/JSON reports
- `unsupported-features.json` - globally unsupported features across all adapters

## Determinism

**Results must be deterministic** - running compliance twice with no code changes must produce identical output files.

When adding parallel execution or processing unordered data (maps, goroutines), always sort results before writing. Example locations that required sorting:
- `runDraftParallel`: sorts `UnsupportedFeatures.Items` after parallel keyword processing
- `runDraft`: sorts keywords before processing

If you add new parallel code paths or aggregate results from concurrent operations, ensure the final output is sorted.
