package hud

import (
	"client/window"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameStatus struct {
	message *string
}

func (ac *GameStatus) Draw() {
	rl.DrawText(*ac.message, (window.SCREEN_WIDTH/2)-100, 30, 20, rl.White)

}

func newGameStatusHUD(message *string) *GameStatus {
	return &GameStatus{
		message: message,
	}
}
