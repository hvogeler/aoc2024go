package maze2

import "fmt"

type Position struct {
	row int
	col int
	heading Heading
}

func NewPosition(row, col int, heading Heading) Position {
	return Position{row, col, heading}
}

func (pos Position) String() string {
	s := fmt.Sprintf("(%d, %d, %s)", pos.row, pos.col, pos.heading)
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