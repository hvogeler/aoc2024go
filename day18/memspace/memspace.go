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
	memLocations map[Location]*MemLocation
	dimensions   Dimensions
}

type MemLocation struct {
	memType MemType
	pos     Location
	pathLen int
	prev    *MemLocation
}

func NewMemLocation(loc Location, memType MemType) *MemLocation {
	return &MemLocation{
		memType: memType,
		pos:     loc,
		pathLen: -1,
	}
}

func (ms *MemSpace) CorruptMemAt(x, y int) {
	pos := NewLocation(x, y)
	ms.memLocations[pos] = NewMemLocation(pos, Corrupt)
}

func (ms *MemLocation) IsVisited() bool {
	return ms.pathLen >= 0
}

func (ml MemLocation) PathLen() int {
	return ml.pathLen
}

func NewMemspace(dimX int, dimY int) MemSpace {
	return MemSpace{
		memLocations: make(map[Location]*MemLocation),
		dimensions:   Dimensions{dimX: dimX, dimY: dimY},
	}
}

func (ms *MemSpace) GetAtPos(x, y int) *MemLocation {
	return ms.GetAt(NewLocation(x, y))
}

func (ms *MemSpace) GetAt(loc Location) *MemLocation {
	memLoc, exists := ms.memLocations[loc]
	if exists {
		return memLoc
	}
	newMemLoc := NewMemLocation(loc, Unused)
	ms.memLocations[loc] = newMemLoc
	return ms.memLocations[loc]
}

func (ms *MemSpace) StartNode() *MemLocation {
	return ms.memLocations[NewLocation(0, 0)]
}

func (ms *MemSpace) ExitNode() *MemLocation {
	return ms.memLocations[NewLocation(ms.dimensions.dimX-1, ms.dimensions.dimY-1)]
}

func (ms MemSpace) String() string {
	var rows strings.Builder
	for y := 0; y < ms.dimensions.dimY; y++ {
		var row strings.Builder
		for x := 0; x < ms.dimensions.dimX; x++ {
			loc := NewLocation(x, y)
			memLoc, exists := ms.memLocations[loc]
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
	spc.memLocations[NewLocation(0, 0)] = NewMemLocation(NewLocation(0, 0), Start)
	spc.memLocations[NewLocation(0, 0)].pathLen = 0

	spc.memLocations[NewLocation(dimX-1, dimY-1)] = NewMemLocation(NewLocation(dimX-1, dimY-1), Exit)

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
		spc.memLocations[loc] = NewMemLocation(loc, Corrupt)
	}
	return spc
}

type MemType string

const (
	Start   MemType = "S"
	Exit    MemType = "E"
	Corrupt MemType = "#"
	Unused  MemType = "."
	Visited MemType = "O"
)

type Dimensions struct {
	dimX int
	dimY int
}
