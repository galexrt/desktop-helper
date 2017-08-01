package ipaddress

import (
	"net"

	"github.com/galexrt/desktop-helper/config"
	"github.com/galexrt/desktop-helper/poller/triggers"
)

type Trigger struct {
	triggers.Trigger
	cfg   *config.IPAddressConfig
	state map[string]net.Interface
}

type IPAddress struct {
	Address string `yaml:"address"`
	Key     int    `yaml:"key"`
}

func init() {
	triggers.Register("ipaddress", New)
}

func New(cfg config.TriggersConfig) (triggers.Trigger, error) {
	return &Trigger{
		cfg:   cfg.IPAddress,
		state: make(map[string]net.Interface),
	}, nil
}

func (trg *Trigger) GetState() error {
	if trg.cfg != nil {
		for _, iface := range trg.cfg.Interfaces {
			addr, err := net.InterfaceByName(iface)
			if err != nil {
				return err
			}
			trg.state[iface] = *addr
		}
	}
	return nil
}

func (trg *Trigger) Match(opts config.TriggerOption) (bool, error) {
	var match bool
	for iface, desired := range opts.IPAddress.Addresses {
		if state, ok := trg.state[iface]; ok {
			addrs, err := state.Addrs()
			if err != nil {
				return false, err
			}
			if len(addrs) > desired.Key {
				addr := addrs[desired.Key]
				if desired.Address == addr.String() {
					match = true
				} else {
					return false, nil
				}
			} else {
				return false, nil
			}
		}
	}
	return match, nil
}
