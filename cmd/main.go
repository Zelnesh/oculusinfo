package main

import (

	"fmt"
	"flag"
	"net"

	"github.com/zelnesh/oculusinfo/internal/whois"
	"github.com/zelnesh/oculusinfo/internal/dnslookup"
	"github.com/zelnesh/oculusinfo/internal/portscanner"

)


func main() {

	var whoisArg bool
	var dnslookupArg bool
	var dnsserver string
	var ip string
	var ports string

	flag.BoolVar(&whoisArg, "whois", false, "Run IP lookup: oculusinfo -whois IP")
	flag.BoolVar(&whoisArg, "w", false, "Short arg key to run IP lookup: oculusinfo -w IP")
	flag.BoolVar(&dnslookupArg, "dnslookup", false, "Run DNS lookup: oculusinfo -dnslookup example.com")
	flag.BoolVar(&dnslookupArg, "d", false, "Short arg key to run DNS lookup: oculusinfo -d example.com")
	flag.StringVar(&dnsserver, "dns", "", "Custom DNS server: oculusinfo -dns 8.8.8.8 -d example.com")
	flag.StringVar(&ip, "ip", "", "IP address to scan.")
	flag.StringVar(&ports, "ports", "", "Ports to scan, single one or range, avaliable options are: oculusinfo -ip [IP] -ports 8080 || oculusinfo -ip [IP] -ports 80-400 || oculusinfo -ip [IP] -ports 80,4040,8080")
	flag.StringVar(&ports, "p", "", "Ports to scan, short key arg.")



	flag.Parse()

	switch {

		case whoisArg:
			args := flag.Args()
			if len(args) == 0 {
				fmt.Println("Missing IP target...")
				return
			}

			ip := args[0]
			fmt.Println("Running IP lookup...")
			ipLookupResult, err := whois.IPwhoIs(ip)
			if err != nil {
				fmt.Println(err)
				return
			}
			data := fmt.Sprintf("IP: %s\nCity: %s\nCountry: %s\nContinent: %s\nRegion: %s\nPostal: %s\nCalling Code: %s\nLatitude: %f\nLongitude: %f", ipLookupResult.IP, ipLookupResult.City, ipLookupResult.Country, ipLookupResult.Continent, ipLookupResult.Region, ipLookupResult.Postal, ipLookupResult.CallingCode, ipLookupResult.Latitude, ipLookupResult.Longitude)

			fmt.Println(data)

		case dnslookupArg:
			args := flag.Args()
			if len(args) == 0 {
				fmt.Println("Missing domain name target...")
				return
			}

			domain := args[0]
			fmt.Println("Runnins DNS Lookup...")
			dnsLookupResult, err := dnslookup.DnsLookup(domain, dnsserver)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Domain:", dnsLookupResult.Domain)
			fmt.Println("DNSserver:", dnsLookupResult.DNSserver)
			fmt.Printf("QueryTime: %dms", dnsLookupResult.QueryTime.Milliseconds())

			fmt.Println("\nA:")
			for _, v := range dnsLookupResult.A {
				fmt.Println(" -", v)
			}

			fmt.Println("\nAAAA:")
			for _, v := range dnsLookupResult.AAAA {
				fmt.Println(" -", v)
			}

			fmt.Println("\nMX:")
			for _, v := range dnsLookupResult.MX {
				fmt.Println(" -", v)
			}

			fmt.Println("\nNS:")
			for _, v := range dnsLookupResult.NS {
				fmt.Println(" -", v)
			}

			if dnsLookupResult.CNAME != "" {
				fmt.Println("\nCNAME:", dnsLookupResult.CNAME)
			}

			fmt.Println("\nTXT:")
			for _, v := range dnsLookupResult.TXT {
				fmt.Println(" -", v)
			}

		case ip != "":

			parsedIP := net.ParseIP(ip)
			if parsedIP == nil {
				fmt.Println("Wrong IP address")
				return
			}


			scanResults := portscanner.ScanPort(ip, ports)

			for result := range scanResults {
				fmt.Println(result)
			}

	}
}
