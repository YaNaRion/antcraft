package player

import rl "github.com/gen2brain/raylib-go/raylib"

type PlayerAction int

const (
	PlaceRoute PlayerAction = iota
	PlaceVillage
	UpgradeTownCenter
)

type Player struct {
	ColorTeam      *rl.Color
	Action         PlayerAction
	IsActivePlayer bool
}

func NewPlayers(playerCount int) []*Player {
	var colors []*rl.Color
	colors = append(colors, &rl.Green)
	colors = append(colors, &rl.Red)

	var players []*Player
	for i := range playerCount {
		players = append(players, &Player{
			ColorTeam:      colors[i],
			Action:         PlaceRoute,
			IsActivePlayer: false,
		})

	}
	return players
}
