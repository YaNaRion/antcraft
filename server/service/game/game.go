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

// func (g *Game) MoveElement(
// 	elementID ElementID,
// 	playerID PlayerID,
// 	newPos math.Vector2,
// 	oldPos math.Vector2,
// ) error {
// 	element := g.MapGrid.Grid[int(oldPos.X)][int(oldPos.Y)].Element
// 	if element == nil {
// 		// TODO: Mettre une vraie erreur avec un system de gestion d'erreur plus avance
// 		return ErrNotUnitFound
// 	}
//
// 	element.SetPost(newPos)
// 	g.MapGrid.Grid[int(oldPos.X)][int(oldPos.Y)].Element = nil
// 	g.MapGrid.Grid[int(newPos.X)][int(newPos.Y)].Element = element
// 	return nil
// }

// TODO: RETOUR DERREUR
func (g *Game) AddElement(el IElement) error {
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
	for _, element := range g.gameElements {
		// if _, ok := element.(IUnit); ok {
		err := g.MoveUnit(element)
		if err == ErrCannotMoveUnitFurder {
			log.Println(err)
		}
		// }
	}
}

func (g *Game) IsNextPositionWalkable(el *ElementData, moveX, moveY float64) bool {
	newX := el.GetPost().X + moveX
	newY := el.GetPost().Y + moveY
	elSize := el.GetSize()

	// Define new bounding box
	elLeft := newX
	elRight := newX + elSize.X
	elTop := newY
	elBottom := newY + elSize.Y

	log.Printf(
		"EL LEFT: %d, EL RIGHT: %d, EL TOP: %d, EL BOT: %d\n",
		int(elLeft),
		int(elRight),
		int(elTop),
		int(elBottom),
	)

	for _, other := range g.gameElements {
		// Skip self
		if other.GetData().GetID() == el.GetID() {
			continue
		}

		otherPos := other.GetData().GetPost()
		otherSize := other.GetData().GetSize()

		// Define other bounding box
		otherLeft := otherPos.X
		otherRight := otherPos.X + otherSize.X
		otherTop := otherPos.Y
		otherBottom := otherPos.Y + otherSize.Y

		log.Printf(
			"other LEFT: %d, other RIGHT: %d, other TOP: %d, other BOT: %d\n",
			int(otherLeft),
			int(otherRight),
			int(otherTop),
			int(otherBottom),
		)

		if elLeft < otherRight && elLeft > otherLeft {
			log.Println()
			return false
		}

		// Check if rectangles overlap
		if elRight < otherLeft && elLeft > otherRight &&
			elBottom > otherTop && elTop < otherBottom {
			log.Println()
			return false // collision detected
		}
	}

	return true // no collision
}

// Plus utilisé logique tranférer à la struct Game, garde pour trace
func (g *Game) MoveUnit(el IElement) error {
	if el.GetData().GetCurrentObjective() == nil {
		return ErrUnitHasNoObjective
	}

	// if !el.canUnitMove(mapGrid) {
	// 	el.directionVector.X = 0
	// 	el.directionVector.Y = 0
	// 	return nil
	// }

	const speed = 1.0
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

	log.Println("DANS NEXT POSTION NOT WALKABLE")

	el.GetData().SetNewTarget(nil)
	// return error collision mettre direction vector nil et currentobjective nil
	return ErrCannotMoveUnitFurder
}
