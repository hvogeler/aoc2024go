package memspace

import (
	"bufio"
	"strconv"
	"strings"
)

// Note: Dimensions given in puzzle are wrong. Add 1 to each coordinate they give us.
// When they say (6,6) Grid they mean actually (7, 7)
// because they include the 0 position.
// For me a (6,6) grid would only havemax coords (5,5)
type MemSpace struct {
	MemLocations map[Location]MemLocation
	dimensions   Dimensions
}

type MemLocation struct {
	memType MemType
	pos     Location
}

func NewMemspace(dimX int, dimY int) MemSpace {
	return MemSpace{
		MemLocations: make(map[Location]MemLocation),
		dimensions:   Dimensions{dimX: dimX, dimY: dimY},
	}
}

func (ms MemSpace) String() string {
	var rows strings.Builder
	for y := 0; y < ms.dimensions.dimY; y++ {
		var row strings.Builder
		for x := 0; x < ms.dimensions.dimX; x++ {
			loc := NewLocation(x, y)
			memLoc, exists := ms.MemLocations[loc]
			if exists {
				row.WriteString(string(memLoc.memType))
			} else {
				row.WriteString(string(Unused))
			}
		}
		rows.WriteString(row.String())
		rows.WriteString("\n")
	}
	return rows.String()
}

func MemSpaceFromStr(s string, dimX, dimY int, maxCorrupted int) MemSpace {
	spc := NewMemspace(dimX, dimY)
	spc.MemLocations[NewLocation(0, 0)] = MemLocation{
		memType: Start,
		pos:     NewLocation(0, 0),
	}
	spc.MemLocations[NewLocation(dimX-1, dimY-1)] = MemLocation{
		memType: Exit,
		pos:     NewLocation(dimX-1, dimY-1),
	}

	scanner := bufio.NewScanner(strings.NewReader(s))
	for n := 0; scanner.Scan() && n < maxCorrupted; n++ {
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
		loc := NewLocation(x, y)
		spc.MemLocations[loc] = MemLocation{
			memType: Corrupt,
			pos:     loc,
		}
	}
	return spc
}

type MemType string

const (
	Start   MemType = "S"
	Exit    MemType = "E"
	Corrupt MemType = "#"
	Unused  MemType = "."
)

type Dimensions struct {
	dimX int
	dimY int
}
