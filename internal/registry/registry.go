package registry

import (
	"fmt"

	"github.com/duketopceo/nanoclaw/internal/errors"
)

type Tool struct {
	Name        string
	Description string
	Schema      string
	IsWrite     bool
}

type Registry struct {
	Tools map[string]Tool
}

var droids = map[string]Registry{
	"sign": {
		Tools: map[string]Tool{
			"SendDocument": {
				Name:        "SendDocument",
				Description: "Sends a document for signature",
				Schema:      `{"type": "object", "properties": {"document_id": {"type": "string"}}, "required": ["document_id"]}`,
				IsWrite:     true,
			},
			"GetDocumentStatus": {
				Name:        "GetDocumentStatus",
				Description: "Checks status of a document",
				Schema:      `{"type": "object", "properties": {"document_id": {"type": "string"}}, "required": ["document_id"]}`,
				IsWrite:     false,
			},
		},
	},
	"health": {
		Tools: map[string]Tool{
			"PingService": {
				Name:        "PingService",
				Description: "Checks service availability",
				Schema:      `{"type": "object", "properties": {"service": {"type": "string"}}, "required": ["service"]}`,
				IsWrite:     false,
			},
			"RestartService": {
				Name:        "RestartService",
				Description: "Restarts a service",
				Schema:      `{"type": "object", "properties": {"service": {"type": "string"}}, "required": ["service"]}`,
				IsWrite:     true,
			},
			"NotifyAdmin": {
				Name:        "NotifyAdmin",
				Description: "Sends an alert to the operator",
				Schema:      `{"type": "object", "properties": {"message": {"type": "string"}}, "required": ["message"]}`,
				IsWrite:     true,
			},
		},
	},
}

func For(droidName string) (Registry, error) {
	reg, ok := droids[droidName]
	if !ok {
		return Registry{}, fmt.Errorf("%w: %s", errors.ErrToolNotAllowed, droidName)
	}
	return reg, nil
}
