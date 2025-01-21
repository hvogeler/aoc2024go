package warehouse

type Pointer rune

const (
	left  Pointer = '<'
	right         = '>'
	up            = '^'
	down          = 'v'
)

type Path struct {
	pointers []Pointer
	cursor int
}
