package gateway

import (
	"main/service/game"
)

type Worker struct{ pos Vector2 }

// Mettre une sorte de mutex ou sem
func (w *Worker) GetPost() Vector2  { return w.pos }
func (w *Worker) SetPost(v Vector2) { w.pos = v }

func NewMapGridToSyncGameState(grid *game.MapGrid, state GameState) *SyncGameState {
	var element []*Element

	for _, row := range grid.Grid {
		for _, elem := range row {
			if elem != nil {
				element = append(element, iElementToProto(elem))
			}
		}
	}
	return &SyncGameState{
		GameState: state,
		Elements:  element,
	}
}

func NewEventSyncGameState(grid *game.MapGrid, state GameState) *Event_SyncGameState {
	return &Event_SyncGameState{
		SyncGameState: NewMapGridToSyncGameState(grid, state),
	}
}

func iElementToProto(elem game.IElement) *Element {
	if elem == nil {
		return nil
	}

	pos := elem.GetPost()
	et := ElementType_WORKER

	currentObjective := &Vector2{
		X: -1,
		Y: -1,
	}

	if elem.GetCurrentObjective() != nil {
		currentObjective = &Vector2{
			X: int32(elem.GetCurrentObjective().X),
			Y: int32(elem.GetCurrentObjective().Y),
		}
	}

	return &Element{
		Pos: &Vector2{
			X: int32(pos.X),
			Y: int32(pos.Y),
		},
		Size: &Vector2{
			X: int32(elem.GetSize().X),
			Y: int32(elem.GetSize().Y),
		},
		CurrentObjective: currentObjective,

		PlayerId:    string(elem.GetPlayerID()),
		UnitId:      string(elem.GetID()),
		ElementType: et,
		Team:        ColorTeam(elem.GetTeam()),
	}
}
