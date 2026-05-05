package registry

import (
	"testing"

	"github.com/duketopceo/nanoclaw/internal/errors"
)

func TestRegistry_For(t *testing.T) {
	// Test happy path
	tools, err := For("health")
	if err != nil {
		t.Fatalf("expected no error for health droid, got %v", err)
	}
	if len(tools) == 0 {
		t.Error("expected tools for health droid, got 0")
	}

	// Test whitelisting: SignDroid cannot access DeleteUser (not in list)
	_, err = For("unknown")
	if err == nil {
		t.Error("expected error for unknown droid, got nil")
	}
	
	if err.Error() != errors.ErrToolNotAllowed.Error()+": unknown" {
		t.Errorf("expected ErrToolNotAllowed, got %v", err)
	}
}
