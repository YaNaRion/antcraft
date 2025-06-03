package unit

import (
	"client/scene/game/building"
	"client/scene/game/ressource"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Target interface {
	GetRec() rl.Rectangle
}

type Unit interface {
	Draw()
	MoveUnit()
	FindNextTarget(ressources []ressource.Ressource)
}

type Worker struct {
	currentTarget           Target
	closedRessource         ressource.Ressource
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

func (w *Worker) FindNextRessource(ressources []ressource.Ressource) {
	for _, ressource := range ressources {
		totalDistance := math.Abs(float64(ressource.GetRec().X - w.rec.X))
		totalDistance += math.Abs(float64(ressource.GetRec().Y - w.rec.Y))
		fmt.Println("RECHERCHE DE RESSOURCE")
		fmt.Println(ressource.GetRec())
		fmt.Println(w.rec)
		fmt.Println(totalDistance)
		if totalDistance < w.distanceClosedRessource {
			fmt.Println("NOUS SOMMES DANS UNE RESSOURCE TROUVER")
			fmt.Println(ressource.GetRec().X)
			w.closedRessource = ressource
			w.currentTarget = ressource
			w.distanceClosedRessource = totalDistance
		}
	}
}

func (w *Worker) FindNextTarget(ressources []ressource.Ressource) {
	if !w.isCarryingRessource {
		w.FindNextRessource(ressources)
	}

	if w.closedRessource != nil {
		if w.closedRessource.GetRec().X == w.rec.X && w.closedRessource.GetRec().Y == w.rec.Y &&
			!w.isCarryingRessource {
			w.isCarryingRessource = true
			w.closedRessource = nil
			w.currentTarget = w.Base
			w.color = rl.Blue
		}
	} else {
		if w.Base != nil && w.isCarryingRessource {
			if w.Base.GetRec().X == w.rec.X && w.Base.GetRec().Y == w.rec.Y &&
				w.isCarryingRessource {
				w.isCarryingRessource = false
				w.distanceClosedRessource = 100000
				w.color = rl.Pink
				w.FindNextRessource(ressources)
			}
		}
	}

	if len(ressources) <= 0 {
		// Message d'erreur lorsqu"il n'y a plus de ressources, mais en sorte que l'affichage soit dynamique avec Width et Height
		rl.DrawText("IL NY A AUCUNE RESSOURCE", 500, 10, 12, rl.White)
	}
}

func (b *Worker) MoveUnit() {
	if b.currentTarget != nil {
		if b.currentTarget.GetRec().X > b.rec.X {
			b.rec.X += 1.0
		} else if b.currentTarget.GetRec().X < b.rec.X {
			b.rec.X -= 1.0
		}

		if b.currentTarget.GetRec().Y > b.rec.Y {
			b.rec.Y += 1.0
		} else if b.currentTarget.GetRec().Y < b.rec.Y {
			b.rec.Y -= 1.0
		}
	}
}
