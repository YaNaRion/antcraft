package gateway

import (
	"fmt"
	"io"
	"log"
	"main/service/game"
	"net"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

type ClientID uuid.UUID

type Client struct {
	Conn *websocket.Conn
	id   ClientID
	addr net.Addr
}

func newClient(conn *websocket.Conn, id ClientID, addr net.Addr) *Client {
	return &Client{
		Conn: conn,
		id:   id,
		addr: addr,
	}
}

func (c *Client) GetID() ClientID {
	return c.id
}

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

func (s *WebsocketManager) HandleWS(ws *websocket.Conn) {
	s.log.Println("New incoming connection from client:", ws.RemoteAddr())

	newClient := newClient(ws, ClientID(uuid.New()), ws.RemoteAddr())

	// Mettre une protection contre les raiseconditions ou prob de sync
	s.clients[newClient.GetID()] = newClient

	s.readLoop(ws)
}

func (s *WebsocketManager) JoinGameHandler(ws *websocket.Conn, event *Event_GameRequest) {
	s.log.Println("JOIN GAME EVENT")

	var newPlayer *game.Player
	var playerID string
	var gameID string = "Game1"

	var unit *game.Unit = game.NewUnit(game.Vector2{
		X: 300,
		Y: 300,
	}, game.Vector2{
		X: 10,
		Y: 10,
	})

	var unit2 *game.Unit = game.NewUnit(game.Vector2{
		X: 350,
		Y: 350,
	}, game.Vector2{
		X: 10,
		Y: 10,
	})

	playerID = fmt.Sprintf("Player%d", len(s.clients))
	newPlayer = game.NewPlayer(game.NewPlayerConn(ws), game.PlayerID(playerID))
	newPlayer.AddElement(unit)
	newPlayer.AddElement(unit2)

	s.gameManager.AddPlayerToGame(game.GameID(gameID), newPlayer)

	gameMap := s.gameManager.GetGame(game.GameID(gameID))
	err := gameMap.AddElement(unit)
	if err != nil {
		s.log.Println(err)
	}

	err = gameMap.AddElement(unit2)

	if err != nil {
		s.log.Println(err)
	}

	s.StartGame(ws, nil)
}

func (s *WebsocketManager) StartGame(ws *websocket.Conn, event *Event_StartGame) {
	s.log.Println("StartGame Event")
	// gameID := event.StartGame.GameId
	gameID := "Game1"
	go s.GameLoop(gameID)
}

func (s *WebsocketManager) GameLoop(gameID string) {
	s.log.Println("DANS GO FUNC DUPDATE")
	var err error
	gameMap := s.gameManager.GetGame(game.GameID(gameID))

	if gameMap.GameState != game.GAME_STATE_IN_GAME {
		return
	}

	var eventRespond Event
	var players []*Player
	for _, player := range s.gameManager.GetPlayerInAGame(game.GameID(gameID)) {
		players = append(players, &Player{
			UniqueId: player.GetPlayerIDString(),
		})
	}

	mapGrid := NewEventSyncGameState(
		s.gameManager.GetMap(game.GameID(gameID)),
		GameState(gameMap.GameState),
	)

	s.log.Println(mapGrid.SyncGameState.Elements)

	// METTRE MUTEX OU DE QUOI
	eventRespond.Data_Event = mapGrid

	marshalData, err := proto.Marshal(&eventRespond)
	if err != nil {
		s.log.Println("send error:", err)
	}

	for _, client := range s.clients {
		err = websocket.Message.Send(client.Conn, marshalData)
		if err != nil {
			s.log.Println("send error:", err)
		}
	}

	if err != nil {
		s.log.Println("send error:", err)
	}
}

func (s *WebsocketManager) MoveElementHandler(
	ws *websocket.Conn,
	unit *Event_MoveElement,
	gameID, playerID *string,
) {
	s.log.Println("MoveElement Event")
	elementNewPos := game.Vector2{
		X: int(unit.MoveElement.GetPos().XPos),
		Y: int(unit.MoveElement.GetPos().YPos),
	}
	elementOldPos := game.Vector2{
		X: int(unit.MoveElement.GetOldPos().XPos),
		Y: int(unit.MoveElement.GetOldPos().YPos),
	}
	err := s.gameManager.MoveElement(
		game.ElementID(unit.MoveElement.UnitId),
		game.PlayerID(unit.MoveElement.PlayerId),
		game.GameID(*gameID),
		elementNewPos,
		elementOldPos,
	)
	if err != nil {
		s.log.Println(err)
	}
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
		case *Event_GameRequest:
			s.JoinGameHandler(ws, x)
		case *Event_GameRespond:
		case *Event_MoveElement:
			s.MoveElementHandler(ws, x, event.GameID, event.PlayerID)
		case *Event_SyncGameState:
		}
	}
}
