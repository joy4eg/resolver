package resolver

import (
	"net"
	"testing"
)

// NOTE: Tests was created for OpenDNS server (208.67.220.220)

var ipMapA = map[string]net.IP{
	"ys.lc":      net.ParseIP("178.165.50.194"),
	"yap.ru":     net.ParseIP("5.187.1.122"),
	"ukr.net":    net.ParseIP("212.42.76.252"),
	"кто.рф":     net.ParseIP("195.208.25.143"),
	"yandex.ru":  net.ParseIP("5.255.255.50"),
	"reddit.com": net.ParseIP("151.101.1.140"),
	"notfound":   nil,
	"":           nil,
}

var ipMapAAAA = map[string]net.IP{
	"ya.ru":         net.ParseIP("2a02:6b8::3"),
	"wikipedia.com": net.ParseIP("2620:0:862:ed1a::1"),
	"mail.ru":       net.ParseIP("2a00:1148:db00:0:b0b0::1"),
	"notexist":      nil,
	"":              nil,
}

var ipMapCNAME = map[string]net.IP{
	"www.ya.ru": net.ParseIP("2a02:6b8::3"),
	"notexist":  nil,
	"":          nil,
}

var ipMapPTR = map[string]string{
	"178.165.50.194":           "undef-shyshkov-kh.maxnet.ua",
	"37.57.216.91":             "ns.triolan.net.ua",
	"213.180.204.3":            "www.yandex.ru",
	"217.69.139.202":           "mail.ru",
	"2a00:1148:db00:0:b0b0::1": "mail.ru",
	"2a02:6b8:a::a":            "yandex.ru",
	"notexists":                "",
}

var ipALLMap = map[string][]net.IP{
	"ys.lc":         []net.IP{net.ParseIP("178.165.50.194")},
	"pingtools.org": []net.IP{net.ParseIP("104.131.67.217")},
	"ukr.net":       []net.IP{net.ParseIP("212.42.76.252"), net.ParseIP("212.42.76.253")},
	"ok.ru":         []net.IP{net.ParseIP("5.61.23.5"), net.ParseIP("217.20.156.159"), net.ParseIP("217.20.155.58")},
	"wikipedia.com": []net.IP{net.ParseIP("91.198.174.192"), net.ParseIP("2620:0:862:ed1a::1")},
	"notexist":      nil,
	"":              nil,
}

func TestResolvIPv4(t *testing.T) {
	for k, v := range ipMapA {
		ip, err := ResolveIPv4(k)
		if !v.Equal(ip) {
			t.Error(k, ": expected", v, "got", ip, "(error is:", err, ")")
		}
	}
}

func TestResolvIPv6(t *testing.T) {
	for k, v := range ipMapAAAA {
		ip, err := ResolveIPv6(k)
		if !v.Equal(ip) {
			t.Error(k, ": expected", v, "got", ip, "(error is:", err, ")")
		}
	}
}

func TestResolvIPvAUTO(t *testing.T) {
	for k, v := range ipMapCNAME {
		ip, err := ResolveIPvAUTO(k)
		if !v.Equal(ip) {
			t.Error(k, ": expected", v, "got", ip, "(error is:", err, ")")
		}
	}
}

func TestResolvPTR(t *testing.T) {
	for k, v := range ipMapPTR {
		host, err := ResolvePTR(k)
		if host != v {
			t.Error(k, ": expected", v, "got", host, "(error is:", err, ")")
		}
	}
}

func ipsExists(ips []net.IP, ip net.IP) bool {
	for _, v := range ips {
		if v.Equal(ip) {
			return true
		}
	}
	return false
}

func ipsEqual(a, b []net.IP) bool {
	if len(a) != len(b) {
		return false
	}

	for _, v := range a {
		if !ipsExists(b, v) {
			return false
		}
	}

	return true
}

func TestResolvIPvALL(t *testing.T) {
	for k, v := range ipALLMap {
		ips, err := ResolveIPvALL(k)
		if !ipsEqual(v, ips) {
			t.Error(k, ": expected", v, "got", ips, "(error is:", err, ")")
		}
	}
}

func TestResolveIPVersion(t *testing.T) {
	for k, v := range ipMapA {
		ip, err := ResolveIPVersion(IPv4, k)
		if !v.Equal(ip) {
			t.Error(k, ": expected", v, "got", ip, "(error is:", err, ")")
		}
	}
	for k, v := range ipMapAAAA {
		ip, err := ResolveIPVersion(IPv6, k)
		if !v.Equal(ip) {
			t.Error(k, ": expected", v, "got", ip, "(error is:", err, ")")
		}
	}
}
