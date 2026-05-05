package harness

import (
	"context"
	"testing"
)

func TestAgentExecution_Run(t *testing.T) {
	ae := &AgentExecution{
		TenantID:  "test-tenant-1",
		DroidName: "health",
		Input:     "check health",
		Tier:      1,
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
