package maze2

import (
	"bufio"
	"container/heap"
	"fmt"
	"strings"
)

type Maze struct {
	tiles      map[Position]Tile
	startTile  *NodeTile
	finishTile *NodeTile
	dimensions Dimensions
}

// findPath will set the cost of the Finish Tile to the cost of the shortset path.
// It will also point back from the Finish Tile to the preceeding tile.
func (m *Maze) FindPath() {
	srcNode := m.startTile
	pq := &NodeHeap{}
	heap.Init(pq)
	for {
		// Process Phase
		for _, heading := range Headings() {
			neighbor := m.NeighborTile(*srcNode, heading)
			if neighbor.TileType() == WallType || srcNode.heading.IsOpposite(heading) {
				continue
			}
			if neighbor.TileType() == NodeType {
				tgtNode := neighbor.(*NodeTile)
				cost := srcNode.cost + srcNode.heading.Cost(heading)
				if cost < tgtNode.cost {
					tgtNode.cost = cost
					tgtNode.preTile = srcNode
					tgtNode.heading = heading
					heap.Push(pq, tgtNode)
				}
			}
		}
		srcNode.isExplored = true
		// Identification Phase
		if pq.Len() == 0 {
			break
		}
		srcNode = heap.Pop(pq).(*NodeTile)
	}
}

// func (m Maze) Neighbors(node NodeTile) (*NodeTile, Heading) {
// 	// for _, heading :=
// }

func (m Maze) NeighborTile(tile NodeTile, h Heading) Tile {
	switch h {
	case North:
		return m.tiles[NewPosition(tile.pos.row-1, tile.pos.col)]
	case East:
		return m.tiles[NewPosition(tile.pos.row, tile.pos.col+1)]
	case South:
		return m.tiles[NewPosition(tile.pos.row+1, tile.pos.col)]
	case West:
		return m.tiles[NewPosition(tile.pos.row, tile.pos.col-1)]
	default:
		panic("Switch exhausted")
	}
}

func (m Maze) Tile(row int, col int) Tile {
	pos := NewPosition(row, col)
	tile, exists := m.tiles[pos]
	if exists {
		return tile
	}
	panic(fmt.Sprintf("Tile at %s does not exist", pos))
}

func (m Maze) String() string {
	var s strings.Builder
	for row := 0; row < m.dimensions.rows; row++ {
		s.WriteString(fmt.Sprintf("%4d. ", row))
		for col := 0; col < m.dimensions.cols; col++ {
			tile := m.Tile(row, col)
			s.WriteString(fmt.Sprint(tile))
		}
		s.WriteString(fmt.Sprintln())
	}
	return s.String()
}

func MazeFromStr(s string) Maze {
	newMaze := Maze{}
	tiles := make(map[Position]Tile)
	scanner := bufio.NewScanner(strings.NewReader(s))
	var cols int
	var rowno int
	for rowno = 0; scanner.Scan(); rowno++ {
		line := scanner.Text()
		if cols == 0 {
			cols = len(line)
		}
		for colno, tileChar := range strings.Split(line, "") {
			switch TileType(tileChar) {
			case WallType:
				tile := WallTile{}
				tiles[NewPosition(rowno, colno)] = tile
			case StartType:
				newNode := NewNodeTile(rowno, colno)
				newNode.nodeType = Start
				newNode.heading = East
				newNode.cost = 0
				tiles[NewPosition(rowno, colno)] = newNode
				newMaze.startTile = newNode
			case FinishType:
				newNode := NewNodeTile(rowno, colno)
				newNode.nodeType = Finish
				tiles[NewPosition(rowno, colno)] = newNode
				newMaze.finishTile = newNode
			case NodeType:
				tiles[NewPosition(rowno, colno)] = NewNodeTile(rowno, colno)
			}
		}
	}
	newMaze.tiles = tiles
	newMaze.dimensions = Dimensions{
		rows: rowno,
		cols: cols}
	return newMaze
}

type Dimensions struct {
	rows int
	cols int
}
