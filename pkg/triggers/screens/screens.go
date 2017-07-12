package screens

import (
	"context"

	"github.com/galexrt/desktop-helper/pkg/triggers"
)

// Trigger is a simple struct for keeping the current state of the trigger
type Trigger struct {
	triggers.Trigger
}

func init() {
	triggers.Register("screens", New())
}

// New create new Trigger
func New() triggers.Trigger {
	return &Trigger{}
}

// GetState with the given config and return struct
func (trigger Trigger) GetState(ctx context.Context, config interface{}) (map[string]interface{}, error) {
	state := map[string]interface{}{
		"count": 3,
		"connected": []string{
			"DP1-1",
			"DP1-2",
			"DP1-3",
		},
	}
	return state, nil
}
