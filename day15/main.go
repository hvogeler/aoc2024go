package main

import (
	wh "day15/warehouse"
	"fmt"
)

func main() {
	data := wh.ReadData("example.dat")
	fmt.Println(data)
}
