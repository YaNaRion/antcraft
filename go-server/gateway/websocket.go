package gateway

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
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
	log.Println("new incoming connection from client:", ws.RemoteAddr())

	// Mettre une protection contre les raiseconditions
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *WebsocketManager) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	var event Event
	for {
		_, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			s.log.Println("read error", err)
		}
		err = proto.Unmarshal(buf, &event)
		if err != nil {
			s.log.Println(err)
		}
		switch x := event.Data_Event.(type) {
		case *Event_PlayerData:
			s.log.Println(x.PlayerData.UniqueID)
		case *Event_JoinRoomRequest:
		case *Event_JoinRoomResponse:
		case *Event_RoomStatusRequest:
		}
	}
}
