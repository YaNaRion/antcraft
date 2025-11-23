package game

import (
	"main/service/math"

	"github.com/google/uuid"
)

type Worker struct {
	element Element
}

func NewWorker(pos math.Vector2, size math.Vector2, team int) *Worker {
	return &Worker{
		element: Element{
			Stat: ElementStats{
				hitPoint:     1,
				meleeDamage:  1,
				rangeDamage:  1,
				attack_range: 1,
			},
			Data: ElementData{
				pos:           pos,
				size:          size,
				id:            ElementID(uuid.New().String()),
				currentTarget: nil,
				team:          team,
			},
		},
	}
}
func (w *Worker) GetElement() *Element {
	return &w.element
}

func (w *Worker) GetData() *ElementData {
	return &w.element.Data
}

func (w *Worker) GetStat() *ElementStats {
	return &w.element.Stat
}

func (w *Worker) Update() {
}
