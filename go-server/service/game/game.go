package game

import (
	"log"
	"main/service/math"
)

type MapGrid struct {
	Grid     [][]Tile
	GridSpec math.Vector2
}

func NewMapGrid(width, height int) *MapGrid {
	grid := make([][]Tile, height)
	for i, row := range grid {
		grid[i] = make([]Tile, width)
		for j := range row {
			row[j] = NewTile()
		}
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
	Players      map[PlayerID]*Player
	GameState    Enum_Game_State
	gameElements map[ElementID]IElement
}

func NewGame(grid *MapGrid) *Game {
	return &Game{
		MapGrid:      grid,
		Players:      make(map[PlayerID]*Player, 0),
		GameState:    GAME_STATE_IN_GAME,
		gameElements: map[ElementID]IElement{},
	}
}

func (g *Game) GetIElementAt(vec *math.Vector2) Tile {
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

// LA FONCTION EST UNIQUEMENT UTILISE DANS UNE FONCTION QUI N'EST PAS UTILISE
func (g *Game) MoveElement(
	elementID ElementID,
	playerID PlayerID,
	newPos math.Vector2,
	oldPos math.Vector2,
) error {
	element := g.MapGrid.Grid[int(oldPos.X)][int(oldPos.Y)].Element
	if element == nil {
		log.Println("ELEMENT N'A PAS ETE MODIFIE")
		// TODO: Mettre une vraie erreur avec un system de gestion d'erreur plus avance
		return nil
	}

	element.SetPost(newPos)
	g.MapGrid.Grid[int(oldPos.X)][int(oldPos.Y)].Element = nil
	g.MapGrid.Grid[int(newPos.X)][int(newPos.Y)].Element = element
	return nil
}

func (g *Game) AddElement(el IElement) error {
	if g.MapGrid.Grid[int(el.GetPost().X)][int(el.GetPost().Y)].Element == nil {
		g.gameElements[el.GetID()] = el
		g.MapGrid.Grid[int(el.GetPost().X)][int(el.GetPost().Y)].Element = el
	}
	return nil
}

func (g *Game) GetElements() map[ElementID]IElement {
	return g.gameElements
}

func (g *Game) AddPlayer(player *Player) {
	g.Players[player.playerID] = player
}

func (g *Game) RemovePlayer(player *Player) {
	g.Players[player.playerID] = nil
}

func (g *Game) UpdateGameState() {
	for _, element := range g.gameElements {
		element.MoveElement(g.MapGrid.Grid)
	}
}
