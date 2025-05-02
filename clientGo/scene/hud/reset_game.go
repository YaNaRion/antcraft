package hud

import rl "github.com/gen2brain/raylib-go/raylib"

const resetGameButtonWidth float32 = 90

type ResetGame struct {
	Button *RectangleButton
}

func (rg *ResetGame) Draw() {
	rg.Button.Draw()
}

func newResetGameHUD() *ResetGame {
	return &ResetGame{
		Button: newButton(
			300,
			ten,
			resetGameButtonWidth,
			buttonHeight,
			&rl.Pink,
			"RESET GAME",
		),
	}
}
