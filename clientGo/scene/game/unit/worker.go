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
	MoveUnit()
	FindNextTarget(ressources []ressource.RessourceMineral)
	GetRec() rl.Rectangle
	GetStatus() WorkerStatus
}

type Worker struct {
	currentTarget           Target
	closedRessource         ressource.RessourceMineral
	Base                    *building.Base
	status                  WorkerStatus
	distanceClosedRessource float64
	rec                     rl.Rectangle
	Width                   int
	Height                  int
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

const BIG_DISTANCE = 1000000

func NewWorker(x, y float32, width, height int, base *building.Base) *Worker {
	return &Worker{
		rec: rl.Rectangle{
			X:      x,
			Y:      y,
			Width:  float32(width),
			Height: float32(height),
		},
		status:                  IDLE,
		Width:                   10,
		Height:                  10,
		distanceClosedRessource: BIG_DISTANCE,
		Base:                    base,
	}
}

func (w *Worker) Draw() {
	rl.DrawRectangleRec(w.rec, mapStatusToRLColor[w.status])
}

func (w *Worker) GetRec() rl.Rectangle {
	return w.rec
}

func (w *Worker) GetStatus() WorkerStatus {
	return w.status
}

func (w *Worker) FindNextRessource(ressources []ressource.RessourceMineral) error {
	if len(ressources) <= 0 {
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
	w.distanceClosedRessource = BIG_DISTANCE
	_ = w.FindNextRessource(ressources)
}

func (w *Worker) FindNextTarget(ressources []ressource.RessourceMineral) {
	if w.status != CARRYING_RESSOURCE {
		err := w.FindNextRessource(ressources)
		if errors.Is(err, error_no_target_found) {
			w.status = IDLE
			w.closedRessource = nil
			w.currentTarget = w.Base
			w.distanceClosedRessource = BIG_DISTANCE
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

func (w *Worker) MoveUnit() {
	const moveMultiplier = 2
	if w.currentTarget != nil {
		if w.currentTarget.GetRec().X > w.rec.X {
			w.rec.X += 1.0 * moveMultiplier
		} else if w.currentTarget.GetRec().X < w.rec.X {
			w.rec.X -= 1.0 * moveMultiplier
		}

		if w.currentTarget.GetRec().Y > w.rec.Y {
			w.rec.Y += 1.0 * moveMultiplier
		} else if w.currentTarget.GetRec().Y < w.rec.Y {
			w.rec.Y -= 1.0 * moveMultiplier
		}
	}
}
