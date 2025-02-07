package cpu

import "fmt"

var OpExec = map[OpCode]func(c *Cpu, oper Operand){
	adv: Adv,
	bdv: Bdv,
	cdv: Cdv,
	out: Out,
	jnz: Jnz,
	bst: Bst,
	bxl: Bxl,
	bxc: Bxc,
}

func Adv(c *Cpu, oper Operand) {
	c.regA = c.regA >> oper
	c.SetInstrPtr(c.instrPtr + 2)
}

func Bdv(c *Cpu, oper Operand) {
	c.regB = c.regA >> oper
	c.SetInstrPtr(c.instrPtr + 2)
}

func Cdv(c *Cpu, oper Operand) {
	c.regC = c.regA >> oper
	c.SetInstrPtr(c.instrPtr + 2)
}

func Bxl(c *Cpu, oper Operand) {
	c.regB = c.regB ^ int(oper)
	c.SetInstrPtr(c.instrPtr + 2)
}

func Bst(c *Cpu, oper Operand) {
	c.regB = c.Eval(oper) % 8
	c.SetInstrPtr(c.instrPtr + 2)
}

func Bxc(c *Cpu, oper Operand) {
	c.regB = c.regC ^ c.regB
	c.SetInstrPtr(c.instrPtr + 2)
}
func Out(c *Cpu, oper Operand) {
	output := c.Eval(oper) % 8
	c.SetInstrPtr(c.instrPtr + 2)
	if c.outputCount > 0 {
		fmt.Print(",")
	}
	c.outputCount++
	fmt.Printf("%d", output)
}

func Jnz(c *Cpu, oper Operand) {
	if c.regA == 0 {
		c.SetInstrPtr(c.instrPtr + 2)
		return
	}
	if c.regA > 0 && int(oper) < len(c.codeMem) {
		c.SetInstrPtr(int(oper))
		return
	}
	panic("Invalid Operand to Jnz")
}
