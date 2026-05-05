## What changed
- Scaffolded NanoClaw Go module (`github.com/duketopceo/nanoclaw`).
- Implemented Triple-Lock boilerplate: CSI stub, Registry whitelisting, and Audit emitter.
- Added foundational tests for tenant isolation and tool whitelisting.
- Configured CI workflow with race detection and vet.

## Why
To establish a secure, multi-tenant-safe execution harness for Pace HQ agents, following the "Triple-Lock" architecture.

## How tested
- `go vet ./...`
- `go test -race ./...`

## Known risks
- CSI and LLM calls are currently stubs (Phase 1.0 task).
- No actual service integrations yet.

## Checklist
- [x] Tests added in this PR
- [x] No secrets in diff
- [x] No commented-out code
- [x] Conventional commits
- [x] CHANGELOG.md updated
- [x] go vet ./... passes locally
- [x] go test -race ./... passes locally
