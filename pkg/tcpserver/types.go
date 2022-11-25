package tcpserver

import "net"

type TCPServer struct {
	Host   string
	Port   string
	Dialer net.Conn
	Err    error
}
type TCPCommunication struct {
	// Dont forget to set NewMessage to true if you want to send a message to the client
	NewMessage bool
}
