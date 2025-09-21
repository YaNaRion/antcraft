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
	s.log.Println("new incoming connection from client:", ws.RemoteAddr())

	// Mettre une protection contre les raiseconditions
	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *WebsocketManager) readLoop(ws *websocket.Conn) {
	for {
		var msg []byte
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			s.log.Println("read error:", err)
			continue
		}

		var event Event
		err = proto.Unmarshal(msg, &event)
		if err != nil {
			s.log.Println("failed to unmarshal:", err)
			continue
		}

		switch x := event.Data_Event.(type) {
		case *Event_PlayerData:
			s.log.Println(x.PlayerData.UniqueId)
		case *Event_MoveUnit:
			var eventSend Event
			unit := x.MoveUnit
			s.log.Println(unit)

			eventSend.Data_Event = &Event_MoveUnit{
				MoveUnit: &MoveUnit{
					UnitId:   unit.UnitId,
					PlayerId: unit.PlayerId,
					OldPos:   unit.OldPos,
					NewPos:   unit.NewPos,
				},
			}

			s.log.Println("EVENT SEND UNIT")
			s.log.Println(eventSend.Data_Event)

			data, err := proto.Marshal(&eventSend)
			if err != nil {
				s.log.Println(err)
				continue
			}

			for con := range s.conns {
				err = websocket.Message.Send(con, data)
				if err != nil {
					s.log.Println("send error:", err)
				}
			}
		}
	}
}
