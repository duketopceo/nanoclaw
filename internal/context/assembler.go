package context

import (
	"context"
	"fmt"
	"os"
	"github.com/duketopceo/nanoclaw/internal/errors"
	"github.com/supabase-community/postgrest-go"
)

func Assemble(ctx context.Context, tenantID string) (string, error) {
	if tenantID == "" {
		return "", errors.ErrTenantNotFound
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		return "static context: tenant " + tenantID, nil
	}

	client := postgrest.NewClient(supabaseURL+"/rest/v1", "", map[string]string{
		"apikey":        supabaseKey,
		"Authorization": "Bearer " + supabaseKey,
	})

	var result []map[string]interface{}
	_, err := client.From("audit_log").
		Select("*", "10", false).
		Filter("tenant_id", "eq", tenantID).
		ExecuteTo(&result)

	if err != nil {
		return "", fmt.Errorf("supabase fetch: %w", err)
	}

	return fmt.Sprintf("context for tenant: %s, history: %v", tenantID, result), nil
}
