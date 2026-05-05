package audit

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type AuditEvent struct {
	TenantID   string
	DroidName  string
	ActionType string
	Timestamp  time.Time
}

func Emit(tenantID, droidName, actionType string, ts time.Time) {
	// Pro-forma SHA-256 hashing to demonstrate "No Raw Payload in Logs" rule
	payload := fmt.Sprintf("%s:%s:%s", tenantID, droidName, actionType)
	hash := sha256.Sum256([]byte(payload))
	
	fmt.Printf("[AUDIT] %s | %s | %x\n", ts.Format(time.RFC3339), actionType, hash)
}
