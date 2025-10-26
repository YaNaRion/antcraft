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
			if elem.Element != nil {
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

func iElementToProto(tile game.Tile) *Element {
	if tile.Element == nil {
		return nil
	}

	pos := tile.Element.GetPost()
	et := ElementType_WORKER

	currentObjective := &Vector2{
		X: -1,
		Y: -1,
	}

	if tile.Element.GetCurrentObjective() != nil {
		currentObjective = &Vector2{
			X: int32(tile.Element.GetCurrentObjective().X),
			Y: int32(tile.Element.GetCurrentObjective().Y),
		}
	}

	return &Element{
		Pos: &Vector2{
			X: int32(pos.X),
			Y: int32(pos.Y),
		},
		Size: &Vector2{
			X: int32(tile.Element.GetSize().X),
			Y: int32(tile.Element.GetSize().Y),
		},
		CurrentObjective: currentObjective,

		PlayerId:    string(tile.Element.GetPlayerID()),
		UnitId:      string(tile.Element.GetID()),
		ElementType: et,
		Team:        ColorTeam(tile.Element.GetTeam()),
	}
}
