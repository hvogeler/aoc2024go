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

func (space Space) String() string {
	s := ""
	for y := 0; y < space.dimensions.tilesY; y++ {
		for x := 0; x < space.dimensions.tilesX; x++ {
			if tile, exists := space.tiles[Location{x, y}]; exists {
				s += fmt.Sprintf("%d", tile.countRobots())
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
	quadrantDimensions := Dimensions{(space.dimensions.tilesX - 1) / 2, (space.dimensions.tilesY - 1) / 2}
	switch quadrant {
	case topLeft:
		return Location{0, 0}, quadrantDimensions
	case topRight:
		return Location{quadrantDimensions.tilesX + 1, 0}, quadrantDimensions
	case bottomLeft:
		return Location{0, quadrantDimensions.tilesY + 1}, quadrantDimensions
	case bottomRight:
		return Location{quadrantDimensions.tilesX + 1, quadrantDimensions.tilesY + 1}, quadrantDimensions
	}
	panic("Switch exhausted")
}

func (space Space) CountQuadrant(quadrant Quadrant) int {
	loc, dim := space.QuadrantCoords(quadrant)
	sumRobots := 0
	for y := loc.y; y < loc.y+dim.tilesY; y++ {
		for x := loc.x; x < loc.x+dim.tilesX; x++ {
			sumRobots += space.Tile(x, y).countRobots()
		}
	}
	return sumRobots
}

func (space Space) CountAllQuadrants() int {
	sum := 0
	for i := topLeft; i <= bottomRight; i++ {
		sum += space.CountQuadrant(Quadrant(i))
	}
	return sum
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

		if tile, exists := space.tiles[location]; exists {
			robot.tile = tile
			tile.robots = append(tile.robots, robot)
		} else {
			tile := new(Tile)
			tile.location = location
			tile.robots = append(tile.robots, robot)
			robot.tile = tile
			space.tiles[location] = tile
		}
	}
	return *space
}

type Dimensions struct {
	tilesX int
	tilesY int
}

type Quadrant int

const (
	topLeft = iota
	topRight
	bottomLeft
	bottomRight
)
