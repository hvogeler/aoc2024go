package clawmachine

import "fmt"

type Location struct {
	x int
	y int
}

func (loc Location) String() string {
	return fmt.Sprintf("(%d, %d)", loc.x, loc.y)
}

func (loca Location) Compare(locb Location) int {
	if loca == locb {
		return 0
	}
	if loca.x < locb.x {
		return -1
	}
	if loca.x > locb.x {
		return 1
	}
	if loca.y < locb.y {
		return -1
	}
	if loca.y > locb.y {
		return 1
	}
	panic("Cannot happen")
}

func (loc Location) IsPast(b Location) bool {
	return loc.x > b.x || loc.y > b.y 
}