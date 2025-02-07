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
	fmt.Println(cpu.Output())
	// Part1: 7,5,4,3,4,5,3,4,6

	cpu = cp.InitialProgramLoad(data)
	// It helps to convert the numbers needed to octal to see the pattern
	// regA 0o4     --> 0
	//      0o45    --> 3,0
	//      0o452   --> 3,3,0
	//      0o4526  --> 0,3,3,0
	//      0o45264 --> 5,0,3,3,0
	// This means that one needs to shift the octal digits to the left and only check the interval of 8 numbers
	// for the least significant octal digit (8 numbers). Then shift left again and repeat searching for the next
	// output digit added.
	// The last interval was 0o4526445133267270 - 0o4526445133267300
	regA := cpu.FindRegAVal(0o4526445133267270, 0o4526445133267300, "2,4,1,1,7,5,1,5,4,3,5,5,0,3,3,0")
	// In decimal:
	// regA := cpu.FindRegAVal(164278899142330, 164278899142339, "2,4,1,1,7,5,1,5,4,3,5,5,0,3,3,0")
	fmt.Printf("RegA must be %d to reproduce the input\n", regA)
	// Part2: RegA must be 164278899142333 to reproduce the input
}

