package resolver

import (
	"net"
	"time"
)

// LookupIP looks up host using the configured resolver.
// It returns an array of that host's IPv4 and IPv6 addresses.
func LookupIPTimeout(host string, timeout time.Duration) (ips []net.IP, err error) {
	if host == "" {
		return nil, &net.DNSError{Err: ErrorNoHost.Error(), Name: host}
	}
	ips, err = resolveIPvALL(host, timeout)
	if err != nil {
		return nil, &net.DNSError{Err: err.Error(), Name: host}
	}
	return ips, nil
}

// LookupAddr performs a reverse lookup for the given address, returning a list
// of names mapping to that address.
func LookupAddrTimeout(addr string, timeout time.Duration) (names []string, err error) {
	hosts, err := resolvePTR(addr, timeout)
	if err != nil {
		return nil, &net.DNSError{Err: err.Error(), Name: addr}
	}
	return hosts, nil
}
