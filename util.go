package resolver

import (
	"bytes"
	"net"
)

func chdot(s string) string {
	var dot string = "."

	if l := len(s); string(s[l-1]) == dot {
		return string(s[:l-1])
	}
	return s
}

func firstIP(ips []net.IP) net.IP {
	if len(ips) == 0 {
		return nil
	}

	ip := ips[0]
	for _, v := range ips {
		if bytes.Compare(ip, v) == 1 {
			ip = v
		}
	}
	return ip
}
