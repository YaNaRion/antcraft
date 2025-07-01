package unit

import (
	"client/scene/game/building"
	"client/scene/game/ressource"
	"errors"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Target interface {
	GetRec() rl.Rectangle
}

type Unit interface {
	Draw()
	MoveUnit([]ressource.RessourceMineral)
	FindNextTarget(ressources []ressource.RessourceMineral)
}

type WorkerStatus int

var mapStatusToRLColor = map[WorkerStatus]rl.Color{
	CARRYING_RESSOURCE: rl.Blue,
	GOING_TO_RESSOURCE: rl.Pink,
	IDLE:               rl.White,
}

const (
	CARRYING_RESSOURCE WorkerStatus = iota
	GOING_TO_RESSOURCE
	IDLE
)

type Worker struct {
	currentTarget           Target
	closedRessource         ressource.RessourceMineral
	Base                    *building.Base
	status                  WorkerStatus
	distanceClosedRessource float64
	rec                     rl.Rectangle
	color                   rl.Color
	Width                   int
	Height                  int
}

func NewWorker(x, y float32, width, height int, base *building.Base) *Worker {
	return &Worker{
		rec: rl.Rectangle{
			X:      x,
			Y:      y,
			Width:  float32(width),
			Height: float32(height),
		},
		color:                   rl.Pink,
		status:                  IDLE,
		Width:                   10,
		Height:                  10,
		distanceClosedRessource: 1000000,
		Base:                    base,
	}
}

func (w *Worker) Draw() {
	rl.DrawRectangleRec(w.rec, mapStatusToRLColor[w.status])
}

func (w *Worker) FindNextRessource(ressources []ressource.RessourceMineral) error {
	if len(ressources) <= 0 {
		w.closedRessource = nil
		return error_no_target_found
	}

	for _, ressource := range ressources {
		totalDistance := math.Abs(float64(ressource.GetRec().X - w.rec.X))
		totalDistance += math.Abs(float64(ressource.GetRec().Y - w.rec.Y))
		if totalDistance < w.distanceClosedRessource {
			w.closedRessource = ressource
			w.currentTarget = ressource
			w.distanceClosedRessource = totalDistance
		}
	}

	w.status = GOING_TO_RESSOURCE
	return nil
}

func (w *Worker) handlerGadderRessource() {
	w.status = CARRYING_RESSOURCE
	w.closedRessource.Consume()
	w.closedRessource = nil
	w.currentTarget = w.Base
}

func (w *Worker) handlerReturnRessourceToBase(ressources []ressource.RessourceMineral) {
	w.status = IDLE
	w.distanceClosedRessource = 100000
	_ = w.FindNextRessource(ressources)
}

func (w *Worker) FindNextTarget(ressources []ressource.RessourceMineral) {
	if w.status != CARRYING_RESSOURCE {
		err := w.FindNextRessource(ressources)
		if errors.Is(err, error_no_target_found) {
			w.status = IDLE
		}
	}

	// Lorsque le worker est sur la ressource
	if w.closedRessource != nil {
		if w.closedRessource.GetRec().X == w.rec.X && w.closedRessource.GetRec().Y == w.rec.Y &&
			w.status != CARRYING_RESSOURCE {
			w.handlerGadderRessource()
		}
		return
	}

	// Lorsque le worker retourne a la base
	if w.Base != nil && w.status == CARRYING_RESSOURCE {
		if w.Base.GetRec().X == w.rec.X && w.Base.GetRec().Y == w.rec.Y &&
			w.status == CARRYING_RESSOURCE {
			w.handlerReturnRessourceToBase(ressources)
		}
	}
}

func (w *Worker) MoveUnit(ressources []ressource.RessourceMineral) {
	if w.status != CARRYING_RESSOURCE {
		err := w.FindNextRessource(ressources)
		if errors.Is(err, error_no_target_found) {
			w.closedRessource = nil
			w.currentTarget = nil
			return
		}
	}

	if w.currentTarget != nil {
		if w.currentTarget.GetRec().X > w.rec.X {
			w.rec.X += 1.0
		} else if w.currentTarget.GetRec().X < w.rec.X {
			w.rec.X -= 1.0
		}

		if w.currentTarget.GetRec().Y > w.rec.Y {
			w.rec.Y += 1.0
		} else if w.currentTarget.GetRec().Y < w.rec.Y {
			w.rec.Y -= 1.0
		}
	}
}
