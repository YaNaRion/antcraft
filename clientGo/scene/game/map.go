package game

import (
	"client/scene/game/building"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MapHeight = 10000
	MapWidth  = 10000
)

type Map struct {
	width     int
	heigh     int
	buildings []building.Building
}

func NewMap() *Map {
	defaulBuilding := building.NewBase(
		500,
		500,
		10,
		10,
	)
	buildings := make([]building.Building, 0)
	buildings = append(buildings, defaulBuilding)
	return &Map{
		width:     MapHeight,
		heigh:     MapWidth,
		buildings: buildings,
	}
}

func (m *Map) Draw() {
	// rl.DrawRectangleLines(0, 0, MapWidth, MapHeight, rl.White)
	rl.DrawRectangle(0, 0, MapWidth, MapHeight, rl.White)
	for _, build := range m.buildings {
		build.Draw()
	}
}
