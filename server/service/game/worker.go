package game

import (
	"main/service/math"

	"github.com/google/uuid"
)

type Worker struct {
	data ElementData
}

func NewWorker(pos math.Vector2, size math.Vector2, team int) *Worker {
	return &Worker{
		data: ElementData{
			pos:           pos,
			size:          size,
			id:            ElementID(uuid.New().String()),
			currentTarget: nil,
			team:          team,
		},
	}
}

func (w *Worker) GetData() *ElementData {
	return &w.data
}

func (w *Worker) Update() {
}
