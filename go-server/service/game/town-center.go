package game

import (
	"main/service/math"

	"github.com/google/uuid"
)

var TOWN_CENTER_SIZE_CONST = math.Vector2{
	X: 100.0,
	Y: 100.0,
}

type TownCenter struct {
	pos             math.Vector2
	size            math.Vector2
	playerID        PlayerID
	unitID          ElementID
	currentTarget   *math.Vector2
	directionVector *math.Vector2
	team            int
}

func NewTownCenter(pos math.Vector2, team int) *TownCenter {
	return &TownCenter{
		pos:           pos,
		size:          TOWN_CENTER_SIZE_CONST,
		unitID:        ElementID(uuid.New().String()),
		currentTarget: nil,
		team:          team,
	}
}

func (u *TownCenter) GetPost() math.Vector2  { return u.pos }
func (u *TownCenter) SetPost(v math.Vector2) { u.pos = v }
func (u *TownCenter) GetID() ElementID       { return u.unitID }
func (u *TownCenter) GetSize() math.Vector2  { return u.size }
func (u *TownCenter) GetPlayerID() PlayerID  { return u.playerID }
func (u *TownCenter) GetCurrentTarget() *math.Vector2 {
	return u.currentTarget
}

func (u *TownCenter) GetTeam() int {
	return u.team
}

func (u *TownCenter) SetNewTargetForUnitOut(newTarget math.Vector2) {
	directionVector := math.SubVec2(newTarget, u.pos)
	directionVector.NormalizeVec()

	u.directionVector = &directionVector
	u.currentTarget = &newTarget
}

func (u *TownCenter) CreateUnitFactory() *Unit {
	pos := math.Vector2{
		X: u.pos.X + u.size.X + 10,
		Y: u.pos.Y + u.size.Y + 10,
	}
	unit := NewUnit(pos, math.Vector2{
		X: 10,
		Y: 10,
	}, u.team)

	unit.currentTarget = u.currentTarget
	unit.directionVector = u.directionVector
	return unit
}
