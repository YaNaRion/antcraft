package game

import (
	"github.com/google/uuid"
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
	MoveElement(mapGrid [][]Tile)
	GetCurrentObjective() *math.Vector2
	GetTeam() int
}

type Unit struct {
	pos             math.Vector2
	size            math.Vector2
	playerID        PlayerID
	unitID          ElementID
	currentTarget   *math.Vector2
	directionVector *math.Vector2
	team            int
}

func NewUnit(pos, size math.Vector2, team int) *Unit {
	return &Unit{
		pos:           pos,
		size:          size,
		unitID:        ElementID(uuid.New().String()),
		currentTarget: nil,
		team:          team,
	}
}

func (u *Unit) GetPost() math.Vector2  { return u.pos }
func (u *Unit) SetPost(v math.Vector2) { u.pos = v }
func (u *Unit) GetID() ElementID       { return u.unitID }
func (u *Unit) GetSize() math.Vector2  { return u.size }
func (u *Unit) GetPlayerID() PlayerID  { return u.playerID }
func (u *Unit) GetCurrentObjective() *math.Vector2 {
	return u.currentTarget
}
func (u *Unit) GetTeam() int {
	return u.team
}

func (u *Unit) SetNewTarget(newTarget math.Vector2) {
	u.currentTarget = &newTarget

	directionVector := math.SubVec2(newTarget, u.pos)

	directionVector.NormalizeVec()
	u.directionVector = &directionVector
}

// TODO Mettre collision d'unite
func (u *Unit) UpdatePos(grid [][]Tile, x, y float64) {
	grid[int(u.pos.X)][int(u.pos.Y)].Element = nil
	u.pos.X += x
	u.pos.Y += y
	grid[int(u.pos.X)][int(u.pos.Y)].Element = u
}

func (u *Unit) MoveElement(mapGrid [][]Tile) {
	if u.currentTarget == nil {
		return
	}

	const speed = 1.0
	moveX := u.directionVector.X * speed
	moveY := u.directionVector.Y * speed

	toTarget := math.Vector2{
		X: u.currentTarget.X - u.pos.X,
		Y: u.currentTarget.Y - u.pos.Y,
	}

	if (moveX*toTarget.X + moveY*toTarget.Y) <= 0 {
		moveX = toTarget.X
		moveY = toTarget.Y
	}

	u.UpdatePos(mapGrid, moveX, moveY)
}
