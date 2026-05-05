package tier

import (
	"fmt"
	"github.com/duketopceo/nanoclaw/internal/errors"
)

// Gate checks if a tenant has the required autonomy tier for an action.
type Gate struct {
	TenantID string
}

// Check verifies the autonomy tier.
// Tier 1: Supervised (Always requires confirmation for writes)
// Tier 2: Semi-Autonomous (Low-risk automated)
// Tier 3: Autonomous
func (g *Gate) Check(currentTier int, toolName string, isWrite bool) error {
	if currentTier < 1 {
		return fmt.Errorf("invalid tier: %d", currentTier)
	}

	if isWrite && currentTier == 1 {
		return errors.ErrActionRequiresApproval
	}

	return nil
}
