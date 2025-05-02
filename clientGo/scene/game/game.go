package game

import (
	"client/player"
	"client/scene/hud"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	Players      []*player.Player
	ActivePlayer *player.Player
	HUD          *hud.HUD
	Map          *Map
}

func NewGameScene() *GameScene {
	gameState := "DEBUT DE LA GAME"
	game := &GameScene{
		Players: player.NewPlayers(2),
		HUD:     hud.NewHUD(2, &gameState),
		Map:     NewMap(),
	}
	game.ActivePlayer = game.Players[0]
	return game
}

func (g *GameScene) HandlerInput() {
	// Switch to player0
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.SwitchPlayer.Buttons[0].Rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.ActivePlayer = g.Players[0]
	}

	// Switch to player1
	// if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.SwitchPlayer.Buttons[1].Rec) &&
	// 	rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
	// 	g.ActivePlayer = g.Players[1]
	// }

	// ResetGame
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.ResetGame.Button.Rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.ResetGame()
	}
}

func (g *GameScene) Draw() {
	rl.ClearBackground(rl.Black)
	g.Map.Draw()
	rl.EndMode2D()
	g.DrawHUD()
}

func (g *GameScene) DrawHUD() {
	g.HUD.Draw(g.Players)
}

func (g *GameScene) ResetGame() {
	g.ActivePlayer = g.Players[0]
}
