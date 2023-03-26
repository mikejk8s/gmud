package backend

import "golang.org/x/net/websocket"

const (
	// If you want to change the websocket address, please only change this constant.
	//
	// A lot of things are bound to this little dude.
	WebsocketAddr = "0.0.0.0:5000"
	Origin        = "0.0.0.0"
)

type WebsocketUtil struct {
	Conn *websocket.Conn
}
