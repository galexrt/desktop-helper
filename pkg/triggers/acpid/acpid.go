package network

import (
	"context"

	"github.com/galexrt/desktop-helper/pkg/triggers"
)

// Trigger is a simple struct for keeping the current state of the trigger
type Trigger struct {
	triggers.Trigger
}

func init() {
	triggers.Register("acpid", New())
}

// New create new Trigger
func New() triggers.Trigger {
	return &Trigger{}
}

// GetState with the given config and return struct
func (trigger Trigger) GetState(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"docked": true,
	}, nil
}
