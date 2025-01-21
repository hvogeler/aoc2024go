package main

import (
	"day14/robots"
	"fmt"
)

func main() {
	data := robots.ReadData("testdata.dat")
	space := robots.SpaceFromString(data, robots.Dimensions{TilesX: 101, TilesY: 103})
	fmt.Println(space)
	space.MoveRobots(100)
	safetyFactor := space.SafetyFactor()
	fmt.Printf("Safety Factor for part 1 is %d\n", safetyFactor)
	// 222901875

	// Part2. Xmas tree appears after 6243 seconds
	data = robots.ReadData("testdata.dat")
	space = robots.SpaceFromString(data, robots.Dimensions{TilesX: 101, TilesY: 103})
	// fmt.Println(space)

    Loop:
	for i := 1; i < 100000; i++ {
		space.MoveRobots(1)
        tile, found := space.FindXmasTree()
		if found {
			fmt.Println(space)
			fmt.Printf("Tree candidate at %s at second %d\n", *tile, i)
            break Loop
		}
	}

}
