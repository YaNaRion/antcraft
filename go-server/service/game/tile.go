package game

type TileType int

const (
	WATTER_TILE = iota
	WALL_TILE
	GROUND
)

type Tile struct {
	Element  IElement
	TileType TileType
}

func NewTile() Tile {
	return Tile{
		Element:  nil,
		TileType: GROUND,
	}
}
