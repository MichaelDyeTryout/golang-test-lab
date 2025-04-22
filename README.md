# `golang-test-lab`

## Introduction

This is a small, isolated project for investigating and demonstrating Golang code quality improvement tooling and processes. The code itself will be trivial but still demonstrate testability that is relevant to real-world coding.

## Manual Testing

### Preconditions

* Golang v. 1.24.1 or newer
* GNU `make`, `sed`, `bash`

### Tryout Commands

* `make test` - Execute all test types, exiting on failure
* `make test-integration` - Execute only integration tests
* `make test-integration-cov` - Generate coverage report based on integration test run (note: purposefully avoids up-to-date behavior; re-run `make test-integration` beforehand to ensure newest coverage data is gathered.)
* `make test-unit` - Execute only unit tests
* `make test-unit-cov` - Generate coverage report for unit tests (note: purposefully avoids up-to-date behavior)
* `make test-unit-cov-viz` - Generate HTML coverage visualization; launch in default browser
* `make clean` - Delete all ephemeral, local-project content, esp. coverage report data and instrumented binaries

## Integration with GitHub Actions

TODO...