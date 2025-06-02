package building

import rl "github.com/gen2brain/raylib-go/raylib"

type Building interface {
	Draw()
	GetRec() rl.Rectangle
}
