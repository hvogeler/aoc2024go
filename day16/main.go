package main

import (
	mz "day16/maze"
	"fmt"
)

func main() {
	data := mz.ReadData("testdata.dat")
	fmt.Println(data)
	// data := ReadData("../testdata.dat")
	// fmt.Println(data)
	maze := mz.MazeFromStr(data)
	fmt.Println(maze)
	i := 0
	for i = 0; maze.CountAlive() > 0; i++ {
		maze.MoveReindeer()
		// if i%100 == 0 {
		// 	fmt.Println(maze)
		// 	fmt.Println()
		// }
	}
	fmt.Println(maze)
	// 149556
}
