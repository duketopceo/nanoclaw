package errors

import "fmt"

var (
	ErrTenantNotFound  = fmt.Errorf("tenant not found")
	ErrToolNotAllowed  = fmt.Errorf("tool not allowed for this droid")
	ErrSchemaMismatch = fmt.Errorf("llm output does not match tool schema")
	ErrTierDenied      = fmt.Errorf("action denied by autonomy tier gate")
)

type SchemaMismatchError struct {
	Raw    string
	Reason string
}

func (e SchemaMismatchError) Error() string {
	return fmt.Sprintf("schema mismatch: %s (reason: %s)", e.Raw, e.Reason)
}
