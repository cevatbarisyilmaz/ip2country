package ip2country_test

import (
	"github.com/cevatbarisyilmaz/ip2country"
	"net"
	"testing"
)

func TestCountry(t *testing.T) {
	testCases := map[string]string{
		"2.16.6.8": "DE",
		"8.8.8.8":  "US",
	}
	for ip, actualCountry := range testCases {
		country, err := ip2country.Country(net.ParseIP(ip))
		if err != nil {
			t.Fatal(err)
		}
		if country != actualCountry {
			t.Error("want:", actualCountry, "got:", country)
		}
	}
}
