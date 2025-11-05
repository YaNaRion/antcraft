package game

import (
	"client/scene/game/building"
	"client/scene/game/ressource"
	"client/scene/game/unit"
)

type Hive struct {
	buildings []building.Building
	units     []unit.Unit
}

func newHive() *Hive {
	buildings := make([]building.Building, 0)
	units := make([]unit.Unit, 0)
	return &Hive{
		buildings: buildings,
		units:     units,
	}
}

func (h *Hive) Draw() {
	for _, build := range h.buildings {
		build.Draw()
	}

	for _, unit := range h.units {
		unit.Draw()
	}
}

func (h *Hive) GetUnits() []unit.Unit {
	return h.units
}

func (h *Hive) MoveHiveUnit(ressources []ressource.RessourceMineral) {
	for _, unitIns := range h.units {
		workerStatus := unitIns.GetStatus()
		unitIns.FindNextTarget(ressources)
		unitIns.MoveUnit()
		if workerStatus == unit.IDLE && workerStatus != unitIns.GetStatus() {
			break
		}
	}
}
