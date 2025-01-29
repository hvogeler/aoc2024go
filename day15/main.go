package main

import (
	wh "day15/warehouse"
	wh2 "day15/wahreouse2"
	"fmt"
)

func main() {
	data := wh.ReadData("testdata.dat")
	// warehouse := wh.WarehouseFromStr(data)
	// warehouse.GoRobotGo()
	// fmt.Println(warehouse)
	// fmt.Printf("Sum Part1: %d\n", warehouse.SumBoxCoords())
	// Result PArt1: 1463715

	warehouse2 := wh2.WarehouseFromStr(data)
	// fmt.Println(warehouse2)
	warehouse2.GoRobotGo()
	// fmt.Println(warehouse2)
	fmt.Printf("***** Sum Part2: %d\n", warehouse2.SumBoxCoords())

}
