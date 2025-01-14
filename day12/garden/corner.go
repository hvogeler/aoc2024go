package garden

type Corner struct {
	plot *Plot
	cornerType CornerType
	orientation Orientation
}

type CornerType int

const (
	convex = iota
	concave
)

type Orientation int

const (
	topLeft = iota
	topRight
	bottomLeft
	bottomRight
)
