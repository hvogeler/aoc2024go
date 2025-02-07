package cpu

import "strings"

type OpCode uint8

const (
	adv OpCode = 0
	bxl OpCode = 1
	bst OpCode = 2
	jnz OpCode = 3
	bxc OpCode = 4
	out OpCode = 5
	bdv OpCode = 6
	cdv OpCode = 7
)

func (op OpCode) String() string {
	var s strings.Builder
	switch op {
	case adv:
		s.WriteString("adv")
	case bxl:
		s.WriteString("bxl")
	case bst:
		s.WriteString("bst")
	case jnz:
		s.WriteString("jnz")
	case bxc:
		s.WriteString("bxc")
	case out:
		s.WriteString("out")
	case bdv:
		s.WriteString("bdv")
	case cdv:
		s.WriteString("cdv")
	default:
		panic("Switch exhausted")
	}
	return s.String()
}