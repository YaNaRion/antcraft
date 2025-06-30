package game

import (
	"client/scene/hud"
	"log"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	HUD         *hud.HUD
	Map         *Map
	isSimStated bool
	GenWork     *time.Timer
}

func NewGameScene() *GameScene {
	game := &GameScene{
		HUD:         hud.NewHUD(),
		Map:         NewMap(),
		isSimStated: false,
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
	// StartSim
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.StartSim.Button.Rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.StartSim()
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
	g.isSimStated = false
	g.Map.RestMap()
	g.Map.PopulateDefaultMap()
}

func (g *GameScene) StartSim() {
	log.Println("La simulation est partie")
	g.isSimStated = true
	go func() {
		for range 6 {
			if g.isSimStated {
				g.Map.GenerateNewWorker()
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func (g *GameScene) GenerateNextFrame() {
	if g.isSimStated {
		g.Map.DefaultUnitMove()
	}
}
