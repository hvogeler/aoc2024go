package maze2p2

import (
	"fmt"
	"math"
)

type TileType string

const (
	// Unused TileType = "."
	StartType  TileType = "S"
	FinishType TileType = "E"
	NodeType   TileType = "."
	WallType   TileType = "#"
)

type Heading string

const (
	North      Heading = "^"
	East       Heading = ">"
	South      Heading = "v"
	West       Heading = "<"
	Undefined  Heading = "."
	AnyHeading Heading = "o"
)

func Headings() []Heading {
	return []Heading{North, East, South, West}
}

func (srcHeading Heading) Cost(tgtHeading Heading) int {
	if (srcHeading == North || srcHeading == South) &&
		(tgtHeading == East || tgtHeading == West) {
		return 1000
	}
	if (srcHeading == East || srcHeading == West) &&
		(tgtHeading == North || tgtHeading == South) {
		return 1000
	}
	if srcHeading.IsOpposite(tgtHeading) {
		return 2000
	}
	return 0
}

func (srcHeading Heading) IsOpposite(tgtHeading Heading) bool {
	return (srcHeading == North && tgtHeading == South) ||
		(srcHeading == East && tgtHeading == West) ||
		(srcHeading == South && tgtHeading == North) ||
		(srcHeading == West && tgtHeading == East)
}

type NodeSubType int

const (
	Intermediate NodeSubType = iota
	Start
	Finish
)

type Tile interface {
	TileType() TileType
	String() string
}

type NodeTile struct {
	nodeType   NodeSubType
	cost       int
	isExplored bool
	pos        Position
	heading    Heading
	preTile    []*NodeTile
}

func (n NodeTile) TileType() TileType {
	return NodeType
}

func (n NodeTile) Println(withPreTiles bool) {
	fmt.Printf("Tile: %s - Heading: %s - Cost: %d - PreTiles: %d - IsExplored: %v", n.pos, n.heading, n.cost, len(n.preTile), n.isExplored)
	if withPreTiles {
		fmt.Println("")
		for i, preT := range n.preTile {
			fmt.Printf("   %d. ", i)
			preT.Println(false)
			fmt.Println("")
		}
	}
}

func (n *NodeTile) SetPosition(pos Position) {
	n.pos = pos
}

func (n NodeTile) String() string {
	switch n.nodeType {
	case Intermediate:
		return string(n.heading)
	case Start:
		return string(StartType)
	case Finish:
		return string(FinishType)
	default:
		panic("Switch exhausted")
	}
}

func (n NodeTile) Heading() Heading {
	return n.heading
}

func NewNodeTile(row int, col int, h Heading) *NodeTile {
	return &NodeTile{
		cost:       math.MaxInt,
		isExplored: false,
		heading:    h,
		pos:        NewPosition(row, col, h),
		preTile:    []*NodeTile{},
	}
}

type WallTile struct {
}

func (w WallTile) TileType() TileType {
	return WallType
}

func (w WallTile) String() string {
	return string(w.TileType())
}

// type StartTile struct {
// 	heading Heading
// }

// func (s StartTile) Heading() Heading {
// 	return s.heading
// }

// func (s StartTile) String() string {
// 	return string(s.TileType())
// }

// func (s StartTile) TileType() TileType {
// 	return Start
// }

// type FinishTile struct {
// 	cost int
// }

// func NewFinishTile() FinishTile {
// 	return FinishTile{
// 		cost: math.MaxInt,
// 	}
// }

// func (f FinishTile) TileType() TileType {
// 	return Finish
// }

// func (f FinishTile) String() string {
// 	return string(f.TileType())
// }
