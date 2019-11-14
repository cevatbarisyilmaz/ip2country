/*
Package ip2country uses github.com/oschwald/maxminddb-golang with an embedded GeoLite2 db
to map net.IPs to country ISO codes
*/
package ip2country

import (
	"encoding/base64"
	"github.com/oschwald/maxminddb-golang"
	"net"
)

type record struct {
	Country struct {
		ISOCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
}

var reader *maxminddb.Reader

func init() {
	d2, _ := base64.StdEncoding.DecodeString(data)
	reader, _ = maxminddb.FromBytes(d2)
}

// Country returns ISO code of the country that given IP belongs to.
func Country(ip net.IP) (string, error) {
	r := &record{}
	err := reader.Lookup(ip, &r)
	if err != nil {
		return "", err
	}
	return r.Country.ISOCode, nil
}
