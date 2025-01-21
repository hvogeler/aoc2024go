package warehouse

type Pointer rune

const (
	left  Pointer = '<'
	right Pointer = '>'
	up    Pointer = '^'
	down  Pointer = 'v'
)

func (ptr Pointer) String() string {
	return string(ptr)
}

type Path struct {
	pointers []Pointer
	cursor   int
}

func (p *Path) NextPointer() Pointer {
	p.cursor++
	return p.pointers[p.cursor-1]
}

func (p Path) Length() int {
	return len(p.pointers)
}
