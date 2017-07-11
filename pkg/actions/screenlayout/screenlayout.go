package screenlayout

import (
	"github.com/galexrt/desktop-helper/pkg/actions"
)

// Action contains options
type Action struct {
	actions.Action
}

type ActionOptions struct {
}

func init() {
	actions.Register("screenlayout", New())
}

// New create new ScreenLayout struct
func New() actions.Action {
	return &Action{}
}

// Run the given options
func (action Action) Run(opts interface{}) error {
	return nil
}
