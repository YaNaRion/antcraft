package camera

import (
	"client/scene/game"
	"client/window"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const edgeThreashold int32 = 5
const cameraSpeed float32 = 50.0

type Camera struct {
	Cam *rl.Camera2D
}

var CameraOffSet = rl.Vector2{
	X: float32(window.SCREEN_WIDTH / 2),
	Y: float32(window.SCREEN_HEIGHT / 2),
}

func NewCamera(offset rl.Vector2, target rl.Vector2) *Camera {
	return &Camera{
		Cam: &rl.Camera2D{
			Zoom:   float32(1.0),
			Target: target,
			// Offset: offset,
		},
	}
}

func (c *Camera) HandlerZoom() {
	c.Cam.Zoom += rl.GetMouseWheelMove() * 0.10
	if c.Cam.Zoom < 0.1 {
		c.Cam.Zoom = 0.1
	}
}

func (c *Camera) MoveEdge() {
	mousePosition := rl.GetMousePosition()
	if mousePosition.X < float32(edgeThreashold) {
		if c.Cam.Target.X > 0 {
			c.Cam.Target.X -= cameraSpeed
		}
	} else if mousePosition.X > float32(window.SCREEN_WIDTH-edgeThreashold) {
		if c.Cam.Target.X <= game.MapWidth {
			c.Cam.Target.X += cameraSpeed
		}
	}

	if mousePosition.Y < float32(edgeThreashold) {
		if c.Cam.Target.Y > 0 {
			c.Cam.Target.Y -= cameraSpeed
		}
	} else if mousePosition.Y > float32(window.SCREEN_HEIGHT-edgeThreashold) {
		if c.Cam.Target.Y <= game.MapHeight {
			c.Cam.Target.Y += cameraSpeed
		}
	}

}
