package game

import (
	"client/scene/game/building"
	"client/scene/game/unit"
)

const (
	MapHeight = 10000
	MapWidth  = 10000
)

type Map struct {
	width     int
	heigh     int
	buildings []building.Building
	units     []unit.Unit
}

func NewMap() *Map {
	buildings := make([]building.Building, 0)
	units := make([]unit.Unit, 0)
	return &Map{
		width:     MapHeight,
		heigh:     MapWidth,
		buildings: buildings,
		units:     units,
	}
}

func (m *Map) Draw() {
	for _, build := range m.buildings {
		build.Draw()
	}

	for _, unit := range m.units {
		unit.Draw()
	}
}

func (m *Map) DefaultUnitMove() {
	for _, unit := range m.units {
		unit.MoveUnit()
	}
}

func (m *Map) PopulateDefaultMap() {
	defaulBuilding := building.NewBase(
		500,
		500,
		10,
		10,
	)
	m.buildings = append(m.buildings, defaulBuilding)
	defaultUnit := unit.NewWorker(600, 600, 2, 2)
	m.units = append(m.units, defaultUnit)
}
