package game

import (
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type IElement interface {
	GetPost() Vector2
	SetPost(Vector2)
}

type PlayerConn struct {
	ws *websocket.Conn
}

func NewPlayerConn(ws *websocket.Conn) *PlayerConn {
	return &PlayerConn{ws: ws}
}

type PlayerID string

type Player struct {
	playerID PlayerID
	elements []IElement
	conn     *PlayerConn
}

func NewPlayer(conn *PlayerConn, playerID PlayerID) *Player {
	return &Player{
		elements: make([]IElement, 0),
		conn:     conn,
		playerID: playerID,
	}
}

func (p *Player) GetPlayerIDString() string {
	return string(p.playerID)
}

type Vector2 struct {
	X int
	Y int
}

func NewVector2(x, y int) Vector2 {
	return Vector2{X: x, Y: y}
}

type ElementID uuid.UUID

type Unit struct {
	pos    Vector2
	size   Vector2
	unitID ElementID
}

func NewUnit(pos, size Vector2) *Unit {
	return &Unit{
		pos:    pos,
		size:   size,
		unitID: ElementID(uuid.New()),
	}
}

func (u *Unit) GetPost() Vector2  { return u.pos }
func (u *Unit) SetPost(v Vector2) { u.pos = v }

type MapGrid struct {
	Grid     [][]IElement
	GridSpec Vector2
}

func NewMapGrid(width, height int) *MapGrid {
	grid := make([][]IElement, height)
	for i := range grid {
		grid[i] = make([]IElement, width)
	}
	return &MapGrid{
		Grid:     grid,
		GridSpec: NewVector2(width, height),
	}
}

type Enum_Game_State int

const (
	GAME_STATE_UNKOWN = iota
	GAME_STATE_IN_LOBBY
	GAME_STATE_IS_STARTING
	GAME_STATE_IN_GAME
	GAME_STATE_END_GAME
)

type Game struct {
	grid     *MapGrid
	players  []*Player
	GameStat Enum_Game_State
}

func NewGame(grid *MapGrid) *Game {
	return &Game{
		grid:     grid,
		players:  make([]*Player, 0),
		GameStat: GAME_STATE_IN_GAME,
	}
}

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

func (g *GameManager) AddPlayerToGame(gameID GameID, player *Player) {
	g.games[gameID].players = append(g.games[gameID].players, player)
}

func (g *GameManager) GetPlayerInAGame(gameID GameID) []*Player {
	return g.games[gameID].players
}

func (g *GameManager) GetMap(gameID GameID) *MapGrid {
	return g.games[gameID].grid
}

func NewGameManager() *GameManager {
	gameManager := &GameManager{
		gamesID: make([]GameID, 0),
		games:   make(map[GameID]*Game),
	}

	// Game de setup pour accelerer le dev
	var defaultGameID GameID = "Game1"
	gameManager.gamesID = append(gameManager.gamesID, defaultGameID)
	var mapGrid *MapGrid = NewMapGrid(100, 100)
	gameManager.games[defaultGameID] = NewGame(mapGrid)

	return gameManager
}
