package ip2country_test

import (
	"fmt"
	"github.com/cevatbarisyilmaz/ip2country"
	"net"
)

func ExampleCountry() {
	country, _ := ip2country.Country(net.IPv4(8, 8, 8, 8))
	fmt.Println(country)
	// Output: US
}
