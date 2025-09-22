package gateway

import "main/service/game"

type Worker struct{ pos Vector2 }

// Mettre une sorte de mutex ou sem
func (w *Worker) GetPost() Vector2  { return w.pos }
func (w *Worker) SetPost(v Vector2) { w.pos = v }

func MapGridToProto(grid *game.MapGrid, state GameState) *SyncGameState {
	rows := make([]*Row, len(grid.Grid))

	for y, row := range grid.Grid {
		values := make([]*Element, len(row))
		for x, elem := range row {
			values[x] = iElementToProto(elem)
		}
		rows[y] = &Row{Values: values}
	}

	return &SyncGameState{
		GameState: state,
		Rows:      rows,
	}
}

func iElementToProto(elem game.IElement) *Element {
	if elem == nil {
		return nil
	}

	pos := elem.GetPost()
	var et Element_ElementType
	et = Element_WORKER

	return &Element{
		Element: &MoveElement{
			Pos: &Vector2{
				XPos: int32(pos.X),
				YPos: int32(pos.Y),
			},
		},
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
