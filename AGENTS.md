# NanoClaw — Agent System Constitution

Inherits from [luke-agents/AGENTS.md](https://github.com/duketopceo/luke-agents/blob/main/AGENTS.md). This file specializes; it does not replace.

## Design Philosophy

NanoClaw is the execution engine for Pace HQ agents. It follows the **Triple-Lock** architecture to ensure multi-tenant safety and reasoning integrity.

### The Triple-Lock

1.  **Context-Silo Injection (CSI):** Every prompt is built from a merge of global git-versioned policies and a Supabase-stored tenant context slice. The harness enforces `WHERE tenant_id = ?` on all dynamic context fetches.
2.  **Tool-Registry Whitelisting:** Agents have zero access to raw system tools. They interact only through a compiled `ToolMap` specific to their Droid persona.
3.  **JSONSchemaGuard:** Every response from the LLM is validated against a Go struct schema before execution.

## Operating Rules

1.  **No Direct Writes:** Agents propose intents (typed structs). The service layer executes the mutation.
2.  **Audit First:** Every intent, execution, and outcome is logged to the `agent_events` table before the client receives a response.
3.  **Tier Gate:** Autonomy is earned. Tier 1 (Supervised) requires a human click for every write action. Tier 2/3 allows autonomous low-risk actions.

---

## Droid Personas

- **SignDroid:** Document lifecycle and signing authority.
- **HealthDroid:** Self-healing infrastructure and service monitoring.
- **VaultDroid:** Password hygiene and breach detection (Tier 1 Always).
- **FilesDroid:** Storage management and cleanup.
- **AdminDroid:** User provisioning and offboarding.
