package main

import (
	"day14/robots"
	"fmt"
)

func main() {
	// data := robots.ReadData("testdata.dat")
	// space := robots.SpaceFromString(data, robots.Dimensions{TilesX: 101, TilesY: 103})
	// fmt.Println(space)
	// space.MoveRobots(100)
	// safetyFactor := space.SafetyFactor()
	// fmt.Printf("Safety Factor for part 1 is %d\n", safetyFactor)
	// 222901875

	// Part2
	data := robots.ReadData("testdata.dat")
	space := robots.SpaceFromString(data, robots.Dimensions{TilesX: 101, TilesY: 103})
	// fmt.Println(space)

	for i := 700000; i < 7000000; i++ {
		space.MoveRobots(1)
		if space.Tile(50, 0).CountRobots() > 0 &&
			space.Tile(49, 1).CountRobots() > 0 &&
			space.Tile(51, 1).CountRobots() > 0 &&
			space.Tile(48, 2).CountRobots() > 0 &&
			space.Tile(52, 2).CountRobots() > 0 {
			fmt.Println(space)
			fmt.Println(i)
		}
	}

}
