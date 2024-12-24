package envaddr

import (
	"net"
	"os"
)

func Get(defaultAddr string) string {
	var addr string = defaultAddr
	var port string = ""
	var portFound bool = false
	var host string = ""

	portVar, exists := os.LookupEnv("PORT")
	if exists {
		port = portVar
		portFound = true
	}
	hostVar, exists := os.LookupEnv("HOST")
	if exists {
		host = hostVar
	}
	ipEnv, exists := os.LookupEnv("IP")
	if exists {
		ip := net.ParseIP(ipEnv)
		if ip != nil {
			if ip.To4() != nil {
				host = ip.String()
			} else if ip.To16() != nil {
				host = "[" + ip.String() + "]"
			}
		}
	}

	if portFound {
		addr = host + ":" + port
	}

	return addr
}
