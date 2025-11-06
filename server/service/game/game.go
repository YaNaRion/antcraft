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
}

func NewGame() *Game {
	return &Game{
		Players:       make(map[PlayerID]*Player, 0),
		GameState:     GAME_STATE_IN_GAME,
		gameElements:  map[ElementID]IElement{},
		gameUnits:     map[ElementID]IUnit{},
		gameBuildings: map[ElementID]IBuilding{},
	}
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

	element.GetData().SetNewTarget(&newTarget)
	return nil
}

func (g *Game) AddElement(el IElement) error {
	g.gameElements[el.GetData().GetID()] = el
	return nil
}

func (g *Game) AddUnit(el IElement, u IUnit) error {
	g.gameUnits[el.GetData().GetID()] = u
	g.gameElements[el.GetData().GetID()] = el
	return nil
}

func (g *Game) AddBuilding(el IElement, b IBuilding) error {
	g.gameBuildings[el.GetData().GetID()] = b
	g.gameElements[el.GetData().GetID()] = el
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
	elSize := el.GetSize()

	elLeft := newX
	elRight := newX + elSize.X
	elTop := newY
	elBottom := newY + elSize.Y

	for _, other := range g.gameElements {
		if other.GetData().GetID() == el.GetID() {
			continue
		}

		otherPos := other.GetData().GetPost()
		otherSize := other.GetData().GetSize()

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
	if el.GetData().GetCurrentObjective() == nil {
		return ErrUnitHasNoObjective
	}

	const speed = 2.0
	moveX := el.GetData().directionVector.X * speed
	moveY := el.GetData().directionVector.Y * speed

	toTargetX := el.GetData().GetCurrentObjective().X - el.GetData().GetPost().X
	toTargetY := el.GetData().GetCurrentObjective().Y - el.GetData().GetPost().Y

	if (moveX*toTargetX + moveY*toTargetY) <= 0 {
		moveX = toTargetX
		moveY = toTargetY
	}

	if (moveX*toTargetX + moveY*toTargetY) <= 0 {
		moveX = toTargetX
		moveY = toTargetY
		el.GetData().SetNewTarget(nil)
	}

	if g.IsNextPositionWalkable(el.GetData(), moveX, moveY) {
		el.GetData().UpdatePos(moveX, moveY)
		return nil
	}

	el.GetData().SetNewTarget(nil)
	// return error collision mettre direction vector nil et currentobjective nil
	return ErrCannotMoveUnitFurder
}
