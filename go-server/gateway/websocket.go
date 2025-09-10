package gateway

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

type WebsocketManager struct {
	conns map[*websocket.Conn]bool
	log   *log.Logger
}

func NewWebsocketManager(log *log.Logger) *WebsocketManager {
	return &WebsocketManager{
		conns: make(map[*websocket.Conn]bool),
		log:   log,
	}
}
func (s *WebsocketManager) HandleWS(ws *websocket.Conn) {
	// fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	// Mettre une protection contre les raiseconditions
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *WebsocketManager) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			s.log.Println("read error", err)
		}
		msg := buf[:n]
		log.Println(string(msg))
		_, err = ws.Write([]byte("thx"))
		if err != nil {
			s.log.Println(err)
		}
	}
}
