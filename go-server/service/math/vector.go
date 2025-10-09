package math

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(x, y float64) Vector2 {
	return Vector2{X: x, Y: y}
}

func SubVec2(first, second Vector2) Vector2 {
	return NewVector2(first.X-second.X, first.Y-second.Y)
}

func (v *Vector2) NormalizeVec() {
	factor := math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
	v.X = v.X / factor
	v.Y = v.Y / factor
}
