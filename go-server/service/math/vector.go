package math

import "math"

type Vector2 struct {
	X int
	Y int
}

func NewVector2(x, y int) Vector2 {
	return Vector2{X: x, Y: y}
}

func SubVec2(first, second Vector2) Vector2 {
	return NewVector2(first.X-second.X, first.Y-second.Y)
}

func (v *Vector2) NormalizeVec() {
	factor := math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2))
	v.X = int(float64(v.X) / factor)
	v.Y = int(float64(v.Y) / factor)
}
