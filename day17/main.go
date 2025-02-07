package main

import (
	cp "day17/cpu"
	"fmt"
)

func main() {
	data := cp.ReadData("testdata.dat")
	cpu := cp.InitialProgramLoad(data)
	fmt.Println(cpu.DisAssemble(-1))
	cpu.Run()
}
