package game

import (
	"log"
	"main/service/math"
)

type MapGrid struct {
	Grid     [][]IElement
	GridSpec math.Vector2
}

func NewMapGrid(width, height int) *MapGrid {
	grid := make([][]IElement, height)
	for i := range grid {
		grid[i] = make([]IElement, width)
	}
	return &MapGrid{
		Grid:     grid,
		GridSpec: math.NewVector2(float64(width), float64(height)),
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

func (g *Game) GetIElementAt(vec *math.Vector2) IElement {
	return g.MapGrid.Grid[int(vec.X)][int(vec.Y)]
}

// TODO AJOUTER UNE GESTION DERREUR SI ON NE TROUVE PAS DELEMENT
func (g *Game) AddTargetToElement(
	elementID ElementID,
	playerID PlayerID,
	newTarget math.Vector2,
	currentPosition math.Vector2,
) error {
	// element := g.MapGrid.Grid[int(currentPosition.X)][int(currentPosition.Y)]
	element := g.gameElements[elementID]
	if element != nil {
		element.SetNewTarget(newTarget)
		return nil
	}
	log.Panic("ELEMENT NEST PAS TROUVE")
	return nil
}

func (g *Game) MoveElement(
	elementID ElementID,
	playerID PlayerID,
	newPos math.Vector2,
	oldPos math.Vector2,
) error {
	element := g.MapGrid.Grid[int(oldPos.X)][int(oldPos.Y)]
	if element == nil {
		log.Println("ELEMENT N'A PAS ETE MODIFIE")
		return nil
	}

	element.SetPost(newPos)
	g.MapGrid.Grid[int(oldPos.X)][int(oldPos.Y)] = nil
	g.MapGrid.Grid[int(newPos.X)][int(newPos.Y)] = element
	return nil
}

func (g *Game) AddElement(el IElement) error {
	if g.MapGrid.Grid[int(el.GetPost().X)][int(el.GetPost().Y)] == nil {
		g.gameElements[el.GetID()] = el
		g.MapGrid.Grid[int(el.GetPost().X)][int(el.GetPost().Y)] = el
	}
	return nil
}

func (g *Game) GetElements() map[ElementID]IElement {
	return g.gameElements
}

func (g *Game) AddPlayer(player *Player) {
	g.players[player.playerID] = player
}

func (g *Game) UpdateGameState() {
	for _, element := range g.gameElements {
		element.MoveElement(g.MapGrid.Grid)
	}
}
