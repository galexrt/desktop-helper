package network

import "github.com/galexrt/desktop-helper/pkg/triggers"

func init() {
	triggers.Register("iprange", NewIPRange())
	triggers.Register("iprange", NewInterfaces())
}
