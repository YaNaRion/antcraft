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

	return &Element{
		Pos: &Vector2{
			X: int32(pos.X),
			Y: int32(pos.Y),
		},
		Size: &Vector2{
			X: int32(elem.GetSize().X),
			Y: int32(elem.GetSize().Y),
		},
		PlayerId:    string(elem.GetPlayerID()),
		UnitId:      string(elem.GetID()),
		ElementType: et,
	}
}

// type Marine struct{ pos Vector2 }
//
// func (m *Marine) GetPost() Vector2  { return m.pos }
// func (m *Marine) SetPost(v Vector2) { m.pos = v }
//
// type Base struct{ pos Vector2 }
//
// func (b *Base) GetPost() Vector2  { return b.pos }
// func (b *Base) SetPost(v Vector2) { b.pos = v }
//
// type Barrack struct{ pos Vector2 }
//
// func (br *Barrack) GetPost() Vector2  { return br.pos }
// func (br *Barrack) SetPost(v Vector2) { br.pos = v }
//
// type Supply struct{ pos Vector2 }
//
// func (s *Supply) GetPost() Vector2  { return s.pos }
// func (s *Supply) SetPost(v Vector2) { s.pos = v }
//
// type Mineral struct{ pos Vector2 }
//
// func (m *Mineral) GetPost() Vector2  { return m.pos }
// func (m *Mineral) SetPost(v Vector2) { m.pos = v }
