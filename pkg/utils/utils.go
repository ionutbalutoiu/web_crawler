package utils

import (
	"net"
	"strings"
)

// GetHostDomain returns the domain from a host string.
// If the given host string contains a port, it will be ignored.
func GetHostDomain(host string) (string, error) {
	if strings.Contains(host, ":") {
		domain, _, err := net.SplitHostPort(host)
		if err != nil {
			return "", err
		}
		return domain, nil
	}
	return host, nil
}
