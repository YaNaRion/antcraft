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
	grid := make([][]Tile, width)
	for i := range width {
		grid = append(grid, []Tile{})
		for range height {
			grid[i] = append(grid[i], NewTile())
		}
	}

	return &MapGrid{
		Grid:     grid,
		GridSpec: math.NewVector2(float64(width), float64(height)),
	}
}

func (m *MapGrid) ToggleCollisionTiles(el IElement, IsWalkable bool) {
	for i := int(el.GetPost().X); i < int(el.GetPost().X+el.GetSize().X); i++ {
		for j := int(el.GetPost().Y); j < int(el.GetPost().Y+el.GetSize().Y); j++ {
			m.Grid[i][j].IsWalkable = IsWalkable
		}
	}
}

func (m *MapGrid) AddElement(el IElement) error {
	if m.Grid[int(el.GetPost().X)][int(el.GetPost().Y)].Element == nil {
		m.Grid[int(el.GetPost().X)][int(el.GetPost().Y)].Element = el
		m.ToggleCollisionTiles(el, false)
	}
	return nil
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
func (g *Game) StartGame() {}

func (g *Game) GetIElementAt(vec *math.Vector2) *Tile {
	return &g.MapGrid.Grid[int(vec.X)][int(vec.Y)]
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
	if element == nil {
		log.Println("ELEMENT NEST PAS TROUVE")
		return nil
	}
	_, ok := element.(IUnit)
	if ok {
		element.SetNewTarget(&newTarget)
		g.MoveUnit(element)
	}
	return nil
}

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

// TODO: RETOUR DERREUR
func (g *Game) AddElement(el IElement) error {
	err := g.MapGrid.AddElement(el)
	if err != nil {
		return nil
	}

	g.gameElements[el.GetID()] = el
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
		if _, ok := element.(IUnit); ok {
			g.MoveUnit(element)
		}
	}
}

func (m *MapGrid) CheckNextPosition(el IElement, moveX, moveY float64) bool {
	return m.Grid[int(el.GetPost().X+moveX)][int(el.GetPost().Y+moveY)].IsWalkable
}

// Plus utilisé logique tranférer à la struct Game, garde pour trace
func (g *Game) MoveUnit(el IElement) error {
	if el.GetCurrentObjective() == nil {
		return nil
	}

	// if !el.canUnitMove(mapGrid) {
	// 	el.directionVector.X = 0
	// 	el.directionVector.Y = 0
	// 	return nil
	// }

	const speed = 1.0
	moveX := el.GetDirectionVector().X * speed
	moveY := el.GetDirectionVector().Y * speed

	toTarget := math.Vector2{
		X: el.GetCurrentObjective().X - el.GetPost().X,
		Y: el.GetCurrentObjective().Y - el.GetPost().Y,
	}

	if (moveX*toTarget.X + moveY*toTarget.Y) <= 0 {
		moveX = toTarget.X
		moveY = toTarget.Y
	}

	if g.MapGrid.CheckNextPosition(el, moveX, moveY) {
		g.MapGrid.ToggleCollisionTiles(el, true)
		el.UpdatePos(g.MapGrid.Grid, moveX, moveY)
		g.MapGrid.ToggleCollisionTiles(el, false)
		return nil
	}

	el.SetNewTarget(nil)
	// return error collision mettre direction vector nil et currentobjective nil
	return nil
}
