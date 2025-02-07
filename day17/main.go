package main

import (
	cp "day17/cpu"
	"fmt"
)

func main() {
	data := cp.ReadData("testdata.dat")
	cpu := cp.InitialProgramLoad(data)
	fmt.Println(cpu.DisAssemble(-1))
	// cpu.Debug()
	cpu.Run()
	// Part1: 7,5,4,3,4,5,3,4,6
	
}
