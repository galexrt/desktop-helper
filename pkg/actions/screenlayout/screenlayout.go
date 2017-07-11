package screenlayout

import (
	"github.com/galexrt/desktop-helper/pkg/actions"
)

// ScreenLayout contains options
type ScreenLayout struct {
	actions.Action
}

func init() {
	actions.Register("screenlayout", NewScreenLayout())
}

// NewScreenLayout create new ScreenLayout struct
func NewScreenLayout() actions.Action {
	return &ScreenLayout{}
}

// Run the given options
func (screenlayout ScreenLayout) Run(config interface{}) error {
	return nil
}
