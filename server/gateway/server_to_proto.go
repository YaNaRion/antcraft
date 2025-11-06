package gateway

import (
	"main/service/game"
)

type Worker struct{ pos Vector2 }

// Mettre une sorte de mutex ou sem
func (w *Worker) GetPost() Vector2  { return w.pos }
func (w *Worker) SetPost(v Vector2) { w.pos = v }

func NewMapGridToSyncGameState(game *game.Game, state GameState) *SyncGameState {
	var element []*Element

	for _, el := range game.GetElements() {
		element = append(element, iElementToProto(el))

	}

	return &SyncGameState{
		GameState: state,
		Elements:  element,
	}
}

func NewEventSyncGameState(game *game.Game, state GameState) *Event_SyncGameState {
	return &Event_SyncGameState{
		SyncGameState: NewMapGridToSyncGameState(game, state),
	}
}

func iElementToProto(el game.IElement) *Element {
	pos := el.GetData().GetPost()
	var et ElementType

	switch el.(type) {
	case *game.Worker:
		et = ElementType_WORKER
	case *game.TownCenter:
		et = ElementType_BASE
	}

	currentObjective := &Vector2{
		X: -1,
		Y: -1,
	}

	if el.GetData().GetCurrentObjective() != nil {
		currentObjective = &Vector2{
			X: int32(el.GetData().GetCurrentObjective().X),
			Y: int32(el.GetData().GetCurrentObjective().Y),
		}
	}

	return &Element{
		Pos: &Vector2{
			X: int32(pos.X),
			Y: int32(pos.Y),
		},
		Size: &Vector2{
			X: int32(el.GetData().GetSize().X),
			Y: int32(el.GetData().GetSize().Y),
		},
		CurrentObjective: currentObjective,

		PlayerId:    string(el.GetData().GetPlayerID()),
		UnitId:      string(el.GetData().GetID()),
		ElementType: et,
		Team:        ColorTeam(el.GetData().GetTeam()),
	}
}
