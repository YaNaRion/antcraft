package unit

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Unit interface {
	Draw()
	MoveUnit()
}

type Worker struct {
	rec    rl.Rectangle
	color  rl.Color
	Width  int
	Height int
}

func NewWorker(x, y float32, width, height int) *Worker {
	return &Worker{
		rec:    rl.Rectangle{X: x, Y: y, Width: float32(width), Height: float32(height)},
		color:  rl.Pink,
		Width:  10,
		Height: 10,
	}
}

func (b *Worker) Draw() {
	rl.DrawRectangleRec(b.rec, b.color)
}

func (b *Worker) MoveUnit() {
	b.rec.X += 1.0
	b.rec.Y += 1.0
}
