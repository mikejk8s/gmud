package tcpserver

import "net"

// TCPServer is used for communication between players.
//
// A TCP Server is initialized when program is first run on TCP_SERVER_HOST and TCP_SERVER_PORT
// environment variables that can be set on docker-composer.yml.
//
// You can set a TCPServer struct to that same values and call NewTCPDialer() to dial to the TCP server that is running in the background.
//
// When TCPServer.NewTCPDialer() is called, do your all work on TCPServer.Dialer and dont forget to close it with TCPServer.CloseTCPDialer() when you are done.
type TCPServer struct {
	Host   string
	Port   string
	Dialer net.Conn
	Err    error
}

// This variables are initialized when first TCP server is created and you can use them to dial to the TCP server.
var (
	TCPHost string
	TCPPort string
)
