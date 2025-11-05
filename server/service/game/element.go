package game

import (
	"main/service/math"
)

type ElementID string

type IElement interface {
	GetPost() math.Vector2
	GetSize() math.Vector2
	GetPlayerID() PlayerID
	SetPost(math.Vector2)
	GetID() ElementID
	SetNewTarget(*math.Vector2)
	GetCurrentObjective() *math.Vector2
	GetTeam() int

	GetDirectionVector() *math.Vector2
	UpdatePos(grid [][]Tile, x, y int)
}

type IUnit interface {
	MoveElement(mapGrid [][]Tile) error
}

type IBuilding interface {
	CreateUnitFactory() *Unit
	SetNewTargetForUnitOut(newTarget math.Vector2)
}

type Unit struct {
	pos             math.Vector2
	size            math.Vector2
	playerID        PlayerID
	id              ElementID
	currentTarget   *math.Vector2
	directionVector *math.Vector2
	team            int
}

type TownCenter struct {
	pos             math.Vector2
	size            math.Vector2
	playerID        PlayerID
	id              ElementID
	currentTarget   *math.Vector2
	directionVector *math.Vector2
	team            int
}
