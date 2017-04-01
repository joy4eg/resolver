package resolver

import (
	"errors"
	"time"

	"github.com/miekg/dns"
	"golang.org/x/net/idna"
)

// Exchange performs a synchronous UDP query. It sends the message m to the address
// contained in a and waits for a reply. Exchange does not retry a failed query, nor
// will it fall back to TCP in case of truncation.
// See https://godoc.org/github.com/miekg/dns#Exchange for more information.
func dnsExchange(m *dns.Msg, a string, t time.Duration) (r *dns.Msg, err error) {
	var co *dns.Conn
	co, err = dns.DialTimeout("udp", a, t)
	if err != nil {
		return nil, err
	}

	defer co.Close()

	opt := m.IsEdns0()
	// If EDNS0 is used use that for size.
	if opt != nil && opt.UDPSize() >= dns.MinMsgSize {
		co.UDPSize = opt.UDPSize()
	}

	deadline := time.Now().Add(t)
	co.SetWriteDeadline(deadline)
	if err = co.WriteMsg(m); err != nil {
		return nil, err
	}

	co.SetReadDeadline(deadline)
	r, err = co.ReadMsg()
	if err == nil && r.Id != m.Id {
		err = dns.ErrId
	}
	return r, err
}

func dnsExchangeOne(m *dns.Msg, a string, t time.Duration) (*dns.Msg, error) {
	in, err := dnsExchange(m, a+":53", t)
	if err != nil {
		return nil, err
	}

	if in != nil {
		if in.Rcode != dns.RcodeSuccess {
			return nil, errors.New(dns.RcodeToString[in.Rcode])
		} else if len(in.Answer) == 0 {
			return nil, ErrorNoHost
		}
	}

	return in, nil
}

func dnsMakeQuestion(host string, Qtype uint16) *dns.Msg {
	if newHost, err := idna.ToASCII(host); err == nil {
		host = newHost
	}

	return &dns.Msg{
		Question: []dns.Question{
			dns.Question{
				Name:   dns.Fqdn(host),
				Qtype:  Qtype,
				Qclass: dns.ClassINET,
			},
		},
		MsgHdr: dns.MsgHdr{
			Id:               dns.Id(),
			RecursionDesired: true,
		},
	}
}
