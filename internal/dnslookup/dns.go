package dnslookup


import (

	"time"

	"github.com/miekg/dns"
)

type DNSLookupResult struct {
	Domain string
	DNSserver string
	QueryTime time.Duration

	A []string
	AAAA []string
	MX []string
	NS []string
	CNAME string
	TXT []string

}

func DnsLookup (domain, customdns string) (*DNSLookupResult, error) {

	if customdns == "" {
		customdns = "1.1.1.1"
	}

	dnslookupresult := &DNSLookupResult{}
	dnslookupresult.Domain = domain
	dnslookupresult.DNSserver = customdns

	msgA := new(dns.Msg)
	msgA.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	msgMX := new(dns.Msg)
	msgMX.SetQuestion(dns.Fqdn(domain), dns.TypeMX)
	msgNS := new(dns.Msg)
	msgNS.SetQuestion(dns.Fqdn(domain), dns.TypeNS)
	msgCNAME := new(dns.Msg)
	msgCNAME.SetQuestion(dns.Fqdn(domain), dns.TypeCNAME)
	msgTXT := new(dns.Msg)
	msgTXT.SetQuestion(dns.Fqdn(domain), dns.TypeTXT)

	client := new(dns.Client)

	start := time.Now()

	respIP, _, err := client.Exchange(msgA, customdns + ":53")
	if err != nil {
		return nil, err
	}

	for _, ans := range respIP.Answer {
		switch version := ans.(type){
			case *dns.A:
				dnslookupresult.A = append(dnslookupresult.A, version.A.String())
			case *dns.AAAA:
				dnslookupresult.AAAA = append(dnslookupresult.AAAA, version.AAAA.String())
		}
	}

	respMX, _, err := client.Exchange(msgMX, customdns + ":53")
	if err != nil {
		return nil, err
	}

	for _, ans := range respMX.Answer {
		if mx, ok := ans.(*dns.MX); ok{
			dnslookupresult.MX = append(dnslookupresult.MX, mx.Mx)
		}
	}

	respNS, _, err := client.Exchange(msgNS, customdns + ":53")
	if err != nil {
		return nil, err
	}

	for _, ans := range respNS.Answer {
		if ns, ok := ans.(*dns.NS); ok{
			dnslookupresult.NS = append(dnslookupresult.NS, ns.Ns)
		}
	}

	respCNAME, _, err := client.Exchange(msgCNAME, customdns + ":53")
	if err != nil {
		return nil, err
	}

	for _, ans := range respCNAME.Answer {
		if cname, ok := ans.(*dns.CNAME); ok{
			dnslookupresult.CNAME = cname.Target
		}
	}

	respTXT, _, err := client.Exchange(msgTXT, customdns + ":53")
	if err != nil {
		return nil, err
	}

	if respTXT.Truncated {
		client.Net = "tcp"
		respTXT, _, err = client.Exchange(msgTXT, customdns + ":53")
		if err != nil {
			return nil, err
		}
	}

	for _, ans := range respTXT.Answer {
		if txt, ok := ans.(*dns.TXT); ok{
			dnslookupresult.TXT = append(dnslookupresult.TXT, txt.Txt...)
		}
	}

	for _, ans := range respTXT.Extra {
		if txt, ok := ans.(*dns.TXT); ok{
			dnslookupresult.TXT = append(dnslookupresult.TXT, txt.Txt...)
		}
	}

	dnslookupresult.QueryTime = time.Since(start)

	return dnslookupresult, nil


}
