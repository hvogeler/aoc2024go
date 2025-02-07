package cpu

import (
	"fmt"
	"strconv"
)

type Instruction struct {
	op OpCode
	operand Operand
}

func NewInstruction(op uint8, operand uint8) Instruction {
	return Instruction{OpCode(op), Operand(operand)}
}

func (i Instruction) String() string {
	return fmt.Sprintf("%5s %d\n", i.op, i.operand)
}

func (i Instruction) DisAssemble() string {
	return fmt.Sprintf("%s %s\n", i.op, i.operand)
}

type Operand uint8

func (o Operand) String() string {
	switch o {
	case 0, 1, 2, 3:
		return strconv.Itoa(int(o))
	case 4:
		return "RegA"
	case 5:
		return "RegB"
	case 6:
		return "RegC"
	default:
		panic("Invalid Operand")
	}

}