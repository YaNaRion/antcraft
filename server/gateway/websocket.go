package gateway

import (
	"io"
	"log"
	"main/service/game"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

type WebsocketManager struct {
	clients     map[ClientID]*Client
	log         *log.Logger
	gameManager *game.GameManager
	// eventManager *EventManager
}

func NewWebsocketManager(log *log.Logger) *WebsocketManager {
	return &WebsocketManager{
		clients:     make(map[ClientID]*Client),
		log:         log,
		gameManager: game.NewGameManager(),
		// eventManager: newEventManager(),
	}
}

func (s *WebsocketManager) CleanClient() {
	s.clients = make(map[ClientID]*Client)
}

// ON CONNECT FONCTION
func (s *WebsocketManager) HandleWS(ws *websocket.Conn) {
	s.log.Println("New incoming connection from client:", ws.RemoteAddr())
	newClient := newClient(ws, ClientID(uuid.New()), ws.RemoteAddr())

	// Mettre une protection contre les raiseconditions ou prob de sync
	s.clients[newClient.GetID()] = newClient

	// Fonction that lisen to message
	s.readLoop(ws)
}

// READ LOOP ON THE SOCKET
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

		switch x := event.DataEvent.(type) {
		case *Event_JoinGame:
			s.log.Println("NEW EVENT: JOIN GAME")
			s.JoinGameHandler(ws, x, event.GameId)
		case *Event_MoveElement:
			s.log.Println("NEW EVENT: MOVE ELEMENT")
			s.MoveElementHandler(ws, x, &event.GameId, &event.PlayerInfo.PlayerId)
		case *Event_StartGame:
			s.log.Println("NEW EVENT: MOVE START GAME")
			s.StartGameHandler(ws, game.GameID(event.GameId))
		case *Event_AttackElement:
			s.log.Println("NEW EVENT: ATTACKING ELEMENT")
			s.AttackHandler(ws, game.GameID(event.GameId))
		}
	}
}
