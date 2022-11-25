package tcpserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// CreateListener is always run on background when program first starts.
//
// This function creates a TCP listener on the host and port that is set on environment variables. You can set them on docker composer.
//
// TCP servers are used for communication between players.
func (s *TCPServer) CreateListener() {
	// Start a new tcp listener on the specified host and port
	log.Println("Starting TCP server on " + s.Host + ":" + s.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Host, s.Port))
	if err != nil {
		log.Println(err)
		return
	}
	defer listener.Close()
	connection, err := listener.Accept()
	if err != nil {
		log.Println(err)
		return
	}
	for {
		switch msg, err := bufio.NewReader(connection).ReadString('\n'); {
		case strings.TrimSpace((msg)) == "STOP":
			fmt.Println("Exiting TCP server!")
			return
		case msg != "" && err == nil:
			fmt.Println("Message Received:", msg)
		}
	}
}
