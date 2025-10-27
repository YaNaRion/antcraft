package game

import "main/service/math"

type Building interface {
	CreateUnitFactory() *Unit
	SetNewTargetForUnitOut(newTarget math.Vector2)
}
