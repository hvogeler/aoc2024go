package memspace

import (
	"bufio"
	"strconv"
	"strings"
)

type InputArray Location

func (ia InputArray) X() int {
	return ia.x
}

func (ia InputArray) Y() int {
	return ia.y
}

func NewInputArray(s string) []Location {
	locs := []Location{}
	scanner := bufio.NewScanner(strings.NewReader(s))
	for n := 0; scanner.Scan(); n++ {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			panic("Invalid Coordinate Format")
		}
		x, errX := strconv.Atoi(parts[0])
		y, errY := strconv.Atoi(parts[1])
		if errX != nil || errY != nil {
			panic("Invalid Coordinate Numeric")
		}
		locs = append(locs, NewLocation(x, y))
	}
	return locs
}

