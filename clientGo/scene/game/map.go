package game

import (
	"client/scene/game/building"
	"client/scene/game/ressource"
	"client/scene/game/unit"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MapHeight = 10000
	MapWidth  = 10000
)

type MapItem struct {
	ressource *ressource.RessourceMap
	hives     []*Hive
}

type Map struct {
	width   int
	heigh   int
	mapItem *MapItem
}

func NewMap() *Map {
	hives := make([]*Hive, 0)
	ressources := make([]ressource.RessourceMineral, 0)
	return &Map{
		width: MapHeight,
		heigh: MapWidth,
		mapItem: &MapItem{
			hives:     hives,
			ressource: ressource.NewRessourceMap(ressources),
		},
	}
}

func (m *Map) Draw() {
	m.mapItem.ressource.ClearEmptyRessource()

	for _, hive := range m.mapItem.hives {
		hive.Draw()
	}

	for _, ressource := range m.mapItem.ressource.Ressources {
		ressource.Draw()
	}
}

func (m *Map) DefaultUnitMove() {
	for _, hive := range m.mapItem.hives {
		for _, unit := range hive.units {
			// unit.
			unit.FindNextTarget(m.mapItem.ressource.Ressources)
			unit.MoveUnit(m.mapItem.ressource.Ressources)
		}
	}
}

func (m *Map) RestMap() {
	m.mapItem.hives = make([]*Hive, 0)
	m.mapItem.ressource.Ressources = make([]ressource.RessourceMineral, 0)
}

func (m *Map) PopulateDefaultMap() {
	hive := newHive()
	m.mapItem.hives = append(m.mapItem.hives, hive)

	defaulBuilding := building.NewBase(
		500,
		500,
		10,
		10,
	)
	hive.buildings = append(hive.buildings, defaulBuilding)

	m.mapItem.ressource.Ressources = append(
		m.mapItem.ressource.Ressources,
		ressource.NewDefaultFood(30, rl.Rectangle{X: 800, Y: 900}, rl.Yellow),
	)

	m.mapItem.ressource.Ressources = append(
		m.mapItem.ressource.Ressources,
		ressource.NewDefaultFood(30, rl.Rectangle{X: 700, Y: 800}, rl.Yellow),
	)

	go func() {
		time.Sleep(20 * time.Second)
		m.mapItem.ressource.Ressources = append(
			m.mapItem.ressource.Ressources,
			ressource.NewDefaultFood(30, rl.Rectangle{X: 1200, Y: 800}, rl.Yellow),
		)
	}()
}

func (m *Map) GenerateNewWorker() {
	defaultUnit := unit.NewWorker(850, 900, 2, 2, m.mapItem.hives[0].buildings[0].(*building.Base))
	m.mapItem.hives[0].units = append(m.mapItem.hives[0].units, defaultUnit)
}
