package scene

type Scene int

var StateName = map[Scene]string{
	MENU: "Menu",
	GAME: "Game",
}

const (
	MENU Scene = iota
	GAME
)
