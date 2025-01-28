package warehouse

type Pointer rune

const (
	Left  Pointer = '<'
	Right Pointer = '>'
	Up    Pointer = '^'
	Down  Pointer = 'v'
)

func (ptr Pointer) Orientation() Orientation {
	if ptr == Left || ptr == Right {
		return Horizontal
	} else {
		return Vertical
	}
}

func (ptr Pointer) String() string {
	return string(ptr)
}

type Orientation int
const (
	Horizontal Orientation = iota
	Vertical
)

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

func (p *Path) AddPointer(ptr Pointer) {
	p.pointers = append(p.pointers, ptr)
}
