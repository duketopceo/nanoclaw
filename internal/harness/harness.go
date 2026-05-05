package harness

import (
	"context"
	"fmt"
	"time"

	"github.com/duketopceo/nanoclaw/internal/audit"
	"github.com/duketopceo/nanoclaw/internal/llm"
	"github.com/duketopceo/nanoclaw/internal/registry"
	"github.com/duketopceo/nanoclaw/internal/schema"
)

// AgentExecution represents a single unit of work through the NanoClaw harness.
type AgentExecution struct {
	TenantID  string
	DroidName string
	Input     string
	Tier      int
}

// ExecutionResult captures the outcome of an agent run.
type ExecutionResult struct {
	Action  string
	Status  string
	Payload any
}

func (ae *AgentExecution) Run(ctx context.Context) (*ExecutionResult, error) {
	// Lock 1 — CSI (Context-Silo Injection)
	// TODO: Implement internal/context/assembler.go
	assembledContext, err := ae.assembleContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("context assembly: %w", err)
	}

	// Lock 2 — Whitelist
	reg, err := registry.For(ae.DroidName)
	if err != nil {
		return nil, fmt.Errorf("tool registry: %w", err)
	}

	// LLM Call
	client := llm.NewOpenRouterClient("standard")
	resp, err := client.Generate(ctx, assembledContext, ae.Input)
	if err != nil {
		return nil, fmt.Errorf("llm generation: %w", err)
	}

	// Lock 3 — Schema
	guard := &schema.Guard{}
	for _, tc := range resp.ToolCalls {
		tool, exists := reg.Tools[tc.Name]
		if !exists {
			return nil, fmt.Errorf("unauthorized tool: %s", tc.Name)
		}

		if err := guard.Validate(tool.Schema, tc.Arguments); err != nil {
			return nil, fmt.Errorf("schema violation for tool %s: %w", tc.Name, err)
		}
	}

	// Tier Gate
	// TODO: Implement internal/tier/gate.go

	// Audit — fires regardless of outcome
	audit.Emit(ae.TenantID, ae.DroidName, "execution_started", time.Now())

	return &ExecutionResult{
		Status:  "completed",
		Payload: resp,
	}, nil
}

func (ae *AgentExecution) assembleContext(ctx context.Context) (string, error) {
	// Placeholder for CSI logic
	return "base context", nil
}
