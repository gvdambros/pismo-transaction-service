package echo

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

// URL a running server's URL (used in integration tests)
func (s *Server) URL() string {
	if s.url == "" {
		s.url = fmt.Sprintf("http://%s:%s", s.ip(), s.port())
	}

	return s.url
}

func (s *Server) ip() string {
	host, _ := os.Hostname()

	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}

	return "::"
}

func (s *Server) port() string {
	tcpInfo := s.Listener.Addr().(*net.TCPAddr)
	return strconv.Itoa(tcpInfo.Port)
}
