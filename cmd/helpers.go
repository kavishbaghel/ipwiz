package cmd

import (
	"inet.af/netaddr"
)

func CompareIps(startIP, endIP netaddr.IP) (result int) {
	return startIP.Compare(endIP)
}
