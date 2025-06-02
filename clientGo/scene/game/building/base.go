package building

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Base struct {
	Position rl.Vector2
	rec      rl.Rectangle
	Width    int
	Height   int
}

func NewBase(x, y float32, width, height int) *Base {
	return &Base{
		Position: rl.Vector2{X: x, Y: y},
		rec:      rl.Rectangle{X: x, Y: y, Width: float32(width), Height: float32(height)},
		Width:    10,
		Height:   10,
	}
}

func (b *Base) Draw() {
	rl.DrawRectangleRec(b.rec, rl.Red)
}

func (b *Base) GetRec() rl.Rectangle {
	return b.rec
}
