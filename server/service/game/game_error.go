package game

import (
	"errors"
)

var (
	ErrUnitHasNoObjective   = errors.New("the unit does not have any objective")
	ErrCannotMoveUnitFurder = errors.New("path for the unit movement is block")
	ErrNotUnitFound         = errors.New("no unit found")
)
