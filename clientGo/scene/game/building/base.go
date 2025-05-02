package building

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Building interface {
	Draw()
}

type Base struct {
	Position rl.Vector2
	rec      rl.Rectangle
	Width    int
	Height   int
}

func NewBase(x, y float32, width, height int) *Base {
	return &Base{
		Position: rl.Vector2{X: x, Y: y},
		rec:      rl.Rectangle{X: x, Y: y, Width: 10, Height: 10},
		Width:    10,
		Height:   10,
	}
}

func (b *Base) Draw() {
	rl.DrawRectangleRec(b.rec, rl.Red)
	rl.DrawText("JE SUIS UNE BASE", 15+b.rec.ToInt32().X, b.rec.ToInt32().Y, 10, rl.Red)
}
