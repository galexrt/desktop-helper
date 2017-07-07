package network

import (
	"github.com/galexrt/desktop-helper/pkg/triggers"
)

// IPRange contains options
type IPRange struct {
	triggers.Trigger
}

// NewIPRange create new IPRange struct
func NewIPRange() triggers.Trigger {
	return &IPRange{}
}

// Match against the given options
func (iprange IPRange) Match(struct{}) (bool, error) {
	return true, nil
}
