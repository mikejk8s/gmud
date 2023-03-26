package backend

import (
	"fmt"
	"net/http"

	"github.com/olahol/melody"
	"golang.org/x/net/websocket"
)

func StartWSServer() {
	m := melody.New()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})
	m.HandleConnect(func(s *melody.Session) {
		s.Write([]byte("sex"))
	})
	http.ListenAndServe(WebsocketAddr, nil)
}

func NewWebsocketUtil() (*WebsocketUtil, error) {
	conn, err := websocket.Dial(fmt.Sprintf("ws://%s", WebsocketAddr), "", Origin)
	if err != nil {
		return nil, err
	}
	return &WebsocketUtil{Conn: conn}, nil
}

func (ws *WebsocketUtil) Close() error {
	return ws.Conn.Close()
}
