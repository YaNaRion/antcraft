package game

import (
	"client/scene/hud"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	HUD *hud.HUD
	Map *Map
}

func NewGameScene() *GameScene {
	game := &GameScene{
		HUD: hud.NewHUD(),
		Map: NewMap(),
	}
	game.Map.PopulateDefaultMap()
	return game
}

func (g *GameScene) HandlerInput() {
	// ResetSim
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
	g.HUD.Draw()
}

func (g *GameScene) ResetGame() {
	log.Println("La simulation est reset")
	g.Map.PopulateDefaultMap()
}

func (g *GameScene) GenerateNextFrame() {
	g.Map.DefaultUnitMove()
}
