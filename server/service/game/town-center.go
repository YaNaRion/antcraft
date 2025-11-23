package game

import (
	"main/service/math"
)

var TOWN_CENTER_SIZE_CONST = math.Vector2{
	X: 100.0,
	Y: 100.0,
}

type TownCenter struct {
	element Element
}

func NewTownCenter(pos math.Vector2, team int) *TownCenter {
	return &TownCenter{
		element: Element{
			Data: NewElementDate(pos, TOWN_CENTER_SIZE_CONST, team),
			Stat: ElementStats{
				hitPoint:     1,
				meleeDamage:  1,
				rangeDamage:  1,
				attack_range: 1,
			},
		},
	}
}

func (tc *TownCenter) GetElement() *Element {
	return &tc.element
}

func (tc *TownCenter) Update() {}

func (tc *TownCenter) SetNewTargetForUnitOut(newTarget math.Vector2) {
	directionVector := math.SubVec2(newTarget, tc.element.Data.pos)
	directionVector.NormalizeVec()

	tc.element.Data.directionVector = &directionVector
	tc.element.Data.currentTarget = &newTarget
}

func (tc *TownCenter) CreateUnitFactory() *Worker {
	pos := math.Vector2{
		X: tc.element.Data.pos.X + tc.element.Data.size.X + 10,
		Y: tc.element.Data.pos.Y + tc.element.Data.size.Y + 10,
	}
	unit := NewWorker(pos, math.Vector2{
		X: 10,
		Y: 10,
	}, tc.element.Data.team)

	unit.element.Data.currentTarget = tc.element.Data.currentTarget
	unit.element.Data.directionVector = tc.element.Data.directionVector
	return unit
}
