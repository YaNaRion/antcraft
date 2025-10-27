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
	SetNewTarget(math.Vector2)
	MoveElement(mapGrid [][]Tile) error
	GetCurrentObjective() *math.Vector2
	GetTeam() int
}
