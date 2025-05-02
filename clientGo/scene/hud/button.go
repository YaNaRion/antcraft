package hud

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectangleButton struct {
	Rec     *rl.Rectangle
	color   *rl.Color
	message string
}

func (b *RectangleButton) Draw() {
	gui.Button(*b.Rec, b.message)
}

func newButton(
	x float32,
	y float32,
	width float32,
	height float32,
	color *rl.Color,
	message string,
) *RectangleButton {
	rec := rl.NewRectangle(
		x,
		y,
		width,
		height,
	)
	return &RectangleButton{
		Rec:     &rec,
		color:   color,
		message: message,
	}
}
