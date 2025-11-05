package game

import (
	"main/service/math"

	"github.com/google/uuid"
)

func NewUnit(pos, size math.Vector2, team int) *Unit {
	return &Unit{
		pos:           pos,
		size:          size,
		id:            ElementID(uuid.New().String()),
		currentTarget: nil,
		team:          team,
	}
}

func (u *Unit) GetPost() math.Vector2  { return u.pos }
func (u *Unit) SetPost(v math.Vector2) { u.pos = v }
func (u *Unit) GetID() ElementID       { return u.id }
func (u *Unit) GetSize() math.Vector2  { return u.size }
func (u *Unit) GetPlayerID() PlayerID  { return u.playerID }
func (u *Unit) GetCurrentObjective() *math.Vector2 {
	return u.currentTarget
}
func (u *Unit) GetDirectionVector() *math.Vector2 {
	return u.directionVector
}
func (u *Unit) GetTeam() int {
	return u.team
}

func (u *Unit) SetNewTarget(newTarget *math.Vector2) {
	u.currentTarget = newTarget

	if newTarget == nil {
		u.directionVector = nil
		return
	}
	directionVector := math.SubVec2(*newTarget, u.pos)

	directionVector.NormalizeVec()
	u.directionVector = &directionVector
}

// TODO Mettre collision d'unite
// Il u a une approximation de la position lors du cast de float a int, ne devrait pas etre grave avec des tres petites tiles comme ici
// Note pour potentiel future bug
func (u *Unit) UpdatePos(grid [][]Tile, x, y int) {
	grid[int(u.pos.X)][int(u.pos.Y)].Element = nil
	u.pos.X += x
	u.pos.Y += y
	grid[int(u.pos.X)][int(u.pos.Y)].Element = u
}

// // fonction que regarde si l'unit peut se deplacer a sa prochaine destination
// func (u *Unit) canUnitMove(mapGrid [][]Tile) bool {
// 	return true
// }
//
// // Avec des murs, il va falloir mettre en place du pathing pour que l'unité arrive à se rendre à sa destination
// func (u *Unit) MoveElement(mapGrid [][]Tile) error {
// 	if u.currentTarget == nil {
// 		return nil
// 	}
//
// 	// if !u.canUnitMove(mapGrid) {
// 	// 	u.directionVector.X = 0
// 	// 	u.directionVector.Y = 0
// 	// 	return nil
// 	// }
//
// 	const speed = 1.0
// 	moveX := u.directionVector.X * speed
// 	moveY := u.directionVector.Y * speed
//
// 	toTarget := math.Vector2{
// 		X: u.currentTarget.X - u.pos.X,
// 		Y: u.currentTarget.Y - u.pos.Y,
// 	}
//
// 	if (moveX*toTarget.X + moveY*toTarget.Y) <= 0 {
// 		moveX = toTarget.X
// 		moveY = toTarget.Y
// 	}
//
// 	u.UpdatePos(mapGrid, moveX, moveY)
// 	return nil
// }
