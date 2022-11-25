package tcpserver

import (
	"fmt"
	"net"
)

// NewTCPDialer assigns a new TCP dialer to the TCPServer struct.
func (s *TCPServer) NewTCPDialer() {
	s.Dialer, s.Err = net.Dial("tcp", fmt.Sprintf("%s:%s", s.Host, s.Port))
}

// Dont put \n in the message, it will be added automatically.
func (s *TCPServer) Writer(connMessage string) {
	fmt.Fprintf(s.Dialer, fmt.Sprintf("%s\n", connMessage))
}
