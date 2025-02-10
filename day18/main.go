package main

import (
	ms "day18/memspace"
	"fmt"
)

func main() {
	data := ms.ReadData("testdata.dat")
	fmt.Println(data)
	ms := ms.MemSpaceFromStr(data, 71, 71, 1024)
	fmt.Println(ms)
	ms.BfsWalk()
	fmt.Printf("Steps for shortest path: %d\n", ms.ExitNode().PathLen())

}

