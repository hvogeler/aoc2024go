package main

import (
	"fmt"
	o "day19/onsen"
)

func main() {
	data := o.ReadData("testdata.dat")
	fmt.Println(data)

	onsen := o.OnsenFromStr(data)
	onsen.CheckDesigns()
	all, possible, impossible := onsen.CountDesigns()

	fmt.Printf("Part1: All Designs: %d, Possible %d, Impossible: %d\n", all, possible, impossible)

	for i, design := range onsen.Designs() {
		if !design.IsPatternPossible() {
			fmt.Println(i)
			fmt.Println(design)
			fmt.Println()
		}
	}
}

