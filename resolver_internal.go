package resolver

import (
	"net"
	"time"

	"github.com/miekg/dns"
)

func exchange(m *dns.Msg, t time.Duration) (a *dns.Msg, err error) {
	srvIdx := 0

	for i := 0; i < config.Attempts; i++ {
		a, err = dnsExchangeOne(m, config.Servers[srvIdx], t)
		if err == nil {
			return
		}

		srvIdx++
		if srvIdx >= len(config.Servers) {
			srvIdx = 0
		}
	}

	return nil, ErrorNoHost
}

func resolveIPv4(addr string, timeout time.Duration) (ips []net.IP, err error) {
	m := dnsMakeQuestion(addr, dns.TypeA)
	in, err := exchange(m, timeout)

	if err != nil {
		return nil, err
	}

	for _, record := range in.Answer {
		switch t := record.(type) {
		case *dns.CNAME:
			return resolveIPv4(t.Target, timeout)
		case *dns.A:
			ips = append(ips, t.A)
		}
	}

	if len(ips) > 0 {
		return ips, nil
	}
	return nil, ErrorNoHost
}

func resolveIPv6(addr string, timeout time.Duration) (ips []net.IP, err error) {
	m := dnsMakeQuestion(addr, dns.TypeAAAA)
	in, err := exchange(m, timeout)

	if err != nil {
		return nil, err
	}

	for _, record := range in.Answer {
		switch t := record.(type) {
		case *dns.CNAME:
			return resolveIPv6(t.Target, timeout)
		case *dns.AAAA:
			ips = append(ips, t.AAAA)
		}
	}

	if len(ips) > 0 {
		return ips, nil
	}
	return nil, ErrorNoHost
}

func resolveIPvAUTO(addr string, timeout time.Duration) (ips []net.IP, err error) {
	ips, err = resolveIPv6(addr, timeout)
	if err != nil {
		return resolveIPv4(addr, timeout)
	}
	return ips, err
}

func resolveIPvALL(addr string, timeout time.Duration) (ips []net.IP, err error) {
	ips4, _ := resolveIPv6(addr, timeout)
	ips6, _ := resolveIPv4(addr, timeout)

	ips = append(ips, ips4...)
	ips = append(ips, ips6...)
	if len(ips) > 0 {
		return ips, nil
	}
	return nil, ErrorNoHost
}

func resolvePTR(ip string, timeout time.Duration) (hosts []string, err error) {
	addr, err := dns.ReverseAddr(ip)
	if err != nil {
		return nil, err
	}

	m := dnsMakeQuestion(addr, dns.TypePTR)
	in, err := exchange(m, timeout)

	if err != nil {
		return nil, err
	}

	for _, record := range in.Answer {
		if t, ok := record.(*dns.PTR); ok {
			hosts = append(hosts, chdot(t.Ptr))
		}
	}

	if len(hosts) > 0 {
		return hosts, nil
	}
	return nil, ErrorNoHost
}
