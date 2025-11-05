package hud

const ten float32 = 30

const (
	buttonWidth  float32 = 60
	buttonHeight float32 = 30
)

type HUD struct {
	ResetGame *ResetGame
	StartSim  *StartGame
}

func NewHUD() *HUD {
	return &HUD{
		ResetGame: newResetGameHUD(),
		StartSim:  newStartGameHUD(),
	}
}

func (h *HUD) Draw() {
	h.ResetGame.Draw()
	h.StartSim.Draw()
}
