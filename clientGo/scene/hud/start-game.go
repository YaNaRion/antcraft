package hud

import rl "github.com/gen2brain/raylib-go/raylib"

type StartGame struct {
	Button *RectangleButton
}

func (rg *StartGame) Draw() {
	rg.Button.Draw()
}

func newStartGameHUD() *StartGame {
	return &StartGame{
		Button: newButton(
			300,
			ten,
			gameStatusActionButtonWidth,
			buttonHeight,
			&rl.Pink,
			"START GAME",
		),
	}
}
