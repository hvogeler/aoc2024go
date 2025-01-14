package main

import (
	"fmt"
	"os"
	"day12/garden"
)

func main() {
	bytes, err := os.ReadFile("testdata.dat")
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	// fmt.Println(data)

	garden := garden.GardenFromStr(data)
	
	fenceCost := garden.FenceCost()
	fmt.Printf("Garden Fencecost = %d\n", fenceCost)
    // fence cost of Part 1 is 1396298
}
