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
	buildings []building.Building
	units     []unit.Unit
	ressource *ressource.RessourceMap
}

type Map struct {
	width   int
	heigh   int
	mapItem *MapItem
}

func NewMap() *Map {
	buildings := make([]building.Building, 0)
	units := make([]unit.Unit, 0)
	ressources := make([]ressource.RessourceMineral, 0)
	return &Map{
		width: MapHeight,
		heigh: MapWidth,
		mapItem: &MapItem{
			buildings: buildings,
			units:     units,
			ressource: ressource.NewRessourceMap(ressources),
		},
	}
}

func (m *Map) Draw() {
	m.mapItem.ressource.ClearEmptyRessource()
	for _, build := range m.mapItem.buildings {
		build.Draw()
	}

	for _, unit := range m.mapItem.units {
		unit.Draw()
	}

	for _, ressource := range m.mapItem.ressource.Ressources {
		ressource.Draw()
	}
}

func (m *Map) DefaultUnitMove() {
	for _, unit := range m.mapItem.units {
		// unit.
		unit.FindNextTarget(m.mapItem.ressource.Ressources)
		unit.MoveUnit(m.mapItem.ressource.Ressources)
	}
}

func (m *Map) RestMap() {
	m.mapItem.buildings = make([]building.Building, 0)
	m.mapItem.units = make([]unit.Unit, 0)
	m.mapItem.ressource.Ressources = make([]ressource.RessourceMineral, 0)
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
	m.mapItem.ressource.Ressources = append(
		m.mapItem.ressource.Ressources,
		ressource.NewDefaultFood(30, rl.Rectangle{X: 800, Y: 900}, rl.Yellow),
	)
	m.mapItem.ressource.Ressources = append(
		m.mapItem.ressource.Ressources,
		ressource.NewDefaultFood(30, rl.Rectangle{X: 700, Y: 800}, rl.Yellow),
	)
}

func (m *Map) GenerateNewWorker() {
	defaultUnit := unit.NewWorker(850, 900, 2, 2, m.mapItem.buildings[0].(*building.Base))
	m.mapItem.units = append(m.mapItem.units, defaultUnit)
}
