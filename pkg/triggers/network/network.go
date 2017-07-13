package network

import (
	"context"
	"net"
	"strconv"
	"strings"

	"github.com/galexrt/desktop-helper/pkg/triggers"
)

// Trigger is a simple struct for keeping the current state of the trigger
type Trigger struct {
	triggers.Trigger
}

func init() {
	triggers.Register("network", New())
}

// New create new Trigger
func New() triggers.Trigger {
	return &Trigger{}
}

// GetState with the given config and return struct
func (trigger Trigger) GetState(ctx context.Context, name string, config map[string]interface{}) (map[string]interface{}, error) {
	state := map[string]interface{}{}
	iface := strings.SplitN(name, "_", 2)[1]
	netIface, err := net.InterfaceByName(iface)
	if err != nil {
		return state, err
	}
	var options map[interface{}]interface{}
	if config["options"] != nil {
		options = config["options"].(map[interface{}]interface{})
		delete(config, "options")
	} else {
		options = map[interface{}]interface{}{
			"addrsKey": "0",
		}
	}
	for key := range config {
		switch key {
		case "ip":
			addrs, err := netIface.Addrs()
			if err != nil {
				return state, nil
			}
			if len(addrs) > 0 {
				addrsKey, _ := strconv.ParseInt(options["addrsKey"].(string), 10, 16)
				ip := addrs[addrsKey]
				state = map[string]interface{}{
					"ip": ip.String(),
				}
			}
		case "subnet":
			// TODO
		case "network_name":
			// TODO
		case "hardware_addr":
			// TODO
		}
	}
	return state, nil
}
