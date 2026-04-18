package main

import (

	"fmt"
	"flag"

	"github.com/zelnesh/oculusinfo/internal/whois"

)


func main() {

	var whoisArg bool

	flag.BoolVar(&whoisArg, "whois", false, "Run IP lookup: culusinfo -whois IP")
	flag.BoolVar(&whoisArg, "w", false, "Short key to run IP lookup: culusinfo -w IP")

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

	}
}
