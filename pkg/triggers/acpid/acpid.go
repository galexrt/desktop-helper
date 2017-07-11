package network

import (
	"context"

	"github.com/galexrt/desktop-helper/pkg/triggers"
)

func init() {
	triggers.Register("acpid", NewACPID())
}

// ACPID contains options
type ACPID struct {
	triggers.Trigger
}

// NewIPRange create new IPRange struct
func NewACPID() triggers.Trigger {
	return &ACPID{}
}

// Match against the given options
func (acpid ACPID) GetState(ctx context.Context, config interface{}) (bool, error) {
	return true, nil
}
