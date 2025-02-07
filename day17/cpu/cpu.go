package cpu

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Cpu struct {
	regA, regB, regC   int
	codeMem            []uint8
	instrPtr           int
	currentInstruction Instruction
}

func (c Cpu) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("\nRegisters:\n  A: %d, B: %d, C: %d - Next Instruction at (%d): %s\n",
		c.regA, c.regB, c.regC, c.instrPtr, c.currentInstruction))
    return s.String()
}

func (c Cpu) DisAssemble(addr int) string {
	var s strings.Builder
    for i := 0; i<len(c.codeMem); i+=2 {
        addrIndicator := "  "
        if addr >= 0 && addr % 2 == 0 && addr == i {
            addrIndicator = "->"
        }
        instr := c.InstrAt(i)
        s.WriteString(fmt.Sprintf("%s %s", addrIndicator, instr.DisAssemble()))
    }
    return s.String()
}

func (c Cpu) Eval(o Operand) int {
	switch o {
	case 0, 1, 2, 3:
		return int(o)
	case 4:
		return c.regA
	case 5:
		return c.regB
	case 6:
		return c.regC
	default:
		panic("Invalid Operand")
	}
}

func (c *Cpu) SetRegA(v int) {
	c.regA = v
}

func (c *Cpu) SetRegB(v int) {
	c.regB = v
}

func (c *Cpu) SetRegC(v int) {
	c.regC = v
}

func (c *Cpu) SetInstrPtr(addr int) {
	c.instrPtr = addr
	c.currentInstruction = c.InstrAt(addr)
}

func (c Cpu) InstrAt(addr int) Instruction {
    if addr % 2 != 0 {
        panic("Invalid Program Address")
    }
    return NewInstruction(c.codeMem[addr], c.codeMem[addr+1])
}

func InitialProgramLoad(program string) Cpu {
	cpu := new(Cpu)
	cpu.instrPtr = 0
	scanner := bufio.NewScanner(strings.NewReader(program))

	for lineno := 0; scanner.Scan(); lineno++ {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			panic("Program can not be parsed")
		}
		switch parts[0] {
		case "Register A":
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(fmt.Sprint("Syntax Error ", err))
			}
			cpu.SetRegA(v)
		case "Register B":
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(fmt.Sprint("Syntax Error ", err))
			}
			cpu.SetRegB(v)
		case "Register C":
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(fmt.Sprint("Syntax Error ", err))
			}
			cpu.SetRegC(v)
		case "Program":
			instrs := strings.Split(parts[1], ",")
			cpu.codeMem = make([]uint8, len(instrs))
			for i, instr := range instrs {
				v, err := strconv.Atoi(instr)
				if err != nil {
					panic(fmt.Sprint("Syntax Error ", err))
				}
				cpu.codeMem[i] = uint8(v)
			}
		default:
			panic("Program can not be parsed")
		}
	}
	cpu.SetInstrPtr(0)
	return *cpu
}
