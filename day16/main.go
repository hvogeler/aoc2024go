package main

import (
	mz "day16/maze2p2"
	"fmt"
)

func main() {
	data := mz.ReadData("testdata.dat")
	fmt.Println(data)
	m := mz.MazeFromStr(data)
	fmt.Println(m)
	m.FindPath()
	fmt.Println(m.PrintPath(m.ShortestPaths()[0]))
	fmt.Printf("Cost: %d\n", m.Score())
	// Part1: 133584
}
