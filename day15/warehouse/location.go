package warehouse

import (
	"fmt"
	"strconv"
	"strings"
)

type Location struct {
	x int
	y int
}

func NewLocation(x, y int) Location {
	return Location{x, y}
}

func (loc Location) String() string {
	return fmt.Sprintf("(%d, %d)", loc.x, loc.y)
}

func (loc Location) X() int {
	return loc.x
}

func (loc Location) Y() int {
	return loc.y
}

func (loc *Location) SetX(x int) {
	loc.x = x
}

func (loc *Location) SetY(y int) {
	loc.y = y
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

// s contains the position in this form: "p=0,4"
func LocationFromString(s string) Location {
	parts := strings.Split(s, "=")
	coords := strings.Split(parts[1], ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		panic("Robot coordinate x not a number")
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic("Robot coordinate y not a number")
	}
	return Location{x, y}
}
