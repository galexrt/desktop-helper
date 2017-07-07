package screens

import (
	"github.com/galexrt/desktop-helper/pkg/triggers"
)

// Screens contains options
type Screens struct {
	triggers.Trigger
}

func init() {
	triggers.Register("screens", NewScreens())
}

// NewScreens create new Screens struct
func NewScreens() triggers.Trigger {
	return &Screens{}
}

// Match against the given options
func (screens Screens) Match(struct{}) (bool, error) {
	return true, nil
}
