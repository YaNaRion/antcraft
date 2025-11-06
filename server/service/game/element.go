package game

import (
	"main/service/math"

	"github.com/google/uuid"
)

type ElementData struct {
	pos             math.Vector2
	size            math.Vector2
	playerID        PlayerID
	id              ElementID
	currentTarget   *math.Vector2
	directionVector *math.Vector2
	team            int
}

type ElementID string

type IElement interface {
	GetData() *ElementData
	Update()
}

type IUnit interface {
	GetData() *ElementData
	// MoveElement() error
}

type IBuilding interface {
	GetData() *ElementData
	// CreateUnitFactory() *IUnit
	// SetNewTargetForUnitOut(newTarget math.Vector2)
}

func NewElementDate(pos math.Vector2, size math.Vector2, team int) ElementData {
	return ElementData{
		pos:           pos,
		size:          size,
		id:            ElementID(uuid.New().String()),
		currentTarget: nil,
		team:          team,
	}
}

func (u *ElementData) GetPost() math.Vector2  { return u.pos }
func (u *ElementData) SetPost(v math.Vector2) { u.pos = v }
func (u *ElementData) GetID() ElementID       { return u.id }
func (u *ElementData) GetSize() math.Vector2  { return u.size }
func (u *ElementData) GetPlayerID() PlayerID  { return u.playerID }
func (u *ElementData) GetCurrentObjective() *math.Vector2 {
	return u.currentTarget
}
func (u *ElementData) GetDirectionVector() *math.Vector2 {
	return u.directionVector
}
func (u *ElementData) GetTeam() int {
	return u.team
}

func (u *ElementData) SetNewTarget(newTarget *math.Vector2) {
	u.currentTarget = newTarget

	if newTarget == nil {
		u.directionVector = nil
		return
	}

	targetVector := math.SubVec2(*newTarget, u.pos)

	directionVector := math.NewVector2(targetVector.X, targetVector.Y)
	directionVector.NormalizeVec()
	u.directionVector = &directionVector
}

// TODO Mettre collision d'unite
// Il u a une approximation de la position lors du cast de float a int, ne devrait pas etre grave avec des tres petites tiles comme ici
// Note pour potentiel future bug
func (u *ElementData) UpdatePos(x, y float64) {
	u.pos.X += x
	u.pos.Y += y
	u.SetNewTarget(u.currentTarget)
}
