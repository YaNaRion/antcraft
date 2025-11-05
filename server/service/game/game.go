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
	for x := range width {
		grid[x] = make([]Tile, height)
		for y := range height {
			grid[x][y] = NewTile()
		}
	}

	return &MapGrid{
		Grid:     grid,
		GridSpec: math.NewVector2(width, height),
	}
}

func (m *MapGrid) ToggleCollisionTiles(el IElement, IsWalkable bool) {
	for i := el.GetPost().X; i < el.GetPost().X+el.GetSize().X; i++ {
		for j := el.GetPost().Y; j < el.GetPost().Y+el.GetSize().Y; j++ {
			m.Grid[j][i].IsWalkable = IsWalkable
		}
	}
}

func (m *MapGrid) AddElement(el IElement) error {
	if m.Grid[int(el.GetPost().X)][int(el.GetPost().Y)].Element == nil {
		m.Grid[int(el.GetPost().X)][int(el.GetPost().Y)].Element = el
		// m.ToggleCollisionTiles(el, false)
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
		return ErrNotUnitFound
	}

	// _, ok := element.(IUnit)
	// if ok {
	// log.Println("DANS SET NEW TARGET FOR UNIT")
	element.SetNewTarget(&newTarget)
	// }
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
		// if _, ok := element.(IUnit); ok {
		err := g.MoveUnit(element)
		if err == ErrCannotMoveUnitFurder {
			log.Println(err)
		}
		// }
	}
}

<<<<<<< HEAD:server/service/game/game.go
func (g *Game) IsNextPositionWalkable(el IElement, moveX, moveY float64) bool {
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
		if other.GetID() == el.GetID() {
			continue
		}

		otherPos := other.GetPost()
		otherSize := other.GetSize()

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

func (m *MapGrid) IsNextPositionWalkable(el IElement, moveX, moveY float64) bool {
	return m.Grid[int(el.GetPost().X+moveX)][int(el.GetPost().Y+moveY)].IsWalkable
=======
func (m *MapGrid) IsNextPostionwalkable(el IElement, moveX, moveY float64) bool {
	return m.Grid[int(float64(el.GetPost().X)+moveX)][int(float64(el.GetPost().Y)+moveY)].IsWalkable
>>>>>>> 64fdb68e278de4d7faac8aae7e5988bd60756730:go-server/service/game/game.go
}

// Plus utilisé logique tranférer à la struct Game, garde pour trace
func (g *Game) MoveUnit(el IElement) error {
	if el.GetCurrentObjective() == nil {
		return ErrUnitHasNoObjective
	}

	// if !el.canUnitMove(mapGrid) {
	// 	el.directionVector.X = 0
	// 	el.directionVector.Y = 0
	// 	return nil
	// }

	const speed = 1.0
	moveX := float64(el.GetDirectionVector().X) * speed
	moveY := float64(el.GetDirectionVector().Y) * speed

	toTargetX := float64(el.GetCurrentObjective().X - el.GetPost().X)
	toTargetY := float64(el.GetCurrentObjective().Y - el.GetPost().Y)

	if (moveX*toTargetX + moveY*toTargetY) <= 0 {
		moveX = toTargetX
		moveY = toTargetY
	}

<<<<<<< HEAD:server/service/game/game.go
	if (moveX*toTarget.X + moveY*toTarget.Y) <= 0 {
		moveX = toTarget.X
		moveY = toTarget.Y
	}

	if g.IsNextPositionWalkable(el, moveX, moveY) {
		// g.MapGrid.ToggleCollisionTiles(el, true)
		el.UpdatePos(g.MapGrid.Grid, moveX, moveY)
		// g.MapGrid.ToggleCollisionTiles(el, false)
=======
	if g.MapGrid.IsNextPostionwalkable(el, moveX, moveY) {
		log.Println("DANS IF WALK")
		g.MapGrid.ToggleCollisionTiles(el, true)
		el.UpdatePos(g.MapGrid.Grid, int(moveX), int(moveY))
		g.MapGrid.ToggleCollisionTiles(el, false)
>>>>>>> 64fdb68e278de4d7faac8aae7e5988bd60756730:go-server/service/game/game.go
		return nil
	}

	log.Println("DANS NEXT POSTION NOT WALKABLE")

	el.SetNewTarget(nil)
	// return error collision mettre direction vector nil et currentobjective nil

	log.Println(el.GetPost().X + int(moveX))
	log.Println(el.GetPost().Y + int(moveY))
	log.Println(g.MapGrid.Grid[el.GetPost().X+int(moveX)][el.GetPost().Y+int(moveY)])
	return ErrCannotMoveUnitFurder
}
