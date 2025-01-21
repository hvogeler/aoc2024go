package main

import (
	wh "day15/warehouse"
	"fmt"
)

func main() {
	data := wh.ReadData("testdata.dat")
	warehouse := wh.WarehouseFromStr(data)
	warehouse.GoRobotGo()
	fmt.Printf("Sum Part1: %d\n", warehouse.SumBoxCoords())
	// Result PArt1: 1463715
}
