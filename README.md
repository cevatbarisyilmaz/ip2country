# ip2country
[![GoDoc](https://godoc.org/github.com/cevatbarisyilmaz/ip2country?status.svg)](https://godoc.org/github.com/cevatbarisyilmaz/ip2country)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/cevatbarisyilmaz/ip2country?sort=semver)](https://github.com/cevatbarisyilmaz/ip2country/releases)
[![GitHub](https://img.shields.io/github/license/cevatbarisyilmaz/ip2country)](https://github.com/cevatbarisyilmaz/ip2country/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/cevatbarisyilmaz/ip2country)](https://goreportcard.com/report/github.com/cevatbarisyilmaz/ip2country)

Package ip2country uses https://github.com/oschwald/maxminddb-golang with an embedded GeoLite2 DB to map `net.IP`s to country ISO codes without dealing with DB files.

## Example

```go
country, _ := ip2country.Country(net.IPv4(8, 8, 8, 8))
fmt.Println(country)
// Output: US
```
This product includes GeoLite2 data created by MaxMind, available from https://www.maxmind.com
