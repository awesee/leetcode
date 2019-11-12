package problem468

import (
	"net"
	"regexp"
)

func validIPAddress(IP string) string {
	if p4 := net.ParseIP(IP).To4(); len(p4) == net.IPv4len && p4.String() == IP {
		return "IPv4"
	} else if regexp.MustCompile(`^([0-9A-Fa-f]{1,4}:){7}[0-9A-Fa-f]{1,4}$`).MatchString(IP) {
		return "IPv6"
	}
	return "Neither"
}
