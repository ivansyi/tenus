package tenus

import (
	"github.com/milosgajdos83/libcontainer/netlink"
)

type NetworkOptions struct {
	IpAddr string
	Gw     string
	Routes []netlink.Route
}
