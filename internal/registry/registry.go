package registry

import (
	"fmt"

	"github.com/duketopceo/nanoclaw/internal/errors"
)

type Tool struct {
	Name        string
	Description string
}

var droids = map[string][]Tool{
	"sign": {
		{Name: "SendDocument", Description: "Sends a document for signature"},
		{Name: "GetDocumentStatus", Description: "Checks status of a document"},
	},
	"health": {
		{Name: "PingService", Description: "Checks service availability"},
		{Name: "RestartService", Description: "Restarts a service"},
		{Name: "NotifyAdmin", Description: "Sends an alert to the operator"},
	},
}

func For(droidName string) ([]Tool, error) {
	tools, ok := droids[droidName]
	if !ok {
		return nil, fmt.Errorf("%w: %s", errors.ErrToolNotAllowed, droidName)
	}
	return tools, nil
}
