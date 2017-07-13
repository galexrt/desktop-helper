package screenlayout

import (
	"context"

	"github.com/galexrt/desktop-helper/pkg/actions"
)

// Action contains options
type Action struct {
	actions.Action
}

func init() {
	actions.Register("screenlayout", New())
}

// New create new ScreenLayout struct
func New() actions.Action {
	return &Action{}
}

// Run the given options
func (action Action) Run(ctx context.Context, opts map[string]interface{}) (string, error) {
	return "", nil
}
