package hud

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type DropDownBox struct {
	pos                 *rl.Vector2
	option              []string
	rec                 rl.Rectangle
	active              int32
	isDropDownBoxActive bool
}

func newDropDown(
	x float32,
	y float32,
	width float32,
	height float32,
	options []string,
) *DropDownBox {
	rec := rl.NewRectangle(
		x,
		y,
		width,
		height,
	)
	return &DropDownBox{
		pos: &rl.Vector2{
			X: x,
			Y: y,
		},
		option:              options,
		rec:                 rec,
		isDropDownBoxActive: false,
		active:              0,
	}
}

func (dd *DropDownBox) Draw() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), dd.rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		dd.isDropDownBoxActive = !dd.isDropDownBoxActive
	}
	// for i, s := range dd.option {
	// rec := rl.NewRectangle(
	// 	dd.pos.X,
	// 	dd.pos.Y+ten*float32(i),
	// 	dd.rec.Width,
	// 	dd.rec.Height,
	// )
	gui.DropdownBox(dd.rec, dd.option[0], nil, dd.isDropDownBoxActive)
	// gui.Label(dd.rec, "TEST LABEL")
	// gui.DropdownBox(dd.rec, dd.option[1], &dd.active, dd.isDropDownBoxActive)
	// }
}
