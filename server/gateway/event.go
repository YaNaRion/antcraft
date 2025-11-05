package gateway

import (
	"fmt"
	"main/service/game"
	"main/service/math"
	// "math/rand"
	"time"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

func (s *WebsocketManager) JoinGameHandler(
	ws *websocket.Conn,
	event *Event_JoinGame,
	gameID string,
) {
	s.log.Println("JOIN GAME EVENT")

	gameMap := s.gameManager.GetGame(game.GameID(gameID))
	if gameMap == nil {
		s.gameManager.PopulateWithDefaultGame()
		gameMap = s.gameManager.GetGame(game.GameID(gameID))
	}
	var newPlayer *game.Player

	// Faire un join game custom pour join la game voulu
	playerID := fmt.Sprintf("Player%d", len(s.clients))
	newPlayer = game.NewPlayer(game.NewPlayerConn(ws), game.PlayerID(playerID))

	// TEAM 1 est rouge TEAM 2 est bleu
	unit := game.NewUnit(math.Vector2{
		X: float64(100),
		Y: float64(10),
	}, math.Vector2{
		X: 10,
		Y: 10,
	},
		len(gameMap.Players)+1,
	)

	// unit2 := game.NewUnit(math.Vector2{
	// 	X: float64(rand.Int() % 500),
	// 	Y: float64(rand.Int() % 500),
	// }, math.Vector2{
	// 	X: 10,
	// 	Y: 10,
	// },
	// 	len(gameMap.Players)+1,
	// )

	cc := game.NewTownCenter(math.Vector2{
		X: float64(100),
		Y: float64(200),
	},
		len(gameMap.Players)+1,
	)

	newPlayer.AddElement(unit)
	// newPlayer.AddElement(unit2)

	newPlayer.AddElement(cc)

	err := gameMap.AddElement(unit)
	if err != nil {
		s.log.Println(err)
	}

	// err = gameMap.AddElement(unit2)
	// if err != nil {
	// 	s.log.Println(err)
	// }
	//
	err = gameMap.AddElement(cc)
	if err != nil {
		s.log.Println(err)
	}

	var eventRespond Event
	eventRespond.GameId = gameID
	var teamColor ColorTeam
	if len(gameMap.Players) == 0 {
		teamColor = ColorTeam_RED_PROTO
	} else if len(gameMap.Players) == 1 {
		teamColor = ColorTeam_BLUE_PROTO
	} else {
		s.log.Println("MAX 2 PLAYER")
		return
	}

	s.gameManager.AddPlayerToGame(game.GameID(gameID), newPlayer)

	eventRespond.PlayerInfo = &Player{
		Color:    &teamColor,
		PlayerId: playerID,
	}
	eventRespond.DataEvent = &Event_JoinGame{}

	marshalData, err := proto.Marshal(&eventRespond)
	if err != nil {
		s.log.Println("Error occured when marsharl game data")
	}

	err = websocket.Message.Send(ws, marshalData)
	if err != nil {
		s.log.Println("Error occured when sending game info")
	}

	// if len(gameMap.Players) == 2 {
	// 	s.log.Println("START GAME")
	s.StartGame(ws, game.GameID(gameID))
	// 	s.log.Printf("GAME STARTED")
	// }

}

func (s *WebsocketManager) StartGame(ws *websocket.Conn, gameID game.GameID) {
	s.log.Println("StartGame Event")
	s.gameManager.StartGame(game.GameID(gameID))
	go s.GameLoop(gameID)
}

func (s *WebsocketManager) GameLoop(gameID game.GameID) {
	s.log.Printf("GAME LOOP FOR GAME: %s", gameID)
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

		// METTRE MUTEX OU DE QUOI
		eventRespond.DataEvent = mapGrid

		marshalData, err := proto.Marshal(&eventRespond)
		if err != nil {
			s.log.Println("send error:", err)
		}

		countNilClient := 0
		for _, client := range s.clients {
			err = websocket.Message.Send(client.Conn, marshalData)
			if err != nil {
				s.log.Println("send error:", err)
				s.gameManager.RemoveGame(gameID)
				s.CleanClient()
				return
			}
		}

		if len(s.clients) == countNilClient {
			return
		}

		// Mise a jours a 60 HZ, 16, pas vraiment vrai car cest 16ms apres le fin dexecution de la boucle INF3610 style
		time.Sleep(16 * time.Millisecond)
	}
}

func (s *WebsocketManager) MoveElementHandler(
	ws *websocket.Conn,
	unit *Event_MoveElement,
	gameID, playerID *string,
) {
	s.log.Println("MoveElement Event")
	// s.log.Printf("MoveElementFrom: %s", ws.RemoteAddr())
	// s.log.Printf("PlayerID: %s", *playerID)
	// s.log.Printf("Moving: %s", unit.MoveElement.GetUnitId())

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

func (s *WebsocketManager) StartGameHandler(
	ws *websocket.Conn,
	gameID game.GameID,
) {
	s.gameManager.StartGame(game.GameID(gameID))
}
