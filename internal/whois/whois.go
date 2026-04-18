package whois

import (

	"net/http"
	"encoding/json"
	"fmt"
)


type WhoIs struct {

	IP string `json:ip`
	City string `json:city`
	Country string `json:country`
	Continent string `json:continent`
	Region string `json:region`
	Postal string `json:postal`
	CallingCode string `json:calling_code`
	Latitude float64 `json:latitude`
	Longitude float64 `json:longitude`
}


func IPwhoIs (ip string) (*WhoIs, error) {

	url := fmt.Sprintf("https://ipwho.is/%s", ip)
	resp, err := http.Get(url)

	if err !=nil {
		return nil, err
	}

	defer resp.Body.Close()

	var whois WhoIs

	err = json.NewDecoder(resp.Body).Decode(&whois)
	if err != nil {
		return nil, err
	}

	return &whois, nil


}
