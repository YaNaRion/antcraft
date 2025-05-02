package hud

import (
	"client/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SwitchPlayer struct {
	Buttons []*RectangleButton
}

func newSwitchPlayerHUD(playerAmount int) *SwitchPlayer {
	return &SwitchPlayer{
		Buttons: newSwitchPlayerButton(playerAmount),
	}
}

func (sh *SwitchPlayer) Draw(players []*player.Player) {
	for _, button := range sh.Buttons {
		button.Draw()
	}
}

func newSwitchPlayerButton(playerAmount int) []*RectangleButton {
	var col []*rl.Color
	col = append(col, &rl.Green)
	col = append(col, &rl.Red)

	var team []string
	team = append(team, "VERT")
	team = append(team, "ROUGE")

	var buttons []*RectangleButton
	for i := range playerAmount {
		buttons = append(
			buttons,
			newButton(
				buttonWidth+buttonWidth*2*float32(i),
				ten,
				float32(buttonWidth),
				float32(buttonHeight),
				col[i],
				team[i],
			),
		)
	}
	return buttons
}
