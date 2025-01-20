package robots

import (
	"fmt"
	"strconv"
	"strings"
)

type Velocity struct {
	x int
	y int
}

func (v Velocity) String() string {
	return fmt.Sprintf("(%d, %d)", v.x, v.y)
}

// s contains the position in this form: "p=0,4"
func VelocityFromString(s string) Location {
	parts := strings.Split(s, "=")
	coords := strings.Split(parts[1], ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		panic("Robot velocity x not a number")
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic("Robot velocity y not a number")
	}
	return Location{x, y}
}
