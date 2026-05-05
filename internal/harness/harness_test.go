package harness

import (
	"context"
	"testing"

	"github.com/duketopceo/nanoclaw/internal/llm"
)

type MockLLM struct{}

func (m *MockLLM) Chat(ctx context.Context, model string, messages []llm.Message, tools []any) (*llm.Message, error) {
	return &llm.Message{
		Role:    "assistant",
		Content: "Mock response",
	}, nil
}

func TestAgentExecution_Run(t *testing.T) {
	ae := &AgentExecution{
		TenantID:  "test-tenant-1",
		DroidName: "health",
		Input:     "check health",
		Tier:      1,
		LLMClient: &MockLLM{},
	}

	ctx := context.Background()
	res, err := ae.Run(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if res.Status != "completed" {
		t.Errorf("expected status 'completed', got %s", res.Status)
	}
}
