package maze2

import "math"

type TileType string

const (
	// Unused TileType = "."
	Start  TileType = "S"
	Finish TileType = "E"
	Node   TileType = "."
	Wall   TileType = "#"
)

type Heading string

const (
	North     Heading = "^"
	East      Heading = ">"
	South     Heading = "v"
	West      Heading = "<"
	Undefined Heading = "."
)

type Tile interface {
	TileType() TileType
	String() string
}

type NodeTile struct {
	cost       int
	isExplored bool
	pos        Position
	heading    Heading
	preTile    *Tile
}

func (n NodeTile) TileType() TileType {
	return Node
}

func (n NodeTile) String() string {
	return string(n.heading)
}

func (n NodeTile) Heading() Heading {
	return n.heading
}

func NewNodeTile(row int, col int) NodeTile {
	return NodeTile{
		cost:       math.MaxInt,
		isExplored: false,
		heading:    Undefined,
		pos:        NewPosition(row, col),
		preTile:    nil,
	}
}

type WallTile struct {
}

func (w WallTile) TileType() TileType {
	return Wall
}

func (w WallTile) String() string {
	return string(w.TileType())
}

type StartTile struct {
	heading Heading
}

func (s StartTile) Heading() Heading {
	return s.heading
}

func (s StartTile) String() string {
	return string(s.TileType())
}

func (s StartTile) TileType() TileType {
	return Start
}

type FinishTile struct {
	cost int
}

func NewFinishTile() FinishTile {
	return FinishTile{
		cost: math.MaxInt,
	}
}

func (f FinishTile) TileType() TileType {
	return Finish
}

func (f FinishTile) String() string {
	return string(f.TileType())
}
