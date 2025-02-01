package maze2

import "fmt"

type Position struct {
	row int
	col int
}

func NewPosition(row, col int) Position {
	return Position{row, col}
}

func (pos Position) String() string {
	s := fmt.Sprintf("(%d, %d)", pos.row, pos.col)
	return s
}

func (pos Position) IsInside(dimensions Dimensions) bool {
	if pos.row < 0 || pos.row >= dimensions.rows || pos.col < 0 || pos.col >= dimensions.cols {
		return false
	}
	return true
}

func (pos Position) Coords() (int, int) {
	return pos.row, pos.col
}