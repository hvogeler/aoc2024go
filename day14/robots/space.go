package robots

import (
	"bufio"
	"fmt"
	"strings"
)

type Space struct {
	robots     []*Robot
	tiles      map[Location](*Tile)
	dimensions Dimensions
}

func (space *Space) MoveRobots(seconds int) {
	for i := range space.robots {
		space.MoveRobot(i, seconds)
	}
}

func (space *Space) MoveRobot(robotIdx int, seconds int) {
	if robotIdx > len(space.robots) {
		return
	}
	robot := space.robots[robotIdx]
	pos := robot.tile.location
	newX := MathMod(pos.x+robot.velocity.x*seconds, space.dimensions.TilesX)
	newY := MathMod(pos.y+robot.velocity.y*seconds, space.dimensions.TilesY)
	space.Tile(pos.x, pos.y).RemoveRobot(robot)
	if space.Tile(pos.x, pos.y).CountRobots() == 0 {
		delete(space.tiles, Location{pos.x, pos.y})
	}
	space.PlaceRobotOnTile(robot, Location{newX, newY})
}

// in Go modulo keeps the sign :-(
func MathMod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

func (space Space) String() string {
	s := ""
	for y := 0; y < space.dimensions.TilesY; y++ {
		for x := 0; x < space.dimensions.TilesX; x++ {
			if tile, exists := space.tiles[Location{x, y}]; exists {
				s += fmt.Sprintf("%d", tile.CountRobots())
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (space Space) Tile(x, y int) *Tile {
	if tile, exists := space.tiles[Location{x, y}]; exists {
		return tile
	} else {
		return new(Tile)
	}
}

func (space Space) QuadrantCoords(quadrant Quadrant) (Location, Dimensions) {
	quadrantDimensions := Dimensions{(space.dimensions.TilesX - 1) / 2, (space.dimensions.TilesY - 1) / 2}
	switch quadrant {
	case topLeft:
		return Location{0, 0}, quadrantDimensions
	case topRight:
		return Location{quadrantDimensions.TilesX + 1, 0}, quadrantDimensions
	case bottomLeft:
		return Location{0, quadrantDimensions.TilesY + 1}, quadrantDimensions
	case bottomRight:
		return Location{quadrantDimensions.TilesX + 1, quadrantDimensions.TilesY + 1}, quadrantDimensions
	}
	panic("Switch exhausted")
}

func (space Space) CountQuadrant(quadrant Quadrant) int {
	loc, dim := space.QuadrantCoords(quadrant)
	sumRobots := 0
	for y := loc.y; y < loc.y+dim.TilesY; y++ {
		for x := loc.x; x < loc.x+dim.TilesX; x++ {
			sumRobots += space.Tile(x, y).CountRobots()
		}
	}
	return sumRobots
}

func (space Space) SafetyFactor() int {
	sum := 1
	for i := topLeft; i <= bottomRight; i++ {
		sum *= space.CountQuadrant(Quadrant(i))
	}
	return sum
}

func (space *Space) PlaceRobotOnTile(robot *Robot, location Location) {
	if tile, exists := space.tiles[location]; exists {
		robot.tile = tile
		// tile.robots = append(tile.robots, robot)
		tile.AddRobot(robot)
	} else {
		tile := new(Tile)
		tile.location = location
		// tile.robots = append(tile.robots, robot)
		tile.AddRobot(robot)
		robot.tile = tile
		space.tiles[location] = tile
	}
}

func SpaceFromString(s string, dim Dimensions) Space {
	space := new(Space)
	space.dimensions = dim
	space.tiles = make(map[Location]*Tile)
	r := bufio.NewScanner(strings.NewReader(s))
	for id := 0; r.Scan(); id++ {
		line := r.Text()
		parts := strings.Split(line, " ")
		location := LocationFromString(parts[0])
		velocity := VelocityFromString(parts[1])
		robot := NewRobot(id, Velocity(velocity))
		space.robots = append(space.robots, robot)
		space.PlaceRobotOnTile(robot, location)
		// if tile, exists := space.tiles[location]; exists {
		// 	robot.tile = tile
		// 	// tile.robots = append(tile.robots, robot)
		// 	tile.AddRobot(robot)
		// } else {
		// 	tile := new(Tile)
		// 	tile.location = location
		// 	// tile.robots = append(tile.robots, robot)
		// 	tile.AddRobot(robot)
		// 	robot.tile = tile
		// 	space.tiles[location] = tile
		// }
	}
	return *space
}

type Dimensions struct {
	TilesX int
	TilesY int
}

type Quadrant int

const (
	topLeft = iota
	topRight
	bottomLeft
	bottomRight
)
