package hud

const ten float32 = 30

const (
	buttonWidth  float32 = 60
	buttonHeight float32 = 30
)

type HUD struct {
	SwitchPlayer *SwitchPlayer
	ResetGame    *ResetGame
	Action       *Action
	GameStatus   *GameStatus
}

func NewHUD() *HUD {
	return &HUD{
		ResetGame: newResetGameHUD(),
	}
}

func (h *HUD) Draw() {
	h.ResetGame.Draw()
}
