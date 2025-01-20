package robots

import (
	"bufio"
	// "fmt"
	"strings"
)

type robotCount int

type Space struct {
	robots     []Robot
	tiles map[Location](robotCount)
	dimensions Dimensions
}

func (space Space) String() string {
	return ""
}

func SpaceFromString(s string, dim Dimensions) Space {
	space := new(Space)
	space.dimensions = dim
	r := bufio.NewScanner(strings.NewReader(s))
	for r.Scan() {
		line := r.Text()
		parts := strings.Split(line, " ")
		location := LocationFromString(parts[0])
		velocity := VelocityFromString(parts[1])
		robot := NewRobot(location, Velocity(velocity))
		space.robots = append(space.robots, robot)
	}
	return *space
}

func (space *Space) makeTiles() {
	space.tiles = make(map[Location](robotCount))
	for _, robot := range space.robots {
		
	}
}

type Dimensions struct {
	tilesX int
	tilesY int
}

