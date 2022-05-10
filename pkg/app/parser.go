package app

import (
	"net/netip"
	"strings"
)

func ParseRanges(ranges []string) ([]netip.Addr, error) {
	var ipResult []netip.Addr
	for _, cidr := range ranges {
		ip, err := ParseHosts(cidr)
		if err != nil {
			return nil, err
		}
		ipResult = append(ipResult, ip...)
	}
	return ipResult, nil
}

func ParseHosts(cidr string) ([]netip.Addr, error) {
	if !strings.Contains(cidr, "/") {
		singleIP, err := netip.ParseAddr(cidr)
		if err != nil {
			return nil, err
		}
		return []netip.Addr{singleIP}, nil

	}
	prefix, err := netip.ParsePrefix(cidr)
	if err != nil {
		return nil, err
	}

	var ips []netip.Addr
	for addr := prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
		ips = append(ips, addr)
	}

	if len(ips) < 2 {
		return ips, nil
	}

	return ips[1 : len(ips)-1], nil
}
