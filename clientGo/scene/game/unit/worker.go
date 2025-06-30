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

// type WorkerStatus int
//
// const (
// 	CARRYING_RESSOURCE = iota
// 	LOOKING_FOR_RESSOURCE
// )

type Worker struct {
	currentTarget           Target
	closedRessource         ressource.RessourceMineral
	Base                    *building.Base
	isCarryingRessource     bool
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
		isCarryingRessource:     false,
		Width:                   10,
		Height:                  10,
		distanceClosedRessource: 1000000,
		Base:                    base,
	}
}

func (b *Worker) Draw() {
	rl.DrawRectangleRec(b.rec, b.color)
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
	return nil
}

func (w *Worker) handlerGadderRessource() {
	w.isCarryingRessource = true
	w.closedRessource.Consume()
	w.closedRessource = nil
	w.currentTarget = w.Base
	w.color = rl.Blue

}

func (w *Worker) handlerReturnToBase(ressources []ressource.RessourceMineral) {
	w.color = rl.Pink
	w.isCarryingRessource = false
	w.distanceClosedRessource = 100000
	err := w.FindNextRessource(ressources)
	if errors.Is(err, error_no_target_found) {
		w.color = rl.White
	}
}

func (w *Worker) FindNextTarget(ressources []ressource.RessourceMineral) {
	if !w.isCarryingRessource {
		err := w.FindNextRessource(ressources)
		if errors.Is(err, error_no_target_found) {
			w.color = rl.White
		}
	}

	if w.closedRessource != nil {
		if w.closedRessource.GetRec().X == w.rec.X && w.closedRessource.GetRec().Y == w.rec.Y &&
			!w.isCarryingRessource {
			w.handlerGadderRessource()
		}
	} else {
		if w.Base != nil && w.isCarryingRessource {
			if w.Base.GetRec().X == w.rec.X && w.Base.GetRec().Y == w.rec.Y && w.isCarryingRessource {
				w.handlerReturnToBase(ressources)
			}
		}
	}

	// if len(ressources) <= 0 {
	// 	// Message d'erreur lorsqu"il n'y a plus de ressources, mais en sorte que l'affichage soit dynamique avec Width et Height
	// 	// rl.DrawText("IL NY A AUCUNE RESSOURCE", 500, 10, 12, rl.White)
	// }
}

func (w *Worker) MoveUnit(ressources []ressource.RessourceMineral) {
	if !w.isCarryingRessource {
		err := w.FindNextRessource(ressources)
		if errors.Is(err, error_no_target_found) {
			w.closedRessource = nil
			w.currentTarget = nil
			w.color = rl.White
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
