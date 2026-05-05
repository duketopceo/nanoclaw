package schema

import (
	"encoding/json"
	"fmt"
)

// Guard validates tool call arguments against a provided JSON schema.
type Guard struct{}

// Validate checks if the provided arguments match the given JSON schema.
// For Phase 0.8, we implement a strict structure check using standard encoding/json.
// In Phase 1.0, this will be upgraded to a full JSON Schema validator.
func (g *Guard) Validate(schemaJSON string, argsJSON string) error {
	var schema map[string]interface{}
	if err := json.Unmarshal([]byte(schemaJSON), &schema); err != nil {
		return fmt.Errorf("invalid schema: %w", err)
	}

	var args map[string]interface{}
	if err := json.Unmarshal([]byte(argsJSON), &args); err != nil {
		return fmt.Errorf("invalid arguments JSON: %w", err)
	}

	// Basic check: Ensure all required fields in schema exist in args
	if required, ok := schema["required"].([]interface{}); ok {
		for _, req := range required {
			field := req.(string)
			if _, exists := args[field]; !exists {
				return fmt.Errorf("missing required field: %s", field)
			}
		}
	}

	return nil
}
