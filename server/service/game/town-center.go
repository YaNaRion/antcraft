package game

import (
	"main/service/math"
)

var TOWN_CENTER_SIZE_CONST = math.Vector2{
	X: 100.0,
	Y: 100.0,
}

type TownCenter struct {
	data ElementData
}

func NewTownCenter(pos math.Vector2, team int) *TownCenter {
	return &TownCenter{
		data: NewElementDate(pos, TOWN_CENTER_SIZE_CONST, team),
	}
}

func (tc *TownCenter) GetData() *ElementData {
	return &tc.data
}

func (tc *TownCenter) Update() {}

func (tc *TownCenter) SetNewTargetForUnitOut(newTarget math.Vector2) {
	directionVector := math.SubVec2(newTarget, tc.data.pos)
	directionVector.NormalizeVec()

	tc.data.directionVector = &directionVector
	tc.data.currentTarget = &newTarget
}

func (tc *TownCenter) CreateUnitFactory() *Worker {
	pos := math.Vector2{
		X: tc.data.pos.X + tc.data.size.X + 10,
		Y: tc.data.pos.Y + tc.data.size.Y + 10,
	}
	unit := NewWorker(pos, math.Vector2{
		X: 10,
		Y: 10,
	}, tc.data.team)

	unit.data.currentTarget = tc.data.currentTarget
	unit.data.directionVector = tc.data.directionVector
	return unit
}
