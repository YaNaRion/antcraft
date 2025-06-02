package ressource

import rl "github.com/gen2brain/raylib-go/raylib"

var quantityToDimension = map[int]int{
	0:  5,  // Quantity < 10
	10: 10, // Quantity < 20
	20: 15, // Quantity < 30
	30: 20, // Quantity < 40
}

func findDimension(quantity int) int {
	return quantityToDimension[(quantity/10)*10]
}

type Ressource interface {
	Consume()
	Draw()
	GetRec() rl.Rectangle
}

type Food struct {
	Quantity  int
	Rec       rl.Rectangle
	dimension float32
	color     rl.Color
}

func NewFood(quantity int, rec rl.Rectangle, color rl.Color) *Food {
	rec.Height = float32(findDimension(quantity))
	rec.Width = rec.Height
	return &Food{
		Quantity: quantity,
		Rec:      rec,
		color:    color,
	}
}

func (f *Food) Consume() {
	f.Quantity -= 1
	f.dimension = float32(findDimension(f.Quantity))
}

func (f *Food) Draw() {
	rl.DrawRectangleRec(f.Rec, f.color)
}

func (f *Food) GetRec() rl.Rectangle {
	return f.Rec
}
