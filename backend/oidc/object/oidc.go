package object

import (
	"fmt"
	"net"
	"strings"
)

type OidcDiscovery struct {
	Issuer           string `json:"issuer"`
	UserinfoEndpoint string `json:"userinfo_endpoint"`
	JwksUri          string `json:"jwks_uri"`
}

func GetOidcDiscovery(host string) OidcDiscovery {
	originBackend := getOriginFromHost(host)
	oidcDiscovery := OidcDiscovery{
		Issuer:           originBackend,
		UserinfoEndpoint: fmt.Sprintf("%s/api/userinfo", originBackend),
		JwksUri:          fmt.Sprintf("%s/.well-known/jwks.json", originBackend),
	}

	return oidcDiscovery
}

func getOriginFromHost(host string) string {
	protocol := "https://"
	if !strings.Contains(host, ".") || isIpAddress(host) {
		protocol = "http://"
	}

	return fmt.Sprintf("%s%s", protocol, host)
}

func isIpAddress(host string) bool {
	// Attempt to split the host and port, ignoring the error
	hostWithoutPort, _, err := net.SplitHostPort(host)
	if err != nil {
		// If an error occurs, it might be because there's no port
		// In that case, use the original host string
		hostWithoutPort = host
	}

	// Attempt to parse the host as an IP address (both IPv4 and IPv6)
	ip := net.ParseIP(hostWithoutPort)
	// if host is not nil is an IP address else is not an IP address
	return ip != nil
}
