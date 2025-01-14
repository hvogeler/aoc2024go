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
	fmt.Printf("Garden Fencecost 1 = %d\n", fenceCost)
    // fence cost of Part 1 is 1396298

    fenceCost2 := garden.FenceCost2()
    fmt.Printf("Garden Fencecost 2 = %d\n", fenceCost2)
    // fence cost of Part 2 is 853588
}
