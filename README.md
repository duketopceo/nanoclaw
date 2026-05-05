# NanoClaw

Tenant-scoped agentic execution harness for Pace HQ.

## Overview

NanoClaw is the "engine" that powers Pace HQ Droids. It implements a secure, triple-locked pipeline for LLM-based tool execution, ensuring that an agent can never access data or execute actions outside its strictly defined tenant scope.

## Architecture

- **CSI (Context-Silo Injection):** Merges policy and tenant data safely.
- **Tool Registry:** Whitelists tools per Droid persona.
- **JSONSchemaGuard:** Validates LLM outputs against Go types.
- **Tier Gate:** Enforces human-in-the-loop based on autonomy tiers.

## Quick Start

```bash
go run cmd/nanoclaw/main.go --tenant [id] --droid health --input "check system health"
```

## Status: v0.1.0 (Scaffold)
- [x] Directory structure
- [x] AgentExecution core struct
- [x] Triple-Lock boilerplate
- [x] Cross-tenant isolation tests
