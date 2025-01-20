package main

import (
	"fmt"
	"day14/robots"
)

func main() {
	data := robots.ReadData("testdata.dat")
	space := robots.SpaceFromString(data, robots.Dimensions{TilesX: 101, TilesY: 103})
	fmt.Println(space)
	space.MoveRobots(100)
	safetyFactor := space.SafetyFactor()
	fmt.Printf("Safety Factor for part 1 is %d\n", safetyFactor)
	// 222901875
}

