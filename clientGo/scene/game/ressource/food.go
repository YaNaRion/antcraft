package ressource

import rl "github.com/gen2brain/raylib-go/raylib"

const defaultConsumeValue = 5

var quantityToDimension = map[int]int{
	0:  5,  // Quantity < 10
	10: 10, // Quantity < 20
	20: 15, // Quantity < 30
	30: 20, // Quantity < 40
}

func findDimension(quantity int) int {
	return quantityToDimension[(quantity/10)*10]
}

type RessourceMineral interface {
	Consume() int
	Draw()
	GetRec() rl.Rectangle
	GetQuantity() int
}

type DefaultFood struct {
	Quantity  int
	Rec       rl.Rectangle
	dimension float32
	color     rl.Color
}

func NewDefaultFood(quantity int, rec rl.Rectangle, color rl.Color) *DefaultFood {
	dimension := float32(findDimension(quantity))
	rec.Height = dimension
	rec.Width = rec.Height
	return &DefaultFood{
		Quantity:  quantity,
		Rec:       rec,
		color:     color,
		dimension: dimension,
	}
}

func (f *DefaultFood) GetQuantity() int {
	return f.Quantity

}

func (f *DefaultFood) updateStatus() {
	f.dimension = float32(findDimension(f.Quantity))
	f.Rec.Height = f.dimension
	f.Rec.Width = f.dimension
}

func (f *DefaultFood) Consume() int {
	f.Quantity -= defaultConsumeValue
	f.updateStatus()
	return defaultConsumeValue
}

func (f *DefaultFood) Draw() {
	rl.DrawRectangleRec(f.Rec, f.color)
}

func (f *DefaultFood) GetRec() rl.Rectangle {
	return f.Rec
}
