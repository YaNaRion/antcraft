package game

import (
	"client/scene/game/building"
	"client/scene/game/ressource"
	"client/scene/game/unit"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MapHeight = 10000
	MapWidth  = 10000
)

type MapItem struct {
	buildings  []building.Building
	units      []unit.Unit
	ressources []ressource.Ressource
}

type Map struct {
	width   int
	heigh   int
	mapItem *MapItem
}

func NewMap() *Map {
	buildings := make([]building.Building, 0)
	units := make([]unit.Unit, 0)
	ressources := make([]ressource.Ressource, 0)
	return &Map{
		width: MapHeight,
		heigh: MapWidth,
		mapItem: &MapItem{
			buildings:  buildings,
			units:      units,
			ressources: ressources,
		},
	}
}

func (m *Map) Draw() {
	for _, build := range m.mapItem.buildings {
		build.Draw()
	}

	for _, unit := range m.mapItem.units {
		unit.Draw()
	}

	for _, ressource := range m.mapItem.ressources {
		ressource.Draw()
	}
}

func (m *Map) DefaultUnitMove() {
	for _, unit := range m.mapItem.units {
		// unit.
		unit.FindNextTarget(m.mapItem.ressources)
		unit.MoveUnit()
	}
}

func (m *Map) RestMap() {
	m.mapItem.buildings = make([]building.Building, 0)
	m.mapItem.units = make([]unit.Unit, 0)
	m.mapItem.ressources = make([]ressource.Ressource, 0)
}

func (m *Map) PopulateDefaultMap() {
	defaulBuilding := building.NewBase(
		500,
		500,
		10,
		10,
	)
	m.mapItem.buildings = append(m.mapItem.buildings, defaulBuilding)
	defaultUnit := unit.NewWorker(850, 900, 2, 2, defaulBuilding)
	m.mapItem.units = append(m.mapItem.units, defaultUnit)
	m.mapItem.ressources = append(
		m.mapItem.ressources,
		ressource.NewFood(30, rl.Rectangle{X: 800, Y: 900}, rl.Yellow),
	)
	m.mapItem.ressources = append(
		m.mapItem.ressources,
		ressource.NewFood(30, rl.Rectangle{X: 700, Y: 800}, rl.Yellow),
	)
}
