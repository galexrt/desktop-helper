package network

import (
	"github.com/galexrt/desktop-helper/pkg/triggers"
)

// Interfaces contains options
type Interfaces struct {
	triggers.Trigger
}

// NewInterfaces create new Interfaces struct
func NewInterfaces() triggers.Trigger {
	return &Interfaces{}
}

// Match against the given options
func (interfaces Interfaces) Match(struct{}) (bool, error) {
	return true, nil
}
