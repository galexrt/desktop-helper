package network

import (
	"context"

	"github.com/galexrt/desktop-helper/pkg/triggers"
)

func init() {
	triggers.Register("network", NewNetworkTrigger())
}

// NetworkTrigger contains options
type NetworkTrigger struct {
	triggers.Trigger
}

// NewInterfaces create new Interfaces struct
func NewNetworkTrigger() triggers.Trigger {
	return &NetworkTrigger{}
}

// Match against the given options
func (networkTrigger NetworkTrigger) GetState(ctx context.Context, config interface{}) (bool, error) {
	return true, nil
}
