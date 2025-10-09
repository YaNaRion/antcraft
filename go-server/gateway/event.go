package gateway

import (
	"fmt"
	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
	"main/service/game"
	"main/service/math"
	"math/rand"
	"time"
)

func (s *WebsocketManager) JoinGameHandler(
	ws *websocket.Conn,
	event *Event_JoinGame,
	gameID string,
) {
	s.log.Println("JOIN GAME EVENT")

	var newPlayer *game.Player
	var playerID string

	// Faire un join game custom pour join la game voulu
	playerID = fmt.Sprintf("Player%d", len(s.clients))
	newPlayer = game.NewPlayer(game.NewPlayerConn(ws), game.PlayerID(playerID))

	var unit *game.Unit = game.NewUnit(math.Vector2{
		X: float64(rand.Int() % 500),
		Y: float64(rand.Int() % 500),
	}, math.Vector2{
		X: 10,
		Y: 10,
	})

	var unit2 *game.Unit = game.NewUnit(math.Vector2{
		X: float64(rand.Int() % 500),
		Y: float64(rand.Int() % 500),
	}, math.Vector2{
		X: 10,
		Y: 10,
	})

	newPlayer.AddElement(unit)
	newPlayer.AddElement(unit2)

	gameMap := s.gameManager.GetGame(game.GameID(gameID))

	err := gameMap.AddElement(unit)
	if err != nil {
		s.log.Println(err)
	}

	err = gameMap.AddElement(unit2)
	if err != nil {
		s.log.Println(err)
	}

	s.gameManager.AddPlayerToGame(game.GameID(gameID), newPlayer)
	s.StartGame(ws, game.GameID(gameID))
}

func (s *WebsocketManager) StartGame(ws *websocket.Conn, gameID game.GameID) {
	s.log.Println("StartGame Event")
	s.gameManager.StartGame(game.GameID(gameID))
	go s.GameLoop(gameID)
}

func (s *WebsocketManager) GameLoop(gameID game.GameID) {
	s.log.Println("GAME LOOP")
	gameMap := s.gameManager.GetGame(gameID)
	for {
		var err error
		if gameMap.GameState != game.GAME_STATE_IN_GAME {
			return
		}
		gameMap.UpdateGameState()

		var eventRespond Event
		mapGrid := NewEventSyncGameState(
			s.gameManager.GetMap(gameID),
			GameState(gameMap.GameState),
		)

		// s.log.Println(mapGrid.SyncGameState.Elements)
		// METTRE MUTEX OU DE QUOI
		eventRespond.DataEvent = mapGrid

		marshalData, err := proto.Marshal(&eventRespond)
		if err != nil {
			s.log.Println("send error:", err)
		}

		// log.Printf("Nombre client: %d", len(s.clients))
		countNilClient := 0
		for _, client := range s.clients {
			err = websocket.Message.Send(client.Conn, marshalData)
			if err != nil {
				s.log.Println("send error:", err)
				s.gameManager.RemoveGame(gameID)
				return
			}
		}

		if len(s.clients) == countNilClient {
			return
		}

		// Mise a jours a 60 HZ, 16
		time.Sleep(16 * time.Millisecond)
	}
}

func (s *WebsocketManager) MoveElementHandler(
	ws *websocket.Conn,
	unit *Event_MoveElement,
	gameID, playerID *string,
) {
	s.log.Println("MoveElement Event")
	s.log.Printf("MoveElementFrom: %s", ws.RemoteAddr())
	s.log.Printf("PlayerID: %s", *playerID)
	s.log.Printf("Moving: %s", unit.MoveElement.GetUnitId())

	elementNewPos := math.Vector2{
		X: float64(unit.MoveElement.GetNewPos().X),
		Y: float64(unit.MoveElement.GetNewPos().Y),
	}

	elementOldPos := math.Vector2{
		X: float64(unit.MoveElement.GetOldPos().X),
		Y: float64(unit.MoveElement.GetOldPos().Y),
	}

	err := s.gameManager.AddTargetToElement(
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
