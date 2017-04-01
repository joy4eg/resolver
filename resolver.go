package resolver

import (
	"errors"
	"net"
	"sort"
)

var (
	ErrorNoHost           = errors.New("no such host")
	ErrorInvalidIPversion = errors.New("invalid IP version given")
)

func ResolvePTR(ip string) (host string, err error) {
	hosts, err := resolvePTR(ip, config.Timeout)
	if err != nil {
		return "", err
	}

	sort.Strings(hosts)
	return hosts[0], nil
}

func ResolveIPv4(addr string) (ip net.IP, err error) {
	ips, err := resolveIPv4(addr, config.Timeout)
	if err != nil {
		return nil, err
	}

	return firstIP(ips), nil
}

func ResolveIPv6(addr string) (ip net.IP, err error) {
	ips, err := resolveIPv6(addr, config.Timeout)
	if err != nil {
		return nil, err
	}
	return firstIP(ips), nil
}

func ResolveIPvAUTO(addr string) (ip net.IP, err error) {
	ips, err := resolveIPvAUTO(addr, config.Timeout)
	if err != nil {
		return nil, err
	}
	return firstIP(ips), nil
}

func ResolveIPvALL(addr string) (ips []net.IP, err error) {
	return resolveIPvALL(addr, config.Timeout)
}

type IPVersion string

const (
	IPv4    IPVersion = "ipv4"
	IPv6    IPVersion = "ipv6"
	IPvAUTO IPVersion = "auto"
)

func ResolveIPVersion(version IPVersion, addr string) (ip net.IP, err error) {
	if ip = net.ParseIP(addr); ip != nil {
		return ip, nil
	}

	switch version {
	case IPvAUTO:
		return ResolveIPvAUTO(addr)
	case IPv6:
		return ResolveIPv6(addr)
	case IPv4:
		return ResolveIPv4(addr)
	}

	return nil, ErrorInvalidIPversion
}
