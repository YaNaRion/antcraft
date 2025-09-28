package game

import (
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

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

func (p *Player) AddElement(el IElement) {
	p.elements = append(p.elements, el)
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

type ElementID string

type IElement interface {
	GetPost() Vector2
	GetSize() Vector2
	GetPlayerID() PlayerID
	SetPost(Vector2)
	GetID() ElementID
}

type Unit struct {
	pos      Vector2
	size     Vector2
	playerID PlayerID
	unitID   ElementID
}

func NewUnit(pos, size Vector2) *Unit {
	return &Unit{
		pos:    pos,
		size:   size,
		unitID: ElementID(uuid.New().String()),
	}
}

func (u *Unit) GetPost() Vector2      { return u.pos }
func (u *Unit) SetPost(v Vector2)     { u.pos = v }
func (u *Unit) GetID() ElementID      { return u.unitID }
func (u *Unit) GetSize() Vector2      { return u.size }
func (u *Unit) GetPlayerID() PlayerID { return u.playerID }

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
	MapGrid      *MapGrid
	players      map[PlayerID]*Player
	GameState    Enum_Game_State
	gameElements map[ElementID]IElement
}

func NewGame(grid *MapGrid) *Game {
	return &Game{
		MapGrid:      grid,
		players:      make(map[PlayerID]*Player, 0),
		GameState:    GAME_STATE_IN_GAME,
		gameElements: map[ElementID]IElement{},
	}
}

func (g *Game) GetIElementAt(vec *Vector2) IElement {
	return g.MapGrid.Grid[vec.X][vec.Y]
}

func (g *Game) MoveElement(
	elementID ElementID,
	playerID PlayerID,
	newPos Vector2,
	oldPos Vector2,
) error {
	element := g.MapGrid.Grid[oldPos.X][oldPos.Y]
	if element == nil {
		// Faire une erreur
		return nil
	}
	element.SetPost(newPos)
	g.MapGrid.Grid[oldPos.X][oldPos.Y] = nil
	g.MapGrid.Grid[newPos.X][newPos.X] = element

	return nil
}

func (g *Game) AddElement(el IElement) error {
	if g.MapGrid.Grid[el.GetPost().X][el.GetPost().Y] == nil {
		g.gameElements[el.GetID()] = el
		g.MapGrid.Grid[el.GetPost().X][el.GetPost().Y] = el

	}
	return nil
}

func (g *Game) GetElements() map[ElementID]IElement {
	return g.gameElements
}

func (g *Game) AddPlayer(player *Player) {
	g.players[player.playerID] = player
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

func (g *GameManager) MoveElement(
	elementID ElementID,
	playerID PlayerID,
	gameID GameID,
	newPos Vector2,
	oldPos Vector2,
) error {
	err := g.games[gameID].MoveElement(elementID, playerID, newPos, oldPos)
	if err != nil {
		// Mettre une erreur
		return nil
	}
	return nil
}

func (g *GameManager) AddPlayerToGame(gameID GameID, player *Player) {
	g.games[gameID].AddPlayer(player)
}

func (g *GameManager) GetPlayerInAGame(gameID GameID) map[PlayerID]*Player {
	return g.games[gameID].players
}

func (g *GameManager) GetMap(gameID GameID) *MapGrid {
	return g.games[gameID].MapGrid
}

func (g *GameManager) StartGame(gameID GameID) {
	g.games[gameID].GameState = GAME_STATE_IN_GAME
}

func NewGameManager() *GameManager {
	gameManager := &GameManager{
		gamesID: make([]GameID, 0),
		games:   make(map[GameID]*Game),
	}

	// Game de setup pour accelerer le dev
	var defaultGameID GameID = "Game1"
	gameManager.gamesID = append(gameManager.gamesID, defaultGameID)
	var mapGrid *MapGrid = NewMapGrid(1000, 1000)
	gameManager.games[defaultGameID] = NewGame(mapGrid)

	return gameManager
}
