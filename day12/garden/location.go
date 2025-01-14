package garden

import "fmt"

type Location struct {
	row int
	col int
}

func (loc Location) String() string {
	return fmt.Sprintf("(%d, %d)", loc.row, loc.col)
}

func (loca Location) Compare(locb Location) int {
	if loca == locb {
		return 0
	}
	if loca.row < locb.row {
		return -1
	}
	if loca.row > locb.row {
		return 1
	}
	if loca.col < locb.col {
		return -1
	}
	if loca.col > locb.col {
		return 1
	}
	panic("Cannot happen")
}

