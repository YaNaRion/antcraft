package hud

import (
	"client/window"
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Dice struct {
	Button  *RectangleButton
	Resulat string
}

func (ac *Dice) Draw() {
	ac.Button.Draw()
	ac.drawResult()
}

func (ac *Dice) drawResult() {
	rl.DrawText(
		ac.Resulat,
		ac.Button.Rec.ToInt32().X+10,
		ac.Button.Rec.ToInt32().Y-20,
		12,
		rl.White,
	)
}

func (ac *Dice) RolDice() int {
	rol := rand.Intn(12) + 1
	ac.Resulat = fmt.Sprintf("%d", rol)
	return rol
}

func newDiceHUD() *Dice {
	var buttons *RectangleButton = newButton(
		float32(window.SCREEN_WIDTH-250),
		float32(window.SCREEN_HEIGHT/2),
		float32(buttonWidth+90),
		float32(buttonHeight),
		nil,
		"LANCER LES DEES",
	)
	return &Dice{
		Button:  buttons,
		Resulat: "WESH LES MECS",
	}
}
