package main

import (
	ms "day18/memspace"
	"fmt"
)

func main() {
	data := ms.ReadData("testdata.dat")
	fmt.Println(data)
	memSpace := ms.MemSpaceFromStr(data, 71, 71, 1024)
	fmt.Println(memSpace)
	memSpace.BfsWalk()
	fmt.Printf("Steps for shortest path: %d\n", memSpace.ExitNode().PathLen())
	// Part 438 steps

	inputArray := ms.NewInputArray(data)
	for i := 1024; i < len(inputArray); i++ {
		memSpace.CorruptMemAt(inputArray[i].X(), inputArray[i].Y())
		memSpace.ResetBfsWalk()
		fmt.Println(memSpace)
		memSpace.BfsWalk()
		fmt.Printf("Corrupted bytes: %d, Steps for shortest path: %d\n\n", i+1, memSpace.ExitNode().PathLen())
		if memSpace.ExitNode().PathLen() < 0 {
			fmt.Printf("Exit cannot be reached after %d. corrupted byte at %s\n", i, inputArray[i])

			break
		}
	}

}

