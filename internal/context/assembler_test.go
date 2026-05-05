package context

import (
	"context"
	"testing"

	"github.com/duketopceo/nanoclaw/internal/errors"
)

func TestAssemble_Isolation(t *testing.T) {
	ctx := context.Background()

	// Test valid tenant
	res, err := Assemble(ctx, "tenant-a")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res != "context for tenant: tenant-a" {
		t.Errorf("expected context for tenant-a, got %s", res)
	}

	// Test missing tenant (Isolation guard)
	_, err = Assemble(ctx, "")
	if err != errors.ErrTenantNotFound {
		t.Errorf("expected ErrTenantNotFound, got %v", err)
	}
}
