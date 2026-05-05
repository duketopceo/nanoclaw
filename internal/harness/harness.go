package harness

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/duketopceo/nanoclaw/internal/audit"
	ncContext "github.com/duketopceo/nanoclaw/internal/context"
	"github.com/duketopceo/nanoclaw/internal/llm"
	"github.com/duketopceo/nanoclaw/internal/registry"
	"github.com/duketopceo/nanoclaw/internal/schema"
	"github.com/duketopceo/nanoclaw/internal/tier"
)

// AgentExecution represents a single unit of work through the NanoClaw harness.
type AgentExecution struct {
	TenantID  string
	DroidName string
	Input     string
	Tier      int
	LLMClient llm.Client // Optional: will use OpenRouter if nil
}

// ExecutionResult captures the outcome of an agent run.
type ExecutionResult struct {
	Action  string
	Status  string
	Payload any
}

func (ae *AgentExecution) Run(ctx context.Context) (*ExecutionResult, error) {
	// Lock 1 — CSI (Context-Silo Injection)
	assembledContext, err := ncContext.Assemble(ctx, ae.TenantID)
	if err != nil {
		return nil, fmt.Errorf("context assembly: %w", err)
	}

	// Lock 2 — Whitelist
	reg, err := registry.For(ae.DroidName)
	if err != nil {
		return nil, fmt.Errorf("tool registry: %w", err)
	}

	// Prepare tools for LLM
	var llmTools []any
	for name, tool := range reg.Tools {
		llmTools = append(llmTools, map[string]any{
			"type": "function",
			"function": map[string]any{
				"name":        name,
				"description": tool.Description,
				"parameters":  json.RawMessage(tool.Schema),
			},
		})
	}

	// LLM Call
	var client llm.Client
	if ae.LLMClient != nil {
		client = ae.LLMClient
	} else {
		apiKey := os.Getenv("OPENROUTER_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("OPENROUTER_API_KEY not set")
		}
		client = llm.NewOpenRouterClient(apiKey)
	}
	
	messages := []llm.Message{
		{Role: "system", Content: assembledContext},
		{Role: "user", Content: ae.Input},
	}

	resp, err := client.Chat(ctx, "anthropic/claude-3.5-sonnet", messages, llmTools)
	if err != nil {
		return nil, fmt.Errorf("llm chat: %w", err)
	}

	// Lock 3 — Schema
	guard := &schema.Guard{}
	for _, tc := range resp.ToolCalls {
		tool, exists := reg.Tools[tc.Function.Name]
		if !exists {
			return nil, fmt.Errorf("unauthorized tool: %s", tc.Function.Name)
		}

		if err := guard.Validate(tool.Schema, tc.Function.Arguments); err != nil {
			return nil, fmt.Errorf("schema violation for tool %s: %w", tc.Function.Name, err)
		}
	}

	// Tier Gate
	gate := &tier.Gate{TenantID: ae.TenantID}
	for _, tc := range resp.ToolCalls {
		tool := reg.Tools[tc.Function.Name]
		if err := gate.Check(ae.Tier, tc.Function.Name, tool.IsWrite); err != nil {
			return nil, fmt.Errorf("tier gate: %w", err)
		}
	}

	// Audit — fires regardless of outcome
	audit.Emit(ae.TenantID, ae.DroidName, "execution_started", time.Now())

	return &ExecutionResult{
		Status:  "completed",
		Payload: resp,
	}, nil
}
