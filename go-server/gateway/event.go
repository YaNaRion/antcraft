package gateway

import (
	"fmt"
	"main/service/game"
	"time"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

func (s *WebsocketManager) JoinGameHandler(ws *websocket.Conn, event *Event_JoinGame) {
	s.log.Println("JOIN GAME EVENT")

	var newPlayer *game.Player
	var playerID string

	// Faire un join game custom pour join la game voulu
	var gameID string = "Game1"

	playerID = fmt.Sprintf("Player%d", len(s.clients))
	newPlayer = game.NewPlayer(game.NewPlayerConn(ws), game.PlayerID(playerID))

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
	s.StartGame(ws, nil)
}

func (s *WebsocketManager) StartGame(ws *websocket.Conn, event *Event_StartGame) {
	s.log.Println("StartGame Event")
	gameID := event.StartGame.GameId
	s.gameManager.StartGame(game.GameID(gameID))
	go s.GameLoop(gameID)
}

func (s *WebsocketManager) GameLoop(gameID string) {
	s.log.Println("GAME LOOP")

	var err error
	gameMap := s.gameManager.GetGame(game.GameID(gameID))

	if gameMap.GameState != game.GAME_STATE_IN_GAME {
		return
	}

	var eventRespond Event

	mapGrid := NewEventSyncGameState(
		s.gameManager.GetMap(game.GameID(gameID)),
		GameState(gameMap.GameState),
	)

	s.log.Println(mapGrid.SyncGameState.Elements)

	// METTRE MUTEX OU DE QUOI
	eventRespond.DataEvent = mapGrid

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

	// Mise a jours a 60 HZ
	time.Sleep(16 * time.Millisecond)
}

func (s *WebsocketManager) MoveElementHandler(
	ws *websocket.Conn,
	unit *Event_MoveElement,
	gameID, playerID *string,
) {
	s.log.Println("MoveElement Event")
	elementNewPos := game.Vector2{
		X: int(unit.MoveElement.GetNewPos().X),
		Y: int(unit.MoveElement.GetNewPos().Y),
	}
	elementOldPos := game.Vector2{
		X: int(unit.MoveElement.GetOldPos().X),
		Y: int(unit.MoveElement.GetOldPos().Y),
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
