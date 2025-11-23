package game

import (
	"log"
	"main/service/math"
)

type Enum_Game_State int

const (
	GAME_STATE_UNKOWN = iota
	GAME_STATE_IN_LOBBY
	GAME_STATE_IS_STARTING
	GAME_STATE_IN_GAME
	GAME_STATE_END_GAME
)

type Game struct {
	Players       map[PlayerID]*Player
	GameState     Enum_Game_State
	gameElements  map[ElementID]IElement
	gameUnits     map[ElementID]IUnit
	gameBuildings map[ElementID]IBuilding
	// gameGrid      [][]Tile
}

func NewGame() *Game {
	return &Game{
		Players:       make(map[PlayerID]*Player, 0),
		GameState:     GAME_STATE_IN_GAME,
		gameElements:  map[ElementID]IElement{},
		gameUnits:     map[ElementID]IUnit{},
		gameBuildings: map[ElementID]IBuilding{},
		// gameGrid:      PopulationDefaultGameGrid(),
	}
}

func PopulationDefaultGameGrid() [][]Tile {
	width := 1500
	height := 1500
	grid := make([][]Tile, width)
	for x := range width {
		grid[x] = make([]Tile, height)
		for y := range height {
			grid[x][y] = NewTile()
		}
	}
	return grid
}

func (g *Game) StartGame() {
	g.GameState = GAME_STATE_IN_GAME
}

func (g *Game) AddTargetToElement(
	elementID ElementID,
	playerID PlayerID,
	newTarget math.Vector2,
	currentPosition math.Vector2,
) error {
	element := g.gameElements[elementID]
	if element == nil {
		log.Println("ELEMENT NEST PAS TROUVE")
		return ErrNotUnitFound
	}

	element.GetElement().Data.SetNewTarget(&newTarget)
	return nil
}

func (g *Game) AddElement(el IElement) error {
	g.gameElements[el.GetElement().Data.GetID()] = el
	return nil
}

func (g *Game) AddUnit(el IElement, u IUnit) error {
	g.gameUnits[el.GetElement().Data.GetID()] = u
	g.gameElements[el.GetElement().Data.GetID()] = el
	return nil
}

func (g *Game) AddBuilding(el IElement, b IBuilding) error {
	g.gameBuildings[el.GetElement().Data.GetID()] = b
	g.gameElements[el.GetElement().Data.GetID()] = el
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
	for _, unit := range g.gameUnits {
		err := g.MoveUnit(unit)
		if err == ErrCannotMoveUnitFurder {
			log.Println(err)
		}
	}
}

func (g *Game) IsNextPositionWalkable(el *ElementData, moveX, moveY float64) bool {
	newX := el.GetPost().X + moveX
	newY := el.GetPost().Y + moveY

	// if !g.gameGrid[int(newX)][int(newY)].IsWalkable {
	// 	return false
	// }

	elSize := el.GetSize()

	elLeft := newX
	elRight := newX + elSize.X
	elTop := newY
	elBottom := newY + elSize.Y

	for _, other := range g.gameElements {
		if other.GetElement().Data.GetID() == el.GetID() {
			continue
		}

		otherPos := other.GetElement().Data.GetPost()
		otherSize := other.GetElement().Data.GetSize()

		otherLeft := otherPos.X
		otherRight := otherPos.X + otherSize.X
		otherTop := otherPos.Y
		otherBottom := otherPos.Y + otherSize.Y

		if elRight > otherLeft && elLeft < otherRight &&
			elBottom > otherTop && elTop < otherBottom {
			return false // collision
		}

	}
	return true // no collision
}

// Plus utilisé logique tranférer à la struct Game, garde pour trace
func (g *Game) MoveUnit(el IUnit) error {
	if el.GetElement().Data.GetCurrentObjective() == nil {
		return ErrUnitHasNoObjective
	}

	const speed = 2.0
	moveX := el.GetElement().Data.directionVector.X * speed
	moveY := el.GetElement().Data.directionVector.Y * speed

	toTargetX := el.GetElement().Data.GetCurrentObjective().X - el.GetElement().Data.GetPost().X
	toTargetY := el.GetElement().Data.GetCurrentObjective().Y - el.GetElement().Data.GetPost().Y

	if (moveX*toTargetX + moveY*toTargetY) <= 0 {
		moveX = toTargetX
		moveY = toTargetY
	}

	if (moveX*toTargetX + moveY*toTargetY) <= 0 {
		moveX = toTargetX
		moveY = toTargetY
		el.GetElement().Data.SetNewTarget(nil)
	}

	if g.IsNextPositionWalkable(&(el.GetElement().Data), moveX, moveY) {
		el.GetElement().Data.UpdatePos(moveX, moveY)
		return nil
	}

	el.GetElement().Data.SetNewTarget(nil)
	// return error collision mettre direction vector nil et currentobjective nil
	return ErrCannotMoveUnitFurder
}
