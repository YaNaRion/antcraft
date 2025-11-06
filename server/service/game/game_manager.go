package game

import (
	"main/service/math"
)

type GameID string

type GameManager struct {
	gamesID []GameID
	games   map[GameID]*Game
}

func (g *GameManager) GetGamesID() []GameID {
	return g.gamesID
}

func (g *GameManager) GetGame(gameID GameID) *Game {
	return g.games[gameID]
}

func (g *GameManager) AddTargetToElement(
	elementID ElementID,
	playerID PlayerID,
	gameID GameID,
	newTarget math.Vector2,
	currentPos math.Vector2,
) error {
	err := g.games[gameID].AddTargetToElement(elementID, playerID, newTarget, currentPos)
	if err != nil {
		return ErrNotUnitFound
	}
	return nil
}

// FONCTION NEST PAS ACTUELLEMENT UTILISE
// func (g *GameManager) MoveElement(
// 	elementID ElementID,
// 	playerID PlayerID,
// 	gameID GameID,
// 	newPos math.Vector2,
// 	oldPos math.Vector2,
// ) error {
// 	err := g.games[gameID].MoveElement(elementID, playerID, newPos, oldPos)
// 	if err != nil {
// 		log.Panicln("TROUVE PAS LELEMENT A DEPLACER")
// 		// Mettre une erreur
// 		return nil
// 	}
// 	return nil
// }

func (g *GameManager) AddPlayerToGame(gameID GameID, player *Player) {
	g.games[gameID].AddPlayer(player)
}

func (g *GameManager) GetPlayerInAGame(gameID GameID) map[PlayerID]*Player {
	return g.games[gameID].Players
}

func (g *GameManager) StartGame(gameID GameID) {
	g.games[gameID].StartGame()
}

func (g *GameManager) RemoveGame(gameID GameID) {
	g.games[gameID] = nil
}

func NewGameManager() *GameManager {
	gameManager := &GameManager{
		gamesID: make([]GameID, 0),
		games:   make(map[GameID]*Game),
	}

	// Game de setup pour accelerer le dev
	return gameManager
}

func (g *GameManager) PopulateWithDefaultGame() {
	var defaultGameID GameID = "Game1"
	g.gamesID = append(g.gamesID, defaultGameID)
	g.games[defaultGameID] = NewGame()
}
