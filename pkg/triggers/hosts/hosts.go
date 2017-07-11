package network

import (
	"context"

	"github.com/galexrt/desktop-helper/pkg/triggers"
)

type HostsTrigger struct {
	triggers.Trigger
}

func init() {
	triggers.Register("hosts", NewHosts())
}

func NewHosts() triggers.Trigger {
	return &HostsTrigger{}
}

func (hostsTrigger HostsTrigger) GetState(ctx context.Context, config interface{}) (bool, error) {
	return true, nil
}
